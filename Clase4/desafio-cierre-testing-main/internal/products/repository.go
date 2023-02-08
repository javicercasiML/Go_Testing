package products

import "errors"

type Repository interface {
	GetAllBySeller(sellerID string) ([]Product, error)
}

type repository struct {
	storage []Product
}

func NewRepository(storage []Product) Repository {
	return &repository{storage: storage}
}

func (r *repository) GetAllBySeller(sellerID string) ([]Product, error) {
	var response []Product
	for _, prod := range r.storage {
		if prod.SellerID == sellerID {
			response = append(response, prod)
		}
	}
	if response == nil {
		return nil, errors.New("seller ID not exist")
	}
	return response, nil
}
