package mocks

import (
	"testmock/internal/domain"

	"github.com/stretchr/testify/mock"
)

// constructor
func NewStorageTestify() *storageTestify {
	return &storageTestify{}
}

// controller
type storageTestify struct {
	mock.Mock
}
func (st *storageTestify) GetByName(name string) (i domain.Item, err error) {
	args := st.Called(name)
	return args.Get(0).(domain.Item), args.Error(1)
}
func (st *storageTestify) UpdateByName(name string, i domain.Item) (err error) {
	args := st.Called(name, i)
	return args.Error(0)
}