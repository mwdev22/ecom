package dbConn

import (
	"fmt"
	"log"

	"github.com/mwdev22/ecom/app/routes/auth"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DbOpen(connString string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	fmt.Println("Connected to PostgreSQL database...")
	return db, nil
}

func InitConn(db *gorm.DB) {
	sqlDB, err := db.DB()
	db.AutoMigrate(models...)
	if err != nil {
		log.Fatal("Failed to get DB from GORM:", err)
	}

	err = sqlDB.Ping()
	if err != nil {
		sqlDB.Close()
		log.Fatal(err)
	}
}

var models = []interface{}{&auth.User{}}
