package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Args1 struct {
	X, Y int
}

func main() {
	// 建立http连接
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:9091")
	if err != nil {
		log.Fatalln("rpc.DialHTTP err:", err)
	}
	// 1. 同步调用
	args := &Args1{
		X: 12,
		Y: 3,
	}
	var reply int
	err = client.Call("ServiceA.Add", args, &reply)
	if err != nil {
		log.Fatalln("ServiceA.Add err:", err)
	}
	fmt.Println("ServiceA.Add(args, &reply) =", reply)

	// 2. 异步调用
	var reply2 int
	divCall := client.Go("ServiceA.Add", args, &reply2, nil)
	replyCall := <-divCall.Done // 接收结果
	fmt.Println(replyCall.Error)
	fmt.Println(reply2)
}
