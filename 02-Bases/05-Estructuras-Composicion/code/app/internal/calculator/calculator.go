package calculator

// Operation represents an operation for potential registration
// in a calculator
type Operation struct {
	// Name is the name of the operation
	Name   string
	// Input is the input of the operation
	Input  [2]float64
	// Output is the output of the operation
	Output float64
}

// Calculator represents a calculator
type Calculator struct {
	// Count is the count of operations
	Count int
	// Limit is the limit of operations
	Limit int
	// Operations is a list of operations
	Operations []Operation
}

// Add returns the sum of two float64 numbers
func (c *Calculator) Add(a, b float64) (r float64, err string) {
	// check if the limit is reached
	if c.Count >= c.Limit {
		err = "limit reached"
		return
	}

	// calculate the sum
	r = a + b
	
	// register the operation
	op := Operation{
		Name:   "Add",
		Input:  [2]float64{a, b},
		Output: r,
	}
	// append the operation to the list
	c.Operations = append(c.Operations, op)
	return
}

// Sub returns the difference of two float64 numbers
func (c *Calculator) Sub(a, b float64) (r float64, err string) {
	// check if the limit is reached
	if c.Count >= c.Limit {
		err = "limit reached"
		return
	}

	// calculate the difference
	r = a - b

	// register the operation
	op := Operation{
		Name:   "Add",
		Input:  [2]float64{a, b},
		Output: r,
	}
	// append the operation to the list
	c.Operations = append(c.Operations, op)
	return
}

// Mul returns the product of two float64 numbers
func (c *Calculator) Mul(a, b float64) (r float64, err string) {
	// check if the limit is reached
	if c.Count >= c.Limit {
		err = "limit reached"
		return
	}

	// calculate the product
	r = a * b

	// register the operation
	op := Operation{
		Name:   "Add",
		Input:  [2]float64{a, b},
		Output: r,
	}
	// append the operation to the list
	c.Operations = append(c.Operations, op)
	
	return
}

// Div returns the quotient of two float64 numbers
func (c *Calculator) Div(a, b float64) (r float64, err string) {
	// check if the limit is reached
	if c.Count >= c.Limit {
		err = "limit reached"
		return
	}

	// calculate the quotient
	r = a / b

	// register the operation
	op := Operation{
		Name:   "Add",
		Input:  [2]float64{a, b},
		Output: r,
	}
	// append the operation to the list
	c.Operations = append(c.Operations, op)
	return
}