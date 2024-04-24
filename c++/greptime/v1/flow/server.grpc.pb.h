// Generated by the gRPC C++ plugin.
// If you make any local change, they will be lost.
// source: greptime/v1/flow/server.proto
// Original file comments:
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
//
#ifndef GRPC_greptime_2fv1_2fflow_2fserver_2eproto__INCLUDED
#define GRPC_greptime_2fv1_2fflow_2fserver_2eproto__INCLUDED

#include "greptime/v1/flow/server.pb.h"

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
namespace flow {

class Flow final {
 public:
  static constexpr char const* service_full_name() {
    return "greptime.v1.flow.Flow";
  }
  class StubInterface {
   public:
    virtual ~StubInterface() {}
    // Handle the control plane request for creating or removing a flow.
    virtual ::grpc::Status HandleCreateRemove(::grpc::ClientContext* context, const ::greptime::v1::flow::FlowRequest& request, ::greptime::v1::flow::FlowResponse* response) = 0;
    std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::greptime::v1::flow::FlowResponse>> AsyncHandleCreateRemove(::grpc::ClientContext* context, const ::greptime::v1::flow::FlowRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::greptime::v1::flow::FlowResponse>>(AsyncHandleCreateRemoveRaw(context, request, cq));
    }
    std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::greptime::v1::flow::FlowResponse>> PrepareAsyncHandleCreateRemove(::grpc::ClientContext* context, const ::greptime::v1::flow::FlowRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::greptime::v1::flow::FlowResponse>>(PrepareAsyncHandleCreateRemoveRaw(context, request, cq));
    }
    // Handle the data plane request for inserting or deleting rows
    // only expect `RegionRequest` to be one of `InsertRequests` or
    // `DeleteRequests` other types of `RegionRequest` will be ignored
    virtual ::grpc::Status HandleMirrorRequest(::grpc::ClientContext* context, const ::greptime::v1::flow::InsertRequests& request, ::greptime::v1::flow::FlowResponse* response) = 0;
    std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::greptime::v1::flow::FlowResponse>> AsyncHandleMirrorRequest(::grpc::ClientContext* context, const ::greptime::v1::flow::InsertRequests& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::greptime::v1::flow::FlowResponse>>(AsyncHandleMirrorRequestRaw(context, request, cq));
    }
    std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::greptime::v1::flow::FlowResponse>> PrepareAsyncHandleMirrorRequest(::grpc::ClientContext* context, const ::greptime::v1::flow::InsertRequests& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::greptime::v1::flow::FlowResponse>>(PrepareAsyncHandleMirrorRequestRaw(context, request, cq));
    }
    class async_interface {
     public:
      virtual ~async_interface() {}
      // Handle the control plane request for creating or removing a flow.
      virtual void HandleCreateRemove(::grpc::ClientContext* context, const ::greptime::v1::flow::FlowRequest* request, ::greptime::v1::flow::FlowResponse* response, std::function<void(::grpc::Status)>) = 0;
      virtual void HandleCreateRemove(::grpc::ClientContext* context, const ::greptime::v1::flow::FlowRequest* request, ::greptime::v1::flow::FlowResponse* response, ::grpc::ClientUnaryReactor* reactor) = 0;
      // Handle the data plane request for inserting or deleting rows
      // only expect `RegionRequest` to be one of `InsertRequests` or
      // `DeleteRequests` other types of `RegionRequest` will be ignored
      virtual void HandleMirrorRequest(::grpc::ClientContext* context, const ::greptime::v1::flow::InsertRequests* request, ::greptime::v1::flow::FlowResponse* response, std::function<void(::grpc::Status)>) = 0;
      virtual void HandleMirrorRequest(::grpc::ClientContext* context, const ::greptime::v1::flow::InsertRequests* request, ::greptime::v1::flow::FlowResponse* response, ::grpc::ClientUnaryReactor* reactor) = 0;
    };
    typedef class async_interface experimental_async_interface;
    virtual class async_interface* async() { return nullptr; }
    class async_interface* experimental_async() { return async(); }
   private:
    virtual ::grpc::ClientAsyncResponseReaderInterface< ::greptime::v1::flow::FlowResponse>* AsyncHandleCreateRemoveRaw(::grpc::ClientContext* context, const ::greptime::v1::flow::FlowRequest& request, ::grpc::CompletionQueue* cq) = 0;
    virtual ::grpc::ClientAsyncResponseReaderInterface< ::greptime::v1::flow::FlowResponse>* PrepareAsyncHandleCreateRemoveRaw(::grpc::ClientContext* context, const ::greptime::v1::flow::FlowRequest& request, ::grpc::CompletionQueue* cq) = 0;
    virtual ::grpc::ClientAsyncResponseReaderInterface< ::greptime::v1::flow::FlowResponse>* AsyncHandleMirrorRequestRaw(::grpc::ClientContext* context, const ::greptime::v1::flow::InsertRequests& request, ::grpc::CompletionQueue* cq) = 0;
    virtual ::grpc::ClientAsyncResponseReaderInterface< ::greptime::v1::flow::FlowResponse>* PrepareAsyncHandleMirrorRequestRaw(::grpc::ClientContext* context, const ::greptime::v1::flow::InsertRequests& request, ::grpc::CompletionQueue* cq) = 0;
  };
  class Stub final : public StubInterface {
   public:
    Stub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options = ::grpc::StubOptions());
    ::grpc::Status HandleCreateRemove(::grpc::ClientContext* context, const ::greptime::v1::flow::FlowRequest& request, ::greptime::v1::flow::FlowResponse* response) override;
    std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::greptime::v1::flow::FlowResponse>> AsyncHandleCreateRemove(::grpc::ClientContext* context, const ::greptime::v1::flow::FlowRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::greptime::v1::flow::FlowResponse>>(AsyncHandleCreateRemoveRaw(context, request, cq));
    }
    std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::greptime::v1::flow::FlowResponse>> PrepareAsyncHandleCreateRemove(::grpc::ClientContext* context, const ::greptime::v1::flow::FlowRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::greptime::v1::flow::FlowResponse>>(PrepareAsyncHandleCreateRemoveRaw(context, request, cq));
    }
    ::grpc::Status HandleMirrorRequest(::grpc::ClientContext* context, const ::greptime::v1::flow::InsertRequests& request, ::greptime::v1::flow::FlowResponse* response) override;
    std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::greptime::v1::flow::FlowResponse>> AsyncHandleMirrorRequest(::grpc::ClientContext* context, const ::greptime::v1::flow::InsertRequests& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::greptime::v1::flow::FlowResponse>>(AsyncHandleMirrorRequestRaw(context, request, cq));
    }
    std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::greptime::v1::flow::FlowResponse>> PrepareAsyncHandleMirrorRequest(::grpc::ClientContext* context, const ::greptime::v1::flow::InsertRequests& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::greptime::v1::flow::FlowResponse>>(PrepareAsyncHandleMirrorRequestRaw(context, request, cq));
    }
    class async final :
      public StubInterface::async_interface {
     public:
      void HandleCreateRemove(::grpc::ClientContext* context, const ::greptime::v1::flow::FlowRequest* request, ::greptime::v1::flow::FlowResponse* response, std::function<void(::grpc::Status)>) override;
      void HandleCreateRemove(::grpc::ClientContext* context, const ::greptime::v1::flow::FlowRequest* request, ::greptime::v1::flow::FlowResponse* response, ::grpc::ClientUnaryReactor* reactor) override;
      void HandleMirrorRequest(::grpc::ClientContext* context, const ::greptime::v1::flow::InsertRequests* request, ::greptime::v1::flow::FlowResponse* response, std::function<void(::grpc::Status)>) override;
      void HandleMirrorRequest(::grpc::ClientContext* context, const ::greptime::v1::flow::InsertRequests* request, ::greptime::v1::flow::FlowResponse* response, ::grpc::ClientUnaryReactor* reactor) override;
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
    ::grpc::ClientAsyncResponseReader< ::greptime::v1::flow::FlowResponse>* AsyncHandleCreateRemoveRaw(::grpc::ClientContext* context, const ::greptime::v1::flow::FlowRequest& request, ::grpc::CompletionQueue* cq) override;
    ::grpc::ClientAsyncResponseReader< ::greptime::v1::flow::FlowResponse>* PrepareAsyncHandleCreateRemoveRaw(::grpc::ClientContext* context, const ::greptime::v1::flow::FlowRequest& request, ::grpc::CompletionQueue* cq) override;
    ::grpc::ClientAsyncResponseReader< ::greptime::v1::flow::FlowResponse>* AsyncHandleMirrorRequestRaw(::grpc::ClientContext* context, const ::greptime::v1::flow::InsertRequests& request, ::grpc::CompletionQueue* cq) override;
    ::grpc::ClientAsyncResponseReader< ::greptime::v1::flow::FlowResponse>* PrepareAsyncHandleMirrorRequestRaw(::grpc::ClientContext* context, const ::greptime::v1::flow::InsertRequests& request, ::grpc::CompletionQueue* cq) override;
    const ::grpc::internal::RpcMethod rpcmethod_HandleCreateRemove_;
    const ::grpc::internal::RpcMethod rpcmethod_HandleMirrorRequest_;
  };
  static std::unique_ptr<Stub> NewStub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options = ::grpc::StubOptions());

  class Service : public ::grpc::Service {
   public:
    Service();
    virtual ~Service();
    // Handle the control plane request for creating or removing a flow.
    virtual ::grpc::Status HandleCreateRemove(::grpc::ServerContext* context, const ::greptime::v1::flow::FlowRequest* request, ::greptime::v1::flow::FlowResponse* response);
    // Handle the data plane request for inserting or deleting rows
    // only expect `RegionRequest` to be one of `InsertRequests` or
    // `DeleteRequests` other types of `RegionRequest` will be ignored
    virtual ::grpc::Status HandleMirrorRequest(::grpc::ServerContext* context, const ::greptime::v1::flow::InsertRequests* request, ::greptime::v1::flow::FlowResponse* response);
  };
  template <class BaseClass>
  class WithAsyncMethod_HandleCreateRemove : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithAsyncMethod_HandleCreateRemove() {
      ::grpc::Service::MarkMethodAsync(0);
    }
    ~WithAsyncMethod_HandleCreateRemove() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status HandleCreateRemove(::grpc::ServerContext* /*context*/, const ::greptime::v1::flow::FlowRequest* /*request*/, ::greptime::v1::flow::FlowResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    void RequestHandleCreateRemove(::grpc::ServerContext* context, ::greptime::v1::flow::FlowRequest* request, ::grpc::ServerAsyncResponseWriter< ::greptime::v1::flow::FlowResponse>* response, ::grpc::CompletionQueue* new_call_cq, ::grpc::ServerCompletionQueue* notification_cq, void *tag) {
      ::grpc::Service::RequestAsyncUnary(0, context, request, response, new_call_cq, notification_cq, tag);
    }
  };
  template <class BaseClass>
  class WithAsyncMethod_HandleMirrorRequest : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithAsyncMethod_HandleMirrorRequest() {
      ::grpc::Service::MarkMethodAsync(1);
    }
    ~WithAsyncMethod_HandleMirrorRequest() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status HandleMirrorRequest(::grpc::ServerContext* /*context*/, const ::greptime::v1::flow::InsertRequests* /*request*/, ::greptime::v1::flow::FlowResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    void RequestHandleMirrorRequest(::grpc::ServerContext* context, ::greptime::v1::flow::InsertRequests* request, ::grpc::ServerAsyncResponseWriter< ::greptime::v1::flow::FlowResponse>* response, ::grpc::CompletionQueue* new_call_cq, ::grpc::ServerCompletionQueue* notification_cq, void *tag) {
      ::grpc::Service::RequestAsyncUnary(1, context, request, response, new_call_cq, notification_cq, tag);
    }
  };
  typedef WithAsyncMethod_HandleCreateRemove<WithAsyncMethod_HandleMirrorRequest<Service > > AsyncService;
  template <class BaseClass>
  class WithCallbackMethod_HandleCreateRemove : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithCallbackMethod_HandleCreateRemove() {
      ::grpc::Service::MarkMethodCallback(0,
          new ::grpc::internal::CallbackUnaryHandler< ::greptime::v1::flow::FlowRequest, ::greptime::v1::flow::FlowResponse>(
            [this](
                   ::grpc::CallbackServerContext* context, const ::greptime::v1::flow::FlowRequest* request, ::greptime::v1::flow::FlowResponse* response) { return this->HandleCreateRemove(context, request, response); }));}
    void SetMessageAllocatorFor_HandleCreateRemove(
        ::grpc::MessageAllocator< ::greptime::v1::flow::FlowRequest, ::greptime::v1::flow::FlowResponse>* allocator) {
      ::grpc::internal::MethodHandler* const handler = ::grpc::Service::GetHandler(0);
      static_cast<::grpc::internal::CallbackUnaryHandler< ::greptime::v1::flow::FlowRequest, ::greptime::v1::flow::FlowResponse>*>(handler)
              ->SetMessageAllocator(allocator);
    }
    ~WithCallbackMethod_HandleCreateRemove() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status HandleCreateRemove(::grpc::ServerContext* /*context*/, const ::greptime::v1::flow::FlowRequest* /*request*/, ::greptime::v1::flow::FlowResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    virtual ::grpc::ServerUnaryReactor* HandleCreateRemove(
      ::grpc::CallbackServerContext* /*context*/, const ::greptime::v1::flow::FlowRequest* /*request*/, ::greptime::v1::flow::FlowResponse* /*response*/)  { return nullptr; }
  };
  template <class BaseClass>
  class WithCallbackMethod_HandleMirrorRequest : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithCallbackMethod_HandleMirrorRequest() {
      ::grpc::Service::MarkMethodCallback(1,
          new ::grpc::internal::CallbackUnaryHandler< ::greptime::v1::flow::InsertRequests, ::greptime::v1::flow::FlowResponse>(
            [this](
                   ::grpc::CallbackServerContext* context, const ::greptime::v1::flow::InsertRequests* request, ::greptime::v1::flow::FlowResponse* response) { return this->HandleMirrorRequest(context, request, response); }));}
    void SetMessageAllocatorFor_HandleMirrorRequest(
        ::grpc::MessageAllocator< ::greptime::v1::flow::InsertRequests, ::greptime::v1::flow::FlowResponse>* allocator) {
      ::grpc::internal::MethodHandler* const handler = ::grpc::Service::GetHandler(1);
      static_cast<::grpc::internal::CallbackUnaryHandler< ::greptime::v1::flow::InsertRequests, ::greptime::v1::flow::FlowResponse>*>(handler)
              ->SetMessageAllocator(allocator);
    }
    ~WithCallbackMethod_HandleMirrorRequest() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status HandleMirrorRequest(::grpc::ServerContext* /*context*/, const ::greptime::v1::flow::InsertRequests* /*request*/, ::greptime::v1::flow::FlowResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    virtual ::grpc::ServerUnaryReactor* HandleMirrorRequest(
      ::grpc::CallbackServerContext* /*context*/, const ::greptime::v1::flow::InsertRequests* /*request*/, ::greptime::v1::flow::FlowResponse* /*response*/)  { return nullptr; }
  };
  typedef WithCallbackMethod_HandleCreateRemove<WithCallbackMethod_HandleMirrorRequest<Service > > CallbackService;
  typedef CallbackService ExperimentalCallbackService;
  template <class BaseClass>
  class WithGenericMethod_HandleCreateRemove : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithGenericMethod_HandleCreateRemove() {
      ::grpc::Service::MarkMethodGeneric(0);
    }
    ~WithGenericMethod_HandleCreateRemove() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status HandleCreateRemove(::grpc::ServerContext* /*context*/, const ::greptime::v1::flow::FlowRequest* /*request*/, ::greptime::v1::flow::FlowResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
  };
  template <class BaseClass>
  class WithGenericMethod_HandleMirrorRequest : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithGenericMethod_HandleMirrorRequest() {
      ::grpc::Service::MarkMethodGeneric(1);
    }
    ~WithGenericMethod_HandleMirrorRequest() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status HandleMirrorRequest(::grpc::ServerContext* /*context*/, const ::greptime::v1::flow::InsertRequests* /*request*/, ::greptime::v1::flow::FlowResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
  };
  template <class BaseClass>
  class WithRawMethod_HandleCreateRemove : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithRawMethod_HandleCreateRemove() {
      ::grpc::Service::MarkMethodRaw(0);
    }
    ~WithRawMethod_HandleCreateRemove() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status HandleCreateRemove(::grpc::ServerContext* /*context*/, const ::greptime::v1::flow::FlowRequest* /*request*/, ::greptime::v1::flow::FlowResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    void RequestHandleCreateRemove(::grpc::ServerContext* context, ::grpc::ByteBuffer* request, ::grpc::ServerAsyncResponseWriter< ::grpc::ByteBuffer>* response, ::grpc::CompletionQueue* new_call_cq, ::grpc::ServerCompletionQueue* notification_cq, void *tag) {
      ::grpc::Service::RequestAsyncUnary(0, context, request, response, new_call_cq, notification_cq, tag);
    }
  };
  template <class BaseClass>
  class WithRawMethod_HandleMirrorRequest : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithRawMethod_HandleMirrorRequest() {
      ::grpc::Service::MarkMethodRaw(1);
    }
    ~WithRawMethod_HandleMirrorRequest() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status HandleMirrorRequest(::grpc::ServerContext* /*context*/, const ::greptime::v1::flow::InsertRequests* /*request*/, ::greptime::v1::flow::FlowResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    void RequestHandleMirrorRequest(::grpc::ServerContext* context, ::grpc::ByteBuffer* request, ::grpc::ServerAsyncResponseWriter< ::grpc::ByteBuffer>* response, ::grpc::CompletionQueue* new_call_cq, ::grpc::ServerCompletionQueue* notification_cq, void *tag) {
      ::grpc::Service::RequestAsyncUnary(1, context, request, response, new_call_cq, notification_cq, tag);
    }
  };
  template <class BaseClass>
  class WithRawCallbackMethod_HandleCreateRemove : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithRawCallbackMethod_HandleCreateRemove() {
      ::grpc::Service::MarkMethodRawCallback(0,
          new ::grpc::internal::CallbackUnaryHandler< ::grpc::ByteBuffer, ::grpc::ByteBuffer>(
            [this](
                   ::grpc::CallbackServerContext* context, const ::grpc::ByteBuffer* request, ::grpc::ByteBuffer* response) { return this->HandleCreateRemove(context, request, response); }));
    }
    ~WithRawCallbackMethod_HandleCreateRemove() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status HandleCreateRemove(::grpc::ServerContext* /*context*/, const ::greptime::v1::flow::FlowRequest* /*request*/, ::greptime::v1::flow::FlowResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    virtual ::grpc::ServerUnaryReactor* HandleCreateRemove(
      ::grpc::CallbackServerContext* /*context*/, const ::grpc::ByteBuffer* /*request*/, ::grpc::ByteBuffer* /*response*/)  { return nullptr; }
  };
  template <class BaseClass>
  class WithRawCallbackMethod_HandleMirrorRequest : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithRawCallbackMethod_HandleMirrorRequest() {
      ::grpc::Service::MarkMethodRawCallback(1,
          new ::grpc::internal::CallbackUnaryHandler< ::grpc::ByteBuffer, ::grpc::ByteBuffer>(
            [this](
                   ::grpc::CallbackServerContext* context, const ::grpc::ByteBuffer* request, ::grpc::ByteBuffer* response) { return this->HandleMirrorRequest(context, request, response); }));
    }
    ~WithRawCallbackMethod_HandleMirrorRequest() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status HandleMirrorRequest(::grpc::ServerContext* /*context*/, const ::greptime::v1::flow::InsertRequests* /*request*/, ::greptime::v1::flow::FlowResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    virtual ::grpc::ServerUnaryReactor* HandleMirrorRequest(
      ::grpc::CallbackServerContext* /*context*/, const ::grpc::ByteBuffer* /*request*/, ::grpc::ByteBuffer* /*response*/)  { return nullptr; }
  };
  template <class BaseClass>
  class WithStreamedUnaryMethod_HandleCreateRemove : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithStreamedUnaryMethod_HandleCreateRemove() {
      ::grpc::Service::MarkMethodStreamed(0,
        new ::grpc::internal::StreamedUnaryHandler<
          ::greptime::v1::flow::FlowRequest, ::greptime::v1::flow::FlowResponse>(
            [this](::grpc::ServerContext* context,
                   ::grpc::ServerUnaryStreamer<
                     ::greptime::v1::flow::FlowRequest, ::greptime::v1::flow::FlowResponse>* streamer) {
                       return this->StreamedHandleCreateRemove(context,
                         streamer);
                  }));
    }
    ~WithStreamedUnaryMethod_HandleCreateRemove() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable regular version of this method
    ::grpc::Status HandleCreateRemove(::grpc::ServerContext* /*context*/, const ::greptime::v1::flow::FlowRequest* /*request*/, ::greptime::v1::flow::FlowResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    // replace default version of method with streamed unary
    virtual ::grpc::Status StreamedHandleCreateRemove(::grpc::ServerContext* context, ::grpc::ServerUnaryStreamer< ::greptime::v1::flow::FlowRequest,::greptime::v1::flow::FlowResponse>* server_unary_streamer) = 0;
  };
  template <class BaseClass>
  class WithStreamedUnaryMethod_HandleMirrorRequest : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithStreamedUnaryMethod_HandleMirrorRequest() {
      ::grpc::Service::MarkMethodStreamed(1,
        new ::grpc::internal::StreamedUnaryHandler<
          ::greptime::v1::flow::InsertRequests, ::greptime::v1::flow::FlowResponse>(
            [this](::grpc::ServerContext* context,
                   ::grpc::ServerUnaryStreamer<
                     ::greptime::v1::flow::InsertRequests, ::greptime::v1::flow::FlowResponse>* streamer) {
                       return this->StreamedHandleMirrorRequest(context,
                         streamer);
                  }));
    }
    ~WithStreamedUnaryMethod_HandleMirrorRequest() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable regular version of this method
    ::grpc::Status HandleMirrorRequest(::grpc::ServerContext* /*context*/, const ::greptime::v1::flow::InsertRequests* /*request*/, ::greptime::v1::flow::FlowResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    // replace default version of method with streamed unary
    virtual ::grpc::Status StreamedHandleMirrorRequest(::grpc::ServerContext* context, ::grpc::ServerUnaryStreamer< ::greptime::v1::flow::InsertRequests,::greptime::v1::flow::FlowResponse>* server_unary_streamer) = 0;
  };
  typedef WithStreamedUnaryMethod_HandleCreateRemove<WithStreamedUnaryMethod_HandleMirrorRequest<Service > > StreamedUnaryService;
  typedef Service SplitStreamedService;
  typedef WithStreamedUnaryMethod_HandleCreateRemove<WithStreamedUnaryMethod_HandleMirrorRequest<Service > > StreamedService;
};

}  // namespace flow
}  // namespace v1
}  // namespace greptime


#endif  // GRPC_greptime_2fv1_2fflow_2fserver_2eproto__INCLUDED
