package inmemory

import (
	"context"
	"errors"
	"time"

	"github.com/skus-finder-psql/internal/core/domain/products"
	"github.com/skus-finder-psql/internal/shared/messages"
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
	return products.Product{}, errors.New(messages.ProductSKUNotFound)
}

func (repository *Repository) SaveProduct(_ context.Context, p products.Product) (products.Product, error) {
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()

	_, exist := repository.list[p.Sku]
	if exist {
		return products.Product{}, errors.New(messages.ProductSKUAlreadyExists)
	}
	repository.list[p.Sku] = p
	return p, nil
}

func (repository *Repository) UpdateProduct(_ context.Context, p products.Product) (products.Product, error) {
	p.UpdatedAt = time.Now()

	_, exist := repository.list[p.Sku]
	if !exist {
		return products.Product{}, errors.New(messages.ProductSKUDoesNotExist)
	}
	repository.list[p.Sku] = p
	return p, nil
}

func (repository *Repository) DeleteProductBySKU(_ context.Context, prodSKU string) (bool, error) {
	var found bool
	for key := range repository.list {
		if key == prodSKU {
			found = true
			delete(repository.list, prodSKU)
		}
	}

	if found {
		return true, nil
	} else {
		return false, errors.New(messages.ProductSKUNotFound)
	}
}

func NewInMemoryRepository() products.Repository {
	return &Repository{
		list: map[string]products.Product{},
	}
}
