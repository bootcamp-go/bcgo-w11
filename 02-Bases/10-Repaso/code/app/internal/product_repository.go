package internal

type ProductRepository interface {
	FindAll() ([]Product, error)
}