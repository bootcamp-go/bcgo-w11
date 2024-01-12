package internal

import "errors"

var (
	// ErrTaskNotFound is returned when the task is not found.
	ErrTaskNotFound = errors.New("task not found")
)

// TaskRepository is an interface that represents a task repository.
type TaskRepository interface {
	// FindByID returns the task with the given ID.
	FindByID(id int) (Task, error)
}