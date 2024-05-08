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
// source: greptime/v1/meta/cluster.proto

package meta

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

type MetasrvPeersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *RequestHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
}

func (x *MetasrvPeersRequest) Reset() {
	*x = MetasrvPeersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greptime_v1_meta_cluster_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetasrvPeersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetasrvPeersRequest) ProtoMessage() {}

func (x *MetasrvPeersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_greptime_v1_meta_cluster_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetasrvPeersRequest.ProtoReflect.Descriptor instead.
func (*MetasrvPeersRequest) Descriptor() ([]byte, []int) {
	return file_greptime_v1_meta_cluster_proto_rawDescGZIP(), []int{0}
}

func (x *MetasrvPeersRequest) GetHeader() *RequestHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

type MetasrvPeersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header    *ResponseHeader    `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Leader    *MetasrvNodeInfo   `protobuf:"bytes,2,opt,name=leader,proto3" json:"leader,omitempty"`
	Followers []*MetasrvNodeInfo `protobuf:"bytes,3,rep,name=followers,proto3" json:"followers,omitempty"`
}

func (x *MetasrvPeersResponse) Reset() {
	*x = MetasrvPeersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greptime_v1_meta_cluster_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetasrvPeersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetasrvPeersResponse) ProtoMessage() {}

func (x *MetasrvPeersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_greptime_v1_meta_cluster_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetasrvPeersResponse.ProtoReflect.Descriptor instead.
func (*MetasrvPeersResponse) Descriptor() ([]byte, []int) {
	return file_greptime_v1_meta_cluster_proto_rawDescGZIP(), []int{1}
}

func (x *MetasrvPeersResponse) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *MetasrvPeersResponse) GetLeader() *MetasrvNodeInfo {
	if x != nil {
		return x.Leader
	}
	return nil
}

func (x *MetasrvPeersResponse) GetFollowers() []*MetasrvNodeInfo {
	if x != nil {
		return x.Followers
	}
	return nil
}

type MetasrvNodeInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Peer      *Peer  `protobuf:"bytes,1,opt,name=peer,proto3" json:"peer,omitempty"`
	Version   string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	GitCommit string `protobuf:"bytes,3,opt,name=git_commit,json=gitCommit,proto3" json:"git_commit,omitempty"`
	// The node start timestamp in milliseconds.
	StartTimeMs uint64 `protobuf:"varint,4,opt,name=start_time_ms,json=startTimeMs,proto3" json:"start_time_ms,omitempty"`
}

func (x *MetasrvNodeInfo) Reset() {
	*x = MetasrvNodeInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greptime_v1_meta_cluster_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetasrvNodeInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetasrvNodeInfo) ProtoMessage() {}

func (x *MetasrvNodeInfo) ProtoReflect() protoreflect.Message {
	mi := &file_greptime_v1_meta_cluster_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetasrvNodeInfo.ProtoReflect.Descriptor instead.
func (*MetasrvNodeInfo) Descriptor() ([]byte, []int) {
	return file_greptime_v1_meta_cluster_proto_rawDescGZIP(), []int{2}
}

func (x *MetasrvNodeInfo) GetPeer() *Peer {
	if x != nil {
		return x.Peer
	}
	return nil
}

func (x *MetasrvNodeInfo) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *MetasrvNodeInfo) GetGitCommit() string {
	if x != nil {
		return x.GitCommit
	}
	return ""
}

func (x *MetasrvNodeInfo) GetStartTimeMs() uint64 {
	if x != nil {
		return x.StartTimeMs
	}
	return 0
}

var File_greptime_v1_meta_cluster_proto protoreflect.FileDescriptor

var file_greptime_v1_meta_cluster_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65,
	0x74, 0x61, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x10, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65,
	0x74, 0x61, 0x1a, 0x1d, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x76, 0x31, 0x2f,
	0x6d, 0x65, 0x74, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d,
	0x65, 0x74, 0x61, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x4e, 0x0a, 0x13, 0x4d, 0x65, 0x74, 0x61, 0x73, 0x72, 0x76, 0x50, 0x65, 0x65, 0x72, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x22,
	0xcc, 0x01, 0x0a, 0x14, 0x4d, 0x65, 0x74, 0x61, 0x73, 0x72, 0x76, 0x50, 0x65, 0x65, 0x72, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74,
	0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x12, 0x39, 0x0a, 0x06, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x21, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x73, 0x72, 0x76, 0x4e, 0x6f, 0x64,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x3f, 0x0a,
	0x09, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x21, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d,
	0x65, 0x74, 0x61, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x73, 0x72, 0x76, 0x4e, 0x6f, 0x64, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x09, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x22, 0x9a,
	0x01, 0x0a, 0x0f, 0x4d, 0x65, 0x74, 0x61, 0x73, 0x72, 0x76, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x2a, 0x0a, 0x04, 0x70, 0x65, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d,
	0x65, 0x74, 0x61, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x52, 0x04, 0x70, 0x65, 0x65, 0x72, 0x12, 0x18,
	0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x69, 0x74, 0x5f,
	0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x69,
	0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x12, 0x22, 0x0a, 0x0d, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x4d, 0x73, 0x32, 0x85, 0x02, 0x0a, 0x07,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x51, 0x0a, 0x08, 0x42, 0x61, 0x74, 0x63, 0x68,
	0x47, 0x65, 0x74, 0x12, 0x21, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x05, 0x52, 0x61,
	0x6e, 0x67, 0x65, 0x12, 0x1e, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5d, 0x0a, 0x0c, 0x4d, 0x65, 0x74, 0x61, 0x73, 0x72, 0x76, 0x50,
	0x65, 0x65, 0x72, 0x73, 0x12, 0x25, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x73, 0x72, 0x76, 0x50,
	0x65, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x67, 0x72,
	0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x4d,
	0x65, 0x74, 0x61, 0x73, 0x72, 0x76, 0x50, 0x65, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x47, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x54, 0x65, 0x61, 0x6d, 0x2f, 0x67,
	0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f,
	0x2f, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74,
	0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_greptime_v1_meta_cluster_proto_rawDescOnce sync.Once
	file_greptime_v1_meta_cluster_proto_rawDescData = file_greptime_v1_meta_cluster_proto_rawDesc
)

func file_greptime_v1_meta_cluster_proto_rawDescGZIP() []byte {
	file_greptime_v1_meta_cluster_proto_rawDescOnce.Do(func() {
		file_greptime_v1_meta_cluster_proto_rawDescData = protoimpl.X.CompressGZIP(file_greptime_v1_meta_cluster_proto_rawDescData)
	})
	return file_greptime_v1_meta_cluster_proto_rawDescData
}

var file_greptime_v1_meta_cluster_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_greptime_v1_meta_cluster_proto_goTypes = []interface{}{
	(*MetasrvPeersRequest)(nil),  // 0: greptime.v1.meta.MetasrvPeersRequest
	(*MetasrvPeersResponse)(nil), // 1: greptime.v1.meta.MetasrvPeersResponse
	(*MetasrvNodeInfo)(nil),      // 2: greptime.v1.meta.MetasrvNodeInfo
	(*RequestHeader)(nil),        // 3: greptime.v1.meta.RequestHeader
	(*ResponseHeader)(nil),       // 4: greptime.v1.meta.ResponseHeader
	(*Peer)(nil),                 // 5: greptime.v1.meta.Peer
	(*BatchGetRequest)(nil),      // 6: greptime.v1.meta.BatchGetRequest
	(*RangeRequest)(nil),         // 7: greptime.v1.meta.RangeRequest
	(*BatchGetResponse)(nil),     // 8: greptime.v1.meta.BatchGetResponse
	(*RangeResponse)(nil),        // 9: greptime.v1.meta.RangeResponse
}
var file_greptime_v1_meta_cluster_proto_depIdxs = []int32{
	3, // 0: greptime.v1.meta.MetasrvPeersRequest.header:type_name -> greptime.v1.meta.RequestHeader
	4, // 1: greptime.v1.meta.MetasrvPeersResponse.header:type_name -> greptime.v1.meta.ResponseHeader
	2, // 2: greptime.v1.meta.MetasrvPeersResponse.leader:type_name -> greptime.v1.meta.MetasrvNodeInfo
	2, // 3: greptime.v1.meta.MetasrvPeersResponse.followers:type_name -> greptime.v1.meta.MetasrvNodeInfo
	5, // 4: greptime.v1.meta.MetasrvNodeInfo.peer:type_name -> greptime.v1.meta.Peer
	6, // 5: greptime.v1.meta.Cluster.BatchGet:input_type -> greptime.v1.meta.BatchGetRequest
	7, // 6: greptime.v1.meta.Cluster.Range:input_type -> greptime.v1.meta.RangeRequest
	0, // 7: greptime.v1.meta.Cluster.MetasrvPeers:input_type -> greptime.v1.meta.MetasrvPeersRequest
	8, // 8: greptime.v1.meta.Cluster.BatchGet:output_type -> greptime.v1.meta.BatchGetResponse
	9, // 9: greptime.v1.meta.Cluster.Range:output_type -> greptime.v1.meta.RangeResponse
	1, // 10: greptime.v1.meta.Cluster.MetasrvPeers:output_type -> greptime.v1.meta.MetasrvPeersResponse
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_greptime_v1_meta_cluster_proto_init() }
func file_greptime_v1_meta_cluster_proto_init() {
	if File_greptime_v1_meta_cluster_proto != nil {
		return
	}
	file_greptime_v1_meta_common_proto_init()
	file_greptime_v1_meta_store_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_greptime_v1_meta_cluster_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetasrvPeersRequest); i {
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
		file_greptime_v1_meta_cluster_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetasrvPeersResponse); i {
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
		file_greptime_v1_meta_cluster_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetasrvNodeInfo); i {
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
			RawDescriptor: file_greptime_v1_meta_cluster_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_greptime_v1_meta_cluster_proto_goTypes,
		DependencyIndexes: file_greptime_v1_meta_cluster_proto_depIdxs,
		MessageInfos:      file_greptime_v1_meta_cluster_proto_msgTypes,
	}.Build()
	File_greptime_v1_meta_cluster_proto = out.File
	file_greptime_v1_meta_cluster_proto_rawDesc = nil
	file_greptime_v1_meta_cluster_proto_goTypes = nil
	file_greptime_v1_meta_cluster_proto_depIdxs = nil
}
