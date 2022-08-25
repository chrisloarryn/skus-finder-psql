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
		formatResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	if len(products) == 0 {
		formatResponse(ctx, http.StatusNoContent, "ok", nil)
	}

	formatResponse(ctx, http.StatusOK, "ok", products)

}

// Error getting response;java.net UnknownHostException: middlwareuat.falabella.cl
