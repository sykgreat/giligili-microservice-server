package logic

import (
	"context"
	"giligili/app/user/utils/password"
	"giligili/common/xerr"
	"github.com/pkg/errors"

	"giligili/app/user/rpc/internal/svc"
	"giligili/app/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangePasswordByPasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewChangePasswordByPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangePasswordByPasswordLogic {
	return &ChangePasswordByPasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ChangePasswordByPasswordLogic) ChangePasswordByPassword(in *pb.ChangePasswordByPasswordRequest) (*pb.Response, error) {
	// 查询用户是否存在
	email, err := l.svcCtx.UserModel.FindOneByEmail(l.ctx, in.Email)
	if err != nil {
		return nil, errors.Wrapf(
			xerr.NewErrMsg("根据邮箱查询用户失败, 请重新修改！"),
			"find user by email err: %s", err,
		)
	}

	// 验证原密码是否正确
	if err = password.ComparePassword(in.OrdPassword, email.Password); err != nil {
		return nil, errors.Wrapf(
			xerr.NewErrMsg("原密码错误, 请重新输入！"),
			"compare password err: %s", err,
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
