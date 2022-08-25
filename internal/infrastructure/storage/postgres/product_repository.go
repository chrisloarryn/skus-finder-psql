package postgres

import (
	"context"
	"fmt"
	"github.com/skus-finder-psql/internal/core/domain/products"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"sync"
	"time"
)

var (
	db                  *gorm.DB
	once                sync.Once
	mysqlDataSourceName = "root:mysql@tcp(localhost:3308)/mysqldbd?parseTime=true"
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
			dbPort = "5432"
		}
		dbName := os.Getenv(DB_NAME)
		if len(dbName) == 0 {
			dbName = string(Postgres)
		}

		// pgDataSourceName := "postgres://postgres:postgres@products-sku-api-db:65432/postgres?sslmode=disable"
		pgDataSourceName := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPass, dbHost, dbPort, dbName)
		db, err = gorm.Open(postgres.Open(pgDataSourceName)) //"postgres", pgDataSourceName)
		if err != nil {
			log.Fatalf("can't open db: %v", err)
		}

		fmt.Println("connected to postgres")
	})
}

func (repository *Repository) newMySQLDB() {
	once.Do(func() {
		var err error
		db, err = gorm.Open(mysql.Open(mysqlDataSourceName))
		if err != nil {
			log.Fatalf("can't open db: %v", err)
		}

		fmt.Println("connected to mysql")
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
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	var results []products.Product

	res := repository.DB().Find(&results)

	if err := res.Error; err != nil {
		return []products.Product{}, fmt.Errorf("an error has ocurred: %s", err.Error())
	}

	return results, nil
}

func (repository *Repository) FindProductBySKU(ctx context.Context, prodSKU string) (products.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	var result products.Product
	res := repository.DB().Where("sku = ?", prodSKU).First(&result)

	if err := res.Error; err != nil {
		return products.Product{}, fmt.Errorf("an error has ocurred %s", err.Error())
	}
	return result, nil
}

func (repository *Repository) DeleteProductBySKU(ctx context.Context, prodSKU string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	var found products.Product
	resu := repository.DB().Where("sku = ?", prodSKU).First(&found)

	if err := resu.Error; err != nil {
		if err.Error() != "record not found" {
			return false, fmt.Errorf("an error has ocurred %s", err.Error())
		}
	}

	if found.ID == 0 {
		return false, fmt.Errorf("sku does not exists")
	}

	res := repository.DB().Where("sku = ?", prodSKU).Delete(products.Product{Sku: prodSKU})

	if err := res.Error; err != nil {
		return false, fmt.Errorf("an error has ocurred %s", err.Error())
	}

	return true, nil
}

func (repository *Repository) SaveProduct(ctx context.Context, p products.Product) (products.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	var result products.Product
	resu := repository.DB().Where("sku = ?", p.Sku).First(&result)

	if err := resu.Error; err != nil {
		if err.Error() != "record not found" {
			return products.Product{}, fmt.Errorf("an error has ocurred %s", err.Error())
		}
	}

	if result.Sku == p.Sku {
		return products.Product{}, fmt.Errorf("sku already exists")
	}

	res := repository.DB().Create(&p)

	if err := res.Error; err != nil {
		return products.Product{}, fmt.Errorf("an error has ocurred: %s", err.Error())
	}

	var afterSave products.Product
	resul := repository.DB().Where("sku = ?", p.Sku).First(&afterSave)

	if err := resul.Error; err != nil {
		return products.Product{}, fmt.Errorf("an error has ocurred %s", err.Error())
	}

	return afterSave, nil
}

func (repository *Repository) UpdateProduct(ctx context.Context, p products.Product) (products.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	var result products.Product
	resu := repository.DB().Where("sku = ?", p.Sku).First(&result)

	if err := resu.Error; err != nil {
		if err.Error() != "record not found" {
			return products.Product{}, fmt.Errorf("an error has ocurred %s", err.Error())
		}
	}

	if len(result.Sku) == 0 {
		return products.Product{}, fmt.Errorf("sku does not exists")
	}

	res := db.Model(products.Product{}).Where("sku = ?", p.Sku).Updates(p)

	if err := res.Error; err != nil {
		return products.Product{}, fmt.Errorf("an error has ocurred %s", err.Error())
	}

	var updatedResult products.Product
	resul := repository.DB().Where("sku = ?", p.Sku).First(&updatedResult)

	if err := resul.Error; err != nil {
		return products.Product{}, fmt.Errorf("an error has ocurred %s", err.Error())
	}
	return updatedResult, nil
}

func NewRepository() products.Repository {
	return &Repository{
		db,
	}
}
