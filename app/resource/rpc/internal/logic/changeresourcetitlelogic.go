package logic

import (
	"context"
	"database/sql"
	"giligili/common/xerr"
	"github.com/pkg/errors"

	"giligili/app/resource/rpc/internal/svc"
	"giligili/app/resource/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangeResourceTitleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewChangeResourceTitleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangeResourceTitleLogic {
	return &ChangeResourceTitleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ChangeResourceTitleLogic) ChangeResourceTitle(in *pb.ChangeResourceTitleRequest) (*pb.BaseResponse, error) {
	// 查询视频是否存在
	video, err := l.svcCtx.ResourceModel.FindOneByVideoId(l.ctx, in.VideoId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("找不到该视频"), "change resource title failed: %v", err)
	}

	// 修改视频标题
	video.Title = sql.NullString{
		String: in.Title,
		Valid:  true,
	}
	err = l.svcCtx.ResourceModel.Update(l.ctx, video)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("修改视频标题失败"), "change resource title failed: %v", err)
	}
	return &pb.BaseResponse{}, nil
}
