package internal

import "errors"

var (
	// ErrInvalidQuery is an error that occurs when a query is invalid
	ErrInvalidQuery = errors.New("repository: invalid query")
)

// ItemRepository is an interface that represents a repository of items
type ItemRepository interface {
	// FindAll returns all items
	FindAll(query map[string]any, limit int) (map[int]Item, error)
}