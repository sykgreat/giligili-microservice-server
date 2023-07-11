package logic

import (
	"context"

	"giligili/app/resource/rpc/internal/svc"
	"giligili/app/resource/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUploadVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadVideoLogic {
	return &UploadVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UploadVideoLogic) UploadVideo(stream pb.UploadService_UploadVideoServer) error {
	// todo: add your logic here and delete this line

	return nil
}
