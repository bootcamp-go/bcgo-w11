package main

import (
	"fmt"
	"os"
)

func main() {
	text := []byte("lorem ipsum dolor sit amet")

	// write to buffer
	_, err := os.Stdout.Write(text)
	if err != nil {
		fmt.Println(err)
		return
	}
}