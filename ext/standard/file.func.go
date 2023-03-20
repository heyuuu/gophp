// <<generate>>

package standard

import (
	b "sik/builtin"
	r "sik/builtin/file"
	"sik/core"
	"sik/core/streams"
	"sik/sapi/cli"
	"sik/zend"
	"sik/zend/argparse"
	"sik/zend/faults"
	"sik/zend/types"
)

func FG(v __auto__) __auto__ { return FileGlobals.v }
func PHP_STREAM_TO_ZVAL(stream *core.PhpStream, arg *types.Zval) {
	b.Assert(arg.IsType(types.IS_RESOURCE))
	core.PhpStreamFromRes(stream, arg.GetRes())
}
func PhpLeStreamContext() int { return LeStreamContext }
func FileContextDtor(res *types.ZendResource) {
	var context *core.PhpStreamContext = (*core.PhpStreamContext)(res.GetPtr())
	if context.GetOptions().GetType() != types.IS_UNDEF {
		zend.ZvalPtrDtor(context.GetOptions())
		context.GetOptions().SetUndef()
	}
	streams.PhpStreamContextFree(context)
}
func FileGlobalsCtor(file_globals_p *PhpFileGlobals) {
	memset(file_globals_p, 0, b.SizeOf("php_file_globals"))
	file_globals_p.SetDefChunkSize(core.PHP_SOCK_CHUNK_SIZE)
}
func FileGlobalsDtor(file_globals_p *PhpFileGlobals) {}
func ZmStartupFile(type_ int, module_number int) int {
	LeStreamContext = zend.ZendRegisterListDestructorsEx(FileContextDtor, nil, "stream-context", module_number)
	FileGlobalsCtor(&FileGlobals)
	zend.REGISTER_INI_ENTRIES(module_number)
	zend.RegisterLongConstant("SEEK_SET", r.SEEK_SET, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("SEEK_CUR", r.SEEK_CUR, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("SEEK_END", r.SEEK_END, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("LOCK_SH", PHP_LOCK_SH, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("LOCK_EX", PHP_LOCK_EX, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("LOCK_UN", PHP_LOCK_UN, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("LOCK_NB", PHP_LOCK_NB, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("STREAM_NOTIFY_CONNECT", streams.PHP_STREAM_NOTIFY_CONNECT, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("STREAM_NOTIFY_AUTH_REQUIRED", streams.PHP_STREAM_NOTIFY_AUTH_REQUIRED, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("STREAM_NOTIFY_AUTH_RESULT", streams.PHP_STREAM_NOTIFY_AUTH_RESULT, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("STREAM_NOTIFY_MIME_TYPE_IS", streams.PHP_STREAM_NOTIFY_MIME_TYPE_IS, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("STREAM_NOTIFY_FILE_SIZE_IS", streams.PHP_STREAM_NOTIFY_FILE_SIZE_IS, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("STREAM_NOTIFY_REDIRECTED", streams.PHP_STREAM_NOTIFY_REDIRECTED, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("STREAM_NOTIFY_PROGRESS", streams.PHP_STREAM_NOTIFY_PROGRESS, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("STREAM_NOTIFY_FAILURE", streams.PHP_STREAM_NOTIFY_FAILURE, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("STREAM_NOTIFY_COMPLETED", streams.PHP_STREAM_NOTIFY_COMPLETED, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("STREAM_NOTIFY_RESOLVE", streams.PHP_STREAM_NOTIFY_RESOLVE, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("STREAM_NOTIFY_SEVERITY_INFO", streams.PHP_STREAM_NOTIFY_SEVERITY_INFO, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("STREAM_NOTIFY_SEVERITY_WARN", streams.PHP_STREAM_NOTIFY_SEVERITY_WARN, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("STREAM_NOTIFY_SEVERITY_ERR", streams.PHP_STREAM_NOTIFY_SEVERITY_ERR, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("STREAM_FILTER_READ", streams.PHP_STREAM_FILTER_READ, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("STREAM_FILTER_WRITE", streams.PHP_STREAM_FILTER_WRITE, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("STREAM_FILTER_ALL", streams.PHP_STREAM_FILTER_ALL, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("STREAM_CLIENT_PERSISTENT", PHP_STREAM_CLIENT_PERSISTENT, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("STREAM_CLIENT_ASYNC_CONNECT", PHP_STREAM_CLIENT_ASYNC_CONNECT, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("STREAM_CLIENT_CONNECT", PHP_STREAM_CLIENT_CONNECT, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("STREAM_CRYPTO_METHOD_ANY_CLIENT", streams.STREAM_CRYPTO_METHOD_ANY_CLIENT, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("STREAM_CRYPTO_METHOD_SSLv2_CLIENT", streams.STREAM_CRYPTO_METHOD_SSLv2_CLIENT, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("STREAM_CRYPTO_METHOD_SSLv3_CLIENT", streams.STREAM_CRYPTO_METHOD_SSLv3_CLIENT, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("STREAM_CRYPTO_METHOD_SSLv23_CLIENT", streams.STREAM_CRYPTO_METHOD_SSLv23_CLIENT, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("STREAM_CRYPTO_METHOD_TLS_CLIENT", streams.STREAM_CRYPTO_METHOD_TLS_CLIENT, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("STREAM_CRYPTO_METHOD_TLSv1_0_CLIENT", streams.STREAM_CRYPTO_METHOD_TLSv1_0_CLIENT, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("STREAM_CRYPTO_METHOD_TLSv1_1_CLIENT", streams.STREAM_CRYPTO_METHOD_TLSv1_1_CLIENT, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("STREAM_CRYPTO_METHOD_TLSv1_2_CLIENT", streams.STREAM_CRYPTO_METHOD_TLSv1_2_CLIENT, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("STREAM_CRYPTO_METHOD_TLSv1_3_CLIENT", streams.STREAM_CRYPTO_METHOD_TLSv1_3_CLIENT, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("STREAM_CRYPTO_METHOD_ANY_SERVER", streams.STREAM_CRYPTO_METHOD_ANY_SERVER, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("STREAM_CRYPTO_METHOD_SSLv2_SERVER", streams.STREAM_CRYPTO_METHOD_SSLv2_SERVER, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("STREAM_CRYPTO_METHOD_SSLv3_SERVER", streams.STREAM_CRYPTO_METHOD_SSLv3_SERVER, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("STREAM_CRYPTO_METHOD_SSLv23_SERVER", streams.STREAM_CRYPTO_METHOD_SSLv23_SERVER, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("STREAM_CRYPTO_METHOD_TLS_SERVER", streams.STREAM_CRYPTO_METHOD_TLS_SERVER, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("STREAM_CRYPTO_METHOD_TLSv1_0_SERVER", streams.STREAM_CRYPTO_METHOD_TLSv1_0_SERVER, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("STREAM_CRYPTO_METHOD_TLSv1_1_SERVER", streams.STREAM_CRYPTO_METHOD_TLSv1_1_SERVER, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("STREAM_CRYPTO_METHOD_TLSv1_2_SERVER", streams.STREAM_CRYPTO_METHOD_TLSv1_2_SERVER, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("STREAM_CRYPTO_METHOD_TLSv1_3_SERVER", streams.STREAM_CRYPTO_METHOD_TLSv1_3_SERVER, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("STREAM_CRYPTO_PROTO_SSLv3", streams.STREAM_CRYPTO_METHOD_SSLv3_SERVER, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("STREAM_CRYPTO_PROTO_TLSv1_0", streams.STREAM_CRYPTO_METHOD_TLSv1_0_SERVER, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("STREAM_CRYPTO_PROTO_TLSv1_1", streams.STREAM_CRYPTO_METHOD_TLSv1_1_SERVER, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("STREAM_CRYPTO_PROTO_TLSv1_2", streams.STREAM_CRYPTO_METHOD_TLSv1_2_SERVER, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("STREAM_CRYPTO_PROTO_TLSv1_3", streams.STREAM_CRYPTO_METHOD_TLSv1_3_SERVER, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("STREAM_SHUT_RD", streams.STREAM_SHUT_RD, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("STREAM_SHUT_WR", streams.STREAM_SHUT_WR, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("STREAM_SHUT_RDWR", streams.STREAM_SHUT_RDWR, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("STREAM_SOCK_STREAM", SOCK_STREAM, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("STREAM_SOCK_DGRAM", SOCK_DGRAM, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("STREAM_PEEK", streams.STREAM_PEEK, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("STREAM_OOB", streams.STREAM_OOB, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("STREAM_SERVER_BIND", streams.STREAM_XPORT_BIND, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("STREAM_SERVER_LISTEN", streams.STREAM_XPORT_LISTEN, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("FILE_USE_INCLUDE_PATH", PHP_FILE_USE_INCLUDE_PATH, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("FILE_IGNORE_NEW_LINES", PHP_FILE_IGNORE_NEW_LINES, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("FILE_SKIP_EMPTY_LINES", PHP_FILE_SKIP_EMPTY_LINES, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("FILE_APPEND", PHP_FILE_APPEND, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("FILE_NO_DEFAULT_CONTEXT", PHP_FILE_NO_DEFAULT_CONTEXT, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("FILE_TEXT", 0, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("FILE_BINARY", 0, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("FNM_NOESCAPE", FNM_NOESCAPE, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("FNM_PATHNAME", FNM_PATHNAME, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("FNM_PERIOD", FNM_PERIOD, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	return types.SUCCESS
}
func ZmShutdownFile(type_ int, module_number int) int {
	FileGlobalsDtor(&FileGlobals)
	return types.SUCCESS
}
func ZifFlock(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var res *types.Zval
	var wouldblock *types.Zval = nil
	var act int
	var stream *core.PhpStream
	var operation zend.ZendLong = 0
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 3

		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			res = fp.ParseResource()
			operation = fp.ParseLong()
			fp.StartOptional()
			wouldblock = fp.ParseZval()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	PHP_STREAM_TO_ZVAL(stream, res)
	act = operation & 3
	if act < 1 || act > 3 {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Illegal operation argument")
		return_value.SetFalse()
		return
	}
	if wouldblock != nil {
		zend.ZEND_TRY_ASSIGN_REF_LONG(wouldblock, 0)
	}

	/* flock_values contains all possible actions if (operation & 4) we won't block on the lock */

	act = FlockValues[act-1] | b.Cond((operation&PHP_LOCK_NB) != 0, LOCK_NB, 0)
	if core.PhpStreamLock(stream, act) != 0 {
		if operation != 0 && errno == core.EWOULDBLOCK && wouldblock != nil {
			zend.ZEND_TRY_ASSIGN_REF_LONG(wouldblock, 1)
		}
		return_value.SetFalse()
		return
	}
	return_value.SetTrue()
	return
}
func ZifGetMetaTags(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var filename *byte
	var filename_len int
	var use_include_path types.ZendBool = 0
	var in_tag int = 0
	var done int = 0
	var looking_for_val int = 0
	var have_name int = 0
	var have_content int = 0
	var saw_name int = 0
	var saw_content int = 0
	var name *byte = nil
	var value *byte = nil
	var temp *byte = nil
	var tok PhpMetaTagsToken
	var tok_last PhpMetaTagsToken
	var md PhpMetaTagsData

	/* Initiailize our structure */

	memset(&md, 0, b.SizeOf("md"))

	/* Parse arguments */

	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 2

		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			filename, filename_len = fp.ParsePath()
			fp.StartOptional()
			use_include_path = fp.ParseBool()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	md.SetStream(core.PhpStreamOpenWrapper(filename, "rb", b.Cond(use_include_path != 0, core.USE_PATH, 0)|core.REPORT_ERRORS, nil))
	if md.GetStream() == nil {
		return_value.SetFalse()
		return
	}
	zend.ArrayInit(return_value)
	tok_last = TOK_EOF
	for done == 0 && b.Assign(&tok, PhpNextMetaToken(&md)) != TOK_EOF {
		if tok == TOK_ID {
			if tok_last == TOK_OPENTAG {
				md.SetInMeta(!(strcasecmp("meta", md.GetTokenData())))
			} else if tok_last == TOK_SLASH && in_tag != 0 {
				if strcasecmp("head", md.GetTokenData()) == 0 {

					/* We are done here! */

					done = 1

					/* We are done here! */

				}
			} else if tok_last == TOK_EQUAL && looking_for_val != 0 {
				if saw_name != 0 {
					if name != nil {
						zend.Efree(name)
					}

					/* Get the NAME attr (Single word attr, non-quoted) */

					name = zend.Estrndup(md.GetTokenData(), md.GetTokenLen())
					temp = name
					for temp != nil && (*temp) {
						if strchr(PHP_META_UNSAFE, *temp) {
							*temp = '_'
						}
						temp++
					}
					have_name = 1
				} else if saw_content != 0 {
					if value != nil {
						zend.Efree(value)
					}
					value = zend.Estrndup(md.GetTokenData(), md.GetTokenLen())
					have_content = 1
				}
				looking_for_val = 0
			} else {
				if md.GetInMeta() != 0 {
					if strcasecmp("name", md.GetTokenData()) == 0 {
						saw_name = 1
						saw_content = 0
						looking_for_val = 1
					} else if strcasecmp("content", md.GetTokenData()) == 0 {
						saw_name = 0
						saw_content = 1
						looking_for_val = 1
					}
				}
			}
		} else if tok == TOK_STRING && tok_last == TOK_EQUAL && looking_for_val != 0 {
			if saw_name != 0 {
				if name != nil {
					zend.Efree(name)
				}

				/* Get the NAME attr (Quoted single/double) */

				name = zend.Estrndup(md.GetTokenData(), md.GetTokenLen())
				temp = name
				for temp != nil && (*temp) {
					if strchr(PHP_META_UNSAFE, *temp) {
						*temp = '_'
					}
					temp++
				}
				have_name = 1
			} else if saw_content != 0 {
				if value != nil {
					zend.Efree(value)
				}
				value = zend.Estrndup(md.GetTokenData(), md.GetTokenLen())
				have_content = 1
			}
			looking_for_val = 0
		} else if tok == TOK_OPENTAG {
			if looking_for_val != 0 {
				looking_for_val = 0
				saw_name = 0
				have_name = saw_name
				saw_content = 0
				have_content = saw_content
			}
			in_tag = 1
		} else if tok == TOK_CLOSETAG {
			if have_name != 0 {

				/* For BC */

				PhpStrtolower(name, strlen(name))
				if have_content != 0 {
					zend.AddAssocString(return_value, name, value)
				} else {
					zend.AddAssocString(return_value, name, "")
				}
				zend.Efree(name)
				if value != nil {
					zend.Efree(value)
				}
			} else if have_content != 0 {
				zend.Efree(value)
			}
			value = nil
			name = value

			/* Reset all of our flags */

			looking_for_val = 0
			in_tag = looking_for_val
			saw_name = 0
			have_name = saw_name
			saw_content = 0
			have_content = saw_content
			md.SetInMeta(0)
		}
		tok_last = tok
		if md.GetTokenData() != nil {
			zend.Efree(md.GetTokenData())
		}
		md.SetTokenData(nil)
	}
	if value != nil {
		zend.Efree(value)
	}
	if name != nil {
		zend.Efree(name)
	}
	core.PhpStreamClose(md.GetStream())
}
func ZifFileGetContents(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var filename *byte
	var filename_len int
	var use_include_path types.ZendBool = 0
	var stream *core.PhpStream
	var offset zend.ZendLong = 0
	var maxlen zend.ZendLong = ssize_t(core.PHP_STREAM_COPY_ALL)
	var zcontext *types.Zval = nil
	var context *core.PhpStreamContext = nil
	var contents *types.ZendString

	/* Parse arguments */

	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 5

		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			filename, filename_len = fp.ParsePath()
			fp.StartOptional()
			use_include_path = fp.ParseBool()
			zcontext = fp.ParseResourceEx(true)
			offset = fp.ParseLong()
			maxlen = fp.ParseLong()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	if executeData.NumArgs() == 5 && maxlen < 0 {
		core.PhpErrorDocref(nil, faults.E_WARNING, "length must be greater than or equal to zero")
		return_value.SetFalse()
		return
	}
	context = streams.PhpStreamContextFromZval(zcontext, 0)
	stream = core.PhpStreamOpenWrapperEx(filename, "rb", b.Cond(use_include_path != 0, core.USE_PATH, 0)|core.REPORT_ERRORS, nil, context)
	if stream == nil {
		return_value.SetFalse()
		return
	}
	if offset != 0 && core.PhpStreamSeek(stream, offset, b.Cond(offset > 0, r.SEEK_SET, r.SEEK_END)) < 0 {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Failed to seek to position "+zend.ZEND_LONG_FMT+" in the stream", offset)
		core.PhpStreamClose(stream)
		return_value.SetFalse()
		return
	}
	if b.Assign(&contents, core.PhpStreamCopyToMem(stream, maxlen, 0)) != nil {
		return_value.SetString(contents)
	} else {
		zend.ZVAL_EMPTY_STRING(return_value)
	}
	core.PhpStreamClose(stream)
}
func ZifFilePutContents(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var stream *core.PhpStream
	var filename *byte
	var filename_len int
	var data *types.Zval
	var numbytes ssize_t = 0
	var flags zend.ZendLong = 0
	var zcontext *types.Zval = nil
	var context *core.PhpStreamContext = nil
	var srcstream *core.PhpStream = nil
	var mode []byte = "wb"
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 4

		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			filename, filename_len = fp.ParsePath()
			data = fp.ParseZval()
			fp.StartOptional()
			flags = fp.ParseLong()
			zcontext = fp.ParseResourceEx(true)
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	if data.IsType(types.IS_RESOURCE) {
		core.PhpStreamFromZval(srcstream, data)
	}
	context = streams.PhpStreamContextFromZval(zcontext, flags&PHP_FILE_NO_DEFAULT_CONTEXT)
	if (flags & PHP_FILE_APPEND) != 0 {
		mode[0] = 'a'
	} else if (flags & LOCK_EX) != 0 {

		/* check to make sure we are dealing with a regular file */

		if core.PhpMemnstr(filename, "://", b.SizeOf("\"://\"")-1, filename+filename_len) {
			if strncasecmp(filename, "file://", b.SizeOf("\"file://\"")-1) {
				core.PhpErrorDocref(nil, faults.E_WARNING, "Exclusive locks may only be set for regular files")
				return_value.SetFalse()
				return
			}
		}
		mode[0] = 'c'
	}
	mode[2] = '0'
	stream = core.PhpStreamOpenWrapperEx(filename, mode, b.Cond((flags&PHP_FILE_USE_INCLUDE_PATH) != 0, core.USE_PATH, 0)|core.REPORT_ERRORS, nil, context)
	if stream == nil {
		return_value.SetFalse()
		return
	}
	if (flags&LOCK_EX) != 0 && (core.PhpStreamSupportsLock(stream) == 0 || core.PhpStreamLock(stream, LOCK_EX) != 0) {
		core.PhpStreamClose(stream)
		core.PhpErrorDocref(nil, faults.E_WARNING, "Exclusive locks are not supported for this stream")
		return_value.SetFalse()
		return
	}
	if mode[0] == 'c' {
		core.PhpStreamTruncateSetSize(stream, 0)
	}
	switch data.GetType() {
	case types.IS_RESOURCE:
		var len_ int
		if core.PhpStreamCopyToStreamEx(srcstream, stream, core.PHP_STREAM_COPY_ALL, &len_) != types.SUCCESS {
			numbytes = -1
		} else {
			if len_ > zend.ZEND_LONG_MAX {
				core.PhpErrorDocref(nil, faults.E_WARNING, "content truncated from %zu to "+zend.ZEND_LONG_FMT+" bytes", len_, zend.ZEND_LONG_MAX)
				len_ = zend.ZEND_LONG_MAX
			}
			numbytes = len_
		}
	case types.IS_NULL:
		fallthrough
	case types.IS_LONG:
		fallthrough
	case types.IS_DOUBLE:
		fallthrough
	case types.IS_FALSE:
		fallthrough
	case types.IS_TRUE:
		zend.ConvertToStringEx(data)
		fallthrough
	case types.IS_STRING:
		if data.GetStr().GetLen() != 0 {
			numbytes = core.PhpStreamWrite(stream, data.GetStr().GetVal(), data.GetStr().GetLen())
			if numbytes != data.GetStr().GetLen() {
				core.PhpErrorDocref(nil, faults.E_WARNING, "Only %zd of %zd bytes written, possibly out of free disk space", numbytes, data.GetStr().GetLen())
				numbytes = -1
			}
		}
	case types.IS_ARRAY:
		if types.Z_ARRVAL_P(data).GetNNumOfElements() {
			var bytes_written ssize_t
			var tmp *types.Zval
			var __ht *types.HashTable = data.GetArr()
			for _, _p := range __ht.foreachData() {
				var _z *types.Zval = _p.GetVal()

				tmp = _z
				var t *types.ZendString
				var str *types.ZendString = zend.ZvalGetTmpString(tmp, &t)
				if str.GetLen() != 0 {
					numbytes += str.GetLen()
					bytes_written = core.PhpStreamWrite(stream, str.GetVal(), str.GetLen())
					if bytes_written != str.GetLen() {
						core.PhpErrorDocref(nil, faults.E_WARNING, "Failed to write %zd bytes to %s", str.GetLen(), filename)
						zend.ZendTmpStringRelease(t)
						numbytes = -1
						break
					}
				}
				zend.ZendTmpStringRelease(t)
			}
		}
	case types.IS_OBJECT:
		if types.Z_OBJ_HT_P(data) != nil {
			var out types.Zval
			if zend.ZendStdCastObjectTostring(data, &out, types.IS_STRING) == types.SUCCESS {
				numbytes = core.PhpStreamWrite(stream, out.GetStr().GetVal(), out.GetStr().GetLen())
				if numbytes != out.GetStr().GetLen() {
					core.PhpErrorDocref(nil, faults.E_WARNING, "Only %zd of %zd bytes written, possibly out of free disk space", numbytes, out.GetStr().GetLen())
					numbytes = -1
				}
				zend.ZvalPtrDtorStr(&out)
				break
			}
		}
		fallthrough
	default:
		numbytes = -1
	}
	core.PhpStreamClose(stream)
	if numbytes < 0 {
		return_value.SetFalse()
		return
	}
	return_value.SetLong(numbytes)
	return
}
func ZifFile(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var filename *byte
	var filename_len int
	var p *byte
	var s *byte
	var e *byte
	var i int = 0
	var eol_marker byte = '\n'
	var flags zend.ZendLong = 0
	var use_include_path types.ZendBool
	var include_new_line types.ZendBool
	var skip_blank_lines types.ZendBool
	var stream *core.PhpStream
	var zcontext *types.Zval = nil
	var context *core.PhpStreamContext = nil
	var target_buf *types.ZendString

	/* Parse arguments */

	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 3

		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			filename, filename_len = fp.ParsePath()
			fp.StartOptional()
			flags = fp.ParseLong()
			zcontext = fp.ParseResourceEx(true)
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	if flags < 0 || flags > (PHP_FILE_USE_INCLUDE_PATH|PHP_FILE_IGNORE_NEW_LINES|PHP_FILE_SKIP_EMPTY_LINES|PHP_FILE_NO_DEFAULT_CONTEXT) {
		core.PhpErrorDocref(nil, faults.E_WARNING, "'"+zend.ZEND_LONG_FMT+"' flag is not supported", flags)
		return_value.SetFalse()
		return
	}
	use_include_path = flags & PHP_FILE_USE_INCLUDE_PATH
	include_new_line = !(flags & PHP_FILE_IGNORE_NEW_LINES)
	skip_blank_lines = flags & PHP_FILE_SKIP_EMPTY_LINES
	context = streams.PhpStreamContextFromZval(zcontext, flags&PHP_FILE_NO_DEFAULT_CONTEXT)
	stream = core.PhpStreamOpenWrapperEx(filename, "rb", b.Cond(use_include_path != 0, core.USE_PATH, 0)|core.REPORT_ERRORS, nil, context)
	if stream == nil {
		return_value.SetFalse()
		return
	}

	/* Initialize return array */

	zend.ArrayInit(return_value)
	if b.Assign(&target_buf, core.PhpStreamCopyToMem(stream, core.PHP_STREAM_COPY_ALL, 0)) != nil {
		s = target_buf.GetVal()
		e = target_buf.GetVal() + target_buf.GetLen()
		if !(b.Assign(&p, (*byte)(streams.PhpStreamLocateEol(stream, target_buf)))) {
			p = e
			goto parse_eol
		}
		if stream.HasFlags(core.PHP_STREAM_FLAG_EOL_MAC) {
			eol_marker = '\r'
		}

		/* for performance reasons the code is duplicated, so that the if (include_new_line)
		 * will not need to be done for every single line in the file. */

		if include_new_line != 0 {
			for {
				p++
			parse_eol:
				zend.AddIndexStringl(return_value, b.PostInc(&i), s, p-s)
				s = p
				if !(b.Assign(&p, memchr(p, eol_marker, e-p))) {
					break
				}
			}
		} else {
			for {
				var windows_eol int = 0
				if p != target_buf.GetVal() && eol_marker == '\n' && (*(p - 1)) == '\r' {
					windows_eol++
				}
				if skip_blank_lines != 0 && p-s-windows_eol == 0 {
					p++
					s = p
					continue
				}
				zend.AddIndexStringl(return_value, b.PostInc(&i), s, p-s-windows_eol)
				p++
				s = p
				if !(b.Assign(&p, memchr(p, eol_marker, e-p))) {
					break
				}
			}
		}

		/* handle any left overs of files without new lines */

		if s != e {
			p = e
			goto parse_eol
		}

		/* handle any left overs of files without new lines */

	}
	if target_buf != nil {
		types.ZendStringFree(target_buf)
	}
	core.PhpStreamClose(stream)
}
func ZifTempnam(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var dir *byte
	var prefix *byte
	var dir_len int
	var prefix_len int
	var opened_path *types.ZendString
	var fd int
	var p *types.ZendString
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 2

		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			dir, dir_len = fp.ParsePath()
			prefix, prefix_len = fp.ParsePath()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	p = PhpBasename(prefix, prefix_len, nil, 0)
	if p.GetLen() > 64 {
		p.GetVal()[63] = '0'
	}
	return_value.SetFalse()
	if b.Assign(&fd, core.PhpOpenTemporaryFdEx(dir, p.GetVal(), &opened_path, core.PHP_TMP_FILE_OPEN_BASEDIR_CHECK_ALWAYS)) >= 0 {
		close(fd)
		return_value.SetString(opened_path)
	}
	types.ZendStringReleaseEx(p, 0)
}
func PhpIfTmpfile(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var stream *core.PhpStream
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	stream = streams._phpStreamFopenTmpfile(0)
	if stream != nil {
		core.PhpStreamToZval(stream, return_value)
	} else {
		return_value.SetFalse()
		return
	}
}
func PhpIfFopen(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var filename *byte
	var mode *byte
	var filename_len int
	var mode_len int
	var use_include_path types.ZendBool = 0
	var zcontext *types.Zval = nil
	var stream *core.PhpStream
	var context *core.PhpStreamContext = nil
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 4

		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			filename, filename_len = fp.ParsePath()
			mode, mode_len = fp.ParseString()
			fp.StartOptional()
			use_include_path = fp.ParseBool()
			zcontext = fp.ParseResourceEx(true)
			if fp.HasError() {
				fp.HandleError()
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	context = streams.PhpStreamContextFromZval(zcontext, 0)
	stream = core.PhpStreamOpenWrapperEx(filename, mode, b.Cond(use_include_path != 0, core.USE_PATH, 0)|core.REPORT_ERRORS, nil, context)
	if stream == nil {
		return_value.SetFalse()
		return
	}
	core.PhpStreamToZval(stream, return_value)
}
func ZifFclose(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var res *types.Zval
	var stream *core.PhpStream
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1

		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			res = fp.ParseResource()
			if fp.HasError() {
				fp.HandleError()
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	PHP_STREAM_TO_ZVAL(stream, res)
	if stream.HasFlags(core.PHP_STREAM_FLAG_NO_FCLOSE) {
		core.PhpErrorDocref(nil, faults.E_WARNING, "%d is not a valid stream resource", stream.GetRes().GetHandle())
		return_value.SetFalse()
		return
	}
	core.PhpStreamFree(stream, core.PHP_STREAM_FREE_KEEP_RSRC|b.Cond(stream.GetIsPersistent() != 0, core.PHP_STREAM_FREE_CLOSE_PERSISTENT, core.PHP_STREAM_FREE_CLOSE))
	return_value.SetTrue()
	return
}
func ZifPopen(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var command *byte
	var mode *byte
	var command_len int
	var mode_len int
	var fp *r.FILE
	var stream *core.PhpStream
	var posix_mode *byte
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 2

		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			command, command_len = fp.ParsePath()
			mode, mode_len = fp.ParseString()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	posix_mode = zend.Estrndup(mode, mode_len)
	var z *byte = memchr(posix_mode, 'b', mode_len)
	if z != nil {
		memmove(z, z+1, mode_len-(z-posix_mode))
	}
	fp = zend.VCWD_POPEN(command, posix_mode)
	if fp == nil {
		core.PhpErrorDocref2(nil, command, posix_mode, faults.E_WARNING, "%s", strerror(errno))
		zend.Efree(posix_mode)
		return_value.SetFalse()
		return
	}
	stream = streams.PhpStreamFopenFromPipe(fp, mode)
	if stream == nil {
		core.PhpErrorDocref2(nil, command, mode, faults.E_WARNING, "%s", strerror(errno))
		return_value.SetFalse()
	} else {
		core.PhpStreamToZval(stream, return_value)
	}
	zend.Efree(posix_mode)
}
func ZifPclose(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var res *types.Zval
	var stream *core.PhpStream
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1

		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			res = fp.ParseResource()
			if fp.HasError() {
				fp.HandleError()
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	PHP_STREAM_TO_ZVAL(stream, res)
	FG(pclose_wait) = 1
	zend.ZendListClose(stream.GetRes())
	FG(pclose_wait) = 0
	return_value.SetLong(FG(pclose_ret))
	return
}
func ZifFeof(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var res *types.Zval
	var stream *core.PhpStream
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1

		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			res = fp.ParseResource()
			if fp.HasError() {
				fp.HandleError()
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	PHP_STREAM_TO_ZVAL(stream, res)
	if core.PhpStreamEof(stream) != 0 {
		return_value.SetTrue()
		return
	} else {
		return_value.SetFalse()
		return
	}
}
func ZifFgets(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var res *types.Zval
	var len_ zend.ZendLong = 1024
	var buf *byte = nil
	var argc int = executeData.NumArgs()
	var line_len int = 0
	var str *types.ZendString
	var stream *core.PhpStream
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 2

		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			res = fp.ParseResource()
			fp.StartOptional()
			len_ = fp.ParseLong()
			if fp.HasError() {
				fp.HandleError()
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	PHP_STREAM_TO_ZVAL(stream, res)
	if argc == 1 {

		/* ask streams to give us a buffer of an appropriate size */

		buf = core.PhpStreamGetLine(stream, nil, 0, &line_len)
		if buf == nil {
			return_value.SetFalse()
			return
		}

		// TODO: avoid reallocation ???

		return_value.SetRawString(b.CastStr(buf, line_len))
		zend.Efree(buf)
	} else if argc > 1 {
		if len_ <= 0 {
			core.PhpErrorDocref(nil, faults.E_WARNING, "Length parameter must be greater than 0")
			return_value.SetFalse()
			return
		}
		str = types.ZendStringAlloc(len_, 0)
		if core.PhpStreamGetLine(stream, str.GetVal(), len_, &line_len) == nil {
			types.ZendStringEfree(str)
			return_value.SetFalse()
			return
		}

		/* resize buffer if it's much larger than the result.
		 * Only needed if the user requested a buffer size. */

		if line_len < int(len_/2) {
			str = types.ZendStringTruncate(str, line_len, 0)
		} else {
			str.SetLen(line_len)
		}
		return_value.SetString(str)
		return
	}
}
func ZifFgetc(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var res *types.Zval
	var buf []byte
	var result int
	var stream *core.PhpStream
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1

		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			res = fp.ParseResource()
			if fp.HasError() {
				fp.HandleError()
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	PHP_STREAM_TO_ZVAL(stream, res)
	result = core.PhpStreamGetc(stream)
	if result == r.EOF {
		return_value.SetFalse()
	} else {
		buf[0] = result
		buf[1] = '0'
		return_value.SetRawString(b.CastStr(buf, 1))
		return
	}
}
func ZifFgetss(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var fd *types.Zval
	var bytes zend.ZendLong = 0
	var len_ int = 0
	var actual_len int
	var retval_len int
	var buf *byte = nil
	var retval *byte
	var stream *core.PhpStream
	var allowed_tags *byte = nil
	var allowed_tags_len int = 0
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 3

		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			fd = fp.ParseResource()
			fp.StartOptional()
			bytes = fp.ParseLong()
			allowed_tags, allowed_tags_len = fp.ParseString()
			if fp.HasError() {
				fp.HandleError()
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	PHP_STREAM_TO_ZVAL(stream, fd)
	if executeData.NumArgs() >= 2 {
		if bytes <= 0 {
			core.PhpErrorDocref(nil, faults.E_WARNING, "Length parameter must be greater than 0")
			return_value.SetFalse()
			return
		}
		len_ = int(bytes)
		buf = zend.SafeEmalloc(b.SizeOf("char"), len_+1, 0)

		/*needed because recv doesn't set null char at end*/

		memset(buf, 0, len_+1)

		/*needed because recv doesn't set null char at end*/

	}
	if b.Assign(&retval, core.PhpStreamGetLine(stream, buf, len_, &actual_len)) == nil {
		if buf != nil {
			zend.Efree(buf)
		}
		return_value.SetFalse()
		return
	}
	retval_len = PhpStripTags(retval, actual_len, stream.GetFgetssState(), allowed_tags, allowed_tags_len)

	// TODO: avoid reallocation ???

	return_value.SetRawString(b.CastStr(retval, retval_len))
	zend.Efree(retval)
}
func ZifFscanf(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var result int
	var argc int = 0
	var format_len int
	var args *types.Zval = nil
	var file_handle *types.Zval
	var buf *byte
	var format *byte
	var len_ int
	var what any
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = -1

		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			file_handle = fp.ParseResource()
			format, format_len = fp.ParseString()
			args, argc = fp.ParseVariadic0()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	what = zend.ZendFetchResource2(file_handle.GetRes(), "File-Handle", streams.PhpFileLeStream(), streams.PhpFileLePstream())

	/* we can't do a ZEND_VERIFY_RESOURCE(what), otherwise we end up
	 * with a leak if we have an invalid filehandle. This needs changing
	 * if the code behind ZEND_VERIFY_RESOURCE changed. - cc */

	if !what {
		return_value.SetFalse()
		return
	}
	buf = core.PhpStreamGetLine((*core.PhpStream)(what), nil, 0, &len_)
	if buf == nil {
		return_value.SetFalse()
		return
	}
	result = PhpSscanfInternal(buf, format, argc, args, 0, return_value)
	zend.Efree(buf)
	if SCAN_ERROR_WRONG_PARAM_COUNT == result {
		zend.ZendWrongParamCount()
		return
	}
}
func ZifFwrite(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var res *types.Zval
	var input *byte
	var inputlen int
	var ret ssize_t
	var num_bytes int
	var maxlen zend.ZendLong = 0
	var stream *core.PhpStream
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 3

		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			res = fp.ParseResource()
			input, inputlen = fp.ParseString()
			fp.StartOptional()
			maxlen = fp.ParseLong()
			if fp.HasError() {
				fp.HandleError()
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	if executeData.NumArgs() == 2 {
		num_bytes = inputlen
	} else if maxlen <= 0 {
		num_bytes = 0
	} else {
		num_bytes = cli.MIN(int(maxlen), inputlen)
	}
	if num_bytes == 0 {
		return_value.SetLong(0)
		return
	}
	PHP_STREAM_TO_ZVAL(stream, res)
	ret = core.PhpStreamWrite(stream, input, num_bytes)
	if ret < 0 {
		return_value.SetFalse()
		return
	}
	return_value.SetLong(ret)
	return
}
func ZifFflush(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var res *types.Zval
	var ret int
	var stream *core.PhpStream
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1

		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			res = fp.ParseResource()
			if fp.HasError() {
				fp.HandleError()
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	PHP_STREAM_TO_ZVAL(stream, res)
	ret = core.PhpStreamFlush(stream)
	if ret != 0 {
		return_value.SetFalse()
		return
	}
	return_value.SetTrue()
	return
}
func ZifRewind(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var res *types.Zval
	var stream *core.PhpStream
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1

		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			res = fp.ParseResource()
			if fp.HasError() {
				fp.HandleError()
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	PHP_STREAM_TO_ZVAL(stream, res)
	if -1 == core.PhpStreamRewind(stream) {
		return_value.SetFalse()
		return
	}
	return_value.SetTrue()
	return
}
func ZifFtell(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var res *types.Zval
	var ret zend.ZendLong
	var stream *core.PhpStream
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1

		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			res = fp.ParseResource()
			if fp.HasError() {
				fp.HandleError()
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	PHP_STREAM_TO_ZVAL(stream, res)
	ret = stream.GetPosition()
	if ret == -1 {
		return_value.SetFalse()
		return
	}
	return_value.SetLong(ret)
	return
}
func ZifFseek(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var res *types.Zval
	var offset zend.ZendLong
	var whence zend.ZendLong = r.SEEK_SET
	var stream *core.PhpStream
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 3

		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			res = fp.ParseResource()
			offset = fp.ParseLong()
			fp.StartOptional()
			whence = fp.ParseLong()
			if fp.HasError() {
				fp.HandleError()
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	PHP_STREAM_TO_ZVAL(stream, res)
	return_value.SetLong(core.PhpStreamSeek(stream, offset, int(whence)))
	return
}
func PhpMkdirEx(dir *byte, mode zend.ZendLong, options int) int {
	var ret int
	if core.PhpCheckOpenBasedir(dir) != 0 {
		return -1
	}
	if b.Assign(&ret, zend.VCWD_MKDIR(dir, mode_t(mode))) < 0 && (options&core.REPORT_ERRORS) != 0 {
		core.PhpErrorDocref(nil, faults.E_WARNING, "%s", strerror(errno))
	}
	return ret
}
func PhpMkdir(dir *byte, mode zend.ZendLong) int {
	return PhpMkdirEx(dir, mode, core.REPORT_ERRORS)
}
func ZifMkdir(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var dir *byte
	var dir_len int
	var zcontext *types.Zval = nil
	var mode zend.ZendLong = 0777
	var recursive types.ZendBool = 0
	var context *core.PhpStreamContext
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 4

		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			dir, dir_len = fp.ParsePath()
			fp.StartOptional()
			mode = fp.ParseLong()
			recursive = fp.ParseBool()
			zcontext = fp.ParseResourceEx(true)
			if fp.HasError() {
				fp.HandleError()
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	context = streams.PhpStreamContextFromZval(zcontext, 0)
	types.ZVAL_BOOL(return_value, core.PhpStreamMkdir(dir, int(mode), b.Cond(recursive != 0, core.PHP_STREAM_MKDIR_RECURSIVE, 0)|core.REPORT_ERRORS, context) != 0)
	return
}
func ZifRmdir(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var dir *byte
	var dir_len int
	var zcontext *types.Zval = nil
	var context *core.PhpStreamContext
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 2

		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			dir, dir_len = fp.ParsePath()
			fp.StartOptional()
			zcontext = fp.ParseResourceEx(true)
			if fp.HasError() {
				fp.HandleError()
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	context = streams.PhpStreamContextFromZval(zcontext, 0)
	types.ZVAL_BOOL(return_value, core.PhpStreamRmdir(dir, core.REPORT_ERRORS, context) != 0)
	return
}
func ZifReadfile(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var filename *byte
	var filename_len int
	var size int = 0
	var use_include_path types.ZendBool = 0
	var zcontext *types.Zval = nil
	var stream *core.PhpStream
	var context *core.PhpStreamContext = nil
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 3

		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			filename, filename_len = fp.ParsePath()
			fp.StartOptional()
			use_include_path = fp.ParseBool()
			zcontext = fp.ParseResourceEx(true)
			if fp.HasError() {
				fp.HandleError()
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	context = streams.PhpStreamContextFromZval(zcontext, 0)
	stream = core.PhpStreamOpenWrapperEx(filename, "rb", b.Cond(use_include_path != 0, core.USE_PATH, 0)|core.REPORT_ERRORS, nil, context)
	if stream != nil {
		size = core.PhpStreamPassthru(stream)
		core.PhpStreamClose(stream)
		return_value.SetLong(size)
		return
	}
	return_value.SetFalse()
	return
}
func ZifUmask(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var mask zend.ZendLong = 0
	var oldumask int
	oldumask = umask(077)
	if BG__().umask == -1 {
		BG__().umask = oldumask
	}
	for {
		var _flags int = 0
		var _min_num_args int = 0
		var _max_num_args int = 1

		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			fp.StartOptional()
			mask = fp.ParseLong()
			if fp.HasError() {
				fp.HandleError()
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	if executeData.NumArgs() == 0 {
		umask(oldumask)
	} else {
		umask(int(mask))
	}
	return_value.SetLong(oldumask)
	return
}
func ZifFpassthru(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var res *types.Zval
	var size int
	var stream *core.PhpStream
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1

		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			res = fp.ParseResource()
			if fp.HasError() {
				fp.HandleError()
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	PHP_STREAM_TO_ZVAL(stream, res)
	size = core.PhpStreamPassthru(stream)
	return_value.SetLong(size)
	return
}
func ZifRename(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var old_name *byte
	var new_name *byte
	var old_name_len int
	var new_name_len int
	var zcontext *types.Zval = nil
	var wrapper *core.PhpStreamWrapper
	var context *core.PhpStreamContext
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 3

		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			old_name, old_name_len = fp.ParsePath()
			new_name, new_name_len = fp.ParsePath()
			fp.StartOptional()
			zcontext = fp.ParseResourceEx(true)
			if fp.HasError() {
				fp.HandleError()
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	wrapper = streams.PhpStreamLocateUrlWrapper(old_name, nil, 0)
	if wrapper == nil || wrapper.GetWops() == nil {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Unable to locate stream wrapper")
		return_value.SetFalse()
		return
	}
	if wrapper.GetWops().GetRename() == nil {
		core.PhpErrorDocref(nil, faults.E_WARNING, "%s wrapper does not support renaming", b.CondF1(wrapper.GetWops().GetLabel() != nil, func() *byte { return wrapper.GetWops().GetLabel() }, "Source"))
		return_value.SetFalse()
		return
	}
	if wrapper != streams.PhpStreamLocateUrlWrapper(new_name, nil, 0) {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Cannot rename a file across wrapper types")
		return_value.SetFalse()
		return
	}
	context = streams.PhpStreamContextFromZval(zcontext, 0)
	types.ZVAL_BOOL(return_value, wrapper.GetWops().GetRename()(wrapper, old_name, new_name, 0, context) != 0)
	return
}
func ZifUnlink(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var filename *byte
	var filename_len int
	var wrapper *core.PhpStreamWrapper
	var zcontext *types.Zval = nil
	var context *core.PhpStreamContext = nil
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 2

		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			filename, filename_len = fp.ParsePath()
			fp.StartOptional()
			zcontext = fp.ParseResourceEx(true)
			if fp.HasError() {
				fp.HandleError()
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	context = streams.PhpStreamContextFromZval(zcontext, 0)
	wrapper = streams.PhpStreamLocateUrlWrapper(filename, nil, 0)
	if wrapper == nil || wrapper.GetWops() == nil {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Unable to locate stream wrapper")
		return_value.SetFalse()
		return
	}
	if wrapper.GetWops().GetUnlink() == nil {
		core.PhpErrorDocref(nil, faults.E_WARNING, "%s does not allow unlinking", b.CondF1(wrapper.GetWops().GetLabel() != nil, func() *byte { return wrapper.GetWops().GetLabel() }, "Wrapper"))
		return_value.SetFalse()
		return
	}
	types.ZVAL_BOOL(return_value, wrapper.GetWops().GetUnlink()(wrapper, filename, core.REPORT_ERRORS, context) != 0)
	return
}
func PhpIfFtruncate(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var fp *types.Zval
	var size zend.ZendLong
	var stream *core.PhpStream
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 2

		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			fp = fp.ParseResource()
			size = fp.ParseLong()
			if fp.HasError() {
				fp.HandleError()
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	if size < 0 {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Negative size is not supported")
		return_value.SetFalse()
		return
	}
	PHP_STREAM_TO_ZVAL(stream, fp)
	if core.PhpStreamTruncateSupported(stream) == 0 {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Can't truncate this stream!")
		return_value.SetFalse()
		return
	}
	types.ZVAL_BOOL(return_value, 0 == core.PhpStreamTruncateSetSize(stream, size))
	return
}
func PhpIfFstat(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var fp *types.Zval
	var stat_dev types.Zval
	var stat_ino types.Zval
	var stat_mode types.Zval
	var stat_nlink types.Zval
	var stat_uid types.Zval
	var stat_gid types.Zval
	var stat_rdev types.Zval
	var stat_size types.Zval
	var stat_atime types.Zval
	var stat_mtime types.Zval
	var stat_ctime types.Zval
	var stat_blksize types.Zval
	var stat_blocks types.Zval
	var stream *core.PhpStream
	var stat_ssb core.PhpStreamStatbuf
	var stat_sb_names []*byte = []*byte{"dev", "ino", "mode", "nlink", "uid", "gid", "rdev", "size", "atime", "mtime", "ctime", "blksize", "blocks"}
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1

		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			fp = fp.ParseResource()
			if fp.HasError() {
				fp.HandleError()
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	PHP_STREAM_TO_ZVAL(stream, fp)
	if core.PhpStreamStat(stream, &stat_ssb) != 0 {
		return_value.SetFalse()
		return
	}
	zend.ArrayInit(return_value)
	stat_dev.SetLong(stat_ssb.GetSb().st_dev)
	stat_ino.SetLong(stat_ssb.GetSb().st_ino)
	stat_mode.SetLong(stat_ssb.GetSb().st_mode)
	stat_nlink.SetLong(stat_ssb.GetSb().st_nlink)
	stat_uid.SetLong(stat_ssb.GetSb().st_uid)
	stat_gid.SetLong(stat_ssb.GetSb().st_gid)
	stat_rdev.SetLong(stat_ssb.GetSb().st_rdev)
	stat_size.SetLong(stat_ssb.GetSb().st_size)
	stat_atime.SetLong(stat_ssb.GetSb().st_atime)
	stat_mtime.SetLong(stat_ssb.GetSb().st_mtime)
	stat_ctime.SetLong(stat_ssb.GetSb().st_ctime)
	stat_blksize.SetLong(stat_ssb.GetSb().st_blksize)
	stat_blocks.SetLong(stat_ssb.GetSb().st_blocks)

	/* Store numeric indexes in proper order */

	return_value.GetArr().NextIndexInsert(&stat_dev)
	return_value.GetArr().NextIndexInsert(&stat_ino)
	return_value.GetArr().NextIndexInsert(&stat_mode)
	return_value.GetArr().NextIndexInsert(&stat_nlink)
	return_value.GetArr().NextIndexInsert(&stat_uid)
	return_value.GetArr().NextIndexInsert(&stat_gid)
	return_value.GetArr().NextIndexInsert(&stat_rdev)
	return_value.GetArr().NextIndexInsert(&stat_size)
	return_value.GetArr().NextIndexInsert(&stat_atime)
	return_value.GetArr().NextIndexInsert(&stat_mtime)
	return_value.GetArr().NextIndexInsert(&stat_ctime)
	return_value.GetArr().NextIndexInsert(&stat_blksize)
	return_value.GetArr().NextIndexInsert(&stat_blocks)

	/* Store string indexes referencing the same zval*/

	return_value.GetArr().KeyAddNew(b.CastStrAuto(stat_sb_names[0]), &stat_dev)
	return_value.GetArr().KeyAddNew(b.CastStrAuto(stat_sb_names[1]), &stat_ino)
	return_value.GetArr().KeyAddNew(b.CastStrAuto(stat_sb_names[2]), &stat_mode)
	return_value.GetArr().KeyAddNew(b.CastStrAuto(stat_sb_names[3]), &stat_nlink)
	return_value.GetArr().KeyAddNew(b.CastStrAuto(stat_sb_names[4]), &stat_uid)
	return_value.GetArr().KeyAddNew(b.CastStrAuto(stat_sb_names[5]), &stat_gid)
	return_value.GetArr().KeyAddNew(b.CastStrAuto(stat_sb_names[6]), &stat_rdev)
	return_value.GetArr().KeyAddNew(b.CastStrAuto(stat_sb_names[7]), &stat_size)
	return_value.GetArr().KeyAddNew(b.CastStrAuto(stat_sb_names[8]), &stat_atime)
	return_value.GetArr().KeyAddNew(b.CastStrAuto(stat_sb_names[9]), &stat_mtime)
	return_value.GetArr().KeyAddNew(b.CastStrAuto(stat_sb_names[10]), &stat_ctime)
	return_value.GetArr().KeyAddNew(b.CastStrAuto(stat_sb_names[11]), &stat_blksize)
	return_value.GetArr().KeyAddNew(b.CastStrAuto(stat_sb_names[12]), &stat_blocks)
}
func ZifCopy(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var source *byte
	var target *byte
	var source_len int
	var target_len int
	var zcontext *types.Zval = nil
	var context *core.PhpStreamContext
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 3

		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			source, source_len = fp.ParsePath()
			target, target_len = fp.ParsePath()
			fp.StartOptional()
			zcontext = fp.ParseResourceEx(true)
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	if streams.PhpStreamLocateUrlWrapper(source, nil, 0) == &PhpPlainFilesWrapper && core.PhpCheckOpenBasedir(source) != 0 {
		return_value.SetFalse()
		return
	}
	context = streams.PhpStreamContextFromZval(zcontext, 0)
	if PhpCopyFileCtx(source, target, 0, context) == types.SUCCESS {
		return_value.SetTrue()
		return
	} else {
		return_value.SetFalse()
		return
	}
}
func PhpCopyFile(src *byte, dest *byte) int { return PhpCopyFileCtx(src, dest, 0, nil) }
func PhpCopyFileEx(src *byte, dest *byte, src_flg int) int {
	return PhpCopyFileCtx(src, dest, src_flg, nil)
}
func PhpCopyFileCtx(src *byte, dest *byte, src_flg int, ctx *core.PhpStreamContext) int {
	var srcstream *core.PhpStream = nil
	var deststream *core.PhpStream = nil
	var ret int = types.FAILURE
	var src_s core.PhpStreamStatbuf
	var dest_s core.PhpStreamStatbuf
	switch core.PhpStreamStatPathEx(src, 0, &src_s, ctx) {
	case -1:

		/* non-statable stream */

		goto safe_to_copy
	case 0:

	default:
		return ret
	}
	if zend.S_ISDIR(src_s.GetSb().st_mode) {
		core.PhpErrorDocref(nil, faults.E_WARNING, "The first argument to copy() function cannot be a directory")
		return types.FAILURE
	}
	switch core.PhpStreamStatPathEx(dest, core.PHP_STREAM_URL_STAT_QUIET|core.PHP_STREAM_URL_STAT_NOCACHE, &dest_s, ctx) {
	case -1:

		/* non-statable stream */

		goto safe_to_copy
	case 0:

	default:
		return ret
	}
	if zend.S_ISDIR(dest_s.GetSb().st_mode) {
		core.PhpErrorDocref(nil, faults.E_WARNING, "The second argument to copy() function cannot be a directory")
		return types.FAILURE
	}
	if !(src_s.GetSb().st_ino) || !(dest_s.GetSb().st_ino) {
		goto no_stat
	}
	if src_s.GetSb().st_ino == dest_s.GetSb().st_ino && src_s.GetSb().st_dev == dest_s.GetSb().st_dev {
		return ret
	} else {
		goto safe_to_copy
	}
no_stat:
	var sp *byte
	var dp *byte
	var res int
	if b.Assign(&sp, core.ExpandFilepath(src, nil)) == nil {
		return ret
	}
	if b.Assign(&dp, core.ExpandFilepath(dest, nil)) == nil {
		zend.Efree(sp)
		goto safe_to_copy
	}
	res = !(strcmp(sp, dp))
	zend.Efree(sp)
	zend.Efree(dp)
	if res != 0 {
		return ret
	}
safe_to_copy:
	srcstream = core.PhpStreamOpenWrapperEx(src, "rb", src_flg|core.REPORT_ERRORS, nil, ctx)
	if srcstream == nil {
		return ret
	}
	deststream = core.PhpStreamOpenWrapperEx(dest, "wb", core.REPORT_ERRORS, nil, ctx)
	if srcstream != nil && deststream != nil {
		ret = core.PhpStreamCopyToStreamEx(srcstream, deststream, core.PHP_STREAM_COPY_ALL, nil)
	}
	if srcstream != nil {
		core.PhpStreamClose(srcstream)
	}
	if deststream != nil {
		core.PhpStreamClose(deststream)
	}
	return ret
}
func ZifFread(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var res *types.Zval
	var len_ zend.ZendLong
	var stream *core.PhpStream
	var str *types.ZendString
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 2

		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			res = fp.ParseResource()
			len_ = fp.ParseLong()
			if fp.HasError() {
				fp.HandleError()
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	PHP_STREAM_TO_ZVAL(stream, res)
	if len_ <= 0 {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Length parameter must be greater than 0")
		return_value.SetFalse()
		return
	}
	str = streams.PhpStreamReadToStr(stream, len_)
	if str == nil {
		zend.ZvalPtrDtorStr(return_value)
		return_value.SetFalse()
		return
	}
	return_value.SetString(str)
	return
}
func PhpFgetcsvLookupTrailingSpaces(ptr *byte, len_ int, delimiter byte) *byte {
	var inc_len int
	var last_chars []uint8 = []uint8{0, 0}
	for len_ > 0 {
		if (*ptr) == '0' {
			inc_len = 1
		} else {
			inc_len = PhpMblen(ptr, len_)
		}
		switch inc_len {
		case -2:
			fallthrough
		case -1:
			inc_len = 1
			core.PhpIgnoreValue(mblen(nil, 0))
		case 0:
			goto quit_loop
		case 1:
			fallthrough
		default:
			last_chars[0] = last_chars[1]
			last_chars[1] = *ptr
		}
		ptr += inc_len
		len_ -= inc_len
	}
quit_loop:
	switch last_chars[1] {
	case '\n':
		if last_chars[0] == '\r' {
			return ptr - 2
		}
		fallthrough
	case '\r':
		return ptr - 1
	}
	return ptr
}
func FPUTCSV_FLD_CHK(c __auto__) __auto__ {
	return memchr(field_str.GetVal(), c, field_str.GetLen())
}
func ZifFputcsv(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var delimiter byte = ','
	var enclosure byte = '"'
	var escape_char int = uint8('\\')
	var stream *core.PhpStream
	var fp *types.Zval = nil
	var fields *types.Zval = nil
	var ret ssize_t
	var delimiter_str *byte = nil
	var enclosure_str *byte = nil
	var escape_str *byte = nil
	var delimiter_str_len int = 0
	var enclosure_str_len int = 0
	var escape_str_len int = 0
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 5

		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			fp = fp.ParseResource()
			fields = fp.ParseArray()
			fp.StartOptional()
			delimiter_str, delimiter_str_len = fp.ParseString()
			enclosure_str, enclosure_str_len = fp.ParseString()
			escape_str, escape_str_len = fp.ParseString()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	if delimiter_str != nil {

		/* Make sure that there is at least one character in string */

		if delimiter_str_len < 1 {
			core.PhpErrorDocref(nil, faults.E_WARNING, "delimiter must be a character")
			return_value.SetFalse()
			return
		} else if delimiter_str_len > 1 {
			core.PhpErrorDocref(nil, faults.E_NOTICE, "delimiter must be a single character")
		}

		/* use first character from string */

		delimiter = *delimiter_str

		/* use first character from string */

	}
	if enclosure_str != nil {
		if enclosure_str_len < 1 {
			core.PhpErrorDocref(nil, faults.E_WARNING, "enclosure must be a character")
			return_value.SetFalse()
			return
		} else if enclosure_str_len > 1 {
			core.PhpErrorDocref(nil, faults.E_NOTICE, "enclosure must be a single character")
		}

		/* use first character from string */

		enclosure = *enclosure_str

		/* use first character from string */

	}
	if escape_str != nil {
		if escape_str_len > 1 {
			core.PhpErrorDocref(nil, faults.E_NOTICE, "escape must be empty or a single character")
		}
		if escape_str_len < 1 {
			escape_char = PHP_CSV_NO_ESCAPE
		} else {

			/* use first character from string */

			escape_char = uint8(*escape_str)

			/* use first character from string */

		}
	}
	PHP_STREAM_TO_ZVAL(stream, fp)
	ret = PhpFputcsv(stream, fields, delimiter, enclosure, escape_char)
	if ret < 0 {
		return_value.SetFalse()
		return
	}
	return_value.SetLong(ret)
	return
}
func PhpFputcsv(stream *core.PhpStream, fields *types.Zval, delimiter byte, enclosure byte, escape_char int) ssize_t {
	var count int
	var i int = 0
	var ret int
	var field_tmp *types.Zval
	var csvline zend.SmartStr = zend.MakeSmartStr(0)
	b.Assert(escape_char >= 0 && escape_char <= UCHAR_MAX || escape_char == PHP_CSV_NO_ESCAPE)
	count = types.Z_ARRVAL_P(fields).GetNNumOfElements()
	var __ht *types.HashTable = fields.GetArr()
	for _, _p := range __ht.foreachData() {
		var _z *types.Zval = _p.GetVal()

		field_tmp = _z
		var tmp_field_str *types.ZendString
		var field_str *types.ZendString = zend.ZvalGetTmpString(field_tmp, &tmp_field_str)

		/* enclose a field that contains a delimiter, an enclosure character, or a newline */

		if FPUTCSV_FLD_CHK(delimiter) || FPUTCSV_FLD_CHK(enclosure) || escape_char != PHP_CSV_NO_ESCAPE && FPUTCSV_FLD_CHK(escape_char) || FPUTCSV_FLD_CHK('\n') || FPUTCSV_FLD_CHK('\r') || FPUTCSV_FLD_CHK('\t') || FPUTCSV_FLD_CHK(' ') {
			var ch *byte = field_str.GetVal()
			var end *byte = ch + field_str.GetLen()
			var escaped int = 0
			csvline.AppendByte(enclosure)
			for ch < end {
				if escape_char != PHP_CSV_NO_ESCAPE && (*ch) == escape_char {
					escaped = 1
				} else if escaped == 0 && (*ch) == enclosure {
					csvline.AppendByte(enclosure)
				} else {
					escaped = 0
				}
				csvline.AppendByte(*ch)
				ch++
			}
			csvline.AppendByte(enclosure)
		} else {
			csvline.AppendString(field_str.GetStr())
		}
		if b.PreInc(&i) != count {
			csvline.AppendString(b.CastStr(&delimiter, 1))
		}
		zend.ZendTmpStringRelease(tmp_field_str)
	}
	csvline.AppendByte('\n')
	csvline.ZeroTail()
	ret = core.PhpStreamWrite(stream, csvline.GetS().GetVal(), csvline.GetS().GetLen())
	csvline.Free()
	return ret
}
func ZifFgetcsv(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var delimiter byte = ','
	var enclosure byte = '"'
	var escape int = uint8('\\')

	/* first section exactly as php_fgetss */

	var len_ zend.ZendLong = 0
	var buf_len int
	var buf *byte
	var stream *core.PhpStream
	var fd *types.Zval
	var len_zv *types.Zval = nil
	var delimiter_str *byte = nil
	var delimiter_str_len int = 0
	var enclosure_str *byte = nil
	var enclosure_str_len int = 0
	var escape_str *byte = nil
	var escape_str_len int = 0
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 5

		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			fd = fp.ParseResource()
			fp.StartOptional()
			len_zv = fp.ParseZval()
			delimiter_str, delimiter_str_len = fp.ParseString()
			enclosure_str, enclosure_str_len = fp.ParseString()
			escape_str, escape_str_len = fp.ParseString()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	if delimiter_str != nil {

		/* Make sure that there is at least one character in string */

		if delimiter_str_len < 1 {
			core.PhpErrorDocref(nil, faults.E_WARNING, "delimiter must be a character")
			return_value.SetFalse()
			return
		} else if delimiter_str_len > 1 {
			core.PhpErrorDocref(nil, faults.E_NOTICE, "delimiter must be a single character")
		}

		/* use first character from string */

		delimiter = delimiter_str[0]

		/* use first character from string */

	}
	if enclosure_str != nil {
		if enclosure_str_len < 1 {
			core.PhpErrorDocref(nil, faults.E_WARNING, "enclosure must be a character")
			return_value.SetFalse()
			return
		} else if enclosure_str_len > 1 {
			core.PhpErrorDocref(nil, faults.E_NOTICE, "enclosure must be a single character")
		}

		/* use first character from string */

		enclosure = enclosure_str[0]

		/* use first character from string */

	}
	if escape_str != nil {
		if escape_str_len > 1 {
			core.PhpErrorDocref(nil, faults.E_NOTICE, "escape must be empty or a single character")
		}
		if escape_str_len < 1 {
			escape = PHP_CSV_NO_ESCAPE
		} else {
			escape = uint8(escape_str[0])
		}
	}
	if len_zv != nil && len_zv.GetType() != types.IS_NULL {
		len_ = zend.ZvalGetLong(len_zv)
		if len_ < 0 {
			core.PhpErrorDocref(nil, faults.E_WARNING, "Length parameter may not be negative")
			return_value.SetFalse()
			return
		} else if len_ == 0 {
			len_ = -1
		}
	} else {
		len_ = -1
	}
	PHP_STREAM_TO_ZVAL(stream, fd)
	if len_ < 0 {
		if b.Assign(&buf, core.PhpStreamGetLine(stream, nil, 0, &buf_len)) == nil {
			return_value.SetFalse()
			return
		}
	} else {
		buf = zend.Emalloc(len_ + 1)
		if core.PhpStreamGetLine(stream, buf, len_+1, &buf_len) == nil {
			zend.Efree(buf)
			return_value.SetFalse()
			return
		}
	}
	PhpFgetcsv(stream, delimiter, enclosure, escape, buf_len, buf, return_value)
}
func PhpFgetcsv(
	stream *core.PhpStream,
	delimiter byte,
	enclosure byte,
	escape_char int,
	buf_len int,
	buf *byte,
	return_value *types.Zval,
) {
	var temp *byte
	var tptr *byte
	var bptr *byte
	var line_end *byte
	var limit *byte
	var temp_len int
	var line_end_len int
	var inc_len int
	var first_field types.ZendBool = 1
	b.Assert(escape_char >= 0 && escape_char <= UCHAR_MAX || escape_char == PHP_CSV_NO_ESCAPE)

	/* initialize internal state */

	core.PhpIgnoreValue(mblen(nil, 0))

	/* Now into new section that parses buf for delimiter/enclosure fields */

	bptr = buf
	tptr = (*byte)(PhpFgetcsvLookupTrailingSpaces(buf, buf_len, delimiter))
	line_end_len = buf_len - size_t(tptr-buf)
	limit = tptr
	line_end = limit

	/* reserve workspace for building each individual field */

	temp_len = buf_len
	temp = zend.Emalloc(temp_len + line_end_len + 1)

	/* Initialize return array */

	zend.ArrayInit(return_value)

	/* Main loop to read CSV fields */

	for {
		var comp_end *byte
		var hunk_begin *byte
		tptr = temp
		if bptr < limit {
			if (*bptr) == '0' {
				inc_len = 1
			} else {
				inc_len = PhpMblen(bptr, limit-bptr)
			}
		} else {
			inc_len = 0
		}
		if inc_len == 1 {
			var tmp *byte = bptr
			for (*tmp) != delimiter && isspace(int(*((*uint8)(tmp)))) {
				tmp++
			}
			if (*tmp) == enclosure {
				bptr = tmp
			}
		}
		if first_field != 0 && bptr == line_end {
			zend.AddNextIndexNull(return_value)
			break
		}
		first_field = 0

		/* 2. Read field, leaving bptr pointing at start of next field */

		if inc_len != 0 && (*bptr) == enclosure {
			var state int = 0
			bptr++
			hunk_begin = bptr

			/* 2A. handle enclosure delimited field */

			for {
				switch inc_len {
				case 0:
					switch state {
					case 2:
						memcpy(tptr, hunk_begin, bptr-hunk_begin-1)
						tptr += bptr - hunk_begin - 1
						hunk_begin = bptr
						goto quit_loop_2
					case 1:
						memcpy(tptr, hunk_begin, bptr-hunk_begin)
						tptr += bptr - hunk_begin
						hunk_begin = bptr
						fallthrough
					case 0:
						var new_buf *byte
						var new_len int
						var new_temp *byte
						if hunk_begin != line_end {
							memcpy(tptr, hunk_begin, bptr-hunk_begin)
							tptr += bptr - hunk_begin
							hunk_begin = bptr
						}

						/* add the embedded line end to the field */

						memcpy(tptr, line_end, line_end_len)
						tptr += line_end_len
						if stream == nil {
							goto quit_loop_2
						} else if b.Assign(&new_buf, core.PhpStreamGetLine(stream, nil, 0, &new_len)) == nil {

							/* we've got an unterminated enclosure,
							 * assign all the data from the start of
							 * the enclosure to end of data to the
							 * last element */

							if int(temp_len > size_t(limit-buf)) != 0 {
								goto quit_loop_2
							}
							return_value.GetArr().DestroyEx()
							return_value.SetFalse()
							goto out
						}
						temp_len += new_len
						new_temp = zend.Erealloc(temp, temp_len)
						tptr = new_temp + size_t(tptr-temp)
						temp = new_temp
						zend.Efree(buf)
						buf_len = new_len
						buf = new_buf
						bptr = buf
						hunk_begin = buf
						limit = (*byte)(PhpFgetcsvLookupTrailingSpaces(buf, buf_len, delimiter))
						line_end = limit
						line_end_len = buf_len - size_t(limit-buf)
						state = 0
					}
				case -2:
					fallthrough
				case -1:
					core.PhpIgnoreValue(mblen(nil, 0))
					fallthrough
				case 1:

					/* we need to determine if the enclosure is
					 * 'real' or is it escaped */

					switch state {
					case 1:
						bptr++
						state = 0
					case 2:
						if (*bptr) != enclosure {

							/* real enclosure */

							memcpy(tptr, hunk_begin, bptr-hunk_begin-1)
							tptr += bptr - hunk_begin - 1
							hunk_begin = bptr
							goto quit_loop_2
						}
						memcpy(tptr, hunk_begin, bptr-hunk_begin)
						tptr += bptr - hunk_begin
						bptr++
						hunk_begin = bptr
						state = 0
					default:
						if (*bptr) == enclosure {
							state = 2
						} else if escape_char != PHP_CSV_NO_ESCAPE && (*bptr) == escape_char {
							state = 1
						}
						bptr++
					}
				default:
					switch state {
					case 2:

						/* real enclosure */

						memcpy(tptr, hunk_begin, bptr-hunk_begin-1)
						tptr += bptr - hunk_begin - 1
						hunk_begin = bptr
						goto quit_loop_2
					case 1:
						bptr += inc_len
						memcpy(tptr, hunk_begin, bptr-hunk_begin)
						tptr += bptr - hunk_begin
						hunk_begin = bptr
						state = 0
					default:
						bptr += inc_len
					}
				}
				if bptr < limit {
					if (*bptr) == '0' {
						inc_len = 1
					} else {
						inc_len = PhpMblen(bptr, limit-bptr)
					}
				} else {
					inc_len = 0
				}
			}
		quit_loop_2:

			/* look up for a delimiter */

			for {
				switch inc_len {
				case 0:
					goto quit_loop_3
				case -2:
					fallthrough
				case -1:
					inc_len = 1
					core.PhpIgnoreValue(mblen(nil, 0))
					fallthrough
				case 1:
					if (*bptr) == delimiter {
						goto quit_loop_3
					}
				default:

				}
				bptr += inc_len
				if bptr < limit {
					if (*bptr) == '0' {
						inc_len = 1
					} else {
						inc_len = PhpMblen(bptr, limit-bptr)
					}
				} else {
					inc_len = 0
				}
			}
		quit_loop_3:
			memcpy(tptr, hunk_begin, bptr-hunk_begin)
			tptr += bptr - hunk_begin
			bptr += inc_len
			comp_end = tptr
		} else {

			/* 2B. Handle non-enclosure field */

			hunk_begin = bptr
			for {
				switch inc_len {
				case 0:
					goto quit_loop_4
				case -2:
					fallthrough
				case -1:
					inc_len = 1
					core.PhpIgnoreValue(mblen(nil, 0))
					fallthrough
				case 1:
					if (*bptr) == delimiter {
						goto quit_loop_4
					}
				default:

				}
				bptr += inc_len
				if bptr < limit {
					if (*bptr) == '0' {
						inc_len = 1
					} else {
						inc_len = PhpMblen(bptr, limit-bptr)
					}
				} else {
					inc_len = 0
				}
			}
		quit_loop_4:
			memcpy(tptr, hunk_begin, bptr-hunk_begin)
			tptr += bptr - hunk_begin
			comp_end = (*byte)(PhpFgetcsvLookupTrailingSpaces(temp, tptr-temp, delimiter))
			if (*bptr) == delimiter {
				bptr++
			}
		}

		/* 3. Now pass our field back to php */

		*comp_end = '0'
		zend.AddNextIndexStringl(return_value, temp, comp_end-temp)
		if inc_len <= 0 {
			break
		}
	}
out:
	zend.Efree(temp)
	if stream != nil {
		zend.Efree(buf)
	}
}
func ZifRealpath(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var filename *byte
	var filename_len int
	var resolved_path_buff []byte
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1

		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			filename, filename_len = fp.ParsePath()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	if zend.VCWD_REALPATH(filename, resolved_path_buff) != nil {
		if core.PhpCheckOpenBasedir(resolved_path_buff) != 0 {
			return_value.SetFalse()
			return
		}
		return_value.SetRawString(b.CastStrAuto(resolved_path_buff))
		return
	} else {
		return_value.SetFalse()
		return
	}
}
func PhpNextMetaToken(md *PhpMetaTagsData) PhpMetaTagsToken {
	var ch int = 0
	var compliment int
	var buff []byte
	memset(any(buff), 0, META_DEF_BUFSIZE+1)
	for md.GetUlc() != 0 || core.PhpStreamEof(md.GetStream()) == 0 && b.Assign(&ch, core.PhpStreamGetc(md.GetStream())) {
		if core.PhpStreamEof(md.GetStream()) != 0 {
			break
		}
		if md.GetUlc() != 0 {
			ch = md.GetLc()
			md.SetUlc(0)
		}
		switch ch {
		case '<':
			return TOK_OPENTAG
		case '>':
			return TOK_CLOSETAG
		case '=':
			return TOK_EQUAL
		case '/':
			return TOK_SLASH
		case '\'':
			fallthrough
		case '"':
			compliment = ch
			md.SetTokenLen(0)
			for core.PhpStreamEof(md.GetStream()) == 0 && b.Assign(&ch, core.PhpStreamGetc(md.GetStream())) && ch != compliment && ch != '<' && ch != '>' {
				buff[b.PostInc(&(md.GetTokenLen()))] = ch
				if md.GetTokenLen() == META_DEF_BUFSIZE {
					break
				}
			}
			if ch == '<' || ch == '>' {

				/* Was just an apostrohpe */

				md.SetUlc(1)
				md.SetLc(ch)
			}

			/* We don't need to alloc unless we are in a meta tag */

			if md.GetInMeta() != 0 {
				md.SetTokenData((*byte)(zend.Emalloc(md.GetTokenLen() + 1)))
				memcpy(md.GetTokenData(), buff, md.GetTokenLen()+1)
			}
			return TOK_STRING
		case '\n':
			fallthrough
		case '\r':
			fallthrough
		case '\t':

		case ' ':
			return TOK_SPACE
		default:
			if isalnum(ch) {
				md.SetTokenLen(0)
				buff[b.PostInc(&(md.GetTokenLen()))] = ch
				for core.PhpStreamEof(md.GetStream()) == 0 && b.Assign(&ch, core.PhpStreamGetc(md.GetStream())) && (isalnum(ch) || strchr(PHP_META_HTML401_CHARS, ch)) {
					buff[b.PostInc(&(md.GetTokenLen()))] = ch
					if md.GetTokenLen() == META_DEF_BUFSIZE {
						break
					}
				}

				/* This is ugly, but we have to replace ungetc */

				if !(isalpha(ch)) && ch != '-' {
					md.SetUlc(1)
					md.SetLc(ch)
				}
				md.SetTokenData((*byte)(zend.Emalloc(md.GetTokenLen() + 1)))
				memcpy(md.GetTokenData(), buff, md.GetTokenLen()+1)
				return TOK_ID
			} else {
				return TOK_OTHER
			}
		}
	}
	return TOK_EOF
}
func ZifFnmatch(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var pattern *byte
	var filename *byte
	var pattern_len int
	var filename_len int
	var flags zend.ZendLong = 0
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 3

		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			pattern, pattern_len = fp.ParsePath()
			filename, filename_len = fp.ParsePath()
			fp.StartOptional()
			flags = fp.ParseLong()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	if filename_len >= core.MAXPATHLEN {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Filename exceeds the maximum allowed length of %d characters", core.MAXPATHLEN)
		return_value.SetFalse()
		return
	}
	if pattern_len >= core.MAXPATHLEN {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Pattern exceeds the maximum allowed length of %d characters", core.MAXPATHLEN)
		return_value.SetFalse()
		return
	}
	types.ZVAL_BOOL(return_value, !(fnmatch(pattern, filename, int(flags))))
	return
}
func ZifSysGetTempDir(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	return_value.SetRawString(b.CastStrAuto((*byte)(core.PhpGetTemporaryDirectory())))
	return
}
