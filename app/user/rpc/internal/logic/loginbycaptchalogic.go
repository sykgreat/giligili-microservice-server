package logic

import (
	"context"
	"giligili/app/captcha/rpc/captchaservice"
	"giligili/common/enum"
	"giligili/common/xerr"
	"github.com/pkg/errors"
	"strconv"

	"giligili/app/user/rpc/internal/svc"
	"giligili/app/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginByCaptchaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginByCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginByCaptchaLogic {
	return &LoginByCaptchaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginByCaptchaLogic) LoginByCaptcha(in *pb.LoginByCaptchaRequest) (*pb.LoginResponse, error) {
	// 验证验证码
	_, err := l.svcCtx.CaptchaRpc.VerifyCaptcha(l.ctx, &captchaservice.VerifyCaptchaReq{
		Email:       in.Email,
		Captcha:     in.Captcha,
		CaptchaType: enum.CaptchaLogin,
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("验证码错误!"), "user login failed! %v", err)
	}

	// 查询用户是否存在
	userInfo, err := l.svcCtx.UserModel.FindOneByEmail(l.ctx, in.Email)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("用户登录失败，请检查账号是否注册!"), "user login failed! %v", err)
	}

	// 生成accessToken
	accessToken, err := l.svcCtx.Jwt.GenerateAccessToken(userInfo.Id, userInfo.Email)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("用户登录失败，验证token生成失败!"), "user login failed! %v", err)
	}

	// 生成refreshToken
	refreshToken, err := l.svcCtx.Jwt.GenerateRefreshToken(userInfo.Id, userInfo.Email)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("用户登录失败，刷新token生成失败!"), "user login failed! %v", err)
	}

	// 将accessToken和refreshToken存入redis
	if err := l.svcCtx.Redis.SetexCtx(
		l.ctx,
		enum.UserModule+enum.Token+strconv.Itoa(int(userInfo.Id))+":"+enum.AccessToken,
		accessToken,
		l.svcCtx.Jwt.AccessTokenExpire,
	); err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("用户登录失败，accessToken存入redis失败!"), "user login failed! %v", err)
	}
	if err := l.svcCtx.Redis.SetexCtx(
		l.ctx,
		enum.UserModule+enum.Token+strconv.Itoa(int(userInfo.Id))+":"+enum.RefreshToken,
		refreshToken,
		l.svcCtx.Jwt.RefreshTokenExpire,
	); err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("用户登录失败，refreshToken存入redis失败!"), "user login failed! %v", err)
	}

	return &pb.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
