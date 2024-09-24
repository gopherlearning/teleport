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
// @generated from file teleport/trait/v1/trait.proto (package teleport.trait.v1, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, messageDesc } from "@bufbuild/protobuf/codegenv1";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file teleport/trait/v1/trait.proto.
 */
export const file_teleport_trait_v1_trait: GenFile = /*@__PURE__*/
  fileDesc("Ch10ZWxlcG9ydC90cmFpdC92MS90cmFpdC5wcm90bxIRdGVsZXBvcnQudHJhaXQudjEiJAoFVHJhaXQSCwoDa2V5GAEgASgJEg4KBnZhbHVlcxgCIAMoCUJOWkxnaXRodWIuY29tL2dyYXZpdGF0aW9uYWwvdGVsZXBvcnQvYXBpL2dlbi9wcm90by9nby90ZWxlcG9ydC90cmFpdC92MTt0cmFpdHYxYgZwcm90bzM");

/**
 * Trait is a trait that can be use in various resources.
 *
 * @generated from message teleport.trait.v1.Trait
 */
export type Trait = Message<"teleport.trait.v1.Trait"> & {
  /**
   * key is the name of the trait.
   *
   * @generated from field: string key = 1;
   */
  key: string;

  /**
   * values is the list of trait values.
   *
   * @generated from field: repeated string values = 2;
   */
  values: string[];
};

/**
 * Describes the message teleport.trait.v1.Trait.
 * Use `create(TraitSchema)` to create a new message.
 */
export const TraitSchema: GenMessage<Trait> = /*@__PURE__*/
  messageDesc(file_teleport_trait_v1_trait, 0);

