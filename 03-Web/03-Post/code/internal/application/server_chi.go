package application

import (
	"code/post/internal"
	"code/post/internal/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// NewServerChi returns a new ServerChi instance
func NewServerChi(address string) *ServerChi {
	// default config / values
	defaultAddress := ":8080"
	if address != "" {
		defaultAddress = address
	}

	return &ServerChi{
		address: defaultAddress,
	}
}

// ServerChi is the server using chi
type ServerChi struct {
	// address is the address to listen on
	address string
}

// Run runs the server
func (s *ServerChi) Run() error {
	// dependencies
	// - db
	movies := make(map[int]internal.Movie, 0)
	lastID := 0
	// - handler
	hd := handlers.NewDefaultMovies(movies, lastID)
	// - router
	rt := chi.NewRouter()
	//   endpoints
	rt.Post("/movies", hd.Create())

	// run server
	return http.ListenAndServe(s.address, rt)
}