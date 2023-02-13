// constants of the whole project
// configure before running the project

package constant

import "time"

var TokenSignedKey = []byte("golang")

// server
const (
	CoreRPCAddress      = "127.0.0.1:8885"
	InteractRPCAddress  = "127.0.0.1:8886"
	SocialRPCAddress    = "127.0.0.1:8887"
	APIServerAddress    = "127.0.0.1:8888"
	CoreServiceName     = "core service"
	InteractServiceName = "interact service"
	SocialServiceName   = "social service"
)

// etcd
const EtcdAddress = "127.0.0.1:2379"

// ServerAddress set according to your host ip address
// configure in the application
const ServerAddress = "192.168.10.249:8888"

const (
	FeedLimit       = 30
	TokenExpireTime = 24 * time.Hour
)

// param name
const (
	LatestTime               = "latest_ime"
	Token                    = "token"
	TokenUserIdentifyKey     = "token_user_id" //use to save user_id taken from token in request context
	UserName                 = "username"
	Password                 = "password"
	UserIdentityKey          = "user_id"
	VideoIdentityKey         = "video_id"
	ActionType               = "action_type"
	ActionTypeFavorite       = 1
	ActionTypeCancelFavorite = 2
	ActionTypeComment        = 1
	ActionTypeDeleteComment  = 2
	Data                     = "data"
	Title                    = "title"
	CommentText              = "comment_text"
	CommentIdentityKey       = "comment_id"
	ToUserIdentityKey        = "to_user_id"
	Content                  = "content" //message的内容
)

// mysql
const (
	MySQLDefaultDSN    = "gorm:gorm@tcp(localhost:9910)/gorm?charset=utf8&parseTime=True&loc=Local"
	TimestampFormatStr = "2006-01-02 15:04:05"
	UserTableName      = "user"
	VideoTableName     = "video"
	FavoriteTableName  = "favorite"
	CommentTableName   = "comment"
	FollowTableName    = "follow"
	MessageTableName   = "message"
)
