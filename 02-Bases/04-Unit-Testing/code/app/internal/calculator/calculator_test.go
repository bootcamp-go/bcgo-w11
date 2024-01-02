package calculator_test

import (
	"app/internal/calculator"
	"testing"

	"github.com/stretchr/testify/require"
)

/*
	Command to run test:
	go test -run ^TestAddition$ -v ./internal/calculator/

	Note:
	- -run flag is used to specify the test function to run
	- -v flag is used to print the test result in verbose mode
	- ./internal/calculator/ is the path to the package to test
	- ^...$ is a regular expression to match the test function name
*/

// TestAddition tests the Addition function
// - subtesting is used to group test cases
func TestAddition(t *testing.T) {
	t.Run("success - case 01: positive number - 1 + 2 = 3", func(t *testing.T) {
		// arrange
		// ...
	
		// act
		a := 1.0
		b := 2.0
		result := calculator.Addition(a, b)
	
		// assert
		expectedResult := 3.0
		if result != expectedResult {
			t.Errorf("Addition(%f, %f) failed, expected %f, got %f", a, b, expectedResult, result)
			return
		}
	})

	t.Run("success - case 02: negative number - (-1)+(-2) = (-3)", func(t *testing.T) {
		// arrange
		// ...
	
		// act
		a := -1.0
		b := -2.0
		result := calculator.Addition(a, b)
	
		// assert
		expectedResult := -3.0
		if result != expectedResult {
			t.Errorf("Addition(%f, %f) failed, expected %f, got %f", a, b, expectedResult, result)
			return
		}
	})
}

// func TestAdditionOneAndTree(t *testing.T) {
// 	// arrange
// 	// ...

// 	// act
// 	a := 1.0
// 	b := 3.0
// 	result := calculator.Addition(a, b)

// 	// assert
// 	expectedResult := 4.0
// 	if result != expectedResult {
// 		t.Errorf("Addition(%f, %f) failed, expected %f, got %f", a, b, expectedResult, result)
// 	}
// }

func TestDivision(t *testing.T) {
	t.Run("success - case 01: floating point number - 1 / 2 = 0.5", func(t *testing.T) {
		// arrange
		// ...


		// act
		n := 1.0
		d := 2.0
		result, err := calculator.Division(n, d)

		// assert
		expectedError := ""
		expectedResult := 0.5
		require.Equal(t, expectedError, err)
		// if err != "" {
		// 	t.Errorf("Division(%f, %f) failed, expected no error, got %s", n, d, err)
		// 	return
		// }
		require.Equal(t, expectedResult, result)
		// if result != expectedResult {
		// 	t.Errorf("Division(%f, %f) failed, expected %f, got %f", n, d, expectedResult, result)
		// 	return
		// }
	})

	t.Run("success - case 02: integer number - 4 / 2 = 2", func(t *testing.T) {
		// arrange

		// act

		// assert
	})

	t.Run("success - case 03: negative number - (-1) / (-2) = 0.5", func(t *testing.T) {
		// arrange

		// act

		// assert
	})

	t.Run("success - case 04: zero divided by n>0 - 0 / 2 = 0", func(t *testing.T) {
		// arrange

		// act

		// assert

	})

	t.Run("failure - case 01: divide by zero - 1 / 0", func(t *testing.T) {
		// arrange
		// ...

		// act
		n := 1.0
		d := 0.0
		result, err := calculator.Division(n, d)

		// assert
		expectedErr := "denominator cannot be zero"
		expectedResult := 0.0
		require.Equal(t, expectedErr, err)
		require.Equal(t, expectedResult, result)
	})

	t.Run("failure - case 02: zero divided by zero - 0 / 0", func(t *testing.T) {
		// arrange

		// act
		n := 0.0
		d := 0.0
		result, err := calculator.Division(n, d)

		// assert
		expectedErr := "zero divided by zero is not possible"
		expectedResult := 0.0
		require.Equal(t, expectedErr, err)
		require.Equal(t, expectedResult, result)
	})
}