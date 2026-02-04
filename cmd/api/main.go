package main

import (
	httpHandler "golang-api-standard-http-lib/internal/delivery/http"
	repository "golang-api-standard-http-lib/internal/repository/postgres"
	"golang-api-standard-http-lib/internal/usecase"
	"golang-api-standard-http-lib/pkg/config"
	"log/slog"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lmittmann/tint"
)

func main() {

	// Create the tint handler for the colors logger
	tintHandler := tint.NewHandler(os.Stdout, &tint.Options{
		Level:      slog.LevelDebug, // Set the minimum log level
		TimeFormat: time.Kitchen,    // Use a shorter time format (e.g., "3:04PM")
		AddSource:  true,            // Shows the file and line number
	})

	logger := slog.New(tintHandler)
	slog.SetDefault(logger)

	// Postgres DB connection setup
	postgresConnectionURL := os.Getenv("POSTGRES_CONNECT_URL")
	repository.ConnectPostgresDB(postgresConnectionURL)

	// load .env file
	config.LoadEnv()

	// setup Gin's application mode (debug | release)
	ginMode := os.Getenv("GIN_MODE")
	gin.SetMode(ginMode)

	// setup Gin router
	r := gin.Default()

	// Create dependencies
	userService := usecase.NewUserService()
	userHandler := httpHandler.NewUserHandler(userService)

	// Map routes
	httpHandler.MapRoutes(r, userHandler)

	port := ":8000"
	slog.Info("Server started successfully", "port", port)

	r.Run(port)
}
