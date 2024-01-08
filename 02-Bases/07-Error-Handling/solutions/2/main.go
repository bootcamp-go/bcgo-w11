package main

import (
	"fmt"
)

// ErrorMinimumTaxable is returned when the salary entered does not reach the taxable minimum
type ErrorMinimumTaxable struct {
	Minimum int
	Salary int
}

func (e ErrorMinimumTaxable) Error() string {
	return fmt.Sprintf("Error: the salary entered does not reach the taxable minimum: %d, salary: %d", e.Minimum, e.Salary)
}

// CalculateTax returns the tax of a salary
func CalculateTax(salary int) (tax int, err error) {
	// check if the salary is taxable
	if salary <= 10000 {
		err = ErrorMinimumTaxable{Minimum: 10000, Salary: salary}
		return
	}

	tax = salary * 20 / 100
	return
}

func main() {
	tax, err := CalculateTax(100000)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(tax)
}