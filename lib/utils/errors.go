/*
 * Teleport
 * Copyright (C) 2023  Gravitational, Inc.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package utils

import (
	"context"
	"errors"
	"io"
	"net"
	"strings"
	"syscall"

	"github.com/gravitational/trace"

	"github.com/gravitational/teleport/api/constants"
)

// IsUseOfClosedNetworkError returns true if the specified error
// indicates the use of a closed network connection.
func IsUseOfClosedNetworkError(err error) bool {
	if err == nil {
		return false
	}

	return errors.Is(err, net.ErrClosed) || strings.Contains(err.Error(), constants.UseOfClosedNetworkConnection)
}

// IsFailedToSendCloseNotifyError returns true if the provided error is the
// "tls: failed to send closeNotify".
func IsFailedToSendCloseNotifyError(err error) bool {
	if err == nil {
		return false
	}
	return strings.Contains(err.Error(), constants.FailedToSendCloseNotify)
}

// IsOKNetworkError returns true if the provided error received from a network
// operation is one of those that usually indicate normal connection close. If
// the error is a trace.Aggregate, all the errors must be OK network errors.
func IsOKNetworkError(err error) bool {
	// trace.Aggregate contains at least one error and all the errors are
	// non-nil
	var a trace.Aggregate
	if errors.As(trace.Unwrap(err), &a) {
		for _, err := range a.Errors() {
			if !IsOKNetworkError(err) {
				return false
			}
		}
		return true
	}
	return errors.Is(err, io.EOF) || IsUseOfClosedNetworkError(err) || IsFailedToSendCloseNotifyError(err)
}

// IsConnectionRefused returns true if the given err is "connection refused" error.
func IsConnectionRefused(err error) bool {
	var errno syscall.Errno
	if errors.As(err, &errno) {
		return errors.Is(errno, syscall.ECONNREFUSED)
	}
	return false
}

// IsUntrustedCertErr checks if an error is an untrusted cert error.
func IsUntrustedCertErr(err error) bool {
	if err == nil {
		return false
	}
	errMsg := err.Error()
	return strings.Contains(errMsg, "x509: certificate is valid for") ||
		strings.Contains(errMsg, "certificate is not trusted")
}

// CanExplainNetworkError returns a simple to understand error message that can
// be used to debug common network and/or protocol errors.
func CanExplainNetworkError(err error) (string, bool) {
	var derr *net.DNSError

	switch {
	// Connection refused errors can be reproduced by attempting to connect to a
	// host:port that no process is listening on. The raw error typically looks
	// like the following:
	//
	// dial tcp 127.0.0.1:8000: connect: connection refused
	case errors.Is(err, syscall.ECONNREFUSED):
		return "Connection refused. Run \"nc -vz a.b.c.d PORT\" on the Teleport " +
			"agent to verify the target application is running and listening on " +
			"the expected host and port.", true
	// Host unreachable errors can be reproduced by running
	// "ip route add unreachable a.b.c.d" to update the routing table to make
	// the host unreachable. Packets will be discarded and an ICMP message
	// will be returned. The raw error typically looks like the following:
	//
	// dial tcp 10.10.10.10:8000: connect: no route to host
	case errors.Is(err, syscall.EHOSTUNREACH):
		return "No route to host. Run \"ip route get a.b.c.d\" on the Teleport " +
			"agent to verify a route to the target application exists in the " +
			"routing table.", true
	// Connection reset errors can be reproduced by creating a HTTP server that
	// accepts requests but closes the connection before writing a response. The
	// raw error typically looks like the following:
	//
	// read tcp 127.0.0.1:49764->127.0.0.1:8000: read: connection reset by peer
	case errors.Is(err, syscall.ECONNRESET):
		return "Connection reset by peer. Run \"curl -v a.b.c.d\" on the Teleport " +
			"agent to verify the target application (or a load balancer in the " +
			"network path) is not abruptly closing the connection after accepting it.", true
	// Slow responses can be reprodued by creating a HTTP server that does a
	// time.Sleep before responding. The raw error typically looks like the following:
	//
	// context deadline exceeded
	case errors.Is(err, context.DeadlineExceeded):
		return "Timeout waiting for response. Run \"curl -v a.b.c.d\" on the " +
			"Teleport agent to verify the target application is not under excessive load.", true
	// No such host errors can be reproduced by attempting to resolve a invalid
	// domain name. The raw error typically looks like the following:
	//
	// dial tcp: lookup fasfasfasf.com: no such host
	case errors.As(err, &derr) && derr.IsNotFound:
		return "No such host. Run \"dig +short fqdn\" on the Teleport agent to " +
			"verify the target application has a valid DNS entry.", true
	}

	return "", false
}

const (
	// SelfSignedCertsMsg is a helper message to point users towards helpful documentation.
	SelfSignedCertsMsg = "Your proxy certificate is not trusted or expired. " +
		"Please update the certificate or follow this guide for self-signed certs: https://goteleport.com/docs/management/admin/self-signed-certs/"
)
