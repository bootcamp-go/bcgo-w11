package internal

type ItemService interface {
	FindAll() ([]Item, error)
}