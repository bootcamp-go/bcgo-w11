package main

import (
	"fmt"
	"os"
)

func main() {
	// start main
	fmt.Println("Iniciando")

	// defer function to recover from panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recuperado en main", r)
		}
	}()

	// open file
	_, err := os.Open("archivo.txt")
	if err != nil {
		panic(err)
	}

	// this line is not executed
	fmt.Println("Finalizando")
}
