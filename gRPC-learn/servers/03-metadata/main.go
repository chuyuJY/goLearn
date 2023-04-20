package main

import (
	"context"
	"fmt"
	"goLearn/gRPC-learn/servers/03-metadata/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"io"
	"net"
	"strconv"
	"time"
)

type server struct {
	pb.UnimplementedGreeterServer
}

// UnarySayHello 普通RPC调用服务端metadata操作
func (s *server) UnarySayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	// 3. 通过defer设置trailer
	defer func() {
		trailer := metadata.Pairs("timestamp", strconv.Itoa(int(time.Now().Unix())))
		grpc.SetTrailer(ctx, trailer)
	}()
	// 1. 从客户端请求上下文中读取metadata
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.DataLoss, "UnarySayHello: failed to get metadata")
	}
	if token, ok := md["token"]; ok {
		fmt.Println("token from metadata:", token)
		if len(token) < 1 || token[0] != "app-test-tiam" {
			return nil, status.Error(codes.Unauthenticated, "UnarySayHello: authentication failed")
		}
	} else {
		return nil, status.Errorf(codes.PermissionDenied, "UnarySayHello: lack of token")
	}

	// 2. 创建和发送header
	header := metadata.New(map[string]string{"location": "ZhengZhou"})
	grpc.SendHeader(ctx, header)

	fmt.Printf("Request received: %v, say hello...\n", in)

	return &pb.HelloResponse{Reply: "hello, " + in.GetName()}, nil
}

// BidirectionalStreaming 流式RPC调用客户端metadata操作
func (s *server) BidirectionalStreaming(stream pb.Greeter_BidirectionalStreamingServer) error {
	// 在defer中创建trailer记录函数的返回时间
	defer func() {
		trailer := metadata.Pairs("timestamp", strconv.Itoa(int(time.Now().Unix())))
		stream.SetTrailer(trailer)
	}()
	// 1. 从client读取metadata
	md, ok := metadata.FromIncomingContext(stream.Context())
	if !ok {
		return status.Errorf(codes.DataLoss, "BidirectionalStreamingSayHello: failed to get metadata")
	}
	if token, ok := md["token"]; ok {
		fmt.Println("token from metadata:", token)
	}

	// 2. 创建和发送header
	header := metadata.New(map[string]string{"location": "Zhengzhou"})
	stream.SendHeader(header)

	// // 3. 读取请求数据并发送响应数据
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		if err := stream.Send(&pb.HelloResponse{Reply: "hello, " + in.GetName()}); err != nil {
			return err
		}
	}
}

func main() {
	listen, err := net.Listen("tcp", ":8972")
	if err != nil {
		fmt.Println(err)
		return
	}
	s := grpc.NewServer()                  // 创建grpc服务器
	pb.RegisterGreeterServer(s, &server{}) // 注册服务
	err = s.Serve(listen)                  // 启动服务
	if err != nil {
		fmt.Println("服务启动失败：", err)
		return
	}
}
