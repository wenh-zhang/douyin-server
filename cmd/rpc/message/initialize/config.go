package initialize

import (
	"douyin/cmd/rpc/message/config"
	"douyin/cmd/rpc/message/global"
	"douyin/shared/constant"
)

func InitConfig() {
	global.EtcdConfig = &config.EtcdConfig{
		Host: constant.EtcdHost,
		Port: constant.EtcdPort,
	}
	global.MySQLConfig = &config.MySQLConfig{
		Host:     constant.MySQLHost,
		Port:     constant.MySQLPort,
		Database: constant.MySQLDatabase,
		User:     constant.MySQLUser,
		Password: constant.MySQLPassword,
	}
	global.RPCConfig = &config.RPCConfig{
		Host: constant.RPCMessageHost,
		Port: constant.RPCMessagePort,
		Name: constant.RPCMessageName,
	}
	global.AmqpConfig = &config.AmqpConfig{
		Host:     constant.AmqpHost,
		Port:     constant.AmqpPort,
		User:     constant.AmqpUser,
		Password: constant.AmqpPassword,
	}
}
