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
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: teleport/userpreferences/v1/cluster_preferences.proto

package userpreferencesv1

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

// PinnedResourcesUserPreferences is a collection of resource IDs that will be
// displayed in the user's pinned resources tab in the Web UI.
type PinnedResourcesUserPreferences struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// resource_ids is a list of unified resource name sort keys.
	ResourceIds []string `protobuf:"bytes,1,rep,name=resource_ids,json=resourceIds,proto3" json:"resource_ids,omitempty"`
}

func (x *PinnedResourcesUserPreferences) Reset() {
	*x = PinnedResourcesUserPreferences{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_userpreferences_v1_cluster_preferences_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PinnedResourcesUserPreferences) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PinnedResourcesUserPreferences) ProtoMessage() {}

func (x *PinnedResourcesUserPreferences) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_userpreferences_v1_cluster_preferences_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PinnedResourcesUserPreferences.ProtoReflect.Descriptor instead.
func (*PinnedResourcesUserPreferences) Descriptor() ([]byte, []int) {
	return file_teleport_userpreferences_v1_cluster_preferences_proto_rawDescGZIP(), []int{0}
}

func (x *PinnedResourcesUserPreferences) GetResourceIds() []string {
	if x != nil {
		return x.ResourceIds
	}
	return nil
}

// ClusterUserPreferences are user preferences saved per cluster.
type ClusterUserPreferences struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// pinned_resources is a list of pinned resources.
	PinnedResources *PinnedResourcesUserPreferences `protobuf:"bytes,1,opt,name=pinned_resources,json=pinnedResources,proto3" json:"pinned_resources,omitempty"`
}

func (x *ClusterUserPreferences) Reset() {
	*x = ClusterUserPreferences{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_userpreferences_v1_cluster_preferences_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterUserPreferences) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterUserPreferences) ProtoMessage() {}

func (x *ClusterUserPreferences) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_userpreferences_v1_cluster_preferences_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterUserPreferences.ProtoReflect.Descriptor instead.
func (*ClusterUserPreferences) Descriptor() ([]byte, []int) {
	return file_teleport_userpreferences_v1_cluster_preferences_proto_rawDescGZIP(), []int{1}
}

func (x *ClusterUserPreferences) GetPinnedResources() *PinnedResourcesUserPreferences {
	if x != nil {
		return x.PinnedResources
	}
	return nil
}

var File_teleport_userpreferences_v1_cluster_preferences_proto protoreflect.FileDescriptor

var file_teleport_userpreferences_v1_cluster_preferences_proto_rawDesc = []byte{
	0x0a, 0x35, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x70,
	0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x22, 0x43, 0x0a, 0x1e, 0x50, 0x69, 0x6e, 0x6e, 0x65, 0x64, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x65, 0x66, 0x65,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x73, 0x22, 0x80, 0x01, 0x0a, 0x16, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x73, 0x12, 0x66, 0x0a, 0x10, 0x70, 0x69, 0x6e, 0x6e, 0x65, 0x64, 0x5f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3b,
	0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x70, 0x72,
	0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x69, 0x6e,
	0x6e, 0x65, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x55, 0x73, 0x65, 0x72,
	0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x52, 0x0f, 0x70, 0x69, 0x6e,
	0x6e, 0x65, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x42, 0x59, 0x5a, 0x57,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x61, 0x76, 0x69,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x75, 0x73, 0x65, 0x72, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x73, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_teleport_userpreferences_v1_cluster_preferences_proto_rawDescOnce sync.Once
	file_teleport_userpreferences_v1_cluster_preferences_proto_rawDescData = file_teleport_userpreferences_v1_cluster_preferences_proto_rawDesc
)

func file_teleport_userpreferences_v1_cluster_preferences_proto_rawDescGZIP() []byte {
	file_teleport_userpreferences_v1_cluster_preferences_proto_rawDescOnce.Do(func() {
		file_teleport_userpreferences_v1_cluster_preferences_proto_rawDescData = protoimpl.X.CompressGZIP(file_teleport_userpreferences_v1_cluster_preferences_proto_rawDescData)
	})
	return file_teleport_userpreferences_v1_cluster_preferences_proto_rawDescData
}

var file_teleport_userpreferences_v1_cluster_preferences_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_teleport_userpreferences_v1_cluster_preferences_proto_goTypes = []interface{}{
	(*PinnedResourcesUserPreferences)(nil), // 0: teleport.userpreferences.v1.PinnedResourcesUserPreferences
	(*ClusterUserPreferences)(nil),         // 1: teleport.userpreferences.v1.ClusterUserPreferences
}
var file_teleport_userpreferences_v1_cluster_preferences_proto_depIdxs = []int32{
	0, // 0: teleport.userpreferences.v1.ClusterUserPreferences.pinned_resources:type_name -> teleport.userpreferences.v1.PinnedResourcesUserPreferences
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_teleport_userpreferences_v1_cluster_preferences_proto_init() }
func file_teleport_userpreferences_v1_cluster_preferences_proto_init() {
	if File_teleport_userpreferences_v1_cluster_preferences_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_teleport_userpreferences_v1_cluster_preferences_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PinnedResourcesUserPreferences); i {
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
		file_teleport_userpreferences_v1_cluster_preferences_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterUserPreferences); i {
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
			RawDescriptor: file_teleport_userpreferences_v1_cluster_preferences_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_teleport_userpreferences_v1_cluster_preferences_proto_goTypes,
		DependencyIndexes: file_teleport_userpreferences_v1_cluster_preferences_proto_depIdxs,
		MessageInfos:      file_teleport_userpreferences_v1_cluster_preferences_proto_msgTypes,
	}.Build()
	File_teleport_userpreferences_v1_cluster_preferences_proto = out.File
	file_teleport_userpreferences_v1_cluster_preferences_proto_rawDesc = nil
	file_teleport_userpreferences_v1_cluster_preferences_proto_goTypes = nil
	file_teleport_userpreferences_v1_cluster_preferences_proto_depIdxs = nil
}
