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

// @generated by protoc-gen-es v2.1.0 with parameter "target=ts"
// @generated from file teleport/secreports/v1/secreports_service.proto (package teleport.secreports.v1, syntax proto3)
/* eslint-disable */

import type { GenEnum, GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv1";
import { enumDesc, fileDesc, messageDesc, serviceDesc } from "@bufbuild/protobuf/codegenv1";
import type { EmptySchema } from "@bufbuild/protobuf/wkt";
import { file_google_protobuf_empty } from "@bufbuild/protobuf/wkt";
import type { ResourceHeader } from "../../header/v1/resourceheader_pb";
import { file_teleport_header_v1_resourceheader } from "../../header/v1/resourceheader_pb";
import type { AuditQuery, AuditQuerySchema, AuditQuerySpec, Report, ReportSchema, ReportStateSchema } from "./secreports_pb";
import { file_teleport_secreports_v1_secreports } from "./secreports_pb";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file teleport/secreports/v1/secreports_service.proto.
 */
export const file_teleport_secreports_v1_secreports_service: GenFile = /*@__PURE__*/
  fileDesc("Ci90ZWxlcG9ydC9zZWNyZXBvcnRzL3YxL3NlY3JlcG9ydHNfc2VydmljZS5wcm90bxIWdGVsZXBvcnQuc2VjcmVwb3J0cy52MSJYChpHZXRBdWRpdFF1ZXJ5UmVzdWx0UmVxdWVzdBIRCglyZXN1bHRfaWQYASABKAkSEgoKbmV4dF90b2tlbhgCIAEoCRITCgttYXhfcmVzdWx0cxgDIAEoBSIzChVRdWVyeVJlc3VsdENvbHVtbkluZm8SDAoEbmFtZRgBIAEoCRIMCgR0eXBlGAIgASgJIh4KDlF1ZXJ5Um93UmVzdWx0EgwKBGRhdGEYASADKAkiigEKDlF1ZXJ5UmVzdWx0U2V0EkIKC2NvbHVtbl9pbmZvGAEgAygLMi0udGVsZXBvcnQuc2VjcmVwb3J0cy52MS5RdWVyeVJlc3VsdENvbHVtbkluZm8SNAoEcm93cxgCIAMoCzImLnRlbGVwb3J0LnNlY3JlcG9ydHMudjEuUXVlcnlSb3dSZXN1bHQifAobR2V0QXVkaXRRdWVyeVJlc3VsdFJlc3BvbnNlEjYKBnJlc3VsdBgBIAEoCzImLnRlbGVwb3J0LnNlY3JlcG9ydHMudjEuUXVlcnlSZXN1bHRTZXQSEgoKbmV4dF90b2tlbhgCIAEoCRIRCglyZXN1bHRfaWQYAyABKAkiLgoQUnVuUmVwb3J0UmVxdWVzdBIMCgRuYW1lGAEgASgJEgwKBGRheXMYAiABKA0iMwoVR2V0UmVwb3J0U3RhdGVSZXF1ZXN0EgwKBG5hbWUYASABKAkSDAoEZGF5cxgCIAEoDSInChdEZWxldGVBdWRpdFF1ZXJ5UmVxdWVzdBIMCgRuYW1lGAEgASgJIiMKE0RlbGV0ZVJlcG9ydFJlcXVlc3QSDAoEbmFtZRgBIAEoCSIzChRSdW5BdWRpdFF1ZXJ5UmVxdWVzdBINCgVxdWVyeRgBIAEoCRIMCgRkYXlzGAIgASgFIlIKF1Vwc2VydEF1ZGl0UXVlcnlSZXF1ZXN0EjcKC2F1ZGl0X3F1ZXJ5GAEgASgLMiIudGVsZXBvcnQuc2VjcmVwb3J0cy52MS5BdWRpdFF1ZXJ5IkUKE1Vwc2VydFJlcG9ydFJlcXVlc3QSLgoGcmVwb3J0GAEgASgLMh4udGVsZXBvcnQuc2VjcmVwb3J0cy52MS5SZXBvcnQiJAoUR2V0QXVkaXRRdWVyeVJlcXVlc3QSDAoEbmFtZRgBIAEoCSIgChBHZXRSZXBvcnRSZXF1ZXN0EgwKBG5hbWUYASABKAkiSgoWR2V0UmVwb3J0UmVzdWx0UmVxdWVzdBIMCgRuYW1lGAEgASgJEgwKBGRheXMYAiABKA0SFAoMZXhlY3V0aW9uX2lkGAMgASgJIkAKF0xpc3RBdWRpdFF1ZXJpZXNSZXF1ZXN0EhEKCXBhZ2Vfc2l6ZRgBIAEoBRISCgpwYWdlX3Rva2VuGAIgASgJIjsKEkxpc3RSZXBvcnRzUmVxdWVzdBIRCglwYWdlX3NpemUYASABKAUSEgoKcGFnZV90b2tlbhgCIAEoCSJoChhMaXN0QXVkaXRRdWVyaWVzUmVzcG9uc2USMwoHcXVlcmllcxgBIAMoCzIiLnRlbGVwb3J0LnNlY3JlcG9ydHMudjEuQXVkaXRRdWVyeRIXCg9uZXh0X3BhZ2VfdG9rZW4YAiABKAkiEgoQR2V0U2NoZW1hUmVxdWVzdCKHAgoRR2V0U2NoZW1hUmVzcG9uc2USQQoFdmlld3MYASADKAsyMi50ZWxlcG9ydC5zZWNyZXBvcnRzLnYxLkdldFNjaGVtYVJlc3BvbnNlLlZpZXdEZXNjGq4BCghWaWV3RGVzYxIMCgRuYW1lGAEgASgJEgwKBGRlc2MYAiABKAkSTgoHY29sdW1ucxgDIAMoCzI9LnRlbGVwb3J0LnNlY3JlcG9ydHMudjEuR2V0U2NoZW1hUmVzcG9uc2UuVmlld0Rlc2MuQ29sdW1uRGVzYxo2CgpDb2x1bW5EZXNjEgwKBG5hbWUYASABKAkSDAoEdHlwZRgCIAEoCRIMCgRkZXNjGAMgASgJIioKFVJ1bkF1ZGl0UXVlcnlSZXNwb25zZRIRCglyZXN1bHRfaWQYASABKAkiXwoTTGlzdFJlcG9ydHNSZXNwb25zZRIvCgdyZXBvcnRzGAEgAygLMh4udGVsZXBvcnQuc2VjcmVwb3J0cy52MS5SZXBvcnQSFwoPbmV4dF9wYWdlX3Rva2VuGAIgASgJIk8KF0dldFJlcG9ydFJlc3VsdFJlc3BvbnNlEjQKBnJlc3VsdBgBIAEoCzIkLnRlbGVwb3J0LnNlY3JlcG9ydHMudjEuUmVwb3J0UmVzdWx0IsQDCgxSZXBvcnRSZXN1bHQSDAoEbmFtZRgBIAEoCRITCgtkZXNjcmlwdGlvbhgCIAEoCRJSChNhdWRpdF9xdWVyeV9yZXN1bHRzGAMgAygLMjUudGVsZXBvcnQuc2VjcmVwb3J0cy52MS5SZXBvcnRSZXN1bHQuQXVkaXRRdWVyeVJlc3VsdBISCgp1cGRhdGVkX2F0GAQgASgJEiYKHnRvdGFsX2V4ZWN1dGlvbl90aW1lX2luX21pbGxpcxgFIAEoAxIjCht0b3RhbF9kYXRhX3NjYW5uZWRfaW5fYnl0ZXMYBiABKAMa2wEKEEF1ZGl0UXVlcnlSZXN1bHQSOwoLYXVkaXRfcXVlcnkYASABKAsyJi50ZWxlcG9ydC5zZWNyZXBvcnRzLnYxLkF1ZGl0UXVlcnlTcGVjEjYKBnJlc3VsdBgCIAEoCzImLnRlbGVwb3J0LnNlY3JlcG9ydHMudjEuUXVlcnlSZXN1bHRTZXQSEQoJcmVzdWx0X2lkGAMgASgJEiAKGGV4ZWN1dGlvbl90aW1lX2luX21pbGxpcxgEIAEoAxIdChVkYXRhX3NjYW5uZWRfaW5fYnl0ZXMYBSABKAMi5AEKClJlcG9ydFNhdGUSMgoGaGVhZGVyGAEgASgLMiIudGVsZXBvcnQuaGVhZGVyLnYxLlJlc291cmNlSGVhZGVyEjcKBXN0YXRlGAIgASgOMigudGVsZXBvcnQuc2VjcmVwb3J0cy52MS5SZXBvcnRTYXRlLlN0YXRlEhIKCnVwZGF0ZWRfYXQYAyABKAkiVQoFU3RhdGUSFQoRU1RBVEVfVU5TUEVDSUZJRUQQABIPCgtTVEFURV9FUlJPUhABEhEKDVNUQVRFX1NVQ0NFU1MQAhIRCg1TVEFURV9SVU5OSU5HEAMyiQsKEVNlY1JlcG9ydHNTZXJ2aWNlElsKEFVwc2VydEF1ZGl0UXVlcnkSLy50ZWxlcG9ydC5zZWNyZXBvcnRzLnYxLlVwc2VydEF1ZGl0UXVlcnlSZXF1ZXN0GhYuZ29vZ2xlLnByb3RvYnVmLkVtcHR5EmEKDUdldEF1ZGl0UXVlcnkSLC50ZWxlcG9ydC5zZWNyZXBvcnRzLnYxLkdldEF1ZGl0UXVlcnlSZXF1ZXN0GiIudGVsZXBvcnQuc2VjcmVwb3J0cy52MS5BdWRpdFF1ZXJ5EnUKEExpc3RBdWRpdFF1ZXJpZXMSLy50ZWxlcG9ydC5zZWNyZXBvcnRzLnYxLkxpc3RBdWRpdFF1ZXJpZXNSZXF1ZXN0GjAudGVsZXBvcnQuc2VjcmVwb3J0cy52MS5MaXN0QXVkaXRRdWVyaWVzUmVzcG9uc2USWwoQRGVsZXRlQXVkaXRRdWVyeRIvLnRlbGVwb3J0LnNlY3JlcG9ydHMudjEuRGVsZXRlQXVkaXRRdWVyeVJlcXVlc3QaFi5nb29nbGUucHJvdG9idWYuRW1wdHkSUwoMVXBzZXJ0UmVwb3J0EisudGVsZXBvcnQuc2VjcmVwb3J0cy52MS5VcHNlcnRSZXBvcnRSZXF1ZXN0GhYuZ29vZ2xlLnByb3RvYnVmLkVtcHR5ElUKCUdldFJlcG9ydBIoLnRlbGVwb3J0LnNlY3JlcG9ydHMudjEuR2V0UmVwb3J0UmVxdWVzdBoeLnRlbGVwb3J0LnNlY3JlcG9ydHMudjEuUmVwb3J0EmYKC0xpc3RSZXBvcnRzEioudGVsZXBvcnQuc2VjcmVwb3J0cy52MS5MaXN0UmVwb3J0c1JlcXVlc3QaKy50ZWxlcG9ydC5zZWNyZXBvcnRzLnYxLkxpc3RSZXBvcnRzUmVzcG9uc2USUwoMRGVsZXRlUmVwb3J0EisudGVsZXBvcnQuc2VjcmVwb3J0cy52MS5EZWxldGVSZXBvcnRSZXF1ZXN0GhYuZ29vZ2xlLnByb3RvYnVmLkVtcHR5EmwKDVJ1bkF1ZGl0UXVlcnkSLC50ZWxlcG9ydC5zZWNyZXBvcnRzLnYxLlJ1bkF1ZGl0UXVlcnlSZXF1ZXN0Gi0udGVsZXBvcnQuc2VjcmVwb3J0cy52MS5SdW5BdWRpdFF1ZXJ5UmVzcG9uc2USfgoTR2V0QXVkaXRRdWVyeVJlc3VsdBIyLnRlbGVwb3J0LnNlY3JlcG9ydHMudjEuR2V0QXVkaXRRdWVyeVJlc3VsdFJlcXVlc3QaMy50ZWxlcG9ydC5zZWNyZXBvcnRzLnYxLkdldEF1ZGl0UXVlcnlSZXN1bHRSZXNwb25zZRJNCglSdW5SZXBvcnQSKC50ZWxlcG9ydC5zZWNyZXBvcnRzLnYxLlJ1blJlcG9ydFJlcXVlc3QaFi5nb29nbGUucHJvdG9idWYuRW1wdHkScgoPR2V0UmVwb3J0UmVzdWx0Ei4udGVsZXBvcnQuc2VjcmVwb3J0cy52MS5HZXRSZXBvcnRSZXN1bHRSZXF1ZXN0Gi8udGVsZXBvcnQuc2VjcmVwb3J0cy52MS5HZXRSZXBvcnRSZXN1bHRSZXNwb25zZRJkCg5HZXRSZXBvcnRTdGF0ZRItLnRlbGVwb3J0LnNlY3JlcG9ydHMudjEuR2V0UmVwb3J0U3RhdGVSZXF1ZXN0GiMudGVsZXBvcnQuc2VjcmVwb3J0cy52MS5SZXBvcnRTdGF0ZRJgCglHZXRTY2hlbWESKC50ZWxlcG9ydC5zZWNyZXBvcnRzLnYxLkdldFNjaGVtYVJlcXVlc3QaKS50ZWxlcG9ydC5zZWNyZXBvcnRzLnYxLkdldFNjaGVtYVJlc3BvbnNlQlhaVmdpdGh1Yi5jb20vZ3Jhdml0YXRpb25hbC90ZWxlcG9ydC9hcGkvZ2VuL3Byb3RvL2dvL3RlbGVwb3J0L3NlY3JlcG9ydHMvdjE7c2VjcmVwb3J0c3YxYgZwcm90bzM", [file_google_protobuf_empty, file_teleport_header_v1_resourceheader, file_teleport_secreports_v1_secreports]);

/**
 * GetAuditQueryResultRequest is a request for GetAuditQueryResult.
 *
 * @generated from message teleport.secreports.v1.GetAuditQueryResultRequest
 */
export type GetAuditQueryResultRequest = Message<"teleport.secreports.v1.GetAuditQueryResultRequest"> & {
  /**
   * result_id is a unique id of the result.
   *
   * @generated from field: string result_id = 1;
   */
  resultId: string;

  /**
   * next_token is a token for pagination.
   *
   * @generated from field: string next_token = 2;
   */
  nextToken: string;

  /**
   * max_results is a maximum number of results to return.
   *
   * @generated from field: int32 max_results = 3;
   */
  maxResults: number;
};

/**
 * Describes the message teleport.secreports.v1.GetAuditQueryResultRequest.
 * Use `create(GetAuditQueryResultRequestSchema)` to create a new message.
 */
export const GetAuditQueryResultRequestSchema: GenMessage<GetAuditQueryResultRequest> = /*@__PURE__*/
  messageDesc(file_teleport_secreports_v1_secreports_service, 0);

/**
 * QueryResultColumnInfo is a column info.
 *
 * @generated from message teleport.secreports.v1.QueryResultColumnInfo
 */
export type QueryResultColumnInfo = Message<"teleport.secreports.v1.QueryResultColumnInfo"> & {
  /**
   * name is name of the column.
   *
   * @generated from field: string name = 1;
   */
  name: string;

  /**
   * type is type of the column.
   *
   * @generated from field: string type = 2;
   */
  type: string;
};

/**
 * Describes the message teleport.secreports.v1.QueryResultColumnInfo.
 * Use `create(QueryResultColumnInfoSchema)` to create a new message.
 */
export const QueryResultColumnInfoSchema: GenMessage<QueryResultColumnInfo> = /*@__PURE__*/
  messageDesc(file_teleport_secreports_v1_secreports_service, 1);

/**
 * QueryRowResult is a row result.
 *
 * @generated from message teleport.secreports.v1.QueryRowResult
 */
export type QueryRowResult = Message<"teleport.secreports.v1.QueryRowResult"> & {
  /**
   * data is a list of values.
   *
   * @generated from field: repeated string data = 1;
   */
  data: string[];
};

/**
 * Describes the message teleport.secreports.v1.QueryRowResult.
 * Use `create(QueryRowResultSchema)` to create a new message.
 */
export const QueryRowResultSchema: GenMessage<QueryRowResult> = /*@__PURE__*/
  messageDesc(file_teleport_secreports_v1_secreports_service, 2);

/**
 * QueryResultSet is a result set.
 *
 * @generated from message teleport.secreports.v1.QueryResultSet
 */
export type QueryResultSet = Message<"teleport.secreports.v1.QueryResultSet"> & {
  /**
   * column_info contains information about columns.
   *
   * @generated from field: repeated teleport.secreports.v1.QueryResultColumnInfo column_info = 1;
   */
  columnInfo: QueryResultColumnInfo[];

  /**
   * rows  is a list of rows containing values.
   *
   * @generated from field: repeated teleport.secreports.v1.QueryRowResult rows = 2;
   */
  rows: QueryRowResult[];
};

/**
 * Describes the message teleport.secreports.v1.QueryResultSet.
 * Use `create(QueryResultSetSchema)` to create a new message.
 */
export const QueryResultSetSchema: GenMessage<QueryResultSet> = /*@__PURE__*/
  messageDesc(file_teleport_secreports_v1_secreports_service, 3);

/**
 * GetAuditQueryResultResponse contains an audit query result.
 *
 * @generated from message teleport.secreports.v1.GetAuditQueryResultResponse
 */
export type GetAuditQueryResultResponse = Message<"teleport.secreports.v1.GetAuditQueryResultResponse"> & {
  /**
   * result is a result set.
   *
   * @generated from field: teleport.secreports.v1.QueryResultSet result = 1;
   */
  result?: QueryResultSet;

  /**
   * next_token is a token for pagination.
   *
   * @generated from field: string next_token = 2;
   */
  nextToken: string;

  /**
   * result_id is a unique id of the result.
   *
   * @generated from field: string result_id = 3;
   */
  resultId: string;
};

/**
 * Describes the message teleport.secreports.v1.GetAuditQueryResultResponse.
 * Use `create(GetAuditQueryResultResponseSchema)` to create a new message.
 */
export const GetAuditQueryResultResponseSchema: GenMessage<GetAuditQueryResultResponse> = /*@__PURE__*/
  messageDesc(file_teleport_secreports_v1_secreports_service, 4);

/**
 * RunReportRequest is a request for RunReport.
 *
 * @generated from message teleport.secreports.v1.RunReportRequest
 */
export type RunReportRequest = Message<"teleport.secreports.v1.RunReportRequest"> & {
  /**
   * name is a name of the security report.
   *
   * @generated from field: string name = 1;
   */
  name: string;

  /**
   * days is a time range is days.
   *
   * @generated from field: uint32 days = 2;
   */
  days: number;
};

/**
 * Describes the message teleport.secreports.v1.RunReportRequest.
 * Use `create(RunReportRequestSchema)` to create a new message.
 */
export const RunReportRequestSchema: GenMessage<RunReportRequest> = /*@__PURE__*/
  messageDesc(file_teleport_secreports_v1_secreports_service, 5);

/**
 * GetStateRequest is a request for GetReportState.
 *
 * @generated from message teleport.secreports.v1.GetReportStateRequest
 */
export type GetReportStateRequest = Message<"teleport.secreports.v1.GetReportStateRequest"> & {
  /**
   * name is a name of the security report.
   *
   * @generated from field: string name = 1;
   */
  name: string;

  /**
   * days is a time range is days.
   *
   * @generated from field: uint32 days = 2;
   */
  days: number;
};

/**
 * Describes the message teleport.secreports.v1.GetReportStateRequest.
 * Use `create(GetReportStateRequestSchema)` to create a new message.
 */
export const GetReportStateRequestSchema: GenMessage<GetReportStateRequest> = /*@__PURE__*/
  messageDesc(file_teleport_secreports_v1_secreports_service, 6);

/**
 * DeleteAuditQueryRequest is audit query delete request.
 *
 * @generated from message teleport.secreports.v1.DeleteAuditQueryRequest
 */
export type DeleteAuditQueryRequest = Message<"teleport.secreports.v1.DeleteAuditQueryRequest"> & {
  /**
   * name is the name of the audit query to delete.
   *
   * @generated from field: string name = 1;
   */
  name: string;
};

/**
 * Describes the message teleport.secreports.v1.DeleteAuditQueryRequest.
 * Use `create(DeleteAuditQueryRequestSchema)` to create a new message.
 */
export const DeleteAuditQueryRequestSchema: GenMessage<DeleteAuditQueryRequest> = /*@__PURE__*/
  messageDesc(file_teleport_secreports_v1_secreports_service, 7);

/**
 * DeleteReportRequest is security report delete request.
 *
 * @generated from message teleport.secreports.v1.DeleteReportRequest
 */
export type DeleteReportRequest = Message<"teleport.secreports.v1.DeleteReportRequest"> & {
  /**
   * name is the name of the security report to delete.
   *
   * @generated from field: string name = 1;
   */
  name: string;
};

/**
 * Describes the message teleport.secreports.v1.DeleteReportRequest.
 * Use `create(DeleteReportRequestSchema)` to create a new message.
 */
export const DeleteReportRequestSchema: GenMessage<DeleteReportRequest> = /*@__PURE__*/
  messageDesc(file_teleport_secreports_v1_secreports_service, 8);

/**
 * RunAuditQueryRequest is audit query run request.
 *
 * @generated from message teleport.secreports.v1.RunAuditQueryRequest
 */
export type RunAuditQueryRequest = Message<"teleport.secreports.v1.RunAuditQueryRequest"> & {
  /**
   * name is the name of the audit query to run.
   *
   * @generated from field: string query = 1;
   */
  query: string;

  /**
   * days is a time range is days.
   *
   * @generated from field: int32 days = 2;
   */
  days: number;
};

/**
 * Describes the message teleport.secreports.v1.RunAuditQueryRequest.
 * Use `create(RunAuditQueryRequestSchema)` to create a new message.
 */
export const RunAuditQueryRequestSchema: GenMessage<RunAuditQueryRequest> = /*@__PURE__*/
  messageDesc(file_teleport_secreports_v1_secreports_service, 9);

/**
 * UpsertAuditQueryRequest is audit query upsert request.
 *
 * @generated from message teleport.secreports.v1.UpsertAuditQueryRequest
 */
export type UpsertAuditQueryRequest = Message<"teleport.secreports.v1.UpsertAuditQueryRequest"> & {
  /**
   * audit_query is the audit query to upsert.
   *
   * @generated from field: teleport.secreports.v1.AuditQuery audit_query = 1;
   */
  auditQuery?: AuditQuery;
};

/**
 * Describes the message teleport.secreports.v1.UpsertAuditQueryRequest.
 * Use `create(UpsertAuditQueryRequestSchema)` to create a new message.
 */
export const UpsertAuditQueryRequestSchema: GenMessage<UpsertAuditQueryRequest> = /*@__PURE__*/
  messageDesc(file_teleport_secreports_v1_secreports_service, 10);

/**
 * UpsertReportRequest is security report upsert request.
 *
 * @generated from message teleport.secreports.v1.UpsertReportRequest
 */
export type UpsertReportRequest = Message<"teleport.secreports.v1.UpsertReportRequest"> & {
  /**
   * report is the security report to upsert.
   *
   * @generated from field: teleport.secreports.v1.Report report = 1;
   */
  report?: Report;
};

/**
 * Describes the message teleport.secreports.v1.UpsertReportRequest.
 * Use `create(UpsertReportRequestSchema)` to create a new message.
 */
export const UpsertReportRequestSchema: GenMessage<UpsertReportRequest> = /*@__PURE__*/
  messageDesc(file_teleport_secreports_v1_secreports_service, 11);

/**
 * GetAuditQueryRequest is audit query get request.
 *
 * @generated from message teleport.secreports.v1.GetAuditQueryRequest
 */
export type GetAuditQueryRequest = Message<"teleport.secreports.v1.GetAuditQueryRequest"> & {
  /**
   * name is the name of the audit query to get.
   *
   * @generated from field: string name = 1;
   */
  name: string;
};

/**
 * Describes the message teleport.secreports.v1.GetAuditQueryRequest.
 * Use `create(GetAuditQueryRequestSchema)` to create a new message.
 */
export const GetAuditQueryRequestSchema: GenMessage<GetAuditQueryRequest> = /*@__PURE__*/
  messageDesc(file_teleport_secreports_v1_secreports_service, 12);

/**
 * GetReportRequest is security report get request.
 *
 * @generated from message teleport.secreports.v1.GetReportRequest
 */
export type GetReportRequest = Message<"teleport.secreports.v1.GetReportRequest"> & {
  /**
   * name is the name of the security report to get.
   *
   * @generated from field: string name = 1;
   */
  name: string;
};

/**
 * Describes the message teleport.secreports.v1.GetReportRequest.
 * Use `create(GetReportRequestSchema)` to create a new message.
 */
export const GetReportRequestSchema: GenMessage<GetReportRequest> = /*@__PURE__*/
  messageDesc(file_teleport_secreports_v1_secreports_service, 13);

/**
 * GetReportResultRequest is report get request.
 *
 * @generated from message teleport.secreports.v1.GetReportResultRequest
 */
export type GetReportResultRequest = Message<"teleport.secreports.v1.GetReportResultRequest"> & {
  /**
   * name is the name of the security report to get.
   *
   * @generated from field: string name = 1;
   */
  name: string;

  /**
   * days is a time range is days.
   *
   * @generated from field: uint32 days = 2;
   */
  days: number;

  /**
   * execution_id is a unique id of the execution.
   *
   * @generated from field: string execution_id = 3;
   */
  executionId: string;
};

/**
 * Describes the message teleport.secreports.v1.GetReportResultRequest.
 * Use `create(GetReportResultRequestSchema)` to create a new message.
 */
export const GetReportResultRequestSchema: GenMessage<GetReportResultRequest> = /*@__PURE__*/
  messageDesc(file_teleport_secreports_v1_secreports_service, 14);

/**
 * ListAuditQueriesRequest is audit query list request.
 *
 * @generated from message teleport.secreports.v1.ListAuditQueriesRequest
 */
export type ListAuditQueriesRequest = Message<"teleport.secreports.v1.ListAuditQueriesRequest"> & {
  /**
   * page_size is the number of results to return.
   *
   * @generated from field: int32 page_size = 1;
   */
  pageSize: number;

  /**
   * page_token is the next_token value returned from a previous List request if any.
   *
   * @generated from field: string page_token = 2;
   */
  pageToken: string;
};

/**
 * Describes the message teleport.secreports.v1.ListAuditQueriesRequest.
 * Use `create(ListAuditQueriesRequestSchema)` to create a new message.
 */
export const ListAuditQueriesRequestSchema: GenMessage<ListAuditQueriesRequest> = /*@__PURE__*/
  messageDesc(file_teleport_secreports_v1_secreports_service, 15);

/**
 * ListAuditQueryResponse is audit query list response.
 *
 * @generated from message teleport.secreports.v1.ListReportsRequest
 */
export type ListReportsRequest = Message<"teleport.secreports.v1.ListReportsRequest"> & {
  /**
   * page_size is the number of results to return.
   *
   * @generated from field: int32 page_size = 1;
   */
  pageSize: number;

  /**
   * page_token is the next_token value returned from a previous List request if any.
   *
   * @generated from field: string page_token = 2;
   */
  pageToken: string;
};

/**
 * Describes the message teleport.secreports.v1.ListReportsRequest.
 * Use `create(ListReportsRequestSchema)` to create a new message.
 */
export const ListReportsRequestSchema: GenMessage<ListReportsRequest> = /*@__PURE__*/
  messageDesc(file_teleport_secreports_v1_secreports_service, 16);

/**
 * ListAuditQueriesResponse is audit query list response.
 *
 * @generated from message teleport.secreports.v1.ListAuditQueriesResponse
 */
export type ListAuditQueriesResponse = Message<"teleport.secreports.v1.ListAuditQueriesResponse"> & {
  /**
   * queries is a list of audit queries.
   *
   * @generated from field: repeated teleport.secreports.v1.AuditQuery queries = 1;
   */
  queries: AuditQuery[];

  /**
   * next_page_token is the next page token. If there are no more results, it will be empty.
   *
   * @generated from field: string next_page_token = 2;
   */
  nextPageToken: string;
};

/**
 * Describes the message teleport.secreports.v1.ListAuditQueriesResponse.
 * Use `create(ListAuditQueriesResponseSchema)` to create a new message.
 */
export const ListAuditQueriesResponseSchema: GenMessage<ListAuditQueriesResponse> = /*@__PURE__*/
  messageDesc(file_teleport_secreports_v1_secreports_service, 17);

/**
 * GetSchemaRequest is a request for GetSchema.
 *
 * @generated from message teleport.secreports.v1.GetSchemaRequest
 */
export type GetSchemaRequest = Message<"teleport.secreports.v1.GetSchemaRequest"> & {
};

/**
 * Describes the message teleport.secreports.v1.GetSchemaRequest.
 * Use `create(GetSchemaRequestSchema)` to create a new message.
 */
export const GetSchemaRequestSchema: GenMessage<GetSchemaRequest> = /*@__PURE__*/
  messageDesc(file_teleport_secreports_v1_secreports_service, 18);

/**
 * GetSchemaResponse is a response for GetSchema.
 *
 * @generated from message teleport.secreports.v1.GetSchemaResponse
 */
export type GetSchemaResponse = Message<"teleport.secreports.v1.GetSchemaResponse"> & {
  /**
   * views is the list of views.
   *
   * @generated from field: repeated teleport.secreports.v1.GetSchemaResponse.ViewDesc views = 1;
   */
  views: GetSchemaResponse_ViewDesc[];
};

/**
 * Describes the message teleport.secreports.v1.GetSchemaResponse.
 * Use `create(GetSchemaResponseSchema)` to create a new message.
 */
export const GetSchemaResponseSchema: GenMessage<GetSchemaResponse> = /*@__PURE__*/
  messageDesc(file_teleport_secreports_v1_secreports_service, 19);

/**
 * ViewDesc is a description of view.
 *
 * @generated from message teleport.secreports.v1.GetSchemaResponse.ViewDesc
 */
export type GetSchemaResponse_ViewDesc = Message<"teleport.secreports.v1.GetSchemaResponse.ViewDesc"> & {
  /**
   * name is the name of the view.
   *
   * @generated from field: string name = 1;
   */
  name: string;

  /**
   * desc is the description of the view.
   *
   * @generated from field: string desc = 2;
   */
  desc: string;

  /**
   * columns is the list of columns.
   *
   * @generated from field: repeated teleport.secreports.v1.GetSchemaResponse.ViewDesc.ColumnDesc columns = 3;
   */
  columns: GetSchemaResponse_ViewDesc_ColumnDesc[];
};

/**
 * Describes the message teleport.secreports.v1.GetSchemaResponse.ViewDesc.
 * Use `create(GetSchemaResponse_ViewDescSchema)` to create a new message.
 */
export const GetSchemaResponse_ViewDescSchema: GenMessage<GetSchemaResponse_ViewDesc> = /*@__PURE__*/
  messageDesc(file_teleport_secreports_v1_secreports_service, 19, 0);

/**
 * ColumnDesc is a description of column.
 *
 * @generated from message teleport.secreports.v1.GetSchemaResponse.ViewDesc.ColumnDesc
 */
export type GetSchemaResponse_ViewDesc_ColumnDesc = Message<"teleport.secreports.v1.GetSchemaResponse.ViewDesc.ColumnDesc"> & {
  /**
   * name is the name of the column.
   *
   * @generated from field: string name = 1;
   */
  name: string;

  /**
   * type is the type of the column.
   *
   * @generated from field: string type = 2;
   */
  type: string;

  /**
   * desc is the description of the column.
   *
   * @generated from field: string desc = 3;
   */
  desc: string;
};

/**
 * Describes the message teleport.secreports.v1.GetSchemaResponse.ViewDesc.ColumnDesc.
 * Use `create(GetSchemaResponse_ViewDesc_ColumnDescSchema)` to create a new message.
 */
export const GetSchemaResponse_ViewDesc_ColumnDescSchema: GenMessage<GetSchemaResponse_ViewDesc_ColumnDesc> = /*@__PURE__*/
  messageDesc(file_teleport_secreports_v1_secreports_service, 19, 0, 0);

/**
 * RunAuditQueryResponse is audit query run response.
 *
 * @generated from message teleport.secreports.v1.RunAuditQueryResponse
 */
export type RunAuditQueryResponse = Message<"teleport.secreports.v1.RunAuditQueryResponse"> & {
  /**
   * result_id is a unique id of the result.
   *
   * @generated from field: string result_id = 1;
   */
  resultId: string;
};

/**
 * Describes the message teleport.secreports.v1.RunAuditQueryResponse.
 * Use `create(RunAuditQueryResponseSchema)` to create a new message.
 */
export const RunAuditQueryResponseSchema: GenMessage<RunAuditQueryResponse> = /*@__PURE__*/
  messageDesc(file_teleport_secreports_v1_secreports_service, 20);

/**
 * ListReportResponse is security report list response.
 *
 * @generated from message teleport.secreports.v1.ListReportsResponse
 */
export type ListReportsResponse = Message<"teleport.secreports.v1.ListReportsResponse"> & {
  /**
   * reports is a list of security reports.
   *
   * @generated from field: repeated teleport.secreports.v1.Report reports = 1;
   */
  reports: Report[];

  /**
   * next_page_token is the next page token. If there are no more results, it will be empty.
   *
   * @generated from field: string next_page_token = 2;
   */
  nextPageToken: string;
};

/**
 * Describes the message teleport.secreports.v1.ListReportsResponse.
 * Use `create(ListReportsResponseSchema)` to create a new message.
 */
export const ListReportsResponseSchema: GenMessage<ListReportsResponse> = /*@__PURE__*/
  messageDesc(file_teleport_secreports_v1_secreports_service, 21);

/**
 * GetReportResultResponse is report result get response.
 *
 * @generated from message teleport.secreports.v1.GetReportResultResponse
 */
export type GetReportResultResponse = Message<"teleport.secreports.v1.GetReportResultResponse"> & {
  /**
   * result is a report execution result.
   *
   * @generated from field: teleport.secreports.v1.ReportResult result = 1;
   */
  result?: ReportResult;
};

/**
 * Describes the message teleport.secreports.v1.GetReportResultResponse.
 * Use `create(GetReportResultResponseSchema)` to create a new message.
 */
export const GetReportResultResponseSchema: GenMessage<GetReportResultResponse> = /*@__PURE__*/
  messageDesc(file_teleport_secreports_v1_secreports_service, 22);

/**
 * Report is the result of security report.
 *
 * @generated from message teleport.secreports.v1.ReportResult
 */
export type ReportResult = Message<"teleport.secreports.v1.ReportResult"> & {
  /**
   * name is a name of the security report.
   *
   * @generated from field: string name = 1;
   */
  name: string;

  /**
   * description is a description of the security report.
   *
   * @generated from field: string description = 2;
   */
  description: string;

  /**
   * audit_query_results is a list of audit query results.
   *
   * @generated from field: repeated teleport.secreports.v1.ReportResult.AuditQueryResult audit_query_results = 3;
   */
  auditQueryResults: ReportResult_AuditQueryResult[];

  /**
   * updated_at is a time when the security report was updated.
   *
   * @generated from field: string updated_at = 4;
   */
  updatedAt: string;

  /**
   * total_execution_time_in_millis is a time in milliseconds when the security report was executed.
   *
   * @generated from field: int64 total_execution_time_in_millis = 5;
   */
  totalExecutionTimeInMillis: bigint;

  /**
   * total_data_scanned_in_bytes is a number of bytes scanned.
   *
   * @generated from field: int64 total_data_scanned_in_bytes = 6;
   */
  totalDataScannedInBytes: bigint;
};

/**
 * Describes the message teleport.secreports.v1.ReportResult.
 * Use `create(ReportResultSchema)` to create a new message.
 */
export const ReportResultSchema: GenMessage<ReportResult> = /*@__PURE__*/
  messageDesc(file_teleport_secreports_v1_secreports_service, 23);

/**
 * AuditQueryResult is a result of audit query.
 *
 * @generated from message teleport.secreports.v1.ReportResult.AuditQueryResult
 */
export type ReportResult_AuditQueryResult = Message<"teleport.secreports.v1.ReportResult.AuditQueryResult"> & {
  /**
   * audit_query is the audit query spec containing information about audit query.
   *
   * @generated from field: teleport.secreports.v1.AuditQuerySpec audit_query = 1;
   */
  auditQuery?: AuditQuerySpec;

  /**
   * result is the result set.
   *
   * @generated from field: teleport.secreports.v1.QueryResultSet result = 2;
   */
  result?: QueryResultSet;

  /**
   * result_id is a unique id of the result.
   *
   * @generated from field: string result_id = 3;
   */
  resultId: string;

  /**
   * execution_time_in_millis is a time in milliseconds when the audit query was executed.
   *
   * @generated from field: int64 execution_time_in_millis = 4;
   */
  executionTimeInMillis: bigint;

  /**
   * data_scanned_in_bytes is a number of bytes scanned.
   *
   * @generated from field: int64 data_scanned_in_bytes = 5;
   */
  dataScannedInBytes: bigint;
};

/**
 * Describes the message teleport.secreports.v1.ReportResult.AuditQueryResult.
 * Use `create(ReportResult_AuditQueryResultSchema)` to create a new message.
 */
export const ReportResult_AuditQueryResultSchema: GenMessage<ReportResult_AuditQueryResult> = /*@__PURE__*/
  messageDesc(file_teleport_secreports_v1_secreports_service, 23, 0);

/**
 * Report is a security report.
 *
 * @generated from message teleport.secreports.v1.ReportSate
 */
export type ReportSate = Message<"teleport.secreports.v1.ReportSate"> & {
  /**
   * header is a resource header.
   *
   * @generated from field: teleport.header.v1.ResourceHeader header = 1;
   */
  header?: ResourceHeader;

  /**
   * state is a state of the security report.
   *
   * @generated from field: teleport.secreports.v1.ReportSate.State state = 2;
   */
  state: ReportSate_State;

  /**
   * updated_at is a time when the security report state was updated.
   *
   * @generated from field: string updated_at = 3;
   */
  updatedAt: string;
};

/**
 * Describes the message teleport.secreports.v1.ReportSate.
 * Use `create(ReportSateSchema)` to create a new message.
 */
export const ReportSateSchema: GenMessage<ReportSate> = /*@__PURE__*/
  messageDesc(file_teleport_secreports_v1_secreports_service, 24);

/**
 * name is a name of the security report.
 *
 * @generated from enum teleport.secreports.v1.ReportSate.State
 */
export enum ReportSate_State {
  /**
   * STATE_UNSPECIFIED is an unspecified state.
   *
   * @generated from enum value: STATE_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * STATE_ERROR is an error state.
   *
   * @generated from enum value: STATE_ERROR = 1;
   */
  ERROR = 1,

  /**
   * STATE_SUCCESS is a success state.
   *
   * @generated from enum value: STATE_SUCCESS = 2;
   */
  SUCCESS = 2,

  /**
   * STATE_RUNNING is a running state.
   *
   * @generated from enum value: STATE_RUNNING = 3;
   */
  RUNNING = 3,
}

/**
 * Describes the enum teleport.secreports.v1.ReportSate.State.
 */
export const ReportSate_StateSchema: GenEnum<ReportSate_State> = /*@__PURE__*/
  enumDesc(file_teleport_secreports_v1_secreports_service, 24, 0);

/**
 * SecReportsService is a service that manages security reports.
 *
 * @generated from service teleport.secreports.v1.SecReportsService
 */
export const SecReportsService: GenService<{
  /**
   * UpsertAuditQuery upsets an audit query.
   *
   * @generated from rpc teleport.secreports.v1.SecReportsService.UpsertAuditQuery
   */
  upsertAuditQuery: {
    methodKind: "unary";
    input: typeof UpsertAuditQueryRequestSchema;
    output: typeof EmptySchema;
  },
  /**
   * GetAuditQuery returns an audit query.
   *
   * @generated from rpc teleport.secreports.v1.SecReportsService.GetAuditQuery
   */
  getAuditQuery: {
    methodKind: "unary";
    input: typeof GetAuditQueryRequestSchema;
    output: typeof AuditQuerySchema;
  },
  /**
   * ListAuditQueries returns a paginated list of all Okta import rule resources.
   *
   * @generated from rpc teleport.secreports.v1.SecReportsService.ListAuditQueries
   */
  listAuditQueries: {
    methodKind: "unary";
    input: typeof ListAuditQueriesRequestSchema;
    output: typeof ListAuditQueriesResponseSchema;
  },
  /**
   * DeleteAuditQuery deletes an audit query.
   *
   * @generated from rpc teleport.secreports.v1.SecReportsService.DeleteAuditQuery
   */
  deleteAuditQuery: {
    methodKind: "unary";
    input: typeof DeleteAuditQueryRequestSchema;
    output: typeof EmptySchema;
  },
  /**
   * UpsertReport upsets a report.
   *
   * @generated from rpc teleport.secreports.v1.SecReportsService.UpsertReport
   */
  upsertReport: {
    methodKind: "unary";
    input: typeof UpsertReportRequestSchema;
    output: typeof EmptySchema;
  },
  /**
   * GetReport returns a report.
   *
   * @generated from rpc teleport.secreports.v1.SecReportsService.GetReport
   */
  getReport: {
    methodKind: "unary";
    input: typeof GetReportRequestSchema;
    output: typeof ReportSchema;
  },
  /**
   * ListReports returns a paginated list of all Okta import rule resources.
   *
   * @generated from rpc teleport.secreports.v1.SecReportsService.ListReports
   */
  listReports: {
    methodKind: "unary";
    input: typeof ListReportsRequestSchema;
    output: typeof ListReportsResponseSchema;
  },
  /**
   * DeleteReport deletes a security report.
   *
   * @generated from rpc teleport.secreports.v1.SecReportsService.DeleteReport
   */
  deleteReport: {
    methodKind: "unary";
    input: typeof DeleteReportRequestSchema;
    output: typeof EmptySchema;
  },
  /**
   * RunAuditQuery runs an audit query.
   *
   * @generated from rpc teleport.secreports.v1.SecReportsService.RunAuditQuery
   */
  runAuditQuery: {
    methodKind: "unary";
    input: typeof RunAuditQueryRequestSchema;
    output: typeof RunAuditQueryResponseSchema;
  },
  /**
   * GetAuditQueryResult returns an audit query result.
   *
   * @generated from rpc teleport.secreports.v1.SecReportsService.GetAuditQueryResult
   */
  getAuditQueryResult: {
    methodKind: "unary";
    input: typeof GetAuditQueryResultRequestSchema;
    output: typeof GetAuditQueryResultResponseSchema;
  },
  /**
   * RunReport runs a security report.
   *
   * @generated from rpc teleport.secreports.v1.SecReportsService.RunReport
   */
  runReport: {
    methodKind: "unary";
    input: typeof RunReportRequestSchema;
    output: typeof EmptySchema;
  },
  /**
   * GetReportResult returns a security report result.
   *
   * @generated from rpc teleport.secreports.v1.SecReportsService.GetReportResult
   */
  getReportResult: {
    methodKind: "unary";
    input: typeof GetReportResultRequestSchema;
    output: typeof GetReportResultResponseSchema;
  },
  /**
   * GetReportState returns a security report state.
   *
   * @generated from rpc teleport.secreports.v1.SecReportsService.GetReportState
   */
  getReportState: {
    methodKind: "unary";
    input: typeof GetReportStateRequestSchema;
    output: typeof ReportStateSchema;
  },
  /**
   * GetSchema returns a schema of audit query.
   *
   * @generated from rpc teleport.secreports.v1.SecReportsService.GetSchema
   */
  getSchema: {
    methodKind: "unary";
    input: typeof GetSchemaRequestSchema;
    output: typeof GetSchemaResponseSchema;
  },
}> = /*@__PURE__*/
  serviceDesc(file_teleport_secreports_v1_secreports_service, 0);

