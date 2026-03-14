package storage

import (
	"log"

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
		if err := bd.AutoMigrate(&products.Product{}); err != nil {
			log.Fatalf("migrate products table: %v", err)
		}
		return postgres.NewRepository()
	default:
		return inmemory.NewInMemoryRepository()
	}
}
