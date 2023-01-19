// <<generate>>

package standard

import (
	g "sik/runtime/grammar"
	"sik/zend"
)

// Source: <ext/standard/http.c>

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
   | Authors: Sara Golemon <pollita@php.net>                              |
   +----------------------------------------------------------------------+
*/

// # include "php_http.h"

// # include "php_ini.h"

// # include "url.h"

// #define URL_DEFAULT_ARG_SEP       "&"

/* {{{ php_url_encode_hash */

func PhpUrlEncodeHashEx(ht *zend.HashTable, formstr *zend.SmartStr, num_prefix *byte, num_prefix_len int, key_prefix *byte, key_prefix_len int, key_suffix string, key_suffix_len int, type_ *zend.Zval, arg_sep *byte, enc_type int) int {
	var key *zend.ZendString = nil
	var newprefix *byte
	var p *byte
	var prop_name *byte
	var arg_sep_len int
	var newprefix_len int
	var prop_len int
	var idx zend.ZendUlong
	var zdata *zend.Zval = nil
	if ht == nil {
		return zend.FAILURE
	}
	if (zend.ZvalGcFlags(ht.gc.u.type_info) & 1 << 5) != 0 {

		/* Prevent recursion */

		return zend.SUCCESS

		/* Prevent recursion */

	}
	if arg_sep == nil {
		arg_sep = zend.ZendIniStringEx("arg_separator.output", g.SizeOf("\"arg_separator.output\"")-1, 0, nil)
		if arg_sep == nil || !(strlen(arg_sep)) {
			arg_sep = "&"
		}
	}
	arg_sep_len = strlen(arg_sep)
	for {
		var __ht *zend.HashTable = ht
		var _p *zend.Bucket = __ht.arData
		var _end *zend.Bucket = _p + __ht.nNumUsed
		for ; _p != _end; _p++ {
			var _z *zend.Zval = &_p.val

			if _z.u1.v.type_ == 0 {
				continue
			}
			idx = _p.h
			key = _p.key
			zdata = _z
			var is_dynamic zend.ZendBool = 1
			if zdata.u1.v.type_ == 13 {
				zdata = zdata.value.zv
				if zdata.u1.v.type_ == 0 {
					continue
				}
				is_dynamic = 0
			}

			/* handling for private & protected object properties */

			if key != nil {
				prop_name = key.val
				prop_len = key.len_
				if type_ != nil && zend.ZendCheckPropertyAccess(type_.value.obj, key, is_dynamic) != zend.SUCCESS {

					/* property not visible in this scope */

					continue

					/* property not visible in this scope */

				}
				if key.val[0] == '0' && type_ != nil {
					var tmp *byte
					zend.ZendUnmanglePropertyNameEx(key, &tmp, &prop_name, &prop_len)
				} else {
					prop_name = key.val
					prop_len = key.len_
				}
			} else {
				prop_name = nil
				prop_len = 0
			}
			if zdata.u1.v.type_ == 10 {
				zdata = &(*zdata).value.ref.val
			}
			if zdata.u1.v.type_ == 7 || zdata.u1.v.type_ == 8 {
				if key != nil {
					var ekey *zend.ZendString
					if enc_type == 2 {
						ekey = PhpRawUrlEncode(prop_name, prop_len)
					} else {
						ekey = PhpUrlEncode(prop_name, prop_len)
					}
					newprefix_len = key_suffix_len + ekey.len_ + key_prefix_len + 3
					newprefix = zend._emalloc(newprefix_len + 1)
					p = newprefix
					if key_prefix != nil {
						memcpy(p, key_prefix, key_prefix_len)
						p += key_prefix_len
					}
					memcpy(p, ekey.val, ekey.len_)
					p += ekey.len_
					zend.ZendStringFree(ekey)
					if key_suffix {
						memcpy(p, key_suffix, key_suffix_len)
						p += key_suffix_len
					}
					*(g.PostInc(&p)) = '%'
					*(g.PostInc(&p)) = '5'
					*(g.PostInc(&p)) = 'B'
					*p = '0'
				} else {
					var ekey *byte
					var ekey_len int

					/* Is an integer key */

					ekey_len = zend.ZendSpprintf(&ekey, 0, "%"+"lld", idx)
					newprefix_len = key_prefix_len + num_prefix_len + ekey_len + key_suffix_len + 3
					newprefix = zend._emalloc(newprefix_len + 1)
					p = newprefix
					if key_prefix != nil {
						memcpy(p, key_prefix, key_prefix_len)
						p += key_prefix_len
					}
					if num_prefix != nil {
						memcpy(p, num_prefix, num_prefix_len)
						p += num_prefix_len
					}
					memcpy(p, ekey, ekey_len)
					p += ekey_len
					zend._efree(ekey)
					if key_suffix {
						memcpy(p, key_suffix, key_suffix_len)
						p += key_suffix_len
					}
					*(g.PostInc(&p)) = '%'
					*(g.PostInc(&p)) = '5'
					*(g.PostInc(&p)) = 'B'
					*p = '0'
				}
				if (zend.ZvalGcFlags(ht.gc.u.type_info) & 1 << 6) == 0 {
					ht.gc.u.type_info |= 1 << 5 << 0
				}
				PhpUrlEncodeHashEx(g.CondF(zdata.u1.v.type_ == 7, func() *zend.ZendArray { return zdata.value.arr }, func() __auto__ {
					if zdata.u1.v.type_ == 8 {
						return zdata.value.obj.handlers.get_properties(zdata)
					} else {
						return nil
					}
				}), formstr, nil, 0, newprefix, newprefix_len, "%5D", 3, g.Cond(zdata.u1.v.type_ == 8, zdata, nil), arg_sep, enc_type)
				if (zend.ZvalGcFlags(ht.gc.u.type_info) & 1 << 6) == 0 {
					ht.gc.u.type_info &= ^(1 << 5 << 0)
				}
				zend._efree(newprefix)
			} else if zdata.u1.v.type_ == 1 || zdata.u1.v.type_ == 9 {

				/* Skip these types */

				continue

				/* Skip these types */

			} else {
				if formstr.s != nil {
					zend.SmartStrAppendlEx(formstr, arg_sep, arg_sep_len, 0)
				}

				/* Simple key=value */

				if key_prefix != nil {
					zend.SmartStrAppendlEx(formstr, key_prefix, key_prefix_len, 0)
				}
				if key != nil {
					var ekey *zend.ZendString
					if enc_type == 2 {
						ekey = PhpRawUrlEncode(prop_name, prop_len)
					} else {
						ekey = PhpUrlEncode(prop_name, prop_len)
					}
					zend.SmartStrAppendEx(formstr, ekey, 0)
					zend.ZendStringFree(ekey)
				} else {

					/* Numeric key */

					if num_prefix != nil {
						zend.SmartStrAppendlEx(formstr, num_prefix, num_prefix_len, 0)
					}
					zend.SmartStrAppendLongEx(formstr, idx, 0)
				}
				if key_suffix {
					zend.SmartStrAppendlEx(formstr, key_suffix, key_suffix_len, 0)
				}
				zend.SmartStrAppendlEx(formstr, "=", 1, 0)
				switch zdata.u1.v.type_ {
				case 6:
					var ekey *zend.ZendString
					if enc_type == 2 {
						ekey = PhpRawUrlEncode(zdata.value.str.val, zdata.value.str.len_)
					} else {
						ekey = PhpUrlEncode(zdata.value.str.val, zdata.value.str.len_)
					}
					zend.SmartStrAppendEx(formstr, ekey, 0)
					zend.ZendStringFree(ekey)
					break
				case 4:
					zend.SmartStrAppendLongEx(formstr, zdata.value.lval, 0)
					break
				case 2:
					zend.SmartStrAppendlEx(formstr, "0", g.SizeOf("\"0\"")-1, 0)
					break
				case 3:
					zend.SmartStrAppendlEx(formstr, "1", g.SizeOf("\"1\"")-1, 0)
					break
				default:
					var ekey *zend.ZendString
					var tmp *zend.ZendString
					var str *zend.ZendString = zend.ZvalGetTmpString(zdata, &tmp)
					if enc_type == 2 {
						ekey = PhpRawUrlEncode(str.val, str.len_)
					} else {
						ekey = PhpUrlEncode(str.val, str.len_)
					}
					zend.SmartStrAppendEx(formstr, ekey, 0)
					zend.ZendTmpStringRelease(tmp)
					zend.ZendStringFree(ekey)
				}
			}
		}
		break
	}
	return zend.SUCCESS
}

/* }}} */

func ZifHttpBuildQuery(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var formdata *zend.Zval
	var prefix *byte = nil
	var arg_sep *byte = nil
	var arg_sep_len int = 0
	var prefix_len int = 0
	var formstr zend.SmartStr = zend.SmartStr{0}
	var enc_type zend.ZendLong = 1
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
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgArray(_arg, &formdata, 0, 1) == 0 {
				_expected_type = zend.Z_EXPECTED_ARRAY
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

			if zend.ZendParseArgString(_arg, &prefix, &prefix_len, 0) == 0 {
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

			if zend.ZendParseArgString(_arg, &arg_sep, &arg_sep_len, 0) == 0 {
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

			if zend.ZendParseArgLong(_arg, &enc_type, &_dummy, 0, 0) == 0 {
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
	if PhpUrlEncodeHashEx(g.CondF(formdata.u1.v.type_ == 7, func() *zend.ZendArray { return formdata.value.arr }, func() __auto__ {
		if formdata.u1.v.type_ == 8 {
			return formdata.value.obj.handlers.get_properties(formdata)
		} else {
			return nil
		}
	}), &formstr, prefix, prefix_len, nil, 0, nil, 0, g.Cond(formdata.u1.v.type_ == 8, formdata, nil), arg_sep, int(enc_type)) == zend.FAILURE {
		if formstr.s != nil {
			zend.SmartStrFreeEx(&formstr, 0)
		}
		return_value.u1.type_info = 2
		return
	}
	if formstr.s == nil {
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = zend.ZendEmptyString
		__z.value.str = __s
		__z.u1.type_info = 6
		return
	}
	zend.SmartStr0(&formstr)
	var __z *zend.Zval = return_value
	var __s *zend.ZendString = formstr.s
	__z.value.str = __s
	__z.u1.type_info = 6 | 1<<0<<8
	return
}

/* }}} */
