package taxcalculator

// TaxCalculation calculates the tax of a salary
func TaxCalculation(salary float64) (tax float64) {
	if salary > 150000 {
		tax = salary * 0.27
	} else if salary > 50000 {
		tax = salary * 0.17
	}
	return
}