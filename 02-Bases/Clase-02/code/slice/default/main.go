package main

import "fmt"

func main() {
	// default value
	var sl []int
	fmt.Println(sl == nil)
	fmt.Printf("sl: %v, len: %d, cap: %d\n", sl, len(sl), cap(sl))

	// add new number
	sl = append(sl, 1)
	fmt.Printf("sl: %v, len: %d, cap: %d\n", sl, len(sl), cap(sl))
}