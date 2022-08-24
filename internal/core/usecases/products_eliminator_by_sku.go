package usecases

import (
	"context"
	"github.com/skus-finder-psql/internal/core/domain/products"
)

// EliminatorProductsBySKU is the use case than find all products
type EliminatorProductsBySKU struct {
	productsRepository products.Repository
}

func NewEliminatorProductsBySKU(repository products.Repository) *EliminatorProductsBySKU {
	return &EliminatorProductsBySKU{
		repository,
	}
}

// Execute finder a product by his ID in the repository of products
func (prodEliminator *EliminatorProductsBySKU) Execute(ctx context.Context, prodSKU string) (bool, error) {
	var ok bool
	if err := products.ValidateProductID(prodSKU); err != nil {
		return ok, err
	}
	ok, err := prodEliminator.productsRepository.DeleteProductBySKU(ctx, prodSKU)
	if err != nil {
		return ok, err
	}
	return ok, nil
}
