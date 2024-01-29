package internal

import (
	"errors"
	"time"
)

// Task is a struct that represents a task
type Task struct {
	// ID is the unique identifier of the task
	// - should be unique (auto-incremented)
	ID int
	// Name is the name of the task
	// - should be unique
	Name string
	// Description is the description of the task
	// - 65535 characters max
	Description string
	// Completed is a boolean that represents if the task is completed or not
	// - default value: false
	Completed bool
	// CreatedAt is the date when the task was created
	// - auto-generated
	CreatedAt time.Time
}

var (
	// ErrTaskNotFound is an error that is returned when a task is not found
	ErrTaskNotFound = errors.New("task not found")
	// ErrTaskAlreadyExists is an error that is returned when a task already exists
	ErrTaskAlreadyExists = errors.New("task already exists")
)

// TaskRepository is an interface that represents a task repository
type TaskRepository interface {
	// FindByID returns a task by its ID
	FindByID(id int) (task Task, err error)
	// FindAll returns all tasks
	FindAll() (tasks []Task, err error)

	// Save saves a task
	Save(task *Task) (err error)
	// Update updates a task
	Update(task *Task) (err error)
	// Delete deletes a task
	Delete(id int) (err error)
}
