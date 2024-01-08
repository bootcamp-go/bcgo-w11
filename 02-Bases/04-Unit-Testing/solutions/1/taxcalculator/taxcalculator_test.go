package taxcalculator_test

import (
	"testing"

	"app/unit-testing/first/taxcalculator"

	"github.com/stretchr/testify/require"
)

// TestTaxCalculation tests the TaxCalculation function
func TestTaxCalculation(t *testing.T) {
	t.Run("success - case 01: salary greater than 150000", func(t *testing.T) {
		// arrange
		// ...

		// act
		salary := 200000.0
		tax := taxcalculator.TaxCalculation(salary)

		// assert
		require.Equal(t, 54000.0, tax)
	})

	t.Run("success - case 02: salary greater than 50000", func(t *testing.T) {
		// arrange
		// ...

		// act
		salary := 100000.0
		tax := taxcalculator.TaxCalculation(salary)

		// assert
		require.Equal(t, 17000.0, tax)
	})

	t.Run("success - case 03: salary less than 50000", func(t *testing.T) {
		// arrange
		// ...

		// act
		salary := 40000.0
		tax := taxcalculator.TaxCalculation(salary)

		// assert
		require.Equal(t, 0.0, tax)
	})
}