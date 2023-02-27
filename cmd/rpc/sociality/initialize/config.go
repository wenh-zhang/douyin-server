package initialize

import (
	"douyin/cmd/rpc/sociality/config"
	"douyin/cmd/rpc/sociality/global"
	"douyin/shared/constant"
)
func InitConfig(){
	global.EtcdConfig = &config.EtcdConfig{
		Host: constant.EtcdHost,
		Port: constant.EtcdPort,
	}
	global.MySQLConfig = &config.MySQLConfig{
		Host: constant.MySQLHost,
		Port: constant.MySQLPort,
		Database: constant.MySQLDatabase,
		User: constant.MySQLUser,
		Password: constant.MySQLPassword,
	}
	global.RPCConfig = &config.RPCConfig{
		Host: constant.RPCSocialityHost,
		Port: constant.RPCSocialityPort,
		Name: constant.RPCSocialityeName,
	}
}

