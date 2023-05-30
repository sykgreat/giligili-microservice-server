package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf

	Email struct {
		User     string
		Name     string
		Password string
		Host     string
		MailType string
		Port     string
	}
}
