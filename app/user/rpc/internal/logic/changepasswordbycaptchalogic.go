package logic

import (
	"context"
	"giligili/app/captcha/rpc/captchaservice"
	"giligili/common/enum"
	"giligili/common/password"
	xerr "giligili/common/xerr"
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
	// 验证验证码 是否正确
	_, err := l.svcCtx.CaptchaRpc.VerifyCaptcha(l.ctx, &captchaservice.VerifyCaptchaReq{
		Email:       in.Email,
		Captcha:     in.Captcha,
		CaptchaType: enum.CaptchaResetPassword,
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("验证码输入错误！"), "change password failed! %v", err)
	}

	// 查询用户是否存在
	user, err := l.svcCtx.UserModel.FindOneByEmail(l.ctx, in.Email)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("该用户不存在，请检查邮箱是否输入错误！"), "change password failed! %v", err)
	}

	// 将密码进行加密
	generatePassword, err := password.GeneratePassword(in.NewPassword)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("密码加密失败"), "change password failed! %v", err)
	}

	// 修改密码
	user.Password = generatePassword
	if err = l.svcCtx.UserModel.Update(l.ctx, user); err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("修改密码失败，请重新尝试！"), "change password failed! %v", err)
	}

	return &pb.Response{}, nil
}
