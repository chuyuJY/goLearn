package main

import (
	"bufio"
	"context"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	pb2 "goLearn/gRPC-learn/authorize/pb"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"golang.org/x/oauth2"
	"google.golang.org/genproto/googleapis/rpc/code"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	defaultName = "world"
)

var (
	addr    = flag.String("addr", "127.0.0.1:8972", "the address to connect to")
	name    = flag.String("name", defaultName, "who to greet")
	restful = flag.String("restful", ":8080", "the port to restful serve on")
)

func getSayHello(c pb2.GreeterClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	reply, err := c.SayHello(ctx, &pb2.HelloRequest{Name: *name})
	// 处理接收到的错误
	if err != nil {
		// 将err转为status
		s := status.Convert(err)
		// 获取details
		for _, d := range s.Details() {
			switch info := d.(type) {
			case *errdetails.QuotaFailure:
				fmt.Printf("Quota failure: %s\n", info)
			default:
				fmt.Printf("Unexpected type: %s\n", info)
			}
		}
		fmt.Printf("c.SayHello failed, err:%v\n", err)
		return
	}
	// 获取到了RPC响应
	log.Println("Greeting:", reply.GetReply())
}

func getLotsOfHello(c pb2.GreeterClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	stream, err := c.LotsOfHello(ctx, &pb2.HelloRequest{Name: *name})
	if err != nil {
		log.Fatalln("c.LotsOfHello err:", err)
		return
	}
	for {
		// 接收服务端流式数据，当收到错误或io.EOF时退出
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		log.Println("get hello:", res.GetReply())
	}
}

func getLotsOfName(c pb2.GreeterClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	// 客户端流式请求
	stream, err := c.LotsOfName(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	names := []string{"world", "chuyu", "tima"}
	for _, name := range names {
		if err := stream.Send(&pb2.HelloRequest{Name: name}); err != nil {
			log.Fatalln(err)
		}
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("get namesHello:", res.GetReply())
}

func getBidiHello(c pb2.GreeterClient) {
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
				// 保证goroutine退出后，主线程再退出
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
		if err := stream.Send(&pb2.HelloRequest{Name: cmd}); err != nil {
			log.Fatalln(err)
		}
	}
	stream.CloseSend()
	// 若goroutine未退出，则主线程阻塞在此处
	<-waitc
}

// unaryGetWithMetadata 普通RPC调用客户端metadata操作
func unaryGetWithMetadata(c pb2.GreeterClient, name string) {
	fmt.Println("--- UnarySayHello client---")
	// 创建metadata
	md := metadata.Pairs(
		"token", "app-test-tiam",
		"request_id", "1234567",
	)
	// 基于metadata创建context
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	// RPC调用 ctx会携带客户端的metadata
	var header, trailer metadata.MD
	// 数据交互发生在此处！
	reply, err := c.UnarySayHello(
		ctx,
		&pb2.HelloRequest{Name: name},
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
func bidirectionalWithMetadata(c pb2.GreeterClient, names []string) {
	fmt.Println("--- BidirectionalWithMetadata client---")
	md := metadata.Pairs(
		"token", "app-test-tiam",
		"request_id", "1234567",
	)
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	var header, trailer metadata.MD
	// 使用带有metadata的context执行RPC调用.
	stream, err := c.BidirectionalStreaming(ctx)
	if err != nil {
		log.Fatalln("failed to call BidirectionalStreaming: %v\n", err)
	}

	go func() {
		// 1. 读取header
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

		// 2. 发送请求数据到server
		for _, name := range names {
			time.Sleep(time.Second)
			if err := stream.Send(&pb2.HelloRequest{Name: name}); err != nil {
				log.Fatalf("failed to send streaming: %v\n", err)
			}
		}
		stream.CloseSend()
	}()

	// 3. 接收所有的响应
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

	// 3. RPC调用结束，获取trailer
	trailer = stream.Trailer()
	if t, ok := trailer["timestamp"]; ok {
		fmt.Printf("timestamp from trailer: %s\n", t)
	} else {
		log.Printf("timestamp expected but doesn't exist in trailer\n")
	}
}

// unaryInterceptor 客户端一元拦截器
func unaryInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	var credsConfigured bool
	for _, o := range opts {
		_, ok := o.(grpc.PerRPCCredsCallOption)
		if ok {
			credsConfigured = true
			break
		}
	}
	if !credsConfigured {
		// "golang.org/x/oauth2"
		opts = append(opts, grpc.PerRPCCredentials(oauth.NewOauthAccess(&oauth2.Token{
			AccessToken: "some-secret-token",
		})))
	}
	// pre-processing
	start := time.Now()
	err := invoker(ctx, method, req, reply, cc, opts...)
	// post-processing
	end := time.Now()
	fmt.Printf("RPC: %s, start time: %s, end time: %s, err: %v\n", method, start.Format(time.RFC3339), end.Format(time.RFC3339), err)
	return err
}

type wrappedStream struct {
	grpc.ClientStream
}

func (ws *wrappedStream) RecvMsg(m interface{}) error {
	fmt.Printf("Receive a message (Type: %T) at %v\n", m, time.Now().Format(time.RFC3339))
	return ws.ClientStream.RecvMsg(m)
}

func (ws *wrappedStream) SendMsg(m interface{}) error {
	fmt.Printf("Send a message (Type: %T) at %v\n", m, time.Now().Format(time.RFC3339))
	return ws.ClientStream.SendMsg(m)
}

func newWrappedStream(s grpc.ClientStream) grpc.ClientStream {
	return &wrappedStream{s}
}

// 客户端流式拦截器
func streamInterceptor(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	var credsConfigured bool
	for _, o := range opts {
		if _, ok := o.(*grpc.PerRPCCredsCallOption); ok {
			credsConfigured = true
			break
		}
	}
	if !credsConfigured {
		opts = append(opts, grpc.PerRPCCredentials(oauth.NewOauthAccess(&oauth2.Token{
			AccessToken: "some-secret-token",
		})))
	}
	s, err := streamer(ctx, desc, cc, method, opts...)
	if err != nil {
		return nil, err
	}
	return newWrappedStream(s), err
}

// unaryEcho 一元echo
func unaryEcho(c pb2.GreeterClient, requestID int, message string, want code.Code) {
	// 指定1s超时
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb2.EchoRequest{Message: message}

	_, err := c.UnaryEcho(ctx, req)
	got := status.Code(err)
	fmt.Printf("[%v] wanted = %v, got = %v\n", requestID, want, got)
}

func main() {
	flag.Parse()

	// 1.
	//certFile := "../ca.crt"
	//creds, _ := credentials.NewClientTLSFromFile(certFile, "www.seuzjy.cn")
	//conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(creds))
	//if err != nil {
	//	log.Fatalln("did not connect:", err)
	//}
	//defer conn.Close()
	//c := pb.NewGreeterClient(conn)

	// 2.
	certFile, keyFile := "./client.pem", "./client.key"
	certificate, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		log.Println("err:", err)
	}
	// 构建CertPool以校验服务端证书有效性
	certPool := x509.NewCertPool()
	ca, _ := ioutil.ReadFile("../ca.pem")
	certPool.AppendCertsFromPEM(ca)
	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{certificate}, //客户端证书
		ServerName:   "seuzjy.cn",                    //注意这里的参数为配置文件中所允许的ServerName，也就是其中配置的DNS...
		RootCAs:      certPool,
	})
	// 建立连接时指定要加载的拦截器
	conn, err := grpc.Dial(*addr,
		grpc.WithTransportCredentials(creds),
		grpc.WithUnaryInterceptor(unaryInterceptor),
		grpc.WithStreamInterceptor(streamInterceptor),
	)
	if err != nil {
		log.Fatalln("did not connect:", err)
	}
	defer conn.Close()
	c := pb2.NewGreeterClient(conn)
	// 执行rpc调用
	getSayHello(c)
	//getLotsOfHello(c)
	//getLotsOfName(c)
	//getBidiHello(c)
	//unaryGetWithMetadata(c, *name)
	//bidirectionalWithMetadata(c, []string{"world", "tiam", "chuyu"})

	//unaryEcho(c, 1, "hello world", code.Code_OK)
	//unaryEcho(c, 2, "delay", code.Code_DEADLINE_EXCEEDED)
	//unaryEcho(c, 3, "[propagate me]nihao", code.Code_OK)
}
