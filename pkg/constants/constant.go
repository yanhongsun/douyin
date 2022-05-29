package constants

const (
	//表名
	VideoTableName = "videos"
	UserTableName  = "users"

	SecretKey   = "secret key"
	IdentityKey = "id"

	//api调用参数字段
	//ViedoKey         = "video_id"
	//ActionType       = "action_type"
	Total            = "total"
	Notes            = "notes"
	NoteID           = "note_id"
	ApiServiceName   = "demoapi"
	ThumbServiceName = "like"
	MySQLDefaultDSN  = "localhost:@tcp(localhost:9910)/douyin?charset=utf8&parseTime=True&loc=Local"
	EtcdAddress      = "127.0.0.1:2379"

	CPURateLimit float64 = 80.0
	DefaultLimit         = 10
)
