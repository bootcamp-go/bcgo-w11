package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Person struct {
	Name   string	`json:"name"`
	Age    int		`json:"age"`
	Cars   []string	`json:"cars"`
}

func main() {
	// strings reader
	reader := strings.NewReader(`{"name":"John","age":30,"cars":["Ford","BMW","Fiat"]}`)

	dc := json.NewDecoder(reader)

	// decode into person struct
	var p Person
	if err := dc.Decode(&p); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", p)
}