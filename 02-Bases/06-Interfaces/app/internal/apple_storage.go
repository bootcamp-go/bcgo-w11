package internal

var (
	// ErrAppleNotFound is returned when an apple is not found
	ErrAppleNotFound = "apple not found"
)

// StorageApple is an interface for storing apples
type StorageApple interface {
	// GetApple returns an apple by its ID
	GetApple(id int) (apple Apple, err string)
}