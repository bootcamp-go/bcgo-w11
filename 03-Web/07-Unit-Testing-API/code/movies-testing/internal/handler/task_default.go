package handler

import (
	"errors"
	"movies-testing/internal"
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/response"

	"github.com/go-chi/chi/v5"
)

// NewTaskDefault returns a new TaskDefault.
func NewTaskDefault(rp internal.TaskRepository) *TaskDefault {
	return &TaskDefault{
		rp: rp,
	}
}

// TaskJSON is an struct that represents a task in JSON format.
type TaskJSON struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

// TaskDefault is an implementation of handlers for tasks.
type TaskDefault struct {
	// rp is the task repository.
	rp internal.TaskRepository
}

// GetTask returns the task with the given ID.
func (h *TaskDefault) GetTask() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get the id from the context
		// r.Context().Value("id")
		// r.Context().Value(chi.RouteCtxKey) -> map[string]string -> "id"
		// request
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Error(w, http.StatusBadRequest, "invalid id")
			return
		}

		// get the task
		task, err := h.rp.FindByID(id)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrTaskNotFound):
				response.Error(w, http.StatusNotFound, "task not found")
			default:
				response.Error(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		// write the task in the response body
		data := TaskJSON{
			ID:          task.ID,
			Name:        task.Name,
			Description: task.Description,
			Completed:   task.Completed,
		}
		response.JSON(w, http.StatusOK, data)
	}
}