package postgres

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/skus-finder-psql/internal/core/domain/products"
	"log"
	"sync"
	"time"
)

var (
	db                  *gorm.DB
	once                sync.Once
	pgDataSourceName    = "postgres://postgres:postgres@localhost:65432/postgres?sslmode=disable"
	mysqlDataSourceName = "root:mysql@tcp(localhost:3308)/mysqldbd?parseTime=true"
)

type Driver string

const (
	MySQL    Driver = "mysql"
	Postgres Driver = "postgres"
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
		db, err = gorm.Open("postgres", pgDataSourceName)
		if err != nil {
			log.Fatalf("can't open db: %v", err)
		}

		fmt.Println("connected to postgres")
	})
}

func (repository *Repository) newMySQLDB() {
	once.Do(func() {
		var err error
		db, err = gorm.Open("mysql", mysqlDataSourceName)
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
	productsRes := make([]products.Product, 0)
	op := repository.DB().Find(productsRes)

	val := op.Value

	r, _ := json.Marshal(val)
	fmt.Println(string(r))

	return []products.Product{}, op.Error
}

func (repository *Repository) FindProductBySKU(ctx context.Context, prodSKU string) (products.Product, error) {
	return products.Product{}, nil
}

func (repository *Repository) DeleteProductBySKU(ctx context.Context, prodSKU string) (bool, error) {
	return false, nil
}

func (repository *Repository) SaveProduct(ctx context.Context, product products.Product) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err()) // prints "context deadline exceeded"
	}
	op := repository.DB().Create(&product)

	value := op.Value
	err := op.Error

	fmt.Printf("%v, %v\n", value, err)

	return nil
}

func (repository *Repository) UpdateProduct(ctx context.Context, product products.Product) (products.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err()) // prints "context deadline exceeded"
	}
	op := repository.DB().Create(&product)

	value := op.Value
	err := op.Error

	fmt.Printf("%v, %v\n", value, err)

	return products.Product{}, nil
}

func NewRepository() products.Repository {
	return &Repository{
		db,
	}
}
