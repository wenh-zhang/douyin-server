package initialize

import (
	"douyin/cmd/api/config"
	"douyin/cmd/api/global"
	"douyin/shared/constant"
)

func initConfig() {
	global.EtcdConfig = &config.EtcdConfig{
		Host: constant.EtcdHost,
		Port: constant.EtcdPort,
	}
	global.MinioConfig = &config.MinioConfig{
		Host: constant.MinioHost,
		Port: constant.MinioPort,
		AccessKeyID:     constant.MinioAccessKeyID,
		SecretAccessKey: constant.MinioSecretAccessKey,
		Bucket:          constant.MinioBucket,
		UserSSL:         constant.MinioUseSSL,
	}
}
