package initialize

func Init() {
	InitConfig()
	InitDB()
	initAmqp()
}
