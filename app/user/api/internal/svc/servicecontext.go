package svc

import (
	"giligili/app/user/api/internal/config"
	"giligili/app/user/api/internal/middleware"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config            config.Config
	JwtAuthMiddleware rest.Middleware
	GetUserIp         rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:            c,
		JwtAuthMiddleware: middleware.NewJwtAuthMiddleware().Handle,
		GetUserIp:         middleware.NewGetUserIpMiddleware().Handle,
	}
}
