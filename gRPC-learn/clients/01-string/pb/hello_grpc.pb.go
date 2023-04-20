// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.15.5
// source: pb/hello.proto

package client_pb

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

const (
	StringService_Concat_FullMethodName = "/pb.StringService/Concat"
	StringService_Diff_FullMethodName   = "/pb.StringService/Diff"
)

// StringServiceClient is the client API for StringService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StringServiceClient interface {
	Concat(ctx context.Context, in *StringRequest, opts ...grpc.CallOption) (*StringResp, error)
	Diff(ctx context.Context, in *StringRequest, opts ...grpc.CallOption) (*StringResp, error)
}

type stringServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStringServiceClient(cc grpc.ClientConnInterface) StringServiceClient {
	return &stringServiceClient{cc}
}

func (c *stringServiceClient) Concat(ctx context.Context, in *StringRequest, opts ...grpc.CallOption) (*StringResp, error) {
	out := new(StringResp)
	err := c.cc.Invoke(ctx, StringService_Concat_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stringServiceClient) Diff(ctx context.Context, in *StringRequest, opts ...grpc.CallOption) (*StringResp, error) {
	out := new(StringResp)
	err := c.cc.Invoke(ctx, StringService_Diff_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StringServiceServer is the server API for StringService service.
// All implementations must embed UnimplementedStringServiceServer
// for forward compatibility
type StringServiceServer interface {
	Concat(context.Context, *StringRequest) (*StringResp, error)
	Diff(context.Context, *StringRequest) (*StringResp, error)
	mustEmbedUnimplementedStringServiceServer()
}

// UnimplementedStringServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStringServiceServer struct {
}

func (UnimplementedStringServiceServer) Concat(context.Context, *StringRequest) (*StringResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Concat not implemented")
}
func (UnimplementedStringServiceServer) Diff(context.Context, *StringRequest) (*StringResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Diff not implemented")
}
func (UnimplementedStringServiceServer) mustEmbedUnimplementedStringServiceServer() {}

// UnsafeStringServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StringServiceServer will
// result in compilation errors.
type UnsafeStringServiceServer interface {
	mustEmbedUnimplementedStringServiceServer()
}

func RegisterStringServiceServer(s grpc.ServiceRegistrar, srv StringServiceServer) {
	s.RegisterService(&StringService_ServiceDesc, srv)
}

func _StringService_Concat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StringServiceServer).Concat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StringService_Concat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StringServiceServer).Concat(ctx, req.(*StringRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StringService_Diff_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StringServiceServer).Diff(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StringService_Diff_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StringServiceServer).Diff(ctx, req.(*StringRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StringService_ServiceDesc is the grpc.ServiceDesc for StringService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StringService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.StringService",
	HandlerType: (*StringServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Concat",
			Handler:    _StringService_Concat_Handler,
		},
		{
			MethodName: "Diff",
			Handler:    _StringService_Diff_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/hello.proto",
}