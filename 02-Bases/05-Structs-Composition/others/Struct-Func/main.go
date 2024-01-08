package main

import "fmt"

type Operation struct {
	Name string
	Do   func()
}

func main() {
	// default
	var op Operation
	fmt.Println(op.Name)
	fmt.Println(op.Do)

	ShowTime := Operation{
		Name: "ShowTime",
		Do: func() {
			println("ShowTime")
		},
	}
	ShowTime.Do()

	ShowDate := Operation{
		Name: "ShowDate",
		Do: func() {
			println("ShowDate")
		},
	}
	ShowDate.Do()
}