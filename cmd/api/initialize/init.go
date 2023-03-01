package initialize

import (
	"douyin/cmd/api/global"
	"douyin/cmd/api/service"
)

func Init() {
	initConfig()
	initRPC()
	initMinio()
	initAmqp()
	global.UploadService = service.NewUpload(global.MinioClient, global.MinioConfig, global.AmqpConn)
}
