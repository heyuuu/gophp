package phpparse

import (
	"encoding/json"
	"fmt"
	"github.com/heyuuu/gophp/compile/ast"
	"strings"
)

type result struct {
	Ok    bool   `json:"ok"`
	Data  string `json:"data"`
	Error string `json:"error"`
}

func decodeAstData(binData string) (file *ast.File, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("php parse decode ast data failed: %v", e)
		}
	}()

	// json decode
	var data any
	decoder := json.NewDecoder(strings.NewReader(binData))
	decoder.UseNumber()
	if err = decoder.Decode(&data); err != nil {
		return nil, err
	}

	// build node
	value, err := decodeData(data)
	if err != nil {
		return nil, err
	}

	stmts := asStmtList(value)
	return buildAstFile(stmts), nil
}

func buildAstFile(stmts []ast.Stmt) *ast.File {
	// 拆分 declare 语句、全局代码和命名空间代码
	var declareStmts []*ast.DeclareStmt
	var globalStmts []ast.Stmt
	var namespaceStmts []*ast.NamespaceStmt
	for _, astStmt := range stmts {
		switch s := astStmt.(type) {
		case *ast.DeclareStmt:
			declareStmts = append(declareStmts, s)
		case *ast.NamespaceStmt:
			namespaceStmts = append(namespaceStmts, s)
		default:
			globalStmts = append(globalStmts, s)
		}
	}
	if len(globalStmts) > 0 && len(namespaceStmts) > 0 {
		panic("Global code should be enclosed in global namespace declaration")
	}

	if len(namespaceStmts) == 0 {
		namespaceStmts = append(namespaceStmts, &ast.NamespaceStmt{
			Name:  nil,
			Stmts: globalStmts,
		})
	}

	return &ast.File{
		Declares:   declareStmts,
		Namespaces: namespaceStmts,
	}
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
func asSliceItemNullable[T any](data any) []T {
	if data == nil {
		return nil
	}

	var arr = data.([]any)
	var items = make([]T, len(arr))
	for i, item := range arr {
		if item != nil {
			items[i] = item.(T)
		}
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
