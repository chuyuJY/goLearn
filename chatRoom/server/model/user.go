package model

// User 定义用户info
type User struct {
	UserId   int    `json:"userId,omitempty"`
	UserPwd  string `json:"userPwd,omitempty"`
	UserName string `json:"userName,omitempty"`
}
