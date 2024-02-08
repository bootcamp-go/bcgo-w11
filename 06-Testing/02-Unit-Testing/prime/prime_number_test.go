package prime_test

import (
	"testing"
	"unit-testing-primes/prime"

	"github.com/stretchr/testify/require"
)

func TestIsPrime(t *testing.T) {
	t.Run("is prime number - 7", func(t *testing.T) {
		// arrange / given
		n := 7
		expectedResult := true

		// act / when
		result := prime.IsPrime(n)

		// assert / then
		if result != expectedResult {
			t.Errorf("IsPrime(%d) = %v; want %v", n, result, expectedResult)
			return
		}
	})

	t.Run("is not prime number - 4", func(t *testing.T) {
		// arrange / given
		n := 4
		expectedResult := false

		// act / when
		result := prime.IsPrime(n)

		// assert / then
		// if result != expectedResult {
		// 	t.Errorf("IsPrime(%d) = %v; want %v", n, result, expectedResult)
		// 	return
		// }
		require.Equal(t, expectedResult, result)
	})
	
	t.Run("is not prime number - 1", func(t *testing.T) {
		// arrange / given
		n := 1
		expectedResult := false

		// act / when
		result := prime.IsPrime(n)

		// assert / then
		// if result != expectedResult {
		// 	t.Errorf("IsPrime(%d) = %v; want %v", n, result, expectedResult)
		// 	return
		// }
		require.Equal(t, expectedResult, result)
	})
}

func TestIsPrime_TableDrivenTest(t *testing.T) {
	type testCase struct {
		name     string
		input    int
		expected bool
	}
	cases := []testCase{
		{name: "is not prime number - -1", input: -1, expected: false},
		{name: "is not prime number - 0", input: 0, expected: false},
		{name: "is not prime number - 1", input: 1, expected: false},
		{name: "is prime number - 2", input: 2, expected: true},
		{name: "is prime number - 3", input: 3, expected: true},
		{name: "is not prime number - 4", input: 4, expected: false},
		{name: "is prime number - 5", input: 5, expected: true},
	}

	// run test cases
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			// arrange / given
			// ...

			// act / when
			result := prime.IsPrime(tc.input)

			// assert / then
			require.Equal(t, tc.expected, result)
		})
	}
}
