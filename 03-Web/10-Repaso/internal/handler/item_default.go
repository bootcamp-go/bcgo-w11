package handler

import (
	"app/10-repaso/internal"
	"errors"
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/response"
)

// NewItemDefault returns a new ItemDefault
func NewItemDefault(rp internal.ItemRepository) *ItemDefault {
	return &ItemDefault{
		rp: rp,
	}
}

// ItemDefault is an implementation of ItemHandler
type ItemDefault struct {
	// rp is the repository of items
	rp internal.ItemRepository
}

// FindAll returns all items
func (h *ItemDefault) FindAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - query: price_min, price_max
		query := make(map[string]any)
		priceMin, ok := r.URL.Query()["price_min"]
		if ok {
			// parse price_min to float64
			priceMinFloat, err := strconv.ParseFloat(priceMin[0], 64)
			if err != nil {
				response.Text(w, http.StatusBadRequest, "price_min must be a number")
				return
			}
			// add price_min to query
			query["price_min"] = priceMinFloat
		}
		priceMax, ok := r.URL.Query()["price_max"]
		if ok {
			// parse price_max to float64
			priceMaxFloat, err := strconv.ParseFloat(priceMax[0], 64)
			if err != nil {
				response.Text(w, http.StatusBadRequest, "price_max must be a number")
				return	
			}
			// add price_max to query
			query["price_max"] = priceMaxFloat
		}

		// process
		items, err := h.rp.FindAll(query, 5)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrInvalidQuery):
				response.Text(w, http.StatusBadRequest, "invalid query")
			default:
				response.Text(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		// response
		response.JSON(w, http.StatusOK, items)
	}
}