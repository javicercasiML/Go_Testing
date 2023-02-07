package shark

import (
	"fmt"
	"testdoubles/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

// El tiburón logra cazar el atún al ser más veloz y al estar en una distancia corta.
// Hacer un assert de que el método GetLinearDistance fue llamado.

func TestHunt_OK(t *testing.T) {
	// arrange
	simulatorMock := mocks.NewSimulatorMock(10)
	whiteShark := &whiteShark{
		speed:     20,
		position:  [2]float64{5, 1},
		simulator: simulatorMock,
	}
	prey := mocks.NewPreyStub()
	// act
	err := whiteShark.Hunt(prey)
	// assert
	assert.NoError(t, err)
	assert.True(t, simulatorMock.Spy)
}

// El tiburón no logra cazar el atún al ser más lento
func TestHunt_ErrSpeed(t *testing.T) {
	// arrange
	simulatorMock := mocks.NewSimulatorMock(10)
	whiteShark := &whiteShark{
		speed:     5,
		position:  [2]float64{5, 1},
		simulator: simulatorMock,
	}
	ErrNotHunt := fmt.Errorf("could not hunt the prey")
	prey := mocks.NewPreyStub()
	// act
	err := whiteShark.Hunt(prey)
	// assert
	assert.Error(t, err)
	assert.Equal(t, err, ErrNotHunt)
	assert.True(t, simulatorMock.Spy)
}

// El tiburón no logra cazar el atún por estar a una distancia
// muy larga, a pesar de ser más veloz
func TestHunt_ErrDistance(t *testing.T) {
	// arrange
	simulatorMock := mocks.NewSimulatorMock(10)
	whiteShark := &whiteShark{
		speed:     20,
		position:  [2]float64{5, 100},
		simulator: simulatorMock,
	}
	ErrNotHunt := fmt.Errorf("could not hunt the prey")
	prey := mocks.NewPreyStub()
	// act
	err := whiteShark.Hunt(prey)
	// assert
	assert.Error(t, err)
	assert.Equal(t, err, ErrNotHunt)
	assert.True(t, simulatorMock.Spy)
}
