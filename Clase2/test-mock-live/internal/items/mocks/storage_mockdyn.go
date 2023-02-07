package mocks

import (
	"testmock/internal/domain"
)

// constructor
func NewStorageMockDyn() *storageMockDyn {
	return &storageMockDyn{}
}

// controller
type storageMockDyn struct {
	Data domain.Item
	Err map[string]error
	Spy map[string]bool
}
func (st *storageMockDyn) GetByName(name string) (i domain.Item, err error) {
	st.Spy["GetByName"] = true
	
	err = st.Err["GetByName"]
	if err != nil {
		return
	}

	i = st.Data
	return
}
func (st *storageMockDyn) UpdateByName(name string, i domain.Item) (err error) {
	st.Spy["UpdateByName"] = true

	err = st.Err["UpdateByName"]
	return
}

func (st *storageMockDyn) Reset() {
	st.Data = domain.Item{}
	st.Err = make(map[string]error)
	st.Spy = make(map[string]bool)
}