package mocks

import (
	"errors"
	"testmock/internal/domain"
)

// constructor
func NewStorageFake() *storageFake {
	return &storageFake{}
}

// controller
var (
	ErrFakeNotFound = errors.New("error not found")
)

type storageFake struct {
	Db []domain.Item
}
func (st *storageFake) GetByName(name string) (i domain.Item, err error) {
	for _, it := range st.Db {
		if it.Name == name {
			i = it
			return
		}
	}

	err = ErrFakeNotFound
	return
}
func (st *storageFake) UpdateByName(name string, i domain.Item) (err error) {
	var item domain.Item; var ix int; var exists bool
	for i, it := range st.Db {
		if it.Name == name {
			ix = i
			item = it
			exists = true
			break
		}
	}

	if !exists {
		err = ErrFakeNotFound
		return
	}

	item.Weight = i.Weight
	item.Price = i.Price
	item.Release = i.Release
	
	st.Db[ix] = item
	return
}