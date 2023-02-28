package initialize

import (
	"douyin/cmd/rpc/sociality/global"
	"douyin/shared/constant"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func initRedis() {
	config := global.RedisConfig
	addr := fmt.Sprintf("%s:%d", config.Host, config.Port)
	global.RedisFollowClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: config.Password,
		DB:       constant.RedisFollowDB,
	})
}
