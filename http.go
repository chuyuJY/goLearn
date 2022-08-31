package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/go-querystring/query"
	"io/ioutil"
	"net/http"
	"net/url"
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
	GetUserInfoPath = "WeBASE-Sign/user/%s/userInfo"
	DeleteUserPath  = "WeBASE-Sign/user"
	GetVersionPath  = "WeBASE-Sign/version"
)

type Data struct {
	SignUserId  string `json:"signUserId,omitempty"`
	AppId       string `json:"appId,omitempty"`
	Address     string `json:"address,omitempty"`
	PublicKey   string `json:"publicKey,omitempty"`
	PrivateKey  string `json:"privateKey,omitempty"`
	Description string `json:"description,omitempty"`
	EncryptType int    `json:"encryptType,omitempty"`
}

type NewWsInfo struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Data    `json:"data,omitempty"`
}

type Register struct {
	SignUserId  string `url:"signUserId"`
	AppId       string `url:"appId"`
	EncryptType int    `url:"encryptType"`
}

type WeBase struct {
	Host string
}

func InitWeBaseInstance(weBaseHost string) (WeBase, error) {
	weBase := WeBase{Host: weBaseHost}
	_, err := weBase.GetVersion()
	if err != nil {
		return WeBase{}, errors.New("Init WeBaseInstance Failed")
	}
	return weBase, nil
}

// GetNewUser 新增WeBase-Sign用户
func (ws *WeBase) GetNewUser(studentID string) error {
	type Register struct {
		SignUserId       string `url:"signUserId"`
		AppId            string `url:"appId"`
		EncryptType      int    `url:"encryptType"`
		ReturnPrivateKey bool   `url:"returnPrivateKey"`
	}
	wsAccount := &Register{
		SignUserId:       studentID,
		AppId:            AppId,
		EncryptType:      EncryptType,
		ReturnPrivateKey: ReturnPrivateKey, // 默认false
	}
	body, _ := query.Values(wsAccount)
	wsUrl := url.URL{
		Scheme:   "http",
		Host:     ws.Host, // 47.100.21.147:5004
		Path:     GetNewUserPath,
		RawQuery: body.Encode(),
	}
	resp, err := http.Get(wsUrl.String())
	if err != nil {
		return errors.New("Request WeBase-Sign Failed")
	}
	defer resp.Body.Close()

	respBody, _ := ioutil.ReadAll(resp.Body)
	var respInfo map[string]interface{}
	json.Unmarshal(respBody, &respInfo)
	if result := respInfo["message"].(string); result != "success" {
		return errors.New(result)
	}

	return nil
}

func (ws *WeBase) PostNewUser(studentID string, privateKey string) {}

func (ws *WeBase) GetUserInfo(studentID string) {

}

func (ws *WeBase) DeleteUser(studentID string) (bool, error) {

}

// GetVersion 获取版本号
func (ws *WeBase) GetVersion() (string, error) {
	wsUrl := url.URL{
		Scheme: "http",
		Host:   ws.Host,
		Path:   GetVersionPath,
	}
	resp, err := http.Get(wsUrl.String())
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
	webase, err := InitWeBaseInstance("47.100.21.147:5004")
	if err != nil {
		fmt.Println(err)
	}
	err = webase.GetNewUser("102312346")
	if err != nil {
		fmt.Println(err)
	}
}
