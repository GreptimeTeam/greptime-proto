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
// source: greptime/v1/index/inverted_index.proto

package index

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

// BitmapType defines the type of bitmap used for indexing.
type BitmapType int32

const (
	BitmapType_BitVec  BitmapType = 0
	BitmapType_Roaring BitmapType = 1
)

// Enum value maps for BitmapType.
var (
	BitmapType_name = map[int32]string{
		0: "BitVec",
		1: "Roaring",
	}
	BitmapType_value = map[string]int32{
		"BitVec":  0,
		"Roaring": 1,
	}
)

func (x BitmapType) Enum() *BitmapType {
	p := new(BitmapType)
	*p = x
	return p
}

func (x BitmapType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BitmapType) Descriptor() protoreflect.EnumDescriptor {
	return file_greptime_v1_index_inverted_index_proto_enumTypes[0].Descriptor()
}

func (BitmapType) Type() protoreflect.EnumType {
	return &file_greptime_v1_index_inverted_index_proto_enumTypes[0]
}

func (x BitmapType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BitmapType.Descriptor instead.
func (BitmapType) EnumDescriptor() ([]byte, []int) {
	return file_greptime_v1_index_inverted_index_proto_rawDescGZIP(), []int{0}
}

// InvertedIndexMetas defines the metadata for an inverted index
// within an inverted index blob.
type InvertedIndexMetas struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A map of tag names to their respective metadata corresponding to an individual
	// inverted index within the inverted index blob.
	Metas map[string]*InvertedIndexMeta `protobuf:"bytes,1,rep,name=metas,proto3" json:"metas,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// The total count of rows within the inverted index blob.
	// This is used to determine the number of segments within the bitmap.
	TotalRowCount uint64 `protobuf:"varint,2,opt,name=total_row_count,json=totalRowCount,proto3" json:"total_row_count,omitempty"`
	// The number of rows per group for bitmap indexing which determines how rows are
	// batched for indexing. Each batch corresponds to a segment in the bitmap and allows
	// for efficient retrieval during queries by reducing the search space.
	SegmentRowCount uint64 `protobuf:"varint,3,opt,name=segment_row_count,json=segmentRowCount,proto3" json:"segment_row_count,omitempty"`
}

func (x *InvertedIndexMetas) Reset() {
	*x = InvertedIndexMetas{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greptime_v1_index_inverted_index_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvertedIndexMetas) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvertedIndexMetas) ProtoMessage() {}

func (x *InvertedIndexMetas) ProtoReflect() protoreflect.Message {
	mi := &file_greptime_v1_index_inverted_index_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvertedIndexMetas.ProtoReflect.Descriptor instead.
func (*InvertedIndexMetas) Descriptor() ([]byte, []int) {
	return file_greptime_v1_index_inverted_index_proto_rawDescGZIP(), []int{0}
}

func (x *InvertedIndexMetas) GetMetas() map[string]*InvertedIndexMeta {
	if x != nil {
		return x.Metas
	}
	return nil
}

func (x *InvertedIndexMetas) GetTotalRowCount() uint64 {
	if x != nil {
		return x.TotalRowCount
	}
	return 0
}

func (x *InvertedIndexMetas) GetSegmentRowCount() uint64 {
	if x != nil {
		return x.SegmentRowCount
	}
	return 0
}

// InvertedIndexMeta contains the metadata for a specific tag's inverted index.
type InvertedIndexMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the tag associated with the inverted index.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The base byte offset for this tag's inverted index data within the blob.
	// Other offsets in this message are relative to this base offset.
	BaseOffset uint64 `protobuf:"varint,2,opt,name=base_offset,json=baseOffset,proto3" json:"base_offset,omitempty"`
	// The total size in bytes of this tag's inverted index data, including bitmaps,
	// FST data.
	InvertedIndexSize uint64 `protobuf:"varint,3,opt,name=inverted_index_size,json=invertedIndexSize,proto3" json:"inverted_index_size,omitempty"`
	// The byte offset of the Finite State Transducer (FST) data relative to the `base_offset`.
	RelativeFstOffset uint32 `protobuf:"varint,4,opt,name=relative_fst_offset,json=relativeFstOffset,proto3" json:"relative_fst_offset,omitempty"`
	// The size in bytes of the FST data.
	FstSize uint32 `protobuf:"varint,5,opt,name=fst_size,json=fstSize,proto3" json:"fst_size,omitempty"`
	// The byte offset relative to the `base_offset` where the null bitmap for this tag
	// starts.
	RelativeNullBitmapOffset uint32 `protobuf:"varint,6,opt,name=relative_null_bitmap_offset,json=relativeNullBitmapOffset,proto3" json:"relative_null_bitmap_offset,omitempty"`
	// The size in bytes of the null bitmap.
	NullBitmapSize uint32 `protobuf:"varint,7,opt,name=null_bitmap_size,json=nullBitmapSize,proto3" json:"null_bitmap_size,omitempty"`
	// Statistical information about the tag's inverted index.
	Stats *InvertedIndexStats `protobuf:"bytes,8,opt,name=stats,proto3" json:"stats,omitempty"`
	// The type of bitmap used for indexing.
	BitmapType BitmapType `protobuf:"varint,9,opt,name=bitmap_type,json=bitmapType,proto3,enum=greptime.v1.index.BitmapType" json:"bitmap_type,omitempty"`
}

func (x *InvertedIndexMeta) Reset() {
	*x = InvertedIndexMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greptime_v1_index_inverted_index_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvertedIndexMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvertedIndexMeta) ProtoMessage() {}

func (x *InvertedIndexMeta) ProtoReflect() protoreflect.Message {
	mi := &file_greptime_v1_index_inverted_index_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvertedIndexMeta.ProtoReflect.Descriptor instead.
func (*InvertedIndexMeta) Descriptor() ([]byte, []int) {
	return file_greptime_v1_index_inverted_index_proto_rawDescGZIP(), []int{1}
}

func (x *InvertedIndexMeta) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *InvertedIndexMeta) GetBaseOffset() uint64 {
	if x != nil {
		return x.BaseOffset
	}
	return 0
}

func (x *InvertedIndexMeta) GetInvertedIndexSize() uint64 {
	if x != nil {
		return x.InvertedIndexSize
	}
	return 0
}

func (x *InvertedIndexMeta) GetRelativeFstOffset() uint32 {
	if x != nil {
		return x.RelativeFstOffset
	}
	return 0
}

func (x *InvertedIndexMeta) GetFstSize() uint32 {
	if x != nil {
		return x.FstSize
	}
	return 0
}

func (x *InvertedIndexMeta) GetRelativeNullBitmapOffset() uint32 {
	if x != nil {
		return x.RelativeNullBitmapOffset
	}
	return 0
}

func (x *InvertedIndexMeta) GetNullBitmapSize() uint32 {
	if x != nil {
		return x.NullBitmapSize
	}
	return 0
}

func (x *InvertedIndexMeta) GetStats() *InvertedIndexStats {
	if x != nil {
		return x.Stats
	}
	return nil
}

func (x *InvertedIndexMeta) GetBitmapType() BitmapType {
	if x != nil {
		return x.BitmapType
	}
	return BitmapType_BitVec
}

// InvertedIndexStats provides statistical data on a tag's inverted index.
type InvertedIndexStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The count of null entries within the tag's column.
	NullCount uint64 `protobuf:"varint,1,opt,name=null_count,json=nullCount,proto3" json:"null_count,omitempty"`
	// The number of distinct values within the tag's column.
	DistinctCount uint64 `protobuf:"varint,2,opt,name=distinct_count,json=distinctCount,proto3" json:"distinct_count,omitempty"`
	// The minimum value found within the tag's column, encoded as bytes.
	MinValue []byte `protobuf:"bytes,3,opt,name=min_value,json=minValue,proto3" json:"min_value,omitempty"`
	// The maximum value found within the tag's column, encoded as bytes.
	MaxValue []byte `protobuf:"bytes,4,opt,name=max_value,json=maxValue,proto3" json:"max_value,omitempty"`
}

func (x *InvertedIndexStats) Reset() {
	*x = InvertedIndexStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greptime_v1_index_inverted_index_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvertedIndexStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvertedIndexStats) ProtoMessage() {}

func (x *InvertedIndexStats) ProtoReflect() protoreflect.Message {
	mi := &file_greptime_v1_index_inverted_index_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvertedIndexStats.ProtoReflect.Descriptor instead.
func (*InvertedIndexStats) Descriptor() ([]byte, []int) {
	return file_greptime_v1_index_inverted_index_proto_rawDescGZIP(), []int{2}
}

func (x *InvertedIndexStats) GetNullCount() uint64 {
	if x != nil {
		return x.NullCount
	}
	return 0
}

func (x *InvertedIndexStats) GetDistinctCount() uint64 {
	if x != nil {
		return x.DistinctCount
	}
	return 0
}

func (x *InvertedIndexStats) GetMinValue() []byte {
	if x != nil {
		return x.MinValue
	}
	return nil
}

func (x *InvertedIndexStats) GetMaxValue() []byte {
	if x != nil {
		return x.MaxValue
	}
	return nil
}

var File_greptime_v1_index_inverted_index_proto protoreflect.FileDescriptor

var file_greptime_v1_index_inverted_index_proto_rawDesc = []byte{
	0x0a, 0x26, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69,
	0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x90, 0x02, 0x0a, 0x12,
	0x49, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x64, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x4d, 0x65, 0x74,
	0x61, 0x73, 0x12, 0x46, 0x0a, 0x05, 0x6d, 0x65, 0x74, 0x61, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x30, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x49, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x64, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x4d, 0x65, 0x74, 0x61, 0x73, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x05, 0x6d, 0x65, 0x74, 0x61, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x5f, 0x72, 0x6f, 0x77, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x6f, 0x77, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x11, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x6f,
	0x77, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0f, 0x73,
	0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x6f, 0x77, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x5e,
	0x0a, 0x0a, 0x4d, 0x65, 0x74, 0x61, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x3a,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e,
	0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x2e, 0x49, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x64, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x4d,
	0x65, 0x74, 0x61, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xa9,
	0x03, 0x0a, 0x11, 0x49, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x64, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x4d, 0x65, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x61, 0x73, 0x65,
	0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x62,
	0x61, 0x73, 0x65, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x2e, 0x0a, 0x13, 0x69, 0x6e, 0x76,
	0x65, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x5f, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x11, 0x69, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x64,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x72, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x66, 0x73, 0x74, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x11, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65,
	0x46, 0x73, 0x74, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x73, 0x74,
	0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x66, 0x73, 0x74,
	0x53, 0x69, 0x7a, 0x65, 0x12, 0x3d, 0x0a, 0x1b, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65,
	0x5f, 0x6e, 0x75, 0x6c, 0x6c, 0x5f, 0x62, 0x69, 0x74, 0x6d, 0x61, 0x70, 0x5f, 0x6f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x18, 0x72, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x76, 0x65, 0x4e, 0x75, 0x6c, 0x6c, 0x42, 0x69, 0x74, 0x6d, 0x61, 0x70, 0x4f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x6e, 0x75, 0x6c, 0x6c, 0x5f, 0x62, 0x69, 0x74, 0x6d,
	0x61, 0x70, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x6e,
	0x75, 0x6c, 0x6c, 0x42, 0x69, 0x74, 0x6d, 0x61, 0x70, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x3b, 0x0a,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x67,
	0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x2e, 0x49, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x64, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x53, 0x74,
	0x61, 0x74, 0x73, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x12, 0x3e, 0x0a, 0x0b, 0x62, 0x69,
	0x74, 0x6d, 0x61, 0x70, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1d, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x2e, 0x42, 0x69, 0x74, 0x6d, 0x61, 0x70, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a,
	0x62, 0x69, 0x74, 0x6d, 0x61, 0x70, 0x54, 0x79, 0x70, 0x65, 0x22, 0x94, 0x01, 0x0a, 0x12, 0x49,
	0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x64, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x53, 0x74, 0x61, 0x74,
	0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x75, 0x6c, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x6e, 0x75, 0x6c, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x25, 0x0a, 0x0e, 0x64, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x63, 0x74, 0x5f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x64, 0x69, 0x73, 0x74, 0x69, 0x6e,
	0x63, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x69, 0x6e, 0x5f, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6d, 0x69, 0x6e, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x78, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6d, 0x61, 0x78, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x2a, 0x25, 0x0a, 0x0a, 0x42, 0x69, 0x74, 0x6d, 0x61, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x0a, 0x0a, 0x06, 0x42, 0x69, 0x74, 0x56, 0x65, 0x63, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x52,
	0x6f, 0x61, 0x72, 0x69, 0x6e, 0x67, 0x10, 0x01, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x47, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x54,
	0x65, 0x61, 0x6d, 0x2f, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2d, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_greptime_v1_index_inverted_index_proto_rawDescOnce sync.Once
	file_greptime_v1_index_inverted_index_proto_rawDescData = file_greptime_v1_index_inverted_index_proto_rawDesc
)

func file_greptime_v1_index_inverted_index_proto_rawDescGZIP() []byte {
	file_greptime_v1_index_inverted_index_proto_rawDescOnce.Do(func() {
		file_greptime_v1_index_inverted_index_proto_rawDescData = protoimpl.X.CompressGZIP(file_greptime_v1_index_inverted_index_proto_rawDescData)
	})
	return file_greptime_v1_index_inverted_index_proto_rawDescData
}

var file_greptime_v1_index_inverted_index_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_greptime_v1_index_inverted_index_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_greptime_v1_index_inverted_index_proto_goTypes = []interface{}{
	(BitmapType)(0),            // 0: greptime.v1.index.BitmapType
	(*InvertedIndexMetas)(nil), // 1: greptime.v1.index.InvertedIndexMetas
	(*InvertedIndexMeta)(nil),  // 2: greptime.v1.index.InvertedIndexMeta
	(*InvertedIndexStats)(nil), // 3: greptime.v1.index.InvertedIndexStats
	nil,                        // 4: greptime.v1.index.InvertedIndexMetas.MetasEntry
}
var file_greptime_v1_index_inverted_index_proto_depIdxs = []int32{
	4, // 0: greptime.v1.index.InvertedIndexMetas.metas:type_name -> greptime.v1.index.InvertedIndexMetas.MetasEntry
	3, // 1: greptime.v1.index.InvertedIndexMeta.stats:type_name -> greptime.v1.index.InvertedIndexStats
	0, // 2: greptime.v1.index.InvertedIndexMeta.bitmap_type:type_name -> greptime.v1.index.BitmapType
	2, // 3: greptime.v1.index.InvertedIndexMetas.MetasEntry.value:type_name -> greptime.v1.index.InvertedIndexMeta
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_greptime_v1_index_inverted_index_proto_init() }
func file_greptime_v1_index_inverted_index_proto_init() {
	if File_greptime_v1_index_inverted_index_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_greptime_v1_index_inverted_index_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InvertedIndexMetas); i {
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
		file_greptime_v1_index_inverted_index_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InvertedIndexMeta); i {
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
		file_greptime_v1_index_inverted_index_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InvertedIndexStats); i {
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
			RawDescriptor: file_greptime_v1_index_inverted_index_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_greptime_v1_index_inverted_index_proto_goTypes,
		DependencyIndexes: file_greptime_v1_index_inverted_index_proto_depIdxs,
		EnumInfos:         file_greptime_v1_index_inverted_index_proto_enumTypes,
		MessageInfos:      file_greptime_v1_index_inverted_index_proto_msgTypes,
	}.Build()
	File_greptime_v1_index_inverted_index_proto = out.File
	file_greptime_v1_index_inverted_index_proto_rawDesc = nil
	file_greptime_v1_index_inverted_index_proto_goTypes = nil
	file_greptime_v1_index_inverted_index_proto_depIdxs = nil
}
