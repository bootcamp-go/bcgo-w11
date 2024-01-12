package handler

import "net/http"

// Ping is a handler for ping endpoint
func Ping(w http.ResponseWriter, r *http.Request) {
	// response
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}

// func Ping() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		// response
// 		w.Header().Set("Content-Type", "text/plain")
// 		w.WriteHeader(http.StatusOK)
// 		w.Write([]byte("pong"))
// 	}
// }