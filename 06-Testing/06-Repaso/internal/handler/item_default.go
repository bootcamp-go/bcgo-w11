package handler

import (
	"go-testing/repaso/internal"
	"net/http"

	"github.com/bootcamp-go/web/response"
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
	rp internal.RepositoryItem
}

// NewItemDefault returns a new ItemDefault
func NewItemDefault(rp internal.RepositoryItem) *ItemDefault {
	return &ItemDefault{rp}
}

// FindAll returns all items in the database
func (h *ItemDefault) FindAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...

		// process
		i, err := h.rp.FindAll(6)
		if err != nil {
			response.Text(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		
		// on purpose
		h.rp.FindAll(7)

		// response
		data := make([]ItemJSON, len(i))
		for k, v := range i {
			data[k] = ItemJSON{
				ID:          v.ID,
				Name:        v.Name,
				Description: v.Description,
				Price:       v.Price,
			}
		}
		response.JSON(w, http.StatusOK, data)
	}
}

// type URLFetcher interface {
// 	FetchUrlPathParam(r *http.Request, key string) (value string)
// }

// func FetchUrlPathParam(url, pattern, key string) (value string) {
// 	// ...
// 	return
// }