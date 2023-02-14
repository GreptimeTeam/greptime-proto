// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: greptime/v1/meta/lock.proto

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

// LockClient is the client API for Lock service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LockClient interface {
	// Lock acquires a distributed shared lock on a given named lock. On success, it
	// will return a unique key that exists so long as the lock is held by the caller.
	Lock(ctx context.Context, in *LockRequest, opts ...grpc.CallOption) (*LockResponse, error)
	// Unlock takes a key returned by Lock and releases the hold on lock.
	Unlock(ctx context.Context, in *UnlockRequest, opts ...grpc.CallOption) (*UnlockResponse, error)
}

type lockClient struct {
	cc grpc.ClientConnInterface
}

func NewLockClient(cc grpc.ClientConnInterface) LockClient {
	return &lockClient{cc}
}

func (c *lockClient) Lock(ctx context.Context, in *LockRequest, opts ...grpc.CallOption) (*LockResponse, error) {
	out := new(LockResponse)
	err := c.cc.Invoke(ctx, "/greptime.v1.meta.Lock/Lock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lockClient) Unlock(ctx context.Context, in *UnlockRequest, opts ...grpc.CallOption) (*UnlockResponse, error) {
	out := new(UnlockResponse)
	err := c.cc.Invoke(ctx, "/greptime.v1.meta.Lock/Unlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LockServer is the server API for Lock service.
// All implementations must embed UnimplementedLockServer
// for forward compatibility
type LockServer interface {
	// Lock acquires a distributed shared lock on a given named lock. On success, it
	// will return a unique key that exists so long as the lock is held by the caller.
	Lock(context.Context, *LockRequest) (*LockResponse, error)
	// Unlock takes a key returned by Lock and releases the hold on lock.
	Unlock(context.Context, *UnlockRequest) (*UnlockResponse, error)
	mustEmbedUnimplementedLockServer()
}

// UnimplementedLockServer must be embedded to have forward compatible implementations.
type UnimplementedLockServer struct {
}

func (UnimplementedLockServer) Lock(context.Context, *LockRequest) (*LockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Lock not implemented")
}
func (UnimplementedLockServer) Unlock(context.Context, *UnlockRequest) (*UnlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unlock not implemented")
}
func (UnimplementedLockServer) mustEmbedUnimplementedLockServer() {}

// UnsafeLockServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LockServer will
// result in compilation errors.
type UnsafeLockServer interface {
	mustEmbedUnimplementedLockServer()
}

func RegisterLockServer(s grpc.ServiceRegistrar, srv LockServer) {
	s.RegisterService(&Lock_ServiceDesc, srv)
}

func _Lock_Lock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LockServer).Lock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greptime.v1.meta.Lock/Lock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LockServer).Lock(ctx, req.(*LockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Lock_Unlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LockServer).Unlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greptime.v1.meta.Lock/Unlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LockServer).Unlock(ctx, req.(*UnlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Lock_ServiceDesc is the grpc.ServiceDesc for Lock service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Lock_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "greptime.v1.meta.Lock",
	HandlerType: (*LockServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Lock",
			Handler:    _Lock_Lock_Handler,
		},
		{
			MethodName: "Unlock",
			Handler:    _Lock_Unlock_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "greptime/v1/meta/lock.proto",
}
