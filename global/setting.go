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
		panic(fmt.Errorf("router config setup failed %s", err))
	}
	err = setting.ReadConfig("Pgsql", &PgsqlConfig)
	if err != nil {
		panic(fmt.Errorf("pgsql config setup failed %s", err))
	}
	err = setting.ReadConfig("SSOGrpc", &SSOGrpcConfigs)
	if err != nil {
		panic(fmt.Errorf("sso grpc config setup failed %s", err))
	}

}
