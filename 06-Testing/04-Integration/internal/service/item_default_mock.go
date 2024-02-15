package service

import (
	"go-testing/integration/internal"

	"github.com/stretchr/testify/mock"
)

// ItemDefaultMock is a mock for ItemDefault
type ItemDefaultMock struct {
	mock.Mock
}

// NewItemDefaultMock creates a new ItemDefaultMock
func NewItemDefaultMock() *ItemDefaultMock {
	return &ItemDefaultMock{}
}

// FindById is a mock for FindById
func (m *ItemDefaultMock) FindById(id int) (i internal.Item, err error) {
	args := m.Called(id)
	return args.Get(0).(internal.Item), args.Error(1)
}
