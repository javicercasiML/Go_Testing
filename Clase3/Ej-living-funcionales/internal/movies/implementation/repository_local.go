package implementation

import (
	"api/internal/models"
	"api/internal/movies"

	"github.com/google/uuid"
)

// constructor: create a new repository
func NewRepositoryLocal(st movies.Storage) movies.Repository {
	return &repositoryLocal{st: st}
}

// controller
type repositoryLocal struct {
	st movies.Storage
}
func (r *repositoryLocal) Create(movie models.Movie) (id string, err error) {
	// get movies
	var mv []models.Movie
	mv, err = r.st.Read()
	if err != nil {
		err = movies.ErrRepositoryInternal
		return
	}

	// check if movie already exists
	for _, m := range mv {
		if m.Title == movie.Title {
			err = movies.ErrRepositoryNotUnique
			return
		}
	}

	// default movie
	id = uuid.NewString()
	movie.ID = id

	// add and save movie
	mv = append(mv, movie)
	
	if err = r.st.Write(mv); err != nil {
		err = movies.ErrRepositoryInternal
	}

	return
}