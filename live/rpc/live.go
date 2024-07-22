package main

import (
	"flag"
	"fmt"

	"zjhx.com/live/rpc/common"
	"zjhx.com/live/rpc/internal/config"
	"zjhx.com/live/rpc/internal/server"
	"zjhx.com/live/rpc/internal/svc"
	"zjhx.com/live/rpc/live"
	"zjhx.com/pkg/interceptors"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/live.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		live.RegisterLiveServer(grpcServer, server.NewLiveServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	_ = consul.RegisterService(c.ListenOn, c.Consul)
	defer s.Stop()
	s.AddUnaryInterceptors(interceptors.ServerErrorInterceptor())

	//轮询生成直播封面
	
	go common.GetLiveCover(ctx)

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
