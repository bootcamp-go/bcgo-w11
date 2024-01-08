package average_test

import (
	"app/unit-testing/second/average"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestAverage tests the average function
func TestAverage(t *testing.T) {
	t.Run("success - case 01: average of 5 and 10 is 7 (rounded)", func(t *testing.T) {
		// arrange
		// ...

		// act
		result, err := average.Average(5, 10)

		// assert
		require.Equal(t, "", err)
		require.Equal(t, uint(7), result)
	})

	t.Run("success - case 02: average of 10, 20 and 30 is 20 (rounded)", func(t *testing.T) {
		// arrange
		// ...

		// act
		result, err := average.Average(10, 20, 30)

		// assert
		require.Equal(t, "", err)
		require.Equal(t, uint(20), result)
	})

	t.Run("failure - case 01: empty grades", func(t *testing.T) {
		// arrange
		// ...

		// act
		result, err := average.Average()

		// assert
		require.Equal(t, "No grades provided", err)
		require.Equal(t, uint(0), result)
	})
}