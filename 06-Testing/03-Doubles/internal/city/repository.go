package city

import "errors"

var (
	// ErrCityDuplicated is returned when the city is duplicated.
	ErrCityDuplicated = errors.New("city duplicated")
)

// RepositoryWriter is the interface that wraps the basic methods to write cities in a repository.
// types of writes:
// - write operations (pure write and single-step)
// - read-write operations
type RepositoryWriter interface {
	// SaveCity saves the given city in the repository.
	SaveCity(c *City) (err error)
}