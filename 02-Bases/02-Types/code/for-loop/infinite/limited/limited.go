package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// infinite (5 iterations)
	i := 0
	for i < 5 {
		i++
	}

	// limited (5 iterations)
	for i := 0; i < 5; i++ {
		fmt.Println("iteration:", i)
		fmt.Println("this is a random number:", rand.Intn(10))
	}
}