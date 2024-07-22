package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"zjhx.com/pkg/interceptors"
	"zjhx.com/record/api/internal/config"
	"zjhx.com/record/rpc/recordclient"
)

type ServiceContext struct {
	Config    config.Config
	RecordRpc recordclient.Record
}

func NewServiceContext(c config.Config) *ServiceContext {
	RecordRpc := zrpc.MustNewClient(c.RecordRpc, zrpc.WithUnaryClientInterceptor(interceptors.ClientErrorInterceptor()))
	return &ServiceContext{
		Config:    c,
		RecordRpc: recordclient.NewRecord(RecordRpc),
	}
}
