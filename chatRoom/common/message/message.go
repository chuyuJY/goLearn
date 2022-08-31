package common

const (
	LoginInfoType     = "LoginInfo"
	LoginResponseType = "LoginResponse"
	RegisterType      = "Register"
)

type Message struct {
	Type string `json:"type,omitempty"`
	Data string `json:"data,omitempty"`
}

// 定义消息
type LoginInfo struct {
	UserId  int    `json:"userId,omitempty"`
	UserPwd string `json:"userPwd,omitempty"`
}

type LoginResponse struct {
	Code  int    `json:"code,omitempty"`  // 返回状态码
	Error string `json:"error,omitempty"` // 返回错误信息
}

type Register struct {
}
