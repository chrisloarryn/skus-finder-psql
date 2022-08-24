package inmemory

import (
	"context"
	"fmt"
	"github.com/skus-finder-psql/internal/core/domain/products"
	"time"
)

// Repository is the struct when you choose the in memory storage
type Repository struct {
	list map[string]products.Product
}

func (repository *Repository) FindAllProducts(_ context.Context) ([]products.Product, error) {
	var result []products.Product

	for _, prod := range repository.list {
		result = append(result, prod)
	}

	return result, nil
}

func (repository *Repository) FindProductBySKU(_ context.Context, prodSKU string) (products.Product, error) {

	for key, prod := range repository.list {
		if key == prodSKU {
			return prod, nil
		}
	}
	return products.Product{}, fmt.Errorf("product ID doesn't exist")
}

func (repository *Repository) SaveProduct(_ context.Context, p products.Product) error {
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()

	_, exist := repository.list[p.Sku]
	if exist {
		return fmt.Errorf("the product SKU already exists")
	}
	repository.list[p.Sku] = p
	return nil
}

func (repository *Repository) UpdateProduct(_ context.Context, p products.Product) (products.Product, error) {
	p.UpdatedAt = time.Now()

	_, exist := repository.list[p.Sku]
	if !exist {
		return products.Product{}, fmt.Errorf("the product SKU does not exists")
	}
	repository.list[p.Sku] = p
	return p, nil
}

func (repository *Repository) DeleteProductBySKU(_ context.Context, prodSKU string) (bool, error) {
	var found bool
	for key, _ := range repository.list {
		fmt.Println(key, prodSKU)
		if key == prodSKU {
			found = true
			delete(repository.list, prodSKU)
		}
	}

	fmt.Println("FOUND", found)
	if found {
		return true, nil
	} else {
		return false, fmt.Errorf("product SKU doesn't exist")
	}
}

func NewInMemoryRepository() products.Repository {
	return &Repository{
		list: map[string]products.Product{},
	}
}
