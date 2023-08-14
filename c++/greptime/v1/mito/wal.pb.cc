// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: greptime/v1/mito/wal.proto

#include "greptime/v1/mito/wal.pb.h"

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
namespace mito {
PROTOBUF_CONSTEXPR Mutation::Mutation(
    ::_pbi::ConstantInitialized): _impl_{
    /*decltype(_impl_.rows_)*/nullptr
  , /*decltype(_impl_.sequence_)*/uint64_t{0u}
  , /*decltype(_impl_.op_type_)*/0
  , /*decltype(_impl_._cached_size_)*/{}} {}
struct MutationDefaultTypeInternal {
  PROTOBUF_CONSTEXPR MutationDefaultTypeInternal()
      : _instance(::_pbi::ConstantInitialized{}) {}
  ~MutationDefaultTypeInternal() {}
  union {
    Mutation _instance;
  };
};
PROTOBUF_ATTRIBUTE_NO_DESTROY PROTOBUF_CONSTINIT PROTOBUF_ATTRIBUTE_INIT_PRIORITY1 MutationDefaultTypeInternal _Mutation_default_instance_;
PROTOBUF_CONSTEXPR WalEntry::WalEntry(
    ::_pbi::ConstantInitialized): _impl_{
    /*decltype(_impl_.mutations_)*/{}
  , /*decltype(_impl_._cached_size_)*/{}} {}
struct WalEntryDefaultTypeInternal {
  PROTOBUF_CONSTEXPR WalEntryDefaultTypeInternal()
      : _instance(::_pbi::ConstantInitialized{}) {}
  ~WalEntryDefaultTypeInternal() {}
  union {
    WalEntry _instance;
  };
};
PROTOBUF_ATTRIBUTE_NO_DESTROY PROTOBUF_CONSTINIT PROTOBUF_ATTRIBUTE_INIT_PRIORITY1 WalEntryDefaultTypeInternal _WalEntry_default_instance_;
}  // namespace mito
}  // namespace v1
}  // namespace greptime
static ::_pb::Metadata file_level_metadata_greptime_2fv1_2fmito_2fwal_2eproto[2];
static const ::_pb::EnumDescriptor* file_level_enum_descriptors_greptime_2fv1_2fmito_2fwal_2eproto[1];
static constexpr ::_pb::ServiceDescriptor const** file_level_service_descriptors_greptime_2fv1_2fmito_2fwal_2eproto = nullptr;

const uint32_t TableStruct_greptime_2fv1_2fmito_2fwal_2eproto::offsets[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) = {
  ~0u,  // no _has_bits_
  PROTOBUF_FIELD_OFFSET(::greptime::v1::mito::Mutation, _internal_metadata_),
  ~0u,  // no _extensions_
  ~0u,  // no _oneof_case_
  ~0u,  // no _weak_field_map_
  ~0u,  // no _inlined_string_donated_
  PROTOBUF_FIELD_OFFSET(::greptime::v1::mito::Mutation, _impl_.op_type_),
  PROTOBUF_FIELD_OFFSET(::greptime::v1::mito::Mutation, _impl_.sequence_),
  PROTOBUF_FIELD_OFFSET(::greptime::v1::mito::Mutation, _impl_.rows_),
  ~0u,  // no _has_bits_
  PROTOBUF_FIELD_OFFSET(::greptime::v1::mito::WalEntry, _internal_metadata_),
  ~0u,  // no _extensions_
  ~0u,  // no _oneof_case_
  ~0u,  // no _weak_field_map_
  ~0u,  // no _inlined_string_donated_
  PROTOBUF_FIELD_OFFSET(::greptime::v1::mito::WalEntry, _impl_.mutations_),
};
static const ::_pbi::MigrationSchema schemas[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) = {
  { 0, -1, -1, sizeof(::greptime::v1::mito::Mutation)},
  { 9, -1, -1, sizeof(::greptime::v1::mito::WalEntry)},
};

static const ::_pb::Message* const file_default_instances[] = {
  &::greptime::v1::mito::_Mutation_default_instance_._instance,
  &::greptime::v1::mito::_WalEntry_default_instance_._instance,
};

const char descriptor_table_protodef_greptime_2fv1_2fmito_2fwal_2eproto[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) =
  "\n\032greptime/v1/mito/wal.proto\022\020greptime.v"
  "1.mito\032\025greptime/v1/row.proto\"h\n\010Mutatio"
  "n\022)\n\007op_type\030\001 \001(\0162\030.greptime.v1.mito.Op"
  "Type\022\020\n\010sequence\030\002 \001(\004\022\037\n\004rows\030\003 \001(\0132\021.g"
  "reptime.v1.Rows\"9\n\010WalEntry\022-\n\tmutations"
  "\030\001 \003(\0132\032.greptime.v1.mito.Mutation*\035\n\006Op"
  "Type\022\n\n\006DELETE\020\000\022\007\n\003PUT\020\001B<Z:github.com/"
  "GreptimeTeam/greptime-proto/go/greptime/"
  "v1/mitob\006proto3"
  ;
static const ::_pbi::DescriptorTable* const descriptor_table_greptime_2fv1_2fmito_2fwal_2eproto_deps[1] = {
  &::descriptor_table_greptime_2fv1_2frow_2eproto,
};
static ::_pbi::once_flag descriptor_table_greptime_2fv1_2fmito_2fwal_2eproto_once;
const ::_pbi::DescriptorTable descriptor_table_greptime_2fv1_2fmito_2fwal_2eproto = {
    false, false, 335, descriptor_table_protodef_greptime_2fv1_2fmito_2fwal_2eproto,
    "greptime/v1/mito/wal.proto",
    &descriptor_table_greptime_2fv1_2fmito_2fwal_2eproto_once, descriptor_table_greptime_2fv1_2fmito_2fwal_2eproto_deps, 1, 2,
    schemas, file_default_instances, TableStruct_greptime_2fv1_2fmito_2fwal_2eproto::offsets,
    file_level_metadata_greptime_2fv1_2fmito_2fwal_2eproto, file_level_enum_descriptors_greptime_2fv1_2fmito_2fwal_2eproto,
    file_level_service_descriptors_greptime_2fv1_2fmito_2fwal_2eproto,
};
PROTOBUF_ATTRIBUTE_WEAK const ::_pbi::DescriptorTable* descriptor_table_greptime_2fv1_2fmito_2fwal_2eproto_getter() {
  return &descriptor_table_greptime_2fv1_2fmito_2fwal_2eproto;
}

// Force running AddDescriptors() at dynamic initialization time.
PROTOBUF_ATTRIBUTE_INIT_PRIORITY2 static ::_pbi::AddDescriptorsRunner dynamic_init_dummy_greptime_2fv1_2fmito_2fwal_2eproto(&descriptor_table_greptime_2fv1_2fmito_2fwal_2eproto);
namespace greptime {
namespace v1 {
namespace mito {
const ::PROTOBUF_NAMESPACE_ID::EnumDescriptor* OpType_descriptor() {
  ::PROTOBUF_NAMESPACE_ID::internal::AssignDescriptors(&descriptor_table_greptime_2fv1_2fmito_2fwal_2eproto);
  return file_level_enum_descriptors_greptime_2fv1_2fmito_2fwal_2eproto[0];
}
bool OpType_IsValid(int value) {
  switch (value) {
    case 0:
    case 1:
      return true;
    default:
      return false;
  }
}


// ===================================================================

class Mutation::_Internal {
 public:
  static const ::greptime::v1::Rows& rows(const Mutation* msg);
};

const ::greptime::v1::Rows&
Mutation::_Internal::rows(const Mutation* msg) {
  return *msg->_impl_.rows_;
}
void Mutation::clear_rows() {
  if (GetArenaForAllocation() == nullptr && _impl_.rows_ != nullptr) {
    delete _impl_.rows_;
  }
  _impl_.rows_ = nullptr;
}
Mutation::Mutation(::PROTOBUF_NAMESPACE_ID::Arena* arena,
                         bool is_message_owned)
  : ::PROTOBUF_NAMESPACE_ID::Message(arena, is_message_owned) {
  SharedCtor(arena, is_message_owned);
  // @@protoc_insertion_point(arena_constructor:greptime.v1.mito.Mutation)
}
Mutation::Mutation(const Mutation& from)
  : ::PROTOBUF_NAMESPACE_ID::Message() {
  Mutation* const _this = this; (void)_this;
  new (&_impl_) Impl_{
      decltype(_impl_.rows_){nullptr}
    , decltype(_impl_.sequence_){}
    , decltype(_impl_.op_type_){}
    , /*decltype(_impl_._cached_size_)*/{}};

  _internal_metadata_.MergeFrom<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(from._internal_metadata_);
  if (from._internal_has_rows()) {
    _this->_impl_.rows_ = new ::greptime::v1::Rows(*from._impl_.rows_);
  }
  ::memcpy(&_impl_.sequence_, &from._impl_.sequence_,
    static_cast<size_t>(reinterpret_cast<char*>(&_impl_.op_type_) -
    reinterpret_cast<char*>(&_impl_.sequence_)) + sizeof(_impl_.op_type_));
  // @@protoc_insertion_point(copy_constructor:greptime.v1.mito.Mutation)
}

inline void Mutation::SharedCtor(
    ::_pb::Arena* arena, bool is_message_owned) {
  (void)arena;
  (void)is_message_owned;
  new (&_impl_) Impl_{
      decltype(_impl_.rows_){nullptr}
    , decltype(_impl_.sequence_){uint64_t{0u}}
    , decltype(_impl_.op_type_){0}
    , /*decltype(_impl_._cached_size_)*/{}
  };
}

Mutation::~Mutation() {
  // @@protoc_insertion_point(destructor:greptime.v1.mito.Mutation)
  if (auto *arena = _internal_metadata_.DeleteReturnArena<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>()) {
  (void)arena;
    return;
  }
  SharedDtor();
}

inline void Mutation::SharedDtor() {
  GOOGLE_DCHECK(GetArenaForAllocation() == nullptr);
  if (this != internal_default_instance()) delete _impl_.rows_;
}

void Mutation::SetCachedSize(int size) const {
  _impl_._cached_size_.Set(size);
}

void Mutation::Clear() {
// @@protoc_insertion_point(message_clear_start:greptime.v1.mito.Mutation)
  uint32_t cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  if (GetArenaForAllocation() == nullptr && _impl_.rows_ != nullptr) {
    delete _impl_.rows_;
  }
  _impl_.rows_ = nullptr;
  ::memset(&_impl_.sequence_, 0, static_cast<size_t>(
      reinterpret_cast<char*>(&_impl_.op_type_) -
      reinterpret_cast<char*>(&_impl_.sequence_)) + sizeof(_impl_.op_type_));
  _internal_metadata_.Clear<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>();
}

const char* Mutation::_InternalParse(const char* ptr, ::_pbi::ParseContext* ctx) {
#define CHK_(x) if (PROTOBUF_PREDICT_FALSE(!(x))) goto failure
  while (!ctx->Done(&ptr)) {
    uint32_t tag;
    ptr = ::_pbi::ReadTag(ptr, &tag);
    switch (tag >> 3) {
      // .greptime.v1.mito.OpType op_type = 1;
      case 1:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 8)) {
          uint64_t val = ::PROTOBUF_NAMESPACE_ID::internal::ReadVarint64(&ptr);
          CHK_(ptr);
          _internal_set_op_type(static_cast<::greptime::v1::mito::OpType>(val));
        } else
          goto handle_unusual;
        continue;
      // uint64 sequence = 2;
      case 2:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 16)) {
          _impl_.sequence_ = ::PROTOBUF_NAMESPACE_ID::internal::ReadVarint64(&ptr);
          CHK_(ptr);
        } else
          goto handle_unusual;
        continue;
      // .greptime.v1.Rows rows = 3;
      case 3:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 26)) {
          ptr = ctx->ParseMessage(_internal_mutable_rows(), ptr);
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

uint8_t* Mutation::_InternalSerialize(
    uint8_t* target, ::PROTOBUF_NAMESPACE_ID::io::EpsCopyOutputStream* stream) const {
  // @@protoc_insertion_point(serialize_to_array_start:greptime.v1.mito.Mutation)
  uint32_t cached_has_bits = 0;
  (void) cached_has_bits;

  // .greptime.v1.mito.OpType op_type = 1;
  if (this->_internal_op_type() != 0) {
    target = stream->EnsureSpace(target);
    target = ::_pbi::WireFormatLite::WriteEnumToArray(
      1, this->_internal_op_type(), target);
  }

  // uint64 sequence = 2;
  if (this->_internal_sequence() != 0) {
    target = stream->EnsureSpace(target);
    target = ::_pbi::WireFormatLite::WriteUInt64ToArray(2, this->_internal_sequence(), target);
  }

  // .greptime.v1.Rows rows = 3;
  if (this->_internal_has_rows()) {
    target = ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::
      InternalWriteMessage(3, _Internal::rows(this),
        _Internal::rows(this).GetCachedSize(), target, stream);
  }

  if (PROTOBUF_PREDICT_FALSE(_internal_metadata_.have_unknown_fields())) {
    target = ::_pbi::WireFormat::InternalSerializeUnknownFieldsToArray(
        _internal_metadata_.unknown_fields<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(::PROTOBUF_NAMESPACE_ID::UnknownFieldSet::default_instance), target, stream);
  }
  // @@protoc_insertion_point(serialize_to_array_end:greptime.v1.mito.Mutation)
  return target;
}

size_t Mutation::ByteSizeLong() const {
// @@protoc_insertion_point(message_byte_size_start:greptime.v1.mito.Mutation)
  size_t total_size = 0;

  uint32_t cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  // .greptime.v1.Rows rows = 3;
  if (this->_internal_has_rows()) {
    total_size += 1 +
      ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::MessageSize(
        *_impl_.rows_);
  }

  // uint64 sequence = 2;
  if (this->_internal_sequence() != 0) {
    total_size += ::_pbi::WireFormatLite::UInt64SizePlusOne(this->_internal_sequence());
  }

  // .greptime.v1.mito.OpType op_type = 1;
  if (this->_internal_op_type() != 0) {
    total_size += 1 +
      ::_pbi::WireFormatLite::EnumSize(this->_internal_op_type());
  }

  return MaybeComputeUnknownFieldsSize(total_size, &_impl_._cached_size_);
}

const ::PROTOBUF_NAMESPACE_ID::Message::ClassData Mutation::_class_data_ = {
    ::PROTOBUF_NAMESPACE_ID::Message::CopyWithSourceCheck,
    Mutation::MergeImpl
};
const ::PROTOBUF_NAMESPACE_ID::Message::ClassData*Mutation::GetClassData() const { return &_class_data_; }


void Mutation::MergeImpl(::PROTOBUF_NAMESPACE_ID::Message& to_msg, const ::PROTOBUF_NAMESPACE_ID::Message& from_msg) {
  auto* const _this = static_cast<Mutation*>(&to_msg);
  auto& from = static_cast<const Mutation&>(from_msg);
  // @@protoc_insertion_point(class_specific_merge_from_start:greptime.v1.mito.Mutation)
  GOOGLE_DCHECK_NE(&from, _this);
  uint32_t cached_has_bits = 0;
  (void) cached_has_bits;

  if (from._internal_has_rows()) {
    _this->_internal_mutable_rows()->::greptime::v1::Rows::MergeFrom(
        from._internal_rows());
  }
  if (from._internal_sequence() != 0) {
    _this->_internal_set_sequence(from._internal_sequence());
  }
  if (from._internal_op_type() != 0) {
    _this->_internal_set_op_type(from._internal_op_type());
  }
  _this->_internal_metadata_.MergeFrom<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(from._internal_metadata_);
}

void Mutation::CopyFrom(const Mutation& from) {
// @@protoc_insertion_point(class_specific_copy_from_start:greptime.v1.mito.Mutation)
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

bool Mutation::IsInitialized() const {
  return true;
}

void Mutation::InternalSwap(Mutation* other) {
  using std::swap;
  _internal_metadata_.InternalSwap(&other->_internal_metadata_);
  ::PROTOBUF_NAMESPACE_ID::internal::memswap<
      PROTOBUF_FIELD_OFFSET(Mutation, _impl_.op_type_)
      + sizeof(Mutation::_impl_.op_type_)
      - PROTOBUF_FIELD_OFFSET(Mutation, _impl_.rows_)>(
          reinterpret_cast<char*>(&_impl_.rows_),
          reinterpret_cast<char*>(&other->_impl_.rows_));
}

::PROTOBUF_NAMESPACE_ID::Metadata Mutation::GetMetadata() const {
  return ::_pbi::AssignDescriptors(
      &descriptor_table_greptime_2fv1_2fmito_2fwal_2eproto_getter, &descriptor_table_greptime_2fv1_2fmito_2fwal_2eproto_once,
      file_level_metadata_greptime_2fv1_2fmito_2fwal_2eproto[0]);
}

// ===================================================================

class WalEntry::_Internal {
 public:
};

WalEntry::WalEntry(::PROTOBUF_NAMESPACE_ID::Arena* arena,
                         bool is_message_owned)
  : ::PROTOBUF_NAMESPACE_ID::Message(arena, is_message_owned) {
  SharedCtor(arena, is_message_owned);
  // @@protoc_insertion_point(arena_constructor:greptime.v1.mito.WalEntry)
}
WalEntry::WalEntry(const WalEntry& from)
  : ::PROTOBUF_NAMESPACE_ID::Message() {
  WalEntry* const _this = this; (void)_this;
  new (&_impl_) Impl_{
      decltype(_impl_.mutations_){from._impl_.mutations_}
    , /*decltype(_impl_._cached_size_)*/{}};

  _internal_metadata_.MergeFrom<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(from._internal_metadata_);
  // @@protoc_insertion_point(copy_constructor:greptime.v1.mito.WalEntry)
}

inline void WalEntry::SharedCtor(
    ::_pb::Arena* arena, bool is_message_owned) {
  (void)arena;
  (void)is_message_owned;
  new (&_impl_) Impl_{
      decltype(_impl_.mutations_){arena}
    , /*decltype(_impl_._cached_size_)*/{}
  };
}

WalEntry::~WalEntry() {
  // @@protoc_insertion_point(destructor:greptime.v1.mito.WalEntry)
  if (auto *arena = _internal_metadata_.DeleteReturnArena<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>()) {
  (void)arena;
    return;
  }
  SharedDtor();
}

inline void WalEntry::SharedDtor() {
  GOOGLE_DCHECK(GetArenaForAllocation() == nullptr);
  _impl_.mutations_.~RepeatedPtrField();
}

void WalEntry::SetCachedSize(int size) const {
  _impl_._cached_size_.Set(size);
}

void WalEntry::Clear() {
// @@protoc_insertion_point(message_clear_start:greptime.v1.mito.WalEntry)
  uint32_t cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  _impl_.mutations_.Clear();
  _internal_metadata_.Clear<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>();
}

const char* WalEntry::_InternalParse(const char* ptr, ::_pbi::ParseContext* ctx) {
#define CHK_(x) if (PROTOBUF_PREDICT_FALSE(!(x))) goto failure
  while (!ctx->Done(&ptr)) {
    uint32_t tag;
    ptr = ::_pbi::ReadTag(ptr, &tag);
    switch (tag >> 3) {
      // repeated .greptime.v1.mito.Mutation mutations = 1;
      case 1:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 10)) {
          ptr -= 1;
          do {
            ptr += 1;
            ptr = ctx->ParseMessage(_internal_add_mutations(), ptr);
            CHK_(ptr);
            if (!ctx->DataAvailable(ptr)) break;
          } while (::PROTOBUF_NAMESPACE_ID::internal::ExpectTag<10>(ptr));
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

uint8_t* WalEntry::_InternalSerialize(
    uint8_t* target, ::PROTOBUF_NAMESPACE_ID::io::EpsCopyOutputStream* stream) const {
  // @@protoc_insertion_point(serialize_to_array_start:greptime.v1.mito.WalEntry)
  uint32_t cached_has_bits = 0;
  (void) cached_has_bits;

  // repeated .greptime.v1.mito.Mutation mutations = 1;
  for (unsigned i = 0,
      n = static_cast<unsigned>(this->_internal_mutations_size()); i < n; i++) {
    const auto& repfield = this->_internal_mutations(i);
    target = ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::
        InternalWriteMessage(1, repfield, repfield.GetCachedSize(), target, stream);
  }

  if (PROTOBUF_PREDICT_FALSE(_internal_metadata_.have_unknown_fields())) {
    target = ::_pbi::WireFormat::InternalSerializeUnknownFieldsToArray(
        _internal_metadata_.unknown_fields<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(::PROTOBUF_NAMESPACE_ID::UnknownFieldSet::default_instance), target, stream);
  }
  // @@protoc_insertion_point(serialize_to_array_end:greptime.v1.mito.WalEntry)
  return target;
}

size_t WalEntry::ByteSizeLong() const {
// @@protoc_insertion_point(message_byte_size_start:greptime.v1.mito.WalEntry)
  size_t total_size = 0;

  uint32_t cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  // repeated .greptime.v1.mito.Mutation mutations = 1;
  total_size += 1UL * this->_internal_mutations_size();
  for (const auto& msg : this->_impl_.mutations_) {
    total_size +=
      ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::MessageSize(msg);
  }

  return MaybeComputeUnknownFieldsSize(total_size, &_impl_._cached_size_);
}

const ::PROTOBUF_NAMESPACE_ID::Message::ClassData WalEntry::_class_data_ = {
    ::PROTOBUF_NAMESPACE_ID::Message::CopyWithSourceCheck,
    WalEntry::MergeImpl
};
const ::PROTOBUF_NAMESPACE_ID::Message::ClassData*WalEntry::GetClassData() const { return &_class_data_; }


void WalEntry::MergeImpl(::PROTOBUF_NAMESPACE_ID::Message& to_msg, const ::PROTOBUF_NAMESPACE_ID::Message& from_msg) {
  auto* const _this = static_cast<WalEntry*>(&to_msg);
  auto& from = static_cast<const WalEntry&>(from_msg);
  // @@protoc_insertion_point(class_specific_merge_from_start:greptime.v1.mito.WalEntry)
  GOOGLE_DCHECK_NE(&from, _this);
  uint32_t cached_has_bits = 0;
  (void) cached_has_bits;

  _this->_impl_.mutations_.MergeFrom(from._impl_.mutations_);
  _this->_internal_metadata_.MergeFrom<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(from._internal_metadata_);
}

void WalEntry::CopyFrom(const WalEntry& from) {
// @@protoc_insertion_point(class_specific_copy_from_start:greptime.v1.mito.WalEntry)
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

bool WalEntry::IsInitialized() const {
  return true;
}

void WalEntry::InternalSwap(WalEntry* other) {
  using std::swap;
  _internal_metadata_.InternalSwap(&other->_internal_metadata_);
  _impl_.mutations_.InternalSwap(&other->_impl_.mutations_);
}

::PROTOBUF_NAMESPACE_ID::Metadata WalEntry::GetMetadata() const {
  return ::_pbi::AssignDescriptors(
      &descriptor_table_greptime_2fv1_2fmito_2fwal_2eproto_getter, &descriptor_table_greptime_2fv1_2fmito_2fwal_2eproto_once,
      file_level_metadata_greptime_2fv1_2fmito_2fwal_2eproto[1]);
}

// @@protoc_insertion_point(namespace_scope)
}  // namespace mito
}  // namespace v1
}  // namespace greptime
PROTOBUF_NAMESPACE_OPEN
template<> PROTOBUF_NOINLINE ::greptime::v1::mito::Mutation*
Arena::CreateMaybeMessage< ::greptime::v1::mito::Mutation >(Arena* arena) {
  return Arena::CreateMessageInternal< ::greptime::v1::mito::Mutation >(arena);
}
template<> PROTOBUF_NOINLINE ::greptime::v1::mito::WalEntry*
Arena::CreateMaybeMessage< ::greptime::v1::mito::WalEntry >(Arena* arena) {
  return Arena::CreateMessageInternal< ::greptime::v1::mito::WalEntry >(arena);
}
PROTOBUF_NAMESPACE_CLOSE

// @@protoc_insertion_point(global_scope)
#include <google/protobuf/port_undef.inc>
