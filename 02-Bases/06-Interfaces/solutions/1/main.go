package main

import "fmt"

// Student is a representation of a student
type Student struct {
	FirstName    string
	LastName     string
	DNI          string
	RegisterDate string
}

// Info prints the information of a student
func (s Student) Info() {
	fmt.Printf("Name: %s\nLast Name: %s\nDNI: %s\nRegister Date: %s\n", s.FirstName, s.LastName, s.DNI, s.RegisterDate)
}

// Strings returns a string representation of a student (implements fmt.Stringer interface)
func (s Student) String() string {
	return fmt.Sprintf("Name: %s\nLast Name: %s\nDNI: %s\nRegister Date: %s\n", s.FirstName, s.LastName, s.DNI, s.RegisterDate)
}

func main() {
	// Create a new student
	student := Student{
		FirstName:    "John",
		LastName:     "Doe",
		DNI:          "12345678",
		RegisterDate: "01/01/2000",
	}

	// Print the student information
	fmt.Println("Student information:")
	student.Info()

	// Print the student information using fmt.Stringer interface
	fmt.Println("\nStudent information (fmt.Stringer):")
	fmt.Println(student)
}