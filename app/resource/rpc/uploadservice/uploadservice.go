// Code generated by goctl. DO NOT EDIT.
// Source: resource.proto

package uploadservice

import (
	"context"

	"giligili/app/resource/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	BaseResponse               = pb.BaseResponse
	ChangeResourceTitleRequest = pb.ChangeResourceTitleRequest
	DeleteResourceRequest      = pb.DeleteResourceRequest
	UploadImgRequest           = pb.UploadImgRequest
	UploadImgResponse          = pb.UploadImgResponse
	UploadVideoRequest         = pb.UploadVideoRequest

	UploadService interface {
		UploadVideo(ctx context.Context, opts ...grpc.CallOption) (pb.UploadService_UploadVideoClient, error)
		UploadImg(ctx context.Context, opts ...grpc.CallOption) (pb.UploadService_UploadImgClient, error)
		ChangeResourceTitle(ctx context.Context, in *ChangeResourceTitleRequest, opts ...grpc.CallOption) (*BaseResponse, error)
		DeleteResource(ctx context.Context, in *DeleteResourceRequest, opts ...grpc.CallOption) (*BaseResponse, error)
	}

	defaultUploadService struct {
		cli zrpc.Client
	}
)

func NewUploadService(cli zrpc.Client) UploadService {
	return &defaultUploadService{
		cli: cli,
	}
}

func (m *defaultUploadService) UploadVideo(ctx context.Context, opts ...grpc.CallOption) (pb.UploadService_UploadVideoClient, error) {
	client := pb.NewUploadServiceClient(m.cli.Conn())
	return client.UploadVideo(ctx, opts...)
}

func (m *defaultUploadService) UploadImg(ctx context.Context, opts ...grpc.CallOption) (pb.UploadService_UploadImgClient, error) {
	client := pb.NewUploadServiceClient(m.cli.Conn())
	return client.UploadImg(ctx, opts...)
}

func (m *defaultUploadService) ChangeResourceTitle(ctx context.Context, in *ChangeResourceTitleRequest, opts ...grpc.CallOption) (*BaseResponse, error) {
	client := pb.NewUploadServiceClient(m.cli.Conn())
	return client.ChangeResourceTitle(ctx, in, opts...)
}

func (m *defaultUploadService) DeleteResource(ctx context.Context, in *DeleteResourceRequest, opts ...grpc.CallOption) (*BaseResponse, error) {
	client := pb.NewUploadServiceClient(m.cli.Conn())
	return client.DeleteResource(ctx, in, opts...)
}