package service

import (
	"app/scaffolding/internal"
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