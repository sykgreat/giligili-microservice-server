package user

import (
	"context"

	"giligili/app/user/api/internal/svc"
	"giligili/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangePasswordByPasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChangePasswordByPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangePasswordByPasswordLogic {
	return &ChangePasswordByPasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChangePasswordByPasswordLogic) ChangePasswordByPassword(req *types.ChangePasswordByPasswordRequest) (resp *types.BaseResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
