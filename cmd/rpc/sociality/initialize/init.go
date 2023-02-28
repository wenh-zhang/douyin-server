package initialize

func Init() {
	InitConfig()
	InitDB()
	initRedis()
}
