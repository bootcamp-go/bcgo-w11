package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// MyHandler is a handler with a map of users as data
type MyHandler struct {
	data map[string]string
}

// NewHandler returns a new MyHandler
func NewHandler() *MyHandler {
	return &MyHandler{
		data: map[string]string{
			"1": "Juan",
			"2": "Pedro",
			"3": "Maria",
			"4": "Jose",
		},
	}
}

// Get returns a handler for the GET / route
func (h *MyHandler) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	}
}

// MyResponse is an struct for the response
type MyResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// GetById returns a handler for the GET /users/{userId} route
func (h *MyHandler) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - get the id from the url
		id := chi.URLParam(r, "userId")

		// - get the user from the map
		name, ok := h.data[id]
		if !ok {
			code := http.StatusNotFound
			body := MyResponse{Message: "User not found", Data: nil}
			w.WriteHeader(code)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(body)
			return
		}

		// response
		code := http.StatusOK
		body := MyResponse{Message: "User found", Data: name}
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(body)
	}
}

// GetByQuery returns a handler for the GET /users?userId={id} route
func (h *MyHandler) GetByQuery() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - get the id from the query params
		id := r.URL.Query().Get("userId")
		hobby := r.URL.Query().Get("hobby")

		// - get the user from the map
		name, ok := h.data[id]
		if !ok {
			code := http.StatusNotFound
			body := MyResponse{Message: "User not found", Data: nil}
			w.WriteHeader(code)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(body)
			return
		}

		// response
		code := http.StatusOK
		body := MyResponse{Message: "User found", Data: name + " " + hobby}
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(body)
	}
}
