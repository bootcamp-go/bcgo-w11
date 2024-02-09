package repository

import (
	"go-testing/doubles/internal/city"

	"github.com/stretchr/testify/mock"
)

/*
	Declare expected calls
	mk.On("SaveCity", city).Return(nil)
	|
	|
	struct {
		methodName string
		expectedInput []any
		output []any
	}
*/
// NewMock returns a new mock for the city repository writer.
func NewMock() *Mock {
	return &Mock{
		FuncSaveCity: func(c *city.City) {},
	}
}

// Mock is the mock for the city repository writer.
type Mock struct {
	// mock - testify
	mock.Mock
	// FuncSaveCity is the function to save a city.
	FuncSaveCity func(c *city.City)
}

func (m *Mock) SaveCity(c *city.City) (err error) {
	// ask output arguments
	output := m.Called(c)

	// process
	if m.FuncSaveCity != nil {
		m.FuncSaveCity(c)
	}

	// output
	err = output.Error(0)
	return
	// err = output[0].(error)
	// return
}