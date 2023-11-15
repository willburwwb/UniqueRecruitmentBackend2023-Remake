package global

import (
	"UniqueRecruitmentBackend/configs"
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/xylonx/zapx"
	"go.uber.org/zap"
)

var redisCli *redis.Client

func GetRedisCli() *redis.Client {
	return redisCli
}

func setupRedis() {
	redisOptions, err := redis.ParseURL(configs.Config.Redis.Dsn)
	if err != nil {
		zapx.With(zap.Error(err)).Error("parse redis dsn rerror")
		panic(err)
	}

	rdb := redis.NewClient(redisOptions)
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		zapx.With(zap.Error(err)).Error("connect to redis rerror")
		panic(err)
	}
	redisCli = rdb
}
