package logic

import (
	"context"
	"encoding/json"
	"giligili/common/times"
	"giligili/common/xerr"
	"strconv"
	"time"

	"github.com/pkg/errors"

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
	// 查询用户是否存在
	one, err := l.svcCtx.UserModel.FindOne(l.ctx, in.UserId)
	if err != nil {
		return nil, errors.Wrapf(
			xerr.NewErrMsg("根据id查询用户失败, 请重新修改！"),
			"find user by id err: %s", err,
		)
	}

	// 查看用户名是否需要修改
	if in.Username != "" {
		// 查询用户名是否已被人使用
		user, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, in.Username)
		if err == nil || user != nil {
			return nil, errors.Wrapf(
				xerr.NewErrMsg("该用户名已被人使用, 请重新输入！"),
				"find user by username err: %s", err,
			)
		}

		one.Username = in.Username
	}

	// 查看邮箱是否需要修改
	if in.Avatar != "" {
		one.Avatar = in.Avatar
	}
	// 查看生日是否需要修改
	if in.Birthday != 0 {
		one.Birthday = times.UnixNanoToTime(in.Birthday)
	}
	// 查看空间背景图是否需要修改
	if in.SpaceCover != "" {
		one.SpaceCover = in.SpaceCover
	}
	// 查看个性签名是否需要修改
	if in.Sign != "" {
		one.Sign = in.Sign
	}
	// 查看性别是否需要修改
	if in.Gender != 0 {
		one.Gender = in.Gender
	}
	// 修改时间
	one.UpdatedTime = time.Now()

	// 修改用户信息
	if err = l.svcCtx.UserModel.Update(l.ctx, one); err != nil {
		return nil, errors.Wrapf(
			xerr.NewErrMsg("修改用户信息失败, 请重新修改！"),
			"update user err: %s", err,
		)
	}

	// 删除缓存中的用户信息
	ctx, err := l.svcCtx.Redis.DelCtx(
		l.ctx,
		l.svcCtx.Config.Redis.Key+":info:"+strconv.Itoa(int(in.UserId)),
	)
	if err != nil || ctx == 0 {
		return nil, errors.Wrapf(
			xerr.NewErrMsg("删除缓存失败, 请重新修改！"),
			"del redis err: %s", err,
		)
	}

	// 序列化用户信息
	marshal, err := json.Marshal(one)
	if err != nil {
		return nil, errors.Wrapf(
			xerr.NewErrMsg("序列化失败, 请重新修改！"),
			"json marshal err: %s", err,
		)
	}

	// 设置缓存
	if err := l.svcCtx.Redis.SetCtx(
		l.ctx,
		l.svcCtx.Config.Redis.Key+":info:"+strconv.Itoa(int(in.UserId)),
		string(marshal),
	); err != nil {
		return nil, errors.Wrapf(
			xerr.NewErrMsg("设置缓存失败, 请重新修改！"),
			"set redis err: %s", err,
		)
	}

	return &pb.Response{
		Result: 200,
	}, nil
}
