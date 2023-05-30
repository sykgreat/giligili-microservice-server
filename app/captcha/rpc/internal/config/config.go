package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf

	Captcha struct {
		Length int8
		Chars  string
		Expire int
	}

	EmailRpcConf zrpc.RpcClientConf
}
