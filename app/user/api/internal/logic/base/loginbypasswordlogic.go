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
	// 验证邮箱格式
	if format := common.VerifyEmailFormat(req.Email); !format {
		return nil, errors.Wrapf(xerr.NewErrMsg("邮箱格式不正确!"), "email format is incorrect!")
	}

	// 调用密码登陆rpc服务
	password, err := l.svcCtx.UserRpc.LoginByPassword(l.ctx, &userservice.LoginByPasswordRequest{
		Email:    req.Email,
		Password: req.Password,
		ClientIp: l.ctx.Value("client_ip").(string),
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("登录失败!"), "login by password failed! %v", err)
	}

	resp = &types.LoginResponse{
		Data: types.Token{
			AccessToken:  password.AccessToken,
			RefreshToken: password.RefreshToken,
		},
	}
	return resp, nil
}
