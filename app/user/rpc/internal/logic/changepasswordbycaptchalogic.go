package logic

import (
	"context"
	"giligili/app/captcha/rpc/captchaservice"
	"giligili/app/user/utils/password"
	"giligili/common/xerr"
	"github.com/pkg/errors"

	"giligili/app/user/rpc/internal/svc"
	"giligili/app/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangePasswordByCaptchaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewChangePasswordByCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangePasswordByCaptchaLogic {
	return &ChangePasswordByCaptchaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ChangePasswordByCaptchaLogic) ChangePasswordByCaptcha(in *pb.ChangePasswordByCaptchaRequest) (*pb.Response, error) {
	// 通过captcha rpc服务验证验证码是否正确
	captcha, err := l.svcCtx.CaptchaRpc.VerifyCaptcha(
		l.ctx,
		&captchaservice.VerifyCaptchaReq{
			Email:   in.Email,
			Captcha: in.Captcha,
		},
	)
	if err != nil || captcha.Result != 200 {
		return nil, errors.Wrapf(
			xerr.NewErrMsg("验证码错误, 请重新输入！"),
			"verify captcha err: %s", err,
		)
	}

	// 查询用户是否存在
	email, err := l.svcCtx.UserModel.FindOneByEmail(l.ctx, in.Email)
	if err != nil {
		return nil, errors.Wrapf(
			xerr.NewErrMsg("根据邮箱查询用户失败, 请重新修改！"),
			"find user by email err: %s", err,
		)
	}

	// 生成密码
	generatePassword, err := password.GeneratePassword(in.NewPassword)
	if err != nil {
		return nil, errors.Wrapf(
			xerr.NewErrMsg("密码加密失败, 请重新修改！"),
			"generate password err: %s", err,
		)
	}

	// 修改密码
	email.Password = generatePassword

	// 更新用户
	if err = l.svcCtx.UserModel.Update(l.ctx, email); err != nil {
		return nil, errors.Wrapf(
			xerr.NewErrMsg("修改密码失败, 请重新修改！"),
			"update user err: %s", err,
		)
	}

	return &pb.Response{
		Result: 200,
	}, nil
}
