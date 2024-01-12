package repository

import "movies-testing/internal"

// NewTaskMap returns a new TaskMap.
func NewTaskMap(m map[int]internal.Task) *TaskMap {
	// default values
	defaultMap := make(map[int]internal.Task)
	if m != nil {
		defaultMap = m
	}

	return &TaskMap{
		m: defaultMap,
	}
}

// TaskMap implements TaskRepository using a map.
type TaskMap struct {
	// m is the map that stores tasks.
	// - key: task id
	// - value: task
	m map[int]internal.Task
}

// FindByID returns the task with the given ID.
func (r *TaskMap) FindByID(id int) (task internal.Task, err error) {
	// check if the task exists
	task, ok := r.m[id]
	if !ok {
		err = internal.ErrTaskNotFound
		return
	}

	return
}