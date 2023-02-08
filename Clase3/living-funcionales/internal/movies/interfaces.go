package movies

import (
	"api/internal/models"
	"errors"
)

// ____________________________________________________________
// Storage: access to consistent data
type Storage interface {
	// Create: create a new movie
	Read() (mv []models.Movie, err error)
	Write(mv []models.Movie) (err error)
}

var (
	ErrStorageInternal = errors.New("internal storage error")
)

// ____________________________________________________________
// Repository: access to storage
type Repository interface {
	// Create: create a new movie
	Create(movie models.Movie) (id string, err error)
}

var (
	ErrRepositoryInternal  = errors.New("internal repository error")
	ErrRepositoryNotUnique = errors.New("movie already exists")
)

// ____________________________________________________________
// Service: business logic
type Service interface {
	// Create: create a new movie
	Create(title string, rating float64, year int) (mv models.Movie, err error)
}

var (
	ErrServiceInternal  = errors.New("internal service error")
	ErrServiceInvalid   = errors.New("invalid movie")
	ErrServiceNotUnique = errors.New("movie already exists")
)