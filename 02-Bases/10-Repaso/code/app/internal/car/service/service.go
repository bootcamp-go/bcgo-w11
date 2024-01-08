package service

type Car struct {
	ID          int64
	Name        string
	Description string
}

type CarService interface {
	FindAll() ([]Car, error)
}