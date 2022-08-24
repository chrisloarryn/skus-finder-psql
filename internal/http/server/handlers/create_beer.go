package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/skus-finder-psql/internal/core/domain/products"
	"github.com/skus-finder-psql/internal/core/usecases"
	"github.com/skus-finder-psql/internal/infrastructure/dependencies"
	"net/http"
)

type CreateProductHandler struct {
	uc *usecases.CreateProduct
}

func NewCreateProductHandler(container dependencies.Container) *CreateProductHandler {
	return &CreateProductHandler{
		uc: usecases.NewCreateProduct(container.ProductsRepository()),
	}
}

func (handler *CreateProductHandler) CreateProduct(ctx *gin.Context) {
	product := products.Product{}
	if err := ctx.BindJSON(&product); err != nil {
		formatResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	err := handler.uc.Execute(ctx, product)
	if err != nil {
		formatResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	formatResponse(ctx, http.StatusOK, "ok", product)
}
