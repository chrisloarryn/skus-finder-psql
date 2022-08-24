package handlers

import (
	"fmt"
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
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Invalid product SKU: %s", productSKU))
		return
	}
	product, err := handler.uc.Execute(ctx, productSKU)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, product)

}
