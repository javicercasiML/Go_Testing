package products

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestService_Get(t *testing.T) {
	// Arrange.
	productToSearch := "abc"
	expectedResult := &Product{
		ID:   "abc",
		Name: "Test Product",
	}

	service := NewService(&FakeRepository{
		ResultOnGet: &Product{
			ID:   "abc",
			Name: "Test Product",
		},
	})

	// Act.
	obtainedResult, err := service.Get(productToSearch)

	// Assert.
	assert.NoError(t, err)
	assert.Equal(t, expectedResult, obtainedResult)
}

func TestService_Get_ErrNotFound(t *testing.T) {
	// Arrange.
	productToSearch := "non_existing_product"
	expectedErr := ErrProductNotFound

	service := NewService(&FakeRepository{
		ErrOnGet: ErrProductNotFound,
	})

	// Act.
	obtainedResult, err := service.Get(productToSearch)

	// Assert.
	assert.EqualError(t, err, expectedErr.Error())
	assert.Nil(t, obtainedResult)
}

func TestService_Get_ErrUnexpected(t *testing.T) {
	// Arrange.
	productToSearch := "non_existing_product"
	expectedErr := ErrUnexpected

	service := NewService(&FakeRepository{
		ErrOnGet: errors.New("i'm an error :P"),
	})

	// Act.
	obtainedResult, err := service.Get(productToSearch)

	// Assert.
	assert.EqualError(t, err, expectedErr.Error())
	assert.Nil(t, obtainedResult)
}
