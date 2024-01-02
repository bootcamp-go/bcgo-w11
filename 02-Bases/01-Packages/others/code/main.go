package main

import "fmt"

// GlobalVariable is a global variable
var GlobalVariable = "Global Variable"

func main() {
	// fullName is a local variable

	result := Sum(1, 2)
	fmt.Println(result)
}

func Sum(a, b int) int {
	return a + b
}