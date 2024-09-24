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
// @generated from file teleport/attestation/v1/attestation.proto (package teleport.attestation.v1, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, messageDesc } from "@bufbuild/protobuf/codegenv1";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file teleport/attestation/v1/attestation.proto.
 */
export const file_teleport_attestation_v1_attestation: GenFile = /*@__PURE__*/
  fileDesc("Cil0ZWxlcG9ydC9hdHRlc3RhdGlvbi92MS9hdHRlc3RhdGlvbi5wcm90bxIXdGVsZXBvcnQuYXR0ZXN0YXRpb24udjEijgEKFEF0dGVzdGF0aW9uU3RhdGVtZW50El0KHXl1YmlrZXlfYXR0ZXN0YXRpb25fc3RhdGVtZW50GAEgASgLMjQudGVsZXBvcnQuYXR0ZXN0YXRpb24udjEuWXViaUtleUF0dGVzdGF0aW9uU3RhdGVtZW50SABCFwoVYXR0ZXN0YXRpb25fc3RhdGVtZW50IkoKG1l1YmlLZXlBdHRlc3RhdGlvblN0YXRlbWVudBIRCglzbG90X2NlcnQYASABKAwSGAoQYXR0ZXN0YXRpb25fY2VydBgCIAEoDEJRWk9naXRodWIuY29tL2dyYXZpdGF0aW9uYWwvdGVsZXBvcnQvYXBpL2dlbi9wcm90by9nby9hdHRlc3RhdGlvbi92MTthdHRlc3RhdGlvbnYxYgZwcm90bzM");

/**
 * AttestationStatement is an attestation statement for a hardware private key.
 *
 * @generated from message teleport.attestation.v1.AttestationStatement
 */
export type AttestationStatement = Message<"teleport.attestation.v1.AttestationStatement"> & {
  /**
   * @generated from oneof teleport.attestation.v1.AttestationStatement.attestation_statement
   */
  attestationStatement: {
    /**
     * yubikey_attestation_statement is an attestation statement for a specific YubiKey PIV slot.
     *
     * @generated from field: teleport.attestation.v1.YubiKeyAttestationStatement yubikey_attestation_statement = 1;
     */
    value: YubiKeyAttestationStatement;
    case: "yubikeyAttestationStatement";
  } | { case: undefined; value?: undefined };
};

/**
 * Describes the message teleport.attestation.v1.AttestationStatement.
 * Use `create(AttestationStatementSchema)` to create a new message.
 */
export const AttestationStatementSchema: GenMessage<AttestationStatement> = /*@__PURE__*/
  messageDesc(file_teleport_attestation_v1_attestation, 0);

/**
 * YubiKeyAttestationStatement is an attestation statement for a specific YubiKey PIV slot.
 *
 * @generated from message teleport.attestation.v1.YubiKeyAttestationStatement
 */
export type YubiKeyAttestationStatement = Message<"teleport.attestation.v1.YubiKeyAttestationStatement"> & {
  /**
   * slot_cert is an attestation certificate generated from a YubiKey PIV
   * slot's public key and signed by the YubiKey's attestation certificate.
   *
   * @generated from field: bytes slot_cert = 1;
   */
  slotCert: Uint8Array;

  /**
   * attestation_cert is the YubiKey's unique attestation certificate, signed by a Yubico CA.
   *
   * @generated from field: bytes attestation_cert = 2;
   */
  attestationCert: Uint8Array;
};

/**
 * Describes the message teleport.attestation.v1.YubiKeyAttestationStatement.
 * Use `create(YubiKeyAttestationStatementSchema)` to create a new message.
 */
export const YubiKeyAttestationStatementSchema: GenMessage<YubiKeyAttestationStatement> = /*@__PURE__*/
  messageDesc(file_teleport_attestation_v1_attestation, 1);

