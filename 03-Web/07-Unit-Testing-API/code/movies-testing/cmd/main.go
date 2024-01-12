package main

import (
	"fmt"
	"movies-testing/internal/application"
)

func main() {
	// env
	// ...

	// application
	// - config
	app := application.NewDefault(":8080")
	// - run
	if err := app.Run(); err != nil {
		fmt.Println(err)
		return
	}
}