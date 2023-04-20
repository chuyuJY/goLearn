package main

import (
	"context"
	"errors"
	"fmt"
	pb "goLearn/gRPC-learn/servers/01-string/pb"
	"google.golang.org/grpc"
	"net"
)

type StringService struct {
	pb.UnimplementedStringServiceServer
}

func (s StringService) Concat(_ context.Context, request *pb.StringRequest) (*pb.StringResp, error) {
	if len(request.A+request.B) > 10 {
		return &pb.StringResp{Err: "too long strings"}, errors.New("too long strings")
	}
	return &pb.StringResp{Ret: request.B + request.A}, nil
}

func (s StringService) Diff(_ context.Context, request *pb.StringRequest) (*pb.StringResp, error) {
	if len(request.A+request.B) > 10 {
		return &pb.StringResp{Err: "too long strings"}, errors.New("too long strings")
	}
	if request.A == request.B {
		return &pb.StringResp{Ret: "two strings is same"}, nil
	} else {
		return &pb.StringResp{Ret: "two strings is not same"}, nil
	}
}

func main() {
	listen, err := net.Listen("tcp", ":8972")
	if err != nil {
		fmt.Println(err)
		return
	}
	s := grpc.NewServer()
	pb.RegisterStringServiceServer(s, &StringService{})
	err = s.Serve(listen)
	if err != nil {
		fmt.Println("服务启动失败: ", err)
	}
}
