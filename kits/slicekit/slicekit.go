package slicekit

import (
	"github.com/heyuuu/gophp/shim/slices"
)

func Map[T any, R any](slice []T, mapper func(T) R) []R {
	if len(slice) == 0 {
		return nil
	}

	result := make([]R, len(slice))
	for i, item := range slice {
		result[i] = mapper(item)
	}
	return result
}

func Each[T any](slice []T, handler func(T)) {
	for _, item := range slice {
		handler(item)
	}
}

func EachReserve[T any](slice []T, handler func(T)) {
	for i := len(slice) - 1; i >= 0; i-- {
		item := slice[i]
		handler(item)
	}
}

func EachEx[T any](slice []T, handler func(T) error) error {
	for _, item := range slice {
		err := handler(item)
		if err != nil {
			return err
		}
	}
	return nil
}

func EachReserveEx[T any](slice []T, handler func(T) error) error {
	for i := len(slice) - 1; i >= 0; i-- {
		item := slice[i]
		err := handler(item)
		if err != nil {
			return err
		}
	}
	return nil
}

func First[T any](slice []T) (T, bool) {
	if len(slice) > 0 {
		return slice[0], true
	}
	var temp T
	return temp, false
}

func FirstFunc[T any](slice []T, predicate func(T) bool) (T, bool) {
	for _, item := range slice {
		if predicate(item) {
			return item, true
		}
	}
	var temp T
	return temp, false
}

func Last[T any](slice []T) (T, bool) {
	if len(slice) > 0 {
		return slice[len(slice)-1], true
	}
	var temp T
	return temp, false
}

func LastFunc[T any](slice []T, predicate func(T) bool) (T, bool) {
	for i := len(slice) - 1; i >= 0; i-- {
		item := slice[i]
		if predicate(item) {
			return item, true
		}
	}
	var temp T
	return temp, false
}

func Push[T any](slicePtr *[]T, item T) {
	*slicePtr = append(*slicePtr, item)
}

func Pop[T any](slicePtr *[]T) (T, bool) {
	slice := *slicePtr
	if len(slice) == 0 {
		var tmp T
		return tmp, false
	}

	top := slice[len(slice)-1]
	*slicePtr = slice[:len(slice)-1]
	return top, true
}

func All[T any](slice []T, handler func(T) bool) bool {
	for _, item := range slice {
		if !handler(item) {
			return false
		}
	}
	return true
}

func Any[T any](slice []T, handler func(T) bool) bool {
	for _, item := range slice {
		if handler(item) {
			return true
		}
	}
	return false
}

func Filter[T any](slice []T, handler func(T) bool) []T {
	return slices.DeleteFunc(slice, func(item T) bool {
		return !handler(item)
	})
}

func Unique[T comparable](slice []T) []T {
	if len(slice) == 0 {
		return nil
	}

	existsSet := make(map[T]struct{}, len(slice))
	return slices.DeleteFunc(slice, func(item T) bool {
		if _, exists := existsSet[item]; exists {
			return true
		} else {
			existsSet[item] = struct{}{}
			return false
		}
	})
}

// Concat returns a new slice concatenating the passed in slices.
func Concat[S ~[]E, E any](sliceArgs ...S) S {
	size := 0
	for _, s := range sliceArgs {
		size += len(s)
		if size < 0 {
			panic("len out of range")
		}
	}
	newslice := slices.Grow[S](nil, size)
	for _, s := range sliceArgs {
		newslice = append(newslice, s...)
	}
	return newslice
}
