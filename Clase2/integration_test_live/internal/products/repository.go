package products

import "errors"

var (
	ErrUnexpected      = errors.New("unexpected error")
	ErrProductNotFound = errors.New("product not found")
)

type Repository interface {
	Get(id string) (*Product, error)
}
