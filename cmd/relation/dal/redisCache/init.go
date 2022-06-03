package redisCache

import (
	"douyin/cmd/relation/dal/redisCache/cache"
	"time"
)

var Caches *cache.Cache

func InitCache() {
	// 默认开启singleflight
	// WithAutoGC : 设置定期全部清理的时间 （默认不清理），2.WithAutoRandGC: 设置定期随机清理的时间
	//3. WithSegmentSize: 设置分段个数，默认分段128个   4. WithMapSize: 设置每个分段里map的大小，默认每个map的大小为128个对象
	Caches = cache.NewCache(cache.WithAutoGC(10*time.Minute), cache.WithAutoRandGC(time.Second),
		cache.WithSegmentSize(64), cache.WithMapSize(128))
	//caches = cache.NewCache(cache.WithSegmentSize(128), cache.WithMapSize(200))
}
