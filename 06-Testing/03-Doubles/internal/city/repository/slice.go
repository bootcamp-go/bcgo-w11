package repository

import "go-testing/doubles/internal/city"

// NewSlice returns a new instance of the Slice repository.
func NewSlice(cities []city.City, lastID int) *Slice {
	// default config
	defaultCities := make([]city.City, 0)
	defaultLastID := 0
	if cities != nil {
		defaultCities = cities
	}
	if lastID != 0 {
		defaultLastID = lastID
	}

	return &Slice{
		cities: defaultCities,
		lastID: defaultLastID,
	}
}

// Slice is a repository that stores the cities in memory.
type Slice struct {
	// cities is the in-memory storage for the cities.
	cities []city.City
	// lastID is the last used identifier for a city.
	lastID int
}

// SaveCity saves the given city in the repository.
func (r *Slice) SaveCity(c *city.City) (err error) {
	// stateful validation
	// - check if the city is duplicated
	for _, v := range r.cities {
		if v.Name == c.Name && v.Country == c.Country {
			err = city.ErrCityDuplicated
			return
		}
	}

	// set lastID
	r.lastID++
	c.ID = r.lastID

	// save city
	r.cities = append(r.cities, *c)

	return
}