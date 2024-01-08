package operations_test

import (
	"app/unit-testing/fourth/operations"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestMinimum tests the Minimum function
func TestMinimum(t *testing.T) {
	t.Run("success - case 01: only one value", func(t *testing.T) {
		// arrange
		// ...

		// act
		result, err := operations.Minimum(1)

		// assert
		require.Equal(t, "", err)
		require.Equal(t, 1, result)
	})

	t.Run("success - case 02: multiple values", func(t *testing.T) {
		// arrange
		// ...

		// act
		result, err := operations.Minimum(1, 2, 3, 4, 5)

		// assert
		require.Equal(t, "", err)
		require.Equal(t, 1, result)
	})

	t.Run("error - case 01: empty values", func(t *testing.T) {
		// arrange
		// ...

		// act
		result, err := operations.Minimum()

		// assert
		require.Equal(t, "Empty values", err)
		require.Equal(t, 0, result)
	})
}

// TestAverage tests the Average function
func TestAverage(t *testing.T) {
	t.Run("success - case 01: only one value", func(t *testing.T) {
		// arrange
		// ...

		// act
		result, err := operations.Average(1)

		// assert
		require.Equal(t, "", err)
		require.Equal(t, 1, result)
	})

	t.Run("success - case 02: multiple values", func(t *testing.T) {
		// arrange
		// ...

		// act
		result, err := operations.Average(1, 2, 3, 4, 5)

		// assert
		require.Equal(t, "", err)
		require.Equal(t, 3, result)
	})

	t.Run("error - case 01: empty values", func(t *testing.T) {
		// arrange
		// ...

		// act
		result, err := operations.Average()

		// assert
		require.Equal(t, "Empty values", err)
		require.Equal(t, 0, result)
	})
}

// TestMaximum tests the Maximum function
func TestMaximum(t *testing.T) {
	t.Run("success - case 01: only one value", func(t *testing.T) {
		// arrange
		// ...

		// act
		result, err := operations.Maximum(1)

		// assert
		require.Equal(t, "", err)
		require.Equal(t, 1, result)
	})

	t.Run("success - case 02: multiple values", func(t *testing.T) {
		// arrange
		// ...

		// act
		result, err := operations.Maximum(1, 2, 3, 4, 5)

		// assert
		require.Equal(t, "", err)
		require.Equal(t, 5, result)
	})

	t.Run("error - case 01: empty values", func(t *testing.T) {
		// arrange
		// ...

		// act
		result, err := operations.Maximum()

		// assert
		require.Equal(t, "Empty values", err)
		require.Equal(t, 0, result)
	})
}

// TestOrchestrator tests the Orchestrator function
func TestOrchestrator(t *testing.T) {
	t.Run("success - case 01: minimum", func(t *testing.T) {
		// arrange
		// ...

		// act
		result, err := operations.Orchestrator(operations.MinimumOperator)

		// assert
		require.Equal(t, "", err)
		require.Equal(t, reflect.ValueOf(operations.Minimum).Pointer(), reflect.ValueOf(result).Pointer())
	})

	t.Run("success - case 02: average", func(t *testing.T) {
		// arrange
		// ...

		// act
		result, err := operations.Orchestrator(operations.AverageOperator)

		// assert
		require.Equal(t, "", err)
		require.Equal(t, reflect.ValueOf(operations.Average).Pointer(), reflect.ValueOf(result).Pointer())
	})

	t.Run("success - case 03: maximum", func(t *testing.T) {
		// arrange
		// ...

		// act
		result, err := operations.Orchestrator(operations.MaximumOperator)

		// assert
		require.Equal(t, "", err)
		require.Equal(t, reflect.ValueOf(operations.Maximum).Pointer(), reflect.ValueOf(result).Pointer())
	})

	t.Run("error - case 01: invalid operator", func(t *testing.T) {
		// arrange
		// ...

		// act
		result, err := operations.Orchestrator(3)

		// assert
		require.Equal(t, "Invalid operator", err)
		require.Nil(t, result)
	})
}