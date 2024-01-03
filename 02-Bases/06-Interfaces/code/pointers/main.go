package main

import "fmt"

type Person struct {
	Name *string
	Age int
}

func main() {
	// name is a variable of type string
	var name string = "John"
	// |
	// ptr is a variable of type *string (pointer to string)
	var ptr *string = &name // string: 0xc000010200
	// |
	// ptr2 -> shallow copy
	var ptr2 *string = ptr // string: 0xc000010200

	// update name
	UpdateName(ptr2)

	// Person
	p := Person{Name: new(string), Age: 20}
	(*p.Name) = "John"
	// -> shallow copy (i'm copying the same address)
	// both p and p1 are referencing the same address
	p1 := p
	newName := "Jane"
	(*p1.Name) = newName
	// -> deep copy (i'm creating a new address and copying the value)
	p2 := Person{Name: new(string), Age: 20}
	(*p2.Name) = *p.Name
	(*p2.Name) = "Gabriel"

	fmt.Println((*p.Name))
}

func ShowName(name *string) {
	fmt.Println(*name)
}

func UpdateName(name *string) {
	(*name) = "Jane"
}


/*
	slice: []int -> [1, 2, 3]			// slice: 0x0008 (sub-array: 0x0010)
	- len: 3
	- cap: 3

	slice2: shallow copy -> slice		// slice2:0x0009 (sub-array: 0x0010)
	- len: 3
	- cap: 3

	// event: exceed the capacity of the slice
	// - create a new array with double the capacity -> 0x0020
	slice: []int -> [1, 2, 3, 4, 5]		// slice: 0x0008 (sub-array: 0x0020) - re-allocated
	- len: 5
	- cap: 6

	slice2: shallow copy -> slice		// slice2:0x0009 (sub-array: 0x0010) - keeps old sub-array
*/