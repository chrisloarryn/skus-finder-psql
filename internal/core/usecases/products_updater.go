package usecases

import (
	"context"
	"github.com/skus-finder-psql/internal/core/domain/products"
)

// UpdateProduct is the use case that update a product
type UpdateProduct struct {
	productsRepository products.Repository
}

// NewUpdateProduct constructor
func NewUpdateProduct(repository products.Repository) *UpdateProduct {
	return &UpdateProduct{
		repository,
	}
}

// Execute finder in the repository of products
func (prodUpdater *UpdateProduct) Execute(ctx context.Context, p products.Product) (products.Product, error) {
	if validateError := products.ValidateProduct(p); validateError != nil {
		return products.Product{}, validateError
	}
	productUpdated, err := prodUpdater.productsRepository.UpdateProduct(ctx, p)
	if err != nil {
		return products.Product{}, err
	}
	return productUpdated, nil
}
