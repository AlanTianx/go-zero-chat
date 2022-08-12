package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"go-zero-chat/apps/app/api/internal/config"
	"go-zero-chat/apps/app/api/internal/handler"
	"go-zero-chat/apps/app/api/internal/handler/ws"
	"go-zero-chat/apps/app/api/internal/svc"
)

var configFile = flag.String("f", "etc/api-api.yaml", "the config file")

func main() {

	flag.Parse()

	_ = logx.SetUp(logx.LogConf{
		ServiceName: "api-api",
		Mode:        "file",
		Encoding:    "default",
		Path:        "logs",
		KeepDays:    3,
	})

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	// 初始化配置ws
	ws.WsServerInit(ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}