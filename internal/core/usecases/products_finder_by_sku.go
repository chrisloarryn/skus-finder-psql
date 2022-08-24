package usecases

import (
	"context"
	"github.com/skus-finder-psql/internal/core/domain/products"
)

// FinderProductsBySKU is the use case than find all products
type FinderProductsBySKU struct {
	productsRepository products.Repository
}

func NewFinderProductsBySKU(repository products.Repository) *FinderProductsBySKU {
	return &FinderProductsBySKU{
		repository,
	}
}

// Execute finder a product by his ID in the repository of products
func (prodFinder *FinderProductsBySKU) Execute(ctx context.Context, prodSKU string) (products.Product, error) {
	if err := products.ValidateProductID(prodSKU); err != nil {
		return products.Product{}, err
	}
	productResult, err := prodFinder.productsRepository.FindProductBySKU(ctx, prodSKU)
	if err != nil {
		return products.Product{}, err
	}
	return productResult, nil
}
