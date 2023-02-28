package initialize

import (
	"douyin/cmd/api/global"
	"douyin/cmd/api/service"
)

func Init() {
	initConfig()
	initRPC()
	initMinio()
	global.UploadService = service.NewUpload(global.MinioClient, global.MinioConfig)
}
