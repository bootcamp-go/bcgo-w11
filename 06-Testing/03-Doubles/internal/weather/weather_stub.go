package weather

// NewWeatherStub returns a new instance of the WeatherStub.
func NewWeatherStub() *WeatherStub {
	// default config
	defaultFuncGetTemperature := func(city string) (degrees float64, err error) {
		return 0.0, nil
	}

	return &WeatherStub{
		FuncGetTemperature: defaultFuncGetTemperature,
	}
}

// weather := weather.NewWeatherStub()
// weather.Degrees = 20.0
// weather.Err = nil

// weather.FuncGetTemperature = func(city string) (degrees float64, err error) {
// 	return 20.0, nil
// }

// weather.FuncGetTemperature = func(city string) (degrees float64, err error) {
// 	return 0.0, errors.New("error getting temperature")
// }

// WeatherStub is a stub for the weather API
// for testing purposes.
type WeatherStub struct {
	// // Degrees is the temperature in degrees Celsius.
	// Degrees float64
	// // Err is the error to return.
	// Err error

	// FuncGetTemperature is the function to get the temperature.
	// - externalize the function to be able to test it
	FuncGetTemperature func(city string) (degrees float64, err error)
}

func (w *WeatherStub) GetTemperature(city string) (degrees float64, err error) {
	// // return set values
	// return w.Degrees, w.Err

	// call the function
	degrees, err = w.FuncGetTemperature(city)
	return
}