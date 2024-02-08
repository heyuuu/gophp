package php

import (
	"github.com/heyuuu/gophp/php/perr"
	"github.com/heyuuu/gophp/php/types"
	"strings"
)

type ClassTable = *types.Table[*types.Class]
type FunctionTable = *types.Table[*types.Function]
type ConstantTable = *types.Table[*types.Constant]

// ExecutorGlobals
type ExecutorGlobals struct {
	ctx            *Context
	errorSuppress  int
	errorReporting int
	precision      int

	constantTable ConstantTable
	functionTable FunctionTable
	classTable    ClassTable

	currentExecuteData *ExecuteData

	nextObjectHandle uint
}

func (eg *ExecutorGlobals) ErrorReporting() int                  { return eg.errorReporting }
func (eg *ExecutorGlobals) SetErrorReporting(errorReporting int) { eg.errorReporting = errorReporting }
func (eg *ExecutorGlobals) Precision() int                       { return eg.precision }
func (eg *ExecutorGlobals) SetPrecision(precision int)           { eg.precision = precision }

func (eg *ExecutorGlobals) Init(ctx *Context, base *ExecutorGlobals) {
	eg.ctx = ctx
	eg.errorReporting = int(perr.E_ALL)
	if base != nil {
		eg.constantTable = base.constantTable.Clone()
		eg.functionTable = base.functionTable.Clone()
		eg.classTable = base.classTable.Clone()
	} else {
		eg.constantTable = types.NewTable[*types.Constant]()
		eg.functionTable = types.NewTable[*types.Function]()
		eg.classTable = types.NewTable[*types.Class]()
	}
	// todo init by ini
	eg.precision = 14
}

func (eg *ExecutorGlobals) ConstantTable() ConstantTable { return eg.constantTable }
func (eg *ExecutorGlobals) FunctionTable() FunctionTable { return eg.functionTable }
func (eg *ExecutorGlobals) ClassTable() ClassTable       { return eg.classTable }

func (eg *ExecutorGlobals) FindFunction(name string) *types.Function {
	// todo 完善 caseIgnore 及命名空间处理
	lcName := strings.ToLower(name)
	return eg.functionTable.Get(lcName)
}

func (eg *ExecutorGlobals) HasException() bool {
	// todo
	return false
}
func (eg *ExecutorGlobals) NoException() bool { return !eg.HasException() }

func (eg *ExecutorGlobals) ErrorSuppress() bool {
	return eg.errorSuppress > 0
}
func (eg *ExecutorGlobals) ErrorSuppressScope(block func()) {
	eg.errorSuppress += 1
	defer func() {
		if eg.errorSuppress > 0 {
			eg.errorSuppress--
		}
	}()

	block()
}

func (eg *ExecutorGlobals) CurrentExecuteData() *ExecuteData {
	return eg.currentExecuteData
}
func (eg *ExecutorGlobals) SetCurrentExecuteData(currentExecuteData *ExecuteData) {
	eg.currentExecuteData = currentExecuteData
}

func (eg *ExecutorGlobals) PushExecuteData(ex *ExecuteData) {
	ex.prev = eg.currentExecuteData
	eg.currentExecuteData = ex
}

func (eg *ExecutorGlobals) PopExecuteData() *ExecuteData {
	result := eg.currentExecuteData
	if result != nil {
		eg.currentExecuteData = result.prev
	}
	return result
}

func (eg *ExecutorGlobals) NextObjectHandle() uint {
	eg.nextObjectHandle++
	return eg.nextObjectHandle
}
