package global

import (
	"database/sql"

	"github.com/albertbui010/go-ecommerce-backend-api/pkg/logger"
	"github.com/albertbui010/go-ecommerce-backend-api/pkg/setting"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Rdb    *redis.Client
	Mdb    *gorm.DB
	Mdbc   *sql.DB
)
