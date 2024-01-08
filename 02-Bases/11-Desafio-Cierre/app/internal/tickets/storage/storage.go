package storage

import "time"

// Ticket is an struct that represents a ticket
type Ticket struct {
	// ID is the ticket identifier
	ID int
	// Passenger is the passenger name
	Passenger string
	// Email is the passenger email
	Email string
	// Destination is the destination string
	Destination string
	// Departure is the departure hour between 0 and 24 hours
	Departure time.Time
	// Price is the ticket price
	Price float64
}

// Reader is an interface that represents a ticket reader
type Reader interface {
	// Read returns all the tickets
	Read() (t []Ticket, err error)
}
