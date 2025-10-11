package core

import "time"

var GlobalConfig Config

type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Mysql    []MysqlConfig  `mapstructure:"mysql"`
	Postgres PostgresConfig `mapstructure:"postgres"`
	Jwt      JwtConfig      `mapstructure:"jwt"`
	Redis    RedisConfig    `mapstructure:"redis"`
	Logger   LoggerConfig   `mapstructure:"logger"`
}

type ServerConfig struct {
	Addr         string        `mapstructure:"addr"`
	ReadTimeout  time.Duration `mapstructure:"read_timeout"`
	WriteTimeout time.Duration `mapstructure:"write_timeout"`
	IdleTimeout  time.Duration `mapstructure:"idle_timeout"`
}

type PostgresConfig struct {
	Dsn           string        `mapstructure:"dsn"`
	TraceLog      bool          `mapstructure:"trace_log"`
	SlowThreshold time.Duration `mapstructure:"slow_threshold"`
	Log           string        `mapstructure:"log"`
	LogLevel      int           `mapstructure:"log_level"`
}
type MysqlConfig struct {
	Instance      string        `mapstructure:"instance"`
	Dsn           string        `mapstructure:"dsn"`
	TraceLog      bool          `mapstructure:"trace_log"`
	SlowThreshold time.Duration `mapstructure:"slow_threshold"`
}

type JwtConfig struct {
	ApiSecret  string `mapstructure:"api_secret"`
	ExpireTime int    `mapstructure:"expire_time"`
}
type RedisConfig struct {
	Addr         string        `mapstructure:"addr"`
	Password     string        `mapstructure:"password"`
	Db           int           `mapstructure:"db"`
	DialTimeout  time.Duration `mapstructure:"dial_timeout"`
	ReadTimeout  time.Duration `mapstructure:"read_timeout"`
	WriteTimeout time.Duration `mapstructure:"write_timeout"`
	PoolSize     int           `mapstructure:"pool_size"`
	MinIdleConns int           `mapstructure:"min_idle_conns"`
}

type LoggerConfig struct {
	LogFile  string `mapstructure:"log_file"`
	LogLevel string `mapstructure:"log_level"`
}
