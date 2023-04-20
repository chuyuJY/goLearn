// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.15.5
// source: pb/hello.proto

package pb

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
	Greeter_LotsOfHello_FullMethodName = "/pb.Greeter/LotsOfHello"
	Greeter_LotsOfName_FullMethodName  = "/pb.Greeter/LotsOfName"
	Greeter_BidiHello_FullMethodName   = "/pb.Greeter/BidiHello"
)

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreeterClient interface {
	// 服务端单向流
	LotsOfHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (Greeter_LotsOfHelloClient, error)
	// 客户端单向流
	LotsOfName(ctx context.Context, opts ...grpc.CallOption) (Greeter_LotsOfNameClient, error)
	// 双向流
	BidiHello(ctx context.Context, opts ...grpc.CallOption) (Greeter_BidiHelloClient, error)
}

type greeterClient struct {
	cc grpc.ClientConnInterface
}

func NewGreeterClient(cc grpc.ClientConnInterface) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) LotsOfHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (Greeter_LotsOfHelloClient, error) {
	stream, err := c.cc.NewStream(ctx, &Greeter_ServiceDesc.Streams[0], Greeter_LotsOfHello_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterLotsOfHelloClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Greeter_LotsOfHelloClient interface {
	Recv() (*HelloResponse, error)
	grpc.ClientStream
}

type greeterLotsOfHelloClient struct {
	grpc.ClientStream
}

func (x *greeterLotsOfHelloClient) Recv() (*HelloResponse, error) {
	m := new(HelloResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greeterClient) LotsOfName(ctx context.Context, opts ...grpc.CallOption) (Greeter_LotsOfNameClient, error) {
	stream, err := c.cc.NewStream(ctx, &Greeter_ServiceDesc.Streams[1], Greeter_LotsOfName_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterLotsOfNameClient{stream}
	return x, nil
}

type Greeter_LotsOfNameClient interface {
	Send(*HelloRequest) error
	CloseAndRecv() (*HelloResponse, error)
	grpc.ClientStream
}

type greeterLotsOfNameClient struct {
	grpc.ClientStream
}

func (x *greeterLotsOfNameClient) Send(m *HelloRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greeterLotsOfNameClient) CloseAndRecv() (*HelloResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(HelloResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greeterClient) BidiHello(ctx context.Context, opts ...grpc.CallOption) (Greeter_BidiHelloClient, error) {
	stream, err := c.cc.NewStream(ctx, &Greeter_ServiceDesc.Streams[2], Greeter_BidiHello_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterBidiHelloClient{stream}
	return x, nil
}

type Greeter_BidiHelloClient interface {
	Send(*HelloRequest) error
	Recv() (*HelloResponse, error)
	grpc.ClientStream
}

type greeterBidiHelloClient struct {
	grpc.ClientStream
}

func (x *greeterBidiHelloClient) Send(m *HelloRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greeterBidiHelloClient) Recv() (*HelloResponse, error) {
	m := new(HelloResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreeterServer is the server API for Greeter service.
// All implementations must embed UnimplementedGreeterServer
// for forward compatibility
type GreeterServer interface {
	// 服务端单向流
	LotsOfHello(*HelloRequest, Greeter_LotsOfHelloServer) error
	// 客户端单向流
	LotsOfName(Greeter_LotsOfNameServer) error
	// 双向流
	BidiHello(Greeter_BidiHelloServer) error
	mustEmbedUnimplementedGreeterServer()
}

// UnimplementedGreeterServer must be embedded to have forward compatible implementations.
type UnimplementedGreeterServer struct {
}

func (UnimplementedGreeterServer) LotsOfHello(*HelloRequest, Greeter_LotsOfHelloServer) error {
	return status.Errorf(codes.Unimplemented, "method LotsOfHello not implemented")
}
func (UnimplementedGreeterServer) LotsOfName(Greeter_LotsOfNameServer) error {
	return status.Errorf(codes.Unimplemented, "method LotsOfName not implemented")
}
func (UnimplementedGreeterServer) BidiHello(Greeter_BidiHelloServer) error {
	return status.Errorf(codes.Unimplemented, "method BidiHello not implemented")
}
func (UnimplementedGreeterServer) mustEmbedUnimplementedGreeterServer() {}

// UnsafeGreeterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreeterServer will
// result in compilation errors.
type UnsafeGreeterServer interface {
	mustEmbedUnimplementedGreeterServer()
}

func RegisterGreeterServer(s grpc.ServiceRegistrar, srv GreeterServer) {
	s.RegisterService(&Greeter_ServiceDesc, srv)
}

func _Greeter_LotsOfHello_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(HelloRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GreeterServer).LotsOfHello(m, &greeterLotsOfHelloServer{stream})
}

type Greeter_LotsOfHelloServer interface {
	Send(*HelloResponse) error
	grpc.ServerStream
}

type greeterLotsOfHelloServer struct {
	grpc.ServerStream
}

func (x *greeterLotsOfHelloServer) Send(m *HelloResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Greeter_LotsOfName_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreeterServer).LotsOfName(&greeterLotsOfNameServer{stream})
}

type Greeter_LotsOfNameServer interface {
	SendAndClose(*HelloResponse) error
	Recv() (*HelloRequest, error)
	grpc.ServerStream
}

type greeterLotsOfNameServer struct {
	grpc.ServerStream
}

func (x *greeterLotsOfNameServer) SendAndClose(m *HelloResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greeterLotsOfNameServer) Recv() (*HelloRequest, error) {
	m := new(HelloRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Greeter_BidiHello_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreeterServer).BidiHello(&greeterBidiHelloServer{stream})
}

type Greeter_BidiHelloServer interface {
	Send(*HelloResponse) error
	Recv() (*HelloRequest, error)
	grpc.ServerStream
}

type greeterBidiHelloServer struct {
	grpc.ServerStream
}

func (x *greeterBidiHelloServer) Send(m *HelloResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greeterBidiHelloServer) Recv() (*HelloRequest, error) {
	m := new(HelloRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Greeter_ServiceDesc is the grpc.ServiceDesc for Greeter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Greeter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "LotsOfHello",
			Handler:       _Greeter_LotsOfHello_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "LotsOfName",
			Handler:       _Greeter_LotsOfName_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "BidiHello",
			Handler:       _Greeter_BidiHello_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "pb/hello.proto",
}