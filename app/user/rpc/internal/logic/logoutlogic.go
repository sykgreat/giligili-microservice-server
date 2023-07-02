package logic

import (
	"context"
	"giligili/common/xerr"
	"strconv"

	"giligili/app/user/rpc/internal/svc"
	"giligili/app/user/rpc/pb"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type LogoutLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LogoutLogic) Logout(in *pb.LogoutRequest) (out *pb.Response, err error) {
	if _, err = l.svcCtx.Redis.DelCtx(
		l.ctx,
		l.svcCtx.Config.Redis.Key+":token:"+strconv.Itoa(int(in.UserId)),
	); err != nil {
		return nil, errors.Wrapf(
			xerr.NewErrMsg("用户登出失败"),
			"redis del err: %s", err,
		)
	}

	if _, err = l.svcCtx.Redis.DelCtx(
		l.ctx,
		l.svcCtx.Config.Redis.Key+":info:"+strconv.Itoa(int(in.UserId))); err != nil {
		return nil, errors.Wrapf(
			xerr.NewErrMsg("用户登出失败"),
			"redis del err: %s", err,
		)
	}
	return &pb.Response{
		Result: 200,
	}, nil
}
