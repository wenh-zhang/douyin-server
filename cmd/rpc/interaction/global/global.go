package global

import (
	"douyin/cmd/rpc/interaction/config"
	user "douyin/shared/kitex_gen/user/userservice"
	video "douyin/shared/kitex_gen/video/videoservice"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
	UserClient user.Client
	VideoClient video.Client

	EtcdConfig *config.EtcdConfig
	MySQLConfig *config.MySQLConfig
	RPCConfig *config.RPCConfig
)
