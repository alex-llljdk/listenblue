package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"zjhx.com/pkg/interceptors"
	"zjhx.com/video/api/internal/config"
	"zjhx.com/video/rpc/videoclient"
)

type ServiceContext struct {
	Config   config.Config
	VideoRpc videoclient.Video
}

func NewServiceContext(c config.Config) *ServiceContext {

	VideoRPC := zrpc.MustNewClient(c.VideoRpc, zrpc.WithUnaryClientInterceptor(interceptors.ClientErrorInterceptor()))
	return &ServiceContext{
		Config:   c,
		VideoRpc: videoclient.NewVideo(VideoRPC),
	}
}
