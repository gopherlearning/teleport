// Copyright 2022 Gravitational, Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// @generated by protoc-gen-es v2.1.0 with parameter "target=ts"
// @generated from file teleport/legacy/client/proto/joinservice.proto (package proto, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, messageDesc, serviceDesc } from "@bufbuild/protobuf/codegenv1";
import type { Certs } from "./certs_pb";
import { file_teleport_legacy_client_proto_certs } from "./certs_pb";
import type { RegisterUsingTokenRequest } from "../../types/types_pb";
import { file_teleport_legacy_types_types } from "../../types/types_pb";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file teleport/legacy/client/proto/joinservice.proto.
 */
export const file_teleport_legacy_client_proto_joinservice: GenFile = /*@__PURE__*/
  fileDesc("Ci50ZWxlcG9ydC9sZWdhY3kvY2xpZW50L3Byb3RvL2pvaW5zZXJ2aWNlLnByb3RvEgVwcm90byKFAQodUmVnaXN0ZXJVc2luZ0lBTU1ldGhvZFJlcXVlc3QSRgoccmVnaXN0ZXJfdXNpbmdfdG9rZW5fcmVxdWVzdBgBIAEoCzIgLnR5cGVzLlJlZ2lzdGVyVXNpbmdUb2tlblJlcXVlc3QSHAoUc3RzX2lkZW50aXR5X3JlcXVlc3QYAiABKAwiUAoeUmVnaXN0ZXJVc2luZ0lBTU1ldGhvZFJlc3BvbnNlEhEKCWNoYWxsZW5nZRgBIAEoCRIbCgVjZXJ0cxgCIAEoCzIMLnByb3RvLkNlcnRzIpYBCh9SZWdpc3RlclVzaW5nQXp1cmVNZXRob2RSZXF1ZXN0EkYKHHJlZ2lzdGVyX3VzaW5nX3Rva2VuX3JlcXVlc3QYASABKAsyIC50eXBlcy5SZWdpc3RlclVzaW5nVG9rZW5SZXF1ZXN0EhUKDWF0dGVzdGVkX2RhdGEYAiABKAwSFAoMYWNjZXNzX3Rva2VuGAMgASgJIlIKIFJlZ2lzdGVyVXNpbmdBenVyZU1ldGhvZFJlc3BvbnNlEhEKCWNoYWxsZW5nZRgBIAEoCRIbCgVjZXJ0cxgCIAEoCzIMLnByb3RvLkNlcnRzIjsKJ1JlZ2lzdGVyVXNpbmdUUE1NZXRob2RDaGFsbGVuZ2VSZXNwb25zZRIQCghzb2x1dGlvbhgBIAEoDCLGAQokUmVnaXN0ZXJVc2luZ1RQTU1ldGhvZEluaXRpYWxSZXF1ZXN0EjYKDGpvaW5fcmVxdWVzdBgBIAEoCzIgLnR5cGVzLlJlZ2lzdGVyVXNpbmdUb2tlblJlcXVlc3QSEQoHZWtfY2VydBgCIAEoDEgAEhAKBmVrX2tleRgDIAEoDEgAEjsKEmF0dGVzdGF0aW9uX3BhcmFtcxgEIAEoCzIfLnByb3RvLlRQTUF0dGVzdGF0aW9uUGFyYW1ldGVyc0IECgJlayK1AQodUmVnaXN0ZXJVc2luZ1RQTU1ldGhvZFJlcXVlc3QSOwoEaW5pdBgBIAEoCzIrLnByb3RvLlJlZ2lzdGVyVXNpbmdUUE1NZXRob2RJbml0aWFsUmVxdWVzdEgAEkwKEmNoYWxsZW5nZV9yZXNwb25zZRgCIAEoCzIuLnByb3RvLlJlZ2lzdGVyVXNpbmdUUE1NZXRob2RDaGFsbGVuZ2VSZXNwb25zZUgAQgkKB3BheWxvYWQihgEKHlJlZ2lzdGVyVXNpbmdUUE1NZXRob2RSZXNwb25zZRI6ChFjaGFsbGVuZ2VfcmVxdWVzdBgBIAEoCzIdLnByb3RvLlRQTUVuY3J5cHRlZENyZWRlbnRpYWxIABIdCgVjZXJ0cxgCIAEoCzIMLnByb3RvLkNlcnRzSABCCQoHcGF5bG9hZCJ1ChhUUE1BdHRlc3RhdGlvblBhcmFtZXRlcnMSDgoGcHVibGljGAEgASgMEhMKC2NyZWF0ZV9kYXRhGAIgASgMEhoKEmNyZWF0ZV9hdHRlc3RhdGlvbhgDIAEoDBIYChBjcmVhdGVfc2lnbmF0dXJlGAQgASgMIkEKFlRQTUVuY3J5cHRlZENyZWRlbnRpYWwSFwoPY3JlZGVudGlhbF9ibG9iGAEgASgMEg4KBnNlY3JldBgCIAEoDDLUAgoLSm9pblNlcnZpY2USaQoWUmVnaXN0ZXJVc2luZ0lBTU1ldGhvZBIkLnByb3RvLlJlZ2lzdGVyVXNpbmdJQU1NZXRob2RSZXF1ZXN0GiUucHJvdG8uUmVnaXN0ZXJVc2luZ0lBTU1ldGhvZFJlc3BvbnNlKAEwARJvChhSZWdpc3RlclVzaW5nQXp1cmVNZXRob2QSJi5wcm90by5SZWdpc3RlclVzaW5nQXp1cmVNZXRob2RSZXF1ZXN0GicucHJvdG8uUmVnaXN0ZXJVc2luZ0F6dXJlTWV0aG9kUmVzcG9uc2UoATABEmkKFlJlZ2lzdGVyVXNpbmdUUE1NZXRob2QSJC5wcm90by5SZWdpc3RlclVzaW5nVFBNTWV0aG9kUmVxdWVzdBolLnByb3RvLlJlZ2lzdGVyVXNpbmdUUE1NZXRob2RSZXNwb25zZSgBMAFCNFoyZ2l0aHViLmNvbS9ncmF2aXRhdGlvbmFsL3RlbGVwb3J0L2FwaS9jbGllbnQvcHJvdG9iBnByb3RvMw", [file_teleport_legacy_client_proto_certs, file_teleport_legacy_types_types]);

/**
 * TODO(nklaassen): Document me.
 *
 * @generated from message proto.RegisterUsingIAMMethodRequest
 */
export type RegisterUsingIAMMethodRequest = Message<"proto.RegisterUsingIAMMethodRequest"> & {
  /**
   * RegisterUsingTokenRequest holds registration parameters common to all
   * join methods.
   *
   * @generated from field: types.RegisterUsingTokenRequest register_using_token_request = 1;
   */
  registerUsingTokenRequest?: RegisterUsingTokenRequest;

  /**
   * StsIdentityRequest is a signed HTTP request to the AWS
   * sts:GetCallerIdentity API endpoint used to prove the AWS identity of a
   * joining node. It must include the challenge string as a signed header.
   *
   * @generated from field: bytes sts_identity_request = 2;
   */
  stsIdentityRequest: Uint8Array;
};

/**
 * Describes the message proto.RegisterUsingIAMMethodRequest.
 * Use `create(RegisterUsingIAMMethodRequestSchema)` to create a new message.
 */
export const RegisterUsingIAMMethodRequestSchema: GenMessage<RegisterUsingIAMMethodRequest> = /*@__PURE__*/
  messageDesc(file_teleport_legacy_client_proto_joinservice, 0);

/**
 * RegisterUsingIAMMethodResponse is a stream response and will contain either a
 * Challenge or signed Certs to join the cluster.
 *
 * @generated from message proto.RegisterUsingIAMMethodResponse
 */
export type RegisterUsingIAMMethodResponse = Message<"proto.RegisterUsingIAMMethodResponse"> & {
  /**
   * Challenge is a crypto-random string that should be included in the signed
   * sts:GetCallerIdentity request.
   *
   * @generated from field: string challenge = 1;
   */
  challenge: string;

  /**
   * Certs is the returned signed certs.
   *
   * @generated from field: proto.Certs certs = 2;
   */
  certs?: Certs;
};

/**
 * Describes the message proto.RegisterUsingIAMMethodResponse.
 * Use `create(RegisterUsingIAMMethodResponseSchema)` to create a new message.
 */
export const RegisterUsingIAMMethodResponseSchema: GenMessage<RegisterUsingIAMMethodResponse> = /*@__PURE__*/
  messageDesc(file_teleport_legacy_client_proto_joinservice, 1);

/**
 * RegisterUsingAzureMethodRequest is the request for registration via the Azure
 * join method.
 *
 * @generated from message proto.RegisterUsingAzureMethodRequest
 */
export type RegisterUsingAzureMethodRequest = Message<"proto.RegisterUsingAzureMethodRequest"> & {
  /**
   * RegisterUsingTokenRequest holds registration parameters common to all
   * join methods.
   *
   * @generated from field: types.RegisterUsingTokenRequest register_using_token_request = 1;
   */
  registerUsingTokenRequest?: RegisterUsingTokenRequest;

  /**
   * AttestedData is a signed JSON document from an Azure VM's attested data
   * metadata endpoint used to prove the identity of a joining node. It must
   * include the challenge string as the nonce.
   *
   * @generated from field: bytes attested_data = 2;
   */
  attestedData: Uint8Array;

  /**
   * AccessToken is a JWT signed by Azure, used to prove the identity of a
   * joining node.
   *
   * @generated from field: string access_token = 3;
   */
  accessToken: string;
};

/**
 * Describes the message proto.RegisterUsingAzureMethodRequest.
 * Use `create(RegisterUsingAzureMethodRequestSchema)` to create a new message.
 */
export const RegisterUsingAzureMethodRequestSchema: GenMessage<RegisterUsingAzureMethodRequest> = /*@__PURE__*/
  messageDesc(file_teleport_legacy_client_proto_joinservice, 2);

/**
 * RegisterUsingAzureMethodResponse is a stream response and will contain either
 * a Challenge or signed Certs to join the cluster.
 *
 * @generated from message proto.RegisterUsingAzureMethodResponse
 */
export type RegisterUsingAzureMethodResponse = Message<"proto.RegisterUsingAzureMethodResponse"> & {
  /**
   * Challenge is a crypto-random string that should be included in the signed
   * attested data.
   *
   * @generated from field: string challenge = 1;
   */
  challenge: string;

  /**
   * Certs is the returned signed certs.
   *
   * @generated from field: proto.Certs certs = 2;
   */
  certs?: Certs;
};

/**
 * Describes the message proto.RegisterUsingAzureMethodResponse.
 * Use `create(RegisterUsingAzureMethodResponseSchema)` to create a new message.
 */
export const RegisterUsingAzureMethodResponseSchema: GenMessage<RegisterUsingAzureMethodResponse> = /*@__PURE__*/
  messageDesc(file_teleport_legacy_client_proto_joinservice, 3);

/**
 * The enrollment challenge response containing the solution returned by
 * calling the TPM2.0 `ActivateCredential` command on the client with the
 * parameters provided in `TPMEncryptedCredential`.
 *
 * @generated from message proto.RegisterUsingTPMMethodChallengeResponse
 */
export type RegisterUsingTPMMethodChallengeResponse = Message<"proto.RegisterUsingTPMMethodChallengeResponse"> & {
  /**
   * The client's solution to `TPMEncryptedCredential` included in
   * `TPMEncryptedCredential` using ActivateCredential.
   *
   * @generated from field: bytes solution = 1;
   */
  solution: Uint8Array;
};

/**
 * Describes the message proto.RegisterUsingTPMMethodChallengeResponse.
 * Use `create(RegisterUsingTPMMethodChallengeResponseSchema)` to create a new message.
 */
export const RegisterUsingTPMMethodChallengeResponseSchema: GenMessage<RegisterUsingTPMMethodChallengeResponse> = /*@__PURE__*/
  messageDesc(file_teleport_legacy_client_proto_joinservice, 4);

/**
 * The initial payload sent from the client to the server during a TPM join
 * request.
 *
 * @generated from message proto.RegisterUsingTPMMethodInitialRequest
 */
export type RegisterUsingTPMMethodInitialRequest = Message<"proto.RegisterUsingTPMMethodInitialRequest"> & {
  /**
   * Holds the registration parameters shared by all join methods.
   *
   * @generated from field: types.RegisterUsingTokenRequest join_request = 1;
   */
  joinRequest?: RegisterUsingTokenRequest;

  /**
   * @generated from oneof proto.RegisterUsingTPMMethodInitialRequest.ek
   */
  ek: {
    /**
     * The device's endorsement certificate in X509, ASN.1 DER form. This
     * certificate contains the public key of the endorsement key. This is
     * preferred to ek_key.
     *
     * @generated from field: bytes ek_cert = 2;
     */
    value: Uint8Array;
    case: "ekCert";
  } | {
    /**
     * The device's public endorsement key in PKIX, ASN.1 DER form. This is
     * used when a TPM does not contain any endorsement certificates.
     *
     * @generated from field: bytes ek_key = 3;
     */
    value: Uint8Array;
    case: "ekKey";
  } | { case: undefined; value?: undefined };

  /**
   * The attestation key and the parameters necessary to remotely verify it as
   * related to the endorsement key.
   *
   * @generated from field: proto.TPMAttestationParameters attestation_params = 4;
   */
  attestationParams?: TPMAttestationParameters;
};

/**
 * Describes the message proto.RegisterUsingTPMMethodInitialRequest.
 * Use `create(RegisterUsingTPMMethodInitialRequestSchema)` to create a new message.
 */
export const RegisterUsingTPMMethodInitialRequestSchema: GenMessage<RegisterUsingTPMMethodInitialRequest> = /*@__PURE__*/
  messageDesc(file_teleport_legacy_client_proto_joinservice, 5);

/**
 * RegisterUsingTPMMethodRequest is the streaming request type for the
 * RegisterUsingTPMMethod RPC.
 *
 * @generated from message proto.RegisterUsingTPMMethodRequest
 */
export type RegisterUsingTPMMethodRequest = Message<"proto.RegisterUsingTPMMethodRequest"> & {
  /**
   * @generated from oneof proto.RegisterUsingTPMMethodRequest.payload
   */
  payload: {
    /**
     * Initial information sent from the client to the server.
     *
     * @generated from field: proto.RegisterUsingTPMMethodInitialRequest init = 1;
     */
    value: RegisterUsingTPMMethodInitialRequest;
    case: "init";
  } | {
    /**
     * The challenge response required to complete the TPM join process. This is
     * sent in response to the servers challenge.
     *
     * @generated from field: proto.RegisterUsingTPMMethodChallengeResponse challenge_response = 2;
     */
    value: RegisterUsingTPMMethodChallengeResponse;
    case: "challengeResponse";
  } | { case: undefined; value?: undefined };
};

/**
 * Describes the message proto.RegisterUsingTPMMethodRequest.
 * Use `create(RegisterUsingTPMMethodRequestSchema)` to create a new message.
 */
export const RegisterUsingTPMMethodRequestSchema: GenMessage<RegisterUsingTPMMethodRequest> = /*@__PURE__*/
  messageDesc(file_teleport_legacy_client_proto_joinservice, 6);

/**
 * RegisterUsingTPMMethodResponse is the streaming response type for the
 * RegisterUsingTPMMethod RPC.
 *
 * @generated from message proto.RegisterUsingTPMMethodResponse
 */
export type RegisterUsingTPMMethodResponse = Message<"proto.RegisterUsingTPMMethodResponse"> & {
  /**
   * @generated from oneof proto.RegisterUsingTPMMethodResponse.payload
   */
  payload: {
    /**
     * The challenge required to complete the TPM join process. This is sent to
     * the client in response to the initial request.
     *
     * @generated from field: proto.TPMEncryptedCredential challenge_request = 1;
     */
    value: TPMEncryptedCredential;
    case: "challengeRequest";
  } | {
    /**
     * The signed certificates resulting from the join process.
     *
     * @generated from field: proto.Certs certs = 2;
     */
    value: Certs;
    case: "certs";
  } | { case: undefined; value?: undefined };
};

/**
 * Describes the message proto.RegisterUsingTPMMethodResponse.
 * Use `create(RegisterUsingTPMMethodResponseSchema)` to create a new message.
 */
export const RegisterUsingTPMMethodResponseSchema: GenMessage<RegisterUsingTPMMethodResponse> = /*@__PURE__*/
  messageDesc(file_teleport_legacy_client_proto_joinservice, 7);

/**
 * The attestation key and the parameters necessary to remotely verify it as
 * related to the endorsement key.
 * See https://pkg.go.dev/github.com/google/go-attestation/attest#AttestationParameters.
 * This message excludes the `UseTCSDActivationFormat` field from the link above
 * as it is TMP 1.x specific and always false.
 *
 * @generated from message proto.TPMAttestationParameters
 */
export type TPMAttestationParameters = Message<"proto.TPMAttestationParameters"> & {
  /**
   * The encoded TPMT_PUBLIC structure containing the attestation public key
   * and signing parameters.
   *
   * @generated from field: bytes public = 1;
   */
  public: Uint8Array;

  /**
   * The properties of the attestation key, encoded as a TPMS_CREATION_DATA
   * structure.
   *
   * @generated from field: bytes create_data = 2;
   */
  createData: Uint8Array;

  /**
   * An assertion as to the details of the key, encoded as a TPMS_ATTEST
   * structure.
   *
   * @generated from field: bytes create_attestation = 3;
   */
  createAttestation: Uint8Array;

  /**
   * A signature of create_attestation, encoded as a TPMT_SIGNATURE structure.
   *
   * @generated from field: bytes create_signature = 4;
   */
  createSignature: Uint8Array;
};

/**
 * Describes the message proto.TPMAttestationParameters.
 * Use `create(TPMAttestationParametersSchema)` to create a new message.
 */
export const TPMAttestationParametersSchema: GenMessage<TPMAttestationParameters> = /*@__PURE__*/
  messageDesc(file_teleport_legacy_client_proto_joinservice, 8);

/**
 * These values are used by the TPM2.0 `ActivateCredential` command to produce
 * the solution which proves possession of the EK and AK.
 *
 * For a more in-depth description see:
 * - https://pkg.go.dev/github.com/google/go-attestation/attest#EncryptedCredential
 * - https://trustedcomputinggroup.org/wp-content/uploads/TCG_TPM2_r1p59_Part3_Commands_code_pub.pdf (Heading 12.5.1 "TPM2_ActivateCredential" "General Description")
 * - https://github.com/google/go-attestation/blob/v0.4.3/attest/activation.go#L199
 * - https://github.com/google/go-tpm/blob/v0.3.3/tpm2/credactivation/credential_activation.go#L61
 *
 * @generated from message proto.TPMEncryptedCredential
 */
export type TPMEncryptedCredential = Message<"proto.TPMEncryptedCredential"> & {
  /**
   * The `credential_blob` parameter to be used with the `ActivateCredential`
   * command. This is used with the decrypted value of `secret` in a
   * cryptographic process to decrypt the solution.
   *
   * @generated from field: bytes credential_blob = 1;
   */
  credentialBlob: Uint8Array;

  /**
   * The `secret` parameter to be used with `ActivateCredential`. This is a
   * seed which can be decrypted with the EK. The decrypted seed is then used
   * when decrypting `credential_blob`.
   *
   * @generated from field: bytes secret = 2;
   */
  secret: Uint8Array;
};

/**
 * Describes the message proto.TPMEncryptedCredential.
 * Use `create(TPMEncryptedCredentialSchema)` to create a new message.
 */
export const TPMEncryptedCredentialSchema: GenMessage<TPMEncryptedCredential> = /*@__PURE__*/
  messageDesc(file_teleport_legacy_client_proto_joinservice, 9);

/**
 * JoinService provides methods which allow Teleport nodes, proxies, and other
 * services to join the Teleport cluster by fetching signed cluster
 * certificates. It is implemented on both the Auth and Proxy servers to serve
 * the needs of both nodes connecting directly to the Auth server and IoT mode
 * nodes connecting only to the Proxy.
 *
 * @generated from service proto.JoinService
 */
export const JoinService: GenService<{
  /**
   * RegisterUsingIAMMethod is used to register a new node to the cluster using
   * the IAM join method.
   *
   * @generated from rpc proto.JoinService.RegisterUsingIAMMethod
   */
  registerUsingIAMMethod: {
    methodKind: "bidi_streaming";
    input: typeof RegisterUsingIAMMethodRequestSchema;
    output: typeof RegisterUsingIAMMethodResponseSchema;
  },
  /**
   * RegisterUsingAzureMethod is used to register a new node to the cluster
   * using the Azure join method.
   *
   * @generated from rpc proto.JoinService.RegisterUsingAzureMethod
   */
  registerUsingAzureMethod: {
    methodKind: "bidi_streaming";
    input: typeof RegisterUsingAzureMethodRequestSchema;
    output: typeof RegisterUsingAzureMethodResponseSchema;
  },
  /**
   * RegisterUsingTPMMethod allows registration of a new agent or Bot to the
   * cluster using a known TPM.
   *
   * @generated from rpc proto.JoinService.RegisterUsingTPMMethod
   */
  registerUsingTPMMethod: {
    methodKind: "bidi_streaming";
    input: typeof RegisterUsingTPMMethodRequestSchema;
    output: typeof RegisterUsingTPMMethodResponseSchema;
  },
}> = /*@__PURE__*/
  serviceDesc(file_teleport_legacy_client_proto_joinservice, 0);

