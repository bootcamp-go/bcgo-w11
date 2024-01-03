package main

import (
	"fmt"
	"strconv"
)

func main() {
	number, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println(number)
}