package initialize

import (
	"github.com/albertbui010/go-ecommerce-backend-api/global"
	"go.uber.org/zap"
)

func Run() {
	// Load config
	LoadConfig()
	InitLogger()
	global.Logger.Info("config ok", zap.String("ok", "success"))
	InitMysql()
	InitRedis()

	r := InitRouter()

	r.Run(":8081")
}
