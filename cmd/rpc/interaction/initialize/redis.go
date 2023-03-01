package initialize

import (
	"douyin/cmd/rpc/interaction/global"
	"douyin/shared/constant"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func initRedis() {
	config := global.RedisConfig
	addr := fmt.Sprintf("%s:%d", config.Host, config.Port)
	global.RedisFavoriteClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: config.Password,
		DB:       constant.RedisFavoriteDB,
	})
	global.RedisCommentClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: config.Password,
		DB:       constant.RedisCommentDB,
	})
}
