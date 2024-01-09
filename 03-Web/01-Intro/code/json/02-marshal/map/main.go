package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	m := map[string]any{
		"name":   "john",
		"age":    30,
		"height": 2.0,
		"wight":  80.5,
	}

	// convert into json bytes.
	bytes, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
		return
	}

	// convert into json string.
	data := string(bytes)
	fmt.Println(data)
}