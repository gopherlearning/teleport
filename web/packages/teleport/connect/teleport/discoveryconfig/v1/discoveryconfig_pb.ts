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
// @generated from file teleport/discoveryconfig/v1/discoveryconfig.proto (package teleport.discoveryconfig.v1, syntax proto3)
/* eslint-disable */

import type { GenEnum, GenFile, GenMessage } from "@bufbuild/protobuf/codegenv1";
import { enumDesc, fileDesc, messageDesc } from "@bufbuild/protobuf/codegenv1";
import type { Timestamp } from "@bufbuild/protobuf/wkt";
import { file_google_protobuf_timestamp } from "@bufbuild/protobuf/wkt";
import type { ResourceHeader } from "../../header/v1/resourceheader_pb";
import { file_teleport_header_v1_resourceheader } from "../../header/v1/resourceheader_pb";
import type { AccessGraphSync, AWSMatcher, AzureMatcher, GCPMatcher, KubernetesMatcher } from "../../legacy/types/types_pb";
import { file_teleport_legacy_types_types } from "../../legacy/types/types_pb";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file teleport/discoveryconfig/v1/discoveryconfig.proto.
 */
export const file_teleport_discoveryconfig_v1_discoveryconfig: GenFile = /*@__PURE__*/
  fileDesc("CjF0ZWxlcG9ydC9kaXNjb3Zlcnljb25maWcvdjEvZGlzY292ZXJ5Y29uZmlnLnByb3RvEht0ZWxlcG9ydC5kaXNjb3Zlcnljb25maWcudjEiyQEKD0Rpc2NvdmVyeUNvbmZpZxIyCgZoZWFkZXIYASABKAsyIi50ZWxlcG9ydC5oZWFkZXIudjEuUmVzb3VyY2VIZWFkZXISPgoEc3BlYxgCIAEoCzIwLnRlbGVwb3J0LmRpc2NvdmVyeWNvbmZpZy52MS5EaXNjb3ZlcnlDb25maWdTcGVjEkIKBnN0YXR1cxgDIAEoCzIyLnRlbGVwb3J0LmRpc2NvdmVyeWNvbmZpZy52MS5EaXNjb3ZlcnlDb25maWdTdGF0dXMi6AEKE0Rpc2NvdmVyeUNvbmZpZ1NwZWMSFwoPZGlzY292ZXJ5X2dyb3VwGAEgASgJEh4KA2F3cxgCIAMoCzIRLnR5cGVzLkFXU01hdGNoZXISIgoFYXp1cmUYAyADKAsyEy50eXBlcy5BenVyZU1hdGNoZXISHgoDZ2NwGAQgAygLMhEudHlwZXMuR0NQTWF0Y2hlchImCgRrdWJlGAUgAygLMhgudHlwZXMuS3ViZXJuZXRlc01hdGNoZXISLAoMYWNjZXNzX2dyYXBoGAYgASgLMhYudHlwZXMuQWNjZXNzR3JhcGhTeW5jIoMEChVEaXNjb3ZlcnlDb25maWdTdGF0dXMSQAoFc3RhdGUYASABKA4yMS50ZWxlcG9ydC5kaXNjb3Zlcnljb25maWcudjEuRGlzY292ZXJ5Q29uZmlnU3RhdGUSGgoNZXJyb3JfbWVzc2FnZRgCIAEoCUgAiAEBEhwKFGRpc2NvdmVyZWRfcmVzb3VyY2VzGAMgASgEEjIKDmxhc3Rfc3luY190aW1lGAQgASgLMhouZ29vZ2xlLnByb3RvYnVmLlRpbWVzdGFtcBKAAQogaW50ZWdyYXRpb25fZGlzY292ZXJlZF9yZXNvdXJjZXMYBiADKAsyVi50ZWxlcG9ydC5kaXNjb3Zlcnljb25maWcudjEuRGlzY292ZXJ5Q29uZmlnU3RhdHVzLkludGVncmF0aW9uRGlzY292ZXJlZFJlc291cmNlc0VudHJ5GoABCiNJbnRlZ3JhdGlvbkRpc2NvdmVyZWRSZXNvdXJjZXNFbnRyeRILCgNrZXkYASABKAkSSAoFdmFsdWUYAiABKAsyOS50ZWxlcG9ydC5kaXNjb3Zlcnljb25maWcudjEuSW50ZWdyYXRpb25EaXNjb3ZlcmVkU3VtbWFyeToCOAFCEAoOX2Vycm9yX21lc3NhZ2VKBAgFEAZSHGF3c19lYzJfaW5zdGFuY2VzX2Rpc2NvdmVyZWQiaAocSW50ZWdyYXRpb25EaXNjb3ZlcmVkU3VtbWFyeRJICgdhd3NfZWMyGAEgASgLMjcudGVsZXBvcnQuZGlzY292ZXJ5Y29uZmlnLnYxLlJlc291cmNlc0Rpc2NvdmVyZWRTdW1tYXJ5Ik0KGlJlc291cmNlc0Rpc2NvdmVyZWRTdW1tYXJ5Eg0KBWZvdW5kGAEgASgEEhAKCGVucm9sbGVkGAIgASgEEg4KBmZhaWxlZBgDIAEoBCqoAQoURGlzY292ZXJ5Q29uZmlnU3RhdGUSJgoiRElTQ09WRVJZX0NPTkZJR19TVEFURV9VTlNQRUNJRklFRBAAEiIKHkRJU0NPVkVSWV9DT05GSUdfU1RBVEVfUlVOTklORxABEiAKHERJU0NPVkVSWV9DT05GSUdfU1RBVEVfRVJST1IQAhIiCh5ESVNDT1ZFUllfQ09ORklHX1NUQVRFX1NZTkNJTkcQA0JiWmBnaXRodWIuY29tL2dyYXZpdGF0aW9uYWwvdGVsZXBvcnQvYXBpL2dlbi9wcm90by9nby90ZWxlcG9ydC9kaXNjb3Zlcnljb25maWcvdjE7ZGlzY292ZXJ5Y29uZmlndjFiBnByb3RvMw", [file_google_protobuf_timestamp, file_teleport_header_v1_resourceheader, file_teleport_legacy_types_types]);

/**
 * DiscoveryConfig is a resource that has Discovery Resource Matchers and a Discovery Group.
 *
 * Teleport Discovery Services will load the dynamic DiscoveryConfigs whose Discovery Group matches the discovery_group defined in their configuration.
 *
 * @generated from message teleport.discoveryconfig.v1.DiscoveryConfig
 */
export type DiscoveryConfig = Message<"teleport.discoveryconfig.v1.DiscoveryConfig"> & {
  /**
   * Header is the resource header.
   *
   * @generated from field: teleport.header.v1.ResourceHeader header = 1;
   */
  header?: ResourceHeader;

  /**
   * Spec is an DiscoveryConfig specification.
   *
   * @generated from field: teleport.discoveryconfig.v1.DiscoveryConfigSpec spec = 2;
   */
  spec?: DiscoveryConfigSpec;

  /**
   * Status is the resource Status
   *
   * @generated from field: teleport.discoveryconfig.v1.DiscoveryConfigStatus status = 3;
   */
  status?: DiscoveryConfigStatus;
};

/**
 * Describes the message teleport.discoveryconfig.v1.DiscoveryConfig.
 * Use `create(DiscoveryConfigSchema)` to create a new message.
 */
export const DiscoveryConfigSchema: GenMessage<DiscoveryConfig> = /*@__PURE__*/
  messageDesc(file_teleport_discoveryconfig_v1_discoveryconfig, 0);

/**
 * DiscoveryConfigSpec contains properties required to create matchers to be used by discovery_service.
 * Those matchers are used by discovery_service to watch for cloud resources and create them in Teleport.
 *
 * @generated from message teleport.discoveryconfig.v1.DiscoveryConfigSpec
 */
export type DiscoveryConfigSpec = Message<"teleport.discoveryconfig.v1.DiscoveryConfigSpec"> & {
  /**
   * DiscoveryGroup is used by discovery_service to add extra matchers.
   * All the discovery_services that have the same discovery_group, will load the matchers of this resource.
   *
   * @generated from field: string discovery_group = 1;
   */
  discoveryGroup: string;

  /**
   * AWS is a list of AWS Matchers.
   *
   * @generated from field: repeated types.AWSMatcher aws = 2;
   */
  aws: AWSMatcher[];

  /**
   * Azure is a list of Azure Matchers.
   *
   * @generated from field: repeated types.AzureMatcher azure = 3;
   */
  azure: AzureMatcher[];

  /**
   * GCP is a list of GCP Matchers.
   *
   * @generated from field: repeated types.GCPMatcher gcp = 4;
   */
  gcp: GCPMatcher[];

  /**
   * Kube is a list of Kubernetes Matchers.
   *
   * @generated from field: repeated types.KubernetesMatcher kube = 5;
   */
  kube: KubernetesMatcher[];

  /**
   * AccessGraph is the configurations for syncing Cloud accounts into Access Graph.
   *
   * @generated from field: types.AccessGraphSync access_graph = 6;
   */
  accessGraph?: AccessGraphSync;
};

/**
 * Describes the message teleport.discoveryconfig.v1.DiscoveryConfigSpec.
 * Use `create(DiscoveryConfigSpecSchema)` to create a new message.
 */
export const DiscoveryConfigSpecSchema: GenMessage<DiscoveryConfigSpec> = /*@__PURE__*/
  messageDesc(file_teleport_discoveryconfig_v1_discoveryconfig, 1);

/**
 * DiscoveryConfigStatus holds dynamic information about the discovery configuration
 * running status such as errors, state and count of the resources.
 *
 * @generated from message teleport.discoveryconfig.v1.DiscoveryConfigStatus
 */
export type DiscoveryConfigStatus = Message<"teleport.discoveryconfig.v1.DiscoveryConfigStatus"> & {
  /**
   * State reports the Discovery config state.
   *
   * @generated from field: teleport.discoveryconfig.v1.DiscoveryConfigState state = 1;
   */
  state: DiscoveryConfigState;

  /**
   * error_message holds the error message when state is DISCOVERY_CONFIG_STATE_ERROR.
   *
   * @generated from field: optional string error_message = 2;
   */
  errorMessage?: string;

  /**
   * discovered_resources holds the count of the discovered resources in the previous iteration.
   *
   * @generated from field: uint64 discovered_resources = 3;
   */
  discoveredResources: bigint;

  /**
   * last_sync_time is the timestamp when the Discovery Config was last sync.
   *
   * @generated from field: google.protobuf.Timestamp last_sync_time = 4;
   */
  lastSyncTime?: Timestamp;

  /**
   * IntegrationDiscoveredResources maps an integration to discovered resources summary.
   *
   * @generated from field: map<string, teleport.discoveryconfig.v1.IntegrationDiscoveredSummary> integration_discovered_resources = 6;
   */
  integrationDiscoveredResources: { [key: string]: IntegrationDiscoveredSummary };
};

/**
 * Describes the message teleport.discoveryconfig.v1.DiscoveryConfigStatus.
 * Use `create(DiscoveryConfigStatusSchema)` to create a new message.
 */
export const DiscoveryConfigStatusSchema: GenMessage<DiscoveryConfigStatus> = /*@__PURE__*/
  messageDesc(file_teleport_discoveryconfig_v1_discoveryconfig, 2);

/**
 * IntegrationDiscoveredSummary contains the a summary for each resource type that was discovered.
 *
 * @generated from message teleport.discoveryconfig.v1.IntegrationDiscoveredSummary
 */
export type IntegrationDiscoveredSummary = Message<"teleport.discoveryconfig.v1.IntegrationDiscoveredSummary"> & {
  /**
   * AWSEC2 contains the summary for the AWS EC2 discovered instances.
   *
   * @generated from field: teleport.discoveryconfig.v1.ResourcesDiscoveredSummary aws_ec2 = 1;
   */
  awsEc2?: ResourcesDiscoveredSummary;
};

/**
 * Describes the message teleport.discoveryconfig.v1.IntegrationDiscoveredSummary.
 * Use `create(IntegrationDiscoveredSummarySchema)` to create a new message.
 */
export const IntegrationDiscoveredSummarySchema: GenMessage<IntegrationDiscoveredSummary> = /*@__PURE__*/
  messageDesc(file_teleport_discoveryconfig_v1_discoveryconfig, 3);

/**
 * ResourcesDiscoveredSummary represents the AWS resources that were discovered.
 *
 * @generated from message teleport.discoveryconfig.v1.ResourcesDiscoveredSummary
 */
export type ResourcesDiscoveredSummary = Message<"teleport.discoveryconfig.v1.ResourcesDiscoveredSummary"> & {
  /**
   * Found holds the count of resources found.
   * After a resource is found, it starts the sync process and ends in either an enrolled or a failed resource.
   *
   * @generated from field: uint64 found = 1;
   */
  found: bigint;

  /**
   * Enrolled holds the count of the resources that were successfully enrolled.
   *
   * @generated from field: uint64 enrolled = 2;
   */
  enrolled: bigint;

  /**
   * Failed holds the count of the resources that failed to enroll.
   *
   * @generated from field: uint64 failed = 3;
   */
  failed: bigint;
};

/**
 * Describes the message teleport.discoveryconfig.v1.ResourcesDiscoveredSummary.
 * Use `create(ResourcesDiscoveredSummarySchema)` to create a new message.
 */
export const ResourcesDiscoveredSummarySchema: GenMessage<ResourcesDiscoveredSummary> = /*@__PURE__*/
  messageDesc(file_teleport_discoveryconfig_v1_discoveryconfig, 4);

/**
 * DiscoveryConfigState is the state of the discovery config resource.
 *
 * @generated from enum teleport.discoveryconfig.v1.DiscoveryConfigState
 */
export enum DiscoveryConfigState {
  /**
   * @generated from enum value: DISCOVERY_CONFIG_STATE_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * DISCOVERY_CONFIG_STATE_RUNNING is used when the operation doesn't report
   * incidents.
   *
   * @generated from enum value: DISCOVERY_CONFIG_STATE_RUNNING = 1;
   */
  RUNNING = 1,

  /**
   * DISCOVERY_CONFIG_STATE_ERROR is used when the operation reports
   * incidents.
   *
   * @generated from enum value: DISCOVERY_CONFIG_STATE_ERROR = 2;
   */
  ERROR = 2,

  /**
   * DISCOVERY_CONFIG_STATE_SYNCING is used when the discovery process has started but didn't finished yet.
   *
   * @generated from enum value: DISCOVERY_CONFIG_STATE_SYNCING = 3;
   */
  SYNCING = 3,
}

/**
 * Describes the enum teleport.discoveryconfig.v1.DiscoveryConfigState.
 */
export const DiscoveryConfigStateSchema: GenEnum<DiscoveryConfigState> = /*@__PURE__*/
  enumDesc(file_teleport_discoveryconfig_v1_discoveryconfig, 0);

