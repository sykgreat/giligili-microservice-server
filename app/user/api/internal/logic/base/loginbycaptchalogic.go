package base

import (
	"context"

	"giligili/app/user/api/internal/svc"
	"giligili/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginByCaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginByCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginByCaptchaLogic {
	return &LoginByCaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginByCaptchaLogic) LoginByCaptcha(req *types.LoginByCaptchaRequest) (resp *types.LoginResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
