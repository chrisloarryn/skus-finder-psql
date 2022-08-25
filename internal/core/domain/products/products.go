package products

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
)

// Product represents the product data
type Product struct {
	gorm.Model
	Sku            string    `gorm:"type:varchar(50); not null" json:"sku"`
	Name           string    `gorm:"type:varchar(50); not null" json:"name"`
	Size           *string   `gorm:"type:varchar(50)" json:"size"`
	Price          float64   `gorm:"type:double precision;not null" json:"price"`
	PrincipalImage *string   `gorm:"type:varchar(50);not null" json:"principal_image"`
	OtherImages    []*string `gorm:"type:varchar(50)" json:"other_images"`
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
	SaveProduct(ctx context.Context, product Product) (Product, error)
}
