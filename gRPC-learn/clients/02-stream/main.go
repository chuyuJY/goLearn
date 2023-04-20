package main

import (
	"bufio"
	"context"
	"flag"
	"goLearn/gRPC-learn/clients/02-stream/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
	"strings"

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

func getLotsOfHello(c pb.GreeterClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	// 得到服务端的发送流
	stream, err := c.LotsOfHello(ctx, &pb.HelloRequest{Name: *name})
	if err != nil {
		log.Fatalln("c.LotsOfHello err:", err)
		return
	}
	for {
		// 接收服务端流式数据，当收到错误或io.EOF时退出
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		log.Println("get hello:", resp.GetReply())
	}
}

func sendLotsOfName(c pb.GreeterClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	// 客户端流式请求
	stream, err := c.LotsOfName(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	names := []string{"world", "chuyu", "tima"}
	for _, name := range names {
		if err := stream.Send(&pb.HelloRequest{Name: name}); err != nil {
			log.Fatalln(err)
		}
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("get namesHello: ", res.GetReply())
}

func getBidiHello(c pb.GreeterClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	// 双向流模式
	stream, err := c.BidiHello(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	// 用于等待goroutine退出
	waitc := make(chan struct{})
	// 接收服务端返回的响应
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				// 保证goroutine退出后，再退出主线程
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalln(err)
			}
			log.Println("getBidiHello:", in.GetReply())
		}
	}()
	// 发送客户端的请求，从标准输入中获取用户的请求
	reader := bufio.NewReader(os.Stdin)
	for {
		cmd, _ := reader.ReadString('\n') // 读到换行
		cmd = strings.TrimSpace(cmd)
		if len(cmd) == 0 {
			continue
		}
		if strings.ToUpper(cmd) == "EXIT" {
			break
		}
		if err := stream.Send(&pb.HelloRequest{Name: cmd}); err != nil {
			log.Fatalln(err)
		}
	}
	stream.CloseSend()
	// 若goroutine未退出，则主线程阻塞在此处
	<-waitc
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
	//getLotsOfHello(c)
	//sendLotsOfName(c)
	getBidiHello(c)
}
