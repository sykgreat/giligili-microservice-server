package config

import (
	"giligili/common"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf

	Mysql struct {
		DataSource string
	}

	Cache cache.CacheConf

	// Snowflake
	Snowflake common.Snowflake

	// Upload
	Upload struct {
		ImgPath   string
		VideoPath string
	}
}
