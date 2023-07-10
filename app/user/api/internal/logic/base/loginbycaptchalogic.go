package base

import (
	"context"
	"giligili/app/user/rpc/userservice"
	"giligili/common"
	"giligili/common/xerr"
	"github.com/pkg/errors"

	"giligili/app/user/api/internal/svc"
	"giligili/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginByCaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginByCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginByCaptchaLogic {
	return &LoginByCaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginByCaptchaLogic) LoginByCaptcha(req *types.LoginByCaptchaRequest) (resp *types.LoginResponse, err error) {
	// 验证邮箱格式
	if format := common.VerifyEmailFormat(req.Email); !format {
		return nil, errors.Wrapf(xerr.NewErrMsg("邮箱格式不正确!"), "email format is incorrect!")
	}

	// 调用验证码登陆rpc服务
	captcha, err := l.svcCtx.UserRpc.LoginByCaptcha(l.ctx, &userservice.LoginByCaptchaRequest{
		Email:    req.Email,
		Captcha:  req.Captcha,
		ClientIp: l.ctx.Value("client_ip").(string),
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("登录失败!"), "login by captcha failed! %v", err)
	}

	resp = &types.LoginResponse{
		Data: types.Token{
			AccessToken:  captcha.AccessToken,
			RefreshToken: captcha.RefreshToken,
		},
	}
	return resp, nil
}
