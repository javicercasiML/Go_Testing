package hello

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Salute_OK(t *testing.T) {

	//arrange
	f := NewPerson("ricardo")

	//act
	s, err := f.Salute() // Usamos el metodo salute de la variable p

	//assert
	assert.NotZero(t, len(s))
	assert.Nil(t, err)
}

func Test_Salute_Error(t *testing.T) {

	//arrange
	p := NewPerson("")

	//act
	q, err := p.Salute()

	//assert
	assert.Zero(t, len(q))
	assert.NotNil(t, err)
}
