package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"goLearn/gRPC-learn/servers/05-authorize/pb"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
	"io/ioutil"
	"log"
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
	listen, err := net.Listen("tcp", "0.0.0.0:8972")
	if err != nil {
		fmt.Println(err)
		return
	}
	// TLS 认证
	certFile, keyFile := "../gRPC-learn/servers/05-authorize/server.crt", "../gRPC-learn/servers/05-authorize/server.key"

	// 从证书相关文件中读取和解析信息，得到证书公钥、密钥对
	certificate, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		log.Fatalf("failed to load key pair: %s", err)
	}
	// 创建CertPool，后续就用池里的证书来校验客户端证书有效性
	// 所以如果有多个客户端 可以给每个客户端使用不同的 CA 证书，来实现分别校验的目的
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("../gRPC-learn/ca_central/ca.pem")
	if err != nil {
		log.Println("err:", err)
	}
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Println("err:", err)
	}
	creds := credentials.NewTLS(&tls.Config{
		// 设置证书链，允许包含一个或多个
		Certificates: []tls.Certificate{certificate},
		// 要求必须校验客户端的证书 可以根据实际情况选用其他参数
		ClientAuth: tls.RequireAndVerifyClientCert, // NOTE: this is optional!
		// 设置根证书的集合，校验方式使用 ClientAuth 中设定的模式
		ClientCAs: certPool,
	})
	s := grpc.NewServer(grpc.Creds(creds))                        // 创建grpc服务器
	pb.RegisterGreeterServer(s, &server{count: map[string]int{}}) // 注册服务
	err = s.Serve(listen)                                         // 启动服务
	if err != nil {
		fmt.Println("服务启动失败：", err)
		return
	}
}
