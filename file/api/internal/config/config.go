package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	FileRpc zrpc.RpcClientConf
	Auth    struct {
		AccessSecret string
		AccessExpire int64
	}
	FilePath string
	Salt     string
}
