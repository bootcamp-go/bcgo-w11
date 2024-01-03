package storage

import "app/internal"

func NewAppleMap(apples map[int]internal.Apple, lastId int) *AppleMap {
	// default values
	defaultApples := make(map[int]internal.Apple)
	defaultLastID := 0
	if apples != nil {
		defaultApples = apples
	}
	if lastId != 0 {
		defaultLastID = lastId
	}

	return &AppleMap{
		apples: defaultApples,
		lastID: defaultLastID,
	}
}

// AppleMap is a storage for apples that implements StorageApple
type AppleMap struct {
	// apples is a map of apples
	// - key: apple ID
	// - value: apple
	apples map[int]internal.Apple

	// lastID is the last apple ID
	lastID int
}

func (a *AppleMap) GetApple(id int) (apple internal.Apple, err string) {
	// check if apple exists
	apple, ok := a.apples[id]
	if !ok {
		err = internal.ErrAppleNotFound
		return
	}

	return
}