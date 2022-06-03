package mysqldb

import (
	"douyin/cmd/comment/pack/configdata"
	"douyin/cmd/comment/pack/zapcomment"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormopentracing "gorm.io/plugin/opentracing"
)

var DB *gorm.DB

func Init() {
	var err error
	s := "%s:%s@tcp(%s)/%s?charset=%s&parseTime=%s&loc=Local"

	DB, err = gorm.Open(mysql.Open(fmt.Sprintf(s,
		configdata.MysqlDatabaseConfig.User,
		configdata.MysqlDatabaseConfig.Password,
		configdata.MysqlDatabaseConfig.Host,
		configdata.MysqlDatabaseConfig.Name,
		configdata.MysqlDatabaseConfig.Charset,
		configdata.MysqlDatabaseConfig.ParseTime)),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		})

	if err != nil {
		zapcomment.Logger.Panic("mysql initialization err: " + err.Error())
	}

	if err = DB.Use(gormopentracing.New()); err != nil {
		zapcomment.Logger.Panic("gorm opentracing initialization err: " + err.Error())
	}

	if err = DB.AutoMigrate(&Comment{}); err != nil {
		zapcomment.Logger.Panic("gorm comments initialization err: " + err.Error())
	}

	if err = DB.AutoMigrate(&CommentIndex{}); err != nil {
		zapcomment.Logger.Panic("gorm commentindeies initialization err: " + err.Error())
	}
	zapcomment.Logger.Info("mysql initialization succeeded")
}
