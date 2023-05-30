package logic

import (
	"context"
	"errors"
	"giligili/app/captcha/utils/captcha"
	"giligili/app/email/rpc/emailservice"

	"giligili/app/captcha/rpc/internal/svc"
	"giligili/app/captcha/rpc/pb"

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
	if ctx, err := l.svcCtx.Redis.GetCtx(l.ctx, l.svcCtx.Config.Redis.Key+":"+in.Email); err != nil {
		logx.Error("redis get error: ", err)
		return nil, err
	} else if len(ctx) != 0 {
		return nil, errors.New("验证码已发送，请勿重复发送")
	}

	// 生成验证码，并存入redis
	generate, i := captcha.Captcha.Generate()
	if err := l.svcCtx.Redis.SetexCtx(l.ctx, l.svcCtx.Config.Redis.Key+":"+in.Email, generate, i); err != nil {
		logx.Error("redis set error: ", err)
		return nil, err
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
			logx.Error("redis del error: ", err)
			return nil, err
		} else if ctx == 0 {
			return nil, errors.New("redis删除失败")
		}

		logx.Error("发送邮件失败: ", err)
		return nil, err
	}

	return &pb.GetCaptchaByEmailResp{
		Result: email.Result,
	}, nil
}
