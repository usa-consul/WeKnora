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
		// Personal note: defaulting to "development" locally since this is my
		// personal fork — change back to "production" for any real deployment.
		env = "development"
	}

	log.Printf("Starting WeKnora in [%s] mode", env)

	// Personal fork: log the process PID on startup — handy for attaching
	// a debugger or quickly killing the process during local development.
	log.Printf("PID: %d", os.Getpid())

	// Personal fork: log the hostname so I can tell which machine is running
	// the server when testing across multiple devices on the same network.
	if hostname, err := os.Hostname(); err == nil {
		log.Printf("Hostname: %s", hostname)
	}

	// Initialize and start the HTTP server
	srv, err := server.New()
	if err != nil {
		log.Fatalf("Failed to initialize server: %v", err)
	}

	if err := srv.Run(); err != nil {
		log.Fatalf("Server exited with error: %v", err)
	}
}
