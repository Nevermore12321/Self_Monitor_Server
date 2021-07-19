package main

import (
	"flag"
	"fmt"
	"github.com/tal-tech/go-zero/rest/httpx"

	Herr "api/errors"
	Gmiddleware "api/global-middleware"
	"api/internal/config"
	"api/internal/handler"
	"api/internal/svc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"
)

var configFile = flag.String("f", "etc/api-gateway.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	// 自定义错误
	httpx.SetErrorHandler(Herr.ErrHandler)

	// 添加全局的 跨域访问 中间件
	server.Use(Gmiddleware.CorssDomain)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
