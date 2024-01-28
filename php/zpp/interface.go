package zpp

import (
	"github.com/heyuuu/gophp/php/types"
)

type IParser interface {
	HasError() bool
	CheckNumArgs()
	StartOptional()

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
	ParseArrayOrObjectZval() types.Zval
	//ParseClass(baseCe *types.Class) *types.Class
	ParseObject() *types.Object
	ParseObjectNullable() *types.Object
	ParseResource() types.Zval
	ParseResourceNullable() *types.Zval
	ParseCallable() *types.UserCallable

	// zval
	ParseZval() types.Zval
	ParseZvalNullable() *types.Zval
	ParseVariadic(postVarargs uint) []types.Zval

	// ref type
	ParseRefZval() types.RefZval
	ParseRefZvalNullable() types.RefZval
	ParseRefArrayOrObject() types.RefZval
	ParseRefArray() *types.Array
	ParseRefArrayNullable() *types.Array
	ParseRefVariadic(postVarargs uint) []types.RefZval
}
