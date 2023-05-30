package main

import (
	"flag"
	"fmt"
	"giligili/app/captcha/utils/captcha"

	"giligili/app/captcha/rpc/internal/config"
	"giligili/app/captcha/rpc/internal/server"
	"giligili/app/captcha/rpc/internal/svc"
	"giligili/app/captcha/rpc/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/captcha.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	captcha.NewCaptcha(c.Captcha.Length, c.Captcha.Chars, c.Captcha.Expire)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterCaptchaServiceServer(grpcServer, server.NewCaptchaServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
