package user

import (
	"context"

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
	// todo: add your logic here and delete this line

	return
}
