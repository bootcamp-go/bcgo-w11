package handler

import (
	"go-testing/integration/internal"
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
)

// ItemJSON is a struct that represents the JSON response
type ItemJSON struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

// ItemDefault is a struct that represents the methods that a handler
type ItemDefault struct {
	// rp is the repository
	sv internal.ItemService
}

// NewItemDefault returns a new ItemDefault
func NewItemDefault(sv internal.ItemService) *ItemDefault {
	return &ItemDefault{sv}
}

// FindById returns an item by its ID
func (h *ItemDefault) FindById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Error(w, http.StatusBadRequest, "invalid id")
			return
		}

		// process
		i, err := h.sv.FindById(id)
		if err != nil {
			switch err {
			case internal.ErrServiceNotFound:
				response.Error(w, http.StatusNotFound, "item not found")
			default:
				response.Error(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		// response
		// - deserialize the item to JSON
		data := ItemJSON{
			ID:          i.ID,
			Name:        i.Name,
			Description: i.Description,
			Price:       i.Price,
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "item found",
			"data":    data,
		})
	}
}
