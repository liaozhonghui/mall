package cache

import (
	"context"
	"sync"
	"time"

	"github.com/allegro/bigcache/v3"
)

var cache *bigcache.BigCache
var once sync.Once

func initBigCache() {
	inst, err := bigcache.New(context.Background(), bigcache.DefaultConfig(time.Minute*10))
	if err != nil {
		panic(err)
	}
	cache = inst
}

func GetInstance() *bigcache.BigCache {
	if cache != nil {
		return cache
	}
	once.Do(func() {
		initBigCache()
	})
	return cache
}
