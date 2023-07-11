package user

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

type ChangePasswordByPasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChangePasswordByPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangePasswordByPasswordLogic {
	return &ChangePasswordByPasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChangePasswordByPasswordLogic) ChangePasswordByPassword(req *types.ChangePasswordByPasswordRequest) (resp *types.BaseResponse, err error) {
	// 校验邮箱格式
	if format := common.VerifyEmailFormat(req.Email); !format {
		return nil, errors.Wrapf(xerr.NewErrMsg("邮箱格式错误"), "email format is incorrect!")
	}

	// 校验新旧密码是否一致
	if req.OldPassword == req.NewPassword {
		return nil, errors.Wrapf(xerr.NewErrMsg("邮箱格式错误"), "The old and new passwords cannot be the same!")
	}

	// 调用修改密码的rpc
	_, err = l.svcCtx.UserRpc.ChangePasswordByPassword(l.ctx, &userservice.ChangePasswordByPasswordRequest{
		Email:       req.Email,
		OrdPassword: req.OldPassword,
		NewPassword: req.NewPassword,
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("修改密码失败"), "change password faild! %v", err)
	}
	return
}
