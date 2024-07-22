package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"zjhx.com/file/api/internal/config"
	"zjhx.com/file/rpc/fileclient"
	"zjhx.com/pkg/interceptors"
)

type ServiceContext struct {
	Config  config.Config
	FileRpc fileclient.File
}

func NewServiceContext(c config.Config) *ServiceContext {
	FileRpc := zrpc.MustNewClient(c.FileRpc, zrpc.WithUnaryClientInterceptor(interceptors.ClientErrorInterceptor()))
	return &ServiceContext{
		Config:  c,
		FileRpc: fileclient.NewFile(FileRpc),
	}
}
