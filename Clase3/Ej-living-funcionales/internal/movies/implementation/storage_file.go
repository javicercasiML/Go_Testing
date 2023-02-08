package implementation

import (
	"api/internal/models"
	"api/internal/movies"
	"encoding/json"
	"os"
)

// constructor: create a new storage
func NewStorageFile(file string) movies.Storage {
	return &storageFile{file: file}
}

// controller
type storageFile struct {
	file string
}
func (s *storageFile) Read() (mv []models.Movie, err error) {
	// open file
	var f *os.File
	f, err = os.Open(s.file)
	if err != nil {
		err = movies.ErrStorageInternal
		return
	}
	defer f.Close()

	// parse
	dec := json.NewDecoder(f)
	if err = dec.Decode(&mv); err != nil {
		err = movies.ErrStorageInternal
	}
	
	return
}
func (s *storageFile) Write(mv []models.Movie) (err error) {
	// open file to overwrite content
	var f *os.File
	f, err = os.OpenFile(s.file, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		err = movies.ErrStorageInternal
		return
	}

	// parse
	enc := json.NewEncoder(f)
	if err = enc.Encode(mv); err != nil {
		err = movies.ErrStorageInternal
	}

	return
}