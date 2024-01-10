package handlers

import (
	"code/post/internal"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// NewDefaultMovies returns a new DefaultMovies instance
func NewDefaultMovies(movies map[int]internal.Movie, lastId int) *DefaultMovies {
	// default config / values
	// ...

	return &DefaultMovies{
		movies: movies,
		lastID: lastId,
	}
}

// DefaultMovies is an implementation with handlers for the Movies storage
type DefaultMovies struct {
	// movies is a map of movie ID to Movie
	// - key: id of the movie
	// - value: movie
	movies map[int]internal.Movie
	// lastID is the last used movie ID
	lastID int
}

// MovieJSON is the movie in JSON format
type MovieJSON struct {
	ID        int     `json:"id"`
	Title     string  `json:"title"`
	Year      int     `json:"year"`
	Rating    float64 `json:"rating"`
	Published bool    `json:"published"`
}

// BodyRequestMovieJSON is the body of the request for a movie in JSON format
type BodyRequestMovieJSON struct {
	Title     string  `json:"title"`
	Year      int     `json:"year"`
	Rating    float64 `json:"rating"`
	Published bool    `json:"published"`
}

// Create creates a new movie
func (d *DefaultMovies) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// token
		token := r.Header.Get("Authorization")
		if token != "123456" {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("invalid token"))
			return
		}

		// request
		// - validate if client request (unstructured data - json is dynamic) adapts to the structured data (struct - static)
		// this validation must be made at dynamic level (unstructured data - json) and not at static level (struct) due to existance or absence of fields
		// bytes, err := io.ReadAll(r.Body)
		// if err != nil {
		// 	w.Header().Set("Content-Type", "text/plain")
		// 	w.WriteHeader(http.StatusBadRequest)
		// 	w.Write([]byte("invalid body"))
		// 	return
		// }
		// var mp map[string]any
		// if err := json.Unmarshal(bytes, &mp); err != nil {
		// 	w.Header().Set("Content-Type", "text/plain")
		// 	w.WriteHeader(http.StatusBadRequest)
		// 	w.Write([]byte("invalid body"))
		// 	return
		// }
		// if err := ValidateKeyExistance(mp, "title", "year", "rating", "published"); err != nil {
		// 	w.Header().Set("Content-Type", "text/plain")
		// 	w.WriteHeader(http.StatusBadRequest)
		// 	w.Write([]byte("invalid body" + err.Error()))
		// 	return
		// }

		// - parse request body (from json/dynamic to struct/static)
			// contentType := r.Header.Get("Content-Type")
			// switch contentType {
			// case "application/json":
			// 	// decoding json to struct
			// case "application/xml":
			// 	// decoding xml to struct
			// case "application/x-www-form-urlencoded":
			// 	// decoding form to struct
			// default:
			// 	// ...
			// }
		var body BodyRequestMovieJSON
		// if err := json.Unmarshal(bytes, &body); err != nil {
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("invalid body"))
			return
		}

		// process
		// - serialize internal.Movie
		movie := internal.Movie{
			Title:     body.Title,
			Year:      body.Year,
			Rating:    body.Rating,
			Published: body.Published,
		}
		// - autoincrement id
		(*d).lastID++
		// - set id
		movie.ID = (*d).lastID
		
		// - validate business rules
		if err := ValidateBussinessRule(&movie); err != nil {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("invalid movie"))
			return
		}
		// - validate if movie already exists
		for _, mv := range (*d).movies {
			if mv.Title == movie.Title {
				w.Header().Set("Content-Type", "text/plain")
				w.WriteHeader(http.StatusConflict)
				w.Write([]byte("movie already exists"))
				return
			}
		}

		// - store movie
		(*d).movies[movie.ID] = movie

		// response
		// - serialize MovieJSON
		data := MovieJSON{
			ID:        movie.ID,
			Title:     movie.Title,
			Year:      movie.Year,
			Rating:    movie.Rating,
			Published: movie.Published,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]any{
			"message": "movie created",
			"data": data,
		})
	}
}

func ValidateKeyExistance(mp map[string]any, keys ...string) error {
	for _, key := range keys {
		if _, ok := mp[key]; !ok {
			return fmt.Errorf("key %s does not exist", key)
		}
	}
	return nil
}

func ValidateBussinessRule(mv *internal.Movie) error {
	// validate movie (business rules)
	// - required
	if mv.Title == "" {
		return errors.New("title is required")
	}
	if mv.Year == 0 {
		return errors.New("year is required")
	}
	// - quality
	if len(mv.Title) < 3 || len(mv.Title) > 150 {
		return errors.New("title must be between 3 and 150 characters")
	}
	if mv.Year < 1888 || mv.Year > 2021 {
		return errors.New("year must be between 1888 and 2021")
	}
	// regex ...

	return nil
}