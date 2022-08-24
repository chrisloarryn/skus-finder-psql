package usecases

import (
	"context"
	"github.com/skus-finder-psql/internal/core/domain/products"
)

// CreateProduct is the use case that create a product
type CreateProduct struct {
	productsRepository products.Repository
}

// NewCreateProduct constructor
func NewCreateProduct(repository products.Repository) *CreateProduct {
	return &CreateProduct{
		repository,
	}
}

// Execute finder in the repository of products
func (prodCreator *CreateProduct) Execute(ctx context.Context, p products.Product) error {
	if validateError := products.ValidateProduct(p); validateError != nil {
		return validateError
	}
	err := prodCreator.productsRepository.SaveProduct(ctx, p)
	if err != nil {
		return err
	}
	return nil
}
