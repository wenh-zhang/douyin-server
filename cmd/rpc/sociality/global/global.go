package global

import (
	"douyin/cmd/rpc/sociality/config"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
	RedisFollowClient *redis.Client

	EtcdConfig *config.EtcdConfig
	MySQLConfig *config.MySQLConfig
	RedisConfig *config.RedisConfig
	RPCConfig *config.RPCConfig
)
