package svc

import (
	"giligili/app/captcha/api/internal/config"
	"giligili/app/captcha/rpc/captchaservice"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	CaptchaRpc captchaservice.CaptchaService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		CaptchaRpc: captchaservice.NewCaptchaService(zrpc.MustNewClient(c.CaptchaRpcConf)),
	}
}
