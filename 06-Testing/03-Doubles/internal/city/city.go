package city

import "time"

// City is an struct that represents a city.
type City struct {
	// ID is the unique identifier for the city.
	ID int
	// Name is the name of the city.
	Name string
	// Country is the country where the city is located.
	Country string
	// Population is the number of people living in the city.
	Population int
	// Temperature is the temperature in the city. The temperature is in degrees Celsius.
	Temperature float64
	// Date is the date when the temperature was measured.
	Date time.Time
}