// <<generate>>

package standard

import (
	"sik/core"
	"sik/core/streams"
	g "sik/runtime/grammar"
	"sik/zend"
)

// Source: <ext/standard/streamsfuncs.h>

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
  | Authors: Wez Furlong <wez@thebrainroom.com>                          |
  +----------------------------------------------------------------------+
*/

// #define PHP_STREAM_CLIENT_PERSISTENT       1

// #define PHP_STREAM_CLIENT_ASYNC_CONNECT       2

// #define PHP_STREAM_CLIENT_CONNECT       4

var ZifStreamWrapperRegister func(execute_data *zend.ZendExecuteData, return_value *zend.Zval)
var ZifStreamWrapperUnregister func(execute_data *zend.ZendExecuteData, return_value *zend.Zval)
var ZifStreamWrapperRestore func(execute_data *zend.ZendExecuteData, return_value *zend.Zval)

// Source: <ext/standard/streamsfuncs.c>

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
  | Authors: Wez Furlong <wez@thebrainroom.com>                          |
  |          Sara Golemon <pollita@php.net>                              |
  +----------------------------------------------------------------------+
*/

// # include "php.h"

// # include "php_globals.h"

// # include "ext/standard/flock_compat.h"

// # include "ext/standard/file.h"

// # include "ext/standard/php_filestat.h"

// # include "php_open_temporary_file.h"

// # include "ext/standard/basic_functions.h"

// # include "php_ini.h"

// # include "streamsfuncs.h"

// # include "php_network.h"

// # include "php_string.h"

// # include < unistd . h >

// #define php_select(m,r,w,e,t) select ( m , r , w , e , t )

type PhpTimeoutUll = unsigned__long__long

// #define GET_CTX_OPT(stream,wrapper,name,val) ( PHP_STREAM_CONTEXT ( stream ) && NULL != ( val = php_stream_context_get_option ( PHP_STREAM_CONTEXT ( stream ) , wrapper , name ) ) )

/* Streams based network functions */

/* {{{ proto array stream_socket_pair(int domain, int type, int protocol)
   Creates a pair of connected, indistinguishable socket streams */

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

			if zend.ZendParseArgLong(_arg, &domain, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
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

			if zend.ZendParseArgLong(_arg, &type_, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
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

			if zend.ZendParseArgLong(_arg, &protocol, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
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
	if 0 != socketpair(int(domain), int(type_), int(protocol), pair) {
		var errbuf []byte
		core.PhpErrorDocref(nil, 1<<1, "failed to create sockets: [%d]: %s", errno, core.PhpSocketStrerror(errno, errbuf, g.SizeOf("errbuf")))
		return_value.u1.type_info = 2
		return
	}
	var __arr *zend.ZendArray = zend._zendNewArray(0)
	var __z *zend.Zval = return_value
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	s1 = core._phpStreamSockOpenFromSocket(pair[0], 0)
	s2 = core._phpStreamSockOpenFromSocket(pair[1], 0)

	/* set the __exposed flag.
	 * php_stream_to_zval() does, add_next_index_resource() does not */

	s1.__exposed = 1
	s2.__exposed = 1
	zend.AddNextIndexResource(return_value, s1.res)
	zend.AddNextIndexResource(return_value, s2.res)
}

/* }}} */

/* {{{ proto resource stream_socket_client(string remoteaddress [, int &errcode [, string &errstring [, double timeout [, int flags [, resource context]]]]])
   Open a client connection to a remote address */

func ZifStreamSocketClient(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var host *zend.ZendString
	var zerrno *zend.Zval = nil
	var zerrstr *zend.Zval = nil
	var zcontext *zend.Zval = nil
	var timeout float64 = float64(FileGlobals.GetDefaultSocketTimeout())
	var conv PhpTimeoutUll
	var tv __struct__timeval
	var hashkey *byte = nil
	var stream *core.PhpStream = nil
	var err int
	var flags zend.ZendLong = 4
	var errstr *zend.ZendString = nil
	var context *core.PhpStreamContext = nil
	return_value.u1.type_info = 2
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 6
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

			if zend.ZendParseArgStr(_arg, &host, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = 4
				break
			}
			_optional = 1
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

			zend.ZendParseArgZvalDeref(_arg, &zerrno, 0)
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

			zend.ZendParseArgZvalDeref(_arg, &zerrstr, 0)
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

			if zend.ZendParseArgDouble(_arg, &timeout, &_dummy, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_DOUBLE
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

			if zend.ZendParseArgLong(_arg, &flags, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
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

			if zend.ZendParseArgResource(_arg, &zcontext, 0) == 0 {
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
	if g.CondF2(g.CondF1(zcontext != nil, func() any { return zend.ZendFetchResourceEx(zcontext, "Stream-Context", PhpLeStreamContext()) }, flags&16), nil, func() *core.PhpStreamContext { return FileGlobals.GetDefaultContext() }) {
		context = FileGlobals.GetDefaultContext()
	} else {
		FileGlobals.SetDefaultContext(streams.PhpStreamContextAlloc())
		context = FileGlobals.GetDefaultContext()
	}
	if (flags & 1) != 0 {
		zend.ZendSpprintf(&hashkey, 0, "stream_socket_client__%s", host.val)
	}

	/* prepare the timeout value for use */

	conv = php_timeout_ull(timeout * 1000000.0)
	tv.tv_sec = conv / 1000000
	tv.tv_usec = conv % 1000000
	if zerrno != nil {
		for {
			assert(zerrno.u1.v.type_ == 10)
			for {
				var _zv *zend.Zval = zerrno
				var ref *zend.ZendReference = _zv.value.ref
				if ref.sources.ptr != nil {
					zend.ZendTryAssignTypedRefLong(ref, 0)
					break
				}
				_zv = &ref.val
				zend.ZvalPtrDtor(_zv)
				var __z *zend.Zval = _zv
				__z.value.lval = 0
				__z.u1.type_info = 4
				break
			}
			break
		}
	}
	if zerrstr != nil {
		for {
			assert(zerrstr.u1.v.type_ == 10)
			for {
				var _zv *zend.Zval = zerrstr
				var ref *zend.ZendReference = _zv.value.ref
				if ref.sources.ptr != nil {
					zend.ZendTryAssignTypedRefEmptyString(ref)
					break
				}
				_zv = &ref.val
				zend.ZvalPtrDtor(_zv)
				var __z *zend.Zval = _zv
				var __s *zend.ZendString = zend.ZendEmptyString
				__z.value.str = __s
				__z.u1.type_info = 6
				break
			}
			break
		}
	}
	stream = streams._phpStreamXportCreate(host.val, host.len_, 0x8, 0|g.Cond((flags&4) != 0, 2, 0)|g.Cond((flags&2) != 0, 16, 0), hashkey, &tv, context, &errstr, &err)
	if stream == nil {

		/* host might contain binary characters */

		var quoted_host *zend.ZendString = PhpAddslashes(host)
		core.PhpErrorDocref(nil, 1<<1, "unable to connect to %s (%s)", quoted_host.val, g.CondF2(errstr == nil, "Unknown error", func() []byte { return errstr.val }))
		zend.ZendStringReleaseEx(quoted_host, 0)
	}
	if hashkey != nil {
		zend._efree(hashkey)
	}
	if stream == nil {
		if zerrno != nil {
			for {
				assert(zerrno.u1.v.type_ == 10)
				for {
					var _zv *zend.Zval = zerrno
					var ref *zend.ZendReference = _zv.value.ref
					if ref.sources.ptr != nil {
						zend.ZendTryAssignTypedRefLong(ref, err)
						break
					}
					_zv = &ref.val
					zend.ZvalPtrDtor(_zv)
					var __z *zend.Zval = _zv
					__z.value.lval = err
					__z.u1.type_info = 4
					break
				}
				break
			}
		}
		if zerrstr != nil && errstr != nil {
			for {
				assert(zerrstr.u1.v.type_ == 10)
				for {
					var _zv *zend.Zval = zerrstr
					var ref *zend.ZendReference = _zv.value.ref
					if ref.sources.ptr != nil {
						zend.ZendTryAssignTypedRefStr(ref, errstr)
						break
					}
					_zv = &ref.val
					zend.ZvalPtrDtor(_zv)
					var __z *zend.Zval = _zv
					var __s *zend.ZendString = errstr
					__z.value.str = __s
					if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
						__z.u1.type_info = 6
					} else {
						__z.u1.type_info = 6 | 1<<0<<8
					}
					break
				}
				break
			}
		} else if errstr != nil {
			zend.ZendStringReleaseEx(errstr, 0)
		}
		return_value.u1.type_info = 2
		return
	}
	if errstr != nil {
		zend.ZendStringReleaseEx(errstr, 0)
	}
	var __z *zend.Zval = return_value
	__z.value.res = stream.res
	__z.u1.type_info = 9 | 1<<0<<8
	stream.__exposed = 1
}

/* }}} */

func ZifStreamSocketServer(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var host *byte
	var host_len int
	var zerrno *zend.Zval = nil
	var zerrstr *zend.Zval = nil
	var zcontext *zend.Zval = nil
	var stream *core.PhpStream = nil
	var err int = 0
	var flags zend.ZendLong = 4 | 8
	var errstr *zend.ZendString = nil
	var context *core.PhpStreamContext = nil
	return_value.u1.type_info = 2
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 5
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

			if zend.ZendParseArgString(_arg, &host, &host_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = 4
				break
			}
			_optional = 1
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

			zend.ZendParseArgZvalDeref(_arg, &zerrno, 0)
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

			zend.ZendParseArgZvalDeref(_arg, &zerrstr, 0)
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

			if zend.ZendParseArgLong(_arg, &flags, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
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

			if zend.ZendParseArgResource(_arg, &zcontext, 0) == 0 {
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
	if g.CondF2(g.CondF1(zcontext != nil, func() any { return zend.ZendFetchResourceEx(zcontext, "Stream-Context", PhpLeStreamContext()) }, flags&16), nil, func() *core.PhpStreamContext { return FileGlobals.GetDefaultContext() }) {
		context = FileGlobals.GetDefaultContext()
	} else {
		FileGlobals.SetDefaultContext(streams.PhpStreamContextAlloc())
		context = FileGlobals.GetDefaultContext()
	}
	if context != nil {
		zend.ZendGcAddref(&(context.res).gc)
	}
	if zerrno != nil {
		for {
			assert(zerrno.u1.v.type_ == 10)
			for {
				var _zv *zend.Zval = zerrno
				var ref *zend.ZendReference = _zv.value.ref
				if ref.sources.ptr != nil {
					zend.ZendTryAssignTypedRefLong(ref, 0)
					break
				}
				_zv = &ref.val
				zend.ZvalPtrDtor(_zv)
				var __z *zend.Zval = _zv
				__z.value.lval = 0
				__z.u1.type_info = 4
				break
			}
			break
		}
	}
	if zerrstr != nil {
		for {
			assert(zerrstr.u1.v.type_ == 10)
			for {
				var _zv *zend.Zval = zerrstr
				var ref *zend.ZendReference = _zv.value.ref
				if ref.sources.ptr != nil {
					zend.ZendTryAssignTypedRefEmptyString(ref)
					break
				}
				_zv = &ref.val
				zend.ZvalPtrDtor(_zv)
				var __z *zend.Zval = _zv
				var __s *zend.ZendString = zend.ZendEmptyString
				__z.value.str = __s
				__z.u1.type_info = 6
				break
			}
			break
		}
	}
	stream = streams._phpStreamXportCreate(host, host_len, 0x8, 1|int(flags), nil, nil, context, &errstr, &err)
	if stream == nil {
		core.PhpErrorDocref(nil, 1<<1, "unable to connect to %s (%s)", host, g.CondF2(errstr == nil, "Unknown error", func() []byte { return errstr.val }))
	}
	if stream == nil {
		if zerrno != nil {
			for {
				assert(zerrno.u1.v.type_ == 10)
				for {
					var _zv *zend.Zval = zerrno
					var ref *zend.ZendReference = _zv.value.ref
					if ref.sources.ptr != nil {
						zend.ZendTryAssignTypedRefLong(ref, err)
						break
					}
					_zv = &ref.val
					zend.ZvalPtrDtor(_zv)
					var __z *zend.Zval = _zv
					__z.value.lval = err
					__z.u1.type_info = 4
					break
				}
				break
			}
		}
		if zerrstr != nil && errstr != nil {
			for {
				assert(zerrstr.u1.v.type_ == 10)
				for {
					var _zv *zend.Zval = zerrstr
					var ref *zend.ZendReference = _zv.value.ref
					if ref.sources.ptr != nil {
						zend.ZendTryAssignTypedRefStr(ref, errstr)
						break
					}
					_zv = &ref.val
					zend.ZvalPtrDtor(_zv)
					var __z *zend.Zval = _zv
					var __s *zend.ZendString = errstr
					__z.value.str = __s
					if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
						__z.u1.type_info = 6
					} else {
						__z.u1.type_info = 6 | 1<<0<<8
					}
					break
				}
				break
			}
		} else if errstr != nil {
			zend.ZendStringReleaseEx(errstr, 0)
		}
		return_value.u1.type_info = 2
		return
	}
	if errstr != nil {
		zend.ZendStringReleaseEx(errstr, 0)
	}
	var __z *zend.Zval = return_value
	__z.value.res = stream.res
	__z.u1.type_info = 9 | 1<<0<<8
	stream.__exposed = 1
}

/* }}} */

func ZifStreamSocketAccept(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var timeout float64 = float64(FileGlobals.GetDefaultSocketTimeout())
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

			if zend.ZendParseArgResource(_arg, &zstream, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_RESOURCE
				_error_code = 4
				break
			}
			_optional = 1
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

			if zend.ZendParseArgDouble(_arg, &timeout, &_dummy, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_DOUBLE
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

			zend.ZendParseArgZvalDeref(_arg, &zpeername, 0)
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

	/* prepare the timeout value for use */

	conv = php_timeout_ull(timeout * 1000000.0)
	tv.tv_sec = conv / 1000000
	tv.tv_usec = conv % 1000000
	if 0 == streams.PhpStreamXportAccept(stream, &clistream, g.Cond(zpeername != nil, &peername, nil), nil, nil, &tv, &errstr) && clistream != nil {
		if peername != nil {
			for {
				assert(zpeername.u1.v.type_ == 10)
				for {
					var _zv *zend.Zval = zpeername
					var ref *zend.ZendReference = _zv.value.ref
					if ref.sources.ptr != nil {
						zend.ZendTryAssignTypedRefStr(ref, peername)
						break
					}
					_zv = &ref.val
					zend.ZvalPtrDtor(_zv)
					var __z *zend.Zval = _zv
					var __s *zend.ZendString = peername
					__z.value.str = __s
					if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
						__z.u1.type_info = 6
					} else {
						__z.u1.type_info = 6 | 1<<0<<8
					}
					break
				}
				break
			}
		}
		var __z *zend.Zval = return_value
		__z.value.res = clistream.res
		__z.u1.type_info = 9 | 1<<0<<8
		clistream.__exposed = 1
	} else {
		if peername != nil {
			zend.ZendStringRelease(peername)
		}
		core.PhpErrorDocref(nil, 1<<1, "accept failed: %s", g.CondF1(errstr != nil, func() []byte { return errstr.val }, "Unknown error"))
		return_value.u1.type_info = 2
	}
	if errstr != nil {
		zend.ZendStringReleaseEx(errstr, 0)
	}
}

/* }}} */

func ZifStreamSocketGetName(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var stream *core.PhpStream
	var zstream *zend.Zval
	var want_peer zend.ZendBool
	var name *zend.ZendString = nil
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

			if zend.ZendParseArgResource(_arg, &zstream, 0) == 0 {
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

			if zend.ZendParseArgBool(_arg, &want_peer, &_dummy, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_BOOL
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
	if 0 != streams.PhpStreamXportGetName(stream, want_peer, &name, nil, nil) || name == nil {
		return_value.u1.type_info = 2
		return
	}
	if name.len_ == 0 || name.val[0] == 0 {
		zend.ZendStringReleaseEx(name, 0)
		return_value.u1.type_info = 2
		return
	}
	var __z *zend.Zval = return_value
	var __s *zend.ZendString = name
	__z.value.str = __s
	if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
		__z.u1.type_info = 6
	} else {
		__z.u1.type_info = 6 | 1<<0<<8
	}
}

/* }}} */

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

			if zend.ZendParseArgResource(_arg, &zstream, 0) == 0 {
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

			if zend.ZendParseArgString(_arg, &data, &datalen, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = 4
				break
			}
			_optional = 1
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

			if zend.ZendParseArgLong(_arg, &flags, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
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

			if zend.ZendParseArgString(_arg, &target_addr, &target_addr_len, 0) == 0 {
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
	if target_addr_len != 0 {

		/* parse the address */

		if zend.FAILURE == core.PhpNetworkParseNetworkAddressWithPort(target_addr, target_addr_len, (*__struct__sockaddr)(&sa), &sl) {
			core.PhpErrorDocref(nil, 1<<1, "Failed to parse `%s' into a valid network address", target_addr)
			return_value.u1.type_info = 2
			return
		}

		/* parse the address */

	}
	var __z *zend.Zval = return_value
	__z.value.lval = streams.PhpStreamXportSendto(stream, data, datalen, int(flags), g.Cond(target_addr_len != 0, &sa, nil), sl)
	__z.u1.type_info = 4
	return
}

/* }}} */

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

			if zend.ZendParseArgResource(_arg, &zstream, 0) == 0 {
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

			if zend.ZendParseArgLong(_arg, &to_read, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = 4
				break
			}
			_optional = 1
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

			if zend.ZendParseArgLong(_arg, &flags, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
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

			zend.ZendParseArgZvalDeref(_arg, &zremote, 0)
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
	if zremote != nil {
		for {
			assert(zremote.u1.v.type_ == 10)
			for {
				var _zv *zend.Zval = zremote
				var ref *zend.ZendReference = _zv.value.ref
				if ref.sources.ptr != nil {
					zend.ZendTryAssignTypedRefNull(ref)
					break
				}
				_zv = &ref.val
				zend.ZvalPtrDtor(_zv)
				_zv.u1.type_info = 1
				break
			}
			break
		}
	}
	if to_read <= 0 {
		core.PhpErrorDocref(nil, 1<<1, "Length parameter must be greater than 0")
		return_value.u1.type_info = 2
		return
	}
	read_buf = zend.ZendStringAlloc(to_read, 0)
	recvd = streams.PhpStreamXportRecvfrom(stream, read_buf.val, to_read, int(flags), nil, nil, g.Cond(zremote != nil, &remote_addr, nil))
	if recvd >= 0 {
		if zremote != nil && remote_addr != nil {
			for {
				assert(zremote.u1.v.type_ == 10)
				for {
					var _zv *zend.Zval = zremote
					var ref *zend.ZendReference = _zv.value.ref
					if ref.sources.ptr != nil {
						zend.ZendTryAssignTypedRefStr(ref, remote_addr)
						break
					}
					_zv = &ref.val
					zend.ZvalPtrDtor(_zv)
					var __z *zend.Zval = _zv
					var __s *zend.ZendString = remote_addr
					__z.value.str = __s
					if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
						__z.u1.type_info = 6
					} else {
						__z.u1.type_info = 6 | 1<<0<<8
					}
					break
				}
				break
			}
		}
		read_buf.val[recvd] = '0'
		read_buf.len_ = recvd
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = read_buf
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
		return
	}
	zend.ZendStringEfree(read_buf)
	return_value.u1.type_info = 2
	return
}

/* }}} */

func ZifStreamGetContents(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var stream *core.PhpStream
	var zsrc *zend.Zval
	var maxlen zend.ZendLong = ssize_t(size_t - 1)
	var desiredpos zend.ZendLong = -1
	var contents *zend.ZendString
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 3
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

			if zend.ZendParseArgResource(_arg, &zsrc, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_RESOURCE
				_error_code = 4
				break
			}
			_optional = 1
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

			if zend.ZendParseArgLong(_arg, &maxlen, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
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

			if zend.ZendParseArgLong(_arg, &desiredpos, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
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
	if maxlen < 0 && maxlen != ssize_t(size_t-1) {
		core.PhpErrorDocref(nil, 1<<1, "Length must be greater than or equal to zero, or -1")
		return_value.u1.type_info = 2
		return
	}
	if g.Assign(&stream, (*core.PhpStream)(zend.ZendFetchResource2Ex(zsrc, "stream", streams.PhpFileLeStream(), streams.PhpFileLePstream()))) == nil {
		return_value.u1.type_info = 2
		return
	}
	if desiredpos >= 0 {
		var seek_res int = 0
		var position zend.ZendOffT
		position = streams._phpStreamTell(stream)
		if position >= 0 && desiredpos > position {

			/* use SEEK_CUR to allow emulation in streams that don't support seeking */

			seek_res = streams._phpStreamSeek(stream, desiredpos-position, SEEK_CUR)

			/* use SEEK_CUR to allow emulation in streams that don't support seeking */

		} else if desiredpos < position {

			/* desired position before position or error on tell */

			seek_res = streams._phpStreamSeek(stream, desiredpos, SEEK_SET)

			/* desired position before position or error on tell */

		}
		if seek_res != 0 {
			core.PhpErrorDocref(nil, 1<<1, "Failed to seek to position "+"%"+"lld"+" in the stream", desiredpos)
			return_value.u1.type_info = 2
			return
		}
	}
	if g.Assign(&contents, streams._phpStreamCopyToMem(stream, maxlen, 0)) {
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = contents
		__z.value.str = __s
		if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
			__z.u1.type_info = 6
		} else {
			__z.u1.type_info = 6 | 1<<0<<8
		}
		return
	} else {
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = zend.ZendEmptyString
		__z.value.str = __s
		__z.u1.type_info = 6
		return
	}
}

/* }}} */

func ZifStreamCopyToStream(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var src *core.PhpStream
	var dest *core.PhpStream
	var zsrc *zend.Zval
	var zdest *zend.Zval
	var maxlen zend.ZendLong = size_t - 1
	var pos zend.ZendLong = 0
	var len_ int
	var ret int
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 4
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

			if zend.ZendParseArgResource(_arg, &zsrc, 0) == 0 {
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

			if zend.ZendParseArgResource(_arg, &zdest, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_RESOURCE
				_error_code = 4
				break
			}
			_optional = 1
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

			if zend.ZendParseArgLong(_arg, &maxlen, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
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

			if zend.ZendParseArgLong(_arg, &pos, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
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
	if g.Assign(&src, (*core.PhpStream)(zend.ZendFetchResource2Ex(zsrc, "stream", streams.PhpFileLeStream(), streams.PhpFileLePstream()))) == nil {
		return_value.u1.type_info = 2
		return
	}
	if g.Assign(&dest, (*core.PhpStream)(zend.ZendFetchResource2Ex(zdest, "stream", streams.PhpFileLeStream(), streams.PhpFileLePstream()))) == nil {
		return_value.u1.type_info = 2
		return
	}
	if pos > 0 && streams._phpStreamSeek(src, pos, SEEK_SET) < 0 {
		core.PhpErrorDocref(nil, 1<<1, "Failed to seek to position "+"%"+"lld"+" in the stream", pos)
		return_value.u1.type_info = 2
		return
	}
	ret = streams._phpStreamCopyToStreamEx(src, dest, maxlen, &len_)
	if ret != zend.SUCCESS {
		return_value.u1.type_info = 2
		return
	}
	var __z *zend.Zval = return_value
	__z.value.lval = len_
	__z.u1.type_info = 4
	return
}

/* }}} */

func ZifStreamGetMetaData(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var zstream *zend.Zval
	var stream *core.PhpStream
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

			if zend.ZendParseArgResource(_arg, &zstream, 0) == 0 {
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
			return
		}
		break
	}
	if g.Assign(&stream, (*core.PhpStream)(zend.ZendFetchResource2Ex(zstream, "stream", streams.PhpFileLeStream(), streams.PhpFileLePstream()))) == nil {
		return_value.u1.type_info = 2
		return
	}
	var __arr *zend.ZendArray = zend._zendNewArray(0)
	var __z *zend.Zval = return_value
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	if !(g.Cond(streams._phpStreamSetOption(stream, 11, 0, return_value) == 0, 1, 0)) {
		zend.AddAssocBoolEx(return_value, "timed_out", strlen("timed_out"), 0)
		zend.AddAssocBoolEx(return_value, "blocked", strlen("blocked"), 1)
		zend.AddAssocBoolEx(return_value, "eof", strlen("eof"), streams._phpStreamEof(stream))
	}
	if stream.wrapperdata.u1.v.type_ != 0 {
		zend.ZvalAddrefP(&stream.wrapperdata)
		zend.AddAssocZvalEx(return_value, "wrapper_data", strlen("wrapper_data"), &stream.wrapperdata)
	}
	if stream.wrapper != nil {
		zend.AddAssocStringEx(return_value, "wrapper_type", strlen("wrapper_type"), (*byte)(stream.wrapper.wops.label))
	}
	zend.AddAssocStringEx(return_value, "stream_type", strlen("stream_type"), (*byte)(stream.ops.label))
	zend.AddAssocStringEx(return_value, "mode", strlen("mode"), stream.mode)
	zend.AddAssocLongEx(return_value, "unread_bytes", strlen("unread_bytes"), stream.writepos-stream.readpos)
	zend.AddAssocBoolEx(return_value, "seekable", strlen("seekable"), stream.ops.seek != nil && (stream.flags&0x1) == 0)
	if stream.orig_path != nil {
		zend.AddAssocStringEx(return_value, "uri", strlen("uri"), stream.orig_path)
	}
}

/* }}} */

func ZifStreamGetTransports(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var stream_xport_hash *zend.HashTable
	var stream_xport *zend.ZendString
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	if g.Assign(&stream_xport_hash, streams.PhpStreamXportGetHash()) {
		var __arr *zend.ZendArray = zend._zendNewArray(0)
		var __z *zend.Zval = return_value
		__z.value.arr = __arr
		__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
		for {
			var __ht *zend.HashTable = stream_xport_hash
			var _p *zend.Bucket = __ht.arData
			var _end *zend.Bucket = _p + __ht.nNumUsed
			for ; _p != _end; _p++ {
				var _z *zend.Zval = &_p.val

				if _z.u1.v.type_ == 0 {
					continue
				}
				stream_xport = _p.key
				zend.AddNextIndexStr(return_value, zend.ZendStringCopy(stream_xport))
			}
			break
		}
	} else {
		return_value.u1.type_info = 2
		return
	}
}

/* }}} */

func ZifStreamGetWrappers(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var url_stream_wrappers_hash *zend.HashTable
	var stream_protocol *zend.ZendString
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	if g.Assign(&url_stream_wrappers_hash, streams._phpStreamGetUrlStreamWrappersHash()) {
		var __arr *zend.ZendArray = zend._zendNewArray(0)
		var __z *zend.Zval = return_value
		__z.value.arr = __arr
		__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
		for {
			var __ht *zend.HashTable = url_stream_wrappers_hash
			var _p *zend.Bucket = __ht.arData
			var _end *zend.Bucket = _p + __ht.nNumUsed
			for ; _p != _end; _p++ {
				var _z *zend.Zval = &_p.val

				if _z.u1.v.type_ == 0 {
					continue
				}
				stream_protocol = _p.key
				if stream_protocol != nil {
					zend.AddNextIndexStr(return_value, zend.ZendStringCopy(stream_protocol))
				}
			}
			break
		}
	} else {
		return_value.u1.type_info = 2
		return
	}
}

/* }}} */

func StreamArrayToFdSet(stream_array *zend.Zval, fds *fd_set, max_fd *core.PhpSocketT) int {
	var elem *zend.Zval
	var stream *core.PhpStream
	var cnt int = 0
	if stream_array.u1.v.type_ != 7 {
		return 0
	}
	for {
		var __ht *zend.HashTable = stream_array.value.arr
		var _p *zend.Bucket = __ht.arData
		var _end *zend.Bucket = _p + __ht.nNumUsed
		for ; _p != _end; _p++ {
			var _z *zend.Zval = &_p.val

			if _z.u1.v.type_ == 0 {
				continue
			}
			elem = _z

			/* Temporary int fd is needed for the STREAM data type on windows, passing this_fd directly to php_stream_cast()
			   would eventually bring a wrong result on x64. php_stream_cast() casts to int internally, and this will leave
			   the higher bits of a SOCKET variable uninitialized on systems with little endian. */

			var this_fd core.PhpSocketT
			if elem.u1.v.type_ == 10 {
				elem = &(*elem).value.ref.val
			}
			stream = (*core.PhpStream)(zend.ZendFetchResource2Ex(elem, "stream", streams.PhpFileLeStream(), streams.PhpFileLePstream()))
			if stream == nil {
				continue
			}

			/* get the fd.
			 * NB: Most other code will NOT use the PHP_STREAM_CAST_INTERNAL flag
			 * when casting.  It is only used here so that the buffered data warning
			 * is not displayed.
			 * */

			if zend.SUCCESS == streams._phpStreamCast(stream, 3|0x20000000, any(&this_fd), 1) && this_fd != -1 {
				if this_fd < FD_SETSIZE {
					FD_SET(this_fd, fds)
				}
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
		break
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
	if stream_array.u1.v.type_ != 7 {
		return 0
	}
	ht = zend._zendNewArray(stream_array.value.arr.nNumOfElements)
	for {
		var __ht *zend.HashTable = stream_array.value.arr
		var _p *zend.Bucket = __ht.arData
		var _end *zend.Bucket = _p + __ht.nNumUsed
		for ; _p != _end; _p++ {
			var _z *zend.Zval = &_p.val

			if _z.u1.v.type_ == 0 {
				continue
			}
			num_ind = _p.h
			key = _p.key
			elem = _z
			var this_fd core.PhpSocketT
			if elem.u1.v.type_ == 10 {
				elem = &(*elem).value.ref.val
			}
			stream = (*core.PhpStream)(zend.ZendFetchResource2Ex(elem, "stream", streams.PhpFileLeStream(), streams.PhpFileLePstream()))
			if stream == nil {
				continue
			}

			/* get the fd
			 * NB: Most other code will NOT use the PHP_STREAM_CAST_INTERNAL flag
			 * when casting.  It is only used here so that the buffered data warning
			 * is not displayed.
			 */

			if zend.SUCCESS == streams._phpStreamCast(stream, 3|0x20000000, any(&this_fd), 1) && this_fd != -1 {
				if this_fd < FD_SETSIZE && FD_ISSET(this_fd, fds) {
					if key == nil {
						dest_elem = zend.ZendHashIndexUpdate(ht, num_ind, elem)
					} else {
						dest_elem = zend.ZendHashUpdate(ht, key, elem)
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
		break
	}

	/* destroy old array and add new one */

	zend.ZvalPtrDtor(stream_array)
	var __arr *zend.ZendArray = ht
	var __z *zend.Zval = stream_array
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
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
	if stream_array.u1.v.type_ != 7 {
		return 0
	}
	ht = zend._zendNewArray(stream_array.value.arr.nNumOfElements)
	for {
		var __ht *zend.HashTable = stream_array.value.arr
		var _p *zend.Bucket = __ht.arData
		var _end *zend.Bucket = _p + __ht.nNumUsed
		for ; _p != _end; _p++ {
			var _z *zend.Zval = &_p.val

			if _z.u1.v.type_ == 0 {
				continue
			}
			num_ind = _p.h
			key = _p.key
			elem = _z
			if elem.u1.v.type_ == 10 {
				elem = &(*elem).value.ref.val
			}
			stream = (*core.PhpStream)(zend.ZendFetchResource2Ex(elem, "stream", streams.PhpFileLeStream(), streams.PhpFileLePstream()))
			if stream == nil {
				continue
			}
			if stream.writepos-stream.readpos > 0 {

				/* allow readable non-descriptor based streams to participate in stream_select.
				 * Non-descriptor streams will only "work" if they have previously buffered the
				 * data.  Not ideal, but better than nothing.
				 * This branch of code also allows blocking streams with buffered data to
				 * operate correctly in stream_select.
				 * */

				if key == nil {
					dest_elem = zend.ZendHashIndexUpdate(ht, num_ind, elem)
				} else {
					dest_elem = zend.ZendHashUpdate(ht, key, elem)
				}
				zend.ZvalAddRef(dest_elem)
				ret++
				continue
			}
		}
		break
	}
	if ret > 0 {

		/* destroy old array and add new one */

		zend.ZvalPtrDtor(stream_array)
		var __arr *zend.ZendArray = ht
		var __z *zend.Zval = stream_array
		__z.value.arr = __arr
		__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	} else {
		zend.ZendArrayDestroy(ht)
	}
	return ret
}

/* }}} */

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
			if _arg.u1.v.type_ == 10 {
				_arg = &(*_arg).value.ref.val
			}

			if zend.ZendParseArgArray(_arg, &r_array, 1, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_ARRAY
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
			if _arg.u1.v.type_ == 10 {
				_arg = &(*_arg).value.ref.val
			}

			if zend.ZendParseArgArray(_arg, &w_array, 1, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_ARRAY
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
			if _arg.u1.v.type_ == 10 {
				_arg = &(*_arg).value.ref.val
			}

			if zend.ZendParseArgArray(_arg, &e_array, 1, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_ARRAY
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

			if zend.ZendParseArgLong(_arg, &sec, &secnull, 1, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = 4
				break
			}
			_optional = 1
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

			if zend.ZendParseArgLong(_arg, &usec, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
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
		core.PhpErrorDocref(nil, 1<<1, "No stream arrays were passed")
		return_value.u1.type_info = 2
		return
	}
	if max_fd >= FD_SETSIZE {
		core._phpEmitFdSetsizeWarning(max_fd)
		max_fd = FD_SETSIZE - 1
	}

	/* If seconds is not set to null, build the timeval, else we wait indefinitely */

	if secnull == 0 {
		if sec < 0 {
			core.PhpErrorDocref(nil, 1<<1, "The seconds parameter must be greater than 0")
			return_value.u1.type_info = 2
			return
		} else if usec < 0 {
			core.PhpErrorDocref(nil, 1<<1, "The microseconds parameter must be greater than 0")
			return_value.u1.type_info = 2
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
				var __z *zend.Zval = w_array
				__z.value.arr = (*zend.ZendArray)(&zend.ZendEmptyArray)
				__z.u1.type_info = 7
			}
			if e_array != nil {
				zend.ZvalPtrDtor(e_array)
				var __z *zend.Zval = e_array
				__z.value.arr = (*zend.ZendArray)(&zend.ZendEmptyArray)
				__z.u1.type_info = 7
			}
			var __z *zend.Zval = return_value
			__z.value.lval = retval
			__z.u1.type_info = 4
			return
		}
	}
	retval = select_(max_fd+1, &rfds, &wfds, &efds, tv_p)
	if retval == -1 {
		core.PhpErrorDocref(nil, 1<<1, "unable to select [%d]: %s (max_fd=%d)", errno, strerror(errno), max_fd)
		return_value.u1.type_info = 2
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
	var __z *zend.Zval = return_value
	__z.value.lval = retval
	__z.u1.type_info = 4
	return
}

/* }}} */

func UserSpaceStreamNotifier(context *core.PhpStreamContext, notifycode int, severity int, xmsg *byte, xcode int, bytes_sofar int, bytes_max int, ptr any) {
	var callback *zend.Zval = &context.notifier.ptr
	var retval zend.Zval
	var zvs []zend.Zval
	var i int
	var __z *zval = &zvs[0]
	__z.value.lval = notifycode
	__z.u1.type_info = 4
	var __z *zval = &zvs[1]
	__z.value.lval = severity
	__z.u1.type_info = 4
	if xmsg != nil {
		var _s *byte = xmsg
		var __z *zend.Zval = &zvs[2]
		var __s *zend.ZendString = zend.ZendStringInit(_s, strlen(_s), 0)
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
	} else {
		&zvs[2].u1.type_info = 1
	}
	var __z *zval = &zvs[3]
	__z.value.lval = xcode
	__z.u1.type_info = 4
	var __z *zval = &zvs[4]
	__z.value.lval = bytes_sofar
	__z.u1.type_info = 4
	var __z *zend.Zval = &zvs[5]
	__z.value.lval = bytes_max
	__z.u1.type_info = 4
	if zend.FAILURE == zend._callUserFunctionEx(nil, callback, &retval, 6, zvs, 0) {
		core.PhpErrorDocref(nil, 1<<1, "failed to call user notifier")
	}
	for i = 0; i < 6; i++ {
		zend.ZvalPtrDtor(&zvs[i])
	}
	zend.ZvalPtrDtor(&retval)
}
func UserSpaceStreamNotifierDtor(notifier *streams.PhpStreamNotifier) {
	if notifier != nil && notifier.ptr.u1.v.type_ != 0 {
		zend.ZvalPtrDtor(&notifier.ptr)
		&notifier.ptr.u1.type_info = 0
	}
}
func ParseContextOptions(context *core.PhpStreamContext, options *zend.Zval) int {
	var wval *zend.Zval
	var oval *zend.Zval
	var wkey *zend.ZendString
	var okey *zend.ZendString
	var ret int = zend.SUCCESS
	for {
		var __ht *zend.HashTable = options.value.arr
		var _p *zend.Bucket = __ht.arData
		var _end *zend.Bucket = _p + __ht.nNumUsed
		for ; _p != _end; _p++ {
			var _z *zend.Zval = &_p.val

			if _z.u1.v.type_ == 0 {
				continue
			}
			wkey = _p.key
			wval = _z
			if wval.u1.v.type_ == 10 {
				wval = &(*wval).value.ref.val
			}
			if wkey != nil && wval.u1.v.type_ == 7 {
				for {
					var __ht *zend.HashTable = wval.value.arr
					var _p *zend.Bucket = __ht.arData
					var _end *zend.Bucket = _p + __ht.nNumUsed
					for ; _p != _end; _p++ {
						var _z *zend.Zval = &_p.val

						if _z.u1.v.type_ == 0 {
							continue
						}
						okey = _p.key
						oval = _z
						if okey != nil {
							streams.PhpStreamContextSetOption(context, wkey.val, okey.val, oval)
						}
					}
					break
				}
			} else {
				core.PhpErrorDocref(nil, 1<<1, "options should have the form [\"wrappername\"][\"optionname\"] = $value")
			}
		}
		break
	}
	return ret
}
func ParseContextParams(context *core.PhpStreamContext, params *zend.Zval) int {
	var ret int = zend.SUCCESS
	var tmp *zend.Zval
	if nil != g.Assign(&tmp, zend.ZendHashStrFind(params.value.arr, "notification", g.SizeOf("\"notification\"")-1)) {
		if context.notifier != nil {
			streams.PhpStreamNotificationFree(context.notifier)
			context.notifier = nil
		}
		context.notifier = streams.PhpStreamNotificationAlloc()
		context.notifier.func_ = UserSpaceStreamNotifier
		var _z1 *zend.Zval = &context.notifier.ptr
		var _z2 *zend.Zval = tmp
		var _gc *zend.ZendRefcounted = _z2.value.counted
		var _t uint32 = _z2.u1.type_info
		_z1.value.counted = _gc
		_z1.u1.type_info = _t
		if (_t & 0xff00) != 0 {
			zend.ZendGcAddref(&_gc.gc)
		}
		context.notifier.dtor = UserSpaceStreamNotifierDtor
	}
	if nil != g.Assign(&tmp, zend.ZendHashStrFind(params.value.arr, "options", g.SizeOf("\"options\"")-1)) {
		if tmp.u1.v.type_ == 7 {
			ParseContextOptions(context, tmp)
		} else {
			core.PhpErrorDocref(nil, 1<<1, "Invalid stream/context parameter")
		}
	}
	return ret
}

/* given a zval which is either a stream or a context, return the underlying
 * stream_context.  If it is a stream that does not have a context assigned, it
 * will create and assign a context and return that.  */

func DecodeContextParam(contextresource *zend.Zval) *core.PhpStreamContext {
	var context *core.PhpStreamContext = nil
	context = zend.ZendFetchResourceEx(contextresource, nil, PhpLeStreamContext())
	if context == nil {
		var stream *core.PhpStream
		stream = zend.ZendFetchResource2Ex(contextresource, nil, streams.PhpFileLeStream(), streams.PhpFileLePstream())
		if stream != nil {
			context = (*core.PhpStreamContext)(g.CondF1(stream.ctx != nil, func() any { return stream.ctx.ptr }, nil))
			if context == nil {

				/* Only way this happens is if file is opened with NO_DEFAULT_CONTEXT
				   param, but then something is called which requires a context.
				   Don't give them the default one though since they already said they
				    didn't want it. */

				context = streams.PhpStreamContextAlloc()
				stream.ctx = context.res
			}
		}
	}
	return context
}

/* }}} */

func ZifStreamContextGetOptions(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var zcontext *zend.Zval
	var context *core.PhpStreamContext
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

			if zend.ZendParseArgResource(_arg, &zcontext, 0) == 0 {
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
	context = DecodeContextParam(zcontext)
	if context == nil {
		core.PhpErrorDocref(nil, 1<<1, "Invalid stream/context parameter")
		return_value.u1.type_info = 2
		return
	}
	var _z1 *zend.Zval = return_value
	var _z2 *zend.Zval = &context.options
	var _gc *zend.ZendRefcounted = _z2.value.counted
	var _t uint32 = _z2.u1.type_info
	_z1.value.counted = _gc
	_z1.u1.type_info = _t
	if (_t & 0xff00) != 0 {
		zend.ZendGcAddref(&_gc.gc)
	}
}

/* }}} */

func ZifStreamContextSetOption(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var zcontext *zend.Zval = nil
	var context *core.PhpStreamContext
	if execute_data.This.u2.num_args == 2 {
		var options *zend.Zval
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

				if zend.ZendParseArgResource(_arg, &zcontext, 0) == 0 {
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

				if zend.ZendParseArgArray(_arg, &options, 0, 0) == 0 {
					_expected_type = zend.Z_EXPECTED_ARRAY
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

		/* figure out where the context is coming from exactly */

		if !(g.Assign(&context, DecodeContextParam(zcontext))) {
			core.PhpErrorDocref(nil, 1<<1, "Invalid stream/context parameter")
			return_value.u1.type_info = 2
			return
		}
		if ParseContextOptions(context, options) == zend.SUCCESS {
			return_value.u1.type_info = 3
		} else {
			return_value.u1.type_info = 2
		}
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

				if zend.ZendParseArgResource(_arg, &zcontext, 0) == 0 {
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

				if zend.ZendParseArgString(_arg, &wrappername, &wrapperlen, 0) == 0 {
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

				if zend.ZendParseArgString(_arg, &optionname, &optionlen, 0) == 0 {
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

				zend.ZendParseArgZvalDeref(_arg, &zvalue, 0)
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

		/* figure out where the context is coming from exactly */

		if !(g.Assign(&context, DecodeContextParam(zcontext))) {
			core.PhpErrorDocref(nil, 1<<1, "Invalid stream/context parameter")
			return_value.u1.type_info = 2
			return
		}
		if streams.PhpStreamContextSetOption(context, wrappername, optionname, zvalue) == zend.SUCCESS {
			return_value.u1.type_info = 3
		} else {
			return_value.u1.type_info = 2
		}
		return
	}
}

/* }}} */

func ZifStreamContextSetParams(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var params *zend.Zval
	var zcontext *zend.Zval
	var context *core.PhpStreamContext
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

			if zend.ZendParseArgResource(_arg, &zcontext, 0) == 0 {
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

			if zend.ZendParseArgArray(_arg, &params, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_ARRAY
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
	context = DecodeContextParam(zcontext)
	if context == nil {
		core.PhpErrorDocref(nil, 1<<1, "Invalid stream/context parameter")
		return_value.u1.type_info = 2
		return
	}
	if ParseContextParams(context, params) == zend.SUCCESS {
		return_value.u1.type_info = 3
	} else {
		return_value.u1.type_info = 2
	}
}

/* }}} */

func ZifStreamContextGetParams(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var zcontext *zend.Zval
	var context *core.PhpStreamContext
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

			if zend.ZendParseArgResource(_arg, &zcontext, 0) == 0 {
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
	context = DecodeContextParam(zcontext)
	if context == nil {
		core.PhpErrorDocref(nil, 1<<1, "Invalid stream/context parameter")
		return_value.u1.type_info = 2
		return
	}
	var __arr *zend.ZendArray = zend._zendNewArray(0)
	var __z *zend.Zval = return_value
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	if context.notifier != nil && context.notifier.ptr.u1.v.type_ != 0 && context.notifier.func_ == UserSpaceStreamNotifier {
		if &(context.notifier.ptr).u1.v.type_flags != 0 {
			zend.ZvalAddrefP(&(context.notifier.ptr))
		}
		zend.AddAssocZvalEx(return_value, "notification", g.SizeOf("\"notification\"")-1, &context.notifier.ptr)
	}
	if &(context.options).u1.v.type_flags != 0 {
		zend.ZvalAddrefP(&(context.options))
	}
	zend.AddAssocZvalEx(return_value, "options", g.SizeOf("\"options\"")-1, &context.options)
}

/* }}} */

func ZifStreamContextGetDefault(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var params *zend.Zval = nil
	var context *core.PhpStreamContext
	for {
		var _flags int = 0
		var _min_num_args int = 0
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
			_optional = 1
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

			if zend.ZendParseArgArray(_arg, &params, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_ARRAY
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
	if FileGlobals.GetDefaultContext() == nil {
		FileGlobals.SetDefaultContext(streams.PhpStreamContextAlloc())
	}
	context = FileGlobals.GetDefaultContext()
	if params != nil {
		ParseContextOptions(context, params)
	}
	var __z *zend.Zval = return_value
	__z.value.res = context.res
	__z.u1.type_info = 9 | 1<<0<<8
	zend.ZendGcAddref(&(context.res).gc)
}

/* }}} */

func ZifStreamContextSetDefault(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var options *zend.Zval = nil
	var context *core.PhpStreamContext
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

			if zend.ZendParseArgArray(_arg, &options, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_ARRAY
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
			return
		}
		break
	}
	if FileGlobals.GetDefaultContext() == nil {
		FileGlobals.SetDefaultContext(streams.PhpStreamContextAlloc())
	}
	context = FileGlobals.GetDefaultContext()
	ParseContextOptions(context, options)
	var __z *zend.Zval = return_value
	__z.value.res = context.res
	__z.u1.type_info = 9 | 1<<0<<8
	zend.ZendGcAddref(&(context.res).gc)
}

/* }}} */

func ZifStreamContextCreate(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var options *zend.Zval = nil
	var params *zend.Zval = nil
	var context *core.PhpStreamContext
	for {
		var _flags int = 0
		var _min_num_args int = 0
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
			_optional = 1
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

			if zend.ZendParseArgArray(_arg, &options, 1, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_ARRAY
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

			if zend.ZendParseArgArray(_arg, &params, 1, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_ARRAY
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
	context = streams.PhpStreamContextAlloc()
	if options != nil {
		ParseContextOptions(context, options)
	}
	if params != nil {
		ParseContextParams(context, params)
	}
	var __z *zend.Zval = return_value
	__z.value.res = context.res
	__z.u1.type_info = 9 | 1<<0<<8
	return
}

/* }}} */

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

			if zend.ZendParseArgResource(_arg, &zstream, 0) == 0 {
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

			if zend.ZendParseArgString(_arg, &filtername, &filternamelen, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = 4
				break
			}
			_optional = 1
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

			if zend.ZendParseArgLong(_arg, &read_write, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
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

			zend.ZendParseArgZvalDeref(_arg, &filterparams, 0)
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
	if (read_write & (0x1 | 0x2)) == 0 {

		/* Chain not specified.
		 * Examine stream->mode to determine which filters are needed
		 * There's no harm in attaching a filter to an unused chain,
		 * but why waste the memory and clock cycles?
		 */

		if strchr(stream.mode, 'r') || strchr(stream.mode, '+') {
			read_write |= 0x1
		}
		if strchr(stream.mode, 'w') || strchr(stream.mode, '+') || strchr(stream.mode, 'a') {
			read_write |= 0x2
		}
	}
	if (read_write & 0x1) != 0 {
		filter = streams.PhpStreamFilterCreate(filtername, filterparams, stream.is_persistent)
		if filter == nil {
			return_value.u1.type_info = 2
			return
		}
		if append != 0 {
			ret = streams.PhpStreamFilterAppendEx(&stream.readfilters, filter)
		} else {
			ret = streams.PhpStreamFilterPrependEx(&stream.readfilters, filter)
		}
		if ret != zend.SUCCESS {
			streams.PhpStreamFilterRemove(filter, 1)
			return_value.u1.type_info = 2
			return
		}
	}
	if (read_write & 0x2) != 0 {
		filter = streams.PhpStreamFilterCreate(filtername, filterparams, stream.is_persistent)
		if filter == nil {
			return_value.u1.type_info = 2
			return
		}
		if append != 0 {
			ret = streams.PhpStreamFilterAppendEx(&stream.writefilters, filter)
		} else {
			ret = streams.PhpStreamFilterPrependEx(&stream.writefilters, filter)
		}
		if ret != zend.SUCCESS {
			streams.PhpStreamFilterRemove(filter, 1)
			return_value.u1.type_info = 2
			return
		}
	}
	if filter != nil {
		filter.res = zend.ZendRegisterResource(filter, streams.PhpFileLeStreamFilter())
		zend.ZendGcAddref(&(filter.res).gc)
		var __z *zend.Zval = return_value
		__z.value.res = filter.res
		__z.u1.type_info = 9 | 1<<0<<8
		return
	} else {
		return_value.u1.type_info = 2
		return
	}
}

/* }}} */

func ZifStreamFilterPrepend(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	ApplyFilterToStream(0, execute_data, return_value)
}

/* }}} */

func ZifStreamFilterAppend(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	ApplyFilterToStream(1, execute_data, return_value)
}

/* }}} */

func ZifStreamFilterRemove(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var zfilter *zend.Zval
	var filter *core.PhpStreamFilter
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

			if zend.ZendParseArgResource(_arg, &zfilter, 0) == 0 {
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
	filter = zend.ZendFetchResource(zfilter.value.res, nil, streams.PhpFileLeStreamFilter())
	if filter == nil {
		core.PhpErrorDocref(nil, 1<<1, "Invalid resource given, not a stream filter")
		return_value.u1.type_info = 2
		return
	}
	if streams._phpStreamFilterFlush(filter, 1) == zend.FAILURE {
		core.PhpErrorDocref(nil, 1<<1, "Unable to flush filter, not removing")
		return_value.u1.type_info = 2
		return
	}
	if zend.ZendListClose(zfilter.value.res) == zend.FAILURE {
		core.PhpErrorDocref(nil, 1<<1, "Could not invalidate filter, not removing")
		return_value.u1.type_info = 2
		return
	} else {
		streams.PhpStreamFilterRemove(filter, 1)
		return_value.u1.type_info = 3
		return
	}
}

/* }}} */

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

			if zend.ZendParseArgResource(_arg, &zstream, 0) == 0 {
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

			if zend.ZendParseArgLong(_arg, &max_length, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = 4
				break
			}
			_optional = 1
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

			if zend.ZendParseArgString(_arg, &str, &str_len, 0) == 0 {
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
	if max_length < 0 {
		core.PhpErrorDocref(nil, 1<<1, "The maximum allowed length must be greater than or equal to zero")
		return_value.u1.type_info = 2
		return
	}
	if max_length == 0 {
		max_length = 8192
	}
	if g.Assign(&stream, (*core.PhpStream)(zend.ZendFetchResource2Ex(zstream, "stream", streams.PhpFileLeStream(), streams.PhpFileLePstream()))) == nil {
		return_value.u1.type_info = 2
		return
	}
	if g.Assign(&buf, streams.PhpStreamGetRecord(stream, max_length, str, str_len)) {
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = buf
		__z.value.str = __s
		if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
			__z.u1.type_info = 6
		} else {
			__z.u1.type_info = 6 | 1<<0<<8
		}
		return
	} else {
		return_value.u1.type_info = 2
		return
	}
}

/* }}} */

func ZifStreamSetBlocking(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var zstream *zend.Zval
	var block zend.ZendBool
	var stream *core.PhpStream
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

			if zend.ZendParseArgResource(_arg, &zstream, 0) == 0 {
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

			if zend.ZendParseArgBool(_arg, &block, &_dummy, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_BOOL
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
			return
		}
		break
	}
	if g.Assign(&stream, (*core.PhpStream)(zend.ZendFetchResource2Ex(zstream, "stream", streams.PhpFileLeStream(), streams.PhpFileLePstream()))) == nil {
		return_value.u1.type_info = 2
		return
	}
	if streams._phpStreamSetOption(stream, 1, block, nil) == -1 {
		return_value.u1.type_info = 2
		return
	}
	return_value.u1.type_info = 3
	return
}

/* }}} */

func ZifStreamSetTimeout(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var socket *zend.Zval
	var seconds zend.ZendLong
	var microseconds zend.ZendLong = 0
	var t __struct__timeval
	var stream *core.PhpStream
	var argc int = execute_data.This.u2.num_args
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 3
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

			if zend.ZendParseArgResource(_arg, &socket, 0) == 0 {
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

			if zend.ZendParseArgLong(_arg, &seconds, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = 4
				break
			}
			_optional = 1
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

			if zend.ZendParseArgLong(_arg, &microseconds, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
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
			return
		}
		break
	}
	if g.Assign(&stream, (*core.PhpStream)(zend.ZendFetchResource2Ex(socket, "stream", streams.PhpFileLeStream(), streams.PhpFileLePstream()))) == nil {
		return_value.u1.type_info = 2
		return
	}
	t.tv_sec = seconds
	if argc == 3 {
		t.tv_usec = microseconds % 1000000
		t.tv_sec += microseconds / 1000000
	} else {
		t.tv_usec = 0
	}
	if 0 == streams._phpStreamSetOption(stream, 4, 0, &t) {
		return_value.u1.type_info = 3
		return
	}
	return_value.u1.type_info = 2
	return
}

/* }}} */

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

			if zend.ZendParseArgResource(_arg, &arg1, 0) == 0 {
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

			if zend.ZendParseArgLong(_arg, &arg2, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
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
	if g.Assign(&stream, (*core.PhpStream)(zend.ZendFetchResource2Ex(arg1, "stream", streams.PhpFileLeStream(), streams.PhpFileLePstream()))) == nil {
		return_value.u1.type_info = 2
		return
	}
	buff = arg2

	/* if buff is 0 then set to non-buffered */

	if buff == 0 {
		ret = streams._phpStreamSetOption(stream, 3, 0, nil)
	} else {
		ret = streams._phpStreamSetOption(stream, 3, 2, &buff)
	}
	var __z *zend.Zval = return_value
	if ret == 0 {
		__z.value.lval = 0
	} else {
		__z.value.lval = EOF
	}
	__z.u1.type_info = 4
	return
}

/* }}} */

func ZifStreamSetChunkSize(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var ret int
	var csize zend.ZendLong
	var zstream *zend.Zval
	var stream *core.PhpStream
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

			if zend.ZendParseArgResource(_arg, &zstream, 0) == 0 {
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

			if zend.ZendParseArgLong(_arg, &csize, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
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
	if csize <= 0 {
		core.PhpErrorDocref(nil, 1<<1, "The chunk size must be a positive integer, given "+"%"+"lld", csize)
		return_value.u1.type_info = 2
		return
	}

	/* stream.chunk_size is actually a size_t, but php_stream_set_option
	 * can only use an int to accept the new value and return the old one.
	 * In any case, values larger than INT_MAX for a chunk size make no sense.
	 */

	if csize > 2147483647 {
		core.PhpErrorDocref(nil, 1<<1, "The chunk size cannot be larger than %d", 2147483647)
		return_value.u1.type_info = 2
		return
	}
	if g.Assign(&stream, (*core.PhpStream)(zend.ZendFetchResource2Ex(zstream, "stream", streams.PhpFileLeStream(), streams.PhpFileLePstream()))) == nil {
		return_value.u1.type_info = 2
		return
	}
	ret = streams._phpStreamSetOption(stream, 5, int(csize), nil)
	var __z *zend.Zval = return_value
	if ret > 0 {
		__z.value.lval = zend.ZendLong(ret)
	} else {
		__z.value.lval = zend.ZendLong(EOF)
	}
	__z.u1.type_info = 4
	return
}

/* }}} */

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

			if zend.ZendParseArgResource(_arg, &arg1, 0) == 0 {
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

			if zend.ZendParseArgLong(_arg, &arg2, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
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
	if g.Assign(&stream, (*core.PhpStream)(zend.ZendFetchResource2Ex(arg1, "stream", streams.PhpFileLeStream(), streams.PhpFileLePstream()))) == nil {
		return_value.u1.type_info = 2
		return
	}
	buff = arg2

	/* if buff is 0 then set to non-buffered */

	if buff == 0 {
		ret = streams._phpStreamSetOption(stream, 2, 0, nil)
	} else {
		ret = streams._phpStreamSetOption(stream, 2, 2, &buff)
	}
	var __z *zend.Zval = return_value
	if ret == 0 {
		__z.value.lval = 0
	} else {
		__z.value.lval = EOF
	}
	__z.u1.type_info = 4
	return
}

/* }}} */

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

			if zend.ZendParseArgResource(_arg, &zstream, 0) == 0 {
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

			if zend.ZendParseArgBool(_arg, &enable, &_dummy, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_BOOL
				_error_code = 4
				break
			}
			_optional = 1
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

			if zend.ZendParseArgLong(_arg, &cryptokind, &cryptokindnull, 1, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
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

			if zend.ZendParseArgResource(_arg, &zsessstream, 0) == 0 {
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
	if g.Assign(&stream, (*core.PhpStream)(zend.ZendFetchResource2Ex(zstream, "stream", streams.PhpFileLeStream(), streams.PhpFileLePstream()))) == nil {
		return_value.u1.type_info = 2
		return
	}
	if enable != 0 {
		if cryptokindnull != 0 {
			var val *zend.Zval
			if !((*core.PhpStreamContext)(g.CondF1(stream.ctx != nil, func() any { return stream.ctx.ptr }, nil)) != nil && nil != g.Assign(&val, streams.PhpStreamContextGetOption((*core.PhpStreamContext)(g.CondF1(stream.ctx != nil, func() any { return stream.ctx.ptr }, nil)), "ssl", "crypto_method"))) {
				core.PhpErrorDocref(nil, 1<<1, "When enabling encryption you must specify the crypto type")
				return_value.u1.type_info = 2
				return
			}
			cryptokind = val.value.lval
		}
		if zsessstream != nil {
			if g.Assign(&sessstream, (*core.PhpStream)(zend.ZendFetchResource2Ex(zsessstream, "stream", streams.PhpFileLeStream(), streams.PhpFileLePstream()))) == nil {
				return_value.u1.type_info = 2
				return
			}
		}
		if streams.PhpStreamXportCryptoSetup(stream, cryptokind, sessstream) < 0 {
			return_value.u1.type_info = 2
			return
		}
	}
	ret = streams.PhpStreamXportCryptoEnable(stream, enable)
	switch ret {
	case -1:
		return_value.u1.type_info = 2
		return
	case 0:
		var __z *zend.Zval = return_value
		__z.value.lval = 0
		__z.u1.type_info = 4
		return
	default:
		return_value.u1.type_info = 3
		return
	}
}

/* }}} */

func ZifStreamResolveIncludePath(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var filename *byte
	var filename_len int
	var resolved_path *zend.ZendString
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

			if zend.ZendParseArgPath(_arg, &filename, &filename_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_PATH
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
			return
		}
		break
	}
	resolved_path = zend.ZendResolvePath(filename, filename_len)
	if resolved_path != nil {
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = resolved_path
		__z.value.str = __s
		if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
			__z.u1.type_info = 6
		} else {
			__z.u1.type_info = 6 | 1<<0<<8
		}
		return
	}
	return_value.u1.type_info = 2
	return
}

/* }}} */

func ZifStreamIsLocal(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var zstream *zend.Zval
	var stream *core.PhpStream = nil
	var wrapper *core.PhpStreamWrapper = nil
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

			zend.ZendParseArgZvalDeref(_arg, &zstream, 0)
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
	if zstream.u1.v.type_ == 9 {
		if g.Assign(&stream, (*core.PhpStream)(zend.ZendFetchResource2Ex(zstream, "stream", streams.PhpFileLeStream(), streams.PhpFileLePstream()))) == nil {
			return_value.u1.type_info = 2
			return
		}
		if stream == nil {
			return_value.u1.type_info = 2
			return
		}
		wrapper = stream.wrapper
	} else {
		if zend.TryConvertToString(zstream) == 0 {
			return
		}
		wrapper = streams.PhpStreamLocateUrlWrapper(zstream.value.str.val, nil, 0)
	}
	if wrapper == nil {
		return_value.u1.type_info = 2
		return
	}
	if wrapper.is_url == 0 {
		return_value.u1.type_info = 3
	} else {
		return_value.u1.type_info = 2
	}
	return
}

/* }}} */

func ZifStreamSupportsLock(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var stream *core.PhpStream
	var zsrc *zend.Zval
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

			if zend.ZendParseArgResource(_arg, &zsrc, 0) == 0 {
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
	if g.Assign(&stream, (*core.PhpStream)(zend.ZendFetchResource2Ex(zsrc, "stream", streams.PhpFileLeStream(), streams.PhpFileLePstream()))) == nil {
		return_value.u1.type_info = 2
		return
	}
	if !(g.Cond(streams._phpStreamSetOption(stream, 6, 0, any(1)) == 0, 1, 0)) {
		return_value.u1.type_info = 2
		return
	}
	return_value.u1.type_info = 3
	return
}

/* {{{ proto bool stream_isatty(resource stream)
Check if a stream is a TTY.
*/

func ZifStreamIsatty(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var zsrc *zend.Zval
	var stream *core.PhpStream
	var fileno core.PhpSocketT
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

			if zend.ZendParseArgResource(_arg, &zsrc, 0) == 0 {
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
	if g.Assign(&stream, (*core.PhpStream)(zend.ZendFetchResource2Ex(zsrc, "stream", streams.PhpFileLeStream(), streams.PhpFileLePstream()))) == nil {
		return_value.u1.type_info = 2
		return
	}
	if streams._phpStreamCast(stream, 3, nil, 0) == zend.SUCCESS {
		streams._phpStreamCast(stream, 3, any(&fileno), 0)
	} else if streams._phpStreamCast(stream, 1, nil, 0) == zend.SUCCESS {
		streams._phpStreamCast(stream, 1, any(&fileno), 0)
	} else {
		return_value.u1.type_info = 2
		return
	}

	/* Check if the file descriptor identifier is a terminal */

	if zend.Isatty(fileno) != 0 {
		return_value.u1.type_info = 3
	} else {
		return_value.u1.type_info = 2
	}

	/* Check if the file descriptor identifier is a terminal */
}

/* {{{ proto int stream_socket_shutdown(resource stream, int how)
   causes all or part of a full-duplex connection on the socket associated
   with stream to be shut down.  If how is SHUT_RD,  further receptions will
   be disallowed. If how is SHUT_WR, further transmissions will be disallowed.
   If how is SHUT_RDWR,  further  receptions and transmissions will be
   disallowed. */

func ZifStreamSocketShutdown(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var how zend.ZendLong
	var zstream *zend.Zval
	var stream *core.PhpStream
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

			if zend.ZendParseArgResource(_arg, &zstream, 0) == 0 {
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

			if zend.ZendParseArgLong(_arg, &how, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
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
	if how != streams.STREAM_SHUT_RD && how != streams.STREAM_SHUT_WR && how != streams.STREAM_SHUT_RDWR {
		core.PhpErrorDocref(nil, 1<<1, "Second parameter $how needs to be one of STREAM_SHUT_RD, STREAM_SHUT_WR or STREAM_SHUT_RDWR")
		return_value.u1.type_info = 2
		return
	}
	if g.Assign(&stream, (*core.PhpStream)(zend.ZendFetchResource2Ex(zstream, "stream", streams.PhpFileLeStream(), streams.PhpFileLePstream()))) == nil {
		return_value.u1.type_info = 2
		return
	}
	if streams.PhpStreamXportShutdown(stream, streams.StreamShutdownT(how)) == 0 {
		return_value.u1.type_info = 3
	} else {
		return_value.u1.type_info = 2
	}
	return
}

/* }}} */
