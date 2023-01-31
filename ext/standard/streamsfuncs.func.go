// <<generate>>

package standard

import (
	b "sik/builtin"
	"sik/core"
	"sik/core/streams"
	r "sik/runtime"
	"sik/zend"
)

func PhpSelect(m core.PhpSocketT, r *fd_set, w *fd_set, e *fd_set, t *__struct__timeval) __auto__ {
	return select_(m, r, w, e, t)
}
func GET_CTX_OPT(stream *core.PhpStream, wrapper string, name string, val *zend.Zval) bool {
	return core.PHP_STREAM_CONTEXT(stream) != nil && nil != b.Assign(&val, streams.PhpStreamContextGetOption(core.PHP_STREAM_CONTEXT(stream), wrapper, name))
}
func ZifStreamSocketPair(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var domain zend.ZendLong
	var type_ zend.ZendLong
	var protocol zend.ZendLong
	var s1 *core.PhpStream
	var s2 *core.PhpStream
	var pair []core.PhpSocketT
	for {
		var _flags int = 0
		var _min_num_args int = 3
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
			if zend.ZendParseArgLong(_arg, &domain, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgLong(_arg, &type_, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgLong(_arg, &protocol, &_dummy, 0, 0) == 0 {
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
	if 0 != socketpair(int(domain), int(type_), int(protocol), pair) {
		var errbuf []byte
		core.PhpErrorDocref(nil, zend.E_WARNING, "failed to create sockets: [%d]: %s", core.PhpSocketErrno(), core.PhpSocketStrerror(core.PhpSocketErrno(), errbuf, b.SizeOf("errbuf")))
		zend.RETVAL_FALSE
		return
	}
	zend.ArrayInit(return_value)
	s1 = core.PhpStreamSockOpenFromSocket(pair[0], 0)
	s2 = core.PhpStreamSockOpenFromSocket(pair[1], 0)

	/* set the __exposed flag.
	 * php_stream_to_zval() does, add_next_index_resource() does not */

	core.PhpStreamAutoCleanup(s1)
	core.PhpStreamAutoCleanup(s2)
	zend.AddNextIndexResource(return_value, s1.GetRes())
	zend.AddNextIndexResource(return_value, s2.GetRes())
}
func ZifStreamSocketClient(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var host *zend.ZendString
	var zerrno *zend.Zval = nil
	var zerrstr *zend.Zval = nil
	var zcontext *zend.Zval = nil
	var timeout float64 = float64(FG(default_socket_timeout))
	var conv PhpTimeoutUll
	var tv __struct__timeval
	var hashkey *byte = nil
	var stream *core.PhpStream = nil
	var err int
	var flags zend.ZendLong = PHP_STREAM_CLIENT_CONNECT
	var errstr *zend.ZendString = nil
	var context *core.PhpStreamContext = nil
	zend.RETVAL_FALSE
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 6
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
			if zend.ZendParseArgStr(_arg, &host, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			zend.ZendParseArgZvalDeref(_arg, &zerrno, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			zend.ZendParseArgZvalDeref(_arg, &zerrstr, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgDouble(_arg, &timeout, &_dummy, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_DOUBLE
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgLong(_arg, &flags, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgResource(_arg, &zcontext, 0) == 0 {
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
	context = streams.PhpStreamContextFromZval(zcontext, flags&PHP_FILE_NO_DEFAULT_CONTEXT)
	if (flags & PHP_STREAM_CLIENT_PERSISTENT) != 0 {
		core.Spprintf(&hashkey, 0, "stream_socket_client__%s", host.GetVal())
	}

	/* prepare the timeout value for use */

	conv = php_timeout_ull(timeout * 1000000.0)
	tv.tv_sec = conv / 1000000
	tv.tv_usec = conv % 1000000
	if zerrno != nil {
		zend.ZEND_TRY_ASSIGN_REF_LONG(zerrno, 0)
	}
	if zerrstr != nil {
		zend.ZEND_TRY_ASSIGN_REF_EMPTY_STRING(zerrstr)
	}
	stream = streams.PhpStreamXportCreate(host.GetVal(), host.GetLen(), core.REPORT_ERRORS, streams.STREAM_XPORT_CLIENT|b.Cond((flags&PHP_STREAM_CLIENT_CONNECT) != 0, streams.STREAM_XPORT_CONNECT, 0)|b.Cond((flags&PHP_STREAM_CLIENT_ASYNC_CONNECT) != 0, streams.STREAM_XPORT_CONNECT_ASYNC, 0), hashkey, &tv, context, &errstr, &err)
	if stream == nil {

		/* host might contain binary characters */

		var quoted_host *zend.ZendString = PhpAddslashes(host)
		core.PhpErrorDocref(nil, zend.E_WARNING, "unable to connect to %s (%s)", quoted_host.GetVal(), b.CondF2(errstr == nil, "Unknown error", func() []byte { return errstr.GetVal() }))
		zend.ZendStringReleaseEx(quoted_host, 0)
	}
	if hashkey != nil {
		zend.Efree(hashkey)
	}
	if stream == nil {
		if zerrno != nil {
			zend.ZEND_TRY_ASSIGN_REF_LONG(zerrno, err)
		}
		if zerrstr != nil && errstr != nil {
			zend.ZEND_TRY_ASSIGN_REF_STR(zerrstr, errstr)
		} else if errstr != nil {
			zend.ZendStringReleaseEx(errstr, 0)
		}
		zend.RETVAL_FALSE
		return
	}
	if errstr != nil {
		zend.ZendStringReleaseEx(errstr, 0)
	}
	core.PhpStreamToZval(stream, return_value)
}
func ZifStreamSocketServer(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var host *byte
	var host_len int
	var zerrno *zend.Zval = nil
	var zerrstr *zend.Zval = nil
	var zcontext *zend.Zval = nil
	var stream *core.PhpStream = nil
	var err int = 0
	var flags zend.ZendLong = streams.STREAM_XPORT_BIND | streams.STREAM_XPORT_LISTEN
	var errstr *zend.ZendString = nil
	var context *core.PhpStreamContext = nil
	zend.RETVAL_FALSE
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
			if zend.ZendParseArgString(_arg, &host, &host_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			zend.ZendParseArgZvalDeref(_arg, &zerrno, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			zend.ZendParseArgZvalDeref(_arg, &zerrstr, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgLong(_arg, &flags, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgResource(_arg, &zcontext, 0) == 0 {
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
	context = streams.PhpStreamContextFromZval(zcontext, flags&PHP_FILE_NO_DEFAULT_CONTEXT)
	if context != nil {
		context.GetRes().AddRefcount()
	}
	if zerrno != nil {
		zend.ZEND_TRY_ASSIGN_REF_LONG(zerrno, 0)
	}
	if zerrstr != nil {
		zend.ZEND_TRY_ASSIGN_REF_EMPTY_STRING(zerrstr)
	}
	stream = streams.PhpStreamXportCreate(host, host_len, core.REPORT_ERRORS, streams.STREAM_XPORT_SERVER|int(flags), nil, nil, context, &errstr, &err)
	if stream == nil {
		core.PhpErrorDocref(nil, zend.E_WARNING, "unable to connect to %s (%s)", host, b.CondF2(errstr == nil, "Unknown error", func() []byte { return errstr.GetVal() }))
	}
	if stream == nil {
		if zerrno != nil {
			zend.ZEND_TRY_ASSIGN_REF_LONG(zerrno, err)
		}
		if zerrstr != nil && errstr != nil {
			zend.ZEND_TRY_ASSIGN_REF_STR(zerrstr, errstr)
		} else if errstr != nil {
			zend.ZendStringReleaseEx(errstr, 0)
		}
		zend.RETVAL_FALSE
		return
	}
	if errstr != nil {
		zend.ZendStringReleaseEx(errstr, 0)
	}
	core.PhpStreamToZval(stream, return_value)
}
func ZifStreamSocketAccept(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var timeout float64 = float64(FG(default_socket_timeout))
	var zpeername *zend.Zval = nil
	var peername *zend.ZendString = nil
	var conv PhpTimeoutUll
	var tv __struct__timeval
	var stream *core.PhpStream = nil
	var clistream *core.PhpStream = nil
	var zstream *zend.Zval
	var errstr *zend.ZendString = nil
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
			if zend.ZendParseArgResource(_arg, &zstream, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_RESOURCE
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgDouble(_arg, &timeout, &_dummy, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_DOUBLE
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			zend.ZendParseArgZvalDeref(_arg, &zpeername, 0)
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
	core.PhpStreamFromZval(stream, zstream)

	/* prepare the timeout value for use */

	conv = php_timeout_ull(timeout * 1000000.0)
	tv.tv_sec = conv / 1000000
	tv.tv_usec = conv % 1000000
	if 0 == streams.PhpStreamXportAccept(stream, &clistream, b.Cond(zpeername != nil, &peername, nil), nil, nil, &tv, &errstr) && clistream != nil {
		if peername != nil {
			zend.ZEND_TRY_ASSIGN_REF_STR(zpeername, peername)
		}
		core.PhpStreamToZval(clistream, return_value)
	} else {
		if peername != nil {
			zend.ZendStringRelease(peername)
		}
		core.PhpErrorDocref(nil, zend.E_WARNING, "accept failed: %s", b.CondF1(errstr != nil, func() []byte { return errstr.GetVal() }, "Unknown error"))
		zend.RETVAL_FALSE
	}
	if errstr != nil {
		zend.ZendStringReleaseEx(errstr, 0)
	}
}
func ZifStreamSocketGetName(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var stream *core.PhpStream
	var zstream *zend.Zval
	var want_peer zend.ZendBool
	var name *zend.ZendString = nil
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
			if zend.ZendParseArgResource(_arg, &zstream, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_RESOURCE
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgBool(_arg, &want_peer, &_dummy, 0) == 0 {
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
			zend.RETVAL_FALSE
			return
		}
		break
	}
	core.PhpStreamFromZval(stream, zstream)
	if 0 != streams.PhpStreamXportGetName(stream, want_peer, &name, nil, nil) || name == nil {
		zend.RETVAL_FALSE
		return
	}
	if name.GetLen() == 0 || name.GetVal()[0] == 0 {
		zend.ZendStringReleaseEx(name, 0)
		zend.RETVAL_FALSE
		return
	}
	zend.RETVAL_STR(name)
}
func ZifStreamSocketSendto(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var stream *core.PhpStream
	var zstream *zend.Zval
	var flags zend.ZendLong = 0
	var data *byte
	var target_addr *byte = nil
	var datalen int
	var target_addr_len int = 0
	var sa core.PhpSockaddrStorage
	var sl socklen_t = 0
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
			if zend.ZendParseArgResource(_arg, &zstream, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_RESOURCE
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgString(_arg, &data, &datalen, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
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
			if zend.ZendParseArgString(_arg, &target_addr, &target_addr_len, 0) == 0 {
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
	core.PhpStreamFromZval(stream, zstream)
	if target_addr_len != 0 {

		/* parse the address */

		if zend.FAILURE == core.PhpNetworkParseNetworkAddressWithPort(target_addr, target_addr_len, (*__struct__sockaddr)(&sa), &sl) {
			core.PhpErrorDocref(nil, zend.E_WARNING, "Failed to parse `%s' into a valid network address", target_addr)
			zend.RETVAL_FALSE
			return
		}

		/* parse the address */

	}
	zend.RETVAL_LONG(streams.PhpStreamXportSendto(stream, data, datalen, int(flags), b.Cond(target_addr_len != 0, &sa, nil), sl))
	return
}
func ZifStreamSocketRecvfrom(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var stream *core.PhpStream
	var zstream *zend.Zval
	var zremote *zend.Zval = nil
	var remote_addr *zend.ZendString = nil
	var to_read zend.ZendLong = 0
	var read_buf *zend.ZendString
	var flags zend.ZendLong = 0
	var recvd int
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
			if zend.ZendParseArgResource(_arg, &zstream, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_RESOURCE
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgLong(_arg, &to_read, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
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
			zend.ZendParseArgZvalDeref(_arg, &zremote, 0)
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
	core.PhpStreamFromZval(stream, zstream)
	if zremote != nil {
		zend.ZEND_TRY_ASSIGN_REF_NULL(zremote)
	}
	if to_read <= 0 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Length parameter must be greater than 0")
		zend.RETVAL_FALSE
		return
	}
	read_buf = zend.ZendStringAlloc(to_read, 0)
	recvd = streams.PhpStreamXportRecvfrom(stream, read_buf.GetVal(), to_read, int(flags), nil, nil, b.Cond(zremote != nil, &remote_addr, nil))
	if recvd >= 0 {
		if zremote != nil && remote_addr != nil {
			zend.ZEND_TRY_ASSIGN_REF_STR(zremote, remote_addr)
		}
		read_buf.GetVal()[recvd] = '0'
		read_buf.SetLen(recvd)
		zend.RETVAL_NEW_STR(read_buf)
		return
	}
	zend.ZendStringEfree(read_buf)
	zend.RETVAL_FALSE
	return
}
func ZifStreamGetContents(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var stream *core.PhpStream
	var zsrc *zend.Zval
	var maxlen zend.ZendLong = ssize_t(core.PHP_STREAM_COPY_ALL)
	var desiredpos zend.ZendLong = -1
	var contents *zend.ZendString
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
			if zend.ZendParseArgResource(_arg, &zsrc, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_RESOURCE
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
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgLong(_arg, &desiredpos, &_dummy, 0, 0) == 0 {
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
	if maxlen < 0 && maxlen != ssize_t(core.PHP_STREAM_COPY_ALL) {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Length must be greater than or equal to zero, or -1")
		zend.RETVAL_FALSE
		return
	}
	core.PhpStreamFromZval(stream, zsrc)
	if desiredpos >= 0 {
		var seek_res int = 0
		var position zend.ZendOffT
		position = core.PhpStreamTell(stream)
		if position >= 0 && desiredpos > position {

			/* use SEEK_CUR to allow emulation in streams that don't support seeking */

			seek_res = core.PhpStreamSeek(stream, desiredpos-position, r.SEEK_CUR)

			/* use SEEK_CUR to allow emulation in streams that don't support seeking */

		} else if desiredpos < position {

			/* desired position before position or error on tell */

			seek_res = core.PhpStreamSeek(stream, desiredpos, r.SEEK_SET)

			/* desired position before position or error on tell */

		}
		if seek_res != 0 {
			core.PhpErrorDocref(nil, zend.E_WARNING, "Failed to seek to position "+zend.ZEND_LONG_FMT+" in the stream", desiredpos)
			zend.RETVAL_FALSE
			return
		}
	}
	if b.Assign(&contents, core.PhpStreamCopyToMem(stream, maxlen, 0)) {
		zend.RETVAL_STR(contents)
		return
	} else {
		zend.RETVAL_EMPTY_STRING()
		return
	}
}
func ZifStreamCopyToStream(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var src *core.PhpStream
	var dest *core.PhpStream
	var zsrc *zend.Zval
	var zdest *zend.Zval
	var maxlen zend.ZendLong = core.PHP_STREAM_COPY_ALL
	var pos zend.ZendLong = 0
	var len_ int
	var ret int
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
			if zend.ZendParseArgResource(_arg, &zsrc, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_RESOURCE
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgResource(_arg, &zdest, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_RESOURCE
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
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgLong(_arg, &pos, &_dummy, 0, 0) == 0 {
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
	core.PhpStreamFromZval(src, zsrc)
	core.PhpStreamFromZval(dest, zdest)
	if pos > 0 && core.PhpStreamSeek(src, pos, r.SEEK_SET) < 0 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Failed to seek to position "+zend.ZEND_LONG_FMT+" in the stream", pos)
		zend.RETVAL_FALSE
		return
	}
	ret = core.PhpStreamCopyToStreamEx(src, dest, maxlen, &len_)
	if ret != zend.SUCCESS {
		zend.RETVAL_FALSE
		return
	}
	zend.RETVAL_LONG(len_)
	return
}
func ZifStreamGetMetaData(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var zstream *zend.Zval
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
			if zend.ZendParseArgResource(_arg, &zstream, 0) == 0 {
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
	core.PhpStreamFromZval(stream, zstream)
	zend.ArrayInit(return_value)
	if core.PhpStreamPopulateMetaData(stream, return_value) == 0 {
		zend.AddAssocBool(return_value, "timed_out", 0)
		zend.AddAssocBool(return_value, "blocked", 1)
		zend.AddAssocBool(return_value, "eof", core.PhpStreamEof(stream))
	}
	if !(zend.Z_ISUNDEF(stream.GetWrapperdata())) {
		zend.Z_ADDREF_P(stream.GetWrapperdata())
		zend.AddAssocZval(return_value, "wrapper_data", stream.GetWrapperdata())
	}
	if stream.GetWrapper() != nil {
		zend.AddAssocString(return_value, "wrapper_type", (*byte)(stream.GetWrapper().GetWops().GetLabel()))
	}
	zend.AddAssocString(return_value, "stream_type", (*byte)(stream.GetOps().GetLabel()))
	zend.AddAssocString(return_value, "mode", stream.GetMode())
	zend.AddAssocLong(return_value, "unread_bytes", stream.GetWritepos()-stream.GetReadpos())
	zend.AddAssocBool(return_value, "seekable", stream.GetOps().GetSeek() != nil && !stream.HasFlags(core.PHP_STREAM_FLAG_NO_SEEK))
	if stream.GetOrigPath() != nil {
		zend.AddAssocString(return_value, "uri", stream.GetOrigPath())
	}
}
func ZifStreamGetTransports(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var stream_xport_hash *zend.HashTable
	var stream_xport *zend.ZendString
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	if b.Assign(&stream_xport_hash, streams.PhpStreamXportGetHash()) {
		zend.ArrayInit(return_value)
		var __ht *zend.HashTable = stream_xport_hash
		for _, _p := range __ht.foreachData() {
			var _z *zend.Zval = _p.GetVal()

			stream_xport = _p.GetKey()
			zend.AddNextIndexStr(return_value, stream_xport.Copy())
		}
	} else {
		zend.RETVAL_FALSE
		return
	}
}
func ZifStreamGetWrappers(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var url_stream_wrappers_hash *zend.HashTable
	var stream_protocol *zend.ZendString
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	if b.Assign(&url_stream_wrappers_hash, core.PhpStreamGetUrlStreamWrappersHash()) {
		zend.ArrayInit(return_value)
		var __ht *zend.HashTable = url_stream_wrappers_hash
		for _, _p := range __ht.foreachData() {
			var _z *zend.Zval = _p.GetVal()

			stream_protocol = _p.GetKey()
			if stream_protocol != nil {
				zend.AddNextIndexStr(return_value, stream_protocol.Copy())
			}
		}
	} else {
		zend.RETVAL_FALSE
		return
	}
}
func StreamArrayToFdSet(stream_array *zend.Zval, fds *fd_set, max_fd *core.PhpSocketT) int {
	var elem *zend.Zval
	var stream *core.PhpStream
	var cnt int = 0
	if stream_array.GetType() != zend.IS_ARRAY {
		return 0
	}
	var __ht *zend.HashTable = stream_array.GetArr()
	for _, _p := range __ht.foreachData() {
		var _z *zend.Zval = _p.GetVal()

		elem = _z

		/* Temporary int fd is needed for the STREAM data type on windows, passing this_fd directly to php_stream_cast()
		   would eventually bring a wrong result on x64. php_stream_cast() casts to int internally, and this will leave
		   the higher bits of a SOCKET variable uninitialized on systems with little endian. */

		var this_fd core.PhpSocketT
		zend.ZVAL_DEREF(elem)
		core.PhpStreamFromZvalNoVerify(stream, elem)
		if stream == nil {
			continue
		}

		/* get the fd.
		 * NB: Most other code will NOT use the PHP_STREAM_CAST_INTERNAL flag
		 * when casting.  It is only used here so that the buffered data warning
		 * is not displayed.
		 * */

		if zend.SUCCESS == core.PhpStreamCast(stream, core.PHP_STREAM_AS_FD_FOR_SELECT|core.PHP_STREAM_CAST_INTERNAL, any(&this_fd), 1) && this_fd != -1 {
			core.PHP_SAFE_FD_SET(this_fd, fds)
			if this_fd > (*max_fd) {
				*max_fd = this_fd
			}
			cnt++
		}

		/* get the fd.
		 * NB: Most other code will NOT use the PHP_STREAM_CAST_INTERNAL flag
		 * when casting.  It is only used here so that the buffered data warning
		 * is not displayed.
		 * */

	}
	if cnt != 0 {
		return 1
	} else {
		return 0
	}
}
func StreamArrayFromFdSet(stream_array *zend.Zval, fds *fd_set) int {
	var elem *zend.Zval
	var dest_elem *zend.Zval
	var ht *zend.HashTable
	var stream *core.PhpStream
	var ret int = 0
	var key *zend.ZendString
	var num_ind zend.ZendUlong
	if stream_array.GetType() != zend.IS_ARRAY {
		return 0
	}
	ht = zend.ZendNewArray(zend.Z_ARRVAL_P(stream_array).GetNNumOfElements())
	var __ht *zend.HashTable = stream_array.GetArr()
	for _, _p := range __ht.foreachData() {
		var _z *zend.Zval = _p.GetVal()

		num_ind = _p.GetH()
		key = _p.GetKey()
		elem = _z
		var this_fd core.PhpSocketT
		zend.ZVAL_DEREF(elem)
		core.PhpStreamFromZvalNoVerify(stream, elem)
		if stream == nil {
			continue
		}

		/* get the fd
		 * NB: Most other code will NOT use the PHP_STREAM_CAST_INTERNAL flag
		 * when casting.  It is only used here so that the buffered data warning
		 * is not displayed.
		 */

		if zend.SUCCESS == core.PhpStreamCast(stream, core.PHP_STREAM_AS_FD_FOR_SELECT|core.PHP_STREAM_CAST_INTERNAL, any(&this_fd), 1) && this_fd != core.SOCK_ERR {
			if core.PHP_SAFE_FD_ISSET(this_fd, fds) {
				if key == nil {
					dest_elem = ht.IndexUpdateH(num_ind, elem)
				} else {
					dest_elem = ht.KeyUpdate(key.GetStr(), elem)
				}
				zend.ZvalAddRef(dest_elem)
				ret++
				continue
			}
		}

		/* get the fd
		 * NB: Most other code will NOT use the PHP_STREAM_CAST_INTERNAL flag
		 * when casting.  It is only used here so that the buffered data warning
		 * is not displayed.
		 */

	}

	/* destroy old array and add new one */

	zend.ZvalPtrDtor(stream_array)
	zend.ZVAL_ARR(stream_array, ht)
	return ret
}
func StreamArrayEmulateReadFdSet(stream_array *zend.Zval) int {
	var elem *zend.Zval
	var dest_elem *zend.Zval
	var ht *zend.HashTable
	var stream *core.PhpStream
	var ret int = 0
	var num_ind zend.ZendUlong
	var key *zend.ZendString
	if stream_array.GetType() != zend.IS_ARRAY {
		return 0
	}
	ht = zend.ZendNewArray(zend.Z_ARRVAL_P(stream_array).GetNNumOfElements())
	var __ht *zend.HashTable = stream_array.GetArr()
	for _, _p := range __ht.foreachData() {
		var _z *zend.Zval = _p.GetVal()

		num_ind = _p.GetH()
		key = _p.GetKey()
		elem = _z
		zend.ZVAL_DEREF(elem)
		core.PhpStreamFromZvalNoVerify(stream, elem)
		if stream == nil {
			continue
		}
		if stream.GetWritepos()-stream.GetReadpos() > 0 {

			/* allow readable non-descriptor based streams to participate in stream_select.
			 * Non-descriptor streams will only "work" if they have previously buffered the
			 * data.  Not ideal, but better than nothing.
			 * This branch of code also allows blocking streams with buffered data to
			 * operate correctly in stream_select.
			 * */

			if key == nil {
				dest_elem = ht.IndexUpdateH(num_ind, elem)
			} else {
				dest_elem = ht.KeyUpdate(key.GetStr(), elem)
			}
			zend.ZvalAddRef(dest_elem)
			ret++
			continue
		}
	}
	if ret > 0 {

		/* destroy old array and add new one */

		zend.ZvalPtrDtor(stream_array)
		zend.ZVAL_ARR(stream_array, ht)
	} else {
		ht.DestroyEx()
	}
	return ret
}
func ZifStreamSelect(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var r_array *zend.Zval
	var w_array *zend.Zval
	var e_array *zend.Zval
	var tv __struct__timeval
	var tv_p *__struct__timeval = nil
	var rfds fd_set
	var wfds fd_set
	var efds fd_set
	var max_fd core.PhpSocketT = 0
	var retval int
	var sets int = 0
	var sec zend.ZendLong
	var usec zend.ZendLong = 0
	var secnull zend.ZendBool
	var set_count int
	var max_set_count int = 0
	for {
		var _flags int = 0
		var _min_num_args int = 4
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
			zend.Z_PARAM_PROLOGUE(1, 0)
			if zend.ZendParseArgArray(_arg, &r_array, 1, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(1, 0)
			if zend.ZendParseArgArray(_arg, &w_array, 1, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(1, 0)
			if zend.ZendParseArgArray(_arg, &e_array, 1, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgLong(_arg, &sec, &secnull, 1, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgLong(_arg, &usec, &_dummy, 0, 0) == 0 {
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
	FD_ZERO(&rfds)
	FD_ZERO(&wfds)
	FD_ZERO(&efds)
	if r_array != nil {
		set_count = StreamArrayToFdSet(r_array, &rfds, &max_fd)
		if set_count > max_set_count {
			max_set_count = set_count
		}
		sets += set_count
	}
	if w_array != nil {
		set_count = StreamArrayToFdSet(w_array, &wfds, &max_fd)
		if set_count > max_set_count {
			max_set_count = set_count
		}
		sets += set_count
	}
	if e_array != nil {
		set_count = StreamArrayToFdSet(e_array, &efds, &max_fd)
		if set_count > max_set_count {
			max_set_count = set_count
		}
		sets += set_count
	}
	if sets == 0 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "No stream arrays were passed")
		zend.RETVAL_FALSE
		return
	}
	core.PHP_SAFE_MAX_FD(max_fd, max_set_count)

	/* If seconds is not set to null, build the timeval, else we wait indefinitely */

	if secnull == 0 {
		if sec < 0 {
			core.PhpErrorDocref(nil, zend.E_WARNING, "The seconds parameter must be greater than 0")
			zend.RETVAL_FALSE
			return
		} else if usec < 0 {
			core.PhpErrorDocref(nil, zend.E_WARNING, "The microseconds parameter must be greater than 0")
			zend.RETVAL_FALSE
			return
		}

		/* Windows, Solaris and BSD do not like microsecond values which are >= 1 sec */

		tv.tv_sec = long(sec + usec/1000000)
		tv.tv_usec = long(usec % 1000000)
		tv_p = &tv
	}

	/* slight hack to support buffered data; if there is data sitting in the
	 * read buffer of any of the streams in the read array, let's pretend
	 * that we selected, but return only the readable sockets */

	if r_array != nil {
		retval = StreamArrayEmulateReadFdSet(r_array)
		if retval > 0 {
			if w_array != nil {
				zend.ZvalPtrDtor(w_array)
				zend.ZVAL_EMPTY_ARRAY(w_array)
			}
			if e_array != nil {
				zend.ZvalPtrDtor(e_array)
				zend.ZVAL_EMPTY_ARRAY(e_array)
			}
			zend.RETVAL_LONG(retval)
			return
		}
	}
	retval = PhpSelect(max_fd+1, &rfds, &wfds, &efds, tv_p)
	if retval == -1 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "unable to select [%d]: %s (max_fd=%d)", errno, strerror(errno), max_fd)
		zend.RETVAL_FALSE
		return
	}
	if r_array != nil {
		StreamArrayFromFdSet(r_array, &rfds)
	}
	if w_array != nil {
		StreamArrayFromFdSet(w_array, &wfds)
	}
	if e_array != nil {
		StreamArrayFromFdSet(e_array, &efds)
	}
	zend.RETVAL_LONG(retval)
	return
}
func UserSpaceStreamNotifier(context *core.PhpStreamContext, notifycode int, severity int, xmsg *byte, xcode int, bytes_sofar int, bytes_max int, ptr any) {
	var callback *zend.Zval = context.GetNotifier().GetPtr()
	var retval zend.Zval
	var zvs []zend.Zval
	var i int
	zend.ZVAL_LONG(&zvs[0], notifycode)
	zend.ZVAL_LONG(&zvs[1], severity)
	if xmsg != nil {
		zend.ZVAL_STRING(&zvs[2], xmsg)
	} else {
		zend.ZVAL_NULL(&zvs[2])
	}
	zend.ZVAL_LONG(&zvs[3], xcode)
	zend.ZVAL_LONG(&zvs[4], bytes_sofar)
	zend.ZVAL_LONG(&zvs[5], bytes_max)
	if zend.FAILURE == zend.CallUserFunctionEx(nil, nil, callback, &retval, 6, zvs, 0, nil) {
		core.PhpErrorDocref(nil, zend.E_WARNING, "failed to call user notifier")
	}
	for i = 0; i < 6; i++ {
		zend.ZvalPtrDtor(&zvs[i])
	}
	zend.ZvalPtrDtor(&retval)
}
func UserSpaceStreamNotifierDtor(notifier *streams.PhpStreamNotifier) {
	if notifier != nil && notifier.GetPtr().GetType() != zend.IS_UNDEF {
		zend.ZvalPtrDtor(notifier.GetPtr())
		zend.ZVAL_UNDEF(notifier.GetPtr())
	}
}
func ParseContextOptions(context *core.PhpStreamContext, options *zend.Zval) int {
	var wval *zend.Zval
	var oval *zend.Zval
	var wkey *zend.ZendString
	var okey *zend.ZendString
	var ret int = zend.SUCCESS
	var __ht *zend.HashTable = options.GetArr()
	for _, _p := range __ht.foreachData() {
		var _z *zend.Zval = _p.GetVal()

		wkey = _p.GetKey()
		wval = _z
		zend.ZVAL_DEREF(wval)
		if wkey != nil && wval.IsType(zend.IS_ARRAY) {
			var __ht *zend.HashTable = wval.GetArr()
			for _, _p := range __ht.foreachData() {
				var _z *zend.Zval = _p.GetVal()

				okey = _p.GetKey()
				oval = _z
				if okey != nil {
					streams.PhpStreamContextSetOption(context, wkey.GetVal(), okey.GetVal(), oval)
				}
			}
		} else {
			core.PhpErrorDocref(nil, zend.E_WARNING, "options should have the form [\"wrappername\"][\"optionname\"] = $value")
		}
	}
	return ret
}
func ParseContextParams(context *core.PhpStreamContext, params *zend.Zval) int {
	var ret int = zend.SUCCESS
	var tmp *zend.Zval
	if nil != b.Assign(&tmp, params.GetArr().KeyFind(b.CastStr("notification", b.SizeOf("\"notification\"")-1))) {
		if context.GetNotifier() != nil {
			streams.PhpStreamNotificationFree(context.GetNotifier())
			context.SetNotifier(nil)
		}
		context.SetNotifier(streams.PhpStreamNotificationAlloc())
		context.GetNotifier().SetFunc(UserSpaceStreamNotifier)
		zend.ZVAL_COPY(context.GetNotifier().GetPtr(), tmp)
		context.GetNotifier().SetDtor(UserSpaceStreamNotifierDtor)
	}
	if nil != b.Assign(&tmp, params.GetArr().KeyFind(b.CastStr("options", b.SizeOf("\"options\"")-1))) {
		if tmp.IsType(zend.IS_ARRAY) {
			ParseContextOptions(context, tmp)
		} else {
			core.PhpErrorDocref(nil, zend.E_WARNING, "Invalid stream/context parameter")
		}
	}
	return ret
}
func DecodeContextParam(contextresource *zend.Zval) *core.PhpStreamContext {
	var context *core.PhpStreamContext = nil
	context = zend.ZendFetchResourceEx(contextresource, nil, PhpLeStreamContext())
	if context == nil {
		var stream *core.PhpStream
		stream = zend.ZendFetchResource2Ex(contextresource, nil, streams.PhpFileLeStream(), streams.PhpFileLePstream())
		if stream != nil {
			context = core.PHP_STREAM_CONTEXT(stream)
			if context == nil {

				/* Only way this happens is if file is opened with NO_DEFAULT_CONTEXT
				   param, but then something is called which requires a context.
				   Don't give them the default one though since they already said they
				    didn't want it. */

				context = streams.PhpStreamContextAlloc()
				stream.SetCtx(context.GetRes())
			}
		}
	}
	return context
}
func ZifStreamContextGetOptions(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var zcontext *zend.Zval
	var context *core.PhpStreamContext
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
			if zend.ZendParseArgResource(_arg, &zcontext, 0) == 0 {
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
	context = DecodeContextParam(zcontext)
	if context == nil {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Invalid stream/context parameter")
		zend.RETVAL_FALSE
		return
	}
	zend.ZVAL_COPY(return_value, context.GetOptions())
}
func ZifStreamContextSetOption(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var zcontext *zend.Zval = nil
	var context *core.PhpStreamContext
	if zend.ZEND_NUM_ARGS() == 2 {
		var options *zend.Zval
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
				if zend.ZendParseArgResource(_arg, &zcontext, 0) == 0 {
					_expected_type = zend.Z_EXPECTED_RESOURCE
					_error_code = zend.ZPP_ERROR_WRONG_ARG
					break
				}
				zend.Z_PARAM_PROLOGUE(0, 0)
				if zend.ZendParseArgArray(_arg, &options, 0, 0) == 0 {
					_expected_type = zend.Z_EXPECTED_ARRAY
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

		/* figure out where the context is coming from exactly */

		if !(b.Assign(&context, DecodeContextParam(zcontext))) {
			core.PhpErrorDocref(nil, zend.E_WARNING, "Invalid stream/context parameter")
			zend.RETVAL_FALSE
			return
		}
		zend.RETVAL_BOOL(ParseContextOptions(context, options) == zend.SUCCESS)
		return
	} else {
		var zvalue *zend.Zval
		var wrappername *byte
		var optionname *byte
		var wrapperlen int
		var optionlen int
		for {
			var _flags int = 0
			var _min_num_args int = 4
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
				if zend.ZendParseArgResource(_arg, &zcontext, 0) == 0 {
					_expected_type = zend.Z_EXPECTED_RESOURCE
					_error_code = zend.ZPP_ERROR_WRONG_ARG
					break
				}
				zend.Z_PARAM_PROLOGUE(0, 0)
				if zend.ZendParseArgString(_arg, &wrappername, &wrapperlen, 0) == 0 {
					_expected_type = zend.Z_EXPECTED_STRING
					_error_code = zend.ZPP_ERROR_WRONG_ARG
					break
				}
				zend.Z_PARAM_PROLOGUE(0, 0)
				if zend.ZendParseArgString(_arg, &optionname, &optionlen, 0) == 0 {
					_expected_type = zend.Z_EXPECTED_STRING
					_error_code = zend.ZPP_ERROR_WRONG_ARG
					break
				}
				zend.Z_PARAM_PROLOGUE(0, 0)
				zend.ZendParseArgZvalDeref(_arg, &zvalue, 0)
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

		/* figure out where the context is coming from exactly */

		if !(b.Assign(&context, DecodeContextParam(zcontext))) {
			core.PhpErrorDocref(nil, zend.E_WARNING, "Invalid stream/context parameter")
			zend.RETVAL_FALSE
			return
		}
		zend.RETVAL_BOOL(streams.PhpStreamContextSetOption(context, wrappername, optionname, zvalue) == zend.SUCCESS)
		return
	}
}
func ZifStreamContextSetParams(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var params *zend.Zval
	var zcontext *zend.Zval
	var context *core.PhpStreamContext
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
			if zend.ZendParseArgResource(_arg, &zcontext, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_RESOURCE
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgArray(_arg, &params, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_ARRAY
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
	context = DecodeContextParam(zcontext)
	if context == nil {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Invalid stream/context parameter")
		zend.RETVAL_FALSE
		return
	}
	zend.RETVAL_BOOL(ParseContextParams(context, params) == zend.SUCCESS)
}
func ZifStreamContextGetParams(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var zcontext *zend.Zval
	var context *core.PhpStreamContext
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
			if zend.ZendParseArgResource(_arg, &zcontext, 0) == 0 {
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
	context = DecodeContextParam(zcontext)
	if context == nil {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Invalid stream/context parameter")
		zend.RETVAL_FALSE
		return
	}
	zend.ArrayInit(return_value)
	if context.GetNotifier() != nil && context.GetNotifier().GetPtr().GetType() != zend.IS_UNDEF && context.GetNotifier().GetFunc() == UserSpaceStreamNotifier {
		zend.Z_TRY_ADDREF(context.GetNotifier().GetPtr())
		zend.AddAssocZvalEx(return_value, "notification", b.SizeOf("\"notification\"")-1, context.GetNotifier().GetPtr())
	}
	zend.Z_TRY_ADDREF(context.GetOptions())
	zend.AddAssocZvalEx(return_value, "options", b.SizeOf("\"options\"")-1, context.GetOptions())
}
func ZifStreamContextGetDefault(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var params *zend.Zval = nil
	var context *core.PhpStreamContext
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
			if zend.ZendParseArgArray(_arg, &params, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_ARRAY
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
	if FG(default_context) == nil {
		FG(default_context) = streams.PhpStreamContextAlloc()
	}
	context = FG(default_context)
	if params != nil {
		ParseContextOptions(context, params)
	}
	streams.PhpStreamContextToZval(context, return_value)
}
func ZifStreamContextSetDefault(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var options *zend.Zval = nil
	var context *core.PhpStreamContext
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
			if zend.ZendParseArgArray(_arg, &options, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_ARRAY
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
	if FG(default_context) == nil {
		FG(default_context) = streams.PhpStreamContextAlloc()
	}
	context = FG(default_context)
	ParseContextOptions(context, options)
	streams.PhpStreamContextToZval(context, return_value)
}
func ZifStreamContextCreate(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var options *zend.Zval = nil
	var params *zend.Zval = nil
	var context *core.PhpStreamContext
	for {
		var _flags int = 0
		var _min_num_args int = 0
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
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgArray(_arg, &options, 1, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgArray(_arg, &params, 1, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_ARRAY
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
	context = streams.PhpStreamContextAlloc()
	if options != nil {
		ParseContextOptions(context, options)
	}
	if params != nil {
		ParseContextParams(context, params)
	}
	zend.RETVAL_RES(context.GetRes())
	return
}
func ApplyFilterToStream(append int, execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var zstream *zend.Zval
	var stream *core.PhpStream
	var filtername *byte
	var filternamelen int
	var read_write zend.ZendLong = 0
	var filterparams *zend.Zval = nil
	var filter *core.PhpStreamFilter = nil
	var ret int
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
			if zend.ZendParseArgResource(_arg, &zstream, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_RESOURCE
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgString(_arg, &filtername, &filternamelen, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgLong(_arg, &read_write, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			zend.ZendParseArgZvalDeref(_arg, &filterparams, 0)
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
	core.PhpStreamFromZval(stream, zstream)
	if (read_write & streams.PHP_STREAM_FILTER_ALL) == 0 {

		/* Chain not specified.
		 * Examine stream->mode to determine which filters are needed
		 * There's no harm in attaching a filter to an unused chain,
		 * but why waste the memory and clock cycles?
		 */

		if strchr(stream.GetMode(), 'r') || strchr(stream.GetMode(), '+') {
			read_write |= streams.PHP_STREAM_FILTER_READ
		}
		if strchr(stream.GetMode(), 'w') || strchr(stream.GetMode(), '+') || strchr(stream.GetMode(), 'a') {
			read_write |= streams.PHP_STREAM_FILTER_WRITE
		}
	}
	if (read_write & streams.PHP_STREAM_FILTER_READ) != 0 {
		filter = streams.PhpStreamFilterCreate(filtername, filterparams, stream.GetIsPersistent())
		if filter == nil {
			zend.RETVAL_FALSE
			return
		}
		if append != 0 {
			ret = streams.PhpStreamFilterAppendEx(stream.GetReadfilters(), filter)
		} else {
			ret = streams.PhpStreamFilterPrependEx(stream.GetReadfilters(), filter)
		}
		if ret != zend.SUCCESS {
			streams.PhpStreamFilterRemove(filter, 1)
			zend.RETVAL_FALSE
			return
		}
	}
	if (read_write & streams.PHP_STREAM_FILTER_WRITE) != 0 {
		filter = streams.PhpStreamFilterCreate(filtername, filterparams, stream.GetIsPersistent())
		if filter == nil {
			zend.RETVAL_FALSE
			return
		}
		if append != 0 {
			ret = streams.PhpStreamFilterAppendEx(stream.GetWritefilters(), filter)
		} else {
			ret = streams.PhpStreamFilterPrependEx(stream.GetWritefilters(), filter)
		}
		if ret != zend.SUCCESS {
			streams.PhpStreamFilterRemove(filter, 1)
			zend.RETVAL_FALSE
			return
		}
	}
	if filter != nil {
		filter.SetRes(zend.ZendRegisterResource(filter, streams.PhpFileLeStreamFilter()))
		filter.GetRes().AddRefcount()
		zend.RETVAL_RES(filter.GetRes())
		return
	} else {
		zend.RETVAL_FALSE
		return
	}
}
func ZifStreamFilterPrepend(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	ApplyFilterToStream(0, execute_data, return_value)
}
func ZifStreamFilterAppend(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	ApplyFilterToStream(1, execute_data, return_value)
}
func ZifStreamFilterRemove(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var zfilter *zend.Zval
	var filter *core.PhpStreamFilter
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
			if zend.ZendParseArgResource(_arg, &zfilter, 0) == 0 {
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
	filter = zend.ZendFetchResource(zfilter.GetRes(), nil, streams.PhpFileLeStreamFilter())
	if filter == nil {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Invalid resource given, not a stream filter")
		zend.RETVAL_FALSE
		return
	}
	if streams.PhpStreamFilterFlush(filter, 1) == zend.FAILURE {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Unable to flush filter, not removing")
		zend.RETVAL_FALSE
		return
	}
	if zend.ZendListClose(zfilter.GetRes()) == zend.FAILURE {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Could not invalidate filter, not removing")
		zend.RETVAL_FALSE
		return
	} else {
		streams.PhpStreamFilterRemove(filter, 1)
		zend.RETVAL_TRUE
		return
	}
}
func ZifStreamGetLine(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var str *byte = nil
	var str_len int = 0
	var max_length zend.ZendLong
	var zstream *zend.Zval
	var buf *zend.ZendString
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
			if zend.ZendParseArgResource(_arg, &zstream, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_RESOURCE
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgLong(_arg, &max_length, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgString(_arg, &str, &str_len, 0) == 0 {
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
	if max_length < 0 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "The maximum allowed length must be greater than or equal to zero")
		zend.RETVAL_FALSE
		return
	}
	if max_length == 0 {
		max_length = core.PHP_SOCK_CHUNK_SIZE
	}
	core.PhpStreamFromZval(stream, zstream)
	if b.Assign(&buf, streams.PhpStreamGetRecord(stream, max_length, str, str_len)) {
		zend.RETVAL_STR(buf)
		return
	} else {
		zend.RETVAL_FALSE
		return
	}
}
func ZifStreamSetBlocking(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var zstream *zend.Zval
	var block zend.ZendBool
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
			if zend.ZendParseArgResource(_arg, &zstream, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_RESOURCE
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgBool(_arg, &block, &_dummy, 0) == 0 {
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
	core.PhpStreamFromZval(stream, zstream)
	if core.PhpStreamSetOption(stream, core.PHP_STREAM_OPTION_BLOCKING, block, nil) == -1 {
		zend.RETVAL_FALSE
		return
	}
	zend.RETVAL_TRUE
	return
}
func ZifStreamSetTimeout(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var socket *zend.Zval
	var seconds zend.ZendLong
	var microseconds zend.ZendLong = 0
	var t __struct__timeval
	var stream *core.PhpStream
	var argc int = zend.ZEND_NUM_ARGS()
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
			if zend.ZendParseArgResource(_arg, &socket, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_RESOURCE
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgLong(_arg, &seconds, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgLong(_arg, &microseconds, &_dummy, 0, 0) == 0 {
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
	core.PhpStreamFromZval(stream, socket)
	t.tv_sec = seconds
	if argc == 3 {
		t.tv_usec = microseconds % 1000000
		t.tv_sec += microseconds / 1000000
	} else {
		t.tv_usec = 0
	}
	if core.PHP_STREAM_OPTION_RETURN_OK == core.PhpStreamSetOption(stream, core.PHP_STREAM_OPTION_READ_TIMEOUT, 0, &t) {
		zend.RETVAL_TRUE
		return
	}
	zend.RETVAL_FALSE
	return
}
func ZifStreamSetWriteBuffer(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var arg1 *zend.Zval
	var ret int
	var arg2 zend.ZendLong
	var buff int
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
			if zend.ZendParseArgResource(_arg, &arg1, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_RESOURCE
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgLong(_arg, &arg2, &_dummy, 0, 0) == 0 {
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
	core.PhpStreamFromZval(stream, arg1)
	buff = arg2

	/* if buff is 0 then set to non-buffered */

	if buff == 0 {
		ret = core.PhpStreamSetOption(stream, core.PHP_STREAM_OPTION_WRITE_BUFFER, core.PHP_STREAM_BUFFER_NONE, nil)
	} else {
		ret = core.PhpStreamSetOption(stream, core.PHP_STREAM_OPTION_WRITE_BUFFER, core.PHP_STREAM_BUFFER_FULL, &buff)
	}
	zend.RETVAL_LONG(b.Cond(ret == 0, 0, r.EOF))
	return
}
func ZifStreamSetChunkSize(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var ret int
	var csize zend.ZendLong
	var zstream *zend.Zval
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
			if zend.ZendParseArgResource(_arg, &zstream, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_RESOURCE
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgLong(_arg, &csize, &_dummy, 0, 0) == 0 {
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
	if csize <= 0 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "The chunk size must be a positive integer, given "+zend.ZEND_LONG_FMT, csize)
		zend.RETVAL_FALSE
		return
	}

	/* stream.chunk_size is actually a size_t, but php_stream_set_option
	 * can only use an int to accept the new value and return the old one.
	 * In any case, values larger than INT_MAX for a chunk size make no sense.
	 */

	if csize > core.INT_MAX {
		core.PhpErrorDocref(nil, zend.E_WARNING, "The chunk size cannot be larger than %d", core.INT_MAX)
		zend.RETVAL_FALSE
		return
	}
	core.PhpStreamFromZval(stream, zstream)
	ret = core.PhpStreamSetOption(stream, core.PHP_STREAM_OPTION_SET_CHUNK_SIZE, int(csize), nil)
	zend.RETVAL_LONG(b.CondF(ret > 0, func() zend.ZendLong { return zend.ZendLong(ret) }, func() zend.ZendLong { return zend.ZendLong(r.EOF) }))
	return
}
func ZifStreamSetReadBuffer(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var arg1 *zend.Zval
	var ret int
	var arg2 zend.ZendLong
	var buff int
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
			if zend.ZendParseArgResource(_arg, &arg1, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_RESOURCE
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgLong(_arg, &arg2, &_dummy, 0, 0) == 0 {
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
	core.PhpStreamFromZval(stream, arg1)
	buff = arg2

	/* if buff is 0 then set to non-buffered */

	if buff == 0 {
		ret = core.PhpStreamSetOption(stream, core.PHP_STREAM_OPTION_READ_BUFFER, core.PHP_STREAM_BUFFER_NONE, nil)
	} else {
		ret = core.PhpStreamSetOption(stream, core.PHP_STREAM_OPTION_READ_BUFFER, core.PHP_STREAM_BUFFER_FULL, &buff)
	}
	zend.RETVAL_LONG(b.Cond(ret == 0, 0, r.EOF))
	return
}
func ZifStreamSocketEnableCrypto(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var cryptokind zend.ZendLong = 0
	var zstream *zend.Zval
	var zsessstream *zend.Zval = nil
	var stream *core.PhpStream
	var sessstream *core.PhpStream = nil
	var enable zend.ZendBool
	var cryptokindnull zend.ZendBool = 1
	var ret int
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
			if zend.ZendParseArgResource(_arg, &zstream, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_RESOURCE
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgBool(_arg, &enable, &_dummy, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_BOOL
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgLong(_arg, &cryptokind, &cryptokindnull, 1, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgResource(_arg, &zsessstream, 0) == 0 {
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
	core.PhpStreamFromZval(stream, zstream)
	if enable != 0 {
		if cryptokindnull != 0 {
			var val *zend.Zval
			if !(GET_CTX_OPT(stream, "ssl", "crypto_method", val)) {
				core.PhpErrorDocref(nil, zend.E_WARNING, "When enabling encryption you must specify the crypto type")
				zend.RETVAL_FALSE
				return
			}
			cryptokind = val.GetLval()
		}
		if zsessstream != nil {
			core.PhpStreamFromZval(sessstream, zsessstream)
		}
		if streams.PhpStreamXportCryptoSetup(stream, cryptokind, sessstream) < 0 {
			zend.RETVAL_FALSE
			return
		}
	}
	ret = streams.PhpStreamXportCryptoEnable(stream, enable)
	switch ret {
	case -1:
		zend.RETVAL_FALSE
		return
	case 0:
		zend.RETVAL_LONG(0)
		return
	default:
		zend.RETVAL_TRUE
		return
	}
}
func ZifStreamResolveIncludePath(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var filename *byte
	var filename_len int
	var resolved_path *zend.ZendString
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
	resolved_path = zend.ZendResolvePath(filename, filename_len)
	if resolved_path != nil {
		zend.RETVAL_STR(resolved_path)
		return
	}
	zend.RETVAL_FALSE
	return
}
func ZifStreamIsLocal(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var zstream *zend.Zval
	var stream *core.PhpStream = nil
	var wrapper *core.PhpStreamWrapper = nil
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
			zend.ZendParseArgZvalDeref(_arg, &zstream, 0)
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
	if zstream.IsType(zend.IS_RESOURCE) {
		core.PhpStreamFromZval(stream, zstream)
		if stream == nil {
			zend.RETVAL_FALSE
			return
		}
		wrapper = stream.GetWrapper()
	} else {
		if zend.TryConvertToString(zstream) == 0 {
			return
		}
		wrapper = streams.PhpStreamLocateUrlWrapper(zend.Z_STRVAL_P(zstream), nil, 0)
	}
	if wrapper == nil {
		zend.RETVAL_FALSE
		return
	}
	zend.RETVAL_BOOL(wrapper.GetIsUrl() == 0)
	return
}
func ZifStreamSupportsLock(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var stream *core.PhpStream
	var zsrc *zend.Zval
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
			if zend.ZendParseArgResource(_arg, &zsrc, 0) == 0 {
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
	core.PhpStreamFromZval(stream, zsrc)
	if core.PhpStreamSupportsLock(stream) == 0 {
		zend.RETVAL_FALSE
		return
	}
	zend.RETVAL_TRUE
	return
}
func ZifStreamIsatty(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var zsrc *zend.Zval
	var stream *core.PhpStream
	var fileno core.PhpSocketT
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
			if zend.ZendParseArgResource(_arg, &zsrc, 0) == 0 {
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
	core.PhpStreamFromZval(stream, zsrc)
	if core.PhpStreamCanCast(stream, core.PHP_STREAM_AS_FD_FOR_SELECT) == zend.SUCCESS {
		core.PhpStreamCast(stream, core.PHP_STREAM_AS_FD_FOR_SELECT, any(&fileno), 0)
	} else if core.PhpStreamCanCast(stream, core.PHP_STREAM_AS_FD) == zend.SUCCESS {
		core.PhpStreamCast(stream, core.PHP_STREAM_AS_FD, any(&fileno), 0)
	} else {
		zend.RETVAL_FALSE
		return
	}

	/* Check if the file descriptor identifier is a terminal */

	zend.RETVAL_BOOL(zend.Isatty(fileno) != 0)

	/* Check if the file descriptor identifier is a terminal */
}
func ZifStreamSocketShutdown(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var how zend.ZendLong
	var zstream *zend.Zval
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
			if zend.ZendParseArgResource(_arg, &zstream, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_RESOURCE
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgLong(_arg, &how, &_dummy, 0, 0) == 0 {
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
	if how != streams.STREAM_SHUT_RD && how != streams.STREAM_SHUT_WR && how != streams.STREAM_SHUT_RDWR {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Second parameter $how needs to be one of STREAM_SHUT_RD, STREAM_SHUT_WR or STREAM_SHUT_RDWR")
		zend.RETVAL_FALSE
		return
	}
	core.PhpStreamFromZval(stream, zstream)
	zend.RETVAL_BOOL(streams.PhpStreamXportShutdown(stream, streams.StreamShutdownT(how)) == 0)
	return
}
