package repository

import (
	"app/repaso/internal"
)

type ItemSlice struct {
	items []internal.Item
}

func (s *ItemSlice) FindAll() ([]internal.Item, error) {
	// make a copy of the slice
	items := make([]internal.Item, len(s.items))
	copy(items, s.items)
	return items, nil
}	