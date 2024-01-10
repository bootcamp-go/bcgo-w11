package response

import "net/http"

// Text writes text response
func Text(w http.ResponseWriter, code int, body string) {
	// set header
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	// set status code
	w.WriteHeader(code)

	// write body
	w.Write([]byte(body))
}