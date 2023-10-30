package php

import (
	"errors"
	"github.com/heyuuu/gophp/php/operators"
)

var vmError = errors.New("vm error")

func vmBinaryOp(ctx *Context, op1, op2 Val, handler func(Val, Val) (Val, bool)) (Val, error) {
	ret, ok := handler(op1, op2)
	if !ok {
		return nil, vmError
	}

	return ret, nil
}

func vmAdd(ctx *Context, op1, op2 Val) (Val, error) {
	ret, ok := operators.Add(op1, op2)
	if !ok {
		return nil, vmError
	}

	return ret, nil
}

func vmEcho(ctx *Context, zv Val) {
	str := ZvalGetString(zv)
	if len(str) > 0 {
		ctx.WriteString(str)
	}
}
