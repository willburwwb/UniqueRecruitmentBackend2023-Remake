package configs

import (
	"fmt"
	"github.com/spf13/viper"
)

var Config settings

func Setup() {
	v := viper.New()
	v.AddConfigPath(".")
	v.SetConfigName(configName)
	v.SetConfigType(configType)
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("read config error, %v", err))
	}
	if err := v.Unmarshal(&Config); err != nil {
		panic(fmt.Sprintf("reload config error, %v", err))
	}
}

func init() {
	Setup()
}
