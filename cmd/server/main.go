package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/lordaris/gotth-boilerplate/internal/app"
	"github.com/lordaris/gotth-boilerplate/internal/handlers"
	"github.com/lordaris/gotth-boilerplate/internal/routes"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Warning: .env file not found, using environment variables")
	}

	// Setup configuration
	var cfg app.Config
	flag.StringVar(&cfg.DB.Host, "db-host", getEnvOrDefault("DB_HOST", "localhost"), "PostgreSQL host")
	flag.StringVar(&cfg.DB.Port, "db-port", getEnvOrDefault("DB_PORT", "5432"), "PostgreSQL port")
	flag.StringVar(&cfg.DB.User, "db-user", getEnvOrDefault("DB_USER", "postgres"), "PostgreSQL user")
	flag.StringVar(&cfg.DB.Password, "db-password", getEnvOrDefault("DB_PASSWORD", "postgres"), "PostgreSQL password")
	flag.StringVar(&cfg.DB.Name, "db-name", getEnvOrDefault("DB_NAME", "gotth-boilerplate"), "PostgreSQL database name")
	flag.StringVar(&cfg.Server.Port, "port", getEnvOrDefault("PORT", "8080"), "Server port")
	flag.Parse()

	// Initialize application
	application := app.NewApplication(cfg)

	// Connect to PostgreSQL
	if err := application.ConnectToDatabase(); err != nil {
		log.Fatal("Could not connect to PostgreSQL: ", err)
	}

	// Setup handlers
	handlers := handlers.NewHandlers(application)

	// Setup router
	router := routes.SetupRouter(handlers)

	// Start server
	port := cfg.Server.Port
	fmt.Printf("Server started at http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func getEnvOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
