package shark

import (
	"fmt"
	"integrationtests/pkg/mocks"
	"integrationtests/prey"
	"integrationtests/simulator"
	"testing"

	"github.com/stretchr/testify/assert"
)

// El tiburón logra cazar el atún al ser más veloz y al estar en una distancia corta.
// Hacer un assert de que el método GetLinearDistance fue llamado.

func TestHunt_OK(t *testing.T) {
	// arrange
	data := map[string]interface{}{
		"white_shark_speed": 10.1,
		"white_shark_x":     1.1,
		"white_shark_y":     2.2,
		"tuna_speed":        5.4,
	}
	simulator := simulator.NewCatchSimulator(10.1)
	storageMock := mocks.NewStorageMock(data)
	whiteShark := CreateWhiteShark(simulator, storageMock)
	prey := prey.CreateTuna(storageMock)
	// act
	err := whiteShark.Hunt(prey)
	// assert
	assert.NoError(t, err)
	assert.True(t, storageMock.Spy)
}

// El tiburón no logra cazar el atún al ser más lento
func TestHunt_ErrSpeed(t *testing.T) {
	// arrange
	data := map[string]interface{}{
		"white_shark_speed": 4.1,
		"white_shark_x":     1.1,
		"white_shark_y":     2.2,
		"tuna_speed":        5.4,
	}
	ErrNotHunt := fmt.Errorf("could not hunt the prey")
	simulator := simulator.NewCatchSimulator(10.1)
	storageMock := mocks.NewStorageMock(data)
	whiteShark := CreateWhiteShark(simulator, storageMock)
	prey := prey.CreateTuna(storageMock)
	// act
	err := whiteShark.Hunt(prey)
	// assert

	assert.Error(t, err)
	assert.Equal(t, err, ErrNotHunt)
	assert.True(t, storageMock.Spy)
}

// El tiburón no logra cazar el atún por estar a una distancia
// muy larga, a pesar de ser más veloz
func TestHunt_ErrDistance(t *testing.T) {
	// arrange
	data := map[string]interface{}{
		"white_shark_speed": 10.1,
		"white_shark_x":     100.1,
		"white_shark_y":     2.2,
		"tuna_speed":        5.4,
	}
	ErrNotHunt := fmt.Errorf("could not hunt the prey")
	simulator := simulator.NewCatchSimulator(10.1)
	storageMock := mocks.NewStorageMock(data)
	whiteShark := CreateWhiteShark(simulator, storageMock)
	prey := prey.CreateTuna(storageMock)
	// act
	err := whiteShark.Hunt(prey)
	// assert

	assert.Error(t, err)
	assert.Equal(t, err, ErrNotHunt)
	assert.True(t, storageMock.Spy)
}
