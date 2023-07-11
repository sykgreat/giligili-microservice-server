package user

import (
	"context"
	"giligili/app/user/rpc/userservice"
	"giligili/common/xerr"
	"github.com/pkg/errors"

	"giligili/app/user/api/internal/svc"
	"giligili/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogoutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LogoutLogic) Logout(req *types.BaseRequest) (resp *types.BaseResponse, err error) {
	// 调用用户登出rpc
	_, err = l.svcCtx.UserRpc.Logout(l.ctx, &userservice.LogoutRequest{
		UserId: l.ctx.Value(`userId`).(int64),
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("用户登出失败！"), "user logout failed! %v", err)
	}
	return
}
