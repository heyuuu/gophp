package faults

import "errors"

type CoreError struct {
}

func (e CoreError) Error() string {
	//TODO implement me
	panic("implement me")
}

func IsError(err error) *CoreError {
	var target CoreError
	if errors.As(err, &target) {
		return &target
	}
	return nil
}
