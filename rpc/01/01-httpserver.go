package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Args struct {
	X, Y int
}

type ServiceA struct {
	X, Y int
}

func (sa *ServiceA) Add(args *Args, reply *int) error {
	*reply = args.X + args.Y
	return nil
}

func main() {
	sa := new(ServiceA)
	rpc.Register(sa) // 注册Rpc服务
	rpc.HandleHTTP() // 基于Http服务
	l, err := net.Listen("tcp", ":9091")
	if err != nil {
		log.Fatalln(err)
	}
	http.Serve(l, nil)
}
