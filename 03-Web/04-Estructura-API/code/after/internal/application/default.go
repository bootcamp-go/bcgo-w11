package application

import (
	"app/scaffolding/internal"
	"app/scaffolding/internal/handler"
	"app/scaffolding/internal/repository"
	"app/scaffolding/internal/service"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// NewDefaultHTTP creates a new instance of a default http server
func NewDefaultHTTP(addr string) *DefaultHTTP {
	// default config / values
	// ...

	return &DefaultHTTP{
		addr: addr,
	}
}

// DefaultHTTP is a struct that represents the default implementation of a http server
type DefaultHTTP struct {
	// addr is the address of the http server
	addr string
}

// Run runs the http server
func (h *DefaultHTTP) Run() (err error) {
	// initialize dependencies
	// - repository
	rp := repository.NewMovieMap(make(map[int]internal.Movie), 0)
	// - service
	sv := service.NewMovieDefault(rp)
	// - handler
	hd := handler.NewDefaultMovies(sv)
	// - router
	rt := chi.NewRouter()
	//   endpoints
	rt.Post("/movies", hd.Create())

	// run http server
	err = http.ListenAndServe(h.addr, rt)
	return
}