package types

import (
	"github.com/heyuuu/gophp/compile/ast"
	"github.com/heyuuu/gophp/php/contracts"
)

type FunctionType uint8

const (
	TypeInternalFunction FunctionType = 1
	TypeUserFunction     FunctionType = 2
	TypeEvalCode         FunctionType = 4
)

/* zend_internal_function_handler */
type ZifHandler func(executeData contracts.IExecuteData, returnValue *Zval)

type Function struct {
	typ          FunctionType
	functionName string
	scope        *Class
	argInfos     []ArgInfo

	// fields for internal function
	handler ZifHandler

	// fields for user function
	stmts   []ast.Stmt
	astFile *ast.File
}

func (f *Function) Type() FunctionType { return f.typ }

func (f *Function) IsInternalFunction() bool { return f.typ == TypeInternalFunction }
func (f *Function) IsUserFunction() bool     { return f.typ == TypeUserFunction }

func (f *Function) Name() string        { return f.functionName }
func (f *Function) ArgInfos() []ArgInfo { return f.argInfos }

func (f *Function) Handler() ZifHandler           { return f.handler }
func (f *Function) SetHandler(handler ZifHandler) { f.handler = handler }
