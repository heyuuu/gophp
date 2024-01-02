package types

import (
	"github.com/heyuuu/gophp/compile/ast"
)

type FunctionType uint8

const (
	TypeInternalFunction FunctionType = 1
	TypeUserFunction     FunctionType = 2
	TypeEvalCode         FunctionType = 4
)

type Function struct {
	typ          FunctionType
	functionName string
	scope        *Class
	argInfos     []ArgInfo

	// fields for internal function
	moduleNumber int
	handler      any

	// fields for user function
	stmts   []ast.Stmt
	astFile *ast.File
}

func NewInternalFunction(name string, handler any, moduleNumber int) *Function {
	return &Function{
		typ:          TypeInternalFunction,
		functionName: name,
		handler:      handler,
		moduleNumber: moduleNumber,
	}
}

func (f *Function) Type() FunctionType { return f.typ }

func (f *Function) IsInternalFunction() bool { return f.typ == TypeInternalFunction }
func (f *Function) IsUserFunction() bool     { return f.typ == TypeUserFunction }

func (f *Function) Name() string        { return f.functionName }
func (f *Function) ArgInfos() []ArgInfo { return f.argInfos }

func (f *Function) Handler() any { return f.handler }
