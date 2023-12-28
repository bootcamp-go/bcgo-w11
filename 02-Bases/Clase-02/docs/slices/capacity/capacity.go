package main

import "fmt"

func main() {
	// slice
	// - capacity: 0 (increases by 1)
	sl1 := make([]int, 0)
	sl1 = append(sl1, 1)
	fmt.Printf("sl1: %v - len: %d - cap: %d\n", sl1, len(sl1), cap(sl1))	// sl1: [1] - len: 1 - cap: 1

	// - capacity: 1 (increases by 1)
	sl2 := make([]int, 0, 1)
	sl2 = append(sl2, 1, 2)
	fmt.Printf("sl2: %v - len: %d - cap: %d\n", sl2, len(sl2), cap(sl2))	// sl2: [1 2] - len: 2 - cap: 2

	// - capacity: n (increases by doubling the current capacity)
	// > 2
	sl3 := make([]int, 0, 2)
	sl3 = append(sl3, 1, 2, 3)
	fmt.Printf("sl3: %v - len: %d - cap: %d\n", sl3, len(sl3), cap(sl3))	// sl3: [1 2 3] - len: 3 - cap: 4
	// > 5
	sl4 := make([]int, 1, 5)
	sl4 = append(sl4, 1, 2, 3, 4, 5, 6)
	fmt.Printf("sl4: %v - len: %d - cap: %d\n", sl4, len(sl4), cap(sl4))	// sl4: [1 1 2 3 4 5 6] - len: 7 - cap: 10
	sl4 = append(sl4, 7, 8, 9, 10, 11, 12)
	fmt.Printf("sl4: %v - len: %d - cap: %d\n", sl4, len(sl4), cap(sl4))	// sl4: [1 1 2 3 4 5 6 7 8 9 10 11 12] - len: 13 - cap: 20 (!10)
}