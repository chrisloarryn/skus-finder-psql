package storage

import (
	"github.com/skus-finder-psql/internal/core/domain/products"
	"github.com/skus-finder-psql/internal/infrastructure/storage/inmemory"
	"github.com/skus-finder-psql/internal/infrastructure/storage/postgres"
)

const (
	PROD = "PRODUCTION"
)

func New(environment string) products.Repository {
	switch environment {
	case PROD:
		repository := &postgres.Repository{}
		bd := repository.New(postgres.Postgres)
		bd.AutoMigrate(&products.Product{})
		return postgres.NewRepository()
	default:
		return inmemory.NewInMemoryRepository()
	}
}
