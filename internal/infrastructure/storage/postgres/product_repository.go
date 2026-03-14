package postgres

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/skus-finder-psql/internal/core/domain/products"
	"github.com/skus-finder-psql/internal/shared/messages"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

type Driver string
type DbEnv string

const (
	MySQL       Driver = "mysql"
	Postgres    Driver = "postgres"
	DB_HOST            = "DB_HOST"
	DB_USER            = "DB_USER"
	DB_PASSWORD        = "DB_PASSWORD"
	DB_PORT            = "DB_PORT"
	DB_NAME            = "DB_NAME"
)

// New create a new instance of db
func (repository *Repository) New(d Driver) *gorm.DB {
	switch d {
	case MySQL:
		repository.newMySQLDB()
	case Postgres:
		repository.newPostgresDB()
	}
	return db
}

func (repository *Repository) newPostgresDB() {
	once.Do(func() {
		var err error
		dbHost := os.Getenv(DB_HOST)
		if len(dbHost) == 0 {
			dbHost = string(Postgres)
		}
		dbUser := os.Getenv(DB_USER)
		if len(dbUser) == 0 {
			dbUser = string(Postgres)
		}
		dbPass := os.Getenv(DB_PASSWORD)
		if len(dbPass) == 0 {
			dbPass = string(Postgres)
		}
		dbPort := os.Getenv(DB_PORT)
		if len(dbPort) == 0 {
			dbPort = defaultPostgresPort
		}
		dbName := os.Getenv(DB_NAME)
		if len(dbName) == 0 {
			dbName = string(Postgres)
		}

		// pgDataSourceName := "postgres://postgres:postgres@products-sku-api-db:65432/postgres?sslmode=disable"
		pgDataSourceName := fmt.Sprintf(postgresDataSourceFormat, dbUser, dbPass, dbHost, dbPort, dbName)
		db, err = gorm.Open(postgres.Open(pgDataSourceName)) //"postgres", pgDataSourceName)
		if err != nil {
			log.Fatalf(openDBErrorFormat, err)
		}

		fmt.Println(postgresConnectedMessage)
	})
}

func (repository *Repository) newMySQLDB() {
	once.Do(func() {
		var err error
		db, err = gorm.Open(mysql.Open(defaultMySQLDataSourceName))
		if err != nil {
			log.Fatalf(openDBErrorFormat, err)
		}

		fmt.Println(mysqlConnectedMessage)
	})
}

// DB return a unique instance of db
func (repository *Repository) DB() *gorm.DB {
	return db
}

type Repository struct {
	db *gorm.DB
}

func (repository *Repository) FindAllProducts(ctx context.Context) ([]products.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, queryTimeout)
	defer cancel()

	var results []products.Product

	res := repository.DB().WithContext(ctx).Find(&results)

	if err := res.Error; err != nil {
		return []products.Product{}, fmt.Errorf(messages.ErrorOccurredAltWithColon, err.Error())
	}

	return results, nil
}

func (repository *Repository) FindProductBySKU(ctx context.Context, prodSKU string) (products.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, queryTimeout)
	defer cancel()

	var result products.Product
	res := repository.DB().WithContext(ctx).Where(skuWhereClause, prodSKU).First(&result)

	if err := res.Error; err != nil {
		return products.Product{}, fmt.Errorf(messages.ErrorOccurredFormat, err.Error())
	}
	return result, nil
}

func (repository *Repository) DeleteProductBySKU(ctx context.Context, prodSKU string) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, queryTimeout)
	defer cancel()

	var found products.Product
	resu := repository.DB().WithContext(ctx).Where(skuWhereClause, prodSKU).First(&found)

	if err := resu.Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return false, fmt.Errorf(messages.ErrorOccurredFormat, err.Error())
		}
	}

	if found.ID == 0 {
		return false, fmt.Errorf(messages.SKUDoesNotExist)
	}

	res := repository.DB().WithContext(ctx).Where(skuWhereClause, prodSKU).Delete(products.Product{Sku: prodSKU})

	if err := res.Error; err != nil {
		return false, fmt.Errorf(messages.ErrorOccurredAltFormat, err.Error())
	}

	return true, nil
}

func (repository *Repository) SaveProduct(ctx context.Context, p products.Product) (products.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, queryTimeout)
	defer cancel()

	var result products.Product
	resu := repository.DB().WithContext(ctx).Where(skuWhereClause, p.Sku).First(&result)

	if err := resu.Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return products.Product{}, fmt.Errorf(messages.ErrorOccurredFormat, err.Error())
		}
	}

	if result.Sku == p.Sku {
		return products.Product{}, fmt.Errorf(messages.SKUAlreadyExists)
	}

	res := repository.DB().WithContext(ctx).Create(&p)

	if err := res.Error; err != nil {
		return products.Product{}, fmt.Errorf(messages.ErrorOccurredWithColon, err.Error())
	}

	var afterSave products.Product
	resul := repository.DB().WithContext(ctx).Where(skuWhereClause, p.Sku).First(&afterSave)

	if err := resul.Error; err != nil {
		return products.Product{}, fmt.Errorf(messages.ErrorOccurredFormat, err.Error())
	}

	return afterSave, nil
}

func (repository *Repository) UpdateProduct(ctx context.Context, p products.Product) (products.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, queryTimeout)
	defer cancel()

	var result products.Product
	resu := repository.DB().WithContext(ctx).Where(skuWhereClause, p.Sku).First(&result)

	if err := resu.Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return products.Product{}, fmt.Errorf(messages.ErrorOccurredFormat, err.Error())
		}
	}

	if len(result.Sku) == 0 {
		return products.Product{}, fmt.Errorf(messages.SKUDoesNotExist)
	}

	res := repository.DB().WithContext(ctx).Model(products.Product{}).Where(skuWhereClause, p.Sku).Updates(p)

	if err := res.Error; err != nil {
		return products.Product{}, fmt.Errorf(messages.ErrorOccurredFormat, err.Error())
	}

	var updatedResult products.Product
	resul := repository.DB().WithContext(ctx).Where(skuWhereClause, p.Sku).First(&updatedResult)

	if err := resul.Error; err != nil {
		return products.Product{}, fmt.Errorf(messages.ErrorOccurredFormat, err.Error())
	}
	return updatedResult, nil
}

func NewRepository() products.Repository {
	return &Repository{
		db,
	}
}
