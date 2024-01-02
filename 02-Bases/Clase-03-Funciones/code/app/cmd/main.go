package main

import (
	"app/internal/calculator"
	"fmt"
)

func main() {
	// numbers
	n1, n2 := 10.0, 5.0
	
	// sum
	// - fetch the addition function
	addition, err := calculator.Orchestrator(calculator.AdditionSymbol)
	if err != "" {
		fmt.Println(err)
		return
	}
	// - add the numbers
	result, err := addition(n1, n2)
	if err != "" {
		fmt.Println(err)
		return
	}

	fmt.Println(result)

	// divide
	// - fetch the division function
	division, err := calculator.Orchestrator(calculator.DivisionSymbol)
	if err != "" {
		fmt.Println(err)
		return
	}
	// - divide the result by 1
	result, err = division(result, 1)
	if err != "" {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}

// AnonimusFunction returns a function that returns a string
func AnonimusFunction() func(firstName, lastName string) (fullname string) {
	return func(firstName, lastName string) (fullname string) {
		fullname = firstName + " " + lastName
		return
	}
}

// ClosureFunction returns a function that returns a string
func ClosureFunction() func(firstName, lastName string) (fullname string) {
	// separator
	separator := "-"

	return func(firstName, lastName string) (fullname string) {
		// default values
		defaultFirstName := "John"
		defaultLastName := "Doe"
		if firstName != "" {
			defaultFirstName = firstName
		}
		if lastName != "" {
			defaultLastName = lastName
		}

		fullname = defaultFirstName + separator + defaultLastName
		return
	}
}