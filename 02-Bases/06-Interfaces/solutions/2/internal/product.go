package internal

// Product is a representation of a product
type Product interface {
	// Price returns the price of a product
	Price() float64
}