package main

import (
	"app/interfaces/internal"
	"app/interfaces/internal/storage"
	"fmt"
)

func main() {
	// env
	// ...

	// app
	var apple internal.StorageApple = storage.NewAppleMap(nil, 0)

	a, err := apple.GetApple(1)
	if err != "" {
		fmt.Println(err)
		return
	}

	fmt.Println(a)
}