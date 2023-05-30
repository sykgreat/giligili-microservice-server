package svc

import (
	"giligili/app/captcha/rpc/internal/config"
	"giligili/app/email/rpc/emailservice"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	Redis    *redis.Redis
	EmailRpc emailservice.EmailService
}

func NewServiceContext(c config.Config) *ServiceContext {
	redisC, err := redis.NewRedis(c.Redis.RedisConf)
	if err != nil {
		logx.Error("redis connect error: %v", err)
		return nil
	}

	return &ServiceContext{
		Config:   c,
		Redis:    redisC,
		EmailRpc: emailservice.NewEmailService(zrpc.MustNewClient(c.EmailRpcConf)),
	}
}
