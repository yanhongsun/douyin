package constants

const (
	EmptyUserId                 = -1
	MaxTime                     = 9223372036854775807
	VideoResourceIpPort         = "http://219.216.86.30:8086"
	VideoSavePath               = "../../cmd/api/resource/videos"
	VideoCoverSavePath          = "../../cmd/api/resource/cover"
	VideoUrlPath                = "resource/videos"
	VideoCoverUrlPath           = "resource/cover"
	FavoriteTableName           = "favorite"
	RelationTableName           = "relation"
	VideoLimitNum               = 30
	UserTableName               = "users"
	VideoTableName              = "videos"
	UserServiceName             = "user"
	VideoServiceName            = "video"
	MySQLDefaultDSN             = "root:111111@tcp(localhost:3306)/douyin?charset=utf8&parseTime=True&loc=Local"
	EtcdAddress                 = "127.0.0.1:2379"
	CPURateLimit        float64 = 80.0
	DefaultLimit                = 10
	RelationServiceName         = "relation"
	UserServHost                = "127.0.0.1:8889"
	// JWT
	Secret = "this is a secret"
	Issuer = "douyin"
	Expire = 7200
)
