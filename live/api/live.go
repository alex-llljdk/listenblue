package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	_ "github.com/zeromicro/zero-contrib/zrpc/registry/consul"
	"zjhx.com/live/api/chat"
	"zjhx.com/live/api/internal/config"
	"zjhx.com/live/api/internal/handler"
	"zjhx.com/live/api/internal/svc"
	"zjhx.com/live/api/subtitle"
	"zjhx.com/pkg/xcode"
)

var configFile = flag.String("f", "etc/live-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf, rest.WithUnauthorizedCallback(xcode.JwtUnauthorizedResult), rest.WithCustomCors(nil, func(w http.ResponseWriter) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		w.Header().Set("Access-Control-Expose-Headers", "Content-Length, Content-Type, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
	}, "*"))
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	hub := chat.NewHub()
	go hub.Run()

	handler.RegisterHandlers(server, ctx, hub)

	server.AddRoute(rest.Route{
		Method: http.MethodGet,
		Path:   "/live/chat",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			chat.ServeWs(hub, w, r)
		},
	})
	fmt.Println(1)
	server.AddRoute(rest.Route{
		Method: http.MethodGet,
		Path:   "/live/subtitle",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			subtitle.ServeWs(w, r)
		},
	})
	fmt.Println(2)
	//自定义错误处理方法
	// httpx.SetErrorHandler(xcode.ErrHandler)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
