package usecases

import (
	"context"
	"github.com/skus-finder-psql/internal/core/domain/products"
)

// FinderAllProducts is the use case than find all products
type FinderAllProducts struct {
	productsRepository products.Repository
}

func NewFinderAllProducts(repository products.Repository) *FinderAllProducts {
	return &FinderAllProducts{
		repository,
	}
}

// Execute finder in the repository of products
func (prodFinder *FinderAllProducts) Execute(ctx context.Context) ([]products.Product, error) {
	productList, err := prodFinder.productsRepository.FindAllProducts(ctx)
	if err != nil {
		return nil, err
	}
	return productList, nil
}
