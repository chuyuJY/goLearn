package model

import "errors"

// 根据业务逻辑的需要，定义错误信息
var (
	ERROR_USER_NOTEXISTS  = errors.New("用户不存在")
	ERROR_USER_EXISTS     = errors.New("用户已存在")
	ERROR_USER_PWDINVALID = errors.New("密码错误")
)
