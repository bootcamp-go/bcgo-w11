package main

import "fmt"

func main() {
	json := map[string]any{
		"name":    "John",
		"age":     30,
		"height":  1.8,
		"married": true,
	}

	// get age
	age := json["age"]

	// type assertion -> int
	ageInt, ok := age.(int)
	if !ok {
		fmt.Println("age is not an int")
		return
	}

	// increment age
	ageInt++

	// set age
	json["age"] = ageInt

	fmt.Println(json)
}