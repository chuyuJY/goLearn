package process

import (
	"encoding/json"
	"fmt"
	utils2 "goLearn/chatRoom/client/utils"
	common "goLearn/chatRoom/common/message"
	"net"
)

type UserProcess struct {
}

func (up *UserProcess) Login(userId int, userPwd string) (err error) {
	conn, err := net.Dial("tcp", "192.168.0.105:8889")
	if err != nil {
		fmt.Println("net.Dial err =", err)
		return err
	}
	defer conn.Close()
	data, err := json.Marshal(common.LoginInfo{
		UserId:  userId,
		UserPwd: userPwd,
	})
	if err != nil {
		fmt.Println("json.Marshal err =", err)
		return err
	}
	message := common.Message{
		Type: common.LoginInfoType,
		Data: string(data),
	}
	data, err = json.Marshal(message)
	if err != nil {
		fmt.Println("json.Marshal err =", err)
		return err
	}
	tf := &utils2.Transfer{
		Conn: conn,
		Buf:  [4096]byte{},
	}
	err = tf.WriteMessage(data)
	if err != nil {
		fmt.Println("tf.WriteMessage err =", err)
		return err
	}
	msg, err := tf.ReadMessage()
	if err != nil {
		fmt.Println("tf.ReadMessage err=", err)
		return err
	}
	loginResult := common.LoginResponse{}
	err = json.Unmarshal([]byte(msg.Data), &loginResult)
	fmt.Println(loginResult)
	if loginResult.Code == 200 {
		go serverProcessMsg(conn)
		for {
			ShowMenu()
		}
	} else if loginResult.Code == 500 {
		fmt.Println(loginResult.Error)
	}
	return err
}
