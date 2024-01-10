package internal

import "errors"

var (
	// ErrMovieTitleAlreadyExists is the error returned when a movie title already exists
	ErrMovieTitleAlreadyExists = errors.New("movie title already exists")
)

// MovieRepository is an interface that represents a movie repository
// - database
// - cache
type MovieRepository interface {
	// Save saves a movie to the repository
	Save(movie *Movie) (err error)
}