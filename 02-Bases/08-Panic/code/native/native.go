package main

import "fmt"

func main() {
	// PARTS OF A PANIC
	// word panic + attempt action + error message

	// panic: runtime error: index out of range
	s := []int{1, 2, 3}
	fmt.Println(s[3])

	// panic: runtime error: integer divide by zero
	a, b := 1, 0
	fmt.Println(a / b)

	// panic: assignment to entry in nil map
	var mapa map[string]string
	mapa["hola"] = "mundo"

	// panic: runtime error: invalid memory address or nil pointer dereference
	var p *int
	*p = 1

	// panic: interface conversion: interface {} is int, not string
	var i interface{} = 1
	str := i.(string)
	fmt.Println(str)

}
