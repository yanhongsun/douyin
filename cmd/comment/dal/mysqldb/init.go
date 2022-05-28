package mysqldb

import (
	"douyin/cmd/comment/configdata"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
