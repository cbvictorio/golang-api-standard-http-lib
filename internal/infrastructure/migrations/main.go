package main

import (
	"golang-api-standard-http-lib/internal/domain"
	postgresRepository "golang-api-standard-http-lib/internal/repository/postgres"
	"golang-api-standard-http-lib/pkg/config"
	"log/slog"
	"os"
)

func main() {
	config.LoadEnv()
	postgresConnectionURL := os.Getenv("POSTGRES_CONNECT_URL")
	postgresClient, err := postgresRepository.ConnectPostgresDB(postgresConnectionURL)

	if err != nil {
		slog.Error("The Postgres database connection failed", "error", err)
		os.Exit(1)
	}

	postgresClient.Migrator().AutoMigrate(&domain.User{})
	slog.Info("Migrations executed successfully")
}
