package logic

import (
	"context"
	"errors"
	"giligili/app/email/utils/email"
	"giligili/common"

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
	if !common.VerifyEmailFormat(in.Email) {
		return nil, errors.New("邮箱格式不正确")
	}
	err := email.Email.Send(in.Email, in.Subject, in.Body)
	if err != nil {
		logx.Error("发送邮件失败: ", err)
		return nil, err
	}
	return &pb.SendEmailResponse{
		Result: 200,
	}, nil
}
