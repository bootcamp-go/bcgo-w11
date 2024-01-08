package main

import (
	"app/interfaces/s2/internal/products"
	"fmt"
)

func main() {
	// app
	// - create small product
	sm, err := products.Factory(products.SmallSize, 1000.0)
	if err != "" {
		fmt.Println(err)
		return
	}
	fmt.Println("Small product price:", sm.Price())
	
	// - create medium product
	md, err := products.Factory(products.MediumSize, 1000.0)
	if err != "" {
		fmt.Println(err)
		return
	}
	fmt.Println("Medium product price:", md.Price())

	// - create large product
	lg, err := products.Factory(products.LargeSize, 1000.0)
	if err != "" {
		fmt.Println(err)
		return
	}
	fmt.Println("Large product price:", lg.Price())
}