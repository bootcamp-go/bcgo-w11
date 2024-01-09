package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	// without subrouter
	r.Get("/users/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("users"))
	})
	r.Get("/users/{userId}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("user by id"))
	})
	r.Get("/users/{userId}/posts", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("posts by user id"))
	})

	// with subrouter
	r.Route("/users", func(r chi.Router) {
		// users
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("users"))

		})
		// user by id
		r.Get("/{userId}", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("user by id"))
		})
		// posts by user id
		r.Get("/{userId}/posts", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("posts by user id"))
		})

	})

	http.ListenAndServe(":8080", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome"))
}

// Request -> Handler -> Service -> Repository -> Map
