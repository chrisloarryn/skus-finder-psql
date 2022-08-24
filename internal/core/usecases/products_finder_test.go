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
	"github.com/stretchr/testify/require"
)

func TestFinderAllProducts_Execute_ShouldReturnsAProductList(t *testing.T) {
	t.Log("Should returns a product list")
	// Setup
	controller := gomock.NewController(t)

	repository := productsmocks.NewMockRepository(controller)
	productsList := []products.Product{
		{
			ID:             123,
			Sku:            "213321123123",
			Name:           "product",
			Size:           commons.StringPointer("XL"),
			Price:          1000.01,
			PrincipalImage: commons.StringPointer("http://image.png"),
			OtherImages: []*string{
				commons.StringPointer("http://other_image.png"),
			},
		},
		{},
	}
	repository.EXPECT().FindAllProducts(gomock.Any()).Return(productsList, nil).Times(1)

	operation := usecases.NewFinderAllProducts(repository)

	// Execute
	result, err := operation.Execute(context.TODO())

	// Verify
	require.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, productsList, result)
}

func TestFinderAllProducts_Execute_ShouldReturnsAnErrorInRepository(t *testing.T) {
	t.Log("Should returns a product list")
	// Setup
	controller := gomock.NewController(t)

	repository := productsmocks.NewMockRepository(controller)
	customError := fmt.Errorf("this is a custom error")
	repository.EXPECT().FindAllProducts(gomock.Any()).Return(nil, customError).Times(1)

	finderAllProducts := usecases.NewFinderAllProducts(repository)

	// Execute
	result, err := finderAllProducts.Execute(context.TODO())

	// Verify
	require.Nil(t, result)
	assert.EqualError(t, err, customError.Error())
}
