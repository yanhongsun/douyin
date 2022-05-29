package db

import (
	"douyin/pkg/constants"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Init init DB
func Init() {
	var err error
	DB, err = gorm.Open(mysql.Open(constants.MySQLDefaultDSN),
		&gorm.Config{
			PrepareStmt:            true, // 预编译缓存
			SkipDefaultTransaction: true, // 跳过默认事务
		},
	)
	if err != nil {
		panic(err)
	}

	// if err = DB.Use(gormopentracing.New()); err != nil {
	// 	panic(err)
	// }
	// //数据库迁移功能
	// m := DB.Migrator()
	// if m.HasTable(&Note{}) {
	// 	return
	// }
	// if err = m.CreateTable(&Note{}); err != nil {
	// 	panic(err)
	// }
	//TODO  建表
}
