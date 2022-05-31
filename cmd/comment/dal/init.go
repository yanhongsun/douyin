package dal

import (
	"douyin/cmd/comment/dal/mysqldb"
	"douyin/cmd/comment/dal/redisdb"
)

func Init() {
	mysqldb.Init()
	redisdb.Init()
}
