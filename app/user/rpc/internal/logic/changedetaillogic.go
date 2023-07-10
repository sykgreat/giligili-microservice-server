package logic

import (
	"context"
	"giligili/common/times"
	"giligili/common/xerr"
	"github.com/pkg/errors"
	"time"

	"giligili/app/user/rpc/internal/svc"
	"giligili/app/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangeDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewChangeDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangeDetailLogic {
	return &ChangeDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ChangeDetailLogic) ChangeDetail(in *pb.ChangeDetailRequest) (*pb.Response, error) {
	// 查询用户信息
	one, err := l.svcCtx.UserModel.FindOne(l.ctx, l.ctx.Value("userId").(int64))
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("查询失败！"), "change detail failed! %v", err)
	}

	// 修改用户名
	if one.Username != in.Username {
		one.Username = in.Username
	}

	// 修改头像
	if one.Avatar != in.Avatar {
		one.Avatar = in.Avatar
	}

	// 修改空间背景
	if one.SpaceCover != in.SpaceCover {
		one.SpaceCover = in.SpaceCover
	}

	// 修改性别
	if one.Gender != in.Gender {
		one.Gender = in.Gender
	}

	// 修改生日
	birthday := times.UnixMilliToTime(in.Birthday)
	if one.Birthday != birthday {
		one.Birthday = birthday
	}

	// 修改用户信息
	one.UpdatedTime = time.Now()
	err = l.svcCtx.UserModel.Update(l.ctx, one)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("修改用户信息失败！"), "change user detail failed! %v", err)
	}
	return &pb.Response{}, nil
}
