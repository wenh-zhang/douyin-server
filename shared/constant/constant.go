// constants of the whole project
// configure before running the project

package constant

import "time"


// server
const (
	RPCInteractionHost = "127.0.0.1"
	RPCInteractionPort = 8881
	RPCInteractionName = "interaction service"
	RPCSocialityHost   = "127.0.0.1"
	RPCSocialityPort   = 8882
	RPCSocialityeName  = "sociality service"
	RPCUserHost        = "127.0.0.1"
	RPCUserPort        = 8883
	RPCUserName        = "user service"
	RPCVideoHost       = "127.0.0.1"
	RPCVideoPort       = 8884
	RPCVideoName       = "video service"
	RPCMessageHost     = "127.0.0.1"
	RPCMessagePort     = 8885
	RPCMessageName     = "Message service"
)

// etcd
const (
	EtcdHost = "127.0.0.1"
	EtcdPort = 2379
)

const (
	MinioPublicHost = "192.168.10.249"
	MinioHost        = "127.0.0.1"
	MinioPort = 9000
	MinioAccessKeyID     = "minio"
	MinioSecretAccessKey = "minio123"
	MinioBucket          = "douyin"
	MinioUseSSL          = false
)

// mysql
const (
	MySQLDefaultDSN   = "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	MySQLHost         = "127.0.0.1"
	MySQLPort         = 9910
	MySQLDatabase     = "douyin"
	MySQLUser         = "gorm"
	MySQLPassword     = "gorm"
	UserTableName     = "user"
	VideoTableName    = "video"
	FavoriteTableName = "favorite"
	CommentTableName  = "comment"
	FollowTableName   = "follow"
	MessageTableName  = "message"
)

// redis
const(
	RedisHost = "127.0.0.1"
	RedisPort = 6379
	RedisPassword = "douyin"
	RedisFavoriteDB = 0
	RedisCommentDB = 1
	RedisFollowDB = 2
)

const (
	TokenSignedKey = "golang"
	FeedLimit       = 30
	TokenExpireTime = 30 * 24 * time.Hour
)

// param name
const (
	Token                = "token"
	TokenUserIdentifyKey = "token_user_id" //use to save user_id taken from token in request context

	ActionTypeFavorite       = 1
	ActionTypeCancelFavorite = 2
	ActionTypeComment        = 1
	ActionTypeDeleteComment  = 2
	ActionTypeFollow         = 1
	ActionTypeCancelFollow   = 2
	OptionFollow             = 0
	OptionFollower           = 1
	OptionFriend             = 2
	MsgTypeReceive           = 0
	MsgTypeSend              = 1
)

const (
	SnowFlakeDataCenterId    = 1
	MinioSnowFlakeWorkerId   = 1
	UserSnowFlakeWorkerId    = 2
	CommentSnowFlakeWorkerId = 3
	VideoSnowFlakeWorkerId   = 4
)


// ServerAddress set according to your host ip address
// configure in the application
const (
	ServerHost = "192.168.10.249"
	ServerPort = 8888
)