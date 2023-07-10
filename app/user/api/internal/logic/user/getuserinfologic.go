package user

import (
	"context"
	"giligili/app/user/rpc/userservice"
	"giligili/common/xerr"
	"github.com/pkg/errors"
	"time"

	"giligili/app/user/api/internal/svc"
	"giligili/app/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo(req *types.BaseRequest) (resp *types.GetUserInfoResponse, err error) {
	// 获取用户id
	userId := l.ctx.Value("userId").(int64)

	// 调用获取用户信息rpc
	userInfo, err := l.svcCtx.UserRpc.GetDetail(l.ctx, &userservice.GetDetailRequest{
		UserId: userId,
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("获取用户信息失败"), "get user info err: %v", err)
	}
	resp = &types.GetUserInfoResponse{
		User: types.User{
			Email:      userInfo.Email,
			Username:   userInfo.Username,
			Avatar:     userInfo.Avatar,
			SpaceCover: userInfo.SpaceCover,
			Sign:       userInfo.Sign,
			Birthday:   time.Unix(userInfo.Birthday, 0).Format("2006-01-02 15:04:05"),
			Gender:     userInfo.Gender,
			Id:         userInfo.UserId,
			CreateTime: time.Unix(userInfo.CreatedAt, 0).Format("2006-01-02 15:04:05"),
			UpdateTime: time.Unix(userInfo.UpdatedAt, 0).Format("2006-01-02 15:04:05"),
		},
	}
	return
}
