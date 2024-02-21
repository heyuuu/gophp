package php

import (
	"github.com/heyuuu/gophp/compile/ast"
	"github.com/heyuuu/gophp/compile/parser"
	"github.com/heyuuu/gophp/php/types"
)

func CompileFile(ctx *Context, fileHandle *FileHandle, skipShebang bool) (*types.Function, error) {
	code, err := fileHandle.ReadAllEx()
	if err != nil {
		return nil, err
	}

	astFile, err := parser.ParseCodeEx(code, skipShebang)
	if err != nil {
		return nil, err
	}

	topStmts := sortTopStmts(astFile)
	topFn := types.NewAstFunction("", nil, topStmts)
	topFn.SetFilename(fileHandle.OpenedPath())

	return topFn, nil
}

func sortTopStmts(file *ast.File) []ast.Stmt {
	var preloads []ast.Stmt
	var stmts []ast.Stmt
	for _, namespace := range file.Namespaces {
		for _, stmt := range namespace.Stmts {
			switch stmt.(type) {
			case *ast.FunctionStmt, *ast.ClassStmt, *ast.InterfaceStmt, *ast.TraitStmt:
				preloads = append(preloads, stmt)
			default:
				stmts = append(stmts, stmt)
			}
		}
	}
	return append(preloads, stmts...)
}
