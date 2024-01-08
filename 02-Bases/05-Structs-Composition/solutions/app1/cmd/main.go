package main

import (
	"app/structs/s1/internal"
	"app/structs/s1/internal/repository"
	"fmt"
)

func main() {
	// app
	// - dependencies
	rp := repository.NewProductMap(nil, 0)
	
	// - save a product
	product := internal.Product{
		Name:        "product 1",
		Price:       10.5,
		Description: "description 1",
		Category:    "category 1",
	}
	product, err := rp.Save(product)
	if err != "" {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("product:", product)

	// - find all products
	products, err := rp.FindAll()
	if err != "" {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("products:", products)

	// - find a product by ID
	product, err = rp.FindByID(1)
	if err != "" {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("product:", product)
}