package phpparse

import (
	"errors"
	"fmt"
	"gophp/php/ast"
)

func decodeAstData(data any) ([]ast.Stmt, error) {
	value, err := decodeData(data)
	if err != nil {
		return nil, err
	}
	return asSlice[ast.Stmt](value)
}

func decodeData(data any) (any, error) {
	switch value := data.(type) {
	case []any:
		for i, item := range value {
			if newItem, err := decodeData(item); err == nil {
				value[i] = newItem
			} else {
				return nil, err
			}
		}
	case map[string]any:
		for key, item := range value {
			if newItem, err := decodeData(item); err == nil {
				value[key] = newItem
			} else {
				return nil, err
			}
		}
		if nodeType, ok := value["nodeType"].(string); ok {
			node, err := decodeNode(value)
			if node == nil || err != nil {
				return nil, errors.New("node decode failed: nodeType=" + nodeType)
			}
			data = node
		}
	}

	return data, nil
}

var (
	TypeError = errors.New("type error")
)

func asType[T any](data any) (T, error) {
	val, ok := data.(T)
	if !ok {
		var tmp T
		return tmp, fmt.Errorf("data is not expected type: f=decode.Type(), err=%w", TypeError)
	}
	return val, nil
}

func asSlice[T any](data any) ([]T, error) {
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

func asMap[T any](data any) (map[string]T, error) {
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
