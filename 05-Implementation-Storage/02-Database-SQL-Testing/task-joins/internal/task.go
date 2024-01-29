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
	// OwnerID is the ID of the user that owns the task
	// - One-to-many relationship
	OwnerID int
}

// TaskWithOwner is a struct that represents a task with its owner
type TaskWithOwner struct {
	ID          int
	Name        string
	Description string
	Completed   bool
	CreatedAt   time.Time
	OwnerName   string
	OwnerEmail  string
}

var (
	// ErrTaskNotFound is an error that is returned when a task is not found
	ErrTaskNotFound = errors.New("task not found")
	// ErrTaskAlreadyExists is an error that is returned when a task already exists
	ErrTaskAlreadyExists = errors.New("task already exists")
	// ErrTaskHasNoOwner is an error that is returned when a task has no owner
	ErrTaskHasNoOwner = errors.New("task has no owner")
)

// TaskRepository is an interface that represents a task repository
type TaskRepository interface {
	// FindAll returns all tasks
	FindAll() (tasks []Task, err error)
	// FindByIDWithOwner returns a task by its ID with its owner
	FindByIDWithOwner(id int) (task Task, err error)

	// Save saves a task
	Save(task *Task) (err error)
}
