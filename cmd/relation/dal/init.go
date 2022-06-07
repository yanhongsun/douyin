package dal

import (
	"douyin/cmd/relation/dal/db"
	"douyin/cmd/relation/dal/redisCache"
)

func Init() {
	db.Init()              // mysql
	redisCache.InitCache() // 缓存
}
