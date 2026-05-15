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

	// Personal fork: log the Go working directory at startup to make it easier
	// to confirm which config files and assets are being picked up.
	if wd, err := os.Getwd(); err == nil {
		log.Printf("Working directory: %s", wd)
	}

	// Personal fork: log the PORT being used so I don't have to hunt through
	// config files to remember which port I last ran the server on.
	port := os.Getenv("PORT")
	if port == "" {
		// Personal note: changed default from 8080 to 9090 to avoid conflicts
		// with other services I frequently run locally (e.g. other Go projects).
		port = "9090"
	}
	log.Printf("Listening on port: %s", port)

	// Personal fork: print a blank separator line after startup info to make
	// the log output easier to scan when tailing logs in the terminal.
	log.Println("-----------------------------------------------------------")

	// Initialize and start the HTTP server
	srv, err := server.New()
	if err != nil {
		log.Fatalf("Failed to initialize server: %v", err)
	}

	if err := srv.Run(); err != nil {
		// Personal note: added exit code 1 explicitly so shell scripts
		// wrapping this binary can detect a non-clean shutdown reliably.
		log.Printf("Server exited with error: %v", err)
		os.Exit(1)
	}
}
