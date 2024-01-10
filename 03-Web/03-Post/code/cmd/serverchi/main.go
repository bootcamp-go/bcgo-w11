package main

import (
	"code/post/internal/application"
	"fmt"
)

func main() {
	// env
	// ...

	// app
	// - config
	app := application.NewServerChi("localhost:8080")
	// - run
	if err := app.Run(); err != nil {
		fmt.Println(err)
		return
	}
}