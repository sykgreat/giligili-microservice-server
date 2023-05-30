package main

import (
	"flag"
	"fmt"
	"giligili/app/email/utils/email"

	"giligili/app/email/rpc/internal/config"
	"giligili/app/email/rpc/internal/server"
	"giligili/app/email/rpc/internal/svc"
	"giligili/app/email/rpc/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/email.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	email.NewEmail(c.Email.User, c.Email.Name, c.Email.Password, c.Email.Host+":"+c.Email.Port, c.Email.MailType) // 初始化邮件服务

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterEmailServiceServer(grpcServer, server.NewEmailServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
