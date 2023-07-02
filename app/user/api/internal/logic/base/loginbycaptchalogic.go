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
	if format := common.VerifyEmailFormat(req.Email); !format {
		resp = &types.LoginResponse{
			Code: -1,
			Msg:  "邮箱格式不正确",
		}
		return
	}

	captcha, err := l.svcCtx.UserRpc.LoginByCaptcha(
		l.ctx,
		&userservice.LoginByCaptchaRequest{
			Email:    req.Email,
			Captcha:  req.Captcha,
			ClientIp: l.ctx.Value("X-Real-IP").(string),
		},
	)
	if err != nil {
		err = errors.Wrapf(xerr.NewErrMsg("登录失败, 账号或验证码输入错误"), "login by captcha err: %s", err)
		return
	}
	resp = &types.LoginResponse{
		Code: 200,
		Msg:  "登录成功",
		Data: types.Token{
			AccessToken:  captcha.AccessToken,
			AccessExpire: captcha.AccessExpire,
			RefreshAfter: captcha.RefreshAfter,
		},
	}
	return
}
