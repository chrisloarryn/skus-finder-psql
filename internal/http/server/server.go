package server

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/skus-finder-psql/internal/http/server/handlers"
	"github.com/skus-finder-psql/internal/infrastructure/dependencies"
)

type ServerHTTP struct{}

func Run(container dependencies.Container) {
	r := gin.Default()

	r.GET(rootPath, pingpong)
	r.GET(pingPath, pingpong)

	api := r.Group(apiPath)
	v1 := api.Group(apiV1Path)

	findAllProductsHandler := handlers.NewFindAllProductsHandler(container)
	createProductHandler := handlers.NewCreateProductHandler(container)
	getOneProductHandler := handlers.NewFindOneProductHandler(container)
	updateProductHandler := handlers.NewUpdateProductHandler(container)
	deleteOneProductHandler := handlers.NewDeleteOneProductHandler(container)

	v1.GET(productsPath, findAllProductsHandler.GetAllProducts)
	v1.POST(productsPath, createProductHandler.CreateProduct)
	v1.GET(productSKUPath, getOneProductHandler.FindOneProduct)
	v1.PATCH(productSKUPath, updateProductHandler.UpdateProduct)
	v1.DELETE(productSKUPath, deleteOneProductHandler.DeleteOneProduct)

	port := os.Getenv(portKey)

	if len(port) == 0 {
		port = defaultPort
	}

	if err := r.Run(listenAddrPrefx + port); err != nil {
		log.Fatalf("start HTTP server: %v", err)
	}
}

func pingpong(c *gin.Context) {
	formatResponse(c, http.StatusOK, pongMessage, nil)
}

func formatResponse(ctx *gin.Context, sc int, msg string, data interface{}) {
	ctx.JSON(sc, gin.H{
		"message": msg,
		"data":    data,
	})
}
