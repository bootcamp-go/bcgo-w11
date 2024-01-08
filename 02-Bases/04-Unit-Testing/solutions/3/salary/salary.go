package salary

type Category int

const (
	CategoryA Category = iota
	CategoryB
	CategoryC
)

func PerHourAndCategory(hours float64, category Category) (amount float64, err string) {
	// calculate amount
	switch category {
	case CategoryA:
		amount = 3000 * hours
		amount *= 1.5
	case CategoryB:
		amount = 1500 * hours
		amount *= 1.2
	case CategoryC:
		amount = 1000 * hours
	default:
		err = "Invalid category"
	}
	return
}