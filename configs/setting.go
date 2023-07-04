package configs

import (
	"time"
)

const (
	configName string = "config"
	configType string = "yaml"
)

type ServerConfigs struct {
	RunMode      string        `mapstructure:"run_mode" json:"run_mode" yaml:"run_mode"`
	Addr         string        `mapstructure:"addr" json:"addr" yaml:"addr"`                            //
	ReadTimeout  time.Duration `mapstructure:"read_timeout" json:"read_timeout" yaml:"read_timeout"`    //
	WriteTimeout time.Duration `mapstructure:"write_timeout" json:"write_timeout" yaml:"write_timeout"` //
}

type PgsqlConfigs struct {
	Host           string `mapstructure:"host" json:"host" yaml:"host"`                                      // 服务器地址
	Port           string `mapstructure:"port" json:"port" yaml:"port"`                                      // 端口
	Dbname         string `mapstructure:"dbname" json:"dbname" yaml:"dbname"`                                // 数据库名
	User           string `mapstructure:"user" json:"user" yaml:"user"`                                      // 数据库用户名
	Password       string `mapstructure:"password" json:"password" yaml:"password"`                          // 数据库密码
	MaxIdleConns   int    `mapstructure:"max_idle_conns" json:"max_idle_conns" yaml:"max_idle_conns"`        // 空闲中的最大连接数
	MaxOpenConns   int    `mapstructure:"max_open_conns" json:"max_open_conns" yaml:"max_open_conns"`        // 打开到数据库的最大连接数
	MaxLifeSeconds int64  `mapstructure:"max_life_seconds" json:"max_life_seconds" yaml:"max_life_seconds" ` // 数据库连接最长生命周期
}

type RedisConfigs struct {
	Addr     string `mapstructure:"addr" json:"addr" yaml:"addr"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	DB       int    `mapstructure:"db" json:"db" yaml:"db"`
}

type SessConfigs struct {
	Secret string `mapstructure:"secret" json:"secret" yaml:"secret"`
	Domain string `mapstructure:"domain" json:"domain" yaml:"domain"`
}

type SSOGrpcConfigs struct {
	Addr string `mapstructure:"addr" json:"addr" yaml:"addr"`
}

type SMSConfigs struct {
	Token string `mapstructure:"token" json:"token" yaml:"token"`
}

type settings struct {
	Server  ServerConfigs  `mapstructure:"Server" yaml:"Server"`
	Pgsql   PgsqlConfigs   `mapstructure:"Pgsql" yaml:"Pgsql"`
	Redis   RedisConfigs   `mapstructure:"Redis" yaml:"Redis"`
	Sess    SessConfigs    `mapstructure:"Sess" yaml:"Sess"`
	SSOGrpc SSOGrpcConfigs `mapstructure:"SSOGrpc" yaml:"SSOGrpc"`
	SMS     SMSConfigs     `mapstructure:"SMS" yaml:"SMS"`
}
