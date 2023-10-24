package php

import "github.com/heyuuu/gophp/php/types"

// misc
type ConstantTable = types.Table[*types.Function]
type FunctionTable = types.Table[*types.Function]
type ClassTable = types.Table[*types.Function]

// ExecutorGlobals
type ExecutorGlobals struct {
	constantTable map[string]*types.Constant
	functionTable *FunctionTable
	classTable    map[string]*types.Class
}
