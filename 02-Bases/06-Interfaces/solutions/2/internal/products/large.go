package products

import "app/interfaces/s2/internal"

// NewLarge returns a new large product
func NewLarge(product internal.Product, storageCost float64, shippingCost float64) Large {
	// default values
	defaultStorageCost := 0.6
	defaultShippingCost := 2500.0
	if storageCost != 0 {
		defaultStorageCost = storageCost
	}
	if shippingCost != 0 {
		defaultShippingCost = shippingCost
	}

	return Large{
		product:      product,
		storageCost:  defaultStorageCost,
		shippingCost: defaultShippingCost,
	}
}

// Large is a representation of a large product
type Large struct {
	// product is other product
	product internal.Product
	// storageCost is the storage cost of a large product (in percentage)
	storageCost float64
	// shippingCost is the shipping cost of a large product (amount)
	shippingCost float64
}

// Price returns the price of a large product
func (l Large) Price() float64 {
	return l.product.Price() * (1 + l.storageCost) + l.shippingCost
}