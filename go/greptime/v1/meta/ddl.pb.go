// Copyright 2023 Greptime Team
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: greptime/v1/meta/ddl.proto

package meta

import (
	v1 "github.com/GreptimeTeam/greptime-proto/go/greptime/v1"
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

type DdlTaskType int32

const (
	DdlTaskType_Create DdlTaskType = 0
)

// Enum value maps for DdlTaskType.
var (
	DdlTaskType_name = map[int32]string{
		0: "Create",
	}
	DdlTaskType_value = map[string]int32{
		"Create": 0,
	}
)

func (x DdlTaskType) Enum() *DdlTaskType {
	p := new(DdlTaskType)
	*p = x
	return p
}

func (x DdlTaskType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DdlTaskType) Descriptor() protoreflect.EnumDescriptor {
	return file_greptime_v1_meta_ddl_proto_enumTypes[0].Descriptor()
}

func (DdlTaskType) Type() protoreflect.EnumType {
	return &file_greptime_v1_meta_ddl_proto_enumTypes[0]
}

func (x DdlTaskType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DdlTaskType.Descriptor instead.
func (DdlTaskType) EnumDescriptor() ([]byte, []int) {
	return file_greptime_v1_meta_ddl_proto_rawDescGZIP(), []int{0}
}

type Status int32

const (
	Status_Done      Status = 0
	Status_Failed    Status = 1
	Status_Cancelled Status = 2
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "Done",
		1: "Failed",
		2: "Cancelled",
	}
	Status_value = map[string]int32{
		"Done":      0,
		"Failed":    1,
		"Cancelled": 2,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_greptime_v1_meta_ddl_proto_enumTypes[1].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_greptime_v1_meta_ddl_proto_enumTypes[1]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_greptime_v1_meta_ddl_proto_rawDescGZIP(), []int{1}
}

type CreateTableTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreateTable *v1.CreateTableExpr `protobuf:"bytes,1,opt,name=create_table,json=createTable,proto3" json:"create_table,omitempty"`
	Partitions  []*Partition        `protobuf:"bytes,2,rep,name=partitions,proto3" json:"partitions,omitempty"`
}

func (x *CreateTableTask) Reset() {
	*x = CreateTableTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greptime_v1_meta_ddl_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTableTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTableTask) ProtoMessage() {}

func (x *CreateTableTask) ProtoReflect() protoreflect.Message {
	mi := &file_greptime_v1_meta_ddl_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTableTask.ProtoReflect.Descriptor instead.
func (*CreateTableTask) Descriptor() ([]byte, []int) {
	return file_greptime_v1_meta_ddl_proto_rawDescGZIP(), []int{0}
}

func (x *CreateTableTask) GetCreateTable() *v1.CreateTableExpr {
	if x != nil {
		return x.CreateTable
	}
	return nil
}

func (x *CreateTableTask) GetPartitions() []*Partition {
	if x != nil {
		return x.Partitions
	}
	return nil
}

type SubmitDdlTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *RequestHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// The following information plays a bigger role in making messages traceable
	// and facilitating debugging.
	Subject string `protobuf:"bytes,2,opt,name=subject,proto3" json:"subject,omitempty"`
	From    string `protobuf:"bytes,3,opt,name=from,proto3" json:"from,omitempty"`
	To      string `protobuf:"bytes,4,opt,name=to,proto3" json:"to,omitempty"`
	// The unix timestamp in milliseconds.
	TimestampMillis int64 `protobuf:"varint,5,opt,name=timestamp_millis,json=timestampMillis,proto3" json:"timestamp_millis,omitempty"`
	// Types that are assignable to Task:
	//
	//	*SubmitDdlTaskRequest_CreateTableTask
	Task isSubmitDdlTaskRequest_Task `protobuf_oneof:"task"`
}

func (x *SubmitDdlTaskRequest) Reset() {
	*x = SubmitDdlTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greptime_v1_meta_ddl_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitDdlTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitDdlTaskRequest) ProtoMessage() {}

func (x *SubmitDdlTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_greptime_v1_meta_ddl_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitDdlTaskRequest.ProtoReflect.Descriptor instead.
func (*SubmitDdlTaskRequest) Descriptor() ([]byte, []int) {
	return file_greptime_v1_meta_ddl_proto_rawDescGZIP(), []int{1}
}

func (x *SubmitDdlTaskRequest) GetHeader() *RequestHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *SubmitDdlTaskRequest) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *SubmitDdlTaskRequest) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *SubmitDdlTaskRequest) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *SubmitDdlTaskRequest) GetTimestampMillis() int64 {
	if x != nil {
		return x.TimestampMillis
	}
	return 0
}

func (m *SubmitDdlTaskRequest) GetTask() isSubmitDdlTaskRequest_Task {
	if m != nil {
		return m.Task
	}
	return nil
}

func (x *SubmitDdlTaskRequest) GetCreateTableTask() *CreateTableTask {
	if x, ok := x.GetTask().(*SubmitDdlTaskRequest_CreateTableTask); ok {
		return x.CreateTableTask
	}
	return nil
}

type isSubmitDdlTaskRequest_Task interface {
	isSubmitDdlTaskRequest_Task()
}

type SubmitDdlTaskRequest_CreateTableTask struct {
	CreateTableTask *CreateTableTask `protobuf:"bytes,6,opt,name=create_table_task,json=createTableTask,proto3,oneof"`
}

func (*SubmitDdlTaskRequest_CreateTableTask) isSubmitDdlTaskRequest_Task() {}

type SubmitDdlTaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// Key is the identifier for the ddl task.
	Key    []byte `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Status Status `protobuf:"varint,3,opt,name=status,proto3,enum=greptime.v1.meta.Status" json:"status,omitempty"`
}

func (x *SubmitDdlTaskResponse) Reset() {
	*x = SubmitDdlTaskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greptime_v1_meta_ddl_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitDdlTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitDdlTaskResponse) ProtoMessage() {}

func (x *SubmitDdlTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_greptime_v1_meta_ddl_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitDdlTaskResponse.ProtoReflect.Descriptor instead.
func (*SubmitDdlTaskResponse) Descriptor() ([]byte, []int) {
	return file_greptime_v1_meta_ddl_proto_rawDescGZIP(), []int{2}
}

func (x *SubmitDdlTaskResponse) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *SubmitDdlTaskResponse) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *SubmitDdlTaskResponse) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_Done
}

var File_greptime_v1_meta_ddl_proto protoreflect.FileDescriptor

var file_greptime_v1_meta_ddl_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65,
	0x74, 0x61, 0x2f, 0x64, 0x64, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x67, 0x72,
	0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x1a, 0x1d,
	0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x61,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67,
	0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2f,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x67, 0x72, 0x65,
	0x70, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x64, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x8f, 0x01, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x62,
	0x6c, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x3f, 0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67,
	0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x45, 0x78, 0x70, 0x72, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x3b, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x74, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x72,
	0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x50,
	0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x22, 0x91, 0x02, 0x0a, 0x14, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x44,
	0x64, 0x6c, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a,
	0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e,
	0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x74, 0x61,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x74, 0x6f, 0x12, 0x29, 0x0a, 0x10, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x5f, 0x6d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x4d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x12,
	0x4f, 0x0a, 0x11, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f,
	0x74, 0x61, 0x73, 0x6b, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x67, 0x72, 0x65,
	0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x48, 0x00, 0x52,
	0x0f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x61, 0x73, 0x6b,
	0x42, 0x06, 0x0a, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x22, 0x95, 0x01, 0x0a, 0x15, 0x53, 0x75, 0x62,
	0x6d, 0x69, 0x74, 0x44, 0x64, 0x6c, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x38, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x30,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18,
	0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x74,
	0x61, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x2a, 0x19, 0x0a, 0x0b, 0x44, 0x64, 0x6c, 0x54, 0x61, 0x73, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x0a, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x10, 0x00, 0x2a, 0x2d, 0x0a, 0x06, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x6f, 0x6e, 0x65, 0x10, 0x00, 0x12,
	0x0a, 0x0a, 0x06, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x43,
	0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x65, 0x64, 0x10, 0x02, 0x32, 0x6b, 0x0a, 0x07, 0x44, 0x64,
	0x6c, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x60, 0x0a, 0x0d, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x44,
	0x64, 0x6c, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x26, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74,
	0x44, 0x64, 0x6c, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27,
	0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x74,
	0x61, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x44, 0x64, 0x6c, 0x54, 0x61, 0x73, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x47, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x54, 0x65,
	0x61, 0x6d, 0x2f, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2d, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x6d, 0x65, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_greptime_v1_meta_ddl_proto_rawDescOnce sync.Once
	file_greptime_v1_meta_ddl_proto_rawDescData = file_greptime_v1_meta_ddl_proto_rawDesc
)

func file_greptime_v1_meta_ddl_proto_rawDescGZIP() []byte {
	file_greptime_v1_meta_ddl_proto_rawDescOnce.Do(func() {
		file_greptime_v1_meta_ddl_proto_rawDescData = protoimpl.X.CompressGZIP(file_greptime_v1_meta_ddl_proto_rawDescData)
	})
	return file_greptime_v1_meta_ddl_proto_rawDescData
}

var file_greptime_v1_meta_ddl_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_greptime_v1_meta_ddl_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_greptime_v1_meta_ddl_proto_goTypes = []interface{}{
	(DdlTaskType)(0),              // 0: greptime.v1.meta.DdlTaskType
	(Status)(0),                   // 1: greptime.v1.meta.Status
	(*CreateTableTask)(nil),       // 2: greptime.v1.meta.CreateTableTask
	(*SubmitDdlTaskRequest)(nil),  // 3: greptime.v1.meta.SubmitDdlTaskRequest
	(*SubmitDdlTaskResponse)(nil), // 4: greptime.v1.meta.SubmitDdlTaskResponse
	(*v1.CreateTableExpr)(nil),    // 5: greptime.v1.CreateTableExpr
	(*Partition)(nil),             // 6: greptime.v1.meta.Partition
	(*RequestHeader)(nil),         // 7: greptime.v1.meta.RequestHeader
	(*ResponseHeader)(nil),        // 8: greptime.v1.meta.ResponseHeader
}
var file_greptime_v1_meta_ddl_proto_depIdxs = []int32{
	5, // 0: greptime.v1.meta.CreateTableTask.create_table:type_name -> greptime.v1.CreateTableExpr
	6, // 1: greptime.v1.meta.CreateTableTask.partitions:type_name -> greptime.v1.meta.Partition
	7, // 2: greptime.v1.meta.SubmitDdlTaskRequest.header:type_name -> greptime.v1.meta.RequestHeader
	2, // 3: greptime.v1.meta.SubmitDdlTaskRequest.create_table_task:type_name -> greptime.v1.meta.CreateTableTask
	8, // 4: greptime.v1.meta.SubmitDdlTaskResponse.header:type_name -> greptime.v1.meta.ResponseHeader
	1, // 5: greptime.v1.meta.SubmitDdlTaskResponse.status:type_name -> greptime.v1.meta.Status
	3, // 6: greptime.v1.meta.DdlTask.SubmitDdlTask:input_type -> greptime.v1.meta.SubmitDdlTaskRequest
	4, // 7: greptime.v1.meta.DdlTask.SubmitDdlTask:output_type -> greptime.v1.meta.SubmitDdlTaskResponse
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_greptime_v1_meta_ddl_proto_init() }
func file_greptime_v1_meta_ddl_proto_init() {
	if File_greptime_v1_meta_ddl_proto != nil {
		return
	}
	file_greptime_v1_meta_common_proto_init()
	file_greptime_v1_meta_route_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_greptime_v1_meta_ddl_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTableTask); i {
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
		file_greptime_v1_meta_ddl_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitDdlTaskRequest); i {
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
		file_greptime_v1_meta_ddl_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitDdlTaskResponse); i {
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
	file_greptime_v1_meta_ddl_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*SubmitDdlTaskRequest_CreateTableTask)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_greptime_v1_meta_ddl_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_greptime_v1_meta_ddl_proto_goTypes,
		DependencyIndexes: file_greptime_v1_meta_ddl_proto_depIdxs,
		EnumInfos:         file_greptime_v1_meta_ddl_proto_enumTypes,
		MessageInfos:      file_greptime_v1_meta_ddl_proto_msgTypes,
	}.Build()
	File_greptime_v1_meta_ddl_proto = out.File
	file_greptime_v1_meta_ddl_proto_rawDesc = nil
	file_greptime_v1_meta_ddl_proto_goTypes = nil
	file_greptime_v1_meta_ddl_proto_depIdxs = nil
}
