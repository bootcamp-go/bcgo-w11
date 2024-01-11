package handler

import (
	"app/put/internal"
	"app/put/platform/web/request"
	"app/put/platform/web/response"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// NewDefaultMovies returns a new DefaultMovies instance
func NewDefaultMovies(sv internal.MovieService) *DefaultMovies {
	return &DefaultMovies{
		sv: sv,
	}
}

// DefaultMovies is an implementation with handlers for the Movies storage
type DefaultMovies struct {
	// sv is a movie service
	sv internal.MovieService
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
// func (d *DefaultMovies) Create(w http.ResponseWriter, r *http.Request) {

// }
func (d *DefaultMovies) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		var body BodyRequestMovieJSON
		if err := request.JSON(r, &body); err != nil {
			response.Text(w, http.StatusBadRequest, "invalid body")
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
		// - save movie
		if err := d.sv.Save(&movie); err != nil {
			switch {
			case errors.Is(err, internal.ErrFieldRequired), errors.Is(err, internal.ErrFieldQuality):
				// w.Header().Set("Content-Type", "text/plain")
				// w.WriteHeader(http.StatusBadRequest)
				// w.Write([]byte("invalid body"))
				response.Text(w, http.StatusBadRequest, "invalid body")
			case errors.Is(err, internal.ErrMovieAlreadyExists):
				response.Text(w, http.StatusConflict, "movie already exists")
			default:
				response.Text(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		// response
		// - deserialize MovieJSON
		data := MovieJSON{
			ID:        movie.ID,
			Title:     movie.Title,
			Year:      movie.Year,
			Rating:    movie.Rating,
			Published: movie.Published,
		}
		// w.Header().Set("Content-Type", "application/json")
		// w.WriteHeader(http.StatusCreated)
		// json.NewEncoder(w).Encode(map[string]any{
		// 	"message": "movie created",
		// 	"data": data,
		// })
		response.JSON(w, http.StatusCreated, map[string]any{
			"message": "movie created",
			"data":    data,
		})
	}
}

// GetByID returns a movie by id
func (d *DefaultMovies) GetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - get id from path
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Text(w, http.StatusBadRequest, "invalid id")
			return
		}

		// process
		// - get movie
		movie, err := d.sv.GetByID(id)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrMovieNotFound):
				response.Text(w, http.StatusNotFound, "movie not found")
			default:
				response.Text(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		// response
		// - serialize MovieJSON
		data := MovieJSON{
			ID:        movie.ID,
			Title:     movie.Title,
			Year:      movie.Year,
			Rating:    movie.Rating,
			Published: movie.Published,
		}
		// - response
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "movie found",
			"data":    data,
		})
	}
}

// Update updates a movie
func (d *DefaultMovies) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - get id from path
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Text(w, http.StatusBadRequest, "invalid id")
			return
		}

		// - get body to []byte
		bytes, err := io.ReadAll(r.Body)
		if err != nil {
			response.Text(w, http.StatusBadRequest, "invalid body")
			return
		}

		// - get body to map[string]any
		var bodyMap map[string]any
		if err := json.Unmarshal(bytes, &bodyMap); err != nil {
			response.Text(w, http.StatusBadRequest, "invalid body")
			return
		}

		// - validate body
		if err := ValidateKeyExistante(bodyMap, "title", "year", "rating", "published"); err != nil {
			response.Text(w, http.StatusBadRequest, "invalid body")
			return
		}

		// - get body
		var body BodyRequestMovieJSON
		if err := json.Unmarshal(bytes, &body); err != nil {
			// if err := request.JSON(r, &body); err != nil {
			response.Text(w, http.StatusBadRequest, "invalid body")
			return
		}

		// process
		// - serialize internal.Movie
		movie := internal.Movie{
			ID:        id,
			Title:     body.Title,
			Year:      body.Year,
			Rating:    body.Rating,
			Published: body.Published,
		}

		// - update movie
		if err := d.sv.Update(&movie); err != nil {
			switch {
			case errors.Is(err, internal.ErrMovieNotFound):
				response.Text(w, http.StatusNotFound, "movie not found")
			case errors.Is(err, internal.ErrFieldRequired), errors.Is(err, internal.ErrFieldQuality):
				response.Text(w, http.StatusBadRequest, "invalid body")
			default:
				response.Text(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		// response
		// - deserialize MovieJSON
		data := MovieJSON{
			ID:        movie.ID,
			Title:     movie.Title,
			Year:      movie.Year,
			Rating:    movie.Rating,
			Published: movie.Published,
		}

		// - response
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "movie updated",
			"data":    data,
		})
	}
}

// UpdatePartial updates a movie
func (d *DefaultMovies) UpdatePartial() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - get id from path
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Text(w, http.StatusBadRequest, "invalid id")
			return
		}

		// - get the movie from the service
		movie, err := d.sv.GetByID(id)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrMovieNotFound):
				response.Text(w, http.StatusNotFound, "movie not found")
			default:
				response.Text(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		// process
		// - serialize internal.Movie to BodyRequestMovieJSON
		reqBody := BodyRequestMovieJSON{
			Title:     movie.Title,
			Year:      movie.Year,
			Rating:    movie.Rating,
			Published: movie.Published,
		}

		// - get body
		if err := request.JSON(r, &reqBody); err != nil {
			response.Text(w, http.StatusBadRequest, "invalid body")
			return
		}

		// - serialize internal.Movie
		movie = internal.Movie{
			ID:        id,
			Title:     reqBody.Title,
			Year:      reqBody.Year,
			Rating:    reqBody.Rating,
			Published: reqBody.Published,
		}

		// - update movie
		if err := d.sv.Update(&movie); err != nil {
			switch {
			case errors.Is(err, internal.ErrMovieNotFound):
				response.Text(w, http.StatusNotFound, "movie not found")
			case errors.Is(err, internal.ErrFieldRequired), errors.Is(err, internal.ErrFieldQuality):
				response.Text(w, http.StatusBadRequest, "invalid body")
			default:
				response.Text(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		// response
		// - deserialize MovieJSON
		data := MovieJSON{
			ID:        id,
			Title:     reqBody.Title,
			Year:      reqBody.Year,
			Rating:    reqBody.Rating,
			Published: reqBody.Published,
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "movie updated",
			"data":    data,
		})

	}
}

func ValidateKeyExistante(mp map[string]any, keys ...string) (err error) {
	for _, k := range keys {
		if _, ok := mp[k]; !ok {
			return fmt.Errorf("key %s not found", k)
		}
	}
	return
}

// Delete deletes a movie
func (d *DefaultMovies) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - get id from path
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": "invalid id",
				"data":    nil,
			})
			return
		}

		// proces
		// - delete movie
		if err := d.sv.Delete(id); err != nil {
			switch {
			case errors.Is(err, internal.ErrMovieNotFound):
				response.Text(w, http.StatusNotFound, "movie not found")
			default:
				response.Text(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		// response
		response.Text(w, http.StatusNoContent, "movie deleted")
	}
}
