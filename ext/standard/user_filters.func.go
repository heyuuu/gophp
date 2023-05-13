package standard

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/core/streams"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/operators"
	"github.com/heyuuu/gophp/zend/zpp"
)

func ZifUserFilterNop(executeData zpp.Ex, return_value zpp.Ret) {}
func PhpBucketDtor(res *types.ZendResource) {
	var bucket *streams.PhpStreamBucket = (*streams.PhpStreamBucket)(res.GetPtr())
	if bucket != nil {
		streams.PhpStreamBucketDelref(bucket)
		bucket = nil
	}
}
func ZmStartupUserFilters(type_ int, module_number int) int {
	/* init the filter class ancestor */

	var phpUserFilter = zend.RegisterClass("php_user_filter", nil, UserFilterClassFuncs)
	zend.ZendDeclarePropertyString(phpUserFilter, "filtername", b.SizeOf("\"filtername\"")-1, "", zend.AccPublic)
	zend.ZendDeclarePropertyString(phpUserFilter, "params", b.SizeOf("\"params\"")-1, "", zend.AccPublic)

	/* init the filter resource; it has no dtor, as streams will always clean it up
	 * at the correct time */

	LeUserfilters = zend.ZendRegisterListDestructorsEx(nil, nil, PHP_STREAM_FILTER_RES_NAME, 0)
	if LeUserfilters == types.FAILURE {
		return types.FAILURE
	}

	/* Filters will dispose of their brigades */

	LeBucketBrigade = zend.ZendRegisterListDestructorsEx(nil, nil, PHP_STREAM_BRIGADE_RES_NAME, module_number)

	/* Brigades will dispose of their buckets */

	LeBucket = zend.ZendRegisterListDestructorsEx(PhpBucketDtor, nil, PHP_STREAM_BUCKET_RES_NAME, module_number)
	if LeBucketBrigade == types.FAILURE {
		return types.FAILURE
	}
	zend.RegisterLongConstant("PSFS_PASS_ON", streams.PSFS_PASS_ON, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("PSFS_FEED_ME", streams.PSFS_FEED_ME, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("PSFS_ERR_FATAL", streams.PSFS_ERR_FATAL, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("PSFS_FLAG_NORMAL", streams.PSFS_FLAG_NORMAL, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("PSFS_FLAG_FLUSH_INC", streams.PSFS_FLAG_FLUSH_INC, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("PSFS_FLAG_FLUSH_CLOSE", streams.PSFS_FLAG_FLUSH_CLOSE, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	return types.SUCCESS
}
func ZmDeactivateUserFilters(type_ int, module_number int) int {
	BG__().UserFilterMap = nil
	return types.SUCCESS
}
func UserfilterDtor(thisfilter *core.PhpStreamFilter) {
	var obj *types.Zval = thisfilter.GetAbstract()
	var func_name types.Zval
	var retval types.Zval
	if obj == nil {

		/* If there's no object associated then there's nothing to dispose of */

		return

		/* If there's no object associated then there's nothing to dispose of */

	}
	func_name.SetStringVal("onclose")
	zend.CallUserFunction(obj, &func_name, &retval, 0, nil)
	// zend.ZvalPtrDtor(&retval)
	// zend.ZvalPtrDtor(&func_name)

	/* kill the object */

	// zend.ZvalPtrDtor(obj)

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
	var obj *types.Zval = thisfilter.GetAbstract()
	var func_name types.Zval
	var retval types.Zval
	var args []types.Zval
	var zpropname types.Zval
	var call_result int

	/* the userfilter object probably doesn't exist anymore */

	if zend.CG__().GetUncleanShutdown() != 0 {
		return ret
	}

	/* Make sure the stream is not closed while the filter callback executes. */

	var orig_no_fclose uint32 = stream.GetFlags() & core.PHP_STREAM_FLAG_NO_FCLOSE
	stream.AddFlags(core.PHP_STREAM_FLAG_NO_FCLOSE)
	if !(types.Z_OBJPROP_P(obj).KeyExistsIndirect("stream")) {
		var tmp types.Zval

		/* Give the userfilter class a hook back to the stream */

		core.PhpStreamToZval(stream, &tmp)
		// 		tmp.AddRefcount()
		zend.AddPropertyZval(obj, "stream", &tmp)

		/* add_property_zval increments the refcount which is unwanted here */

		// zend.ZvalPtrDtor(&tmp)

		/* add_property_zval increments the refcount which is unwanted here */

	}
	func_name.SetStringVal("filter")

	/* Setup calling arguments */

	args[0].SetResource(zend.ZendRegisterResource(buckets_in, LeBucketBrigade))
	args[1].SetResource(zend.ZendRegisterResource(buckets_out, LeBucketBrigade))
	if bytes_consumed != nil {
		args[2].SetLong(*bytes_consumed)
	} else {
		args[2].SetNull()
	}
	args[3].SetBool((flags & streams.PSFS_FLAG_FLUSH_CLOSE) != 0)
	call_result = zend.CallUserFunctionEx(obj, &func_name, &retval, 4, args, 0)
	// zend.ZvalPtrDtor(&func_name)
	if call_result == types.SUCCESS && retval.IsNotUndef() {
		operators.ConvertToLong(&retval)
		ret = int(retval.Long())
	} else if call_result == types.FAILURE {
		core.PhpErrorDocref(nil, faults.E_WARNING, "failed to call filter function")
	}
	if bytes_consumed != nil {
		*bytes_consumed = operators.ZvalGetLong(&args[2])
	}
	if buckets_in.GetHead() != nil {
		var bucket *streams.PhpStreamBucket = buckets_in.GetHead()
		core.PhpErrorDocref(nil, faults.E_WARNING, "Unprocessed filter buckets remaining on input brigade")
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

	zpropname.SetStringVal("stream")
	obj.Object().Handlers().GetUnsetProperty()(obj, &zpropname, nil)
	// zend.ZvalPtrDtor(&zpropname)
	// zend.ZvalPtrDtor(&args[3])
	// zend.ZvalPtrDtor(&args[2])
	// zend.ZvalPtrDtor(&args[1])
	// zend.ZvalPtrDtor(&args[0])
	stream.SubFlags(core.PHP_STREAM_FLAG_NO_FCLOSE)
	stream.AddFlags(orig_no_fclose)
	return ret
}
func UserFilterFactoryCreate(filtername *byte, filterparams *types.Zval, persistent uint8) *core.PhpStreamFilter {
	var fdat *PhpUserFilterData = nil
	var filter *core.PhpStreamFilter
	var obj types.Zval
	var zfilter types.Zval
	var func_name types.Zval
	var retval types.Zval
	var len_ int

	/* some sanity checks */

	if persistent != 0 {
		core.PhpErrorDocref(nil, faults.E_WARNING, "cannot use a user-space filter with a persistent stream")
		return nil
	}
	len_ = strlen(filtername)

	/* determine the classname/class entry */
	fdat = BG__().UserFilterMap[b.CastStrAuto(filtername)]
	if nil == fdat {
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
				b.Assert(period[0] == '.')
				period[1] = '*'
				period[2] = '0'
				fdat = BG__().UserFilterMap[wildcard]
				if nil != fdat {
					period = nil
				} else {
					*period = '0'
					period = strrchr(wildcard, '.')
				}
			}
			zend.Efree(wildcard)
		}
		if fdat == nil {
			core.PhpErrorDocref(nil, faults.E_WARNING, "Err, filter \"%s\" is not in the user-filter map, but somehow the user-filter-factory was invoked for it!?", filtername)
			return nil
		}
	}

	/* bind the classname to the actual class */

	if fdat.GetCe() == nil {
		if nil == b.Assign(&(fdat.GetCe()), zend.ZendLookupClass(fdat.GetClassname())) {
			core.PhpErrorDocref(nil, faults.E_WARNING, "user-filter \"%s\" requires class \"%s\", but that class is not defined", filtername, fdat.GetClassname())
			return nil
		}
	}

	/* create the object */

	if zend.ObjectInitEx(&obj, fdat.GetCe()) == types.FAILURE {
		return nil
	}
	filter = streams.PhpStreamFilterAlloc(&UserfilterOps, nil, 0)
	if filter == nil {
		// zend.ZvalPtrDtor(&obj)
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

	func_name.SetStringVal("oncreate")
	zend.CallUserFunction(&obj, &func_name, &retval, 0, nil)
	if retval.IsNotUndef() {
		if retval.IsType(types.IS_FALSE) {

			/* User reported filter creation error "return false;" */

			// zend.ZvalPtrDtor(&retval)

			/* Kill the filter (safely) */

			filter.GetAbstract().SetUndef()
			streams.PhpStreamFilterFree(filter)

			/* Kill the object */

			// zend.ZvalPtrDtor(&obj)

			/* Report failure to filter_alloc */

			return nil

			/* Report failure to filter_alloc */

		}
		// zend.ZvalPtrDtor(&retval)
	}
	// zend.ZvalPtrDtor(&func_name)

	/* set the filter property, this will be used during cleanup */

	zfilter.SetResource(zend.ZendRegisterResource(filter, LeUserfilters))
	filter.GetAbstract().SetObject(obj.Object())
	zend.AddPropertyZval(&obj, "filter", &zfilter)

	/* add_property_zval increments the refcount which is unwanted here */

	// zend.ZvalPtrDtor(&zfilter)
	return filter
}
func FilterItemDtor(zv *types.Zval) {
	var fdat *PhpUserFilterData = zv.Ptr()
	// types.ZendStringReleaseEx(fdat.GetClassname(), 0)
	zend.Efree(fdat)
}
func ZifStreamBucketMakeWriteable(executeData zpp.Ex, return_value zpp.Ret, brigade *types.Zval) {
	var zbrigade *types.Zval
	var zbucket types.Zval
	var brigade *streams.PhpStreamBucketBrigade
	var bucket *streams.PhpStreamBucket
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 1, 0)
			zbrigade = fp.ParseResource()
			if fp.HasError() {
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	if b.Assign(&brigade, (*streams.PhpStreamBucketBrigade)(zend.ZendFetchResource(zbrigade.Resource(), PHP_STREAM_BRIGADE_RES_NAME, LeBucketBrigade))) == nil {
		return_value.SetFalse()
		return
	}
	return_value.SetNull()
	if brigade.GetHead() != nil && b.Assign(&bucket, streams.PhpStreamBucketMakeWriteable(brigade.GetHead())) {
		zbucket.SetResource(zend.ZendRegisterResource(bucket, LeBucket))
		zend.ObjectInit(return_value)
		zend.AddPropertyZval(return_value, "bucket", &zbucket)

		/* add_property_zval increments the refcount which is unwanted here */

		// zend.ZvalPtrDtor(&zbucket)
		zend.AddPropertyStringl(return_value, "data", b.CastStr(bucket.GetBuf(), bucket.GetBuflen()))
		zend.AddPropertyLong(return_value, "datalen", bucket.GetBuflen())
	}
}
func PhpStreamBucketAttach(append int, executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var zbrigade *types.Zval
	var zobject *types.Zval
	var pzbucket *types.Zval
	var pzdata *types.Zval
	var brigade *streams.PhpStreamBucketBrigade
	var bucket *streams.PhpStreamBucket
	for {
		for {
			fp := zpp.FastParseStart(executeData, 2, 2, 0)
			zbrigade = fp.ParseResource()
			zobject = fp.ParseObject()
			if fp.HasError() {
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	if nil == b.Assign(&pzbucket, types.ZendHashStrFindDeref(types.Z_OBJPROP_P(zobject), "bucket")) {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Object has no bucket property")
		return_value.SetFalse()
		return
	}
	if b.Assign(&brigade, (*streams.PhpStreamBucketBrigade)(zend.ZendFetchResource(zbrigade.Resource(), PHP_STREAM_BRIGADE_RES_NAME, LeBucketBrigade))) == nil {
		return_value.SetFalse()
		return
	}
	if b.Assign(&bucket, (*streams.PhpStreamBucket)(zend.ZendFetchResourceEx(pzbucket, PHP_STREAM_BUCKET_RES_NAME, LeBucket))) == nil {
		return_value.SetFalse()
		return
	}
	if nil != b.Assign(&pzdata, types.ZendHashStrFindDeref(types.Z_OBJPROP_P(zobject), "data")) && pzdata.IsString() {
		if bucket.GetOwnBuf() == 0 {
			bucket = streams.PhpStreamBucketMakeWriteable(bucket)
		}
		if bucket.GetBuflen() != pzdata.String().GetLen() {
			bucket.SetBuf(zend.Perealloc(bucket.GetBuf(), pzdata.String().GetLen()))
			bucket.SetBuflen(pzdata.String().GetLen())
		}
		memcpy(bucket.GetBuf(), pzdata.String().GetVal(), bucket.GetBuflen())
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
func ZifStreamBucketPrepend(executeData zpp.Ex, return_value zpp.Ret, brigade *types.Zval, bucket *types.Zval) {
	PhpStreamBucketAttach(0, executeData, return_value)
}
func ZifStreamBucketAppend(executeData zpp.Ex, return_value zpp.Ret, brigade *types.Zval, bucket *types.Zval) {
	PhpStreamBucketAttach(1, executeData, return_value)
}
func ZifStreamBucketNew(executeData zpp.Ex, return_value zpp.Ret, stream *types.Zval, buffer *types.Zval) {
	var zstream *types.Zval
	var zbucket types.Zval
	var stream *core.PhpStream
	var buffer *byte
	var pbuffer *byte
	var buffer_len int
	var bucket *streams.PhpStreamBucket
	for {
		for {
			fp := zpp.FastParseStart(executeData, 2, 2, 0)
			zstream = fp.ParseZval()
			buffer, buffer_len = fp.ParseString()
			if fp.HasError() {
				return_value.SetFalse()
				return
			}
			break
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

	// zend.ZvalPtrDtor(&zbucket)
	zend.AddPropertyStringl(return_value, "data", b.CastStr(bucket.GetBuf(), bucket.GetBuflen()))
	zend.AddPropertyLong(return_value, "datalen", bucket.GetBuflen())
}
func ZifStreamGetFilters(executeData zpp.Ex, return_value zpp.Ret) {
	var filter_name *types.String
	var filters_hash *types.Array
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	zend.ArrayInit(return_value)
	filters_hash = core.PhpGetStreamFiltersHash()
	if filters_hash != nil {
		var __ht *types.Array = filters_hash
		for _, _p := range __ht.ForeachData() {
			var _z *types.Zval = _p.GetVal()

			filter_name = _p.GetKey()
			if filter_name != nil {
				zend.AddNextIndexStr(return_value, filter_name.Copy())
			}
		}
	}
}
func ZifStreamFilterRegister(filtername string, classname string) bool {
	var fdat *PhpUserFilterData
	if filtername == "" {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Filter name cannot be empty")
		return false
	}
	if classname == "" {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Class name cannot be empty")
		return false
	}
	if BG__().UserFilterMap == nil {
		BG__().UserFilterMap = make(map[string]*PhpUserFilterData)
	}
	fdat = NewPhpUserFilterData(nil, classname)

	if _, exist := BG__().UserFilterMap[filtername]; exist {
		return false
	}
	BG__().UserFilterMap[filtername] = fdat

	if streams.PhpStreamFilterRegisterFactoryVolatile(filtername, &UserFilterFactory) != types.SUCCESS {
		return false
	}

	return true
}
