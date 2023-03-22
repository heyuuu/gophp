// <<generate>>

package streams

import (
	b "sik/builtin"
	"sik/core"
	"sik/ext/standard"
	"sik/zend"
	"sik/zend/faults"
	"sik/zend/types"
)

func PhpGetStreamFiltersHashGlobal() *types.Array { return &StreamFiltersHash }
func _phpGetStreamFiltersHash() *types.Array {
	if standard.FG(stream_filters) {
		return standard.FG(stream_filters)
	} else {
		return &StreamFiltersHash
	}
}
func PhpStreamFilterRegisterFactory(filterpattern *byte, factory *PhpStreamFilterFactory) int {
	var ret int
	var str *types.String = types.ZendStringInitInterned(filterpattern, strlen(filterpattern), 1)
	if types.ZendHashAddPtr(&StreamFiltersHash, str.GetStr(), any(factory)) {
		ret = types.SUCCESS
	} else {
		ret = types.FAILURE
	}
	types.ZendStringReleaseEx(str, 1)
	return ret
}
func PhpStreamFilterUnregisterFactory(filterpattern *byte) int {
	return types.ZendHashStrDel(&StreamFiltersHash, filterpattern)
}
func PhpStreamFilterRegisterFactoryVolatile(filterpattern *types.String, factory *PhpStreamFilterFactory) int {
	if !(standard.FG(stream_filters)) {
		zend.ALLOC_HASHTABLE(standard.FG(stream_filters))
		standard.FG(stream_filters) = types.MakeArrayEx(StreamFiltersHash.GetNNumOfElements()+1, nil, 0)
		types.ZendHashCopy(standard.FG(stream_filters), &StreamFiltersHash, nil)
	}
	if types.ZendHashAddPtr(standard.FG(stream_filters), filterpattern.GetStr(), any(factory)) {
		return types.SUCCESS
	} else {
		return types.FAILURE
	}
}
func PhpStreamBucketNew(stream *core.PhpStream, buf *byte, buflen int, own_buf uint8, buf_persistent uint8) *PhpStreamBucket {
	var is_persistent int = stream.GetIsPersistent()
	var bucket *PhpStreamBucket
	bucket = (*PhpStreamBucket)(zend.Pemalloc(b.SizeOf("php_stream_bucket"), is_persistent))
	bucket.SetPrev(nil)
	bucket.SetNext(bucket.GetPrev())
	if is_persistent != 0 && buf_persistent == 0 {

		/* all data in a persistent bucket must also be persistent */

		bucket.SetBuf(zend.Pemalloc(buflen, 1))
		memcpy(bucket.GetBuf(), buf, buflen)
		bucket.SetBuflen(buflen)
		bucket.SetOwnBuf(1)
	} else {
		bucket.SetBuf(buf)
		bucket.SetBuflen(buflen)
		bucket.SetOwnBuf(own_buf)
	}
	bucket.SetIsPersistent(is_persistent)
	bucket.SetRefcount(1)
	bucket.SetBrigade(nil)
	return bucket
}
func PhpStreamBucketMakeWriteable(bucket *PhpStreamBucket) *PhpStreamBucket {
	var retval *PhpStreamBucket
	PhpStreamBucketUnlink(bucket)
	if bucket.GetRefcount() == 1 && bucket.GetOwnBuf() != 0 {
		return bucket
	}
	retval = (*PhpStreamBucket)(zend.Pemalloc(b.SizeOf("php_stream_bucket"), bucket.GetIsPersistent()))
	memcpy(retval, bucket, b.SizeOf("* retval"))
	retval.SetBuf(zend.Pemalloc(retval.GetBuflen(), retval.GetIsPersistent()))
	memcpy(retval.GetBuf(), bucket.GetBuf(), retval.GetBuflen())
	retval.SetRefcount(1)
	retval.SetOwnBuf(1)
	PhpStreamBucketDelref(bucket)
	return retval
}
func PhpStreamBucketSplit(in *PhpStreamBucket, left **PhpStreamBucket, right **PhpStreamBucket, length int) int {
	*left = (*PhpStreamBucket)(zend.Pecalloc(1, b.SizeOf("php_stream_bucket"), in.GetIsPersistent()))
	*right = (*PhpStreamBucket)(zend.Pecalloc(1, b.SizeOf("php_stream_bucket"), in.GetIsPersistent()))
	left.SetBuf(zend.Pemalloc(length, in.GetIsPersistent()))
	left.SetBuflen(length)
	memcpy(left.GetBuf(), in.GetBuf(), length)
	left.SetRefcount(1)
	left.SetOwnBuf(1)
	left.SetIsPersistent(in.GetIsPersistent())
	right.SetBuflen(in.GetBuflen() - length)
	right.SetBuf(zend.Pemalloc(right.GetBuflen(), in.GetIsPersistent()))
	memcpy(right.GetBuf(), in.GetBuf()+length, right.GetBuflen())
	right.SetRefcount(1)
	right.SetOwnBuf(1)
	right.SetIsPersistent(in.GetIsPersistent())
	return types.SUCCESS
}
func PhpStreamBucketDelref(bucket *PhpStreamBucket) {
	if b.PreDec(&(bucket.GetRefcount())) == 0 {
		if bucket.GetOwnBuf() != 0 {
			zend.Pefree(bucket.GetBuf(), bucket.GetIsPersistent())
		}
		zend.Pefree(bucket, bucket.GetIsPersistent())
	}
}
func PhpStreamBucketPrepend(brigade *PhpStreamBucketBrigade, bucket *PhpStreamBucket) {
	bucket.SetNext(brigade.GetHead())
	bucket.SetPrev(nil)
	if brigade.GetHead() != nil {
		brigade.GetHead().SetPrev(bucket)
	} else {
		brigade.SetTail(bucket)
	}
	brigade.SetHead(bucket)
	bucket.SetBrigade(brigade)
}
func PhpStreamBucketAppend(brigade *PhpStreamBucketBrigade, bucket *PhpStreamBucket) {
	if brigade.GetTail() == bucket {
		return
	}
	bucket.SetPrev(brigade.GetTail())
	bucket.SetNext(nil)
	if brigade.GetTail() != nil {
		brigade.GetTail().SetNext(bucket)
	} else {
		brigade.SetHead(bucket)
	}
	brigade.SetTail(bucket)
	bucket.SetBrigade(brigade)
}
func PhpStreamBucketUnlink(bucket *PhpStreamBucket) {
	if bucket.GetPrev() != nil {
		bucket.GetPrev().SetNext(bucket.GetNext())
	} else if bucket.GetBrigade() != nil {
		bucket.GetBrigade().SetHead(bucket.GetNext())
	}
	if bucket.GetNext() != nil {
		bucket.GetNext().SetPrev(bucket.GetPrev())
	} else if bucket.GetBrigade() != nil {
		bucket.GetBrigade().SetTail(bucket.GetPrev())
	}
	bucket.SetBrigade(nil)
	bucket.SetPrev(nil)
	bucket.SetNext(bucket.GetPrev())
}
func PhpStreamFilterCreate(filtername *byte, filterparams *types.Zval, persistent uint8) *core.PhpStreamFilter {
	var filter_hash *types.Array = b.CondF1(standard.FG(stream_filters), func() __auto__ { return standard.FG(stream_filters) }, &StreamFiltersHash)
	var factory *PhpStreamFilterFactory = nil
	var filter *core.PhpStreamFilter = nil
	var n int
	var period *byte
	n = strlen(filtername)
	if nil != b.Assign(&factory, types.ZendHashStrFindPtr(filter_hash, b.CastStr(filtername, n))) {
		filter = factory.GetCreateFilter()(filtername, filterparams, persistent)
	} else if b.Assign(&period, strrchr(filtername, '.')) {

		/* try a wildcard */

		var wildname *byte
		wildname = zend.SafeEmalloc(1, n, 3)
		memcpy(wildname, filtername, n+1)
		period = wildname + (period - filtername)
		for period != nil && filter == nil {
			b.Assert(period[0] == '.')
			period[1] = '*'
			period[2] = '0'
			if nil != b.Assign(&factory, types.ZendHashStrFindPtr(filter_hash, wildname)) {
				filter = factory.GetCreateFilter()(filtername, filterparams, persistent)
			}
			*period = '0'
			period = strrchr(wildname, '.')
		}
		zend.Efree(wildname)
	}
	if filter == nil {

		/* TODO: these need correct docrefs */

		if factory == nil {
			core.PhpErrorDocref(nil, faults.E_WARNING, "unable to locate filter \"%s\"", filtername)
		} else {
			core.PhpErrorDocref(nil, faults.E_WARNING, "unable to create or locate filter \"%s\"", filtername)
		}

		/* TODO: these need correct docrefs */

	}
	return filter
}
func _phpStreamFilterAlloc(fops *PhpStreamFilterOps, abstract any, persistent uint8) *core.PhpStreamFilter {
	var filter *core.PhpStreamFilter
	filter = (*core.PhpStreamFilter)(PemallocRelOrig(b.SizeOf("php_stream_filter"), persistent))
	memset(filter, 0, b.SizeOf("php_stream_filter"))
	filter.SetFops(fops)
	filter.GetAbstract().GetPtr() = abstract
	filter.SetIsPersistent(persistent)
	return filter
}
func PhpStreamFilterFree(filter *core.PhpStreamFilter) {
	if filter.GetFops().GetDtor() != nil {
		filter.GetFops().GetDtor()(filter)
	}
	zend.Pefree(filter, filter.GetIsPersistent())
}
func PhpStreamFilterPrependEx(chain *PhpStreamFilterChain, filter *core.PhpStreamFilter) int {
	filter.SetNext(chain.GetHead())
	filter.SetPrev(nil)
	if chain.GetHead() != nil {
		chain.GetHead().SetPrev(filter)
	} else {
		chain.SetTail(filter)
	}
	chain.SetHead(filter)
	filter.SetChain(chain)
	return types.SUCCESS
}
func _phpStreamFilterPrepend(chain *PhpStreamFilterChain, filter *core.PhpStreamFilter) {
	PhpStreamFilterPrependEx(chain, filter)
}
func PhpStreamFilterAppendEx(chain *PhpStreamFilterChain, filter *core.PhpStreamFilter) int {
	var stream *core.PhpStream = chain.GetStream()
	filter.SetPrev(chain.GetTail())
	filter.SetNext(nil)
	if chain.GetTail() != nil {
		chain.GetTail().SetNext(filter)
	} else {
		chain.SetHead(filter)
	}
	chain.SetTail(filter)
	filter.SetChain(chain)
	if &(stream.GetReadfilters()) == chain && stream.GetWritepos()-stream.GetReadpos() > 0 {

		/* Let's going ahead and wind anything in the buffer through this filter */

		var brig_in PhpStreamBucketBrigade = MakePhpStreamBucketBrigade(nil, nil)
		var brig_out PhpStreamBucketBrigade = MakePhpStreamBucketBrigade(nil, nil)
		var brig_inp *PhpStreamBucketBrigade = &brig_in
		var brig_outp *PhpStreamBucketBrigade = &brig_out
		var status PhpStreamFilterStatusT
		var bucket *PhpStreamBucket
		var consumed int = 0
		bucket = PhpStreamBucketNew(stream, (*byte)(stream.GetReadbuf()+stream.GetReadpos()), stream.GetWritepos()-stream.GetReadpos(), 0, 0)
		PhpStreamBucketAppend(brig_inp, bucket)
		status = filter.GetFops().GetFilter()(stream, filter, brig_inp, brig_outp, &consumed, PSFS_FLAG_NORMAL)
		if stream.GetReadpos()+consumed > uint32(stream.GetWritepos()) {

			/* No behaving filter should cause this. */

			status = PSFS_ERR_FATAL

			/* No behaving filter should cause this. */

		}
		switch status {
		case PSFS_ERR_FATAL:
			for brig_in.GetHead() != nil {
				bucket = brig_in.GetHead()
				PhpStreamBucketUnlink(bucket)
				PhpStreamBucketDelref(bucket)
			}
			for brig_out.GetHead() != nil {
				bucket = brig_out.GetHead()
				PhpStreamBucketUnlink(bucket)
				PhpStreamBucketDelref(bucket)
			}
			core.PhpErrorDocref(nil, faults.E_WARNING, "Filter failed to process pre-buffered data")
			return types.FAILURE
		case PSFS_FEED_ME:

			/* We don't actually need data yet,
			   leave this filter in a feed me state until data is needed.
			   Reset stream's internal read buffer since the filter is "holding" it. */

			stream.SetReadpos(0)
			stream.SetWritepos(0)
		case PSFS_PASS_ON:

			/* If any data is consumed, we cannot rely upon the existing read buffer,
			   as the filtered data must replace the existing data, so invalidate the cache */

			stream.SetWritepos(0)
			stream.SetReadpos(0)
			for brig_outp.GetHead() != nil {
				bucket = brig_outp.GetHead()

				/* Grow buffer to hold this bucket if need be.
				   TODO: See warning in main/stream/streams.c::php_stream_fill_read_buffer */

				if stream.GetReadbuflen()-stream.GetWritepos() < bucket.GetBuflen() {
					stream.SetReadbuflen(stream.GetReadbuflen() + bucket.GetBuflen())
					stream.SetReadbuf(zend.Perealloc(stream.GetReadbuf(), stream.GetReadbuflen(), stream.GetIsPersistent()))
				}
				memcpy(stream.GetReadbuf()+stream.GetWritepos(), bucket.GetBuf(), bucket.GetBuflen())
				stream.SetWritepos(stream.GetWritepos() + bucket.GetBuflen())
				PhpStreamBucketUnlink(bucket)
				PhpStreamBucketDelref(bucket)
			}
		}
	}
	return types.SUCCESS
}
func _phpStreamFilterAppend(chain *PhpStreamFilterChain, filter *core.PhpStreamFilter) {
	if PhpStreamFilterAppendEx(chain, filter) != types.SUCCESS {
		if chain.GetHead() == filter {
			chain.SetHead(nil)
			chain.SetTail(nil)
		} else {
			filter.GetPrev().SetNext(nil)
			chain.SetTail(filter.GetPrev())
		}
	}
}
func _phpStreamFilterFlush(filter *core.PhpStreamFilter, finish int) int {
	var brig_a PhpStreamBucketBrigade = MakePhpStreamBucketBrigade(nil, nil)
	var brig_b PhpStreamBucketBrigade = MakePhpStreamBucketBrigade(nil, nil)
	var inp *PhpStreamBucketBrigade = &brig_a
	var outp *PhpStreamBucketBrigade = &brig_b
	var brig_temp *PhpStreamBucketBrigade
	var bucket *PhpStreamBucket
	var chain *PhpStreamFilterChain
	var current *core.PhpStreamFilter
	var stream *core.PhpStream
	var flushed_size int = 0
	var flags long = b.Cond(finish != 0, PSFS_FLAG_FLUSH_CLOSE, PSFS_FLAG_FLUSH_INC)
	if filter.GetChain() == nil || filter.GetChain().GetStream() == nil {

		/* Filter is not attached to a chain, or chain is somehow not part of a stream */

		return types.FAILURE

		/* Filter is not attached to a chain, or chain is somehow not part of a stream */

	}
	chain = filter.GetChain()
	stream = chain.GetStream()
	for current = filter; current != nil; current = current.GetNext() {
		var status PhpStreamFilterStatusT
		status = current.GetFops().GetFilter()(stream, current, inp, outp, nil, flags)
		if status == PSFS_FEED_ME {

			/* We've flushed the data far enough */

			return types.SUCCESS

			/* We've flushed the data far enough */

		}
		if status == PSFS_ERR_FATAL {
			return types.FAILURE
		}

		/* Otherwise we have data available to PASS_ON
		   Swap the brigades and continue */

		brig_temp = inp
		inp = outp
		outp = brig_temp
		outp.SetHead(nil)
		outp.SetTail(nil)
		flags = PSFS_FLAG_NORMAL
	}

	/* Last filter returned data via PSFS_PASS_ON
	   Do something with it */

	for bucket = inp.GetHead(); bucket != nil; bucket = bucket.GetNext() {
		flushed_size += bucket.GetBuflen()
	}
	if flushed_size == 0 {

		/* Unlikely, but possible */

		return types.SUCCESS

		/* Unlikely, but possible */

	}
	if chain == &(stream.GetReadfilters()) {

		/* Dump any newly flushed data to the read buffer */

		if stream.GetReadpos() > 0 {

			/* Back the buffer up */

			memcpy(stream.GetReadbuf(), stream.GetReadbuf()+stream.GetReadpos(), stream.GetWritepos()-stream.GetReadpos())
			stream.SetReadpos(0)
			stream.SetWritepos(stream.GetWritepos() - stream.GetReadpos())
		}
		if flushed_size > stream.GetReadbuflen()-stream.GetWritepos() {

			/* Grow the buffer */

			stream.SetReadbuf(zend.Perealloc(stream.GetReadbuf(), stream.GetWritepos()+flushed_size+stream.GetChunkSize(), stream.GetIsPersistent()))

			/* Grow the buffer */

		}
		for b.Assign(&bucket, inp.GetHead()) {
			memcpy(stream.GetReadbuf()+stream.GetWritepos(), bucket.GetBuf(), bucket.GetBuflen())
			stream.SetWritepos(stream.GetWritepos() + bucket.GetBuflen())
			PhpStreamBucketUnlink(bucket)
			PhpStreamBucketDelref(bucket)
		}
	} else if chain == &(stream.GetWritefilters()) {

		/* Send flushed data to the stream */

		for b.Assign(&bucket, inp.GetHead()) {
			var count ssize_t = stream.GetOps().GetWrite()(stream, bucket.GetBuf(), bucket.GetBuflen())
			if count > 0 {
				stream.SetPosition(stream.GetPosition() + count)
			}
			PhpStreamBucketUnlink(bucket)
			PhpStreamBucketDelref(bucket)
		}

		/* Send flushed data to the stream */

	}
	return types.SUCCESS
}
func PhpStreamFilterRemove(filter *core.PhpStreamFilter, call_dtor int) *core.PhpStreamFilter {
	if filter.GetPrev() != nil {
		filter.GetPrev().SetNext(filter.GetNext())
	} else {
		filter.GetChain().SetHead(filter.GetNext())
	}
	if filter.GetNext() != nil {
		filter.GetNext().SetPrev(filter.GetPrev())
	} else {
		filter.GetChain().SetTail(filter.GetPrev())
	}
	if filter.GetRes() != nil {
		zend.ZendListDelete(filter.GetRes())
	}
	if call_dtor != 0 {
		PhpStreamFilterFree(filter)
		return nil
	}
	return filter
}
