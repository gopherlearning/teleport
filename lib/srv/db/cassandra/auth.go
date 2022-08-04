/*
Copyright 2022 Gravitational, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cassandra

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws/arn"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	awssession "github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sigv4-auth-cassandra-gocql-driver-plugin/sigv4"
	"github.com/datastax/go-cassandra-native-protocol/frame"
	"github.com/datastax/go-cassandra-native-protocol/message"
	"github.com/datastax/go-cassandra-native-protocol/primitive"
	"github.com/gocql/gocql"
	"github.com/gravitational/trace"

	"github.com/gravitational/teleport/lib/srv/db/cassandra/protocol"
	"github.com/gravitational/teleport/lib/srv/db/common"
)

const (
	passwordAuthenticator = "org.apache.cassandra.auth.PasswordAuthenticator"
)

type handshakeHandler interface {
	handleHandshake(clientConn, serverConn *protocol.Conn) error
}

type basicHandshake struct {
	ses *common.Session
}

func (pp *basicHandshake) handleHandshake(clientConn, serverConn *protocol.Conn) error {
	for {
		req, err := clientConn.ReadPacket()
		if err != nil {
			return trace.Wrap(err)
		}

		if req.Header().OpCode == primitive.OpCodeAuthResponse {
			if err := handleAuthResponse(clientConn, pp.ses, req); err != nil {
				return trace.Wrap(err)
			}
		}
		if err := serverConn.WriteFrame(req.Frame()); err != nil {
			return trace.Wrap(err)
		}
		rcv, err := serverConn.ReadPacket()
		if err != nil {
			return trace.Wrap(err)
		}
		if _, err := clientConn.Write(rcv.Raw()); err != nil {
			return trace.Wrap(err)
		}
		switch rcv.Header().OpCode {
		case primitive.OpCodeReady, primitive.OpCodeAuthSuccess:
			return nil
		}
	}
}

type failedHandshake struct {
	error error
}

func (h failedHandshake) handshake(clientConn, _ *protocol.Conn) error {
	for {
		packet, err := clientConn.ReadPacket()
		if err != nil {
			return trace.Wrap(err)
		}

		switch packet.Header().OpCode {
		case primitive.OpCodeStartup:
			fr := frame.NewFrame(
				packet.Header().Version,
				packet.Header().StreamId,
				&message.Authenticate{Authenticator: passwordAuthenticator},
			)
			if err := clientConn.WriteFrame(fr); err != nil {
				return trace.Wrap(err)
			}
		case primitive.OpCodeAuthResponse:
			var msg message.Message
			if trace.IsAccessDenied(err) {
				msg = &message.AuthenticationError{ErrorMessage: h.error.Error()}
			} else {
				msg = &message.ServerError{ErrorMessage: h.error.Error()}
			}
			fr := frame.NewFrame(packet.Header().Version, packet.Header().StreamId, msg)
			if err := clientConn.WriteFrame(fr); err != nil {
				return trace.Wrap(err)
			}
			return nil

		case primitive.OpCodeOptions:
			fr := frame.NewFrame(
				packet.Header().Version,
				packet.Header().StreamId,
				&message.Supported{
					Options: map[string][]string{
						"CQL_VERSION": {"3.4.5"},
						"COMPRESSION": {},
					},
				},
			)
			if err := clientConn.WriteFrame(fr); err != nil {
				return trace.Wrap(err)
			}
		}
	}
}

func handleAuthResponse(clientConn *protocol.Conn, ses *common.Session, pkt *protocol.Packet) error {
	msg, ok := pkt.FrameBody().Message.(*message.AuthResponse)
	if !ok {
		return trace.BadParameter("got unexpected packet")
	}
	vErr := validateCassandraUsername(ses, msg)
	if vErr == nil {
		return nil
	}
	if err := sendAuthenticationErrorMessage(vErr, clientConn, pkt.Frame()); err != nil {
		return trace.NewAggregate(vErr, err)
	}
	return vErr
}

func sendAuthenticationErrorMessage(authErr error, clientConn *protocol.Conn, incoming *frame.Frame) error {
	authErrMsg := frame.NewFrame(incoming.Header.Version, incoming.Header.StreamId, &message.AuthenticationError{
		ErrorMessage: authErr.Error(),
	})
	if err := clientConn.WriteFrame(authErrMsg); err != nil {
		return trace.Wrap(err)
	}
	return nil
}

type authAWSSigV4Auth struct {
	ses      *common.Session
	authFunc func(username string, ses *common.Session) (gocql.Authenticator, error)
}

func (a *authAWSSigV4Auth) getSigV4Authenticator(username string) (gocql.Authenticator, error) {
	session, err := awssession.NewSessionWithOptions(awssession.Options{
		SharedConfigState: awssession.SharedConfigEnable,
	})
	if err != nil {
		return nil, trace.Wrap(err)
	}
	cred, err := stscreds.NewCredentials(session, a.buildRoleARN(username)).Get()
	if err != nil {
		return nil, trace.Wrap(err)
	}
	auth := sigv4.NewAwsAuthenticator()
	auth.Region = a.ses.Database.GetAWS().Region
	auth.AccessKeyId = cred.AccessKeyID
	auth.SessionToken = cred.SessionToken
	auth.SecretAccessKey = cred.SecretAccessKey
	return auth, nil
}

func (a *authAWSSigV4Auth) buildRoleARN(username string) string {
	if arn.IsARN(username) {
		return username
	}
	resource := username
	if !strings.Contains(resource, "/") {
		resource = fmt.Sprintf("role/%s", username)
	}

	return arn.ARN{
		Partition: "aws",
		Service:   "iam",
		AccountID: a.ses.Database.GetAWS().AccountID,
		Resource:  resource,
	}.String()
}

func (a *authAWSSigV4Auth) initPasswordAuth(clientConn *protocol.Conn, req *protocol.Packet) (*protocol.Packet, error) {
	authMsg := frame.NewFrame(
		req.Header().Version,
		req.Header().StreamId,
		&message.Authenticate{Authenticator: passwordAuthenticator},
	)
	if err := clientConn.WriteFrame(authMsg); err != nil {
		return nil, trace.Wrap(err)
	}
	pkt, err := clientConn.ReadPacket()
	if err != nil {
		return nil, trace.Wrap(err)
	}
	return pkt, nil
}

func (a *authAWSSigV4Auth) handleStartupMessage(clientConn, serverConn *protocol.Conn, req *protocol.Packet) error {
	authResp, err := a.initPasswordAuth(clientConn, req)
	if err != nil {
		return trace.Wrap(err)
	}
	if err := handleAuthResponse(clientConn, a.ses, authResp); err != nil {
		return trace.Wrap(err)
	}

	if err := serverConn.WriteFrame(req.Frame()); err != nil {
		return trace.Wrap(err)
	}
	authFrame, err := serverConn.ReadPacket()
	if err != nil {
		return trace.Wrap(err)
	}
	if err := a.handleAuth(clientConn, serverConn, authFrame.Frame()); err != nil {
		return trace.Wrap(err)
	}

	readyFrame := frame.NewFrame(authResp.Header().Version, authResp.Header().StreamId, &message.AuthSuccess{})
	if err := clientConn.WriteFrame(readyFrame); err != nil {
		return trace.Wrap(err)
	}
	return nil
}

func (a *authAWSSigV4Auth) handleHandshake(clientConn, serverConn *protocol.Conn) error {
	for {
		req, err := clientConn.ReadPacket()
		if err != nil {
			return trace.Wrap(err)
		}
		switch req.Header().OpCode {
		case primitive.OpCodeStartup:
			if err := a.handleStartupMessage(clientConn, serverConn, req); err != nil {
				return trace.Wrap(err)
			}
			return nil
		default:
			if err := serverConn.WriteFrame(req.Frame()); err != nil {
				return trace.Wrap(err)
			}
			resp, err := serverConn.ReadPacket()
			if err != nil {
				return trace.Wrap(err)
			}
			if _, err := clientConn.Write(resp.Raw()); err != nil {
				return trace.Wrap(err)
			}
		}
	}
}

func (a *authAWSSigV4Auth) handleAuth(_, serverConn *protocol.Conn, fr *frame.Frame) error {
	authMsg, ok := fr.Body.Message.(*message.Authenticate)
	if !ok {
		return trace.BadParameter("unexpected message type %T", fr.Body.Message)
	}
	awsAuth, err := a.getSigV4Authenticator(a.ses.DatabaseUser)
	if err != nil {
		return trace.Wrap(err)
	}

	data, challenger, err := awsAuth.Challenge([]byte(authMsg.Authenticator))
	if err != nil {
		return trace.Wrap(err)
	}

	authRespFrame := frame.NewFrame(
		fr.Header.Version,
		fr.Header.StreamId,
		&message.AuthResponse{Token: data},
	)
	for {
		if err := serverConn.WriteFrame(authRespFrame); err != nil {
			return trace.Wrap(err)
		}
		rcv, err := serverConn.ReadPacket()
		if err != nil {
			return trace.Wrap(err)
		}

		switch v := rcv.FrameBody().Message.(type) {
		case *message.AuthSuccess:
			if challenger != nil {
				if err = challenger.Success(v.Token); err != nil {
					return trace.Wrap(err)
				}
			}
			return nil
		case *message.AuthChallenge:
			data, challenger, err = challenger.Challenge(v.Token)
			if err != nil {
				return err
			}
			authRespFrame = frame.NewFrame(
				rcv.Header().Version,
				rcv.Header().StreamId,
				&message.AuthResponse{
					Token: data,
				})
		default:
			return fmt.Errorf("unknown frame response during authentication: %v", v)
		}
	}
}
