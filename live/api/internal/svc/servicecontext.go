package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"zjhx.com/live/api/internal/config"
	"zjhx.com/live/rpc/liveclient"
	"zjhx.com/pkg/interceptors"
)

type ServiceContext struct {
	Config  config.Config
	LiveRpc liveclient.Live
}

func NewServiceContext(c config.Config) *ServiceContext {
	LiveRPC := zrpc.MustNewClient(c.LiveRpc, zrpc.WithUnaryClientInterceptor(interceptors.ClientErrorInterceptor()))
	return &ServiceContext{
		Config:  c,
		LiveRpc: liveclient.NewLive(LiveRPC),
	}
}
