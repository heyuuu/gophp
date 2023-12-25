package zpp

import (
	"github.com/heyuuu/gophp/php/types"
)

type ExecuteData interface {
	CalleeName() string
	NumArgs() int
	Arg(pos int) *types.Zval
	IsArgUseStrictTypes() bool
}

type IParser interface {
	HasError() bool

	// exactly type
	ParseBool() bool
	ParseBoolNullable() *bool
	ParseLong() int
	ParseLongNullable() *int
	ParseStrictLong() int
	ParseStrictLongNullable() *int
	ParseDouble() float64
	ParseDoubleNullable() *float64
	ParseString() string
	ParseStringNullable() *string
	ParsePath() string
	ParsePathNullable() *string
	ParseArray() *types.Array
	ParseArrayNullable() *types.Array
	ParseArrayOrObjectHt() *types.Array
	ParseArrayOrObject() *types.Zval
	ParseClass(baseCe *types.Class) *types.Class
	ParseObject() *types.Object
	ParseObjectNullable() *types.Object
	//ParseResource() *types.Resource

	// zval
	ParseZval() *types.Zval
	ParseZvalNullable() *types.Zval
	ParseZvalDeref() *types.Zval
	ParseVariadic() []*types.Zval

	// ref type
	//ParseRefZval() *types.Zval
}
