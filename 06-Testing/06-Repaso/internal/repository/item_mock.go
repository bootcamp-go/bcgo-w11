package repository

import "go-testing/repaso/internal"

// NewItemMock returns a new ItemMock
func NewItemMock() *ItemMock {
	return &ItemMock{
		FindAllFunc: func() (i []internal.Item, err error) {return},
		Spy: Spy{
			MethodCalls: make(map[string]int),
			MethodArgs:  make(map[string][]Args),
		},
	}
}

type Args []any

// Spy allows us to have more control over the code execution
type Spy struct {
	// MethodCalls
	// - key: method name
	// - value: number of times the method was called
	MethodCalls map[string]int

	// MethodArgs
	// - key: method name
	// - value: calls with their arguments
	MethodArgs map[string][]Args
}

// ItemMock is a struct that represents the methods that a repository
type ItemMock struct {
	// ________________________________________________________
	// externalize returns
	// FindAllItems
	// FindAllItems []internal.Item
	// FindAllError error
	FindAllFunc func() ([]internal.Item, error)
	// ...
	// .
	// .
	// ...
	// .
	// .
	Spy
}

func (mk *ItemMock) FindAll(limit int) (i []internal.Item, err error) {
	// register the method call
	mk.MethodCalls["FindAll"]++

	// register the method arguments
	// MethodArgs := map[string][]Args{
	// 	"FindAll": {
	// 		{6},	-> 0 first registry
	//		{7},	-> 1 second registry
	//		{8},	-> 2 third registry
	// 	},
	mk.MethodArgs["FindAll"] = append(mk.MethodArgs["FindAll"], []any{limit})

	// ________________________________________________________
	// static approach
	// i = []internal.Item{
	// 	{ID: 1, Name: "Item 1", Description: "Description 1", Price: 1.1},
	// 	{ID: 2, Name: "Item 2", Description: "Description 2", Price: 2.2},
	// 	{ID: 3, Name: "Item 3", Description: "Description 3", Price: 3.3},
	// }
	// return

	// ________________________________________________________
	// dynamic approach (not the best)
	// i = mk.FindAllItems
	// err = mk.FindAllError
	// return

	// ________________________________________________________
	// dynamic approach
	i, err = mk.FindAllFunc()
	return
}