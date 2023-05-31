package logic

import (
	"context"
	"encoding/json"
	"strconv"

	"giligili/app/captcha/rpc/captchaservice"
	"giligili/app/user/rpc/internal/svc"
	"giligili/app/user/rpc/pb"
	"giligili/common/xerr"

	"github.com/pkg/errors"
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
	// 通过邮箱查询用户信息
	user, err := l.svcCtx.UserModel.FindOneByEmail(l.ctx, in.Email)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "根据邮箱查询用户信息失败, email:%s,err:%v", in.Email, err)
	}
	if user == nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("用户不存在"), "email:%s", in.Email)
	}

	// 通过captchaRpc验证码校验
	if verifyResult, err := l.svcCtx.CaptchaRpc.VerifyCaptcha(l.ctx, &captchaservice.VerifyCaptchaReq{
		Email:   in.Email,
		Captcha: in.Captcha,
	}); err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("验证码校验失败"), "VerifyCaptcha, err:%v", err)
	} else if verifyResult.Result != 200 {
		return nil, errors.Wrapf(xerr.NewErrMsg("验证码校验失败"), "验证码输入错误, 请重新输入!")
	}

	// 生成token
	generateTokenLogic := NewGenerateTokenLogic(l.ctx, l.svcCtx)
	token, err := generateTokenLogic.GenerateToken(&pb.GenerateTokenReq{
		UserId: user.Id,
		Email:  user.Email,
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("生成token失败"), "GenerateToken userId : %d", user.Id)
	}

	// 将token存储到redis
	if err = l.svcCtx.Redis.SetexCtx(
		l.ctx,
		l.svcCtx.Config.Redis.Key+":token:"+strconv.Itoa(int(user.Id)),
		token.AccessToken,
		int(token.AccessExpire),
	); err != nil {
		return nil, errors.Wrapf(
			xerr.NewErrMsg("redis存储token失败"),
			"SetexCtx token:%s, expire:%d", token.AccessToken, token.AccessExpire,
		)
	}

	// 序列化用户对象
	marshal, err := json.Marshal(user)
	if err != nil {
		return nil, errors.Wrapf(
			xerr.NewErrMsg("用户序列化失败！"),
			"json Marshal err: %s", err,
		)
	}

	// 存储用户信息
	if err := l.svcCtx.Redis.SetCtx(
		l.ctx,
		l.svcCtx.Config.Redis.Key+":info:"+strconv.Itoa(int(user.Id)),
		string(marshal),
	); err != nil {
		return nil, errors.Wrapf(
			xerr.NewErrMsg("redis存用户信息失败"),
			"redis set user info err: %s, ", err,
		)
	}

	return &pb.LoginResponse{
		AccessToken:  token.AccessToken,
		AccessExpire: token.AccessExpire,
		RefreshAfter: token.RefreshAfter,
	}, nil
}
