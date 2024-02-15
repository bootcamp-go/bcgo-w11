package calculator

import "errors"

var (
	// DivisionErrorIndeterminateForm is returned when both numbers are zero
	// DivisionErrorIndeterminateForm = errors.New("Error: indeterminate form.")
	// ErrDivisionIndeterminateForm is returned when both numbers are zero
	ErrDivisionIndeterminateForm = errors.New("error: indeterminate form")
	// DivisionErrorZeroDenominator is returned when the denominator is zero
	// DivisionErrorZeroDenominator = errors.New("Error: division by zero.")
	// ErrDivisionZeroDenominator is returned when the denominator is zero
	ErrDivisionZeroDenominator = errors.New("error: division by zero")
)

// Division returns the division of two numbers
func Division(a, b int) (result int, err error) {
	// check if both numbers are zero
	if a == 0 && b == 0 {
		err = ErrDivisionIndeterminateForm
		return
	}

	// check if the denominator is zero
	if b == 0 {
		err = ErrDivisionZeroDenominator
		return
	}

	result = a / b
	return
}

func Multiply(a, b int) int {
	return a * b
}