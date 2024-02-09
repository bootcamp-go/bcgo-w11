package weather

// NewWeatherMock returns a new instance of the WeatherMock.
func NewWeatherMock() *WeatherMock {
	return &WeatherMock{
		FuncGetTemperature: func(city string) (degrees float64, err error) {
			return
		},
	}
}

// WeatherMock is a mock for the WeatherAPI interface.
type WeatherMock struct {
	// FuncGetTemperature is the function to externally set the GetTemperature method.
	FuncGetTemperature func(city string) (degrees float64, err error)
	// Spy
	Calls struct {
		// GetTemperature is the number of calls to the GetTemperature method.
		GetTemperature    int
		CurrentParamCity  string
	}
}

// GetTemperature returns the temperature for the given city. The temperature is in degrees Celsius.
func (m *WeatherMock) GetTemperature(city string) (degrees float64, err error) {
	// increment the number of calls
	m.Calls.GetTemperature++

	// register the current parameter
	m.Calls.CurrentParamCity = city

	// call the function
	degrees, err = m.FuncGetTemperature(city)
	return
}