package zpp

import (
	"github.com/heyuuu/gophp/php/perr"
	"github.com/heyuuu/gophp/php/types"
)

var _ IParser = (*FastParser)(nil)

type FastParser struct {
	ex         ExecuteData
	minNumArgs int
	maxNumArgs int
	flags      int
	err        error
	argIndex   int
}

func NewFastParser(ex ExecuteData, minNumArgs int, maxNumArgs int, flags int) IParser {
	return &FastParser{
		ex:         ex,
		minNumArgs: minNumArgs,
		maxNumArgs: maxNumArgs,
		flags:      flags,
	}
}

func (p *FastParser) HasError() bool {
	return p.err != nil
}

func (p *FastParser) nextArg() *types.Zval {
	if p.err != nil || p.argIndex >= p.ex.NumArgs() {
		return nil
	}

	arg := p.ex.Arg(p.argIndex)
	p.argIndex++
	return arg
}

func (p *FastParser) ParseBool() bool {
	//TODO implement me
	panic(perr.Todo())
}

func (p *FastParser) ParseBoolNullable() *bool {
	//TODO implement me
	panic(perr.Todo())
}

func (p *FastParser) ParseLong() int {
	//TODO implement me
	panic(perr.Todo())
}

func (p *FastParser) ParseLongNullable() *int {
	//TODO implement me
	panic(perr.Todo())
}

func (p *FastParser) ParseStrictLong() int {
	//TODO implement me
	panic(perr.Todo())
}

func (p *FastParser) ParseStrictLongNullable() *int {
	//TODO implement me
	panic(perr.Todo())
}

func (p *FastParser) ParseDouble() float64 {
	//TODO implement me
	panic(perr.Todo())
}

func (p *FastParser) ParseDoubleNullable() *float64 {
	//TODO implement me
	panic(perr.Todo())
}

func (p *FastParser) ParseString() string {
	//TODO implement me
	panic(perr.Todo())
}

func (p *FastParser) ParseStringNullable() *string {
	//TODO implement me
	panic(perr.Todo())
}

func (p *FastParser) ParsePath() string {
	//TODO implement me
	panic(perr.Todo())
}

func (p *FastParser) ParsePathNullable() *string {
	//TODO implement me
	panic(perr.Todo())
}

func (p *FastParser) ParseArray() *types.Array {
	//TODO implement me
	panic(perr.Todo())
}

func (p *FastParser) ParseArrayNullable() *types.Array {
	//TODO implement me
	panic(perr.Todo())
}

func (p *FastParser) ParseArrayOrObjectHt() *types.Array {
	//TODO implement me
	panic(perr.Todo())
}

func (p *FastParser) ParseArrayOrObject() *types.Zval {
	//TODO implement me
	panic(perr.Todo())
}

func (p *FastParser) ParseClass(baseCe *types.Class) *types.Class {
	//TODO implement me
	panic(perr.Todo())
}

func (p *FastParser) ParseObject() *types.Object {
	//TODO implement me
	panic(perr.Todo())
}

func (p *FastParser) ParseObjectNullable() *types.Object {

	//TODO implement me
	panic(perr.Todo())
}

func (p *FastParser) ParseZval() *types.Zval {
	arg := p.nextArg()
	if arg == nil {
		return nil
	}

	return arg
}

func (p *FastParser) ParseZvalNullable() *types.Zval {
	arg := p.nextArg()
	if arg == nil {
		return nil
	}

	if arg.IsNull() {
		return nil
	}
	return arg
}

func (p *FastParser) ParseZvalDeref() *types.Zval {
	//TODO implement me
	panic(perr.Todo())
}

func (p *FastParser) ParseVariadic() []*types.Zval {
	var args []*types.Zval
	for {
		arg := p.nextArg()
		if arg == nil {
			break
		}
		args = append(args, arg)
	}
	return args
}
