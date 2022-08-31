package main

import (
	"errors"
	"fmt"
	common "goLearn/chatRoom/common/message"
	process2 "goLearn/chatRoom/server/process"
	"goLearn/chatRoom/server/utils"
	"net"
)

type Processor struct {
	Conn net.Conn
}

// 根据消息种类不同，调用不同函数
func (p *Processor) serverProcessMsg(msg *common.Message) (err error) {

	up := &process2.UserProcess{Conn: p.Conn}
	switch msg.Type {
	case common.LoginInfoType:
		err = up.ServiceProcessLogin(msg)
		if err != nil {
			return err
		}
	case common.RegisterType:
	default:
		err = errors.New("消息类型不存在")
	}
	return err
}

func (p *Processor) process2() (err error) {
	// 循环读取
	for {
		tf := &utils.Transfer{
			Conn: p.Conn,
			Buf:  [4096]byte{},
		}
		msg, err := tf.ReadMessage()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("msg =", msg)
		err = p.serverProcessMsg(&msg)
		if err != nil {
			return err
		}
	}
}
