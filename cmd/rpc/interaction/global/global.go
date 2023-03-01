package global

import (
	"douyin/cmd/rpc/interaction/config"
	user "douyin/shared/kitex_gen/user/userservice"
	video "douyin/shared/kitex_gen/video/videoservice"
	"github.com/go-redis/redis/v8"
	"github.com/streadway/amqp"
	"gorm.io/gorm"
)

var (
	DB                  *gorm.DB
	RedisFavoriteClient *redis.Client
	RedisCommentClient  *redis.Client
	UserClient          user.Client
	VideoClient         video.Client
	AmqpConn            *amqp.Connection

	EtcdConfig  *config.EtcdConfig
	MySQLConfig *config.MySQLConfig
	RedisConfig *config.RedisConfig
	RPCConfig   *config.RPCConfig
	AmqpConfig  *config.AmqpConfig
)
