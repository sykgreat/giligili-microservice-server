package logic

import (
	"context"
	"giligili/common/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"strconv"

	"giligili/app/resource/rpc/internal/svc"
	"giligili/app/resource/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadImgLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUploadImgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadImgLogic {
	return &UploadImgLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UploadImgLogic) UploadImg(stream pb.UploadService_UploadImgServer) error {
	name := strconv.FormatInt(l.svcCtx.Snowflake.NextVal(), 10) + ".jpg"
	file := storage.NewFile(name)

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			if err := l.svcCtx.UploadImg.Store(file); err != nil {
				return status.Error(codes.Internal, err.Error())
			}
			return stream.SendAndClose(&pb.UploadImgResponse{Url: name})
		}
		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}

		if err := file.Write(req.Content); err != nil {
			return status.Error(codes.Internal, err.Error())
		}
	}
}
