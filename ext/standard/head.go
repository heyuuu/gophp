// <<generate>>

package standard

import (
	"sik/core"
	r "sik/runtime"
	g "sik/runtime/grammar"
	"sik/zend"
)

// Source: <ext/standard/head.h>

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
   | Author: Rasmus Lerdorf <rasmus@lerdorf.on.ca>                        |
   +----------------------------------------------------------------------+
*/

// #define HEAD_H

// #define COOKIE_EXPIRES       "; expires="

// #define COOKIE_MAX_AGE       "; Max-Age="

// #define COOKIE_DOMAIN       "; domain="

// #define COOKIE_PATH       "; path="

// #define COOKIE_SECURE       "; secure"

// #define COOKIE_HTTPONLY       "; HttpOnly"

// #define COOKIE_SAMESITE       "; SameSite="

var ZmActivateHead func(type_ int, module_number int) int

// Source: <ext/standard/head.c>

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
   | Author: Rasmus Lerdorf <rasmus@lerdorf.on.ca>                        |
   +----------------------------------------------------------------------+
*/

// # include < stdio . h >

// # include "php.h"

// # include "ext/standard/php_standard.h"

// failed # include "ext/date/php_date.h"

// # include "SAPI.h"

// # include "php_main.h"

// # include "head.h"

// # include < time . h >

// # include "php_globals.h"

// # include "zend_smart_str.h"

/* Implementation of the language Header() function */

func ZifHeader(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var rep zend.ZendBool = 1
	var ctr core.SapiHeaderLine = core.SapiHeaderLine{0}
	var len_ int
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
			r.Assert(_i <= _min_num_args || _optional == 1)
			r.Assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgString(_arg, &ctr.line, &len_, 0) == 0 {
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

			if zend.ZendParseArgBool(_arg, &rep, &_dummy, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_BOOL
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

			if zend.ZendParseArgLong(_arg, &ctr.response_code, &_dummy, 0, 0) == 0 {
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
	ctr.line_len = uint32(len_)
	core.SapiHeaderOp(g.Cond(rep != 0, core.SAPI_HEADER_REPLACE, 1<<0), &ctr)
}

/* }}} */

func ZifHeaderRemove(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var ctr core.SapiHeaderLine = core.SapiHeaderLine{0}
	var len_ int = 0
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
			r.Assert(_i <= _min_num_args || _optional == 1)
			r.Assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgString(_arg, &ctr.line, &len_, 0) == 0 {
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
			return
		}
		break
	}
	ctr.line_len = uint32(len_)
	core.SapiHeaderOp(g.Cond(execute_data.This.u2.num_args == 0, core.SAPI_HEADER_DELETE_ALL, core.SAPI_HEADER_DELETE), &ctr)
}

/* }}} */

func PhpHeader() int {
	if core.SapiSendHeaders() == zend.FAILURE || core.sapi_globals.request_info.headers_only != 0 {
		return 0
	} else {
		return 1
	}
}
func PhpSetcookie(name *zend.ZendString, value *zend.ZendString, expires int64, path *zend.ZendString, domain *zend.ZendString, secure int, httponly int, samesite *zend.ZendString, url_encode int) int {
	var dt *zend.ZendString
	var ctr core.SapiHeaderLine = core.SapiHeaderLine{0}
	var result int
	var buf zend.SmartStr = zend.SmartStr{0}
	if name.len_ == 0 {
		zend.ZendError(1<<1, "Cookie names must not be empty")
		return zend.FAILURE
	} else if strpbrk(name.val, "=,; \t\r\n013014") != nil {
		zend.ZendError(1<<1, "Cookie names cannot contain any of the following '=,; \\t\\r\\n\\013\\014'")
		return zend.FAILURE
	}
	if url_encode == 0 && value != nil && strpbrk(value.val, ",; \t\r\n013014") != nil {
		zend.ZendError(1<<1, "Cookie values cannot contain any of the following ',; \\t\\r\\n\\013\\014'")
		return zend.FAILURE
	}
	if path != nil && strpbrk(path.val, ",; \t\r\n013014") != nil {
		zend.ZendError(1<<1, "Cookie paths cannot contain any of the following ',; \\t\\r\\n\\013\\014'")
		return zend.FAILURE
	}
	if domain != nil && strpbrk(domain.val, ",; \t\r\n013014") != nil {
		zend.ZendError(1<<1, "Cookie domains cannot contain any of the following ',; \\t\\r\\n\\013\\014'")
		return zend.FAILURE
	}
	if value == nil || value.len_ == 0 {

		/*
		 * MSIE doesn't delete a cookie when you set it to a null value
		 * so in order to force cookies to be deleted, even on MSIE, we
		 * pick an expiry date in the past
		 */

		dt = php_format_date("D, d-M-Y H:i:s T", g.SizeOf("\"D, d-M-Y H:i:s T\"")-1, 1, 0)
		zend.SmartStrAppendlEx(&buf, "Set-Cookie: ", strlen("Set-Cookie: "), 0)
		zend.SmartStrAppendEx(&buf, name, 0)
		zend.SmartStrAppendlEx(&buf, "=deleted; expires=", strlen("=deleted; expires="), 0)
		zend.SmartStrAppendEx(&buf, dt, 0)
		zend.SmartStrAppendlEx(&buf, "; Max-Age=0", strlen("; Max-Age=0"), 0)
		zend.ZendStringFree(dt)
	} else {
		zend.SmartStrAppendlEx(&buf, "Set-Cookie: ", strlen("Set-Cookie: "), 0)
		zend.SmartStrAppendEx(&buf, name, 0)
		zend.SmartStrAppendcEx(&buf, '=', 0)
		if url_encode != 0 {
			var encoded_value *zend.ZendString = PhpRawUrlEncode(value.val, value.len_)
			zend.SmartStrAppendEx(&buf, encoded_value, 0)
			zend.ZendStringReleaseEx(encoded_value, 0)
		} else {
			zend.SmartStrAppendEx(&buf, value, 0)
		}
		if expires > 0 {
			var p *byte
			var diff float64
			zend.SmartStrAppendlEx(&buf, "; expires=", strlen("; expires="), 0)
			dt = php_format_date("D, d-M-Y H:i:s T", g.SizeOf("\"D, d-M-Y H:i:s T\"")-1, expires, 0)

			/* check to make sure that the year does not exceed 4 digits in length */

			p = zend.ZendMemrchr(dt.val, '-', dt.len_)
			if p == nil || (*(p + 5)) != ' ' {
				zend.ZendStringFree(dt)
				zend.SmartStrFreeEx(&buf, 0)
				zend.ZendError(1<<1, "Expiry date cannot have a year greater than 9999")
				return zend.FAILURE
			}
			zend.SmartStrAppendEx(&buf, dt, 0)
			zend.ZendStringFree(dt)
			diff = difftime(expires, php_time())
			if diff < 0 {
				diff = 0
			}
			zend.SmartStrAppendlEx(&buf, "; Max-Age=", strlen("; Max-Age="), 0)
			zend.SmartStrAppendLongEx(&buf, zend.ZendLong(diff), 0)
		}
	}
	if path != nil && path.len_ != 0 {
		zend.SmartStrAppendlEx(&buf, "; path=", strlen("; path="), 0)
		zend.SmartStrAppendEx(&buf, path, 0)
	}
	if domain != nil && domain.len_ != 0 {
		zend.SmartStrAppendlEx(&buf, "; domain=", strlen("; domain="), 0)
		zend.SmartStrAppendEx(&buf, domain, 0)
	}
	if secure != 0 {
		zend.SmartStrAppendlEx(&buf, "; secure", strlen("; secure"), 0)
	}
	if httponly != 0 {
		zend.SmartStrAppendlEx(&buf, "; HttpOnly", strlen("; HttpOnly"), 0)
	}
	if samesite != nil && samesite.len_ != 0 {
		zend.SmartStrAppendlEx(&buf, "; SameSite=", strlen("; SameSite="), 0)
		zend.SmartStrAppendEx(&buf, samesite, 0)
	}
	ctr.line = buf.s.val
	ctr.line_len = uint32(buf.s).len_
	result = core.SapiHeaderOp(1<<0, &ctr)
	zend.ZendStringRelease(buf.s)
	return result
}
func PhpHeadParseCookieOptionsArray(options *zend.Zval, expires *zend.ZendLong, path **zend.ZendString, domain **zend.ZendString, secure *zend.ZendBool, httponly *zend.ZendBool, samesite **zend.ZendString) {
	var found int = 0
	var key *zend.ZendString
	var value *zend.Zval
	for {
		var __ht *zend.HashTable = options.value.arr
		var _p *zend.Bucket = __ht.arData
		var _end *zend.Bucket = _p + __ht.nNumUsed
		for ; _p != _end; _p++ {
			var _z *zend.Zval = &_p.val

			if _z.u1.v.type_ == 0 {
				continue
			}
			key = _p.key
			value = _z
			if key != nil {
				if key.len_ == g.SizeOf("\"expires\"")-1 && zend.ZendBinaryStrcasecmp(key.val, key.len_, "expires", g.SizeOf("\"expires\"")-1) == 0 {
					*expires = zend.ZvalGetLong(value)
					found++
				} else if key.len_ == g.SizeOf("\"path\"")-1 && zend.ZendBinaryStrcasecmp(key.val, key.len_, "path", g.SizeOf("\"path\"")-1) == 0 {
					*path = zend.ZvalGetString(value)
					found++
				} else if key.len_ == g.SizeOf("\"domain\"")-1 && zend.ZendBinaryStrcasecmp(key.val, key.len_, "domain", g.SizeOf("\"domain\"")-1) == 0 {
					*domain = zend.ZvalGetString(value)
					found++
				} else if key.len_ == g.SizeOf("\"secure\"")-1 && zend.ZendBinaryStrcasecmp(key.val, key.len_, "secure", g.SizeOf("\"secure\"")-1) == 0 {
					*secure = zend.ZendIsTrue(value)
					found++
				} else if key.len_ == g.SizeOf("\"httponly\"")-1 && zend.ZendBinaryStrcasecmp(key.val, key.len_, "httponly", g.SizeOf("\"httponly\"")-1) == 0 {
					*httponly = zend.ZendIsTrue(value)
					found++
				} else if key.len_ == g.SizeOf("\"samesite\"")-1 && zend.ZendBinaryStrcasecmp(key.val, key.len_, "samesite", g.SizeOf("\"samesite\"")-1) == 0 {
					*samesite = zend.ZvalGetString(value)
					found++
				} else {
					core.PhpErrorDocref(nil, 1<<1, "Unrecognized key '%s' found in the options array", key.val)
				}
			} else {
				core.PhpErrorDocref(nil, 1<<1, "Numeric key found in the options array")
			}
		}
		break
	}

	/* Array is not empty but no valid keys were found */

	if found == 0 && options.value.arr.nNumOfElements > 0 {
		core.PhpErrorDocref(nil, 1<<1, "No valid options were found in the given array")
	}

	/* Array is not empty but no valid keys were found */
}

/* {{{ proto bool setcookie(string name [, string value [, int expires [, string path [, string domain [, bool secure[, bool httponly]]]]]])
               setcookie(string name [, string value [, array options]])
Send a cookie */

func ZifSetcookie(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var expires_or_options *zend.Zval = nil
	var name *zend.ZendString
	var value *zend.ZendString = nil
	var path *zend.ZendString = nil
	var domain *zend.ZendString = nil
	var samesite *zend.ZendString = nil
	var expires zend.ZendLong = 0
	var secure zend.ZendBool = 0
	var httponly zend.ZendBool = 0
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 7
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

			if zend.ZendParseArgStr(_arg, &name, 0) == 0 {
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

			if zend.ZendParseArgStr(_arg, &value, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
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

			zend.ZendParseArgZvalDeref(_arg, &expires_or_options, 0)
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

			if zend.ZendParseArgStr(_arg, &path, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
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

			if zend.ZendParseArgStr(_arg, &domain, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
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

			if zend.ZendParseArgBool(_arg, &secure, &_dummy, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_BOOL
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

			if zend.ZendParseArgBool(_arg, &httponly, &_dummy, 0) == 0 {
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
	if expires_or_options != nil {
		if expires_or_options.u1.v.type_ == 7 {
			if execute_data.This.u2.num_args > 3 {
				core.PhpErrorDocref(nil, 1<<1, "Cannot pass arguments after the options array")
				return_value.u1.type_info = 2
				return
			}
			PhpHeadParseCookieOptionsArray(expires_or_options, &expires, &path, &domain, &secure, &httponly, &samesite)
		} else {
			expires = zend.ZvalGetLong(expires_or_options)
		}
	}
	if zend.EG.exception == nil {
		if PhpSetcookie(name, value, expires, path, domain, secure, httponly, samesite, 1) == zend.SUCCESS {
			return_value.u1.type_info = 3
		} else {
			return_value.u1.type_info = 2
		}
	}
	if expires_or_options != nil && expires_or_options.u1.v.type_ == 7 {
		if path != nil {
			zend.ZendStringRelease(path)
		}
		if domain != nil {
			zend.ZendStringRelease(domain)
		}
		if samesite != nil {
			zend.ZendStringRelease(samesite)
		}
	}
}

/* }}} */

func ZifSetrawcookie(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var expires_or_options *zend.Zval = nil
	var name *zend.ZendString
	var value *zend.ZendString = nil
	var path *zend.ZendString = nil
	var domain *zend.ZendString = nil
	var samesite *zend.ZendString = nil
	var expires zend.ZendLong = 0
	var secure zend.ZendBool = 0
	var httponly zend.ZendBool = 0
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 7
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

			if zend.ZendParseArgStr(_arg, &name, 0) == 0 {
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

			if zend.ZendParseArgStr(_arg, &value, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
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

			zend.ZendParseArgZvalDeref(_arg, &expires_or_options, 0)
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

			if zend.ZendParseArgStr(_arg, &path, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
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

			if zend.ZendParseArgStr(_arg, &domain, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
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

			if zend.ZendParseArgBool(_arg, &secure, &_dummy, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_BOOL
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

			if zend.ZendParseArgBool(_arg, &httponly, &_dummy, 0) == 0 {
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
	if expires_or_options != nil {
		if expires_or_options.u1.v.type_ == 7 {
			if execute_data.This.u2.num_args > 3 {
				core.PhpErrorDocref(nil, 1<<1, "Cannot pass arguments after the options array")
				return_value.u1.type_info = 2
				return
			}
			PhpHeadParseCookieOptionsArray(expires_or_options, &expires, &path, &domain, &secure, &httponly, &samesite)
		} else {
			expires = zend.ZvalGetLong(expires_or_options)
		}
	}
	if zend.EG.exception == nil {
		if PhpSetcookie(name, value, expires, path, domain, secure, httponly, samesite, 0) == zend.SUCCESS {
			return_value.u1.type_info = 3
		} else {
			return_value.u1.type_info = 2
		}
	}
	if expires_or_options != nil && expires_or_options.u1.v.type_ == 7 {
		if path != nil {
			zend.ZendStringRelease(path)
		}
		if domain != nil {
			zend.ZendStringRelease(domain)
		}
		if samesite != nil {
			zend.ZendStringRelease(samesite)
		}
	}
}

/* }}} */

func ZifHeadersSent(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var arg1 *zend.Zval = nil
	var arg2 *zend.Zval = nil
	var file *byte = ""
	var line int = 0
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
			r.Assert(_i <= _min_num_args || _optional == 1)
			r.Assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			zend.ZendParseArgZvalDeref(_arg, &arg1, 0)
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

			zend.ZendParseArgZvalDeref(_arg, &arg2, 0)
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
	if core.sapi_globals.headers_sent != 0 {
		line = core.PhpOutputGetStartLineno()
		file = core.PhpOutputGetStartFilename()
	}
	switch execute_data.This.u2.num_args {
	case 2:
		for {
			r.Assert(arg2.u1.v.type_ == 10)
			for {
				var _zv *zend.Zval = arg2
				var ref *zend.ZendReference = _zv.value.ref
				if ref.sources.ptr != nil {
					zend.ZendTryAssignTypedRefLong(ref, line)
					break
				}
				_zv = &ref.val
				zend.ZvalPtrDtor(_zv)
				var __z *zend.Zval = _zv
				__z.value.lval = line
				__z.u1.type_info = 4
				break
			}
			break
		}
	case 1:
		if file != nil {
			for {
				r.Assert(arg1.u1.v.type_ == 10)
				for {
					var _zv *zend.Zval = arg1
					var ref *zend.ZendReference = _zv.value.ref
					if ref.sources.ptr != nil {
						zend.ZendTryAssignTypedRefString(ref, file)
						break
					}
					_zv = &ref.val
					zend.ZvalPtrDtor(_zv)
					var _s *byte = file
					var __z *zend.Zval = _zv
					var __s *zend.ZendString = zend.ZendStringInit(_s, strlen(_s), 0)
					__z.value.str = __s
					__z.u1.type_info = 6 | 1<<0<<8
					break
				}
				break
			}
		} else {
			for {
				r.Assert(arg1.u1.v.type_ == 10)
				for {
					var _zv *zend.Zval = arg1
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
		break
	}
	if core.sapi_globals.headers_sent != 0 {
		return_value.u1.type_info = 3
		return
	} else {
		return_value.u1.type_info = 2
		return
	}
}

/* }}} */

func PhpHeadApplyHeaderListToHash(data any, arg any) {
	var sapi_header *core.SapiHeader = (*core.SapiHeader)(data)
	if arg && sapi_header != nil {
		zend.AddNextIndexString((*zend.Zval)(arg), (*byte)(sapi_header.header))
	}
}

/* {{{ proto array headers_list(void)
   Return list of headers to be sent / already sent */

func ZifHeadersList(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	var __arr *zend.ZendArray = zend._zendNewArray(0)
	var __z *zend.Zval = return_value
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	zend.ZendLlistApplyWithArgument(&(core.sapi_globals.sapi_headers).headers, PhpHeadApplyHeaderListToHash, return_value)
}

/* }}} */

func ZifHttpResponseCode(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var response_code zend.ZendLong = 0
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
			r.Assert(_i <= _min_num_args || _optional == 1)
			r.Assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgLong(_arg, &response_code, &_dummy, 0, 0) == 0 {
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
	if response_code != 0 {
		var old_response_code zend.ZendLong
		old_response_code = core.sapi_globals.sapi_headers.http_response_code
		core.sapi_globals.sapi_headers.http_response_code = int(response_code)
		if old_response_code != 0 {
			var __z *zend.Zval = return_value
			__z.value.lval = old_response_code
			__z.u1.type_info = 4
			return
		}
		return_value.u1.type_info = 3
		return
	}
	if core.sapi_globals.sapi_headers.http_response_code == 0 {
		return_value.u1.type_info = 2
		return
	}
	var __z *zend.Zval = return_value
	__z.value.lval = core.sapi_globals.sapi_headers.http_response_code
	__z.u1.type_info = 4
	return
}

/* }}} */
