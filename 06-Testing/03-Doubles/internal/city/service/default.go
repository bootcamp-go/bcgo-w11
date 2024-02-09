package service

import (
	"errors"
	"go-testing/doubles/internal/city"
	"go-testing/doubles/internal/logger"
	"go-testing/doubles/internal/weather"
	"time"
)

var (
	// ErrCityInvalid is returned when the city is invalid.
	ErrCityInvalid = errors.New("city invalid")
)

// NewDefault returns a new instance of the Default service.
func NewDefault(cityRpWriter city.RepositoryWriter, weatherAPI weather.WeatherAPI, lg logger.Logger) *Default {
	return &Default{
		cityRpWriter: cityRpWriter,
		weatherAPI:   weatherAPI,
		lg:           lg,
	}
}

// Default is the default service for the cities.
type Default struct {
	// cityRpWriter is the repository writer for the cities.
	cityRpWriter city.RepositoryWriter
	// weatherAPI is the weather API.
	weatherAPI weather.WeatherAPI
	// logger is the logger for the service.
	lg logger.Logger
}

// AddCity adds a new city to the repository.
func (s *Default) AddCity(name, country string, population int, date time.Time) (c city.City, err error) {
	defer func () {
		// log if there is an error
		if err != nil {
			s.lg.Logf("error adding city: %v", err)
		}
	}()

	// create city
	c = city.City{
		Name:       name,
		Country:    country,
		Population: population,
		Date:       date,
	}
	// - stateless validation
	if c.Name == "" || c.Country == "" || c.Population < 0 {
		err = ErrCityInvalid
		return
	}

	// - fetch temperature
	c.Temperature, err = s.weatherAPI.GetTemperature(c.Name)
	if err != nil {
		return
	}

	// save city
	err = s.cityRpWriter.SaveCity(&c)
	if err != nil {
		return
	}

	return
}