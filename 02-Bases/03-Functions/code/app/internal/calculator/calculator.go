package calculator

type MathOperator string

const (
	// AdditionSymbol is the symbol used for addition.
	AdditionSymbol MathOperator = "+"
	// SubtractionSymbol is the symbol used for subtraction.
	SubtractionSymbol = "-"
	// MultiplicationSymbol is the symbol used for multiplication.
	MultiplicationSymbol = "*"
	// DivisionSymbol is the symbol used for division.
	DivisionSymbol = "/"
)

type MathOperation func(float64, float64) (float64, string)

// decorator
// func ShowTime(mo MathOperation) MathOperation {
// 	return func(n1, n2 float64) (result float64, err string) {
// 		// before
// 		println("before")
// 		// call the original function
// 		result, err = mo(n1, n2)
// 		// after
// 		println("after")
// 		return
// 	}
// }

// Addition returns the sum of 2 float64 values.
func Addition(n1, n2 float64) (float64, string) {
	return n1 + n2, ""
}

// Subtraction returns the difference of 2 float64 values.
func Subtraction(n1, n2 float64) (result float64, err string) {
	result = n1 - n2
	return
}

// Multiplication returns the product of 2 float64 values.
func Multiplication(n1, n2 float64) (result float64, err string) {
	result = n1 * n2
	return
}

// Division returns the quotient of 2 float64 values.
func Division(n1, n2 float64) (result float64, err string) {
	// check if denominator is 0
	if n2 == 0 {
		err = "cannot divide by 0"
		return
	}

	// divide
	result = n1 / n2
	return result, ""
}