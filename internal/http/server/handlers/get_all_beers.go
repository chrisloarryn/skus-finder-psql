package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/skus-finder-psql/internal/core/usecases"
	"github.com/skus-finder-psql/internal/infrastructure/dependencies"
	"net/http"
)

type FindAllProductsHandler struct {
	uc *usecases.FinderAllProducts
}

func NewFindAllProductsHandler(container dependencies.Container) *FindAllProductsHandler {
	return &FindAllProductsHandler{
		uc: usecases.NewFinderAllProducts(container.ProductsRepository()),
	}
}

func (handler *FindAllProductsHandler) GetAllProducts(ctx *gin.Context) {
	products, err := handler.uc.Execute(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, products)

}

// Error getting response;java.net UnknownHostException: middlwareuat.falabella.cl
