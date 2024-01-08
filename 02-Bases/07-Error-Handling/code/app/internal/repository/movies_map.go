package repository

import (
	"app/error-handling/first/internal"
	"errors"
	"fmt"
)	

var (
	ErrInvalidMovie = errors.New("invalid movie")
	ErrConstraintViolation = errors.New("constraint violation")
	ErrLimitReached = errors.New("limit reached")
)

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
		err = fmt.Errorf("%w. limit: %d", ErrLimitReached, m.limit)
		err = fmt.Errorf("%w. extra info", err)
		err = fmt.Errorf("%w. extra info", err)
		err = fmt.Errorf("%w. extra info", err)
		err = fmt.Errorf("%w. extra info", err)
		err = fmt.Errorf("%w. extra info", err)
		err = fmt.Errorf("%w. extra info 17", err)
		err = fmt.Errorf("%w. extra info", err)
		return
	}

	// check logic business
	// - required fields
	if movie.Title == "" {
		err = fmt.Errorf("%w. field: title", ErrInvalidMovie)
		return
	}
	// - quality of fields
	if !(movie.Rating >= 0 && movie.Rating <= 10) {
		err = fmt.Errorf("%w. field: rating", ErrInvalidMovie)
		return
	}
	if movie.Year < 1900 || movie.Year > 2100 {
		err = fmt.Errorf("%w. field: year", ErrInvalidMovie)
		return
	}

	// check constraints db
	// - unique fields
	for _, m := range m.movies {
		if m.Title == movie.Title {
			err = ErrConstraintViolation
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