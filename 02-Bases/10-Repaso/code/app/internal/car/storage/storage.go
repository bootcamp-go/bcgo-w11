package storage

import "errors"

var (
	ErrNotFound = errors.New("not found")
)

type Car struct {
	ID          int64
	Name        string
	Description string
}

type CarStorage interface {
	FindAll() ([]Car, error)
}