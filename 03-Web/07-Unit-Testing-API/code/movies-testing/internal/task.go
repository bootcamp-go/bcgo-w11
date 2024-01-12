package internal

// Task is an struct that represents a task.
type Task struct {
	// ID is the unique identifier of the task.
	ID          int
	// Name is the name of the task.
	Name        string
	// Description is the description of the task.
	Description string
	// Completed is the status of the task.
	Completed   bool
}