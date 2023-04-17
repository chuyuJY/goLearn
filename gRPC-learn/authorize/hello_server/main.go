package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	pb2 "goLearn/gRPC-learn/authorize/pb"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"google.golang.org/grpc/credentials/insecure"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type server struct {
	pb2.UnimplementedGreeterServer
	mu    sync.Mutex     // count的并发锁
	count map[string]int // 记录每个name访问的次数
}

// SayHello 重写
func (s *server) SayHello(ctx context.Context, in *pb2.HelloRequest) (*pb2.HelloResponse, error) {
	//s.mu.Lock()
	//defer s.mu.Unlock()
	//s.count[in.GetName()]++ // 记录用户访问次数
	//// 错误处理
	//if s.count[in.GetName()] > 1 {
	//	st := status.New(codes.ResourceExhausted, "Request limit exceeded")
	//	ds, err := st.WithDetails(
	//		&errdetails.QuotaFailure{
	//			Violations: []*errdetails.QuotaFailure_Violation{{
	//				Subject:     fmt.Sprintf("name: %s", in.GetName()),
	//				Description: "每个name限制访问次数为1次",
	//			}},
	//		})
	//	if err != nil {
	//		return nil, st.Err()
	//	}
	//	return nil, ds.Err()
	//}
	return &pb2.HelloResponse{Reply: "Hello, " + in.Name}, nil
}

// LotsOfHello 流式响应客户端的请求
func (s *server) LotsOfHello(in *pb2.HelloRequest, stream pb2.Greeter_LotsOfHelloServer) error {
	words := []string{
		"你好",
		"hello",
		"こんにちは",
		"안녕하세요",
	}
	for _, word := range words {
		data := &pb2.HelloResponse{
			Reply: strings.Join([]string{word, in.GetName()}, ", "),
		}
		if err := stream.Send(data); err != nil {
			return err
		}
	}
	return nil
}

// LotsOfName 一次性回复客户端的流式请求
func (s *server) LotsOfName(stream pb2.Greeter_LotsOfNameServer) error {
	reply := "你好呀！"
	names := []string{}
	for {
		// 接收客户端流式数据
		res, err := stream.Recv()
		if err == io.EOF {
			// 接收完毕，一次性回复
			return stream.SendAndClose(&pb2.HelloResponse{Reply: reply + strings.Join(names, ", ")})
		}
		if err != nil {
			return err
		}
		names = append(names, res.GetName())
	}
}

// BidiHello 双向流式Hello
func (s *server) BidiHello(stream pb2.Greeter_BidiHelloServer) error {
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
		if err := stream.Send(&pb2.HelloResponse{Reply: reply}); err != nil {
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

// UnarySayHello 普通RPC调用服务端metadata操作
func (s *server) UnarySayHello(ctx context.Context, in *pb2.HelloRequest) (*pb2.HelloResponse, error) {
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
	return &pb2.HelloResponse{Reply: in.Name}, nil
}

// BidirectionalStreaming 流式RPC调用客户端metadata操作
func (s *server) BidirectionalStreaming(stream pb2.Greeter_BidirectionalStreamingServer) error {
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
	header := metadata.New(map[string]string{"location": "ZhengZhou"})
	stream.SendHeader(header)

	// 3. 读取请求数据并发送响应数据
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		if err := stream.Send(&pb2.HelloResponse{Reply: in.GetName()}); err != nil {
			return err
		}
	}
}

// unaryInterceptor 服务端一元拦截器
func unaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// authentication (token verification)
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.InvalidArgument, "missing metadata")
	}
	if !valid(md["authorization"]) {
		return nil, status.Errorf(codes.Unauthenticated, "invalid token")
	}

	start := time.Now()
	m, err := handler(ctx, req)
	if err != nil {
		fmt.Println("RPC failed with err %v\n", err)
	}
	end := time.Now()
	fmt.Printf("RPC: %s, start time: %s, end time: %s, err: %v\n", info.FullMethod, start.Format(time.RFC3339), end.Format(time.RFC3339), err)
	return m, err
}

type wrappedStream struct {
	grpc.ServerStream
}

func (ws *wrappedStream) RecvMsg(m interface{}) error {
	fmt.Printf("Receive a message (Type: %T) at %v\n", m, time.Now().Format(time.RFC3339))
	return ws.ServerStream.RecvMsg(m)
}

func (ws *wrappedStream) SendMsg(m interface{}) error {
	fmt.Printf("Send a message (Type: %T) at %v\n", m, time.Now().Format(time.RFC3339))
	return ws.ServerStream.SendMsg(m)
}

func newWrappedStream(s grpc.ServerStream) grpc.ServerStream {
	return &wrappedStream{s}
}

// streamInterceptor 服务端流拦截器
func streamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	// authentication (token verification)
	md, ok := metadata.FromIncomingContext(ss.Context())
	if !ok {
		return status.Errorf(codes.InvalidArgument, "missing metadata")
	}
	if !valid(md["authorization"]) {
		return status.Errorf(codes.Unauthenticated, "invalid token")
	}
	start := time.Now()
	err := handler(srv, newWrappedStream(ss))
	if err != nil {
		fmt.Println("RPC failed with err %v\n", err)
	}
	end := time.Now()
	fmt.Printf("RPC: %s, start time: %s, end time: %s, err: %v\n", info.FullMethod, start.Format(time.RFC3339), end.Format(time.RFC3339), err)
	return err
}

func valid(authorization []string) bool {
	if len(authorization) < 1 {
		return false
	}
	//fmt.Println(authorization[0]) // Bearer some-secret-token
	token := strings.TrimPrefix(authorization[0], "Bearer ")
	// 执行token的认证逻辑
	// 为了演示方便简单判断token是否与"some-secret-token"相等
	return token == "some-secret-token"
}

// UnaryEcho 一元Echo
func (s *server) UnaryEcho(ctx context.Context, in *pb2.EchoRequest) (*pb2.EchoResponse, error) {
	message := in.Message
	if strings.HasPrefix(message, "[propagate me]") {
		time.Sleep(800 * time.Millisecond)
		message = strings.TrimPrefix(message, "[propagate me]")
		return &pb2.EchoResponse{Message: message}, nil
	}
	if message == "delay" {
		time.Sleep(2 * time.Second)
	}
	return &pb2.EchoResponse{Message: message}, nil
}

func main() {
	listen, err := net.Listen("tcp", "0.0.0.0:8972")
	if err != nil {
		fmt.Println(err)
		return
	}
	// TLS 认证
	certFile, keyFile := "./server.pem", "./server.key"

	// 1.
	//creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
	//if err != nil {
	//	grpclog.Fatalf("Failed to generate credentials, err: %v", err)
	//}

	// 2.
	// 从证书相关文件中读取和解析信息，得到证书公钥、密钥对
	certificate, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		log.Fatalf("failed to load key pair: %s", err)
	}
	// 创建CertPool，后续就用池里的证书来校验客户端证书有效性
	// 所以如果有多个客户端 可以给每个客户端使用不同的 CA 证书，来实现分别校验的目的
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("../ca.pem")
	if err != nil {
		log.Println("err:", err)
	}
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Println("err:", err)
	}
	_ = credentials.NewTLS(&tls.Config{
		// 设置证书链，允许包含一个或多个
		Certificates: []tls.Certificate{certificate},
		// 要求必须校验客户端的证书 可以根据实际情况选用其他参数
		ClientAuth: tls.RequireAndVerifyClientCert, // NOTE: this is optional!
		// 设置根证书的集合，校验方式使用 ClientAuth 中设定的模式
		ClientCAs: certPool,
	})
	// 创建grpc服务器
	s := grpc.NewServer(
	//grpc.Creds(creds),
	//grpc.UnaryInterceptor(unaryInterceptor),
	//grpc.StreamInterceptor(streamInterceptor),
	)

	pb2.RegisterGreeterServer(s, &server{count: map[string]int{}}) // 注册服务, 注意初始化count
	log.Println("Serving gRPC on 0.0.0.0:8972")
	// go程启动gRPC服务
	go func() {
		log.Fatalln(s.Serve(listen))
	}()

	// 创建一个连接到我们刚刚启动的 gRPC 服务器的客户端连接
	// gRPC-Gateway 就是通过它来代理请求（将HTTP请求转为RPC请求）
	conn, err := grpc.DialContext(
		context.Background(),
		"127.0.0.1:8972",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server")
	}
	gwmux := runtime.NewServeMux()

	// 注册Greeter
	err = pb2.RegisterGreeterHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway, err:", err)
	}
	gwServer := &http.Server{
		Addr:    ":8080",
		Handler: gwmux,
	}
	// 8080端口提供gRPC-Gateway服务
	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8080")
	log.Fatalln(gwServer.ListenAndServe())
}
