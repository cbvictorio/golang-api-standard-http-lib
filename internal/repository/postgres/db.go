package repository

import (
	"log/slog"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var PostgresClient *gorm.DB

func ConnectPostgresDB(dsn string) error {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		slog.Error("Postgres Database connection failed")
		return err
	}

	PostgresClient = db
	slog.Info("Postgres Database connection succesful")

	return nil
}
