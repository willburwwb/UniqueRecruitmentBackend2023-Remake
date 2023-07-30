package global

import (
	"UniqueRecruitmentBackend/configs"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	goredis "github.com/redis/go-redis/v9"
	"strconv"
)

var SessStore sessions.Store

func setupSess() {
	rdsOpt, err := goredis.ParseURL(configs.Config.Redis.Dsn)
	if err != nil {
		panic(err)
	}
	SessStore, err = redis.NewStoreWithDB(
		10, rdsOpt.Network, rdsOpt.Addr, rdsOpt.Password,
		strconv.FormatInt(int64(rdsOpt.DB), 10),
		[]byte(configs.Config.Server.SessionSecret),
	)
	if err != nil {
		panic(err)
	}
	SessStore.Options(sessions.Options{Path: "/", Domain: configs.Config.Server.SessionDomain, HttpOnly: true})
}
