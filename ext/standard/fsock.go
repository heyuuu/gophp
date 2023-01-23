// <<generate>>

package standard

import (
	"sik/core"
	"sik/core/streams"
	r "sik/runtime"
	g "sik/runtime/grammar"
	"sik/zend"
)

// Source: <ext/standard/fsock.h>

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
   | Authors: Paul Panotzki - Bunyip Information Systems                  |
   |          Jim Winstead <jimw@php.net>                                 |
   |          Wez Furlong                                                 |
   +----------------------------------------------------------------------+
*/

// #define FSOCK_H

// # include "file.h"

// # include "php_network.h"

// Source: <ext/standard/fsock.c>

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
   | Authors: Paul Panotzki - Bunyip Information Systems                  |
   |          Jim Winstead <jimw@php.net>                                 |
   |          Sascha Schumann <sascha@schumann.cx>                        |
   +----------------------------------------------------------------------+
*/

// # include "php.h"

// # include "php_globals.h"

// # include < stdlib . h >

// # include < stddef . h >

// # include "php_network.h"

// # include "file.h"

/* {{{ php_fsockopen() */

func PhpFsockopenStream(execute_data *zend.ZendExecuteData, return_value *zend.Zval, persistent int) {
	var host *byte
	var host_len int
	var port zend.ZendLong = -1
	var zerrno *zend.Zval = nil
	var zerrstr *zend.Zval = nil
	var timeout float64 = float64(FileGlobals.GetDefaultSocketTimeout())
	var conv int64
	var tv __struct__timeval
	var hashkey *byte = nil
	var stream *core.PhpStream = nil
	var err int
	var hostname *byte = nil
	var hostname_len int
	var errstr *zend.ZendString = nil
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
			r.Assert(_i <= _min_num_args || _optional == 1)
			r.Assert(_i > _min_num_args || _optional == 0)
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
			r.Assert(_i <= _min_num_args || _optional == 1)
			r.Assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgLong(_arg, &port, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = 4
				break
			}
			_i++
			r.Assert(_i <= _min_num_args || _optional == 1)
			r.Assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			zend.ZendParseArgZvalDeref(_arg, &zerrno, 0)
			_i++
			r.Assert(_i <= _min_num_args || _optional == 1)
			r.Assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			zend.ZendParseArgZvalDeref(_arg, &zerrstr, 0)
			_i++
			r.Assert(_i <= _min_num_args || _optional == 1)
			r.Assert(_i > _min_num_args || _optional == 0)
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
	if persistent != 0 {
		zend.ZendSpprintf(&hashkey, 0, "pfsockopen__%s:"+"%"+"lld", host, port)
	}
	if port > 0 {
		hostname_len = zend.ZendSpprintf(&hostname, 0, "%s:"+"%"+"lld", host, port)
	} else {
		hostname_len = host_len
		hostname = host
	}

	/* prepare the timeout value for use */

	conv = time_t(timeout * 1000000.0)
	tv.tv_sec = conv / 1000000
	tv.tv_usec = conv % 1000000
	stream = streams._phpStreamXportCreate(hostname, hostname_len, 0x8, 0|2, hashkey, &tv, nil, &errstr, &err)
	if port > 0 {
		zend._efree(hostname)
	}
	if stream == nil {
		core.PhpErrorDocref(nil, 1<<1, "unable to connect to %s:"+"%"+"lld"+" (%s)", host, port, g.CondF2(errstr == nil, "Unknown error", func() []byte { return errstr.val }))
	}
	if hashkey != nil {
		zend._efree(hashkey)
	}
	if stream == nil {
		if zerrno != nil {
			for {
				r.Assert(zerrno.u1.v.type_ == 10)
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
		if errstr != nil {
			if zerrstr != nil {
				for {
					r.Assert(zerrstr.u1.v.type_ == 10)
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
			} else {
				zend.ZendStringRelease(errstr)
			}
		}
		return_value.u1.type_info = 2
		return
	}
	if zerrno != nil {
		for {
			r.Assert(zerrno.u1.v.type_ == 10)
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
			r.Assert(zerrstr.u1.v.type_ == 10)
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
	if errstr != nil {
		zend.ZendStringReleaseEx(errstr, 0)
	}
	var __z *zend.Zval = return_value
	__z.value.res = stream.res
	__z.u1.type_info = 9 | 1<<0<<8
	stream.__exposed = 1
}

/* }}} */

func ZifFsockopen(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpFsockopenStream(execute_data, return_value, 0)
}

/* }}} */

func ZifPfsockopen(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpFsockopenStream(execute_data, return_value, 1)
}

/* }}} */
