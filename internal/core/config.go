package core

import "time"

var GlobalConfig MallConfig

type MallConfig struct {
	Server ServerConfig  `mapstructure:"server"`
	Logger LoggerConfig  `mapstructure:"logger"`
	Mysql  []MysqlConfig `mapstructure:"mysql"`
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
