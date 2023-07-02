package logic

import (
	"context"
	"time"

	"giligili/app/user/rpc/internal/svc"
	"giligili/app/user/rpc/pb"
	"giligili/common/ctxdata"
	"giligili/common/xerr"

	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GenerateTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGenerateTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateTokenLogic {
	return &GenerateTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GenerateTokenLogic) GenerateToken(in *pb.GenerateTokenReq) (*pb.GenerateTokenResp, error) {
	now := time.Now().Unix()                             // 当前时间戳
	accessExpire := l.svcCtx.Config.JwtAuth.AccessExpire // 有效期
	// 生成token
	accessToken, err := l.getJwtToken(l.svcCtx.Config.JwtAuth.AccessSecret, now, accessExpire, in.UserId, in.Email)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("生成token失败"), "getJwtToken err userId:%d , err:%v", in.UserId, err)
	}

	return &pb.GenerateTokenResp{
		AccessToken:  accessToken,
		AccessExpire: now + accessExpire,
		RefreshAfter: now + accessExpire/2,
	}, nil
}

// getJwtToken 生成jwt
func (l *GenerateTokenLogic) getJwtToken(secretKey string, iat, seconds, userId int64, email string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims[ctxdata.CtxKeyJwtUserId] = userId
	claims["email"] = email
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
