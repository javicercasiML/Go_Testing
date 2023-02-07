package prey

import (
	"testdoubles/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSpeed_OK(t *testing.T) {
	// arange
	st := mocks.NewPreyStub()
	expectedResult := 100.0

	// act
	result := st.GetSpeed()

	// assert
	assert.Equal(t, result, expectedResult)
}
