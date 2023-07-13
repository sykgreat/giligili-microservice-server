package base

import (
	"context"
	"giligili/app/user/rpc/userservice"
	"giligili/common/Verify"
	"giligili/common/xerr"
	"github.com/pkg/errors"

	"giligili/app/user/api/internal/svc"
	"giligili/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterRequest) (resp *types.BaseRequest, err error) {
	// 验证邮箱格式
	if format := Verify.VerifyEmailFormat(req.Email); !format {
		return nil, errors.Wrapf(xerr.NewErrMsg("邮箱格式不正确!"), "email format is incorrect!")
	}

	// 调用用户注册rpc
	_, err = l.svcCtx.UserRpc.Register(l.ctx, &userservice.RegisterRequest{
		Email:    req.Email,
		Password: req.Password,
		Captcha:  req.Captcha,
		ClientIp: l.ctx.Value("client_ip").(string),
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("用户注册失败!"), "user register failed! %v", err)
	}
	return
}
