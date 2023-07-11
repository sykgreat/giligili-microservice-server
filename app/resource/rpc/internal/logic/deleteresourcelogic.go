package logic

import (
	"context"
	"giligili/common/xerr"
	"github.com/pkg/errors"
	"os"
	"strconv"

	"giligili/app/resource/rpc/internal/svc"
	"giligili/app/resource/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteResourceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteResourceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteResourceLogic {
	return &DeleteResourceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteResourceLogic) DeleteResource(in *pb.DeleteResourceRequest) (*pb.BaseResponse, error) {
	// 查询视频是否存在
	video, err := l.svcCtx.ResourceModel.FindOneByVideoId(l.ctx, in.VideoId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("找不到该视频"), "delete resource failed: %v", err)
	}

	// 删除视频
	err = l.svcCtx.ResourceModel.Delete(l.ctx, video.Id)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("删除视频失败"), "delete resource failed: %v", err)
	}

	err = os.RemoveAll(l.svcCtx.Config.Upload.VideoPath + "/" + strconv.FormatInt(video.Uid.Int64, 10) + "/" + strconv.FormatInt(video.Id, 10))
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("删除视频失败"), "delete resource failed: %v", err)
	}
	return &pb.BaseResponse{}, nil
}
