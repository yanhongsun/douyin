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
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}
	//TODO自动建表
	// m := DB.Migrator()
	// if m.HasTable(&Note{}) {
	// 	return
	// }
	// if err = m.CreateTable(&Note{}); err != nil {
	// 	panic(err)
	// }
}
