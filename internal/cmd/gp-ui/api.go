package main

import (
	"encoding/json"
	"errors"
	"github.com/heyuuu/gophp/compile/ast"
	"github.com/heyuuu/gophp/compile/ir"
	"github.com/heyuuu/gophp/compile/parser"
	"github.com/heyuuu/gophp/shim/maps"
	"github.com/heyuuu/gophp/utils/vardumper"
	"net/http"
	"sort"
	"strings"
)

type ApiResponse[T any] struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
	Data  T      `json:"data"`
}

func ApiWrapHandler(handler func(r *http.Request) (any, error)) http.HandlerFunc {
	return wrapHandler(func(request *http.Request) ([]byte, error) {
		var res ApiResponse[any]
		if data, err := handler(request); err == nil {
			res.Code = 0
			res.Data = data
		} else {
			res.Code = 1
			res.Error = err.Error()
		}
		return json.Marshal(res)
	})
}

//
const (
	TypeAst       = "AST"
	TypeAstPrint  = "AST-print"
	TypeIr        = "IR"
	TypeIrPrint   = "IR-print"
	TypeIrProject = "IR-project"
)

type ApiTypeResult struct {
	Type    string `json:"type"`
	Content string `json:"content"`
}

func apiHandler2(request *http.Request) (data any, err error) {
	err = request.ParseForm()
	if err != nil {
		return
	}

	input := request.FormValue("input")
	if input == "" {
		return nil, errors.New("input is empty")
	}

	result, err := parseCodeEx(input)
	if err != nil {
		return nil, err
	}

	return struct {
		Input  string          `json:"input"`
		Result []ApiTypeResult `json:"result"`
	}{
		Input:  input,
		Result: result,
	}, nil
}

func parseCodeEx(code string) (result []ApiTypeResult, err error) {
	// Ast
	astNodes, err := parser.ParseCode(code)
	if err != nil {
		return
	}

	astDump := vardumper.Sprint(astNodes)
	result = append(result, ApiTypeResult{Type: TypeAst, Content: astDump})

	astPrint, err := ast.PrintFile(astNodes)
	if err != nil {
		return
	}
	result = append(result, ApiTypeResult{Type: TypeAstPrint, Content: astPrint})

	// IR
	irFile, err := ir.ParseAstFile(astNodes)
	result = append(result, ApiTypeResult{Type: TypeIr, Content: vardumper.Sprint(irFile)})
	if err != nil {
		return
	}

	irPrint, err := printIrFileAsProject(irFile)
	if err != nil {
		return
	}
	result = append(result, ApiTypeResult{Type: TypeIrPrint, Content: irPrint})

	return
}

func printIrFileAsProject(irFile *ir.File) (string, error) {
	irProj := ir.NewProject()
	_ = irProj.AddFile("__main__", irFile)

	contents, err := ir.PrintProject(irProj)
	if err != nil {
		return "", err
	}

	keys := maps.Keys(contents)
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	var buf strings.Builder
	for _, key := range keys {
		content := contents[key]
		buf.WriteString(content)
	}

	return buf.String(), nil
}
