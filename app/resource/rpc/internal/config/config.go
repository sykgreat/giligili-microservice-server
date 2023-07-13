package config

import (
	"giligili/common/Snowflake"
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
	Snowflake Snowflake.Snowflake

	// Upload
	Upload struct {
		ImgPath   string
		VideoPath string
	}
}
