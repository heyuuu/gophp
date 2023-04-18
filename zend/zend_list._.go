package zend

import (
	types2 "github.com/heyuuu/gophp/php/types"
)

type RsrcDtorFuncT func(res *types2.ZendResource)

var ListDestructors types2.Array
