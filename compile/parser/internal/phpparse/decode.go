package phpparse

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/heyuuu/gophp/compile/ast"
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
	stmts = asStmtList(value)
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
		if _, ok := value["nodeType"].(string); ok {
			if node, err := decodeNode(value); err == nil {
				data = node
			} else {
				return nil, err
			}
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

func asFlags(data any) ast.Flags {
	if data == nil {
		return 0
	}
	val, err := data.(json.Number).Int64()
	if err != nil {
		panic(err)
	}
	return ast.Flags(val)
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

func asStmtList(data any) []ast.Stmt {
	var stmts []ast.Stmt
	for _, stmt := range asSlice[ast.Stmt](data) {
		switch stmt.(type) {
		case *ast.BlockStmt:
			stmts = append(stmts, stmt.(*ast.BlockStmt).List...)
		default:
			stmts = append(stmts, stmt)
		}
	}
	return stmts
}

func asTypeHint(data any) ast.TypeHint {
	if data == nil {
		return nil
	}

	switch node := data.(type) {
	case *ast.Ident:
		return &ast.SimpleType{
			Name: ast.NewName(node.Name),
		}
	case *ast.Name:
		return &ast.SimpleType{Name: node}
	default:
		return data.(ast.TypeHint)
	}
}

func asTypeHints(data any) []ast.TypeHint {
	if data == nil {
		return nil
	}

	var items []ast.TypeHint
	for _, item := range data.([]any) {
		items = append(items, asTypeHint(item))
	}
	return items
}

func concatName(name1 *ast.Name, name2 *ast.Name) *ast.Name {
	// 合并 Parts
	parts := append(append([]string{}, name1.Parts...), name2.Parts...)

	// newName 继承 name1 的其他属性
	return &ast.Name{Kind: name1.Kind, Parts: parts}
}
