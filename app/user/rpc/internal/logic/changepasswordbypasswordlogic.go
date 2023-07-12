package logic

import (
	"context"
	"giligili/common/password"
	"giligili/common/xerr"
	"github.com/pkg/errors"
	"time"

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
	user, err := l.svcCtx.UserModel.FindOneByEmail(l.ctx, in.Email)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("该用户不存在，请检查邮箱是否输入错误！"), "change password failed! %v", err)
	}

	// 查看原密码是否输入正确
	err = password.ComparePassword(user.Password, in.OrdPassword)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("原密码输入错误，请重新输入！"), "change password failed! %v", err)
	}

	generatePassword, err := password.GeneratePassword(in.NewPassword)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("密码生成失败"), "change password failed! %v", err)
	}

	// 修改密码
	user.Password = generatePassword
	user.UpdateTime = time.Now()
	if _, err = l.svcCtx.UserModel.Update(l.ctx, nil, user); err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("修改密码失败，请重新尝试！"), "change password failed! %v", err)
	}
	return &pb.Response{}, nil
}
