package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/skus-finder-psql/internal/core/usecases"
	"github.com/skus-finder-psql/internal/infrastructure/dependencies"
	"net/http"
)

type DeleteOneProductHandler struct {
	uc *usecases.EliminatorProductsBySKU
}

func NewDeleteOneProductHandler(container dependencies.Container) *DeleteOneProductHandler {
	return &DeleteOneProductHandler{
		uc: usecases.NewEliminatorProductsBySKU(container.ProductsRepository()),
	}
}

func (handler *DeleteOneProductHandler) DeleteOneProduct(ctx *gin.Context) {
	productSKU := ctx.Param("productSKU")
	if len(productSKU) == 0 {
		formatResponse(ctx, http.StatusBadRequest, "product sku not provided", nil)
		return
	}
	ok, err := handler.uc.Execute(ctx, productSKU)
	if err != nil {
		formatResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	formatResponse(ctx, http.StatusOK, "ok", ok)
}
