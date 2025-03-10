// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: greptime/v1/meta/region.proto

#include "greptime/v1/meta/region.pb.h"

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
PROTOBUF_CONSTEXPR MigrateRegionRequest::MigrateRegionRequest(
    ::_pbi::ConstantInitialized): _impl_{
    /*decltype(_impl_.header_)*/nullptr
  , /*decltype(_impl_.region_id_)*/uint64_t{0u}
  , /*decltype(_impl_.from_peer_)*/uint64_t{0u}
  , /*decltype(_impl_.to_peer_)*/uint64_t{0u}
  , /*decltype(_impl_.timeout_secs_)*/0u
  , /*decltype(_impl_._cached_size_)*/{}} {}
struct MigrateRegionRequestDefaultTypeInternal {
  PROTOBUF_CONSTEXPR MigrateRegionRequestDefaultTypeInternal()
      : _instance(::_pbi::ConstantInitialized{}) {}
  ~MigrateRegionRequestDefaultTypeInternal() {}
  union {
    MigrateRegionRequest _instance;
  };
};
PROTOBUF_ATTRIBUTE_NO_DESTROY PROTOBUF_CONSTINIT PROTOBUF_ATTRIBUTE_INIT_PRIORITY1 MigrateRegionRequestDefaultTypeInternal _MigrateRegionRequest_default_instance_;
PROTOBUF_CONSTEXPR MigrateRegionResponse::MigrateRegionResponse(
    ::_pbi::ConstantInitialized): _impl_{
    /*decltype(_impl_.header_)*/nullptr
  , /*decltype(_impl_.pid_)*/nullptr
  , /*decltype(_impl_._cached_size_)*/{}} {}
struct MigrateRegionResponseDefaultTypeInternal {
  PROTOBUF_CONSTEXPR MigrateRegionResponseDefaultTypeInternal()
      : _instance(::_pbi::ConstantInitialized{}) {}
  ~MigrateRegionResponseDefaultTypeInternal() {}
  union {
    MigrateRegionResponse _instance;
  };
};
PROTOBUF_ATTRIBUTE_NO_DESTROY PROTOBUF_CONSTINIT PROTOBUF_ATTRIBUTE_INIT_PRIORITY1 MigrateRegionResponseDefaultTypeInternal _MigrateRegionResponse_default_instance_;
}  // namespace meta
}  // namespace v1
}  // namespace greptime
static ::_pb::Metadata file_level_metadata_greptime_2fv1_2fmeta_2fregion_2eproto[2];
static constexpr ::_pb::EnumDescriptor const** file_level_enum_descriptors_greptime_2fv1_2fmeta_2fregion_2eproto = nullptr;
static constexpr ::_pb::ServiceDescriptor const** file_level_service_descriptors_greptime_2fv1_2fmeta_2fregion_2eproto = nullptr;

const uint32_t TableStruct_greptime_2fv1_2fmeta_2fregion_2eproto::offsets[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) = {
  ~0u,  // no _has_bits_
  PROTOBUF_FIELD_OFFSET(::greptime::v1::meta::MigrateRegionRequest, _internal_metadata_),
  ~0u,  // no _extensions_
  ~0u,  // no _oneof_case_
  ~0u,  // no _weak_field_map_
  ~0u,  // no _inlined_string_donated_
  PROTOBUF_FIELD_OFFSET(::greptime::v1::meta::MigrateRegionRequest, _impl_.header_),
  PROTOBUF_FIELD_OFFSET(::greptime::v1::meta::MigrateRegionRequest, _impl_.region_id_),
  PROTOBUF_FIELD_OFFSET(::greptime::v1::meta::MigrateRegionRequest, _impl_.from_peer_),
  PROTOBUF_FIELD_OFFSET(::greptime::v1::meta::MigrateRegionRequest, _impl_.to_peer_),
  PROTOBUF_FIELD_OFFSET(::greptime::v1::meta::MigrateRegionRequest, _impl_.timeout_secs_),
  ~0u,  // no _has_bits_
  PROTOBUF_FIELD_OFFSET(::greptime::v1::meta::MigrateRegionResponse, _internal_metadata_),
  ~0u,  // no _extensions_
  ~0u,  // no _oneof_case_
  ~0u,  // no _weak_field_map_
  ~0u,  // no _inlined_string_donated_
  PROTOBUF_FIELD_OFFSET(::greptime::v1::meta::MigrateRegionResponse, _impl_.header_),
  PROTOBUF_FIELD_OFFSET(::greptime::v1::meta::MigrateRegionResponse, _impl_.pid_),
};
static const ::_pbi::MigrationSchema schemas[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) = {
  { 0, -1, -1, sizeof(::greptime::v1::meta::MigrateRegionRequest)},
  { 11, -1, -1, sizeof(::greptime::v1::meta::MigrateRegionResponse)},
};

static const ::_pb::Message* const file_default_instances[] = {
  &::greptime::v1::meta::_MigrateRegionRequest_default_instance_._instance,
  &::greptime::v1::meta::_MigrateRegionResponse_default_instance_._instance,
};

const char descriptor_table_protodef_greptime_2fv1_2fmeta_2fregion_2eproto[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) =
  "\n\035greptime/v1/meta/region.proto\022\020greptim"
  "e.v1.meta\032\035greptime/v1/meta/common.proto"
  "\"\224\001\n\024MigrateRegionRequest\022/\n\006header\030\001 \001("
  "\0132\037.greptime.v1.meta.RequestHeader\022\021\n\tre"
  "gion_id\030\003 \001(\004\022\021\n\tfrom_peer\030\004 \001(\004\022\017\n\007to_p"
  "eer\030\005 \001(\004\022\024\n\014timeout_secs\030\006 \001(\r\"u\n\025Migra"
  "teRegionResponse\0220\n\006header\030\001 \001(\0132 .grept"
  "ime.v1.meta.ResponseHeader\022*\n\003pid\030\002 \001(\0132"
  "\035.greptime.v1.meta.ProcedureIdB<Z:github"
  ".com/GreptimeTeam/greptime-proto/go/grep"
  "time/v1/metab\006proto3"
  ;
static const ::_pbi::DescriptorTable* const descriptor_table_greptime_2fv1_2fmeta_2fregion_2eproto_deps[1] = {
  &::descriptor_table_greptime_2fv1_2fmeta_2fcommon_2eproto,
};
static ::_pbi::once_flag descriptor_table_greptime_2fv1_2fmeta_2fregion_2eproto_once;
const ::_pbi::DescriptorTable descriptor_table_greptime_2fv1_2fmeta_2fregion_2eproto = {
    false, false, 420, descriptor_table_protodef_greptime_2fv1_2fmeta_2fregion_2eproto,
    "greptime/v1/meta/region.proto",
    &descriptor_table_greptime_2fv1_2fmeta_2fregion_2eproto_once, descriptor_table_greptime_2fv1_2fmeta_2fregion_2eproto_deps, 1, 2,
    schemas, file_default_instances, TableStruct_greptime_2fv1_2fmeta_2fregion_2eproto::offsets,
    file_level_metadata_greptime_2fv1_2fmeta_2fregion_2eproto, file_level_enum_descriptors_greptime_2fv1_2fmeta_2fregion_2eproto,
    file_level_service_descriptors_greptime_2fv1_2fmeta_2fregion_2eproto,
};
PROTOBUF_ATTRIBUTE_WEAK const ::_pbi::DescriptorTable* descriptor_table_greptime_2fv1_2fmeta_2fregion_2eproto_getter() {
  return &descriptor_table_greptime_2fv1_2fmeta_2fregion_2eproto;
}

// Force running AddDescriptors() at dynamic initialization time.
PROTOBUF_ATTRIBUTE_INIT_PRIORITY2 static ::_pbi::AddDescriptorsRunner dynamic_init_dummy_greptime_2fv1_2fmeta_2fregion_2eproto(&descriptor_table_greptime_2fv1_2fmeta_2fregion_2eproto);
namespace greptime {
namespace v1 {
namespace meta {

// ===================================================================

class MigrateRegionRequest::_Internal {
 public:
  static const ::greptime::v1::meta::RequestHeader& header(const MigrateRegionRequest* msg);
};

const ::greptime::v1::meta::RequestHeader&
MigrateRegionRequest::_Internal::header(const MigrateRegionRequest* msg) {
  return *msg->_impl_.header_;
}
void MigrateRegionRequest::clear_header() {
  if (GetArenaForAllocation() == nullptr && _impl_.header_ != nullptr) {
    delete _impl_.header_;
  }
  _impl_.header_ = nullptr;
}
MigrateRegionRequest::MigrateRegionRequest(::PROTOBUF_NAMESPACE_ID::Arena* arena,
                         bool is_message_owned)
  : ::PROTOBUF_NAMESPACE_ID::Message(arena, is_message_owned) {
  SharedCtor(arena, is_message_owned);
  // @@protoc_insertion_point(arena_constructor:greptime.v1.meta.MigrateRegionRequest)
}
MigrateRegionRequest::MigrateRegionRequest(const MigrateRegionRequest& from)
  : ::PROTOBUF_NAMESPACE_ID::Message() {
  MigrateRegionRequest* const _this = this; (void)_this;
  new (&_impl_) Impl_{
      decltype(_impl_.header_){nullptr}
    , decltype(_impl_.region_id_){}
    , decltype(_impl_.from_peer_){}
    , decltype(_impl_.to_peer_){}
    , decltype(_impl_.timeout_secs_){}
    , /*decltype(_impl_._cached_size_)*/{}};

  _internal_metadata_.MergeFrom<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(from._internal_metadata_);
  if (from._internal_has_header()) {
    _this->_impl_.header_ = new ::greptime::v1::meta::RequestHeader(*from._impl_.header_);
  }
  ::memcpy(&_impl_.region_id_, &from._impl_.region_id_,
    static_cast<size_t>(reinterpret_cast<char*>(&_impl_.timeout_secs_) -
    reinterpret_cast<char*>(&_impl_.region_id_)) + sizeof(_impl_.timeout_secs_));
  // @@protoc_insertion_point(copy_constructor:greptime.v1.meta.MigrateRegionRequest)
}

inline void MigrateRegionRequest::SharedCtor(
    ::_pb::Arena* arena, bool is_message_owned) {
  (void)arena;
  (void)is_message_owned;
  new (&_impl_) Impl_{
      decltype(_impl_.header_){nullptr}
    , decltype(_impl_.region_id_){uint64_t{0u}}
    , decltype(_impl_.from_peer_){uint64_t{0u}}
    , decltype(_impl_.to_peer_){uint64_t{0u}}
    , decltype(_impl_.timeout_secs_){0u}
    , /*decltype(_impl_._cached_size_)*/{}
  };
}

MigrateRegionRequest::~MigrateRegionRequest() {
  // @@protoc_insertion_point(destructor:greptime.v1.meta.MigrateRegionRequest)
  if (auto *arena = _internal_metadata_.DeleteReturnArena<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>()) {
  (void)arena;
    return;
  }
  SharedDtor();
}

inline void MigrateRegionRequest::SharedDtor() {
  GOOGLE_DCHECK(GetArenaForAllocation() == nullptr);
  if (this != internal_default_instance()) delete _impl_.header_;
}

void MigrateRegionRequest::SetCachedSize(int size) const {
  _impl_._cached_size_.Set(size);
}

void MigrateRegionRequest::Clear() {
// @@protoc_insertion_point(message_clear_start:greptime.v1.meta.MigrateRegionRequest)
  uint32_t cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  if (GetArenaForAllocation() == nullptr && _impl_.header_ != nullptr) {
    delete _impl_.header_;
  }
  _impl_.header_ = nullptr;
  ::memset(&_impl_.region_id_, 0, static_cast<size_t>(
      reinterpret_cast<char*>(&_impl_.timeout_secs_) -
      reinterpret_cast<char*>(&_impl_.region_id_)) + sizeof(_impl_.timeout_secs_));
  _internal_metadata_.Clear<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>();
}

const char* MigrateRegionRequest::_InternalParse(const char* ptr, ::_pbi::ParseContext* ctx) {
#define CHK_(x) if (PROTOBUF_PREDICT_FALSE(!(x))) goto failure
  while (!ctx->Done(&ptr)) {
    uint32_t tag;
    ptr = ::_pbi::ReadTag(ptr, &tag);
    switch (tag >> 3) {
      // .greptime.v1.meta.RequestHeader header = 1;
      case 1:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 10)) {
          ptr = ctx->ParseMessage(_internal_mutable_header(), ptr);
          CHK_(ptr);
        } else
          goto handle_unusual;
        continue;
      // uint64 region_id = 3;
      case 3:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 24)) {
          _impl_.region_id_ = ::PROTOBUF_NAMESPACE_ID::internal::ReadVarint64(&ptr);
          CHK_(ptr);
        } else
          goto handle_unusual;
        continue;
      // uint64 from_peer = 4;
      case 4:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 32)) {
          _impl_.from_peer_ = ::PROTOBUF_NAMESPACE_ID::internal::ReadVarint64(&ptr);
          CHK_(ptr);
        } else
          goto handle_unusual;
        continue;
      // uint64 to_peer = 5;
      case 5:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 40)) {
          _impl_.to_peer_ = ::PROTOBUF_NAMESPACE_ID::internal::ReadVarint64(&ptr);
          CHK_(ptr);
        } else
          goto handle_unusual;
        continue;
      // uint32 timeout_secs = 6;
      case 6:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 48)) {
          _impl_.timeout_secs_ = ::PROTOBUF_NAMESPACE_ID::internal::ReadVarint32(&ptr);
          CHK_(ptr);
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

uint8_t* MigrateRegionRequest::_InternalSerialize(
    uint8_t* target, ::PROTOBUF_NAMESPACE_ID::io::EpsCopyOutputStream* stream) const {
  // @@protoc_insertion_point(serialize_to_array_start:greptime.v1.meta.MigrateRegionRequest)
  uint32_t cached_has_bits = 0;
  (void) cached_has_bits;

  // .greptime.v1.meta.RequestHeader header = 1;
  if (this->_internal_has_header()) {
    target = ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::
      InternalWriteMessage(1, _Internal::header(this),
        _Internal::header(this).GetCachedSize(), target, stream);
  }

  // uint64 region_id = 3;
  if (this->_internal_region_id() != 0) {
    target = stream->EnsureSpace(target);
    target = ::_pbi::WireFormatLite::WriteUInt64ToArray(3, this->_internal_region_id(), target);
  }

  // uint64 from_peer = 4;
  if (this->_internal_from_peer() != 0) {
    target = stream->EnsureSpace(target);
    target = ::_pbi::WireFormatLite::WriteUInt64ToArray(4, this->_internal_from_peer(), target);
  }

  // uint64 to_peer = 5;
  if (this->_internal_to_peer() != 0) {
    target = stream->EnsureSpace(target);
    target = ::_pbi::WireFormatLite::WriteUInt64ToArray(5, this->_internal_to_peer(), target);
  }

  // uint32 timeout_secs = 6;
  if (this->_internal_timeout_secs() != 0) {
    target = stream->EnsureSpace(target);
    target = ::_pbi::WireFormatLite::WriteUInt32ToArray(6, this->_internal_timeout_secs(), target);
  }

  if (PROTOBUF_PREDICT_FALSE(_internal_metadata_.have_unknown_fields())) {
    target = ::_pbi::WireFormat::InternalSerializeUnknownFieldsToArray(
        _internal_metadata_.unknown_fields<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(::PROTOBUF_NAMESPACE_ID::UnknownFieldSet::default_instance), target, stream);
  }
  // @@protoc_insertion_point(serialize_to_array_end:greptime.v1.meta.MigrateRegionRequest)
  return target;
}

size_t MigrateRegionRequest::ByteSizeLong() const {
// @@protoc_insertion_point(message_byte_size_start:greptime.v1.meta.MigrateRegionRequest)
  size_t total_size = 0;

  uint32_t cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  // .greptime.v1.meta.RequestHeader header = 1;
  if (this->_internal_has_header()) {
    total_size += 1 +
      ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::MessageSize(
        *_impl_.header_);
  }

  // uint64 region_id = 3;
  if (this->_internal_region_id() != 0) {
    total_size += ::_pbi::WireFormatLite::UInt64SizePlusOne(this->_internal_region_id());
  }

  // uint64 from_peer = 4;
  if (this->_internal_from_peer() != 0) {
    total_size += ::_pbi::WireFormatLite::UInt64SizePlusOne(this->_internal_from_peer());
  }

  // uint64 to_peer = 5;
  if (this->_internal_to_peer() != 0) {
    total_size += ::_pbi::WireFormatLite::UInt64SizePlusOne(this->_internal_to_peer());
  }

  // uint32 timeout_secs = 6;
  if (this->_internal_timeout_secs() != 0) {
    total_size += ::_pbi::WireFormatLite::UInt32SizePlusOne(this->_internal_timeout_secs());
  }

  return MaybeComputeUnknownFieldsSize(total_size, &_impl_._cached_size_);
}

const ::PROTOBUF_NAMESPACE_ID::Message::ClassData MigrateRegionRequest::_class_data_ = {
    ::PROTOBUF_NAMESPACE_ID::Message::CopyWithSourceCheck,
    MigrateRegionRequest::MergeImpl
};
const ::PROTOBUF_NAMESPACE_ID::Message::ClassData*MigrateRegionRequest::GetClassData() const { return &_class_data_; }


void MigrateRegionRequest::MergeImpl(::PROTOBUF_NAMESPACE_ID::Message& to_msg, const ::PROTOBUF_NAMESPACE_ID::Message& from_msg) {
  auto* const _this = static_cast<MigrateRegionRequest*>(&to_msg);
  auto& from = static_cast<const MigrateRegionRequest&>(from_msg);
  // @@protoc_insertion_point(class_specific_merge_from_start:greptime.v1.meta.MigrateRegionRequest)
  GOOGLE_DCHECK_NE(&from, _this);
  uint32_t cached_has_bits = 0;
  (void) cached_has_bits;

  if (from._internal_has_header()) {
    _this->_internal_mutable_header()->::greptime::v1::meta::RequestHeader::MergeFrom(
        from._internal_header());
  }
  if (from._internal_region_id() != 0) {
    _this->_internal_set_region_id(from._internal_region_id());
  }
  if (from._internal_from_peer() != 0) {
    _this->_internal_set_from_peer(from._internal_from_peer());
  }
  if (from._internal_to_peer() != 0) {
    _this->_internal_set_to_peer(from._internal_to_peer());
  }
  if (from._internal_timeout_secs() != 0) {
    _this->_internal_set_timeout_secs(from._internal_timeout_secs());
  }
  _this->_internal_metadata_.MergeFrom<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(from._internal_metadata_);
}

void MigrateRegionRequest::CopyFrom(const MigrateRegionRequest& from) {
// @@protoc_insertion_point(class_specific_copy_from_start:greptime.v1.meta.MigrateRegionRequest)
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

bool MigrateRegionRequest::IsInitialized() const {
  return true;
}

void MigrateRegionRequest::InternalSwap(MigrateRegionRequest* other) {
  using std::swap;
  _internal_metadata_.InternalSwap(&other->_internal_metadata_);
  ::PROTOBUF_NAMESPACE_ID::internal::memswap<
      PROTOBUF_FIELD_OFFSET(MigrateRegionRequest, _impl_.timeout_secs_)
      + sizeof(MigrateRegionRequest::_impl_.timeout_secs_)
      - PROTOBUF_FIELD_OFFSET(MigrateRegionRequest, _impl_.header_)>(
          reinterpret_cast<char*>(&_impl_.header_),
          reinterpret_cast<char*>(&other->_impl_.header_));
}

::PROTOBUF_NAMESPACE_ID::Metadata MigrateRegionRequest::GetMetadata() const {
  return ::_pbi::AssignDescriptors(
      &descriptor_table_greptime_2fv1_2fmeta_2fregion_2eproto_getter, &descriptor_table_greptime_2fv1_2fmeta_2fregion_2eproto_once,
      file_level_metadata_greptime_2fv1_2fmeta_2fregion_2eproto[0]);
}

// ===================================================================

class MigrateRegionResponse::_Internal {
 public:
  static const ::greptime::v1::meta::ResponseHeader& header(const MigrateRegionResponse* msg);
  static const ::greptime::v1::meta::ProcedureId& pid(const MigrateRegionResponse* msg);
};

const ::greptime::v1::meta::ResponseHeader&
MigrateRegionResponse::_Internal::header(const MigrateRegionResponse* msg) {
  return *msg->_impl_.header_;
}
const ::greptime::v1::meta::ProcedureId&
MigrateRegionResponse::_Internal::pid(const MigrateRegionResponse* msg) {
  return *msg->_impl_.pid_;
}
void MigrateRegionResponse::clear_header() {
  if (GetArenaForAllocation() == nullptr && _impl_.header_ != nullptr) {
    delete _impl_.header_;
  }
  _impl_.header_ = nullptr;
}
void MigrateRegionResponse::clear_pid() {
  if (GetArenaForAllocation() == nullptr && _impl_.pid_ != nullptr) {
    delete _impl_.pid_;
  }
  _impl_.pid_ = nullptr;
}
MigrateRegionResponse::MigrateRegionResponse(::PROTOBUF_NAMESPACE_ID::Arena* arena,
                         bool is_message_owned)
  : ::PROTOBUF_NAMESPACE_ID::Message(arena, is_message_owned) {
  SharedCtor(arena, is_message_owned);
  // @@protoc_insertion_point(arena_constructor:greptime.v1.meta.MigrateRegionResponse)
}
MigrateRegionResponse::MigrateRegionResponse(const MigrateRegionResponse& from)
  : ::PROTOBUF_NAMESPACE_ID::Message() {
  MigrateRegionResponse* const _this = this; (void)_this;
  new (&_impl_) Impl_{
      decltype(_impl_.header_){nullptr}
    , decltype(_impl_.pid_){nullptr}
    , /*decltype(_impl_._cached_size_)*/{}};

  _internal_metadata_.MergeFrom<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(from._internal_metadata_);
  if (from._internal_has_header()) {
    _this->_impl_.header_ = new ::greptime::v1::meta::ResponseHeader(*from._impl_.header_);
  }
  if (from._internal_has_pid()) {
    _this->_impl_.pid_ = new ::greptime::v1::meta::ProcedureId(*from._impl_.pid_);
  }
  // @@protoc_insertion_point(copy_constructor:greptime.v1.meta.MigrateRegionResponse)
}

inline void MigrateRegionResponse::SharedCtor(
    ::_pb::Arena* arena, bool is_message_owned) {
  (void)arena;
  (void)is_message_owned;
  new (&_impl_) Impl_{
      decltype(_impl_.header_){nullptr}
    , decltype(_impl_.pid_){nullptr}
    , /*decltype(_impl_._cached_size_)*/{}
  };
}

MigrateRegionResponse::~MigrateRegionResponse() {
  // @@protoc_insertion_point(destructor:greptime.v1.meta.MigrateRegionResponse)
  if (auto *arena = _internal_metadata_.DeleteReturnArena<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>()) {
  (void)arena;
    return;
  }
  SharedDtor();
}

inline void MigrateRegionResponse::SharedDtor() {
  GOOGLE_DCHECK(GetArenaForAllocation() == nullptr);
  if (this != internal_default_instance()) delete _impl_.header_;
  if (this != internal_default_instance()) delete _impl_.pid_;
}

void MigrateRegionResponse::SetCachedSize(int size) const {
  _impl_._cached_size_.Set(size);
}

void MigrateRegionResponse::Clear() {
// @@protoc_insertion_point(message_clear_start:greptime.v1.meta.MigrateRegionResponse)
  uint32_t cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  if (GetArenaForAllocation() == nullptr && _impl_.header_ != nullptr) {
    delete _impl_.header_;
  }
  _impl_.header_ = nullptr;
  if (GetArenaForAllocation() == nullptr && _impl_.pid_ != nullptr) {
    delete _impl_.pid_;
  }
  _impl_.pid_ = nullptr;
  _internal_metadata_.Clear<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>();
}

const char* MigrateRegionResponse::_InternalParse(const char* ptr, ::_pbi::ParseContext* ctx) {
#define CHK_(x) if (PROTOBUF_PREDICT_FALSE(!(x))) goto failure
  while (!ctx->Done(&ptr)) {
    uint32_t tag;
    ptr = ::_pbi::ReadTag(ptr, &tag);
    switch (tag >> 3) {
      // .greptime.v1.meta.ResponseHeader header = 1;
      case 1:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 10)) {
          ptr = ctx->ParseMessage(_internal_mutable_header(), ptr);
          CHK_(ptr);
        } else
          goto handle_unusual;
        continue;
      // .greptime.v1.meta.ProcedureId pid = 2;
      case 2:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 18)) {
          ptr = ctx->ParseMessage(_internal_mutable_pid(), ptr);
          CHK_(ptr);
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

uint8_t* MigrateRegionResponse::_InternalSerialize(
    uint8_t* target, ::PROTOBUF_NAMESPACE_ID::io::EpsCopyOutputStream* stream) const {
  // @@protoc_insertion_point(serialize_to_array_start:greptime.v1.meta.MigrateRegionResponse)
  uint32_t cached_has_bits = 0;
  (void) cached_has_bits;

  // .greptime.v1.meta.ResponseHeader header = 1;
  if (this->_internal_has_header()) {
    target = ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::
      InternalWriteMessage(1, _Internal::header(this),
        _Internal::header(this).GetCachedSize(), target, stream);
  }

  // .greptime.v1.meta.ProcedureId pid = 2;
  if (this->_internal_has_pid()) {
    target = ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::
      InternalWriteMessage(2, _Internal::pid(this),
        _Internal::pid(this).GetCachedSize(), target, stream);
  }

  if (PROTOBUF_PREDICT_FALSE(_internal_metadata_.have_unknown_fields())) {
    target = ::_pbi::WireFormat::InternalSerializeUnknownFieldsToArray(
        _internal_metadata_.unknown_fields<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(::PROTOBUF_NAMESPACE_ID::UnknownFieldSet::default_instance), target, stream);
  }
  // @@protoc_insertion_point(serialize_to_array_end:greptime.v1.meta.MigrateRegionResponse)
  return target;
}

size_t MigrateRegionResponse::ByteSizeLong() const {
// @@protoc_insertion_point(message_byte_size_start:greptime.v1.meta.MigrateRegionResponse)
  size_t total_size = 0;

  uint32_t cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  // .greptime.v1.meta.ResponseHeader header = 1;
  if (this->_internal_has_header()) {
    total_size += 1 +
      ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::MessageSize(
        *_impl_.header_);
  }

  // .greptime.v1.meta.ProcedureId pid = 2;
  if (this->_internal_has_pid()) {
    total_size += 1 +
      ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::MessageSize(
        *_impl_.pid_);
  }

  return MaybeComputeUnknownFieldsSize(total_size, &_impl_._cached_size_);
}

const ::PROTOBUF_NAMESPACE_ID::Message::ClassData MigrateRegionResponse::_class_data_ = {
    ::PROTOBUF_NAMESPACE_ID::Message::CopyWithSourceCheck,
    MigrateRegionResponse::MergeImpl
};
const ::PROTOBUF_NAMESPACE_ID::Message::ClassData*MigrateRegionResponse::GetClassData() const { return &_class_data_; }


void MigrateRegionResponse::MergeImpl(::PROTOBUF_NAMESPACE_ID::Message& to_msg, const ::PROTOBUF_NAMESPACE_ID::Message& from_msg) {
  auto* const _this = static_cast<MigrateRegionResponse*>(&to_msg);
  auto& from = static_cast<const MigrateRegionResponse&>(from_msg);
  // @@protoc_insertion_point(class_specific_merge_from_start:greptime.v1.meta.MigrateRegionResponse)
  GOOGLE_DCHECK_NE(&from, _this);
  uint32_t cached_has_bits = 0;
  (void) cached_has_bits;

  if (from._internal_has_header()) {
    _this->_internal_mutable_header()->::greptime::v1::meta::ResponseHeader::MergeFrom(
        from._internal_header());
  }
  if (from._internal_has_pid()) {
    _this->_internal_mutable_pid()->::greptime::v1::meta::ProcedureId::MergeFrom(
        from._internal_pid());
  }
  _this->_internal_metadata_.MergeFrom<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(from._internal_metadata_);
}

void MigrateRegionResponse::CopyFrom(const MigrateRegionResponse& from) {
// @@protoc_insertion_point(class_specific_copy_from_start:greptime.v1.meta.MigrateRegionResponse)
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

bool MigrateRegionResponse::IsInitialized() const {
  return true;
}

void MigrateRegionResponse::InternalSwap(MigrateRegionResponse* other) {
  using std::swap;
  _internal_metadata_.InternalSwap(&other->_internal_metadata_);
  ::PROTOBUF_NAMESPACE_ID::internal::memswap<
      PROTOBUF_FIELD_OFFSET(MigrateRegionResponse, _impl_.pid_)
      + sizeof(MigrateRegionResponse::_impl_.pid_)
      - PROTOBUF_FIELD_OFFSET(MigrateRegionResponse, _impl_.header_)>(
          reinterpret_cast<char*>(&_impl_.header_),
          reinterpret_cast<char*>(&other->_impl_.header_));
}

::PROTOBUF_NAMESPACE_ID::Metadata MigrateRegionResponse::GetMetadata() const {
  return ::_pbi::AssignDescriptors(
      &descriptor_table_greptime_2fv1_2fmeta_2fregion_2eproto_getter, &descriptor_table_greptime_2fv1_2fmeta_2fregion_2eproto_once,
      file_level_metadata_greptime_2fv1_2fmeta_2fregion_2eproto[1]);
}

// @@protoc_insertion_point(namespace_scope)
}  // namespace meta
}  // namespace v1
}  // namespace greptime
PROTOBUF_NAMESPACE_OPEN
template<> PROTOBUF_NOINLINE ::greptime::v1::meta::MigrateRegionRequest*
Arena::CreateMaybeMessage< ::greptime::v1::meta::MigrateRegionRequest >(Arena* arena) {
  return Arena::CreateMessageInternal< ::greptime::v1::meta::MigrateRegionRequest >(arena);
}
template<> PROTOBUF_NOINLINE ::greptime::v1::meta::MigrateRegionResponse*
Arena::CreateMaybeMessage< ::greptime::v1::meta::MigrateRegionResponse >(Arena* arena) {
  return Arena::CreateMessageInternal< ::greptime::v1::meta::MigrateRegionResponse >(arena);
}
PROTOBUF_NAMESPACE_CLOSE

// @@protoc_insertion_point(global_scope)
#include <google/protobuf/port_undef.inc>
