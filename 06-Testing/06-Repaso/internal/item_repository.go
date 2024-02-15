package internal

import "errors"

var (
	// ErrNotFound is returned when an item is not found in the database
	ErrNotFound = errors.New("item not found")
)

// RepositoryItem is an interface that represents the methods that a repository
// must implement to interact with the database
type RepositoryItem interface {
	// FindAll returns all items in the database
	FindAll(limit int) (i []Item, err error)
}