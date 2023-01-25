// <<generate>>

package streams

import (
	b "sik/builtin"
	"sik/core"
	"sik/ext/standard"
	r "sik/runtime"
	"sik/sapi/cli"
	"sik/zend"
)

func PhpFileLeStream() int       { return LeStream }
func PhpFileLePstream() int      { return LePstream }
func PhpFileLeStreamFilter() int { return LeStreamFilter }
func _phpStreamGetUrlStreamWrappersHash() *zend.HashTable {
	if standard.FG(stream_wrappers) {
		return standard.FG(stream_wrappers)
	} else {
		return &UrlStreamWrappersHash
	}
}
func PhpStreamGetUrlStreamWrappersHashGlobal() *zend.HashTable { return &UrlStreamWrappersHash }
func ForgetPersistentResourceIdNumbers(el *zend.Zval) int {
	var stream *core.PhpStream
	var rsrc *zend.ZendResource = zend.Z_RES_P(el)
	if rsrc.type_ != LePstream {
		return 0
	}
	stream = (*core.PhpStream)(rsrc.ptr)
	stream.res = nil
	if stream.ctx != nil {
		zend.ZendListDelete(stream.ctx)
		stream.ctx = nil
	}
	return 0
}
func ZmDeactivateStreams(type_ int, module_number int) int {
	var el *zend.Zval
	for {
		var __ht *zend.HashTable = &(zend.ExecutorGlobals.persistent_list)
		var _p *zend.Bucket = __ht.arData
		var _end *zend.Bucket = _p + __ht.nNumUsed
		for ; _p != _end; _p++ {
			var _z *zend.Zval = &_p.val

			if zend.UNEXPECTED(zend.Z_TYPE_P(_z) == zend.IS_UNDEF) {
				continue
			}
			el = _z
			ForgetPersistentResourceIdNumbers(el)
		}
		break
	}
	return zend.SUCCESS
}
func PhpStreamEncloses(enclosing *core.PhpStream, enclosed *core.PhpStream) *core.PhpStream {
	var orig *core.PhpStream = enclosed.enclosing_stream
	core.PhpStreamAutoCleanup(enclosed)
	enclosed.enclosing_stream = enclosing
	return orig
}
func PhpStreamFromPersistentId(persistent_id *byte, stream **core.PhpStream) int {
	var le *zend.ZendResource
	if b.Assign(&le, zend.ZendHashStrFindPtr(&(zend.ExecutorGlobals.persistent_list), persistent_id, strlen(persistent_id))) != nil {
		if le.type_ == LePstream {
			if stream != nil {
				var regentry *zend.ZendResource = nil

				/* see if this persistent resource already has been loaded to the
				 * regular list; allowing the same resource in several entries in the
				 * regular list causes trouble (see bug #54623) */

				*stream = (*core.PhpStream)(le.ptr)
				for {
					var __ht *zend.HashTable = &(zend.ExecutorGlobals.regular_list)
					var _p *zend.Bucket = __ht.arData
					var _end *zend.Bucket = _p + __ht.nNumUsed
					for ; _p != _end; _p++ {
						var _z *zend.Zval = &_p.val

						if zend.UNEXPECTED(zend.Z_TYPE_P(_z) == zend.IS_UNDEF) {
							continue
						}
						regentry = zend.Z_PTR_P(_z)
						if regentry.ptr == le.ptr {
							zend.GC_ADDREF(regentry)
							(*stream).res = regentry
							return core.PHP_STREAM_PERSISTENT_SUCCESS
						}
					}
					break
				}
				zend.GC_ADDREF(le)
				(*stream).res = zend.ZendRegisterResource(*stream, LePstream)
			}
			return core.PHP_STREAM_PERSISTENT_SUCCESS
		}
		return core.PHP_STREAM_PERSISTENT_FAILURE
	}
	return core.PHP_STREAM_PERSISTENT_NOT_EXIST
}
func PhpGetWrapperErrorsList(wrapper *core.PhpStreamWrapper) *zend.ZendLlist {
	if !(standard.FG(wrapper_errors)) {
		return nil
	} else {
		return (*zend.ZendLlist)(zend.ZendHashStrFindPtr(standard.FG(wrapper_errors), (*byte)(&wrapper), b.SizeOf("wrapper")))
	}
}
func PhpStreamDisplayWrapperErrors(wrapper *core.PhpStreamWrapper, path *byte, caption string) {
	var tmp *byte
	var msg *byte
	var free_msg int = 0
	if zend.ExecutorGlobals.exception != nil {

		/* Don't emit additional warnings if an exception has already been thrown. */

		return

		/* Don't emit additional warnings if an exception has already been thrown. */

	}
	tmp = zend.Estrdup(path)
	if wrapper != nil {
		var err_list *zend.ZendLlist = PhpGetWrapperErrorsList(wrapper)
		if err_list != nil {
			var l int = 0
			var brlen int
			var i int
			var count int = int(zend.ZendLlistCount(err_list))
			var br *byte
			var err_buf_p **byte
			var pos zend.ZendLlistPosition
			if core.PG(html_errors) {
				brlen = 7
				br = "<br />\n"
			} else {
				brlen = 1
				br = "\n"
			}
			err_buf_p = zend.ZendLlistGetFirstEx(err_list, &pos)
			i = 0
			for err_buf_p != nil {
				l += strlen(*err_buf_p)
				if i < count-1 {
					l += brlen
				}
				err_buf_p = zend.ZendLlistGetNextEx(err_list, &pos)
				i++
			}
			msg = zend.Emalloc(l + 1)
			msg[0] = '0'
			err_buf_p = zend.ZendLlistGetFirstEx(err_list, &pos)
			i = 0
			for err_buf_p != nil {
				strcat(msg, *err_buf_p)
				if i < count-1 {
					strcat(msg, br)
				}
				err_buf_p = zend.ZendLlistGetNextEx(err_list, &pos)
				i++
			}
			free_msg = 1
		} else {
			if wrapper == &PhpPlainFilesWrapper {
				msg = strerror(errno)
			} else {
				msg = "operation failed"
			}
		}
	} else {
		msg = "no suitable wrapper could be found"
	}
	core.PhpStripUrlPasswd(tmp)
	core.PhpErrorDocref1(nil, tmp, zend.E_WARNING, "%s: %s", caption, msg)
	zend.Efree(tmp)
	if free_msg != 0 {
		zend.Efree(msg)
	}
}
func PhpStreamTidyWrapperErrorLog(wrapper *core.PhpStreamWrapper) {
	if wrapper != nil && standard.FG(wrapper_errors) {
		zend.ZendHashStrDel(standard.FG(wrapper_errors), (*byte)(&wrapper), b.SizeOf("wrapper"))
	}
}
func WrapperErrorDtor(error any) { zend.Efree(*((**byte)(error))) }
func WrapperListDtor(item *zend.Zval) {
	var list *zend.ZendLlist = (*zend.ZendLlist)(zend.Z_PTR_P(item))
	zend.ZendLlistDestroy(list)
	zend.Efree(list)
}
func PhpStreamWrapperLogError(wrapper *core.PhpStreamWrapper, options int, fmt string, _ ...any) {
	var args va_list
	var buffer *byte = nil
	va_start(args, fmt)
	core.Vspprintf(&buffer, 0, fmt, args)
	va_end(args)
	if (options&core.REPORT_ERRORS) != 0 || wrapper == nil {
		core.PhpErrorDocref(nil, zend.E_WARNING, "%s", buffer)
		zend.Efree(buffer)
	} else {
		var list *zend.ZendLlist = nil
		if !(standard.FG(wrapper_errors)) {
			zend.ALLOC_HASHTABLE(standard.FG(wrapper_errors))
			zend.ZendHashInit(standard.FG(wrapper_errors), 8, nil, WrapperListDtor, 0)
		} else {
			list = zend.ZendHashStrFindPtr(standard.FG(wrapper_errors), (*byte)(&wrapper), b.SizeOf("wrapper"))
		}
		if list == nil {
			var new_list zend.ZendLlist
			zend.ZendLlistInit(&new_list, b.SizeOf("buffer"), WrapperErrorDtor, 0)
			list = zend.ZendHashStrUpdateMem(standard.FG(wrapper_errors), (*byte)(&wrapper), b.SizeOf("wrapper"), &new_list, b.SizeOf("new_list"))
		}

		/* append to linked list */

		zend.ZendLlistAddElement(list, &buffer)

		/* append to linked list */

	}
}
func _phpStreamAlloc(ops *core.PhpStreamOps, abstract any, persistent_id *byte, mode *byte) *core.PhpStream {
	var ret *core.PhpStream
	ret = (*core.PhpStream)(PemallocRelOrig(b.SizeOf("php_stream"), b.Cond(persistent_id != nil, 1, 0)))
	memset(ret, 0, b.SizeOf("php_stream"))
	ret.readfilters.SetStream(ret)
	ret.writefilters.SetStream(ret)
	ret.ops = ops
	ret.abstract = abstract
	if persistent_id != nil {
		ret.is_persistent = 1
	} else {
		ret.is_persistent = 0
	}
	ret.chunk_size = standard.FG(def_chunk_size)
	if standard.FG(auto_detect_line_endings) {
		ret.flags |= core.PHP_STREAM_FLAG_DETECT_EOL
	}
	if persistent_id != nil {
		if nil == zend.ZendRegisterPersistentResource(persistent_id, strlen(persistent_id), ret, LePstream) {
			zend.Pefree(ret, 1)
			return nil
		}
	}
	ret.res = zend.ZendRegisterResource(ret, b.Cond(persistent_id != nil, LePstream, LeStream))
	strlcpy(ret.mode, mode, b.SizeOf("ret -> mode"))
	ret.wrapper = nil
	ret.wrapperthis = nil
	zend.ZVAL_UNDEF(&ret.wrapperdata)
	ret.stdiocast = nil
	ret.orig_path = nil
	ret.ctx = nil
	ret.readbuf = nil
	ret.enclosing_stream = nil
	return ret
}
func _phpStreamFreeEnclosed(stream_enclosed *core.PhpStream, close_options int) int {
	return core.PhpStreamFree(stream_enclosed, close_options|core.PHP_STREAM_FREE_IGNORE_ENCLOSING)
}
func _phpStreamFreePersistent(zv *zend.Zval, pStream any) int {
	var le *zend.ZendResource = zend.Z_RES_P(zv)
	return le.ptr == pStream
}
func _phpStreamFree(stream *core.PhpStream, close_options int) int {
	var ret int = 1
	var preserve_handle int = b.Cond((close_options&core.PHP_STREAM_FREE_PRESERVE_HANDLE) != 0, 1, 0)
	var release_cast int = 1
	var context *core.PhpStreamContext

	/* During shutdown resources may be released before other resources still holding them.
	 * When only resoruces are referenced this is not a problem, because they are refcounted
	 * and will only be fully freed once the refcount drops to zero. However, if php_stream*
	 * is held directly, we don't have this guarantee. To avoid use-after-free we ignore all
	 * stream free operations in shutdown unless they come from the resource list destruction,
	 * or by freeing an enclosed stream (in which case resource list destruction will not have
	 * freed it). */

	if (zend.ExecutorGlobals.flags&zend.EG_FLAGS_IN_RESOURCE_SHUTDOWN) != 0 && (close_options&(core.PHP_STREAM_FREE_RSRC_DTOR|core.PHP_STREAM_FREE_IGNORE_ENCLOSING)) == 0 {
		return 1
	}
	context = core.PHP_STREAM_CONTEXT(stream)
	if (stream.flags & core.PHP_STREAM_FLAG_NO_CLOSE) != 0 {
		preserve_handle = 1
	}
	if stream.in_free != 0 {

		/* hopefully called recursively from the enclosing stream; the pointer was NULLed below */

		if stream.in_free == 1 && (close_options&core.PHP_STREAM_FREE_IGNORE_ENCLOSING) != 0 && stream.enclosing_stream == nil {
			close_options |= core.PHP_STREAM_FREE_RSRC_DTOR
		} else {
			return 1
		}

		/* hopefully called recursively from the enclosing stream; the pointer was NULLed below */

	}
	stream.in_free++

	/* force correct order on enclosing/enclosed stream destruction (only from resource
	 * destructor as in when reverse destroying the resource list) */

	if (close_options&core.PHP_STREAM_FREE_RSRC_DTOR) != 0 && (close_options&core.PHP_STREAM_FREE_IGNORE_ENCLOSING) == 0 && (close_options&(core.PHP_STREAM_FREE_CALL_DTOR|core.PHP_STREAM_FREE_RELEASE_STREAM)) != 0 && stream.enclosing_stream != nil {
		var enclosing_stream *core.PhpStream = stream.enclosing_stream
		stream.enclosing_stream = nil

		/* we force PHP_STREAM_CALL_DTOR because that's from where the
		 * enclosing stream can free this stream. */

		return core.PhpStreamFree(enclosing_stream, (close_options|core.PHP_STREAM_FREE_CALL_DTOR|core.PHP_STREAM_FREE_KEEP_RSRC) & ^core.PHP_STREAM_FREE_RSRC_DTOR)

		/* we force PHP_STREAM_CALL_DTOR because that's from where the
		 * enclosing stream can free this stream. */

	}

	/* if we are releasing the stream only (and preserving the underlying handle),
	 * we need to do things a little differently.
	 * We are only ever called like this when the stream is cast to a FILE*
	 * for include (or other similar) purposes.
	 * */

	if preserve_handle != 0 {
		if stream.fclose_stdiocast == core.PHP_STREAM_FCLOSE_FOPENCOOKIE {

			/* If the stream was fopencookied, we must NOT touch anything
			 * here, as the cookied stream relies on it all.
			 * Instead, mark the stream as OK to auto-clean */

			core.PhpStreamAutoCleanup(stream)
			stream.in_free--
			return 0
		}

		/* otherwise, make sure that we don't close the FILE* from a cast */

		release_cast = 0

		/* otherwise, make sure that we don't close the FILE* from a cast */

	}
	if (stream.flags&core.PHP_STREAM_FLAG_WAS_WRITTEN) != 0 || stream.writefilters.GetHead() != nil {

		/* make sure everything is saved */

		_phpStreamFlush(stream, 1)

		/* make sure everything is saved */

	}

	/* If not called from the resource dtor, remove the stream from the resource list. */

	if (close_options&core.PHP_STREAM_FREE_RSRC_DTOR) == 0 && stream.res != nil {

		/* Close resource, but keep it in resource list */

		zend.ZendListClose(stream.res)
		if (close_options & core.PHP_STREAM_FREE_KEEP_RSRC) == 0 {

			/* Completely delete zend_resource, if not referenced */

			zend.ZendListDelete(stream.res)
			stream.res = nil
		}
	}
	if (close_options & core.PHP_STREAM_FREE_CALL_DTOR) != 0 {
		if release_cast != 0 && stream.fclose_stdiocast == core.PHP_STREAM_FCLOSE_FOPENCOOKIE {

			/* calling fclose on an fopencookied stream will ultimately
			   call this very same function.  If we were called via fclose,
			   the cookie_closer unsets the fclose_stdiocast flags, so
			   we can be sure that we only reach here when PHP code calls
			   php_stream_free.
			   Lets let the cookie code clean it all up.
			*/

			stream.in_free = 0
			return r.Fclose(stream.stdiocast)
		}
		ret = stream.ops.close(stream, b.Cond(preserve_handle != 0, 0, 1))
		stream.abstract = nil

		/* tidy up any FILE* that might have been fdopened */

		if release_cast != 0 && stream.fclose_stdiocast == core.PHP_STREAM_FCLOSE_FDOPEN && stream.stdiocast != nil {
			r.Fclose(stream.stdiocast)
			stream.stdiocast = nil
			stream.fclose_stdiocast = core.PHP_STREAM_FCLOSE_NONE
		}

		/* tidy up any FILE* that might have been fdopened */

	}
	if (close_options & core.PHP_STREAM_FREE_RELEASE_STREAM) != 0 {
		for stream.readfilters.GetHead() != nil {
			if stream.readfilters.GetHead().GetRes() != nil {
				zend.ZendListClose(stream.readfilters.GetHead().GetRes())
			}
			PhpStreamFilterRemove(stream.readfilters.GetHead(), 1)
		}
		for stream.writefilters.GetHead() != nil {
			if stream.writefilters.GetHead().GetRes() != nil {
				zend.ZendListClose(stream.writefilters.GetHead().GetRes())
			}
			PhpStreamFilterRemove(stream.writefilters.GetHead(), 1)
		}
		if stream.wrapper != nil && stream.wrapper.wops != nil && stream.wrapper.wops.stream_closer != nil {
			stream.wrapper.wops.stream_closer(stream.wrapper, stream)
			stream.wrapper = nil
		}
		if zend.Z_TYPE(stream.wrapperdata) != zend.IS_UNDEF {
			zend.ZvalPtrDtor(&stream.wrapperdata)
			zend.ZVAL_UNDEF(&stream.wrapperdata)
		}
		if stream.readbuf != nil {
			zend.Pefree(stream.readbuf, stream.is_persistent)
			stream.readbuf = nil
		}
		if stream.is_persistent != 0 && (close_options&core.PHP_STREAM_FREE_PERSISTENT) != 0 {

			/* we don't work with *stream but need its value for comparison */

			zend.ZendHashApplyWithArgument(&(zend.ExecutorGlobals.persistent_list), _phpStreamFreePersistent, stream)

			/* we don't work with *stream but need its value for comparison */

		}
		if stream.orig_path != nil {
			zend.Pefree(stream.orig_path, stream.is_persistent)
			stream.orig_path = nil
		}
		zend.Pefree(stream, stream.is_persistent)
	}
	if context != nil {
		zend.ZendListDelete(context.GetRes())
	}
	return ret
}
func _phpStreamFillReadBuffer(stream *core.PhpStream, size int) int {
	/* allocate/fill the buffer */

	if stream.readfilters.GetHead() != nil {
		var to_read_now int = cli.MIN(size, stream.chunk_size)
		var chunk_buf *byte
		var brig_in PhpStreamBucketBrigade = PhpStreamBucketBrigade{nil, nil}
		var brig_out PhpStreamBucketBrigade = PhpStreamBucketBrigade{nil, nil}
		var brig_inp *PhpStreamBucketBrigade = &brig_in
		var brig_outp *PhpStreamBucketBrigade = &brig_out
		var brig_swap *PhpStreamBucketBrigade

		/* allocate a buffer for reading chunks */

		chunk_buf = zend.Emalloc(stream.chunk_size)
		for stream.eof == 0 && stream.writepos-stream.readpos < zend.ZendOffT(to_read_now) {
			var justread ssize_t = 0
			var flags int
			var bucket *PhpStreamBucket
			var status PhpStreamFilterStatusT = PSFS_ERR_FATAL
			var filter *core.PhpStreamFilter

			/* read a chunk into a bucket */

			justread = stream.ops.read(stream, chunk_buf, stream.chunk_size)
			if justread < 0 && stream.writepos == stream.readpos {
				zend.Efree(chunk_buf)
				return zend.FAILURE
			} else if justread > 0 {
				bucket = PhpStreamBucketNew(stream, chunk_buf, justread, 0, 0)

				/* after this call, bucket is owned by the brigade */

				PhpStreamBucketAppend(brig_inp, bucket)
				if stream.eof != 0 {
					flags = PSFS_FLAG_FLUSH_CLOSE
				} else {
					flags = PSFS_FLAG_NORMAL
				}
			} else {
				if stream.eof != 0 {
					flags = PSFS_FLAG_FLUSH_CLOSE
				} else {
					flags = PSFS_FLAG_FLUSH_INC
				}
			}

			/* wind the handle... */

			for filter = stream.readfilters.GetHead(); filter != nil; filter = filter.GetNext() {
				status = filter.GetFops().GetFilter()(stream, filter, brig_inp, brig_outp, nil, flags)
				if status != PSFS_PASS_ON {
					break
				}

				/* brig_out becomes brig_in.
				 * brig_in will always be empty here, as the filter MUST attach any un-consumed buckets
				 * to its own brigade */

				brig_swap = brig_inp
				brig_inp = brig_outp
				brig_outp = brig_swap
				memset(brig_outp, 0, b.SizeOf("* brig_outp"))
			}
			switch status {
			case PSFS_PASS_ON:

				/* we get here when the last filter in the chain has data to pass on.
				 * in this situation, we are passing the brig_in brigade into the
				 * stream read buffer */

				for brig_inp.GetHead() != nil {
					bucket = brig_inp.GetHead()

					/* reduce buffer memory consumption if possible, to avoid a realloc */

					if stream.readbuf != nil && stream.readbuflen-stream.writepos < bucket.GetBuflen() {
						if stream.writepos > stream.readpos {
							memmove(stream.readbuf, stream.readbuf+stream.readpos, stream.writepos-stream.readpos)
						}
						stream.writepos -= stream.readpos
						stream.readpos = 0
					}

					/* grow buffer to hold this bucket */

					if stream.readbuflen-stream.writepos < bucket.GetBuflen() {
						stream.readbuflen += bucket.GetBuflen()
						stream.readbuf = zend.Perealloc(stream.readbuf, stream.readbuflen, stream.is_persistent)
					}
					if bucket.GetBuflen() != 0 {
						memcpy(stream.readbuf+stream.writepos, bucket.GetBuf(), bucket.GetBuflen())
					}
					stream.writepos += bucket.GetBuflen()
					PhpStreamBucketUnlink(bucket)
					PhpStreamBucketDelref(bucket)
				}
				break
			case PSFS_FEED_ME:

				/* when a filter needs feeding, there is no brig_out to deal with.
				 * we simply continue the loop; if the caller needs more data,
				 * we will read again, otherwise out job is done here */

				break
			case PSFS_ERR_FATAL:

				/* some fatal error. Theoretically, the stream is borked, so all
				 * further reads should fail. */

				stream.eof = 1
				zend.Efree(chunk_buf)
				return zend.FAILURE
			}
			if justread <= 0 {
				break
			}
		}
		zend.Efree(chunk_buf)
		return zend.SUCCESS
	} else {

		/* is there enough data in the buffer ? */

		if stream.writepos-stream.readpos < zend.ZendOffT(size) {
			var justread ssize_t = 0

			/* reduce buffer memory consumption if possible, to avoid a realloc */

			if stream.readbuf != nil && stream.readbuflen-stream.writepos < stream.chunk_size {
				if stream.writepos > stream.readpos {
					memmove(stream.readbuf, stream.readbuf+stream.readpos, stream.writepos-stream.readpos)
				}
				stream.writepos -= stream.readpos
				stream.readpos = 0
			}

			/* grow the buffer if required
			 * TODO: this can fail for persistent streams */

			if stream.readbuflen-stream.writepos < stream.chunk_size {
				stream.readbuflen += stream.chunk_size
				stream.readbuf = zend.Perealloc(stream.readbuf, stream.readbuflen, stream.is_persistent)
			}
			justread = stream.ops.read(stream, (*byte)(stream.readbuf+stream.writepos), stream.readbuflen-stream.writepos)
			if justread < 0 {
				return zend.FAILURE
			}
			stream.writepos += justread
		}
		return zend.SUCCESS
	}

	/* allocate/fill the buffer */
}
func _phpStreamRead(stream *core.PhpStream, buf *byte, size int) ssize_t {
	var toread ssize_t = 0
	var didread ssize_t = 0
	for size > 0 {

		/* take from the read buffer first.
		 * It is possible that a buffered stream was switched to non-buffered, so we
		 * drain the remainder of the buffer before using the "raw" read mode for
		 * the excess */

		if stream.writepos > stream.readpos {
			toread = stream.writepos - stream.readpos
			if toread > size {
				toread = size
			}
			memcpy(buf, stream.readbuf+stream.readpos, toread)
			stream.readpos += toread
			size -= toread
			buf += toread
			didread += toread
		}

		/* ignore eof here; the underlying state might have changed */

		if size == 0 {
			break
		}
		if stream.readfilters.GetHead() == nil && ((stream.flags&core.PHP_STREAM_FLAG_NO_BUFFER) != 0 || stream.chunk_size == 1) {
			toread = stream.ops.read(stream, buf, size)
			if toread < 0 {

				/* Report an error if the read failed and we did not read any data
				 * before that. Otherwise return the data we did read. */

				if didread == 0 {
					return toread
				}
				break
			}
		} else {
			if core.PhpStreamFillReadBuffer(stream, size) != zend.SUCCESS {
				if didread == 0 {
					return -1
				}
				break
			}
			toread = stream.writepos - stream.readpos
			if int(toread > size) != 0 {
				toread = size
			}
			if toread > 0 {
				memcpy(buf, stream.readbuf+stream.readpos, toread)
				stream.readpos += toread
			}
		}
		if toread > 0 {
			didread += toread
			buf += toread
			size -= toread
		} else {

			/* EOF, or temporary end of data (for non-blocking mode). */

			break

			/* EOF, or temporary end of data (for non-blocking mode). */

		}

		/* just break anyway, to avoid greedy read for file://, php://memory, and php://temp */

		if stream.wrapper != &PhpPlainFilesWrapper && stream.ops != &PhpStreamMemoryOps && stream.ops != &PhpStreamTempOps {
			break
		}

		/* just break anyway, to avoid greedy read for file://, php://memory, and php://temp */

	}
	if didread > 0 {
		stream.position += didread
	}
	return didread
}
func PhpStreamReadToStr(stream *core.PhpStream, len_ int) *zend.ZendString {
	var str *zend.ZendString = zend.ZendStringAlloc(len_, 0)
	var read ssize_t = core.PhpStreamRead(stream, zend.ZSTR_VAL(str), len_)
	if read < 0 {
		zend.ZendStringEfree(str)
		return nil
	}
	zend.ZSTR_LEN(str) = read
	zend.ZSTR_VAL(str)[read] = 0
	if int(read < len_/2) != 0 {
		return zend.ZendStringTruncate(str, read, 0)
	}
	return str
}
func _phpStreamEof(stream *core.PhpStream) int {
	/* if there is data in the buffer, it's not EOF */

	if stream.writepos-stream.readpos > 0 {
		return 0
	}

	/* use the configured timeout when checking eof */

	if stream.eof == 0 && core.PHP_STREAM_OPTION_RETURN_ERR == core.PhpStreamSetOption(stream, core.PHP_STREAM_OPTION_CHECK_LIVENESS, 0, nil) {
		stream.eof = 1
	}
	return stream.eof
}
func _phpStreamPutc(stream *core.PhpStream, c int) int {
	var buf uint8 = c
	if core.PhpStreamWrite(stream, (*byte)(&buf), 1) > 0 {
		return 1
	}
	return r.EOF
}
func _phpStreamGetc(stream *core.PhpStream) int {
	var buf byte
	if core.PhpStreamRead(stream, &buf, 1) > 0 {
		return buf & 0xff
	}
	return r.EOF
}
func _phpStreamPuts(stream *core.PhpStream, buf *byte) int {
	var len_ int
	var newline []byte = "\n"
	len_ = strlen(buf)
	if len_ > 0 && core.PhpStreamWrite(stream, buf, len_) > 0 && core.PhpStreamWrite(stream, newline, 1) > 0 {
		return 1
	}
	return 0
}
func _phpStreamStat(stream *core.PhpStream, ssb *core.PhpStreamStatbuf) int {
	memset(ssb, 0, b.SizeOf("* ssb"))

	/* if the stream was wrapped, allow the wrapper to stat it */

	if stream.wrapper != nil && stream.wrapper.wops.stream_stat != nil {
		return stream.wrapper.wops.stream_stat(stream.wrapper, stream, ssb)
	}

	/* if the stream doesn't directly support stat-ing, return with failure.
	 * We could try and emulate this by casting to a FD and fstat-ing it,
	 * but since the fd might not represent the actual underlying content
	 * this would give bogus results. */

	if stream.ops.stat == nil {
		return -1
	}
	return stream.ops.stat(stream, ssb)
}
func PhpStreamLocateEol(stream *core.PhpStream, buf *zend.ZendString) *byte {
	var avail int
	var cr *byte
	var lf *byte
	var eol *byte = nil
	var readptr *byte
	if buf == nil {
		readptr = (*byte)(stream.readbuf + stream.readpos)
		avail = stream.writepos - stream.readpos
	} else {
		readptr = zend.ZSTR_VAL(buf)
		avail = zend.ZSTR_LEN(buf)
	}

	/* Look for EOL */

	if (stream.flags & core.PHP_STREAM_FLAG_DETECT_EOL) != 0 {
		cr = memchr(readptr, '\r', avail)
		lf = memchr(readptr, '\n', avail)
		if cr != nil && lf != cr+1 && !(lf != nil && lf < cr) {

			/* mac */

			stream.flags ^= core.PHP_STREAM_FLAG_DETECT_EOL
			stream.flags |= core.PHP_STREAM_FLAG_EOL_MAC
			eol = cr
		} else if cr != nil && lf != nil && cr == lf-1 || lf != nil {

			/* dos or unix endings */

			stream.flags ^= core.PHP_STREAM_FLAG_DETECT_EOL
			eol = lf
		}
	} else if (stream.flags & core.PHP_STREAM_FLAG_EOL_MAC) != 0 {
		eol = memchr(readptr, '\r', avail)
	} else {

		/* unix (and dos) line endings */

		eol = memchr(readptr, '\n', avail)

		/* unix (and dos) line endings */

	}
	return eol
}
func _phpStreamGetLine(stream *core.PhpStream, buf *byte, maxlen int, returned_len *int) *byte {
	var avail int = 0
	var current_buf_size int = 0
	var total_copied int = 0
	var grow_mode int = 0
	var bufstart *byte = buf
	if buf == nil {
		grow_mode = 1
	} else if maxlen == 0 {
		return nil
	}

	/*
	 * If the underlying stream operations block when no new data is readable,
	 * we need to take extra precautions.
	 *
	 * If there is buffered data available, we check for a EOL. If it exists,
	 * we pass the data immediately back to the caller. This saves a call
	 * to the read implementation and will not block where blocking
	 * is not necessary at all.
	 *
	 * If the stream buffer contains more data than the caller requested,
	 * we can also avoid that costly step and simply return that data.
	 */

	for {
		avail = stream.writepos - stream.readpos
		if avail > 0 {
			var cpysz int = 0
			var readptr *byte
			var eol *byte
			var done int = 0
			readptr = (*byte)(stream.readbuf + stream.readpos)
			eol = PhpStreamLocateEol(stream, nil)
			if eol != nil {
				cpysz = eol - readptr + 1
				done = 1
			} else {
				cpysz = avail
			}
			if grow_mode != 0 {

				/* allow room for a NUL. If this realloc is really a realloc
				 * (ie: second time around), we get an extra byte. In most
				 * cases, with the default chunk size of 8K, we will only
				 * incur that overhead once.  When people have lines longer
				 * than 8K, we waste 1 byte per additional 8K or so.
				 * That seems acceptable to me, to avoid making this code
				 * hard to follow */

				bufstart = zend.Erealloc(bufstart, current_buf_size+cpysz+1)
				current_buf_size += cpysz + 1
				buf = bufstart + total_copied
			} else {
				if cpysz >= maxlen-1 {
					cpysz = maxlen - 1
					done = 1
				}
			}
			memcpy(buf, readptr, cpysz)
			stream.position += cpysz
			stream.readpos += cpysz
			buf += cpysz
			maxlen -= cpysz
			total_copied += cpysz
			if done != 0 {
				break
			}
		} else if stream.eof != 0 {
			break
		} else {

			/* XXX: Should be fine to always read chunk_size */

			var toread int
			if grow_mode != 0 {
				toread = stream.chunk_size
			} else {
				toread = maxlen - 1
				if toread > stream.chunk_size {
					toread = stream.chunk_size
				}
			}
			core.PhpStreamFillReadBuffer(stream, toread)
			if stream.writepos-stream.readpos == 0 {
				break
			}
		}
	}
	if total_copied == 0 {
		if grow_mode != 0 {
			r.Assert(bufstart == nil)
		}
		return nil
	}
	buf[0] = '0'
	if returned_len != nil {
		*returned_len = total_copied
	}
	return bufstart
}
func STREAM_BUFFERED_AMOUNT(stream *core.PhpStream) __auto__ {
	return size_t(stream.writepos - stream.readpos)
}
func _phpStreamSearchDelim(stream *core.PhpStream, maxlen int, skiplen int, delim *byte, delim_len int) *byte {
	var seek_len int

	/* set the maximum number of bytes we're allowed to read from buffer */

	seek_len = cli.MIN(STREAM_BUFFERED_AMOUNT(stream), maxlen)
	if seek_len <= skiplen {
		return nil
	}
	if delim_len == 1 {
		return memchr(&stream.readbuf[stream.readpos+skiplen], delim[0], seek_len-skiplen)
	} else {
		return core.PhpMemnstr((*byte)(&stream.readbuf[stream.readpos+skiplen]), delim, delim_len, (*byte)(&stream.readbuf[stream.readpos+seek_len]))
	}
}
func PhpStreamGetRecord(stream *core.PhpStream, maxlen int, delim *byte, delim_len int) *zend.ZendString {
	var ret_buf *zend.ZendString
	var found_delim *byte = nil
	var buffered_len int
	var tent_ret_len int
	var has_delim int = delim_len > 0
	if maxlen == 0 {
		return nil
	}
	if has_delim != 0 {
		found_delim = _phpStreamSearchDelim(stream, maxlen, 0, delim, delim_len)
	}
	buffered_len = STREAM_BUFFERED_AMOUNT(stream)

	/* try to read up to maxlen length bytes while we don't find the delim */

	for found_delim == nil && buffered_len < maxlen {
		var just_read int
		var to_read_now int
		to_read_now = cli.MIN(maxlen-buffered_len, stream.chunk_size)
		core.PhpStreamFillReadBuffer(stream, buffered_len+to_read_now)
		just_read = STREAM_BUFFERED_AMOUNT(stream) - buffered_len

		/* Assume the stream is temporarily or permanently out of data */

		if just_read == 0 {
			break
		}
		if has_delim != 0 {

			/* search for delimiter, but skip buffered_len (the number of bytes
			 * buffered before this loop iteration), as they have already been
			 * searched for the delimiter.
			 * The left part of the delimiter may still remain in the buffer,
			 * so subtract up to <delim_len - 1> from buffered_len, which is
			 * the amount of data we skip on this search  as an optimization
			 */

			found_delim = _phpStreamSearchDelim(stream, maxlen, b.Cond(buffered_len >= delim_len-1, buffered_len-(delim_len-1), 0), delim, delim_len)
			if found_delim != nil {
				break
			}
		}
		buffered_len += just_read
	}
	if has_delim != 0 && found_delim != nil {
		tent_ret_len = found_delim - (*byte)(&stream.readbuf[stream.readpos])
	} else if has_delim == 0 && STREAM_BUFFERED_AMOUNT(stream) >= maxlen {
		tent_ret_len = maxlen
	} else {

		/* return with error if the delimiter string (if any) was not found, we
		 * could not completely fill the read buffer with maxlen bytes and we
		 * don't know we've reached end of file. Added with non-blocking streams
		 * in mind, where this situation is frequent */

		if STREAM_BUFFERED_AMOUNT(stream) < maxlen && stream.eof == 0 {
			return nil
		} else if STREAM_BUFFERED_AMOUNT(stream) == 0 && stream.eof != 0 {

			/* refuse to return an empty string just because by accident
			 * we knew of EOF in a read that returned no data */

			return nil

			/* refuse to return an empty string just because by accident
			 * we knew of EOF in a read that returned no data */

		} else {
			tent_ret_len = cli.MIN(STREAM_BUFFERED_AMOUNT(stream), maxlen)
		}

		/* return with error if the delimiter string (if any) was not found, we
		 * could not completely fill the read buffer with maxlen bytes and we
		 * don't know we've reached end of file. Added with non-blocking streams
		 * in mind, where this situation is frequent */

	}
	ret_buf = zend.ZendStringAlloc(tent_ret_len, 0)

	/* php_stream_read will not call ops->read here because the necessary
	 * data is guaranteedly buffered */

	zend.ZSTR_LEN(ret_buf) = core.PhpStreamRead(stream, zend.ZSTR_VAL(ret_buf), tent_ret_len)
	if found_delim != nil {
		stream.readpos += delim_len
		stream.position += delim_len
	}
	zend.ZSTR_VAL(ret_buf)[zend.ZSTR_LEN(ret_buf)] = '0'
	return ret_buf
}
func _phpStreamWriteBuffer(stream *core.PhpStream, buf *byte, count int) ssize_t {
	var didwrite ssize_t = 0
	var justwrote ssize_t

	/* if we have a seekable stream we need to ensure that data is written at the
	 * current stream->position. This means invalidating the read buffer and then
	* performing a low-level seek */

	if stream.ops.seek != nil && (stream.flags&core.PHP_STREAM_FLAG_NO_SEEK) == 0 && stream.readpos != stream.writepos {
		stream.writepos = 0
		stream.readpos = stream.writepos
		stream.ops.seek(stream, stream.position, r.SEEK_SET, &stream.position)
	}
	for count > 0 {
		var towrite int = count
		if towrite > stream.chunk_size {
			towrite = stream.chunk_size
		}
		justwrote = stream.ops.write(stream, buf, towrite)
		if justwrote <= 0 {

			/* If we already successfully wrote some bytes and a write error occurred
			 * later, report the successfully written bytes. */

			if didwrite == 0 {
				return justwrote
			}
			return didwrite
		}
		buf += justwrote
		count -= justwrote
		didwrite += justwrote
		stream.position += justwrote
	}
	return didwrite
}
func _phpStreamWriteFiltered(stream *core.PhpStream, buf *byte, count int, flags int) ssize_t {
	var consumed int = 0
	var bucket *PhpStreamBucket
	var brig_in PhpStreamBucketBrigade = PhpStreamBucketBrigade{nil, nil}
	var brig_out PhpStreamBucketBrigade = PhpStreamBucketBrigade{nil, nil}
	var brig_inp *PhpStreamBucketBrigade = &brig_in
	var brig_outp *PhpStreamBucketBrigade = &brig_out
	var brig_swap *PhpStreamBucketBrigade
	var status PhpStreamFilterStatusT = PSFS_ERR_FATAL
	var filter *core.PhpStreamFilter
	if buf != nil {
		bucket = PhpStreamBucketNew(stream, (*byte)(buf), count, 0, 0)
		PhpStreamBucketAppend(&brig_in, bucket)
	}
	for filter = stream.writefilters.GetHead(); filter != nil; filter = filter.GetNext() {

		/* for our return value, we are interested in the number of bytes consumed from
		 * the first filter in the chain */

		status = filter.GetFops().GetFilter()(stream, filter, brig_inp, brig_outp, b.Cond(filter == stream.writefilters.GetHead(), &consumed, nil), flags)
		if status != PSFS_PASS_ON {
			break
		}

		/* brig_out becomes brig_in.
		 * brig_in will always be empty here, as the filter MUST attach any un-consumed buckets
		 * to its own brigade */

		brig_swap = brig_inp
		brig_inp = brig_outp
		brig_outp = brig_swap
		memset(brig_outp, 0, b.SizeOf("* brig_outp"))
	}
	switch status {
	case PSFS_PASS_ON:

		/* filter chain generated some output; push it through to the
		 * underlying stream */

		for brig_inp.GetHead() != nil {
			bucket = brig_inp.GetHead()
			if _phpStreamWriteBuffer(stream, bucket.GetBuf(), bucket.GetBuflen()) < 0 {
				consumed = ssize_t - 1
			}

			/* Potential error situation - eg: no space on device. Perhaps we should keep this brigade
			 * hanging around and try to write it later.
			 * At the moment, we just drop it on the floor
			 * */

			PhpStreamBucketUnlink(bucket)
			PhpStreamBucketDelref(bucket)
		}
		break
	case PSFS_FEED_ME:

		/* need more data before we can push data through to the stream */

		break
	case PSFS_ERR_FATAL:

		/* some fatal error.  Theoretically, the stream is borked, so all
		 * further writes should fail. */

		return ssize_t - 1
	}
	return consumed
}
func _phpStreamFlush(stream *core.PhpStream, closing int) int {
	var ret int = 0
	if stream.writefilters.GetHead() != nil {
		_phpStreamWriteFiltered(stream, nil, 0, b.Cond(closing != 0, PSFS_FLAG_FLUSH_CLOSE, PSFS_FLAG_FLUSH_INC))
	}
	stream.flags &= ^core.PHP_STREAM_FLAG_WAS_WRITTEN
	if stream.ops.flush != nil {
		ret = stream.ops.flush(stream)
	}
	return ret
}
func _phpStreamWrite(stream *core.PhpStream, buf *byte, count int) ssize_t {
	var bytes ssize_t
	if count == 0 {
		return 0
	}
	zend.ZEND_ASSERT(buf != nil)
	if stream.ops.write == nil {
		core.PhpErrorDocref(nil, zend.E_NOTICE, "Stream is not writable")
		return ssize_t - 1
	}
	if stream.writefilters.GetHead() != nil {
		bytes = _phpStreamWriteFiltered(stream, buf, count, PSFS_FLAG_NORMAL)
	} else {
		bytes = _phpStreamWriteBuffer(stream, buf, count)
	}
	if bytes {
		stream.flags |= core.PHP_STREAM_FLAG_WAS_WRITTEN
	}
	return bytes
}
func _phpStreamPrintf(stream *core.PhpStream, fmt *byte, _ ...any) ssize_t {
	var count ssize_t
	var buf *byte
	var ap va_list
	va_start(ap, fmt)
	count = core.Vspprintf(&buf, 0, fmt, ap)
	va_end(ap)
	if buf == nil {
		return -1
	}
	count = core.PhpStreamWrite(stream, buf, count)
	zend.Efree(buf)
	return count
}
func _phpStreamTell(stream *core.PhpStream) zend.ZendOffT { return stream.position }
func _phpStreamSeek(stream *core.PhpStream, offset zend.ZendOffT, whence int) int {
	if stream.fclose_stdiocast == core.PHP_STREAM_FCLOSE_FOPENCOOKIE {

		/* flush to commit data written to the fopencookie FILE* */

		r.Fflush(stream.stdiocast)

		/* flush to commit data written to the fopencookie FILE* */

	}

	/* handle the case where we are in the buffer */

	if (stream.flags & core.PHP_STREAM_FLAG_NO_BUFFER) == 0 {
		switch whence {
		case r.SEEK_CUR:
			if offset > 0 && offset <= stream.writepos-stream.readpos {
				stream.readpos += offset
				stream.position += offset
				stream.eof = 0
				return 0
			}
			break
		case r.SEEK_SET:
			if offset > stream.position && offset <= stream.position+stream.writepos-stream.readpos {
				stream.readpos += offset - stream.position
				stream.position = offset
				stream.eof = 0
				return 0
			}
			break
		}
	}
	if stream.ops.seek != nil && (stream.flags&core.PHP_STREAM_FLAG_NO_SEEK) == 0 {
		var ret int
		if stream.writefilters.GetHead() != nil {
			_phpStreamFlush(stream, 0)
		}
		switch whence {
		case r.SEEK_CUR:
			offset = stream.position + offset
			whence = r.SEEK_SET
			break
		}
		ret = stream.ops.seek(stream, offset, whence, &stream.position)
		if (stream.flags&core.PHP_STREAM_FLAG_NO_SEEK) == 0 || ret == 0 {
			if ret == 0 {
				stream.eof = 0
			}

			/* invalidate the buffer contents */

			stream.writepos = 0
			stream.readpos = stream.writepos
			return ret
		}
	}

	/* emulate forward moving seeks with reads */

	if whence == r.SEEK_CUR && offset >= 0 {
		var tmp []byte
		var didread ssize_t
		for offset > 0 {
			if b.Assign(&didread, core.PhpStreamRead(stream, tmp, cli.MIN(offset, b.SizeOf("tmp")))) <= 0 {
				return -1
			}
			offset -= didread
		}
		stream.eof = 0
		return 0
	}
	core.PhpErrorDocref(nil, zend.E_WARNING, "stream does not support seeking")
	return -1
}
func _phpStreamSetOption(stream *core.PhpStream, option int, value int, ptrparam any) int {
	var ret int = core.PHP_STREAM_OPTION_RETURN_NOTIMPL
	if stream.ops.set_option != nil {
		ret = stream.ops.set_option(stream, option, value, ptrparam)
	}
	if ret == core.PHP_STREAM_OPTION_RETURN_NOTIMPL {
		switch option {
		case core.PHP_STREAM_OPTION_SET_CHUNK_SIZE:

			/* XXX chunk size itself is of size_t, that might be ok or not for a particular case*/

			if stream.chunk_size > core.INT_MAX {
				ret = core.INT_MAX
			} else {
				ret = int(stream.chunk_size)
			}
			stream.chunk_size = value
			return ret
		case core.PHP_STREAM_OPTION_READ_BUFFER:

			/* try to match the buffer mode as best we can */

			if value == core.PHP_STREAM_BUFFER_NONE {
				stream.flags |= core.PHP_STREAM_FLAG_NO_BUFFER
			} else if (stream.flags & core.PHP_STREAM_FLAG_NO_BUFFER) != 0 {
				stream.flags ^= core.PHP_STREAM_FLAG_NO_BUFFER
			}
			ret = core.PHP_STREAM_OPTION_RETURN_OK
			break
		default:

		}
	}
	return ret
}
func _phpStreamTruncateSetSize(stream *core.PhpStream, newsize int) int {
	return core.PhpStreamSetOption(stream, core.PHP_STREAM_OPTION_TRUNCATE_API, core.PHP_STREAM_TRUNCATE_SET_SIZE, &newsize)
}
func _phpStreamPassthru(stream *core.PhpStream) ssize_t {
	var bcount int = 0
	var buf []byte
	var b ssize_t
	if PhpStreamMmapPossible(stream) {
		var p *byte
		var mapped int
		p = _phpStreamMmapRange(stream, core.PhpStreamTell(stream), PHP_STREAM_MMAP_ALL, PHP_STREAM_MAP_MODE_SHARED_READONLY, &mapped)
		if p != nil {
			for {

				/* output functions return int, so pass in int max */

				if 0 < b.Assign(&b, core.PHPWRITE(p+bcount, cli.MIN(mapped-bcount, core.INT_MAX))) {
					bcount += b
				}

				/* output functions return int, so pass in int max */

				if !(b > 0 && mapped > bcount) {
					break
				}
			}
			PhpStreamMmapUnmapEx(stream, mapped)
			return bcount
		}
	}
	for b.Assign(&b, core.PhpStreamRead(stream, buf, b.SizeOf("buf"))) > 0 {
		core.PHPWRITE(buf, b)
		bcount += b
	}
	if b < 0 && bcount == 0 {
		return b
	}
	return bcount
}
func _phpStreamCopyToMem(src *core.PhpStream, maxlen int, persistent int) *zend.ZendString {
	var ret ssize_t = 0
	var ptr *byte
	var len_ int = 0
	var max_len int
	var step int = CHUNK_SIZE
	var min_room int = CHUNK_SIZE / 4
	var ssbuf core.PhpStreamStatbuf
	var result *zend.ZendString
	if maxlen == 0 {
		return zend.ZSTR_EMPTY_ALLOC()
	}
	if maxlen == core.PHP_STREAM_COPY_ALL {
		maxlen = 0
	}
	if maxlen > 0 {
		result = zend.ZendStringAlloc(maxlen, persistent)
		ptr = zend.ZSTR_VAL(result)
		for len_ < maxlen && core.PhpStreamEof(src) == 0 {
			ret = core.PhpStreamRead(src, ptr, maxlen-len_)
			if ret <= 0 {

				// TODO: Propagate error?

				break

				// TODO: Propagate error?

			}
			len_ += ret
			ptr += ret
		}
		if len_ != 0 {
			zend.ZSTR_LEN(result) = len_
			zend.ZSTR_VAL(result)[len_] = '0'

			/* Only truncate if the savings are large enough */

			if len_ < maxlen/2 {
				result = zend.ZendStringTruncate(result, len_, persistent)
			}

			/* Only truncate if the savings are large enough */

		} else {
			zend.ZendStringFree(result)
			result = nil
		}
		return result
	}

	/* avoid many reallocs by allocating a good sized chunk to begin with, if
	 * we can.  Note that the stream may be filtered, in which case the stat
	 * result may be inaccurate, as the filter may inflate or deflate the
	 * number of bytes that we can read.  In order to avoid an upsize followed
	 * by a downsize of the buffer, overestimate by the step size (which is
	 * 8K).  */

	if core.PhpStreamStat(src, &ssbuf) == 0 && ssbuf.sb.st_size > 0 {
		max_len = zend.MAX(ssbuf.sb.st_size-src.position, 0) + step
	} else {
		max_len = step
	}
	result = zend.ZendStringAlloc(max_len, persistent)
	ptr = zend.ZSTR_VAL(result)

	// TODO: Propagate error?

	for b.Assign(&ret, core.PhpStreamRead(src, ptr, max_len-len_)) > 0 {
		len_ += ret
		if len_+min_room >= max_len {
			result = zend.ZendStringExtend(result, max_len+step, persistent)
			max_len += step
			ptr = zend.ZSTR_VAL(result) + len_
		} else {
			ptr += ret
		}
	}
	if len_ != 0 {
		result = zend.ZendStringTruncate(result, len_, persistent)
		zend.ZSTR_VAL(result)[len_] = '0'
	} else {
		zend.ZendStringFree(result)
		result = nil
	}
	return result
}
func _phpStreamCopyToStreamEx(src *core.PhpStream, dest *core.PhpStream, maxlen int, len_ *int) int {
	var buf []byte
	var haveread int = 0
	var towrite int
	var dummy int
	var ssbuf core.PhpStreamStatbuf
	if len_ == nil {
		len_ = &dummy
	}
	if maxlen == 0 {
		*len_ = 0
		return zend.SUCCESS
	}
	if maxlen == core.PHP_STREAM_COPY_ALL {
		maxlen = 0
	}
	if core.PhpStreamStat(src, &ssbuf) == 0 {
		if ssbuf.sb.st_size == 0 && zend.S_ISREG(ssbuf.sb.st_mode) {
			*len_ = 0
			return zend.SUCCESS
		}
	}
	if PhpStreamMmapPossible(src) {
		var p *byte
		for {
			var chunk_size int = b.Cond(maxlen == 0 || maxlen > PHP_STREAM_MMAP_MAX, PHP_STREAM_MMAP_MAX, maxlen)
			var mapped int
			p = _phpStreamMmapRange(src, core.PhpStreamTell(src), chunk_size, PHP_STREAM_MAP_MODE_SHARED_READONLY, &mapped)
			if p != nil {
				var didwrite ssize_t
				if core.PhpStreamSeek(src, mapped, r.SEEK_CUR) != 0 {
					PhpStreamMmapUnmap(src)
					break
				}
				didwrite = core.PhpStreamWrite(dest, p, mapped)
				if didwrite < 0 {
					*len_ = haveread
					return zend.FAILURE
				}
				PhpStreamMmapUnmap(src)
				haveread += didwrite
				*len_ = haveread

				/* we've got at least 1 byte to read
				 * less than 1 is an error
				 * AND read bytes match written */

				if mapped == 0 || mapped != didwrite {
					return zend.FAILURE
				}
				if mapped < chunk_size {
					return zend.SUCCESS
				}
				if maxlen != 0 {
					maxlen -= mapped
					if maxlen == 0 {
						return zend.SUCCESS
					}
				}
			}
			if p == nil {
				break
			}
		}
	}
	for true {
		var readchunk int = b.SizeOf("buf")
		var didread ssize_t
		var writeptr *byte
		if maxlen != 0 && maxlen-haveread < readchunk {
			readchunk = maxlen - haveread
		}
		didread = core.PhpStreamRead(src, buf, readchunk)
		if didread <= 0 {
			*len_ = haveread
			if didread < 0 {
				return zend.FAILURE
			} else {
				return zend.SUCCESS
			}
		}
		towrite = didread
		writeptr = buf
		haveread += didread
		for towrite != 0 {
			var didwrite ssize_t = core.PhpStreamWrite(dest, writeptr, towrite)
			if didwrite <= 0 {
				*len_ = haveread - (didread - towrite)
				return zend.FAILURE
			}
			towrite -= didwrite
			writeptr += didwrite
		}
		if maxlen-haveread == 0 {
			break
		}
	}
	*len_ = haveread

	/* we've got at least 1 byte to read.
	 * less than 1 is an error */

	if haveread > 0 || src.eof != 0 {
		return zend.SUCCESS
	}
	return zend.FAILURE
}
func _phpStreamCopyToStream(src *core.PhpStream, dest *core.PhpStream, maxlen int) int {
	var len_ int
	var ret int = _phpStreamCopyToStreamEx(src, dest, maxlen, &len_)
	if ret == zend.SUCCESS && len_ == 0 && maxlen != 0 {
		return 1
	}
	return len_
}
func StreamResourceRegularDtor(rsrc *zend.ZendResource) {
	var stream *core.PhpStream = (*core.PhpStream)(rsrc.ptr)

	/* set the return value for pclose */

	standard.FG(pclose_ret) = core.PhpStreamFree(stream, core.PHP_STREAM_FREE_CLOSE|core.PHP_STREAM_FREE_RSRC_DTOR)

	/* set the return value for pclose */
}
func StreamResourcePersistentDtor(rsrc *zend.ZendResource) {
	var stream *core.PhpStream = (*core.PhpStream)(rsrc.ptr)
	standard.FG(pclose_ret) = core.PhpStreamFree(stream, core.PHP_STREAM_FREE_CLOSE|core.PHP_STREAM_FREE_RSRC_DTOR)
}
func PhpShutdownStreamHashes() {
	if standard.FG(stream_wrappers) {
		zend.ZendHashDestroy(standard.FG(stream_wrappers))
		zend.Efree(standard.FG(stream_wrappers))
		standard.FG(stream_wrappers) = nil
	}
	if standard.FG(stream_filters) {
		zend.ZendHashDestroy(standard.FG(stream_filters))
		zend.Efree(standard.FG(stream_filters))
		standard.FG(stream_filters) = nil
	}
	if standard.FG(wrapper_errors) {
		zend.ZendHashDestroy(standard.FG(wrapper_errors))
		zend.Efree(standard.FG(wrapper_errors))
		standard.FG(wrapper_errors) = nil
	}
}
func PhpInitStreamWrappers(module_number int) int {
	LeStream = zend.ZendRegisterListDestructorsEx(StreamResourceRegularDtor, nil, "stream", module_number)
	LePstream = zend.ZendRegisterListDestructorsEx(nil, StreamResourcePersistentDtor, "persistent stream", module_number)

	/* Filters are cleaned up by the streams they're attached to */

	LeStreamFilter = zend.ZendRegisterListDestructorsEx(nil, nil, "stream filter", module_number)
	zend.ZendHashInit(&UrlStreamWrappersHash, 8, nil, nil, 1)
	zend.ZendHashInit(PhpGetStreamFiltersHashGlobal(), 8, nil, nil, 1)
	zend.ZendHashInit(PhpStreamXportGetHash(), 8, nil, nil, 1)
	if PhpStreamXportRegister("tcp", PhpStreamGenericSocketFactory) == zend.SUCCESS && PhpStreamXportRegister("udp", PhpStreamGenericSocketFactory) == zend.SUCCESS {
		return zend.SUCCESS
	} else {
		return zend.FAILURE
	}
}
func PhpShutdownStreamWrappers(module_number int) int {
	zend.ZendHashDestroy(&UrlStreamWrappersHash)
	zend.ZendHashDestroy(PhpGetStreamFiltersHashGlobal())
	zend.ZendHashDestroy(PhpStreamXportGetHash())
	return zend.SUCCESS
}
func PhpStreamWrapperSchemeValidate(protocol *byte, protocol_len uint) int {
	var i uint
	for i = 0; i < protocol_len; i++ {
		if !(isalnum(int(protocol[i]))) && protocol[i] != '+' && protocol[i] != '-' && protocol[i] != '.' {
			return zend.FAILURE
		}
	}
	return zend.SUCCESS
}
func PhpRegisterUrlStreamWrapper(protocol string, wrapper *core.PhpStreamWrapper) int {
	var protocol_len uint = uint(strlen(protocol))
	var ret int
	var str *zend.ZendString
	if PhpStreamWrapperSchemeValidate(protocol, protocol_len) == zend.FAILURE {
		return zend.FAILURE
	}
	str = zend.ZendStringInitInterned(protocol, protocol_len, 1)
	if zend.ZendHashAddPtr(&UrlStreamWrappersHash, str, any(wrapper)) {
		ret = zend.SUCCESS
	} else {
		ret = zend.FAILURE
	}
	zend.ZendStringReleaseEx(str, 1)
	return ret
}
func PhpUnregisterUrlStreamWrapper(protocol string) int {
	return zend.ZendHashStrDel(&UrlStreamWrappersHash, protocol, strlen(protocol))
}
func CloneWrapperHash() {
	zend.ALLOC_HASHTABLE(standard.FG(stream_wrappers))
	zend.ZendHashInit(standard.FG(stream_wrappers), zend.ZendHashNumElements(&UrlStreamWrappersHash), nil, nil, 0)
	zend.ZendHashCopy(standard.FG(stream_wrappers), &UrlStreamWrappersHash, nil)
}
func PhpRegisterUrlStreamWrapperVolatile(protocol *zend.ZendString, wrapper *core.PhpStreamWrapper) int {
	if PhpStreamWrapperSchemeValidate(zend.ZSTR_VAL(protocol), zend.ZSTR_LEN(protocol)) == zend.FAILURE {
		return zend.FAILURE
	}
	if !(standard.FG(stream_wrappers)) {
		CloneWrapperHash()
	}
	if zend.ZendHashAddPtr(standard.FG(stream_wrappers), protocol, wrapper) {
		return zend.SUCCESS
	} else {
		return zend.FAILURE
	}
}
func PhpUnregisterUrlStreamWrapperVolatile(protocol *zend.ZendString) int {
	if !(standard.FG(stream_wrappers)) {
		CloneWrapperHash()
	}
	return zend.ZendHashDel(standard.FG(stream_wrappers), protocol)
}
func PhpStreamLocateUrlWrapper(path *byte, path_for_open **byte, options int) *core.PhpStreamWrapper {
	var wrapper_hash *zend.HashTable = b.CondF1(standard.FG(stream_wrappers), func() __auto__ { return standard.FG(stream_wrappers) }, &UrlStreamWrappersHash)
	var wrapper *core.PhpStreamWrapper = nil
	var p *byte
	var protocol *byte = nil
	var n int = 0
	if path_for_open != nil {
		*path_for_open = (*byte)(path)
	}
	if (options & core.IGNORE_URL) != 0 {
		return (*core.PhpStreamWrapper)(b.Cond((options&core.STREAM_LOCATE_WRAPPERS_ONLY) != 0, nil, &PhpPlainFilesWrapper))
	}
	for p = path; isalnum(int(*p)) || (*p) == '+' || (*p) == '-' || (*p) == '.'; p++ {
		n++
	}
	if (*p) == ':' && n > 1 && (!(strncmp("//", p+1, 2)) || n == 4 && !(memcmp("data:", path, 5))) {
		protocol = path
	}
	if protocol != nil {
		if nil == b.Assign(&wrapper, zend.ZendHashStrFindPtr(wrapper_hash, protocol, n)) {
			var tmp *byte = zend.Estrndup(protocol, n)
			standard.PhpStrtolower(tmp, n)
			if nil == b.Assign(&wrapper, zend.ZendHashStrFindPtr(wrapper_hash, tmp, n)) {
				var wrapper_name []byte
				if n >= b.SizeOf("wrapper_name") {
					n = b.SizeOf("wrapper_name") - 1
				}
				core.PHP_STRLCPY(wrapper_name, protocol, b.SizeOf("wrapper_name"), n)
				core.PhpErrorDocref(nil, zend.E_WARNING, "Unable to find the wrapper \"%s\" - did you forget to enable it when you configured PHP?", wrapper_name)
				wrapper = nil
				protocol = nil
			}
			zend.Efree(tmp)
		}
	}

	/* TODO: curl based streams probably support file:// properly */

	if protocol == nil || !(strncasecmp(protocol, "file", n)) {

		/* fall back on regular file access */

		var plain_files_wrapper *core.PhpStreamWrapper = (*core.PhpStreamWrapper)(&PhpPlainFilesWrapper)
		if protocol != nil {
			var localhost int = 0
			if !(strncasecmp(path, "file://localhost/", 17)) {
				localhost = 1
			}
			if localhost == 0 && path[n+3] != '0' && path[n+3] != '/' {
				if (options & core.REPORT_ERRORS) != 0 {
					core.PhpErrorDocref(nil, zend.E_WARNING, "remote host file access not supported, %s", path)
				}
				return nil
			}
			if path_for_open != nil {

				/* skip past protocol and :/, but handle windows correctly */

				*path_for_open = (*byte)(path + n + 1)
				if localhost == 1 {
					*path_for_open += 11
				}
				for (*(b.PreInc(&(*path_for_open)))) == '/' {

				}
				*path_for_open--
			}
		}
		if (options & core.STREAM_LOCATE_WRAPPERS_ONLY) != 0 {
			return nil
		}
		if standard.FG(stream_wrappers) {

			/* The file:// wrapper may have been disabled/overridden */

			if wrapper != nil {

				/* It was found so go ahead and provide it */

				return wrapper

				/* It was found so go ahead and provide it */

			}

			/* Check again, the original check might have not known the protocol name */

			if b.Assign(&wrapper, zend.ZendHashFindExPtr(wrapper_hash, zend.ZSTR_KNOWN(zend.ZEND_STR_FILE), 1)) != nil {
				return wrapper
			}
			if (options & core.REPORT_ERRORS) != 0 {
				core.PhpErrorDocref(nil, zend.E_WARNING, "file:// wrapper is disabled in the server configuration")
			}
			return nil
		}
		return plain_files_wrapper
	}
	if wrapper != nil && wrapper.is_url != 0 && (options&core.STREAM_DISABLE_URL_PROTECTION) == 0 && (!(core.PG(allow_url_fopen)) || ((options&core.STREAM_OPEN_FOR_INCLUDE) != 0 || core.PG(in_user_include)) && !(core.PG(allow_url_include))) {
		if (options & core.REPORT_ERRORS) != 0 {

			/* protocol[n] probably isn't '\0' */

			if !(core.PG(allow_url_fopen)) {
				core.PhpErrorDocref(nil, zend.E_WARNING, "%.*s:// wrapper is disabled in the server configuration by allow_url_fopen=0", int(n), protocol)
			} else {
				core.PhpErrorDocref(nil, zend.E_WARNING, "%.*s:// wrapper is disabled in the server configuration by allow_url_include=0", int(n), protocol)
			}

			/* protocol[n] probably isn't '\0' */

		}
		return nil
	}
	return wrapper
}
func _phpStreamMkdir(path *byte, mode int, options int, context *core.PhpStreamContext) int {
	var wrapper *core.PhpStreamWrapper = nil
	wrapper = PhpStreamLocateUrlWrapper(path, nil, 0)
	if wrapper == nil || wrapper.wops == nil || wrapper.wops.stream_mkdir == nil {
		return 0
	}
	return wrapper.wops.stream_mkdir(wrapper, path, mode, options, context)
}
func _phpStreamRmdir(path *byte, options int, context *core.PhpStreamContext) int {
	var wrapper *core.PhpStreamWrapper = nil
	wrapper = PhpStreamLocateUrlWrapper(path, nil, 0)
	if wrapper == nil || wrapper.wops == nil || wrapper.wops.stream_rmdir == nil {
		return 0
	}
	return wrapper.wops.stream_rmdir(wrapper, path, options, context)
}
func _phpStreamStatPath(path *byte, flags int, ssb *core.PhpStreamStatbuf, context *core.PhpStreamContext) int {
	var wrapper *core.PhpStreamWrapper = nil
	var path_to_open *byte = path
	var ret int
	memset(ssb, 0, b.SizeOf("* ssb"))
	if (flags & core.PHP_STREAM_URL_STAT_NOCACHE) == 0 {

		/* Try to hit the cache first */

		if (flags & core.PHP_STREAM_URL_STAT_LINK) != 0 {
			if standard.BG(CurrentLStatFile) && strcmp(path, standard.BG(CurrentLStatFile)) == 0 {
				memcpy(ssb, &standard.BG(lssb), b.SizeOf("php_stream_statbuf"))
				return 0
			}
		} else {
			if standard.BG(CurrentStatFile) && strcmp(path, standard.BG(CurrentStatFile)) == 0 {
				memcpy(ssb, &standard.BG(ssb), b.SizeOf("php_stream_statbuf"))
				return 0
			}
		}

		/* Try to hit the cache first */

	}
	wrapper = PhpStreamLocateUrlWrapper(path, &path_to_open, 0)
	if wrapper != nil && wrapper.wops.url_stat != nil {
		ret = wrapper.wops.url_stat(wrapper, path_to_open, flags, ssb, context)
		if ret == 0 {
			if (flags & core.PHP_STREAM_URL_STAT_NOCACHE) == 0 {

				/* Drop into cache */

				if (flags & core.PHP_STREAM_URL_STAT_LINK) != 0 {
					if standard.BG(CurrentLStatFile) {
						zend.Efree(standard.BG(CurrentLStatFile))
					}
					standard.BG(CurrentLStatFile) = zend.Estrdup(path)
					memcpy(&standard.BG(lssb), ssb, b.SizeOf("php_stream_statbuf"))
				} else {
					if standard.BG(CurrentStatFile) {
						zend.Efree(standard.BG(CurrentStatFile))
					}
					standard.BG(CurrentStatFile) = zend.Estrdup(path)
					memcpy(&standard.BG(ssb), ssb, b.SizeOf("php_stream_statbuf"))
				}

				/* Drop into cache */

			}
		}
		return ret
	}
	return -1
}
func _phpStreamOpendir(path *byte, options int, context *core.PhpStreamContext) *core.PhpStream {
	var stream *core.PhpStream = nil
	var wrapper *core.PhpStreamWrapper = nil
	var path_to_open *byte
	if path == nil || !(*path) {
		return nil
	}
	path_to_open = path
	wrapper = PhpStreamLocateUrlWrapper(path, &path_to_open, options)
	if wrapper != nil && wrapper.wops.dir_opener != nil {
		stream = wrapper.wops.dir_opener(wrapper, path_to_open, "r", options^core.REPORT_ERRORS, nil, context)
		if stream != nil {
			stream.wrapper = wrapper
			stream.flags |= core.PHP_STREAM_FLAG_NO_BUFFER | core.PHP_STREAM_FLAG_IS_DIR
		}
	} else if wrapper != nil {
		PhpStreamWrapperLogError(wrapper, options^core.REPORT_ERRORS, "not implemented")
	}
	if stream == nil && (options&core.REPORT_ERRORS) != 0 {
		PhpStreamDisplayWrapperErrors(wrapper, path, "failed to open dir")
	}
	PhpStreamTidyWrapperErrorLog(wrapper)
	return stream
}
func _phpStreamReaddir(dirstream *core.PhpStream, ent *core.PhpStreamDirent) *core.PhpStreamDirent {
	if b.SizeOf("php_stream_dirent") == core.PhpStreamRead(dirstream, (*byte)(ent), b.SizeOf("php_stream_dirent")) {
		return ent
	}
	return nil
}
func _phpStreamOpenWrapperEx(path *byte, mode *byte, options int, opened_path **zend.ZendString, context *core.PhpStreamContext) *core.PhpStream {
	var stream *core.PhpStream = nil
	var wrapper *core.PhpStreamWrapper = nil
	var path_to_open *byte
	var persistent int = options & core.STREAM_OPEN_PERSISTENT
	var resolved_path *zend.ZendString = nil
	var copy_of_path *byte = nil
	if opened_path != nil {
		*opened_path = nil
	}
	if path == nil || !(*path) {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Filename cannot be empty")
		return nil
	}
	if (options & core.USE_PATH) != 0 {
		resolved_path = zend.ZendResolvePath(path, strlen(path))
		if resolved_path != nil {
			path = zend.ZSTR_VAL(resolved_path)

			/* we've found this file, don't re-check include_path or run realpath */

			options |= core.STREAM_ASSUME_REALPATH
			options &= ^core.USE_PATH
		}
		if zend.ExecutorGlobals.exception != nil {
			return nil
		}
	}
	path_to_open = path
	wrapper = PhpStreamLocateUrlWrapper(path, &path_to_open, options)
	if (options&core.STREAM_USE_URL) != 0 && (wrapper == nil || wrapper.is_url == 0) {
		core.PhpErrorDocref(nil, zend.E_WARNING, "This function may only be used against URLs")
		if resolved_path != nil {
			zend.ZendStringReleaseEx(resolved_path, 0)
		}
		return nil
	}
	if wrapper != nil {
		if wrapper.wops.stream_opener == nil {
			PhpStreamWrapperLogError(wrapper, options^core.REPORT_ERRORS, "wrapper does not support stream open")
		} else {
			stream = wrapper.wops.stream_opener(wrapper, path_to_open, mode, options^core.REPORT_ERRORS, opened_path, context)
		}

		/* if the caller asked for a persistent stream but the wrapper did not
		 * return one, force an error here */

		if stream != nil && (options&core.STREAM_OPEN_PERSISTENT) != 0 && stream.is_persistent == 0 {
			PhpStreamWrapperLogError(wrapper, options^core.REPORT_ERRORS, "wrapper does not support persistent streams")
			core.PhpStreamClose(stream)
			stream = nil
		}
		if stream != nil {
			stream.wrapper = wrapper
		}
	}
	if stream != nil {
		if opened_path != nil && (*opened_path) == nil && resolved_path != nil {
			*opened_path = resolved_path
			resolved_path = nil
		}
		if stream.orig_path != nil {
			zend.Pefree(stream.orig_path, persistent)
		}
		copy_of_path = zend.Pestrdup(path, persistent)
		stream.orig_path = copy_of_path
	}
	if stream != nil && (options&core.STREAM_MUST_SEEK) != 0 {
		var newstream *core.PhpStream
		switch core.PhpStreamMakeSeekableRel(stream, &newstream, b.Cond((options&core.STREAM_WILL_CAST) != 0, core.PHP_STREAM_PREFER_STDIO, core.PHP_STREAM_NO_PREFERENCE)) {
		case core.PHP_STREAM_UNCHANGED:
			if resolved_path != nil {
				zend.ZendStringReleaseEx(resolved_path, 0)
			}
			return stream
		case core.PHP_STREAM_RELEASED:
			if newstream.orig_path != nil {
				zend.Pefree(newstream.orig_path, persistent)
			}
			newstream.orig_path = zend.Pestrdup(path, persistent)
			if resolved_path != nil {
				zend.ZendStringReleaseEx(resolved_path, 0)
			}
			return newstream
		default:
			core.PhpStreamClose(stream)
			stream = nil
			if (options & core.REPORT_ERRORS) != 0 {
				var tmp *byte = zend.Estrdup(path)
				core.PhpStripUrlPasswd(tmp)
				core.PhpErrorDocref1(nil, tmp, zend.E_WARNING, "could not make seekable - %s", tmp)
				zend.Efree(tmp)
				options ^= core.REPORT_ERRORS
			}
		}
	}
	if stream != nil && stream.ops.seek != nil && (stream.flags&core.PHP_STREAM_FLAG_NO_SEEK) == 0 && strchr(mode, 'a') && stream.position == 0 {
		var newpos zend.ZendOffT = 0

		/* if opened for append, we need to revise our idea of the initial file position */

		if 0 == stream.ops.seek(stream, 0, r.SEEK_CUR, &newpos) {
			stream.position = newpos
		}

		/* if opened for append, we need to revise our idea of the initial file position */

	}
	if stream == nil && (options&core.REPORT_ERRORS) != 0 {
		PhpStreamDisplayWrapperErrors(wrapper, path, "failed to open stream")
		if opened_path != nil && (*opened_path) != nil {
			zend.ZendStringReleaseEx(*opened_path, 0)
			*opened_path = nil
		}
	}
	PhpStreamTidyWrapperErrorLog(wrapper)
	if resolved_path != nil {
		zend.ZendStringReleaseEx(resolved_path, 0)
	}
	return stream
}
func PhpStreamContextSet(stream *core.PhpStream, context *core.PhpStreamContext) *core.PhpStreamContext {
	var oldcontext *core.PhpStreamContext = core.PHP_STREAM_CONTEXT(stream)
	if context != nil {
		stream.ctx = context.GetRes()
		zend.GC_ADDREF(context.GetRes())
	} else {
		stream.ctx = nil
	}
	if oldcontext != nil {
		zend.ZendListDelete(oldcontext.GetRes())
	}
	return oldcontext
}
func PhpStreamNotificationNotify(context *core.PhpStreamContext, notifycode int, severity int, xmsg *byte, xcode int, bytes_sofar int, bytes_max int, ptr any) {
	if context != nil && context.GetNotifier() != nil {
		context.GetNotifier().GetFunc()(context, notifycode, severity, xmsg, xcode, bytes_sofar, bytes_max, ptr)
	}
}
func PhpStreamContextFree(context *core.PhpStreamContext) {
	if zend.Z_TYPE(context.GetOptions()) != zend.IS_UNDEF {
		zend.ZvalPtrDtor(&context.options)
		zend.ZVAL_UNDEF(&context.options)
	}
	if context.GetNotifier() != nil {
		PhpStreamNotificationFree(context.GetNotifier())
		context.SetNotifier(nil)
	}
	zend.Efree(context)
}
func PhpStreamContextAlloc() *core.PhpStreamContext {
	var context *core.PhpStreamContext
	context = zend.Ecalloc(1, b.SizeOf("php_stream_context"))
	context.SetNotifier(nil)
	zend.ArrayInit(&context.options)
	context.SetRes(zend.ZendRegisterResource(context, standard.PhpLeStreamContext()))
	return context
}
func PhpStreamNotificationAlloc() *PhpStreamNotifier {
	return zend.Ecalloc(1, b.SizeOf("php_stream_notifier"))
}
func PhpStreamNotificationFree(notifier *PhpStreamNotifier) {
	if notifier.GetDtor() != nil {
		notifier.GetDtor()(notifier)
	}
	zend.Efree(notifier)
}
func PhpStreamContextGetOption(context *core.PhpStreamContext, wrappername string, optionname string) *zend.Zval {
	var wrapperhash *zend.Zval
	if nil == b.Assign(&wrapperhash, zend.ZendHashStrFind(zend.Z_ARRVAL(context.GetOptions()), wrappername, strlen(wrappername))) {
		return nil
	}
	return zend.ZendHashStrFind(zend.Z_ARRVAL_P(wrapperhash), optionname, strlen(optionname))
}
func PhpStreamContextSetOption(context *core.PhpStreamContext, wrappername *byte, optionname *byte, optionvalue *zend.Zval) int {
	var wrapperhash *zend.Zval
	var category zend.Zval
	zend.SEPARATE_ARRAY(&context.options)
	wrapperhash = zend.ZendHashStrFind(zend.Z_ARRVAL(context.GetOptions()), wrappername, strlen(wrappername))
	if nil == wrapperhash {
		zend.ArrayInit(&category)
		wrapperhash = zend.ZendHashStrUpdate(zend.Z_ARRVAL(context.GetOptions()), (*byte)(wrappername), strlen(wrappername), &category)
	}
	zend.ZVAL_DEREF(optionvalue)
	zend.Z_TRY_ADDREF_P(optionvalue)
	zend.SEPARATE_ARRAY(wrapperhash)
	zend.ZendHashStrUpdate(zend.Z_ARRVAL_P(wrapperhash), optionname, strlen(optionname), optionvalue)
	return zend.SUCCESS
}
func PhpStreamDirentAlphasort(a **zend.ZendString, b **zend.ZendString) int {
	return strcoll(zend.ZSTR_VAL(*a), zend.ZSTR_VAL(*b))
}
func PhpStreamDirentAlphasortr(a **zend.ZendString, b **zend.ZendString) int {
	return strcoll(zend.ZSTR_VAL(*b), zend.ZSTR_VAL(*a))
}
func _phpStreamScandir(dirname *byte, namelist []**zend.ZendString, flags int, context *core.PhpStreamContext, compare func(a **zend.ZendString, b **zend.ZendString) int) int {
	var stream *core.PhpStream
	var sdp core.PhpStreamDirent
	var vector **zend.ZendString = nil
	var vector_size uint = 0
	var nfiles uint = 0
	if !namelist {
		return zend.FAILURE
	}
	stream = core.PhpStreamOpendir(dirname, core.REPORT_ERRORS, context)
	if stream == nil {
		return zend.FAILURE
	}
	for core.PhpStreamReaddir(stream, &sdp) != nil {
		if nfiles == vector_size {
			if vector_size == 0 {
				vector_size = 10
			} else {
				if vector_size*2 < vector_size {

					/* overflow */

					core.PhpStreamClosedir(stream)
					zend.Efree(vector)
					return zend.FAILURE
				}
				vector_size *= 2
			}
			vector = (**zend.ZendString)(zend.SafeErealloc(vector, vector_size, b.SizeOf("char *"), 0))
		}
		vector[nfiles] = zend.ZendStringInit(sdp.d_name, strlen(sdp.d_name), 0)
		nfiles++
		if vector_size < 10 || nfiles == 0 {

			/* overflow */

			core.PhpStreamClosedir(stream)
			zend.Efree(vector)
			return zend.FAILURE
		}
	}
	core.PhpStreamClosedir(stream)
	*namelist = vector
	if nfiles > 0 && compare != nil {
		qsort(*namelist, nfiles, b.SizeOf("zend_string *"), (func(any, any) int)(compare))
	}
	return nfiles
}
