package svc

import (
	"giligili/app/user/api/internal/config"
	"giligili/app/user/api/internal/middleware"
	"giligili/app/user/rpc/userservice"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	GetUserIp rest.Middleware

	UserRpc userservice.UserService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		GetUserIp: middleware.NewGetUserIpMiddleware().Handle,

		UserRpc: userservice.NewUserService(zrpc.MustNewClient(c.UserRpcConf)),
	}
}
