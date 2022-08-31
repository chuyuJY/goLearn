package main

import (
	"goLearn/rpc/02/service"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

func main() {
	stringService := new(service.StringService)
	rpc.Register(stringService)
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", "127.0.0.1:9091")
	if err != nil {
		log.Fatalln(err)
	}
	http.Serve(l, nil)
}
