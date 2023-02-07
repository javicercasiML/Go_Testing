package hunt

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSharkHuntsSuccessfully(t *testing.T) {
	// Arrange
	s := Shark{
		hungry: true,
		tired:  false,
		speed:  100,
	}
	presa := Prey{
		name:  "Pejerrey",
		speed: 50,
	}
	// Act
	err := s.Hunt(&presa)

	// Assert
	assert.NoError(t, err)
	assert.Nil(t, err)
	assert.Equal(t, false, s.hungry)
	assert.Equal(t, true, s.tired)
}

func TestSharkCannotHuntBecauseIsTired(t *testing.T) {
	// Arrange
	s := Shark{
		hungry: false,
		tired:  true,
		speed:  100,
	}
	presa := Prey{
		name:  "Pejerrey",
		speed: 50,
	}
	ErrExpected := errors.New("cannot hunt, i am really tired")

	// Act
	obteinedResult := s.Hunt(&presa)

	// Assert
	assert.Error(t, obteinedResult)
	assert.Equal(t, ErrExpected, obteinedResult)
}

func TestSharkCannotHuntBecaisIsNotHungry(t *testing.T) {
	// Arrange
	s := Shark{
		hungry: false,
		tired:  false,
		speed:  100,
	}
	presa := Prey{
		name:  "Pejerrey",
		speed: 50,
	}
	ErrExpected := errors.New("cannot hunt, i am not hungry")

	// Act
	obteinedResult := s.Hunt(&presa)

	// Assert
	assert.Error(t, obteinedResult)
	assert.Equal(t, ErrExpected, obteinedResult)
}

func TestSharkCannotReachThePrey(t *testing.T) {
	// Arrange
	s := Shark{
		hungry: true,
		tired:  false,
		speed:  100,
	}
	presa := Prey{
		name:  "Pejerrey",
		speed: 150,
	}
	ErrExpected := errors.New("could not catch it")

	// Act
	obteinedResult := s.Hunt(&presa)

	// Assert
	assert.Error(t, obteinedResult)
	assert.Equal(t, ErrExpected, obteinedResult)

}

func TestSharkHuntNilPrey(t *testing.T) {
	// Arrange
	s := Shark{
		hungry: true,
		tired:  false,
		speed:  100,
	}
	presa := Prey{
		name:  "",
		speed: 0,
	}
	ErrExpected := errors.New("invalid Prey")

	// Act
	obteinedResult := s.Hunt(&presa)

	// Assert
	assert.Error(t, obteinedResult)
	assert.Equal(t, ErrExpected, obteinedResult)

}

func TestSharkSuccessfully(t *testing.T) {
	t.Run("Prey: Pejerrey", func(t *testing.T) {

		// Arrange
		s := Shark{
			hungry: true,
			tired:  false,
			speed:  100,
		}
		presa := Prey{
			name:  "Pejerrey",
			speed: 50,
		}
		// Act
		err := s.Hunt(&presa)

		// Assert
		assert.NoError(t, err)
		assert.Nil(t, err)
		assert.Equal(t, false, s.hungry)
		assert.Equal(t, true, s.tired)
	})
	t.Run("Prey: Dientudo", func(t *testing.T) {

		// Arrange
		s := Shark{
			hungry: true,
			tired:  false,
			speed:  100,
		}
		presa := Prey{
			name:  "Dientudo",
			speed: 25,
		}
		// Act
		err := s.Hunt(&presa)

		// Assert
		assert.NoError(t, err)
		assert.Nil(t, err)
		assert.Equal(t, false, s.hungry)
		assert.Equal(t, true, s.tired)
	})
}
