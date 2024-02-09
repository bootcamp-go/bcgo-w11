package repository

import "go-testing/doubles/internal/city"

// NewStub returns a new instance of the Stub.
func NewStub() *Stub {
	// default config
	defaultFuncSaveCity := func(c *city.City) (err error) {
		return
	}

	return &Stub{
		FuncSaveCity: defaultFuncSaveCity,
	}
}

// Stub is a stub repository.
type Stub struct {
	// FuncSaveCity externalize the applied process on the method SaveCity
	FuncSaveCity func(c *city.City) (err error)
}

// SaveCity saves the given city in the repository.
func (s *Stub) SaveCity(c *city.City) (err error) {
	// call the function
	err = s.FuncSaveCity(c)
	return
}