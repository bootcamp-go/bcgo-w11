package service_test

import (
	"go-testing/doubles/internal/city"
	"go-testing/doubles/internal/city/repository"
	"go-testing/doubles/internal/city/service"
	"go-testing/doubles/internal/logger"
	"go-testing/doubles/internal/weather"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestDefaultAddCity(t *testing.T) {
	t.Run("success - add city", func(t *testing.T) {
		// arrange
		// - repository writer
		rp := repository.NewStub()
		rp.FuncSaveCity = func(c *city.City) (err error) {
			c.ID = 1
			return
		}
		// - weather API
		wa := weather.NewWeatherStub()
		wa.FuncGetTemperature = func (city string) (degrees float64, err error) {
			degrees = 20.0
			return
		}
		// - logger
		lg := logger.NewDummy()
		// - service
		sv := service.NewDefault(rp, wa, lg)

		// act
		c, err := sv.AddCity("Paris", "France", 2_200_000, time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC))

		// assert
		expectedCity := city.City{
			ID:         1,
			Name:       "Paris",
			Country:    "France",
			Population: 2_200_000,
			Date:       time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			Temperature: 20.0,
		}
		require.Equal(t, expectedCity, c)
		require.NoError(t, err)
		// require.ErrorIs(t, expectedErr, err)
		// if err != nil {
		// 	require.EqualError(t, err, "error adding city: city invalid")
		// 	require.EqualError(t, err, expectedErr.Error())
		// }
	})

	t.Run("success 02 - add other city", func(t *testing.T) {
		// arrange
		// - repository writer
		rp := repository.NewStub()
		rp.FuncSaveCity = func(c *city.City) (err error) {
			c.ID = 1
			return
		}
		// - weather API
		wa := weather.NewWeatherMock()
		wa.FuncGetTemperature = func (city string) (degrees float64, err error) {
			degrees = 20.0
			return
		}
		// - logger
		lg := logger.NewDummy()
		// - service
		sv := service.NewDefault(rp, wa, lg)

		// act
		c, err := sv.AddCity("Paris", "France", 2_200_000, time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC))

		// assert
		expectedCity := city.City{
			ID:         1,
			Name:       "Paris",
			Country:    "France",
			Population: 2_200_000,
			Date:       time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			Temperature: 20.0,
		}
		require.Equal(t, expectedCity, c)
		require.NoError(t, err)
		require.Equal(t, 1, wa.Calls.GetTemperature)
		require.Equal(t, "Paris", wa.Calls.CurrentParamCity)
	})

	t.Run("error - invalid city", func(t *testing.T) {
		// arrange
		// - repository writer
		// ...
		// - weather API
		// ...
		// - logger
		lg := logger.NewDummy()
		// - service
		sv := service.NewDefault(nil, nil, lg)

		// act
		c, err := sv.AddCity("", "France", 2_200_000, time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC))

		// assert
		expectedCity := city.City{
			ID:         0,
			Name:       "",
			Country:    "France",
			Population: 2_200_000,
			Date:       time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		}
		expectedErr := service.ErrCityInvalid
		require.Equal(t, expectedCity, c)
		require.ErrorIs(t, expectedErr, err)
		require.EqualError(t, err, expectedErr.Error())
	})

	t.Run("error - failed to fetch temperature", func(t *testing.T) {
		// arrange

		// act

		// assert

	})

	t.Run("error - failed to save city", func(t *testing.T) {
		// arrange

		// act

		// assert

	})
}