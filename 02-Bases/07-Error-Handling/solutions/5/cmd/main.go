package main

import (
	"app/error-handling/s5/internal/salary"
	"errors"
	"fmt"
)

func main() {
	// create a calculator
	cl := salary.NewCalculator(0.1, 150000.0, 80.0)

	// work rate
	hours := 100.0
	rate := 10.0

	// calculate the salary
	payment, err := cl.Calculate(hours, rate)
	if err != nil {
		switch {
		case errors.Is(err, salary.ErrHoursNotAcceptable):
			fmt.Println("Error: the hours entered are not acceptable")
		default:
			fmt.Println("Error: unknown error")
		}
		return
	}

	fmt.Println(payment)
}