package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name     string  `json:"name"`
	Age      int     `json:"age"`
	Height   float64 `json:"height,omitempty"`
	Weight   float64 `json:"weight"`
	Password string  `json:"-"`
}

func main() {
	// create person struct
	p := Person{
		Name:     "john",
		Age:      30,
		Height:   0.0,
		Weight:   80.5,
		Password: "123456",
	}

	// convert into json bytes.
	bytes, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
		return
	}

	// convert into json string.
	data := string(bytes)
	fmt.Println(data)
}
