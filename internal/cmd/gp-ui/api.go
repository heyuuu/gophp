package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/heyuuu/gophp/compile/ast"
	"github.com/heyuuu/gophp/compile/parser"
	"github.com/heyuuu/gophp/kits/vardumper"
	"github.com/heyuuu/gophp/php"
	"net/http"
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
	TypeAst      = "AST"
	TypeAstPrint = "AST-print"
	TypeRun      = "Run"
)

type ApiTypeResult struct {
	Type    string `json:"type"`
	Content string `json:"content"`
}

func apiHandler(request *http.Request) (data any, err error) {
	err = request.ParseForm()
	if err != nil {
		return
	}

	input := request.FormValue("input")
	if input == "" {
		return nil, errors.New("input is empty")
	}

	result, err := parseCode(input)
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

func parseCode(code string) (result []ApiTypeResult, err error) {
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

	// run code
	output := runCode(code)
	result = append(result, ApiTypeResult{Type: TypeRun, Content: output})
	return
}

var engine = php.NewEngine()

func runCode(code string) (output string) {
	defer func() {
		if e := recover(); e != nil {
			output = fmt.Sprintf("Execute panic: %s", e)
		}
	}()

	var buf strings.Builder

	ctx := engine.NewContext()

	buf.WriteString(">>> output start:\n")
	ctx.OG().PushHandler(&buf)

	fileHandle := php.NewFileHandleByString(code)
	retval, err := php.ExecuteScript(ctx, fileHandle, false)

	buf.WriteString("\n>>> output end\n\n")
	if err != nil {
		buf.WriteString("Execute failed: " + err.Error())
	} else {
		buf.WriteString(fmt.Sprintf("Execute succed, retval = %v", retval))
	}

	return buf.String()
}
