package main

import "fmt"

func main() {
	// default value
	var names map[string]int
	fmt.Println(names == nil)

	names = make(map[string]int)

	// add new name
	names["John"] = 18  // panic: assignment to entry in nil map if not initialized
	fmt.Printf("names: %v, len: %d\n", names, len(names))
}