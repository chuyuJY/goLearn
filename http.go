package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/go-querystring/query"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// Webase请求配置
const (
	AppId            = "invoice"
	EncryptType      = 0
	ReturnPrivateKey = false // 默认false
)

// Webase接口路径
const (
	GetNewUserPath  = "WeBASE-Sign/user/newUser"
	PostNewUserPath = "WeBASE-Sign/user/newUser"
	GetUserInfoPath = "WeBASE-Sign/user/{signUserId}/userInfo"
	DeleteUserPath  = "WeBASE-Sign/user"
	GetVersionPath  = "WeBASE-Sign/version"
)

type WeBase struct {
	Host string
}

func InitWeBaseInstance(weBaseHost string) (*WeBase, error) {
	weBase := &WeBase{Host: weBaseHost}
	_, err := weBase.GetVersion()
	if err != nil {
		return &WeBase{}, errors.New("Init WeBaseInstance Failed")
	}
	return weBase, nil
}

// GetNewUser 新增WeBase-Sign用户
func (ws *WeBase) GetNewUser(studentID string) error {
	type options struct {
		signUserId       string
		appId            string
		encryptType      int
		returnPrivateKey bool
	}
	opt := &options{
		signUserId:       studentID,
		appId:            AppId,
		encryptType:      EncryptType,
		returnPrivateKey: ReturnPrivateKey, // 默认false
	}
	body, _ := query.Values(opt)
	wsUrl, _ := url.Parse(ws.Host + GetNewUserPath)
	wsUrl.RawQuery = body.Encode()
	resp, err := http.Get(wsUrl.String())
	if err != nil {
		return errors.New("Request WeBase-Sign Failed")
	}

	respBody, _ := ioutil.ReadAll(resp.Body)
	var respInfo map[string]interface{}
	json.Unmarshal(respBody, &respInfo)
	if result := respInfo["message"].(string); result != "success" {
		return errors.New(result)
	}
	return nil
}

func (ws *WeBase) PostNewUser(studentID string, privateKey string) {

}

// GetUserInfo 查询用户信息 TODO: 身份验证
func (ws *WeBase) GetUserInfo(studentID string, returnPK bool) error {
	wsUrl, _ := url.Parse(ws.Host + regexp.MustCompile("{.*?}").ReplaceAllString(GetUserInfoPath, studentID))
	q := wsUrl.Query()
	q.Set("returnPrivateKey", strconv.FormatBool(returnPK))
	wsUrl.RawQuery = q.Encode()
	resp, err := http.Get(wsUrl.String())
	if err != nil {
		return err
	}
	respBody, _ := ioutil.ReadAll(resp.Body)
	var respInfo map[string]interface{}
	json.Unmarshal(respBody, &respInfo)

	// TODO: 处理返回信息
	fmt.Println(respInfo)
	return nil
}

// DeleteUser 注销用户（未删除）
func (ws *WeBase) DeleteUser(studentID string) (bool, error) {
	wsUrl := ws.Host + DeleteUserPath
	data, err := json.Marshal(map[string]interface{}{"signUserId": studentID})
	if err != nil {
		return false, err
	}

	req, err := http.NewRequest("DELETE", wsUrl, bytes.NewReader(data))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return false, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	var respInfo map[string]interface{}
	json.Unmarshal(body, &respInfo)
	if result := respInfo["message"].(string); result != "success" {
		return false, errors.New(result)
	}
	return true, nil
}

// GetVersion 获取版本号
func (ws *WeBase) GetVersion() (string, error) {
	wsUrl := ws.Host + GetVersionPath
	resp, err := http.Get(wsUrl)
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func main() {
	webase, err := InitWeBaseInstance("http://47.100.21.147:5004/")
	if err != nil {
		fmt.Println(err)
	}

	// 1. 测试新建用户
	//err = webase.GetNewUser("102312346")
	//if err != nil {
	//	fmt.Println(err)
	//}

	// 2. 测试注销用户
	//ok, err := webase.DeleteUser("102312346")
	//if !ok || err != nil {
	//	fmt.Println(err)
	//}

	// 3. 测试查看用户
	webase.GetUserInfo("20177830226", false)
}
