package main

func main() {
	// apples
	apples := [2]string{"red", "green"}

	// read operation
	println("first apple", apples[0])
	println("second apple", apples[1])

	// write operation
	// green := apples[1]	// does not work
	// green = "yellow"

	apples[1] = "yellow"
	println("update second apple")
	println("second apple", apples[1])
}