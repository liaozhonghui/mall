package core

import "time"

var GlobalConfig MallConfig

type MallConfig struct {
	Server ServerConfig  `mapstructure:"server"`
	Logger LoggerConfig  `mapstructure:"logger"`
	Mysql  []MysqlConfig `mapstructure:"mysql"`
	Jwt    JwtConfig     `mapstructure:"jwt"`
	Redis  []RedisConfig `mapstructure:"redis"`
}

type ServerConfig struct {
	Addr         string        `mapstructure:"addr"`
	ReadTimeout  time.Duration `mapstructure:"readTimeOut"`
	WriteTimeout time.Duration `mapstructure:"writeTimeOut"`
	IdleTimeout  time.Duration `mapstructure:"idleTimeOut"`
}

type MysqlConfig struct {
	Instance      string        `mapstructure:"instance"`
	Dsn           string        `mapstructure:"dsn"`
	TraceLog      bool          `mapstructure:"trace_log"`
	SlowThreshold time.Duration `mapstructure:"slow_threshold"`
}

type LoggerConfig struct {
	LogFile  string `mapstructure:"logFile"`
	LogLevel string `mapstructure:"logLevel"`
}

type JwtConfig struct {
	ApiSecret   string        `mapstructure:"api_secret"`
	ExpireTime  time.Duration `mapstructure:"expireTime"`
	AdminSecret string        `mapstructure:"admin_secret"`
}

type RedisConfig struct {
	Instance     string `mapstructure:"instance"`
	Addr         string `mapstructure:"addr"`
	Password     string `mapstructure:"password"`
	DB           int    `mapstructure:"db"`
	DialTimeout  int    `mapstructure:"dial_timeout"`
	ReadTimeout  int    `mapstructure:"read_timeout"`
	WriteTimeout int    `mapstructure:"write_timeout"`
}
