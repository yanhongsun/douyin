package dal

import (
	"douyin/cmd/video/dal/db"
	"douyin/cmd/video/dal/minicache"
)

func Init() {

	db.Init()
	minicache.Init()
}
