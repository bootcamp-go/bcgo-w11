package storage

import (
	"encoding/csv"
	"os"
	"strconv"
	"time"
)

// CSV is an struct that represents a CSV storage
type CSV struct {
	file *os.File
}

// NewCSV returns a new CSV instance
func NewCSV(file *os.File) *CSV {
	return &CSV{file: file}
}

// Read returns all the tickets
func (c *CSV) Read() (t []Ticket, err error) {
	// create a new reader
	rd := csv.NewReader(c.file)

	// read the header
	_, err = rd.Read()
	if err != nil {
		return
	}

	// read all the records
	records, err := rd.ReadAll()
	if err != nil {
		return
	}

	// iterate over the records
	for _, v := range records {
		// create a new ticket
		// - id
		var id int
		id, err = strconv.Atoi(v[0])
		if err != nil {
			return
		}
		// - departure
		var departure time.Time
		departure, err = time.Parse("15:04", v[4])
		if err != nil {
			return
		}
		// - price
		var price float64
		price, err = strconv.ParseFloat(v[5], 64)
		if err != nil {
			return
		}

		// append the ticket
		t = append(t, Ticket{
			ID:          id,
			Passenger:   v[1],
			Email:       v[2],
			Destination: v[3],
			Departure:   departure,
			Price:       price,
		})
	}

	// return the tickets
	return
}