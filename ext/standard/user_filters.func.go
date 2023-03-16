// <<generate>>

package standard

import (
	b "sik/builtin"
	"sik/core"
	"sik/core/streams"
	"sik/zend"
)

func ZifUserFilterNop(executeData *zend.ZendExecuteData, return_value *zend.Zval) {}
func PhpBucketDtor(res *zend.ZendResource) {
	var bucket *streams.PhpStreamBucket = (*streams.PhpStreamBucket)(res.GetPtr())
	if bucket != nil {
		streams.PhpStreamBucketDelref(bucket)
		bucket = nil
	}
}
func ZmStartupUserFilters(type_ int, module_number int) int {
	var php_user_filter *zend.ZendClassEntry

	/* init the filter class ancestor */

	memset(&UserFilterClassEntry, 0, b.SizeOf("zend_class_entry"))
	UserFilterClassEntry.SetName(zend.ZendStringInitInterned("php_user_filter", b.SizeOf("\"php_user_filter\"")-1, 1))
	UserFilterClassEntry.SetBuiltinFunctions(UserFilterClassFuncs)
	if b.Assign(&php_user_filter, zend.ZendRegisterInternalClass(&UserFilterClassEntry)) == nil {
		return zend.FAILURE
	}
	zend.ZendDeclarePropertyString(php_user_filter, "filtername", b.SizeOf("\"filtername\"")-1, "", zend.ZEND_ACC_PUBLIC)
	zend.ZendDeclarePropertyString(php_user_filter, "params", b.SizeOf("\"params\"")-1, "", zend.ZEND_ACC_PUBLIC)

	/* init the filter resource; it has no dtor, as streams will always clean it up
	 * at the correct time */

	LeUserfilters = zend.ZendRegisterListDestructorsEx(nil, nil, PHP_STREAM_FILTER_RES_NAME, 0)
	if LeUserfilters == zend.FAILURE {
		return zend.FAILURE
	}

	/* Filters will dispose of their brigades */

	LeBucketBrigade = zend.ZendRegisterListDestructorsEx(nil, nil, PHP_STREAM_BRIGADE_RES_NAME, module_number)

	/* Brigades will dispose of their buckets */

	LeBucket = zend.ZendRegisterListDestructorsEx(PhpBucketDtor, nil, PHP_STREAM_BUCKET_RES_NAME, module_number)
	if LeBucketBrigade == zend.FAILURE {
		return zend.FAILURE
	}
	zend.REGISTER_LONG_CONSTANT("PSFS_PASS_ON", streams.PSFS_PASS_ON, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("PSFS_FEED_ME", streams.PSFS_FEED_ME, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("PSFS_ERR_FATAL", streams.PSFS_ERR_FATAL, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("PSFS_FLAG_NORMAL", streams.PSFS_FLAG_NORMAL, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("PSFS_FLAG_FLUSH_INC", streams.PSFS_FLAG_FLUSH_INC, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("PSFS_FLAG_FLUSH_CLOSE", streams.PSFS_FLAG_FLUSH_CLOSE, zend.CONST_CS|zend.CONST_PERSISTENT)
	return zend.SUCCESS
}
func ZmDeactivateUserFilters(type_ int, module_number int) int {
	if BG(user_filter_map) {
		BG(user_filter_map).Destroy()
		zend.Efree(BG(user_filter_map))
		BG(user_filter_map) = nil
	}
	return zend.SUCCESS
}
func UserfilterDtor(thisfilter *core.PhpStreamFilter) {
	var obj *zend.Zval = thisfilter.GetAbstract()
	var func_name zend.Zval
	var retval zend.Zval
	if obj == nil {

		/* If there's no object associated then there's nothing to dispose of */

		return

		/* If there's no object associated then there's nothing to dispose of */

	}
	func_name.SetRawString("onclose")
	zend.CallUserFunction(obj, &func_name, &retval, 0, nil)
	zend.ZvalPtrDtor(&retval)
	zend.ZvalPtrDtor(&func_name)

	/* kill the object */

	zend.ZvalPtrDtor(obj)

	/* kill the object */
}
func UserfilterFilter(
	stream *core.PhpStream,
	thisfilter *core.PhpStreamFilter,
	buckets_in *streams.PhpStreamBucketBrigade,
	buckets_out *streams.PhpStreamBucketBrigade,
	bytes_consumed *int,
	flags int,
) streams.PhpStreamFilterStatusT {
	var ret int = streams.PSFS_ERR_FATAL
	var obj *zend.Zval = thisfilter.GetAbstract()
	var func_name zend.Zval
	var retval zend.Zval
	var args []zend.Zval
	var zpropname zend.Zval
	var call_result int

	/* the userfilter object probably doesn't exist anymore */

	if zend.CG__().GetUncleanShutdown() != 0 {
		return ret
	}

	/* Make sure the stream is not closed while the filter callback executes. */

	var orig_no_fclose uint32 = stream.GetFlags() & core.PHP_STREAM_FLAG_NO_FCLOSE
	stream.AddFlags(core.PHP_STREAM_FLAG_NO_FCLOSE)
	if !(zend.Z_OBJPROP_P(obj).KeyExistsInd("stream")) {
		var tmp zend.Zval

		/* Give the userfilter class a hook back to the stream */

		core.PhpStreamToZval(stream, &tmp)
		tmp.AddRefcount()
		zend.AddPropertyZval(obj, "stream", &tmp)

		/* add_property_zval increments the refcount which is unwanted here */

		zend.ZvalPtrDtor(&tmp)

		/* add_property_zval increments the refcount which is unwanted here */

	}
	func_name.SetRawString("filter")

	/* Setup calling arguments */

	args[0].SetResource(zend.ZendRegisterResource(buckets_in, LeBucketBrigade))
	args[1].SetResource(zend.ZendRegisterResource(buckets_out, LeBucketBrigade))
	if bytes_consumed != nil {
		args[2].SetLong(*bytes_consumed)
	} else {
		args[2].SetNull()
	}
	zend.ZVAL_BOOL(&args[3], (flags&streams.PSFS_FLAG_FLUSH_CLOSE) != 0)
	call_result = zend.CallUserFunctionEx(obj, &func_name, &retval, 4, args, 0)
	zend.ZvalPtrDtor(&func_name)
	if call_result == zend.SUCCESS && retval.GetType() != zend.IS_UNDEF {
		zend.ConvertToLong(&retval)
		ret = int(retval.GetLval())
	} else if call_result == zend.FAILURE {
		core.PhpErrorDocref(nil, zend.E_WARNING, "failed to call filter function")
	}
	if bytes_consumed != nil {
		*bytes_consumed = zend.ZvalGetLong(&args[2])
	}
	if buckets_in.GetHead() != nil {
		var bucket *streams.PhpStreamBucket = buckets_in.GetHead()
		core.PhpErrorDocref(nil, zend.E_WARNING, "Unprocessed filter buckets remaining on input brigade")
		for b.Assign(&bucket, buckets_in.GetHead()) {

			/* Remove unconsumed buckets from the brigade */

			streams.PhpStreamBucketUnlink(bucket)
			streams.PhpStreamBucketDelref(bucket)
		}
	}
	if ret != streams.PSFS_PASS_ON {
		var bucket *streams.PhpStreamBucket = buckets_out.GetHead()
		for bucket != nil {
			streams.PhpStreamBucketUnlink(bucket)
			streams.PhpStreamBucketDelref(bucket)
			bucket = buckets_out.GetHead()
		}
	}

	/* filter resources are cleaned up by the stream destructor,
	 * keeping a reference to the stream resource here would prevent it
	 * from being destroyed properly */

	zpropname.SetRawString("stream")
	zend.Z_OBJ_HT(*obj).GetUnsetProperty()(obj, &zpropname, nil)
	zend.ZvalPtrDtor(&zpropname)
	zend.ZvalPtrDtor(&args[3])
	zend.ZvalPtrDtor(&args[2])
	zend.ZvalPtrDtor(&args[1])
	zend.ZvalPtrDtor(&args[0])
	stream.SubFlags(core.PHP_STREAM_FLAG_NO_FCLOSE)
	stream.AddFlags(orig_no_fclose)
	return ret
}
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
		core.PhpErrorDocref(nil, zend.E_WARNING, "cannot use a user-space filter with a persistent stream")
		return nil
	}
	len_ = strlen(filtername)

	/* determine the classname/class entry */

	if nil == b.Assign(&fdat, zend.ZendHashStrFindPtr(BG(user_filter_map), (*byte)(filtername), len_)) {
		var period *byte

		/* Userspace Filters using ambiguous wildcards could cause problems.
		   i.e.: myfilter.foo.bar will always call into myfilter.foo.*
		         never seeing myfilter.*
		   TODO: Allow failed userfilter creations to continue
		         scanning through the list */

		if b.Assign(&period, strrchr(filtername, '.')) {
			var wildcard *byte = zend.SafeEmalloc(len_, 1, 3)

			/* Search for wildcard matches instead */

			memcpy(wildcard, filtername, len_+1)
			period = wildcard + (period - filtername)
			for period != nil {
				zend.ZEND_ASSERT(period[0] == '.')
				period[1] = '*'
				period[2] = '0'
				if nil != b.Assign(&fdat, zend.ZendHashStrFindPtr(BG(user_filter_map), wildcard, strlen(wildcard))) {
					period = nil
				} else {
					*period = '0'
					period = strrchr(wildcard, '.')
				}
			}
			zend.Efree(wildcard)
		}
		if fdat == nil {
			core.PhpErrorDocref(nil, zend.E_WARNING, "Err, filter \"%s\" is not in the user-filter map, but somehow the user-filter-factory was invoked for it!?", filtername)
			return nil
		}
	}

	/* bind the classname to the actual class */

	if fdat.GetCe() == nil {
		if nil == b.Assign(&(fdat.GetCe()), zend.ZendLookupClass(fdat.GetClassname())) {
			core.PhpErrorDocref(nil, zend.E_WARNING, "user-filter \"%s\" requires class \"%s\", but that class is not defined", filtername, fdat.GetClassname().GetVal())
			return nil
		}
	}

	/* create the object */

	if zend.ObjectInitEx(&obj, fdat.GetCe()) == zend.FAILURE {
		return nil
	}
	filter = streams.PhpStreamFilterAlloc(&UserfilterOps, nil, 0)
	if filter == nil {
		zend.ZvalPtrDtor(&obj)
		return nil
	}

	/* filtername */

	zend.AddPropertyString(&obj, "filtername", b.CastStrAuto(filtername))

	/* and the parameters, if any */

	if filterparams != nil {
		zend.AddPropertyZval(&obj, "params", filterparams)
	} else {
		zend.AddPropertyNull(&obj, "params")
	}

	/* invoke the constructor */

	func_name.SetRawString("oncreate")
	zend.CallUserFunction(&obj, &func_name, &retval, 0, nil)
	if retval.GetType() != zend.IS_UNDEF {
		if retval.IsType(zend.IS_FALSE) {

			/* User reported filter creation error "return false;" */

			zend.ZvalPtrDtor(&retval)

			/* Kill the filter (safely) */

			filter.GetAbstract().SetUndef()
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

	zfilter.SetResource(zend.ZendRegisterResource(filter, LeUserfilters))
	filter.GetAbstract().SetObject(obj.GetObj())
	zend.AddPropertyZval(&obj, "filter", &zfilter)

	/* add_property_zval increments the refcount which is unwanted here */

	zend.ZvalPtrDtor(&zfilter)
	return filter
}
func FilterItemDtor(zv *zend.Zval) {
	var fdat *PhpUserFilterData = zv.GetPtr()
	zend.ZendStringReleaseEx(fdat.GetClassname(), 0)
	zend.Efree(fdat)
}
func ZifStreamBucketMakeWriteable(executeData *zend.ZendExecuteData, return_value *zend.Zval) {
	var zbrigade *zend.Zval
	var zbucket zend.Zval
	var brigade *streams.PhpStreamBucketBrigade
	var bucket *streams.PhpStreamBucket
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.CheckNumArgsException(_min_num_args, _max_num_args)
					} else {
						zend.CheckNumArgsError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(executeData, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgResource(_arg, &zbrigade, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_RESOURCE
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if _error_code != zend.ZPP_ERROR_OK {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return_value.SetFalse()
			return
		}
		break
	}
	if b.Assign(&brigade, (*streams.PhpStreamBucketBrigade)(zend.ZendFetchResource(zbrigade.GetRes(), PHP_STREAM_BRIGADE_RES_NAME, LeBucketBrigade))) == nil {
		return_value.SetFalse()
		return
	}
	return_value.SetNull()
	if brigade.GetHead() != nil && b.Assign(&bucket, streams.PhpStreamBucketMakeWriteable(brigade.GetHead())) {
		zbucket.SetResource(zend.ZendRegisterResource(bucket, LeBucket))
		zend.ObjectInit(return_value)
		zend.AddPropertyZval(return_value, "bucket", &zbucket)

		/* add_property_zval increments the refcount which is unwanted here */

		zend.ZvalPtrDtor(&zbucket)
		zend.AddPropertyStringl(return_value, "data", b.CastStr(bucket.GetBuf(), bucket.GetBuflen()))
		zend.AddPropertyLong(return_value, "datalen", bucket.GetBuflen())
	}
}
func PhpStreamBucketAttach(append int, executeData *zend.ZendExecuteData, return_value *zend.Zval) {
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
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.CheckNumArgsException(_min_num_args, _max_num_args)
					} else {
						zend.CheckNumArgsError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(executeData, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgResource(_arg, &zbrigade, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_RESOURCE
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgObject(_arg, &zobject, nil, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_OBJECT
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if _error_code != zend.ZPP_ERROR_OK {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return_value.SetFalse()
			return
		}
		break
	}
	if nil == b.Assign(&pzbucket, zend.ZendHashStrFindDeref(zend.Z_OBJPROP_P(zobject), "bucket", b.SizeOf("\"bucket\"")-1)) {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Object has no bucket property")
		return_value.SetFalse()
		return
	}
	if b.Assign(&brigade, (*streams.PhpStreamBucketBrigade)(zend.ZendFetchResource(zbrigade.GetRes(), PHP_STREAM_BRIGADE_RES_NAME, LeBucketBrigade))) == nil {
		return_value.SetFalse()
		return
	}
	if b.Assign(&bucket, (*streams.PhpStreamBucket)(zend.ZendFetchResourceEx(pzbucket, PHP_STREAM_BUCKET_RES_NAME, LeBucket))) == nil {
		return_value.SetFalse()
		return
	}
	if nil != b.Assign(&pzdata, zend.ZendHashStrFindDeref(zend.Z_OBJPROP_P(zobject), "data", b.SizeOf("\"data\"")-1)) && pzdata.IsType(zend.IS_STRING) {
		if bucket.GetOwnBuf() == 0 {
			bucket = streams.PhpStreamBucketMakeWriteable(bucket)
		}
		if bucket.GetBuflen() != zend.Z_STRLEN_P(pzdata) {
			bucket.SetBuf(zend.Perealloc(bucket.GetBuf(), zend.Z_STRLEN_P(pzdata), bucket.GetIsPersistent()))
			bucket.SetBuflen(zend.Z_STRLEN_P(pzdata))
		}
		memcpy(bucket.GetBuf(), zend.Z_STRVAL_P(pzdata), bucket.GetBuflen())
	}
	if append != 0 {
		streams.PhpStreamBucketAppend(brigade, bucket)
	} else {
		streams.PhpStreamBucketPrepend(brigade, bucket)
	}

	/* This is a hack necessary to accommodate situations where bucket is appended to the stream
	 * multiple times. See bug35916.phpt for reference.
	 */

	if bucket.GetRefcount() == 1 {
		bucket.GetRefcount()++
	}

	/* This is a hack necessary to accommodate situations where bucket is appended to the stream
	 * multiple times. See bug35916.phpt for reference.
	 */
}
func ZifStreamBucketPrepend(executeData *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpStreamBucketAttach(0, executeData, return_value)
}
func ZifStreamBucketAppend(executeData *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpStreamBucketAttach(1, executeData, return_value)
}
func ZifStreamBucketNew(executeData *zend.ZendExecuteData, return_value *zend.Zval) {
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
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.CheckNumArgsException(_min_num_args, _max_num_args)
					} else {
						zend.CheckNumArgsError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(executeData, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			zend.ZendParseArgZvalDeref(_arg, &zstream, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgString(_arg, &buffer, &buffer_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if _error_code != zend.ZPP_ERROR_OK {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return_value.SetFalse()
			return
		}
		break
	}
	core.PhpStreamFromZval(stream, zstream)
	pbuffer = zend.Pemalloc(buffer_len, stream.GetIsPersistent())
	memcpy(pbuffer, buffer, buffer_len)
	bucket = streams.PhpStreamBucketNew(stream, pbuffer, buffer_len, 1, stream.GetIsPersistent())
	if bucket == nil {
		return_value.SetFalse()
		return
	}
	zbucket.SetResource(zend.ZendRegisterResource(bucket, LeBucket))
	zend.ObjectInit(return_value)
	zend.AddPropertyZval(return_value, "bucket", &zbucket)

	/* add_property_zval increments the refcount which is unwanted here */

	zend.ZvalPtrDtor(&zbucket)
	zend.AddPropertyStringl(return_value, "data", b.CastStr(bucket.GetBuf(), bucket.GetBuflen()))
	zend.AddPropertyLong(return_value, "datalen", bucket.GetBuflen())
}
func ZifStreamGetFilters(executeData *zend.ZendExecuteData, return_value *zend.Zval) {
	var filter_name *zend.ZendString
	var filters_hash *zend.HashTable
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	zend.ArrayInit(return_value)
	filters_hash = core.PhpGetStreamFiltersHash()
	if filters_hash != nil {
		var __ht *zend.HashTable = filters_hash
		for _, _p := range __ht.foreachData() {
			var _z *zend.Zval = _p.GetVal()

			filter_name = _p.GetKey()
			if filter_name != nil {
				zend.AddNextIndexStr(return_value, filter_name.Copy())
			}
		}
	}
}
func ZifStreamFilterRegister(executeData *zend.ZendExecuteData, return_value *zend.Zval) {
	var filtername *zend.ZendString
	var classname *zend.ZendString
	var fdat *PhpUserFilterData
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 2
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.CheckNumArgsException(_min_num_args, _max_num_args)
					} else {
						zend.CheckNumArgsError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(executeData, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgStr(_arg, &filtername, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgStr(_arg, &classname, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if _error_code != zend.ZPP_ERROR_OK {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return_value.SetFalse()
			return
		}
		break
	}
	return_value.SetFalse()
	if filtername.GetLen() == 0 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Filter name cannot be empty")
		return
	}
	if classname.GetLen() == 0 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Class name cannot be empty")
		return
	}
	if !(BG(user_filter_map)) {
		BG(user_filter_map) = (*zend.HashTable)(zend.Emalloc(b.SizeOf("HashTable")))
		zend.ZendHashInit(BG(user_filter_map), 8, nil, zend.DtorFuncT(FilterItemDtor), 0)
	}
	fdat = zend.Ecalloc(1, b.SizeOf("struct php_user_filter_data"))
	fdat.SetClassname(classname.Copy())
	if zend.ZendHashAddPtr(BG(user_filter_map), filtername, fdat) != nil && streams.PhpStreamFilterRegisterFactoryVolatile(filtername, &UserFilterFactory) == zend.SUCCESS {
		return_value.SetTrue()
	} else {
		zend.ZendStringReleaseEx(classname, 0)
		zend.Efree(fdat)
	}
}
