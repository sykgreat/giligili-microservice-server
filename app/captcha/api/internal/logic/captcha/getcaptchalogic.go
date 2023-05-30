package captcha

import (
	"context"
	"errors"
	"giligili/app/captcha/rpc/captchaservice"
	"giligili/common"

	"giligili/app/captcha/api/internal/svc"
	"giligili/app/captcha/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCaptchaLogic {
	return &GetCaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCaptchaLogic) GetCaptcha(req *types.GetCaptchaRequest) (resp *types.GetCaptchaResponse, err error) {
	if !common.VerifyEmailFormat(req.Email) {
		return nil, errors.New("邮箱格式不正确")
	}
	email, err := l.svcCtx.CaptchaRpc.GetCaptchaByEmail(l.ctx, &captchaservice.GetCaptchaByEmailReq{
		Email: req.Email,
	})
	if err != nil {
		logx.Error("获取验证码失败: ", err)
		return nil, err
	}
	return &types.GetCaptchaResponse{
		Code: email.Result,
		Msg:  "验证码已发送至邮箱，请注意查收！",
	}, nil
}
