// <<generate>>

package standard

import (
	b "sik/builtin"
	"sik/core"
	"sik/core/streams"
	"sik/zend"
	"sik/zend/types"
)

func PhpFsockopenStream(executeData *zend.ZendExecuteData, return_value *types.Zval, persistent int) {
	var host *byte
	var host_len int
	var port zend.ZendLong = -1
	var zerrno *types.Zval = nil
	var zerrstr *types.Zval = nil
	var timeout float64 = float64(FG(default_socket_timeout))
	var conv int64
	var tv __struct__timeval
	var hashkey *byte = nil
	var stream *core.PhpStream = nil
	var err int
	var hostname *byte = nil
	var hostname_len int
	var errstr *types.ZendString = nil
	return_value.SetFalse()
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 5
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
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
			_real_arg = executeData.Arg(0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgString(_arg, &host, &host_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if !zend.ZendParseArgLong00(_arg, &port) {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
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
	if persistent != 0 {
		core.Spprintf(&hashkey, 0, "pfsockopen__%s:"+zend.ZEND_LONG_FMT, host, port)
	}
	if port > 0 {
		hostname_len = core.Spprintf(&hostname, 0, "%s:"+zend.ZEND_LONG_FMT, host, port)
	} else {
		hostname_len = host_len
		hostname = host
	}

	/* prepare the timeout value for use */

	conv = time_t(timeout * 1000000.0)
	tv.tv_sec = conv / 1000000
	tv.tv_usec = conv % 1000000
	stream = streams.PhpStreamXportCreate(hostname, hostname_len, core.REPORT_ERRORS, streams.STREAM_XPORT_CLIENT|streams.STREAM_XPORT_CONNECT, hashkey, &tv, nil, &errstr, &err)
	if port > 0 {
		zend.Efree(hostname)
	}
	if stream == nil {
		core.PhpErrorDocref(nil, zend.E_WARNING, "unable to connect to %s:"+zend.ZEND_LONG_FMT+" (%s)", host, port, b.CondF2(errstr == nil, "Unknown error", func() []byte { return errstr.GetVal() }))
	}
	if hashkey != nil {
		zend.Efree(hashkey)
	}
	if stream == nil {
		if zerrno != nil {
			zend.ZEND_TRY_ASSIGN_REF_LONG(zerrno, err)
		}
		if errstr != nil {
			if zerrstr != nil {
				zend.ZEND_TRY_ASSIGN_REF_STR(zerrstr, errstr)
			} else {
				types.ZendStringRelease(errstr)
			}
		}
		return_value.SetFalse()
		return
	}
	if zerrno != nil {
		zend.ZEND_TRY_ASSIGN_REF_LONG(zerrno, 0)
	}
	if zerrstr != nil {
		zend.ZEND_TRY_ASSIGN_REF_EMPTY_STRING(zerrstr)
	}
	if errstr != nil {
		types.ZendStringReleaseEx(errstr, 0)
	}
	core.PhpStreamToZval(stream, return_value)
}
func ZifFsockopen(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	PhpFsockopenStream(executeData, return_value, 0)
}
func ZifPfsockopen(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	PhpFsockopenStream(executeData, return_value, 1)
}
