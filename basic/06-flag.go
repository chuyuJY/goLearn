package main

import (
	"flag"
	"fmt"
)

func parseCommand() {
	// 1. 定义几个变量用来存放命令行参数
	var (
		user string
		pwd  string
		host string
		port int
	)
	// 2. 转换
	/*
		&users: 接收用户命令行输入的 -u 后面的参数值
		"u": -u 指定参数
		"": 默认参数
		"用户名，默认为空": 说明
	*/
	flag.StringVar(&user, "u", "", "用户名，默认为空")
	flag.StringVar(&pwd, "pwd", "", "密码，默认为空")
	flag.StringVar(&host, "h", "localhost", "主机名，默认为localhost")
	flag.IntVar(&port, "port", 3306, "端口号，默认为3306")
	flag.Parse() // 一定要调用
	// 输出结果
	fmt.Printf("users=%v pwd=%v host=%v port=%v", user, pwd, host, port)
}

func main() {
	parseCommand()
}
