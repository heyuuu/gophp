package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/heyuuu/gophp/compile/ast"
	"github.com/heyuuu/gophp/compile/parser"
	"github.com/heyuuu/gophp/kits/vardumper"
	"github.com/heyuuu/gophp/php"
	_ "github.com/heyuuu/gophp/php/boot"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
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
	TypeRawRun   = "Run-Raw"
	TypeDiffRun  = "Run-Diff"
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

	// raw run code
	rawOutput := rawRunCode(code)
	result = append(result, ApiTypeResult{Type: TypeRawRun, Content: rawOutput})

	// diff run
	diff := diffOutput(output, rawOutput)
	result = append(result, ApiTypeResult{Type: TypeDiffRun, Content: diff})

	return
}

func runCode(code string) (output string) {
	var buf strings.Builder
	defer func() {
		if e := recover(); e != nil {
			buf.WriteString(fmt.Sprintf(">>> Execute panic: %v", e))
			log.Printf("%+v\n", e)
		}

		output = buf.String()
	}()

	engine := php.NewEngine()
	err := engine.Start()
	if err != nil {
		buf.WriteString("engine start failed: " + err.Error())
		return
	}

	ctx := engine.NewContext(nil, nil)
	engine.HandleContext(ctx, func(ctx *php.Context) {
		ctx.OG().PushHandler(&buf)

		fileHandle := php.NewFileHandleByString(code)
		_, err = php.ExecuteScript(ctx, fileHandle, false)

		if err != nil {
			buf.WriteString("Execute failed: " + err.Error())
		}
	})

	return
}

func rawRunCode(code string) string {
	if strings.HasPrefix(code, "<?php\n") {
		code = code[6:]
	} else {
		code = "?>" + code
	}

	output, err := runCommand(5*time.Second, "php", "-r", code)
	if err != nil {
		return output + "\n" + err.Error()
	}
	return output
}

func diffOutput(text1 string, text2 string) string {
	f1, err := createTmpFile(text1)
	if err != nil {
		return err.Error()
	}

	f2, err := createTmpFile(text2)
	if err != nil {
		return err.Error()
	}

	output, err := runCommand(5*time.Second, "diff", "-a", f1, f2)
	if output == "" && err != nil {
		return "run diff command failed: " + err.Error()
	}

	return output
}

func createTmpFile(content string) (string, error) {
	f, err := os.CreateTemp(os.TempDir(), "gophp-")
	if err != nil {
		return "", fmt.Errorf("create tmp file failed: %w", err)
	}
	name := f.Name()
	f.WriteString(content)
	f.Close()
	return name, nil
}

func runCommand(timeout time.Duration, name string, args ...string) (string, error) {
	// 超时控制
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, name, args...)
	log.Printf("run command: %s\n", cmd.String())
	if output, err := cmd.CombinedOutput(); err == nil {
		return string(output), nil
	} else if ctx.Err() != nil && errors.Is(ctx.Err(), context.DeadlineExceeded) {
		return string(output), fmt.Errorf("run timeout: %w", err)
	} else {
		return string(output), fmt.Errorf("run fail: %w", err)
	}
}
