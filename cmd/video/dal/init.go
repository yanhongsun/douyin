package dal

import (
	"douyin/cmd/video/dal/cache"
	"douyin/cmd/video/dal/db"
)

func Init() {

	db.Init()
	cache.Init()
}
