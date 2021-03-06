package configdata

import (
	"douyin/cmd/comment/pack/zapcomment"

	"github.com/spf13/viper"
)

type Setting struct {
	vp *viper.Viper
}

func NewSetting() (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("commentconfig")
	vp.AddConfigPath("./config/")
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return &Setting{vp}, nil
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	return nil
}

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
	UserServName    string
	VideoServName   string
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

type TencentCloud struct {
	TextModeration string
	SecretId       string
	SecretKey      string
}

var (
	MysqlDatabaseConfig *MysqlDatabase
	CommentServerConfig *CommentServer
	RedisDatabaseConfig *RedisDatabase
	KafkaConfig         *Kafka
	TencentCloudConfig  *TencentCloud
)

func SetupSetting() error {
	setting, err := NewSetting()
	if err != nil {
		zapcomment.Logger.Error("reading config err: " + err.Error())
		return err
	}

	err = setting.ReadSection("MysqlDatabase", &MysqlDatabaseConfig)
	if err != nil {
		zapcomment.Logger.Error("reading mysqlConfig err: " + err.Error())
		return err
	}

	err = setting.ReadSection("CommentServer", &CommentServerConfig)
	if err != nil {
		zapcomment.Logger.Error("reading commentServerConfig err: " + err.Error())
		return err
	}

	err = setting.ReadSection("RedisDatabase", &RedisDatabaseConfig)
	if err != nil {
		zapcomment.Logger.Error("reading redisConfig err: " + err.Error())
		return err
	}

	err = setting.ReadSection("Kafka", &KafkaConfig)
	if err != nil {
		zapcomment.Logger.Error("reading kafkaConfig err: " + err.Error())
		return err
	}

	err = setting.ReadSection("TencentCloud", &TencentCloudConfig)
	if err != nil {
		zapcomment.Logger.Error("reading tencentConfig err: " + err.Error())
		return err
	}

	zapcomment.Logger.Info("configuration initialization succeeded")
	return nil
}
