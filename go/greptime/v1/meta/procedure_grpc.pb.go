// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: greptime/v1/meta/procedure.proto

package meta

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ProcedureServiceClient is the client API for ProcedureService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProcedureServiceClient interface {
	// Query a submitted procedure state
	Query(ctx context.Context, in *QueryProcedureRequest, opts ...grpc.CallOption) (*ProcedureStateResponse, error)
	// Submits a DDL task
	Ddl(ctx context.Context, in *DdlTaskRequest, opts ...grpc.CallOption) (*DdlTaskResponse, error)
	// Submits a region migration task
	Migrate(ctx context.Context, in *MigrateRegionRequest, opts ...grpc.CallOption) (*MigrateRegionResponse, error)
	// Query all submitted procedures details
	Details(ctx context.Context, in *ProcedureDetailRequest, opts ...grpc.CallOption) (*ProcedureDetailResponse, error)
}

type procedureServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProcedureServiceClient(cc grpc.ClientConnInterface) ProcedureServiceClient {
	return &procedureServiceClient{cc}
}

func (c *procedureServiceClient) Query(ctx context.Context, in *QueryProcedureRequest, opts ...grpc.CallOption) (*ProcedureStateResponse, error) {
	out := new(ProcedureStateResponse)
	err := c.cc.Invoke(ctx, "/greptime.v1.meta.ProcedureService/query", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *procedureServiceClient) Ddl(ctx context.Context, in *DdlTaskRequest, opts ...grpc.CallOption) (*DdlTaskResponse, error) {
	out := new(DdlTaskResponse)
	err := c.cc.Invoke(ctx, "/greptime.v1.meta.ProcedureService/ddl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *procedureServiceClient) Migrate(ctx context.Context, in *MigrateRegionRequest, opts ...grpc.CallOption) (*MigrateRegionResponse, error) {
	out := new(MigrateRegionResponse)
	err := c.cc.Invoke(ctx, "/greptime.v1.meta.ProcedureService/migrate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *procedureServiceClient) Details(ctx context.Context, in *ProcedureDetailRequest, opts ...grpc.CallOption) (*ProcedureDetailResponse, error) {
	out := new(ProcedureDetailResponse)
	err := c.cc.Invoke(ctx, "/greptime.v1.meta.ProcedureService/details", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProcedureServiceServer is the server API for ProcedureService service.
// All implementations must embed UnimplementedProcedureServiceServer
// for forward compatibility
type ProcedureServiceServer interface {
	// Query a submitted procedure state
	Query(context.Context, *QueryProcedureRequest) (*ProcedureStateResponse, error)
	// Submits a DDL task
	Ddl(context.Context, *DdlTaskRequest) (*DdlTaskResponse, error)
	// Submits a region migration task
	Migrate(context.Context, *MigrateRegionRequest) (*MigrateRegionResponse, error)
	// Query all submitted procedures details
	Details(context.Context, *ProcedureDetailRequest) (*ProcedureDetailResponse, error)
	mustEmbedUnimplementedProcedureServiceServer()
}

// UnimplementedProcedureServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProcedureServiceServer struct {
}

func (UnimplementedProcedureServiceServer) Query(context.Context, *QueryProcedureRequest) (*ProcedureStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query not implemented")
}
func (UnimplementedProcedureServiceServer) Ddl(context.Context, *DdlTaskRequest) (*DdlTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ddl not implemented")
}
func (UnimplementedProcedureServiceServer) Migrate(context.Context, *MigrateRegionRequest) (*MigrateRegionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Migrate not implemented")
}
func (UnimplementedProcedureServiceServer) Details(context.Context, *ProcedureDetailRequest) (*ProcedureDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Details not implemented")
}
func (UnimplementedProcedureServiceServer) mustEmbedUnimplementedProcedureServiceServer() {}

// UnsafeProcedureServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProcedureServiceServer will
// result in compilation errors.
type UnsafeProcedureServiceServer interface {
	mustEmbedUnimplementedProcedureServiceServer()
}

func RegisterProcedureServiceServer(s grpc.ServiceRegistrar, srv ProcedureServiceServer) {
	s.RegisterService(&ProcedureService_ServiceDesc, srv)
}

func _ProcedureService_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryProcedureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProcedureServiceServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greptime.v1.meta.ProcedureService/query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProcedureServiceServer).Query(ctx, req.(*QueryProcedureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProcedureService_Ddl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DdlTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProcedureServiceServer).Ddl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greptime.v1.meta.ProcedureService/ddl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProcedureServiceServer).Ddl(ctx, req.(*DdlTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProcedureService_Migrate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MigrateRegionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProcedureServiceServer).Migrate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greptime.v1.meta.ProcedureService/migrate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProcedureServiceServer).Migrate(ctx, req.(*MigrateRegionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProcedureService_Details_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProcedureDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProcedureServiceServer).Details(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greptime.v1.meta.ProcedureService/details",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProcedureServiceServer).Details(ctx, req.(*ProcedureDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProcedureService_ServiceDesc is the grpc.ServiceDesc for ProcedureService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProcedureService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "greptime.v1.meta.ProcedureService",
	HandlerType: (*ProcedureServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "query",
			Handler:    _ProcedureService_Query_Handler,
		},
		{
			MethodName: "ddl",
			Handler:    _ProcedureService_Ddl_Handler,
		},
		{
			MethodName: "migrate",
			Handler:    _ProcedureService_Migrate_Handler,
		},
		{
			MethodName: "details",
			Handler:    _ProcedureService_Details_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "greptime/v1/meta/procedure.proto",
}
