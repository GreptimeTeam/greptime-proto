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
// source: greptime/v1/notify.proto

package v1

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

type NotificationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CatalogName    string `protobuf:"bytes,1,opt,name=catalog_name,json=catalogName,proto3" json:"catalog_name,omitempty"`
	SchemaName     string `protobuf:"bytes,2,opt,name=schema_name,json=schemaName,proto3" json:"schema_name,omitempty"`
	TableName      string `protobuf:"bytes,3,opt,name=table_name,json=tableName,proto3" json:"table_name,omitempty"`
	BackedFilePath string `protobuf:"bytes,4,opt,name=backed_file_path,json=backedFilePath,proto3" json:"backed_file_path,omitempty"`
	StartOffset    uint32 `protobuf:"varint,5,opt,name=start_offset,json=startOffset,proto3" json:"start_offset,omitempty"`
	EndOffset      uint32 `protobuf:"varint,6,opt,name=end_offset,json=endOffset,proto3" json:"end_offset,omitempty"`
	CheckSum       uint32 `protobuf:"varint,7,opt,name=check_sum,json=checkSum,proto3" json:"check_sum,omitempty"`
}

func (x *NotificationRequest) Reset() {
	*x = NotificationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greptime_v1_notify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotificationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationRequest) ProtoMessage() {}

func (x *NotificationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_greptime_v1_notify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationRequest.ProtoReflect.Descriptor instead.
func (*NotificationRequest) Descriptor() ([]byte, []int) {
	return file_greptime_v1_notify_proto_rawDescGZIP(), []int{0}
}

func (x *NotificationRequest) GetCatalogName() string {
	if x != nil {
		return x.CatalogName
	}
	return ""
}

func (x *NotificationRequest) GetSchemaName() string {
	if x != nil {
		return x.SchemaName
	}
	return ""
}

func (x *NotificationRequest) GetTableName() string {
	if x != nil {
		return x.TableName
	}
	return ""
}

func (x *NotificationRequest) GetBackedFilePath() string {
	if x != nil {
		return x.BackedFilePath
	}
	return ""
}

func (x *NotificationRequest) GetStartOffset() uint32 {
	if x != nil {
		return x.StartOffset
	}
	return 0
}

func (x *NotificationRequest) GetEndOffset() uint32 {
	if x != nil {
		return x.EndOffset
	}
	return 0
}

func (x *NotificationRequest) GetCheckSum() uint32 {
	if x != nil {
		return x.CheckSum
	}
	return 0
}

type NotificationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *NotificationResponse) Reset() {
	*x = NotificationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greptime_v1_notify_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotificationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationResponse) ProtoMessage() {}

func (x *NotificationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_greptime_v1_notify_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationResponse.ProtoReflect.Descriptor instead.
func (*NotificationResponse) Descriptor() ([]byte, []int) {
	return file_greptime_v1_notify_proto_rawDescGZIP(), []int{1}
}

func (x *NotificationResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_greptime_v1_notify_proto protoreflect.FileDescriptor

var file_greptime_v1_notify_proto_rawDesc = []byte{
	0x0a, 0x18, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x67, 0x72, 0x65, 0x70,
	0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x81, 0x02, 0x0a, 0x13, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x21, 0x0a, 0x0c, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x64, 0x5f, 0x66, 0x69, 0x6c,
	0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x62, 0x61,
	0x63, 0x6b, 0x65, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x21, 0x0a, 0x0c,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x65, 0x6e, 0x64, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x09, 0x65, 0x6e, 0x64, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x73, 0x75, 0x6d, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x08, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x75, 0x6d, 0x22, 0x30, 0x0a, 0x14, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0x56, 0x0a,
	0x03, 0x53, 0x68, 0x6d, 0x12, 0x4f, 0x0a, 0x06, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x20,
	0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x21, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x47, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x54, 0x65, 0x61, 0x6d,
	0x2f, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x6f, 0x2f, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_greptime_v1_notify_proto_rawDescOnce sync.Once
	file_greptime_v1_notify_proto_rawDescData = file_greptime_v1_notify_proto_rawDesc
)

func file_greptime_v1_notify_proto_rawDescGZIP() []byte {
	file_greptime_v1_notify_proto_rawDescOnce.Do(func() {
		file_greptime_v1_notify_proto_rawDescData = protoimpl.X.CompressGZIP(file_greptime_v1_notify_proto_rawDescData)
	})
	return file_greptime_v1_notify_proto_rawDescData
}

var file_greptime_v1_notify_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_greptime_v1_notify_proto_goTypes = []interface{}{
	(*NotificationRequest)(nil),  // 0: greptime.v1.NotificationRequest
	(*NotificationResponse)(nil), // 1: greptime.v1.NotificationResponse
}
var file_greptime_v1_notify_proto_depIdxs = []int32{
	0, // 0: greptime.v1.Shm.notify:input_type -> greptime.v1.NotificationRequest
	1, // 1: greptime.v1.Shm.notify:output_type -> greptime.v1.NotificationResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_greptime_v1_notify_proto_init() }
func file_greptime_v1_notify_proto_init() {
	if File_greptime_v1_notify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_greptime_v1_notify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotificationRequest); i {
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
		file_greptime_v1_notify_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotificationResponse); i {
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
			RawDescriptor: file_greptime_v1_notify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_greptime_v1_notify_proto_goTypes,
		DependencyIndexes: file_greptime_v1_notify_proto_depIdxs,
		MessageInfos:      file_greptime_v1_notify_proto_msgTypes,
	}.Build()
	File_greptime_v1_notify_proto = out.File
	file_greptime_v1_notify_proto_rawDesc = nil
	file_greptime_v1_notify_proto_goTypes = nil
	file_greptime_v1_notify_proto_depIdxs = nil
}
