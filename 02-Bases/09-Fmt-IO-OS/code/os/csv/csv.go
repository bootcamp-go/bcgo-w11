package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Person struct {
	Id      int
	Name    string
	Genre   string
	Age     int
	Country string
}

func main() {
	// open file
	f, err := os.Open("./code/os/csv/data.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	// reader csv
	rd := csv.NewReader(f)

	// skip header
	rd.Read()

	// read csv
	for {
		record, err := rd.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			fmt.Println(err)
			return
		}

		// serialize person
		// - binding
		id, err := strconv.Atoi(record[0])
		if err != nil {
			fmt.Println(err)
			return
		}
		age, err := strconv.Atoi(record[3])
		if err != nil {
			fmt.Println(err)
			return
		}
		// - serialize
		person := Person{
			Id:      id,
			Name:    record[1],
			Genre:   record[2],
			Age:     age,
			Country: record[4],
		}

		fmt.Println(person)
	}
}
