package perr

import (
	"errors"
)

func Panic(message string) {
	panic(NewInternal(message))
}

func NewInternal(message string) error {
	err := errors.New(message)
	return WithStack(err)
}
