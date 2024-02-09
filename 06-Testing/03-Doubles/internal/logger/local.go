package logger

import "log"

// NewLocal returns a new Local logger.
func NewLocal() *Local {
	return &Local{}
}

// Local is a Local logger.
type Local struct {}

func (d *Local) Logf(format string, args ...interface{}) {
	log.Printf(format, args...)
}

func (d *Local) Warnf(format string, args ...interface{}) {
	log.Printf("WARNING: "+format, args...)
}