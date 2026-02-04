package config

import (
	"log/slog"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()

	if err != nil {
		slog.Error("No .env file found, falling back to system environment variables\n")
		return
	}

	slog.Info(".env loaded successfully\n")
}
