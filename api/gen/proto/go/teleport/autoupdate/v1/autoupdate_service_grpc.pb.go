// Copyright 2024 Gravitational, Inc.
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
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: teleport/autoupdate/v1/autoupdate_service.proto

package autoupdate

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	AutoupdateService_GetAutoupdateConfig_FullMethodName     = "/teleport.autoupdate.v1.AutoupdateService/GetAutoupdateConfig"
	AutoupdateService_UpsertAutoupdateConfig_FullMethodName  = "/teleport.autoupdate.v1.AutoupdateService/UpsertAutoupdateConfig"
	AutoupdateService_DeleteAutoupdateConfig_FullMethodName  = "/teleport.autoupdate.v1.AutoupdateService/DeleteAutoupdateConfig"
	AutoupdateService_GetAutoupdateVersion_FullMethodName    = "/teleport.autoupdate.v1.AutoupdateService/GetAutoupdateVersion"
	AutoupdateService_UpsertAutoupdateVersion_FullMethodName = "/teleport.autoupdate.v1.AutoupdateService/UpsertAutoupdateVersion"
	AutoupdateService_DeleteAutoupdateVersion_FullMethodName = "/teleport.autoupdate.v1.AutoupdateService/DeleteAutoupdateVersion"
)

// AutoupdateServiceClient is the client API for AutoupdateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// AutoupdateService provides an API to manage autoupdates.
type AutoupdateServiceClient interface {
	// GetAutoupdateConfig gets the current autoupdate config singleton.
	GetAutoupdateConfig(ctx context.Context, in *GetAutoupdateConfigRequest, opts ...grpc.CallOption) (*AutoupdateConfig, error)
	// UpsertAutoupdateConfig creates a new AutoupdateConfig or replaces an existing AutoupdateConfig.
	UpsertAutoupdateConfig(ctx context.Context, in *UpsertAutoupdateConfigRequest, opts ...grpc.CallOption) (*AutoupdateConfig, error)
	// DeleteAutoupdateConfig hard deletes the specified AutoupdateConfig.
	DeleteAutoupdateConfig(ctx context.Context, in *DeleteAutoupdateConfigRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// GetAutoupdateVersion gets the current autoupdate version singleton.
	GetAutoupdateVersion(ctx context.Context, in *GetAutoupdateVersionRequest, opts ...grpc.CallOption) (*AutoupdateVersion, error)
	// UpsertAutoupdateVersion creates a new AutoupdateVersion or replaces an existing AutoupdateVersion.
	UpsertAutoupdateVersion(ctx context.Context, in *UpsertAutoupdateVersionRequest, opts ...grpc.CallOption) (*AutoupdateVersion, error)
	// DeleteAutoupdateVersion hard deletes the specified AutoupdateVersionRequest.
	DeleteAutoupdateVersion(ctx context.Context, in *DeleteAutoupdateVersionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type autoupdateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAutoupdateServiceClient(cc grpc.ClientConnInterface) AutoupdateServiceClient {
	return &autoupdateServiceClient{cc}
}

func (c *autoupdateServiceClient) GetAutoupdateConfig(ctx context.Context, in *GetAutoupdateConfigRequest, opts ...grpc.CallOption) (*AutoupdateConfig, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AutoupdateConfig)
	err := c.cc.Invoke(ctx, AutoupdateService_GetAutoupdateConfig_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *autoupdateServiceClient) UpsertAutoupdateConfig(ctx context.Context, in *UpsertAutoupdateConfigRequest, opts ...grpc.CallOption) (*AutoupdateConfig, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AutoupdateConfig)
	err := c.cc.Invoke(ctx, AutoupdateService_UpsertAutoupdateConfig_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *autoupdateServiceClient) DeleteAutoupdateConfig(ctx context.Context, in *DeleteAutoupdateConfigRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, AutoupdateService_DeleteAutoupdateConfig_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *autoupdateServiceClient) GetAutoupdateVersion(ctx context.Context, in *GetAutoupdateVersionRequest, opts ...grpc.CallOption) (*AutoupdateVersion, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AutoupdateVersion)
	err := c.cc.Invoke(ctx, AutoupdateService_GetAutoupdateVersion_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *autoupdateServiceClient) UpsertAutoupdateVersion(ctx context.Context, in *UpsertAutoupdateVersionRequest, opts ...grpc.CallOption) (*AutoupdateVersion, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AutoupdateVersion)
	err := c.cc.Invoke(ctx, AutoupdateService_UpsertAutoupdateVersion_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *autoupdateServiceClient) DeleteAutoupdateVersion(ctx context.Context, in *DeleteAutoupdateVersionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, AutoupdateService_DeleteAutoupdateVersion_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AutoupdateServiceServer is the server API for AutoupdateService service.
// All implementations must embed UnimplementedAutoupdateServiceServer
// for forward compatibility.
//
// AutoupdateService provides an API to manage autoupdates.
type AutoupdateServiceServer interface {
	// GetAutoupdateConfig gets the current autoupdate config singleton.
	GetAutoupdateConfig(context.Context, *GetAutoupdateConfigRequest) (*AutoupdateConfig, error)
	// UpsertAutoupdateConfig creates a new AutoupdateConfig or replaces an existing AutoupdateConfig.
	UpsertAutoupdateConfig(context.Context, *UpsertAutoupdateConfigRequest) (*AutoupdateConfig, error)
	// DeleteAutoupdateConfig hard deletes the specified AutoupdateConfig.
	DeleteAutoupdateConfig(context.Context, *DeleteAutoupdateConfigRequest) (*emptypb.Empty, error)
	// GetAutoupdateVersion gets the current autoupdate version singleton.
	GetAutoupdateVersion(context.Context, *GetAutoupdateVersionRequest) (*AutoupdateVersion, error)
	// UpsertAutoupdateVersion creates a new AutoupdateVersion or replaces an existing AutoupdateVersion.
	UpsertAutoupdateVersion(context.Context, *UpsertAutoupdateVersionRequest) (*AutoupdateVersion, error)
	// DeleteAutoupdateVersion hard deletes the specified AutoupdateVersionRequest.
	DeleteAutoupdateVersion(context.Context, *DeleteAutoupdateVersionRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedAutoupdateServiceServer()
}

// UnimplementedAutoupdateServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAutoupdateServiceServer struct{}

func (UnimplementedAutoupdateServiceServer) GetAutoupdateConfig(context.Context, *GetAutoupdateConfigRequest) (*AutoupdateConfig, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAutoupdateConfig not implemented")
}
func (UnimplementedAutoupdateServiceServer) UpsertAutoupdateConfig(context.Context, *UpsertAutoupdateConfigRequest) (*AutoupdateConfig, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertAutoupdateConfig not implemented")
}
func (UnimplementedAutoupdateServiceServer) DeleteAutoupdateConfig(context.Context, *DeleteAutoupdateConfigRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAutoupdateConfig not implemented")
}
func (UnimplementedAutoupdateServiceServer) GetAutoupdateVersion(context.Context, *GetAutoupdateVersionRequest) (*AutoupdateVersion, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAutoupdateVersion not implemented")
}
func (UnimplementedAutoupdateServiceServer) UpsertAutoupdateVersion(context.Context, *UpsertAutoupdateVersionRequest) (*AutoupdateVersion, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertAutoupdateVersion not implemented")
}
func (UnimplementedAutoupdateServiceServer) DeleteAutoupdateVersion(context.Context, *DeleteAutoupdateVersionRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAutoupdateVersion not implemented")
}
func (UnimplementedAutoupdateServiceServer) mustEmbedUnimplementedAutoupdateServiceServer() {}
func (UnimplementedAutoupdateServiceServer) testEmbeddedByValue()                           {}

// UnsafeAutoupdateServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AutoupdateServiceServer will
// result in compilation errors.
type UnsafeAutoupdateServiceServer interface {
	mustEmbedUnimplementedAutoupdateServiceServer()
}

func RegisterAutoupdateServiceServer(s grpc.ServiceRegistrar, srv AutoupdateServiceServer) {
	// If the following call pancis, it indicates UnimplementedAutoupdateServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AutoupdateService_ServiceDesc, srv)
}

func _AutoupdateService_GetAutoupdateConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAutoupdateConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutoupdateServiceServer).GetAutoupdateConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AutoupdateService_GetAutoupdateConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutoupdateServiceServer).GetAutoupdateConfig(ctx, req.(*GetAutoupdateConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AutoupdateService_UpsertAutoupdateConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertAutoupdateConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutoupdateServiceServer).UpsertAutoupdateConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AutoupdateService_UpsertAutoupdateConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutoupdateServiceServer).UpsertAutoupdateConfig(ctx, req.(*UpsertAutoupdateConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AutoupdateService_DeleteAutoupdateConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAutoupdateConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutoupdateServiceServer).DeleteAutoupdateConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AutoupdateService_DeleteAutoupdateConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutoupdateServiceServer).DeleteAutoupdateConfig(ctx, req.(*DeleteAutoupdateConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AutoupdateService_GetAutoupdateVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAutoupdateVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutoupdateServiceServer).GetAutoupdateVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AutoupdateService_GetAutoupdateVersion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutoupdateServiceServer).GetAutoupdateVersion(ctx, req.(*GetAutoupdateVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AutoupdateService_UpsertAutoupdateVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertAutoupdateVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutoupdateServiceServer).UpsertAutoupdateVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AutoupdateService_UpsertAutoupdateVersion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutoupdateServiceServer).UpsertAutoupdateVersion(ctx, req.(*UpsertAutoupdateVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AutoupdateService_DeleteAutoupdateVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAutoupdateVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutoupdateServiceServer).DeleteAutoupdateVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AutoupdateService_DeleteAutoupdateVersion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutoupdateServiceServer).DeleteAutoupdateVersion(ctx, req.(*DeleteAutoupdateVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AutoupdateService_ServiceDesc is the grpc.ServiceDesc for AutoupdateService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AutoupdateService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "teleport.autoupdate.v1.AutoupdateService",
	HandlerType: (*AutoupdateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAutoupdateConfig",
			Handler:    _AutoupdateService_GetAutoupdateConfig_Handler,
		},
		{
			MethodName: "UpsertAutoupdateConfig",
			Handler:    _AutoupdateService_UpsertAutoupdateConfig_Handler,
		},
		{
			MethodName: "DeleteAutoupdateConfig",
			Handler:    _AutoupdateService_DeleteAutoupdateConfig_Handler,
		},
		{
			MethodName: "GetAutoupdateVersion",
			Handler:    _AutoupdateService_GetAutoupdateVersion_Handler,
		},
		{
			MethodName: "UpsertAutoupdateVersion",
			Handler:    _AutoupdateService_UpsertAutoupdateVersion_Handler,
		},
		{
			MethodName: "DeleteAutoupdateVersion",
			Handler:    _AutoupdateService_DeleteAutoupdateVersion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "teleport/autoupdate/v1/autoupdate_service.proto",
}
