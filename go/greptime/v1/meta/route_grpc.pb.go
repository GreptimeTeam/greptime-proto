// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: greptime/v1/meta/route.proto

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

// RouterClient is the client API for Router service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RouterClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*RouteResponse, error)
	// Fetch routing information for tables. The smallest unit is the complete
	// routing information(all regions) of a table.
	//
	// ```text
	// table_1
	//    table_name
	//    table_schema
	//    regions
	//      region_1
	//        leader_peer
	//        follower_peer_1, follower_peer_2
	//      region_2
	//        leader_peer
	//        follower_peer_1, follower_peer_2, follower_peer_3
	//      region_xxx
	// table_2
	//    ...
	// ```
	//
	Route(ctx context.Context, in *RouteRequest, opts ...grpc.CallOption) (*RouteResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*RouteResponse, error)
}

type routerClient struct {
	cc grpc.ClientConnInterface
}

func NewRouterClient(cc grpc.ClientConnInterface) RouterClient {
	return &routerClient{cc}
}

func (c *routerClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*RouteResponse, error) {
	out := new(RouteResponse)
	err := c.cc.Invoke(ctx, "/greptime.v1.meta.Router/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routerClient) Route(ctx context.Context, in *RouteRequest, opts ...grpc.CallOption) (*RouteResponse, error) {
	out := new(RouteResponse)
	err := c.cc.Invoke(ctx, "/greptime.v1.meta.Router/Route", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routerClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*RouteResponse, error) {
	out := new(RouteResponse)
	err := c.cc.Invoke(ctx, "/greptime.v1.meta.Router/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RouterServer is the server API for Router service.
// All implementations must embed UnimplementedRouterServer
// for forward compatibility
type RouterServer interface {
	Create(context.Context, *CreateRequest) (*RouteResponse, error)
	// Fetch routing information for tables. The smallest unit is the complete
	// routing information(all regions) of a table.
	//
	// ```text
	// table_1
	//    table_name
	//    table_schema
	//    regions
	//      region_1
	//        leader_peer
	//        follower_peer_1, follower_peer_2
	//      region_2
	//        leader_peer
	//        follower_peer_1, follower_peer_2, follower_peer_3
	//      region_xxx
	// table_2
	//    ...
	// ```
	//
	Route(context.Context, *RouteRequest) (*RouteResponse, error)
	Delete(context.Context, *DeleteRequest) (*RouteResponse, error)
	mustEmbedUnimplementedRouterServer()
}

// UnimplementedRouterServer must be embedded to have forward compatible implementations.
type UnimplementedRouterServer struct {
}

func (UnimplementedRouterServer) Create(context.Context, *CreateRequest) (*RouteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedRouterServer) Route(context.Context, *RouteRequest) (*RouteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Route not implemented")
}
func (UnimplementedRouterServer) Delete(context.Context, *DeleteRequest) (*RouteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedRouterServer) mustEmbedUnimplementedRouterServer() {}

// UnsafeRouterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RouterServer will
// result in compilation errors.
type UnsafeRouterServer interface {
	mustEmbedUnimplementedRouterServer()
}

func RegisterRouterServer(s grpc.ServiceRegistrar, srv RouterServer) {
	s.RegisterService(&Router_ServiceDesc, srv)
}

func _Router_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouterServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greptime.v1.meta.Router/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouterServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Router_Route_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RouteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouterServer).Route(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greptime.v1.meta.Router/Route",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouterServer).Route(ctx, req.(*RouteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Router_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouterServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greptime.v1.meta.Router/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouterServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Router_ServiceDesc is the grpc.ServiceDesc for Router service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Router_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "greptime.v1.meta.Router",
	HandlerType: (*RouterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Router_Create_Handler,
		},
		{
			MethodName: "Route",
			Handler:    _Router_Route_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Router_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "greptime/v1/meta/route.proto",
}
