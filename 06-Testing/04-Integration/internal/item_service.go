package internal

import "errors"

var (
	// ErrServiceNotFound is returned when an item is not found in the database
	ErrServiceNotFound = errors.New("service: item not found")
	// ErrServiceInternal is returned when an internal error occurs
	ErrServiceInternal = errors.New("service: internal error")
)

// ItemService is an interface that represents the methods that a service
type ItemService interface {
	// FindById returns all items in the database
	FindById(id int) (i Item, err error)
}
