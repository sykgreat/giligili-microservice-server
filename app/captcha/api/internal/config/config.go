package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf

	// captchaRpc服务配置
	CaptchaRpcConf zrpc.RpcClientConf
}
