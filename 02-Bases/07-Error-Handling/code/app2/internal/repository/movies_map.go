package repository

import (
	"app/error-handling/second/internal"
	"fmt"
)

// ______________________________________________________________
// custom error
type ErrorMovieLimit struct {
	Msg string
	Limit int
}

func (e *ErrorMovieLimit) Error() string {
	return fmt.Sprintf("%s. limit: %d", e.Msg, e.Limit)
}


type ErrorMovieInvalid struct {
	Msg string
	Feature string
}

func (e *ErrorMovieInvalid) Error() string {
	return fmt.Sprintf("%s. field: %s", e.Msg, e.Feature)
}

// ______________________________________________________________
func NewMovieMap(limit int) *MovieMap {
	return &MovieMap{
		movies: make(map[int]internal.Movie),
		limit: limit,
	}
}

type MovieMap struct {
	movies map[int]internal.Movie
	lastId int	// auto-increment
	limit, count int
}

func (m *MovieMap) Save(movie *internal.Movie) (err error) {
	// check limit
	if m.count >= m.limit {
		err = &ErrorMovieLimit{
			Msg: "limit reached",
			Limit: m.limit,
		}
		err = fmt.Errorf("%w. extra info", err)
		err = fmt.Errorf("%w. extra info", err)
		err = fmt.Errorf("%w. extra info", err)
		err = fmt.Errorf("%w. extra info", err)
		return
	}

	// check logic business
	// - required fields
	if movie.Title == "" {
		// err = fmt.Errorf("%w. field: title", ErrInvalidMovie)
		err = &ErrorMovieInvalid{
			Msg: "invalid movie",
			Feature: "title",
		}
		err = fmt.Errorf("%w. extra info", err)
		return
	}
	// - quality of fields
	if !(movie.Rating >= 0 && movie.Rating <= 10) {
		// err = fmt.Errorf("%w. field: rating", ErrInvalidMovie)
		err = &ErrorMovieInvalid{
			Msg: "invalid movie",
			Feature: "rating",
		}
		err = fmt.Errorf("%w. extra info", err)
		return
	}
	if movie.Year < 1900 || movie.Year > 2100 {
		// err = fmt.Errorf("%w. field: year", ErrInvalidMovie)
		err = &ErrorMovieInvalid{
			Msg: "invalid movie",
			Feature: "year",
		}
		err = fmt.Errorf("%w. extra info", err)
		return
	}

	// check constraints db
	// - unique fields
	for _, m := range m.movies {
		if m.Title == movie.Title {
			// err = fmt.Errorf("%w. field: title", ErrConstraintViolation)
			err = &ErrorMovieInvalid{
				Msg: "constraint violation",
				Feature: "title",
			}
			err = fmt.Errorf("%w. extra info", err)
			return
		}
	}

	// save
	// - increment id
	(*m).lastId++
	// - set id
	(*movie).ID = (*m).lastId
	// - save
	m.movies[movie.ID] = *movie

	// increment count
	m.count++

	return
}