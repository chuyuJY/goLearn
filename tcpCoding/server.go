package main

import (
	"fmt"
	"net"
)

func serverDemo() {
	fmt.Println("服务器开始监听......")
	// 1. 开启一个监听端口
	listener, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("listen err =", err)
		return
	}
	defer listener.Close()
	fmt.Println("listener is:", listener)
	// 2. 等待连接
	for {
		fmt.Println("等待客户端连接......")
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept err =", err)
		}
		fmt.Println("conn is:", conn)
		fmt.Println("该连接的客户端ip =", conn.RemoteAddr()) // 获得客户端的ip地址
		// 接收数据
		go serverProcess(conn)
	}
}

// 接收客户端数据
func serverProcess(conn net.Conn) {
	// 循环接收
	defer conn.Close() // 一定要记得关闭
	fmt.Printf("server 等待 客户端:%v 发送信息......\n", conn.RemoteAddr())
	for {
		// 创建切片，用来接收
		buf := make([]byte, 1024)
		// 1. 等待客户端通过conn发送信息
		// 2. 如果客户端没有write[发送信息]，那么协程就阻塞在这里
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("server conn.Read err =", err)
			return
		}
		//fmt.Println("读取message长度为:", n)
		// 显示message
		message := string(buf[:n]) // 注意要[:n]
		fmt.Printf("信息: %v 来自客户端:%v\n", message, conn.RemoteAddr())
		if message == "exit" {
			break
		}
	}
	fmt.Printf("客户端: %v 退出......\n", conn.RemoteAddr())
}

func main() {
	serverDemo()
}
