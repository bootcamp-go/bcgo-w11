package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name string   `json:"name"`
	Age  int      `json:"age"`
	Cars []string `json:"cars"`
}

func main() {
	// create person struct
	p := Person{
		Name: "john",
		Age:  30,
		Cars: []string{"Ford", "BMW", "Fiat"},
	}

	// create encoder
	ec := json.NewEncoder(os.Stdout)
	if err := ec.Encode(p); err != nil {
		fmt.Println(err)
		return
	}
}