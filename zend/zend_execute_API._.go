package zend

import "sik/zend/types"

/* true globals */

var EmptyFcallInfo types.ZendFcallInfo = types.MakeZendFcallInfo(0, types.Zval{}, nil, nil, nil, 0, 0)
var EmptyFcallInfoCache types.ZendFcallInfoCache = types.MakeZendFcallInfoCache(nil, nil, nil, nil)

/* This one doesn't exists on QNX */

const SIGPROF = 27
