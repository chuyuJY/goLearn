package process

import (
	"fmt"
	utils2 "goLearn/chatRoom/client/utils"
	"net"
)

func ShowMenu() {
	fmt.Println("-----------------恭喜xxx登录成功-----------------")
	fmt.Println("-----------------1. 显示在线用户列表-----------------")
	fmt.Println("-----------------2. 发送消息-----------------")
	fmt.Println("-----------------3. 信息列表-----------------")
	fmt.Println("-----------------4. 退出系统-----------------")
	fmt.Printf("请选择【1-4】：")
	key := 0
	fmt.Scanln(&key)
	switch key {
	case 1:
		fmt.Println("显示")
	case 2:
		fmt.Println("发送消息")
	case 3:
		fmt.Println("信息列表")
	case 4:
		fmt.Println("退出系统")
	default:
		fmt.Println("输入有误，请重新输入...")
	}
}

func serverProcessMsg(conn net.Conn) {
	tf := &utils2.Transfer{
		Conn: conn,
		Buf:  [4096]byte{},
	}
	for {
		fmt.Printf("客户端: %v 正在等待服务器的消息...", conn.RemoteAddr())
		msg, err := tf.ReadMessage()
		if err != nil {
			fmt.Println("tf.ReadMessage err =", err)
			return
		}
		fmt.Println("msg =", msg)
	}

}
