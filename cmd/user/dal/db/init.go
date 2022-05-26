package db

import (
	"fmt"
	"github.com/douyin/cmd/user/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/opentracing"
)

var DB *gorm.DB

func Init() {
	var err error
	// TODO: add mysql url
	s := "%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local"
	DB, err = gorm.Open(mysql.Open(fmt.Sprintf(s,
		global.DatabaseSetting.UserName,
		global.DatabaseSetting.Password,
		global.DatabaseSetting.Host,
		global.DatabaseSetting.DBName,
		global.DatabaseSetting.Charset,
		global.DatabaseSetting.ParseTime)),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		})
	if err != nil {
		panic(err)
	}

	if err = DB.Use(gormopentracing.New()); err != nil {
		panic(err)
	}

	m := DB.Migrator()
	if m.HasTable(&User{}) {
		return
	}
	if err = m.CreateTable(&User{}); err != nil {
		panic(err)
	}
}
