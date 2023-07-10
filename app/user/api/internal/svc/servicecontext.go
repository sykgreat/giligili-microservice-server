package svc

import (
	"giligili/app/user/api/internal/config"
	"giligili/app/user/api/internal/middleware"
	"giligili/app/user/rpc/userservice"
	"giligili/common/jwt"
	common "giligili/common/middleware"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config            config.Config
	JwtAuthMiddleware rest.Middleware
	GetUserIp         rest.Middleware
	Jwt               *jwt.Jwt
	UserRpc           userservice.UserService
}

func NewServiceContext(c config.Config) *ServiceContext {
	j := &c.Jwt
	return &ServiceContext{
		Config:            c,
		JwtAuthMiddleware: common.NewJwtAuthMiddleware(j, redis.MustNewRedis(c.Redis)).Handle,
		GetUserIp:         middleware.NewGetUserIpMiddleware().Handle,
		Jwt:               j,
		UserRpc:           userservice.NewUserService(zrpc.MustNewClient(c.UserRpcConf)),
	}
}
