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

type LoginByPasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginByPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginByPasswordLogic {
	return &LoginByPasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginByPasswordLogic) LoginByPassword(req *types.LoginByPasswordRequest) (resp *types.LoginResponse, err error) {
	if format := common.VerifyEmailFormat(req.Email); !format {
		resp = &types.LoginResponse{
			Code: -1,
			Msg:  "邮箱格式不正确",
		}
		return
	}

	password, err := l.svcCtx.UserRpc.LoginByPassword(
		l.ctx,
		&userservice.LoginByPasswordRequest{
			Email:    req.Email,
			Password: req.Password,
		},
	)
	if err != nil {
		err = errors.Wrapf(xerr.NewErrMsg("登录失败, 账号或密码输入错误"), "login by password err: %s", err)
		return
	}
	resp = &types.LoginResponse{
		Code: 200,
		Msg:  "登录成功",
		Data: types.Token{
			AccessToken:  password.AccessToken,
			AccessExpire: password.AccessExpire,
			RefreshAfter: password.RefreshAfter,
		},
	}
	return
}
