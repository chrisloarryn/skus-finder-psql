package server

import (
	"github.com/gin-gonic/gin"
	"github.com/skus-finder-psql/internal/http/server/handlers"
	"github.com/skus-finder-psql/internal/infrastructure/dependencies"
	"net/http"
)

type ServerHTTP struct {
}

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

	r.Run(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
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
