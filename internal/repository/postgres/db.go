package repository

import (
	"log/slog"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresClient struct {
	*gorm.DB
}

func ConnectPostgresDB(dsn string) (*PostgresClient, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		slog.Error("Postgres Database connection failed")
		return nil, err
	}

	slog.Info("Postgres Database connection succesful")

	return &PostgresClient{
		DB: db,
	}, nil
}
