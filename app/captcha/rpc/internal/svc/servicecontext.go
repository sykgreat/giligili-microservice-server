package svc

import (
	"giligili/app/captcha/rpc/internal/config"
	"giligili/app/email/rpc/emailservice"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config             // 基本配置文件
	Redis    *redis.Redis              // 定义redis
	EmailRpc emailservice.EmailService // 定义emailRpc服务
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 初始化redis
	redisC, err := redis.NewRedis(c.Redis.RedisConf)
	if err != nil {
		logx.Error("redis connect error: %v", err)
		return nil
	}

	return &ServiceContext{
		Config:   c,
		Redis:    redisC,
		EmailRpc: emailservice.NewEmailService(zrpc.MustNewClient(c.EmailRpcConf)), // 初始化emailRpc服务
	}
}
