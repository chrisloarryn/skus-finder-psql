package main

import (
	"fmt"
	_ "github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/skus-finder-psql/internal/http/server"
	"github.com/skus-finder-psql/internal/infrastructure/dependencies"
	"os"
)

func main() {
	fmt.Println("Running...")

	// load environment variables from .env file if environment is not set
	if os.Getenv(dependencies.EnvironmentKey) == "" {
		godotenv.Load()
	}

	container := dependencies.NewContainer()

	server.Run(container)
}
