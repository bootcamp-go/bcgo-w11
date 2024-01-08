package storage

type Slice struct {
	cars []Car
}

func (s *Slice) FindAll() ([]Car, error) {
	// make a copy of the slice
	cars := make([]Car, len(s.cars))
	copy(cars, s.cars)
	return cars, nil
}