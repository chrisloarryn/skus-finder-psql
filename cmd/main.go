package main

import (
	"fmt"
	_ "github.com/gin-gonic/gin"
	"github.com/skus-finder-psql/internal/http/server"
	"github.com/skus-finder-psql/internal/infrastructure/dependencies"
)

func main() {
	fmt.Println("Running...")

	container := dependencies.NewContainer()

	server.Run(container)
}
