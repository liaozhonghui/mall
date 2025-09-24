package redis

import (
	"mall/internal/core"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

var rdbInstance map[string]*redis.Client
var once sync.Once

func initRedis() {
	redisCfgs := core.GlobalConfig.Redis

	tmpInstance := make(map[string]*redis.Client, 0)
	for _, redisCfg := range redisCfgs {
		rdb := redis.NewClient(&redis.Options{
			Addr:         redisCfg.Addr,
			Password:     redisCfg.Password,
			DB:           redisCfg.DB,
			DialTimeout:  time.Duration(redisCfg.DialTimeout) * time.Millisecond,
			ReadTimeout:  time.Duration(redisCfg.ReadTimeout) * time.Millisecond,
			WriteTimeout: time.Duration(redisCfg.WriteTimeout) * time.Millisecond,
		})
		tmpInstance[redisCfg.Instance] = rdb
	}
	rdbInstance = tmpInstance
}

func GetRedisInstance(instance string) *redis.Client {
	if instance == "" {
		instance = "default"
	}
	if rdb, ok := rdbInstance[instance]; ok {
		return rdb
	}

	once.Do(func() {
		initRedis()
	})

	return rdbInstance[instance]
}
