package config

import (
	"giligili/common"
	"giligili/common/jwt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf

	Mysql struct {
		DataSource string
	}

	Cache cache.CacheConf

	CaptchaRpcConf zrpc.RpcClientConf

	// Snowflake
	Snowflake common.Snowflake

	// Jwt
	Jwt jwt.Jwt
}
