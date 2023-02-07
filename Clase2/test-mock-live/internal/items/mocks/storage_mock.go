package mocks

import (
	"testmock/internal/domain"
)

// constructor
func NewStorageMock() *storageMock {
	return &storageMock{}
}

// controller
type storageMock struct {
	Data domain.Item
	Err error
	Spy bool
}
func (st *storageMock) GetByName(name string) (i domain.Item, err error) {
	st.Spy = true

	if st.Err != nil {
		err = st.Err
		return
	}
	
	i = st.Data
	return
}
func (st *storageMock) UpdateByName(name string, i domain.Item) (err error) {
	st.Spy = true

	err = st.Err
	return
}

func (st *storageMock) Reset() {
	st.Data = domain.Item{}
	st.Err = nil
}