package process

import (
	"encoding/json"
	"errors"
	"fmt"
	common "goLearn/chatRoom/common/message"
	"goLearn/chatRoom/server/model"
	"goLearn/chatRoom/server/utils"
	"net"
)

type UserProcess struct {
	Conn net.Conn
}

func (up *UserProcess) ServiceProcessLogin(msg *common.Message) (err error) {
	loginInfo := common.LoginInfo{}
	err = json.Unmarshal([]byte(msg.Data), &loginInfo)
	if err != nil {
		err = errors.New("data json.Unmarshal err")
		return err
	}

	// 返回的登录结果
	resultMsg := common.Message{
		Type: common.LoginResponseType,
		Data: "",
	}
	loginResponse := common.LoginResponse{
		Code:  0,
		Error: "",
	}

	//if loginInfo.UserId == 100 && loginInfo.UserPwd == "root" {
	//	loginResponse.Code = 200
	//} else {
	//	loginResponse.Code = 500 // 表示用户不存在
	//	loginResponse.Error = "该用户不存在，请先注册......"
	//}
	user, err := model.MyUserDao.Login(loginInfo.UserId, loginInfo.UserPwd)
	if err != nil {
		loginResponse.Code = 500
		loginResponse.Error = "该用户不存在，请先注册..."
	} else {
		loginResponse.Code = 200
	}
	fmt.Printf("%v 登录成功\n", user)

	data, err := json.Marshal(loginResponse)
	if err != nil {
		fmt.Println("json.Marshal err =", err)
		return
	}
	resultMsg.Data = string(data)
	data, err = json.Marshal(resultMsg)
	if err != nil {
		fmt.Println("json.Marshal err =", err)
		return
	}

	tf := &utils.Transfer{
		Conn: up.Conn,
		Buf:  [4096]byte{},
	}
	err = tf.WriteMessage(data)
	if err != nil {
		fmt.Println("writeMessage err =", err)
	}

	return
}
