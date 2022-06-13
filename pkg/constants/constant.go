package constants

/**
 * @Description
 * @Author 拥抱漏风
 * @Date 2022/5/20 16:09
 **/

const (
	NoteTableName              = "note"
	UserTableName              = "user"
	SecretKey                  = "secret key"
	IdentityKey                = "id"
	Total                      = "total"
	Notes                      = "notes"
	NoteID                     = "note_id"
	ApiServiceName             = "api"
	VideoServiceName           = "video"
	UserServiceName            = "user"
	CommentServiceName         = "Comment"
	MySQLDefaultDSN            = "lzq:lzq@tcp(localhost:9910)/douyin?charset=utf8&parseTime=True&loc=Local"
	EtcdAddress                = "127.0.0.1:2379"
	RedisAddress               = "localhost:6379"
	CPURateLimit       float64 = 80.0
	LocalIP                    = "10.9.42.122"
	DefaultLimit               = 10
)
