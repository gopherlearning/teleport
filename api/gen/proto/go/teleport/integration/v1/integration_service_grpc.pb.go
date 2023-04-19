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

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: teleport/integration/v1/integration_service.proto

package v1

import (
	context "context"
	types "github.com/gravitational/teleport/api/types"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	IntegrationService_ListIntegrations_FullMethodName      = "/teleport.integration.v1.IntegrationService/ListIntegrations"
	IntegrationService_GetIntegration_FullMethodName        = "/teleport.integration.v1.IntegrationService/GetIntegration"
	IntegrationService_CreateIntegration_FullMethodName     = "/teleport.integration.v1.IntegrationService/CreateIntegration"
	IntegrationService_UpdateIntegration_FullMethodName     = "/teleport.integration.v1.IntegrationService/UpdateIntegration"
	IntegrationService_DeleteIntegration_FullMethodName     = "/teleport.integration.v1.IntegrationService/DeleteIntegration"
	IntegrationService_DeleteAllIntegrations_FullMethodName = "/teleport.integration.v1.IntegrationService/DeleteAllIntegrations"
)

// IntegrationServiceClient is the client API for IntegrationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IntegrationServiceClient interface {
	// ListIntegrations returns a paginated list of Integration resources.
	ListIntegrations(ctx context.Context, in *ListIntegrationsRequest, opts ...grpc.CallOption) (*ListIntegrationsResponse, error)
	// GetIntegration returns the specified Integration resource.
	GetIntegration(ctx context.Context, in *GetIntegrationRequest, opts ...grpc.CallOption) (*types.IntegrationV1, error)
	// CreateIntegration creates a new Integration resource.
	CreateIntegration(ctx context.Context, in *CreateIntegrationRequest, opts ...grpc.CallOption) (*types.IntegrationV1, error)
	// UpdateIntegration updates an existing Integration resource.
	UpdateIntegration(ctx context.Context, in *UpdateIntegrationRequest, opts ...grpc.CallOption) (*types.IntegrationV1, error)
	// DeleteIntegration removes the specified Integration resource.
	DeleteIntegration(ctx context.Context, in *DeleteIntegrationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// DeleteAllIntegrations removes all Integrations.
	DeleteAllIntegrations(ctx context.Context, in *DeleteAllIntegrationsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type integrationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewIntegrationServiceClient(cc grpc.ClientConnInterface) IntegrationServiceClient {
	return &integrationServiceClient{cc}
}

func (c *integrationServiceClient) ListIntegrations(ctx context.Context, in *ListIntegrationsRequest, opts ...grpc.CallOption) (*ListIntegrationsResponse, error) {
	out := new(ListIntegrationsResponse)
	err := c.cc.Invoke(ctx, IntegrationService_ListIntegrations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *integrationServiceClient) GetIntegration(ctx context.Context, in *GetIntegrationRequest, opts ...grpc.CallOption) (*types.IntegrationV1, error) {
	out := new(types.IntegrationV1)
	err := c.cc.Invoke(ctx, IntegrationService_GetIntegration_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *integrationServiceClient) CreateIntegration(ctx context.Context, in *CreateIntegrationRequest, opts ...grpc.CallOption) (*types.IntegrationV1, error) {
	out := new(types.IntegrationV1)
	err := c.cc.Invoke(ctx, IntegrationService_CreateIntegration_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *integrationServiceClient) UpdateIntegration(ctx context.Context, in *UpdateIntegrationRequest, opts ...grpc.CallOption) (*types.IntegrationV1, error) {
	out := new(types.IntegrationV1)
	err := c.cc.Invoke(ctx, IntegrationService_UpdateIntegration_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *integrationServiceClient) DeleteIntegration(ctx context.Context, in *DeleteIntegrationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, IntegrationService_DeleteIntegration_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *integrationServiceClient) DeleteAllIntegrations(ctx context.Context, in *DeleteAllIntegrationsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, IntegrationService_DeleteAllIntegrations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IntegrationServiceServer is the server API for IntegrationService service.
// All implementations must embed UnimplementedIntegrationServiceServer
// for forward compatibility
type IntegrationServiceServer interface {
	// ListIntegrations returns a paginated list of Integration resources.
	ListIntegrations(context.Context, *ListIntegrationsRequest) (*ListIntegrationsResponse, error)
	// GetIntegration returns the specified Integration resource.
	GetIntegration(context.Context, *GetIntegrationRequest) (*types.IntegrationV1, error)
	// CreateIntegration creates a new Integration resource.
	CreateIntegration(context.Context, *CreateIntegrationRequest) (*types.IntegrationV1, error)
	// UpdateIntegration updates an existing Integration resource.
	UpdateIntegration(context.Context, *UpdateIntegrationRequest) (*types.IntegrationV1, error)
	// DeleteIntegration removes the specified Integration resource.
	DeleteIntegration(context.Context, *DeleteIntegrationRequest) (*emptypb.Empty, error)
	// DeleteAllIntegrations removes all Integrations.
	DeleteAllIntegrations(context.Context, *DeleteAllIntegrationsRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedIntegrationServiceServer()
}

// UnimplementedIntegrationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedIntegrationServiceServer struct {
}

func (UnimplementedIntegrationServiceServer) ListIntegrations(context.Context, *ListIntegrationsRequest) (*ListIntegrationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListIntegrations not implemented")
}
func (UnimplementedIntegrationServiceServer) GetIntegration(context.Context, *GetIntegrationRequest) (*types.IntegrationV1, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIntegration not implemented")
}
func (UnimplementedIntegrationServiceServer) CreateIntegration(context.Context, *CreateIntegrationRequest) (*types.IntegrationV1, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateIntegration not implemented")
}
func (UnimplementedIntegrationServiceServer) UpdateIntegration(context.Context, *UpdateIntegrationRequest) (*types.IntegrationV1, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateIntegration not implemented")
}
func (UnimplementedIntegrationServiceServer) DeleteIntegration(context.Context, *DeleteIntegrationRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteIntegration not implemented")
}
func (UnimplementedIntegrationServiceServer) DeleteAllIntegrations(context.Context, *DeleteAllIntegrationsRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAllIntegrations not implemented")
}
func (UnimplementedIntegrationServiceServer) mustEmbedUnimplementedIntegrationServiceServer() {}

// UnsafeIntegrationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IntegrationServiceServer will
// result in compilation errors.
type UnsafeIntegrationServiceServer interface {
	mustEmbedUnimplementedIntegrationServiceServer()
}

func RegisterIntegrationServiceServer(s grpc.ServiceRegistrar, srv IntegrationServiceServer) {
	s.RegisterService(&IntegrationService_ServiceDesc, srv)
}

func _IntegrationService_ListIntegrations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListIntegrationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IntegrationServiceServer).ListIntegrations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IntegrationService_ListIntegrations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IntegrationServiceServer).ListIntegrations(ctx, req.(*ListIntegrationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IntegrationService_GetIntegration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIntegrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IntegrationServiceServer).GetIntegration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IntegrationService_GetIntegration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IntegrationServiceServer).GetIntegration(ctx, req.(*GetIntegrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IntegrationService_CreateIntegration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateIntegrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IntegrationServiceServer).CreateIntegration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IntegrationService_CreateIntegration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IntegrationServiceServer).CreateIntegration(ctx, req.(*CreateIntegrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IntegrationService_UpdateIntegration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateIntegrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IntegrationServiceServer).UpdateIntegration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IntegrationService_UpdateIntegration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IntegrationServiceServer).UpdateIntegration(ctx, req.(*UpdateIntegrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IntegrationService_DeleteIntegration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteIntegrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IntegrationServiceServer).DeleteIntegration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IntegrationService_DeleteIntegration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IntegrationServiceServer).DeleteIntegration(ctx, req.(*DeleteIntegrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IntegrationService_DeleteAllIntegrations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAllIntegrationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IntegrationServiceServer).DeleteAllIntegrations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IntegrationService_DeleteAllIntegrations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IntegrationServiceServer).DeleteAllIntegrations(ctx, req.(*DeleteAllIntegrationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// IntegrationService_ServiceDesc is the grpc.ServiceDesc for IntegrationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IntegrationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "teleport.integration.v1.IntegrationService",
	HandlerType: (*IntegrationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListIntegrations",
			Handler:    _IntegrationService_ListIntegrations_Handler,
		},
		{
			MethodName: "GetIntegration",
			Handler:    _IntegrationService_GetIntegration_Handler,
		},
		{
			MethodName: "CreateIntegration",
			Handler:    _IntegrationService_CreateIntegration_Handler,
		},
		{
			MethodName: "UpdateIntegration",
			Handler:    _IntegrationService_UpdateIntegration_Handler,
		},
		{
			MethodName: "DeleteIntegration",
			Handler:    _IntegrationService_DeleteIntegration_Handler,
		},
		{
			MethodName: "DeleteAllIntegrations",
			Handler:    _IntegrationService_DeleteAllIntegrations_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "teleport/integration/v1/integration_service.proto",
}
