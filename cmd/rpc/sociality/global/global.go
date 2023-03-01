package global

import (
	"douyin/cmd/rpc/sociality/config"
	"github.com/go-redis/redis/v8"
	"github.com/streadway/amqp"
	"gorm.io/gorm"
)

var (
	DB                *gorm.DB
	RedisFollowClient *redis.Client
	AmqpConn          *amqp.Connection

	EtcdConfig  *config.EtcdConfig
	MySQLConfig *config.MySQLConfig
	RedisConfig *config.RedisConfig
	RPCConfig   *config.RPCConfig
	AmqpConfig  *config.AmqpConfig
)
