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

// @generated by protoc-gen-connect-es v1.5.0 with parameter "target=ts"
// @generated from file teleport/accessmonitoringrules/v1/access_monitoring_rules_service.proto (package teleport.accessmonitoringrules.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { AccessMonitoringRule, CreateAccessMonitoringRuleRequest, DeleteAccessMonitoringRuleRequest, GetAccessMonitoringRuleRequest, ListAccessMonitoringRulesRequest, ListAccessMonitoringRulesResponse, ListAccessMonitoringRulesWithFilterRequest, ListAccessMonitoringRulesWithFilterResponse, UpdateAccessMonitoringRuleRequest, UpsertAccessMonitoringRuleRequest } from "./access_monitoring_rules_pb.js";
import { Empty, MethodKind } from "@bufbuild/protobuf";

/**
 * AccessMonitoringRulesService provides CRUD methods for Access Monitoring Rules resources.
 *
 * @generated from service teleport.accessmonitoringrules.v1.AccessMonitoringRulesService
 */
export const AccessMonitoringRulesService = {
  typeName: "teleport.accessmonitoringrules.v1.AccessMonitoringRulesService",
  methods: {
    /**
     * CreateAccessMonitoringRule creates the specified access monitoring rule.
     *
     * @generated from rpc teleport.accessmonitoringrules.v1.AccessMonitoringRulesService.CreateAccessMonitoringRule
     */
    createAccessMonitoringRule: {
      name: "CreateAccessMonitoringRule",
      I: CreateAccessMonitoringRuleRequest,
      O: AccessMonitoringRule,
      kind: MethodKind.Unary,
    },
    /**
     * UpdateAccessMonitoringRule updates the specified access monitoring rule.
     *
     * @generated from rpc teleport.accessmonitoringrules.v1.AccessMonitoringRulesService.UpdateAccessMonitoringRule
     */
    updateAccessMonitoringRule: {
      name: "UpdateAccessMonitoringRule",
      I: UpdateAccessMonitoringRuleRequest,
      O: AccessMonitoringRule,
      kind: MethodKind.Unary,
    },
    /**
     * UpsertAccessMonitoringRule upserts the specified access monitoring rule.
     *
     * @generated from rpc teleport.accessmonitoringrules.v1.AccessMonitoringRulesService.UpsertAccessMonitoringRule
     */
    upsertAccessMonitoringRule: {
      name: "UpsertAccessMonitoringRule",
      I: UpsertAccessMonitoringRuleRequest,
      O: AccessMonitoringRule,
      kind: MethodKind.Unary,
    },
    /**
     * GetAccessMonitoringRule gets the specified access monitoring rule.
     *
     * @generated from rpc teleport.accessmonitoringrules.v1.AccessMonitoringRulesService.GetAccessMonitoringRule
     */
    getAccessMonitoringRule: {
      name: "GetAccessMonitoringRule",
      I: GetAccessMonitoringRuleRequest,
      O: AccessMonitoringRule,
      kind: MethodKind.Unary,
    },
    /**
     * DeleteAccessMonitoringRule deletes the specified access monitoring rule.
     *
     * @generated from rpc teleport.accessmonitoringrules.v1.AccessMonitoringRulesService.DeleteAccessMonitoringRule
     */
    deleteAccessMonitoringRule: {
      name: "DeleteAccessMonitoringRule",
      I: DeleteAccessMonitoringRuleRequest,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * ListAccessMonitoringRules lists current access monitoring rules.
     *
     * @generated from rpc teleport.accessmonitoringrules.v1.AccessMonitoringRulesService.ListAccessMonitoringRules
     */
    listAccessMonitoringRules: {
      name: "ListAccessMonitoringRules",
      I: ListAccessMonitoringRulesRequest,
      O: ListAccessMonitoringRulesResponse,
      kind: MethodKind.Unary,
    },
    /**
     * ListAccessMonitoringRulesWithFilter lists current access monitoring rules that match the provided filter.
     *
     * @generated from rpc teleport.accessmonitoringrules.v1.AccessMonitoringRulesService.ListAccessMonitoringRulesWithFilter
     */
    listAccessMonitoringRulesWithFilter: {
      name: "ListAccessMonitoringRulesWithFilter",
      I: ListAccessMonitoringRulesWithFilterRequest,
      O: ListAccessMonitoringRulesWithFilterResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

