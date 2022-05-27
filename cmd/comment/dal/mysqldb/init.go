package mysqldb

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	var err error
	DB, err = gorm.Open(mysql.Open("gorm:gorm@tcp(localhost:3306)/douyin?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
	})

	if err != nil {
		panic(err)
	}

	m := DB.Migrator()

	if m.HasTable(&Comment{}) {
		return
	}

	if err = m.CreateTable(&Comment{}); err != nil {
		panic(err)
	}

	if m.HasTable(&CommentIndex{}) {
		return
	}

	if err = m.CreateTable(&CommentIndex{}); err != nil {
		panic(err)
	}
}
