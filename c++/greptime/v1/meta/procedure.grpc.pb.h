// Generated by the gRPC C++ plugin.
// If you make any local change, they will be lost.
// source: greptime/v1/meta/procedure.proto
// Original file comments:
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
#ifndef GRPC_greptime_2fv1_2fmeta_2fprocedure_2eproto__INCLUDED
#define GRPC_greptime_2fv1_2fmeta_2fprocedure_2eproto__INCLUDED

#include "greptime/v1/meta/procedure.pb.h"

#include <functional>
#include <grpcpp/generic/async_generic_service.h>
#include <grpcpp/support/async_stream.h>
#include <grpcpp/support/async_unary_call.h>
#include <grpcpp/support/client_callback.h>
#include <grpcpp/client_context.h>
#include <grpcpp/completion_queue.h>
#include <grpcpp/support/message_allocator.h>
#include <grpcpp/support/method_handler.h>
#include <grpcpp/impl/codegen/proto_utils.h>
#include <grpcpp/impl/rpc_method.h>
#include <grpcpp/support/server_callback.h>
#include <grpcpp/impl/codegen/server_callback_handlers.h>
#include <grpcpp/server_context.h>
#include <grpcpp/impl/service_type.h>
#include <grpcpp/impl/codegen/status.h>
#include <grpcpp/support/stub_options.h>
#include <grpcpp/support/sync_stream.h>

namespace greptime {
namespace v1 {
namespace meta {

class Procedure final {
 public:
  static constexpr char const* service_full_name() {
    return "greptime.v1.meta.Procedure";
  }
  class StubInterface {
   public:
    virtual ~StubInterface() {}
    // Query a submitted procedure state
    virtual ::grpc::Status query(::grpc::ClientContext* context, const ::greptime::v1::meta::QueryProcedureRequest& request, ::greptime::v1::meta::ProcedureStateResponse* response) = 0;
    std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::greptime::v1::meta::ProcedureStateResponse>> Asyncquery(::grpc::ClientContext* context, const ::greptime::v1::meta::QueryProcedureRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::greptime::v1::meta::ProcedureStateResponse>>(AsyncqueryRaw(context, request, cq));
    }
    std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::greptime::v1::meta::ProcedureStateResponse>> PrepareAsyncquery(::grpc::ClientContext* context, const ::greptime::v1::meta::QueryProcedureRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::greptime::v1::meta::ProcedureStateResponse>>(PrepareAsyncqueryRaw(context, request, cq));
    }
    // Submits a DDL task
    virtual ::grpc::Status ddl(::grpc::ClientContext* context, const ::greptime::v1::meta::DdlTaskRequest& request, ::greptime::v1::meta::DdlTaskResponse* response) = 0;
    std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::greptime::v1::meta::DdlTaskResponse>> Asyncddl(::grpc::ClientContext* context, const ::greptime::v1::meta::DdlTaskRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::greptime::v1::meta::DdlTaskResponse>>(AsyncddlRaw(context, request, cq));
    }
    std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::greptime::v1::meta::DdlTaskResponse>> PrepareAsyncddl(::grpc::ClientContext* context, const ::greptime::v1::meta::DdlTaskRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::greptime::v1::meta::DdlTaskResponse>>(PrepareAsyncddlRaw(context, request, cq));
    }
    // Submits a region migration task
    virtual ::grpc::Status migrate(::grpc::ClientContext* context, const ::greptime::v1::meta::MigrateRegionRequest& request, ::greptime::v1::meta::MigrateRegionResponse* response) = 0;
    std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::greptime::v1::meta::MigrateRegionResponse>> Asyncmigrate(::grpc::ClientContext* context, const ::greptime::v1::meta::MigrateRegionRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::greptime::v1::meta::MigrateRegionResponse>>(AsyncmigrateRaw(context, request, cq));
    }
    std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::greptime::v1::meta::MigrateRegionResponse>> PrepareAsyncmigrate(::grpc::ClientContext* context, const ::greptime::v1::meta::MigrateRegionRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::greptime::v1::meta::MigrateRegionResponse>>(PrepareAsyncmigrateRaw(context, request, cq));
    }
    class async_interface {
     public:
      virtual ~async_interface() {}
      // Query a submitted procedure state
      virtual void query(::grpc::ClientContext* context, const ::greptime::v1::meta::QueryProcedureRequest* request, ::greptime::v1::meta::ProcedureStateResponse* response, std::function<void(::grpc::Status)>) = 0;
      virtual void query(::grpc::ClientContext* context, const ::greptime::v1::meta::QueryProcedureRequest* request, ::greptime::v1::meta::ProcedureStateResponse* response, ::grpc::ClientUnaryReactor* reactor) = 0;
      // Submits a DDL task
      virtual void ddl(::grpc::ClientContext* context, const ::greptime::v1::meta::DdlTaskRequest* request, ::greptime::v1::meta::DdlTaskResponse* response, std::function<void(::grpc::Status)>) = 0;
      virtual void ddl(::grpc::ClientContext* context, const ::greptime::v1::meta::DdlTaskRequest* request, ::greptime::v1::meta::DdlTaskResponse* response, ::grpc::ClientUnaryReactor* reactor) = 0;
      // Submits a region migration task
      virtual void migrate(::grpc::ClientContext* context, const ::greptime::v1::meta::MigrateRegionRequest* request, ::greptime::v1::meta::MigrateRegionResponse* response, std::function<void(::grpc::Status)>) = 0;
      virtual void migrate(::grpc::ClientContext* context, const ::greptime::v1::meta::MigrateRegionRequest* request, ::greptime::v1::meta::MigrateRegionResponse* response, ::grpc::ClientUnaryReactor* reactor) = 0;
    };
    typedef class async_interface experimental_async_interface;
    virtual class async_interface* async() { return nullptr; }
    class async_interface* experimental_async() { return async(); }
   private:
    virtual ::grpc::ClientAsyncResponseReaderInterface< ::greptime::v1::meta::ProcedureStateResponse>* AsyncqueryRaw(::grpc::ClientContext* context, const ::greptime::v1::meta::QueryProcedureRequest& request, ::grpc::CompletionQueue* cq) = 0;
    virtual ::grpc::ClientAsyncResponseReaderInterface< ::greptime::v1::meta::ProcedureStateResponse>* PrepareAsyncqueryRaw(::grpc::ClientContext* context, const ::greptime::v1::meta::QueryProcedureRequest& request, ::grpc::CompletionQueue* cq) = 0;
    virtual ::grpc::ClientAsyncResponseReaderInterface< ::greptime::v1::meta::DdlTaskResponse>* AsyncddlRaw(::grpc::ClientContext* context, const ::greptime::v1::meta::DdlTaskRequest& request, ::grpc::CompletionQueue* cq) = 0;
    virtual ::grpc::ClientAsyncResponseReaderInterface< ::greptime::v1::meta::DdlTaskResponse>* PrepareAsyncddlRaw(::grpc::ClientContext* context, const ::greptime::v1::meta::DdlTaskRequest& request, ::grpc::CompletionQueue* cq) = 0;
    virtual ::grpc::ClientAsyncResponseReaderInterface< ::greptime::v1::meta::MigrateRegionResponse>* AsyncmigrateRaw(::grpc::ClientContext* context, const ::greptime::v1::meta::MigrateRegionRequest& request, ::grpc::CompletionQueue* cq) = 0;
    virtual ::grpc::ClientAsyncResponseReaderInterface< ::greptime::v1::meta::MigrateRegionResponse>* PrepareAsyncmigrateRaw(::grpc::ClientContext* context, const ::greptime::v1::meta::MigrateRegionRequest& request, ::grpc::CompletionQueue* cq) = 0;
  };
  class Stub final : public StubInterface {
   public:
    Stub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options = ::grpc::StubOptions());
    ::grpc::Status query(::grpc::ClientContext* context, const ::greptime::v1::meta::QueryProcedureRequest& request, ::greptime::v1::meta::ProcedureStateResponse* response) override;
    std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::greptime::v1::meta::ProcedureStateResponse>> Asyncquery(::grpc::ClientContext* context, const ::greptime::v1::meta::QueryProcedureRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::greptime::v1::meta::ProcedureStateResponse>>(AsyncqueryRaw(context, request, cq));
    }
    std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::greptime::v1::meta::ProcedureStateResponse>> PrepareAsyncquery(::grpc::ClientContext* context, const ::greptime::v1::meta::QueryProcedureRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::greptime::v1::meta::ProcedureStateResponse>>(PrepareAsyncqueryRaw(context, request, cq));
    }
    ::grpc::Status ddl(::grpc::ClientContext* context, const ::greptime::v1::meta::DdlTaskRequest& request, ::greptime::v1::meta::DdlTaskResponse* response) override;
    std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::greptime::v1::meta::DdlTaskResponse>> Asyncddl(::grpc::ClientContext* context, const ::greptime::v1::meta::DdlTaskRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::greptime::v1::meta::DdlTaskResponse>>(AsyncddlRaw(context, request, cq));
    }
    std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::greptime::v1::meta::DdlTaskResponse>> PrepareAsyncddl(::grpc::ClientContext* context, const ::greptime::v1::meta::DdlTaskRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::greptime::v1::meta::DdlTaskResponse>>(PrepareAsyncddlRaw(context, request, cq));
    }
    ::grpc::Status migrate(::grpc::ClientContext* context, const ::greptime::v1::meta::MigrateRegionRequest& request, ::greptime::v1::meta::MigrateRegionResponse* response) override;
    std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::greptime::v1::meta::MigrateRegionResponse>> Asyncmigrate(::grpc::ClientContext* context, const ::greptime::v1::meta::MigrateRegionRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::greptime::v1::meta::MigrateRegionResponse>>(AsyncmigrateRaw(context, request, cq));
    }
    std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::greptime::v1::meta::MigrateRegionResponse>> PrepareAsyncmigrate(::grpc::ClientContext* context, const ::greptime::v1::meta::MigrateRegionRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::greptime::v1::meta::MigrateRegionResponse>>(PrepareAsyncmigrateRaw(context, request, cq));
    }
    class async final :
      public StubInterface::async_interface {
     public:
      void query(::grpc::ClientContext* context, const ::greptime::v1::meta::QueryProcedureRequest* request, ::greptime::v1::meta::ProcedureStateResponse* response, std::function<void(::grpc::Status)>) override;
      void query(::grpc::ClientContext* context, const ::greptime::v1::meta::QueryProcedureRequest* request, ::greptime::v1::meta::ProcedureStateResponse* response, ::grpc::ClientUnaryReactor* reactor) override;
      void ddl(::grpc::ClientContext* context, const ::greptime::v1::meta::DdlTaskRequest* request, ::greptime::v1::meta::DdlTaskResponse* response, std::function<void(::grpc::Status)>) override;
      void ddl(::grpc::ClientContext* context, const ::greptime::v1::meta::DdlTaskRequest* request, ::greptime::v1::meta::DdlTaskResponse* response, ::grpc::ClientUnaryReactor* reactor) override;
      void migrate(::grpc::ClientContext* context, const ::greptime::v1::meta::MigrateRegionRequest* request, ::greptime::v1::meta::MigrateRegionResponse* response, std::function<void(::grpc::Status)>) override;
      void migrate(::grpc::ClientContext* context, const ::greptime::v1::meta::MigrateRegionRequest* request, ::greptime::v1::meta::MigrateRegionResponse* response, ::grpc::ClientUnaryReactor* reactor) override;
     private:
      friend class Stub;
      explicit async(Stub* stub): stub_(stub) { }
      Stub* stub() { return stub_; }
      Stub* stub_;
    };
    class async* async() override { return &async_stub_; }

   private:
    std::shared_ptr< ::grpc::ChannelInterface> channel_;
    class async async_stub_{this};
    ::grpc::ClientAsyncResponseReader< ::greptime::v1::meta::ProcedureStateResponse>* AsyncqueryRaw(::grpc::ClientContext* context, const ::greptime::v1::meta::QueryProcedureRequest& request, ::grpc::CompletionQueue* cq) override;
    ::grpc::ClientAsyncResponseReader< ::greptime::v1::meta::ProcedureStateResponse>* PrepareAsyncqueryRaw(::grpc::ClientContext* context, const ::greptime::v1::meta::QueryProcedureRequest& request, ::grpc::CompletionQueue* cq) override;
    ::grpc::ClientAsyncResponseReader< ::greptime::v1::meta::DdlTaskResponse>* AsyncddlRaw(::grpc::ClientContext* context, const ::greptime::v1::meta::DdlTaskRequest& request, ::grpc::CompletionQueue* cq) override;
    ::grpc::ClientAsyncResponseReader< ::greptime::v1::meta::DdlTaskResponse>* PrepareAsyncddlRaw(::grpc::ClientContext* context, const ::greptime::v1::meta::DdlTaskRequest& request, ::grpc::CompletionQueue* cq) override;
    ::grpc::ClientAsyncResponseReader< ::greptime::v1::meta::MigrateRegionResponse>* AsyncmigrateRaw(::grpc::ClientContext* context, const ::greptime::v1::meta::MigrateRegionRequest& request, ::grpc::CompletionQueue* cq) override;
    ::grpc::ClientAsyncResponseReader< ::greptime::v1::meta::MigrateRegionResponse>* PrepareAsyncmigrateRaw(::grpc::ClientContext* context, const ::greptime::v1::meta::MigrateRegionRequest& request, ::grpc::CompletionQueue* cq) override;
    const ::grpc::internal::RpcMethod rpcmethod_query_;
    const ::grpc::internal::RpcMethod rpcmethod_ddl_;
    const ::grpc::internal::RpcMethod rpcmethod_migrate_;
  };
  static std::unique_ptr<Stub> NewStub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options = ::grpc::StubOptions());

  class Service : public ::grpc::Service {
   public:
    Service();
    virtual ~Service();
    // Query a submitted procedure state
    virtual ::grpc::Status query(::grpc::ServerContext* context, const ::greptime::v1::meta::QueryProcedureRequest* request, ::greptime::v1::meta::ProcedureStateResponse* response);
    // Submits a DDL task
    virtual ::grpc::Status ddl(::grpc::ServerContext* context, const ::greptime::v1::meta::DdlTaskRequest* request, ::greptime::v1::meta::DdlTaskResponse* response);
    // Submits a region migration task
    virtual ::grpc::Status migrate(::grpc::ServerContext* context, const ::greptime::v1::meta::MigrateRegionRequest* request, ::greptime::v1::meta::MigrateRegionResponse* response);
  };
  template <class BaseClass>
  class WithAsyncMethod_query : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithAsyncMethod_query() {
      ::grpc::Service::MarkMethodAsync(0);
    }
    ~WithAsyncMethod_query() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status query(::grpc::ServerContext* /*context*/, const ::greptime::v1::meta::QueryProcedureRequest* /*request*/, ::greptime::v1::meta::ProcedureStateResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    void Requestquery(::grpc::ServerContext* context, ::greptime::v1::meta::QueryProcedureRequest* request, ::grpc::ServerAsyncResponseWriter< ::greptime::v1::meta::ProcedureStateResponse>* response, ::grpc::CompletionQueue* new_call_cq, ::grpc::ServerCompletionQueue* notification_cq, void *tag) {
      ::grpc::Service::RequestAsyncUnary(0, context, request, response, new_call_cq, notification_cq, tag);
    }
  };
  template <class BaseClass>
  class WithAsyncMethod_ddl : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithAsyncMethod_ddl() {
      ::grpc::Service::MarkMethodAsync(1);
    }
    ~WithAsyncMethod_ddl() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status ddl(::grpc::ServerContext* /*context*/, const ::greptime::v1::meta::DdlTaskRequest* /*request*/, ::greptime::v1::meta::DdlTaskResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    void Requestddl(::grpc::ServerContext* context, ::greptime::v1::meta::DdlTaskRequest* request, ::grpc::ServerAsyncResponseWriter< ::greptime::v1::meta::DdlTaskResponse>* response, ::grpc::CompletionQueue* new_call_cq, ::grpc::ServerCompletionQueue* notification_cq, void *tag) {
      ::grpc::Service::RequestAsyncUnary(1, context, request, response, new_call_cq, notification_cq, tag);
    }
  };
  template <class BaseClass>
  class WithAsyncMethod_migrate : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithAsyncMethod_migrate() {
      ::grpc::Service::MarkMethodAsync(2);
    }
    ~WithAsyncMethod_migrate() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status migrate(::grpc::ServerContext* /*context*/, const ::greptime::v1::meta::MigrateRegionRequest* /*request*/, ::greptime::v1::meta::MigrateRegionResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    void Requestmigrate(::grpc::ServerContext* context, ::greptime::v1::meta::MigrateRegionRequest* request, ::grpc::ServerAsyncResponseWriter< ::greptime::v1::meta::MigrateRegionResponse>* response, ::grpc::CompletionQueue* new_call_cq, ::grpc::ServerCompletionQueue* notification_cq, void *tag) {
      ::grpc::Service::RequestAsyncUnary(2, context, request, response, new_call_cq, notification_cq, tag);
    }
  };
  typedef WithAsyncMethod_query<WithAsyncMethod_ddl<WithAsyncMethod_migrate<Service > > > AsyncService;
  template <class BaseClass>
  class WithCallbackMethod_query : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithCallbackMethod_query() {
      ::grpc::Service::MarkMethodCallback(0,
          new ::grpc::internal::CallbackUnaryHandler< ::greptime::v1::meta::QueryProcedureRequest, ::greptime::v1::meta::ProcedureStateResponse>(
            [this](
                   ::grpc::CallbackServerContext* context, const ::greptime::v1::meta::QueryProcedureRequest* request, ::greptime::v1::meta::ProcedureStateResponse* response) { return this->query(context, request, response); }));}
    void SetMessageAllocatorFor_query(
        ::grpc::MessageAllocator< ::greptime::v1::meta::QueryProcedureRequest, ::greptime::v1::meta::ProcedureStateResponse>* allocator) {
      ::grpc::internal::MethodHandler* const handler = ::grpc::Service::GetHandler(0);
      static_cast<::grpc::internal::CallbackUnaryHandler< ::greptime::v1::meta::QueryProcedureRequest, ::greptime::v1::meta::ProcedureStateResponse>*>(handler)
              ->SetMessageAllocator(allocator);
    }
    ~WithCallbackMethod_query() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status query(::grpc::ServerContext* /*context*/, const ::greptime::v1::meta::QueryProcedureRequest* /*request*/, ::greptime::v1::meta::ProcedureStateResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    virtual ::grpc::ServerUnaryReactor* query(
      ::grpc::CallbackServerContext* /*context*/, const ::greptime::v1::meta::QueryProcedureRequest* /*request*/, ::greptime::v1::meta::ProcedureStateResponse* /*response*/)  { return nullptr; }
  };
  template <class BaseClass>
  class WithCallbackMethod_ddl : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithCallbackMethod_ddl() {
      ::grpc::Service::MarkMethodCallback(1,
          new ::grpc::internal::CallbackUnaryHandler< ::greptime::v1::meta::DdlTaskRequest, ::greptime::v1::meta::DdlTaskResponse>(
            [this](
                   ::grpc::CallbackServerContext* context, const ::greptime::v1::meta::DdlTaskRequest* request, ::greptime::v1::meta::DdlTaskResponse* response) { return this->ddl(context, request, response); }));}
    void SetMessageAllocatorFor_ddl(
        ::grpc::MessageAllocator< ::greptime::v1::meta::DdlTaskRequest, ::greptime::v1::meta::DdlTaskResponse>* allocator) {
      ::grpc::internal::MethodHandler* const handler = ::grpc::Service::GetHandler(1);
      static_cast<::grpc::internal::CallbackUnaryHandler< ::greptime::v1::meta::DdlTaskRequest, ::greptime::v1::meta::DdlTaskResponse>*>(handler)
              ->SetMessageAllocator(allocator);
    }
    ~WithCallbackMethod_ddl() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status ddl(::grpc::ServerContext* /*context*/, const ::greptime::v1::meta::DdlTaskRequest* /*request*/, ::greptime::v1::meta::DdlTaskResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    virtual ::grpc::ServerUnaryReactor* ddl(
      ::grpc::CallbackServerContext* /*context*/, const ::greptime::v1::meta::DdlTaskRequest* /*request*/, ::greptime::v1::meta::DdlTaskResponse* /*response*/)  { return nullptr; }
  };
  template <class BaseClass>
  class WithCallbackMethod_migrate : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithCallbackMethod_migrate() {
      ::grpc::Service::MarkMethodCallback(2,
          new ::grpc::internal::CallbackUnaryHandler< ::greptime::v1::meta::MigrateRegionRequest, ::greptime::v1::meta::MigrateRegionResponse>(
            [this](
                   ::grpc::CallbackServerContext* context, const ::greptime::v1::meta::MigrateRegionRequest* request, ::greptime::v1::meta::MigrateRegionResponse* response) { return this->migrate(context, request, response); }));}
    void SetMessageAllocatorFor_migrate(
        ::grpc::MessageAllocator< ::greptime::v1::meta::MigrateRegionRequest, ::greptime::v1::meta::MigrateRegionResponse>* allocator) {
      ::grpc::internal::MethodHandler* const handler = ::grpc::Service::GetHandler(2);
      static_cast<::grpc::internal::CallbackUnaryHandler< ::greptime::v1::meta::MigrateRegionRequest, ::greptime::v1::meta::MigrateRegionResponse>*>(handler)
              ->SetMessageAllocator(allocator);
    }
    ~WithCallbackMethod_migrate() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status migrate(::grpc::ServerContext* /*context*/, const ::greptime::v1::meta::MigrateRegionRequest* /*request*/, ::greptime::v1::meta::MigrateRegionResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    virtual ::grpc::ServerUnaryReactor* migrate(
      ::grpc::CallbackServerContext* /*context*/, const ::greptime::v1::meta::MigrateRegionRequest* /*request*/, ::greptime::v1::meta::MigrateRegionResponse* /*response*/)  { return nullptr; }
  };
  typedef WithCallbackMethod_query<WithCallbackMethod_ddl<WithCallbackMethod_migrate<Service > > > CallbackService;
  typedef CallbackService ExperimentalCallbackService;
  template <class BaseClass>
  class WithGenericMethod_query : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithGenericMethod_query() {
      ::grpc::Service::MarkMethodGeneric(0);
    }
    ~WithGenericMethod_query() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status query(::grpc::ServerContext* /*context*/, const ::greptime::v1::meta::QueryProcedureRequest* /*request*/, ::greptime::v1::meta::ProcedureStateResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
  };
  template <class BaseClass>
  class WithGenericMethod_ddl : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithGenericMethod_ddl() {
      ::grpc::Service::MarkMethodGeneric(1);
    }
    ~WithGenericMethod_ddl() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status ddl(::grpc::ServerContext* /*context*/, const ::greptime::v1::meta::DdlTaskRequest* /*request*/, ::greptime::v1::meta::DdlTaskResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
  };
  template <class BaseClass>
  class WithGenericMethod_migrate : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithGenericMethod_migrate() {
      ::grpc::Service::MarkMethodGeneric(2);
    }
    ~WithGenericMethod_migrate() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status migrate(::grpc::ServerContext* /*context*/, const ::greptime::v1::meta::MigrateRegionRequest* /*request*/, ::greptime::v1::meta::MigrateRegionResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
  };
  template <class BaseClass>
  class WithRawMethod_query : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithRawMethod_query() {
      ::grpc::Service::MarkMethodRaw(0);
    }
    ~WithRawMethod_query() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status query(::grpc::ServerContext* /*context*/, const ::greptime::v1::meta::QueryProcedureRequest* /*request*/, ::greptime::v1::meta::ProcedureStateResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    void Requestquery(::grpc::ServerContext* context, ::grpc::ByteBuffer* request, ::grpc::ServerAsyncResponseWriter< ::grpc::ByteBuffer>* response, ::grpc::CompletionQueue* new_call_cq, ::grpc::ServerCompletionQueue* notification_cq, void *tag) {
      ::grpc::Service::RequestAsyncUnary(0, context, request, response, new_call_cq, notification_cq, tag);
    }
  };
  template <class BaseClass>
  class WithRawMethod_ddl : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithRawMethod_ddl() {
      ::grpc::Service::MarkMethodRaw(1);
    }
    ~WithRawMethod_ddl() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status ddl(::grpc::ServerContext* /*context*/, const ::greptime::v1::meta::DdlTaskRequest* /*request*/, ::greptime::v1::meta::DdlTaskResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    void Requestddl(::grpc::ServerContext* context, ::grpc::ByteBuffer* request, ::grpc::ServerAsyncResponseWriter< ::grpc::ByteBuffer>* response, ::grpc::CompletionQueue* new_call_cq, ::grpc::ServerCompletionQueue* notification_cq, void *tag) {
      ::grpc::Service::RequestAsyncUnary(1, context, request, response, new_call_cq, notification_cq, tag);
    }
  };
  template <class BaseClass>
  class WithRawMethod_migrate : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithRawMethod_migrate() {
      ::grpc::Service::MarkMethodRaw(2);
    }
    ~WithRawMethod_migrate() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status migrate(::grpc::ServerContext* /*context*/, const ::greptime::v1::meta::MigrateRegionRequest* /*request*/, ::greptime::v1::meta::MigrateRegionResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    void Requestmigrate(::grpc::ServerContext* context, ::grpc::ByteBuffer* request, ::grpc::ServerAsyncResponseWriter< ::grpc::ByteBuffer>* response, ::grpc::CompletionQueue* new_call_cq, ::grpc::ServerCompletionQueue* notification_cq, void *tag) {
      ::grpc::Service::RequestAsyncUnary(2, context, request, response, new_call_cq, notification_cq, tag);
    }
  };
  template <class BaseClass>
  class WithRawCallbackMethod_query : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithRawCallbackMethod_query() {
      ::grpc::Service::MarkMethodRawCallback(0,
          new ::grpc::internal::CallbackUnaryHandler< ::grpc::ByteBuffer, ::grpc::ByteBuffer>(
            [this](
                   ::grpc::CallbackServerContext* context, const ::grpc::ByteBuffer* request, ::grpc::ByteBuffer* response) { return this->query(context, request, response); }));
    }
    ~WithRawCallbackMethod_query() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status query(::grpc::ServerContext* /*context*/, const ::greptime::v1::meta::QueryProcedureRequest* /*request*/, ::greptime::v1::meta::ProcedureStateResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    virtual ::grpc::ServerUnaryReactor* query(
      ::grpc::CallbackServerContext* /*context*/, const ::grpc::ByteBuffer* /*request*/, ::grpc::ByteBuffer* /*response*/)  { return nullptr; }
  };
  template <class BaseClass>
  class WithRawCallbackMethod_ddl : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithRawCallbackMethod_ddl() {
      ::grpc::Service::MarkMethodRawCallback(1,
          new ::grpc::internal::CallbackUnaryHandler< ::grpc::ByteBuffer, ::grpc::ByteBuffer>(
            [this](
                   ::grpc::CallbackServerContext* context, const ::grpc::ByteBuffer* request, ::grpc::ByteBuffer* response) { return this->ddl(context, request, response); }));
    }
    ~WithRawCallbackMethod_ddl() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status ddl(::grpc::ServerContext* /*context*/, const ::greptime::v1::meta::DdlTaskRequest* /*request*/, ::greptime::v1::meta::DdlTaskResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    virtual ::grpc::ServerUnaryReactor* ddl(
      ::grpc::CallbackServerContext* /*context*/, const ::grpc::ByteBuffer* /*request*/, ::grpc::ByteBuffer* /*response*/)  { return nullptr; }
  };
  template <class BaseClass>
  class WithRawCallbackMethod_migrate : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithRawCallbackMethod_migrate() {
      ::grpc::Service::MarkMethodRawCallback(2,
          new ::grpc::internal::CallbackUnaryHandler< ::grpc::ByteBuffer, ::grpc::ByteBuffer>(
            [this](
                   ::grpc::CallbackServerContext* context, const ::grpc::ByteBuffer* request, ::grpc::ByteBuffer* response) { return this->migrate(context, request, response); }));
    }
    ~WithRawCallbackMethod_migrate() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status migrate(::grpc::ServerContext* /*context*/, const ::greptime::v1::meta::MigrateRegionRequest* /*request*/, ::greptime::v1::meta::MigrateRegionResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    virtual ::grpc::ServerUnaryReactor* migrate(
      ::grpc::CallbackServerContext* /*context*/, const ::grpc::ByteBuffer* /*request*/, ::grpc::ByteBuffer* /*response*/)  { return nullptr; }
  };
  template <class BaseClass>
  class WithStreamedUnaryMethod_query : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithStreamedUnaryMethod_query() {
      ::grpc::Service::MarkMethodStreamed(0,
        new ::grpc::internal::StreamedUnaryHandler<
          ::greptime::v1::meta::QueryProcedureRequest, ::greptime::v1::meta::ProcedureStateResponse>(
            [this](::grpc::ServerContext* context,
                   ::grpc::ServerUnaryStreamer<
                     ::greptime::v1::meta::QueryProcedureRequest, ::greptime::v1::meta::ProcedureStateResponse>* streamer) {
                       return this->Streamedquery(context,
                         streamer);
                  }));
    }
    ~WithStreamedUnaryMethod_query() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable regular version of this method
    ::grpc::Status query(::grpc::ServerContext* /*context*/, const ::greptime::v1::meta::QueryProcedureRequest* /*request*/, ::greptime::v1::meta::ProcedureStateResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    // replace default version of method with streamed unary
    virtual ::grpc::Status Streamedquery(::grpc::ServerContext* context, ::grpc::ServerUnaryStreamer< ::greptime::v1::meta::QueryProcedureRequest,::greptime::v1::meta::ProcedureStateResponse>* server_unary_streamer) = 0;
  };
  template <class BaseClass>
  class WithStreamedUnaryMethod_ddl : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithStreamedUnaryMethod_ddl() {
      ::grpc::Service::MarkMethodStreamed(1,
        new ::grpc::internal::StreamedUnaryHandler<
          ::greptime::v1::meta::DdlTaskRequest, ::greptime::v1::meta::DdlTaskResponse>(
            [this](::grpc::ServerContext* context,
                   ::grpc::ServerUnaryStreamer<
                     ::greptime::v1::meta::DdlTaskRequest, ::greptime::v1::meta::DdlTaskResponse>* streamer) {
                       return this->Streamedddl(context,
                         streamer);
                  }));
    }
    ~WithStreamedUnaryMethod_ddl() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable regular version of this method
    ::grpc::Status ddl(::grpc::ServerContext* /*context*/, const ::greptime::v1::meta::DdlTaskRequest* /*request*/, ::greptime::v1::meta::DdlTaskResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    // replace default version of method with streamed unary
    virtual ::grpc::Status Streamedddl(::grpc::ServerContext* context, ::grpc::ServerUnaryStreamer< ::greptime::v1::meta::DdlTaskRequest,::greptime::v1::meta::DdlTaskResponse>* server_unary_streamer) = 0;
  };
  template <class BaseClass>
  class WithStreamedUnaryMethod_migrate : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithStreamedUnaryMethod_migrate() {
      ::grpc::Service::MarkMethodStreamed(2,
        new ::grpc::internal::StreamedUnaryHandler<
          ::greptime::v1::meta::MigrateRegionRequest, ::greptime::v1::meta::MigrateRegionResponse>(
            [this](::grpc::ServerContext* context,
                   ::grpc::ServerUnaryStreamer<
                     ::greptime::v1::meta::MigrateRegionRequest, ::greptime::v1::meta::MigrateRegionResponse>* streamer) {
                       return this->Streamedmigrate(context,
                         streamer);
                  }));
    }
    ~WithStreamedUnaryMethod_migrate() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable regular version of this method
    ::grpc::Status migrate(::grpc::ServerContext* /*context*/, const ::greptime::v1::meta::MigrateRegionRequest* /*request*/, ::greptime::v1::meta::MigrateRegionResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    // replace default version of method with streamed unary
    virtual ::grpc::Status Streamedmigrate(::grpc::ServerContext* context, ::grpc::ServerUnaryStreamer< ::greptime::v1::meta::MigrateRegionRequest,::greptime::v1::meta::MigrateRegionResponse>* server_unary_streamer) = 0;
  };
  typedef WithStreamedUnaryMethod_query<WithStreamedUnaryMethod_ddl<WithStreamedUnaryMethod_migrate<Service > > > StreamedUnaryService;
  typedef Service SplitStreamedService;
  typedef WithStreamedUnaryMethod_query<WithStreamedUnaryMethod_ddl<WithStreamedUnaryMethod_migrate<Service > > > StreamedService;
};

}  // namespace meta
}  // namespace v1
}  // namespace greptime


#endif  // GRPC_greptime_2fv1_2fmeta_2fprocedure_2eproto__INCLUDED