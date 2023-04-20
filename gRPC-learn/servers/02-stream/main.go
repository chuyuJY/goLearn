package main

import (
	"context"
	"fmt"
	"goLearn/gRPC-learn/servers/02-stream/pb"
	"google.golang.org/grpc"
	"io"
	"net"
	"strings"
)

type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello 重写
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Reply: "Hello, " + in.Name}, nil
}

// LotsOfHello 流式响应客户端的请求
func (s *server) LotsOfHello(in *pb.HelloRequest, stream pb.Greeter_LotsOfHelloServer) error {
	words := []string{
		"你好",
		"hello",
		"こんにちは",
		"안녕하세요",
	}
	for _, word := range words {
		data := &pb.HelloResponse{
			Reply: strings.Join([]string{word, in.GetName()}, ", "),
		}
		if err := stream.Send(data); err != nil {
			return err
		}
	}
	return nil
}

// LotsOfName 一次性回复客户端的流式请求
func (s *server) LotsOfName(stream pb.Greeter_LotsOfNameServer) error {
	reply := "你好呀！"
	names := []string{}
	for {
		// 接收客户端流式数据
		resp, err := stream.Recv()
		if err == io.EOF {
			// 接收完毕，一次性回复
			return stream.SendAndClose(&pb.HelloResponse{Reply: reply + strings.Join(names, ", ")})
		}
		if err != nil {
			return err
		}
		names = append(names, resp.GetName())
	}
}

// BidiHello 双向流式Hello
func (s *server) BidiHello(stream pb.Greeter_BidiHelloServer) error {
	for {
		// 接收流式请求
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		reply := magic(in.GetName()) // 处理接收到数据

		// 返回流式响应
		if err := stream.Send(&pb.HelloResponse{Reply: reply}); err != nil {
			return err
		}
	}
}

// magic 处理接收的数据：一段价值连城的人工智能代码
func magic(s string) string {
	// 一段价值连城的人工智能代码
	s = strings.ReplaceAll(s, "吗", "")
	s = strings.ReplaceAll(s, "吧", "")
	s = strings.ReplaceAll(s, "你", "我")
	s = strings.ReplaceAll(s, "？", "!")
	s = strings.ReplaceAll(s, "?", "!")
	return s
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
