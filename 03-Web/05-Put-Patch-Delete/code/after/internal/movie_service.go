package internal

import (
	"errors"
)

var (
	ErrFieldRequired      = errors.New("field required")
	ErrFieldQuality       = errors.New("field quality")
	ErrMovieAlreadyExists = errors.New("movie already exists")
)

// MovieService is an interface that represents a movie service
// - business logic
// - validation
// - external services (e.g. apis, databases, etc.)
type MovieService interface {
	// Save saves a movie
	Save(movie *Movie) (err error)
	// GetByID returns a movie by id
	GetByID(id int) (movie Movie, err error)
	// Update updates a movie
	Update(movie *Movie) (err error)
	// Delete deletes a movie by id
	Delete(id int) (err error)
}
