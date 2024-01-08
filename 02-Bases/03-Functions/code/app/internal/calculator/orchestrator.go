package calculator

func Orchestrator(operator MathOperator) (mo MathOperation, err string) {
	switch operator {
	case AdditionSymbol:
		mo = Addition
	case SubtractionSymbol:
		mo = Subtraction
	case MultiplicationSymbol:
		mo = Multiplication
	case DivisionSymbol:
		mo = Division
	default:
		err = "invalid operator"
	}
	return
}