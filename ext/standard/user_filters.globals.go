// <<generate>>

package standard

import (
	b "sik/builtin"
	"sik/core/streams"
	"sik/zend"
)

const PHP_STREAM_BRIGADE_RES_NAME = "userfilter.bucket brigade"
const PHP_STREAM_BUCKET_RES_NAME = "userfilter.bucket"
const PHP_STREAM_FILTER_RES_NAME = "userfilter.filter"

var LeUserfilters int
var LeBucketBrigade int
var LeBucket int
var ArginfoPhpUserFilterFilter []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"in", 0, 0, 0},
	{"out", 0, 0, 0},
	{"consumed", 0, 1, 0},
	{"closing", 0, 0, 0},
}
var arginfo_php_user_filter_onCreate []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
}
var arginfo_php_user_filter_onClose []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
}
var UserFilterClassFuncs []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"filter",
		ZifUserFilterNop,
		ArginfoPhpUserFilterFilter,
		uint32_t(b.SizeOf("arginfo_php_user_filter_filter")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"onCreate",
		ZifUserFilterNop,
		arginfo_php_user_filter_onCreate,
		uint32_t(b.SizeOf("arginfo_php_user_filter_onCreate")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"onClose",
		ZifUserFilterNop,
		arginfo_php_user_filter_onClose,
		uint32_t(b.SizeOf("arginfo_php_user_filter_onClose")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{nil, nil, nil, 0, 0},
}
var UserFilterClassEntry zend.ZendClassEntry
var UserfilterOps streams.PhpStreamFilterOps = streams.PhpStreamFilterOps{UserfilterFilter, UserfilterDtor, "user-filter"}
var UserFilterFactory streams.PhpStreamFilterFactory = streams.PhpStreamFilterFactory{UserFilterFactoryCreate}
