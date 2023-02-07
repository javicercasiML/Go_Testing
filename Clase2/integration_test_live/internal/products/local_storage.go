package products

type LocalStorage struct {
	Data []Product
}

func (storage *LocalStorage) Get(id string) (product *Product, err error) {
	for index := range storage.Data {
		if storage.Data[index].ID == id {
			product = &storage.Data[index]
			break
		}
	}

	return
}
