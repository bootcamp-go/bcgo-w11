package logger

// NewDummy returns a new instance of the Dummy logger.
func NewDummy() *Dummy {
	return &Dummy{}
}

// Dummy is a dummy logger that does nothing.
type Dummy struct{}

// Logf logs a message.
func (d *Dummy) Logf(format string, v ...interface{}) {
	// do nothing
}

// Warnf logs a warning message.
func (d *Dummy) Warnf(format string, v ...interface{}) {
	// do nothing
}