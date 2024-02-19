package php

import (
	"github.com/heyuuu/gophp/php/assert"
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
	symbolTable    ISymtable      `get:""`
	errorReporting perr.ErrorType `prop:""`
	exitStatus     int            `prop:""`
	precision      int            `prop:""`

	constantTable ConstantTable `get:""`
	functionTable FunctionTable `get:""`
	classTable    ClassTable    `get:""`

	currentExecuteData *ExecuteData `prop:""`

	nextObjectHandle uint
}

func (eg *ExecutorGlobals) InitBase(ctx *Context) {
	*eg = ExecutorGlobals{
		ctx:            ctx,
		errorReporting: perr.E_ALL, // todo init by ini
		precision:      14,         // todo init by ini
		constantTable:  types.NewTable[*types.Constant](),
		functionTable:  types.NewTable[*types.Function](),
		classTable:     types.NewTable[*types.Class](),
	}
}

func (eg *ExecutorGlobals) Init(ctx *Context, base *ExecutorGlobals) {
	assert.Assert(base != nil)
	*eg = ExecutorGlobals{
		ctx:            ctx,
		symbolTable:    NewSymtable(),
		errorReporting: base.errorReporting,
		exitStatus:     0,
		precision:      base.precision,
		constantTable:  base.constantTable.Clone(),
		functionTable:  base.functionTable.Clone(),
		classTable:     base.classTable.Clone(),
	}
}

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
