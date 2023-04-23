package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

const (
	configName string = "config"
	configType string = "yaml"
)

type Setting struct {
	v *viper.Viper
}

func NewSetting() (*Setting, error) {
	setting := new(Setting)
	v := viper.New()
	v.AddConfigPath(".")
	v.SetConfigName(configName)
	v.SetConfigType(configType)
	setting.v = v
	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("fatal error config file: %s ", err)
	}
	return setting, nil
}
