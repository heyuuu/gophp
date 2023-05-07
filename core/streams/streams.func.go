package streams

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/builtin/ascii"
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/ext/standard"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/sapi/cli"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
)

func PhpFileLeStream() int       { return LeStream }
func PhpFileLePstream() int      { return LePstream }
func PhpFileLeStreamFilter() int { return LeStreamFilter }
func ZmDeactivateStreams(type_ int, module_number int) int {
	zend.EG__().PersistentList().Foreach(func(_ string, rsrc *types.ZendResource) {
		if rsrc.GetType() != LePstream {
			return
		}
		var stream *core.PhpStream = rsrc.GetPtr().(*core.PhpStream)
		stream.SetRes(nil)
		if stream.GetCtx() != nil {
			zend.ZendListDelete(stream.GetCtx())
			stream.SetCtx(nil)
		}
	})

	return types.SUCCESS
}
func PhpStreamEncloses(enclosing *core.PhpStream, enclosed *core.PhpStream) *core.PhpStream {
	var orig *core.PhpStream = enclosed.GetEnclosingStream()
	core.PhpStreamAutoCleanup(enclosed)
	enclosed.SetEnclosingStream(enclosing)
	return orig
}
func PhpStreamFromPersistentId(persistent_id string, stream **core.PhpStream) int {
	var le = zend.EG__().PersistentList().Get(persistent_id)
	if le != nil {
		if le.GetType() == LePstream {
			if stream != nil {
				var regentry *types.ZendResource = nil

				/* see if this persistent resource already has been loaded to the
				 * regular list; allowing the same resource in several entries in the
				 * regular list causes trouble (see bug #54623) */

				*stream = (*core.PhpStream)(le.GetPtr())
				for iter := zend.EG__().GetRegularList().Iterator(); iter.Valid(); iter.Next() {
					value := iter.Current()
					regentry = value.Ptr()
					if regentry.GetPtr() == le.GetPtr() {
						stream.SetRes(regentry)
						return core.PHP_STREAM_PERSISTENT_SUCCESS
					}
				}

				stream.SetRes(zend.ZendRegisterResource(*stream, LePstream))
			}
			return core.PHP_STREAM_PERSISTENT_SUCCESS
		}
		return core.PHP_STREAM_PERSISTENT_FAILURE
	}
	return core.PHP_STREAM_PERSISTENT_NOT_EXIST
}
func PhpStreamWrapperLogError(wrapper *core.PhpStreamWrapper, options int, fmt string, _ ...any) {
	var args va_list
	var buffer *byte = nil
	va_start(args, fmt)
	core.Vspprintf(&buffer, 0, fmt, args)
	va_end(args)
	if (options&core.REPORT_ERRORS) != 0 || wrapper == nil {
		core.PhpErrorDocref(nil, faults.E_WARNING, "%s", buffer)
		zend.Efree(buffer)
	} else {
		//var list *zend.ZendLlist = nil
		//if !(standard.FG__().GetWrapperErrors()) {
		//	zend.ALLOC_HASHTABLE(standard.FG__().GetWrapperErrors())
		//	standard.FG__().GetWrapperErrors().InitEx(8, WrapperListDtor)
		//} else {
		//	list = types.ZendHashStrFindPtr(standard.FG__().GetWrapperErrors(), b.CastStr((*byte)(&wrapper), b.SizeOf("wrapper")))
		//}
		//if list == nil {
		//	var new_list zend.ZendLlist
		//	new_list.Init(b.SizeOf("buffer"), WrapperErrorDtor, 0)
		//	list = types.ZendHashStrUpdateMem(standard.FG__().GetWrapperErrors(), b.CastStr((*byte)(&wrapper), b.SizeOf("wrapper")), &new_list, b.SizeOf("new_list"))
		//}

		/* append to linked list */
		//list.AddElement(&buffer)
	}
}
func PhpStreamReadToStr(stream *core.PhpStream, len_ int) *string {
	return core.PhpStreamReadStr(stream, len_)
}
func PhpStreamLocateEol(stream *core.PhpStream, buf *types.String) *byte {
	var avail int
	var cr *byte
	var lf *byte
	var eol *byte = nil
	var readptr *byte
	if buf == nil {
		readptr = (*byte)(stream.GetReadbuf() + stream.GetReadpos())
		avail = stream.GetWritepos() - stream.GetReadpos()
	} else {
		readptr = buf.GetVal()
		avail = buf.GetLen()
	}

	/* Look for EOL */

	if stream.HasFlags(core.PHP_STREAM_FLAG_DETECT_EOL) {
		cr = memchr(readptr, '\r', avail)
		lf = memchr(readptr, '\n', avail)
		if cr != nil && lf != cr+1 && !(lf != nil && lf < cr) {

			/* mac */

			stream.SetFlags(stream.GetFlags() ^ core.PHP_STREAM_FLAG_DETECT_EOL)
			stream.AddFlags(core.PHP_STREAM_FLAG_EOL_MAC)
			eol = cr
		} else if cr != nil && lf != nil && cr == lf-1 || lf != nil {

			/* dos or unix endings */

			stream.SetFlags(stream.GetFlags() ^ core.PHP_STREAM_FLAG_DETECT_EOL)
			eol = lf
		}
	} else if stream.HasFlags(core.PHP_STREAM_FLAG_EOL_MAC) {
		eol = memchr(readptr, '\r', avail)
	} else {

		/* unix (and dos) line endings */

		eol = memchr(readptr, '\n', avail)

		/* unix (and dos) line endings */

	}
	return eol
}
func STREAM_BUFFERED_AMOUNT(stream *core.PhpStream) __auto__ {
	return size_t(stream.GetWritepos() - stream.GetReadpos())
}
func _phpStreamSearchDelim(stream *core.PhpStream, maxlen int, skiplen int, delim *byte, delim_len int) *byte {
	var seek_len int

	/* set the maximum number of bytes we're allowed to read from buffer */

	seek_len = cli.MIN(STREAM_BUFFERED_AMOUNT(stream), maxlen)
	if seek_len <= skiplen {
		return nil
	}
	if delim_len == 1 {
		return memchr(stream.GetReadbuf()[stream.GetReadpos()+skiplen], delim[0], seek_len-skiplen)
	} else {
		return core.PhpMemnstr((*byte)(stream.GetReadbuf()[stream.GetReadpos()+skiplen]), delim, delim_len, (*byte)(stream.GetReadbuf()[stream.GetReadpos()+seek_len]))
	}
}
func PhpStreamGetRecord(stream *core.PhpStream, maxlen int, delim *byte, delim_len int) *types.String {
	var ret_buf *types.String
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
		to_read_now = cli.MIN(maxlen-buffered_len, stream.GetChunkSize())
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
		tent_ret_len = found_delim - (*byte)(stream.GetReadbuf()[stream.GetReadpos()])
	} else if has_delim == 0 && STREAM_BUFFERED_AMOUNT(stream) >= maxlen {
		tent_ret_len = maxlen
	} else {

		/* return with error if the delimiter string (if any) was not found, we
		 * could not completely fill the read buffer with maxlen bytes and we
		 * don't know we've reached end of file. Added with non-blocking streams
		 * in mind, where this situation is frequent */

		if STREAM_BUFFERED_AMOUNT(stream) < maxlen && stream.GetEof() == 0 {
			return nil
		} else if STREAM_BUFFERED_AMOUNT(stream) == 0 && stream.GetEof() != 0 {

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

	/* php_stream_read will not call ops->read here because the necessary
	 * data is guaranteedly buffered */
	retStr := core.PhpStreamReadStr(stream, tent_ret_len)
	if found_delim != nil {
		stream.SetReadpos(stream.GetReadpos() + delim_len)
		stream.SetPosition(stream.GetPosition() + delim_len)
	}
	return types.NewStringSafe(retStr)
}
func PhpStreamWrapperSchemeValidate(protocol string) bool {
	for _, c := range []byte(protocol) {
		if !ascii.IsAlphaNum(c) && c != '+' && c != '-' && c != '.' {
			return false
		}
	}
	return true
}
func PhpRegisterUrlStreamWrapper(protocol string, wrapper *core.PhpStreamWrapper) {
	if !PhpStreamWrapperSchemeValidate(protocol) {
		return
	}
	UrlStreamWrappersHash[protocol] = wrapper
}
func PhpUnregisterUrlStreamWrapper(protocol string) {
	delete(UrlStreamWrappersHash, protocol)
}
func CloneWrapperHash() {
	standard.FG__().SetStreamWrappers(b.CopyMap(UrlStreamWrappersHash))
}
func PhpRegisterUrlStreamWrapperVolatile(protocol *types.String, wrapper *core.PhpStreamWrapper) int {
	if !PhpStreamWrapperSchemeValidate(protocol.GetStr()) {
		return types.FAILURE
	}
	if standard.FG__().StreamWrappers() == nil {
		CloneWrapperHash()
	}
	if _, exists := standard.FG__().StreamWrappers()[protocol.GetStr()]; exists {
		return types.FAILURE
	} else {
		standard.FG__().StreamWrappers()[protocol.GetStr()] = wrapper
		return types.SUCCESS
	}
}
func PhpUnregisterUrlStreamWrapperVolatile(protocol *types.String) int {
	if standard.FG__().StreamWrappers() != nil {
		CloneWrapperHash()
	}
	if standard.FG__().StreamWrappers()[protocol.GetStr()] != nil {
		return types.FAILURE
	}
	delete(standard.FG__().StreamWrappers(), protocol.GetStr())
	return types.SUCCESS
}
func PhpStreamLocateUrlWrapper(path *byte, path_for_open **byte, options int) *core.PhpStreamWrapper {
	var wrapper_hash map[string]*core.PhpStreamWrapper
	if standard.FG__().StreamWrappers() != nil {
		wrapper_hash = standard.FG__().StreamWrappers()
	} else {
		wrapper_hash = UrlStreamWrappersHash
	}
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
		if nil == wrapper_hash[b.CastStr(protocol, n)] {
			tmp := ascii.StrToLower(b.CastStr(protocol, n))
			wrapper = wrapper_hash[tmp]
			if nil == wrapper {
				var wrapper_name []byte
				if n >= b.SizeOf("wrapper_name") {
					n = b.SizeOf("wrapper_name") - 1
				}
				core.PHP_STRLCPY(wrapper_name, protocol, b.SizeOf("wrapper_name"), n)
				core.PhpErrorDocref(nil, faults.E_WARNING, "Unable to find the wrapper \"%s\" - did you forget to enable it when you configured PHP?", wrapper_name)
				wrapper = nil
				protocol = nil
			}
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
					core.PhpErrorDocref(nil, faults.E_WARNING, "remote host file access not supported, %s", path)
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
		if standard.FG__().StreamWrappers() != nil {
			/* The file:// wrapper may have been disabled/overridden */
			if wrapper != nil {
				/* It was found so go ahead and provide it */
				return wrapper
			}

			/* Check again, the original check might have not known the protocol name */
			wrapper = wrapper_hash[types.STR_FILE]
			if wrapper != nil {
				return wrapper
			}
			if (options & core.REPORT_ERRORS) != 0 {
				core.PhpErrorDocref(nil, faults.E_WARNING, "file:// wrapper is disabled in the server configuration")
			}
			return nil
		}
		return plain_files_wrapper
	}
	if wrapper != nil && wrapper.GetIsUrl() != 0 && (options&core.STREAM_DISABLE_URL_PROTECTION) == 0 && (!(core.PG__().allow_url_fopen) || ((options&core.STREAM_OPEN_FOR_INCLUDE) != 0 || core.PG__().in_user_include) && !(core.PG__().allow_url_include)) {
		if (options & core.REPORT_ERRORS) != 0 {

			/* protocol[n] probably isn't '\0' */

			if !(core.PG__().allow_url_fopen) {
				core.PhpErrorDocref(nil, faults.E_WARNING, "%.*s:// wrapper is disabled in the server configuration by allow_url_fopen=0", int(n), protocol)
			} else {
				core.PhpErrorDocref(nil, faults.E_WARNING, "%.*s:// wrapper is disabled in the server configuration by allow_url_include=0", int(n), protocol)
			}

			/* protocol[n] probably isn't '\0' */

		}
		return nil
	}
	return wrapper
}
func PhpStreamContextSet(stream *core.PhpStream, context *core.PhpStreamContext) *core.PhpStreamContext {
	var oldcontext *core.PhpStreamContext = core.PHP_STREAM_CONTEXT(stream)
	if context != nil {
		stream.SetCtx(context.GetRes())
		//context.GetRes().AddRefcount()
	} else {
		stream.SetCtx(nil)
	}
	if oldcontext != nil {
		zend.ZendListDelete(oldcontext.GetRes())
	}
	return oldcontext
}
func PhpStreamNotificationNotify(
	context *core.PhpStreamContext,
	notifycode int,
	severity int,
	xmsg *byte,
	xcode int,
	bytes_sofar int,
	bytes_max int,
	ptr any,
) {
	if context != nil && context.GetNotifier() != nil {
		context.GetNotifier().GetFunc()(context, notifycode, severity, xmsg, xcode, bytes_sofar, bytes_max, ptr)
	}
}
func PhpStreamContextFree(context *core.PhpStreamContext) {
	if context.GetOptions().IsNotUndef() {
		// zend.ZvalPtrDtor(context.GetOptions())
		context.GetOptions().SetUndef()
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
	zend.ArrayInit(context.GetOptions())
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
func PhpStreamContextGetOption(context *core.PhpStreamContext, wrappername string, optionname string) *types.Zval {
	var wrapperhash *types.Zval
	if nil == b.Assign(&wrapperhash, context.GetOptions().Array().KeyFind(b.CastStrAuto(wrappername))) {
		return nil
	}
	return wrapperhash.Array().KeyFind(b.CastStrAuto(optionname))
}
func PhpStreamContextSetOption(context *core.PhpStreamContext, wrappername *byte, optionname *byte, optionvalue *types.Zval) int {
	var wrapperhash *types.Zval
	var category types.Zval
	types.SeparateArray(context.GetOptions())
	wrapperhash = context.GetOptions().Array().KeyFind(b.CastStrAuto(wrappername))
	if nil == wrapperhash {
		zend.ArrayInit(&category)
		wrapperhash = context.GetOptions().Array().KeyUpdate(b.CastStr((*byte)(wrappername), strlen(wrappername)), &category)
	}
	optionvalue = types.ZVAL_DEREF(optionvalue)
	//optionvalue.TryAddRefcount()
	types.SeparateArray(wrapperhash)
	wrapperhash.Array().KeyUpdate(b.CastStrAuto(optionname), optionvalue)
	return types.SUCCESS
}
func PhpStreamDirentAlphasort(i **types.String, j **types.String) int {
	return b.StrColl(i.GetVal(), j.GetVal())
}
func PhpStreamDirentAlphasortr(i **types.String, j **types.String) int {
	return b.StrColl(j.GetVal(), i.GetVal())
}
