package mocks

import "testmock/internal/domain"

// constructor
func NewStorageDummy() *storageDummy {
	return &storageDummy{}
}

// controller
type storageDummy struct {
	
}
func (st *storageDummy) GetByName(name string) (i domain.Item, err error) {
	return
}
func (st *storageDummy) UpdateByName(name string, i domain.Item) (err error) {
	return
}