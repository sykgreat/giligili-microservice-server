// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"

	"giligili/app/user/rpc/internal/logic"
	"giligili/app/user/rpc/internal/svc"
	"giligili/app/user/rpc/pb"
)

type UserServiceServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedUserServiceServer
}

func NewUserServiceServer(svcCtx *svc.ServiceContext) *UserServiceServer {
	return &UserServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServiceServer) LoginByPassword(ctx context.Context, in *pb.LoginByPasswordRequest) (*pb.LoginResponse, error) {
	l := logic.NewLoginByPasswordLogic(ctx, s.svcCtx)
	return l.LoginByPassword(in)
}

func (s *UserServiceServer) LoginByCaptcha(ctx context.Context, in *pb.LoginByCaptchaRequest) (*pb.LoginResponse, error) {
	l := logic.NewLoginByCaptchaLogic(ctx, s.svcCtx)
	return l.LoginByCaptcha(in)
}

func (s *UserServiceServer) Logout(ctx context.Context, in *pb.LogoutRequest) (*pb.Response, error) {
	l := logic.NewLogoutLogic(ctx, s.svcCtx)
	return l.Logout(in)
}

func (s *UserServiceServer) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.Response, error) {
	l := logic.NewRegisterLogic(ctx, s.svcCtx)
	return l.Register(in)
}

func (s *UserServiceServer) GetDetail(ctx context.Context, in *pb.GetDetailRequest) (*pb.GetDetailResponse, error) {
	l := logic.NewGetDetailLogic(ctx, s.svcCtx)
	return l.GetDetail(in)
}

func (s *UserServiceServer) ChangeDetail(ctx context.Context, in *pb.ChangeDetailRequest) (*pb.Response, error) {
	l := logic.NewChangeDetailLogic(ctx, s.svcCtx)
	return l.ChangeDetail(in)
}

func (s *UserServiceServer) ChangePasswordByCaptcha(ctx context.Context, in *pb.ChangePasswordByCaptchaRequest) (*pb.Response, error) {
	l := logic.NewChangePasswordByCaptchaLogic(ctx, s.svcCtx)
	return l.ChangePasswordByCaptcha(in)
}

func (s *UserServiceServer) ChangePasswordByPassword(ctx context.Context, in *pb.ChangePasswordByPasswordRequest) (*pb.Response, error) {
	l := logic.NewChangePasswordByPasswordLogic(ctx, s.svcCtx)
	return l.ChangePasswordByPassword(in)
}

func (s *UserServiceServer) GenerateToken(ctx context.Context, in *pb.GenerateTokenReq) (*pb.GenerateTokenResp, error) {
	l := logic.NewGenerateTokenLogic(ctx, s.svcCtx)
	return l.GenerateToken(in)
}
