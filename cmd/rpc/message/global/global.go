package global

import (
	"douyin/cmd/rpc/message/config"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB

	EtcdConfig *config.EtcdConfig
	MySQLConfig *config.MySQLConfig
	RPCConfig *config.RPCConfig
)
