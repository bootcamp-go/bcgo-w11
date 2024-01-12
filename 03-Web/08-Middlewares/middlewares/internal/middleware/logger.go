package middleware

import (
	"fmt"
	"net/http"
	"time"
)

func NewLogger() *Logger {
	return &Logger{}
}

type Logger struct{}

func (l *Logger) Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("logger")
		
		// before
		start := time.Now()

		// call
		next.ServeHTTP(w, r)

		// after
		end := time.Since(start)
		fmt.Printf("request duration: %d", end.Nanoseconds())
	})
}