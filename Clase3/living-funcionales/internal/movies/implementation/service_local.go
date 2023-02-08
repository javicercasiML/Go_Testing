package implementation

import (
	"api/internal/models"
	"api/internal/movies"
	"errors"
)

// constructor: create a new service
func NewServiceLocal(rp movies.Repository) movies.Service {
	return &serviceLocal{rp: rp}
}

// controller
type serviceLocal struct {
	rp movies.Repository
}
func (s *serviceLocal) Create(title string, rating float64, year int) (mv models.Movie, err error) {
	// prepare movie
	movie := models.Movie{
		Title:  title,
		Rating: rating,
		Year:   year,
	}

	// validate movie
	if !movie.Valid() {
		err = movies.ErrServiceInvalid
		return
	}

	// create movie
	var id string
	id, err = s.rp.Create(movie)
	if err != nil {
		if errors.Is(err, movies.ErrRepositoryNotUnique) {
			err = movies.ErrServiceNotUnique
		} else {
			err = movies.ErrServiceInternal
		}
		return
	}

	movie.ID = id
	mv = movie
	return
}