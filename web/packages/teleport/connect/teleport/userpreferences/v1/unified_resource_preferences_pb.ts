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
// @generated from file teleport/userpreferences/v1/unified_resource_preferences.proto (package teleport.userpreferences.v1, syntax proto3)
/* eslint-disable */

import type { GenEnum, GenFile, GenMessage } from "@bufbuild/protobuf/codegenv1";
import { enumDesc, fileDesc, messageDesc } from "@bufbuild/protobuf/codegenv1";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file teleport/userpreferences/v1/unified_resource_preferences.proto.
 */
export const file_teleport_userpreferences_v1_unified_resource_preferences: GenFile = /*@__PURE__*/
  fileDesc("Cj50ZWxlcG9ydC91c2VycHJlZmVyZW5jZXMvdjEvdW5pZmllZF9yZXNvdXJjZV9wcmVmZXJlbmNlcy5wcm90bxIbdGVsZXBvcnQudXNlcnByZWZlcmVuY2VzLnYxIrACChpVbmlmaWVkUmVzb3VyY2VQcmVmZXJlbmNlcxI8CgtkZWZhdWx0X3RhYhgBIAEoDjInLnRlbGVwb3J0LnVzZXJwcmVmZXJlbmNlcy52MS5EZWZhdWx0VGFiEjgKCXZpZXdfbW9kZRgCIAEoDjIlLnRlbGVwb3J0LnVzZXJwcmVmZXJlbmNlcy52MS5WaWV3TW9kZRJFChBsYWJlbHNfdmlld19tb2RlGAMgASgOMisudGVsZXBvcnQudXNlcnByZWZlcmVuY2VzLnYxLkxhYmVsc1ZpZXdNb2RlElMKF2F2YWlsYWJsZV9yZXNvdXJjZV9tb2RlGAQgASgOMjIudGVsZXBvcnQudXNlcnByZWZlcmVuY2VzLnYxLkF2YWlsYWJsZVJlc291cmNlTW9kZSpWCgpEZWZhdWx0VGFiEhsKF0RFRkFVTFRfVEFCX1VOU1BFQ0lGSUVEEAASEwoPREVGQVVMVF9UQUJfQUxMEAESFgoSREVGQVVMVF9UQUJfUElOTkVEEAIqTQoIVmlld01vZGUSGQoVVklFV19NT0RFX1VOU1BFQ0lGSUVEEAASEgoOVklFV19NT0RFX0NBUkQQARISCg5WSUVXX01PREVfTElTVBACKnEKDkxhYmVsc1ZpZXdNb2RlEiAKHExBQkVMU19WSUVXX01PREVfVU5TUEVDSUZJRUQQABIdChlMQUJFTFNfVklFV19NT0RFX0VYUEFOREVEEAESHgoaTEFCRUxTX1ZJRVdfTU9ERV9DT0xMQVBTRUQQAirUAQoVQXZhaWxhYmxlUmVzb3VyY2VNb2RlEicKI0FWQUlMQUJMRV9SRVNPVVJDRV9NT0RFX1VOU1BFQ0lGSUVEEAASHwobQVZBSUxBQkxFX1JFU09VUkNFX01PREVfQUxMEAESJgoiQVZBSUxBQkxFX1JFU09VUkNFX01PREVfQUNDRVNTSUJMRRACEicKI0FWQUlMQUJMRV9SRVNPVVJDRV9NT0RFX1JFUVVFU1RBQkxFEAMSIAocQVZBSUxBQkxFX1JFU09VUkNFX01PREVfTk9ORRAEQllaV2dpdGh1Yi5jb20vZ3Jhdml0YXRpb25hbC90ZWxlcG9ydC9hcGkvZ2VuL3Byb3RvL2dvL3VzZXJwcmVmZXJlbmNlcy92MTt1c2VycHJlZmVyZW5jZXN2MWIGcHJvdG8z");

/**
 * UnifiedResourcePreferences are preferences used in the Unified Resource web UI
 *
 * @generated from message teleport.userpreferences.v1.UnifiedResourcePreferences
 */
export type UnifiedResourcePreferences = Message<"teleport.userpreferences.v1.UnifiedResourcePreferences"> & {
  /**
   * default_tab is the default tab selected in the unified resource web UI
   *
   * @generated from field: teleport.userpreferences.v1.DefaultTab default_tab = 1;
   */
  defaultTab: DefaultTab;

  /**
   * view_mode is the view mode selected in the unified resource Web UI
   *
   * @generated from field: teleport.userpreferences.v1.ViewMode view_mode = 2;
   */
  viewMode: ViewMode;

  /**
   * labels_view_mode is whether the labels for resources should all be collapsed or expanded in the unified resource Web UI list view.
   *
   * @generated from field: teleport.userpreferences.v1.LabelsViewMode labels_view_mode = 3;
   */
  labelsViewMode: LabelsViewMode;

  /**
   * available_resource_mode specifies which option in the availability filter menu the user has selected, if any
   *
   * @generated from field: teleport.userpreferences.v1.AvailableResourceMode available_resource_mode = 4;
   */
  availableResourceMode: AvailableResourceMode;
};

/**
 * Describes the message teleport.userpreferences.v1.UnifiedResourcePreferences.
 * Use `create(UnifiedResourcePreferencesSchema)` to create a new message.
 */
export const UnifiedResourcePreferencesSchema: GenMessage<UnifiedResourcePreferences> = /*@__PURE__*/
  messageDesc(file_teleport_userpreferences_v1_unified_resource_preferences, 0);

/**
 * DefaultTab is the default tab selected in the unified resource web UI
 *
 * @generated from enum teleport.userpreferences.v1.DefaultTab
 */
export enum DefaultTab {
  /**
   * @generated from enum value: DEFAULT_TAB_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * ALL is all resources
   *
   * @generated from enum value: DEFAULT_TAB_ALL = 1;
   */
  ALL = 1,

  /**
   * PINNED is only pinned resources
   *
   * @generated from enum value: DEFAULT_TAB_PINNED = 2;
   */
  PINNED = 2,
}

/**
 * Describes the enum teleport.userpreferences.v1.DefaultTab.
 */
export const DefaultTabSchema: GenEnum<DefaultTab> = /*@__PURE__*/
  enumDesc(file_teleport_userpreferences_v1_unified_resource_preferences, 0);

/**
 * ViewMode is the view mode selected in the unified resource Web UI
 *
 * @generated from enum teleport.userpreferences.v1.ViewMode
 */
export enum ViewMode {
  /**
   * @generated from enum value: VIEW_MODE_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * CARD is the card view
   *
   * @generated from enum value: VIEW_MODE_CARD = 1;
   */
  CARD = 1,

  /**
   * LIST is the list view
   *
   * @generated from enum value: VIEW_MODE_LIST = 2;
   */
  LIST = 2,
}

/**
 * Describes the enum teleport.userpreferences.v1.ViewMode.
 */
export const ViewModeSchema: GenEnum<ViewMode> = /*@__PURE__*/
  enumDesc(file_teleport_userpreferences_v1_unified_resource_preferences, 1);

/**
 * * LabelsViewMode is whether the labels for resources should all be collapsed or expanded. This only applies to the list view. 
 *
 * @generated from enum teleport.userpreferences.v1.LabelsViewMode
 */
export enum LabelsViewMode {
  /**
   * @generated from enum value: LABELS_VIEW_MODE_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * EXPANDED is the expanded state which shows all labels for every resource.
   *
   * @generated from enum value: LABELS_VIEW_MODE_EXPANDED = 1;
   */
  EXPANDED = 1,

  /**
   * COLLAPSED is the collapsed state which hides all labels for every resource.
   *
   * @generated from enum value: LABELS_VIEW_MODE_COLLAPSED = 2;
   */
  COLLAPSED = 2,
}

/**
 * Describes the enum teleport.userpreferences.v1.LabelsViewMode.
 */
export const LabelsViewModeSchema: GenEnum<LabelsViewMode> = /*@__PURE__*/
  enumDesc(file_teleport_userpreferences_v1_unified_resource_preferences, 2);

/**
 * * AvailableResourceMode specifies which option in the availability filter menu the user has selected, if any 
 *
 * @generated from enum teleport.userpreferences.v1.AvailableResourceMode
 */
export enum AvailableResourceMode {
  /**
   * @generated from enum value: AVAILABLE_RESOURCE_MODE_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: AVAILABLE_RESOURCE_MODE_ALL = 1;
   */
  ALL = 1,

  /**
   * @generated from enum value: AVAILABLE_RESOURCE_MODE_ACCESSIBLE = 2;
   */
  ACCESSIBLE = 2,

  /**
   * @generated from enum value: AVAILABLE_RESOURCE_MODE_REQUESTABLE = 3;
   */
  REQUESTABLE = 3,

  /**
   * @generated from enum value: AVAILABLE_RESOURCE_MODE_NONE = 4;
   */
  NONE = 4,
}

/**
 * Describes the enum teleport.userpreferences.v1.AvailableResourceMode.
 */
export const AvailableResourceModeSchema: GenEnum<AvailableResourceMode> = /*@__PURE__*/
  enumDesc(file_teleport_userpreferences_v1_unified_resource_preferences, 3);

