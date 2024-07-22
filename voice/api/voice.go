package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
	_ "github.com/zeromicro/zero-contrib/zrpc/registry/consul"
	"zjhx.com/pkg/xcode"
	"zjhx.com/voice/api/internal/config"
	"zjhx.com/voice/api/internal/handler"
	"zjhx.com/voice/api/internal/svc"
	"zjhx.com/voice/api/ws"
)

var configFile = flag.String("f", "etc/voice-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf, rest.WithCustomCors(nil, func(w http.ResponseWriter) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		w.Header().Set("Access-Control-Expose-Headers", "Content-Length, Content-Type, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
	}, "*"))
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	server.AddRoute(rest.Route{
		Method: http.MethodGet,
		Path:   "/voice/ws",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			ws.ServeWs(w, r)
		},
	})
	//自定义错误处理方法
	httpx.SetErrorHandler(xcode.ErrHandler)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
