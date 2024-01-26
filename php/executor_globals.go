package php

import (
	"github.com/heyuuu/gophp/php/types"
	"strings"
)

type ClassTable = *types.Table[*types.Class]
type FunctionTable = *types.Table[*types.Function]
type ConstantTable = *types.Table[*types.Constant]

// ExecutorGlobals
type ExecutorGlobals struct {
	errorSuppress  int
	errorReporting int
	precision      int

	constantTable ConstantTable
	functionTable FunctionTable
	classTable    ClassTable
}

func (e *ExecutorGlobals) ErrorReporting() int                  { return e.errorReporting }
func (e *ExecutorGlobals) SetErrorReporting(errorReporting int) { e.errorReporting = errorReporting }
func (e *ExecutorGlobals) Precision() int                       { return e.precision }
func (e *ExecutorGlobals) SetPrecision(precision int)           { e.precision = precision }

func (e *ExecutorGlobals) Init(base *ExecutorGlobals) {
	if base != nil {
		e.constantTable = base.constantTable.Clone()
		e.functionTable = base.functionTable.Clone()
		e.classTable = base.classTable.Clone()
	} else {
		e.constantTable = types.NewTable[*types.Constant]()
		e.functionTable = types.NewTable[*types.Function]()
		e.classTable = types.NewTable[*types.Class]()
	}
	// todo init by ini
	e.precision = 14
}

func (e *ExecutorGlobals) ConstantTable() ConstantTable { return e.constantTable }
func (e *ExecutorGlobals) FunctionTable() FunctionTable { return e.functionTable }
func (e *ExecutorGlobals) ClassTable() ClassTable       { return e.classTable }

func (e *ExecutorGlobals) FindFunction(name string) *types.Function {
	// todo 完善 caseIgnore 及命名空间处理
	lcName := strings.ToLower(name)
	return e.functionTable.Get(lcName)
}

func (e *ExecutorGlobals) HasException() bool {
	// todo
	return false
}

func (e *ExecutorGlobals) ErrorSuppress() bool {
	return e.errorSuppress > 0
}
func (e *ExecutorGlobals) ErrorSuppressScope(block func()) {
	e.errorSuppress += 1
	defer func() {
		if e.errorSuppress > 0 {
			e.errorSuppress--
		}
	}()

	block()
}
