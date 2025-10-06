package initialize

import (
	"fmt"
	"time"

	"github.com/albertbui010/go-ecommerce-backend-api/global"
	"github.com/albertbui010/go-ecommerce-backend-api/internal/model"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
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
	// genTableDAO()
	migrateTables()
}

func genTableDAO() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./internal/model",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	// // gormdb, _ := gorm.Open(mysql.Open("root:@(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"))
	g.UseDB(global.Mdb) // reuse your gorm db

	// // Generate basic type-safe DAO API for struct `model.User` following conventions
	// g.ApplyBasic(model.User{})

	g.GenerateModel("go_crm_user")

	// // Generate Type Safe API with Dynamic SQL defined on Querier interface for `model.User` and `model.Company`
	// g.ApplyInterface(func(Querier) {}, model.User{}, model.Company{})

	// Generate the code
	g.Execute()
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
		// &po.Role{},
		// &po.User{},
		// &model.GoCrmUserV2{},
		&model.GoCrmUser{},
	)
	if err != nil {
		fmt.Println("Migrating table failed:", err)
	}
}
