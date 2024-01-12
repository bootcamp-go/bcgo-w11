package main

import (
	"fmt"
	"middlewares/internal/application"
	"os"
)

func main() {
	// env
	addr := os.Getenv("ADDR")
	token := os.Getenv("TOKEN")

	// app
	// - config
	app := application.NewServerCHI(addr, token)
	// - run
	if err := app.Run(); err != nil {
		fmt.Println(err)
		return
	}
}