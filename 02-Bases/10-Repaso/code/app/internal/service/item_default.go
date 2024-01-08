package service

import (
	"app/repaso/internal"
)

type ItemDefault struct {
	rp internal.ItemRepository
}

func (s *ItemDefault) FindAll() ([]internal.Item, error) {
	// ...
	// ...
	// ...
	// ...
	// ...
	return s.rp.FindAll()
}