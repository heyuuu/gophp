package perr

import (
	"errors"
)

func New(message string) error {
	err := errors.New(message)
	return WithStack(err)
}
