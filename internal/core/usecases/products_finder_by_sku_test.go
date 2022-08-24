package usecases_test

import (
	"context"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/skus-finder-psql/internal/core/domain/products"
	"github.com/skus-finder-psql/internal/core/domain/products/productsmocks"
	"github.com/skus-finder-psql/internal/core/usecases"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	commons "github.com/validatecl/go-microservices-commons"
	"testing"
)

func TestFinderProductBySKU_Execute_ShouldReturnsAProductData(t *testing.T) {
	t.Log("Should returns a product from his SKU")
	// Setup
	controller := gomock.NewController(t)

	productSKU := "213321123123"
	productResult := products.Product{
		Sku:            "213321123123",
		Name:           "product",
		Size:           commons.StringPointer("XL"),
		Price:          1000.01,
		PrincipalImage: commons.StringPointer("http://image.png"),
		OtherImages: []*string{
			commons.StringPointer("http://other_image.png"),
		},
	}

	repository := productsmocks.NewMockRepository(controller)
	repository.EXPECT().FindProductBySKU(gomock.Any(), productSKU).Return(productResult, nil).Times(1)

	operation := usecases.NewFinderProductsBySKU(repository)

	// Execute
	result, err := operation.Execute(context.TODO(), productSKU)

	// Verify
	require.NoError(t, err)
	assert.Equal(t, productResult, result)
}

func TestFinderProductBySKU_Execute_ShouldReturnsAnErrorFromRepository(t *testing.T) {
	t.Log("Should returns an error from repository")
	// Setup
	controller := gomock.NewController(t)

	productSKU := "123"
	customError := fmt.Errorf("this is a custom error")

	repository := productsmocks.NewMockRepository(controller)
	repository.EXPECT().FindProductBySKU(gomock.Any(), productSKU).Return(products.Product{}, customError).Times(1)

	operation := usecases.NewFinderProductsBySKU(repository)

	// Execute
	result, err := operation.Execute(context.TODO(), productSKU)

	// Verify
	require.Error(t, err, customError.Error())
	assert.Equal(t, products.Product{}, result)
}

func TestFinderProductBySKU_Execute_ShouldReturnsAnErrorForInvalidID(t *testing.T) {
	t.Log("Should returns an error for invalid ID")
	// Setup
	controller := gomock.NewController(t)

	productSKU := ""
	repository := productsmocks.NewMockRepository(controller)

	operation := usecases.NewFinderProductsBySKU(repository)

	// Execute
	result, err := operation.Execute(context.TODO(), productSKU)

	// Verify
	require.Error(t, err, fmt.Sprintf("Invalid ID: %s", productSKU))
	assert.Equal(t, products.Product{}, result)
}
