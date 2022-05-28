package configdata

import (
	"douyin/pkg/config"
)

type MysqlDatabase struct {
	User      string
	Password  string
	Host      string
	Name      string
	Charset   string
	ParseTime string
}

type CommentServer struct {
	EtcdHost        string
	CommentServName string
	CommentServHost string
}

type RedisDatabase struct {
	Host     string
	Password string
}

type Kafka struct {
	Host               string
	TopicComments      string
	TopicCommentNumber string
}

var (
	MysqlDatabaseConfig *MysqlDatabase
	CommentServerConfig *CommentServer
	RedisDatabaseConfig *RedisDatabase
	KafkaConfig         *Kafka
)

func SetupSetting() error {
	setting, err := config.NewSetting()
	if err != nil {
		return err
	}

	err = setting.ReadSection("MysqlDatabase", &MysqlDatabaseConfig)
	if err != nil {
		return err
	}

	err = setting.ReadSection("CommentServer", &CommentServerConfig)
	if err != nil {
		return err
	}

	err = setting.ReadSection("RedisDatabase", &RedisDatabaseConfig)
	if err != nil {
		return err
	}

	err = setting.ReadSection("Kafka", &KafkaConfig)
	if err != nil {
		return err
	}

	return err
}
