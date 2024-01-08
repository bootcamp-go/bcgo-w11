package products

import "app/interfaces/s2/internal"

// NewMedium returns a new medium product
func NewMedium(product internal.Product, storageCost float64) Medium {
	// default values
	defaultStorageCost := 0.3
	if storageCost != 0 {
		defaultStorageCost = storageCost
	}

	return Medium{
		product:     product,
		storageCost: defaultStorageCost,
	}
}

// Medium is a representation of a medium product
type Medium struct {
	// product is other product
	product internal.Product
	// storageCost is the storage cost of a medium product (in percentage)
	storageCost float64
}

// Price returns the price of a medium product
func (m Medium) Price() float64 {
	return m.product.Price() * (1 + m.storageCost)
}