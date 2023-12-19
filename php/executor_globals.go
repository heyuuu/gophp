package php

import "github.com/heyuuu/gophp/php/types"

type ClassTable = *types.Table[*types.Class]
type FunctionTable = *types.Table[*types.Function]
type ConstantTable = *types.Table[*types.Constant]

// ExecutorGlobals
type ExecutorGlobals struct {
	constantTable ConstantTable
	functionTable FunctionTable
	classTable    ClassTable
}

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
}

func (e *ExecutorGlobals) ConstantTable() ConstantTable { return e.constantTable }
func (e *ExecutorGlobals) FunctionTable() FunctionTable { return e.functionTable }
func (e *ExecutorGlobals) ClassTable() ClassTable       { return e.classTable }

func (e *ExecutorGlobals) FindFunction(name string) *types.Function {
	// todo 完善 caseIgnore 及命名空间处理
	return e.functionTable.Get(name)
}
