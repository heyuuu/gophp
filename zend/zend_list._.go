package zend

import "github.com/heyuuu/gophp/zend/types"

type RsrcDtorFuncT func(res *types.ZendResource)

var LeIndexPtr int

/* true global */

var ListDestructors types.Array
