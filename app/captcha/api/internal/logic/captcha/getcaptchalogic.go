package captcha

import (
	"context"
	"giligili/app/captcha/rpc/captchaservice"
	"giligili/common"
	"giligili/common/xerr"
	"github.com/pkg/errors"

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
	// 验证邮箱格式
	if !common.VerifyEmailFormat(req.Email) {
		return nil, errors.New("邮箱格式不正确")
	}

	// 调用rpc服务 获取验证码
	email, err := l.svcCtx.CaptchaRpc.GetCaptchaByEmail(l.ctx, &captchaservice.GetCaptchaByEmailReq{
		Email:       req.Email,
		CaptchaType: req.CaptchaType,
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("验证码已发送，请勿重复发送"), "captcha has been sent!")
	}
	return &types.GetCaptchaResponse{
		Code: email.Result,
	}, nil
}
