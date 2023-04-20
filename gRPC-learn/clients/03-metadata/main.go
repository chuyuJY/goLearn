package main

import (
	"context"
	"flag"
	"fmt"
	"goLearn/gRPC-learn/authorize/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"io"
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

// unaryGetWithMetadata 普通RPC调用客户端metadata操作
func unaryGetWithMetadata(c pb.GreeterClient, name string) {
	fmt.Println("--- UnarySayHello client---")
	// 创建metadata, 携带验证信息
	md := metadata.Pairs(
		"token", "app-test-chuyu",
		"request_id", "1234567",
	)
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	// RPC调用 ctx会携带客户端的metadata
	var header, trailer metadata.MD
	reply, err := c.UnarySayHello(
		ctx,
		&pb.HelloRequest{Name: name},
		grpc.Header(&header),   // 接收服务端的Header
		grpc.Trailer(&trailer), // 接收服务端的Trailer
	)
	if err != nil {
		log.Println("failed to call SayHello:", err)
		return
	}

	// 1. 从header中获取location
	//location := header.Get("location")
	if location, ok := header["location"]; ok {
		fmt.Println("location from header:", location)
	} else {
		log.Println("location expected but doesn't exist in header")
		return
	}

	// 2. 获取响应结果
	fmt.Printf("got response: %s\n", reply.Reply)

	// 3. 从trailer中获取metadata
	if t, ok := trailer["timestamp"]; ok {
		fmt.Println("timestamp from metadata:", t)
	} else {
		log.Println("timestamp expected but doesn't exist in trailer")
		return
	}
}

// bidirectionalWithMetadata 流式RPC调用客户端metadata
func bidirectionalWithMetadata(c pb.GreeterClient, names []string) {
	fmt.Println("--- BidirectionalWithMetadata client---")
	md := metadata.Pairs(
		"token", "app-test-tiam",
		"request_id", "1234567",
	)
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	var header, trailer metadata.MD
	// 1. 使用带有metadata的context执行RPC调用.
	// 此处是调用远程 RPC，获取 Stream
	stream, err := c.BidirectionalStreaming(ctx)
	if err != nil {
		log.Fatalln("failed to call BidirectionalStreaming: %v\n", err)
	}

	// 发送
	go func() {
		// 2. 读取header
		header, err = stream.Header()
		if err != nil {
			log.Fatalln("failed to get header from stream: %v", err)
		}
		if location, ok := header["location"]; ok {
			fmt.Println("location from header:", location)
		} else {
			log.Println("location expected but doesn't exist in header")
			return
		}

		// 3. 发送请求数据到server
		for _, name := range names {
			time.Sleep(time.Second)
			if err := stream.Send(&pb.HelloRequest{Name: name}); err != nil {
				log.Fatalf("failed to send streaming: %v\n", err)
			}
		}
		stream.CloseSend()
	}()

	// 4. 接收所有的响应
	var rpcStatus error
	for {
		reply, err := stream.Recv()
		if err != nil {
			rpcStatus = err
			break
		}
		fmt.Println("got response:", reply.Reply)
	}
	if rpcStatus != io.EOF {
		log.Printf("failed to finish server streaming: %v\n", rpcStatus)
		return
	}

	// 5. RPC调用结束，获取trailer
	trailer = stream.Trailer()
	if t, ok := trailer["timestamp"]; ok {
		fmt.Printf("timestamp from trailer: %s\n", t)
	} else {
		log.Printf("timestamp expected but doesn't exist in trailer\n")
	}
}

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("did not connect:", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)
	//unaryGetWithMetadata(c, defaultName)
	bidirectionalWithMetadata(c, []string{"world", "tiam", "chuyu"})
}
