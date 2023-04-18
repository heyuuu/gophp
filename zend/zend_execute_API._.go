package zend

import (
	types2 "github.com/heyuuu/gophp/php/types"
)

/* true globals */

var EmptyFcallInfo types2.ZendFcallInfo = types2.MakeZendFcallInfo(0, types2.Zval{}, nil, nil, nil, 0, 0)
var EmptyFcallInfoCache types2.ZendFcallInfoCache = types2.MakeZendFcallInfoCache(nil, nil, nil, nil)

/* This one doesn't exists on QNX */

const SIGPROF = 27
