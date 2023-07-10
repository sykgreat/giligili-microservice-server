package svc

import (
	"giligili/app/captcha/rpc/captchaservice"
	"giligili/app/user/rpc/internal/config"
	"giligili/common"
	"giligili/common/jwt"
	"giligili/model/user"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	UserModel  user.UserModel
	Redis      *redis.Redis
	CaptchaRpc captchaservice.CaptchaService
	Snowflake  *common.Snowflake
	Jwt        *jwt.Jwt
}

func NewServiceContext(c config.Config) *ServiceContext {
	snowflake, err := common.NewSnowflake(c.Snowflake.WorkerId, c.Snowflake.DatacenterId, c.Snowflake.Sequence)
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config:     c,
		UserModel:  user.NewUserModel(sqlx.NewMysql(c.Mysql.DataSource), c.Cache),
		Redis:      redis.MustNewRedis(c.Redis.RedisConf),
		CaptchaRpc: captchaservice.NewCaptchaService(zrpc.MustNewClient(c.CaptchaRpcConf)),
		Snowflake:  snowflake,
		Jwt:        &c.Jwt,
	}
}
