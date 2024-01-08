package internal

type ProductService interface {
	FindAll() ([]Product, error)
}