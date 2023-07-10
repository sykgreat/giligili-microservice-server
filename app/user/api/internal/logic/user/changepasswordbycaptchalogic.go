package user

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

type ChangePasswordByCaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChangePasswordByCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangePasswordByCaptchaLogic {
	return &ChangePasswordByCaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChangePasswordByCaptchaLogic) ChangePasswordByCaptcha(req *types.ChangePasswordByCaptchaRequest) (resp *types.BaseResponse, err error) {
	// 校验邮箱格式
	if format := common.VerifyEmailFormat(req.Email); !format {
		return nil, errors.Wrapf(xerr.NewErrMsg("邮箱格式错误"), "email format is incorrect!")
	}

	_, err = l.svcCtx.UserRpc.ChangePasswordByCaptcha(l.ctx, &userservice.ChangePasswordByCaptchaRequest{
		Email:       req.Email,
		Captcha:     req.Captcha,
		NewPassword: req.Password,
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("修改密码失败"), "change password faild! %v", err)
	}
	return
}
