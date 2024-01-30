package service

import "app/internal"

// NewProductsDefault creates new default service for product entity.
func NewProductsDefault(rp internal.RepositoryProduct) *ProductsDefault {
	return &ProductsDefault{rp}
}

// ProductsDefault is the default service implementation for product entity.
type ProductsDefault struct {
	// rp is the repository for product entity.
	rp internal.RepositoryProduct
}

// FindAll returns all products.
func (s *ProductsDefault) FindAll() (p []internal.Product, err error) {
	p, err = s.rp.FindAll()
	return
}

// Save saves the product.
func (s *ProductsDefault) Save(p *internal.Product) (err error) {
	err = s.rp.Save(p)
	return
}