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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: teleport/autoupdate/v1/autoupdate_agent_plan.proto

package autoupdate

import (
	v1 "github.com/gravitational/teleport/api/gen/proto/go/teleport/header/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Schedule type for the rollout
type Schedule int32

const (
	// UNSPECIFIED update schedule
	Schedule_SCHEDULE_UNSPECIFIED Schedule = 0
	// REGULAR update schedule
	Schedule_SCHEDULE_REGULAR Schedule = 1
	// IMMEDIATE update schedule for updating all agents immediately
	Schedule_SCHEDULE_IMMEDIATE Schedule = 2
)

// Enum value maps for Schedule.
var (
	Schedule_name = map[int32]string{
		0: "SCHEDULE_UNSPECIFIED",
		1: "SCHEDULE_REGULAR",
		2: "SCHEDULE_IMMEDIATE",
	}
	Schedule_value = map[string]int32{
		"SCHEDULE_UNSPECIFIED": 0,
		"SCHEDULE_REGULAR":     1,
		"SCHEDULE_IMMEDIATE":   2,
	}
)

func (x Schedule) Enum() *Schedule {
	p := new(Schedule)
	*p = x
	return p
}

func (x Schedule) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Schedule) Descriptor() protoreflect.EnumDescriptor {
	return file_teleport_autoupdate_v1_autoupdate_agent_plan_proto_enumTypes[0].Descriptor()
}

func (Schedule) Type() protoreflect.EnumType {
	return &file_teleport_autoupdate_v1_autoupdate_agent_plan_proto_enumTypes[0]
}

func (x Schedule) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Schedule.Descriptor instead.
func (Schedule) EnumDescriptor() ([]byte, []int) {
	return file_teleport_autoupdate_v1_autoupdate_agent_plan_proto_rawDescGZIP(), []int{0}
}

// Strategy type for the rollout
type Strategy int32

const (
	// UNSPECIFIED update strategy
	Strategy_STRATEGY_UNSPECIFIED Strategy = 0
	// GROUPED update schedule, with no backpressure
	Strategy_STRATEGY_GROUPED Strategy = 1
	// BACKPRESSURE update schedule
	Strategy_STRATEGY_BACKPRESSURE Strategy = 2
)

// Enum value maps for Strategy.
var (
	Strategy_name = map[int32]string{
		0: "STRATEGY_UNSPECIFIED",
		1: "STRATEGY_GROUPED",
		2: "STRATEGY_BACKPRESSURE",
	}
	Strategy_value = map[string]int32{
		"STRATEGY_UNSPECIFIED":  0,
		"STRATEGY_GROUPED":      1,
		"STRATEGY_BACKPRESSURE": 2,
	}
)

func (x Strategy) Enum() *Strategy {
	p := new(Strategy)
	*p = x
	return p
}

func (x Strategy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Strategy) Descriptor() protoreflect.EnumDescriptor {
	return file_teleport_autoupdate_v1_autoupdate_agent_plan_proto_enumTypes[1].Descriptor()
}

func (Strategy) Type() protoreflect.EnumType {
	return &file_teleport_autoupdate_v1_autoupdate_agent_plan_proto_enumTypes[1]
}

func (x Strategy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Strategy.Descriptor instead.
func (Strategy) EnumDescriptor() ([]byte, []int) {
	return file_teleport_autoupdate_v1_autoupdate_agent_plan_proto_rawDescGZIP(), []int{1}
}

// Mode of operation
type Mode int32

const (
	// UNSPECIFIED update mode
	Mode_MODE_UNSPECIFIED Mode = 0
	// DISABLE updates
	Mode_MODE_DISABLE Mode = 1
	// ENABLE updates
	Mode_MODE_ENABLE Mode = 2
	// PAUSE updates
	Mode_MODE_PAUSE Mode = 3
)

// Enum value maps for Mode.
var (
	Mode_name = map[int32]string{
		0: "MODE_UNSPECIFIED",
		1: "MODE_DISABLE",
		2: "MODE_ENABLE",
		3: "MODE_PAUSE",
	}
	Mode_value = map[string]int32{
		"MODE_UNSPECIFIED": 0,
		"MODE_DISABLE":     1,
		"MODE_ENABLE":      2,
		"MODE_PAUSE":       3,
	}
)

func (x Mode) Enum() *Mode {
	p := new(Mode)
	*p = x
	return p
}

func (x Mode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Mode) Descriptor() protoreflect.EnumDescriptor {
	return file_teleport_autoupdate_v1_autoupdate_agent_plan_proto_enumTypes[2].Descriptor()
}

func (Mode) Type() protoreflect.EnumType {
	return &file_teleport_autoupdate_v1_autoupdate_agent_plan_proto_enumTypes[2]
}

func (x Mode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Mode.Descriptor instead.
func (Mode) EnumDescriptor() ([]byte, []int) {
	return file_teleport_autoupdate_v1_autoupdate_agent_plan_proto_rawDescGZIP(), []int{2}
}

// State of the rollout
type State int32

const (
	// UNSPECIFIED state
	State_STATE_UNSPECIFIED State = 0
	// UNSTARTED state
	State_STATE_UNSTARTED State = 1
	// CANARY state
	State_STATE_CANARY State = 2
	// ACTIVE state
	State_STATE_ACTIVE State = 3
	// DONE state
	State_STATE_DONE State = 4
	// ROLLEDBACK state
	State_STATE_ROLLEDBACK State = 5
)

// Enum value maps for State.
var (
	State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "STATE_UNSTARTED",
		2: "STATE_CANARY",
		3: "STATE_ACTIVE",
		4: "STATE_DONE",
		5: "STATE_ROLLEDBACK",
	}
	State_value = map[string]int32{
		"STATE_UNSPECIFIED": 0,
		"STATE_UNSTARTED":   1,
		"STATE_CANARY":      2,
		"STATE_ACTIVE":      3,
		"STATE_DONE":        4,
		"STATE_ROLLEDBACK":  5,
	}
)

func (x State) Enum() *State {
	p := new(State)
	*p = x
	return p
}

func (x State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (State) Descriptor() protoreflect.EnumDescriptor {
	return file_teleport_autoupdate_v1_autoupdate_agent_plan_proto_enumTypes[3].Descriptor()
}

func (State) Type() protoreflect.EnumType {
	return &file_teleport_autoupdate_v1_autoupdate_agent_plan_proto_enumTypes[3]
}

func (x State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use State.Descriptor instead.
func (State) EnumDescriptor() ([]byte, []int) {
	return file_teleport_autoupdate_v1_autoupdate_agent_plan_proto_rawDescGZIP(), []int{3}
}

// AutoUpdateAgentPlan holds dynamic configuration settings for agent autoupdates.
type AutoUpdateAgentPlan struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// kind is the kind of the resource.
	Kind string `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	// sub_kind is the sub kind of the resource.
	SubKind string `protobuf:"bytes,2,opt,name=sub_kind,json=subKind,proto3" json:"sub_kind,omitempty"`
	// version is the version of the resource.
	Version string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	// metadata is the metadata of the resource.
	Metadata *v1.Metadata `protobuf:"bytes,4,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// spec is the spec of the resource.
	Spec *AutoUpdateAgentPlanSpec `protobuf:"bytes,5,opt,name=spec,proto3" json:"spec,omitempty"`
	// status is the status of the resource.
	Status *AutoUpdateAgentPlanStatus `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *AutoUpdateAgentPlan) Reset() {
	*x = AutoUpdateAgentPlan{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_autoupdate_v1_autoupdate_agent_plan_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AutoUpdateAgentPlan) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AutoUpdateAgentPlan) ProtoMessage() {}

func (x *AutoUpdateAgentPlan) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_autoupdate_v1_autoupdate_agent_plan_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AutoUpdateAgentPlan.ProtoReflect.Descriptor instead.
func (*AutoUpdateAgentPlan) Descriptor() ([]byte, []int) {
	return file_teleport_autoupdate_v1_autoupdate_agent_plan_proto_rawDescGZIP(), []int{0}
}

func (x *AutoUpdateAgentPlan) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *AutoUpdateAgentPlan) GetSubKind() string {
	if x != nil {
		return x.SubKind
	}
	return ""
}

func (x *AutoUpdateAgentPlan) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *AutoUpdateAgentPlan) GetMetadata() *v1.Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *AutoUpdateAgentPlan) GetSpec() *AutoUpdateAgentPlanSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *AutoUpdateAgentPlan) GetStatus() *AutoUpdateAgentPlanStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

// AutoUpdateAgentPlanSpec is the spec for the autoupdate version.
type AutoUpdateAgentPlanSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// start_version is the version to update from.
	StartVersion string `protobuf:"bytes,1,opt,name=start_version,json=startVersion,proto3" json:"start_version,omitempty"`
	// target_version is the version to update to.
	TargetVersion string `protobuf:"bytes,2,opt,name=target_version,json=targetVersion,proto3" json:"target_version,omitempty"`
	// schedule to use for the rollout
	Schedule Schedule `protobuf:"varint,3,opt,name=schedule,proto3,enum=teleport.autoupdate.v1.Schedule" json:"schedule,omitempty"`
	// strategy to use for the rollout
	Strategy Strategy `protobuf:"varint,4,opt,name=strategy,proto3,enum=teleport.autoupdate.v1.Strategy" json:"strategy,omitempty"`
	// autoupdate_mode to use for the rollout
	AutoupdateMode Mode `protobuf:"varint,5,opt,name=autoupdate_mode,json=autoupdateMode,proto3,enum=teleport.autoupdate.v1.Mode" json:"autoupdate_mode,omitempty"`
}

func (x *AutoUpdateAgentPlanSpec) Reset() {
	*x = AutoUpdateAgentPlanSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_autoupdate_v1_autoupdate_agent_plan_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AutoUpdateAgentPlanSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AutoUpdateAgentPlanSpec) ProtoMessage() {}

func (x *AutoUpdateAgentPlanSpec) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_autoupdate_v1_autoupdate_agent_plan_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AutoUpdateAgentPlanSpec.ProtoReflect.Descriptor instead.
func (*AutoUpdateAgentPlanSpec) Descriptor() ([]byte, []int) {
	return file_teleport_autoupdate_v1_autoupdate_agent_plan_proto_rawDescGZIP(), []int{1}
}

func (x *AutoUpdateAgentPlanSpec) GetStartVersion() string {
	if x != nil {
		return x.StartVersion
	}
	return ""
}

func (x *AutoUpdateAgentPlanSpec) GetTargetVersion() string {
	if x != nil {
		return x.TargetVersion
	}
	return ""
}

func (x *AutoUpdateAgentPlanSpec) GetSchedule() Schedule {
	if x != nil {
		return x.Schedule
	}
	return Schedule_SCHEDULE_UNSPECIFIED
}

func (x *AutoUpdateAgentPlanSpec) GetStrategy() Strategy {
	if x != nil {
		return x.Strategy
	}
	return Strategy_STRATEGY_UNSPECIFIED
}

func (x *AutoUpdateAgentPlanSpec) GetAutoupdateMode() Mode {
	if x != nil {
		return x.AutoupdateMode
	}
	return Mode_MODE_UNSPECIFIED
}

// AutoUpdateAgentPlanStatus is the status for the AutoUpdateAgentPlan.
type AutoUpdateAgentPlanStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name of the group
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// start_time of the rollout
	StartTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// initial_count is the number of connected agents at the start of the window.
	InitialCount int64 `protobuf:"varint,3,opt,name=initial_count,json=initialCount,proto3" json:"initial_count,omitempty"`
	// present_count is the current number of connected agents.
	PresentCount int64 `protobuf:"varint,4,opt,name=present_count,json=presentCount,proto3" json:"present_count,omitempty"`
	// failed_count specifies the number of failed agents.
	FailedCount int64 `protobuf:"varint,5,opt,name=failed_count,json=failedCount,proto3" json:"failed_count,omitempty"`
	// canaries is a list of canary agents.
	Canaries []*Canary `protobuf:"bytes,6,rep,name=canaries,proto3" json:"canaries,omitempty"`
	// progress is the current progress through the rollout.
	Progress float32 `protobuf:"fixed32,7,opt,name=progress,proto3" json:"progress,omitempty"`
	// state is the current state of the rollout.
	State State `protobuf:"varint,8,opt,name=state,proto3,enum=teleport.autoupdate.v1.State" json:"state,omitempty"`
	// last_update_time is the time of the previous update for this group.
	LastUpdateTime *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=last_update_time,json=lastUpdateTime,proto3" json:"last_update_time,omitempty"`
	// last_update_reason is the trigger for the last update
	LastUpdateReason string `protobuf:"bytes,10,opt,name=last_update_reason,json=lastUpdateReason,proto3" json:"last_update_reason,omitempty"`
}

func (x *AutoUpdateAgentPlanStatus) Reset() {
	*x = AutoUpdateAgentPlanStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_autoupdate_v1_autoupdate_agent_plan_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AutoUpdateAgentPlanStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AutoUpdateAgentPlanStatus) ProtoMessage() {}

func (x *AutoUpdateAgentPlanStatus) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_autoupdate_v1_autoupdate_agent_plan_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AutoUpdateAgentPlanStatus.ProtoReflect.Descriptor instead.
func (*AutoUpdateAgentPlanStatus) Descriptor() ([]byte, []int) {
	return file_teleport_autoupdate_v1_autoupdate_agent_plan_proto_rawDescGZIP(), []int{2}
}

func (x *AutoUpdateAgentPlanStatus) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AutoUpdateAgentPlanStatus) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *AutoUpdateAgentPlanStatus) GetInitialCount() int64 {
	if x != nil {
		return x.InitialCount
	}
	return 0
}

func (x *AutoUpdateAgentPlanStatus) GetPresentCount() int64 {
	if x != nil {
		return x.PresentCount
	}
	return 0
}

func (x *AutoUpdateAgentPlanStatus) GetFailedCount() int64 {
	if x != nil {
		return x.FailedCount
	}
	return 0
}

func (x *AutoUpdateAgentPlanStatus) GetCanaries() []*Canary {
	if x != nil {
		return x.Canaries
	}
	return nil
}

func (x *AutoUpdateAgentPlanStatus) GetProgress() float32 {
	if x != nil {
		return x.Progress
	}
	return 0
}

func (x *AutoUpdateAgentPlanStatus) GetState() State {
	if x != nil {
		return x.State
	}
	return State_STATE_UNSPECIFIED
}

func (x *AutoUpdateAgentPlanStatus) GetLastUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.LastUpdateTime
	}
	return nil
}

func (x *AutoUpdateAgentPlanStatus) GetLastUpdateReason() string {
	if x != nil {
		return x.LastUpdateReason
	}
	return ""
}

// Canary agent
type Canary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// update_uuid of the canary agent
	UpdateUuid string `protobuf:"bytes,1,opt,name=update_uuid,json=updateUuid,proto3" json:"update_uuid,omitempty"`
	// host_uuid of the canary agent
	HostUuid string `protobuf:"bytes,2,opt,name=host_uuid,json=hostUuid,proto3" json:"host_uuid,omitempty"`
	// hostname of the canary agent
	Hostname string `protobuf:"bytes,3,opt,name=hostname,proto3" json:"hostname,omitempty"`
	// success state of the canary agent
	Success bool `protobuf:"varint,4,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *Canary) Reset() {
	*x = Canary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_autoupdate_v1_autoupdate_agent_plan_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Canary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Canary) ProtoMessage() {}

func (x *Canary) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_autoupdate_v1_autoupdate_agent_plan_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Canary.ProtoReflect.Descriptor instead.
func (*Canary) Descriptor() ([]byte, []int) {
	return file_teleport_autoupdate_v1_autoupdate_agent_plan_proto_rawDescGZIP(), []int{3}
}

func (x *Canary) GetUpdateUuid() string {
	if x != nil {
		return x.UpdateUuid
	}
	return ""
}

func (x *Canary) GetHostUuid() string {
	if x != nil {
		return x.HostUuid
	}
	return ""
}

func (x *Canary) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *Canary) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_teleport_autoupdate_v1_autoupdate_agent_plan_proto protoreflect.FileDescriptor

var file_teleport_autoupdate_v1_autoupdate_agent_plan_proto_rawDesc = []byte{
	0x0a, 0x32, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x61,
	0x75, 0x74, 0x6f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x74,
	0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xa8, 0x02, 0x0a, 0x13, 0x41, 0x75, 0x74, 0x6f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41,
	0x67, 0x65, 0x6e, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x19, 0x0a, 0x08,
	0x73, 0x75, 0x62, 0x5f, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x73, 0x75, 0x62, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x38, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x68,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x43, 0x0a, 0x04, 0x73,
	0x70, 0x65, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x74, 0x65, 0x6c, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x6f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x67, 0x65,
	0x6e, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63,
	0x12, 0x49, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x31, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x61, 0x75, 0x74, 0x6f,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x6f, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xa8, 0x02, 0x0a, 0x17,
	0x41, 0x75, 0x74, 0x6f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x50,
	0x6c, 0x61, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x0e,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x3c, 0x0a, 0x08, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x2e, 0x61, 0x75, 0x74, 0x6f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x08, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x12, 0x3c, 0x0a, 0x08, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x61,
	0x75, 0x74, 0x6f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72,
	0x61, 0x74, 0x65, 0x67, 0x79, 0x52, 0x08, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x12,
	0x45, 0x0a, 0x0f, 0x61, 0x75, 0x74, 0x6f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x6f,
	0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x0e, 0x61, 0x75, 0x74, 0x6f, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x22, 0xd8, 0x03, 0x0a, 0x19, 0x41, 0x75, 0x74, 0x6f, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x69, 0x6e, 0x69, 0x74,
	0x69, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72, 0x65, 0x73,
	0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0c, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x21, 0x0a,
	0x0c, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0b, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x3a, 0x0a, 0x08, 0x63, 0x61, 0x6e, 0x61, 0x72, 0x69, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x61, 0x75,
	0x74, 0x6f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x6e, 0x61,
	0x72, 0x79, 0x52, 0x08, 0x63, 0x61, 0x6e, 0x61, 0x72, 0x69, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08,
	0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x33, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x44, 0x0a,
	0x10, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x10, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x22, 0x7c, 0x0a, 0x06, 0x43, 0x61, 0x6e, 0x61, 0x72, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x75, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x68, 0x6f, 0x73, 0x74, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x68, 0x6f, 0x73, 0x74, 0x55, 0x75, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73,
	0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73,
	0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2a,
	0x52, 0x0a, 0x08, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x14, 0x53,
	0x43, 0x48, 0x45, 0x44, 0x55, 0x4c, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x43, 0x48, 0x45, 0x44, 0x55, 0x4c,
	0x45, 0x5f, 0x52, 0x45, 0x47, 0x55, 0x4c, 0x41, 0x52, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x53,
	0x43, 0x48, 0x45, 0x44, 0x55, 0x4c, 0x45, 0x5f, 0x49, 0x4d, 0x4d, 0x45, 0x44, 0x49, 0x41, 0x54,
	0x45, 0x10, 0x02, 0x2a, 0x55, 0x0a, 0x08, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x12,
	0x18, 0x0a, 0x14, 0x53, 0x54, 0x52, 0x41, 0x54, 0x45, 0x47, 0x59, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x54, 0x52,
	0x41, 0x54, 0x45, 0x47, 0x59, 0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x45, 0x44, 0x10, 0x01, 0x12,
	0x19, 0x0a, 0x15, 0x53, 0x54, 0x52, 0x41, 0x54, 0x45, 0x47, 0x59, 0x5f, 0x42, 0x41, 0x43, 0x4b,
	0x50, 0x52, 0x45, 0x53, 0x53, 0x55, 0x52, 0x45, 0x10, 0x02, 0x2a, 0x4f, 0x0a, 0x04, 0x4d, 0x6f,
	0x64, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x4d, 0x4f, 0x44, 0x45,
	0x5f, 0x44, 0x49, 0x53, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x4d, 0x4f,
	0x44, 0x45, 0x5f, 0x45, 0x4e, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x4d,
	0x4f, 0x44, 0x45, 0x5f, 0x50, 0x41, 0x55, 0x53, 0x45, 0x10, 0x03, 0x2a, 0x7d, 0x0a, 0x05, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x53,
	0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x54, 0x41, 0x52, 0x54, 0x45, 0x44, 0x10, 0x01,
	0x12, 0x10, 0x0a, 0x0c, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x43, 0x41, 0x4e, 0x41, 0x52, 0x59,
	0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x41, 0x43, 0x54, 0x49,
	0x56, 0x45, 0x10, 0x03, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x44, 0x4f,
	0x4e, 0x45, 0x10, 0x04, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x52, 0x4f,
	0x4c, 0x4c, 0x45, 0x44, 0x42, 0x41, 0x43, 0x4b, 0x10, 0x05, 0x42, 0x56, 0x5a, 0x54, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x61, 0x76, 0x69, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f,
	0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x75, 0x74, 0x6f, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_teleport_autoupdate_v1_autoupdate_agent_plan_proto_rawDescOnce sync.Once
	file_teleport_autoupdate_v1_autoupdate_agent_plan_proto_rawDescData = file_teleport_autoupdate_v1_autoupdate_agent_plan_proto_rawDesc
)

func file_teleport_autoupdate_v1_autoupdate_agent_plan_proto_rawDescGZIP() []byte {
	file_teleport_autoupdate_v1_autoupdate_agent_plan_proto_rawDescOnce.Do(func() {
		file_teleport_autoupdate_v1_autoupdate_agent_plan_proto_rawDescData = protoimpl.X.CompressGZIP(file_teleport_autoupdate_v1_autoupdate_agent_plan_proto_rawDescData)
	})
	return file_teleport_autoupdate_v1_autoupdate_agent_plan_proto_rawDescData
}

var file_teleport_autoupdate_v1_autoupdate_agent_plan_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_teleport_autoupdate_v1_autoupdate_agent_plan_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_teleport_autoupdate_v1_autoupdate_agent_plan_proto_goTypes = []any{
	(Schedule)(0),                     // 0: teleport.autoupdate.v1.Schedule
	(Strategy)(0),                     // 1: teleport.autoupdate.v1.Strategy
	(Mode)(0),                         // 2: teleport.autoupdate.v1.Mode
	(State)(0),                        // 3: teleport.autoupdate.v1.State
	(*AutoUpdateAgentPlan)(nil),       // 4: teleport.autoupdate.v1.AutoUpdateAgentPlan
	(*AutoUpdateAgentPlanSpec)(nil),   // 5: teleport.autoupdate.v1.AutoUpdateAgentPlanSpec
	(*AutoUpdateAgentPlanStatus)(nil), // 6: teleport.autoupdate.v1.AutoUpdateAgentPlanStatus
	(*Canary)(nil),                    // 7: teleport.autoupdate.v1.Canary
	(*v1.Metadata)(nil),               // 8: teleport.header.v1.Metadata
	(*timestamppb.Timestamp)(nil),     // 9: google.protobuf.Timestamp
}
var file_teleport_autoupdate_v1_autoupdate_agent_plan_proto_depIdxs = []int32{
	8,  // 0: teleport.autoupdate.v1.AutoUpdateAgentPlan.metadata:type_name -> teleport.header.v1.Metadata
	5,  // 1: teleport.autoupdate.v1.AutoUpdateAgentPlan.spec:type_name -> teleport.autoupdate.v1.AutoUpdateAgentPlanSpec
	6,  // 2: teleport.autoupdate.v1.AutoUpdateAgentPlan.status:type_name -> teleport.autoupdate.v1.AutoUpdateAgentPlanStatus
	0,  // 3: teleport.autoupdate.v1.AutoUpdateAgentPlanSpec.schedule:type_name -> teleport.autoupdate.v1.Schedule
	1,  // 4: teleport.autoupdate.v1.AutoUpdateAgentPlanSpec.strategy:type_name -> teleport.autoupdate.v1.Strategy
	2,  // 5: teleport.autoupdate.v1.AutoUpdateAgentPlanSpec.autoupdate_mode:type_name -> teleport.autoupdate.v1.Mode
	9,  // 6: teleport.autoupdate.v1.AutoUpdateAgentPlanStatus.start_time:type_name -> google.protobuf.Timestamp
	7,  // 7: teleport.autoupdate.v1.AutoUpdateAgentPlanStatus.canaries:type_name -> teleport.autoupdate.v1.Canary
	3,  // 8: teleport.autoupdate.v1.AutoUpdateAgentPlanStatus.state:type_name -> teleport.autoupdate.v1.State
	9,  // 9: teleport.autoupdate.v1.AutoUpdateAgentPlanStatus.last_update_time:type_name -> google.protobuf.Timestamp
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_teleport_autoupdate_v1_autoupdate_agent_plan_proto_init() }
func file_teleport_autoupdate_v1_autoupdate_agent_plan_proto_init() {
	if File_teleport_autoupdate_v1_autoupdate_agent_plan_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_teleport_autoupdate_v1_autoupdate_agent_plan_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*AutoUpdateAgentPlan); i {
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
		file_teleport_autoupdate_v1_autoupdate_agent_plan_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*AutoUpdateAgentPlanSpec); i {
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
		file_teleport_autoupdate_v1_autoupdate_agent_plan_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*AutoUpdateAgentPlanStatus); i {
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
		file_teleport_autoupdate_v1_autoupdate_agent_plan_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Canary); i {
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
			RawDescriptor: file_teleport_autoupdate_v1_autoupdate_agent_plan_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_teleport_autoupdate_v1_autoupdate_agent_plan_proto_goTypes,
		DependencyIndexes: file_teleport_autoupdate_v1_autoupdate_agent_plan_proto_depIdxs,
		EnumInfos:         file_teleport_autoupdate_v1_autoupdate_agent_plan_proto_enumTypes,
		MessageInfos:      file_teleport_autoupdate_v1_autoupdate_agent_plan_proto_msgTypes,
	}.Build()
	File_teleport_autoupdate_v1_autoupdate_agent_plan_proto = out.File
	file_teleport_autoupdate_v1_autoupdate_agent_plan_proto_rawDesc = nil
	file_teleport_autoupdate_v1_autoupdate_agent_plan_proto_goTypes = nil
	file_teleport_autoupdate_v1_autoupdate_agent_plan_proto_depIdxs = nil
}
