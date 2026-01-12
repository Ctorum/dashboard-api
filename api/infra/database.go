package infra

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() *gorm.DB {

	var err error
	dsn := os.Getenv("DATABASE_DSN")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	log.Print("Connected to database")
	if err != nil {
		log.Fatal("Failed to connect to database")
	}
	return DB
}
