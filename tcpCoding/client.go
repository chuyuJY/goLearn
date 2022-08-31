package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func clientDemo() {
	conn, err := net.Dial("tcp", "192.168.0.108:8888")
	if err != nil {
		fmt.Println("client dial err =", err)
		return
	}
	fmt.Println("client conn is:", conn)
	clientProcess(conn)
}

func clientProcess(conn net.Conn) {
	defer conn.Close()
	for {
		fmt.Print("请输入要发送的信息: ")
		// 功能1：客户端发送单行数据，然后退出
		reader := bufio.NewReader(os.Stdin) // os.Stdin 代表标准输入（终端）
		// 从终端读取一行用户输入，并准备发给服务器
		message, err := reader.ReadString('\n') // 会多读一个'\n'
		if err != nil {
			fmt.Println("ReadString err =", err)
		}
		message = strings.Trim(message, " \r\n") // 删掉 " \r\n"
		// 再将message发送给服务器
		n, err := conn.Write([]byte(message))
		if err != nil {
			fmt.Println("conn.Write err =", err)
		}
		fmt.Printf("message长度为%v 发送成功......\n", n)
		// 当输入 exit 退出
		if message == "exit" {
			break
		}
	}
	fmt.Println("客户端退出......")
}

func main() {
	clientDemo()
}
