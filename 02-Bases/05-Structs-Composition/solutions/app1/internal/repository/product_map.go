package repository

import "app/structs/s1/internal"

// NewProductMap returns a new ProductMap
func NewProductMap(products map[int]internal.Product, lastId int) *ProductMap {
	// default config
	defaultProducts := make(map[int]internal.Product)
	defaultLastID := 0
	if products != nil {
		defaultProducts = products
	}
	if lastId != 0 {
		defaultLastID = lastId
	}

	return &ProductMap{
		products: defaultProducts,
		lastID:   defaultLastID,
	}
}

// ProductMap is a struct that represents a map of products
type ProductMap struct {
	// products is the map of products
	// - key: product ID
	// - value: product
	products map[int]internal.Product
	// lastID is the last ID assigned to a product
	// - autoincremental
	lastID int
}

// FindAll returns all the products
func (pm *ProductMap) FindAll() (pr map[int]internal.Product, err string) {
	// make a copy of the map
	pr = make(map[int]internal.Product)
	for k, v := range pm.products {
		pr[k] = v
	}

	return
}

// FindByID returns the product with the given ID
func (pm *ProductMap) FindByID(id int) (pr internal.Product, err string) {
	// check if the product exists
	product, ok := pm.products[id]
	if !ok {
		err = "product not found"
		return
	}

	pr = product
	return
}

// Save saves the given product
func (pm *ProductMap) Save(product internal.Product) (pr internal.Product, err string) {
	// autoincremental ID
	pm.lastID++
	product.ID = pm.lastID

	// save the product
	pm.products[product.ID] = product

	pr = product
	return
}