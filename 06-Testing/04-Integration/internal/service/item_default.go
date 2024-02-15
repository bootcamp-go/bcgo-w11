package service

import (
	"go-testing/integration/internal"
)

// ItemDefault is the default implementation of ItemService
type ItemDefault struct {
	rp internal.ItemRepository
}

// NewItemDefault creates a new ItemDefault
func NewItemDefault(rp internal.ItemRepository) *ItemDefault {
	return &ItemDefault{rp}
}

// FindById finds an item by its id
func (s *ItemDefault) FindById(id int) (i internal.Item, err error) {
	i, err = s.rp.FindById(id)
	if err != nil {
		switch err {
		case internal.ErrRepositoryNotFound:
			err = internal.ErrServiceNotFound
		default:
			err = internal.ErrServiceInternal
		}
		return
	}

	return
}
