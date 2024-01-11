package application

import (
	"app/put/internal"
	"app/put/internal/handler"
	"app/put/internal/repository"
	"app/put/internal/service"
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
	// rt.Get("/movies/{id}", hd.GetByID())
	// rt.Post("/movies", hd.Create())
	// rt.Put("/movies/{id}", hd.Update())

	rt.Route("/movies", func(rt chi.Router) {
		rt.Get("/{id}", hd.GetByID())
		rt.Post("/", hd.Create())
		rt.Put("/{id}", hd.Update())
		rt.Patch("/{id}", hd.UpdatePartial())
		rt.Delete("/{id}", hd.Delete())
	})
	// run http server
	err = http.ListenAndServe(h.addr, rt)
	return
}
