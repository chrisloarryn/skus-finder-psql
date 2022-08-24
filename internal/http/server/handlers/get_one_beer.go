package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/skus-finder-psql/internal/core/usecases"
	"github.com/skus-finder-psql/internal/infrastructure/dependencies"
	"net/http"
)

type FindOneProductHandler struct {
	uc *usecases.FinderProductsBySKU
}

func NewFindOneProductHandler(container dependencies.Container) *FindOneProductHandler {
	return &FindOneProductHandler{
		uc: usecases.NewFinderProductsBySKU(container.ProductsRepository()),
	}
}

func (handler *FindOneProductHandler) FindOneProduct(ctx *gin.Context) {
	productSKU := ctx.Param("productSKU")
	if len(productSKU) == 0 {
		formatResponse(ctx, http.StatusBadRequest, "product sku not provided", nil)
		return
	}
	product, err := handler.uc.Execute(ctx, productSKU)
	if err != nil {
		formatResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	formatResponse(ctx, http.StatusOK, "ok", product)
}
