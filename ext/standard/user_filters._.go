// <<generate>>

package standard

import (
	"sik/core/streams"
	"sik/zend"
	"sik/zend/types"
)

// Source: <ext/standard/user_filters.c>

/*
   +----------------------------------------------------------------------+
   | PHP Version 7                                                        |
   +----------------------------------------------------------------------+
   | Copyright (c) The PHP Group                                          |
   +----------------------------------------------------------------------+
   | This source file is subject to version 3.01 of the PHP license,      |
   | that is bundled with this package in the file LICENSE, and is        |
   | available through the world-wide-web at the following url:           |
   | http://www.php.net/license/3_01.txt                                  |
   | If you did not receive a copy of the PHP license and are unable to   |
   | obtain it through the world-wide-web, please send a note to          |
   | license@php.net so we can mail you a copy immediately.               |
   +----------------------------------------------------------------------+
   | Authors:                                                             |
   | Wez Furlong (wez@thebrainroom.com)                                   |
   | Sara Golemon (pollita@php.net)                                       |
   +----------------------------------------------------------------------+
*/

const PHP_STREAM_BRIGADE_RES_NAME = "userfilter.bucket brigade"
const PHP_STREAM_BUCKET_RES_NAME = "userfilter.bucket"
const PHP_STREAM_FILTER_RES_NAME = "userfilter.filter"

/* to provide context for calling into the next filter from user-space */

var LeUserfilters int
var LeBucketBrigade int
var LeBucket int

/* define the __special__  base filter class */

var ArginfoPhpUserFilterFilter []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("in"),
	zend.MakeArgInfo("out"),
	zend.MakeArgInfo("consumed", ArgInfoByRef(1)),
	zend.MakeArgInfo("closing"),
}
var arginfo_php_user_filter_onCreate []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
}
var arginfo_php_user_filter_onClose []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
}
var UserFilterClassFuncs []types.ZendFunctionEntry = []types.ZendFunctionEntry{
	types.MakeZendFunctionEntryEx("filter", 0, ZifUserFilterNop, ArginfoPhpUserFilterFilter),
	types.MakeZendFunctionEntryEx("onCreate", 0, ZifUserFilterNop, arginfo_php_user_filter_onCreate),
	types.MakeZendFunctionEntryEx("onClose", 0, ZifUserFilterNop, arginfo_php_user_filter_onClose),
}
var UserFilterClassEntry types.ClassEntry
var UserfilterOps streams.PhpStreamFilterOps = streams.MakePhpStreamFilterOps(UserfilterFilter, UserfilterDtor, "user-filter")
var UserFilterFactory streams.PhpStreamFilterFactory = streams.MakePhpStreamFilterFactory(UserFilterFactoryCreate)

/* {{{ proto object stream_bucket_make_writeable(resource brigade)
   Return a bucket object from the brigade for operating on */
