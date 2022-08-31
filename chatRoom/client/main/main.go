package main

import (
	"fmt"
	process2 "goLearn/chatRoom/client/process"
)

func main() {
	// 接收用户选择
	key := -1
	// 判断是否继续循环
	userId, userPwd := 0, ""
	for {
		fmt.Println("-----------------欢迎登录多人聊天系统-----------------")
		fmt.Println("                 1. 登录聊天室")
		fmt.Println("                 2. 注册用户")
		fmt.Println("                 3. 退出系统")
		fmt.Println("---------------------------------------------------")
		fmt.Printf("                 请选择【1-3】：")
		fmt.Scanln(&key)
		switch key {
		case 1:
			fmt.Println("登录聊天室")
			fmt.Printf("请输入用户Id：")
			fmt.Scanln(&userId)
			fmt.Printf("请输入用户密码：")
			fmt.Scanln(&userPwd)
			up := &process2.UserProcess{}
			up.Login(userId, userPwd)
		case 2:
			fmt.Println("注册用户")
		case 3:
			fmt.Println("退出系统")
		default:
			fmt.Println("输入有误，请重新输入...")
		}
	}

}
