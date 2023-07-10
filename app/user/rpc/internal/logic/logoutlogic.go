package logic

import (
	"context"
	"giligili/common/enum"
	"giligili/common/xerr"
	"github.com/pkg/errors"
	"strconv"

	"giligili/app/user/rpc/internal/svc"
	"giligili/app/user/rpc/pb"

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

func (l *LogoutLogic) Logout(in *pb.LogoutRequest) (*pb.Response, error) {
	// 删除redis中的用户token
	val, err := l.svcCtx.Redis.DelCtx(l.ctx, enum.UserModule+enum.Token+strconv.Itoa(int(l.ctx.Value(`userId`).(int64)))+":*")
	if err != nil || val == 0 {
		return nil, errors.Wrapf(xerr.NewErrMsg("用户退出登陆失败！"), "user logout failed! %v", err)
	}
	return &pb.Response{}, nil
}
