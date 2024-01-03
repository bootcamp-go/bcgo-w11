package main

import "fmt"

type Preferences struct {
	Drinks []string
	Foods  []string
}

type Person struct {
	Name  string
	Age   int
	Likes Preferences
}

func main() {
	// initialize by default
	// - zero values
	var p Person // p := Person{}
	fmt.Printf("%#v\n", p)

	// write-operations
	fmt.Println("\nwrite-operations")
	p.Name = "John"
	p.Age = 42
	p.Likes.Drinks = []string{"Coffee", "Tea"}
	p.Likes.Foods = []string{"Burger", "Pizza"}
	// p.Likes = Preferences{
	// 	Drinks: []string{"Coffee", "Tea"},
	// 	Foods:  []string{"Burger", "Pizza"},
	// }
	
	// update name (copy)
	UpdateName(p)
	
	fmt.Printf("%#v\n", p)
	
	// read-operations
	fmt.Println("\nread-operations")
	fmt.Println("name -> ", p.Name)
}

func UpdateName(p Person) {
	p.Name = "Jane"
}