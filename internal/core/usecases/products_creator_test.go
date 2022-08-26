package usecases_test

import (
	"context"
	"fmt"
	"github.com/skus-finder-psql/internal/core/domain/products"
	"github.com/skus-finder-psql/internal/core/domain/products/productsmocks"
	"github.com/skus-finder-psql/internal/core/usecases"
	commons "github.com/validatecl/go-microservices-commons"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateProduct_Execute_ShouldCreateAProduct(t *testing.T) {
	t.Log("Should create a Product")
	// Setup
	controller := gomock.NewController(t)

	repository := productsmocks.NewMockRepository(controller)

	newProduct := products.Product{
		Sku:            "FAL-1000009",
		Brand:          "FAL",
		Name:           "product",
		Size:           commons.StringPointer("XL"),
		Price:          1000.01,
		PrincipalImage: commons.StringPointer("http://image.png"),
		OtherImages: []*string{
			commons.StringPointer("http://other_image.png"),
		},
	}
	repository.EXPECT().SaveProduct(gomock.Any(), newProduct).Return(newProduct, nil).Times(1)

	createProductUseCase := usecases.NewCreateProduct(repository)

	// Execute
	_, err := createProductUseCase.Execute(context.TODO(), newProduct)

	// Verify
	assert.NoError(t, err)
}

func TestCreateProduct_Execute_ShouldReturnAnError(t *testing.T) {
	t.Log("Should return an error when try to create a product")
	// Setup
	controller := gomock.NewController(t)
	newProduct := products.Product{
		Sku:            "FAL-1000009",
		Brand:          "FAL",
		Name:           "product",
		Size:           commons.StringPointer("XL"),
		Price:          1000.01,
		PrincipalImage: commons.StringPointer("http://image.png"),
		OtherImages: []*string{
			commons.StringPointer("http://other_image.png"),
		},
	}
	customError := fmt.Errorf("sku: must be in a valid format")

	repository := productsmocks.NewMockRepository(controller)

	repository.EXPECT().SaveProduct(gomock.Any(), newProduct).Return(products.Product{}, customError).Times(1)

	createProductUseCase := usecases.NewCreateProduct(repository)

	// Execute
	_, err := createProductUseCase.Execute(context.TODO(), newProduct)

	// Verify
	assert.EqualError(t, err, customError.Error())
}

func TestCreateProduct_Execute_ShouldReturnAnErrorForInvalidNegative(t *testing.T) {
	t.Log("Should return an error when try to create a product")
	// Setup
	controller := gomock.NewController(t)
	newProduct := products.Product{
		Sku:            "FAL-1000009",
		Brand:          "FAL",
		Name:           "product",
		Size:           commons.StringPointer("XL"),
		Price:          -1000.01,
		PrincipalImage: commons.StringPointer("http://image.png"),
		OtherImages: []*string{
			commons.StringPointer("http://other_image.png"),
		},
	}
	invalidPriceError := fmt.Errorf("price: must be no less than 1.")

	repository := productsmocks.NewMockRepository(controller)

	createProductUseCase := usecases.NewCreateProduct(repository)

	// Execute
	_, err := createProductUseCase.Execute(context.TODO(), newProduct)

	// Verify
	assert.EqualError(t, err, invalidPriceError.Error())
}
