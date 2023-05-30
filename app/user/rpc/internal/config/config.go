package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf

	// mysql配置
	Mysql struct {
		DataSource string
	}
	// redis配置
	Cache cache.CacheConf
	// captchaRpc配置
	CaptchaRpcConf zrpc.RpcClientConf
	// jwt配置
	JwtAuth struct {
		AccessSecret string
		AccessExpire int64
	}
	// Snowflake配置
	Snowflake struct {
		WorkerName     string
		DatacenterName string
		Sequence       int64
	}
}
