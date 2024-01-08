package operations

type Operator int

const (
	MinimumOperator Operator = iota
	AverageOperator
	MaximumOperator
)

// Operation is a function that receives a list of integers and returns an integer and a string
type Operation func(...int) (int, string)

// Minimum returns the minimum value of a list of integers
func Minimum(values ...int) (result int, err string) {
	// check if values is empty
	if len(values) == 0 {
		err = "Empty values"
		return
	}

	// calculate minimum
	var min int
	for ix, value := range values {
		if ix == 0 {
			min = value
			continue
		}
		if value < min {
			min = value
		}
	}
	result = min
	return
}

// Average returns the average value of a list of integers
func Average(values ...int) (result int, err string) {
	// check if values is empty
	if len(values) == 0 {
		err = "Empty values"
		return
	}

	// calculate average
	var sum int
	for _, value := range values {
		sum += value
	}
	result = sum / len(values)
	return
}

// Maximum returns the maximum value of a list of integers
func Maximum(values ...int) (result int, err string) {
	// check if values is empty
	if len(values) == 0 {
		err = "Empty values"
		return
	}

	// calculate maximum
	var max int
	for ix, value := range values {
		if ix == 0 {
			max = value
			continue
		}
		if value > max {
			max = value
		}
	}
	result = max
	return
}

// Orchestrator returns the operation function and an error message
func Orchestrator(operator Operator) (op Operation, err string) {
	switch operator {
	case MinimumOperator:
		op = Minimum
	case AverageOperator:
		op = Average
	case MaximumOperator:
		op = Maximum
	default:
		err = "Invalid operator"
	}
	return
}