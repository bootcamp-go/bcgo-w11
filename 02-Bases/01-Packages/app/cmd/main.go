package main

import (
	"app/packages/internal/calculator"
	"fmt"

	"github.com/fatih/color"
)

/*
	primite types
	- string
	- int
	- float64
	- bool

	based types
	- byte

	> explicit declaration
	> ...

	> implicit declaration
	type inference
	name := "John"	-> string
	age := 25		-> int
	height := 175.5 -> float64
	married := true -> bool
*/

func main() {
	// explicit declaration
	// - var
	var firstName string
	firstName = "John"

	var lastName string = "Doe"

	var hasDriverLicense bool = true

	// implicit declaration (type inference)
	// - var
	// - :=
	var age = 25    // int
	height := 175.5 // float64

	// show info
	fmt.Println("First Name:", firstName)
	fmt.Println("Last Name:", lastName)
	fmt.Println("Age:", age)
	fmt.Println("Height:", height)
	fmt.Println("Has Driver License:", hasDriverLicense)

	// use calculator package
	fmt.Println("Sum:", calculator.SumResult)
	fmt.Println("Sub:", calculator.SubResult)

	// use color package
	color.Red("The program has been finished")
}