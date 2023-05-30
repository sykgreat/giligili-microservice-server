package logic

import (
	"context"

	"giligili/app/user/rpc/internal/svc"
	"giligili/app/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginByCaptchaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginByCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginByCaptchaLogic {
	return &LoginByCaptchaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginByCaptchaLogic) LoginByCaptcha(in *pb.LoginByCaptchaRequest) (*pb.LoginResponse, error) {

	return &pb.LoginResponse{}, nil
}
