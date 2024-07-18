// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: api/whoami.proto

package whoami

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

// WhoAmIClient is the client API for WhoAmI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WhoAmIClient interface {
	GetWhoAmI(ctx context.Context, in *WhoAmIRequest, opts ...grpc.CallOption) (*WhoAmIResponse, error)
}

type whoAmIClient struct {
	cc grpc.ClientConnInterface
}

func NewWhoAmIClient(cc grpc.ClientConnInterface) WhoAmIClient {
	return &whoAmIClient{cc}
}

func (c *whoAmIClient) GetWhoAmI(ctx context.Context, in *WhoAmIRequest, opts ...grpc.CallOption) (*WhoAmIResponse, error) {
	out := new(WhoAmIResponse)
	err := c.cc.Invoke(ctx, "/docker.whoami.service.WhoAmI/GetWhoAmI", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WhoAmIServer is the server API for WhoAmI service.
// All implementations must embed UnimplementedWhoAmIServer
// for forward compatibility
type WhoAmIServer interface {
	GetWhoAmI(context.Context, *WhoAmIRequest) (*WhoAmIResponse, error)
	mustEmbedUnimplementedWhoAmIServer()
}

// UnimplementedWhoAmIServer must be embedded to have forward compatible implementations.
type UnimplementedWhoAmIServer struct {
}

func (UnimplementedWhoAmIServer) GetWhoAmI(context.Context, *WhoAmIRequest) (*WhoAmIResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWhoAmI not implemented")
}
func (UnimplementedWhoAmIServer) mustEmbedUnimplementedWhoAmIServer() {}

// UnsafeWhoAmIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WhoAmIServer will
// result in compilation errors.
type UnsafeWhoAmIServer interface {
	mustEmbedUnimplementedWhoAmIServer()
}

func RegisterWhoAmIServer(s grpc.ServiceRegistrar, srv WhoAmIServer) {
	s.RegisterService(&WhoAmI_ServiceDesc, srv)
}

func _WhoAmI_GetWhoAmI_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WhoAmIRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WhoAmIServer).GetWhoAmI(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/docker.whoami.service.WhoAmI/GetWhoAmI",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WhoAmIServer).GetWhoAmI(ctx, req.(*WhoAmIRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WhoAmI_ServiceDesc is the grpc.ServiceDesc for WhoAmI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WhoAmI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "docker.whoami.service.WhoAmI",
	HandlerType: (*WhoAmIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetWhoAmI",
			Handler:    _WhoAmI_GetWhoAmI_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/whoami.proto",
}