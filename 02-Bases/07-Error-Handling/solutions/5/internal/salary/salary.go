package salary

import (
	"errors"
	"fmt"
)

var (
	ErrHoursNotAcceptable = errors.New("the hours entered are not acceptable")
)

// NewCalculator returns a new calculator
func NewCalculator(tax float64, taxable float64, minimumHours float64) Calculator {
	// default values
	defaultTax := 0.1
	defaultTaxable := 150000.0
	defaultMinimumHours := 80.0
	if tax != 0 {
		defaultTax = tax
	}
	if taxable != 0 {
		defaultTaxable = taxable
	}
	if minimumHours != 0 {
		defaultMinimumHours = minimumHours
	}

	return Calculator{
		Tax:          defaultTax,
		Taxable:      defaultTaxable,
		MinimumHours: defaultMinimumHours,
	}
}

// Calculator is a struct that calculates the salary
type Calculator struct {
	// Tax is the tax percentage
	Tax float64
	// Taxable is the minimum salary that is taxable
	Taxable float64
	// MinimumHours is the minimum hours that are acceptable
	MinimumHours float64
}

// Calculate returns the salary of a person
func (c *Calculator) Calculate(hours float64, rate float64) (salary float64, err error) {
	// check if the hours are acceptable
	if hours < c.MinimumHours {
		err = fmt.Errorf("%w: %f, hours: %f", ErrHoursNotAcceptable, c.MinimumHours, hours)
		return
	}
	// check if salary should have tax deducted
	salary = hours * rate
	if salary < c.Taxable {
		return
	}

	salary *= 1.0 - c.Tax
	return
}