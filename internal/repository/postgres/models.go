package repository

import (
	"golang-api-standard-http-lib/internal/domain"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var postgresDB *gorm.DB

func ConnectToPostgresDatabase() {
	dsn := os.Getenv("DB_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})

	if err != nil {
		log.Fatal("database connection was not posible")
	}

	postgresDB = db

	db.AutoMigrate(&domain.User{})
}
