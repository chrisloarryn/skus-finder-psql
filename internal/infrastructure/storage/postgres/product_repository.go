package postgres

import (
	"context"
	"fmt"
	"github.com/skus-finder-psql/internal/core/domain/products"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

	//repository.New(Postgres)

	var results []products.Product

	// Get all records
	res := repository.DB().Find(&results)
	// SELECT * FROM users;
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
	// SELECT * FROM users WHERE name = 'jinzhu' ORDER BY id LIMIT 1;

	if err := res.Error; err != nil {
		return products.Product{}, fmt.Errorf("an error has ocurred %s", err.Error())
	}
	// SELECT * FROM users WHERE id = 10;
	return result, nil
}

func (repository *Repository) DeleteProductBySKU(ctx context.Context, prodSKU string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	var found products.Product
	resu := repository.DB().Where("sku = ?", prodSKU).First(&found)
	// SELECT * FROM users WHERE name = 'jinzhu' ORDER BY id LIMIT 1;

	//.Model(products.Product{Sku: p.Sku})

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
	// SELECT * FROM users WHERE name = 'jinzhu' ORDER BY id LIMIT 1;

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
	// SELECT * FROM users WHERE name = 'jinzhu' ORDER BY id LIMIT 1;

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
	// SELECT * FROM users WHERE name = 'jinzhu' ORDER BY id LIMIT 1;

	if err := resu.Error; err != nil {
		return products.Product{}, fmt.Errorf("an error has ocurred %s", err.Error())
	}

	if len(result.Sku) == 0 {
		return products.Product{}, fmt.Errorf("sku does not exists")
	}

	return p, nil
}

func NewRepository() products.Repository {
	return &Repository{
		db,
	}
}
