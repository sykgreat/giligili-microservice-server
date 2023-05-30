package logic

import (
	"context"

	"giligili/app/user/rpc/internal/svc"
	"giligili/app/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangePasswordByCaptchaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewChangePasswordByCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangePasswordByCaptchaLogic {
	return &ChangePasswordByCaptchaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ChangePasswordByCaptchaLogic) ChangePasswordByCaptcha(in *pb.ChangePasswordByCaptchaRequest) (*pb.Response, error) {
	// todo: add your logic here and delete this line

	return &pb.Response{}, nil
}
