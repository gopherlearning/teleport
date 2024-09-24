// Copyright 2023 Gravitational, Inc
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

// @generated by protoc-gen-connect-es v1.5.0 with parameter "target=ts"
// @generated from file teleport/trust/v1/trust_service.proto (package teleport.trust.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { DeleteCertAuthorityRequest, GenerateHostCertRequest, GenerateHostCertResponse, GetCertAuthoritiesRequest, GetCertAuthoritiesResponse, GetCertAuthorityRequest, RotateCertAuthorityRequest, RotateCertAuthorityResponse, RotateExternalCertAuthorityRequest, RotateExternalCertAuthorityResponse, UpsertCertAuthorityRequest } from "./trust_service_pb.js";
import { CertAuthorityV2 } from "../../legacy/types/types_pb.js";
import { Empty, MethodKind } from "@bufbuild/protobuf";

/**
 * TrustService provides methods to manage certificate authorities.
 *
 * @generated from service teleport.trust.v1.TrustService
 */
export const TrustService = {
  typeName: "teleport.trust.v1.TrustService",
  methods: {
    /**
     * GetCertAuthority returns a cert authority by type and domain.
     *
     * @generated from rpc teleport.trust.v1.TrustService.GetCertAuthority
     */
    getCertAuthority: {
      name: "GetCertAuthority",
      I: GetCertAuthorityRequest,
      O: CertAuthorityV2,
      kind: MethodKind.Unary,
    },
    /**
     * GetCertAuthorities returns all cert authorities with the specified type.
     *
     * @generated from rpc teleport.trust.v1.TrustService.GetCertAuthorities
     */
    getCertAuthorities: {
      name: "GetCertAuthorities",
      I: GetCertAuthoritiesRequest,
      O: GetCertAuthoritiesResponse,
      kind: MethodKind.Unary,
    },
    /**
     * DeleteCertAuthority deletes the matching cert authority.
     *
     * @generated from rpc teleport.trust.v1.TrustService.DeleteCertAuthority
     */
    deleteCertAuthority: {
      name: "DeleteCertAuthority",
      I: DeleteCertAuthorityRequest,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * UpsertCertAuthority creates or updates the provided cert authority.
     *
     * @generated from rpc teleport.trust.v1.TrustService.UpsertCertAuthority
     */
    upsertCertAuthority: {
      name: "UpsertCertAuthority",
      I: UpsertCertAuthorityRequest,
      O: CertAuthorityV2,
      kind: MethodKind.Unary,
    },
    /**
     * RotateCertAuthority is a request to start rotation of the certificate authority.
     *
     * @generated from rpc teleport.trust.v1.TrustService.RotateCertAuthority
     */
    rotateCertAuthority: {
      name: "RotateCertAuthority",
      I: RotateCertAuthorityRequest,
      O: RotateCertAuthorityResponse,
      kind: MethodKind.Unary,
    },
    /**
     * RotateExternalCertAuthority rotates an external cert authority.
     *
     * @generated from rpc teleport.trust.v1.TrustService.RotateExternalCertAuthority
     */
    rotateExternalCertAuthority: {
      name: "RotateExternalCertAuthority",
      I: RotateExternalCertAuthorityRequest,
      O: RotateExternalCertAuthorityResponse,
      kind: MethodKind.Unary,
    },
    /**
     * GenerateHostCert takes a public key in the OpenSSH `authorized_keys` format and returns
     * a SSH certificate signed by the Host CA.
     *
     * @generated from rpc teleport.trust.v1.TrustService.GenerateHostCert
     */
    generateHostCert: {
      name: "GenerateHostCert",
      I: GenerateHostCertRequest,
      O: GenerateHostCertResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

