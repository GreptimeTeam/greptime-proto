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
// source: substrait_extension/promql_plan.proto

package substrait_extension

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

type EmptyMetric struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Start timestamp in millisecond
	Start int64 `protobuf:"varint,1,opt,name=start,proto3" json:"start,omitempty"`
	// End timestamp in millisecond
	End int64 `protobuf:"varint,2,opt,name=end,proto3" json:"end,omitempty"`
	// Interval in millisecond
	Interval int64 `protobuf:"varint,3,opt,name=interval,proto3" json:"interval,omitempty"`
}

func (x *EmptyMetric) Reset() {
	*x = EmptyMetric{}
	if protoimpl.UnsafeEnabled {
		mi := &file_substrait_extension_promql_plan_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyMetric) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyMetric) ProtoMessage() {}

func (x *EmptyMetric) ProtoReflect() protoreflect.Message {
	mi := &file_substrait_extension_promql_plan_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyMetric.ProtoReflect.Descriptor instead.
func (*EmptyMetric) Descriptor() ([]byte, []int) {
	return file_substrait_extension_promql_plan_proto_rawDescGZIP(), []int{0}
}

func (x *EmptyMetric) GetStart() int64 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *EmptyMetric) GetEnd() int64 {
	if x != nil {
		return x.End
	}
	return 0
}

func (x *EmptyMetric) GetInterval() int64 {
	if x != nil {
		return x.Interval
	}
	return 0
}

type InstantManipulate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Start timestamp in millisecond
	Start int64 `protobuf:"varint,1,opt,name=start,proto3" json:"start,omitempty"`
	// End timestamp in millisecond
	End int64 `protobuf:"varint,2,opt,name=end,proto3" json:"end,omitempty"`
	// Interval in millisecond
	Interval int64 `protobuf:"varint,3,opt,name=interval,proto3" json:"interval,omitempty"`
	// Look-back delta in millisecond
	LookbackDelta int64 `protobuf:"varint,4,opt,name=lookback_delta,json=lookbackDelta,proto3" json:"lookback_delta,omitempty"`
	// Column name of time index column
	TimeIndex string `protobuf:"bytes,5,opt,name=time_index,json=timeIndex,proto3" json:"time_index,omitempty"`
	// Optional field column name for validating staleness
	FieldIndex string `protobuf:"bytes,6,opt,name=field_index,json=fieldIndex,proto3" json:"field_index,omitempty"`
}

func (x *InstantManipulate) Reset() {
	*x = InstantManipulate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_substrait_extension_promql_plan_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstantManipulate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstantManipulate) ProtoMessage() {}

func (x *InstantManipulate) ProtoReflect() protoreflect.Message {
	mi := &file_substrait_extension_promql_plan_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstantManipulate.ProtoReflect.Descriptor instead.
func (*InstantManipulate) Descriptor() ([]byte, []int) {
	return file_substrait_extension_promql_plan_proto_rawDescGZIP(), []int{1}
}

func (x *InstantManipulate) GetStart() int64 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *InstantManipulate) GetEnd() int64 {
	if x != nil {
		return x.End
	}
	return 0
}

func (x *InstantManipulate) GetInterval() int64 {
	if x != nil {
		return x.Interval
	}
	return 0
}

func (x *InstantManipulate) GetLookbackDelta() int64 {
	if x != nil {
		return x.LookbackDelta
	}
	return 0
}

func (x *InstantManipulate) GetTimeIndex() string {
	if x != nil {
		return x.TimeIndex
	}
	return ""
}

func (x *InstantManipulate) GetFieldIndex() string {
	if x != nil {
		return x.FieldIndex
	}
	return ""
}

type SeriesNormalize struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Offset in millisecond
	Offset int64 `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	// Column name of time index column
	TimeIndex string `protobuf:"bytes,2,opt,name=time_index,json=timeIndex,proto3" json:"time_index,omitempty"`
	// Whether to filter out NaN value
	FilterNan bool `protobuf:"varint,3,opt,name=filter_nan,json=filterNan,proto3" json:"filter_nan,omitempty"`
	// Names of tag columns
	TagColumns []string `protobuf:"bytes,4,rep,name=tag_columns,json=tagColumns,proto3" json:"tag_columns,omitempty"`
}

func (x *SeriesNormalize) Reset() {
	*x = SeriesNormalize{}
	if protoimpl.UnsafeEnabled {
		mi := &file_substrait_extension_promql_plan_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeriesNormalize) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeriesNormalize) ProtoMessage() {}

func (x *SeriesNormalize) ProtoReflect() protoreflect.Message {
	mi := &file_substrait_extension_promql_plan_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeriesNormalize.ProtoReflect.Descriptor instead.
func (*SeriesNormalize) Descriptor() ([]byte, []int) {
	return file_substrait_extension_promql_plan_proto_rawDescGZIP(), []int{2}
}

func (x *SeriesNormalize) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *SeriesNormalize) GetTimeIndex() string {
	if x != nil {
		return x.TimeIndex
	}
	return ""
}

func (x *SeriesNormalize) GetFilterNan() bool {
	if x != nil {
		return x.FilterNan
	}
	return false
}

func (x *SeriesNormalize) GetTagColumns() []string {
	if x != nil {
		return x.TagColumns
	}
	return nil
}

type SeriesDivide struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Names of tag columns
	TagColumns []string `protobuf:"bytes,1,rep,name=tag_columns,json=tagColumns,proto3" json:"tag_columns,omitempty"`
}

func (x *SeriesDivide) Reset() {
	*x = SeriesDivide{}
	if protoimpl.UnsafeEnabled {
		mi := &file_substrait_extension_promql_plan_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeriesDivide) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeriesDivide) ProtoMessage() {}

func (x *SeriesDivide) ProtoReflect() protoreflect.Message {
	mi := &file_substrait_extension_promql_plan_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeriesDivide.ProtoReflect.Descriptor instead.
func (*SeriesDivide) Descriptor() ([]byte, []int) {
	return file_substrait_extension_promql_plan_proto_rawDescGZIP(), []int{3}
}

func (x *SeriesDivide) GetTagColumns() []string {
	if x != nil {
		return x.TagColumns
	}
	return nil
}

type RangeManipulate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Start timestamp in millisecond
	Start int64 `protobuf:"varint,1,opt,name=start,proto3" json:"start,omitempty"`
	// End timestamp in millisecond
	End int64 `protobuf:"varint,2,opt,name=end,proto3" json:"end,omitempty"`
	// Interval in millisecond
	Interval int64 `protobuf:"varint,3,opt,name=interval,proto3" json:"interval,omitempty"`
	// Range in millisecond
	Range int64 `protobuf:"varint,4,opt,name=range,proto3" json:"range,omitempty"`
	// Column name of time index column
	TimeIndex string `protobuf:"bytes,5,opt,name=time_index,json=timeIndex,proto3" json:"time_index,omitempty"`
	// Names of tag columns
	TagColumns []string `protobuf:"bytes,6,rep,name=tag_columns,json=tagColumns,proto3" json:"tag_columns,omitempty"`
}

func (x *RangeManipulate) Reset() {
	*x = RangeManipulate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_substrait_extension_promql_plan_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RangeManipulate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RangeManipulate) ProtoMessage() {}

func (x *RangeManipulate) ProtoReflect() protoreflect.Message {
	mi := &file_substrait_extension_promql_plan_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RangeManipulate.ProtoReflect.Descriptor instead.
func (*RangeManipulate) Descriptor() ([]byte, []int) {
	return file_substrait_extension_promql_plan_proto_rawDescGZIP(), []int{4}
}

func (x *RangeManipulate) GetStart() int64 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *RangeManipulate) GetEnd() int64 {
	if x != nil {
		return x.End
	}
	return 0
}

func (x *RangeManipulate) GetInterval() int64 {
	if x != nil {
		return x.Interval
	}
	return 0
}

func (x *RangeManipulate) GetRange() int64 {
	if x != nil {
		return x.Range
	}
	return 0
}

func (x *RangeManipulate) GetTimeIndex() string {
	if x != nil {
		return x.TimeIndex
	}
	return ""
}

func (x *RangeManipulate) GetTagColumns() []string {
	if x != nil {
		return x.TagColumns
	}
	return nil
}

type ScalarCalculate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Start timestamp in millisecond
	Start int64 `protobuf:"varint,1,opt,name=start,proto3" json:"start,omitempty"`
	// End timestamp in millisecond
	End int64 `protobuf:"varint,2,opt,name=end,proto3" json:"end,omitempty"`
	// Interval in millisecond
	Interval int64 `protobuf:"varint,3,opt,name=interval,proto3" json:"interval,omitempty"`
	// Column name of time index column
	TimeIndex string `protobuf:"bytes,5,opt,name=time_index,json=timeIndex,proto3" json:"time_index,omitempty"`
	// Names of tag columns
	TagColumns []string `protobuf:"bytes,6,rep,name=tag_columns,json=tagColumns,proto3" json:"tag_columns,omitempty"`
	// Column name of field column
	FieldColumn string `protobuf:"bytes,7,opt,name=field_column,json=fieldColumn,proto3" json:"field_column,omitempty"`
}

func (x *ScalarCalculate) Reset() {
	*x = ScalarCalculate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_substrait_extension_promql_plan_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScalarCalculate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScalarCalculate) ProtoMessage() {}

func (x *ScalarCalculate) ProtoReflect() protoreflect.Message {
	mi := &file_substrait_extension_promql_plan_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScalarCalculate.ProtoReflect.Descriptor instead.
func (*ScalarCalculate) Descriptor() ([]byte, []int) {
	return file_substrait_extension_promql_plan_proto_rawDescGZIP(), []int{5}
}

func (x *ScalarCalculate) GetStart() int64 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *ScalarCalculate) GetEnd() int64 {
	if x != nil {
		return x.End
	}
	return 0
}

func (x *ScalarCalculate) GetInterval() int64 {
	if x != nil {
		return x.Interval
	}
	return 0
}

func (x *ScalarCalculate) GetTimeIndex() string {
	if x != nil {
		return x.TimeIndex
	}
	return ""
}

func (x *ScalarCalculate) GetTagColumns() []string {
	if x != nil {
		return x.TagColumns
	}
	return nil
}

func (x *ScalarCalculate) GetFieldColumn() string {
	if x != nil {
		return x.FieldColumn
	}
	return ""
}

var File_substrait_extension_promql_plan_proto protoreflect.FileDescriptor

var file_substrait_extension_promql_plan_proto_rawDesc = []byte{
	0x0a, 0x25, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x61, 0x69, 0x74, 0x5f, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x71, 0x6c, 0x5f, 0x70, 0x6c, 0x61,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x61,
	0x69, 0x74, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x51, 0x0a, 0x0b,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x12, 0x14, 0x0a, 0x05, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03,
	0x65, 0x6e, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x22,
	0xbe, 0x01, 0x0a, 0x11, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x4d, 0x61, 0x6e, 0x69, 0x70,
	0x75, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x65,
	0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x25, 0x0a, 0x0e, 0x6c, 0x6f, 0x6f,
	0x6b, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0d, 0x6c, 0x6f, 0x6f, 0x6b, 0x62, 0x61, 0x63, 0x6b, 0x44, 0x65, 0x6c, 0x74, 0x61,
	0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12,
	0x1f, 0x0a, 0x0b, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x22, 0x88, 0x01, 0x0a, 0x0f, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x4e, 0x6f, 0x72, 0x6d, 0x61,
	0x6c, 0x69, 0x7a, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x74, 0x69, 0x6d, 0x65, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1d, 0x0a, 0x0a, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x09, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x61,
	0x67, 0x5f, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0a, 0x74, 0x61, 0x67, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x22, 0x2f, 0x0a, 0x0c, 0x53,
	0x65, 0x72, 0x69, 0x65, 0x73, 0x44, 0x69, 0x76, 0x69, 0x64, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x74,
	0x61, 0x67, 0x5f, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0a, 0x74, 0x61, 0x67, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x22, 0xab, 0x01, 0x0a,
	0x0f, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x4d, 0x61, 0x6e, 0x69, 0x70, 0x75, 0x6c, 0x61, 0x74, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x76, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x76, 0x61, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x69,
	0x6d, 0x65, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x61, 0x67,
	0x5f, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a,
	0x74, 0x61, 0x67, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x22, 0xb8, 0x01, 0x0a, 0x0f, 0x53,
	0x63, 0x61, 0x6c, 0x61, 0x72, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76,
	0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76,
	0x61, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x61, 0x67, 0x5f, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x61, 0x67, 0x43, 0x6f, 0x6c, 0x75, 0x6d,
	0x6e, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x63, 0x6f, 0x6c, 0x75,
	0x6d, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x43,
	0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x47, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x54, 0x65, 0x61, 0x6d,
	0x2f, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x6f, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x61, 0x69, 0x74, 0x5f, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_substrait_extension_promql_plan_proto_rawDescOnce sync.Once
	file_substrait_extension_promql_plan_proto_rawDescData = file_substrait_extension_promql_plan_proto_rawDesc
)

func file_substrait_extension_promql_plan_proto_rawDescGZIP() []byte {
	file_substrait_extension_promql_plan_proto_rawDescOnce.Do(func() {
		file_substrait_extension_promql_plan_proto_rawDescData = protoimpl.X.CompressGZIP(file_substrait_extension_promql_plan_proto_rawDescData)
	})
	return file_substrait_extension_promql_plan_proto_rawDescData
}

var file_substrait_extension_promql_plan_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_substrait_extension_promql_plan_proto_goTypes = []interface{}{
	(*EmptyMetric)(nil),       // 0: substrait_extension.EmptyMetric
	(*InstantManipulate)(nil), // 1: substrait_extension.InstantManipulate
	(*SeriesNormalize)(nil),   // 2: substrait_extension.SeriesNormalize
	(*SeriesDivide)(nil),      // 3: substrait_extension.SeriesDivide
	(*RangeManipulate)(nil),   // 4: substrait_extension.RangeManipulate
	(*ScalarCalculate)(nil),   // 5: substrait_extension.ScalarCalculate
}
var file_substrait_extension_promql_plan_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_substrait_extension_promql_plan_proto_init() }
func file_substrait_extension_promql_plan_proto_init() {
	if File_substrait_extension_promql_plan_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_substrait_extension_promql_plan_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyMetric); i {
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
		file_substrait_extension_promql_plan_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstantManipulate); i {
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
		file_substrait_extension_promql_plan_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SeriesNormalize); i {
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
		file_substrait_extension_promql_plan_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SeriesDivide); i {
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
		file_substrait_extension_promql_plan_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RangeManipulate); i {
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
		file_substrait_extension_promql_plan_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScalarCalculate); i {
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
			RawDescriptor: file_substrait_extension_promql_plan_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_substrait_extension_promql_plan_proto_goTypes,
		DependencyIndexes: file_substrait_extension_promql_plan_proto_depIdxs,
		MessageInfos:      file_substrait_extension_promql_plan_proto_msgTypes,
	}.Build()
	File_substrait_extension_promql_plan_proto = out.File
	file_substrait_extension_promql_plan_proto_rawDesc = nil
	file_substrait_extension_promql_plan_proto_goTypes = nil
	file_substrait_extension_promql_plan_proto_depIdxs = nil
}
