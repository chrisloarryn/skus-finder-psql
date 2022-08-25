package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/skus-finder-psql/internal/core/domain/products"
	"github.com/skus-finder-psql/internal/core/usecases"
	"github.com/skus-finder-psql/internal/infrastructure/dependencies"
	"net/http"
)

type UpdateProductHandler struct {
	uc *usecases.UpdateProduct
}

func NewUpdateProductHandler(container dependencies.Container) *UpdateProductHandler {
	return &UpdateProductHandler{
		uc: usecases.NewUpdateProduct(container.ProductsRepository()),
	}
}

func (handler *UpdateProductHandler) UpdateProduct(ctx *gin.Context) {
	product := products.Product{}
	if err := ctx.BindJSON(&product); err != nil {
		formatResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	resProd, err := handler.uc.Execute(ctx, product)
	if err != nil {
		formatResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	formatResponse(ctx, http.StatusOK, "ok", resProd)
}
