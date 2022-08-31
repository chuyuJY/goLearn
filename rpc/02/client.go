package main

import (
	"fmt"
	"goLearn/rpc/02/service"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:9091")
	if err != nil {
		log.Fatalln(err)
	}
	stringReq := &service.StringRequest{
		A: "A",
		B: "B",
	}
	var reply string
	// 1. 同步调用
	err = client.Call("StringService.Concat", stringReq, &reply)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("StringService.Concat: %s concat %s = %s\n", stringReq.A, stringReq.B, reply)

	// 2. 异步调用
	stringReq = &service.StringRequest{
		A: "ACD",
		B: "BDF",
	}
	call := client.Go("StringService.Diff", stringReq, &reply, nil)
	_ = <-call.Done
	fmt.Printf("StringService Diff: %s diff %s = %s", stringReq.A, stringReq.B, reply)
}
