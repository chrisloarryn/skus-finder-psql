package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/skus-finder-psql/internal/http/server"
	"github.com/skus-finder-psql/internal/infrastructure/dependencies"
	"log"
	"os"
)

func main() {
	fmt.Println("Running...")

	// load environment variables from .env file if environment is not set
	if os.Getenv(dependencies.EnvironmentKey) == "" {
		if err := godotenv.Load(); err != nil {
			log.Printf("warning: unable to load .env: %v", err)
		}
	}

	container := dependencies.NewContainer()

	server.Run(container)
}
