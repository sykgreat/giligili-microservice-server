package logic

import (
	"context"
	"giligili/app/captcha/rpc/captchaservice"
	"giligili/app/user/model"
	"giligili/app/user/rpc/internal/svc"
	"giligili/app/user/rpc/pb"
	"giligili/app/user/utils/password"
	"giligili/common/xerr"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *pb.RegisterRequest) (*pb.Response, error) {
	// 通过邮箱查询用户信息
	user, _ := l.svcCtx.UserModel.FindOneByEmail(l.ctx, in.Email)
	if user != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("该邮箱已经注册！"), "该邮箱已经注册: %s", in.Email)
	}

	// 使用验证码rpc验证验证码
	if verifyResult, err := l.svcCtx.CaptchaRpc.VerifyCaptcha(l.ctx, &captchaservice.VerifyCaptchaReq{
		Email:   in.Email,
		Captcha: in.Captcha,
	}); err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("验证码验证失败"), "VerifyCaptcha err:%v", err)
	} else if verifyResult.Result != 200 {
		return nil, errors.Wrapf(xerr.NewErrMsg("验证码验证失败"), "验证码输入错误, 请重新输入!")
	}

	// 密码加密
	s, err := password.GeneratePassword(in.Password)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("密码加密失败"), "GeneratePassword err:%v", err)
	}

	t := time.Now()
	_, err = l.svcCtx.UserModel.Insert(
		l.ctx,
		&model.User{
			Id:          l.svcCtx.Snowflakes.NextVal(),
			Email:       in.Email,
			Password:    s,
			Username:    strings.Split(in.Email, "@")[0],
			Birthday:    t,
			CreatedTime: t,
			UpdatedTime: t,
			Avatar:      "https://i1.hdslb.com/bfs/face/2c72223afa74b0036daee60cd99c069760b653df.jpg@240w_240h_1c_1s_!web-avatar-space-header.avif",
			SpaceCover:  "https://i0.hdslb.com/bfs/space/cb1c3ef50e22b6096fde67febe863494caefebad.png@2560w_400h_100q_1o.webp",
			Sign:        "这个人很懒, 什么都没有留下!",
			Status:      "1",
			ClientIp:    in.ClientIp,
		},
	)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("用户注册失败"), "Insert err:%v", err)
	}

	return &pb.Response{
		Result: 200,
	}, nil
}
