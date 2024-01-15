package php

import (
	"fmt"
	"github.com/heyuuu/gophp/php/perr"
	"github.com/heyuuu/gophp/php/types"
)

// ExecutorError
type ExecutorError string

func (e ExecutorError) Error() string { return string(e) }

// executor
type Executor struct {
	ctx         *Context
	executeData *ExecuteData
}

func NewExecutor(ctx *Context) *Executor {
	return &Executor{ctx: ctx}
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
	perr.Assert(fn != nil)

	// push && pop executeData
	e.executeData = NewExecuteData(e.ctx, args, e.executeData)
	defer func() {
		e.executeData = e.executeData.prev
	}()

	if fn.IsInternalFunction() {
		var retval types.Zval
		if handler, ok := fn.Handler().(ZifHandler); ok {
			handler(e.executeData, &retval)
		} else {
			perr.Panic(fmt.Sprintf("不支持的内部函数 handler 类型: %T", fn.Handler()))
		}
		return retval
	} else {
		return e.userFunction(fn, args)
	}
}

func (e *Executor) initStringCall(name string) *types.Function {
	// todo ZendInitDynamicCallString
	fn := e.ctx.EG().FindFunction(name)
	if fn == nil {
		ThrowError(e.ctx, nil, fmt.Sprintf("Call to undefined function %s()", name))
		return nil
	}
	return fn
}

func (e *Executor) initNewExecuteData(args []Val) *ExecuteData {
	e.executeData = NewExecuteData(e.ctx, args, e.executeData)
	return e.executeData
}

func (e *Executor) zvalToArrayKey(dim Val) types.ArrayKey {
	if dim.IsLong() {
		return types.IdxKey(dim.Long())
	} else if dim.IsString() {
		return types.StrKey(dim.String())
	} else {
		panic(perr.Internalf("暂不支持的内部数组 dim 类型: %s", types.ZvalGetType(dim)))
	}
}
