package phpparse

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"gophp/php/ast"
)

type result struct {
	Ok    bool   `json:"ok"`
	Data  string `json:"data"`
	Error string `json:"error"`
}

func decodeOutput(output []byte) ([]ast.Stmt, error) {
	var res result
	if err := json.Unmarshal(output, &res); err != nil {
		return nil, err
	}

	if !res.Ok {
		return nil, errors.New(res.Error)
	}

	return decodeAstData([]byte(res.Data))
}

func decodeAstData(binData []byte) (stmts []ast.Stmt, err error) {
	defer func() {
		if fault := recover(); fault != nil {
			err = fmt.Errorf("decode ast data failed: %v", fault)
		}
	}()

	// json decode
	var data any
	decoder := json.NewDecoder(bytes.NewReader(binData))
	decoder.UseNumber()
	if err = decoder.Decode(&data); err != nil {
		return nil, err
	}

	// build node
	value, err := decodeData(data)
	if err != nil {
		return nil, err
	}
	stmts = asSlice[ast.Stmt](value)
	return stmts, nil
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

func asInt(data any) int {
	val, err := data.(json.Number).Int64()
	if err != nil {
		panic(err)
	}
	return int(val)
}

func asFloat(data any) float64 {
	val, err := data.(json.Number).Float64()
	if err != nil {
		panic(err)
	}
	return val
}

func asTypeOrNil[T any](data any) T {
	if data == nil {
		var tmp T
		return tmp
	}
	return data.(T)
}

func asSlice[T any](data any) []T {
	if data == nil {
		return nil
	}

	var items []T
	for _, item := range data.([]any) {
		items = append(items, item.(T))
	}
	return items
}
