package service

import (
	"app/put/internal"
	"fmt"
)

// NewMovieDefault creates a new instance of a movie service
func NewMovieDefault(rp internal.MovieRepository) *MovieDefault {
	return &MovieDefault{
		rp: rp,
	}
}

// MovieDefault is a struct that represents the default implementation of a movie service
type MovieDefault struct {
	// rp is a movie repository
	rp internal.MovieRepository
	// external services
	// ... (weather api, etc.)
}

// Save saves a movie
func (m *MovieDefault) Save(movie *internal.Movie) (err error) {
	// external services
	// ...

	// business logic
	// - validate required fields
	if err = ValidateMovie(movie); err != nil {
		return
	}

	// save movie
	err = m.rp.Save(movie)
	if err != nil {
		switch err {
		case internal.ErrMovieTitleAlreadyExists:
			err = fmt.Errorf("%w: title", internal.ErrMovieAlreadyExists)
		}
		return
	}

	return
}

// GetByID returns a movie by id
func (m *MovieDefault) GetByID(id int) (movie internal.Movie, err error) {
	// get movie
	movie, err = m.rp.GetByID(id)
	if err != nil {
		switch err {
		case internal.ErrMovieNotFound:
			err = fmt.Errorf("%w: id", internal.ErrMovieNotFound)
		}
		return
	}

	return
}

// ValidateMovie validates a movie
func ValidateMovie(movie *internal.Movie) (err error) {
	// - validate required fields
	if (*movie).Title == "" {
		return fmt.Errorf("%w: title", internal.ErrFieldRequired)
	}
	if (*movie).Year == 0 {
		return fmt.Errorf("%w: year", internal.ErrFieldRequired)
	}
	// - validate quality
	if !((*movie).Rating >= 0 && (*movie).Rating <= 10) {
		return fmt.Errorf("%w: rating", internal.ErrFieldQuality)
	}
	if !((*movie).Year >= 1888 && (*movie).Year <= 9999) {
		return fmt.Errorf("%w: year", internal.ErrFieldQuality)
	}

	return
}

// Update updates a movie
func (m *MovieDefault) Update(movie *internal.Movie) (err error) {
	// validate
	if err = ValidateMovie(movie); err != nil {
		return
	}

	// update movie
	err = m.rp.Update(movie)
	if err != nil {
		switch err {
		case internal.ErrMovieNotFound:
			err = fmt.Errorf("%w: id", internal.ErrMovieNotFound)
		}
		return
	}
	return
}

func (m *MovieDefault) Delete(id int) (err error) {
	// delete movie
	err = m.rp.Delete(id)
	if err != nil {
		switch err {
		case internal.ErrMovieNotFound:
			err = fmt.Errorf("%w: id", internal.ErrMovieNotFound)
		}
		return
	}
	return
}
