package products

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test Get: using storage_mock
func TestGetAllBySeller_ok(t *testing.T) {
	// arrange
	repo := NewRepositoryMock()
	sv := NewService(repo)

	// act
	t.Run("happy path", func(t *testing.T) {
		// arrange
		expectedResult := []Product{
			{
				ID:          "1",
				SellerID:    "1",
				Description: "mock",
				Price:       744,
			},
		}

		repo.Data = append(repo.Data, Product{
			ID:          "1",
			SellerID:    "1",
			Description: "mock",
			Price:       744,
		})

		// act
		prods, err := sv.GetAllBySeller("1")

		// assert
		assert.NoError(t, err)
		assert.Equal(t, expectedResult, prods)
		assert.True(t, repo.Spy)
	})

	// act
	t.Run("error", func(t *testing.T) {
		// arrange
		repo.Reset()
		expectedError := errors.New("internal server error")
		repo.Err = expectedError

		// act
		prods, err := sv.GetAllBySeller("1")

		// assert
		assert.Error(t, err)
		assert.ErrorIs(t, err, expectedError)
		assert.Empty(t, prods)
		assert.True(t, repo.Spy)
	})
}
