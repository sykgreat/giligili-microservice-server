package logic

import (
	"context"
	"database/sql"

	"giligili/app/captcha/rpc/captchaservice"
	"giligili/app/user/model"
	"giligili/app/user/rpc/internal/svc"
	"giligili/app/user/rpc/pb"
	"giligili/app/user/utils/password"
	"giligili/common/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *pb.RegisterRequest) (*pb.Response, error) {
	// 使用验证码rpc验证验证码
	if verifyResult, err := l.svcCtx.CaptchaRpc.VerifyCaptcha(l.ctx, &captchaservice.VerifyCaptchaReq{
		Email:   in.Email,
		Captcha: in.Captcha,
	}); err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("验证码验证失败"), "VerifyCaptcha err:%v", err)
	} else if verifyResult.Result != 200 {
		return nil, errors.Wrapf(xerr.NewErrMsg("验证码验证失败"), "验证码输入错误, 请重新输入!")
	}

	s, err := password.GeneratePassword(in.Password)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("密码加密失败"), "GeneratePassword err:%v", err)
	}

	result, err := l.svcCtx.UserModel.Insert(
		l.ctx,
		&model.User{
			Id:    l.svcCtx.Snowflakes.NextVal(),
			Email: in.Email,
			Password: sql.NullString{
				String: s,
				Valid:  true,
			},
		},
	)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("用户注册失败"), "Insert err:%v", err)
	}
	result.LastInsertId()

	return &pb.Response{
		Result: 200,
	}, nil
}
