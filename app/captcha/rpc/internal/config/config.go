package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf

	// 验证码配置
	Captcha struct {
		Length int8
		Chars  string
		Expire int
	}

	// emailRpc服务配置
	EmailRpcConf zrpc.RpcClientConf
}
