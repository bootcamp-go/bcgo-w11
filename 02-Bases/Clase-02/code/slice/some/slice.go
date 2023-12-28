package main

import "fmt"

func main() {
	// slice is a pointer to an array (3)
	// - len: 3
	// - cap: 3 (n > 1, then re-allocation duplicates the current capacity)
	
	// names := []string{}  		  // len: 0, cap: 0
	// names := make([]string, 0)	  // len: 0, cap: 0

	// names := []string{"", "", ""}  // len: 3, cap: 3
	names := make([]string, 3, 3)
	// names := make([]string, 3)	  // len: 3, cap: 3
	names[0] = "John"
	names[1] = "Paul"
	names[2] = "George"
	fmt.Printf("names: %v, len: %d, cap: %d\n", names, len(names), cap(names))

	// now lets add a new name, but capacity is full
	fmt.Println("adding a new name")
	names = append(names, "Ringo", "Pete")
	names = append(names, "keeps current capacity")
	fmt.Printf("names: %v, len: %d, cap: %d\n", names, len(names), cap(names))

	// lets add a new name exceeding capacity
	// - len: 6
	// - cap: 6 -> 12 (n > 1, then re-allocation doubles the current capacity)
	fmt.Println("adding a new name exceeding capacity")
	names = append(names, "Stuart")
	fmt.Printf("names: %v, len: %d, cap: %d\n", names, len(names), cap(names))
}