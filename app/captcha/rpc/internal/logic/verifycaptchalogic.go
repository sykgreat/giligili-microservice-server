package logic

import (
	"context"
	"giligili/common/enum"

	"giligili/app/captcha/rpc/internal/svc"
	"giligili/app/captcha/rpc/pb"
	"giligili/common/xerr"

	"github.com/pkg/errors"
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
	// 查看是否存在验证码
	ctx, err := l.svcCtx.Redis.GetCtx(l.ctx, enum.CaptchaModule+enum.Captcha+in.CaptchaType+in.Email)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("redis获取失败!"), "redis get failed! %s", err)
	} else if ctx != in.Captcha { // 验证码不正确
		return nil, errors.Wrapf(xerr.NewErrMsg("验证码输入错误，请重新输入!"), "captcha is incorrect!")
	}

	// 删除验证码
	del, err := l.svcCtx.Redis.DelCtx(l.ctx, enum.CaptchaModule+enum.Captcha+in.CaptchaType+in.Email)
	if err != nil || del == 0 {
		return nil, errors.Wrapf(xerr.NewErrMsg("redis删除失败!"), "redis del failed! %s", err)
	}

	return &pb.VerifyCaptchaResp{}, nil
}
