package products

type FakeStorage struct {
	ResultOnGet *Product
	ErrOnGet    error
}

func (storage *FakeStorage) Get(id string) (product *Product, err error) {
	return storage.ResultOnGet, storage.ErrOnGet
}
