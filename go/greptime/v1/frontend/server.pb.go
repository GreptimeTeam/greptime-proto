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
// source: greptime/v1/frontend/server.proto

package frontend

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

type ListProcessRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Catalog string `protobuf:"bytes,1,opt,name=catalog,proto3" json:"catalog,omitempty"`
}

func (x *ListProcessRequest) Reset() {
	*x = ListProcessRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greptime_v1_frontend_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListProcessRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProcessRequest) ProtoMessage() {}

func (x *ListProcessRequest) ProtoReflect() protoreflect.Message {
	mi := &file_greptime_v1_frontend_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProcessRequest.ProtoReflect.Descriptor instead.
func (*ListProcessRequest) Descriptor() ([]byte, []int) {
	return file_greptime_v1_frontend_server_proto_rawDescGZIP(), []int{0}
}

func (x *ListProcessRequest) GetCatalog() string {
	if x != nil {
		return x.Catalog
	}
	return ""
}

type ListProcessResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Processes []*ProcessInfo `protobuf:"bytes,1,rep,name=processes,proto3" json:"processes,omitempty"`
}

func (x *ListProcessResponse) Reset() {
	*x = ListProcessResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greptime_v1_frontend_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListProcessResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProcessResponse) ProtoMessage() {}

func (x *ListProcessResponse) ProtoReflect() protoreflect.Message {
	mi := &file_greptime_v1_frontend_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProcessResponse.ProtoReflect.Descriptor instead.
func (*ListProcessResponse) Descriptor() ([]byte, []int) {
	return file_greptime_v1_frontend_server_proto_rawDescGZIP(), []int{1}
}

func (x *ListProcessResponse) GetProcesses() []*ProcessInfo {
	if x != nil {
		return x.Processes
	}
	return nil
}

type KillProcessRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Frontend server address of process.
	ServerAddr string `protobuf:"bytes,1,opt,name=server_addr,json=serverAddr,proto3" json:"server_addr,omitempty"`
	// Catalog of process to kill.
	Catalog string `protobuf:"bytes,2,opt,name=catalog,proto3" json:"catalog,omitempty"`
	// ID of process to kill.
	ProcessId uint32 `protobuf:"varint,3,opt,name=process_id,json=processId,proto3" json:"process_id,omitempty"`
}

func (x *KillProcessRequest) Reset() {
	*x = KillProcessRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greptime_v1_frontend_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KillProcessRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KillProcessRequest) ProtoMessage() {}

func (x *KillProcessRequest) ProtoReflect() protoreflect.Message {
	mi := &file_greptime_v1_frontend_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KillProcessRequest.ProtoReflect.Descriptor instead.
func (*KillProcessRequest) Descriptor() ([]byte, []int) {
	return file_greptime_v1_frontend_server_proto_rawDescGZIP(), []int{2}
}

func (x *KillProcessRequest) GetServerAddr() string {
	if x != nil {
		return x.ServerAddr
	}
	return ""
}

func (x *KillProcessRequest) GetCatalog() string {
	if x != nil {
		return x.Catalog
	}
	return ""
}

func (x *KillProcessRequest) GetProcessId() uint32 {
	if x != nil {
		return x.ProcessId
	}
	return 0
}

type KillProcessResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether targeting process is successfully killed
	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *KillProcessResponse) Reset() {
	*x = KillProcessResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greptime_v1_frontend_server_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KillProcessResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KillProcessResponse) ProtoMessage() {}

func (x *KillProcessResponse) ProtoReflect() protoreflect.Message {
	mi := &file_greptime_v1_frontend_server_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KillProcessResponse.ProtoReflect.Descriptor instead.
func (*KillProcessResponse) Descriptor() ([]byte, []int) {
	return file_greptime_v1_frontend_server_proto_rawDescGZIP(), []int{3}
}

func (x *KillProcessResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type ProcessInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID.
	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Catalog of process.
	Catalog string `protobuf:"bytes,2,opt,name=catalog,proto3" json:"catalog,omitempty"`
	// Involved schemas.
	Schemas []string `protobuf:"bytes,3,rep,name=schemas,proto3" json:"schemas,omitempty"`
	// Query string.
	Query string `protobuf:"bytes,4,opt,name=query,proto3" json:"query,omitempty"`
	// Start time.
	StartTimestamp int64 `protobuf:"varint,5,opt,name=start_timestamp,json=startTimestamp,proto3" json:"start_timestamp,omitempty"`
	// Client info.
	Client string `protobuf:"bytes,6,opt,name=client,proto3" json:"client,omitempty"`
	// Frontend info of process.
	Frontend string `protobuf:"bytes,7,opt,name=frontend,proto3" json:"frontend,omitempty"`
}

func (x *ProcessInfo) Reset() {
	*x = ProcessInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greptime_v1_frontend_server_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcessInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessInfo) ProtoMessage() {}

func (x *ProcessInfo) ProtoReflect() protoreflect.Message {
	mi := &file_greptime_v1_frontend_server_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessInfo.ProtoReflect.Descriptor instead.
func (*ProcessInfo) Descriptor() ([]byte, []int) {
	return file_greptime_v1_frontend_server_proto_rawDescGZIP(), []int{4}
}

func (x *ProcessInfo) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ProcessInfo) GetCatalog() string {
	if x != nil {
		return x.Catalog
	}
	return ""
}

func (x *ProcessInfo) GetSchemas() []string {
	if x != nil {
		return x.Schemas
	}
	return nil
}

func (x *ProcessInfo) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *ProcessInfo) GetStartTimestamp() int64 {
	if x != nil {
		return x.StartTimestamp
	}
	return 0
}

func (x *ProcessInfo) GetClient() string {
	if x != nil {
		return x.Client
	}
	return ""
}

func (x *ProcessInfo) GetFrontend() string {
	if x != nil {
		return x.Frontend
	}
	return ""
}

var File_greptime_v1_frontend_server_proto protoreflect.FileDescriptor

var file_greptime_v1_frontend_server_proto_rawDesc = []byte{
	0x0a, 0x21, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x72,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x14, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x22, 0x2e, 0x0a, 0x12, 0x4c, 0x69, 0x73,
	0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x22, 0x56, 0x0a, 0x13, 0x4c, 0x69, 0x73,
	0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3f, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65,
	0x73, 0x22, 0x6e, 0x0a, 0x12, 0x4b, 0x69, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x61, 0x74, 0x61,
	0x6c, 0x6f, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x61, 0x74, 0x61, 0x6c,
	0x6f, 0x67, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49,
	0x64, 0x22, 0x2f, 0x0a, 0x13, 0x4b, 0x69, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x22, 0xc4, 0x01, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x27, 0x0a, 0x0f,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x32, 0xd2, 0x01, 0x0a, 0x08, 0x46, 0x72,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x12, 0x62, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x12, 0x28, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x29, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x66, 0x72,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x62, 0x0a, 0x0b, 0x4b, 0x69,
	0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x12, 0x28, 0x2e, 0x67, 0x72, 0x65, 0x70,
	0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64,
	0x2e, 0x4b, 0x69, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x4b, 0x69, 0x6c, 0x6c, 0x50,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x61,
	0x0a, 0x17, 0x69, 0x6f, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x42, 0x06, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x47, 0x72,
	0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x2f, 0x67, 0x72, 0x65, 0x70, 0x74,
	0x69, 0x6d, 0x65, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x67, 0x72, 0x65,
	0x70, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_greptime_v1_frontend_server_proto_rawDescOnce sync.Once
	file_greptime_v1_frontend_server_proto_rawDescData = file_greptime_v1_frontend_server_proto_rawDesc
)

func file_greptime_v1_frontend_server_proto_rawDescGZIP() []byte {
	file_greptime_v1_frontend_server_proto_rawDescOnce.Do(func() {
		file_greptime_v1_frontend_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_greptime_v1_frontend_server_proto_rawDescData)
	})
	return file_greptime_v1_frontend_server_proto_rawDescData
}

var file_greptime_v1_frontend_server_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_greptime_v1_frontend_server_proto_goTypes = []interface{}{
	(*ListProcessRequest)(nil),  // 0: greptime.v1.frontend.ListProcessRequest
	(*ListProcessResponse)(nil), // 1: greptime.v1.frontend.ListProcessResponse
	(*KillProcessRequest)(nil),  // 2: greptime.v1.frontend.KillProcessRequest
	(*KillProcessResponse)(nil), // 3: greptime.v1.frontend.KillProcessResponse
	(*ProcessInfo)(nil),         // 4: greptime.v1.frontend.ProcessInfo
}
var file_greptime_v1_frontend_server_proto_depIdxs = []int32{
	4, // 0: greptime.v1.frontend.ListProcessResponse.processes:type_name -> greptime.v1.frontend.ProcessInfo
	0, // 1: greptime.v1.frontend.Frontend.ListProcess:input_type -> greptime.v1.frontend.ListProcessRequest
	2, // 2: greptime.v1.frontend.Frontend.KillProcess:input_type -> greptime.v1.frontend.KillProcessRequest
	1, // 3: greptime.v1.frontend.Frontend.ListProcess:output_type -> greptime.v1.frontend.ListProcessResponse
	3, // 4: greptime.v1.frontend.Frontend.KillProcess:output_type -> greptime.v1.frontend.KillProcessResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_greptime_v1_frontend_server_proto_init() }
func file_greptime_v1_frontend_server_proto_init() {
	if File_greptime_v1_frontend_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_greptime_v1_frontend_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListProcessRequest); i {
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
		file_greptime_v1_frontend_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListProcessResponse); i {
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
		file_greptime_v1_frontend_server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KillProcessRequest); i {
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
		file_greptime_v1_frontend_server_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KillProcessResponse); i {
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
		file_greptime_v1_frontend_server_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProcessInfo); i {
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
			RawDescriptor: file_greptime_v1_frontend_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_greptime_v1_frontend_server_proto_goTypes,
		DependencyIndexes: file_greptime_v1_frontend_server_proto_depIdxs,
		MessageInfos:      file_greptime_v1_frontend_server_proto_msgTypes,
	}.Build()
	File_greptime_v1_frontend_server_proto = out.File
	file_greptime_v1_frontend_server_proto_rawDesc = nil
	file_greptime_v1_frontend_server_proto_goTypes = nil
	file_greptime_v1_frontend_server_proto_depIdxs = nil
}
