package weather

import "errors"

var (
	// ErrCityNotFound is returned when the city is not found.
	ErrCityNotFound = errors.New("city not found")
)

// WeatherAPI is the interface that wraps the basic methods to get the temperature for a given city.
type WeatherAPI interface {
	// GetTemperature returns the temperature for the given city. The temperature is in degrees Celsius.
	GetTemperature(city string) (degrees float64, err error)
}