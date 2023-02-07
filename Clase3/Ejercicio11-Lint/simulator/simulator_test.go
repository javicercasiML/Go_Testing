package simulator

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCancatch_OK(t *testing.T) {
	// arrange
	sm := NewCatchSimulator(100)

	// act
	result := sm.CanCatch(10.1, 10.2, 1.3)

	// assert
	assert.True(t, result)
}

func TestGetLinearDistance_OK(t *testing.T) {
	// arrange
	sm := NewCatchSimulator(100)

	// act
	result := sm.GetLinearDistance([2]float64{1.1, 2.2})

	// assert
	assert.Equal(t, math.Round(result*100)/100, 2.46)
}
