package dao

import (
	"mall/internal/core"
	"mall/internal/dao/redis"
	"testing"
)

func TestRedisConn(t *testing.T) {
	var err error
	err = core.InitConfig("../../configs/config.yaml")
	if err != nil {
		t.Errorf("init config failed: %v", err)
	}
	t.Log("init config success")
	redisDao := redis.GetRedisInstance("")
	if redisDao == nil {
		t.Error("get redis instance failed")
	}
	defer func() {
		_ = redisDao.Close()
	}()

	str, err := redisDao.Ping(t.Context()).Result()
	if err != nil {
		t.Errorf("ping redis failed: %v", err)
	}
	t.Log("redis ping result:", str)
	t.Log("ping redis success")
}
