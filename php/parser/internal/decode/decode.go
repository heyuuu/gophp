package decode

import (
	"errors"
	"fmt"
)

var (
	TypeError = errors.New("type error")
)

func Type[T any](data any) (T, error) {
	val, ok := data.(T)
	if !ok {
		return nil, fmt.Errorf("data is not expected type: f=decode.Type(), err=%w", TypeError)
	}
	return val, nil
}

func Slice[T any](data any) ([]T, error) {
	items, ok := data.([]any)
	if !ok {
		return nil, fmt.Errorf("data is not a slice: f=decode.Slice(), err=%w", TypeError)
	}
	var result []T
	for i, item := range items {
		if val, ok := item.(T); ok {
			result = append(result, val)
		} else {
			return nil, fmt.Errorf("data[%d] is not expected type: f=decode.Slice(), err=%w", i, TypeError)
		}
	}
	return result, nil
}

func Map[T any](data any) (map[string]T, error) {
	items, ok := data.(map[string]any)
	if !ok {
		return nil, fmt.Errorf("data is not a slice: f=decode.Map(), err=%w", TypeError)
	}

	var result = make(map[string]T)
	for key, item := range items {
		if val, ok := item.(T); ok {
			result[key] = val
		} else {
			return nil, fmt.Errorf("data[%s] is not expected type: f=decode.Map(), err=%w", key, TypeError)
		}
	}
	return result, nil
}
