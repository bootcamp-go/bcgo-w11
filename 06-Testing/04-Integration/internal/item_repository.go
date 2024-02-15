package internal

import "errors"

var (
	// ErrRepositoryNotFound is returned when an item is not found in the database
	ErrRepositoryNotFound = errors.New("repository: item not found")
	// ErrRepositoryInternal is returned when an internal error occurs
	ErrRepositoryInternal = errors.New("repository: internal error")
)

// RepositoryItem is an interface that represents the methods that a repository
// must implement to interact with the database
type ItemRepository interface {
	// FindById returns all items in the database
	FindById(id int) (i Item, err error)
}
