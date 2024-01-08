package main

import (
	"fmt"
)

// ErrorMinimumTaxable is returned when the salary entered does not reach the taxable minimum
type ErrorMinimumTaxable struct {}

func (e ErrorMinimumTaxable) Error() string {
	return "Error: the salary entered does not reach the taxable minimum"
}

// CalculateTax returns the tax of a salary
func CalculateTax(salary int) (tax int, err error) {
	// check if the salary is taxable
	if salary < 150000 {
		err = ErrorMinimumTaxable{}
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