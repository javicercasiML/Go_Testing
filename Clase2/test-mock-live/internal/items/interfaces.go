package items

import (
	"errors"
	"testmock/internal/domain"
)

// Storage: ...
// ...
type Storage interface {
	// read
	GetByName(name string) (i domain.Item, err error)
	// write
	UpdateByName(name string, i domain.Item) (err error)
}

var (
	ErrStorageItemNotFound = errors.New("error storage: item was not found")
	ErrStorageRowsAffected = errors.New("error storage: invalid rows affected")	
	ErrStorageInternal 	   = errors.New("error storage: internal")
)

// Service:
// ...
type Service interface {
	// read
	GetTotalByName(name string, quantity int) (total float64, err error)
	// write
	UpdateByName(name string, weight, price *float64, release *string) (i domain.Item, err error)
}

var (
	ErrServiceInvalidDomain = errors.New("error service: item is not valid")
	ErrServiceInternal 	    = errors.New("error storage: internal")
)