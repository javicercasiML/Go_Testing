package internal

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/go-bootcamp/test-integracion/pkg/store"
	"github.com/stretchr/testify/assert"
)

func TestGetAllSuccesfullS(t *testing.T) {
	//Arrange
	var users = []User{
		{
			ID:    1,
			Name:  "Juan",
			Email: "juan@emial.com",
			Age:   32,
		},
		{
			ID:    2,
			Name:  "Pepe",
			Email: "pepe@emial.com",
			Age:   17,
		},
	}
	dataJson, _ := json.Marshal(users)
	dbMock := store.Mock{
		Data: dataJson,
	}

	storeStub := store.FileStore{
		FileName: "",
		Mock:     &dbMock,
	}
	repo := NewRepository(storeStub)
	sv := NewService(repo)

	//Act
	myUsers, _ := sv.GetAll()

	//Assert
	assert.Equal(t, users, myUsers)
}

func TestGetAllFailS(t *testing.T) {
	//Arrange
	var user = User{

		ID:    1,
		Name:  "Juan",
		Email: "juan@emial.com",
		Age:   32,
	}

	expectedError := errors.New("error for Storage")
	dbMock := store.Mock{
		Err: expectedError,
	}
	storeStub := store.FileStore{
		FileName: "",
		Mock:     &dbMock,
	}
	rep := NewRepository(storeStub)
	sv := NewService(rep)

	//act
	result, err := sv.Store(user.ID, user.Name, user.Email, user.Age)

	//Assert
	assert.Equal(t, expectedError, err)
	assert.Equal(t, User{}, result)
}
