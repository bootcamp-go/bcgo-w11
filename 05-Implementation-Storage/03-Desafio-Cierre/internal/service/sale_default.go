package service

import "app/internal"

// NewSalesDefault creates new default service for sale entity.
func NewSalesDefault(rp internal.RepositorySale) *SalesDefault {
	return &SalesDefault{rp}
}

// SalesDefault is the default service implementation for sale entity.
type SalesDefault struct {
	// rp is the repository for sale entity.
	rp internal.RepositorySale
}

// FindAll returns all sales.
func (sv *SalesDefault) FindAll() (s []internal.Sale, err error) {
	s, err = sv.rp.FindAll()
	return
}

// Save saves the sale.
func (sv *SalesDefault) Save(s *internal.Sale) (err error) {
	err = sv.rp.Save(s)
	return
}