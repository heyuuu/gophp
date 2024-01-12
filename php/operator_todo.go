package php

import (
	"github.com/heyuuu/gophp/php/perr"
	"github.com/heyuuu/gophp/php/types"
)

// Operator
type Operator struct {
	OperatorHandler
}

func New(handler OperatorHandler) *Operator {
	return &Operator{OperatorHandler: handler}
}

// OperatorHandler
type OperatorHandler interface {
	Precision() int
	Error(level perr.ErrorType, message string)
	ThrowError(exceptionCe *types.Class, message string)
	ThrowException(exceptionCe *types.Class, message string)
	NewObject(properties *types.Array) *types.Object
	ObjectGetArray(obj *types.Object) *types.Array
	HasException() bool
}

func NewOperator(ctx *Context) *Operator {
	handler := ctxOperatorHandler{ctx: ctx}
	return New(handler)
}

type ctxOperatorHandler struct {
	ctx *Context
}

var _ OperatorHandler = ctxOperatorHandler{}

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
