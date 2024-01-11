package internal

import "errors"

var (
	// ErrMovieTitleAlreadyExists is the error returned when a movie title already exists
	ErrMovieTitleAlreadyExists = errors.New("movie title already exists")
	// ErrMovieNotFound is the error returned when a movie is not found
	ErrMovieNotFound = errors.New("movie not found")
)

// MovieRepository is an interface that represents a movie repository
// - database
// - cache
type MovieRepository interface {
	// Save saves a movie to the repository
	Save(movie *Movie) (err error)
	// GetByID returns a movie by id from the repository
	GetByID(id int) (movie Movie, err error)
	// Update updates a movie in the repository
	Update(movie *Movie) (err error)
	// Delete deletes a movie from the repository
	Delete(id int) (err error)
}
