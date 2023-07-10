package logic

import (
	"context"
	"giligili/common/xerr"
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

func (l *GetDetailLogic) GetDetail(in *pb.GetDetailRequest) (*pb.GetDetailResponse, error) {
	one, err := l.svcCtx.UserModel.FindOne(l.ctx, in.UserId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("用户不存在"), "get user detail err: %v", err)
	}
	gender := "未知"
	switch one.Gender {
	case 1:
		gender = "男"
	case 2:
		gender = "女"
	default:
	}
	return &pb.GetDetailResponse{
		Email:      one.Email,
		Username:   one.Username,
		Avatar:     one.Avatar,
		SpaceCover: one.SpaceCover,
		Sign:       one.Sign,
		Birthday:   one.Birthday.Unix(),
		Gender:     gender,
		UserId:     one.Id,
		CreatedAt:  one.CreatedTime.Unix(),
		UpdatedAt:  one.UpdatedTime.Unix(),
		ClientIp:   one.ClientIp,
	}, nil
}
