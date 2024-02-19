package perr

import (
	"errors"
	"fmt"
)

func Internal(message string) error {
	err := errors.New(message)
	return WithStack(err)
}

func Internalf(format string, a ...any) error {
	message := fmt.Sprintf(format, a...)
	return Internal(message)
}

func Panic(message string) {
	panic(Internal(message))
}

func Unreachable() error {
	return Internal("unreachable")
}

func Todo() error {
	return Internal("todo")
}

func Todof(format string, a ...any) error {
	message := "todo: " + fmt.Sprintf(format, a...)
	return Internal(message)
}

var (
	ErrExit = errors.New("Exit")
)
