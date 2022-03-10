package verror

import (
	"fmt"
)

type CastError struct {
	Value interface{}
	Desc  string
}

func (c CastError) Error() string {
	return fmt.Sprintf("%v %s", c.Value, c.Desc)
}

func NewCastError(value interface{}, desc string) *CastError {
	return &CastError{Value: value, Desc: desc}
}

func CastBoolError(value interface{}) error {
	return NewCastError(value, "cast bool error")
}

func CastDoubleError(value interface{}) error {
	return NewCastError(value, "cast double error")
}

func CastIntegerError(value interface{}) error {
	return NewCastError(value, "cast integer error")
}

func CastStringError(value interface{}) error {
	return NewCastError(value, "cast string error")
}
