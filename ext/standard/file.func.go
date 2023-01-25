// <<generate>>

package standard

import (
	b "sik/builtin"
	"sik/core"
	"sik/core/streams"
	r "sik/runtime"
	"sik/sapi/cli"
	"sik/zend"
)

func FG(v __auto__) __auto__ { return FileGlobals.v }
func PHP_STREAM_TO_ZVAL(stream *core.PhpStream, arg *zend.Zval) {
	zend.ZEND_ASSERT(zend.Z_TYPE_P(arg) == zend.IS_RESOURCE)
	core.PhpStreamFromRes(stream, zend.Z_RES_P(arg))
}
func PhpLeStreamContext() int { return LeStreamContext }
func FileContextDtor(res *zend.ZendResource) {
	var context *core.PhpStreamContext = (*core.PhpStreamContext)(res.GetPtr())
	if context.GetOptions().GetType() != zend.IS_UNDEF {
		zend.ZvalPtrDtor(&context.GetOptions())
		zend.ZVAL_UNDEF(&context.GetOptions())
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
	zend.REGISTER_INI_ENTRIES()
	zend.REGISTER_LONG_CONSTANT("SEEK_SET", r.SEEK_SET, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("SEEK_CUR", r.SEEK_CUR, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("SEEK_END", r.SEEK_END, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("LOCK_SH", PHP_LOCK_SH, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("LOCK_EX", PHP_LOCK_EX, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("LOCK_UN", PHP_LOCK_UN, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("LOCK_NB", PHP_LOCK_NB, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("STREAM_NOTIFY_CONNECT", streams.PHP_STREAM_NOTIFY_CONNECT, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("STREAM_NOTIFY_AUTH_REQUIRED", streams.PHP_STREAM_NOTIFY_AUTH_REQUIRED, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("STREAM_NOTIFY_AUTH_RESULT", streams.PHP_STREAM_NOTIFY_AUTH_RESULT, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("STREAM_NOTIFY_MIME_TYPE_IS", streams.PHP_STREAM_NOTIFY_MIME_TYPE_IS, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("STREAM_NOTIFY_FILE_SIZE_IS", streams.PHP_STREAM_NOTIFY_FILE_SIZE_IS, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("STREAM_NOTIFY_REDIRECTED", streams.PHP_STREAM_NOTIFY_REDIRECTED, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("STREAM_NOTIFY_PROGRESS", streams.PHP_STREAM_NOTIFY_PROGRESS, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("STREAM_NOTIFY_FAILURE", streams.PHP_STREAM_NOTIFY_FAILURE, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("STREAM_NOTIFY_COMPLETED", streams.PHP_STREAM_NOTIFY_COMPLETED, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("STREAM_NOTIFY_RESOLVE", streams.PHP_STREAM_NOTIFY_RESOLVE, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("STREAM_NOTIFY_SEVERITY_INFO", streams.PHP_STREAM_NOTIFY_SEVERITY_INFO, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("STREAM_NOTIFY_SEVERITY_WARN", streams.PHP_STREAM_NOTIFY_SEVERITY_WARN, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("STREAM_NOTIFY_SEVERITY_ERR", streams.PHP_STREAM_NOTIFY_SEVERITY_ERR, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("STREAM_FILTER_READ", streams.PHP_STREAM_FILTER_READ, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("STREAM_FILTER_WRITE", streams.PHP_STREAM_FILTER_WRITE, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("STREAM_FILTER_ALL", streams.PHP_STREAM_FILTER_ALL, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("STREAM_CLIENT_PERSISTENT", PHP_STREAM_CLIENT_PERSISTENT, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("STREAM_CLIENT_ASYNC_CONNECT", PHP_STREAM_CLIENT_ASYNC_CONNECT, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("STREAM_CLIENT_CONNECT", PHP_STREAM_CLIENT_CONNECT, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("STREAM_CRYPTO_METHOD_ANY_CLIENT", streams.STREAM_CRYPTO_METHOD_ANY_CLIENT, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("STREAM_CRYPTO_METHOD_SSLv2_CLIENT", streams.STREAM_CRYPTO_METHOD_SSLv2_CLIENT, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("STREAM_CRYPTO_METHOD_SSLv3_CLIENT", streams.STREAM_CRYPTO_METHOD_SSLv3_CLIENT, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("STREAM_CRYPTO_METHOD_SSLv23_CLIENT", streams.STREAM_CRYPTO_METHOD_SSLv23_CLIENT, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("STREAM_CRYPTO_METHOD_TLS_CLIENT", streams.STREAM_CRYPTO_METHOD_TLS_CLIENT, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("STREAM_CRYPTO_METHOD_TLSv1_0_CLIENT", streams.STREAM_CRYPTO_METHOD_TLSv1_0_CLIENT, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("STREAM_CRYPTO_METHOD_TLSv1_1_CLIENT", streams.STREAM_CRYPTO_METHOD_TLSv1_1_CLIENT, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("STREAM_CRYPTO_METHOD_TLSv1_2_CLIENT", streams.STREAM_CRYPTO_METHOD_TLSv1_2_CLIENT, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("STREAM_CRYPTO_METHOD_TLSv1_3_CLIENT", streams.STREAM_CRYPTO_METHOD_TLSv1_3_CLIENT, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("STREAM_CRYPTO_METHOD_ANY_SERVER", streams.STREAM_CRYPTO_METHOD_ANY_SERVER, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("STREAM_CRYPTO_METHOD_SSLv2_SERVER", streams.STREAM_CRYPTO_METHOD_SSLv2_SERVER, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("STREAM_CRYPTO_METHOD_SSLv3_SERVER", streams.STREAM_CRYPTO_METHOD_SSLv3_SERVER, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("STREAM_CRYPTO_METHOD_SSLv23_SERVER", streams.STREAM_CRYPTO_METHOD_SSLv23_SERVER, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("STREAM_CRYPTO_METHOD_TLS_SERVER", streams.STREAM_CRYPTO_METHOD_TLS_SERVER, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("STREAM_CRYPTO_METHOD_TLSv1_0_SERVER", streams.STREAM_CRYPTO_METHOD_TLSv1_0_SERVER, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("STREAM_CRYPTO_METHOD_TLSv1_1_SERVER", streams.STREAM_CRYPTO_METHOD_TLSv1_1_SERVER, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("STREAM_CRYPTO_METHOD_TLSv1_2_SERVER", streams.STREAM_CRYPTO_METHOD_TLSv1_2_SERVER, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("STREAM_CRYPTO_METHOD_TLSv1_3_SERVER", streams.STREAM_CRYPTO_METHOD_TLSv1_3_SERVER, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("STREAM_CRYPTO_PROTO_SSLv3", streams.STREAM_CRYPTO_METHOD_SSLv3_SERVER, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("STREAM_CRYPTO_PROTO_TLSv1_0", streams.STREAM_CRYPTO_METHOD_TLSv1_0_SERVER, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("STREAM_CRYPTO_PROTO_TLSv1_1", streams.STREAM_CRYPTO_METHOD_TLSv1_1_SERVER, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("STREAM_CRYPTO_PROTO_TLSv1_2", streams.STREAM_CRYPTO_METHOD_TLSv1_2_SERVER, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("STREAM_CRYPTO_PROTO_TLSv1_3", streams.STREAM_CRYPTO_METHOD_TLSv1_3_SERVER, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("STREAM_SHUT_RD", streams.STREAM_SHUT_RD, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("STREAM_SHUT_WR", streams.STREAM_SHUT_WR, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("STREAM_SHUT_RDWR", streams.STREAM_SHUT_RDWR, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("STREAM_SOCK_STREAM", SOCK_STREAM, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("STREAM_SOCK_DGRAM", SOCK_DGRAM, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("STREAM_PEEK", streams.STREAM_PEEK, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("STREAM_OOB", streams.STREAM_OOB, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("STREAM_SERVER_BIND", streams.STREAM_XPORT_BIND, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("STREAM_SERVER_LISTEN", streams.STREAM_XPORT_LISTEN, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("FILE_USE_INCLUDE_PATH", PHP_FILE_USE_INCLUDE_PATH, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("FILE_IGNORE_NEW_LINES", PHP_FILE_IGNORE_NEW_LINES, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("FILE_SKIP_EMPTY_LINES", PHP_FILE_SKIP_EMPTY_LINES, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("FILE_APPEND", PHP_FILE_APPEND, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("FILE_NO_DEFAULT_CONTEXT", PHP_FILE_NO_DEFAULT_CONTEXT, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("FILE_TEXT", 0, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("FILE_BINARY", 0, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("FNM_NOESCAPE", FNM_NOESCAPE, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("FNM_PATHNAME", FNM_PATHNAME, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("FNM_PERIOD", FNM_PERIOD, zend.CONST_CS|zend.CONST_PERSISTENT)
	return zend.SUCCESS
}
func ZmShutdownFile(type_ int, module_number int) int {
	FileGlobalsDtor(&FileGlobals)
	return zend.SUCCESS
}
func ZifFlock(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var res *zend.Zval
	var wouldblock *zend.Zval = nil
	var act int
	var stream *core.PhpStream
	var operation zend.ZendLong = 0
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 3
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
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgResource(_arg, &res, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_RESOURCE
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgLong(_arg, &operation, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			zend.ZendParseArgZvalDeref(_arg, &wouldblock, 0)
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
			return
		}
		break
	}
	PHP_STREAM_TO_ZVAL(stream, res)
	act = operation & 3
	if act < 1 || act > 3 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Illegal operation argument")
		zend.RETVAL_FALSE
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
		zend.RETVAL_FALSE
		return
	}
	zend.RETVAL_TRUE
	return
}
func ZifGetMetaTags(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var filename *byte
	var filename_len int
	var use_include_path zend.ZendBool = 0
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
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgPath(_arg, &filename, &filename_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_PATH
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgBool(_arg, &use_include_path, &_dummy, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_BOOL
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
			return
		}
		break
	}
	md.SetStream(core.PhpStreamOpenWrapper(filename, "rb", b.Cond(use_include_path != 0, core.USE_PATH, 0)|core.REPORT_ERRORS, nil))
	if md.GetStream() == nil {
		zend.RETVAL_FALSE
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
func ZifFileGetContents(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var filename *byte
	var filename_len int
	var use_include_path zend.ZendBool = 0
	var stream *core.PhpStream
	var offset zend.ZendLong = 0
	var maxlen zend.ZendLong = ssize_t(core.PHP_STREAM_COPY_ALL)
	var zcontext *zend.Zval = nil
	var context *core.PhpStreamContext = nil
	var contents *zend.ZendString

	/* Parse arguments */

	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 5
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
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgPath(_arg, &filename, &filename_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_PATH
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgBool(_arg, &use_include_path, &_dummy, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_BOOL
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgResource(_arg, &zcontext, 1) == 0 {
				_expected_type = zend.Z_EXPECTED_RESOURCE
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgLong(_arg, &offset, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgLong(_arg, &maxlen, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
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
			return
		}
		break
	}
	if zend.ZEND_NUM_ARGS() == 5 && maxlen < 0 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "length must be greater than or equal to zero")
		zend.RETVAL_FALSE
		return
	}
	context = streams.PhpStreamContextFromZval(zcontext, 0)
	stream = core.PhpStreamOpenWrapperEx(filename, "rb", b.Cond(use_include_path != 0, core.USE_PATH, 0)|core.REPORT_ERRORS, nil, context)
	if stream == nil {
		zend.RETVAL_FALSE
		return
	}
	if offset != 0 && core.PhpStreamSeek(stream, offset, b.Cond(offset > 0, r.SEEK_SET, r.SEEK_END)) < 0 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Failed to seek to position "+zend.ZEND_LONG_FMT+" in the stream", offset)
		core.PhpStreamClose(stream)
		zend.RETVAL_FALSE
		return
	}
	if b.Assign(&contents, core.PhpStreamCopyToMem(stream, maxlen, 0)) != nil {
		zend.RETVAL_STR(contents)
	} else {
		zend.RETVAL_EMPTY_STRING()
	}
	core.PhpStreamClose(stream)
}
func ZifFilePutContents(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var stream *core.PhpStream
	var filename *byte
	var filename_len int
	var data *zend.Zval
	var numbytes ssize_t = 0
	var flags zend.ZendLong = 0
	var zcontext *zend.Zval = nil
	var context *core.PhpStreamContext = nil
	var srcstream *core.PhpStream = nil
	var mode []byte = "wb"
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 4
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
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgPath(_arg, &filename, &filename_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_PATH
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			zend.ZendParseArgZvalDeref(_arg, &data, 0)
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgLong(_arg, &flags, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgResource(_arg, &zcontext, 1) == 0 {
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
			return
		}
		break
	}
	if zend.Z_TYPE_P(data) == zend.IS_RESOURCE {
		core.PhpStreamFromZval(srcstream, data)
	}
	context = streams.PhpStreamContextFromZval(zcontext, flags&PHP_FILE_NO_DEFAULT_CONTEXT)
	if (flags & PHP_FILE_APPEND) != 0 {
		mode[0] = 'a'
	} else if (flags & LOCK_EX) != 0 {

		/* check to make sure we are dealing with a regular file */

		if core.PhpMemnstr(filename, "://", b.SizeOf("\"://\"")-1, filename+filename_len) {
			if strncasecmp(filename, "file://", b.SizeOf("\"file://\"")-1) {
				core.PhpErrorDocref(nil, zend.E_WARNING, "Exclusive locks may only be set for regular files")
				zend.RETVAL_FALSE
				return
			}
		}
		mode[0] = 'c'
	}
	mode[2] = '0'
	stream = core.PhpStreamOpenWrapperEx(filename, mode, b.Cond((flags&PHP_FILE_USE_INCLUDE_PATH) != 0, core.USE_PATH, 0)|core.REPORT_ERRORS, nil, context)
	if stream == nil {
		zend.RETVAL_FALSE
		return
	}
	if (flags&LOCK_EX) != 0 && (core.PhpStreamSupportsLock(stream) == 0 || core.PhpStreamLock(stream, LOCK_EX) != 0) {
		core.PhpStreamClose(stream)
		core.PhpErrorDocref(nil, zend.E_WARNING, "Exclusive locks are not supported for this stream")
		zend.RETVAL_FALSE
		return
	}
	if mode[0] == 'c' {
		core.PhpStreamTruncateSetSize(stream, 0)
	}
	switch zend.Z_TYPE_P(data) {
	case zend.IS_RESOURCE:
		var len_ int
		if core.PhpStreamCopyToStreamEx(srcstream, stream, core.PHP_STREAM_COPY_ALL, &len_) != zend.SUCCESS {
			numbytes = -1
		} else {
			if len_ > zend.ZEND_LONG_MAX {
				core.PhpErrorDocref(nil, zend.E_WARNING, "content truncated from %zu to "+zend.ZEND_LONG_FMT+" bytes", len_, zend.ZEND_LONG_MAX)
				len_ = zend.ZEND_LONG_MAX
			}
			numbytes = len_
		}
		break
	case zend.IS_NULL:

	case zend.IS_LONG:

	case zend.IS_DOUBLE:

	case zend.IS_FALSE:

	case zend.IS_TRUE:
		zend.ConvertToStringEx(data)
	case zend.IS_STRING:
		if zend.Z_STRLEN_P(data) != 0 {
			numbytes = core.PhpStreamWrite(stream, zend.Z_STRVAL_P(data), zend.Z_STRLEN_P(data))
			if numbytes != zend.Z_STRLEN_P(data) {
				core.PhpErrorDocref(nil, zend.E_WARNING, "Only %zd of %zd bytes written, possibly out of free disk space", numbytes, zend.Z_STRLEN_P(data))
				numbytes = -1
			}
		}
		break
	case zend.IS_ARRAY:
		if zend.ZendHashNumElements(zend.Z_ARRVAL_P(data)) {
			var bytes_written ssize_t
			var tmp *zend.Zval
			for {
				var __ht *zend.HashTable = zend.Z_ARRVAL_P(data)
				var _p *zend.Bucket = __ht.GetArData()
				var _end *zend.Bucket = _p + __ht.GetNNumUsed()
				for ; _p != _end; _p++ {
					var _z *zend.Zval = &_p.GetVal()

					if zend.Z_TYPE_P(_z) == zend.IS_UNDEF {
						continue
					}
					tmp = _z
					var t *zend.ZendString
					var str *zend.ZendString = zend.ZvalGetTmpString(tmp, &t)
					if zend.ZSTR_LEN(str) != 0 {
						numbytes += zend.ZSTR_LEN(str)
						bytes_written = core.PhpStreamWrite(stream, zend.ZSTR_VAL(str), zend.ZSTR_LEN(str))
						if bytes_written != zend.ZSTR_LEN(str) {
							core.PhpErrorDocref(nil, zend.E_WARNING, "Failed to write %zd bytes to %s", zend.ZSTR_LEN(str), filename)
							zend.ZendTmpStringRelease(t)
							numbytes = -1
							break
						}
					}
					zend.ZendTmpStringRelease(t)
				}
				break
			}
		}
		break
	case zend.IS_OBJECT:
		if zend.Z_OBJ_HT_P(data) != nil {
			var out zend.Zval
			if zend.ZendStdCastObjectTostring(data, &out, zend.IS_STRING) == zend.SUCCESS {
				numbytes = core.PhpStreamWrite(stream, zend.Z_STRVAL(out), zend.Z_STRLEN(out))
				if numbytes != zend.Z_STRLEN(out) {
					core.PhpErrorDocref(nil, zend.E_WARNING, "Only %zd of %zd bytes written, possibly out of free disk space", numbytes, zend.Z_STRLEN(out))
					numbytes = -1
				}
				zend.ZvalPtrDtorStr(&out)
				break
			}
		}
	default:
		numbytes = -1
		break
	}
	core.PhpStreamClose(stream)
	if numbytes < 0 {
		zend.RETVAL_FALSE
		return
	}
	zend.RETVAL_LONG(numbytes)
	return
}
func ZifFile(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var filename *byte
	var filename_len int
	var p *byte
	var s *byte
	var e *byte
	var i int = 0
	var eol_marker byte = '\n'
	var flags zend.ZendLong = 0
	var use_include_path zend.ZendBool
	var include_new_line zend.ZendBool
	var skip_blank_lines zend.ZendBool
	var stream *core.PhpStream
	var zcontext *zend.Zval = nil
	var context *core.PhpStreamContext = nil
	var target_buf *zend.ZendString

	/* Parse arguments */

	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 3
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
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgPath(_arg, &filename, &filename_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_PATH
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgLong(_arg, &flags, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgResource(_arg, &zcontext, 1) == 0 {
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
			return
		}
		break
	}
	if flags < 0 || flags > (PHP_FILE_USE_INCLUDE_PATH|PHP_FILE_IGNORE_NEW_LINES|PHP_FILE_SKIP_EMPTY_LINES|PHP_FILE_NO_DEFAULT_CONTEXT) {
		core.PhpErrorDocref(nil, zend.E_WARNING, "'"+zend.ZEND_LONG_FMT+"' flag is not supported", flags)
		zend.RETVAL_FALSE
		return
	}
	use_include_path = flags & PHP_FILE_USE_INCLUDE_PATH
	include_new_line = !(flags & PHP_FILE_IGNORE_NEW_LINES)
	skip_blank_lines = flags & PHP_FILE_SKIP_EMPTY_LINES
	context = streams.PhpStreamContextFromZval(zcontext, flags&PHP_FILE_NO_DEFAULT_CONTEXT)
	stream = core.PhpStreamOpenWrapperEx(filename, "rb", b.Cond(use_include_path != 0, core.USE_PATH, 0)|core.REPORT_ERRORS, nil, context)
	if stream == nil {
		zend.RETVAL_FALSE
		return
	}

	/* Initialize return array */

	zend.ArrayInit(return_value)
	if b.Assign(&target_buf, core.PhpStreamCopyToMem(stream, core.PHP_STREAM_COPY_ALL, 0)) != nil {
		s = zend.ZSTR_VAL(target_buf)
		e = zend.ZSTR_VAL(target_buf) + zend.ZSTR_LEN(target_buf)
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
				if p != zend.ZSTR_VAL(target_buf) && eol_marker == '\n' && (*(p - 1)) == '\r' {
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
		zend.ZendStringFree(target_buf)
	}
	core.PhpStreamClose(stream)
}
func ZifTempnam(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var dir *byte
	var prefix *byte
	var dir_len int
	var prefix_len int
	var opened_path *zend.ZendString
	var fd int
	var p *zend.ZendString
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
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgPath(_arg, &dir, &dir_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_PATH
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgPath(_arg, &prefix, &prefix_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_PATH
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
			return
		}
		break
	}
	p = PhpBasename(prefix, prefix_len, nil, 0)
	if zend.ZSTR_LEN(p) > 64 {
		zend.ZSTR_VAL(p)[63] = '0'
	}
	zend.RETVAL_FALSE
	if b.Assign(&fd, core.PhpOpenTemporaryFdEx(dir, zend.ZSTR_VAL(p), &opened_path, core.PHP_TMP_FILE_OPEN_BASEDIR_CHECK_ALWAYS)) >= 0 {
		close(fd)
		zend.RETVAL_STR(opened_path)
	}
	zend.ZendStringReleaseEx(p, 0)
}
func PhpIfTmpfile(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var stream *core.PhpStream
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	stream = streams._phpStreamFopenTmpfile(0)
	if stream != nil {
		core.PhpStreamToZval(stream, return_value)
	} else {
		zend.RETVAL_FALSE
		return
	}
}
func PhpIfFopen(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var filename *byte
	var mode *byte
	var filename_len int
	var mode_len int
	var use_include_path zend.ZendBool = 0
	var zcontext *zend.Zval = nil
	var stream *core.PhpStream
	var context *core.PhpStreamContext = nil
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 4
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
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgPath(_arg, &filename, &filename_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_PATH
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgString(_arg, &mode, &mode_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgBool(_arg, &use_include_path, &_dummy, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_BOOL
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgResource(_arg, &zcontext, 1) == 0 {
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
			zend.RETVAL_FALSE
			return
		}
		break
	}
	context = streams.PhpStreamContextFromZval(zcontext, 0)
	stream = core.PhpStreamOpenWrapperEx(filename, mode, b.Cond(use_include_path != 0, core.USE_PATH, 0)|core.REPORT_ERRORS, nil, context)
	if stream == nil {
		zend.RETVAL_FALSE
		return
	}
	core.PhpStreamToZval(stream, return_value)
}
func ZifFclose(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var res *zend.Zval
	var stream *core.PhpStream
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
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgResource(_arg, &res, 0) == 0 {
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
			zend.RETVAL_FALSE
			return
		}
		break
	}
	PHP_STREAM_TO_ZVAL(stream, res)
	if stream.HasFlags(core.PHP_STREAM_FLAG_NO_FCLOSE) {
		core.PhpErrorDocref(nil, zend.E_WARNING, "%d is not a valid stream resource", stream.GetRes().GetHandle())
		zend.RETVAL_FALSE
		return
	}
	core.PhpStreamFree(stream, core.PHP_STREAM_FREE_KEEP_RSRC|b.Cond(stream.GetIsPersistent() != 0, core.PHP_STREAM_FREE_CLOSE_PERSISTENT, core.PHP_STREAM_FREE_CLOSE))
	zend.RETVAL_TRUE
	return
}
func ZifPopen(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
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
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgPath(_arg, &command, &command_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_PATH
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgString(_arg, &mode, &mode_len, 0) == 0 {
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
			return
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
		core.PhpErrorDocref2(nil, command, posix_mode, zend.E_WARNING, "%s", strerror(errno))
		zend.Efree(posix_mode)
		zend.RETVAL_FALSE
		return
	}
	stream = streams.PhpStreamFopenFromPipe(fp, mode)
	if stream == nil {
		core.PhpErrorDocref2(nil, command, mode, zend.E_WARNING, "%s", strerror(errno))
		zend.RETVAL_FALSE
	} else {
		core.PhpStreamToZval(stream, return_value)
	}
	zend.Efree(posix_mode)
}
func ZifPclose(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var res *zend.Zval
	var stream *core.PhpStream
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
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgResource(_arg, &res, 0) == 0 {
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
			zend.RETVAL_FALSE
			return
		}
		break
	}
	PHP_STREAM_TO_ZVAL(stream, res)
	FG(pclose_wait) = 1
	zend.ZendListClose(stream.GetRes())
	FG(pclose_wait) = 0
	zend.RETVAL_LONG(FG(pclose_ret))
	return
}
func ZifFeof(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var res *zend.Zval
	var stream *core.PhpStream
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
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgResource(_arg, &res, 0) == 0 {
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
			zend.RETVAL_FALSE
			return
		}
		break
	}
	PHP_STREAM_TO_ZVAL(stream, res)
	if core.PhpStreamEof(stream) != 0 {
		zend.RETVAL_TRUE
		return
	} else {
		zend.RETVAL_FALSE
		return
	}
}
func ZifFgets(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var res *zend.Zval
	var len_ zend.ZendLong = 1024
	var buf *byte = nil
	var argc int = zend.ZEND_NUM_ARGS()
	var line_len int = 0
	var str *zend.ZendString
	var stream *core.PhpStream
	for {
		var _flags int = 0
		var _min_num_args int = 1
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
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgResource(_arg, &res, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_RESOURCE
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgLong(_arg, &len_, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
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
			zend.RETVAL_FALSE
			return
		}
		break
	}
	PHP_STREAM_TO_ZVAL(stream, res)
	if argc == 1 {

		/* ask streams to give us a buffer of an appropriate size */

		buf = core.PhpStreamGetLine(stream, nil, 0, &line_len)
		if buf == nil {
			zend.RETVAL_FALSE
			return
		}

		// TODO: avoid reallocation ???

		zend.RETVAL_STRINGL(buf, line_len)
		zend.Efree(buf)
	} else if argc > 1 {
		if len_ <= 0 {
			core.PhpErrorDocref(nil, zend.E_WARNING, "Length parameter must be greater than 0")
			zend.RETVAL_FALSE
			return
		}
		str = zend.ZendStringAlloc(len_, 0)
		if core.PhpStreamGetLine(stream, zend.ZSTR_VAL(str), len_, &line_len) == nil {
			zend.ZendStringEfree(str)
			zend.RETVAL_FALSE
			return
		}

		/* resize buffer if it's much larger than the result.
		 * Only needed if the user requested a buffer size. */

		if line_len < int(len_/2) {
			str = zend.ZendStringTruncate(str, line_len, 0)
		} else {
			zend.ZSTR_LEN(str) = line_len
		}
		zend.RETVAL_NEW_STR(str)
		return
	}
}
func ZifFgetc(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var res *zend.Zval
	var buf []byte
	var result int
	var stream *core.PhpStream
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
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgResource(_arg, &res, 0) == 0 {
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
			zend.RETVAL_FALSE
			return
		}
		break
	}
	PHP_STREAM_TO_ZVAL(stream, res)
	result = core.PhpStreamGetc(stream)
	if result == r.EOF {
		zend.RETVAL_FALSE
	} else {
		buf[0] = result
		buf[1] = '0'
		zend.RETVAL_STRINGL(buf, 1)
		return
	}
}
func ZifFgetss(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var fd *zend.Zval
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
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgResource(_arg, &fd, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_RESOURCE
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgLong(_arg, &bytes, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgString(_arg, &allowed_tags, &allowed_tags_len, 0) == 0 {
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
			zend.RETVAL_FALSE
			return
		}
		break
	}
	PHP_STREAM_TO_ZVAL(stream, fd)
	if zend.ZEND_NUM_ARGS() >= 2 {
		if bytes <= 0 {
			core.PhpErrorDocref(nil, zend.E_WARNING, "Length parameter must be greater than 0")
			zend.RETVAL_FALSE
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
		zend.RETVAL_FALSE
		return
	}
	retval_len = PhpStripTags(retval, actual_len, &stream.GetFgetssState(), allowed_tags, allowed_tags_len)

	// TODO: avoid reallocation ???

	zend.RETVAL_STRINGL(retval, retval_len)
	zend.Efree(retval)
}
func ZifFscanf(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var result int
	var argc int = 0
	var format_len int
	var args *zend.Zval = nil
	var file_handle *zend.Zval
	var buf *byte
	var format *byte
	var len_ int
	var what any
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = -1
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
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgResource(_arg, &file_handle, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_RESOURCE
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgString(_arg, &format, &format_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			var _num_varargs int = _num_args - _i - 0
			if _num_varargs > 0 {
				args = _real_arg + 1
				argc = _num_varargs
				_i += _num_varargs
				_real_arg += _num_varargs
			} else {
				args = nil
				argc = 0
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
			return
		}
		break
	}
	what = zend.ZendFetchResource2(zend.Z_RES_P(file_handle), "File-Handle", streams.PhpFileLeStream(), streams.PhpFileLePstream())

	/* we can't do a ZEND_VERIFY_RESOURCE(what), otherwise we end up
	 * with a leak if we have an invalid filehandle. This needs changing
	 * if the code behind ZEND_VERIFY_RESOURCE changed. - cc */

	if !what {
		zend.RETVAL_FALSE
		return
	}
	buf = core.PhpStreamGetLine((*core.PhpStream)(what), nil, 0, &len_)
	if buf == nil {
		zend.RETVAL_FALSE
		return
	}
	result = PhpSscanfInternal(buf, format, argc, args, 0, return_value)
	zend.Efree(buf)
	if SCAN_ERROR_WRONG_PARAM_COUNT == result {
		zend.WRONG_PARAM_COUNT
	}
}
func ZifFwrite(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var res *zend.Zval
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
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgResource(_arg, &res, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_RESOURCE
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgString(_arg, &input, &inputlen, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgLong(_arg, &maxlen, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
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
			zend.RETVAL_FALSE
			return
		}
		break
	}
	if zend.ZEND_NUM_ARGS() == 2 {
		num_bytes = inputlen
	} else if maxlen <= 0 {
		num_bytes = 0
	} else {
		num_bytes = cli.MIN(int(maxlen), inputlen)
	}
	if num_bytes == 0 {
		zend.RETVAL_LONG(0)
		return
	}
	PHP_STREAM_TO_ZVAL(stream, res)
	ret = core.PhpStreamWrite(stream, input, num_bytes)
	if ret < 0 {
		zend.RETVAL_FALSE
		return
	}
	zend.RETVAL_LONG(ret)
	return
}
func ZifFflush(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var res *zend.Zval
	var ret int
	var stream *core.PhpStream
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
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgResource(_arg, &res, 0) == 0 {
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
			zend.RETVAL_FALSE
			return
		}
		break
	}
	PHP_STREAM_TO_ZVAL(stream, res)
	ret = core.PhpStreamFlush(stream)
	if ret != 0 {
		zend.RETVAL_FALSE
		return
	}
	zend.RETVAL_TRUE
	return
}
func ZifRewind(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var res *zend.Zval
	var stream *core.PhpStream
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
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgResource(_arg, &res, 0) == 0 {
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
			zend.RETVAL_FALSE
			return
		}
		break
	}
	PHP_STREAM_TO_ZVAL(stream, res)
	if -1 == core.PhpStreamRewind(stream) {
		zend.RETVAL_FALSE
		return
	}
	zend.RETVAL_TRUE
	return
}
func ZifFtell(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var res *zend.Zval
	var ret zend.ZendLong
	var stream *core.PhpStream
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
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgResource(_arg, &res, 0) == 0 {
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
			zend.RETVAL_FALSE
			return
		}
		break
	}
	PHP_STREAM_TO_ZVAL(stream, res)
	ret = core.PhpStreamTell(stream)
	if ret == -1 {
		zend.RETVAL_FALSE
		return
	}
	zend.RETVAL_LONG(ret)
	return
}
func ZifFseek(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var res *zend.Zval
	var offset zend.ZendLong
	var whence zend.ZendLong = r.SEEK_SET
	var stream *core.PhpStream
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 3
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
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgResource(_arg, &res, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_RESOURCE
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgLong(_arg, &offset, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgLong(_arg, &whence, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
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
			zend.RETVAL_FALSE
			return
		}
		break
	}
	PHP_STREAM_TO_ZVAL(stream, res)
	zend.RETVAL_LONG(core.PhpStreamSeek(stream, offset, int(whence)))
	return
}
func PhpMkdirEx(dir *byte, mode zend.ZendLong, options int) int {
	var ret int
	if core.PhpCheckOpenBasedir(dir) != 0 {
		return -1
	}
	if b.Assign(&ret, zend.VCWD_MKDIR(dir, mode_t(mode))) < 0 && (options&core.REPORT_ERRORS) != 0 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "%s", strerror(errno))
	}
	return ret
}
func PhpMkdir(dir *byte, mode zend.ZendLong) int {
	return PhpMkdirEx(dir, mode, core.REPORT_ERRORS)
}
func ZifMkdir(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var dir *byte
	var dir_len int
	var zcontext *zend.Zval = nil
	var mode zend.ZendLong = 0777
	var recursive zend.ZendBool = 0
	var context *core.PhpStreamContext
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 4
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
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgPath(_arg, &dir, &dir_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_PATH
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgLong(_arg, &mode, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgBool(_arg, &recursive, &_dummy, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_BOOL
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgResource(_arg, &zcontext, 1) == 0 {
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
			zend.RETVAL_FALSE
			return
		}
		break
	}
	context = streams.PhpStreamContextFromZval(zcontext, 0)
	zend.RETVAL_BOOL(core.PhpStreamMkdir(dir, int(mode), b.Cond(recursive != 0, core.PHP_STREAM_MKDIR_RECURSIVE, 0)|core.REPORT_ERRORS, context) != 0)
	return
}
func ZifRmdir(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var dir *byte
	var dir_len int
	var zcontext *zend.Zval = nil
	var context *core.PhpStreamContext
	for {
		var _flags int = 0
		var _min_num_args int = 1
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
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgPath(_arg, &dir, &dir_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_PATH
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgResource(_arg, &zcontext, 1) == 0 {
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
			zend.RETVAL_FALSE
			return
		}
		break
	}
	context = streams.PhpStreamContextFromZval(zcontext, 0)
	zend.RETVAL_BOOL(core.PhpStreamRmdir(dir, core.REPORT_ERRORS, context) != 0)
	return
}
func ZifReadfile(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var filename *byte
	var filename_len int
	var size int = 0
	var use_include_path zend.ZendBool = 0
	var zcontext *zend.Zval = nil
	var stream *core.PhpStream
	var context *core.PhpStreamContext = nil
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 3
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
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgPath(_arg, &filename, &filename_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_PATH
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgBool(_arg, &use_include_path, &_dummy, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_BOOL
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgResource(_arg, &zcontext, 1) == 0 {
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
			zend.RETVAL_FALSE
			return
		}
		break
	}
	context = streams.PhpStreamContextFromZval(zcontext, 0)
	stream = core.PhpStreamOpenWrapperEx(filename, "rb", b.Cond(use_include_path != 0, core.USE_PATH, 0)|core.REPORT_ERRORS, nil, context)
	if stream != nil {
		size = core.PhpStreamPassthru(stream)
		core.PhpStreamClose(stream)
		zend.RETVAL_LONG(size)
		return
	}
	zend.RETVAL_FALSE
	return
}
func ZifUmask(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var mask zend.ZendLong = 0
	var oldumask int
	oldumask = umask(077)
	if BG(umask) == -1 {
		BG(umask) = oldumask
	}
	for {
		var _flags int = 0
		var _min_num_args int = 0
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
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgLong(_arg, &mask, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
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
			zend.RETVAL_FALSE
			return
		}
		break
	}
	if zend.ZEND_NUM_ARGS() == 0 {
		umask(oldumask)
	} else {
		umask(int(mask))
	}
	zend.RETVAL_LONG(oldumask)
	return
}
func ZifFpassthru(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var res *zend.Zval
	var size int
	var stream *core.PhpStream
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
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgResource(_arg, &res, 0) == 0 {
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
			zend.RETVAL_FALSE
			return
		}
		break
	}
	PHP_STREAM_TO_ZVAL(stream, res)
	size = core.PhpStreamPassthru(stream)
	zend.RETVAL_LONG(size)
	return
}
func ZifRename(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var old_name *byte
	var new_name *byte
	var old_name_len int
	var new_name_len int
	var zcontext *zend.Zval = nil
	var wrapper *core.PhpStreamWrapper
	var context *core.PhpStreamContext
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 3
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
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgPath(_arg, &old_name, &old_name_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_PATH
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgPath(_arg, &new_name, &new_name_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_PATH
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgResource(_arg, &zcontext, 1) == 0 {
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
			zend.RETVAL_FALSE
			return
		}
		break
	}
	wrapper = streams.PhpStreamLocateUrlWrapper(old_name, nil, 0)
	if wrapper == nil || wrapper.GetWops() == nil {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Unable to locate stream wrapper")
		zend.RETVAL_FALSE
		return
	}
	if wrapper.GetWops().GetRename() == nil {
		core.PhpErrorDocref(nil, zend.E_WARNING, "%s wrapper does not support renaming", b.CondF1(wrapper.GetWops().GetLabel() != nil, func() *byte { return wrapper.GetWops().GetLabel() }, "Source"))
		zend.RETVAL_FALSE
		return
	}
	if wrapper != streams.PhpStreamLocateUrlWrapper(new_name, nil, 0) {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Cannot rename a file across wrapper types")
		zend.RETVAL_FALSE
		return
	}
	context = streams.PhpStreamContextFromZval(zcontext, 0)
	zend.RETVAL_BOOL(wrapper.GetWops().GetRename()(wrapper, old_name, new_name, 0, context) != 0)
	return
}
func ZifUnlink(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var filename *byte
	var filename_len int
	var wrapper *core.PhpStreamWrapper
	var zcontext *zend.Zval = nil
	var context *core.PhpStreamContext = nil
	for {
		var _flags int = 0
		var _min_num_args int = 1
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
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgPath(_arg, &filename, &filename_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_PATH
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgResource(_arg, &zcontext, 1) == 0 {
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
			zend.RETVAL_FALSE
			return
		}
		break
	}
	context = streams.PhpStreamContextFromZval(zcontext, 0)
	wrapper = streams.PhpStreamLocateUrlWrapper(filename, nil, 0)
	if wrapper == nil || wrapper.GetWops() == nil {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Unable to locate stream wrapper")
		zend.RETVAL_FALSE
		return
	}
	if wrapper.GetWops().GetUnlink() == nil {
		core.PhpErrorDocref(nil, zend.E_WARNING, "%s does not allow unlinking", b.CondF1(wrapper.GetWops().GetLabel() != nil, func() *byte { return wrapper.GetWops().GetLabel() }, "Wrapper"))
		zend.RETVAL_FALSE
		return
	}
	zend.RETVAL_BOOL(wrapper.GetWops().GetUnlink()(wrapper, filename, core.REPORT_ERRORS, context) != 0)
	return
}
func PhpIfFtruncate(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var fp *zend.Zval
	var size zend.ZendLong
	var stream *core.PhpStream
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
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgResource(_arg, &fp, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_RESOURCE
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgLong(_arg, &size, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
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
			zend.RETVAL_FALSE
			return
		}
		break
	}
	if size < 0 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Negative size is not supported")
		zend.RETVAL_FALSE
		return
	}
	PHP_STREAM_TO_ZVAL(stream, fp)
	if core.PhpStreamTruncateSupported(stream) == 0 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Can't truncate this stream!")
		zend.RETVAL_FALSE
		return
	}
	zend.RETVAL_BOOL(0 == core.PhpStreamTruncateSetSize(stream, size))
	return
}
func PhpIfFstat(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var fp *zend.Zval
	var stat_dev zend.Zval
	var stat_ino zend.Zval
	var stat_mode zend.Zval
	var stat_nlink zend.Zval
	var stat_uid zend.Zval
	var stat_gid zend.Zval
	var stat_rdev zend.Zval
	var stat_size zend.Zval
	var stat_atime zend.Zval
	var stat_mtime zend.Zval
	var stat_ctime zend.Zval
	var stat_blksize zend.Zval
	var stat_blocks zend.Zval
	var stream *core.PhpStream
	var stat_ssb core.PhpStreamStatbuf
	var stat_sb_names []*byte = []*byte{"dev", "ino", "mode", "nlink", "uid", "gid", "rdev", "size", "atime", "mtime", "ctime", "blksize", "blocks"}
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
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgResource(_arg, &fp, 0) == 0 {
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
			zend.RETVAL_FALSE
			return
		}
		break
	}
	PHP_STREAM_TO_ZVAL(stream, fp)
	if core.PhpStreamStat(stream, &stat_ssb) != 0 {
		zend.RETVAL_FALSE
		return
	}
	zend.ArrayInit(return_value)
	zend.ZVAL_LONG(&stat_dev, stat_ssb.sb.st_dev)
	zend.ZVAL_LONG(&stat_ino, stat_ssb.sb.st_ino)
	zend.ZVAL_LONG(&stat_mode, stat_ssb.sb.st_mode)
	zend.ZVAL_LONG(&stat_nlink, stat_ssb.sb.st_nlink)
	zend.ZVAL_LONG(&stat_uid, stat_ssb.sb.st_uid)
	zend.ZVAL_LONG(&stat_gid, stat_ssb.sb.st_gid)
	zend.ZVAL_LONG(&stat_rdev, stat_ssb.sb.st_rdev)
	zend.ZVAL_LONG(&stat_size, stat_ssb.sb.st_size)
	zend.ZVAL_LONG(&stat_atime, stat_ssb.sb.st_atime)
	zend.ZVAL_LONG(&stat_mtime, stat_ssb.sb.st_mtime)
	zend.ZVAL_LONG(&stat_ctime, stat_ssb.sb.st_ctime)
	zend.ZVAL_LONG(&stat_blksize, stat_ssb.sb.st_blksize)
	zend.ZVAL_LONG(&stat_blocks, stat_ssb.sb.st_blocks)

	/* Store numeric indexes in proper order */

	zend.ZendHashNextIndexInsert(zend.Z_ARRVAL_P(return_value), &stat_dev)
	zend.ZendHashNextIndexInsert(zend.Z_ARRVAL_P(return_value), &stat_ino)
	zend.ZendHashNextIndexInsert(zend.Z_ARRVAL_P(return_value), &stat_mode)
	zend.ZendHashNextIndexInsert(zend.Z_ARRVAL_P(return_value), &stat_nlink)
	zend.ZendHashNextIndexInsert(zend.Z_ARRVAL_P(return_value), &stat_uid)
	zend.ZendHashNextIndexInsert(zend.Z_ARRVAL_P(return_value), &stat_gid)
	zend.ZendHashNextIndexInsert(zend.Z_ARRVAL_P(return_value), &stat_rdev)
	zend.ZendHashNextIndexInsert(zend.Z_ARRVAL_P(return_value), &stat_size)
	zend.ZendHashNextIndexInsert(zend.Z_ARRVAL_P(return_value), &stat_atime)
	zend.ZendHashNextIndexInsert(zend.Z_ARRVAL_P(return_value), &stat_mtime)
	zend.ZendHashNextIndexInsert(zend.Z_ARRVAL_P(return_value), &stat_ctime)
	zend.ZendHashNextIndexInsert(zend.Z_ARRVAL_P(return_value), &stat_blksize)
	zend.ZendHashNextIndexInsert(zend.Z_ARRVAL_P(return_value), &stat_blocks)

	/* Store string indexes referencing the same zval*/

	zend.ZendHashStrAddNew(zend.Z_ARRVAL_P(return_value), stat_sb_names[0], strlen(stat_sb_names[0]), &stat_dev)
	zend.ZendHashStrAddNew(zend.Z_ARRVAL_P(return_value), stat_sb_names[1], strlen(stat_sb_names[1]), &stat_ino)
	zend.ZendHashStrAddNew(zend.Z_ARRVAL_P(return_value), stat_sb_names[2], strlen(stat_sb_names[2]), &stat_mode)
	zend.ZendHashStrAddNew(zend.Z_ARRVAL_P(return_value), stat_sb_names[3], strlen(stat_sb_names[3]), &stat_nlink)
	zend.ZendHashStrAddNew(zend.Z_ARRVAL_P(return_value), stat_sb_names[4], strlen(stat_sb_names[4]), &stat_uid)
	zend.ZendHashStrAddNew(zend.Z_ARRVAL_P(return_value), stat_sb_names[5], strlen(stat_sb_names[5]), &stat_gid)
	zend.ZendHashStrAddNew(zend.Z_ARRVAL_P(return_value), stat_sb_names[6], strlen(stat_sb_names[6]), &stat_rdev)
	zend.ZendHashStrAddNew(zend.Z_ARRVAL_P(return_value), stat_sb_names[7], strlen(stat_sb_names[7]), &stat_size)
	zend.ZendHashStrAddNew(zend.Z_ARRVAL_P(return_value), stat_sb_names[8], strlen(stat_sb_names[8]), &stat_atime)
	zend.ZendHashStrAddNew(zend.Z_ARRVAL_P(return_value), stat_sb_names[9], strlen(stat_sb_names[9]), &stat_mtime)
	zend.ZendHashStrAddNew(zend.Z_ARRVAL_P(return_value), stat_sb_names[10], strlen(stat_sb_names[10]), &stat_ctime)
	zend.ZendHashStrAddNew(zend.Z_ARRVAL_P(return_value), stat_sb_names[11], strlen(stat_sb_names[11]), &stat_blksize)
	zend.ZendHashStrAddNew(zend.Z_ARRVAL_P(return_value), stat_sb_names[12], strlen(stat_sb_names[12]), &stat_blocks)
}
func ZifCopy(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var source *byte
	var target *byte
	var source_len int
	var target_len int
	var zcontext *zend.Zval = nil
	var context *core.PhpStreamContext
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 3
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
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgPath(_arg, &source, &source_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_PATH
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgPath(_arg, &target, &target_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_PATH
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgResource(_arg, &zcontext, 1) == 0 {
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
			return
		}
		break
	}
	if streams.PhpStreamLocateUrlWrapper(source, nil, 0) == &PhpPlainFilesWrapper && core.PhpCheckOpenBasedir(source) != 0 {
		zend.RETVAL_FALSE
		return
	}
	context = streams.PhpStreamContextFromZval(zcontext, 0)
	if PhpCopyFileCtx(source, target, 0, context) == zend.SUCCESS {
		zend.RETVAL_TRUE
		return
	} else {
		zend.RETVAL_FALSE
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
	var ret int = zend.FAILURE
	var src_s core.PhpStreamStatbuf
	var dest_s core.PhpStreamStatbuf
	switch core.PhpStreamStatPathEx(src, 0, &src_s, ctx) {
	case -1:

		/* non-statable stream */

		goto safe_to_copy
		break
	case 0:
		break
	default:
		return ret
	}
	if zend.S_ISDIR(src_s.sb.st_mode) {
		core.PhpErrorDocref(nil, zend.E_WARNING, "The first argument to copy() function cannot be a directory")
		return zend.FAILURE
	}
	switch core.PhpStreamStatPathEx(dest, core.PHP_STREAM_URL_STAT_QUIET|core.PHP_STREAM_URL_STAT_NOCACHE, &dest_s, ctx) {
	case -1:

		/* non-statable stream */

		goto safe_to_copy
		break
	case 0:
		break
	default:
		return ret
	}
	if zend.S_ISDIR(dest_s.sb.st_mode) {
		core.PhpErrorDocref(nil, zend.E_WARNING, "The second argument to copy() function cannot be a directory")
		return zend.FAILURE
	}
	if !(src_s.sb.st_ino) || !(dest_s.sb.st_ino) {
		goto no_stat
	}
	if src_s.sb.st_ino == dest_s.sb.st_ino && src_s.sb.st_dev == dest_s.sb.st_dev {
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
func ZifFread(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var res *zend.Zval
	var len_ zend.ZendLong
	var stream *core.PhpStream
	var str *zend.ZendString
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
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgResource(_arg, &res, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_RESOURCE
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgLong(_arg, &len_, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
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
			zend.RETVAL_FALSE
			return
		}
		break
	}
	PHP_STREAM_TO_ZVAL(stream, res)
	if len_ <= 0 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Length parameter must be greater than 0")
		zend.RETVAL_FALSE
		return
	}
	str = streams.PhpStreamReadToStr(stream, len_)
	if str == nil {
		zend.ZvalPtrDtorStr(return_value)
		zend.RETVAL_FALSE
		return
	}
	zend.RETVAL_STR(str)
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

		case -1:
			inc_len = 1
			core.PhpIgnoreValue(mblen(nil, 0))
			break
		case 0:
			goto quit_loop
		case 1:

		default:
			last_chars[0] = last_chars[1]
			last_chars[1] = *ptr
			break
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
	case '\r':
		return ptr - 1
	}
	return ptr
}
func FPUTCSV_FLD_CHK(c __auto__) __auto__ {
	return memchr(zend.ZSTR_VAL(field_str), c, zend.ZSTR_LEN(field_str))
}
func ZifFputcsv(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var delimiter byte = ','
	var enclosure byte = '"'
	var escape_char int = uint8('\\')
	var stream *core.PhpStream
	var fp *zend.Zval = nil
	var fields *zend.Zval = nil
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
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgResource(_arg, &fp, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_RESOURCE
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgArray(_arg, &fields, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgString(_arg, &delimiter_str, &delimiter_str_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgString(_arg, &enclosure_str, &enclosure_str_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgString(_arg, &escape_str, &escape_str_len, 0) == 0 {
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
			return
		}
		break
	}
	if delimiter_str != nil {

		/* Make sure that there is at least one character in string */

		if delimiter_str_len < 1 {
			core.PhpErrorDocref(nil, zend.E_WARNING, "delimiter must be a character")
			zend.RETVAL_FALSE
			return
		} else if delimiter_str_len > 1 {
			core.PhpErrorDocref(nil, zend.E_NOTICE, "delimiter must be a single character")
		}

		/* use first character from string */

		delimiter = *delimiter_str

		/* use first character from string */

	}
	if enclosure_str != nil {
		if enclosure_str_len < 1 {
			core.PhpErrorDocref(nil, zend.E_WARNING, "enclosure must be a character")
			zend.RETVAL_FALSE
			return
		} else if enclosure_str_len > 1 {
			core.PhpErrorDocref(nil, zend.E_NOTICE, "enclosure must be a single character")
		}

		/* use first character from string */

		enclosure = *enclosure_str

		/* use first character from string */

	}
	if escape_str != nil {
		if escape_str_len > 1 {
			core.PhpErrorDocref(nil, zend.E_NOTICE, "escape must be empty or a single character")
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
		zend.RETVAL_FALSE
		return
	}
	zend.RETVAL_LONG(ret)
	return
}
func PhpFputcsv(stream *core.PhpStream, fields *zend.Zval, delimiter byte, enclosure byte, escape_char int) ssize_t {
	var count int
	var i int = 0
	var ret int
	var field_tmp *zend.Zval
	var csvline zend.SmartStr = zend.SmartStr{0}
	zend.ZEND_ASSERT(escape_char >= 0 && escape_char <= UCHAR_MAX || escape_char == PHP_CSV_NO_ESCAPE)
	count = zend.ZendHashNumElements(zend.Z_ARRVAL_P(fields))
	for {
		var __ht *zend.HashTable = zend.Z_ARRVAL_P(fields)
		var _p *zend.Bucket = __ht.GetArData()
		var _end *zend.Bucket = _p + __ht.GetNNumUsed()
		for ; _p != _end; _p++ {
			var _z *zend.Zval = &_p.GetVal()

			if zend.Z_TYPE_P(_z) == zend.IS_UNDEF {
				continue
			}
			field_tmp = _z
			var tmp_field_str *zend.ZendString
			var field_str *zend.ZendString = zend.ZvalGetTmpString(field_tmp, &tmp_field_str)

			/* enclose a field that contains a delimiter, an enclosure character, or a newline */

			if FPUTCSV_FLD_CHK(delimiter) || FPUTCSV_FLD_CHK(enclosure) || escape_char != PHP_CSV_NO_ESCAPE && FPUTCSV_FLD_CHK(escape_char) || FPUTCSV_FLD_CHK('\n') || FPUTCSV_FLD_CHK('\r') || FPUTCSV_FLD_CHK('\t') || FPUTCSV_FLD_CHK(' ') {
				var ch *byte = zend.ZSTR_VAL(field_str)
				var end *byte = ch + zend.ZSTR_LEN(field_str)
				var escaped int = 0
				zend.SmartStrAppendc(&csvline, enclosure)
				for ch < end {
					if escape_char != PHP_CSV_NO_ESCAPE && (*ch) == escape_char {
						escaped = 1
					} else if escaped == 0 && (*ch) == enclosure {
						zend.SmartStrAppendc(&csvline, enclosure)
					} else {
						escaped = 0
					}
					zend.SmartStrAppendc(&csvline, *ch)
					ch++
				}
				zend.SmartStrAppendc(&csvline, enclosure)
			} else {
				zend.SmartStrAppend(&csvline, field_str)
			}
			if b.PreInc(&i) != count {
				zend.SmartStrAppendl(&csvline, &delimiter, 1)
			}
			zend.ZendTmpStringRelease(tmp_field_str)
		}
		break
	}
	zend.SmartStrAppendc(&csvline, '\n')
	zend.SmartStr0(&csvline)
	ret = core.PhpStreamWrite(stream, zend.ZSTR_VAL(csvline.GetS()), zend.ZSTR_LEN(csvline.GetS()))
	zend.SmartStrFree(&csvline)
	return ret
}
func ZifFgetcsv(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var delimiter byte = ','
	var enclosure byte = '"'
	var escape int = uint8('\\')

	/* first section exactly as php_fgetss */

	var len_ zend.ZendLong = 0
	var buf_len int
	var buf *byte
	var stream *core.PhpStream
	var fd *zend.Zval
	var len_zv *zend.Zval = nil
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
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgResource(_arg, &fd, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_RESOURCE
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			zend.ZendParseArgZvalDeref(_arg, &len_zv, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgString(_arg, &delimiter_str, &delimiter_str_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgString(_arg, &enclosure_str, &enclosure_str_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgString(_arg, &escape_str, &escape_str_len, 0) == 0 {
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
			return
		}
		break
	}
	if delimiter_str != nil {

		/* Make sure that there is at least one character in string */

		if delimiter_str_len < 1 {
			core.PhpErrorDocref(nil, zend.E_WARNING, "delimiter must be a character")
			zend.RETVAL_FALSE
			return
		} else if delimiter_str_len > 1 {
			core.PhpErrorDocref(nil, zend.E_NOTICE, "delimiter must be a single character")
		}

		/* use first character from string */

		delimiter = delimiter_str[0]

		/* use first character from string */

	}
	if enclosure_str != nil {
		if enclosure_str_len < 1 {
			core.PhpErrorDocref(nil, zend.E_WARNING, "enclosure must be a character")
			zend.RETVAL_FALSE
			return
		} else if enclosure_str_len > 1 {
			core.PhpErrorDocref(nil, zend.E_NOTICE, "enclosure must be a single character")
		}

		/* use first character from string */

		enclosure = enclosure_str[0]

		/* use first character from string */

	}
	if escape_str != nil {
		if escape_str_len > 1 {
			core.PhpErrorDocref(nil, zend.E_NOTICE, "escape must be empty or a single character")
		}
		if escape_str_len < 1 {
			escape = PHP_CSV_NO_ESCAPE
		} else {
			escape = uint8(escape_str[0])
		}
	}
	if len_zv != nil && zend.Z_TYPE_P(len_zv) != zend.IS_NULL {
		len_ = zend.ZvalGetLong(len_zv)
		if len_ < 0 {
			core.PhpErrorDocref(nil, zend.E_WARNING, "Length parameter may not be negative")
			zend.RETVAL_FALSE
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
			zend.RETVAL_FALSE
			return
		}
	} else {
		buf = zend.Emalloc(len_ + 1)
		if core.PhpStreamGetLine(stream, buf, len_+1, &buf_len) == nil {
			zend.Efree(buf)
			zend.RETVAL_FALSE
			return
		}
	}
	PhpFgetcsv(stream, delimiter, enclosure, escape, buf_len, buf, return_value)
}
func PhpFgetcsv(stream *core.PhpStream, delimiter byte, enclosure byte, escape_char int, buf_len int, buf *byte, return_value *zend.Zval) {
	var temp *byte
	var tptr *byte
	var bptr *byte
	var line_end *byte
	var limit *byte
	var temp_len int
	var line_end_len int
	var inc_len int
	var first_field zend.ZendBool = 1
	zend.ZEND_ASSERT(escape_char >= 0 && escape_char <= UCHAR_MAX || escape_char == PHP_CSV_NO_ESCAPE)

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
							zend.ZendArrayDestroy(zend.Z_ARR_P(return_value))
							zend.RETVAL_FALSE
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
						break
					}
					break
				case -2:

				case -1:
					core.PhpIgnoreValue(mblen(nil, 0))
				case 1:

					/* we need to determine if the enclosure is
					 * 'real' or is it escaped */

					switch state {
					case 1:
						bptr++
						state = 0
						break
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
						break
					default:
						if (*bptr) == enclosure {
							state = 2
						} else if escape_char != PHP_CSV_NO_ESCAPE && (*bptr) == escape_char {
							state = 1
						}
						bptr++
						break
					}
					break
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
						break
					default:
						bptr += inc_len
						break
					}
					break
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

				case -1:
					inc_len = 1
					core.PhpIgnoreValue(mblen(nil, 0))
				case 1:
					if (*bptr) == delimiter {
						goto quit_loop_3
					}
					break
				default:
					break
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

				case -1:
					inc_len = 1
					core.PhpIgnoreValue(mblen(nil, 0))
				case 1:
					if (*bptr) == delimiter {
						goto quit_loop_4
					}
					break
				default:
					break
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
func ZifRealpath(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var filename *byte
	var filename_len int
	var resolved_path_buff []byte
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
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgPath(_arg, &filename, &filename_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_PATH
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
			return
		}
		break
	}
	if zend.VCWD_REALPATH(filename, resolved_path_buff) != nil {
		if core.PhpCheckOpenBasedir(resolved_path_buff) != 0 {
			zend.RETVAL_FALSE
			return
		}
		zend.RETVAL_STRING(resolved_path_buff)
		return
	} else {
		zend.RETVAL_FALSE
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
			break
		case '>':
			return TOK_CLOSETAG
			break
		case '=':
			return TOK_EQUAL
			break
		case '/':
			return TOK_SLASH
			break
		case '\'':

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
			break
		case '\n':

		case '\r':

		case '\t':
			break
		case ' ':
			return TOK_SPACE
			break
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
			break
		}
	}
	return TOK_EOF
}
func ZifFnmatch(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var pattern *byte
	var filename *byte
	var pattern_len int
	var filename_len int
	var flags zend.ZendLong = 0
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 3
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
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgPath(_arg, &pattern, &pattern_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_PATH
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgPath(_arg, &filename, &filename_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_PATH
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgLong(_arg, &flags, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
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
			return
		}
		break
	}
	if filename_len >= core.MAXPATHLEN {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Filename exceeds the maximum allowed length of %d characters", core.MAXPATHLEN)
		zend.RETVAL_FALSE
		return
	}
	if pattern_len >= core.MAXPATHLEN {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Pattern exceeds the maximum allowed length of %d characters", core.MAXPATHLEN)
		zend.RETVAL_FALSE
		return
	}
	zend.RETVAL_BOOL(!(fnmatch(pattern, filename, int(flags))))
	return
}
func ZifSysGetTempDir(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	zend.RETVAL_STRING((*byte)(core.PhpGetTemporaryDirectory()))
	return
}
