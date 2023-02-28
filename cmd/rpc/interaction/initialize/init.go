package initialize

func Init() {
	initConfig()
	initDB()
	initRedis()
	initRPC()
}
