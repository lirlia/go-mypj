// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package myproto

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

// CloudClient is the client API for Cloud service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CloudClient interface {
	GetDate(ctx context.Context, in *DateRequest, opts ...grpc.CallOption) (*DateReply, error)
}

type cloudClient struct {
	cc grpc.ClientConnInterface
}

func NewCloudClient(cc grpc.ClientConnInterface) CloudClient {
	return &cloudClient{cc}
}

func (c *cloudClient) GetDate(ctx context.Context, in *DateRequest, opts ...grpc.CallOption) (*DateReply, error) {
	out := new(DateReply)
	err := c.cc.Invoke(ctx, "/myproto.Cloud/GetDate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CloudServer is the server API for Cloud service.
// All implementations must embed UnimplementedCloudServer
// for forward compatibility
type CloudServer interface {
	GetDate(context.Context, *DateRequest) (*DateReply, error)
	mustEmbedUnimplementedCloudServer()
}

// UnimplementedCloudServer must be embedded to have forward compatible implementations.
type UnimplementedCloudServer struct {
}

func (UnimplementedCloudServer) GetDate(context.Context, *DateRequest) (*DateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDate not implemented")
}
func (UnimplementedCloudServer) mustEmbedUnimplementedCloudServer() {}

// UnsafeCloudServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CloudServer will
// result in compilation errors.
type UnsafeCloudServer interface {
	mustEmbedUnimplementedCloudServer()
}

func RegisterCloudServer(s grpc.ServiceRegistrar, srv CloudServer) {
	s.RegisterService(&Cloud_ServiceDesc, srv)
}

func _Cloud_GetDate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudServer).GetDate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/myproto.Cloud/GetDate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudServer).GetDate(ctx, req.(*DateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Cloud_ServiceDesc is the grpc.ServiceDesc for Cloud service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Cloud_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "myproto.Cloud",
	HandlerType: (*CloudServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDate",
			Handler:    _Cloud_GetDate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hello.proto",
}
