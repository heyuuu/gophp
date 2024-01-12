package php

import (
	"errors"
)

var vmError = errors.New("vm error")

func vmBinaryOp(ctx *Context, op1, op2 Val, handler func(Val, Val) (Val, bool)) (Val, error) {
	ret, ok := handler(op1, op2)
	if !ok {
		return nil, vmError
	}

	return ret, nil
}

func vmEcho(ctx *Context, zv Val) {
	str := ZvalGetStrVal(ctx, zv)
	if len(str) > 0 {
		ctx.WriteString(str)
	}
}
