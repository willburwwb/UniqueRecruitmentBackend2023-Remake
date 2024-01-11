package configs

import (
	"time"
)

const (
	configName string = "config.local"
	configType string = "yml"
)

type server struct {
	RunMode       string        `mapstructure:"run_mode" json:"run_mode" yaml:"run_mode"`
	Addr          string        `mapstructure:"addr" json:"addr" yaml:"addr"`                            //
	ReadTimeout   time.Duration `mapstructure:"read_timeout" json:"read_timeout" yaml:"read_timeout"`    //
	WriteTimeout  time.Duration `mapstructure:"write_timeout" json:"write_timeout" yaml:"write_timeout"` //
	SessionSecret string        `mapstructure:"session_secret" json:"session_secret" yaml:"session_secret"`
	SessionDomain string        `mapstructure:"session_domain" json:"session_domain" yaml:"session_domain"`
}

type pgsql struct {
	Host           string `mapstructure:"host" json:"host" yaml:"host"`                                      // 服务器地址
	Port           string `mapstructure:"port" json:"port" yaml:"port"`                                      // 端口
	Dbname         string `mapstructure:"dbname" json:"dbname" yaml:"dbname"`                                // 数据库名
	User           string `mapstructure:"user" json:"user" yaml:"user"`                                      // 数据库用户名
	Password       string `mapstructure:"password" json:"password" yaml:"password"`                          // 数据库密码
	MaxIdleConns   int    `mapstructure:"max_idle_conns" json:"max_idle_conns" yaml:"max_idle_conns"`        // 空闲中的最大连接数
	MaxOpenConns   int    `mapstructure:"max_open_conns" json:"max_open_conns" yaml:"max_open_conns"`        // 打开到数据库的最大连接数
	MaxLifeSeconds int64  `mapstructure:"max_life_seconds" json:"max_life_seconds" yaml:"max_life_seconds" ` // 数据库连接最长生命周期
}

type redis struct {
	Dsn string `mapstructure:"dsn" json:"dsn" yaml:"dsn"`
}

type sso struct {
	Addr string `mapstructure:"addr" json:"addr" yaml:"addr"`
}

type grpc struct {
	Addr string `mapstructure:"addr" json:"addr" yaml:"addr"`
}

type sms struct {
	Token string `mapstructure:"token" json:"token" yaml:"token"`
}

type cos struct {
	CosUrl       string `mapstructure:"cos_url" json:"cos_url" yaml:"cos_url"`
	CosSecretID  string `mapstructure:"cos_secret_id" json:"cos_secret_id" yaml:"cos_secret_id"`
	CosSecretKey string `mapstructure:"cos_secret_key" json:"cos_secret_key" yaml:"cos_secret_key"`
}

type apm struct {
	Name          string `mapstructure:"name" json:"name" yaml:"name"`
	ReportBackend string `mapstructure:"report_backend" json:"report_backend" yaml:"report_backend"`
}

type settings struct {
	Server server `mapstructure:"server" yaml:"server"`
	Pgsql  pgsql  `mapstructure:"pgsql" yaml:"pgsql"`
	Redis  redis  `mapstructure:"redis" yaml:"redis"`
	SSO    sso    `mapstructure:"sso" yaml:"sso"`
	Grpc   grpc   `mapstructure:"grpc" yaml:"grpc"`
	SMS    sms    `mapstructure:"sms" yaml:"sms"`
	COS    cos    `mapstructure:"COS" yaml:"COS"`
	Apm    apm    `mapstructure:"apm" yaml:"apm"`
}
