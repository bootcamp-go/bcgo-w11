package salary_test

import (
	"app/unit-testing/third/salary"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestPerHourAndCategory tests the PerHourAndCategory function
func TestPerHourAndCategory(t *testing.T) {
	t.Run("success - case 01: category A - 100 hours", func(t *testing.T) {
		// arrange
		// ...

		// act
		result, err := salary.PerHourAndCategory(100.0, salary.CategoryA)

		// assert
		require.Equal(t, "", err)
		require.Equal(t, 450000.0, result)
	})

	t.Run("success - case 02: category B - 100 hours", func(t *testing.T) {
		// arrange
		// ...

		// act
		result, err := salary.PerHourAndCategory(100.0, salary.CategoryB)

		// assert
		require.Equal(t, "", err)
		require.Equal(t, 180000.0, result)
	})

	t.Run("success - case 03: category C - 100 hours", func(t *testing.T) {
		// arrange
		// ...

		// act
		result, err := salary.PerHourAndCategory(100.0, salary.CategoryC)

		// assert
		require.Equal(t, "", err)
		require.Equal(t, 100000.0, result)
	})

	t.Run("failure - case 01: invalid category", func(t *testing.T) {
		// arrange
		// ...

		// act
		result, err := salary.PerHourAndCategory(100.0, 4)

		// assert
		require.Equal(t, "Invalid category", err)
		require.Equal(t, 0.0, result)
	})
}
