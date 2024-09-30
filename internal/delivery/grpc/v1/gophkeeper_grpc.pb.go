// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.12.4
// source: internal/delivery/grpc/v1/gophkeeper.proto

package gophkeeper

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	GophKeeperHandler_Ping_FullMethodName = "/gophkeeper.GophKeeperHandler/Ping"
)

// GophKeeperHandlerClient is the client API for GophKeeperHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GophKeeperHandlerClient interface {
	Ping(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
}

type gophKeeperHandlerClient struct {
	cc grpc.ClientConnInterface
}

func NewGophKeeperHandlerClient(cc grpc.ClientConnInterface) GophKeeperHandlerClient {
	return &gophKeeperHandlerClient{cc}
}

func (c *gophKeeperHandlerClient) Ping(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, GophKeeperHandler_Ping_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GophKeeperHandlerServer is the server API for GophKeeperHandler service.
// All implementations must embed UnimplementedGophKeeperHandlerServer
// for forward compatibility
type GophKeeperHandlerServer interface {
	Ping(context.Context, *empty.Empty) (*empty.Empty, error)
	mustEmbedUnimplementedGophKeeperHandlerServer()
}

// UnimplementedGophKeeperHandlerServer must be embedded to have forward compatible implementations.
type UnimplementedGophKeeperHandlerServer struct {
}

func (UnimplementedGophKeeperHandlerServer) Ping(context.Context, *empty.Empty) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedGophKeeperHandlerServer) mustEmbedUnimplementedGophKeeperHandlerServer() {}

// UnsafeGophKeeperHandlerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GophKeeperHandlerServer will
// result in compilation errors.
type UnsafeGophKeeperHandlerServer interface {
	mustEmbedUnimplementedGophKeeperHandlerServer()
}

func RegisterGophKeeperHandlerServer(s grpc.ServiceRegistrar, srv GophKeeperHandlerServer) {
	s.RegisterService(&GophKeeperHandler_ServiceDesc, srv)
}

func _GophKeeperHandler_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophKeeperHandlerServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GophKeeperHandler_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophKeeperHandlerServer).Ping(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// GophKeeperHandler_ServiceDesc is the grpc.ServiceDesc for GophKeeperHandler service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GophKeeperHandler_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gophkeeper.GophKeeperHandler",
	HandlerType: (*GophKeeperHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _GophKeeperHandler_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/delivery/grpc/v1/gophkeeper.proto",
}
