package server

import (
	"github.com/gin-gonic/gin"
	"github.com/skus-finder-psql/internal/http/server/handlers"
	"github.com/skus-finder-psql/internal/infrastructure/dependencies"
	"net/http"
	"os"
)

type ServerHTTP struct {
}

const (
	PORT = "PORT"
)

func Run(container dependencies.Container) {
	r := gin.Default()

	r.GET("/", pingpong)
	r.GET("/ping", pingpong)

	api := r.Group("/api")
	v1 := api.Group("/v1")

	findAllProductsHandler := handlers.NewFindAllProductsHandler(container)
	createProductHandler := handlers.NewCreateProductHandler(container)
	getOneProductHandler := handlers.NewFindOneProductHandler(container)
	updateProductHandler := handlers.NewUpdateProductHandler(container)
	deleteOneProductHandler := handlers.NewDeleteOneProductHandler(container)

	v1.GET("/products", findAllProductsHandler.GetAllProducts)
	v1.POST("/products", createProductHandler.CreateProduct)
	v1.GET("/products/:productSKU", getOneProductHandler.FindOneProduct)
	v1.PATCH("/products/:productSKU", updateProductHandler.UpdateProduct)
	v1.DELETE("/products/:productSKU", deleteOneProductHandler.DeleteOneProduct)

	port := os.Getenv(PORT)
	if len(port) == 0 {
		port = "8088"
	}
	r.Run(":" + port)
}

func pingpong(c *gin.Context) {
	formatResponse(c, http.StatusOK, "pong", nil)
}

func formatResponse(ctx *gin.Context, sc int, msg string, data interface{}) {
	ctx.JSON(sc, gin.H{
		"message": msg,
		"data":    data,
	})
}
