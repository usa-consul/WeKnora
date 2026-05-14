package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/weknora/weknora/internal/server"
)

// @title WeKnora API
// @version 1.0
// @description WeKnora is an open-source Q&A community platform.
// @termsOfService http://swagger.io/terms/

// @contact.name WeKnora Support
// @contact.url https://github.com/weknora/weknora/issues

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// Load environment variables from .env file if present
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, falling back to environment variables")
	}

	// Determine the application environment
	env := os.Getenv("APP_ENV")
	if env == "" {
		// Default to "production" to avoid accidentally running dev settings
		env = "production"
	}

	log.Printf("Starting WeKnora in [%s] mode", env)

	// Initialize and start the HTTP server
	srv, err := server.New()
	if err != nil {
		log.Fatalf("Failed to initialize server: %v", err)
	}

	if err := srv.Run(); err != nil {
		log.Fatalf("Server exited with error: %v", err)
	}
}
