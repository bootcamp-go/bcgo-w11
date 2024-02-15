package calculator_test

import (
	"go-testing/coverage-linter/calculator"
	"testing"
)

func TestDivision(t *testing.T) {
	t.Run("failure - both numbers are zero", func(t *testing.T) {
		// arrange
		// ...

		// act
		numerator := 0
		denominator := 0
		result, err := calculator.Division(numerator, denominator)

		// assert
		if err != calculator.ErrDivisionIndeterminateForm {
			t.Errorf("expected: %v, got: %v", calculator.ErrDivisionIndeterminateForm, err)
			return
		}
		if result != 0 {
			t.Errorf("expected: 0, got: %d", result)
			return
		}
	})
}