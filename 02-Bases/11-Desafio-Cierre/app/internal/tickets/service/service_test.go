package service_test

import (
	"app/desafio-cierre/internal/tickets/service"
	"app/desafio-cierre/internal/tickets/storage"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestDefault_GetTotal(t *testing.T) {
	t.Run("success 01 - should return some tickets", func(t *testing.T) {
		// arrange
		// - storage: slice
		st := storage.NewSlice([]storage.Ticket{
			{ID: 1, Passenger: "John", Destination: "London", Departure: time.Date(0, 0, 0, 6, 0, 0, 0, time.UTC), Price: 100},
			{ID: 2, Passenger: "Jane", Destination: "Paris", Departure: time.Date(0, 0, 0, 12, 0, 0, 0, time.UTC), Price: 200},
		})
		// - service: default
		sv := service.NewDefault(st)

		// act
		total, err := sv.GetTotal()

		// assert
		require.NoError(t, err)
		require.Equal(t, 2, total)
	})

	t.Run("success 02 - should return no tickets", func(t *testing.T) {
		// arrange
		// - storage: slice
		st := storage.NewSlice([]storage.Ticket{})
		// - service: default
		sv := service.NewDefault(st)

		// act
		total, err := sv.GetTotal()

		// assert
		require.NoError(t, err)
		require.Equal(t, 0, total)
	})
}

func TestDefault_GetCountByPeriod(t *testing.T) {
	t.Run("success 01 - should return 1", func(t *testing.T) {
		// arrange
		// - storage: slice
		st := storage.NewSlice([]storage.Ticket{
			{ID: 1, Passenger: "John", Destination: "London", Departure: time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC), Price: 100},
			{ID: 2, Passenger: "Jane", Destination: "Paris", Departure: time.Date(0, 0, 0, 18, 0, 0, 0, time.UTC), Price: 200},
		})
		// - service: default
		sv := service.NewDefault(st)

		// act
		count, err := sv.GetCountByPeriod(service.EarlyMorning)

		// assert
		require.NoError(t, err)
		require.Equal(t, 1, count)
	})

	t.Run("success 02 - should return 0", func(t *testing.T) {
		// arrange
		// - storage: slice
		st := storage.NewSlice([]storage.Ticket{
			{ID: 1, Passenger: "John", Destination: "London", Departure: time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC), Price: 100},
			{ID: 2, Passenger: "Jane", Destination: "Paris", Departure: time.Date(0, 0, 0, 18, 0, 0, 0, time.UTC), Price: 200},
		})
		// - service: default
		sv := service.NewDefault(st)

		// act
		count, err := sv.GetCountByPeriod(service.Night)

		// assert
		require.NoError(t, err)
		require.Equal(t, 0, count)
	})

	t.Run("failure 01 - should return an error for an invalid period", func(t *testing.T) {
		// arrange
		// - storage: slice
		st := storage.NewSlice([]storage.Ticket{})
		// - service: default
		sv := service.NewDefault(st)

		// act
		count, err := sv.GetCountByPeriod(service.Period{})

		// assert
		require.Error(t, err)
		require.ErrorIs(t, err, service.ErrInvalidPeriod)
		require.EqualError(t, err, service.ErrInvalidPeriod.Error())
		require.Equal(t, 0, count)
	})
}

func TestDefault_GetPercentageByDestination(t *testing.T) {
	t.Run("success 01 - should return some percentage", func(t *testing.T) {
		// arrange
		// - storage: slice
		st := storage.NewSlice([]storage.Ticket{
			{ID: 1, Passenger: "John", Destination: "London", Departure: time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC), Price: 100},
			{ID: 2, Passenger: "Jane", Destination: "Paris", Departure: time.Date(0, 0, 0, 18, 0, 0, 0, time.UTC), Price: 200},
		})
		// - service: default
		sv := service.NewDefault(st)

		// act
		percentage, err := sv.GetPercentageByDestination("London")

		// assert
		expectedPercentage := 0.5
		require.NoError(t, err)
		require.Equal(t, expectedPercentage, percentage)
	})

	t.Run("success 02 - should return percentage 0 - no destination tickets", func(t *testing.T) {
		// arrange
		// - storage: slice
		st := storage.NewSlice([]storage.Ticket{
			{ID: 1, Passenger: "John", Destination: "London", Departure: time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC), Price: 100},
			{ID: 2, Passenger: "Jane", Destination: "Paris", Departure: time.Date(0, 0, 0, 18, 0, 0, 0, time.UTC), Price: 200},
		})
		// - service: default
		sv := service.NewDefault(st)

		// act
		percentage, err := sv.GetPercentageByDestination("Madrid")

		// assert
		expectedPercentage := 0.0
		require.NoError(t, err)
		require.Equal(t, expectedPercentage, percentage)
	})

	t.Run("success 03 - should return percentage 0 - no tickets", func(t *testing.T) {
		// arrange
		// - storage: slice
		st := storage.NewSlice([]storage.Ticket{})
		// - service: default
		sv := service.NewDefault(st)
		
		// act
		percentage, err := sv.GetPercentageByDestination("London")

		// assert
		expectedPercentage := 0.0
		require.NoError(t, err)
		require.Equal(t, expectedPercentage, percentage)
	})
}