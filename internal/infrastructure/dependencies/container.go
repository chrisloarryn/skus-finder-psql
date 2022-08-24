package dependencies

import (
	"github.com/skus-finder-psql/internal/core/domain/products"
	"github.com/skus-finder-psql/internal/infrastructure/storage"
	"os"
)

type Container interface {
	ProductsRepository() products.Repository
}

type container struct {
	productsRepository products.Repository
}

const EnvironmentKey = "ENVIRONMENT"

func NewContainer() Container {
	environment := os.Getenv(EnvironmentKey)

	return &container{
		productsRepository: storage.New(environment),
	}
}

func (container *container) ProductsRepository() products.Repository {
	return container.productsRepository
}
