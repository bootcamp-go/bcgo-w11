package pets_test

import (
	"app/unit-testing/fifthy/pets"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestDogGramsIntake tests the DogGramsIntake function
func TestDogGramsIntake(t *testing.T) {
	t.Run("success - case 01: 1 dog", func(t *testing.T) {
		// arrange
		// ...

		// act
		result := pets.DogGramsIntake(1)

		// assert
		require.Equal(t, 10000.0, result)			
	})

	t.Run("success - case 02: 3 dogs", func(t *testing.T) {
		// arrange
		// ...

		// act
		result := pets.DogGramsIntake(3)

		// assert
		require.Equal(t, 30000.0, result)
	})
}

// TestCatGramsIntake tests the CatGramsIntake function
func TestCatGramsIntake(t *testing.T) {
	t.Run("success - case 01: 1 cat", func(t *testing.T) {
		// arrange
		// ...

		// act
		result := pets.CatGramsIntake(1)

		// assert
		require.Equal(t, 5000.0, result)
	})

	t.Run("success - case 02: 3 cats", func(t *testing.T) {
		// arrange
		// ...

		// act
		result := pets.CatGramsIntake(3)

		// assert
		require.Equal(t, 15000.0, result)
	})
}

// TestHamsterGramsIntake tests the HamsterGramsIntake function
func TestHamsterGramsIntake(t *testing.T) {
	t.Run("success - case 01: 1 hamster", func(t *testing.T) {
		// arrange
		// ...

		// act
		result := pets.HamsterGramsIntake(1)

		// assert
		require.Equal(t, 250.0, result)
	})

	t.Run("success - case 02: 3 hamsters", func(t *testing.T) {
		// arrange
		// ...

		// act
		result := pets.HamsterGramsIntake(3)

		// assert
		require.Equal(t, 750.0, result)
	})
}

// TestTarantulaGramsIntake tests the TarantulaGramsIntake function
func TestTarantulaGramsIntake(t *testing.T) {
	t.Run("success - case 01: 1 tarantula", func(t *testing.T) {
		// arrange
		// ...

		// act
		result := pets.TarantulaGramsIntake(1)

		// assert
		require.Equal(t, 150.0, result)
	})

	t.Run("success - case 02: 3 tarantulas", func(t *testing.T) {
		// arrange
		// ...

		// act
		result := pets.TarantulaGramsIntake(3)

		// assert
		require.Equal(t, 450.0, result)
	})
}

// TestOrchestrator tests the Orchestrator function
func TestOrchestrator(t *testing.T) {
	t.Run("success - case 01: dog", func(t *testing.T) {
		// arrange
		// ...

		// act
		result, err := pets.Orchestrator(pets.Dog)

		// assert
		require.Equal(t, "", err)
		require.Equal(t, reflect.ValueOf(pets.DogGramsIntake).Pointer(), reflect.ValueOf(result).Pointer())
	})

	t.Run("success - case 02: cat", func(t *testing.T) {
		// arrange
		// ...

		// act
		result, err := pets.Orchestrator(pets.Cat)

		// assert
		require.Equal(t, "", err)
		require.Equal(t, reflect.ValueOf(pets.CatGramsIntake).Pointer(), reflect.ValueOf(result).Pointer())
	})

	t.Run("success - case 03: hamster", func(t *testing.T) {
		// arrange
		// ...

		// act
		result, err := pets.Orchestrator(pets.Hamster)

		// assert
		require.Equal(t, "", err)
		require.Equal(t, reflect.ValueOf(pets.HamsterGramsIntake).Pointer(), reflect.ValueOf(result).Pointer())
	})

	t.Run("success - case 04: tarantula", func(t *testing.T) {
		// arrange
		// ...

		// act
		result, err := pets.Orchestrator(pets.Tarantula)

		// assert
		require.Equal(t, "", err)
		require.Equal(t, reflect.ValueOf(pets.TarantulaGramsIntake).Pointer(), reflect.ValueOf(result).Pointer())
	})

	t.Run("error - case 01: unknown pet", func(t *testing.T) {
		// arrange
		// ...

		// act
		result, err := pets.Orchestrator(999)

		// assert
		require.Equal(t, "Unknown pet", err)
		require.Nil(t, result)
	})
}
