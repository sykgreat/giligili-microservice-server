package logic

import (
	"context"
	"encoding/json"
	"giligili/common/xerr"
	"strconv"

	"github.com/pkg/errors"

	"giligili/app/user/rpc/internal/svc"
	"giligili/app/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDetailLogic {
	return &GetDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetDetailLogic) GetDetail(in *pb.GetDetailRequest) (res *pb.GetDetailResponse, err error) {
	// 从Redis中获取用户信息
	ctx, err := l.svcCtx.Redis.GetCtx(l.ctx, l.svcCtx.Config.Redis.Key+":info:"+strconv.Itoa(int(in.UserId)))
	if err != nil { // 获取失败
		// 从数据库中获取用户信息
		one, err := l.svcCtx.UserModel.FindOne(l.ctx, in.UserId)
		if err != nil {
			return nil, errors.Wrapf(
				xerr.NewErrMsg("获取用户信息失败!"),
				"系统异常, err: %s", err,
			)
		}

		// 序列化用户信息
		marshal, err := json.Marshal(one)
		if err != nil {
			return nil, errors.Wrapf(
				xerr.NewErrMsg("用户信息序列化失败"),
				"user marshal, err: %s", err,
			)
		}

		// 将用户信息存入redis
		if err := l.svcCtx.Redis.SetCtx(
			l.ctx,
			l.svcCtx.Config.Redis.Key+":info:"+strconv.Itoa(int(one.Id)),
			string(marshal),
		); err != nil {
			return nil, errors.Wrapf(
				xerr.NewErrMsg("redis存用户信息失败"),
				"redis set user info err: %s, ", err,
			)
		}

		// 反序列化用户信息
		if err := json.Unmarshal(marshal, res); err != nil {
			return nil, errors.Wrapf(
				xerr.NewErrMsg("反序列化用户信息失败"),
				"json unmarshal, err: %s", err,
			)
		}
	}

	// 反序列化用户信息
	if err := json.Unmarshal([]byte(ctx), res); err != nil {
		return nil, errors.Wrapf(
			xerr.NewErrMsg("反序列化用户信息失败"),
			"json unmarshal, err: %s", err,
		)
	}

	return res, nil
}
