// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: greptime/v1/meta/cluster.proto

#include "greptime/v1/meta/cluster.pb.h"

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
}  // namespace meta
}  // namespace v1
}  // namespace greptime
static constexpr ::_pb::EnumDescriptor const** file_level_enum_descriptors_greptime_2fv1_2fmeta_2fcluster_2eproto = nullptr;
static constexpr ::_pb::ServiceDescriptor const** file_level_service_descriptors_greptime_2fv1_2fmeta_2fcluster_2eproto = nullptr;
const uint32_t TableStruct_greptime_2fv1_2fmeta_2fcluster_2eproto::offsets[1] = {};
static constexpr ::_pbi::MigrationSchema* schemas = nullptr;
static constexpr ::_pb::Message* const* file_default_instances = nullptr;

const char descriptor_table_protodef_greptime_2fv1_2fmeta_2fcluster_2eproto[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) =
  "\n\036greptime/v1/meta/cluster.proto\022\020grepti"
  "me.v1.meta\032\034greptime/v1/meta/store.proto"
  "2\246\001\n\007Cluster\022Q\n\010BatchGet\022!.greptime.v1.m"
  "eta.BatchGetRequest\032\".greptime.v1.meta.B"
  "atchGetResponse\022H\n\005Range\022\036.greptime.v1.m"
  "eta.RangeRequest\032\037.greptime.v1.meta.Rang"
  "eResponseB<Z:github.com/GreptimeTeam/gre"
  "ptime-proto/go/greptime/v1/metab\006proto3"
  ;
static const ::_pbi::DescriptorTable* const descriptor_table_greptime_2fv1_2fmeta_2fcluster_2eproto_deps[1] = {
  &::descriptor_table_greptime_2fv1_2fmeta_2fstore_2eproto,
};
static ::_pbi::once_flag descriptor_table_greptime_2fv1_2fmeta_2fcluster_2eproto_once;
const ::_pbi::DescriptorTable descriptor_table_greptime_2fv1_2fmeta_2fcluster_2eproto = {
    false, false, 319, descriptor_table_protodef_greptime_2fv1_2fmeta_2fcluster_2eproto,
    "greptime/v1/meta/cluster.proto",
    &descriptor_table_greptime_2fv1_2fmeta_2fcluster_2eproto_once, descriptor_table_greptime_2fv1_2fmeta_2fcluster_2eproto_deps, 1, 0,
    schemas, file_default_instances, TableStruct_greptime_2fv1_2fmeta_2fcluster_2eproto::offsets,
    nullptr, file_level_enum_descriptors_greptime_2fv1_2fmeta_2fcluster_2eproto,
    file_level_service_descriptors_greptime_2fv1_2fmeta_2fcluster_2eproto,
};
PROTOBUF_ATTRIBUTE_WEAK const ::_pbi::DescriptorTable* descriptor_table_greptime_2fv1_2fmeta_2fcluster_2eproto_getter() {
  return &descriptor_table_greptime_2fv1_2fmeta_2fcluster_2eproto;
}

// Force running AddDescriptors() at dynamic initialization time.
PROTOBUF_ATTRIBUTE_INIT_PRIORITY2 static ::_pbi::AddDescriptorsRunner dynamic_init_dummy_greptime_2fv1_2fmeta_2fcluster_2eproto(&descriptor_table_greptime_2fv1_2fmeta_2fcluster_2eproto);
namespace greptime {
namespace v1 {
namespace meta {

// @@protoc_insertion_point(namespace_scope)
}  // namespace meta
}  // namespace v1
}  // namespace greptime
PROTOBUF_NAMESPACE_OPEN
PROTOBUF_NAMESPACE_CLOSE

// @@protoc_insertion_point(global_scope)
#include <google/protobuf/port_undef.inc>
