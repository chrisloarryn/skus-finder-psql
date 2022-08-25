package usecases_test

import (
	"context"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/skus-finder-psql/internal/core/domain/products/productsmocks"
	"github.com/skus-finder-psql/internal/core/usecases"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestEliminatorProductsBySKU_Execute_ShouldBoolData(t *testing.T) {
	t.Log("Should returns a product from his ID")
	// Setup
	controller := gomock.NewController(t)

	productSKU := "123"

	repository := productsmocks.NewMockRepository(controller)
	repository.EXPECT().DeleteProductBySKU(gomock.Any(), productSKU).Return(true, nil).Times(1)

	operation := usecases.NewEliminatorProductsBySKU(repository)

	// Execute
	result, err := operation.Execute(context.TODO(), productSKU)

	// Verify
	require.NoError(t, err)
	assert.Equal(t, true, result)
}

func TestEliminatorProductsBySKU_Execute_ShouldReturnsAnErrorFromRepository(t *testing.T) {
	t.Log("Should returns an error from repository")
	// Setup
	controller := gomock.NewController(t)

	productSKU := "123"
	customError := fmt.Errorf("this is a custom error")

	repository := productsmocks.NewMockRepository(controller)
	repository.EXPECT().DeleteProductBySKU(gomock.Any(), productSKU).Return(false, customError).Times(1)

	operation := usecases.NewEliminatorProductsBySKU(repository)

	// Execute
	_, err := operation.Execute(context.TODO(), productSKU)

	// Verify
	require.Error(t, err, customError.Error())
}

func TestEliminatorProductsBySKU_Execute_ShouldReturnsAnErrorForInvalidID(t *testing.T) {
	t.Log("Should returns an error for invalid ID")
	// Setup
	controller := gomock.NewController(t)

	productSKU := ""
	repository := productsmocks.NewMockRepository(controller)

	operation := usecases.NewEliminatorProductsBySKU(repository)

	// Execute
	result, err := operation.Execute(context.TODO(), productSKU)

	// Verify
	require.Error(t, err, fmt.Sprintf("Invalid ID: %s", productSKU))
	assert.Equal(t, false, result)
}
