package logic

import (
	"context"
	"giligili/app/captcha/rpc/captchaservice"
	"giligili/common"
	"giligili/common/enum"
	"giligili/common/password"
	"giligili/common/xerr"
	"giligili/model/user"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"strconv"
	"time"

	"giligili/app/user/rpc/internal/svc"
	"giligili/app/user/rpc/pb"

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
	// 验证邮箱格式
	if format := common.VerifyEmailFormat(in.Email); !format {
		return nil, errors.Wrapf(xerr.NewErrMsg("邮箱格式不正确!"), "邮箱格式不正确!")
	}

	// 验证邮箱是否已被注册
	email, err := l.svcCtx.UserModel.FindOneByEmail(l.ctx, in.Email)
	if err != nil && err != sqlc.ErrNotFound { // 数据库查询错误
		return nil, errors.Wrapf(xerr.NewErrMsg("用户注册失败!"), "用户注册失败!")
	}
	if email != nil { // 邮箱已被注册
		return nil, errors.Wrapf(xerr.NewErrMsg("该邮箱已被注册!"), "该邮箱已被注册!")
	}

	// 验证验证码
	_, err = l.svcCtx.CaptchaRpc.VerifyCaptcha(l.ctx, &captchaservice.VerifyCaptchaReq{
		Email:       in.Email,
		Captcha:     in.Captcha,
		CaptchaType: enum.CaptchaRegister,
	})
	if err != nil { // 验证码错误
		return nil, errors.Wrapf(xerr.NewErrMsg("验证码输入错误，请重新输入!"), "验证码输入错误，请重新输入!")
	}

	// 验证密码
	if common.VerifyPasswordFormat(in.Password) {
		return nil, errors.Wrapf(xerr.NewErrMsg("密码不符合要求!"), "密码不符合要求!")
	}

	// 密码加密
	generatePassword, err := password.GeneratePassword(in.Password)
	if err != nil { // 密码加密失败
		return nil, errors.Wrapf(xerr.NewErrMsg("密码生成失败!"), "密码生成失败!")
	}

	// 用户信息
	snowflakeId := l.svcCtx.Snowflake.NextVal() // 生成用户id
	now := time.Now()
	u := &user.User{
		Id:          snowflakeId,
		Username:    "用户" + strconv.FormatInt(snowflakeId, 10),
		Email:       in.Email,
		Password:    generatePassword,
		Birthday:    now,
		CreatedTime: now,
		UpdatedTime: now,
		Avatar:      "https://i1.hdslb.com/bfs/face/2c72223afa74b0036daee60cd99c069760b653df.jpg@240w_240h_1c_1s_!web-avatar-space-header.avif",
		SpaceCover:  "https://i0.hdslb.com/bfs/space/cb1c3ef50e22b6096fde67febe863494caefebad.png@2560w_400h_100q_1o.webp",
		Sign:        "这个人很懒, 什么都没有留下!",
		Status:      "1",
		ClientIp:    in.ClientIp,
	}

	// 插入用户
	_, err = l.svcCtx.UserModel.Insert(l.ctx, u)
	if err != nil { // 插入用户失败
		return nil, errors.Wrapf(xerr.NewErrMsg("用户注册失败!"), "用户注册失败!")
	}

	return &pb.Response{}, nil
}
