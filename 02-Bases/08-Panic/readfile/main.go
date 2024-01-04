package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// open a file
	f, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	// read bytes from file
	bytes, err := io.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		return
	}

	// print bytes
	data := string(bytes)
	fmt.Println(data)
}