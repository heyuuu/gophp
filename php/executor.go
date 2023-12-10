package php

import (
	"fmt"
	"github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/operators"
	"github.com/heyuuu/gophp/php/types"
)

// ExecutorError
type ExecutorError string

func (e ExecutorError) Error() string { return string(e) }

// executor
type Executor struct {
	ctx         *Context
	executeData *ExecuteData
	operator    *operators.Operator
}

func NewExecutor(ctx *Context) *Executor {
	return &Executor{
		ctx:      ctx,
		operator: ctx.Operator(),
	}
}

func (e *Executor) Execute(fn *types.Function) (retval Val, ret error) {
	//defer func() {
	//	if e := recover(); e != nil {
	//		if err_, ok := e.(ExecutorError); ok {
	//			retval, ret = nil, err_
	//		} else {
	//			panic(fmt.Errorf("%w", e)) // re-panic
	//		}
	//	}
	//}()

	return e.function(fn, nil), nil
}

func (e *Executor) function(fn *types.Function, args []Val) Val {
	lang.Assert(fn != nil)
	if fn.IsInternalFunction() {
		// todo
		var retval Val
		fn.Handler()(nil, retval)
		return retval
	} else {
		return e.userFunction(fn, args)
	}
}

func (e *Executor) initStringCall(name string) *types.Function {
	// todo ZendInitDynamicCallString
	fn := e.ctx.EG().FindFunction(name)
	if fn == nil {
		ThrowError(nil, fmt.Sprintf("Call to undefined function %s()", name))
		return nil
	}
	return fn
}
