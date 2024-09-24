//
// Teleport
// Copyright (C) 2024  Gravitational, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

// @generated by protoc-gen-es v2.1.0 with parameter "target=ts"
// @generated from file accessgraph/v1alpha/gitlab.proto (package accessgraph.v1alpha, syntax proto3)
/* eslint-disable */

import type { GenEnum, GenFile, GenMessage } from "@bufbuild/protobuf/codegenv1";
import { enumDesc, fileDesc, messageDesc } from "@bufbuild/protobuf/codegenv1";
import type { Timestamp } from "@bufbuild/protobuf/wkt";
import { file_google_protobuf_timestamp } from "@bufbuild/protobuf/wkt";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file accessgraph/v1alpha/gitlab.proto.
 */
export const file_accessgraph_v1alpha_gitlab: GenFile = /*@__PURE__*/
  fileDesc("CiBhY2Nlc3NncmFwaC92MWFscGhhL2dpdGxhYi5wcm90bxITYWNjZXNzZ3JhcGgudjFhbHBoYSIVChNHaXRsYWJTeW5jT3BlcmF0aW9uIkwKEkdpdGxhYlJlc291cmNlTGlzdBI2CglyZXNvdXJjZXMYASADKAsyIy5hY2Nlc3NncmFwaC52MWFscGhhLkdpdGxhYlJlc291cmNlIrsCCg5HaXRsYWJSZXNvdXJjZRIxCgVncm91cBgBIAEoCzIgLmFjY2Vzc2dyYXBoLnYxYWxwaGEuR2l0bGFiR3JvdXBIABI1Cgdwcm9qZWN0GAIgASgLMiIuYWNjZXNzZ3JhcGgudjFhbHBoYS5HaXRsYWJQcm9qZWN0SAASQgoOcHJvamVjdF9tZW1iZXIYAyABKAsyKC5hY2Nlc3NncmFwaC52MWFscGhhLkdpdGxhYlByb2plY3RNZW1iZXJIABI+Cgxncm91cF9tZW1iZXIYBCABKAsyJi5hY2Nlc3NncmFwaC52MWFscGhhLkdpdGxhYkdyb3VwTWVtYmVySAASLwoEdXNlchgFIAEoCzIfLmFjY2Vzc2dyYXBoLnYxYWxwaGEuR2l0bGFiVXNlckgAQgoKCHJlc291cmNlIlEKC0dpdGxhYkdyb3VwEgwKBG5hbWUYASABKAkSDAoEcGF0aBgCIAEoCRIRCglmdWxsX25hbWUYAyABKAkSEwoLZGVzY3JpcHRpb24YBCABKAkiQAoNR2l0bGFiUHJvamVjdBIMCgRuYW1lGAEgASgJEgwKBHBhdGgYAiABKAkSEwoLZGVzY3JpcHRpb24YAyABKAkimAEKE0dpdGxhYlByb2plY3RNZW1iZXISEAoIdXNlcm5hbWUYASABKAkSOgoMYWNjZXNzX2xldmVsGAIgASgOMiQuYWNjZXNzZ3JhcGgudjFhbHBoYS5BY2Nlc3NMZXZlbFR5cGUSMwoHcHJvamVjdBgDIAEoCzIiLmFjY2Vzc2dyYXBoLnYxYWxwaGEuR2l0bGFiUHJvamVjdCKSAQoRR2l0bGFiR3JvdXBNZW1iZXISEAoIdXNlcm5hbWUYASABKAkSOgoMYWNjZXNzX2xldmVsGAIgASgOMiQuYWNjZXNzZ3JhcGgudjFhbHBoYS5BY2Nlc3NMZXZlbFR5cGUSLwoFZ3JvdXAYAyABKAsyIC5hY2Nlc3NncmFwaC52MWFscGhhLkdpdGxhYkdyb3VwIqcCCgpHaXRsYWJVc2VyEhAKCHVzZXJuYW1lGAEgASgJEg0KBWVtYWlsGAIgASgJEgwKBG5hbWUYAyABKAkSEAoIaXNfYWRtaW4YBCABKAgSFAoMb3JnYW5pemF0aW9uGAUgASgJEjMKD2xhc3Rfc2lnbl9pbl9hdBgGIAEoCzIaLmdvb2dsZS5wcm90b2J1Zi5UaW1lc3RhbXASGAoQY2FuX2NyZWF0ZV9ncm91cBgHIAEoCBIaChJjYW5fY3JlYXRlX3Byb2plY3QYCCABKAgSGgoSdHdvX2ZhY3Rvcl9lbmFibGVkGAkgASgIEjsKCmlkZW50aXRpZXMYCiADKAsyJy5hY2Nlc3NncmFwaC52MWFscGhhLkdpdGxhYlVzZXJJZGVudGl0eSI6ChJHaXRsYWJVc2VySWRlbnRpdHkSEAoIcHJvdmlkZXIYASABKAkSEgoKZXh0ZXJuX3VpZBgCIAEoCSqWAgoPQWNjZXNzTGV2ZWxUeXBlEiEKHUFDQ0VTU19MRVZFTF9UWVBFX1VOU1BFQ0lGSUVEEAASJAogQUNDRVNTX0xFVkVMX1RZUEVfTk9fUEVSTUlTU0lPTlMQARIdChlBQ0NFU1NfTEVWRUxfVFlQRV9NSU5JTUFMEAISGwoXQUNDRVNTX0xFVkVMX1RZUEVfR1VFU1QQAxIeChpBQ0NFU1NfTEVWRUxfVFlQRV9SRVBPUlRFUhAEEh8KG0FDQ0VTU19MRVZFTF9UWVBFX0RFVkVMT1BFUhAFEiAKHEFDQ0VTU19MRVZFTF9UWVBFX01BSU5UQUlORVIQBhIbChdBQ0NFU1NfTEVWRUxfVFlQRV9PV05FUhAHQldaVWdpdGh1Yi5jb20vZ3Jhdml0YXRpb25hbC90ZWxlcG9ydC9nZW4vcHJvdG8vZ28vYWNjZXNzZ3JhcGgvdjFhbHBoYTthY2Nlc3NncmFwaHYxYWxwaGFiBnByb3RvMw", [file_google_protobuf_timestamp]);

/**
 * GitlabSyncOperation is a request to sync Gitlab resources
 *
 * @generated from message accessgraph.v1alpha.GitlabSyncOperation
 */
export type GitlabSyncOperation = Message<"accessgraph.v1alpha.GitlabSyncOperation"> & {
};

/**
 * Describes the message accessgraph.v1alpha.GitlabSyncOperation.
 * Use `create(GitlabSyncOperationSchema)` to create a new message.
 */
export const GitlabSyncOperationSchema: GenMessage<GitlabSyncOperation> = /*@__PURE__*/
  messageDesc(file_accessgraph_v1alpha_gitlab, 0);

/**
 * GitlabResourceList is a request that contains resources to be sync.
 *
 * @generated from message accessgraph.v1alpha.GitlabResourceList
 */
export type GitlabResourceList = Message<"accessgraph.v1alpha.GitlabResourceList"> & {
  /**
   * resources is a list of gitlab resources to sync.
   *
   * @generated from field: repeated accessgraph.v1alpha.GitlabResource resources = 1;
   */
  resources: GitlabResource[];
};

/**
 * Describes the message accessgraph.v1alpha.GitlabResourceList.
 * Use `create(GitlabResourceListSchema)` to create a new message.
 */
export const GitlabResourceListSchema: GenMessage<GitlabResourceList> = /*@__PURE__*/
  messageDesc(file_accessgraph_v1alpha_gitlab, 1);

/**
 * GitlabResource represents a Gitlab resource
 *
 * @generated from message accessgraph.v1alpha.GitlabResource
 */
export type GitlabResource = Message<"accessgraph.v1alpha.GitlabResource"> & {
  /**
   * @generated from oneof accessgraph.v1alpha.GitlabResource.resource
   */
  resource: {
    /**
     * group represents a gitlab group or subgroup in an organization.
     *
     * @generated from field: accessgraph.v1alpha.GitlabGroup group = 1;
     */
    value: GitlabGroup;
    case: "group";
  } | {
    /**
     * project represents a gitlab repository.
     *
     * @generated from field: accessgraph.v1alpha.GitlabProject project = 2;
     */
    value: GitlabProject;
    case: "project";
  } | {
    /**
     * project_member represents a user with certain access levels to a project.
     *
     * @generated from field: accessgraph.v1alpha.GitlabProjectMember project_member = 3;
     */
    value: GitlabProjectMember;
    case: "projectMember";
  } | {
    /**
     * group_member represents a user with certain access levels to a group and all subgroups/projects within.
     *
     * @generated from field: accessgraph.v1alpha.GitlabGroupMember group_member = 4;
     */
    value: GitlabGroupMember;
    case: "groupMember";
  } | {
    /**
     * user represents a gitlab user.
     *
     * @generated from field: accessgraph.v1alpha.GitlabUser user = 5;
     */
    value: GitlabUser;
    case: "user";
  } | { case: undefined; value?: undefined };
};

/**
 * Describes the message accessgraph.v1alpha.GitlabResource.
 * Use `create(GitlabResourceSchema)` to create a new message.
 */
export const GitlabResourceSchema: GenMessage<GitlabResource> = /*@__PURE__*/
  messageDesc(file_accessgraph_v1alpha_gitlab, 2);

/**
 * GitlabGroup represents a Gitlab group
 *
 * @generated from message accessgraph.v1alpha.GitlabGroup
 */
export type GitlabGroup = Message<"accessgraph.v1alpha.GitlabGroup"> & {
  /**
   * name is the group name.
   *
   * @generated from field: string name = 1;
   */
  name: string;

  /**
   * path is the universal identifier for the group location.
   *
   * @generated from field: string path = 2;
   */
  path: string;

  /**
   * full_name is the group full name.
   *
   * @generated from field: string full_name = 3;
   */
  fullName: string;

  /**
   * description is the group description.
   *
   * @generated from field: string description = 4;
   */
  description: string;
};

/**
 * Describes the message accessgraph.v1alpha.GitlabGroup.
 * Use `create(GitlabGroupSchema)` to create a new message.
 */
export const GitlabGroupSchema: GenMessage<GitlabGroup> = /*@__PURE__*/
  messageDesc(file_accessgraph_v1alpha_gitlab, 3);

/**
 * GitlabProject represents a Gitlab project
 *
 * @generated from message accessgraph.v1alpha.GitlabProject
 */
export type GitlabProject = Message<"accessgraph.v1alpha.GitlabProject"> & {
  /**
   * name is the repository name.
   *
   * @generated from field: string name = 1;
   */
  name: string;

  /**
   * path is the universal identifier for the project location.
   *
   * @generated from field: string path = 2;
   */
  path: string;

  /**
   * description is the project description.
   *
   * @generated from field: string description = 3;
   */
  description: string;
};

/**
 * Describes the message accessgraph.v1alpha.GitlabProject.
 * Use `create(GitlabProjectSchema)` to create a new message.
 */
export const GitlabProjectSchema: GenMessage<GitlabProject> = /*@__PURE__*/
  messageDesc(file_accessgraph_v1alpha_gitlab, 4);

/**
 * GitlabProjectMember represents a Gitlab project member
 *
 * @generated from message accessgraph.v1alpha.GitlabProjectMember
 */
export type GitlabProjectMember = Message<"accessgraph.v1alpha.GitlabProjectMember"> & {
  /**
   * username is the username of the user.
   *
   * @generated from field: string username = 1;
   */
  username: string;

  /**
   * access_level defines the permissions the user has to the project.
   *
   * @generated from field: accessgraph.v1alpha.AccessLevelType access_level = 2;
   */
  accessLevel: AccessLevelType;

  /**
   * project identifies the project that the user is member of.
   *
   * @generated from field: accessgraph.v1alpha.GitlabProject project = 3;
   */
  project?: GitlabProject;
};

/**
 * Describes the message accessgraph.v1alpha.GitlabProjectMember.
 * Use `create(GitlabProjectMemberSchema)` to create a new message.
 */
export const GitlabProjectMemberSchema: GenMessage<GitlabProjectMember> = /*@__PURE__*/
  messageDesc(file_accessgraph_v1alpha_gitlab, 5);

/**
 * GitlabGroupMember represents a Gitlab group member
 *
 * @generated from message accessgraph.v1alpha.GitlabGroupMember
 */
export type GitlabGroupMember = Message<"accessgraph.v1alpha.GitlabGroupMember"> & {
  /**
   * username is the username of the user.
   *
   * @generated from field: string username = 1;
   */
  username: string;

  /**
   * access_level defines the permissions the user has to the group and all projects within.
   *
   * @generated from field: accessgraph.v1alpha.AccessLevelType access_level = 2;
   */
  accessLevel: AccessLevelType;

  /**
   * project identifies the project that the user is member of.
   *
   * @generated from field: accessgraph.v1alpha.GitlabGroup group = 3;
   */
  group?: GitlabGroup;
};

/**
 * Describes the message accessgraph.v1alpha.GitlabGroupMember.
 * Use `create(GitlabGroupMemberSchema)` to create a new message.
 */
export const GitlabGroupMemberSchema: GenMessage<GitlabGroupMember> = /*@__PURE__*/
  messageDesc(file_accessgraph_v1alpha_gitlab, 6);

/**
 * GitlabGroupMember represents a Gitlab user.
 *
 * @generated from message accessgraph.v1alpha.GitlabUser
 */
export type GitlabUser = Message<"accessgraph.v1alpha.GitlabUser"> & {
  /**
   * username is the username of the user.
   *
   * @generated from field: string username = 1;
   */
  username: string;

  /**
   * email is the user's email.
   *
   * @generated from field: string email = 2;
   */
  email: string;

  /**
   * name is the user's name.
   *
   * @generated from field: string name = 3;
   */
  name: string;

  /**
   * is_admin indicates if a user is admin.
   *
   * @generated from field: bool is_admin = 4;
   */
  isAdmin: boolean;

  /**
   * organization is the user's organization.
   *
   * @generated from field: string organization = 5;
   */
  organization: string;

  /**
   * last_sign_in_at identifies the last sign in date.
   *
   * @generated from field: google.protobuf.Timestamp last_sign_in_at = 6;
   */
  lastSignInAt?: Timestamp;

  /**
   * can_create_group identifies if the user can create groups.
   *
   * @generated from field: bool can_create_group = 7;
   */
  canCreateGroup: boolean;

  /**
   * can_create_project identifies if the user can create projects.
   *
   * @generated from field: bool can_create_project = 8;
   */
  canCreateProject: boolean;

  /**
   * two_factor_enabled identifies if the user has two factor authentication enabled.
   *
   * @generated from field: bool two_factor_enabled = 9;
   */
  twoFactorEnabled: boolean;

  /**
   * identities represents the identity source for the user.
   *
   * @generated from field: repeated accessgraph.v1alpha.GitlabUserIdentity identities = 10;
   */
  identities: GitlabUserIdentity[];
};

/**
 * Describes the message accessgraph.v1alpha.GitlabUser.
 * Use `create(GitlabUserSchema)` to create a new message.
 */
export const GitlabUserSchema: GenMessage<GitlabUser> = /*@__PURE__*/
  messageDesc(file_accessgraph_v1alpha_gitlab, 7);

/**
 * GitlabUserIdentity identifies the external identity of the user.
 *
 * @generated from message accessgraph.v1alpha.GitlabUserIdentity
 */
export type GitlabUserIdentity = Message<"accessgraph.v1alpha.GitlabUserIdentity"> & {
  /**
   * provider identifies the identity provider.
   *
   * @generated from field: string provider = 1;
   */
  provider: string;

  /**
   * extern_uid identifies the external uid of the identity.
   *
   * @generated from field: string extern_uid = 2;
   */
  externUid: string;
};

/**
 * Describes the message accessgraph.v1alpha.GitlabUserIdentity.
 * Use `create(GitlabUserIdentitySchema)` to create a new message.
 */
export const GitlabUserIdentitySchema: GenMessage<GitlabUserIdentity> = /*@__PURE__*/
  messageDesc(file_accessgraph_v1alpha_gitlab, 8);

/**
 * AccessLevelType defines the access level a user has
 * to a project or Gitlab group.
 *
 * @generated from enum accessgraph.v1alpha.AccessLevelType
 */
export enum AccessLevelType {
  /**
   * ACCESS_LEVEL_TYPE_UNSPECIFIED is an unspecified permissions.
   *
   * @generated from enum value: ACCESS_LEVEL_TYPE_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * ACCESS_LEVEL_TYPE_NO_PERMISSIONS represents no permissions.
   *
   * @generated from enum value: ACCESS_LEVEL_TYPE_NO_PERMISSIONS = 1;
   */
  NO_PERMISSIONS = 1,

  /**
   * ACCESS_LEVEL_TYPE_MINIMAL represents "minimal" permissions to a project/group.
   *
   * @generated from enum value: ACCESS_LEVEL_TYPE_MINIMAL = 2;
   */
  MINIMAL = 2,

  /**
   * ACCESS_LEVEL_TYPE_GUEST represents "guest" permissions to a project/group.
   *
   * @generated from enum value: ACCESS_LEVEL_TYPE_GUEST = 3;
   */
  GUEST = 3,

  /**
   * ACCESS_LEVEL_TYPE_REPORTER represents "reporter" permissions to a project/group.
   *
   * @generated from enum value: ACCESS_LEVEL_TYPE_REPORTER = 4;
   */
  REPORTER = 4,

  /**
   * ACCESS_LEVEL_TYPE_DEVELOPER represents "developer" permissions to a project/group.
   *
   * @generated from enum value: ACCESS_LEVEL_TYPE_DEVELOPER = 5;
   */
  DEVELOPER = 5,

  /**
   * ACCESS_LEVEL_TYPE_MAINTAINER represents "master" permissions to a project/group.
   *
   * @generated from enum value: ACCESS_LEVEL_TYPE_MAINTAINER = 6;
   */
  MAINTAINER = 6,

  /**
   * ACCESS_LEVEL_TYPE_OWNER represents "owner" permissions to a project/group.
   *
   * @generated from enum value: ACCESS_LEVEL_TYPE_OWNER = 7;
   */
  OWNER = 7,
}

/**
 * Describes the enum accessgraph.v1alpha.AccessLevelType.
 */
export const AccessLevelTypeSchema: GenEnum<AccessLevelType> = /*@__PURE__*/
  enumDesc(file_accessgraph_v1alpha_gitlab, 0);

