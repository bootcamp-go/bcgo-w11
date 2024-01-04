package main

import (
	"errors"
	"strconv"
)

type CustomError struct {
	Msg string
	Code int
}

func (e *CustomError) Error() string {
	return e.Msg + " " + strconv.Itoa(e.Code)
}

var err error = errors.New("this is an error")