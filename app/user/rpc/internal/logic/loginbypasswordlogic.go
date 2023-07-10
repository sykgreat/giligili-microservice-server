package logic

import (
	"context"
	"giligili/app/user/rpc/internal/svc"
	"giligili/app/user/rpc/pb"
	"giligili/common"
	"giligili/common/enum"
	"giligili/common/password"
	"giligili/common/xerr"
	"github.com/pkg/errors"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginByPasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginByPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginByPasswordLogic {
	return &LoginByPasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginByPasswordLogic) LoginByPassword(in *pb.LoginByPasswordRequest) (*pb.LoginResponse, error) {
	// 验证邮箱格式
	if format := common.VerifyEmailFormat(in.Email); !format {
		return nil, errors.Wrapf(xerr.NewErrMsg("邮箱格式不正确!"), "邮箱格式不正确!")
	}

	// 查询用户是否存在
	userInfo, err := l.svcCtx.UserModel.FindOneByEmail(l.ctx, in.Email)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("用户登录失败，请检查账号或密码是否输入正确!"), "用户登录失败，请检查账号或密码是否输入正确!")
	}

	// 验证密码
	if err := password.ComparePassword(userInfo.Password, in.Password); err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("用户登录失败，请检查账号或密码是否输入正确!"), "用户登录失败，请检查账号或密码是否输入正确!")
	}

	// 生成accessToken
	accessToken, err := l.svcCtx.Jwt.GenerateAccessToken(userInfo.Id, userInfo.Email)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("用户登录失败，验证token生成失败!"), "用户登录失败，验证token生成失败!")
	}

	// 生成refreshToken
	refreshToken, err := l.svcCtx.Jwt.GenerateRefreshToken(userInfo.Id, userInfo.Email)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("用户登录失败，刷新token生成失败!"), "用户登录失败，刷新token生成失败!")
	}

	// 将accessToken和refreshToken存入redis
	if err := l.svcCtx.Redis.SetexCtx(
		l.ctx,
		enum.UserModule+":"+enum.Token+":"+strconv.Itoa(int(userInfo.Id))+":"+enum.AccessToken,
		accessToken,
		l.svcCtx.Jwt.AccessTokenExpire,
	); err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("用户登录失败，accessToken存入redis失败!"), "用户登录失败，accessToken存入redis失败!")
	}
	if err := l.svcCtx.Redis.SetexCtx(
		l.ctx,
		enum.UserModule+":"+enum.Token+":"+strconv.Itoa(int(userInfo.Id))+":"+enum.RefreshToken,
		refreshToken,
		l.svcCtx.Jwt.RefreshTokenExpire,
	); err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("用户登录失败，refreshToken存入redis失败!"), "用户登录失败，refreshToken存入redis失败!")
	}

	return &pb.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
