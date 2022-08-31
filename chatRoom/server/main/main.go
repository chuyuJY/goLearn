package main

import (
	"fmt"
	"goLearn/chatRoom/server/model"
	"net"
	"time"
)

func process(conn net.Conn) {
	defer conn.Close()
	processor := &Processor{Conn: conn}
	err := processor.process2()
	if err != nil {
		fmt.Println("main process err =", err)
	}
}

// 完成UserDao的初始化
func initUserDao() {
	model.MyUserDao = model.NewUserDao(pool)
}

func main() {
	// 当服务器启动，就初始化redis pool
	initPool("localhost:6379", 16, 0, 300*time.Second)
	initUserDao()
	// 提示
	fmt.Println("服务器在8889端口监听......")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	defer listen.Close()
	if err != nil {
		fmt.Println("net.Listen err =", err)
		return
	}
	// 一旦监听成功，就等待客户端来连接
	for {
		fmt.Println("等待客户端连接服务器......")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err =", err)
			continue
		}
		// 一旦连接成功，启动协程保持运行
		go process(conn)
	}
}
