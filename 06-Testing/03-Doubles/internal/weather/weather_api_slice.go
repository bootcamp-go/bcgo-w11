package weather

// NewWeatherAPISlice returns a new WeatherAPISlice.
func NewWeatherAPISlice(citiesTemperature map[string]float64) *WeatherAPISlice {
	// default config
	defaultDb := make(map[string]float64)
	if citiesTemperature != nil {
		defaultDb = citiesTemperature
	}

	return &WeatherAPISlice{
		citiesTemperature: defaultDb,
	}
}

// WeatherAPISlice is a WeatherAPI that stores the temperature for each city in a map.
type WeatherAPISlice struct {
	// cityTemperature is a map that stores the temperature for each city.
	citiesTemperature map[string]float64
}

// GetTemperature returns the temperature for the given city. The temperature is in degrees Celsius.
func (w *WeatherAPISlice) GetTemperature(city string) (degrees float64, err error) {
	degrees, ok := w.citiesTemperature[city]
	if !ok {
		err = ErrCityNotFound
		return
	}

	return
}