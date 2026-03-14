package products

import (
	"context"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"

	sharedconstants "github.com/skus-finder-psql/internal/shared/constants"
	"github.com/skus-finder-psql/internal/shared/messages"
	"gorm.io/gorm"
)

// Product represents the product data
type Product struct {
	gorm.Model
	Sku            string    `gorm:"type:varchar(50); not null" json:"sku"`
	Name           string    `gorm:"type:varchar(50); not null" json:"name"`
	Brand          string    `gorm:"type:varchar(50); not null" json:"brand"`
	Size           *string   `gorm:"type:varchar(50)" json:"size"`
	Price          float64   `gorm:"type:double precision;not null" json:"price"`
	PrincipalImage *string   `gorm:"type:varchar(50);not null" json:"principal_image"`
	OtherImages    []*string `gorm:"type:varchar(50)" json:"other_images"`
}

// ValidateProductID just validates the ID value shouldn't be negative
func ValidateProductID(productSKU string) error {
	if len(productSKU) < 1 {
		return fmt.Errorf(messages.InvalidIDFormat, productSKU)
	}
	return nil
}

// ValidatePrice validates the price value
func ValidatePrice(price float64) error {
	if price < sharedconstants.NegativePriceThreshold {
		return fmt.Errorf(messages.InvalidPrice)
	}
	return nil
}

// ValidateProduct validates all field and required fields of product data
func ValidateProduct(p Product) error {
	err := validation.ValidateStruct(&p,
		validation.Field(&p.Sku, validation.Required, validation.Length(sharedconstants.ProductSKUMinLength, sharedconstants.ProductSKUMaxLength), validation.Match(sharedconstants.ProductSKURegexp)),

		validation.Field(&p.Name, validation.Required, validation.Length(sharedconstants.ProductTextMinLength, sharedconstants.ProductTextMaxLength)),
		validation.Field(&p.Brand, validation.Required, validation.Length(sharedconstants.ProductTextMinLength, sharedconstants.ProductTextMaxLength)),
		validation.Field(&p.Price, validation.Required, validation.Min(sharedconstants.ProductPriceMin), validation.Max(sharedconstants.ProductPriceMax)),
		validation.Field(&p.PrincipalImage, validation.Required, is.URL, validation.Required, validation.Length(sharedconstants.ProductTextMinLength, sharedconstants.ProductTextMaxLength)),
	)

	if err != nil {
		return err
	}

	if len(p.Name) == 0 {
		return fmt.Errorf(messages.NameCouldNotBeEmpty)
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
