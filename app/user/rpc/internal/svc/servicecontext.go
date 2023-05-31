package svc

import (
	"encoding/binary"
	"giligili/app/captcha/rpc/captchaservice"
	"giligili/app/user/model"
	"giligili/app/user/rpc/internal/config"
	"giligili/common"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	Redis      *redis.Redis                  // 定义redis
	UserModel  model.UserModel               // 定义用户模型
	CaptchaRpc captchaservice.CaptchaService // 定义验证码rpc
	Snowflakes *common.Snowflake             // 定义雪花算法
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 初始化redis
	r, err := redis.NewRedis(c.Redis.RedisConf)
	if err != nil {
		logx.Error(err)
		return nil
	}

	// workId, err := strconv.ParseInt(c.Snowflake.WorkerName, 10, 64)
	// if err != nil {
	// 	logx.Error("Snowflake worker name parse error: ", err)
	// 	return nil
	// }
	// datacenterId, err := strconv.ParseInt(c.Snowflake.DatacenterName, 10, 64)
	// if err != nil {
	// 	logx.Error("Snowflake worker name parse error: ", err)
	// 	return nil
	// }

	snowflakes, err := common.NewSnowflake(
		int64(binary.BigEndian.Uint32([]byte(c.Snowflake.WorkerName))),
		int64(binary.BigEndian.Uint32([]byte(c.Snowflake.DatacenterName))),
		c.Snowflake.Sequence,
	)
	if err != nil {
		logx.Error("Snowflake init error: ", err)
	}

	return &ServiceContext{
		Config: c,

		Redis:      r,                                                                      // 初始化redis
		UserModel:  model.NewUserModel(sqlx.NewMysql(c.Mysql.DataSource), c.Cache),         // 初始化用户模型
		CaptchaRpc: captchaservice.NewCaptchaService(zrpc.MustNewClient(c.CaptchaRpcConf)), // 初始化验证码rpc
		Snowflakes: snowflakes,                                                             // 初始化雪花算法
	}
}
