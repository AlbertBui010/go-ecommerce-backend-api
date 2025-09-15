package initialize

func Run() {
	// Load config
	LoadConfig()
	InitLogger()
	InitMysql()
	InitRedis()

	r := InitRouter()

	r.Run(":8081")
}
