package minicache

import (
	"douyin/cmd/video/dal/minicache/cache"
	"time"
)

var Caches *cache.Cache

func Init() {
	//配置cache参数
	Caches = cache.NewCache(cache.WithAutoGC(10*time.Minute), cache.WithAutoRandGC(time.Second),
		cache.WithSegmentSize(64), cache.WithMapSize(128))
}
