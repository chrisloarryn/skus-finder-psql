package products

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
)

// Product represents the product data
type Product struct {
	gorm.Model
	Sku            string    `json:"sku" gorm:"type:varchar(50); not null"`
	Name           string    `json:"name" gorm:"type:varchar(50); not null"`
	Size           *string   `json:"size" gorm:"type:varchar(50)"`
	Price          float64   `json:"price" gorm:"type:double precision;not null"`
	PrincipalImage *string   `json:"principal_image" gorm:"type:varchar(50);not null"`
	OtherImages    []*string `json:"other_images" gorm:"type:varchar(50)"`
}

// ValidateProductID just validates the ID value shouldn't be negative
func ValidateProductID(productSKU string) error {
	if len(productSKU) < 1 {
		return fmt.Errorf("invalid ID: %s", productSKU)
	}
	return nil
}

// ValidatePrice validates the price value
func ValidatePrice(price float64) error {
	if price < 0 {
		return fmt.Errorf("invalid price")
	}
	return nil
}

// ValidateProduct validates all field and required fields of product data
func ValidateProduct(p Product) error {
	if len(p.Name) == 0 {
		return fmt.Errorf("name couldn't be empty")
	}
	if err := ValidatePrice(p.Price); err != nil {
		return err
	}
	return nil
}

//go:generate mockgen -package productsmocks -destination productsmocks/products_repository_mocks.go . Repository

// Repository is the storage abstraction
type Repository interface {
	FindAllProducts(ctx context.Context) ([]Product, error)
	FindProductBySKU(ctx context.Context, productSKU string) (Product, error)
	UpdateProduct(ctx context.Context, product Product) (Product, error)
	DeleteProductBySKU(ctx context.Context, productSKU string) (bool, error)
	SaveProduct(ctx context.Context, product Product) error
}
