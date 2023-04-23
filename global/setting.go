package global

import (
	"UniqueRecruitmentBackend/configs"
	"fmt"
)

func init() {
	setting, err := configs.NewSetting()
	if err != nil {
		panic(fmt.Errorf("config setup failed %s", err))
	}
	err = setting.ReadConfig("Server", &ServerConfig)
	if err != nil {
		panic(fmt.Errorf("server config setup failed %s", err))
	}
	err = setting.ReadConfig("Pgsql", &PgsqlConfig)
	if err != nil {
		panic(fmt.Errorf("pgsql config setup failed %s", err))
	}
	err = setting.ReadConfig("Redis", &RedisConfig)
	if err != nil {
		panic(fmt.Errorf("redis config setup failed %s", err))
	}
	err = setting.ReadConfig("Session", &SessConfig)
	if err != nil {
		panic(fmt.Errorf("session config setup failed %s", err))
	}
}
