package main

import (
	"context"
	"flag"
	"fmt"
	"goLearn/gRPC-learn/clients/04-error/pb"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"log"
	"time"
)

const (
	defaultName = "world"
)

var (
	addr    = flag.String("addr", "127.0.0.1:8972", "the address to connect to")
	name    = flag.String("name", defaultName, "who to greet")
	restful = flag.String("restful", ":8080", "the port to restful serve on")
)

func getSayHello(c pb.GreeterClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	reply, err := c.SayHello(ctx, &pb.HelloRequest{Name: defaultName})
	// 处理接收到的错误
	if err != nil {
		fmt.Printf("get c.SayHello failed, err: %v\n", err)
		// 将err转为status
		s := status.Convert(err)
		// 获取details
		for _, detail := range s.Details() {
			switch info := detail.(type) {
			case *errdetails.QuotaFailure:
				fmt.Printf("Quota failure: %s\n", info)
			default:
				fmt.Printf("Unexpected type: %s\n", info)
			}
		}
		return
	}
	// 获取到了RPC响应
	log.Println("Greeting:", reply.GetReply())
}

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("did not connect:", err)
	}
	defer conn.Close()
	// 执行rpc调用
	c := pb.NewGreeterClient(conn)
	getSayHello(c)
}
