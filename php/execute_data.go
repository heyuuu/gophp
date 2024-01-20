package php

import (
	"github.com/heyuuu/gophp/php/perr"
	"github.com/heyuuu/gophp/php/types"
)

type ExecuteData struct {
	ctx     *Context
	args    []types.Zval
	symbols ISymtable
	prev    *ExecuteData
	fn      *types.Function
}

func NewExecuteData(ctx *Context, args []types.Zval, prev *ExecuteData) *ExecuteData {
	return &ExecuteData{
		ctx:     ctx,
		args:    args,
		symbols: NewSymtable(),
		prev:    prev,
	}
}

func (ex *ExecuteData) Ctx() *Context {
	return ex.ctx
}

func (ex *ExecuteData) CalleeName() string {
	//TODO implement me
	panic(perr.Todo())
}

func (ex *ExecuteData) NumArgs() int {
	return len(ex.args)
}

func (ex *ExecuteData) Arg(pos int) types.Zval {
	if pos >= 0 && pos < len(ex.args) {
		return ex.args[pos]
	}
	return types.Undef
}

func (ex *ExecuteData) isStrictTypes() bool {
	return ex != nil && ex.fn != nil && ex.fn.IsStrictTypes()
}
func (ex *ExecuteData) IsCallUseStrictTypes() bool { return ex.isStrictTypes() } // ZEND_RET_USES_STRICT_TYPES
func (ex *ExecuteData) IsRetUseStrictTypes() bool  { return ex.isStrictTypes() }
func (ex *ExecuteData) IsArgUseStrictTypes() bool  { return ex.prev.isStrictTypes() } // ZEND_ARG_USES_STRICT_TYPES
