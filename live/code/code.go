package code

import "zjhx.com/pkg/xcode"

var (
	RoomExistErr    = xcode.New(40001, "该房间已被占用")
	RedisErr        = xcode.New(40002, "Redis错误")
	RoomNotExistErr = xcode.New(40003, "该房间不存在")
	PasswordErr     = xcode.New(40004, "房间密码错误")
	InsertRecordErr = xcode.New(40005, "插入记录错误")
	LiveUrl         = "rtmp://127.0.0.1:1935/live"
	LiveDataUrl     = "C:/Users/fjf/Desktop/backend/data/savel"
	LiveNetUrl      = "http://172.30.230.46/savel"
)
