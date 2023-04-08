package zend

import "github.com/heyuuu/gophp/zend/types"

type RsrcDtorFuncT func(res *types.ZendResource)

var ListDestructors types.Array
