package user

import (
	"context"

	"giligili/app/user/api/internal/svc"
	"giligili/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangeDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChangeDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangeDetailLogic {
	return &ChangeDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChangeDetailLogic) ChangeDetail(req *types.ChangeDetailRequest) (resp *types.BaseResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
