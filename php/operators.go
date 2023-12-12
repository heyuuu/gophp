package php

import (
	"github.com/heyuuu/gophp/php/operators"
	"github.com/heyuuu/gophp/php/perr"
	"github.com/heyuuu/gophp/php/types"
)

func NewOperator(ctx *Context) *operators.Operator {
	handler := ctxOperatorHandler{ctx: ctx}
	return operators.New(handler)
}

type ctxOperatorHandler struct {
	ctx *Context
}

var _ operators.OperatorHandler = ctxOperatorHandler{}

func (c ctxOperatorHandler) Precision() int {
	//TODO implement me
	panic("implement me")
}

func (c ctxOperatorHandler) Error(level perr.ErrorType, message string) {
	//TODO implement me
	panic("implement me")
}

func (c ctxOperatorHandler) ThrowError(exceptionCe *types.Class, message string) {
	//TODO implement me
	panic("implement me")
}

func (c ctxOperatorHandler) ThrowException(exceptionCe *types.Class, message string) {
	//TODO implement me
	panic("implement me")
}

func (c ctxOperatorHandler) NewObject(properties *types.Array) *types.Object {
	//TODO implement me
	panic("implement me")
}

func (c ctxOperatorHandler) ObjectGetArray(obj *types.Object) *types.Array {
	//TODO implement me
	panic("implement me")
}

func (c ctxOperatorHandler) HasException() bool {
	//TODO implement me
	panic("implement me")
}
