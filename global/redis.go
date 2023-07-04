package global

import (
	"UniqueRecruitmentBackend/configs"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var redisCli *redis.Client

func GetRedisCli() *redis.Client {
	return redisCli
}

func setupRedis() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     configs.Config.Redis.Addr,
		Password: configs.Config.Redis.Password,
		DB:       configs.Config.Redis.DB,
	})
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		panic(fmt.Sprintf("connect to redis error, %v", err))
	}
	redisCli = rdb
}
