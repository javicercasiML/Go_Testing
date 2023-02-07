package products

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLocalRepository_Get(t *testing.T) {
	// Arrange.
	productToSearch := "abc123"
	expectedResult := &Product{
		ID:   "abc123",
		Name: "Product 1",
	}

	repository := NewLocalRepository(&FakeStorage{
		ResultOnGet: &Product{
			ID:   "abc123",
			Name: "Product 1",
		},
		ErrOnGet: nil,
	})

	// Act.
	obtainedResult, err := repository.Get(productToSearch)

	// Assert.
	assert.NoError(t, err)
	assert.Equal(t, expectedResult, obtainedResult)
}

func TestLocalRepository_Get_ErrProductNotFound(t *testing.T) {
	// Arrange.
	productToSearch := "non_existing_product"
	expectedErr := ErrProductNotFound

	repository := NewLocalRepository(&FakeStorage{
		ResultOnGet: nil,
	})

	// Act.
	obtainedResult, err := repository.Get(productToSearch)

	// Assert.
	assert.EqualError(t, err, expectedErr.Error())
	assert.Nil(t, obtainedResult)
}

func TestLocalRepository_Get_ErrUnexpectedOnDB(t *testing.T) {
	// Arrange.
	productToSearch := "abc123"
	expectedErr := ErrUnexpected

	repository := NewLocalRepository(&FakeStorage{
		ErrOnGet: errors.New("soy un error que no conoces, romperé tu código"),
	})

	// Act.
	obtainedResult, err := repository.Get(productToSearch)

	// Assert.
	assert.EqualError(t, err, expectedErr.Error())
	assert.Nil(t, obtainedResult)
}
