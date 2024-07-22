package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"zjhx.com/pkg/interceptors"
	"zjhx.com/voice/api/internal/config"
	"zjhx.com/voice/rpc/voiceclient"
)

type ServiceContext struct {
	Config   config.Config
	VoiceRpc voiceclient.Voice
}

func NewServiceContext(c config.Config) *ServiceContext {
	VoiceRPC := zrpc.MustNewClient(c.VoiceRpc, zrpc.WithUnaryClientInterceptor(interceptors.ClientErrorInterceptor()))
	return &ServiceContext{
		Config:   c,
		VoiceRpc: voiceclient.NewVoice(VoiceRPC),
	}
}
