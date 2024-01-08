package internal

type ItemRepository interface {
	FindAll() ([]Item, error)
}