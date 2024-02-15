package calculator_test

import (
	"go-testing/coverage-linter/calculator"
	"testing"
)

func TestIsPrime(t *testing.T) {
	type testCase struct {
		name          string
		inputNumber   int
		expectedPrime bool
	}
	testCases := []testCase{
		// not prime numbers
		{name: "1 is not prime", inputNumber: 1, expectedPrime: false},
		{name: "4 is not prime", inputNumber: 4, expectedPrime: false},
		// prime numbers
		{name: "2 is prime", inputNumber: 2, expectedPrime: true},
		{name: "3 is prime", inputNumber: 3, expectedPrime: true},
	}

	// run the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// arrange
			// ...

			// act
			result := calculator.IsPrime(tc.inputNumber)

			// assert
			if result != tc.expectedPrime {
				t.Errorf("input: %d, expected: %v, got: %v", tc.inputNumber, tc.expectedPrime, result)
				return
			}
		})
	}

}