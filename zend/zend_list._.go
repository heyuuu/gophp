package zend

import (
	"github.com/heyuuu/gophp/php/types"
)

type RsrcDtorFuncT func(res *types.ZendResource)

var ListDestructors types.Array
