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
