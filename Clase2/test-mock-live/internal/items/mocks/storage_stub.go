package mocks

import (
	"testmock/internal/domain"
	"time"
)

// constructor
func NewStorageStub() *storageStub {
	return &storageStub{}
}

// controller
type storageStub struct {
	
}
func (st *storageStub) GetByName(name string) (i domain.Item, err error) {
	i = domain.Item{
		ID: "AAA",
		Name: "Pepsi",
		Weight: 1.5,
		Price: 150,
		Release: time.Now(),
	}
	return
}
func (st *storageStub) UpdateByName(name string, i domain.Item) (err error) {
	return
}