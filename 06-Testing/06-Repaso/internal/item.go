package internal

// Item is a struct that represents a single item in the inventory
type Item struct {
	// ID is the unique identifier for the item
	ID          int
	// Name is the name of the item
	Name        string
	// Description is a brief description of the item
	Description string
	// Price is the cost of the item
	Price       float64
}