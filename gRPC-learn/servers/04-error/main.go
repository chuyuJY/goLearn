package main

import (
	"context"
	"fmt"
	"goLearn/gRPC-learn/servers/04-error/pb"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net"
	"sync"
)

type server struct {
	pb.UnimplementedGreeterServer
	mu    sync.Mutex     // count 的并发锁
	count map[string]int // 记录每个 name 访问的次数
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.count[in.GetName()]++ // 记录用户访问次数
	// 错误处理
	if s.count[in.GetName()] > 1 {
		st := status.New(codes.ResourceExhausted, "Request limit exceeded")
		ds, err := st.WithDetails(
			&errdetails.QuotaFailure{
				Violations: []*errdetails.QuotaFailure_Violation{{
					Subject:     fmt.Sprintf("name: %s", in.GetName()),
					Description: "每个name限制访问次数为1次",
				}},
			})
		if err != nil {
			return nil, st.Err()
		}
		return nil, ds.Err()
	}
	return &pb.HelloResponse{Reply: "Hello, " + in.GetName()}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":8972")
	if err != nil {
		fmt.Println(err)
		return
	}
	s := grpc.NewServer()                                         // 创建grpc服务器
	pb.RegisterGreeterServer(s, &server{count: map[string]int{}}) // 注册服务
	err = s.Serve(listen)                                         // 启动服务
	if err != nil {
		fmt.Println("服务启动失败：", err)
		return
	}
}
