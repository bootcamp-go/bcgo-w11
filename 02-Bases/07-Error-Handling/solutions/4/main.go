package main

import (
	"errors"
	"fmt"
)

var (
	// ErrorMinimumTaxable is returned when the salary entered does not reach the taxable minimum
	ErrorMinimumTaxable = errors.New("Error: the salary entered does not reach the taxable minimum")
)

// CalculateTax returns the tax of a salary
func CalculateTax(salary int) (tax int, err error) {
	// check if the salary is taxable
	if salary <= 150000 {
		err = fmt.Errorf("Error: the salary entered does not reach the taxable minimum: %d, salary: %d", 150000, salary)
		return
	}

	tax = salary * 20 / 100
	return
}

func main() {
	tax, err := CalculateTax(100000)
	if err != nil {
		if errors.Is(err, ErrorMinimumTaxable) {
			fmt.Println("code 01: ", err)
		} else {
			fmt.Println("code 02: ", err)
		}
		return
	}

	fmt.Println(tax)
}