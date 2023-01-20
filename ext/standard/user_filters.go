// <<generate>>

package standard

import (
	"sik/core"
	"sik/core/streams"
	g "sik/runtime/grammar"
	"sik/zend"
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

// # include "php.h"

// # include "php_globals.h"

// # include "ext/standard/basic_functions.h"

// # include "ext/standard/file.h"

// #define PHP_STREAM_BRIGADE_RES_NAME       "userfilter.bucket brigade"

// #define PHP_STREAM_BUCKET_RES_NAME       "userfilter.bucket"

// #define PHP_STREAM_FILTER_RES_NAME       "userfilter.filter"

// @type PhpUserFilterData struct

/* to provide context for calling into the next filter from user-space */

var LeUserfilters int
var LeBucketBrigade int
var LeBucket int

/* define the __special__  base filter class */

func ZifUserFilterNop(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {}

var ArginfoPhpUserFilterFilter []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"in", 0, 0, 0}, {"out", 0, 0, 0}, {"consumed", 0, 1, 0}, {"closing", 0, 0, 0}}
var arginfo_php_user_filter_onCreate []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}}
var arginfo_php_user_filter_onClose []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}}
var UserFilterClassFuncs []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"filter",
		ZifUserFilterNop,
		ArginfoPhpUserFilterFilter,
		uint32(g.SizeOf("arginfo_php_user_filter_filter")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"onCreate",
		ZifUserFilterNop,
		arginfo_php_user_filter_onCreate,
		uint32(g.SizeOf("arginfo_php_user_filter_onCreate")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"onClose",
		ZifUserFilterNop,
		arginfo_php_user_filter_onClose,
		uint32(g.SizeOf("arginfo_php_user_filter_onClose")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{nil, nil, nil, 0, 0},
}
var UserFilterClassEntry zend.ZendClassEntry

func PhpBucketDtor(res *zend.ZendResource) {
	var bucket *streams.PhpStreamBucket = (*streams.PhpStreamBucket)(res.ptr)
	if bucket != nil {
		streams.PhpStreamBucketDelref(bucket)
		bucket = nil
	}
}
func ZmStartupUserFilters(type_ int, module_number int) int {
	var php_user_filter *zend.ZendClassEntry

	/* init the filter class ancestor */

	memset(&UserFilterClassEntry, 0, g.SizeOf("zend_class_entry"))
	UserFilterClassEntry.name = zend.ZendStringInitInterned("php_user_filter", g.SizeOf("\"php_user_filter\"")-1, 1)
	UserFilterClassEntry.info.internal.builtin_functions = UserFilterClassFuncs
	if g.Assign(&php_user_filter, zend.ZendRegisterInternalClass(&UserFilterClassEntry)) == nil {
		return zend.FAILURE
	}
	zend.ZendDeclarePropertyString(php_user_filter, "filtername", g.SizeOf("\"filtername\"")-1, "", 1<<0)
	zend.ZendDeclarePropertyString(php_user_filter, "params", g.SizeOf("\"params\"")-1, "", 1<<0)

	/* init the filter resource; it has no dtor, as streams will always clean it up
	 * at the correct time */

	LeUserfilters = zend.ZendRegisterListDestructorsEx(nil, nil, "userfilter.filter", 0)
	if LeUserfilters == zend.FAILURE {
		return zend.FAILURE
	}

	/* Filters will dispose of their brigades */

	LeBucketBrigade = zend.ZendRegisterListDestructorsEx(nil, nil, "userfilter.bucket brigade", module_number)

	/* Brigades will dispose of their buckets */

	LeBucket = zend.ZendRegisterListDestructorsEx(PhpBucketDtor, nil, "userfilter.bucket", module_number)
	if LeBucketBrigade == zend.FAILURE {
		return zend.FAILURE
	}
	zend.ZendRegisterLongConstant("PSFS_PASS_ON", g.SizeOf("\"PSFS_PASS_ON\"")-1, streams.PSFS_PASS_ON, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("PSFS_FEED_ME", g.SizeOf("\"PSFS_FEED_ME\"")-1, streams.PSFS_FEED_ME, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("PSFS_ERR_FATAL", g.SizeOf("\"PSFS_ERR_FATAL\"")-1, streams.PSFS_ERR_FATAL, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("PSFS_FLAG_NORMAL", g.SizeOf("\"PSFS_FLAG_NORMAL\"")-1, 0, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("PSFS_FLAG_FLUSH_INC", g.SizeOf("\"PSFS_FLAG_FLUSH_INC\"")-1, 1, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("PSFS_FLAG_FLUSH_CLOSE", g.SizeOf("\"PSFS_FLAG_FLUSH_CLOSE\"")-1, 2, 1<<0|1<<1, module_number)
	return zend.SUCCESS
}
func ZmDeactivateUserFilters(type_ int, module_number int) int {
	if BasicGlobals.GetUserFilterMap() != nil {
		zend.ZendHashDestroy(BasicGlobals.GetUserFilterMap())
		zend._efree(BasicGlobals.GetUserFilterMap())
		BasicGlobals.SetUserFilterMap(nil)
	}
	return zend.SUCCESS
}
func UserfilterDtor(thisfilter *core.PhpStreamFilter) {
	var obj *zend.Zval = &thisfilter.abstract
	var func_name zend.Zval
	var retval zend.Zval
	if obj == nil {

		/* If there's no object associated then there's nothing to dispose of */

		return

		/* If there's no object associated then there's nothing to dispose of */

	}
	var __z *zend.Zval = &func_name
	var __s *zend.ZendString = zend.ZendStringInit("onclose", g.SizeOf("\"onclose\"")-1, 0)
	__z.value.str = __s
	__z.u1.type_info = 6 | 1<<0<<8
	zend._callUserFunctionEx(obj, &func_name, &retval, 0, nil, 1)
	zend.ZvalPtrDtor(&retval)
	zend.ZvalPtrDtor(&func_name)

	/* kill the object */

	zend.ZvalPtrDtor(obj)

	/* kill the object */
}
func UserfilterFilter(stream *core.PhpStream, thisfilter *core.PhpStreamFilter, buckets_in *streams.PhpStreamBucketBrigade, buckets_out *streams.PhpStreamBucketBrigade, bytes_consumed *int, flags int) streams.PhpStreamFilterStatusT {
	var ret int = streams.PSFS_ERR_FATAL
	var obj *zend.Zval = &thisfilter.abstract
	var func_name zend.Zval
	var retval zend.Zval
	var args []zend.Zval
	var zpropname zend.Zval
	var call_result int

	/* the userfilter object probably doesn't exist anymore */

	if zend.CG.unclean_shutdown != 0 {
		return ret
	}

	/* Make sure the stream is not closed while the filter callback executes. */

	var orig_no_fclose uint32 = stream.flags & 0x80
	stream.flags |= 0x80
	if zend.ZendHashStrExistsInd(obj.value.obj.handlers.get_properties(&(*obj)), "stream", g.SizeOf("\"stream\"")-1) == 0 {
		var tmp zend.Zval

		/* Give the userfilter class a hook back to the stream */

		var __z *zend.Zval = &tmp
		__z.value.res = stream.res
		__z.u1.type_info = 9 | 1<<0<<8
		stream.__exposed = 1
		zend.ZvalAddrefP(&tmp)
		zend.AddPropertyZvalEx(obj, "stream", strlen("stream"), &tmp)

		/* add_property_zval increments the refcount which is unwanted here */

		zend.ZvalPtrDtor(&tmp)

		/* add_property_zval increments the refcount which is unwanted here */

	}
	var __z *zend.Zval = &func_name
	var __s *zend.ZendString = zend.ZendStringInit("filter", g.SizeOf("\"filter\"")-1, 0)
	__z.value.str = __s
	__z.u1.type_info = 6 | 1<<0<<8

	/* Setup calling arguments */

	var __z *zend.Zval = &args[0]
	__z.value.res = zend.ZendRegisterResource(buckets_in, LeBucketBrigade)
	__z.u1.type_info = 9 | 1<<0<<8
	var __z *zend.Zval = &args[1]
	__z.value.res = zend.ZendRegisterResource(buckets_out, LeBucketBrigade)
	__z.u1.type_info = 9 | 1<<0<<8
	if bytes_consumed != nil {
		var __z *zend.Zval = &args[2]
		__z.value.lval = *bytes_consumed
		__z.u1.type_info = 4
	} else {
		&args[2].u1.type_info = 1
	}
	if (flags & 2) != 0 {
		&args[3].u1.type_info = 3
	} else {
		&args[3].u1.type_info = 2
	}
	call_result = zend._callUserFunctionEx(obj, &func_name, &retval, 4, args, 0)
	zend.ZvalPtrDtor(&func_name)
	if call_result == zend.SUCCESS && retval.u1.v.type_ != 0 {
		zend.ConvertToLong(&retval)
		ret = int(retval.value.lval)
	} else if call_result == zend.FAILURE {
		core.PhpErrorDocref(nil, 1<<1, "failed to call filter function")
	}
	if bytes_consumed != nil {
		*bytes_consumed = zend.ZvalGetLong(&args[2])
	}
	if buckets_in.head != nil {
		var bucket *streams.PhpStreamBucket = buckets_in.head
		core.PhpErrorDocref(nil, 1<<1, "Unprocessed filter buckets remaining on input brigade")
		for g.Assign(&bucket, buckets_in.head) {

			/* Remove unconsumed buckets from the brigade */

			streams.PhpStreamBucketUnlink(bucket)
			streams.PhpStreamBucketDelref(bucket)
		}
	}
	if ret != streams.PSFS_PASS_ON {
		var bucket *streams.PhpStreamBucket = buckets_out.head
		for bucket != nil {
			streams.PhpStreamBucketUnlink(bucket)
			streams.PhpStreamBucketDelref(bucket)
			bucket = buckets_out.head
		}
	}

	/* filter resources are cleaned up by the stream destructor,
	 * keeping a reference to the stream resource here would prevent it
	 * from being destroyed properly */

	var __z *zend.Zval = &zpropname
	var __s *zend.ZendString = zend.ZendStringInit("stream", g.SizeOf("\"stream\"")-1, 0)
	__z.value.str = __s
	__z.u1.type_info = 6 | 1<<0<<8
	obj.value.obj.handlers.unset_property(obj, &zpropname, nil)
	zend.ZvalPtrDtor(&zpropname)
	zend.ZvalPtrDtor(&args[3])
	zend.ZvalPtrDtor(&args[2])
	zend.ZvalPtrDtor(&args[1])
	zend.ZvalPtrDtor(&args[0])
	stream.flags &= ^0x80
	stream.flags |= orig_no_fclose
	return ret
}

var UserfilterOps streams.PhpStreamFilterOps = streams.PhpStreamFilterOps{UserfilterFilter, UserfilterDtor, "user-filter"}

func UserFilterFactoryCreate(filtername *byte, filterparams *zend.Zval, persistent uint8) *core.PhpStreamFilter {
	var fdat *PhpUserFilterData = nil
	var filter *core.PhpStreamFilter
	var obj zend.Zval
	var zfilter zend.Zval
	var func_name zend.Zval
	var retval zend.Zval
	var len_ int

	/* some sanity checks */

	if persistent != 0 {
		core.PhpErrorDocref(nil, 1<<1, "cannot use a user-space filter with a persistent stream")
		return nil
	}
	len_ = strlen(filtername)

	/* determine the classname/class entry */

	if nil == g.Assign(&fdat, zend.ZendHashStrFindPtr(BasicGlobals.GetUserFilterMap(), (*byte)(filtername), len_)) {
		var period *byte

		/* Userspace Filters using ambiguous wildcards could cause problems.
		   i.e.: myfilter.foo.bar will always call into myfilter.foo.*
		         never seeing myfilter.*
		   TODO: Allow failed userfilter creations to continue
		         scanning through the list */

		if g.Assign(&period, strrchr(filtername, '.')) {
			var wildcard *byte = zend._safeEmalloc(len_, 1, 3)

			/* Search for wildcard matches instead */

			memcpy(wildcard, filtername, len_+1)
			period = wildcard + (period - filtername)
			for period != nil {
				assert(period[0] == '.')
				period[1] = '*'
				period[2] = '0'
				if nil != g.Assign(&fdat, zend.ZendHashStrFindPtr(BasicGlobals.GetUserFilterMap(), wildcard, strlen(wildcard))) {
					period = nil
				} else {
					*period = '0'
					period = strrchr(wildcard, '.')
				}
			}
			zend._efree(wildcard)
		}
		if fdat == nil {
			core.PhpErrorDocref(nil, 1<<1, "Err, filter \"%s\" is not in the user-filter map, but somehow the user-filter-factory was invoked for it!?", filtername)
			return nil
		}
	}

	/* bind the classname to the actual class */

	if fdat.GetCe() == nil {
		if nil == g.Assign(&(fdat.GetCe()), zend.ZendLookupClass(fdat.GetClassname())) {
			core.PhpErrorDocref(nil, 1<<1, "user-filter \"%s\" requires class \"%s\", but that class is not defined", filtername, fdat.GetClassname().val)
			return nil
		}
	}

	/* create the object */

	if zend.ObjectInitEx(&obj, fdat.GetCe()) == zend.FAILURE {
		return nil
	}
	filter = streams._phpStreamFilterAlloc(&UserfilterOps, nil, 0)
	if filter == nil {
		zend.ZvalPtrDtor(&obj)
		return nil
	}

	/* filtername */

	zend.AddPropertyStringEx(&obj, "filtername", strlen("filtername"), (*byte)(filtername))

	/* and the parameters, if any */

	if filterparams != nil {
		zend.AddPropertyZvalEx(&obj, "params", strlen("params"), filterparams)
	} else {
		zend.AddPropertyNullEx(&obj, "params", strlen("params"))
	}

	/* invoke the constructor */

	var __z *zend.Zval = &func_name
	var __s *zend.ZendString = zend.ZendStringInit("oncreate", g.SizeOf("\"oncreate\"")-1, 0)
	__z.value.str = __s
	__z.u1.type_info = 6 | 1<<0<<8
	zend._callUserFunctionEx(&obj, &func_name, &retval, 0, nil, 1)
	if retval.u1.v.type_ != 0 {
		if retval.u1.v.type_ == 2 {

			/* User reported filter creation error "return false;" */

			zend.ZvalPtrDtor(&retval)

			/* Kill the filter (safely) */

			&filter.abstract.u1.type_info = 0
			streams.PhpStreamFilterFree(filter)

			/* Kill the object */

			zend.ZvalPtrDtor(&obj)

			/* Report failure to filter_alloc */

			return nil

			/* Report failure to filter_alloc */

		}
		zend.ZvalPtrDtor(&retval)
	}
	zend.ZvalPtrDtor(&func_name)

	/* set the filter property, this will be used during cleanup */

	var __z *zend.Zval = &zfilter
	__z.value.res = zend.ZendRegisterResource(filter, LeUserfilters)
	__z.u1.type_info = 9 | 1<<0<<8
	var __z *zend.Zval = &filter.abstract
	__z.value.obj = obj.value.obj
	__z.u1.type_info = 8 | 1<<0<<8 | 1<<1<<8
	zend.AddPropertyZvalEx(&obj, "filter", strlen("filter"), &zfilter)

	/* add_property_zval increments the refcount which is unwanted here */

	zend.ZvalPtrDtor(&zfilter)
	return filter
}

var UserFilterFactory streams.PhpStreamFilterFactory = streams.PhpStreamFilterFactory{UserFilterFactoryCreate}

func FilterItemDtor(zv *zend.Zval) {
	var fdat *PhpUserFilterData = zv.value.ptr
	zend.ZendStringReleaseEx(fdat.GetClassname(), 0)
	zend._efree(fdat)
}

/* {{{ proto object stream_bucket_make_writeable(resource brigade)
   Return a bucket object from the brigade for operating on */

func ZifStreamBucketMakeWriteable(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var zbrigade *zend.Zval
	var zbucket zend.Zval
	var brigade *streams.PhpStreamBucketBrigade
	var bucket *streams.PhpStreamBucket
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgResource(_arg, &zbrigade, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_RESOURCE
				_error_code = 4
				break
			}
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return_value.u1.type_info = 2
			return
		}
		break
	}
	if g.Assign(&brigade, (*streams.PhpStreamBucketBrigade)(zend.ZendFetchResource(zbrigade.value.res, "userfilter.bucket brigade", LeBucketBrigade))) == nil {
		return_value.u1.type_info = 2
		return
	}
	return_value.u1.type_info = 1
	if brigade.head != nil && g.Assign(&bucket, streams.PhpStreamBucketMakeWriteable(brigade.head)) {
		var __z *zend.Zval = &zbucket
		__z.value.res = zend.ZendRegisterResource(bucket, LeBucket)
		__z.u1.type_info = 9 | 1<<0<<8
		zend.ObjectInit(return_value)
		zend.AddPropertyZvalEx(return_value, "bucket", strlen("bucket"), &zbucket)

		/* add_property_zval increments the refcount which is unwanted here */

		zend.ZvalPtrDtor(&zbucket)
		zend.AddPropertyStringlEx(return_value, "data", strlen("data"), bucket.buf, bucket.buflen)
		zend.AddPropertyLongEx(return_value, "datalen", strlen("datalen"), bucket.buflen)
	}
}

/* }}} */

func PhpStreamBucketAttach(append int, execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var zbrigade *zend.Zval
	var zobject *zend.Zval
	var pzbucket *zend.Zval
	var pzdata *zend.Zval
	var brigade *streams.PhpStreamBucketBrigade
	var bucket *streams.PhpStreamBucket
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 2
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgResource(_arg, &zbrigade, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_RESOURCE
				_error_code = 4
				break
			}
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgObject(_arg, &zobject, nil, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_OBJECT
				_error_code = 4
				break
			}
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return_value.u1.type_info = 2
			return
		}
		break
	}
	if nil == g.Assign(&pzbucket, zend.ZendHashStrFindDeref(zobject.value.obj.handlers.get_properties(&(*zobject)), "bucket", g.SizeOf("\"bucket\"")-1)) {
		core.PhpErrorDocref(nil, 1<<1, "Object has no bucket property")
		return_value.u1.type_info = 2
		return
	}
	if g.Assign(&brigade, (*streams.PhpStreamBucketBrigade)(zend.ZendFetchResource(zbrigade.value.res, "userfilter.bucket brigade", LeBucketBrigade))) == nil {
		return_value.u1.type_info = 2
		return
	}
	if g.Assign(&bucket, (*streams.PhpStreamBucket)(zend.ZendFetchResourceEx(pzbucket, "userfilter.bucket", LeBucket))) == nil {
		return_value.u1.type_info = 2
		return
	}
	if nil != g.Assign(&pzdata, zend.ZendHashStrFindDeref(zobject.value.obj.handlers.get_properties(&(*zobject)), "data", g.SizeOf("\"data\"")-1)) && pzdata.u1.v.type_ == 6 {
		if bucket.own_buf == 0 {
			bucket = streams.PhpStreamBucketMakeWriteable(bucket)
		}
		if bucket.buflen != pzdata.value.str.len_ {
			if bucket.is_persistent != 0 {
				bucket.buf = zend.__zendRealloc(bucket.buf, pzdata.value.str.len_)
			} else {
				bucket.buf = zend._erealloc(bucket.buf, pzdata.value.str.len_)
			}
			bucket.buflen = pzdata.value.str.len_
		}
		memcpy(bucket.buf, pzdata.value.str.val, bucket.buflen)
	}
	if append != 0 {
		streams.PhpStreamBucketAppend(brigade, bucket)
	} else {
		streams.PhpStreamBucketPrepend(brigade, bucket)
	}

	/* This is a hack necessary to accommodate situations where bucket is appended to the stream
	 * multiple times. See bug35916.phpt for reference.
	 */

	if bucket.refcount == 1 {
		bucket.refcount++
	}

	/* This is a hack necessary to accommodate situations where bucket is appended to the stream
	 * multiple times. See bug35916.phpt for reference.
	 */
}

/* }}} */

func ZifStreamBucketPrepend(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpStreamBucketAttach(0, execute_data, return_value)
}

/* }}} */

func ZifStreamBucketAppend(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpStreamBucketAttach(1, execute_data, return_value)
}

/* }}} */

func ZifStreamBucketNew(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var zstream *zend.Zval
	var zbucket zend.Zval
	var stream *core.PhpStream
	var buffer *byte
	var pbuffer *byte
	var buffer_len int
	var bucket *streams.PhpStreamBucket
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 2
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			zend.ZendParseArgZvalDeref(_arg, &zstream, 0)
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgString(_arg, &buffer, &buffer_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = 4
				break
			}
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return_value.u1.type_info = 2
			return
		}
		break
	}
	if g.Assign(&stream, (*core.PhpStream)(zend.ZendFetchResource2Ex(zstream, "stream", streams.PhpFileLeStream(), streams.PhpFileLePstream()))) == nil {
		return_value.u1.type_info = 2
		return
	}
	if stream.is_persistent != 0 {
		pbuffer = zend.__zendMalloc(buffer_len)
	} else {
		pbuffer = zend._emalloc(buffer_len)
	}
	memcpy(pbuffer, buffer, buffer_len)
	bucket = streams.PhpStreamBucketNew(stream, pbuffer, buffer_len, 1, stream.is_persistent)
	if bucket == nil {
		return_value.u1.type_info = 2
		return
	}
	var __z *zend.Zval = &zbucket
	__z.value.res = zend.ZendRegisterResource(bucket, LeBucket)
	__z.u1.type_info = 9 | 1<<0<<8
	zend.ObjectInit(return_value)
	zend.AddPropertyZvalEx(return_value, "bucket", strlen("bucket"), &zbucket)

	/* add_property_zval increments the refcount which is unwanted here */

	zend.ZvalPtrDtor(&zbucket)
	zend.AddPropertyStringlEx(return_value, "data", strlen("data"), bucket.buf, bucket.buflen)
	zend.AddPropertyLongEx(return_value, "datalen", strlen("datalen"), bucket.buflen)
}

/* }}} */

func ZifStreamGetFilters(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var filter_name *zend.ZendString
	var filters_hash *zend.HashTable
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	var __arr *zend.ZendArray = zend._zendNewArray(0)
	var __z *zend.Zval = return_value
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	filters_hash = streams._phpGetStreamFiltersHash()
	if filters_hash != nil {
		for {
			var __ht *zend.HashTable = filters_hash
			var _p *zend.Bucket = __ht.arData
			var _end *zend.Bucket = _p + __ht.nNumUsed
			for ; _p != _end; _p++ {
				var _z *zend.Zval = &_p.val

				if _z.u1.v.type_ == 0 {
					continue
				}
				filter_name = _p.key
				if filter_name != nil {
					zend.AddNextIndexStr(return_value, zend.ZendStringCopy(filter_name))
				}
			}
			break
		}
	}
}

/* }}} */

func ZifStreamFilterRegister(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var filtername *zend.ZendString
	var classname *zend.ZendString
	var fdat *PhpUserFilterData
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 2
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgStr(_arg, &filtername, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = 4
				break
			}
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgStr(_arg, &classname, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = 4
				break
			}
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return_value.u1.type_info = 2
			return
		}
		break
	}
	return_value.u1.type_info = 2
	if filtername.len_ == 0 {
		core.PhpErrorDocref(nil, 1<<1, "Filter name cannot be empty")
		return
	}
	if classname.len_ == 0 {
		core.PhpErrorDocref(nil, 1<<1, "Class name cannot be empty")
		return
	}
	if BasicGlobals.GetUserFilterMap() == nil {
		BasicGlobals.SetUserFilterMap((*zend.HashTable)(zend._emalloc(g.SizeOf("HashTable"))))
		zend._zendHashInit(BasicGlobals.GetUserFilterMap(), 8, zend.DtorFuncT(FilterItemDtor), 0)
	}
	fdat = zend._ecalloc(1, g.SizeOf("struct php_user_filter_data"))
	fdat.SetClassname(zend.ZendStringCopy(classname))
	if zend.ZendHashAddPtr(BasicGlobals.GetUserFilterMap(), filtername, fdat) != nil && streams.PhpStreamFilterRegisterFactoryVolatile(filtername, &UserFilterFactory) == zend.SUCCESS {
		return_value.u1.type_info = 3
	} else {
		zend.ZendStringReleaseEx(classname, 0)
		zend._efree(fdat)
	}
}

/* }}} */
