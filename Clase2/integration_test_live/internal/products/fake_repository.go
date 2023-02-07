package products

type FakeRepository struct {
	ResultOnGet *Product
	ErrOnGet    error
}

func (repository *FakeRepository) Get(id string) (product *Product, err error) {
	return repository.ResultOnGet, repository.ErrOnGet
}
