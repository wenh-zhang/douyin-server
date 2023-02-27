package initialize

import (
	"douyin/cmd/api/config"
	"douyin/cmd/api/global"
	"douyin/shared/constant"
)

func InitConfig(){
	global.EtcdConfig = &config.EtcdConfig{
		Host: constant.EtcdHost,
		Port: constant.EtcdPort,
	}
}


