// <<generate>>

package standard

import (
	"sik/core/streams"
	"sik/zend"
	"sik/zend/types"
)

const PHP_STREAM_BRIGADE_RES_NAME = "userfilter.bucket brigade"
const PHP_STREAM_BUCKET_RES_NAME = "userfilter.bucket"
const PHP_STREAM_FILTER_RES_NAME = "userfilter.filter"

var LeUserfilters int
var LeBucketBrigade int
var LeBucket int

var UserFilterClassFuncs []types.ZendFunctionEntry = []types.ZendFunctionEntry{
	types.MakeZendFunctionEntryEx("filter", 0, ZifUserFilterNop, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgName("in"),
		zend.MakeArgName("out"),
		zend.MakeArgInfo("consumed", zend.ArgInfoByRef(1)),
		zend.MakeArgName("closing"),
	}),
	types.MakeZendFunctionEntryEx("onCreate", 0, ZifUserFilterNop, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("onClose", 0, ZifUserFilterNop, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
}
var UserFilterClassEntry types.ClassEntry
var UserfilterOps streams.PhpStreamFilterOps = streams.MakePhpStreamFilterOps(UserfilterFilter, UserfilterDtor, "user-filter")
var UserFilterFactory streams.PhpStreamFilterFactory = streams.MakePhpStreamFilterFactory(UserFilterFactoryCreate)
