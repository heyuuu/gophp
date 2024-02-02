package php

import (
	"github.com/heyuuu/gophp/php/types"
)

type ExecuteData struct {
	ctx     *Context
	fn      *types.Function
	args    []types.Zval
	symbols ISymtable
	prev    *ExecuteData

	thisClass *types.Class
	thisObj   *types.Object
}

func NewExecuteData(ctx *Context, fn *types.Function, args []types.Zval) *ExecuteData {
	return &ExecuteData{
		ctx:     ctx,
		args:    args,
		symbols: NewSymtable(),
		fn:      fn,
	}
}

func (ex *ExecuteData) Ctx() *Context {
	return ex.ctx
}
func (ex *ExecuteData) Args() []types.Zval { return ex.args }
func (ex *ExecuteData) NumArgs() int {
	return len(ex.args)
}
func (ex *ExecuteData) Arg(pos int) types.Zval {
	if pos >= 0 && pos < len(ex.args) {
		return ex.args[pos]
	}
	return types.Undef
}

func (ex *ExecuteData) Symbols() ISymtable           { return ex.symbols }
func (ex *ExecuteData) SetSymbols(symbols ISymtable) { ex.symbols = symbols }
func (ex *ExecuteData) Prev() *ExecuteData           { return ex.prev }
func (ex *ExecuteData) SetPrev(prev *ExecuteData)    { ex.prev = prev }
func (ex *ExecuteData) Fn() *types.Function          { return ex.fn }
func (ex *ExecuteData) SetFn(fn *types.Function)     { ex.fn = fn }

func (ex *ExecuteData) CalleeName() string {
	//TODO implement me
	if ex.fn == nil {
		return ""
	}

	return ex.fn.Name()
}

// scope
func (ex *ExecuteData) Scope() any {
	if ex.thisObj != nil {
		return ex.thisObj
	} else if ex.thisClass != nil {
		return ex.thisClass
	} else {
		return nil
	}
}

func (ex *ExecuteData) SetScope(scope any) {
	if scope == nil {
		ex.thisClass, ex.thisObj = nil, nil
		return
	}

	switch s := scope.(type) {
	case *types.Object:
		ex.thisClass, ex.thisObj = s.Ce(), s
	case *types.Class:
		ex.thisClass, ex.thisObj = s, nil
	default:
		panic("ExecuteData.SetScope() 只支持 *types.Object、*types.ClassEntry 或 nil 参数")
	}
}
func (ex *ExecuteData) InScope() bool             { return ex.thisClass != nil }
func (ex *ExecuteData) ThisObject() *types.Object { return ex.thisObj }
func (ex *ExecuteData) ThisClass() *types.Class   { return ex.thisClass }

// 临时兼容，后续使用 ThisObject 替代
func (ex *ExecuteData) ThisObjectZval() *types.Zval {
	if ex.thisObj != nil {
		return types.NewZvalObject(ex.thisObj)
	}
	return nil
}

func (ex *ExecuteData) isStrictTypes() bool {
	return ex != nil && ex.fn != nil && ex.fn.IsStrictTypes()
}
func (ex *ExecuteData) IsCallUseStrictTypes() bool { return ex.isStrictTypes() } // ZEND_RET_USES_STRICT_TYPES
func (ex *ExecuteData) IsRetUseStrictTypes() bool  { return ex.isStrictTypes() }
func (ex *ExecuteData) IsArgUseStrictTypes() bool  { return ex.prev.isStrictTypes() } // ZEND_ARG_USES_STRICT_TYPES
