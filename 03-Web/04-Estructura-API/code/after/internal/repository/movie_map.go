package repository

import "app/scaffolding/internal"

// NewMovieMap returns a new MovieMap instance
func NewMovieMap(db map[int]internal.Movie, lastId int) *MovieMap {
	// default config / values
	// ...

	return &MovieMap{
		db:     db,
		lastId: lastId,
	}
}

// MovieMap is a struct that represents a movie repository
type MovieMap struct {
	// db is a map that represents a database
	// - key: id of the movie
	// - value: movie
	db map[int]internal.Movie
	// lastId is the last id of the movie
	lastId int
}

func (m *MovieMap) Save(movie *internal.Movie) (err error) {
	// validate movie (consistency)
	// - title: unique
	for _, v := range (*m).db {
		if v.Title == movie.Title {
			return internal.ErrMovieTitleAlreadyExists
		}
	}

	// autoincrement
	// - increment id
	(*m).lastId++
	// - set id
	(*movie).ID = (*m).lastId

	// store movie
	(*m).db[(*movie).ID] = *movie

	return
}