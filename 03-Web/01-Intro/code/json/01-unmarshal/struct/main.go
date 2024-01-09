package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name   string	`json:"name"`
	Age    int		`json:"age"`
	Height float64	`json:"height"`
}

// type Person struct {
// 	Name   string
// 	Age    int
// 	Height float64
// 	Date   time.Time
// }

// unmarshaller interface
// func (p *Person) UnmarshalJSON(b []byte) error {
// 	// create temp struct
// 	type temp struct {
// 		Name   string
// 		Age    int
// 		Height float64
// 		Date   string
// 	}

// 	// unmarshal into temp struct
// 	var t temp
// 	if err := json.Unmarshal(b, &t); err != nil {
// 		return err
// 	}

// 	// convert temp struct into Person struct
// 	p.Name = t.Name
// 	p.Age = t.Age
// 	p.Height = t.Height

// 	// convert date string into time.Time
// 	date, err := time.Parse("2006-01-02", t.Date)
// 	if err != nil {
// 		return err
// 	}
// 	p.Date = date

// 	return nil
// }



func main() {
	// json bytes
	bytes := []byte(`{"name":"john","age":30,"Height":2.0,"weight":null}`)

	// convert into struct
	var p Person
	if err := json.Unmarshal(bytes, &p); err != nil {
		fmt.Println(err)
		return
	}

	// print struct
	fmt.Println(p)
}