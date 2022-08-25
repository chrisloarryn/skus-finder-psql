package main

import (
	"fmt"
	_ "github.com/gin-gonic/gin"
	"github.com/skus-finder-psql/internal/http/server"
	"github.com/skus-finder-psql/internal/infrastructure/dependencies"
	"os"
)

func main() {
	fmt.Println("Running...")

	err := os.Setenv("ENVIRONMENT", "PRODUCTION")
	if err != nil {
		return
	}

	container := dependencies.NewContainer()

	server.Run(container)
}
