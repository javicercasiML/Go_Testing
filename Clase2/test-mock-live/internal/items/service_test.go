package items

import (
	"errors"
	"testing"
	"testmock/internal/domain"
	"testmock/internal/items/mocks"
	"time"

	"github.com/stretchr/testify/assert"
)

// -----------------------------------------------------------------------------------------
// Test Get: using storage_mock
func TestGetTotalByName(t *testing.T) {
	// arrange
	st := mocks.NewStorageMock()
	sv := NewService(st)

	// act
	t.Run("happy path", func(t *testing.T) {
		// arrange
		d, _ := time.Parse(YYYY_MM_DD, "2022-01-01")
		st.Data = domain.Item{ID: "AAA", Name: "Pepsi", Weight: 1.5, Price: 150, Release: d}

		expected := float64(150 * 5)

		// act
		total, err := sv.GetTotalByName("Pepsi", 5)

		// assert
		assert.NoError(t, err)
		assert.Equal(t, expected, total)
		assert.True(t, st.Spy)
	})

	t.Run("error", func(t *testing.T) {
		// arrange
		st.Reset()
		e := errors.New("internal error")
		st.Err = e

		// act
		total, err := sv.GetTotalByName("Pepsi", 5)

		// assert
		assert.Error(t, err)
		assert.ErrorIs(t, err, e)
		assert.Zero(t, total)
		assert.True(t, st.Spy)
	})
}
func TestGetTotalByName_Fake(t *testing.T) {
	// arrange
	st := mocks.NewStorageFake()
	sv := NewService(st)

	// act
	t.Run("ok", func(t *testing.T) {
		// arrange
		d, _ := time.Parse(YYYY_MM_DD, "2022-01-01")
		db := []domain.Item{{ID: "AAA", Name: "Pepsi", Weight: 1.5, Price: 150, Release: d}}
		st.Db = db

		// act
		total, err := sv.GetTotalByName("Pepsi", 5)

		// assert
		assert.NoError(t, err)
		assert.Equal(t, 750.0, total)
	})

	t.Run("err", func(t *testing.T) {
		// arrange
		d, _ := time.Parse(YYYY_MM_DD, "2022-01-01")
		db := []domain.Item{{ID: "BBB", Name: "Coke", Weight: 1.5, Price: 150, Release: d}}
		st.Db = db

		// act
		total, err := sv.GetTotalByName("Pepsi", 5)

		// assert
		assert.Error(t, err)
		assert.ErrorIs(t, err, mocks.ErrFakeNotFound)
		assert.Equal(t, 0.0, total)
	})
}



// -----------------------------------------------------------------------------------------
// Test Update: using storage_mockdyn
func TestUpdateByName_MockDyn(t *testing.T) {
	// arrange
	st := mocks.NewStorageMockDyn()
	sv := NewService(st)

	// act
	t.Run("ok", func(t *testing.T) {
		// arrange
		st.Reset()
		d, _ := time.Parse(YYYY_MM_DD, "2022-01-01")
		st.Data = domain.Item{ID: "AAA", Name: "Pepsi", Weight: 1.5, Price: 150, Release: d}

		// act
		i, err := sv.UpdateByName("Pepsi", nil, nil, nil)

		// assert
		assert.NoError(t, err)
		assert.Equal(t, st.Data, i)
		assert.True(t, st.Spy["GetByName"])
		assert.True(t, st.Spy["UpdateByName"])
	})

	t.Run("fail get", func(t *testing.T) {
		// arrange
		st.Reset()
		e := errors.New("error internal")
		st.Err["GetByName"] = e

		// act
		i, err := sv.UpdateByName("Pepsi", nil, nil, nil)
		
		// assert
		assert.Error(t, err)
		assert.ErrorIs(t, err, e)
		assert.Empty(t, i)
		assert.True(t, st.Spy["GetByName"])
		assert.False(t, st.Spy["UpdateByName"])
	})
}

// Test Update: using storage_mock testify
func TestUpdateByName_MockTestify(t *testing.T) {
	// arrange

	// act
	t.Run("fail update", func(t *testing.T) {
		// arrange
		st := mocks.NewStorageTestify()
		sv := NewService(st)

		d, _ := time.Parse(YYYY_MM_DD, "2022-01-01")
		data := domain.Item{ID: "AAA", Name: "Pepsi", Weight: 1.5, Price: 150, Release: d}
		e := errors.New("error internal")

		// -> define behavior for the mock
		st.On("GetByName", "Pepsi").Return(data, nil)
		st.On("UpdateByName", "Pepsi", data).Return(e)

		// act
		i, err := sv.UpdateByName("Pepsi", nil, nil, nil)
		
		// assert
		assert.Error(t, err)
		assert.ErrorIs(t, err, e)
		assert.Equal(t, data, i)
		assert.True(t, st.AssertExpectations(t))
	})
}