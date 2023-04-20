package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"goLearn/gRPC-learn/clients/04-error/pb"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
	"io/ioutil"
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
	certFile, keyFile := "../gRPC-learn/clients/05-authorize/client.crt", "../gRPC-learn/clients/05-authorize/client.key"
	certificate, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		log.Println("err:", err)
	}
	// 构建CertPool以校验服务端证书有效性
	certPool := x509.NewCertPool()
	ca, _ := ioutil.ReadFile("../gRPC-learn/ca_central/ca.pem")
	certPool.AppendCertsFromPEM(ca)
	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{certificate}, //客户端证书
		ServerName:   "seuzjy.cn",                    //注意这里的参数为配置文件中所允许的ServerName，也就是其中配置的DNS...
		RootCAs:      certPool,
	})
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalln("did not connect:", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)
	getSayHello(c)
}
