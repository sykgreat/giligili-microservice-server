package logic

import (
	"context"

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
	// todo: add your logic here and delete this line

	return &pb.Response{}, nil
}
