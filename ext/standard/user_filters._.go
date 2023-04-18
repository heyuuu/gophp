package standard

import (
	"github.com/heyuuu/gophp/core/streams"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
)

const PHP_STREAM_BRIGADE_RES_NAME = "userfilter.bucket brigade"
const PHP_STREAM_BUCKET_RES_NAME = "userfilter.bucket"
const PHP_STREAM_FILTER_RES_NAME = "userfilter.filter"

var LeUserfilters int
var LeBucketBrigade int
var LeBucket int

var UserFilterClassFuncs []types.FunctionEntry = []types.FunctionEntry{
	types.MakeZendFunctionEntryEx("filter", 0, ZifUserFilterNop, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgName("in"),
		zend.MakeArgName("out"),
		zend.MakeArgByRef("consumed"),
		zend.MakeArgName("closing"),
	}),
	types.MakeZendFunctionEntryEx("onCreate", 0, ZifUserFilterNop, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("onClose", 0, ZifUserFilterNop, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
}
var UserFilterClassEntry types.ClassEntry
var UserfilterOps streams.PhpStreamFilterOps = streams.MakePhpStreamFilterOps(UserfilterFilter, UserfilterDtor, "user-filter")
var UserFilterFactory streams.PhpStreamFilterFactory = streams.MakePhpStreamFilterFactory(UserFilterFactoryCreate)
