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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: teleport/resourceusage/v1/resourceusage_service.proto

package resourceusagev1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// GetUsageRequest is the request for GetUsage
type GetUsageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetUsageRequest) Reset() {
	*x = GetUsageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_resourceusage_v1_resourceusage_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUsageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUsageRequest) ProtoMessage() {}

func (x *GetUsageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_resourceusage_v1_resourceusage_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUsageRequest.ProtoReflect.Descriptor instead.
func (*GetUsageRequest) Descriptor() ([]byte, []int) {
	return file_teleport_resourceusage_v1_resourceusage_service_proto_rawDescGZIP(), []int{0}
}

// GetUsageResponse is the response for GetUsage
type GetUsageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessRequests *AccessRequestsUsage `protobuf:"bytes,1,opt,name=access_requests,json=accessRequests,proto3" json:"access_requests,omitempty"`
}

func (x *GetUsageResponse) Reset() {
	*x = GetUsageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_resourceusage_v1_resourceusage_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUsageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUsageResponse) ProtoMessage() {}

func (x *GetUsageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_resourceusage_v1_resourceusage_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUsageResponse.ProtoReflect.Descriptor instead.
func (*GetUsageResponse) Descriptor() ([]byte, []int) {
	return file_teleport_resourceusage_v1_resourceusage_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetUsageResponse) GetAccessRequests() *AccessRequestsUsage {
	if x != nil {
		return x.AccessRequests
	}
	return nil
}

// AccessRequestsUsage defines the usage limits for access requests.
// Currently this is limited on the basis of access requests used per calendar month.
type AccessRequestsUsage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// MonthlyLimit is the amount of requests that are allowed per month
	MonthlyLimit int32 `protobuf:"varint,1,opt,name=monthly_limit,json=monthlyLimit,proto3" json:"monthly_limit,omitempty"`
	// MonthlyUsed is the amount of requests that have been used this month
	MonthlyUsed int32 `protobuf:"varint,2,opt,name=monthly_used,json=monthlyUsed,proto3" json:"monthly_used,omitempty"`
}

func (x *AccessRequestsUsage) Reset() {
	*x = AccessRequestsUsage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_resourceusage_v1_resourceusage_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessRequestsUsage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessRequestsUsage) ProtoMessage() {}

func (x *AccessRequestsUsage) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_resourceusage_v1_resourceusage_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessRequestsUsage.ProtoReflect.Descriptor instead.
func (*AccessRequestsUsage) Descriptor() ([]byte, []int) {
	return file_teleport_resourceusage_v1_resourceusage_service_proto_rawDescGZIP(), []int{2}
}

func (x *AccessRequestsUsage) GetMonthlyLimit() int32 {
	if x != nil {
		return x.MonthlyLimit
	}
	return 0
}

func (x *AccessRequestsUsage) GetMonthlyUsed() int32 {
	if x != nil {
		return x.MonthlyUsed
	}
	return 0
}

var File_teleport_resourceusage_v1_resourceusage_service_proto protoreflect.FileDescriptor

var file_teleport_resourceusage_v1_resourceusage_service_proto_rawDesc = []byte{
	0x0a, 0x35, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x75, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x76, 0x31, 0x22, 0x11, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x6b, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x55, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x57, 0x0a, 0x0f, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x55, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x0e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x73, 0x22, 0x5d, 0x0a, 0x13, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x73, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x6f, 0x6e,
	0x74, 0x68, 0x6c, 0x79, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0c, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x21,
	0x0a, 0x0c, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x5f, 0x75, 0x73, 0x65, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x55, 0x73, 0x65,
	0x64, 0x32, 0x7b, 0x0a, 0x14, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x55, 0x73, 0x61,
	0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x63, 0x0a, 0x08, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2a, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2b, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x5e,
	0x5a, 0x5c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x61,
	0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x75, 0x73, 0x61, 0x67, 0x65, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_teleport_resourceusage_v1_resourceusage_service_proto_rawDescOnce sync.Once
	file_teleport_resourceusage_v1_resourceusage_service_proto_rawDescData = file_teleport_resourceusage_v1_resourceusage_service_proto_rawDesc
)

func file_teleport_resourceusage_v1_resourceusage_service_proto_rawDescGZIP() []byte {
	file_teleport_resourceusage_v1_resourceusage_service_proto_rawDescOnce.Do(func() {
		file_teleport_resourceusage_v1_resourceusage_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_teleport_resourceusage_v1_resourceusage_service_proto_rawDescData)
	})
	return file_teleport_resourceusage_v1_resourceusage_service_proto_rawDescData
}

var file_teleport_resourceusage_v1_resourceusage_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_teleport_resourceusage_v1_resourceusage_service_proto_goTypes = []interface{}{
	(*GetUsageRequest)(nil),     // 0: teleport.resourceusage.v1.GetUsageRequest
	(*GetUsageResponse)(nil),    // 1: teleport.resourceusage.v1.GetUsageResponse
	(*AccessRequestsUsage)(nil), // 2: teleport.resourceusage.v1.AccessRequestsUsage
}
var file_teleport_resourceusage_v1_resourceusage_service_proto_depIdxs = []int32{
	2, // 0: teleport.resourceusage.v1.GetUsageResponse.access_requests:type_name -> teleport.resourceusage.v1.AccessRequestsUsage
	0, // 1: teleport.resourceusage.v1.ResourceUsageService.GetUsage:input_type -> teleport.resourceusage.v1.GetUsageRequest
	1, // 2: teleport.resourceusage.v1.ResourceUsageService.GetUsage:output_type -> teleport.resourceusage.v1.GetUsageResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_teleport_resourceusage_v1_resourceusage_service_proto_init() }
func file_teleport_resourceusage_v1_resourceusage_service_proto_init() {
	if File_teleport_resourceusage_v1_resourceusage_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_teleport_resourceusage_v1_resourceusage_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUsageRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_teleport_resourceusage_v1_resourceusage_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUsageResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_teleport_resourceusage_v1_resourceusage_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccessRequestsUsage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_teleport_resourceusage_v1_resourceusage_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_teleport_resourceusage_v1_resourceusage_service_proto_goTypes,
		DependencyIndexes: file_teleport_resourceusage_v1_resourceusage_service_proto_depIdxs,
		MessageInfos:      file_teleport_resourceusage_v1_resourceusage_service_proto_msgTypes,
	}.Build()
	File_teleport_resourceusage_v1_resourceusage_service_proto = out.File
	file_teleport_resourceusage_v1_resourceusage_service_proto_rawDesc = nil
	file_teleport_resourceusage_v1_resourceusage_service_proto_goTypes = nil
	file_teleport_resourceusage_v1_resourceusage_service_proto_depIdxs = nil
}
