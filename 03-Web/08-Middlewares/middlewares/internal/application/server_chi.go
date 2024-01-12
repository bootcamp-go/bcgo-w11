package application

import (
	"middlewares/internal/handler"
	"middlewares/internal/middleware"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// NewServerCHI creates a new chi server
func NewServerCHI(addr string, token string) *ServerCHI {
	// default config
	defaultAddr := ":8080"
	defaultToken := ""
	if addr != "" {
		defaultAddr = addr
	}
	if token != "" {
		defaultToken = token
	}

	return &ServerCHI{
		addr: defaultAddr,
		token: defaultToken,
	}
}

// ServerCHI is a chi server
type ServerCHI struct {
	// addr is the address to listen on, ":http" if empty
	addr string
	// token is the token to authenticate requests
	token string
}

// Run starts the server
func (s *ServerCHI) Run() (err error) {
	// dependencies
	// - handler
	// - middleware
	auth := middleware.NewAuthenticator(s.token)
	logger := middleware.NewLogger()
	// - router
	rt := chi.NewRouter()
	//   middlewares (ping -> auth -> logger) decorator pattern
	rt.Use(logger.Log, auth.Auth)
	//   endpoints
	rt.Get("/ping", handler.Ping)

	// start server
	err = http.ListenAndServe(s.addr, rt)
	return
}
