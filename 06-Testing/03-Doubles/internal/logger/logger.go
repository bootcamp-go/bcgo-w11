package logger

// Logger is the interface that wraps the basic logging methods.
type Logger interface {
	// Logf logs a message with the given format and arguments.
	Logf(format string, args ...interface{})
	// Warnf logs a warning message with the given format and arguments.
	Warnf(format string, args ...interface{})
}