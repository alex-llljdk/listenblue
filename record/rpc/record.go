package main

import (
	"flag"
	"fmt"

	"zjhx.com/pkg/interceptors"
	"zjhx.com/record/rpc/internal/config"
	"zjhx.com/record/rpc/internal/server"
	"zjhx.com/record/rpc/internal/svc"
	"zjhx.com/record/rpc/record"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/record.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		record.RegisterRecordServer(grpcServer, server.NewRecordServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	_ = consul.RegisterService(c.ListenOn, c.Consul)
	defer s.Stop()
	s.AddUnaryInterceptors(interceptors.ServerErrorInterceptor())

	fmt.Printf("Starting record rpc server at %s...\n", c.ListenOn)
	s.Start()
}
