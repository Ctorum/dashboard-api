package infra

import (
	"database/sql"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var DBInterface *sql.DB

func InitDatabase() *gorm.DB {

	var err error
	dsn := os.Getenv("DATABASE_DSN")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}
	DBInterface, err = DB.DB()
	if err != nil {
		log.Fatal("Failed to connect to database")
	}
	DBInterface.SetMaxIdleConns(10)  // Pooling Config for idle connections
	DBInterface.SetMaxOpenConns(100) // Pooling Config for open connections

	return DB
}
