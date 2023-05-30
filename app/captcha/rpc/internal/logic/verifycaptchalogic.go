package logic

import (
	"context"
	"errors"

	"giligili/app/captcha/rpc/internal/svc"
	"giligili/app/captcha/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyCaptchaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewVerifyCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyCaptchaLogic {
	return &VerifyCaptchaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *VerifyCaptchaLogic) VerifyCaptcha(in *pb.VerifyCaptchaReq) (*pb.VerifyCaptchaResp, error) {
	if ctx, err := l.svcCtx.Redis.GetCtx(l.ctx, l.svcCtx.Config.Redis.Key+":"+in.Email); err != nil {
		logx.Error("redis get error: ", err)
		return nil, err
	} else if ctx != in.Captcha {
		return nil, errors.New("验证码输入错误，请重新输入！")
	}

	if delCtx, err := l.svcCtx.Redis.DelCtx(l.ctx, l.svcCtx.Config.Redis.Key+":"+in.Email); err != nil {
		logx.Error("redis del error: ", err)
		return nil, err
	} else if delCtx == 0 {
		return nil, errors.New("验证码已过期，请重新获取！")
	}

	return &pb.VerifyCaptchaResp{
		Result: 200,
	}, nil
}
