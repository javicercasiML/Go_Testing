package mocks

import (
	"api/internal/models"

	"github.com/stretchr/testify/mock"
)

// constructor
func NewStorageMock() *storageMock {
	return &storageMock{}
}

// controller
type storageMock struct {
	mock.Mock
}
func (s *storageMock) Read() (mv []models.Movie, err error) {
	args := s.Called()
	mv = args.Get(0).([]models.Movie)
	err = args.Error(1)
	return
}
func (s *storageMock) Write(mv []models.Movie) (err error) {
	args := s.Called(mv)
	err = args.Error(0)
	return
}

// reset
func (s *storageMock) Reset() {
	s.Mock = mock.Mock{}
}