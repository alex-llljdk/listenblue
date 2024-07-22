package code

import "zjhx.com/pkg/xcode"

var (
	UploadErr    = xcode.New(20001, "文件上传失败")
	CreateRecordErr = xcode.New(20001, "数据库错误")
)
