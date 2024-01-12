package main

import "os"

func main() {
	// read enviroment variable
	goPath := os.Getenv("GOPATH")
	println(goPath)

	// set enviroment variable
	os.Setenv("KEY", "value")
	println(os.Getenv("KEY"))
}