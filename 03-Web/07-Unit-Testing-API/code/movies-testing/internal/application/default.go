package application

import (
	"movies-testing/internal/handler"
	"movies-testing/internal/repository"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// NewDefault returns a new Default.
func NewDefault(addr string) *Default {
	return &Default{
		addr: addr,
	}
}

// Default is an implementation of the Default interface.
type Default struct {
	// addr is the address to listen on.
	addr string
}

// Run runs the application.
func (a *Default) Run() (err error) {
	// dependencies
	// - repository
	rp := repository.NewTaskMap(nil)
	// - handler
	hd := handler.NewTaskDefault(rp)
	// - router
	rt := chi.NewRouter()
	//   endpoints
	// tasks/1
	// m := map[string]string{
	// 	"id": "1",
	// }
	// ctx := context.WithValue(context.TODO(), "chi", m)
	// newR := r.WithContext(context.WithValue(r.Context(), "id", "1"))
	// newR := r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, m))
	rt.Get("/tasks/{id}", hd.GetTask())

	// run the server
	err = http.ListenAndServe(a.addr, rt)
	return
}