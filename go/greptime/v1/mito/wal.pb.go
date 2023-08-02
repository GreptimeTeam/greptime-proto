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
// source: greptime/v1/mito/wal.proto

package mito

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

// Type of operation to rows.
type OpType int32

const (
	// Delete rows.
	OpType_DELETE OpType = 0
	// Put rows.
	OpType_PUT OpType = 1
)

// Enum value maps for OpType.
var (
	OpType_name = map[int32]string{
		0: "DELETE",
		1: "PUT",
	}
	OpType_value = map[string]int32{
		"DELETE": 0,
		"PUT":    1,
	}
)

func (x OpType) Enum() *OpType {
	p := new(OpType)
	*p = x
	return p
}

func (x OpType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OpType) Descriptor() protoreflect.EnumDescriptor {
	return file_greptime_v1_mito_wal_proto_enumTypes[0].Descriptor()
}

func (OpType) Type() protoreflect.EnumType {
	return &file_greptime_v1_mito_wal_proto_enumTypes[0]
}

func (x OpType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OpType.Descriptor instead.
func (OpType) EnumDescriptor() ([]byte, []int) {
	return file_greptime_v1_mito_wal_proto_rawDescGZIP(), []int{0}
}

// Mutation contains updates to a set of rows.
type Mutation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Type of this mutation.
	OpType OpType `protobuf:"varint,1,opt,name=op_type,json=opType,proto3,enum=greptime.v1.mito.OpType" json:"op_type,omitempty"`
	// Start WAL sequence of this mutation.
	Sequence uint64 `protobuf:"varint,2,opt,name=sequence,proto3" json:"sequence,omitempty"`
	// Row updates to write to the WAL.
	Rows *v1.Rows `protobuf:"bytes,3,opt,name=rows,proto3" json:"rows,omitempty"`
}

func (x *Mutation) Reset() {
	*x = Mutation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greptime_v1_mito_wal_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Mutation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Mutation) ProtoMessage() {}

func (x *Mutation) ProtoReflect() protoreflect.Message {
	mi := &file_greptime_v1_mito_wal_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Mutation.ProtoReflect.Descriptor instead.
func (*Mutation) Descriptor() ([]byte, []int) {
	return file_greptime_v1_mito_wal_proto_rawDescGZIP(), []int{0}
}

func (x *Mutation) GetOpType() OpType {
	if x != nil {
		return x.OpType
	}
	return OpType_DELETE
}

func (x *Mutation) GetSequence() uint64 {
	if x != nil {
		return x.Sequence
	}
	return 0
}

func (x *Mutation) GetRows() *v1.Rows {
	if x != nil {
		return x.Rows
	}
	return nil
}

// Mutation to a region, contains a list of row mutations to the region.
type RegionMutation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Id of the region to write.
	RegionId uint64 `protobuf:"varint,1,opt,name=region_id,json=regionId,proto3" json:"region_id,omitempty"`
	// List of mutations.
	Mutations []*Mutation `protobuf:"bytes,2,rep,name=mutations,proto3" json:"mutations,omitempty"`
}

func (x *RegionMutation) Reset() {
	*x = RegionMutation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greptime_v1_mito_wal_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegionMutation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegionMutation) ProtoMessage() {}

func (x *RegionMutation) ProtoReflect() protoreflect.Message {
	mi := &file_greptime_v1_mito_wal_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegionMutation.ProtoReflect.Descriptor instead.
func (*RegionMutation) Descriptor() ([]byte, []int) {
	return file_greptime_v1_mito_wal_proto_rawDescGZIP(), []int{1}
}

func (x *RegionMutation) GetRegionId() uint64 {
	if x != nil {
		return x.RegionId
	}
	return 0
}

func (x *RegionMutation) GetMutations() []*Mutation {
	if x != nil {
		return x.Mutations
	}
	return nil
}

// Entry in the WAL.
type WalEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Mutations to regions.
	RegionMutations []*RegionMutation `protobuf:"bytes,1,rep,name=region_mutations,json=regionMutations,proto3" json:"region_mutations,omitempty"`
}

func (x *WalEntry) Reset() {
	*x = WalEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greptime_v1_mito_wal_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WalEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WalEntry) ProtoMessage() {}

func (x *WalEntry) ProtoReflect() protoreflect.Message {
	mi := &file_greptime_v1_mito_wal_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WalEntry.ProtoReflect.Descriptor instead.
func (*WalEntry) Descriptor() ([]byte, []int) {
	return file_greptime_v1_mito_wal_proto_rawDescGZIP(), []int{2}
}

func (x *WalEntry) GetRegionMutations() []*RegionMutation {
	if x != nil {
		return x.RegionMutations
	}
	return nil
}

var File_greptime_v1_mito_wal_proto protoreflect.FileDescriptor

var file_greptime_v1_mito_wal_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x69,
	0x74, 0x6f, 0x2f, 0x77, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x67, 0x72,
	0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x69, 0x74, 0x6f, 0x1a, 0x15,
	0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x77, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x80, 0x01, 0x0a, 0x08, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x31, 0x0a, 0x07, 0x6f, 0x70, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x6d, 0x69, 0x74, 0x6f, 0x2e, 0x4f, 0x70, 0x54, 0x79, 0x70, 0x65, 0x52, 0x06, 0x6f,
	0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63,
	0x65, 0x12, 0x25, 0x0a, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f,
	0x77, 0x73, 0x52, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x22, 0x67, 0x0a, 0x0e, 0x52, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x09, 0x6d, 0x75, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x72, 0x65,
	0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x69, 0x74, 0x6f, 0x2e, 0x4d, 0x75,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x6d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x22, 0x57, 0x0a, 0x08, 0x57, 0x61, 0x6c, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x4b, 0x0a,
	0x10, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69,
	0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x69, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0f, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2a, 0x1d, 0x0a, 0x06, 0x4f, 0x70,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x00,
	0x12, 0x07, 0x0a, 0x03, 0x50, 0x55, 0x54, 0x10, 0x01, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x47, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65,
	0x54, 0x65, 0x61, 0x6d, 0x2f, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2d, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x6d, 0x69, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_greptime_v1_mito_wal_proto_rawDescOnce sync.Once
	file_greptime_v1_mito_wal_proto_rawDescData = file_greptime_v1_mito_wal_proto_rawDesc
)

func file_greptime_v1_mito_wal_proto_rawDescGZIP() []byte {
	file_greptime_v1_mito_wal_proto_rawDescOnce.Do(func() {
		file_greptime_v1_mito_wal_proto_rawDescData = protoimpl.X.CompressGZIP(file_greptime_v1_mito_wal_proto_rawDescData)
	})
	return file_greptime_v1_mito_wal_proto_rawDescData
}

var file_greptime_v1_mito_wal_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_greptime_v1_mito_wal_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_greptime_v1_mito_wal_proto_goTypes = []interface{}{
	(OpType)(0),            // 0: greptime.v1.mito.OpType
	(*Mutation)(nil),       // 1: greptime.v1.mito.Mutation
	(*RegionMutation)(nil), // 2: greptime.v1.mito.RegionMutation
	(*WalEntry)(nil),       // 3: greptime.v1.mito.WalEntry
	(*v1.Rows)(nil),        // 4: greptime.v1.Rows
}
var file_greptime_v1_mito_wal_proto_depIdxs = []int32{
	0, // 0: greptime.v1.mito.Mutation.op_type:type_name -> greptime.v1.mito.OpType
	4, // 1: greptime.v1.mito.Mutation.rows:type_name -> greptime.v1.Rows
	1, // 2: greptime.v1.mito.RegionMutation.mutations:type_name -> greptime.v1.mito.Mutation
	2, // 3: greptime.v1.mito.WalEntry.region_mutations:type_name -> greptime.v1.mito.RegionMutation
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_greptime_v1_mito_wal_proto_init() }
func file_greptime_v1_mito_wal_proto_init() {
	if File_greptime_v1_mito_wal_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_greptime_v1_mito_wal_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Mutation); i {
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
		file_greptime_v1_mito_wal_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegionMutation); i {
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
		file_greptime_v1_mito_wal_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WalEntry); i {
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
			RawDescriptor: file_greptime_v1_mito_wal_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_greptime_v1_mito_wal_proto_goTypes,
		DependencyIndexes: file_greptime_v1_mito_wal_proto_depIdxs,
		EnumInfos:         file_greptime_v1_mito_wal_proto_enumTypes,
		MessageInfos:      file_greptime_v1_mito_wal_proto_msgTypes,
	}.Build()
	File_greptime_v1_mito_wal_proto = out.File
	file_greptime_v1_mito_wal_proto_rawDesc = nil
	file_greptime_v1_mito_wal_proto_goTypes = nil
	file_greptime_v1_mito_wal_proto_depIdxs = nil
}
