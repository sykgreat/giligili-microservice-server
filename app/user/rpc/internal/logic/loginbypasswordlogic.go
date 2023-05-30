package logic

import (
	"context"

	"giligili/app/user/rpc/internal/svc"
	"giligili/app/user/rpc/pb"
	"giligili/app/user/utils/password"
	"giligili/common/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginByPasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginByPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginByPasswordLogic {
	return &LoginByPasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginByPasswordLogic) LoginByPassword(in *pb.LoginByPasswordRequest) (*pb.LoginResponse, error) {
	user, err := l.svcCtx.UserModel.FindOneByEmail(l.ctx, in.Email)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "根据邮箱查询用户信息失败, email:%s,err:%v", in.Email, err)
	}
	if user == nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("用户不存在"), "email:%s", in.Email)
	}
	if err = password.ComparePassword(user.Password.String, in.Password); err != nil {
		return nil, errors.Wrap(xerr.NewErrMsg("账号或密码不正确"), "秘密输入错误, 请重新输入!")
	}

	generateTokenLogic := NewGenerateTokenLogic(l.ctx, l.svcCtx)
	token, err := generateTokenLogic.GenerateToken(&pb.GenerateTokenReq{
		UserId: user.Id,
		Email:  user.Email,
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("生成token失败"), "GenerateToken userId : %d", user.Id)
	}
	return &pb.LoginResponse{
		AccessToken:  token.AccessToken,
		AccessExpire: token.AccessExpire,
		RefreshAfter: token.RefreshAfter,
	}, nil
}
