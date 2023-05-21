package main

import (
	"context"
	"fmt"
	"goLearn/gRPC-learn/servers/07-context/pb"
	"google.golang.org/grpc"
	"net"
	"strings"
	"time"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) UnaryEcho(ctx context.Context, in *pb.EchoRequest) (*pb.EchoResponse, error) {
	message := in.Message
	if strings.HasPrefix(message, "[propagate me]") {
		time.Sleep(800 * time.Millisecond)
		message = strings.TrimPrefix(message, "[propagate me]")
		return &pb.EchoResponse{Message: message}, nil
	}
	if message == "delay" {
		time.Sleep(2 * time.Second)
	}
	return &pb.EchoResponse{Message: message}, nil
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
