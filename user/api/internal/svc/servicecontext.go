package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"zjhx.com/pkg/interceptors"
	"zjhx.com/user/rpc/userclient"

	"zjhx.com/user/api/internal/config"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	//自定义拦截器
	UserRPC := zrpc.MustNewClient(c.UserRpc, zrpc.WithUnaryClientInterceptor(interceptors.ClientErrorInterceptor()))
	return &ServiceContext{
		Config:  c,
		UserRpc: userclient.NewUser(UserRPC),
		// UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
