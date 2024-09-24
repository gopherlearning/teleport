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
// @generated from file teleport/machineid/v1/bot_instance_service.proto (package teleport.machineid.v1, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, messageDesc, serviceDesc } from "@bufbuild/protobuf/codegenv1";
import type { EmptySchema } from "@bufbuild/protobuf/wkt";
import { file_google_protobuf_empty } from "@bufbuild/protobuf/wkt";
import type { BotInstance, BotInstanceSchema, BotInstanceStatusHeartbeat } from "./bot_instance_pb";
import { file_teleport_machineid_v1_bot_instance } from "./bot_instance_pb";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file teleport/machineid/v1/bot_instance_service.proto.
 */
export const file_teleport_machineid_v1_bot_instance_service: GenFile = /*@__PURE__*/
  fileDesc("CjB0ZWxlcG9ydC9tYWNoaW5laWQvdjEvYm90X2luc3RhbmNlX3NlcnZpY2UucHJvdG8SFXRlbGVwb3J0Lm1hY2hpbmVpZC52MSI+ChVHZXRCb3RJbnN0YW5jZVJlcXVlc3QSEAoIYm90X25hbWUYASABKAkSEwoLaW5zdGFuY2VfaWQYAiABKAkiWQoXTGlzdEJvdEluc3RhbmNlc1JlcXVlc3QSFwoPZmlsdGVyX2JvdF9uYW1lGAEgASgJEhEKCXBhZ2Vfc2l6ZRgCIAEoBRISCgpwYWdlX3Rva2VuGAMgASgJIm4KGExpc3RCb3RJbnN0YW5jZXNSZXNwb25zZRI5Cg1ib3RfaW5zdGFuY2VzGAEgAygLMiIudGVsZXBvcnQubWFjaGluZWlkLnYxLkJvdEluc3RhbmNlEhcKD25leHRfcGFnZV90b2tlbhgCIAEoCSJBChhEZWxldGVCb3RJbnN0YW5jZVJlcXVlc3QSEAoIYm90X25hbWUYASABKAkSEwoLaW5zdGFuY2VfaWQYAiABKAkiXgoWU3VibWl0SGVhcnRiZWF0UmVxdWVzdBJECgloZWFydGJlYXQYASABKAsyMS50ZWxlcG9ydC5tYWNoaW5laWQudjEuQm90SW5zdGFuY2VTdGF0dXNIZWFydGJlYXQiGQoXU3VibWl0SGVhcnRiZWF0UmVzcG9uc2UyvQMKEkJvdEluc3RhbmNlU2VydmljZRJiCg5HZXRCb3RJbnN0YW5jZRIsLnRlbGVwb3J0Lm1hY2hpbmVpZC52MS5HZXRCb3RJbnN0YW5jZVJlcXVlc3QaIi50ZWxlcG9ydC5tYWNoaW5laWQudjEuQm90SW5zdGFuY2UScwoQTGlzdEJvdEluc3RhbmNlcxIuLnRlbGVwb3J0Lm1hY2hpbmVpZC52MS5MaXN0Qm90SW5zdGFuY2VzUmVxdWVzdBovLnRlbGVwb3J0Lm1hY2hpbmVpZC52MS5MaXN0Qm90SW5zdGFuY2VzUmVzcG9uc2USXAoRRGVsZXRlQm90SW5zdGFuY2USLy50ZWxlcG9ydC5tYWNoaW5laWQudjEuRGVsZXRlQm90SW5zdGFuY2VSZXF1ZXN0GhYuZ29vZ2xlLnByb3RvYnVmLkVtcHR5EnAKD1N1Ym1pdEhlYXJ0YmVhdBItLnRlbGVwb3J0Lm1hY2hpbmVpZC52MS5TdWJtaXRIZWFydGJlYXRSZXF1ZXN0Gi4udGVsZXBvcnQubWFjaGluZWlkLnYxLlN1Ym1pdEhlYXJ0YmVhdFJlc3BvbnNlQlZaVGdpdGh1Yi5jb20vZ3Jhdml0YXRpb25hbC90ZWxlcG9ydC9hcGkvZ2VuL3Byb3RvL2dvL3RlbGVwb3J0L21hY2hpbmVpZC92MTttYWNoaW5laWR2MWIGcHJvdG8z", [file_google_protobuf_empty, file_teleport_machineid_v1_bot_instance]);

/**
 * Request for GetBotInstance.
 *
 * @generated from message teleport.machineid.v1.GetBotInstanceRequest
 */
export type GetBotInstanceRequest = Message<"teleport.machineid.v1.GetBotInstanceRequest"> & {
  /**
   * The name of the bot associated with the instance.
   *
   * @generated from field: string bot_name = 1;
   */
  botName: string;

  /**
   * The unique identifier of the bot instance to retrieve.
   *
   * @generated from field: string instance_id = 2;
   */
  instanceId: string;
};

/**
 * Describes the message teleport.machineid.v1.GetBotInstanceRequest.
 * Use `create(GetBotInstanceRequestSchema)` to create a new message.
 */
export const GetBotInstanceRequestSchema: GenMessage<GetBotInstanceRequest> = /*@__PURE__*/
  messageDesc(file_teleport_machineid_v1_bot_instance_service, 0);

/**
 * Request for ListBotInstances.
 *
 * Follows the pagination semantics of
 * https://cloud.google.com/apis/design/standard_methods#list
 *
 * @generated from message teleport.machineid.v1.ListBotInstancesRequest
 */
export type ListBotInstancesRequest = Message<"teleport.machineid.v1.ListBotInstancesRequest"> & {
  /**
   * The name of the Bot to list BotInstances for. If empty, all BotInstances
   * will be listed.
   *
   * @generated from field: string filter_bot_name = 1;
   */
  filterBotName: string;

  /**
   * The maximum number of items to return.
   * The server may impose a different page size at its discretion.
   *
   * @generated from field: int32 page_size = 2;
   */
  pageSize: number;

  /**
   * The page_token value returned from a previous ListBotInstances request, if
   * any.
   *
   * @generated from field: string page_token = 3;
   */
  pageToken: string;
};

/**
 * Describes the message teleport.machineid.v1.ListBotInstancesRequest.
 * Use `create(ListBotInstancesRequestSchema)` to create a new message.
 */
export const ListBotInstancesRequestSchema: GenMessage<ListBotInstancesRequest> = /*@__PURE__*/
  messageDesc(file_teleport_machineid_v1_bot_instance_service, 1);

/**
 * Response for ListBotInstances.
 *
 * @generated from message teleport.machineid.v1.ListBotInstancesResponse
 */
export type ListBotInstancesResponse = Message<"teleport.machineid.v1.ListBotInstancesResponse"> & {
  /**
   * BotInstance that matched the search.
   *
   * @generated from field: repeated teleport.machineid.v1.BotInstance bot_instances = 1;
   */
  botInstances: BotInstance[];

  /**
   * Token to retrieve the next page of results, or empty if there are no
   * more results exist.
   *
   * @generated from field: string next_page_token = 2;
   */
  nextPageToken: string;
};

/**
 * Describes the message teleport.machineid.v1.ListBotInstancesResponse.
 * Use `create(ListBotInstancesResponseSchema)` to create a new message.
 */
export const ListBotInstancesResponseSchema: GenMessage<ListBotInstancesResponse> = /*@__PURE__*/
  messageDesc(file_teleport_machineid_v1_bot_instance_service, 2);

/**
 * Request for DeleteBotInstance.
 *
 * @generated from message teleport.machineid.v1.DeleteBotInstanceRequest
 */
export type DeleteBotInstanceRequest = Message<"teleport.machineid.v1.DeleteBotInstanceRequest"> & {
  /**
   * The name of the BotInstance to delete.
   *
   * @generated from field: string bot_name = 1;
   */
  botName: string;

  /**
   * The unique identifier of the bot instance to delete.
   *
   * @generated from field: string instance_id = 2;
   */
  instanceId: string;
};

/**
 * Describes the message teleport.machineid.v1.DeleteBotInstanceRequest.
 * Use `create(DeleteBotInstanceRequestSchema)` to create a new message.
 */
export const DeleteBotInstanceRequestSchema: GenMessage<DeleteBotInstanceRequest> = /*@__PURE__*/
  messageDesc(file_teleport_machineid_v1_bot_instance_service, 3);

/**
 * The request for SubmitHeartbeat.
 *
 * @generated from message teleport.machineid.v1.SubmitHeartbeatRequest
 */
export type SubmitHeartbeatRequest = Message<"teleport.machineid.v1.SubmitHeartbeatRequest"> & {
  /**
   * The heartbeat data to submit.
   *
   * @generated from field: teleport.machineid.v1.BotInstanceStatusHeartbeat heartbeat = 1;
   */
  heartbeat?: BotInstanceStatusHeartbeat;
};

/**
 * Describes the message teleport.machineid.v1.SubmitHeartbeatRequest.
 * Use `create(SubmitHeartbeatRequestSchema)` to create a new message.
 */
export const SubmitHeartbeatRequestSchema: GenMessage<SubmitHeartbeatRequest> = /*@__PURE__*/
  messageDesc(file_teleport_machineid_v1_bot_instance_service, 4);

/**
 * The response for SubmitHeartbeat.
 *
 * Empty
 *
 * @generated from message teleport.machineid.v1.SubmitHeartbeatResponse
 */
export type SubmitHeartbeatResponse = Message<"teleport.machineid.v1.SubmitHeartbeatResponse"> & {
};

/**
 * Describes the message teleport.machineid.v1.SubmitHeartbeatResponse.
 * Use `create(SubmitHeartbeatResponseSchema)` to create a new message.
 */
export const SubmitHeartbeatResponseSchema: GenMessage<SubmitHeartbeatResponse> = /*@__PURE__*/
  messageDesc(file_teleport_machineid_v1_bot_instance_service, 5);

/**
 * BotInstanceService provides functions to record and manage bot instances.
 *
 * @generated from service teleport.machineid.v1.BotInstanceService
 */
export const BotInstanceService: GenService<{
  /**
   * GetBotInstance returns the specified BotInstance resource.
   *
   * @generated from rpc teleport.machineid.v1.BotInstanceService.GetBotInstance
   */
  getBotInstance: {
    methodKind: "unary";
    input: typeof GetBotInstanceRequestSchema;
    output: typeof BotInstanceSchema;
  },
  /**
   * ListBotInstances returns a page of BotInstance resources.
   *
   * @generated from rpc teleport.machineid.v1.BotInstanceService.ListBotInstances
   */
  listBotInstances: {
    methodKind: "unary";
    input: typeof ListBotInstancesRequestSchema;
    output: typeof ListBotInstancesResponseSchema;
  },
  /**
   * DeleteBotInstance hard deletes the specified BotInstance resource.
   *
   * @generated from rpc teleport.machineid.v1.BotInstanceService.DeleteBotInstance
   */
  deleteBotInstance: {
    methodKind: "unary";
    input: typeof DeleteBotInstanceRequestSchema;
    output: typeof EmptySchema;
  },
  /**
   * SubmitHeartbeat submits a heartbeat for a BotInstance.
   *
   * @generated from rpc teleport.machineid.v1.BotInstanceService.SubmitHeartbeat
   */
  submitHeartbeat: {
    methodKind: "unary";
    input: typeof SubmitHeartbeatRequestSchema;
    output: typeof SubmitHeartbeatResponseSchema;
  },
}> = /*@__PURE__*/
  serviceDesc(file_teleport_machineid_v1_bot_instance_service, 0);

