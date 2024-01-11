package repository

import "app/put/internal"

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
	if err = m.ValidateMovieTitle((*movie).Title); err != nil {
		return
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

func (m *MovieMap) GetByID(id int) (movie internal.Movie, err error) {
	movie, ok := m.db[id]
	if !ok {
		err = internal.ErrMovieNotFound
		return
	}

	return
}

func (m *MovieMap) ValidateMovieTitle(title string) (err error) {
	// validate movie (consistency)
	// - title: unique
	for _, v := range (*m).db {
		if v.Title == title {
			return internal.ErrMovieTitleAlreadyExists
		}
	}

	return
}

func (m *MovieMap) Update(movie *internal.Movie) (err error) {

	_, ok := m.db[(*movie).ID]
	if !ok {
		err = internal.ErrMovieNotFound
		return
	}

	m.db[(*movie).ID] = *movie
	return
}

func (m *MovieMap) Delete(id int) (err error) {
	_, ok := m.db[id]
	if !ok {
		err = internal.ErrMovieNotFound
		return
	}

	delete(m.db, id)

	return
}
