package products

// NewSmall returns a new small product
func NewSmall(price float64) Small {
	return Small{
		price: price,
	}
}

// Small is a representation of a small product
type Small struct {
	// price is the price of a small product
	price float64
}

// Price returns the price of a small product
func (s Small) Price() float64 {
	return s.price
}
