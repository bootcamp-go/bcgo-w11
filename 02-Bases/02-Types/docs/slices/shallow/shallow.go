package main

import (
	"fmt"
)

func main() {
	// ______________________________________________________________________
	// slice: original
	sl := make([]int, 3)

	// slice: copy
	// - shallow (copy with inner references): has same subyacent pointer to the array
	slShallowCopy := sl
	// - deep: (copy with new re-initialized references): has different subyacent pointer to the array (new array) with same values (copied)
	slDeepCopy := make([]int, len(sl))
	copy(slDeepCopy, slShallowCopy)

	// slice: reference
	// - not affected by dynamic re-allocation
	slRef := &sl
	display("initialization", SliceInt{"original", sl}, SliceInt{"shallow-copy", slShallowCopy}, SliceInt{"deep-copy", slDeepCopy}, SliceInt{"reference", *slRef})
	/*  output:
		initialization
		-original: [0 0 0]
		-shallow-copy: [0 0 0]
		-deep-copy: [0 0 0]
		-reference: [0 0 0]
	*/

	// ______________________________________________________________________
	// mutate original: no re-allocation
	sl[0] = 1
	sl[1] = 2
	sl[2] = 3
	display("mutate original: no re-allocation", SliceInt{"original", sl}, SliceInt{"shallow-copy", slShallowCopy}, SliceInt{"deep-copy", slDeepCopy}, SliceInt{"reference", *slRef})
	/*  output:
		mutate original: no re-allocation
		-original: [1 2 3]
		-shallow-copy: [1 2 3]
		-deep-copy: [0 0 0]
		-reference: [1 2 3]
	*/

	// ______________________________________________________________________
	// mutate original: re-allocation
	sl = append(sl, 4)
	display("mutate original: re-allocation", SliceInt{"original", sl}, SliceInt{"shallow-copy", slShallowCopy}, SliceInt{"deep-copy", slDeepCopy}, SliceInt{"reference", *slRef})
	/*  output:
		mutate original: re-allocation
		-original: [1 2 3 4]
		-shallow-copy: [1 2 3]
		-deep-copy: [0 0 0]
		-reference: [1 2 3 4]
	*/

	// ______________________________________________________________________
	// mutate shallow-copy
	slShallowCopy[0] = 10
	display("mutate shallow-copy", SliceInt{"original", sl}, SliceInt{"shallow-copy", slShallowCopy}, SliceInt{"deep-copy", slDeepCopy}, SliceInt{"reference", *slRef})
	/*  output:
		mutate shallow-copy
		-original: [1 2 3 4]
		-shallow-copy: [10 2 3]
		-deep-copy: [0 0 0]
		-reference: [1 2 3 4]
	*/
}

// SliceInt: slice of int
type SliceInt struct {
	name string
	slice []int
}
// display: print slices
func display(title string, slices ...SliceInt) {
	fmt.Println(title)
	for _, v := range slices {
		fmt.Printf("-%s: %v\n", v.name, v.slice)
	}
	fmt.Println()
}