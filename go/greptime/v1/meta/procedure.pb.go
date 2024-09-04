// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: greptime/v1/meta/procedure.proto

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

type ProcedureStatus int32

const (
	ProcedureStatus_Running         ProcedureStatus = 0
	ProcedureStatus_Done            ProcedureStatus = 1
	ProcedureStatus_Retrying        ProcedureStatus = 2
	ProcedureStatus_Failed          ProcedureStatus = 3
	ProcedureStatus_PrepareRollback ProcedureStatus = 4
	ProcedureStatus_RollingBack     ProcedureStatus = 5
)

// Enum value maps for ProcedureStatus.
var (
	ProcedureStatus_name = map[int32]string{
		0: "Running",
		1: "Done",
		2: "Retrying",
		3: "Failed",
		4: "PrepareRollback",
		5: "RollingBack",
	}
	ProcedureStatus_value = map[string]int32{
		"Running":         0,
		"Done":            1,
		"Retrying":        2,
		"Failed":          3,
		"PrepareRollback": 4,
		"RollingBack":     5,
	}
)

func (x ProcedureStatus) Enum() *ProcedureStatus {
	p := new(ProcedureStatus)
	*p = x
	return p
}

func (x ProcedureStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProcedureStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_greptime_v1_meta_procedure_proto_enumTypes[0].Descriptor()
}

func (ProcedureStatus) Type() protoreflect.EnumType {
	return &file_greptime_v1_meta_procedure_proto_enumTypes[0]
}

func (x ProcedureStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProcedureStatus.Descriptor instead.
func (ProcedureStatus) EnumDescriptor() ([]byte, []int) {
	return file_greptime_v1_meta_procedure_proto_rawDescGZIP(), []int{0}
}

type ProcedureMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          *ProcedureId    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	TypeName    string          `protobuf:"bytes,2,opt,name=type_name,json=typeName,proto3" json:"type_name,omitempty"`
	Status      ProcedureStatus `protobuf:"varint,3,opt,name=status,proto3,enum=greptime.v1.meta.ProcedureStatus" json:"status,omitempty"`
	StartTimeMs uint64          `protobuf:"varint,4,opt,name=start_time_ms,json=startTimeMs,proto3" json:"start_time_ms,omitempty"`
	EndTimeMs   uint64          `protobuf:"varint,5,opt,name=end_time_ms,json=endTimeMs,proto3" json:"end_time_ms,omitempty"`
	LockKeys    []string        `protobuf:"bytes,6,rep,name=lock_keys,json=lockKeys,proto3" json:"lock_keys,omitempty"`
	Error       string          `protobuf:"bytes,7,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *ProcedureMeta) Reset() {
	*x = ProcedureMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greptime_v1_meta_procedure_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcedureMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcedureMeta) ProtoMessage() {}

func (x *ProcedureMeta) ProtoReflect() protoreflect.Message {
	mi := &file_greptime_v1_meta_procedure_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcedureMeta.ProtoReflect.Descriptor instead.
func (*ProcedureMeta) Descriptor() ([]byte, []int) {
	return file_greptime_v1_meta_procedure_proto_rawDescGZIP(), []int{0}
}

func (x *ProcedureMeta) GetId() *ProcedureId {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *ProcedureMeta) GetTypeName() string {
	if x != nil {
		return x.TypeName
	}
	return ""
}

func (x *ProcedureMeta) GetStatus() ProcedureStatus {
	if x != nil {
		return x.Status
	}
	return ProcedureStatus_Running
}

func (x *ProcedureMeta) GetStartTimeMs() uint64 {
	if x != nil {
		return x.StartTimeMs
	}
	return 0
}

func (x *ProcedureMeta) GetEndTimeMs() uint64 {
	if x != nil {
		return x.EndTimeMs
	}
	return 0
}

func (x *ProcedureMeta) GetLockKeys() []string {
	if x != nil {
		return x.LockKeys
	}
	return nil
}

func (x *ProcedureMeta) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type QueryProcedureRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *RequestHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Pid    *ProcedureId   `protobuf:"bytes,2,opt,name=pid,proto3" json:"pid,omitempty"`
}

func (x *QueryProcedureRequest) Reset() {
	*x = QueryProcedureRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greptime_v1_meta_procedure_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryProcedureRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryProcedureRequest) ProtoMessage() {}

func (x *QueryProcedureRequest) ProtoReflect() protoreflect.Message {
	mi := &file_greptime_v1_meta_procedure_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryProcedureRequest.ProtoReflect.Descriptor instead.
func (*QueryProcedureRequest) Descriptor() ([]byte, []int) {
	return file_greptime_v1_meta_procedure_proto_rawDescGZIP(), []int{1}
}

func (x *QueryProcedureRequest) GetHeader() *RequestHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *QueryProcedureRequest) GetPid() *ProcedureId {
	if x != nil {
		return x.Pid
	}
	return nil
}

type ProcedureStateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Status ProcedureStatus `protobuf:"varint,2,opt,name=status,proto3,enum=greptime.v1.meta.ProcedureStatus" json:"status,omitempty"`
	Error  string          `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *ProcedureStateResponse) Reset() {
	*x = ProcedureStateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greptime_v1_meta_procedure_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcedureStateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcedureStateResponse) ProtoMessage() {}

func (x *ProcedureStateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_greptime_v1_meta_procedure_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcedureStateResponse.ProtoReflect.Descriptor instead.
func (*ProcedureStateResponse) Descriptor() ([]byte, []int) {
	return file_greptime_v1_meta_procedure_proto_rawDescGZIP(), []int{2}
}

func (x *ProcedureStateResponse) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *ProcedureStateResponse) GetStatus() ProcedureStatus {
	if x != nil {
		return x.Status
	}
	return ProcedureStatus_Running
}

func (x *ProcedureStateResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type ProcedureDetailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *RequestHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
}

func (x *ProcedureDetailRequest) Reset() {
	*x = ProcedureDetailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greptime_v1_meta_procedure_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcedureDetailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcedureDetailRequest) ProtoMessage() {}

func (x *ProcedureDetailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_greptime_v1_meta_procedure_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcedureDetailRequest.ProtoReflect.Descriptor instead.
func (*ProcedureDetailRequest) Descriptor() ([]byte, []int) {
	return file_greptime_v1_meta_procedure_proto_rawDescGZIP(), []int{3}
}

func (x *ProcedureDetailRequest) GetHeader() *RequestHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

type ProcedureDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header     *ResponseHeader  `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Procedures []*ProcedureMeta `protobuf:"bytes,2,rep,name=procedures,proto3" json:"procedures,omitempty"`
}

func (x *ProcedureDetailResponse) Reset() {
	*x = ProcedureDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greptime_v1_meta_procedure_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcedureDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcedureDetailResponse) ProtoMessage() {}

func (x *ProcedureDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_greptime_v1_meta_procedure_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcedureDetailResponse.ProtoReflect.Descriptor instead.
func (*ProcedureDetailResponse) Descriptor() ([]byte, []int) {
	return file_greptime_v1_meta_procedure_proto_rawDescGZIP(), []int{4}
}

func (x *ProcedureDetailResponse) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *ProcedureDetailResponse) GetProcedures() []*ProcedureMeta {
	if x != nil {
		return x.Procedures
	}
	return nil
}

var File_greptime_v1_meta_procedure_proto protoreflect.FileDescriptor

var file_greptime_v1_meta_procedure_proto_rawDesc = []byte{
	0x0a, 0x20, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65,
	0x74, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x64, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x10, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x6d, 0x65, 0x74, 0x61, 0x1a, 0x1d, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x64, 0x64, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1d, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74,
	0x61, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8d,
	0x02, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x64, 0x75, 0x72, 0x65, 0x4d, 0x65, 0x74, 0x61,
	0x12, 0x2d, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x67,
	0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e,
	0x50, 0x72, 0x6f, 0x63, 0x65, 0x64, 0x75, 0x72, 0x65, 0x49, 0x64, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x74, 0x79, 0x70, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x67,
	0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e,
	0x50, 0x72, 0x6f, 0x63, 0x65, 0x64, 0x75, 0x72, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x22, 0x0a, 0x0d, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x4d, 0x73, 0x12, 0x1e, 0x0a, 0x0b, 0x65,
	0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6d, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x09, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x4d, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x6c,
	0x6f, 0x63, 0x6b, 0x5f, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08,
	0x6c, 0x6f, 0x63, 0x6b, 0x4b, 0x65, 0x79, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x81,
	0x01, 0x0a, 0x15, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x64, 0x75, 0x72,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74,
	0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x12, 0x2f, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x74,
	0x61, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x64, 0x75, 0x72, 0x65, 0x49, 0x64, 0x52, 0x03, 0x70,
	0x69, 0x64, 0x22, 0xa3, 0x01, 0x0a, 0x16, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x64, 0x75, 0x72, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a,
	0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x74, 0x61,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52,
	0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x39, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69,
	0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65,
	0x64, 0x75, 0x72, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x51, 0x0a, 0x16, 0x50, 0x72, 0x6f, 0x63,
	0x65, 0x64, 0x75, 0x72, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x37, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x22, 0x94, 0x01, 0x0a, 0x17,
	0x50, 0x72, 0x6f, 0x63, 0x65, 0x64, 0x75, 0x72, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69,
	0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x12, 0x3f, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x64, 0x75, 0x72, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x64, 0x75,
	0x72, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x64, 0x75, 0x72,
	0x65, 0x73, 0x2a, 0x68, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x64, 0x75, 0x72, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67,
	0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x6f, 0x6e, 0x65, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08,
	0x52, 0x65, 0x74, 0x72, 0x79, 0x69, 0x6e, 0x67, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x61,
	0x69, 0x6c, 0x65, 0x64, 0x10, 0x03, 0x12, 0x13, 0x0a, 0x0f, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72,
	0x65, 0x52, 0x6f, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x10, 0x04, 0x12, 0x0f, 0x0a, 0x0b, 0x52,
	0x6f, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x42, 0x61, 0x63, 0x6b, 0x10, 0x05, 0x32, 0xf6, 0x02, 0x0a,
	0x10, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x64, 0x75, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x5a, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x27, 0x2e, 0x67, 0x72, 0x65,
	0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x64, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x64, 0x75, 0x72, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a,
	0x03, 0x64, 0x64, 0x6c, 0x12, 0x20, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x44, 0x64, 0x6c, 0x54, 0x61, 0x73, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x44, 0x64, 0x6c, 0x54, 0x61, 0x73,
	0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5a, 0x0a, 0x07, 0x6d, 0x69, 0x67,
	0x72, 0x61, 0x74, 0x65, 0x12, 0x26, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x67,
	0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e,
	0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5e, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x12, 0x28, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d,
	0x65, 0x74, 0x61, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x64, 0x75, 0x72, 0x65, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x67, 0x72, 0x65,
	0x70, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x50, 0x72,
	0x6f, 0x63, 0x65, 0x64, 0x75, 0x72, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x47, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x54, 0x65, 0x61, 0x6d,
	0x2f, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x6f, 0x2f, 0x67, 0x72, 0x65, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d,
	0x65, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_greptime_v1_meta_procedure_proto_rawDescOnce sync.Once
	file_greptime_v1_meta_procedure_proto_rawDescData = file_greptime_v1_meta_procedure_proto_rawDesc
)

func file_greptime_v1_meta_procedure_proto_rawDescGZIP() []byte {
	file_greptime_v1_meta_procedure_proto_rawDescOnce.Do(func() {
		file_greptime_v1_meta_procedure_proto_rawDescData = protoimpl.X.CompressGZIP(file_greptime_v1_meta_procedure_proto_rawDescData)
	})
	return file_greptime_v1_meta_procedure_proto_rawDescData
}

var file_greptime_v1_meta_procedure_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_greptime_v1_meta_procedure_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_greptime_v1_meta_procedure_proto_goTypes = []interface{}{
	(ProcedureStatus)(0),            // 0: greptime.v1.meta.ProcedureStatus
	(*ProcedureMeta)(nil),           // 1: greptime.v1.meta.ProcedureMeta
	(*QueryProcedureRequest)(nil),   // 2: greptime.v1.meta.QueryProcedureRequest
	(*ProcedureStateResponse)(nil),  // 3: greptime.v1.meta.ProcedureStateResponse
	(*ProcedureDetailRequest)(nil),  // 4: greptime.v1.meta.ProcedureDetailRequest
	(*ProcedureDetailResponse)(nil), // 5: greptime.v1.meta.ProcedureDetailResponse
	(*ProcedureId)(nil),             // 6: greptime.v1.meta.ProcedureId
	(*RequestHeader)(nil),           // 7: greptime.v1.meta.RequestHeader
	(*ResponseHeader)(nil),          // 8: greptime.v1.meta.ResponseHeader
	(*DdlTaskRequest)(nil),          // 9: greptime.v1.meta.DdlTaskRequest
	(*MigrateRegionRequest)(nil),    // 10: greptime.v1.meta.MigrateRegionRequest
	(*DdlTaskResponse)(nil),         // 11: greptime.v1.meta.DdlTaskResponse
	(*MigrateRegionResponse)(nil),   // 12: greptime.v1.meta.MigrateRegionResponse
}
var file_greptime_v1_meta_procedure_proto_depIdxs = []int32{
	6,  // 0: greptime.v1.meta.ProcedureMeta.id:type_name -> greptime.v1.meta.ProcedureId
	0,  // 1: greptime.v1.meta.ProcedureMeta.status:type_name -> greptime.v1.meta.ProcedureStatus
	7,  // 2: greptime.v1.meta.QueryProcedureRequest.header:type_name -> greptime.v1.meta.RequestHeader
	6,  // 3: greptime.v1.meta.QueryProcedureRequest.pid:type_name -> greptime.v1.meta.ProcedureId
	8,  // 4: greptime.v1.meta.ProcedureStateResponse.header:type_name -> greptime.v1.meta.ResponseHeader
	0,  // 5: greptime.v1.meta.ProcedureStateResponse.status:type_name -> greptime.v1.meta.ProcedureStatus
	7,  // 6: greptime.v1.meta.ProcedureDetailRequest.header:type_name -> greptime.v1.meta.RequestHeader
	8,  // 7: greptime.v1.meta.ProcedureDetailResponse.header:type_name -> greptime.v1.meta.ResponseHeader
	1,  // 8: greptime.v1.meta.ProcedureDetailResponse.procedures:type_name -> greptime.v1.meta.ProcedureMeta
	2,  // 9: greptime.v1.meta.ProcedureService.query:input_type -> greptime.v1.meta.QueryProcedureRequest
	9,  // 10: greptime.v1.meta.ProcedureService.ddl:input_type -> greptime.v1.meta.DdlTaskRequest
	10, // 11: greptime.v1.meta.ProcedureService.migrate:input_type -> greptime.v1.meta.MigrateRegionRequest
	4,  // 12: greptime.v1.meta.ProcedureService.details:input_type -> greptime.v1.meta.ProcedureDetailRequest
	3,  // 13: greptime.v1.meta.ProcedureService.query:output_type -> greptime.v1.meta.ProcedureStateResponse
	11, // 14: greptime.v1.meta.ProcedureService.ddl:output_type -> greptime.v1.meta.DdlTaskResponse
	12, // 15: greptime.v1.meta.ProcedureService.migrate:output_type -> greptime.v1.meta.MigrateRegionResponse
	5,  // 16: greptime.v1.meta.ProcedureService.details:output_type -> greptime.v1.meta.ProcedureDetailResponse
	13, // [13:17] is the sub-list for method output_type
	9,  // [9:13] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_greptime_v1_meta_procedure_proto_init() }
func file_greptime_v1_meta_procedure_proto_init() {
	if File_greptime_v1_meta_procedure_proto != nil {
		return
	}
	file_greptime_v1_meta_common_proto_init()
	file_greptime_v1_meta_ddl_proto_init()
	file_greptime_v1_meta_region_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_greptime_v1_meta_procedure_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProcedureMeta); i {
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
		file_greptime_v1_meta_procedure_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryProcedureRequest); i {
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
		file_greptime_v1_meta_procedure_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProcedureStateResponse); i {
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
		file_greptime_v1_meta_procedure_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProcedureDetailRequest); i {
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
		file_greptime_v1_meta_procedure_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProcedureDetailResponse); i {
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
			RawDescriptor: file_greptime_v1_meta_procedure_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_greptime_v1_meta_procedure_proto_goTypes,
		DependencyIndexes: file_greptime_v1_meta_procedure_proto_depIdxs,
		EnumInfos:         file_greptime_v1_meta_procedure_proto_enumTypes,
		MessageInfos:      file_greptime_v1_meta_procedure_proto_msgTypes,
	}.Build()
	File_greptime_v1_meta_procedure_proto = out.File
	file_greptime_v1_meta_procedure_proto_rawDesc = nil
	file_greptime_v1_meta_procedure_proto_goTypes = nil
	file_greptime_v1_meta_procedure_proto_depIdxs = nil
}
