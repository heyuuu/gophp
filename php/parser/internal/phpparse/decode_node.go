package phpparse

import "gophp/php/ast"

func decodeNode(data map[string]any) (node ast.Node, err error) {
	nodeType := data["nodeType"].(string)
	switch nodeType {
	case "Stmt_Expression":
		node = &ast.ExpressionStmt{}
		// todo 待完成
	}

	return nil, nil
}
