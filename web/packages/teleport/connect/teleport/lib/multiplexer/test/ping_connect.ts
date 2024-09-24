//
// Teleport
// Copyright (C) 2023  Gravitational, Inc.
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

// @generated by protoc-gen-connect-es v1.5.0 with parameter "target=ts"
// @generated from file teleport/lib/multiplexer/test/ping.proto (package teleport.lib.multiplexer.test, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { Request, Response } from "./ping_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * Pinger is a service used in tests
 *
 * @generated from service teleport.lib.multiplexer.test.Pinger
 */
export const Pinger = {
  typeName: "teleport.lib.multiplexer.test.Pinger",
  methods: {
    /**
     * @generated from rpc teleport.lib.multiplexer.test.Pinger.Ping
     */
    ping: {
      name: "Ping",
      I: Request,
      O: Response,
      kind: MethodKind.Unary,
    },
  }
} as const;

