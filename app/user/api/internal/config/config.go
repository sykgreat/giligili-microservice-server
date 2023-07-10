package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	Jwt struct {
		AccessTokenSecret  string
		AccessExpire       int64
		RefreshExpire      int64
		RefreshTokenSecret string
	}
}
