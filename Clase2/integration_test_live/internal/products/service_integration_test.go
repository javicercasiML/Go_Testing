package products

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestService_Integration_Get(t *testing.T) {
	// Arrange.
	productToSearch := "abc123"
	expectedResult := &Product{
		ID:    "abc123",
		Name:  "Test Product",
		Price: 123.45,
	}

	storage := &FakeStorage{
		ResultOnGet: &Product{
			ID:    "abc123",
			Name:  "Test Product",
			Price: 123.45,
		},
	}
	repository := NewLocalRepository(storage)
	service := NewService(repository)

	// Act.
	obtainedResult, err := service.Get(productToSearch)

	// Assert.
	assert.NoError(t, err)
	assert.Equal(t, expectedResult, obtainedResult)
}

func TestService_Integration_Get_ErrNotFound(t *testing.T) {
	// Arrange.
	productToSearch := "non_existing_product"
	expectedErr := ErrProductNotFound

	storage := &FakeStorage{
		ResultOnGet: nil,
	}
	repository := NewLocalRepository(storage)
	service := NewService(repository)

	// Act.
	obtainedResult, err := service.Get(productToSearch)

	// Assert.
	assert.Equal(t, expectedErr, err)
	assert.Nil(t, obtainedResult)
}

func TestService_Integration_Get_ErrUnexpected(t *testing.T) {
	// Arrange.
	productToSearch := "non_existing_product"
	expectedErr := ErrUnexpected

	storage := &FakeStorage{
		ErrOnGet: errors.New("i'm an unexpected error"),
	}
	repository := NewLocalRepository(storage)
	service := NewService(repository)

	// Act.
	obtainedResult, err := service.Get(productToSearch)

	// Assert.
	assert.Equal(t, expectedErr, err)
	assert.Nil(t, obtainedResult)
}
