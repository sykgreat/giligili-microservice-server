package base

import (
	"context"
	"giligili/app/user/rpc/userservice"
	"giligili/common"
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

func (l *RegisterLogic) Register(req *types.RegisterRequest) (resp *types.BaseResponse, err error) {
	if format := common.VerifyEmailFormat(req.Email); !format {
		resp = &types.BaseResponse{
			Code: -1,
			Msg:  "邮箱格式不正确",
		}
		return
	}

	_, err = l.svcCtx.UserRpc.Register(
		l.ctx,
		&userservice.RegisterRequest{
			Email:    req.Email,
			Password: req.Password,
			Captcha:  req.Captcha,
		},
	)
	if err != nil {
		err = errors.Wrapf(xerr.NewErrMsg("注册失败, 请稍后再试"), "register err: %s", err)
		return
	}

	resp = &types.BaseResponse{
		Code: 200,
		Msg:  "注册成功",
	}
	return
}
