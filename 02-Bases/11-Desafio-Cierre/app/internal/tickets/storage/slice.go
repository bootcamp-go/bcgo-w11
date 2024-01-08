package storage

// NewSlice returns a new Slice instance
func NewSlice(tickets []Ticket) *Slice {
	return &Slice{tickets: tickets}
}

// Slice is an struct that represents a slice storage
type Slice struct {
	// tickets is the slice of tickets
	tickets []Ticket
}

// Read returns all the tickets
func (s *Slice) Read() (t []Ticket, err error) {
	// return the tickets
	t = s.tickets
	return
}
