package constants

const (
	FeedLimit = 30
)

// param name
const (
	LatestTime         = "latest_ime"
	Token              = "token"
	UserName           = "username"
	Password           = "password"
	UserIdentityKey    = "user_id"
	VideoIdentityKey   = "video_id"
	ActionType         = "action_type"
	IdentityKey        = "id"
	Data               = "data"
	Title              = "title"
	CommentText        = "comment_text"
	CommentIdentityKey = "comment_id"
	ToUserIdentityKey  = "to_user_id"
	Content            = "content" //message的内容
)

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
