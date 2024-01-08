package products

import "app/interfaces/s2/internal"

type ProductSize int

const (
	SmallSize ProductSize = iota
	MediumSize
	LargeSize
)

// Factory is a function that returns a product
func Factory(size ProductSize, cost float64) (pr internal.Product, err string) {
	switch size {
	case SmallSize:
		pr = NewSmall(cost)
	case MediumSize:
		pr = NewMedium(NewSmall(cost), 0.0)
	case LargeSize:
		pr = NewLarge(NewMedium(NewSmall(cost), 0.0), 0.0, 0.0)
	default:
		err = "invalid size"
	}
	return
}