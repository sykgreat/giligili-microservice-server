// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package userservice

import (
	"context"

	"giligili/app/user/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	ChangeDetailRequest             = pb.ChangeDetailRequest
	ChangePasswordByCaptchaRequest  = pb.ChangePasswordByCaptchaRequest
	ChangePasswordByPasswordRequest = pb.ChangePasswordByPasswordRequest
	GenerateTokenReq                = pb.GenerateTokenReq
	GenerateTokenResp               = pb.GenerateTokenResp
	GetDetailRequest                = pb.GetDetailRequest
	GetDetailResponse               = pb.GetDetailResponse
	LoginByCaptchaRequest           = pb.LoginByCaptchaRequest
	LoginByPasswordRequest          = pb.LoginByPasswordRequest
	LoginResponse                   = pb.LoginResponse
	LogoutRequest                   = pb.LogoutRequest
	RegisterRequest                 = pb.RegisterRequest
	Response                        = pb.Response

	UserService interface {
		LoginByPassword(ctx context.Context, in *LoginByPasswordRequest, opts ...grpc.CallOption) (*LoginResponse, error)
		LoginByCaptcha(ctx context.Context, in *LoginByCaptchaRequest, opts ...grpc.CallOption) (*LoginResponse, error)
		Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*Response, error)
		Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*Response, error)
		GetDetail(ctx context.Context, in *GetDetailRequest, opts ...grpc.CallOption) (*GetDetailResponse, error)
		ChangeDetail(ctx context.Context, in *ChangeDetailRequest, opts ...grpc.CallOption) (*Response, error)
		ChangePasswordByCaptcha(ctx context.Context, in *ChangePasswordByCaptchaRequest, opts ...grpc.CallOption) (*Response, error)
		ChangePasswordByPassword(ctx context.Context, in *ChangePasswordByPasswordRequest, opts ...grpc.CallOption) (*Response, error)
	}

	defaultUserService struct {
		cli zrpc.Client
	}
)

func NewUserService(cli zrpc.Client) UserService {
	return &defaultUserService{
		cli: cli,
	}
}

func (m *defaultUserService) LoginByPassword(ctx context.Context, in *LoginByPasswordRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	client := pb.NewUserServiceClient(m.cli.Conn())
	return client.LoginByPassword(ctx, in, opts...)
}

func (m *defaultUserService) LoginByCaptcha(ctx context.Context, in *LoginByCaptchaRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	client := pb.NewUserServiceClient(m.cli.Conn())
	return client.LoginByCaptcha(ctx, in, opts...)
}

func (m *defaultUserService) Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*Response, error) {
	client := pb.NewUserServiceClient(m.cli.Conn())
	return client.Logout(ctx, in, opts...)
}

func (m *defaultUserService) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*Response, error) {
	client := pb.NewUserServiceClient(m.cli.Conn())
	return client.Register(ctx, in, opts...)
}

func (m *defaultUserService) GetDetail(ctx context.Context, in *GetDetailRequest, opts ...grpc.CallOption) (*GetDetailResponse, error) {
	client := pb.NewUserServiceClient(m.cli.Conn())
	return client.GetDetail(ctx, in, opts...)
}

func (m *defaultUserService) ChangeDetail(ctx context.Context, in *ChangeDetailRequest, opts ...grpc.CallOption) (*Response, error) {
	client := pb.NewUserServiceClient(m.cli.Conn())
	return client.ChangeDetail(ctx, in, opts...)
}

func (m *defaultUserService) ChangePasswordByCaptcha(ctx context.Context, in *ChangePasswordByCaptchaRequest, opts ...grpc.CallOption) (*Response, error) {
	client := pb.NewUserServiceClient(m.cli.Conn())
	return client.ChangePasswordByCaptcha(ctx, in, opts...)
}

func (m *defaultUserService) ChangePasswordByPassword(ctx context.Context, in *ChangePasswordByPasswordRequest, opts ...grpc.CallOption) (*Response, error) {
	client := pb.NewUserServiceClient(m.cli.Conn())
	return client.ChangePasswordByPassword(ctx, in, opts...)
}
