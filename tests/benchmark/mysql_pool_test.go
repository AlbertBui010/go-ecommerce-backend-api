package benchmark

import (
	"log"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID   uint
	Name string
}

func insertRecord(b *testing.B, db *gorm.DB) {
	user := User{Name: "Albert"}
	if err := db.Create(&user).Error; err != nil {
		b.Fatal(err)
	}
}

// go test -bench=. -benchmem

// cpu: Apple M2 Pro
// BenchmarkOpenConns1-10              1387        827965 ns/op     5851 B/op         77 allocs/op
func BenchmarkOpenConns1(b *testing.B) {
	dsn := "root:root1234@tcp(127.0.0.1:3306)/shopdevgo?charset=utf8mb4&parseTime=True"

	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Fatalf("failed to conect database: %v", err)
	}

	if db.Migrator().HasTable(&User{}) {
		if err := db.Migrator().DropTable(&User{}); err != nil {

		}
	}

	db.AutoMigrate(&User{})
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get sql.DB from gorm.DB: %v", err)
	}

	sqlDB.SetMaxOpenConns(1)
	defer sqlDB.Close()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			insertRecord(b, db)
		}
	})
}

// cpu: Apple M2 Pro
// BenchmarkOpenConns10-10             3843        324651 ns/op     5795 B/op         75 allocs/op

func BenchmarkOpenConns10(b *testing.B) {
	dsn := "root:root1234@tcp(127.0.0.1:3306)/shopdevgo?charset=utf8mb4&parseTime=True"

	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Fatalf("failed to conect database: %v", err)
	}

	if db.Migrator().HasTable(&User{}) {
		if err := db.Migrator().DropTable(&User{}); err != nil {

		}
	}

	db.AutoMigrate(&User{})
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get sql.DB from gorm.DB: %v", err)
	}

	sqlDB.SetMaxOpenConns(10)
	defer sqlDB.Close()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			insertRecord(b, db)
		}
	})
}
