package main

import (
	"context"
	"flag"
	"fmt"
	"goLearn/gRPC-learn/clients/07-context/pb"
	"google.golang.org/genproto/googleapis/rpc/code"
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

func getUnaryEcho(c pb.GreeterClient, requestID int, message string, want code.Code) {
	// 指定 1s 超时
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.EchoRequest{Message: message}

	_, err := c.UnaryEcho(ctx, req)
	got := status.Code(err)
	fmt.Printf("[%v] wanted = %v, got = %v\n", requestID, want, got)
}

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("did not connect:", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)
	// 执行rpc调用
	getUnaryEcho(c, 1, "hello world", code.Code_OK)
	getUnaryEcho(c, 2, "delay", code.Code_DEADLINE_EXCEEDED)
	getUnaryEcho(c, 3, "[propagate me]nihao", code.Code_OK)
}
