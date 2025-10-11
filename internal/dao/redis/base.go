package redis

import (
	"mall/internal/core"
	"sync"

	"github.com/go-redis/redis/v8"
)

// go-redis
var rdb *redis.Client
var once sync.Once

func initRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:         core.GlobalConfig.Redis.Addr,
		Password:     core.GlobalConfig.Redis.Password,
		DB:           core.GlobalConfig.Redis.Db,
		DialTimeout:  core.GlobalConfig.Redis.DialTimeout,  // 连接建立超时，默认5秒
		ReadTimeout:  core.GlobalConfig.Redis.ReadTimeout,  // 读超时，默认3秒，-1表示取消读超时
		WriteTimeout: core.GlobalConfig.Redis.WriteTimeout, // 写超时，默认等于读超时
		PoolSize:     core.GlobalConfig.Redis.PoolSize,     // 连接池大小，默认10个连接
		MinIdleConns: core.GlobalConfig.Redis.MinIdleConns, // 最小空闲连接数，默认0
	})
}
func GetInstance() *redis.Client {
	if rdb != nil {
		return rdb
	}
	once.Do(func() {
		initRedis()
	})
	return rdb
}
