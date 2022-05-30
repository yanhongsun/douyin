package constants

const (
	EmptyUserId                = -1
	MaxTime                    = 9223372036854775807
	VideoSavePath              = "resource/videos"
	VideoCoverSavePath         = "resource/cover"
	FavoriteTableName          = "favorite"
	RelationTableName          = "relation"
	VideoLimitNum              = 30
	UserTableName              = "users"
	VideoTableName             = "videos"
	VideoServiceName           = "video"
	MySQLDefaultDSN            = "root:111111@tcp(localhost:3306)/douyin?charset=utf8&parseTime=True&loc=Local"
	EtcdAddress                = "127.0.0.1:2379"
	CPURateLimit       float64 = 80.0
	DefaultLimit               = 10
)
