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
// @generated from file teleport/kubewaitingcontainer/v1/kubewaitingcontainer_service.proto (package teleport.kubewaitingcontainer.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { CreateKubernetesWaitingContainerRequest, DeleteKubernetesWaitingContainerRequest, GetKubernetesWaitingContainerRequest, ListKubernetesWaitingContainersRequest, ListKubernetesWaitingContainersResponse } from "./kubewaitingcontainer_service_pb.js";
import { Empty, MethodKind } from "@bufbuild/protobuf";
import { KubernetesWaitingContainer } from "./kubewaitingcontainer_pb.js";

/**
 * KubeWaitingContainersService manages Kubernetes ephemeral
 * containers that are waiting to be created until moderated
 * session conditions are met.
 *
 * @generated from service teleport.kubewaitingcontainer.v1.KubeWaitingContainersService
 */
export const KubeWaitingContainersService = {
  typeName: "teleport.kubewaitingcontainer.v1.KubeWaitingContainersService",
  methods: {
    /**
     * ListKubernetesWaitingContainers returns a Kubernetes ephemeral
     * container that is waiting to be created.
     *
     * @generated from rpc teleport.kubewaitingcontainer.v1.KubeWaitingContainersService.ListKubernetesWaitingContainers
     */
    listKubernetesWaitingContainers: {
      name: "ListKubernetesWaitingContainers",
      I: ListKubernetesWaitingContainersRequest,
      O: ListKubernetesWaitingContainersResponse,
      kind: MethodKind.Unary,
    },
    /**
     * GetKubernetesWaitingContainer returns a Kubernetes ephemeral
     * container that is waiting to be created.
     *
     * @generated from rpc teleport.kubewaitingcontainer.v1.KubeWaitingContainersService.GetKubernetesWaitingContainer
     */
    getKubernetesWaitingContainer: {
      name: "GetKubernetesWaitingContainer",
      I: GetKubernetesWaitingContainerRequest,
      O: KubernetesWaitingContainer,
      kind: MethodKind.Unary,
    },
    /**
     * CreateKubernetesWaitingContainer creates a Kubernetes ephemeral
     * container that is waiting to be created.
     *
     * @generated from rpc teleport.kubewaitingcontainer.v1.KubeWaitingContainersService.CreateKubernetesWaitingContainer
     */
    createKubernetesWaitingContainer: {
      name: "CreateKubernetesWaitingContainer",
      I: CreateKubernetesWaitingContainerRequest,
      O: KubernetesWaitingContainer,
      kind: MethodKind.Unary,
    },
    /**
     * DeleteKubernetesWaitingContainer deletes a Kubernetes ephemeral
     * container that is waiting to be created.
     *
     * @generated from rpc teleport.kubewaitingcontainer.v1.KubeWaitingContainersService.DeleteKubernetesWaitingContainer
     */
    deleteKubernetesWaitingContainer: {
      name: "DeleteKubernetesWaitingContainer",
      I: DeleteKubernetesWaitingContainerRequest,
      O: Empty,
      kind: MethodKind.Unary,
    },
  }
} as const;

