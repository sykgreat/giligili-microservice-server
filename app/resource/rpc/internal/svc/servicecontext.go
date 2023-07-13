package svc

import (
	"giligili/app/resource/rpc/internal/config"
	"giligili/common/Snowflake"
	"giligili/common/storage"
	"giligili/model/resource"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config        config.Config
	ResourceModel resource.ResourceModel
	Redis         *redis.Redis
	Snowflake     *Snowflake.Snowflake
	UploadVideo   *storage.Storage
	UploadImg     *storage.Storage
}

func NewServiceContext(c config.Config) *ServiceContext {
	snowflake, err := Snowflake.NewSnowflake(c.Snowflake.WorkerId, c.Snowflake.DatacenterId, c.Snowflake.Sequence)
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config:        c,
		ResourceModel: resource.NewResourceModel(sqlx.NewMysql(c.Mysql.DataSource), c.Cache),
		Redis:         redis.MustNewRedis(c.Redis.RedisConf),
		Snowflake:     snowflake,
		UploadVideo:   storage.New(c.Upload.VideoPath),
		UploadImg:     storage.New(c.Upload.ImgPath),
	}
}
