// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: greptime/v1/meta/procedure.proto

#include "greptime/v1/meta/procedure.pb.h"

#include <algorithm>

#include <google/protobuf/io/coded_stream.h>
#include <google/protobuf/extension_set.h>
#include <google/protobuf/wire_format_lite.h>
#include <google/protobuf/descriptor.h>
#include <google/protobuf/generated_message_reflection.h>
#include <google/protobuf/reflection_ops.h>
#include <google/protobuf/wire_format.h>
// @@protoc_insertion_point(includes)
#include <google/protobuf/port_def.inc>

PROTOBUF_PRAGMA_INIT_SEG

namespace _pb = ::PROTOBUF_NAMESPACE_ID;
namespace _pbi = _pb::internal;

namespace greptime {
namespace v1 {
namespace meta {
PROTOBUF_CONSTEXPR ProcedureState::ProcedureState(
    ::_pbi::ConstantInitialized): _impl_{
    /*decltype(_impl_.status_)*/0
  , /*decltype(_impl_._cached_size_)*/{}} {}
struct ProcedureStateDefaultTypeInternal {
  PROTOBUF_CONSTEXPR ProcedureStateDefaultTypeInternal()
      : _instance(::_pbi::ConstantInitialized{}) {}
  ~ProcedureStateDefaultTypeInternal() {}
  union {
    ProcedureState _instance;
  };
};
PROTOBUF_ATTRIBUTE_NO_DESTROY PROTOBUF_CONSTINIT PROTOBUF_ATTRIBUTE_INIT_PRIORITY1 ProcedureStateDefaultTypeInternal _ProcedureState_default_instance_;
}  // namespace meta
}  // namespace v1
}  // namespace greptime
static ::_pb::Metadata file_level_metadata_greptime_2fv1_2fmeta_2fprocedure_2eproto[1];
static const ::_pb::EnumDescriptor* file_level_enum_descriptors_greptime_2fv1_2fmeta_2fprocedure_2eproto[1];
static constexpr ::_pb::ServiceDescriptor const** file_level_service_descriptors_greptime_2fv1_2fmeta_2fprocedure_2eproto = nullptr;

const uint32_t TableStruct_greptime_2fv1_2fmeta_2fprocedure_2eproto::offsets[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) = {
  ~0u,  // no _has_bits_
  PROTOBUF_FIELD_OFFSET(::greptime::v1::meta::ProcedureState, _internal_metadata_),
  ~0u,  // no _extensions_
  ~0u,  // no _oneof_case_
  ~0u,  // no _weak_field_map_
  ~0u,  // no _inlined_string_donated_
  PROTOBUF_FIELD_OFFSET(::greptime::v1::meta::ProcedureState, _impl_.status_),
};
static const ::_pbi::MigrationSchema schemas[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) = {
  { 0, -1, -1, sizeof(::greptime::v1::meta::ProcedureState)},
};

static const ::_pb::Message* const file_default_instances[] = {
  &::greptime::v1::meta::_ProcedureState_default_instance_._instance,
};

const char descriptor_table_protodef_greptime_2fv1_2fmeta_2fprocedure_2eproto[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) =
  "\n greptime/v1/meta/procedure.proto\022\020grep"
  "time.v1.meta\032\035greptime/v1/meta/common.pr"
  "oto\032\032greptime/v1/meta/ddl.proto\032\035greptim"
  "e/v1/meta/region.proto\"C\n\016ProcedureState"
  "\0221\n\006status\030\001 \001(\0162!.greptime.v1.meta.Proc"
  "edureStatus*9\n\017ProcedureStatus\022\r\n\tExecut"
  "ing\020\000\022\r\n\tSuspended\020\001\022\010\n\004Done\020\0022\375\001\n\tProce"
  "dure\022H\n\005query\022\035.greptime.v1.meta.Procedu"
  "reId\032 .greptime.v1.meta.ProcedureState\022J"
  "\n\003ddl\022 .greptime.v1.meta.DdlTaskRequest\032"
  "!.greptime.v1.meta.DdlTaskResponse\022Z\n\007mi"
  "grate\022&.greptime.v1.meta.MigrateRegionRe"
  "quest\032\'.greptime.v1.meta.MigrateRegionRe"
  "sponseB<Z:github.com/GreptimeTeam/grepti"
  "me-proto/go/greptime/v1/metab\006proto3"
  ;
static const ::_pbi::DescriptorTable* const descriptor_table_greptime_2fv1_2fmeta_2fprocedure_2eproto_deps[3] = {
  &::descriptor_table_greptime_2fv1_2fmeta_2fcommon_2eproto,
  &::descriptor_table_greptime_2fv1_2fmeta_2fddl_2eproto,
  &::descriptor_table_greptime_2fv1_2fmeta_2fregion_2eproto,
};
static ::_pbi::once_flag descriptor_table_greptime_2fv1_2fmeta_2fprocedure_2eproto_once;
const ::_pbi::DescriptorTable descriptor_table_greptime_2fv1_2fmeta_2fprocedure_2eproto = {
    false, false, 596, descriptor_table_protodef_greptime_2fv1_2fmeta_2fprocedure_2eproto,
    "greptime/v1/meta/procedure.proto",
    &descriptor_table_greptime_2fv1_2fmeta_2fprocedure_2eproto_once, descriptor_table_greptime_2fv1_2fmeta_2fprocedure_2eproto_deps, 3, 1,
    schemas, file_default_instances, TableStruct_greptime_2fv1_2fmeta_2fprocedure_2eproto::offsets,
    file_level_metadata_greptime_2fv1_2fmeta_2fprocedure_2eproto, file_level_enum_descriptors_greptime_2fv1_2fmeta_2fprocedure_2eproto,
    file_level_service_descriptors_greptime_2fv1_2fmeta_2fprocedure_2eproto,
};
PROTOBUF_ATTRIBUTE_WEAK const ::_pbi::DescriptorTable* descriptor_table_greptime_2fv1_2fmeta_2fprocedure_2eproto_getter() {
  return &descriptor_table_greptime_2fv1_2fmeta_2fprocedure_2eproto;
}

// Force running AddDescriptors() at dynamic initialization time.
PROTOBUF_ATTRIBUTE_INIT_PRIORITY2 static ::_pbi::AddDescriptorsRunner dynamic_init_dummy_greptime_2fv1_2fmeta_2fprocedure_2eproto(&descriptor_table_greptime_2fv1_2fmeta_2fprocedure_2eproto);
namespace greptime {
namespace v1 {
namespace meta {
const ::PROTOBUF_NAMESPACE_ID::EnumDescriptor* ProcedureStatus_descriptor() {
  ::PROTOBUF_NAMESPACE_ID::internal::AssignDescriptors(&descriptor_table_greptime_2fv1_2fmeta_2fprocedure_2eproto);
  return file_level_enum_descriptors_greptime_2fv1_2fmeta_2fprocedure_2eproto[0];
}
bool ProcedureStatus_IsValid(int value) {
  switch (value) {
    case 0:
    case 1:
    case 2:
      return true;
    default:
      return false;
  }
}


// ===================================================================

class ProcedureState::_Internal {
 public:
};

ProcedureState::ProcedureState(::PROTOBUF_NAMESPACE_ID::Arena* arena,
                         bool is_message_owned)
  : ::PROTOBUF_NAMESPACE_ID::Message(arena, is_message_owned) {
  SharedCtor(arena, is_message_owned);
  // @@protoc_insertion_point(arena_constructor:greptime.v1.meta.ProcedureState)
}
ProcedureState::ProcedureState(const ProcedureState& from)
  : ::PROTOBUF_NAMESPACE_ID::Message() {
  ProcedureState* const _this = this; (void)_this;
  new (&_impl_) Impl_{
      decltype(_impl_.status_){}
    , /*decltype(_impl_._cached_size_)*/{}};

  _internal_metadata_.MergeFrom<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(from._internal_metadata_);
  _this->_impl_.status_ = from._impl_.status_;
  // @@protoc_insertion_point(copy_constructor:greptime.v1.meta.ProcedureState)
}

inline void ProcedureState::SharedCtor(
    ::_pb::Arena* arena, bool is_message_owned) {
  (void)arena;
  (void)is_message_owned;
  new (&_impl_) Impl_{
      decltype(_impl_.status_){0}
    , /*decltype(_impl_._cached_size_)*/{}
  };
}

ProcedureState::~ProcedureState() {
  // @@protoc_insertion_point(destructor:greptime.v1.meta.ProcedureState)
  if (auto *arena = _internal_metadata_.DeleteReturnArena<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>()) {
  (void)arena;
    return;
  }
  SharedDtor();
}

inline void ProcedureState::SharedDtor() {
  GOOGLE_DCHECK(GetArenaForAllocation() == nullptr);
}

void ProcedureState::SetCachedSize(int size) const {
  _impl_._cached_size_.Set(size);
}

void ProcedureState::Clear() {
// @@protoc_insertion_point(message_clear_start:greptime.v1.meta.ProcedureState)
  uint32_t cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  _impl_.status_ = 0;
  _internal_metadata_.Clear<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>();
}

const char* ProcedureState::_InternalParse(const char* ptr, ::_pbi::ParseContext* ctx) {
#define CHK_(x) if (PROTOBUF_PREDICT_FALSE(!(x))) goto failure
  while (!ctx->Done(&ptr)) {
    uint32_t tag;
    ptr = ::_pbi::ReadTag(ptr, &tag);
    switch (tag >> 3) {
      // .greptime.v1.meta.ProcedureStatus status = 1;
      case 1:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 8)) {
          uint64_t val = ::PROTOBUF_NAMESPACE_ID::internal::ReadVarint64(&ptr);
          CHK_(ptr);
          _internal_set_status(static_cast<::greptime::v1::meta::ProcedureStatus>(val));
        } else
          goto handle_unusual;
        continue;
      default:
        goto handle_unusual;
    }  // switch
  handle_unusual:
    if ((tag == 0) || ((tag & 7) == 4)) {
      CHK_(ptr);
      ctx->SetLastTag(tag);
      goto message_done;
    }
    ptr = UnknownFieldParse(
        tag,
        _internal_metadata_.mutable_unknown_fields<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(),
        ptr, ctx);
    CHK_(ptr != nullptr);
  }  // while
message_done:
  return ptr;
failure:
  ptr = nullptr;
  goto message_done;
#undef CHK_
}

uint8_t* ProcedureState::_InternalSerialize(
    uint8_t* target, ::PROTOBUF_NAMESPACE_ID::io::EpsCopyOutputStream* stream) const {
  // @@protoc_insertion_point(serialize_to_array_start:greptime.v1.meta.ProcedureState)
  uint32_t cached_has_bits = 0;
  (void) cached_has_bits;

  // .greptime.v1.meta.ProcedureStatus status = 1;
  if (this->_internal_status() != 0) {
    target = stream->EnsureSpace(target);
    target = ::_pbi::WireFormatLite::WriteEnumToArray(
      1, this->_internal_status(), target);
  }

  if (PROTOBUF_PREDICT_FALSE(_internal_metadata_.have_unknown_fields())) {
    target = ::_pbi::WireFormat::InternalSerializeUnknownFieldsToArray(
        _internal_metadata_.unknown_fields<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(::PROTOBUF_NAMESPACE_ID::UnknownFieldSet::default_instance), target, stream);
  }
  // @@protoc_insertion_point(serialize_to_array_end:greptime.v1.meta.ProcedureState)
  return target;
}

size_t ProcedureState::ByteSizeLong() const {
// @@protoc_insertion_point(message_byte_size_start:greptime.v1.meta.ProcedureState)
  size_t total_size = 0;

  uint32_t cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  // .greptime.v1.meta.ProcedureStatus status = 1;
  if (this->_internal_status() != 0) {
    total_size += 1 +
      ::_pbi::WireFormatLite::EnumSize(this->_internal_status());
  }

  return MaybeComputeUnknownFieldsSize(total_size, &_impl_._cached_size_);
}

const ::PROTOBUF_NAMESPACE_ID::Message::ClassData ProcedureState::_class_data_ = {
    ::PROTOBUF_NAMESPACE_ID::Message::CopyWithSourceCheck,
    ProcedureState::MergeImpl
};
const ::PROTOBUF_NAMESPACE_ID::Message::ClassData*ProcedureState::GetClassData() const { return &_class_data_; }


void ProcedureState::MergeImpl(::PROTOBUF_NAMESPACE_ID::Message& to_msg, const ::PROTOBUF_NAMESPACE_ID::Message& from_msg) {
  auto* const _this = static_cast<ProcedureState*>(&to_msg);
  auto& from = static_cast<const ProcedureState&>(from_msg);
  // @@protoc_insertion_point(class_specific_merge_from_start:greptime.v1.meta.ProcedureState)
  GOOGLE_DCHECK_NE(&from, _this);
  uint32_t cached_has_bits = 0;
  (void) cached_has_bits;

  if (from._internal_status() != 0) {
    _this->_internal_set_status(from._internal_status());
  }
  _this->_internal_metadata_.MergeFrom<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(from._internal_metadata_);
}

void ProcedureState::CopyFrom(const ProcedureState& from) {
// @@protoc_insertion_point(class_specific_copy_from_start:greptime.v1.meta.ProcedureState)
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

bool ProcedureState::IsInitialized() const {
  return true;
}

void ProcedureState::InternalSwap(ProcedureState* other) {
  using std::swap;
  _internal_metadata_.InternalSwap(&other->_internal_metadata_);
  swap(_impl_.status_, other->_impl_.status_);
}

::PROTOBUF_NAMESPACE_ID::Metadata ProcedureState::GetMetadata() const {
  return ::_pbi::AssignDescriptors(
      &descriptor_table_greptime_2fv1_2fmeta_2fprocedure_2eproto_getter, &descriptor_table_greptime_2fv1_2fmeta_2fprocedure_2eproto_once,
      file_level_metadata_greptime_2fv1_2fmeta_2fprocedure_2eproto[0]);
}

// @@protoc_insertion_point(namespace_scope)
}  // namespace meta
}  // namespace v1
}  // namespace greptime
PROTOBUF_NAMESPACE_OPEN
template<> PROTOBUF_NOINLINE ::greptime::v1::meta::ProcedureState*
Arena::CreateMaybeMessage< ::greptime::v1::meta::ProcedureState >(Arena* arena) {
  return Arena::CreateMessageInternal< ::greptime::v1::meta::ProcedureState >(arena);
}
PROTOBUF_NAMESPACE_CLOSE

// @@protoc_insertion_point(global_scope)
#include <google/protobuf/port_undef.inc>
