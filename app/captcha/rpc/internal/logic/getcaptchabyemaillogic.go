package logic

import (
	"context"
	"giligili/common/captcha"
	"giligili/common/enum"

	"giligili/app/captcha/rpc/internal/svc"
	"giligili/app/captcha/rpc/pb"
	"giligili/app/email/rpc/emailservice"
	"giligili/common/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetCaptchaByEmailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCaptchaByEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCaptchaByEmailLogic {
	return &GetCaptchaByEmailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCaptchaByEmailLogic) GetCaptchaByEmail(in *pb.GetCaptchaByEmailReq) (result *pb.GetCaptchaByEmailResp, err error) {
	// 判断验证码是否已发送
	if ctx, err := l.svcCtx.Redis.GetCtx(l.ctx, enum.CaptchaModule+":"+enum.Captcha+":"+in.CaptchaType+":"+in.Email); err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("redis获取失败"), "redis get failed! %v", err)
	} else if len(ctx) != 0 {
		return nil, errors.Wrapf(xerr.NewErrMsg("验证码已发送，请勿重复发送"), "captcha has been sent!")
	}

	// 生成验证码，并存入redis
	generate, i := captcha.Captcha.Generate() // 生成验证码
	if err := l.svcCtx.Redis.SetexCtx(l.ctx, enum.CaptchaModule+":"+enum.Captcha+":"+in.CaptchaType+":"+in.Email, generate, i); err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("redis存储失败"), "redis set failed! %v", err)
	}

	// 调用emailRpc 发送邮件
	email, err := l.svcCtx.EmailRpc.SendEmail(
		l.ctx,
		&emailservice.SendEmailRequest{
			Email:   in.Email,
			Subject: "验证码",
			Body:    "<h3>尊敬的用户：</h3><p>您好! 您的验证码是 <span style='color:red'> " + generate + "</span>，五分钟内有效，祝您生活愉快！</p>",
		},
	)
	if err != nil {
		// 发送失败，删除redis中的验证码
		if ctx, err := l.svcCtx.Redis.DelCtx(l.ctx, l.svcCtx.Config.Redis.Key+":"+in.Email); err != nil {
			return nil, errors.Wrapf(xerr.NewErrMsg("redis删除失败"), "redis del failed! %v", err)
		} else if ctx == 0 { // 删除失败，说明redis中没有该验证码
			return nil, errors.Wrapf(xerr.NewErrMsg("没有该验证码"), "no captcha!")
		}

		return nil, errors.Wrapf(xerr.NewErrMsg("发送邮件失败"), "email send failed! %v", err)
	}

	return &pb.GetCaptchaByEmailResp{
		Result: email.Result,
	}, nil
}
