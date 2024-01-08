package main

import (
	"fmt"
)

// Person is a struct that represents a person
type Person struct {
	// ID is the unique identifier of the person
	ID int
	// Name is the name of the person
	Name string
	// DateOfBirth is the date of birth of the person
	DateOfBirth string
}

// Employee is a struct that represents an employee
type Employee struct {
	// ID is the unique identifier of the employee
	ID int
	// Position is the position of the employee
	Position string
	// Person are the personal details of the employee
	Person
}

func (e Employee) Info() {
	fmt.Printf("Employee Info:\n- ID: %d\n- Position: %s\nPersonalInfo:\n- Name: %s\n- DateOfBirth: %v\n", e.ID, e.Position, e.Name, e.DateOfBirth)
}

// String returns a string representation of the employee
// - adapts to fmt.Stringer interface
func (e Employee) String() string {
	return fmt.Sprintf("Employee Info:\n- ID: %d\n- Position: %s\nPersonalInfo:\n- Name: %s\n- DateOfBirth: %v\n", e.ID, e.Position, e.Name, e.DateOfBirth)
}

func main() {
	// create a new person
	p := Person{
		ID:          1,
		Name:        "John",
		DateOfBirth: "1990-01-01",
	}

	// create a new employee
	e := Employee{
		ID:       1,
		Position: "Developer",
		Person:   p,
	}

	// print the employee (from the Info method)
	fmt.Println("Info method")
	e.Info()

	fmt.Println("\nString method")
	// print the employee (from the String method - fmt.Stringer interface)
	fmt.Println(e)
}