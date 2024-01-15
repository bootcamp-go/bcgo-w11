package repository

import "app/10-repaso/internal"

// NewItemMap returns a new ItemMap
func NewItemMap(items map[int]internal.Item) *ItemMap {
	return &ItemMap{
		items: items,
	}
}

// ItemMap is an implementation of ItemRepository that uses a map
type ItemMap struct {
	// items is the map of items
	items map[int]internal.Item
}

// FindAll returns all items
func (r *ItemMap) FindAll(query map[string]any, limit int) (items map[int]internal.Item, err error) {
	// items
	items = make(map[int]internal.Item)

	// filter
	querySet := len(query) > 0
	for id, item := range r.items {
		// if query is set, filter
		if querySet {
			// price min
			priceMin, ok := query["price_min"]
			if !ok {
				err = internal.ErrInvalidQuery
				return
			}
			priceMinFloat, ok := priceMin.(float64)
			if !ok {
				err = internal.ErrInvalidQuery
				return
			}

			// price max
			priceMax, ok := query["price_max"]
			if !ok {
				err = internal.ErrInvalidQuery
				return
			}
			priceMaxFloat, ok := priceMax.(float64)
			if !ok {
				err = internal.ErrInvalidQuery
				return
			}

			// filter
			if item.Price < priceMinFloat || item.Price > priceMaxFloat {
				continue
			}
		}

		// by default add item
		items[id] = item
	}

	// limit
	limited := make(map[int]internal.Item)
	if limit > 0 {
		var cc int
		for id, item := range items {
			// if limit is reached, stop
			if cc == limit {
				break
			}

			// add item
			limited[id] = item
		}
	}

	// return
	items = limited
	return
}