package code

import "zjhx.com/pkg/xcode"

var (
	UserNotExistErr  = xcode.New(10001, "邮箱或密码错误")
	UserExistErr     = xcode.New(10002, "该邮箱已被注册")
	UserUnExistError = xcode.New(10003, "该用户不存在")
	GetUserError     = xcode.New(10004, "获取用户信息错误")
)
