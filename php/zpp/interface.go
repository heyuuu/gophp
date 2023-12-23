package zpp

import "github.com/heyuuu/gophp/php/types"

type IExecuteData interface {
}

type IParser interface {
	HasError() bool
	ParseZval() *types.Zval
}
