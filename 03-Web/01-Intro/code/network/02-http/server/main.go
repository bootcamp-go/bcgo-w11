package main

import (
	"fmt"
	"net/http"
)

func main() {
	// register endpoints
	// - pong
	handlerPong := func (w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			w.Write([]byte("get"))
		case "POST":
			w.Write([]byte("post"))
		}
		w.Write([]byte("pong"))
	}
	// - register endpoint ping
	http.HandleFunc("/ping", handlerPong)

	// - handler2
	handler2 := func (w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	}
	// - register endpoint with handler2
	http.HandleFunc("/hello", handler2)

	uniqueHandler := func (w http.ResponseWriter, r *http.Request) {
		// detect method
		switch r.Method {
		case "GET":
			// search path (ping, hello) registered by method (GET)
		case "POST":
			// on this method, search corresponding path (ping, hello)
		}
	}
	http.HandleFunc("/", uniqueHandler)


	// listen and serve
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
		return
	}
}