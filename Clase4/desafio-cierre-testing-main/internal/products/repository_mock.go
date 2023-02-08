package products

// constructor
func NewRepositoryMock() *repositoryMock {
	return &repositoryMock{}
}

// controller
type repositoryMock struct {
	Data []Product
	Err  error
	Spy  bool
}

func (rm *repositoryMock) GetAllBySeller(sellerID string) ([]Product, error) {
	rm.Spy = true

	if rm.Err != nil {
		err := rm.Err
		return []Product{}, err
	}

	return rm.Data, nil
}

func (rm *repositoryMock) Reset() {
	rm.Data = []Product{}
	rm.Err = nil
	rm.Spy = false
}
