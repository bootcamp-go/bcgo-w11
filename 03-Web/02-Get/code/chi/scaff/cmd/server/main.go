package main

import (
	"net/http"
	"scaff/cmd/server/handler"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	h := handler.NewHandler()

	r.Get("/", h.Get())
	r.Get("/users/{userId}", h.GetById())
	r.Get("/users", h.GetByQuery())

	http.ListenAndServe(":8080", r)
}
