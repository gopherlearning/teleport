// Copyright 2024 Gravitational, Inc
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
// @generated from file teleport/machineid/v1/federation_service.proto (package teleport.machineid.v1, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, messageDesc, serviceDesc } from "@bufbuild/protobuf/codegenv1";
import type { EmptySchema } from "@bufbuild/protobuf/wkt";
import { file_google_protobuf_empty } from "@bufbuild/protobuf/wkt";
import type { SPIFFEFederation, SPIFFEFederationSchema } from "./federation_pb";
import { file_teleport_machineid_v1_federation } from "./federation_pb";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file teleport/machineid/v1/federation_service.proto.
 */
export const file_teleport_machineid_v1_federation_service: GenFile = /*@__PURE__*/
  fileDesc("Ci50ZWxlcG9ydC9tYWNoaW5laWQvdjEvZmVkZXJhdGlvbl9zZXJ2aWNlLnByb3RvEhV0ZWxlcG9ydC5tYWNoaW5laWQudjEiKgoaR2V0U1BJRkZFRmVkZXJhdGlvblJlcXVlc3QSDAoEbmFtZRgBIAEoCSJFChxMaXN0U1BJRkZFRmVkZXJhdGlvbnNSZXF1ZXN0EhEKCXBhZ2Vfc2l6ZRgBIAEoBRISCgpwYWdlX3Rva2VuGAIgASgJIn0KHUxpc3RTUElGRkVGZWRlcmF0aW9uc1Jlc3BvbnNlEkMKEnNwaWZmZV9mZWRlcmF0aW9ucxgBIAMoCzInLnRlbGVwb3J0Lm1hY2hpbmVpZC52MS5TUElGRkVGZWRlcmF0aW9uEhcKD25leHRfcGFnZV90b2tlbhgCIAEoCSItCh1EZWxldGVTUElGRkVGZWRlcmF0aW9uUmVxdWVzdBIMCgRuYW1lGAEgASgJImMKHUNyZWF0ZVNQSUZGRUZlZGVyYXRpb25SZXF1ZXN0EkIKEXNwaWZmZV9mZWRlcmF0aW9uGAEgASgLMicudGVsZXBvcnQubWFjaGluZWlkLnYxLlNQSUZGRUZlZGVyYXRpb24y8gMKF1NQSUZGRUZlZGVyYXRpb25TZXJ2aWNlEnEKE0dldFNQSUZGRUZlZGVyYXRpb24SMS50ZWxlcG9ydC5tYWNoaW5laWQudjEuR2V0U1BJRkZFRmVkZXJhdGlvblJlcXVlc3QaJy50ZWxlcG9ydC5tYWNoaW5laWQudjEuU1BJRkZFRmVkZXJhdGlvbhKCAQoVTGlzdFNQSUZGRUZlZGVyYXRpb25zEjMudGVsZXBvcnQubWFjaGluZWlkLnYxLkxpc3RTUElGRkVGZWRlcmF0aW9uc1JlcXVlc3QaNC50ZWxlcG9ydC5tYWNoaW5laWQudjEuTGlzdFNQSUZGRUZlZGVyYXRpb25zUmVzcG9uc2USZgoWRGVsZXRlU1BJRkZFRmVkZXJhdGlvbhI0LnRlbGVwb3J0Lm1hY2hpbmVpZC52MS5EZWxldGVTUElGRkVGZWRlcmF0aW9uUmVxdWVzdBoWLmdvb2dsZS5wcm90b2J1Zi5FbXB0eRJ3ChZDcmVhdGVTUElGRkVGZWRlcmF0aW9uEjQudGVsZXBvcnQubWFjaGluZWlkLnYxLkNyZWF0ZVNQSUZGRUZlZGVyYXRpb25SZXF1ZXN0GicudGVsZXBvcnQubWFjaGluZWlkLnYxLlNQSUZGRUZlZGVyYXRpb25CVlpUZ2l0aHViLmNvbS9ncmF2aXRhdGlvbmFsL3RlbGVwb3J0L2FwaS9nZW4vcHJvdG8vZ28vdGVsZXBvcnQvbWFjaGluZWlkL3YxO21hY2hpbmVpZHYxYgZwcm90bzM", [file_google_protobuf_empty, file_teleport_machineid_v1_federation]);

/**
 * GetSPIFFEFederationRequest is the request message for GetSPIFFEFederation.
 *
 * @generated from message teleport.machineid.v1.GetSPIFFEFederationRequest
 */
export type GetSPIFFEFederationRequest = Message<"teleport.machineid.v1.GetSPIFFEFederationRequest"> & {
  /**
   * The name of the SPIFFEFederation resource to fetch.
   *
   * @generated from field: string name = 1;
   */
  name: string;
};

/**
 * Describes the message teleport.machineid.v1.GetSPIFFEFederationRequest.
 * Use `create(GetSPIFFEFederationRequestSchema)` to create a new message.
 */
export const GetSPIFFEFederationRequestSchema: GenMessage<GetSPIFFEFederationRequest> = /*@__PURE__*/
  messageDesc(file_teleport_machineid_v1_federation_service, 0);

/**
 * Request for ListSPIFFEFederations.
 *
 * Follows the pagination semantics of
 * https://cloud.google.com/apis/design/standard_methods#list
 *
 * @generated from message teleport.machineid.v1.ListSPIFFEFederationsRequest
 */
export type ListSPIFFEFederationsRequest = Message<"teleport.machineid.v1.ListSPIFFEFederationsRequest"> & {
  /**
   * The maximum number of items to return.
   * The server may impose a different page size at its discretion.
   *
   * @generated from field: int32 page_size = 1;
   */
  pageSize: number;

  /**
   * The page_token value returned from a previous ListSPIFFEFederations
   * request, if any.
   *
   * @generated from field: string page_token = 2;
   */
  pageToken: string;
};

/**
 * Describes the message teleport.machineid.v1.ListSPIFFEFederationsRequest.
 * Use `create(ListSPIFFEFederationsRequestSchema)` to create a new message.
 */
export const ListSPIFFEFederationsRequestSchema: GenMessage<ListSPIFFEFederationsRequest> = /*@__PURE__*/
  messageDesc(file_teleport_machineid_v1_federation_service, 1);

/**
 * ListSPIFFEFederationsResponse is the response message for ListSPIFFEFederations.
 *
 * @generated from message teleport.machineid.v1.ListSPIFFEFederationsResponse
 */
export type ListSPIFFEFederationsResponse = Message<"teleport.machineid.v1.ListSPIFFEFederationsResponse"> & {
  /**
   * @generated from field: repeated teleport.machineid.v1.SPIFFEFederation spiffe_federations = 1;
   */
  spiffeFederations: SPIFFEFederation[];

  /**
   * Token to retrieve the next page of results, or empty if there are no
   * more results exist.
   *
   * @generated from field: string next_page_token = 2;
   */
  nextPageToken: string;
};

/**
 * Describes the message teleport.machineid.v1.ListSPIFFEFederationsResponse.
 * Use `create(ListSPIFFEFederationsResponseSchema)` to create a new message.
 */
export const ListSPIFFEFederationsResponseSchema: GenMessage<ListSPIFFEFederationsResponse> = /*@__PURE__*/
  messageDesc(file_teleport_machineid_v1_federation_service, 2);

/**
 * DeleteSPIFFEFederationRequest is the request message for DeleteSPIFFEFederation.
 *
 * @generated from message teleport.machineid.v1.DeleteSPIFFEFederationRequest
 */
export type DeleteSPIFFEFederationRequest = Message<"teleport.machineid.v1.DeleteSPIFFEFederationRequest"> & {
  /**
   * The name of the SPIFFEFederation resource to delete.
   *
   * @generated from field: string name = 1;
   */
  name: string;
};

/**
 * Describes the message teleport.machineid.v1.DeleteSPIFFEFederationRequest.
 * Use `create(DeleteSPIFFEFederationRequestSchema)` to create a new message.
 */
export const DeleteSPIFFEFederationRequestSchema: GenMessage<DeleteSPIFFEFederationRequest> = /*@__PURE__*/
  messageDesc(file_teleport_machineid_v1_federation_service, 3);

/**
 * CreateSPIFFEFederationRequest is the request message for CreateSPIFFEFederation.
 *
 * @generated from message teleport.machineid.v1.CreateSPIFFEFederationRequest
 */
export type CreateSPIFFEFederationRequest = Message<"teleport.machineid.v1.CreateSPIFFEFederationRequest"> & {
  /**
   * The SPIFFEFederation resource to create.
   *
   * @generated from field: teleport.machineid.v1.SPIFFEFederation spiffe_federation = 1;
   */
  spiffeFederation?: SPIFFEFederation;
};

/**
 * Describes the message teleport.machineid.v1.CreateSPIFFEFederationRequest.
 * Use `create(CreateSPIFFEFederationRequestSchema)` to create a new message.
 */
export const CreateSPIFFEFederationRequestSchema: GenMessage<CreateSPIFFEFederationRequest> = /*@__PURE__*/
  messageDesc(file_teleport_machineid_v1_federation_service, 4);

/**
 * SPIFFEFederationService provides methods to manage SPIFFE Federations
 * between trust domains.
 *
 * @generated from service teleport.machineid.v1.SPIFFEFederationService
 */
export const SPIFFEFederationService: GenService<{
  /**
   * GetSPIFFEFederation returns a SPIFFEFederation resource by name.
   *
   * @generated from rpc teleport.machineid.v1.SPIFFEFederationService.GetSPIFFEFederation
   */
  getSPIFFEFederation: {
    methodKind: "unary";
    input: typeof GetSPIFFEFederationRequestSchema;
    output: typeof SPIFFEFederationSchema;
  },
  /**
   * ListSPIFFEFederations returns a list of SPIFFEFederation resources.
   * Follows the pagination semantics of
   * https://cloud.google.com/apis/design/design_patterns#list_pagination
   *
   * @generated from rpc teleport.machineid.v1.SPIFFEFederationService.ListSPIFFEFederations
   */
  listSPIFFEFederations: {
    methodKind: "unary";
    input: typeof ListSPIFFEFederationsRequestSchema;
    output: typeof ListSPIFFEFederationsResponseSchema;
  },
  /**
   * DeleteSPIFFEFederation deletes a SPIFFEFederation resource by name.
   *
   * @generated from rpc teleport.machineid.v1.SPIFFEFederationService.DeleteSPIFFEFederation
   */
  deleteSPIFFEFederation: {
    methodKind: "unary";
    input: typeof DeleteSPIFFEFederationRequestSchema;
    output: typeof EmptySchema;
  },
  /**
   * CreateSPIFFEFederation creates a SPIFFEFederation resource.
   *
   * @generated from rpc teleport.machineid.v1.SPIFFEFederationService.CreateSPIFFEFederation
   */
  createSPIFFEFederation: {
    methodKind: "unary";
    input: typeof CreateSPIFFEFederationRequestSchema;
    output: typeof SPIFFEFederationSchema;
  },
}> = /*@__PURE__*/
  serviceDesc(file_teleport_machineid_v1_federation_service, 0);

