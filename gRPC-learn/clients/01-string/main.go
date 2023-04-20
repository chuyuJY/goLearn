package main

import (
	"context"
	"flag"
	pb "goLearn/gRPC-learn/clients/01-string/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

const (
	defaultA = "hello"
	defaultB = "world"
)

var (
	addr = flag.String("addr", "127.0.0.1:8972", "the address to connect to")
)

func getConcat(c pb.StringServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	reply, err := c.Concat(ctx, &pb.StringRequest{A: defaultA, B: defaultB})
	if err != nil {
		log.Fatalln("could not concat:", err)
	}
	log.Println("concat:", reply.GetRet())
}

func getDiff(c pb.StringServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	reply, err := c.Diff(ctx, &pb.StringRequest{A: defaultA, B: defaultB})
	if err != nil {
		log.Fatalln("could not diff:", err)
	}
	log.Println("diff:", reply.GetRet())
}

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("did not connect: ", err)
		return
	}
	defer conn.Close()
	c := pb.NewStringServiceClient(conn)
	// 执行远程调用
	getConcat(c)
	getDiff(c)
}
