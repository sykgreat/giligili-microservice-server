package logic

import (
	"context"
	"giligili/app/email/utils/email"

	"giligili/app/email/rpc/internal/svc"
	"giligili/app/email/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendEmailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendEmailLogic {
	return &SendEmailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendEmailLogic) SendEmail(in *pb.SendEmailRequest) (*pb.SendEmailResponse, error) {
	// 发送验证码
	err := email.Email.Send(in.Email, in.Subject, in.Body)
	if err != nil {
		logx.Error("发送邮件失败: ", err)
		return nil, err
	}
	return &pb.SendEmailResponse{
		Result: 200,
	}, nil
}
