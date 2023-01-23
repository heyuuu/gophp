// <<generate>>

package standard

import (
	"sik/core"
	r "sik/runtime"
	g "sik/runtime/grammar"
	"sik/zend"
)

// Source: <ext/standard/string.c>

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
   | Authors: Rasmus Lerdorf <rasmus@php.net>                             |
   |          Stig Sæther Bakken <ssb@php.net>                          |
   |          Zeev Suraski <zeev@php.net>                                 |
   +----------------------------------------------------------------------+
*/

// # include < stdio . h >

// # include "php.h"

// # include "php_rand.h"

// # include "php_string.h"

// # include "php_variables.h"

// # include < locale . h >

// # include < langinfo . h >

// # include < monetary . h >

/*
 * This define is __special__  here because some versions of libintl redefine setlocale __special__
 * to point to libintl_setlocale.  That's a ridiculous thing to do as far
 * as I am concerned, but with this define and __special__  the subsequent undef we
 * limit the damage to just the actual setlocale() call in this file
 * without turning zif_setlocale into zif_libintl_setlocale.  -Rasmus
 */

// #define php_my_setlocale       setlocale

// # include "scanf.h"

// # include "zend_API.h"

// # include "zend_execute.h"

// # include "php_globals.h"

// # include "basic_functions.h"

// # include "zend_smart_str.h"

// # include < Zend / zend_exceptions . h >

/* For str_getcsv() support */

// # include "ext/standard/file.h"

/* For php_next_utf8_char() */

// # include "ext/standard/html.h"

// #define STR_PAD_LEFT       0

// #define STR_PAD_RIGHT       1

// #define STR_PAD_BOTH       2

// #define PHP_PATHINFO_DIRNAME       1

// #define PHP_PATHINFO_BASENAME       2

// #define PHP_PATHINFO_EXTENSION       4

// #define PHP_PATHINFO_FILENAME       8

// #define PHP_PATHINFO_ALL       ( PHP_PATHINFO_DIRNAME | PHP_PATHINFO_BASENAME | PHP_PATHINFO_EXTENSION | PHP_PATHINFO_FILENAME )

// #define STR_STRSPN       0

// #define STR_STRCSPN       1

/* {{{ register_string_constants
 */

func RegisterStringConstants(type_ int, module_number int) {
	zend.ZendRegisterLongConstant("STR_PAD_LEFT", g.SizeOf("\"STR_PAD_LEFT\"")-1, 0, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("STR_PAD_RIGHT", g.SizeOf("\"STR_PAD_RIGHT\"")-1, 1, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("STR_PAD_BOTH", g.SizeOf("\"STR_PAD_BOTH\"")-1, 2, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("PATHINFO_DIRNAME", g.SizeOf("\"PATHINFO_DIRNAME\"")-1, 1, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("PATHINFO_BASENAME", g.SizeOf("\"PATHINFO_BASENAME\"")-1, 2, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("PATHINFO_EXTENSION", g.SizeOf("\"PATHINFO_EXTENSION\"")-1, 4, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("PATHINFO_FILENAME", g.SizeOf("\"PATHINFO_FILENAME\"")-1, 8, 1<<0|1<<1, module_number)

	/* If last members of struct lconv equal CHAR_MAX, no grouping is done */

	zend.ZendRegisterLongConstant("CHAR_MAX", g.SizeOf("\"CHAR_MAX\"")-1, CHAR_MAX, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("LC_CTYPE", g.SizeOf("\"LC_CTYPE\"")-1, LC_CTYPE, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("LC_NUMERIC", g.SizeOf("\"LC_NUMERIC\"")-1, LC_NUMERIC, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("LC_TIME", g.SizeOf("\"LC_TIME\"")-1, LC_TIME, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("LC_COLLATE", g.SizeOf("\"LC_COLLATE\"")-1, LC_COLLATE, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("LC_MONETARY", g.SizeOf("\"LC_MONETARY\"")-1, LC_MONETARY, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("LC_ALL", g.SizeOf("\"LC_ALL\"")-1, LC_ALL, 1<<0|1<<1, module_number)
}

/* }}} */

/* this is read-only, so it's ok */

var Hexconvtab []byte = "0123456789abcdef"

/* localeconv mutex */

/* {{{ php_bin2hex
 */

func PhpBin2hex(old *uint8, oldlen int) *zend.ZendString {
	var result *zend.ZendString
	var i int
	var j int
	result = zend.ZendStringSafeAlloc(oldlen, 2*g.SizeOf("char"), 0, 0)
	j = 0
	i = j
	for ; i < oldlen; i++ {
		result.val[g.PostInc(&j)] = Hexconvtab[old[i]>>4]
		result.val[g.PostInc(&j)] = Hexconvtab[old[i]&15]
	}
	result.val[j] = '0'
	return result
}

/* }}} */

func PhpHex2bin(old *uint8, oldlen int) *zend.ZendString {
	var target_length int = oldlen >> 1
	var str *zend.ZendString = zend.ZendStringAlloc(target_length, 0)
	var ret *uint8 = (*uint8)(str.val)
	var i int
	var j int
	j = 0
	i = j
	for ; i < target_length; i++ {
		var c uint8 = old[g.PostInc(&j)]
		var l uint8 = c & ^0x20
		var is_letter int = uint(l-'A'^l-'F'-1)>>8*g.SizeOf("unsigned int") - 1
		var d uint8

		/* basically (c >= '0' && c <= '9') || (l >= 'A' && l <= 'F') */

		if ((c ^ '0') - 10>>8*g.SizeOf("unsigned int") - 1 | is_letter) != 0 {
			d = l - 0x10 - 0x27*is_letter<<4
		} else {
			zend.ZendStringEfree(str)
			return nil
		}
		c = old[g.PostInc(&j)]
		l = c & ^0x20
		is_letter = uint(l-'A'^l-'F'-1)>>8*g.SizeOf("unsigned int") - 1
		if ((c ^ '0') - 10>>8*g.SizeOf("unsigned int") - 1 | is_letter) != 0 {
			d |= l - 0x10 - 0x27*is_letter
		} else {
			zend.ZendStringEfree(str)
			return nil
		}
		ret[i] = d
	}
	ret[i] = '0'
	return str
}

/* }}} */

func LocaleconvR(out *__struct__lconv) *__struct__lconv {
	/*  cur->locinfo is struct __crt_locale_info which implementation is
	    hidden in vc14. TODO revisit this and check if a workaround available
	    and needed. */

	/* localeconv doesn't return an error condition */

	*out = (*localeconv)()
	return out
}

/* }}} */

/* {{{ proto string bin2hex(string data)
   Converts the binary representation of data to hex */

func ZifBin2hex(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var result *zend.ZendString
	var data *zend.ZendString
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
			r.Assert(_i <= _min_num_args || _optional == 1)
			r.Assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgStr(_arg, &data, 0) == 0 {
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
	result = PhpBin2hex((*uint8)(data.val), data.len_)
	if result == nil {
		return_value.u1.type_info = 2
		return
	}
	var __z *zend.Zval = return_value
	var __s *zend.ZendString = result
	__z.value.str = __s
	if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
		__z.u1.type_info = 6
	} else {
		__z.u1.type_info = 6 | 1<<0<<8
	}
	return
}

/* }}} */

func ZifHex2bin(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var result *zend.ZendString
	var data *zend.ZendString
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
			r.Assert(_i <= _min_num_args || _optional == 1)
			r.Assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgStr(_arg, &data, 0) == 0 {
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
	if data.len_%2 != 0 {
		core.PhpErrorDocref(nil, 1<<1, "Hexadecimal input string must have an even length")
		return_value.u1.type_info = 2
		return
	}
	result = PhpHex2bin((*uint8)(data.val), data.len_)
	if result == nil {
		core.PhpErrorDocref(nil, 1<<1, "Input string must be hexadecimal string")
		return_value.u1.type_info = 2
		return
	}
	var __z *zend.Zval = return_value
	var __s *zend.ZendString = result
	__z.value.str = __s
	if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
		__z.u1.type_info = 6
	} else {
		__z.u1.type_info = 6 | 1<<0<<8
	}
}

/* }}} */

func PhpSpnCommonHandler(execute_data *zend.ZendExecuteData, return_value *zend.Zval, behavior int) {
	var s11 *zend.ZendString
	var s22 *zend.ZendString
	var start zend.ZendLong = 0
	var len_ zend.ZendLong = 0
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
			r.Assert(_i <= _min_num_args || _optional == 1)
			r.Assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgStr(_arg, &s11, 0) == 0 {
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

			if zend.ZendParseArgStr(_arg, &s22, 0) == 0 {
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

			if zend.ZendParseArgLong(_arg, &start, &_dummy, 0, 0) == 0 {
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

			if zend.ZendParseArgLong(_arg, &len_, &_dummy, 0, 0) == 0 {
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
	if execute_data.This.u2.num_args < 4 {
		len_ = s11.len_
	}

	/* look at substr() function for more information */

	if start < 0 {
		start += zend_long(s11).len_
		if start < 0 {
			start = 0
		}
	} else if int(start > s11.len_) != 0 {
		return_value.u1.type_info = 2
		return
	}
	if len_ < 0 {
		len_ += s11.len_ - start
		if len_ < 0 {
			len_ = 0
		}
	}
	if len_ > zend_long(s11).len_-start {
		len_ = s11.len_ - start
	}
	if len_ == 0 {
		var __z *zend.Zval = return_value
		__z.value.lval = 0
		__z.u1.type_info = 4
		return
	}
	if behavior == 0 {
		var __z *zend.Zval = return_value
		__z.value.lval = PhpStrspn(s11.val+start, s22.val, s11.val+start+len_, s22.val+s22.len_)
		__z.u1.type_info = 4
		return
	} else if behavior == 1 {
		var __z *zend.Zval = return_value
		__z.value.lval = PhpStrcspn(s11.val+start, s22.val, s11.val+start+len_, s22.val+s22.len_)
		__z.u1.type_info = 4
		return
	}
}

/* }}} */

func ZifStrspn(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpSpnCommonHandler(execute_data, return_value, 0)
}

/* }}} */

func ZifStrcspn(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpSpnCommonHandler(execute_data, return_value, 1)
}

/* }}} */

/* }}} */

func ZifStrcoll(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var s1 *zend.ZendString
	var s2 *zend.ZendString
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
			r.Assert(_i <= _min_num_args || _optional == 1)
			r.Assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgStr(_arg, &s1, 0) == 0 {
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

			if zend.ZendParseArgStr(_arg, &s2, 0) == 0 {
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
	var __z *zend.Zval = return_value
	__z.value.lval = strcoll((*byte)(s1.val), (*byte)(s2.val))
	__z.u1.type_info = 4
	return
}

/* }}} */

func PhpCharmask(input *uint8, len_ int, mask *byte) int {
	var end *uint8
	var c uint8
	var result int = zend.SUCCESS
	memset(mask, 0, 256)
	for end = input + len_; input < end; input++ {
		c = *input
		if input+3 < end && input[1] == '.' && input[2] == '.' && input[3] >= c {
			memset(mask+c, 1, input[3]-c+1)
			input += 3
		} else if input+1 < end && input[0] == '.' && input[1] == '.' {

			/* Error, try to be as helpful as possible:
			   (a range ending/starting with '.' won't be captured here) */

			if end-len_ >= input {
				core.PhpErrorDocref(nil, 1<<1, "Invalid '..'-range, no character to the left of '..'")
				result = zend.FAILURE
				continue
			}
			if input+2 >= end {
				core.PhpErrorDocref(nil, 1<<1, "Invalid '..'-range, no character to the right of '..'")
				result = zend.FAILURE
				continue
			}
			if input[-1] > input[2] {
				core.PhpErrorDocref(nil, 1<<1, "Invalid '..'-range, '..'-range needs to be incrementing")
				result = zend.FAILURE
				continue
			}

			/* FIXME: better error (a..b..c is the only left possibility?) */

			core.PhpErrorDocref(nil, 1<<1, "Invalid '..'-range")
			result = zend.FAILURE
			continue
		} else {
			mask[c] = 1
		}
	}
	return result
}

/* }}} */

func PhpTrimInt(str *zend.ZendString, what *byte, what_len int, mode int) *zend.ZendString {
	var start *byte = str.val
	var end *byte = start + str.len_
	var mask []byte
	if what != nil {
		if what_len == 1 {
			var p byte = *what
			if (mode & 1) != 0 {
				for start != end {
					if (*start) == p {
						start++
					} else {
						break
					}
				}
			}
			if (mode & 2) != 0 {
				for start != end {
					if (*(end - 1)) == p {
						end--
					} else {
						break
					}
				}
			}
		} else {
			PhpCharmask((*uint8)(what), what_len, mask)
			if (mode & 1) != 0 {
				for start != end {
					if mask[uint8(*start)] {
						start++
					} else {
						break
					}
				}
			}
			if (mode & 2) != 0 {
				for start != end {
					if mask[uint8(*(end - 1))] {
						end--
					} else {
						break
					}
				}
			}
		}
	} else {
		if (mode & 1) != 0 {
			for start != end {
				var c uint8 = uint8(*start)
				if c <= ' ' && (c == ' ' || c == '\n' || c == '\r' || c == '\t' || c == 'v' || c == '0') {
					start++
				} else {
					break
				}
			}
		}
		if (mode & 2) != 0 {
			for start != end {
				var c uint8 = uint8(*(end - 1))
				if c <= ' ' && (c == ' ' || c == '\n' || c == '\r' || c == '\t' || c == 'v' || c == '0') {
					end--
				} else {
					break
				}
			}
		}
	}
	if str.len_ == end-start {
		return zend.ZendStringCopy(str)
	} else if end-start == 0 {
		return zend.ZendEmptyString
	} else {
		return zend.ZendStringInit(start, end-start, 0)
	}
}

/* }}} */

func PhpTrim(str *zend.ZendString, what *byte, what_len int, mode int) *zend.ZendString {
	return PhpTrimInt(str, what, what_len, mode)
}

/* }}} */

func PhpDoTrim(execute_data *zend.ZendExecuteData, return_value *zend.Zval, mode int) {
	var str *zend.ZendString
	var what *zend.ZendString = nil
	for {
		var _flags int = 0
		var _min_num_args int = 1
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
			r.Assert(_i <= _min_num_args || _optional == 1)
			r.Assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgStr(_arg, &str, 0) == 0 {
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

			if zend.ZendParseArgStr(_arg, &what, 0) == 0 {
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
	var __z *zend.Zval = return_value
	var __s *zend.ZendString = PhpTrimInt(str, g.CondF1(what != nil, func() []byte { return what.val }, nil), g.CondF1(what != nil, func() int { return what.len_ }, 0), mode)
	__z.value.str = __s
	if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
		__z.u1.type_info = 6
	} else {
		__z.u1.type_info = 6 | 1<<0<<8
	}
}

/* }}} */

func ZifTrim(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpDoTrim(execute_data, return_value, 3)
}

/* }}} */

func ZifRtrim(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpDoTrim(execute_data, return_value, 2)
}

/* }}} */

func ZifLtrim(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpDoTrim(execute_data, return_value, 1)
}

/* }}} */

func ZifWordwrap(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var text *zend.ZendString
	var breakchar *byte = "\n"
	var newtextlen int
	var chk int
	var breakchar_len int = 1
	var alloced int
	var current zend.ZendLong = 0
	var laststart zend.ZendLong = 0
	var lastspace zend.ZendLong = 0
	var linelength zend.ZendLong = 75
	var docut zend.ZendBool = 0
	var newtext *zend.ZendString
	for {
		var _flags int = 0
		var _min_num_args int = 1
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
			r.Assert(_i <= _min_num_args || _optional == 1)
			r.Assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgStr(_arg, &text, 0) == 0 {
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

			if zend.ZendParseArgLong(_arg, &linelength, &_dummy, 0, 0) == 0 {
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

			if zend.ZendParseArgString(_arg, &breakchar, &breakchar_len, 0) == 0 {
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

			if zend.ZendParseArgBool(_arg, &docut, &_dummy, 0) == 0 {
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
	if text.len_ == 0 {
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = zend.ZendEmptyString
		__z.value.str = __s
		__z.u1.type_info = 6
		return
	}
	if breakchar_len == 0 {
		core.PhpErrorDocref(nil, 1<<1, "Break string cannot be empty")
		return_value.u1.type_info = 2
		return
	}
	if linelength == 0 && docut != 0 {
		core.PhpErrorDocref(nil, 1<<1, "Can't force cut when width is zero")
		return_value.u1.type_info = 2
		return
	}

	/* Special case for a single-character break as it needs no
	   additional storage space */

	if breakchar_len == 1 && docut == 0 {
		newtext = zend.ZendStringInit(text.val, text.len_, 0)
		lastspace = 0
		laststart = lastspace
		for current = 0; current < zend_long(text).len_; current++ {
			if text.val[current] == breakchar[0] {
				lastspace = current + 1
				laststart = lastspace
			} else if text.val[current] == ' ' {
				if current-laststart >= linelength {
					newtext.val[current] = breakchar[0]
					laststart = current + 1
				}
				lastspace = current
			} else if current-laststart >= linelength && laststart != lastspace {
				newtext.val[lastspace] = breakchar[0]
				laststart = lastspace + 1
			}
		}
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = newtext
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
		return
	} else {

		/* Multiple character line break or forced cut */

		if linelength > 0 {
			chk = size_t(text.len_/linelength + 1)
			newtext = zend.ZendStringSafeAlloc(chk, breakchar_len, text.len_, 0)
			alloced = text.len_ + chk*breakchar_len + 1
		} else {
			chk = text.len_
			alloced = text.len_*(breakchar_len+1) + 1
			newtext = zend.ZendStringSafeAlloc(text.len_, breakchar_len+1, 0, 0)
		}

		/* now keep track of the actual new text length */

		newtextlen = 0
		lastspace = 0
		laststart = lastspace
		for current = 0; current < zend_long(text).len_; current++ {
			if chk == 0 {
				alloced += size_t(((text.len_-current+1)/linelength+1)*breakchar_len) + 1
				newtext = zend.ZendStringExtend(newtext, alloced, 0)
				chk = size_t((text.len_-current)/linelength) + 1
			}

			/* when we hit an existing break, copy to new buffer, and
			 * fix up laststart and lastspace */

			if text.val[current] == breakchar[0] && current+breakchar_len < text.len_ && !(strncmp(text.val+current, breakchar, breakchar_len)) {
				memcpy(newtext.val+newtextlen, text.val+laststart, current-laststart+breakchar_len)
				newtextlen += current - laststart + breakchar_len
				current += breakchar_len - 1
				lastspace = current + 1
				laststart = lastspace
				chk--
			} else if text.val[current] == ' ' {
				if current-laststart >= linelength {
					memcpy(newtext.val+newtextlen, text.val+laststart, current-laststart)
					newtextlen += current - laststart
					memcpy(newtext.val+newtextlen, breakchar, breakchar_len)
					newtextlen += breakchar_len
					laststart = current + 1
					chk--
				}
				lastspace = current
			} else if current-laststart >= linelength && docut != 0 && laststart >= lastspace {
				memcpy(newtext.val+newtextlen, text.val+laststart, current-laststart)
				newtextlen += current - laststart
				memcpy(newtext.val+newtextlen, breakchar, breakchar_len)
				newtextlen += breakchar_len
				lastspace = current
				laststart = lastspace
				chk--
			} else if current-laststart >= linelength && laststart < lastspace {
				memcpy(newtext.val+newtextlen, text.val+laststart, lastspace-laststart)
				newtextlen += lastspace - laststart
				memcpy(newtext.val+newtextlen, breakchar, breakchar_len)
				newtextlen += breakchar_len
				lastspace = lastspace + 1
				laststart = lastspace
				chk--
			}

			/* when we hit an existing break, copy to new buffer, and
			 * fix up laststart and lastspace */

		}

		/* copy over any stragglers */

		if laststart != current {
			memcpy(newtext.val+newtextlen, text.val+laststart, current-laststart)
			newtextlen += current - laststart
		}
		newtext.val[newtextlen] = '0'

		/* free unused memory */

		newtext = zend.ZendStringTruncate(newtext, newtextlen, 0)
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = newtext
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
		return
	}

	/* Special case for a single-character break as it needs no
	   additional storage space */
}

/* }}} */

func PhpExplode(delim *zend.ZendString, str *zend.ZendString, return_value *zend.Zval, limit zend.ZendLong) {
	var p1 *byte = str.val
	var endp *byte = str.val + str.len_
	var p2 *byte = zend.ZendMemnstr(str.val, delim.val, delim.len_, endp)
	var tmp zend.Zval
	if p2 == nil {
		var __z *zend.Zval = &tmp
		var __s *zend.ZendString = str
		__z.value.str = __s
		if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
			__z.u1.type_info = 6
		} else {
			zend.ZendGcAddref(&__s.gc)
			__z.u1.type_info = 6 | 1<<0<<8
		}
		zend.ZendHashNextIndexInsertNew(return_value.value.arr, &tmp)
	} else {
		for {
			var l int = p2 - p1
			if l == 0 {
				var __z *zend.Zval = &tmp
				var __s *zend.ZendString = zend.ZendEmptyString
				__z.value.str = __s
				__z.u1.type_info = 6
			} else if l == 1 {
				var __z *zend.Zval = &tmp
				var __s *zend.ZendString = zend.ZendOneCharString[zend_uchar(*p1)]
				__z.value.str = __s
				__z.u1.type_info = 6
			} else {
				var __z *zend.Zval = &tmp
				var __s *zend.ZendString = zend.ZendStringInit(p1, p2-p1, 0)
				__z.value.str = __s
				__z.u1.type_info = 6 | 1<<0<<8
			}
			zend.ZendHashNextIndexInsertNew(return_value.value.arr, &tmp)
			p1 = p2 + delim.len_
			p2 = zend.ZendMemnstr(p1, delim.val, delim.len_, endp)
			if !(p2 != nil && g.PreDec(&limit) > 1) {
				break
			}
		}
		if p1 <= endp {
			var __z *zend.Zval = &tmp
			var __s *zend.ZendString = zend.ZendStringInit(p1, endp-p1, 0)
			__z.value.str = __s
			__z.u1.type_info = 6 | 1<<0<<8
			zend.ZendHashNextIndexInsertNew(return_value.value.arr, &tmp)
		}
	}
}

/* }}} */

func PhpExplodeNegativeLimit(delim *zend.ZendString, str *zend.ZendString, return_value *zend.Zval, limit zend.ZendLong) {
	// #define EXPLODE_ALLOC_STEP       64

	var p1 *byte = str.val
	var endp *byte = str.val + str.len_
	var p2 *byte = zend.ZendMemnstr(str.val, delim.val, delim.len_, endp)
	var tmp zend.Zval
	if p2 == nil {

	} else {
		var allocated int = 64
		var found int = 0
		var i zend.ZendLong
		var to_return zend.ZendLong
		var positions **byte = zend._emalloc(allocated * g.SizeOf("char *"))
		positions[g.PostInc(&found)] = p1
		for {
			if found >= allocated {
				allocated = found + 64
				positions = zend._erealloc(positions, allocated*g.SizeOf("char *"))
			}
			p1 = p2 + delim.len_
			positions[g.PostInc(&found)] = p1
			p2 = zend.ZendMemnstr(p1, delim.val, delim.len_, endp)
			if p2 == nil {
				break
			}
		}
		to_return = limit + found

		/* limit is at least -1 therefore no need of bounds checking : i will be always less than found */

		for i = 0; i < to_return; i++ {
			var __z *zend.Zval = &tmp
			var __s *zend.ZendString = zend.ZendStringInit(positions[i], positions[i+1]-delim.len_-positions[i], 0)
			__z.value.str = __s
			__z.u1.type_info = 6 | 1<<0<<8
			zend.ZendHashNextIndexInsertNew(return_value.value.arr, &tmp)
		}
		zend._efree(any(positions))
	}
}

/* }}} */

func ZifExplode(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var str *zend.ZendString
	var delim *zend.ZendString
	var limit zend.ZendLong = INT64_MAX
	var tmp zend.Zval
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
			r.Assert(_i <= _min_num_args || _optional == 1)
			r.Assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgStr(_arg, &delim, 0) == 0 {
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

			if zend.ZendParseArgStr(_arg, &str, 0) == 0 {
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

			if zend.ZendParseArgLong(_arg, &limit, &_dummy, 0, 0) == 0 {
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
	if delim.len_ == 0 {
		core.PhpErrorDocref(nil, 1<<1, "Empty delimiter")
		return_value.u1.type_info = 2
		return
	}
	var __arr *zend.ZendArray = zend._zendNewArray(0)
	var __z *zend.Zval = return_value
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	if str.len_ == 0 {
		if limit >= 0 {
			var __z *zend.Zval = &tmp
			var __s *zend.ZendString = zend.ZendEmptyString
			__z.value.str = __s
			__z.u1.type_info = 6
			zend.ZendHashIndexAddNew(return_value.value.arr, 0, &tmp)
		}
		return
	}
	if limit > 1 {
		PhpExplode(delim, str, return_value, limit)
	} else if limit < 0 {
		PhpExplodeNegativeLimit(delim, str, return_value, limit)
	} else {
		var __z *zend.Zval = &tmp
		var __s *zend.ZendString = str
		__z.value.str = __s
		if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
			__z.u1.type_info = 6
		} else {
			zend.ZendGcAddref(&__s.gc)
			__z.u1.type_info = 6 | 1<<0<<8
		}
		zend.ZendHashIndexAddNew(return_value.value.arr, 0, &tmp)
	}
}

/* }}} */

func PhpImplode(glue *zend.ZendString, pieces *zend.Zval, return_value *zend.Zval) {
	var tmp *zend.Zval
	var numelems int
	var str *zend.ZendString
	var cptr *byte
	var len_ int = 0
	var strings *struct {
		str  *zend.ZendString
		lval zend.ZendLong
	}
	var ptr *struct {
		str  *zend.ZendString
		lval zend.ZendLong
	}
	numelems = pieces.value.arr.nNumOfElements
	if numelems == 0 {
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = zend.ZendEmptyString
		__z.value.str = __s
		__z.u1.type_info = 6
		return
	} else if numelems == 1 {

		/* loop to search the first not undefined element... */

		for {
			var __ht *zend.HashTable = pieces.value.arr
			var _p *zend.Bucket = __ht.arData
			var _end *zend.Bucket = _p + __ht.nNumUsed
			for ; _p != _end; _p++ {
				var _z *zend.Zval = &_p.val
				if _z.u1.v.type_ == 13 {
					_z = _z.value.zv
				}
				if _z.u1.v.type_ == 0 {
					continue
				}
				tmp = _z
				var __z *zend.Zval = return_value
				var __s *zend.ZendString = zend.ZvalGetString(tmp)
				__z.value.str = __s
				if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
					__z.u1.type_info = 6
				} else {
					__z.u1.type_info = 6 | 1<<0<<8
				}
				return
			}
			break
		}

		/* loop to search the first not undefined element... */

	}
	strings = zend._emalloc(g.SizeOf("* strings") * numelems)
	ptr = strings
	for {
		var __ht *zend.HashTable = pieces.value.arr
		var _p *zend.Bucket = __ht.arData
		var _end *zend.Bucket = _p + __ht.nNumUsed
		for ; _p != _end; _p++ {
			var _z *zend.Zval = &_p.val
			if _z.u1.v.type_ == 13 {
				_z = _z.value.zv
			}
			if _z.u1.v.type_ == 0 {
				continue
			}
			tmp = _z
			if tmp.u1.v.type_ == 6 {
				ptr.str = tmp.value.str
				len_ += ptr.str.len_
				ptr.lval = 0
				ptr++
			} else if tmp.u1.v.type_ == 4 {
				var val zend.ZendLong = tmp.value.lval
				ptr.str = nil
				ptr.lval = val
				ptr++
				if val <= 0 {
					len_++
				}
				for val != 0 {
					val /= 10
					len_++
				}
			} else {
				ptr.str = zend.ZvalGetStringFunc(tmp)
				len_ += ptr.str.len_
				ptr.lval = 1
				ptr++
			}
		}
		break
	}

	/* numelems can not be 0, we checked above */

	str = zend.ZendStringSafeAlloc(numelems-1, glue.len_, len_, 0)
	cptr = str.val + str.len_
	*cptr = 0
	for true {
		ptr--
		if ptr.str != nil {
			cptr -= ptr.str.len_
			memcpy(cptr, ptr.str.val, ptr.str.len_)
			if ptr.lval != 0 {
				zend.ZendStringReleaseEx(ptr.str, 0)
			}
		} else {
			var oldPtr *byte = cptr
			var oldVal byte = *cptr
			cptr = zend.ZendPrintLongToBuf(cptr, ptr.lval)
			*oldPtr = oldVal
		}
		if ptr == strings {
			break
		}
		cptr -= glue.len_
		memcpy(cptr, glue.val, glue.len_)
	}
	zend._efree(strings)
	var __z *zend.Zval = return_value
	var __s *zend.ZendString = str
	__z.value.str = __s
	__z.u1.type_info = 6 | 1<<0<<8
	return
}

/* }}} */

func ZifImplode(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var arg1 *zend.Zval
	var arg2 *zend.Zval = nil
	var pieces *zend.Zval
	var glue *zend.ZendString
	var tmp_glue *zend.ZendString
	for {
		var _flags int = 0
		var _min_num_args int = 1
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
	if arg2 == nil {
		if arg1.u1.v.type_ != 7 {
			core.PhpErrorDocref(nil, 1<<1, "Argument must be an array")
			return
		}
		glue = zend.ZendEmptyString
		tmp_glue = nil
		pieces = arg1
	} else {
		if arg1.u1.v.type_ == 7 {
			glue = zend.ZvalGetTmpString(arg2, &tmp_glue)
			pieces = arg1
			core.PhpErrorDocref(nil, 1<<13, "Passing glue string after array is deprecated. Swap the parameters")
		} else if arg2.u1.v.type_ == 7 {
			glue = zend.ZvalGetTmpString(arg1, &tmp_glue)
			pieces = arg2
		} else {
			core.PhpErrorDocref(nil, 1<<1, "Invalid arguments passed")
			return
		}
	}
	PhpImplode(glue, pieces, return_value)
	zend.ZendTmpStringRelease(tmp_glue)
}

/* }}} */

// #define STRTOK_TABLE(p) BG ( strtok_table ) [ ( unsigned char ) * p ]

/* {{{ proto string strtok([string str,] string token)
   Tokenize a string */

func ZifStrtok(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var str *zend.ZendString
	var tok *zend.ZendString = nil
	var token *byte
	var token_end *byte
	var p *byte
	var pe *byte
	var skipped int = 0
	for {
		var _flags int = 0
		var _min_num_args int = 1
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
			r.Assert(_i <= _min_num_args || _optional == 1)
			r.Assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgStr(_arg, &str, 0) == 0 {
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

			if zend.ZendParseArgStr(_arg, &tok, 0) == 0 {
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
	if execute_data.This.u2.num_args == 1 {
		tok = str
	} else {
		zend.ZvalPtrDtor(&(BasicGlobals.GetStrtokZval()))
		var __z *zend.Zval = &(BasicGlobals.GetStrtokZval())
		var __s *zend.ZendString = zend.ZendStringInit(str.val, str.len_, 0)
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
		BasicGlobals.SetStrtokString(BasicGlobals.strtok_zval.value.str.val)
		BasicGlobals.SetStrtokLast(BasicGlobals.GetStrtokString())
		BasicGlobals.SetStrtokLen(str.len_)
	}
	p = BasicGlobals.GetStrtokLast()
	pe = BasicGlobals.GetStrtokString() + BasicGlobals.GetStrtokLen()
	if p == nil || p >= pe {
		return_value.u1.type_info = 2
		return
	}
	token = tok.val
	token_end = token + tok.len_
	for token < token_end {
		BasicGlobals.GetStrtokTable()[uint8(g.PostInc(&(*token)))] = 1
	}

	/* Skip leading delimiters */

	for BasicGlobals.GetStrtokTable()[uint8(*p)] {
		if g.PreInc(&p) >= pe {

			/* no other chars left */

			BasicGlobals.SetStrtokLast(nil)
			return_value.u1.type_info = 2
			goto restore
		}
		skipped++
	}

	/* We know at this place that *p is no delimiter, so skip it */

	for g.PreInc(&p) < pe {
		if BasicGlobals.GetStrtokTable()[uint8(*p)] {
			goto return_token
		}
	}
	if p-BasicGlobals.GetStrtokLast() != 0 {
	return_token:
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = zend.ZendStringInit(BasicGlobals.GetStrtokLast()+skipped, p-BasicGlobals.GetStrtokLast()-skipped, 0)
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
		BasicGlobals.SetStrtokLast(p + 1)
	} else {
		return_value.u1.type_info = 2
		BasicGlobals.SetStrtokLast(nil)
	}

	/* Restore table -- usually faster then memset'ing the table on every invocation */

restore:
	token = tok.val
	for token < token_end {
		BasicGlobals.GetStrtokTable()[uint8(g.PostInc(&(*token)))] = 0
	}
}

/* }}} */

func PhpStrtoupper(s *byte, len_ int) *byte {
	var c *uint8
	var e *uint8
	c = (*uint8)(s)
	e = (*uint8)(c + len_)
	for c < e {
		*c = toupper(*c)
		c++
	}
	return s
}

/* }}} */

func PhpStringToupper(s *zend.ZendString) *zend.ZendString {
	var c *uint8
	var e *uint8
	c = (*uint8)(s.val)
	e = c + s.len_
	for c < e {
		if islower(*c) {
			var r *uint8
			var res *zend.ZendString = zend.ZendStringAlloc(s.len_, 0)
			if c != (*uint8)(s.val) {
				memcpy(res.val, s.val, c-(*uint8)(s.val))
			}
			r = c + (res.val - s.val)
			for c < e {
				*r = toupper(*c)
				r++
				c++
			}
			*r = '0'
			return res
		}
		c++
	}
	return zend.ZendStringCopy(s)
}

/* }}} */

func ZifStrtoupper(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var arg *zend.ZendString
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
			r.Assert(_i <= _min_num_args || _optional == 1)
			r.Assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgStr(_arg, &arg, 0) == 0 {
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
	var __z *zend.Zval = return_value
	var __s *zend.ZendString = PhpStringToupper(arg)
	__z.value.str = __s
	if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
		__z.u1.type_info = 6
	} else {
		__z.u1.type_info = 6 | 1<<0<<8
	}
	return
}

/* }}} */

func PhpStrtolower(s *byte, len_ int) *byte {
	var c *uint8
	var e *uint8
	c = (*uint8)(s)
	e = c + len_
	for c < e {
		*c = tolower(*c)
		c++
	}
	return s
}

/* }}} */

func PhpStringTolower(s *zend.ZendString) *zend.ZendString {
	var c *uint8
	var e *uint8
	c = (*uint8)(s.val)
	e = c + s.len_
	for c < e {
		if isupper(*c) {
			var r *uint8
			var res *zend.ZendString = zend.ZendStringAlloc(s.len_, 0)
			if c != (*uint8)(s.val) {
				memcpy(res.val, s.val, c-(*uint8)(s.val))
			}
			r = c + (res.val - s.val)
			for c < e {
				*r = tolower(*c)
				r++
				c++
			}
			*r = '0'
			return res
		}
		c++
	}
	return zend.ZendStringCopy(s)
}

/* }}} */

func ZifStrtolower(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var str *zend.ZendString
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
			r.Assert(_i <= _min_num_args || _optional == 1)
			r.Assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgStr(_arg, &str, 0) == 0 {
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
	var __z *zend.Zval = return_value
	var __s *zend.ZendString = PhpStringTolower(str)
	__z.value.str = __s
	if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
		__z.u1.type_info = 6
	} else {
		__z.u1.type_info = 6 | 1<<0<<8
	}
	return
}

/* }}} */

func PhpBasename(s *byte, len_ int, suffix *byte, sufflen int) *zend.ZendString {
	var c *byte
	var comp *byte
	var cend *byte
	var inc_len int
	var cnt int
	var state int
	var ret *zend.ZendString
	c = (*byte)(s)
	cend = c
	comp = cend
	cnt = len_
	state = 0
	for cnt > 0 {
		if (*c) == '0' {
			inc_len = 1
		} else {
			inc_len = mblen(c, cnt)
		}
		switch inc_len {
		case -2:

		case -1:
			inc_len = 1
			void(mblen(nil, 0))
			break
		case 0:
			goto quit_loop
		case 1:
			if (*c) == '/' {
				if state == 1 {
					state = 0
					cend = c
				}
			} else {
				if state == 0 {
					comp = c
					state = 1
				}
			}
			break
		default:
			if state == 0 {
				comp = c
				state = 1
			}
			break
		}
		c += inc_len
		cnt -= inc_len
	}
quit_loop:
	if state == 1 {
		cend = c
	}
	if suffix != nil && sufflen < size_t(cend-comp) && memcmp(cend-sufflen, suffix, sufflen) == 0 {
		cend -= sufflen
	}
	len_ = cend - comp
	ret = zend.ZendStringInit(comp, len_, 0)
	return ret
}

/* }}} */

func ZifBasename(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var string *byte
	var suffix *byte = nil
	var string_len int
	var suffix_len int = 0
	for {
		var _flags int = 0
		var _min_num_args int = 1
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
			r.Assert(_i <= _min_num_args || _optional == 1)
			r.Assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgString(_arg, &string, &string_len, 0) == 0 {
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

			if zend.ZendParseArgString(_arg, &suffix, &suffix_len, 0) == 0 {
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
	var __z *zend.Zval = return_value
	var __s *zend.ZendString = PhpBasename(string, string_len, suffix, suffix_len)
	__z.value.str = __s
	if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
		__z.u1.type_info = 6
	} else {
		__z.u1.type_info = 6 | 1<<0<<8
	}
	return
}

/* }}} */

func PhpDirname(path *byte, len_ int) int { return zend.ZendDirname(path, len_) }

/* }}} */

func ZifDirname(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var str *byte
	var str_len int
	var ret *zend.ZendString
	var levels zend.ZendLong = 1
	for {
		var _flags int = 0
		var _min_num_args int = 1
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
			r.Assert(_i <= _min_num_args || _optional == 1)
			r.Assert(_i > _min_num_args || _optional == 0)
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

			if zend.ZendParseArgLong(_arg, &levels, &_dummy, 0, 0) == 0 {
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
	ret = zend.ZendStringInit(str, str_len, 0)
	if levels == 1 {

		/* Default case */

		ret.len_ = zend.ZendDirname(ret.val, str_len)

		/* Default case */

	} else if levels < 1 {
		core.PhpErrorDocref(nil, 1<<1, "Invalid argument, levels must be >= 1")
		zend.ZendStringEfree(ret)
		return
	} else {

		/* Some levels up */

		for {
			ret.len_ = zend.ZendDirname(ret.val, g.Assign(&str_len, ret.len_))
			if !(ret.len_ < str_len && g.PreDec(&levels)) {
				break
			}
		}

		/* Some levels up */

	}
	var __z *zend.Zval = return_value
	var __s *zend.ZendString = ret
	__z.value.str = __s
	__z.u1.type_info = 6 | 1<<0<<8
	return
}

/* }}} */

func ZifPathinfo(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var tmp zend.Zval
	var path *byte
	var dirname *byte
	var path_len int
	var have_basename int
	var opt zend.ZendLong = 1 | 2 | 4 | 8
	var ret *zend.ZendString = nil
	for {
		var _flags int = 0
		var _min_num_args int = 1
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
			r.Assert(_i <= _min_num_args || _optional == 1)
			r.Assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgString(_arg, &path, &path_len, 0) == 0 {
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

			if zend.ZendParseArgLong(_arg, &opt, &_dummy, 0, 0) == 0 {
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
	have_basename = (opt & 2) == 2
	var __arr *zend.ZendArray = zend._zendNewArray(0)
	var __z *zend.Zval = &tmp
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	if (opt & 1) == 1 {
		dirname = zend._estrndup(path, path_len)
		PhpDirname(dirname, path_len)
		if *dirname {
			zend.AddAssocStringEx(&tmp, "dirname", strlen("dirname"), dirname)
		}
		zend._efree(dirname)
	}
	if have_basename != 0 {
		ret = PhpBasename(path, path_len, nil, 0)
		zend.AddAssocStrEx(&tmp, "basename", strlen("basename"), zend.ZendStringCopy(ret))
	}
	if (opt & 4) == 4 {
		var p *byte
		var idx ptrdiff_t
		if have_basename == 0 {
			ret = PhpBasename(path, path_len, nil, 0)
		}
		p = zend.ZendMemrchr(ret.val, '.', ret.len_)
		if p != nil {
			idx = p - ret.val
			zend.AddAssocStringlEx(&tmp, "extension", strlen("extension"), ret.val+idx+1, ret.len_-idx-1)
		}
	}
	if (opt & 8) == 8 {
		var p *byte
		var idx ptrdiff_t

		/* Have we already looked up the basename? */

		if have_basename == 0 && ret == nil {
			ret = PhpBasename(path, path_len, nil, 0)
		}
		p = zend.ZendMemrchr(ret.val, '.', ret.len_)
		if p != nil {
			idx = p - ret.val
		} else {
			idx = ptrdiff_t(ret).len_
		}
		zend.AddAssocStringlEx(&tmp, "filename", strlen("filename"), ret.val, idx)
	}
	if ret != nil {
		zend.ZendStringReleaseEx(ret, 0)
	}
	if opt == (1 | 2 | 4 | 8) {
		var _z1 *zend.Zval = return_value
		var _z2 *zend.Zval = &tmp
		var _gc *zend.ZendRefcounted = _z2.value.counted
		var _t uint32 = _z2.u1.type_info
		_z1.value.counted = _gc
		_z1.u1.type_info = _t
	} else {
		var element *zend.Zval
		if g.Assign(&element, zend.ZendHashGetCurrentDataEx(tmp.value.arr, &(tmp.value.arr).nInternalPointer)) != nil {
			var _z3 *zend.Zval = element
			if (_z3.u1.type_info & 0xff00) != 0 {
				if (_z3.u1.type_info & 0xff) == 10 {
					_z3 = &(*_z3).value.ref.val
					if (_z3.u1.type_info & 0xff00) != 0 {
						zend.ZvalAddrefP(_z3)
					}
				} else {
					zend.ZvalAddrefP(_z3)
				}
			}
			var _z1 *zend.Zval = return_value
			var _z2 *zend.Zval = _z3
			var _gc *zend.ZendRefcounted = _z2.value.counted
			var _t uint32 = _z2.u1.type_info
			_z1.value.counted = _gc
			_z1.u1.type_info = _t
		} else {
			var __z *zend.Zval = return_value
			var __s *zend.ZendString = zend.ZendEmptyString
			__z.value.str = __s
			__z.u1.type_info = 6
		}
		zend.ZvalPtrDtor(&tmp)
	}
}

/* }}} */

func PhpStristr(s *byte, t *byte, s_len int, t_len int) *byte {
	PhpStrtolower(s, s_len)
	PhpStrtolower(t, t_len)
	return (*byte)(zend.ZendMemnstr(s, t, t_len, s+s_len))
}

/* }}} */

func PhpStrspn(s1 *byte, s2 *byte, s1_end *byte, s2_end *byte) int {
	var p *byte = s1
	var spanp *byte
	var c byte = *p
cont:
	for spanp = s2; p != s1_end && spanp != s2_end; {
		if g.PostInc(&(*spanp)) == c {
			c = *(g.PreInc(&p))
			goto cont
		}
	}
	return p - s1
}

/* }}} */

func PhpStrcspn(s1 *byte, s2 *byte, s1_end *byte, s2_end *byte) int {
	var p *byte
	var spanp *byte
	var c byte = *s1
	for p = s1; ; {
		spanp = s2
		for {
			if (*spanp) == c || p == s1_end {
				return p - s1
			}
			if g.PostInc(&spanp) >= s2_end-1 {
				break
			}
		}
		c = *(g.PreInc(&p))
	}
}

/* }}} */

func PhpNeedleChar(needle *zend.Zval, target *byte) int {
	switch needle.u1.v.type_ {
	case 4:
		*target = byte(needle.value.lval)
		return zend.SUCCESS
	case 1:

	case 2:
		*target = '0'
		return zend.SUCCESS
	case 3:
		*target = '1'
		return zend.SUCCESS
	case 5:

	case 8:
		*target = byte(zend.ZvalGetLong(needle))
		return zend.SUCCESS
	default:
		core.PhpErrorDocref(nil, 1<<1, "needle is not a string or an integer")
		return zend.FAILURE
	}
}

/* }}} */

func ZifStristr(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var needle *zend.Zval
	var haystack *zend.ZendString
	var found *byte = nil
	var found_offset int
	var haystack_dup *byte
	var needle_char []byte
	var part zend.ZendBool = 0
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
			r.Assert(_i <= _min_num_args || _optional == 1)
			r.Assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgStr(_arg, &haystack, 0) == 0 {
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

			zend.ZendParseArgZvalDeref(_arg, &needle, 0)
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

			if zend.ZendParseArgBool(_arg, &part, &_dummy, 0) == 0 {
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
	haystack_dup = zend._estrndup(haystack.val, haystack.len_)
	if needle.u1.v.type_ == 6 {
		var orig_needle *byte
		if needle.value.str.len_ == 0 {
			core.PhpErrorDocref(nil, 1<<1, "Empty needle")
			zend._efree(haystack_dup)
			return_value.u1.type_info = 2
			return
		}
		orig_needle = zend._estrndup(needle.value.str.val, needle.value.str.len_)
		found = PhpStristr(haystack_dup, orig_needle, haystack.len_, needle.value.str.len_)
		zend._efree(orig_needle)
	} else {
		if PhpNeedleChar(needle, needle_char) != zend.SUCCESS {
			zend._efree(haystack_dup)
			return_value.u1.type_info = 2
			return
		}
		needle_char[1] = 0
		core.PhpErrorDocref(nil, 1<<13, "Non-string needles will be interpreted as strings in the future. "+"Use an explicit chr() call to preserve the current behavior")
		found = PhpStristr(haystack_dup, needle_char, haystack.len_, 1)
	}
	if found != nil {
		found_offset = found - haystack_dup
		if part != 0 {
			var __z *zend.Zval = return_value
			var __s *zend.ZendString = zend.ZendStringInit(haystack.val, found_offset, 0)
			__z.value.str = __s
			__z.u1.type_info = 6 | 1<<0<<8
		} else {
			var __z *zend.Zval = return_value
			var __s *zend.ZendString = zend.ZendStringInit(haystack.val+found_offset, haystack.len_-found_offset, 0)
			__z.value.str = __s
			__z.u1.type_info = 6 | 1<<0<<8
		}
	} else {
		return_value.u1.type_info = 2
	}
	zend._efree(haystack_dup)
}

/* }}} */

func ZifStrstr(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var needle *zend.Zval
	var haystack *zend.ZendString
	var found *byte = nil
	var needle_char []byte
	var found_offset zend.ZendLong
	var part zend.ZendBool = 0
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
			r.Assert(_i <= _min_num_args || _optional == 1)
			r.Assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgStr(_arg, &haystack, 0) == 0 {
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

			zend.ZendParseArgZvalDeref(_arg, &needle, 0)
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

			if zend.ZendParseArgBool(_arg, &part, &_dummy, 0) == 0 {
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
	if needle.u1.v.type_ == 6 {
		if needle.value.str.len_ == 0 {
			core.PhpErrorDocref(nil, 1<<1, "Empty needle")
			return_value.u1.type_info = 2
			return
		}
		found = zend.ZendMemnstr(haystack.val, needle.value.str.val, needle.value.str.len_, haystack.val+haystack.len_)
	} else {
		if PhpNeedleChar(needle, needle_char) != zend.SUCCESS {
			return_value.u1.type_info = 2
			return
		}
		needle_char[1] = 0
		core.PhpErrorDocref(nil, 1<<13, "Non-string needles will be interpreted as strings in the future. "+"Use an explicit chr() call to preserve the current behavior")
		found = zend.ZendMemnstr(haystack.val, needle_char, 1, haystack.val+haystack.len_)
	}
	if found != nil {
		found_offset = found - haystack.val
		if part != 0 {
			var __z *zend.Zval = return_value
			var __s *zend.ZendString = zend.ZendStringInit(haystack.val, found_offset, 0)
			__z.value.str = __s
			__z.u1.type_info = 6 | 1<<0<<8
			return
		} else {
			var __z *zend.Zval = return_value
			var __s *zend.ZendString = zend.ZendStringInit(found, haystack.len_-found_offset, 0)
			__z.value.str = __s
			__z.u1.type_info = 6 | 1<<0<<8
			return
		}
	}
	return_value.u1.type_info = 2
	return
}

/* }}} */

func ZifStrpos(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var needle *zend.Zval
	var haystack *zend.ZendString
	var found *byte = nil
	var needle_char []byte
	var offset zend.ZendLong = 0
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
			r.Assert(_i <= _min_num_args || _optional == 1)
			r.Assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgStr(_arg, &haystack, 0) == 0 {
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

			zend.ZendParseArgZvalDeref(_arg, &needle, 0)
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

			if zend.ZendParseArgLong(_arg, &offset, &_dummy, 0, 0) == 0 {
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
	if offset < 0 {
		offset += zend_long(haystack).len_
	}
	if offset < 0 || int(offset > haystack.len_) != 0 {
		core.PhpErrorDocref(nil, 1<<1, "Offset not contained in string")
		return_value.u1.type_info = 2
		return
	}
	if needle.u1.v.type_ == 6 {
		if needle.value.str.len_ == 0 {
			core.PhpErrorDocref(nil, 1<<1, "Empty needle")
			return_value.u1.type_info = 2
			return
		}
		found = (*byte)(zend.ZendMemnstr(haystack.val+offset, needle.value.str.val, needle.value.str.len_, haystack.val+haystack.len_))
	} else {
		if PhpNeedleChar(needle, needle_char) != zend.SUCCESS {
			return_value.u1.type_info = 2
			return
		}
		needle_char[1] = 0
		core.PhpErrorDocref(nil, 1<<13, "Non-string needles will be interpreted as strings in the future. "+"Use an explicit chr() call to preserve the current behavior")
		found = (*byte)(zend.ZendMemnstr(haystack.val+offset, needle_char, 1, haystack.val+haystack.len_))
	}
	if found != nil {
		var __z *zend.Zval = return_value
		__z.value.lval = found - haystack.val
		__z.u1.type_info = 4
		return
	} else {
		return_value.u1.type_info = 2
		return
	}
}

/* }}} */

func ZifStripos(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var found *byte = nil
	var haystack *zend.ZendString
	var offset zend.ZendLong = 0
	var needle_char []byte
	var needle *zend.Zval
	var needle_dup *zend.ZendString = nil
	var haystack_dup *zend.ZendString
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
			r.Assert(_i <= _min_num_args || _optional == 1)
			r.Assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgStr(_arg, &haystack, 0) == 0 {
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

			zend.ZendParseArgZvalDeref(_arg, &needle, 0)
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

			if zend.ZendParseArgLong(_arg, &offset, &_dummy, 0, 0) == 0 {
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
	if offset < 0 {
		offset += zend_long(haystack).len_
	}
	if offset < 0 || int(offset > haystack.len_) != 0 {
		core.PhpErrorDocref(nil, 1<<1, "Offset not contained in string")
		return_value.u1.type_info = 2
		return
	}
	if haystack.len_ == 0 {
		return_value.u1.type_info = 2
		return
	}
	if needle.u1.v.type_ == 6 {
		if needle.value.str.len_ == 0 || needle.value.str.len_ > haystack.len_ {
			return_value.u1.type_info = 2
			return
		}
		haystack_dup = PhpStringTolower(haystack)
		needle_dup = PhpStringTolower(needle.value.str)
		found = (*byte)(zend.ZendMemnstr(haystack_dup.val+offset, needle_dup.val, needle_dup.len_, haystack_dup.val+haystack.len_))
	} else {
		if PhpNeedleChar(needle, needle_char) != zend.SUCCESS {
			return_value.u1.type_info = 2
			return
		}
		core.PhpErrorDocref(nil, 1<<13, "Non-string needles will be interpreted as strings in the future. "+"Use an explicit chr() call to preserve the current behavior")
		haystack_dup = PhpStringTolower(haystack)
		needle_char[0] = tolower(needle_char[0])
		needle_char[1] = '0'
		found = (*byte)(zend.ZendMemnstr(haystack_dup.val+offset, needle_char, g.SizeOf("needle_char")-1, haystack_dup.val+haystack.len_))
	}
	if found != nil {
		var __z *zend.Zval = return_value
		__z.value.lval = found - haystack_dup.val
		__z.u1.type_info = 4
	} else {
		return_value.u1.type_info = 2
	}
	zend.ZendStringReleaseEx(haystack_dup, 0)
	if needle_dup != nil {
		zend.ZendStringReleaseEx(needle_dup, 0)
	}
}

/* }}} */

func ZifStrrpos(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var zneedle *zend.Zval
	var haystack *zend.ZendString
	var needle_len int
	var offset zend.ZendLong = 0
	var ord_needle []byte
	var p *byte
	var e *byte
	var found *byte
	var needle *byte
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
			r.Assert(_i <= _min_num_args || _optional == 1)
			r.Assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgStr(_arg, &haystack, 0) == 0 {
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

			zend.ZendParseArgZvalDeref(_arg, &zneedle, 0)
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

			if zend.ZendParseArgLong(_arg, &offset, &_dummy, 0, 0) == 0 {
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
	if zneedle.u1.v.type_ == 6 {
		needle = zneedle.value.str.val
		needle_len = zneedle.value.str.len_
	} else {
		if PhpNeedleChar(zneedle, ord_needle) != zend.SUCCESS {
			return_value.u1.type_info = 2
			return
		}
		core.PhpErrorDocref(nil, 1<<13, "Non-string needles will be interpreted as strings in the future. "+"Use an explicit chr() call to preserve the current behavior")
		ord_needle[1] = '0'
		needle = ord_needle
		needle_len = 1
	}
	if haystack.len_ == 0 || needle_len == 0 {
		return_value.u1.type_info = 2
		return
	}
	if offset >= 0 {
		if int(offset > haystack.len_) != 0 {
			core.PhpErrorDocref(nil, 1<<1, "Offset is greater than the length of haystack string")
			return_value.u1.type_info = 2
			return
		}
		p = haystack.val + int(offset)
		e = haystack.val + haystack.len_
	} else {
		if offset < -2147483647 || size_t(-offset) > haystack.len_ {
			core.PhpErrorDocref(nil, 1<<1, "Offset is greater than the length of haystack string")
			return_value.u1.type_info = 2
			return
		}
		p = haystack.val
		if size_t-offset < needle_len {
			e = haystack.val + haystack.len_
		} else {
			e = haystack.val + haystack.len_ + offset + needle_len
		}
	}
	if g.Assign(&found, zend.ZendMemnrstr(p, needle, needle_len, e)) {
		var __z *zend.Zval = return_value
		__z.value.lval = found - haystack.val
		__z.u1.type_info = 4
		return
	}
	return_value.u1.type_info = 2
	return
}

/* }}} */

func ZifStrripos(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var zneedle *zend.Zval
	var needle *zend.ZendString
	var haystack *zend.ZendString
	var offset zend.ZendLong = 0
	var p *byte
	var e *byte
	var found *byte
	var needle_dup *zend.ZendString
	var haystack_dup *zend.ZendString
	var ord_needle *zend.ZendString = nil
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
			r.Assert(_i <= _min_num_args || _optional == 1)
			r.Assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgStr(_arg, &haystack, 0) == 0 {
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

			zend.ZendParseArgZvalDeref(_arg, &zneedle, 0)
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

			if zend.ZendParseArgLong(_arg, &offset, &_dummy, 0, 0) == 0 {
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
	ord_needle = (*zend.ZendString)(zend._emalloc(zend_long((*byte)(&((*zend.ZendString)(nil).val))-(*byte)(nil)) + 1 + 1 + (8-1) & ^(8-1)))
	zend.ZendGcSetRefcount(&ord_needle.gc, 1)
	ord_needle.gc.u.type_info = 6
	ord_needle.h = 0
	ord_needle.len_ = 1
	if zneedle.u1.v.type_ == 6 {
		needle = zneedle.value.str
	} else {
		if PhpNeedleChar(zneedle, ord_needle.val) != zend.SUCCESS {
			zend._efree(ord_needle)
			return_value.u1.type_info = 2
			return
		}
		core.PhpErrorDocref(nil, 1<<13, "Non-string needles will be interpreted as strings in the future. "+"Use an explicit chr() call to preserve the current behavior")
		ord_needle.val[1] = '0'
		needle = ord_needle
	}
	if haystack.len_ == 0 || needle.len_ == 0 {
		zend._efree(ord_needle)
		return_value.u1.type_info = 2
		return
	}
	if needle.len_ == 1 {

		/* Single character search can shortcut memcmps
		   Can also avoid tolower emallocs */

		if offset >= 0 {
			if int(offset > haystack.len_) != 0 {
				zend._efree(ord_needle)
				core.PhpErrorDocref(nil, 1<<1, "Offset is greater than the length of haystack string")
				return_value.u1.type_info = 2
				return
			}
			p = haystack.val + int(offset)
			e = haystack.val + haystack.len_ - 1
		} else {
			p = haystack.val
			if offset < -2147483647 || size_t(-offset) > haystack.len_ {
				zend._efree(ord_needle)
				core.PhpErrorDocref(nil, 1<<1, "Offset is greater than the length of haystack string")
				return_value.u1.type_info = 2
				return
			}
			e = haystack.val + (haystack.len_ + int(offset))
		}

		/* Borrow that ord_needle buffer to avoid repeatedly tolower()ing needle */

		(*ord_needle).val = tolower((*needle).val)
		for e >= p {
			if tolower(*e) == (*ord_needle).val {
				zend._efree(ord_needle)
				var __z *zend.Zval = return_value
				__z.value.lval = e - p + g.Cond(offset > 0, offset, 0)
				__z.u1.type_info = 4
				return
			}
			e--
		}
		zend._efree(ord_needle)
		return_value.u1.type_info = 2
		return
	}
	haystack_dup = PhpStringTolower(haystack)
	if offset >= 0 {
		if int(offset > haystack.len_) != 0 {
			zend.ZendStringReleaseEx(haystack_dup, 0)
			zend._efree(ord_needle)
			core.PhpErrorDocref(nil, 1<<1, "Offset is greater than the length of haystack string")
			return_value.u1.type_info = 2
			return
		}
		p = haystack_dup.val + offset
		e = haystack_dup.val + haystack.len_
	} else {
		if offset < -2147483647 || size_t(-offset) > haystack.len_ {
			zend.ZendStringReleaseEx(haystack_dup, 0)
			zend._efree(ord_needle)
			core.PhpErrorDocref(nil, 1<<1, "Offset is greater than the length of haystack string")
			return_value.u1.type_info = 2
			return
		}
		p = haystack_dup.val
		if size_t-offset < needle.len_ {
			e = haystack_dup.val + haystack.len_
		} else {
			e = haystack_dup.val + haystack.len_ + offset + needle.len_
		}
	}
	needle_dup = PhpStringTolower(needle)
	if g.Assign(&found, (*byte)(zend.ZendMemnrstr(p, needle_dup.val, needle_dup.len_, e))) {
		var __z *zend.Zval = return_value
		__z.value.lval = found - haystack_dup.val
		__z.u1.type_info = 4
		zend.ZendStringReleaseEx(needle_dup, 0)
		zend.ZendStringReleaseEx(haystack_dup, 0)
		zend._efree(ord_needle)
	} else {
		zend.ZendStringReleaseEx(needle_dup, 0)
		zend.ZendStringReleaseEx(haystack_dup, 0)
		zend._efree(ord_needle)
		return_value.u1.type_info = 2
		return
	}
}

/* }}} */

func ZifStrrchr(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var needle *zend.Zval
	var haystack *zend.ZendString
	var found *byte = nil
	var found_offset zend.ZendLong
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
			r.Assert(_i <= _min_num_args || _optional == 1)
			r.Assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgStr(_arg, &haystack, 0) == 0 {
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

			zend.ZendParseArgZvalDeref(_arg, &needle, 0)
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
	if needle.u1.v.type_ == 6 {
		found = zend.ZendMemrchr(haystack.val, (*(needle.value.str)).val, haystack.len_)
	} else {
		var needle_chr byte
		if PhpNeedleChar(needle, &needle_chr) != zend.SUCCESS {
			return_value.u1.type_info = 2
			return
		}
		core.PhpErrorDocref(nil, 1<<13, "Non-string needles will be interpreted as strings in the future. "+"Use an explicit chr() call to preserve the current behavior")
		found = zend.ZendMemrchr(haystack.val, needle_chr, haystack.len_)
	}
	if found != nil {
		found_offset = found - haystack.val
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = zend.ZendStringInit(found, haystack.len_-found_offset, 0)
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
		return
	} else {
		return_value.u1.type_info = 2
		return
	}
}

/* }}} */

func PhpChunkSplit(src *byte, srclen int, end *byte, endlen int, chunklen int) *zend.ZendString {
	var q *byte
	var p *byte
	var chunks int
	var restlen int
	var out_len int
	var dest *zend.ZendString
	chunks = srclen / chunklen
	restlen = srclen - chunks*chunklen
	if chunks > 2147483647-1 {
		return nil
	}
	out_len = chunks + 1
	if endlen != 0 && out_len > 2147483647/endlen {
		return nil
	}
	out_len *= endlen
	if out_len > 2147483647-srclen-1 {
		return nil
	}
	out_len += srclen + 1
	dest = zend.ZendStringAlloc(out_len*g.SizeOf("char"), 0)
	p = src
	q = dest.val
	for p < src+srclen-chunklen+1 {
		memcpy(q, p, chunklen)
		q += chunklen
		memcpy(q, end, endlen)
		q += endlen
		p += chunklen
	}
	if restlen != 0 {
		memcpy(q, p, restlen)
		q += restlen
		memcpy(q, end, endlen)
		q += endlen
	}
	*q = '0'
	dest.len_ = q - dest.val
	return dest
}

/* }}} */

func ZifChunkSplit(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var str *zend.ZendString
	var end *byte = "\r\n"
	var endlen int = 2
	var chunklen zend.ZendLong = 76
	var result *zend.ZendString
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

			if zend.ZendParseArgStr(_arg, &str, 0) == 0 {
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

			if zend.ZendParseArgLong(_arg, &chunklen, &_dummy, 0, 0) == 0 {
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

			if zend.ZendParseArgString(_arg, &end, &endlen, 0) == 0 {
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
	if chunklen <= 0 {
		core.PhpErrorDocref(nil, 1<<1, "Chunk length should be greater than zero")
		return_value.u1.type_info = 2
		return
	}
	if int(chunklen > str.len_) != 0 {

		/* to maintain BC, we must return original string + ending */

		result = zend.ZendStringSafeAlloc(str.len_, 1, endlen, 0)
		memcpy(result.val, str.val, str.len_)
		memcpy(result.val+str.len_, end, endlen)
		result.val[result.len_] = '0'
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = result
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
		return
	}
	if str.len_ == 0 {
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = zend.ZendEmptyString
		__z.value.str = __s
		__z.u1.type_info = 6
		return
	}
	result = PhpChunkSplit(str.val, str.len_, end, endlen, int(chunklen))
	if result != nil {
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = result
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

func ZifSubstr(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var str *zend.ZendString
	var l zend.ZendLong = 0
	var f zend.ZendLong
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
			r.Assert(_i <= _min_num_args || _optional == 1)
			r.Assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgStr(_arg, &str, 0) == 0 {
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

			if zend.ZendParseArgLong(_arg, &f, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
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

			if zend.ZendParseArgLong(_arg, &l, &_dummy, 0, 0) == 0 {
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
	if f > zend_long(str).len_ {
		return_value.u1.type_info = 2
		return
	} else if f < 0 {

		/* if "from" position is negative, count start position from the end
		 * of the string
		 */

		if size_t-f > str.len_ {
			f = 0
		} else {
			f = zend_long(str).len_ + f
		}
		if argc > 2 {
			if l < 0 {

				/* if "length" position is negative, set it to the length
				 * needed to stop that many chars from the end of the string
				 */

				if size_t(-l) > str.len_-int(f) {
					if size_t(-l) > str.len_ {
						return_value.u1.type_info = 2
						return
					} else {
						l = 0
					}
				} else {
					l = zend_long(str).len_ - f + l
				}

				/* if "length" position is negative, set it to the length
				 * needed to stop that many chars from the end of the string
				 */

			} else if int(l > str.len_-int(f)) != 0 {
				goto truncate_len
			}
		} else {
			goto truncate_len
		}
	} else if argc > 2 {
		if l < 0 {

			/* if "length" position is negative, set it to the length
			 * needed to stop that many chars from the end of the string
			 */

			if size_t(-l) > str.len_-int(f) {
				return_value.u1.type_info = 2
				return
			} else {
				l = zend_long(str).len_ - f + l
			}

			/* if "length" position is negative, set it to the length
			 * needed to stop that many chars from the end of the string
			 */

		} else if int(l > str.len_-int(f)) != 0 {
			goto truncate_len
		}
	} else {
	truncate_len:
		l = zend_long(str).len_ - f
	}
	if l == 0 {
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = zend.ZendEmptyString
		__z.value.str = __s
		__z.u1.type_info = 6
		return
	} else if l == 1 {
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = zend.ZendOneCharString[zend_uchar(str.val[f])]
		__z.value.str = __s
		__z.u1.type_info = 6
		return
	} else if l == str.len_ {
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = str
		__z.value.str = __s
		if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
			__z.u1.type_info = 6
		} else {
			zend.ZendGcAddref(&__s.gc)
			__z.u1.type_info = 6 | 1<<0<<8
		}
		return
	}
	var __z *zend.Zval = return_value
	var __s *zend.ZendString = zend.ZendStringInit(str.val+f, l, 0)
	__z.value.str = __s
	__z.u1.type_info = 6 | 1<<0<<8
	return
}

/* }}} */

func ZifSubstrReplace(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var str *zend.Zval
	var from *zend.Zval
	var len_ *zend.Zval = nil
	var repl *zend.Zval
	var l zend.ZendLong = 0
	var f zend.ZendLong
	var argc int = execute_data.This.u2.num_args
	var result *zend.ZendString
	var from_idx zend.HashPosition
	var repl_idx zend.HashPosition
	var len_idx zend.HashPosition
	var tmp_str *zend.Zval = nil
	var tmp_from *zend.Zval = nil
	var tmp_repl *zend.Zval = nil
	var tmp_len *zend.Zval = nil
	for {
		var _flags int = 0
		var _min_num_args int = 3
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
			r.Assert(_i <= _min_num_args || _optional == 1)
			r.Assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			zend.ZendParseArgZvalDeref(_arg, &str, 0)
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

			zend.ZendParseArgZvalDeref(_arg, &repl, 0)
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

			zend.ZendParseArgZvalDeref(_arg, &from, 0)
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

			zend.ZendParseArgZvalDeref(_arg, &len_, 0)
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
	if str.u1.v.type_ != 7 {
		if str.u1.v.type_ != 6 {
			if str.u1.v.type_ != 6 {
				zend._convertToString(str)
			}
		}
	}
	if repl.u1.v.type_ != 7 {
		if repl.u1.v.type_ != 6 {
			if repl.u1.v.type_ != 6 {
				zend._convertToString(repl)
			}
		}
	}
	if from.u1.v.type_ != 7 {
		if from.u1.v.type_ != 4 {
			zend.ConvertToLong(from)
		}
	}
	if zend.EG.exception != nil {
		return
	}
	if argc > 3 {
		if len_.u1.v.type_ != 7 {
			if len_.u1.v.type_ != 4 {
				zend.ConvertToLong(len_)
			}
			l = len_.value.lval
		}
	} else {
		if str.u1.v.type_ != 7 {
			l = str.value.str.len_
		}
	}
	if str.u1.v.type_ == 6 {
		if argc == 3 && from.u1.v.type_ == 7 || argc == 4 && from.u1.v.type_ != len_.u1.v.type_ {
			core.PhpErrorDocref(nil, 1<<1, "'start' and 'length' should be of same type - numerical or array ")
			var __z *zend.Zval = return_value
			var __s *zend.ZendString = str.value.str
			__z.value.str = __s
			if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
				__z.u1.type_info = 6
			} else {
				zend.ZendGcAddref(&__s.gc)
				__z.u1.type_info = 6 | 1<<0<<8
			}
			return
		}
		if argc == 4 && from.u1.v.type_ == 7 {
			if from.value.arr.nNumOfElements != len_.value.arr.nNumOfElements {
				core.PhpErrorDocref(nil, 1<<1, "'start' and 'length' should have the same number of elements")
				var __z *zend.Zval = return_value
				var __s *zend.ZendString = str.value.str
				__z.value.str = __s
				if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
					__z.u1.type_info = 6
				} else {
					zend.ZendGcAddref(&__s.gc)
					__z.u1.type_info = 6 | 1<<0<<8
				}
				return
			}
		}
	}
	if str.u1.v.type_ != 7 {
		if from.u1.v.type_ != 7 {
			var repl_str *zend.ZendString
			var tmp_repl_str *zend.ZendString = nil
			f = from.value.lval

			/* if "from" position is negative, count start position from the end
			 * of the string
			 */

			if f < 0 {
				f = zend_long(str.value.str).len_ + f
				if f < 0 {
					f = 0
				}
			} else if int(f > str.value.str.len_) != 0 {
				f = str.value.str.len_
			}

			/* if "length" position is negative, set it to the length
			 * needed to stop that many chars from the end of the string
			 */

			if l < 0 {
				l = zend_long(str.value.str).len_ - f + l
				if l < 0 {
					l = 0
				}
			}
			if int(l > str.value.str.len_ || l < 0 && size_t(-l) > str.value.str.len_) != 0 {
				l = str.value.str.len_
			}
			if f+l > zend_long(str.value.str).len_ {
				l = str.value.str.len_ - f
			}
			if repl.u1.v.type_ == 7 {
				repl_idx = 0
				for repl_idx < repl.value.arr.nNumUsed {
					tmp_repl = &(*repl).value.arr.arData[repl_idx].val
					if tmp_repl.u1.v.type_ != 0 {
						break
					}
					repl_idx++
				}
				if repl_idx < repl.value.arr.nNumUsed {
					repl_str = zend.ZvalGetTmpString(tmp_repl, &tmp_repl_str)
				} else {
					repl_str = zend.ZendEmptyString
				}
			} else {
				repl_str = repl.value.str
			}
			result = zend.ZendStringSafeAlloc(1, str.value.str.len_-l+repl_str.len_, 0, 0)
			memcpy(result.val, str.value.str.val, f)
			if repl_str.len_ != 0 {
				memcpy(result.val+f, repl_str.val, repl_str.len_)
			}
			memcpy(result.val+f+repl_str.len_, str.value.str.val+f+l, str.value.str.len_-f-l)
			result.val[result.len_] = '0'
			zend.ZendTmpStringRelease(tmp_repl_str)
			var __z *zend.Zval = return_value
			var __s *zend.ZendString = result
			__z.value.str = __s
			__z.u1.type_info = 6 | 1<<0<<8
			return
		} else {
			core.PhpErrorDocref(nil, 1<<1, "Functionality of 'start' and 'length' as arrays is not implemented")
			var __z *zend.Zval = return_value
			var __s *zend.ZendString = str.value.str
			__z.value.str = __s
			if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
				__z.u1.type_info = 6
			} else {
				zend.ZendGcAddref(&__s.gc)
				__z.u1.type_info = 6 | 1<<0<<8
			}
			return
		}
	} else {
		var str_index *zend.ZendString = nil
		var result_len int
		var num_index zend.ZendUlong
		var __arr *zend.ZendArray = zend._zendNewArray(0)
		var __z *zend.Zval = return_value
		__z.value.arr = __arr
		__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
		repl_idx = 0
		len_idx = repl_idx
		from_idx = len_idx
		for {
			var __ht *zend.HashTable = str.value.arr
			var _p *zend.Bucket = __ht.arData
			var _end *zend.Bucket = _p + __ht.nNumUsed
			for ; _p != _end; _p++ {
				var _z *zend.Zval = &_p.val
				if _z.u1.v.type_ == 13 {
					_z = _z.value.zv
				}
				if _z.u1.v.type_ == 0 {
					continue
				}
				num_index = _p.h
				str_index = _p.key
				tmp_str = _z
				var tmp_orig_str *zend.ZendString
				var orig_str *zend.ZendString = zend.ZvalGetTmpString(tmp_str, &tmp_orig_str)
				if from.u1.v.type_ == 7 {
					for from_idx < from.value.arr.nNumUsed {
						tmp_from = &(*from).value.arr.arData[from_idx].val
						if tmp_from.u1.v.type_ != 0 {
							break
						}
						from_idx++
					}
					if from_idx < from.value.arr.nNumUsed {
						f = zend.ZvalGetLong(tmp_from)
						if f < 0 {
							f = zend_long(orig_str).len_ + f
							if f < 0 {
								f = 0
							}
						} else if f > zend_long(orig_str).len_ {
							f = orig_str.len_
						}
						from_idx++
					} else {
						f = 0
					}
				} else {
					f = from.value.lval
					if f < 0 {
						f = zend_long(orig_str).len_ + f
						if f < 0 {
							f = 0
						}
					} else if f > zend_long(orig_str).len_ {
						f = orig_str.len_
					}
				}
				if argc > 3 && len_.u1.v.type_ == 7 {
					for len_idx < len_.value.arr.nNumUsed {
						tmp_len = &(*len_).value.arr.arData[len_idx].val
						if tmp_len.u1.v.type_ != 0 {
							break
						}
						len_idx++
					}
					if len_idx < len_.value.arr.nNumUsed {
						l = zend.ZvalGetLong(tmp_len)
						len_idx++
					} else {
						l = orig_str.len_
					}
				} else if argc > 3 {
					l = len_.value.lval
				} else {
					l = orig_str.len_
				}
				if l < 0 {
					l = orig_str.len_ - f + l
					if l < 0 {
						l = 0
					}
				}
				r.Assert(0 <= f && f <= INT64_MAX)
				r.Assert(0 <= l && l <= INT64_MAX)
				if int(f+l) > orig_str.len_ {
					l = orig_str.len_ - f
				}
				result_len = orig_str.len_ - l
				if repl.u1.v.type_ == 7 {
					for repl_idx < repl.value.arr.nNumUsed {
						tmp_repl = &(*repl).value.arr.arData[repl_idx].val
						if tmp_repl.u1.v.type_ != 0 {
							break
						}
						repl_idx++
					}
					if repl_idx < repl.value.arr.nNumUsed {
						var tmp_repl_str *zend.ZendString
						var repl_str *zend.ZendString = zend.ZvalGetTmpString(tmp_repl, &tmp_repl_str)
						result_len += repl_str.len_
						repl_idx++
						result = zend.ZendStringSafeAlloc(1, result_len, 0, 0)
						memcpy(result.val, orig_str.val, f)
						memcpy(result.val+f, repl_str.val, repl_str.len_)
						memcpy(result.val+f+repl_str.len_, orig_str.val+f+l, orig_str.len_-f-l)
						zend.ZendTmpStringRelease(tmp_repl_str)
					} else {
						result = zend.ZendStringSafeAlloc(1, result_len, 0, 0)
						memcpy(result.val, orig_str.val, f)
						memcpy(result.val+f, orig_str.val+f+l, orig_str.len_-f-l)
					}
				} else {
					result_len += repl.value.str.len_
					result = zend.ZendStringSafeAlloc(1, result_len, 0, 0)
					memcpy(result.val, orig_str.val, f)
					memcpy(result.val+f, repl.value.str.val, repl.value.str.len_)
					memcpy(result.val+f+repl.value.str.len_, orig_str.val+f+l, orig_str.len_-f-l)
				}
				result.val[result.len_] = '0'
				if str_index != nil {
					var tmp zend.Zval
					var __z *zend.Zval = &tmp
					var __s *zend.ZendString = result
					__z.value.str = __s
					__z.u1.type_info = 6 | 1<<0<<8
					zend.ZendSymtableUpdate(return_value.value.arr, str_index, &tmp)
				} else {
					zend.AddIndexStr(return_value, num_index, result)
				}
				zend.ZendTmpStringRelease(tmp_orig_str)
			}
			break
		}
	}
}

/* }}} */

func ZifQuotemeta(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var old *zend.ZendString
	var old_end *byte
	var p *byte
	var q *byte
	var c byte
	var str *zend.ZendString
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
			r.Assert(_i <= _min_num_args || _optional == 1)
			r.Assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgStr(_arg, &old, 0) == 0 {
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
	old_end = old.val + old.len_
	if old.val == old_end {
		return_value.u1.type_info = 2
		return
	}
	str = zend.ZendStringSafeAlloc(2, old.len_, 0, 0)
	p = old.val
	q = str.val
	for ; p != old_end; p++ {
		c = *p
		switch c {
		case '.':

		case '\\':

		case '+':

		case '*':

		case '?':

		case '[':

		case '^':

		case ']':

		case '$':

		case '(':

		case ')':
			g.PostInc(&(*q)) = '\\'
		default:
			g.PostInc(&(*q)) = c
		}
	}
	*q = '0'
	var __z *zend.Zval = return_value
	var __s *zend.ZendString = zend.ZendStringTruncate(str, q-str.val, 0)
	__z.value.str = __s
	__z.u1.type_info = 6 | 1<<0<<8
	return
}

/* }}} */

func ZifOrd(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var str *zend.ZendString
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
			r.Assert(_i <= _min_num_args || _optional == 1)
			r.Assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgStr(_arg, &str, 0) == 0 {
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
	var __z *zend.Zval = return_value
	__z.value.lval = uint8(str.val[0])
	__z.u1.type_info = 4
	return
}

/* }}} */

func ZifChr(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var c zend.ZendLong
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
			r.Assert(_i <= _min_num_args || _optional == 1)
			r.Assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgLong(_arg, &c, &_dummy, 0, 0) == 0 {
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
			c = 0
		}
		break
	}
	c &= 0xff
	var __z *zend.Zval = return_value
	var __s *zend.ZendString = zend.ZendOneCharString[c]
	__z.value.str = __s
	__z.u1.type_info = 6
}

/* }}} */

func PhpUcfirst(str *zend.ZendString) *zend.ZendString {
	var ch uint8 = str.val[0]
	var r uint8 = toupper(ch)
	if r == ch {
		return zend.ZendStringCopy(str)
	} else {
		var s *zend.ZendString = zend.ZendStringInit(str.val, str.len_, 0)
		s.val[0] = r
		return s
	}
}

/* }}} */

func ZifUcfirst(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var str *zend.ZendString
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
			r.Assert(_i <= _min_num_args || _optional == 1)
			r.Assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgStr(_arg, &str, 0) == 0 {
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
	if str.len_ == 0 {
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = zend.ZendEmptyString
		__z.value.str = __s
		__z.u1.type_info = 6
		return
	}
	var __z *zend.Zval = return_value
	var __s *zend.ZendString = PhpUcfirst(str)
	__z.value.str = __s
	if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
		__z.u1.type_info = 6
	} else {
		__z.u1.type_info = 6 | 1<<0<<8
	}
	return
}

/* }}} */

func PhpLcfirst(str *zend.ZendString) *zend.ZendString {
	var r uint8 = tolower(str.val[0])
	if r == str.val[0] {
		return zend.ZendStringCopy(str)
	} else {
		var s *zend.ZendString = zend.ZendStringInit(str.val, str.len_, 0)
		s.val[0] = r
		return s
	}
}

/* }}} */

func ZifLcfirst(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var str *zend.ZendString
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
			r.Assert(_i <= _min_num_args || _optional == 1)
			r.Assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgStr(_arg, &str, 0) == 0 {
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
	if str.len_ == 0 {
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = zend.ZendEmptyString
		__z.value.str = __s
		__z.u1.type_info = 6
		return
	}
	var __z *zend.Zval = return_value
	var __s *zend.ZendString = PhpLcfirst(str)
	__z.value.str = __s
	if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
		__z.u1.type_info = 6
	} else {
		__z.u1.type_info = 6 | 1<<0<<8
	}
	return
}

/* }}} */

func ZifUcwords(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var str *zend.ZendString
	var delims *byte = " \t\r\nfv"
	var r *byte
	var r_end *byte
	var delims_len int = 6
	var mask []byte
	for {
		var _flags int = 0
		var _min_num_args int = 1
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
			r.Assert(_i <= _min_num_args || _optional == 1)
			r.Assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgStr(_arg, &str, 0) == 0 {
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

			if zend.ZendParseArgString(_arg, &delims, &delims_len, 0) == 0 {
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
	if str.len_ == 0 {
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = zend.ZendEmptyString
		__z.value.str = __s
		__z.u1.type_info = 6
		return
	}
	PhpCharmask((*uint8)(delims), delims_len, mask)
	var __z *zend.Zval = return_value
	var __s *zend.ZendString = zend.ZendStringInit(str.val, str.len_, 0)
	__z.value.str = __s
	__z.u1.type_info = 6 | 1<<0<<8
	r = return_value.value.str.val
	*r = toupper(uint8(*r))
	for r_end = r + return_value.value.str.len_ - 1; r < r_end; {
		if mask[uint8(g.PostInc(&(*r)))] {
			*r = toupper(uint8(*r))
		}
	}
}

/* }}} */

func PhpStrtr(str *byte, len_ int, str_from *byte, str_to *byte, trlen int) *byte {
	var i int
	if trlen < 1 {
		return str
	} else if trlen == 1 {
		var ch_from byte = *str_from
		var ch_to byte = *str_to
		for i = 0; i < len_; i++ {
			if str[i] == ch_from {
				str[i] = ch_to
			}
		}
	} else {
		var xlat []uint8
		var j uint8 = 0
		for {
			xlat[j] = j
			if g.PreInc(&j) == 0 {
				break
			}
		}
		for i = 0; i < trlen; i++ {
			xlat[int(uint8(str_from[i]))] = str_to[i]
		}
		for i = 0; i < len_; i++ {
			str[i] = xlat[int(uint8(str[i]))]
		}
	}
	return str
}

/* }}} */

func PhpStrtrEx(str *zend.ZendString, str_from *byte, str_to *byte, trlen int) *zend.ZendString {
	var new_str *zend.ZendString = nil
	var i int
	if trlen < 1 {
		return zend.ZendStringCopy(str)
	} else if trlen == 1 {
		var ch_from byte = *str_from
		var ch_to byte = *str_to
		for i = 0; i < str.len_; i++ {
			if str.val[i] == ch_from {
				new_str = zend.ZendStringAlloc(str.len_, 0)
				memcpy(new_str.val, str.val, i)
				new_str.val[i] = ch_to
				break
			}
		}
		for ; i < str.len_; i++ {
			if str.val[i] != ch_from {
				new_str.val[i] = str.val[i]
			} else {
				new_str.val[i] = ch_to
			}
		}
	} else {
		var xlat []uint8
		var j uint8 = 0
		for {
			xlat[j] = j
			if g.PreInc(&j) == 0 {
				break
			}
		}
		for i = 0; i < trlen; i++ {
			xlat[int(uint8(str_from[i]))] = str_to[i]
		}
		for i = 0; i < str.len_; i++ {
			if str.val[i] != xlat[int(uint8(str.val[i]))] {
				new_str = zend.ZendStringAlloc(str.len_, 0)
				memcpy(new_str.val, str.val, i)
				new_str.val[i] = xlat[int(uint8(str.val[i]))]
				break
			}
		}
		for ; i < str.len_; i++ {
			new_str.val[i] = xlat[int(uint8(str.val[i]))]
		}
	}
	if new_str == nil {
		return zend.ZendStringCopy(str)
	}
	new_str.val[new_str.len_] = 0
	return new_str
}

/* }}} */

func PhpStrtrArray(return_value *zend.Zval, input *zend.ZendString, pats *zend.HashTable) {
	var str *byte = input.val
	var slen int = input.len_
	var num_key zend.ZendUlong
	var str_key *zend.ZendString
	var len_ int
	var pos int
	var old_pos int
	var num_keys int = 0
	var minlen int = 128 * 1024
	var maxlen int = 0
	var str_hash zend.HashTable
	var entry *zend.Zval
	var key *byte
	var result zend.SmartStr = zend.SmartStr{0}
	var bitset []zend.ZendUlong
	var num_bitset *zend.ZendUlong

	/* we will collect all possible key lengths */

	num_bitset = zend._ecalloc((slen+g.SizeOf("zend_ulong"))/g.SizeOf("zend_ulong"), g.SizeOf("zend_ulong"))
	memset(bitset, 0, g.SizeOf("bitset"))

	/* check if original array has numeric keys */

	for {
		var __ht *zend.HashTable = pats
		var _p *zend.Bucket = __ht.arData
		var _end *zend.Bucket = _p + __ht.nNumUsed
		for ; _p != _end; _p++ {
			var _z *zend.Zval = &_p.val

			if _z.u1.v.type_ == 0 {
				continue
			}
			str_key = _p.key
			if str_key == nil {
				num_keys = 1
			} else {
				len_ = str_key.len_
				if len_ < 1 {
					zend._efree(num_bitset)
					return_value.u1.type_info = 2
					return
				} else if len_ > slen {

					/* skip long patterns */

					continue

					/* skip long patterns */

				}
				if len_ > maxlen {
					maxlen = len_
				}
				if len_ < minlen {
					minlen = len_
				}

				/* remember possible key length */

				num_bitset[len_/g.SizeOf("zend_ulong")] |= 1 << len_ % g.SizeOf("zend_ulong")
				bitset[uint8(str_key.val[0])/g.SizeOf("zend_ulong")] |= 1 << uint8(str_key.val[0]) % g.SizeOf("zend_ulong")
			}
		}
		break
	}
	if num_keys != 0 {
		var key_used *zend.ZendString

		/* we have to rebuild HashTable with numeric keys */

		zend._zendHashInit(&str_hash, pats.nNumOfElements, nil, 0)
		for {
			var __ht *zend.HashTable = pats
			var _p *zend.Bucket = __ht.arData
			var _end *zend.Bucket = _p + __ht.nNumUsed
			for ; _p != _end; _p++ {
				var _z *zend.Zval = &_p.val
				if _z.u1.v.type_ == 13 {
					_z = _z.value.zv
				}
				if _z.u1.v.type_ == 0 {
					continue
				}
				num_key = _p.h
				str_key = _p.key
				entry = _z
				if str_key == nil {
					key_used = zend.ZendLongToStr(num_key)
					len_ = key_used.len_
					if len_ > slen {

						/* skip long patterns */

						zend.ZendStringRelease(key_used)
						continue
					}
					if len_ > maxlen {
						maxlen = len_
					}
					if len_ < minlen {
						minlen = len_
					}

					/* remember possible key length */

					num_bitset[len_/g.SizeOf("zend_ulong")] |= 1 << len_ % g.SizeOf("zend_ulong")
					bitset[uint8(key_used.val[0])/g.SizeOf("zend_ulong")] |= 1 << uint8(key_used.val[0]) % g.SizeOf("zend_ulong")
				} else {
					key_used = str_key
					len_ = key_used.len_
					if len_ > slen {

						/* skip long patterns */

						continue

						/* skip long patterns */

					}
				}
				zend.ZendHashAdd(&str_hash, key_used, entry)
				if str_key == nil {
					zend.ZendStringReleaseEx(key_used, 0)
				}
			}
			break
		}
		pats = &str_hash
	}
	if minlen > maxlen {

		/* return the original string */

		if pats == &str_hash {
			zend.ZendHashDestroy(&str_hash)
		}
		zend._efree(num_bitset)
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = input
		__z.value.str = __s
		if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
			__z.u1.type_info = 6
		} else {
			zend.ZendGcAddref(&__s.gc)
			__z.u1.type_info = 6 | 1<<0<<8
		}
		return
	}
	pos = 0
	old_pos = pos
	for pos <= slen-minlen {
		key = str + pos
		if (bitset[uint8(key[0])/g.SizeOf("zend_ulong")] & 1 << uint8(key[0]) % g.SizeOf("zend_ulong")) != 0 {
			len_ = maxlen
			if len_ > slen-pos {
				len_ = slen - pos
			}
			for len_ >= minlen {
				if (num_bitset[len_/g.SizeOf("zend_ulong")] & 1 << len_ % g.SizeOf("zend_ulong")) != 0 {
					entry = zend.ZendHashStrFind(pats, key, len_)
					if entry != nil {
						var tmp *zend.ZendString
						var s *zend.ZendString = zend.ZvalGetTmpString(entry, &tmp)
						zend.SmartStrAppendlEx(&result, str+old_pos, pos-old_pos, 0)
						zend.SmartStrAppendEx(&result, s, 0)
						old_pos = pos + len_
						pos = old_pos - 1
						zend.ZendTmpStringRelease(tmp)
						break
					}
				}
				len_--
			}
		}
		pos++
	}
	if result.s != nil {
		zend.SmartStrAppendlEx(&result, str+old_pos, slen-old_pos, 0)
		zend.SmartStr0(&result)
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = result.s
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
	} else {
		zend.SmartStrFreeEx(&result, 0)
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = input
		__z.value.str = __s
		if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
			__z.u1.type_info = 6
		} else {
			zend.ZendGcAddref(&__s.gc)
			__z.u1.type_info = 6 | 1<<0<<8
		}
	}
	if pats == &str_hash {
		zend.ZendHashDestroy(&str_hash)
	}
	zend._efree(num_bitset)
}

/* }}} */

func PhpCharToStrEx(str *zend.ZendString, from byte, to *byte, to_len int, case_sensitivity int, replace_count *zend.ZendLong) *zend.ZendString {
	var result *zend.ZendString
	var char_count int = 0
	var lc_from int = 0
	var source *byte
	var source_end *byte = str.val + str.len_
	var target *byte
	if case_sensitivity != 0 {
		var p *byte = str.val
		var e *byte = p + str.len_
		for g.Assign(&p, memchr(p, from, e-p)) {
			char_count++
			p++
		}
	} else {
		lc_from = tolower(from)
		for source = str.val; source < source_end; source++ {
			if tolower(*source) == lc_from {
				char_count++
			}
		}
	}
	if char_count == 0 {
		return zend.ZendStringCopy(str)
	}
	if to_len > 0 {
		result = zend.ZendStringSafeAlloc(char_count, to_len-1, str.len_, 0)
	} else {
		result = zend.ZendStringAlloc(str.len_-char_count, 0)
	}
	target = result.val
	if case_sensitivity != 0 {
		var p *byte = str.val
		var e *byte = p + str.len_
		var s *byte = str.val
		for g.Assign(&p, memchr(p, from, e-p)) {
			memcpy(target, s, p-s)
			target += p - s
			memcpy(target, to, to_len)
			target += to_len
			p++
			s = p
			if replace_count != nil {
				*replace_count += 1
			}
		}
		if s < e {
			memcpy(target, s, e-s)
			target += e - s
		}
	} else {
		for source = str.val; source < source_end; source++ {
			if tolower(*source) == lc_from {
				if replace_count != nil {
					*replace_count += 1
				}
				memcpy(target, to, to_len)
				target += to_len
			} else {
				*target = *source
				target++
			}
		}
	}
	*target = 0
	return result
}

/* }}} */

func PhpStrToStrEx(haystack *zend.ZendString, needle *byte, needle_len int, str *byte, str_len int, replace_count *zend.ZendLong) *zend.ZendString {
	var new_str *zend.ZendString
	if needle_len < haystack.len_ {
		var end *byte
		var p *byte
		var r *byte
		var e *byte
		if needle_len == str_len {
			new_str = nil
			end = haystack.val + haystack.len_
			for p = haystack.val; g.Assign(&r, (*byte)(zend.ZendMemnstr(p, needle, needle_len, end))); p = r + needle_len {
				if new_str == nil {
					new_str = zend.ZendStringInit(haystack.val, haystack.len_, 0)
				}
				memcpy(new_str.val+(r-haystack.val), str, str_len)
				*replace_count++
			}
			if new_str == nil {
				goto nothing_todo
			}
			return new_str
		} else {
			var count int = 0
			var o *byte = haystack.val
			var n *byte = needle
			var endp *byte = o + haystack.len_
			for g.Assign(&o, (*byte)(zend.ZendMemnstr(o, n, needle_len, endp))) {
				o += needle_len
				count++
			}
			if count == 0 {

				/* Needle doesn't occur, shortcircuit the actual replacement. */

				goto nothing_todo

				/* Needle doesn't occur, shortcircuit the actual replacement. */

			}
			if str_len > needle_len {
				new_str = zend.ZendStringSafeAlloc(count, str_len-needle_len, haystack.len_, 0)
			} else {
				new_str = zend.ZendStringAlloc(count*(str_len-needle_len)+haystack.len_, 0)
			}
			e = new_str.val
			end = haystack.val + haystack.len_
			for p = haystack.val; g.Assign(&r, (*byte)(zend.ZendMemnstr(p, needle, needle_len, end))); p = r + needle_len {
				memcpy(e, p, r-p)
				e += r - p
				memcpy(e, str, str_len)
				e += str_len
				*replace_count++
			}
			if p < end {
				memcpy(e, p, end-p)
				e += end - p
			}
			*e = '0'
			return new_str
		}
	} else if needle_len > haystack.len_ || memcmp(haystack.val, needle, haystack.len_) {
	nothing_todo:
		return zend.ZendStringCopy(haystack)
	} else {
		if str_len == 0 {
			new_str = zend.ZendEmptyString
		} else if str_len == 1 {
			new_str = zend.ZendOneCharString[zend_uchar(*str)]
		} else {
			new_str = zend.ZendStringInit(str, str_len, 0)
		}
		*replace_count++
		return new_str
	}
}

/* }}} */

func PhpStrToStrIEx(haystack *zend.ZendString, lc_haystack *byte, needle *zend.ZendString, str *byte, str_len int, replace_count *zend.ZendLong) *zend.ZendString {
	var new_str *zend.ZendString = nil
	var lc_needle *zend.ZendString
	if needle.len_ < haystack.len_ {
		var end *byte
		var p *byte
		var r *byte
		var e *byte
		if needle.len_ == str_len {
			lc_needle = PhpStringTolower(needle)
			end = lc_haystack + haystack.len_
			for p = lc_haystack; g.Assign(&r, (*byte)(zend.ZendMemnstr(p, lc_needle.val, lc_needle.len_, end))); p = r + lc_needle.len_ {
				if new_str == nil {
					new_str = zend.ZendStringInit(haystack.val, haystack.len_, 0)
				}
				memcpy(new_str.val+(r-lc_haystack), str, str_len)
				*replace_count++
			}
			zend.ZendStringReleaseEx(lc_needle, 0)
			if new_str == nil {
				goto nothing_todo
			}
			return new_str
		} else {
			var count int = 0
			var o *byte = lc_haystack
			var n *byte
			var endp *byte = o + haystack.len_
			lc_needle = PhpStringTolower(needle)
			n = lc_needle.val
			for g.Assign(&o, (*byte)(zend.ZendMemnstr(o, n, lc_needle.len_, endp))) {
				o += lc_needle.len_
				count++
			}
			if count == 0 {

				/* Needle doesn't occur, shortcircuit the actual replacement. */

				zend.ZendStringReleaseEx(lc_needle, 0)
				goto nothing_todo
			}
			if str_len > lc_needle.len_ {
				new_str = zend.ZendStringSafeAlloc(count, str_len-lc_needle.len_, haystack.len_, 0)
			} else {
				new_str = zend.ZendStringAlloc(count*(str_len-lc_needle.len_)+haystack.len_, 0)
			}
			e = new_str.val
			end = lc_haystack + haystack.len_
			for p = lc_haystack; g.Assign(&r, (*byte)(zend.ZendMemnstr(p, lc_needle.val, lc_needle.len_, end))); p = r + lc_needle.len_ {
				memcpy(e, haystack.val+(p-lc_haystack), r-p)
				e += r - p
				memcpy(e, str, str_len)
				e += str_len
				*replace_count++
			}
			if p < end {
				memcpy(e, haystack.val+(p-lc_haystack), end-p)
				e += end - p
			}
			*e = '0'
			zend.ZendStringReleaseEx(lc_needle, 0)
			return new_str
		}
	} else if needle.len_ > haystack.len_ {
	nothing_todo:
		return zend.ZendStringCopy(haystack)
	} else {
		lc_needle = PhpStringTolower(needle)
		if memcmp(lc_haystack, lc_needle.val, lc_needle.len_) {
			zend.ZendStringReleaseEx(lc_needle, 0)
			goto nothing_todo
		}
		zend.ZendStringReleaseEx(lc_needle, 0)
		new_str = zend.ZendStringInit(str, str_len, 0)
		*replace_count++
		return new_str
	}
}

/* }}} */

func PhpStrToStr(haystack *byte, length int, needle string, needle_len int, str string, str_len int) *zend.ZendString {
	var new_str *zend.ZendString
	if needle_len < length {
		var end *byte
		var s *byte
		var p *byte
		var e *byte
		var r *byte
		if needle_len == str_len {
			new_str = zend.ZendStringInit(haystack, length, 0)
			end = new_str.val + length
			for p = new_str.val; g.Assign(&r, (*byte)(zend.ZendMemnstr(p, needle, needle_len, end))); p = r + needle_len {
				memcpy(r, str, str_len)
			}
			return new_str
		} else {
			if str_len < needle_len {
				new_str = zend.ZendStringAlloc(length, 0)
			} else {
				var count int = 0
				var o *byte = haystack
				var n *byte = needle
				var endp *byte = o + length
				for g.Assign(&o, (*byte)(zend.ZendMemnstr(o, n, needle_len, endp))) {
					o += needle_len
					count++
				}
				if count == 0 {

					/* Needle doesn't occur, shortcircuit the actual replacement. */

					new_str = zend.ZendStringInit(haystack, length, 0)
					return new_str
				} else {
					if str_len > needle_len {
						new_str = zend.ZendStringSafeAlloc(count, str_len-needle_len, length, 0)
					} else {
						new_str = zend.ZendStringAlloc(count*(str_len-needle_len)+length, 0)
					}
				}
			}
			e = new_str.val
			s = e
			end = haystack + length
			for p = haystack; g.Assign(&r, (*byte)(zend.ZendMemnstr(p, needle, needle_len, end))); p = r + needle_len {
				memcpy(e, p, r-p)
				e += r - p
				memcpy(e, str, str_len)
				e += str_len
			}
			if p < end {
				memcpy(e, p, end-p)
				e += end - p
			}
			*e = '0'
			new_str = zend.ZendStringTruncate(new_str, e-s, 0)
			return new_str
		}
	} else if needle_len > length || memcmp(haystack, needle, length) {
		new_str = zend.ZendStringInit(haystack, length, 0)
		return new_str
	} else {
		new_str = zend.ZendStringInit(str, str_len, 0)
		return new_str
	}
}

/* }}} */

func ZifStrtr(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var from *zend.Zval
	var str *zend.ZendString
	var to *byte = nil
	var to_len int = 0
	var ac int = execute_data.This.u2.num_args
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
			r.Assert(_i <= _min_num_args || _optional == 1)
			r.Assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgStr(_arg, &str, 0) == 0 {
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

			zend.ZendParseArgZvalDeref(_arg, &from, 0)
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

			if zend.ZendParseArgString(_arg, &to, &to_len, 0) == 0 {
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
	if ac == 2 && from.u1.v.type_ != 7 {
		core.PhpErrorDocref(nil, 1<<1, "The second argument is not an array")
		return_value.u1.type_info = 2
		return
	}

	/* shortcut for empty string */

	if str.len_ == 0 {
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = zend.ZendEmptyString
		__z.value.str = __s
		__z.u1.type_info = 6
		return
	}
	if ac == 2 {
		var pats *zend.HashTable = from.value.arr
		if pats.nNumOfElements < 1 {
			var __z *zend.Zval = return_value
			var __s *zend.ZendString = str
			__z.value.str = __s
			if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
				__z.u1.type_info = 6
			} else {
				zend.ZendGcAddref(&__s.gc)
				__z.u1.type_info = 6 | 1<<0<<8
			}
			return
		} else if pats.nNumOfElements == 1 {
			var num_key zend.ZendLong
			var str_key *zend.ZendString
			var tmp_str *zend.ZendString
			var replace *zend.ZendString
			var tmp_replace *zend.ZendString
			var entry *zend.Zval
			for {
				var __ht *zend.HashTable = pats
				var _p *zend.Bucket = __ht.arData
				var _end *zend.Bucket = _p + __ht.nNumUsed
				for ; _p != _end; _p++ {
					var _z *zend.Zval = &_p.val
					if _z.u1.v.type_ == 13 {
						_z = _z.value.zv
					}
					if _z.u1.v.type_ == 0 {
						continue
					}
					num_key = _p.h
					str_key = _p.key
					entry = _z
					tmp_str = nil
					if str_key == nil {
						tmp_str = zend.ZendLongToStr(num_key)
						str_key = tmp_str
					}
					replace = zend.ZvalGetTmpString(entry, &tmp_replace)
					if str_key.len_ < 1 {
						var __z *zend.Zval = return_value
						var __s *zend.ZendString = str
						__z.value.str = __s
						if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
							__z.u1.type_info = 6
						} else {
							zend.ZendGcAddref(&__s.gc)
							__z.u1.type_info = 6 | 1<<0<<8
						}
					} else if str_key.len_ == 1 {
						var __z *zend.Zval = return_value
						var __s *zend.ZendString = PhpCharToStrEx(str, str_key.val[0], replace.val, replace.len_, 1, nil)
						__z.value.str = __s
						if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
							__z.u1.type_info = 6
						} else {
							__z.u1.type_info = 6 | 1<<0<<8
						}
					} else {
						var dummy zend.ZendLong
						var __z *zend.Zval = return_value
						var __s *zend.ZendString = PhpStrToStrEx(str, str_key.val, str_key.len_, replace.val, replace.len_, &dummy)
						__z.value.str = __s
						if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
							__z.u1.type_info = 6
						} else {
							__z.u1.type_info = 6 | 1<<0<<8
						}
					}
					zend.ZendTmpStringRelease(tmp_str)
					zend.ZendTmpStringRelease(tmp_replace)
					return
				}
				break
			}
		} else {
			PhpStrtrArray(return_value, str, pats)
		}
	} else {
		if zend.TryConvertToString(from) == 0 {
			return
		}
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = PhpStrtrEx(str, from.value.str.val, to, g.CondF1(from.value.str.len_ < to_len, func() int { return from.value.str.len_ }, to_len))
		__z.value.str = __s
		if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
			__z.u1.type_info = 6
		} else {
			__z.u1.type_info = 6 | 1<<0<<8
		}
		return
	}
}

/* }}} */

func ZifStrrev(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var str *zend.ZendString
	var s *byte
	var e *byte
	var p *byte
	var n *zend.ZendString
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
			r.Assert(_i <= _min_num_args || _optional == 1)
			r.Assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgStr(_arg, &str, 0) == 0 {
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
	n = zend.ZendStringAlloc(str.len_, 0)
	p = n.val
	s = str.val
	e = s + str.len_
	e--
	for e >= s {
		*e--
		g.PostInc(&(*p)) = (*e) + 1
	}
	*p = '0'
	var __z *zend.Zval = return_value
	var __s *zend.ZendString = n
	__z.value.str = __s
	__z.u1.type_info = 6 | 1<<0<<8
}

/* }}} */

func PhpSimilarStr(txt1 *byte, len1 int, txt2 *byte, len2 int, pos1 *int, pos2 *int, max *int, count *int) {
	var p *byte
	var q *byte
	var end1 *byte = (*byte)(txt1 + len1)
	var end2 *byte = (*byte)(txt2 + len2)
	var l int
	*max = 0
	*count = 0
	for p = (*byte)(txt1); p < end1; p++ {
		for q = (*byte)(txt2); q < end2; q++ {
			for l = 0; p+l < end1 && q+l < end2 && p[l] == q[l]; l++ {

			}
			if l > (*max) {
				*max = l
				*count += 1
				*pos1 = p - txt1
				*pos2 = q - txt2
			}
		}
	}
}

/* }}} */

func PhpSimilarChar(txt1 *byte, len1 int, txt2 *byte, len2 int) int {
	var sum int
	var pos1 int = 0
	var pos2 int = 0
	var max int
	var count int
	PhpSimilarStr(txt1, len1, txt2, len2, &pos1, &pos2, &max, &count)
	if g.Assign(&sum, max) {
		if pos1 != 0 && pos2 != 0 && count > 1 {
			sum += PhpSimilarChar(txt1, pos1, txt2, pos2)
		}
		if pos1+max < len1 && pos2+max < len2 {
			sum += PhpSimilarChar(txt1+pos1+max, len1-pos1-max, txt2+pos2+max, len2-pos2-max)
		}
	}
	return sum
}

/* }}} */

func ZifSimilarText(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var t1 *zend.ZendString
	var t2 *zend.ZendString
	var percent *zend.Zval = nil
	var ac int = execute_data.This.u2.num_args
	var sim int
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
			r.Assert(_i <= _min_num_args || _optional == 1)
			r.Assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgStr(_arg, &t1, 0) == 0 {
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

			if zend.ZendParseArgStr(_arg, &t2, 0) == 0 {
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

			zend.ZendParseArgZvalDeref(_arg, &percent, 0)
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
	if t1.len_+t2.len_ == 0 {
		if ac > 2 {
			for {
				r.Assert(percent.u1.v.type_ == 10)
				for {
					var _zv *zend.Zval = percent
					var ref *zend.ZendReference = _zv.value.ref
					if ref.sources.ptr != nil {
						zend.ZendTryAssignTypedRefDouble(ref, 0)
						break
					}
					_zv = &ref.val
					zend.ZvalPtrDtor(_zv)
					var __z *zend.Zval = _zv
					__z.value.dval = 0
					__z.u1.type_info = 5
					break
				}
				break
			}
		}
		var __z *zend.Zval = return_value
		__z.value.lval = 0
		__z.u1.type_info = 4
		return
	}
	sim = PhpSimilarChar(t1.val, t1.len_, t2.val, t2.len_)
	if ac > 2 {
		for {
			r.Assert(percent.u1.v.type_ == 10)
			for {
				var _zv *zend.Zval = percent
				var ref *zend.ZendReference = _zv.value.ref
				if ref.sources.ptr != nil {
					zend.ZendTryAssignTypedRefDouble(ref, sim*200.0/(t1.len_+t2.len_))
					break
				}
				_zv = &ref.val
				zend.ZvalPtrDtor(_zv)
				var __z *zend.Zval = _zv
				__z.value.dval = sim * 200.0 / (t1.len_ + t2.len_)
				__z.u1.type_info = 5
				break
			}
			break
		}
	}
	var __z *zend.Zval = return_value
	__z.value.lval = sim
	__z.u1.type_info = 4
	return
}

/* }}} */

func ZifAddcslashes(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var str *zend.ZendString
	var what *zend.ZendString
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
			r.Assert(_i <= _min_num_args || _optional == 1)
			r.Assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgStr(_arg, &str, 0) == 0 {
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

			if zend.ZendParseArgStr(_arg, &what, 0) == 0 {
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
	if str.len_ == 0 {
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = zend.ZendEmptyString
		__z.value.str = __s
		__z.u1.type_info = 6
		return
	}
	if what.len_ == 0 {
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = str
		__z.value.str = __s
		if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
			__z.u1.type_info = 6
		} else {
			zend.ZendGcAddref(&__s.gc)
			__z.u1.type_info = 6 | 1<<0<<8
		}
		return
	}
	var __z *zend.Zval = return_value
	var __s *zend.ZendString = PhpAddcslashesStr(str.val, str.len_, what.val, what.len_)
	__z.value.str = __s
	if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
		__z.u1.type_info = 6
	} else {
		__z.u1.type_info = 6 | 1<<0<<8
	}
	return
}

/* }}} */

func ZifAddslashes(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var str *zend.ZendString
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
			r.Assert(_i <= _min_num_args || _optional == 1)
			r.Assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgStr(_arg, &str, 0) == 0 {
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
	if str.len_ == 0 {
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = zend.ZendEmptyString
		__z.value.str = __s
		__z.u1.type_info = 6
		return
	}
	var __z *zend.Zval = return_value
	var __s *zend.ZendString = PhpAddslashes(str)
	__z.value.str = __s
	if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
		__z.u1.type_info = 6
	} else {
		__z.u1.type_info = 6 | 1<<0<<8
	}
	return
}

/* }}} */

func ZifStripcslashes(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var str *zend.ZendString
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
			r.Assert(_i <= _min_num_args || _optional == 1)
			r.Assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgStr(_arg, &str, 0) == 0 {
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
	var __z *zend.Zval = return_value
	var __s *zend.ZendString = zend.ZendStringInit(str.val, str.len_, 0)
	__z.value.str = __s
	__z.u1.type_info = 6 | 1<<0<<8
	PhpStripcslashes(return_value.value.str)
}

/* }}} */

func ZifStripslashes(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var str *zend.ZendString
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
			r.Assert(_i <= _min_num_args || _optional == 1)
			r.Assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgStr(_arg, &str, 0) == 0 {
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
	var __z *zend.Zval = return_value
	var __s *zend.ZendString = zend.ZendStringInit(str.val, str.len_, 0)
	__z.value.str = __s
	__z.u1.type_info = 6 | 1<<0<<8
	PhpStripslashes(return_value.value.str)
}

/* }}} */

func PhpStripcslashes(str *zend.ZendString) {
	var source *byte
	var end *byte
	var target *byte
	var nlen int = str.len_
	var i int
	var numtmp []byte
	source = (*byte)(str.val)
	end = source + str.len_
	target = str.val
	for ; source < end; source++ {
		if (*source) == '\\' && source+1 < end {
			source++
			switch *source {
			case 'n':
				g.PostInc(&(*target)) = '\n'
				nlen--
				break
			case 'r':
				g.PostInc(&(*target)) = '\r'
				nlen--
				break
			case 'a':
				g.PostInc(&(*target)) = 'a'
				nlen--
				break
			case 't':
				g.PostInc(&(*target)) = '\t'
				nlen--
				break
			case 'v':
				g.PostInc(&(*target)) = 'v'
				nlen--
				break
			case 'b':
				g.PostInc(&(*target)) = 'b'
				nlen--
				break
			case 'f':
				g.PostInc(&(*target)) = 'f'
				nlen--
				break
			case '\\':
				g.PostInc(&(*target)) = '\\'
				nlen--
				break
			case 'x':
				if source+1 < end && isxdigit(int(*(source + 1))) {
					numtmp[0] = *(g.PreInc(&source))
					if source+1 < end && isxdigit(int(*(source + 1))) {
						numtmp[1] = *(g.PreInc(&source))
						numtmp[2] = '0'
						nlen -= 3
					} else {
						numtmp[1] = '0'
						nlen -= 2
					}
					g.PostInc(&(*target)) = byte(strtol(numtmp, nil, 16))
					break
				}
			default:
				i = 0
				for source < end && (*source) >= '0' && (*source) <= '7' && i < 3 {
					*source++
					numtmp[g.PostInc(&i)] = (*source) - 1
				}
				if i != 0 {
					numtmp[i] = '0'
					g.PostInc(&(*target)) = byte(strtol(numtmp, nil, 8))
					nlen -= i
					source--
				} else {
					g.PostInc(&(*target)) = *source
					nlen--
				}
			}
		} else {
			g.PostInc(&(*target)) = *source
		}
	}
	if nlen != 0 {
		*target = '0'
	}
	str.len_ = nlen
}

/* }}} */

func PhpAddcslashesStr(str *byte, len_ int, what *byte, wlength int) *zend.ZendString {
	var flags []byte
	var target *byte
	var source *byte
	var end *byte
	var c byte
	var newlen int
	var new_str *zend.ZendString = zend.ZendStringSafeAlloc(4, len_, 0, 0)
	PhpCharmask((*uint8)(what), wlength, flags)
	source = str
	end = source + len_
	target = new_str.val
	for ; source < end; source++ {
		c = *source
		if flags[uint8(c)] {
			if uint8(c < 32 || uint8(c > 126) != 0) != 0 {
				g.PostInc(&(*target)) = '\\'
				switch c {
				case '\n':
					g.PostInc(&(*target)) = 'n'
					break
				case '\t':
					g.PostInc(&(*target)) = 't'
					break
				case '\r':
					g.PostInc(&(*target)) = 'r'
					break
				case 'a':
					g.PostInc(&(*target)) = 'a'
					break
				case 'v':
					g.PostInc(&(*target)) = 'v'
					break
				case 'b':
					g.PostInc(&(*target)) = 'b'
					break
				case 'f':
					g.PostInc(&(*target)) = 'f'
					break
				default:
					target += sprintf(target, "%03o", uint8(c))
				}
				continue
			}
			g.PostInc(&(*target)) = '\\'
		}
		g.PostInc(&(*target)) = c
	}
	*target = 0
	newlen = target - new_str.val
	if newlen < len_*4 {
		new_str = zend.ZendStringTruncate(new_str, newlen, 0)
	}
	return new_str
}

/* }}} */

func PhpAddcslashes(str *zend.ZendString, what string, wlength int) *zend.ZendString {
	return PhpAddcslashesStr(str.val, str.len_, what, wlength)
}

/* }}} */

func PhpAddslashes(str *zend.ZendString) *zend.ZendString {
	/* maximum string length, worst case situation */

	var target *byte
	var source *byte
	var end *byte
	var offset int
	var new_str *zend.ZendString
	if str == nil {
		return zend.ZendEmptyString
	}
	source = str.val
	end = source + str.len_
	for source < end {
		switch *source {
		case '0':

		case '\'':

		case '"':

		case '\\':
			goto do_escape
		default:
			source++
			break
		}
	}
	return zend.ZendStringCopy(str)
do_escape:
	offset = source - (*byte)(str.val)
	new_str = zend.ZendStringSafeAlloc(2, str.len_-offset, offset, 0)
	memcpy(new_str.val, str.val, offset)
	target = new_str.val + offset
	for source < end {
		switch *source {
		case '0':
			g.PostInc(&(*target)) = '\\'
			g.PostInc(&(*target)) = '0'
			break
		case '\'':

		case '"':

		case '\\':
			g.PostInc(&(*target)) = '\\'
		default:
			g.PostInc(&(*target)) = *source
			break
		}
		source++
	}
	*target = '0'
	if new_str.len_-(target-new_str.val) > 16 {
		new_str = zend.ZendStringTruncate(new_str, target-new_str.val, 0)
	} else {
		new_str.len_ = target - new_str.val
	}
	return new_str
}

/* }}} */

func PhpStripslashesImpl(str *byte, out *byte, len_ int) *byte {
	for len_ > 0 {
		if (*str) == '\\' {
			str++
			len_--
			if len_ > 0 {
				if (*str) == '0' {
					g.PostInc(&(*out)) = '0'
					str++
				} else {
					*str++
					g.PostInc(&(*out)) = (*str) - 1
				}
				len_--
			}
		} else {
			*str++
			g.PostInc(&(*out)) = (*str) - 1
			len_--
		}
	}
	return out
}
func PhpStripslashes(str *zend.ZendString) {
	var t *byte = PhpStripslashesImpl(str.val, str.val, str.len_)
	if t != str.val+str.len_ {
		str.len_ = t - str.val
		str.val[str.len_] = '0'
	}
}

/* }}} */

/* }}} */

// #define _HEB_BLOCK_TYPE_ENG       1

// #define _HEB_BLOCK_TYPE_HEB       2

// #define isheb(c) ( ( ( ( ( unsigned char ) c ) >= 224 ) && ( ( ( unsigned char ) c ) <= 250 ) ) ? 1 : 0 )

// #define _isblank(c) ( ( ( ( ( unsigned char ) c ) == ' ' || ( ( unsigned char ) c ) == '\t' ) ) ? 1 : 0 )

// #define _isnewline(c) ( ( ( ( ( unsigned char ) c ) == '\n' || ( ( unsigned char ) c ) == '\r' ) ) ? 1 : 0 )

/* {{{ php_str_replace_in_subject
 */

func PhpStrReplaceInSubject(search *zend.Zval, replace *zend.Zval, subject *zend.Zval, result *zend.Zval, case_sensitivity int) zend.ZendLong {
	var search_entry *zend.Zval
	var tmp_result *zend.ZendString
	var tmp_subject_str *zend.ZendString
	var replace_value *byte = nil
	var replace_len int = 0
	var replace_count zend.ZendLong = 0
	var subject_str *zend.ZendString
	var lc_subject_str *zend.ZendString = nil
	var replace_idx uint32

	/* Make sure we're dealing with strings. */

	subject_str = zend.ZvalGetTmpString(subject, &tmp_subject_str)
	if subject_str.len_ == 0 {
		zend.ZendTmpStringRelease(tmp_subject_str)
		var __z *zend.Zval = result
		var __s *zend.ZendString = zend.ZendEmptyString
		__z.value.str = __s
		__z.u1.type_info = 6
		return 0
	}

	/* If search is an array */

	if search.u1.v.type_ == 7 {

		/* Duplicate subject string for repeated replacement */

		zend.ZendStringAddref(subject_str)
		if replace.u1.v.type_ == 7 {
			replace_idx = 0
		} else {

			/* Set replacement value to the passed one */

			replace_value = replace.value.str.val
			replace_len = replace.value.str.len_
		}

		/* For each entry in the search array, get the entry */

		for {
			var __ht *zend.HashTable = search.value.arr
			var _p *zend.Bucket = __ht.arData
			var _end *zend.Bucket = _p + __ht.nNumUsed
			for ; _p != _end; _p++ {
				var _z *zend.Zval = &_p.val
				if _z.u1.v.type_ == 13 {
					_z = _z.value.zv
				}
				if _z.u1.v.type_ == 0 {
					continue
				}
				search_entry = _z

				/* Make sure we're dealing with strings. */

				var tmp_search_str *zend.ZendString
				var search_str *zend.ZendString = zend.ZvalGetTmpString(search_entry, &tmp_search_str)
				var replace_entry_str *zend.ZendString
				var tmp_replace_entry_str *zend.ZendString = nil

				/* If replace is an array. */

				if replace.u1.v.type_ == 7 {

					/* Get current entry */

					var replace_entry *zend.Zval = nil
					for replace_idx < replace.value.arr.nNumUsed {
						replace_entry = &(*replace).value.arr.arData[replace_idx].val
						if replace_entry.u1.v.type_ != 0 {
							break
						}
						replace_idx++
					}
					if replace_idx < replace.value.arr.nNumUsed {

						/* Make sure we're dealing with strings. */

						replace_entry_str = zend.ZvalGetTmpString(replace_entry, &tmp_replace_entry_str)

						/* Set replacement value to the one we got from array */

						replace_value = replace_entry_str.val
						replace_len = replace_entry_str.len_
						replace_idx++
					} else {

						/* We've run out of replacement strings, so use an empty one. */

						replace_value = ""
						replace_len = 0
					}
				}
				if search_str.len_ == 1 {
					var old_replace_count zend.ZendLong = replace_count
					tmp_result = PhpCharToStrEx(subject_str, search_str.val[0], replace_value, replace_len, case_sensitivity, &replace_count)
					if lc_subject_str != nil && replace_count != old_replace_count {
						zend.ZendStringReleaseEx(lc_subject_str, 0)
						lc_subject_str = nil
					}
				} else if search_str.len_ > 1 {
					if case_sensitivity != 0 {
						tmp_result = PhpStrToStrEx(subject_str, search_str.val, search_str.len_, replace_value, replace_len, &replace_count)
					} else {
						var old_replace_count zend.ZendLong = replace_count
						if lc_subject_str == nil {
							lc_subject_str = PhpStringTolower(subject_str)
						}
						tmp_result = PhpStrToStrIEx(subject_str, lc_subject_str.val, search_str, replace_value, replace_len, &replace_count)
						if replace_count != old_replace_count {
							zend.ZendStringReleaseEx(lc_subject_str, 0)
							lc_subject_str = nil
						}
					}
				} else {
					zend.ZendTmpStringRelease(tmp_search_str)
					zend.ZendTmpStringRelease(tmp_replace_entry_str)
					continue
				}
				zend.ZendTmpStringRelease(tmp_search_str)
				zend.ZendTmpStringRelease(tmp_replace_entry_str)
				if subject_str == tmp_result {
					zend.ZendStringDelref(subject_str)
				} else {
					zend.ZendStringReleaseEx(subject_str, 0)
					subject_str = tmp_result
					if subject_str.len_ == 0 {
						zend.ZendStringReleaseEx(subject_str, 0)
						var __z *zend.Zval = result
						var __s *zend.ZendString = zend.ZendEmptyString
						__z.value.str = __s
						__z.u1.type_info = 6
						if lc_subject_str != nil {
							zend.ZendStringReleaseEx(lc_subject_str, 0)
						}
						zend.ZendTmpStringRelease(tmp_subject_str)
						return replace_count
					}
				}
			}
			break
		}
		var __z *zend.Zval = result
		var __s *zend.ZendString = subject_str
		__z.value.str = __s
		if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
			__z.u1.type_info = 6
		} else {
			__z.u1.type_info = 6 | 1<<0<<8
		}
		if lc_subject_str != nil {
			zend.ZendStringReleaseEx(lc_subject_str, 0)
		}
	} else {
		r.Assert(search.u1.v.type_ == 6)
		if search.value.str.len_ == 1 {
			var __z *zend.Zval = result
			var __s *zend.ZendString = PhpCharToStrEx(subject_str, search.value.str.val[0], replace.value.str.val, replace.value.str.len_, case_sensitivity, &replace_count)
			__z.value.str = __s
			if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
				__z.u1.type_info = 6
			} else {
				__z.u1.type_info = 6 | 1<<0<<8
			}
		} else if search.value.str.len_ > 1 {
			if case_sensitivity != 0 {
				var __z *zend.Zval = result
				var __s *zend.ZendString = PhpStrToStrEx(subject_str, search.value.str.val, search.value.str.len_, replace.value.str.val, replace.value.str.len_, &replace_count)
				__z.value.str = __s
				if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
					__z.u1.type_info = 6
				} else {
					__z.u1.type_info = 6 | 1<<0<<8
				}
			} else {
				lc_subject_str = PhpStringTolower(subject_str)
				var __z *zend.Zval = result
				var __s *zend.ZendString = PhpStrToStrIEx(subject_str, lc_subject_str.val, search.value.str, replace.value.str.val, replace.value.str.len_, &replace_count)
				__z.value.str = __s
				if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
					__z.u1.type_info = 6
				} else {
					__z.u1.type_info = 6 | 1<<0<<8
				}
				zend.ZendStringReleaseEx(lc_subject_str, 0)
			}
		} else {
			var __z *zend.Zval = result
			var __s *zend.ZendString = subject_str
			__z.value.str = __s
			if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
				__z.u1.type_info = 6
			} else {
				zend.ZendGcAddref(&__s.gc)
				__z.u1.type_info = 6 | 1<<0<<8
			}
		}
	}
	zend.ZendTmpStringRelease(tmp_subject_str)
	return replace_count
}

/* }}} */

func PhpStrReplaceCommon(execute_data *zend.ZendExecuteData, return_value *zend.Zval, case_sensitivity int) {
	var subject *zend.Zval
	var search *zend.Zval
	var replace *zend.Zval
	var subject_entry *zend.Zval
	var zcount *zend.Zval = nil
	var result zend.Zval
	var string_key *zend.ZendString
	var num_key zend.ZendUlong
	var count zend.ZendLong = 0
	var argc int = execute_data.This.u2.num_args
	for {
		var _flags int = 0
		var _min_num_args int = 3
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
			r.Assert(_i <= _min_num_args || _optional == 1)
			r.Assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			zend.ZendParseArgZvalDeref(_arg, &search, 0)
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

			zend.ZendParseArgZvalDeref(_arg, &replace, 0)
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

			zend.ZendParseArgZvalDeref(_arg, &subject, 0)
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

			zend.ZendParseArgZvalDeref(_arg, &zcount, 0)
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

	/* Make sure we're dealing with strings and do the replacement. */

	if search.u1.v.type_ != 7 {
		if search.u1.v.type_ != 6 {
			if search.u1.v.type_ != 6 {
				zend._convertToString(search)
			}
		}
		if replace.u1.v.type_ != 6 {
			if replace.u1.v.type_ != 6 {
				if replace.u1.v.type_ != 6 {
					zend._convertToString(replace)
				}
			}
		}
	} else if replace.u1.v.type_ != 7 {
		if replace.u1.v.type_ != 6 {
			if replace.u1.v.type_ != 6 {
				zend._convertToString(replace)
			}
		}
	}
	if zend.EG.exception != nil {
		return
	}

	/* if subject is an array */

	if subject.u1.v.type_ == 7 {
		var __arr *zend.ZendArray = zend._zendNewArray(0)
		var __z *zend.Zval = return_value
		__z.value.arr = __arr
		__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8

		/* For each subject entry, convert it to string, then perform replacement
		   and add the result to the return_value array. */

		for {
			var __ht *zend.HashTable = subject.value.arr
			var _p *zend.Bucket = __ht.arData
			var _end *zend.Bucket = _p + __ht.nNumUsed
			for ; _p != _end; _p++ {
				var _z *zend.Zval = &_p.val
				if _z.u1.v.type_ == 13 {
					_z = _z.value.zv
				}
				if _z.u1.v.type_ == 0 {
					continue
				}
				num_key = _p.h
				string_key = _p.key
				subject_entry = _z
				if subject_entry.u1.v.type_ == 10 {
					subject_entry = &(*subject_entry).value.ref.val
				}
				if subject_entry.u1.v.type_ != 7 && subject_entry.u1.v.type_ != 8 {
					count += PhpStrReplaceInSubject(search, replace, subject_entry, &result, case_sensitivity)
				} else {
					var _z1 *zend.Zval = &result
					var _z2 *zend.Zval = subject_entry
					var _gc *zend.ZendRefcounted = _z2.value.counted
					var _t uint32 = _z2.u1.type_info
					_z1.value.counted = _gc
					_z1.u1.type_info = _t
					if (_t & 0xff00) != 0 {
						zend.ZendGcAddref(&_gc.gc)
					}
				}

				/* Add to return array */

				if string_key != nil {
					zend.ZendHashAddNew(return_value.value.arr, string_key, &result)
				} else {
					zend.ZendHashIndexAddNew(return_value.value.arr, num_key, &result)
				}

				/* Add to return array */

			}
			break
		}

		/* For each subject entry, convert it to string, then perform replacement
		   and add the result to the return_value array. */

	} else {
		count = PhpStrReplaceInSubject(search, replace, subject, return_value, case_sensitivity)
	}
	if argc > 3 {
		for {
			r.Assert(zcount.u1.v.type_ == 10)
			for {
				var _zv *zend.Zval = zcount
				var ref *zend.ZendReference = _zv.value.ref
				if ref.sources.ptr != nil {
					zend.ZendTryAssignTypedRefLong(ref, count)
					break
				}
				_zv = &ref.val
				zend.ZvalPtrDtor(_zv)
				var __z *zend.Zval = _zv
				__z.value.lval = count
				__z.u1.type_info = 4
				break
			}
			break
		}
	}
}

/* }}} */

func ZifStrReplace(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpStrReplaceCommon(execute_data, return_value, 1)
}

/* }}} */

func ZifStrIreplace(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpStrReplaceCommon(execute_data, return_value, 0)
}

/* }}} */

func PhpHebrev(execute_data *zend.ZendExecuteData, return_value *zend.Zval, convert_newlines int) {
	var str *byte
	var heb_str *byte
	var target *byte
	var tmp *byte
	var block_start int
	var block_end int
	var block_type int
	var block_length int
	var i int
	var max_chars zend.ZendLong = 0
	var char_count zend.ZendLong
	var begin int
	var end int
	var orig_begin int
	var str_len int
	var broken_str *zend.ZendString
	for {
		var _flags int = 0
		var _min_num_args int = 1
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
			r.Assert(_i <= _min_num_args || _optional == 1)
			r.Assert(_i > _min_num_args || _optional == 0)
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

			if zend.ZendParseArgLong(_arg, &max_chars, &_dummy, 0, 0) == 0 {
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
	if str_len == 0 {
		return_value.u1.type_info = 2
		return
	}
	tmp = str
	block_end = 0
	block_start = block_end
	heb_str = (*byte)(zend._emalloc(str_len + 1))
	target = heb_str + str_len
	*target = 0
	target--
	block_length = 0
	if g.Cond(uint8(*tmp) >= 224 && uint8(*tmp) <= 250, 1, 0) {
		block_type = 2
	} else {
		block_type = 1
	}
	for {
		if block_type == 2 {
			for (g.Cond(uint8(int(*(tmp + 1))) >= 224 && uint8(int(*(tmp + 1))) <= 250, 1, 0) || g.Cond(uint8(int(*(tmp + 1))) == ' ' || uint8(int(*(tmp + 1))) == '\t', 1, 0) || ispunct(int(*(tmp + 1))) || int((*(tmp + 1)) == '\n') != 0) && block_end < str_len-1 {
				tmp++
				block_end++
				block_length++
			}
			for i = block_start + 1; i <= block_end+1; i++ {
				*target = str[i-1]
				switch *target {
				case '(':
					*target = ')'
					break
				case ')':
					*target = '('
					break
				case '[':
					*target = ']'
					break
				case ']':
					*target = '['
					break
				case '{':
					*target = '}'
					break
				case '}':
					*target = '{'
					break
				case '<':
					*target = '>'
					break
				case '>':
					*target = '<'
					break
				case '\\':
					*target = '/'
					break
				case '/':
					*target = '\\'
					break
				default:
					break
				}
				target--
			}
			block_type = 1
		} else {
			for !(g.Cond(uint8(*(tmp + 1)) >= 224 && uint8(*(tmp + 1)) <= 250, 1, 0)) && int((*(tmp + 1)) != '\n' && block_end < str_len-1) != 0 {
				tmp++
				block_end++
				block_length++
			}
			for (g.Cond(uint8(int(*tmp)) == ' ' || uint8(int(*tmp)) == '\t', 1, 0) || ispunct(int(*tmp))) && (*tmp) != '/' && (*tmp) != '-' && block_end > block_start {
				tmp--
				block_end--
			}
			for i = block_end + 1; i >= block_start+1; i-- {
				*target = str[i-1]
				target--
			}
			block_type = 2
		}
		block_start = block_end + 1
		if block_end >= str_len-1 {
			break
		}
	}
	broken_str = zend.ZendStringAlloc(str_len, 0)
	end = str_len - 1
	begin = end
	target = broken_str.val
	for true {
		char_count = 0
		for (max_chars == 0 || max_chars > 0 && char_count < max_chars) && begin > 0 {
			char_count++
			begin--
			if g.Cond(uint8(heb_str[begin]) == '\n' || uint8(heb_str[begin]) == '\r', 1, 0) {
				for begin > 0 && g.Cond(uint8(heb_str[begin-1]) == '\n' || uint8(heb_str[begin-1]) == '\r', 1, 0) {
					begin--
					char_count++
				}
				break
			}
		}
		if max_chars >= 0 && char_count == max_chars {
			var new_char_count int = char_count
			var new_begin int = begin
			for new_char_count > 0 {
				if g.Cond(uint8(heb_str[new_begin]) == ' ' || uint8(heb_str[new_begin]) == '\t', 1, 0) || g.Cond(uint8(heb_str[new_begin]) == '\n' || uint8(heb_str[new_begin]) == '\r', 1, 0) {
					break
				}
				new_begin++
				new_char_count--
			}
			if new_char_count > 0 {
				begin = new_begin
			}
		}
		orig_begin = begin
		if g.Cond(uint8(heb_str[begin]) == ' ' || uint8(heb_str[begin]) == '\t', 1, 0) {
			heb_str[begin] = '\n'
		}
		for begin <= end && g.Cond(uint8(heb_str[begin]) == '\n' || uint8(heb_str[begin]) == '\r', 1, 0) {
			begin++
		}
		for i = begin; i <= end; i++ {
			*target = heb_str[i]
			target++
		}
		for i = orig_begin; i <= end && g.Cond(uint8(heb_str[i]) == '\n' || uint8(heb_str[i]) == '\r', 1, 0); i++ {
			*target = heb_str[i]
			target++
		}
		begin = orig_begin
		if begin == 0 {
			*target = 0
			break
		}
		begin--
		end = begin
	}
	zend._efree(heb_str)
	if convert_newlines != 0 {
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = PhpCharToStrEx(broken_str, '\n', "<br />\n", 7, 1, nil)
		__z.value.str = __s
		if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
			__z.u1.type_info = 6
		} else {
			__z.u1.type_info = 6 | 1<<0<<8
		}
		zend.ZendStringReleaseEx(broken_str, 0)
	} else {
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = broken_str
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
		return
	}
}

/* }}} */

func ZifHebrev(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpHebrev(execute_data, return_value, 0)
}

/* }}} */

func ZifHebrevc(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpHebrev(execute_data, return_value, 1)
}

/* }}} */

func ZifNl2br(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	/* in brief this inserts <br /> or <br> before matched regexp \n\r?|\r\n? */

	var tmp *byte
	var end *byte
	var str *zend.ZendString
	var target *byte
	var repl_cnt int = 0
	var is_xhtml zend.ZendBool = 1
	var result *zend.ZendString
	for {
		var _flags int = 0
		var _min_num_args int = 1
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
			r.Assert(_i <= _min_num_args || _optional == 1)
			r.Assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgStr(_arg, &str, 0) == 0 {
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

			if zend.ZendParseArgBool(_arg, &is_xhtml, &_dummy, 0) == 0 {
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
	tmp = str.val
	end = str.val + str.len_

	/* it is really faster to scan twice and allocate mem once instead of scanning once
	   and constantly reallocing */

	for tmp < end {
		if (*tmp) == '\r' {
			if (*(tmp + 1)) == '\n' {
				tmp++
			}
			repl_cnt++
		} else if (*tmp) == '\n' {
			if (*(tmp + 1)) == '\r' {
				tmp++
			}
			repl_cnt++
		}
		tmp++
	}
	if repl_cnt == 0 {
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = str
		__z.value.str = __s
		if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
			__z.u1.type_info = 6
		} else {
			zend.ZendGcAddref(&__s.gc)
			__z.u1.type_info = 6 | 1<<0<<8
		}
		return
	}
	var repl_len int = g.CondF(is_xhtml != 0, func() int { return g.SizeOf("\"<br />\"") - 1 }, func() int { return g.SizeOf("\"<br>\"") - 1 })
	result = zend.ZendStringSafeAlloc(repl_cnt, repl_len, str.len_, 0)
	target = result.val
	tmp = str.val
	for tmp < end {
		switch *tmp {
		case '\r':

		case '\n':
			g.PostInc(&(*target)) = '<'
			g.PostInc(&(*target)) = 'b'
			g.PostInc(&(*target)) = 'r'
			if is_xhtml != 0 {
				g.PostInc(&(*target)) = ' '
				g.PostInc(&(*target)) = '/'
			}
			g.PostInc(&(*target)) = '>'
			if (*tmp) == '\r' && (*(tmp + 1)) == '\n' || (*tmp) == '\n' && (*(tmp + 1)) == '\r' {
				*tmp++
				g.PostInc(&(*target)) = (*tmp) - 1
			}
		default:
			g.PostInc(&(*target)) = *tmp
		}
		tmp++
	}
	*target = '0'
	var __z *zend.Zval = return_value
	var __s *zend.ZendString = result
	__z.value.str = __s
	__z.u1.type_info = 6 | 1<<0<<8
	return
}

/* }}} */

func ZifStripTags(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var buf *zend.ZendString
	var str *zend.ZendString
	var allow *zend.Zval = nil
	var allowed_tags *byte = nil
	var allowed_tags_len int = 0
	var tags_ss zend.SmartStr = zend.SmartStr{0}
	for {
		var _flags int = 0
		var _min_num_args int = 1
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
			r.Assert(_i <= _min_num_args || _optional == 1)
			r.Assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgStr(_arg, &str, 0) == 0 {
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

			zend.ZendParseArgZvalDeref(_arg, &allow, 0)
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
	if allow != nil {
		if allow.u1.v.type_ == 7 {
			var tmp *zend.Zval
			var tag *zend.ZendString
			for {
				var __ht *zend.HashTable = allow.value.arr
				var _p *zend.Bucket = __ht.arData
				var _end *zend.Bucket = _p + __ht.nNumUsed
				for ; _p != _end; _p++ {
					var _z *zend.Zval = &_p.val

					if _z.u1.v.type_ == 0 {
						continue
					}
					tmp = _z
					tag = zend.ZvalGetString(tmp)
					zend.SmartStrAppendcEx(&tags_ss, '<', 0)
					zend.SmartStrAppendEx(&tags_ss, tag, 0)
					zend.SmartStrAppendcEx(&tags_ss, '>', 0)
					zend.ZendStringRelease(tag)
				}
				break
			}
			if tags_ss.s != nil {
				zend.SmartStr0(&tags_ss)
				allowed_tags = tags_ss.s.val
				allowed_tags_len = tags_ss.s.len_
			}
		} else {

			/* To maintain a certain BC, we allow anything for the second parameter and return original string */

			if allow.u1.v.type_ != 6 {
				zend._convertToString(allow)
			}
			allowed_tags = allow.value.str.val
			allowed_tags_len = allow.value.str.len_
		}
	}
	buf = zend.ZendStringInit(str.val, str.len_, 0)
	buf.len_ = PhpStripTagsEx(buf.val, str.len_, nil, allowed_tags, allowed_tags_len, 0)
	zend.SmartStrFreeEx(&tags_ss, 0)
	var __z *zend.Zval = return_value
	var __s *zend.ZendString = buf
	__z.value.str = __s
	__z.u1.type_info = 6 | 1<<0<<8
	return
}

/* }}} */

func ZifSetlocale(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var args *zend.Zval = nil
	var plocale *zend.Zval
	var loc *zend.ZendString
	var retval *byte
	var cat zend.ZendLong
	var num_args int
	var i int = 0
	var idx uint32
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = -1
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

			if zend.ZendParseArgLong(_arg, &cat, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = 4
				break
			}
			var _num_varargs int = _num_args - _i - 0
			if _num_varargs > 0 {
				args = _real_arg + 1
				num_args = _num_varargs
				_i += _num_varargs
				_real_arg += _num_varargs
			} else {
				args = nil
				num_args = 0
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
	idx = 0
	for true {
		if args[0].u1.v.type_ == 7 {
			for idx < args[0].value.arr.nNumUsed {
				plocale = &args[0].value.arr.arData[idx].val
				if plocale.u1.v.type_ != 0 {
					break
				}
				idx++
			}
			if idx >= args[0].value.arr.nNumUsed {
				break
			}
		} else {
			plocale = &args[i]
		}
		loc = zend.ZvalTryGetString(plocale)
		if loc == nil {
			return
		}
		if !(strcmp("0", loc.val)) {
			zend.ZendStringReleaseEx(loc, 0)
			loc = nil
		} else {
			if loc.len_ >= 255 {
				core.PhpErrorDocref(nil, 1<<1, "Specified locale name is too long")
				zend.ZendStringReleaseEx(loc, 0)
				break
			}
		}
		retval = setlocale(cat, g.CondF1(loc != nil, func() []byte { return loc.val }, nil))
		if retval != nil {
			if loc != nil {

				/* Remember if locale was changed */

				var len_ int = strlen(retval)
				BasicGlobals.SetLocaleChanged(1)
				if cat == LC_CTYPE || cat == LC_ALL {
					if BasicGlobals.GetLocaleString() != nil {
						zend.ZendStringReleaseEx(BasicGlobals.GetLocaleString(), 0)
					}
					if len_ == loc.len_ && !(memcmp(loc.val, retval, len_)) {
						BasicGlobals.SetLocaleString(zend.ZendStringCopy(loc))
						var __z *zend.Zval = return_value
						var __s *zend.ZendString = BasicGlobals.GetLocaleString()
						__z.value.str = __s
						if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
							__z.u1.type_info = 6
						} else {
							__z.u1.type_info = 6 | 1<<0<<8
						}
						return
					} else {
						BasicGlobals.SetLocaleString(zend.ZendStringInit(retval, len_, 0))
						zend.ZendStringReleaseEx(loc, 0)
						var __z *zend.Zval = return_value
						var __s *zend.ZendString = BasicGlobals.GetLocaleString()
						__z.value.str = __s
						if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
							__z.u1.type_info = 6
						} else {
							zend.ZendGcAddref(&__s.gc)
							__z.u1.type_info = 6 | 1<<0<<8
						}
						return
					}
				} else if len_ == loc.len_ && !(memcmp(loc.val, retval, len_)) {
					var __z *zend.Zval = return_value
					var __s *zend.ZendString = loc
					__z.value.str = __s
					if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
						__z.u1.type_info = 6
					} else {
						__z.u1.type_info = 6 | 1<<0<<8
					}
					return
				}
				zend.ZendStringReleaseEx(loc, 0)
			}
			var _s *byte = retval
			var __z *zend.Zval = return_value
			var __s *zend.ZendString = zend.ZendStringInit(_s, strlen(_s), 0)
			__z.value.str = __s
			__z.u1.type_info = 6 | 1<<0<<8
			return
		}
		if loc != nil {
			zend.ZendStringReleaseEx(loc, 0)
		}
		if args[0].u1.v.type_ == 7 {
			idx++
		} else {
			if g.PreInc(&i) >= num_args {
				break
			}
		}
	}
	return_value.u1.type_info = 2
	return
}

/* }}} */

func ZifParseStr(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var arg *byte
	var arrayArg *zend.Zval = nil
	var res *byte = nil
	var arglen int
	for {
		var _flags int = 0
		var _min_num_args int = 1
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
			r.Assert(_i <= _min_num_args || _optional == 1)
			r.Assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgString(_arg, &arg, &arglen, 0) == 0 {
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

			zend.ZendParseArgZvalDeref(_arg, &arrayArg, 0)
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
	res = zend._estrndup(arg, arglen)
	if arrayArg == nil {
		var tmp zend.Zval
		var symbol_table *zend.ZendArray
		if zend.ZendForbidDynamicCall("parse_str() with a single argument") == zend.FAILURE {
			zend._efree(res)
			return
		}
		core.PhpErrorDocref(nil, 1<<13, "Calling parse_str() without the result argument is deprecated")
		symbol_table = zend.ZendRebuildSymbolTable()
		var __arr *zend.ZendArray = symbol_table
		var __z *zend.Zval = &tmp
		__z.value.arr = __arr
		__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
		core.sapi_module.treat_data(3, res, &tmp)
		if zend.ZendHashDel(symbol_table, zend.ZendKnownStrings[zend.ZEND_STR_THIS]) == zend.SUCCESS {
			zend.ZendThrowError(nil, "Cannot re-assign $this")
		}
	} else {
		arrayArg = zend.ZendTryArrayInit(arrayArg)
		if arrayArg == nil {
			zend._efree(res)
			return
		}
		core.sapi_module.treat_data(3, res, arrayArg)
	}
}

/* }}} */

// #define PHP_TAG_BUF_SIZE       1023

/* {{{ php_tag_find
 *
 * Check if tag is in a set of tags
 *
 * states:
 *
 * 0 start tag
 * 1 first non-whitespace char seen
 */

func PhpTagFind(tag *byte, len_ int, set *byte) int {
	var c byte
	var n *byte
	var t *byte
	var state int = 0
	var done int = 0
	var norm *byte
	if len_ == 0 {
		return 0
	}
	norm = zend._emalloc(len_ + 1)
	n = norm
	t = tag
	c = tolower(*t)

	/*
	   normalize the tag removing leading and trailing whitespace
	   and turn any <a whatever...> into just <a> and any </tag>
	   into <tag>
	*/

	for done == 0 {
		switch c {
		case '<':
			*(g.PostInc(&n)) = c
			break
		case '>':
			done = 1
			break
		default:
			if !(isspace(int(c))) {
				if state == 0 {
					state = 1
				}
				if c != '/' || (*(t - 1)) != '<' && (*(t + 1)) != '>' {
					*(g.PostInc(&n)) = c
				}
			} else {
				if state == 1 {
					done = 1
				}
			}
			break
		}
		c = tolower(*(g.PreInc(&t)))
	}
	*(g.PostInc(&n)) = '>'
	*n = '0'
	if strstr(set, norm) {
		done = 1
	} else {
		done = 0
	}
	zend._efree(norm)
	return done
}

/* }}} */

func PhpStripTags(rbuf *byte, len_ int, stateptr *uint8, allow *byte, allow_len int) int {
	return PhpStripTagsEx(rbuf, len_, stateptr, allow, allow_len, 0)
}

/* }}} */

func PhpStripTagsEx(rbuf *byte, len_ int, stateptr *uint8, allow *byte, allow_len int, allow_tag_spaces zend.ZendBool) int {
	var tbuf *byte
	var tp *byte
	var rp *byte
	var c byte
	var lc byte
	var buf *byte
	var p *byte
	var end *byte
	var br int
	var depth int = 0
	var in_q int = 0
	var state uint8 = 0
	var pos int
	var allow_free *byte = nil
	var is_xml byte = 0
	buf = zend._estrndup(rbuf, len_)
	end = buf + len_
	lc = '0'
	p = buf
	rp = rbuf
	br = 0
	if allow != nil {
		allow_free = zend.ZendStrTolowerDupEx(allow, allow_len)
		if allow_free != nil {
			allow = allow_free
		} else {
			allow = allow
		}
		tbuf = zend._emalloc(1023 + 1)
		tp = tbuf
	} else {
		tp = nil
		tbuf = tp
	}
	if stateptr != nil {
		state = *stateptr
		switch state {
		case 1:
			goto state_1
		case 2:
			goto state_2
		case 3:
			goto state_3
		case 4:
			goto state_4
		default:
			break
		}
	}
state_0:
	if p >= end {
		goto finish
	}
	c = *p
	switch c {
	case '0':
		break
	case '<':
		if in_q != 0 {
			break
		}
		if isspace(*(p + 1)) && allow_tag_spaces == 0 {
			*(g.PostInc(&rp)) = c
			break
		}
		lc = '<'
		state = 1
		if allow != nil {
			if tp-tbuf >= 1023 {
				pos = tp - tbuf
				tbuf = zend._erealloc(tbuf, tp-tbuf+1023+1)
				tp = tbuf + pos
			}
			*(g.PostInc(&tp)) = '<'
		}
		p++
		goto state_1
	case '>':
		if depth != 0 {
			depth--
			break
		}
		if in_q != 0 {
			break
		}
		*(g.PostInc(&rp)) = c
		break
	default:
		*(g.PostInc(&rp)) = c
		break
	}
	p++
	goto state_0
state_1:
	if p >= end {
		goto finish
	}
	c = *p
	switch c {
	case '0':
		break
	case '<':
		if in_q != 0 {
			break
		}
		if isspace(*(p + 1)) && allow_tag_spaces == 0 {
			goto reg_char_1
		}
		depth++
		break
	case '>':
		if depth != 0 {
			depth--
			break
		}
		if in_q != 0 {
			break
		}
		lc = '>'
		if is_xml && p >= buf+1 && (*(p - 1)) == '-' {
			break
		}
		is_xml = 0
		state = is_xml
		in_q = state
		if allow != nil {
			if tp-tbuf >= 1023 {
				pos = tp - tbuf
				tbuf = zend._erealloc(tbuf, tp-tbuf+1023+1)
				tp = tbuf + pos
			}
			*(g.PostInc(&tp)) = '>'
			*tp = '0'
			if PhpTagFind(tbuf, tp-tbuf, allow) != 0 {
				memcpy(rp, tbuf, tp-tbuf)
				rp += tp - tbuf
			}
			tp = tbuf
		}
		p++
		goto state_0
	case '"':

	case '\'':
		if p != buf && (in_q == 0 || (*p) == in_q) {
			if in_q != 0 {
				in_q = 0
			} else {
				in_q = *p
			}
		}
		goto reg_char_1
	case '!':

		/* JavaScript & Other HTML scripting languages */

		if p >= buf+1 && (*(p - 1)) == '<' {
			state = 3
			lc = c
			p++
			goto state_3
		} else {
			goto reg_char_1
		}
		break
	case '?':
		if p >= buf+1 && (*(p - 1)) == '<' {
			br = 0
			state = 2
			p++
			goto state_2
		} else {
			goto reg_char_1
		}
		break
	default:
	reg_char_1:
		if allow != nil {
			if tp-tbuf >= 1023 {
				pos = tp - tbuf
				tbuf = zend._erealloc(tbuf, tp-tbuf+1023+1)
				tp = tbuf + pos
			}
			*(g.PostInc(&tp)) = c
		}
		break
	}
	p++
	goto state_1
state_2:
	if p >= end {
		goto finish
	}
	c = *p
	switch c {
	case '(':
		if lc != '"' && lc != '\'' {
			lc = '('
			br++
		}
		break
	case ')':
		if lc != '"' && lc != '\'' {
			lc = ')'
			br--
		}
		break
	case '>':
		if depth != 0 {
			depth--
			break
		}
		if in_q != 0 {
			break
		}
		if br == 0 && p >= buf+1 && lc != '"' && (*(p - 1)) == '?' {
			state = 0
			in_q = state
			tp = tbuf
			p++
			goto state_0
		}
		break
	case '"':

	case '\'':
		if p >= buf+1 && (*(p - 1)) != '\\' {
			if lc == c {
				lc = '0'
			} else if lc != '\\' {
				lc = c
			}
			if p != buf && (in_q == 0 || (*p) == in_q) {
				if in_q != 0 {
					in_q = 0
				} else {
					in_q = *p
				}
			}
		}
		break
	case 'l':

	case 'L':

		/* swm: If we encounter '<?xml' then we shouldn't be in
		 * state == 2 (PHP). Switch back to HTML.
		 */

		if state == 2 && p > buf+4 && ((*(p - 1)) == 'm' || (*(p - 1)) == 'M') && ((*(p - 2)) == 'x' || (*(p - 2)) == 'X') && (*(p - 3)) == '?' && (*(p - 4)) == '<' {
			state = 1
			is_xml = 1
			p++
			goto state_1
		}
		break
	default:
		break
	}
	p++
	goto state_2
state_3:
	if p >= end {
		goto finish
	}
	c = *p
	switch c {
	case '>':
		if depth != 0 {
			depth--
			break
		}
		if in_q != 0 {
			break
		}
		state = 0
		in_q = state
		tp = tbuf
		p++
		goto state_0
	case '"':

	case '\'':
		if p != buf && (*(p - 1)) != '\\' && (in_q == 0 || (*p) == in_q) {
			if in_q != 0 {
				in_q = 0
			} else {
				in_q = *p
			}
		}
		break
	case '-':
		if p >= buf+2 && (*(p - 1)) == '-' && (*(p - 2)) == '!' {
			state = 4
			p++
			goto state_4
		}
		break
	case 'E':

	case 'e':

		/* !DOCTYPE exception */

		if p > buf+6 && ((*(p - 1)) == 'p' || (*(p - 1)) == 'P') && ((*(p - 2)) == 'y' || (*(p - 2)) == 'Y') && ((*(p - 3)) == 't' || (*(p - 3)) == 'T') && ((*(p - 4)) == 'c' || (*(p - 4)) == 'C') && ((*(p - 5)) == 'o' || (*(p - 5)) == 'O') && ((*(p - 6)) == 'd' || (*(p - 6)) == 'D') {
			state = 1
			p++
			goto state_1
		}
		break
	default:
		break
	}
	p++
	goto state_3
state_4:
	for p < end {
		c = *p
		if c == '>' && in_q == 0 {
			if p >= buf+2 && (*(p - 1)) == '-' && (*(p - 2)) == '-' {
				state = 0
				in_q = state
				tp = tbuf
				p++
				goto state_0
			}
		}
		p++
	}
finish:
	if rp < rbuf+len_ {
		*rp = '0'
	}
	zend._efree(any(buf))
	if tbuf != nil {
		zend._efree(tbuf)
	}
	if allow_free != nil {
		zend._efree(allow_free)
	}
	if stateptr != nil {
		*stateptr = state
	}
	return size_t(rp - rbuf)
}

/* }}} */

func ZifStrGetcsv(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var str *zend.ZendString
	var delim byte = ','
	var enc byte = '"'
	var esc int = uint8('\\')
	var delim_str *byte = nil
	var enc_str *byte = nil
	var esc_str *byte = nil
	var delim_len int = 0
	var enc_len int = 0
	var esc_len int = 0
	for {
		var _flags int = 0
		var _min_num_args int = 1
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
			r.Assert(_i <= _min_num_args || _optional == 1)
			r.Assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgStr(_arg, &str, 0) == 0 {
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

			if zend.ZendParseArgString(_arg, &delim_str, &delim_len, 0) == 0 {
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

			if zend.ZendParseArgString(_arg, &enc_str, &enc_len, 0) == 0 {
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

			if zend.ZendParseArgString(_arg, &esc_str, &esc_len, 0) == 0 {
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
	if delim_len != 0 {
		delim = delim_str[0]
	} else {
		delim = delim
	}
	if enc_len != 0 {
		enc = enc_str[0]
	} else {
		enc = enc
	}
	if esc_str != nil {
		if esc_len != 0 {
			esc = uint8(esc_str[0])
		} else {
			esc = -1
		}
	}
	PhpFgetcsv(nil, delim, enc, esc, str.len_, str.val, return_value)
}

/* }}} */

func ZifStrRepeat(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var input_str *zend.ZendString
	var mult zend.ZendLong
	var result *zend.ZendString
	var result_len int
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
			r.Assert(_i <= _min_num_args || _optional == 1)
			r.Assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgStr(_arg, &input_str, 0) == 0 {
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

			if zend.ZendParseArgLong(_arg, &mult, &_dummy, 0, 0) == 0 {
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
	if mult < 0 {
		core.PhpErrorDocref(nil, 1<<1, "Second argument has to be greater than or equal to 0")
		return
	}

	/* Don't waste our time if it's empty */

	if input_str.len_ == 0 || mult == 0 {
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = zend.ZendEmptyString
		__z.value.str = __s
		__z.u1.type_info = 6
		return
	}

	/* Initialize the result string */

	result = zend.ZendStringSafeAlloc(input_str.len_, mult, 0, 0)
	result_len = input_str.len_ * mult

	/* Heavy optimization for situations where input string is 1 byte long */

	if input_str.len_ == 1 {
		memset(result.val, (*input_str).val, mult)
	} else {
		var s *byte
		var ee *byte
		var e *byte
		var l ptrdiff_t = 0
		memcpy(result.val, input_str.val, input_str.len_)
		s = result.val
		e = result.val + input_str.len_
		ee = result.val + result_len
		for e < ee {
			if e-s < ee-e {
				l = e - s
			} else {
				l = ee - e
			}
			memmove(e, s, l)
			e += l
		}
	}
	result.val[result_len] = '0'
	var __z *zend.Zval = return_value
	var __s *zend.ZendString = result
	__z.value.str = __s
	__z.u1.type_info = 6 | 1<<0<<8
	return
}

/* }}} */

func ZifCountChars(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var input *zend.ZendString
	var chars []int
	var mymode zend.ZendLong = 0
	var buf *uint8
	var inx int
	var retstr []byte
	var retlen int = 0
	var tmp int = 0
	for {
		var _flags int = 0
		var _min_num_args int = 1
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
			r.Assert(_i <= _min_num_args || _optional == 1)
			r.Assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgStr(_arg, &input, 0) == 0 {
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

			if zend.ZendParseArgLong(_arg, &mymode, &_dummy, 0, 0) == 0 {
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
	if mymode < 0 || mymode > 4 {
		core.PhpErrorDocref(nil, 1<<1, "Unknown mode")
		return_value.u1.type_info = 2
		return
	}
	buf = (*uint8)(input.val)
	memset(any(chars), 0, g.SizeOf("chars"))
	for tmp < input.len_ {
		chars[*buf]++
		buf++
		tmp++
	}
	if mymode < 3 {
		var __arr *zend.ZendArray = zend._zendNewArray(0)
		var __z *zend.Zval = return_value
		__z.value.arr = __arr
		__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	}
	for inx = 0; inx < 256; inx++ {
		switch mymode {
		case 0:
			zend.AddIndexLong(return_value, inx, chars[inx])
			break
		case 1:
			if chars[inx] != 0 {
				zend.AddIndexLong(return_value, inx, chars[inx])
			}
			break
		case 2:
			if chars[inx] == 0 {
				zend.AddIndexLong(return_value, inx, chars[inx])
			}
			break
		case 3:
			if chars[inx] != 0 {
				retstr[g.PostInc(&retlen)] = inx
			}
			break
		case 4:
			if chars[inx] == 0 {
				retstr[g.PostInc(&retlen)] = inx
			}
			break
		}
	}
	if mymode >= 3 && mymode <= 4 {
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = zend.ZendStringInit(retstr, retlen, 0)
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
		return
	}
}

/* }}} */

func PhpStrnatcmp(execute_data *zend.ZendExecuteData, return_value *zend.Zval, fold_case int) {
	var s1 *zend.ZendString
	var s2 *zend.ZendString
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
			r.Assert(_i <= _min_num_args || _optional == 1)
			r.Assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgStr(_arg, &s1, 0) == 0 {
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

			if zend.ZendParseArgStr(_arg, &s2, 0) == 0 {
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
	var __z *zend.Zval = return_value
	__z.value.lval = StrnatcmpEx(s1.val, s1.len_, s2.val, s2.len_, fold_case)
	__z.u1.type_info = 4
	return
}

/* }}} */

func StringNaturalCompareFunctionEx(result *zend.Zval, op1 *zend.Zval, op2 *zend.Zval, case_insensitive zend.ZendBool) int {
	var tmp_str1 *zend.ZendString
	var tmp_str2 *zend.ZendString
	var str1 *zend.ZendString = zend.ZvalGetTmpString(op1, &tmp_str1)
	var str2 *zend.ZendString = zend.ZvalGetTmpString(op2, &tmp_str2)
	var __z *zend.Zval = result
	__z.value.lval = StrnatcmpEx(str1.val, str1.len_, str2.val, str2.len_, case_insensitive)
	__z.u1.type_info = 4
	zend.ZendTmpStringRelease(tmp_str1)
	zend.ZendTmpStringRelease(tmp_str2)
	return zend.SUCCESS
}

/* }}} */

func StringNaturalCaseCompareFunction(result *zend.Zval, op1 *zend.Zval, op2 *zend.Zval) int {
	return StringNaturalCompareFunctionEx(result, op1, op2, 1)
}

/* }}} */

func StringNaturalCompareFunction(result *zend.Zval, op1 *zend.Zval, op2 *zend.Zval) int {
	return StringNaturalCompareFunctionEx(result, op1, op2, 0)
}

/* }}} */

func ZifStrnatcmp(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpStrnatcmp(execute_data, return_value, 0)
}

/* }}} */

func ZifLocaleconv(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var grouping zend.Zval
	var mon_grouping zend.Zval
	var len_ int
	var i int

	/* We don't need no stinkin' parameters... */

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
	var __arr *zend.ZendArray = zend._zendNewArray(0)
	var __z *zend.Zval = &grouping
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	var __arr *zend.ZendArray = zend._zendNewArray(0)
	var __z *zend.Zval = &mon_grouping
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	var currlocdata __struct__lconv
	LocaleconvR(&currlocdata)

	/* Grab the grouping data out of the array */

	len_ = int(strlen(currlocdata.grouping))
	for i = 0; i < len_; i++ {
		zend.AddIndexLong(&grouping, i, currlocdata.grouping[i])
	}

	/* Grab the monetary grouping data out of the array */

	len_ = int(strlen(currlocdata.mon_grouping))
	for i = 0; i < len_; i++ {
		zend.AddIndexLong(&mon_grouping, i, currlocdata.mon_grouping[i])
	}
	zend.AddAssocStringEx(return_value, "decimal_point", strlen("decimal_point"), currlocdata.decimal_point)
	zend.AddAssocStringEx(return_value, "thousands_sep", strlen("thousands_sep"), currlocdata.thousands_sep)
	zend.AddAssocStringEx(return_value, "int_curr_symbol", strlen("int_curr_symbol"), currlocdata.int_curr_symbol)
	zend.AddAssocStringEx(return_value, "currency_symbol", strlen("currency_symbol"), currlocdata.currency_symbol)
	zend.AddAssocStringEx(return_value, "mon_decimal_point", strlen("mon_decimal_point"), currlocdata.mon_decimal_point)
	zend.AddAssocStringEx(return_value, "mon_thousands_sep", strlen("mon_thousands_sep"), currlocdata.mon_thousands_sep)
	zend.AddAssocStringEx(return_value, "positive_sign", strlen("positive_sign"), currlocdata.positive_sign)
	zend.AddAssocStringEx(return_value, "negative_sign", strlen("negative_sign"), currlocdata.negative_sign)
	zend.AddAssocLongEx(return_value, "int_frac_digits", strlen("int_frac_digits"), currlocdata.int_frac_digits)
	zend.AddAssocLongEx(return_value, "frac_digits", strlen("frac_digits"), currlocdata.frac_digits)
	zend.AddAssocLongEx(return_value, "p_cs_precedes", strlen("p_cs_precedes"), currlocdata.p_cs_precedes)
	zend.AddAssocLongEx(return_value, "p_sep_by_space", strlen("p_sep_by_space"), currlocdata.p_sep_by_space)
	zend.AddAssocLongEx(return_value, "n_cs_precedes", strlen("n_cs_precedes"), currlocdata.n_cs_precedes)
	zend.AddAssocLongEx(return_value, "n_sep_by_space", strlen("n_sep_by_space"), currlocdata.n_sep_by_space)
	zend.AddAssocLongEx(return_value, "p_sign_posn", strlen("p_sign_posn"), currlocdata.p_sign_posn)
	zend.AddAssocLongEx(return_value, "n_sign_posn", strlen("n_sign_posn"), currlocdata.n_sign_posn)
	zend.ZendHashStrUpdate(return_value.value.arr, "grouping", g.SizeOf("\"grouping\"")-1, &grouping)
	zend.ZendHashStrUpdate(return_value.value.arr, "mon_grouping", g.SizeOf("\"mon_grouping\"")-1, &mon_grouping)
}

/* }}} */

func ZifStrnatcasecmp(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	PhpStrnatcmp(execute_data, return_value, 1)
}

/* }}} */

func ZifSubstrCount(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var haystack *byte
	var needle *byte
	var offset zend.ZendLong = 0
	var length zend.ZendLong = 0
	var ac int = execute_data.This.u2.num_args
	var count zend.ZendLong = 0
	var haystack_len int
	var needle_len int
	var p *byte
	var endp *byte
	var cmp byte
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
			r.Assert(_i <= _min_num_args || _optional == 1)
			r.Assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgString(_arg, &haystack, &haystack_len, 0) == 0 {
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

			if zend.ZendParseArgString(_arg, &needle, &needle_len, 0) == 0 {
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

			if zend.ZendParseArgLong(_arg, &offset, &_dummy, 0, 0) == 0 {
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

			if zend.ZendParseArgLong(_arg, &length, &_dummy, 0, 0) == 0 {
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
	if needle_len == 0 {
		core.PhpErrorDocref(nil, 1<<1, "Empty substring")
		return_value.u1.type_info = 2
		return
	}
	p = haystack
	endp = p + haystack_len
	if offset < 0 {
		offset += zend.ZendLong(haystack_len)
	}
	if offset < 0 || int(offset > haystack_len) != 0 {
		core.PhpErrorDocref(nil, 1<<1, "Offset not contained in string")
		return_value.u1.type_info = 2
		return
	}
	p += offset
	if ac == 4 {
		if length < 0 {
			length += haystack_len - offset
		}
		if length < 0 || int(length > haystack_len-offset) != 0 {
			core.PhpErrorDocref(nil, 1<<1, "Invalid length value")
			return_value.u1.type_info = 2
			return
		}
		endp = p + length
	}
	if needle_len == 1 {
		cmp = needle[0]
		for g.Assign(&p, memchr(p, cmp, endp-p)) {
			count++
			p++
		}
	} else {
		for g.Assign(&p, (*byte)(zend.ZendMemnstr(p, needle, needle_len, endp))) {
			p += needle_len
			count++
		}
	}
	var __z *zend.Zval = return_value
	__z.value.lval = count
	__z.u1.type_info = 4
	return
}

/* }}} */

func ZifStrPad(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	/* Input arguments */

	var input *zend.ZendString
	var pad_length zend.ZendLong

	/* Helper variables */

	var num_pad_chars int
	var pad_str *byte = " "
	var pad_str_len int = 1
	var pad_type_val zend.ZendLong = 1
	var i int
	var left_pad int = 0
	var right_pad int = 0
	var result *zend.ZendString = nil
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
			r.Assert(_i <= _min_num_args || _optional == 1)
			r.Assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgStr(_arg, &input, 0) == 0 {
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

			if zend.ZendParseArgLong(_arg, &pad_length, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
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

			if zend.ZendParseArgString(_arg, &pad_str, &pad_str_len, 0) == 0 {
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

			if zend.ZendParseArgLong(_arg, &pad_type_val, &_dummy, 0, 0) == 0 {
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

	/* If resulting string turns out to be shorter than input string,
	   we simply copy the input and return. */

	if pad_length < 0 || int(pad_length <= input.len_) != 0 {
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = input
		__z.value.str = __s
		if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
			__z.u1.type_info = 6
		} else {
			zend.ZendGcAddref(&__s.gc)
			__z.u1.type_info = 6 | 1<<0<<8
		}
		return
	}
	if pad_str_len == 0 {
		core.PhpErrorDocref(nil, 1<<1, "Padding string cannot be empty")
		return
	}
	if pad_type_val < 0 || pad_type_val > 2 {
		core.PhpErrorDocref(nil, 1<<1, "Padding type has to be STR_PAD_LEFT, STR_PAD_RIGHT, or STR_PAD_BOTH")
		return
	}
	num_pad_chars = pad_length - input.len_
	if num_pad_chars >= 2147483647 {
		core.PhpErrorDocref(nil, 1<<1, "Padding length is too long")
		return
	}
	result = zend.ZendStringSafeAlloc(1, input.len_, num_pad_chars, 0)
	result.len_ = 0

	/* We need to figure out the left/right padding lengths. */

	switch pad_type_val {
	case 1:
		left_pad = 0
		right_pad = num_pad_chars
		break
	case 0:
		left_pad = num_pad_chars
		right_pad = 0
		break
	case 2:
		left_pad = num_pad_chars / 2
		right_pad = num_pad_chars - left_pad
		break
	}

	/* First we pad on the left. */

	for i = 0; i < left_pad; i++ {
		result.val[g.PostInc(&(result.len_))] = pad_str[i%pad_str_len]
	}

	/* Then we copy the input string. */

	memcpy(result.val+result.len_, input.val, input.len_)
	result.len_ += input.len_

	/* Finally, we pad on the right. */

	for i = 0; i < right_pad; i++ {
		result.val[g.PostInc(&(result.len_))] = pad_str[i%pad_str_len]
	}
	result.val[result.len_] = '0'
	var __z *zend.Zval = return_value
	var __s *zend.ZendString = result
	__z.value.str = __s
	__z.u1.type_info = 6 | 1<<0<<8
	return
}

/* }}} */

func ZifSscanf(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var args *zend.Zval = nil
	var str *byte
	var format *byte
	var str_len int
	var format_len int
	var result int
	var num_args int = 0
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = -1
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

			if zend.ZendParseArgString(_arg, &str, &str_len, 0) == 0 {
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

			if zend.ZendParseArgString(_arg, &format, &format_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = 4
				break
			}
			var _num_varargs int = _num_args - _i - 0
			if _num_varargs > 0 {
				args = _real_arg + 1
				num_args = _num_varargs
				_i += _num_varargs
				_real_arg += _num_varargs
			} else {
				args = nil
				num_args = 0
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
	result = PhpSscanfInternal(str, format, num_args, args, 0, return_value)
	if -1-1-1-1 == result {
		zend.ZendWrongParamCount()
		return
	}
}

/* }}} */

func PhpStrRot13(str *zend.ZendString) *zend.ZendString {
	var ret *zend.ZendString
	var p *byte
	var e *byte
	var target *byte
	if str.len_ == 0 {
		return zend.ZendEmptyString
	}
	ret = zend.ZendStringAlloc(str.len_, 0)
	p = str.val
	e = p + str.len_
	target = ret.val
	for p < e {
		if (*p) >= 'a' && (*p) <= 'z' {
			g.PostInc(&(*target)) = 'a' + (g.PostInc(&(*p))-'a'+13)%26
		} else if (*p) >= 'A' && (*p) <= 'Z' {
			g.PostInc(&(*target)) = 'A' + (g.PostInc(&(*p))-'A'+13)%26
		} else {
			*p++
			g.PostInc(&(*target)) = (*p) - 1
		}
	}
	*target = '0'
	return ret
}

/* }}} */

func ZifStrRot13(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var arg *zend.ZendString
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
			r.Assert(_i <= _min_num_args || _optional == 1)
			r.Assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgStr(_arg, &arg, 0) == 0 {
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
	var __z *zend.Zval = return_value
	var __s *zend.ZendString = PhpStrRot13(arg)
	__z.value.str = __s
	if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
		__z.u1.type_info = 6
	} else {
		__z.u1.type_info = 6 | 1<<0<<8
	}
	return
}

/* }}} */

func PhpStringShuffle(str *byte, len_ zend.ZendLong) {
	var n_elems zend.ZendLong
	var rnd_idx zend.ZendLong
	var n_left zend.ZendLong
	var temp byte

	/* The implementation is stolen from array_data_shuffle       */

	n_elems = len_
	if n_elems <= 1 {
		return
	}
	n_left = n_elems
	for g.PreDec(&n_left) {
		rnd_idx = PhpMtRandRange(0, n_left)
		if rnd_idx != n_left {
			temp = str[n_left]
			str[n_left] = str[rnd_idx]
			str[rnd_idx] = temp
		}
	}
}

/* }}} */

func ZifStrShuffle(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var arg *zend.ZendString
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
			r.Assert(_i <= _min_num_args || _optional == 1)
			r.Assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgStr(_arg, &arg, 0) == 0 {
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
	var __z *zend.Zval = return_value
	var __s *zend.ZendString = zend.ZendStringInit(arg.val, arg.len_, 0)
	__z.value.str = __s
	__z.u1.type_info = 6 | 1<<0<<8
	if return_value.value.str.len_ > 1 {
		PhpStringShuffle(return_value.value.str.val, zend_long(return_value.value.str).len_)
	}
}

/* }}} */

func ZifStrWordCount(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var str *zend.ZendString
	var char_list *byte = nil
	var ch []*byte
	var p *byte
	var e *byte
	var s *byte
	var char_list_len int = 0
	var word_count int = 0
	var type_ zend.ZendLong = 0
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

			if zend.ZendParseArgStr(_arg, &str, 0) == 0 {
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

			if zend.ZendParseArgLong(_arg, &type_, &_dummy, 0, 0) == 0 {
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

			if zend.ZendParseArgString(_arg, &char_list, &char_list_len, 0) == 0 {
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
	switch type_ {
	case 1:

	case 2:
		var __arr *zend.ZendArray = zend._zendNewArray(0)
		var __z *zend.Zval = return_value
		__z.value.arr = __arr
		__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
		if str.len_ == 0 {
			return
		}
		break
	case 0:
		if str.len_ == 0 {
			var __z *zend.Zval = return_value
			__z.value.lval = 0
			__z.u1.type_info = 4
			return
		}

		/* nothing to be done */

		break
	default:
		core.PhpErrorDocref(nil, 1<<1, "Invalid format value "+"%"+"lld", type_)
		return_value.u1.type_info = 2
		return
	}
	if char_list != nil {
		PhpCharmask((*uint8)(char_list), char_list_len, ch)
	}
	p = str.val
	e = str.val + str.len_

	/* first character cannot be ' or -, unless explicitly allowed by the user */

	if (*p) == '\'' && (char_list == nil || ch['\''] == nil) || (*p) == '-' && (char_list == nil || ch['-'] == nil) {
		p++
	}

	/* last character cannot be -, unless explicitly allowed by the user */

	if (*(e - 1)) == '-' && (char_list == nil || ch['-'] == nil) {
		e--
	}
	for p < e {
		s = p
		for p < e && (isalpha(uint8(*p)) || char_list != nil && ch[uint8(*p)] != nil || (*p) == '\'' || (*p) == '-') {
			p++
		}
		if p > s {
			switch type_ {
			case 1:
				zend.AddNextIndexStringl(return_value, s, p-s)
				break
			case 2:
				zend.AddIndexStringl(return_value, s-str.val, s, p-s)
				break
			default:
				word_count++
				break
			}
		}
		p++
	}
	if type_ == 0 {
		var __z *zend.Zval = return_value
		__z.value.lval = word_count
		__z.u1.type_info = 4
		return
	}
}

/* }}} */

/* {{{ proto string money_format(string format , float value)
   Convert monetary value(s) to string */

func ZifMoneyFormat(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var format_len int = 0
	var format *byte
	var p *byte
	var e *byte
	var value float64
	var check zend.ZendBool = 0
	var str *zend.ZendString
	var res_len ssize_t
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
			r.Assert(_i <= _min_num_args || _optional == 1)
			r.Assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgString(_arg, &format, &format_len, 0) == 0 {
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

			if zend.ZendParseArgDouble(_arg, &value, &_dummy, 0) == 0 {
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
			return
		}
		break
	}
	p = format
	e = p + format_len
	for g.Assign(&p, memchr(p, '%', e-p)) {
		if (*(p + 1)) == '%' {
			p += 2
		} else if check == 0 {
			check = 1
			p++
		} else {
			core.PhpErrorDocref(nil, 1<<1, "Only a single %%i or %%n token can be used")
			return_value.u1.type_info = 2
			return
		}
	}
	str = zend.ZendStringSafeAlloc(format_len, 1, 1024, 0)
	if g.Assign(&res_len, strfmon(str.val, str.len_, format, value)) < 0 {
		zend.ZendStringEfree(str)
		return_value.u1.type_info = 2
		return
	}
	str.len_ = int(res_len)
	str.val[str.len_] = '0'
	var __z *zend.Zval = return_value
	var __s *zend.ZendString = zend.ZendStringTruncate(str, str.len_, 0)
	__z.value.str = __s
	__z.u1.type_info = 6 | 1<<0<<8
	return
}

/* }}} */

/* {{{ proto array str_split(string str [, int split_length])
   Convert a string to an array. If split_length is specified, break the string down into chunks each split_length characters long. */

func ZifStrSplit(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var str *zend.ZendString
	var split_length zend.ZendLong = 1
	var p *byte
	var n_reg_segments int
	for {
		var _flags int = 0
		var _min_num_args int = 1
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
			r.Assert(_i <= _min_num_args || _optional == 1)
			r.Assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgStr(_arg, &str, 0) == 0 {
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

			if zend.ZendParseArgLong(_arg, &split_length, &_dummy, 0, 0) == 0 {
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
	if split_length <= 0 {
		core.PhpErrorDocref(nil, 1<<1, "The length of each segment must be greater than zero")
		return_value.u1.type_info = 2
		return
	}
	if 0 == str.len_ || int(split_length >= str.len_) != 0 {
		var __arr *zend.ZendArray = zend._zendNewArray(1)
		var __z *zend.Zval = return_value
		__z.value.arr = __arr
		__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
		zend.AddNextIndexStringl(return_value, str.val, str.len_)
		return
	}
	var __arr *zend.ZendArray = zend._zendNewArray(uint32((str.len_-1)/split_length + 1))
	var __z *zend.Zval = return_value
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	n_reg_segments = str.len_ / split_length
	p = str.val
	for g.PostDec(&n_reg_segments) > 0 {
		zend.AddNextIndexStringl(return_value, p, split_length)
		p += split_length
	}
	if p != str.val+str.len_ {
		zend.AddNextIndexStringl(return_value, p, str.val+str.len_-p)
	}
}

/* }}} */

func ZifStrpbrk(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var haystack *zend.ZendString
	var char_list *zend.ZendString
	var haystack_ptr *byte
	var cl_ptr *byte
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
			r.Assert(_i <= _min_num_args || _optional == 1)
			r.Assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgStr(_arg, &haystack, 0) == 0 {
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

			if zend.ZendParseArgStr(_arg, &char_list, 0) == 0 {
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
	if char_list.len_ == 0 {
		core.PhpErrorDocref(nil, 1<<1, "The character list cannot be empty")
		return_value.u1.type_info = 2
		return
	}
	for haystack_ptr = haystack.val; haystack_ptr < haystack.val+haystack.len_; haystack_ptr++ {
		for cl_ptr = char_list.val; cl_ptr < char_list.val+char_list.len_; cl_ptr++ {
			if (*cl_ptr) == (*haystack_ptr) {
				var __z *zend.Zval = return_value
				var __s *zend.ZendString = zend.ZendStringInit(haystack_ptr, haystack.val+haystack.len_-haystack_ptr, 0)
				__z.value.str = __s
				__z.u1.type_info = 6 | 1<<0<<8
				return
			}
		}
	}
	return_value.u1.type_info = 2
	return
}

/* }}} */

func ZifSubstrCompare(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var s1 *zend.ZendString
	var s2 *zend.ZendString
	var offset zend.ZendLong
	var len_ zend.ZendLong = 0
	var len_is_default zend.ZendBool = 1
	var cs zend.ZendBool = 0
	var cmp_len int
	for {
		var _flags int = 0
		var _min_num_args int = 3
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

			if zend.ZendParseArgStr(_arg, &s1, 0) == 0 {
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

			if zend.ZendParseArgStr(_arg, &s2, 0) == 0 {
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

			if zend.ZendParseArgLong(_arg, &offset, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
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

			if zend.ZendParseArgLong(_arg, &len_, &len_is_default, 1, 0) == 0 {
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

			if zend.ZendParseArgBool(_arg, &cs, &_dummy, 0) == 0 {
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
	if len_is_default == 0 && len_ <= 0 {
		if len_ == 0 {
			var __z *zend.Zval = return_value
			__z.value.lval = 0
			__z.u1.type_info = 4
			return
		} else {
			core.PhpErrorDocref(nil, 1<<1, "The length must be greater than or equal to zero")
			return_value.u1.type_info = 2
			return
		}
	}
	if offset < 0 {
		offset = s1.len_ + offset
		if offset < 0 {
			offset = 0
		} else {
			offset = offset
		}
	}
	if int(offset > s1.len_) != 0 {
		core.PhpErrorDocref(nil, 1<<1, "The start position cannot exceed initial string length")
		return_value.u1.type_info = 2
		return
	}
	if len_ != 0 {
		cmp_len = int(len_)
	} else {
		if s2.len_ > s1.len_-offset {
			cmp_len = s2.len_
		} else {
			cmp_len = s1.len_ - offset
		}
	}
	if cs == 0 {
		var __z *zend.Zval = return_value
		__z.value.lval = zend.ZendBinaryStrncmp(s1.val+offset, s1.len_-offset, s2.val, s2.len_, cmp_len)
		__z.u1.type_info = 4
		return
	} else {
		var __z *zend.Zval = return_value
		__z.value.lval = zend.ZendBinaryStrncasecmpL(s1.val+offset, s1.len_-offset, s2.val, s2.len_, cmp_len)
		__z.u1.type_info = 4
		return
	}
}

/* }}} */

func PhpUtf8Encode(s *byte, len_ int) *zend.ZendString {
	var pos int = len_
	var str *zend.ZendString
	var c uint8
	str = zend.ZendStringSafeAlloc(len_, 2, 0, 0)
	str.len_ = 0
	for pos > 0 {

		/* The lower 256 codepoints of Unicode are identical to Latin-1,
		 * so we don't need to do any mapping here. */

		c = uint8(*s)
		if c < 0x80 {
			str.val[g.PostInc(&(str.len_))] = byte(c)
		} else {
			str.val[g.PostInc(&(str.len_))] = 0xc0 | c>>6
			str.val[g.PostInc(&(str.len_))] = 0x80 | c&0x3f
		}
		pos--
		s++
	}
	str.val[str.len_] = '0'
	str = zend.ZendStringTruncate(str, str.len_, 0)
	return str
}

/* }}} */

func PhpUtf8Decode(s *byte, len_ int) *zend.ZendString {
	var pos int = 0
	var c uint
	var str *zend.ZendString
	str = zend.ZendStringAlloc(len_, 0)
	str.len_ = 0
	for pos < len_ {
		var status int = zend.FAILURE
		c = PhpNextUtf8Char((*uint8)(s), int(len_), &pos, &status)

		/* The lower 256 codepoints of Unicode are identical to Latin-1,
		 * so we don't need to do any mapping here beyond replacing non-Latin-1
		 * characters. */

		if status == zend.FAILURE || c > 0xff {
			c = '?'
		}
		str.val[g.PostInc(&(str.len_))] = c
	}
	str.val[str.len_] = '0'
	if str.len_ < len_ {
		str = zend.ZendStringTruncate(str, str.len_, 0)
	}
	return str
}

/* }}} */

func ZifUtf8Encode(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var arg *byte
	var arg_len int
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
			r.Assert(_i <= _min_num_args || _optional == 1)
			r.Assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgString(_arg, &arg, &arg_len, 0) == 0 {
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
	var __z *zend.Zval = return_value
	var __s *zend.ZendString = PhpUtf8Encode(arg, arg_len)
	__z.value.str = __s
	if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
		__z.u1.type_info = 6
	} else {
		__z.u1.type_info = 6 | 1<<0<<8
	}
	return
}

/* }}} */

func ZifUtf8Decode(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var arg *byte
	var arg_len int
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
			r.Assert(_i <= _min_num_args || _optional == 1)
			r.Assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgString(_arg, &arg, &arg_len, 0) == 0 {
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
	var __z *zend.Zval = return_value
	var __s *zend.ZendString = PhpUtf8Decode(arg, arg_len)
	__z.value.str = __s
	if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
		__z.u1.type_info = 6
	} else {
		__z.u1.type_info = 6 | 1<<0<<8
	}
	return
}

/* }}} */
