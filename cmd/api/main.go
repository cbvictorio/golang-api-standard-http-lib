package main

import (
	httpHandler "golang-api-standard-http-lib/internal/delivery/http"
	repository "golang-api-standard-http-lib/internal/repository/postgres"
	"golang-api-standard-http-lib/internal/usecase"
	"golang-api-standard-http-lib/pkg/config"
	"log/slog"
	"os"
	"time"

	"github.com/gin-contrib/cors"
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

	// load .env file
	config.LoadEnv()

	// Postgres DB connection setup
	postgresConnectionURL := os.Getenv("POSTGRES_CONNECT_URL")
	if err := repository.ConnectPostgresDB(postgresConnectionURL); err != nil {
		slog.Error("Failed to connect to database", "error", err)
		os.Exit(1)
	}

	// setup Gin's application mode (debug | release)
	ginMode := os.Getenv("GIN_MODE")
	gin.SetMode(ginMode)

	// setup Gin router
	router := gin.Default()

	// setup CORS config
	frontendAppUrl := os.Getenv("FRONTEND_APP_URL")

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{frontendAppUrl}, // Your Angular dev URL
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Create dependencies
	userService := usecase.NewUserService()
	userHandler := httpHandler.NewUserHandler(userService)

	// Map routes
	httpHandler.MapRoutes(router, userHandler)

	// run the app
	port := ":8000"
	slog.Info("Server started successfully", "port", port)

	router.Run(port)
}
