package middleware

import (
	"fmt"
	"net/http"
)

// NewAuthenticator creates a new authenticator
func NewAuthenticator(token string) *Authenticator {
	return &Authenticator{
		token: token,
	}
}

// Authenticator is a middleware that authenticates the request
type Authenticator struct {
	token string
}

func (a *Authenticator) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("auth")

		// logic before
		token := r.Header.Get("Authorization")
		if token != a.token {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// call next
		next.ServeHTTP(w, r)

		// logic after
		// ...
	})
}