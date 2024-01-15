package main

import "fmt"

func main() {
	m := map[string]any{
		"foo": "bar",
		"bar": 1,
	}

	shallow := m

	shallow["foo"] = "baz"

	fmt.Println(m["foo"])
}