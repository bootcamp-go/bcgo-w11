package calculator

// Addition adds two float64 numbers and returns the result
func Addition(a, b float64) (result float64) {
	result = a + b
	return
}

// Subtraction subtracts two float64 numbers and returns the result
func Subtraction(a, b float64) (result float64) {
	result = a - b
	return
}

// Multiplication multiplies two float64 numbers and returns the result
func Multiplication(a, b float64) (result float64) {
	result = a * b
	return
}

// Division divides two float64 numbers and returns the result
func Division(a, b float64) (result float64, err string) {
	// check if the numerator and denominator is zero
	if a == 0 && b == 0 {
		err = "zero divided by zero is not possible"
		return
	}

	// check if the denominator is zero
	if b == 0 {
		err = "denominator cannot be zero"
		return
	}

	// divide the numbers
	result = a / b
	return
}