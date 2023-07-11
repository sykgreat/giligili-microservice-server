package config

import (
	"giligili/common/jwt"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf

	// Jwt
	Jwt jwt.Jwt

	// Redis
	Redis redis.RedisConf

	UserRpcConf zrpc.RpcClientConf
}
