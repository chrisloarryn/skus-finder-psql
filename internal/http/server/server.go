package server

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/skus-finder-psql/internal/http/server/handlers"
	"github.com/skus-finder-psql/internal/infrastructure/dependencies"
	sharedconstants "github.com/skus-finder-psql/internal/shared/constants"
)

type ServerHTTP struct{}

func Run(container dependencies.Container) {
	r := gin.Default()

	r.GET(sharedconstants.HTTPRootPath, pingpong)
	r.GET(sharedconstants.HTTPPingPath, pingpong)

	api := r.Group(sharedconstants.HTTPAPIPath)
	v1 := api.Group(sharedconstants.HTTPAPIV1Path)

	findAllProductsHandler := handlers.NewFindAllProductsHandler(container)
	createProductHandler := handlers.NewCreateProductHandler(container)
	getOneProductHandler := handlers.NewFindOneProductHandler(container)
	updateProductHandler := handlers.NewUpdateProductHandler(container)
	deleteOneProductHandler := handlers.NewDeleteOneProductHandler(container)

	v1.GET(sharedconstants.HTTPProductsPath, findAllProductsHandler.GetAllProducts)
	v1.POST(sharedconstants.HTTPProductsPath, createProductHandler.CreateProduct)
	v1.GET(sharedconstants.HTTPProductSKUPath, getOneProductHandler.FindOneProduct)
	v1.PATCH(sharedconstants.HTTPProductSKUPath, updateProductHandler.UpdateProduct)
	v1.DELETE(sharedconstants.HTTPProductSKUPath, deleteOneProductHandler.DeleteOneProduct)

	port := os.Getenv(sharedconstants.HTTPPortKey)

	if len(port) == 0 {
		port = sharedconstants.HTTPDefaultPort
	}

	if err := r.Run(sharedconstants.HTTPListenAddrPrefix + port); err != nil {
		log.Fatalf("start HTTP server: %v", err)
	}
}

func pingpong(c *gin.Context) {
	formatResponse(c, http.StatusOK, sharedconstants.HTTPPongMessage, nil)
}

func formatResponse(ctx *gin.Context, sc int, msg string, data interface{}) {
	ctx.JSON(sc, gin.H{
		"message": msg,
		"data":    data,
	})
}
