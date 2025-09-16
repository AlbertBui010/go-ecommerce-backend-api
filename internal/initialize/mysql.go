package initialize

import (
	"fmt"
	"time"

	"github.com/albertbui010/go-ecommerce-backend-api/global"
	"github.com/albertbui010/go-ecommerce-backend-api/internal/po"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func checkErrorPanic(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(err)
	}
}

func InitMysql() {
	m := global.Config.MySql
	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	var s = fmt.Sprintf(
		dsn,
		m.Username,
		m.Password,
		m.Host,
		m.Port,
		m.Dbname,
	)

	db, err := gorm.Open(mysql.Open(s), &gorm.Config{
		SkipDefaultTransaction: false,
	})
	checkErrorPanic(err, "Init MySQL initialization error")
	global.Logger.Info("Initializing MySQL Successfully")
	global.Mdb = db

	// Set pool
	SetPool()
	migrateTables()
}

func SetPool() {
	m := global.Config.MySql
	sqlDB, err := global.Mdb.DB()
	if err != nil {
		fmt.Printf("MySQL error::%s", err)
	}
	sqlDB.SetConnMaxIdleTime(time.Duration(m.MaxIdleConns))
	sqlDB.SetMaxOpenConns(m.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime))
}

func migrateTables() {
	err := global.Mdb.AutoMigrate(
		&po.Role{},
		&po.User{},
	)
	if err != nil {
		fmt.Println("Migrating table failed:", err)
	}
}
