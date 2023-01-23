// <<generate>>

package standard

import (
	"sik/core"
	"sik/core/streams"
	r "sik/runtime"
	g "sik/runtime/grammar"
	"sik/zend"
)

// Source: <ext/standard/url.h>

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
   | Author: Jim Winstead <jimw@php.net>                                  |
   +----------------------------------------------------------------------+
*/

// #define URL_H

// @type PhpUrl struct

// #define PHP_URL_SCHEME       0

// #define PHP_URL_HOST       1

// #define PHP_URL_PORT       2

// #define PHP_URL_USER       3

// #define PHP_URL_PASS       4

// #define PHP_URL_PATH       5

// #define PHP_URL_QUERY       6

// #define PHP_URL_FRAGMENT       7

// #define PHP_QUERY_RFC1738       1

// #define PHP_QUERY_RFC3986       2

// Source: <ext/standard/url.c>

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
   | Author: Jim Winstead <jimw@php.net>                                  |
   +----------------------------------------------------------------------+
*/

// # include < stdlib . h >

// # include < string . h >

// # include < ctype . h >

// # include < sys / types . h >

// # include "php.h"

// # include "url.h"

// # include "file.h"

/* {{{ free_url
 */

func PhpUrlFree(theurl *PhpUrl) {
	if theurl.GetScheme() != nil {
		zend.ZendStringReleaseEx(theurl.GetScheme(), 0)
	}
	if theurl.GetUser() != nil {
		zend.ZendStringReleaseEx(theurl.GetUser(), 0)
	}
	if theurl.GetPass() != nil {
		zend.ZendStringReleaseEx(theurl.GetPass(), 0)
	}
	if theurl.GetHost() != nil {
		zend.ZendStringReleaseEx(theurl.GetHost(), 0)
	}
	if theurl.GetPath() != nil {
		zend.ZendStringReleaseEx(theurl.GetPath(), 0)
	}
	if theurl.GetQuery() != nil {
		zend.ZendStringReleaseEx(theurl.GetQuery(), 0)
	}
	if theurl.GetFragment() != nil {
		zend.ZendStringReleaseEx(theurl.GetFragment(), 0)
	}
	zend._efree(theurl)
}

/* }}} */

func PhpReplaceControlcharsEx(str *byte, len_ int) *byte {
	var s *uint8 = (*uint8)(str)
	var e *uint8 = (*uint8)(str + len_)
	if str == nil {
		return nil
	}
	for s < e {
		if iscntrl(*s) {
			*s = '_'
		}
		s++
	}
	return str
}

/* }}} */

func PhpReplaceControlchars(str *byte) *byte {
	return PhpReplaceControlcharsEx(str, strlen(str))
}
func PhpUrlParse(str *byte) *PhpUrl { return PhpUrlParseEx(str, strlen(str)) }
func BinaryStrcspn(s *byte, e *byte, chars string) *byte {
	for *chars {
		var p *byte = memchr(s, *chars, e-s)
		if p != nil {
			e = p
		}
		chars++
	}
	return e
}

/* {{{ php_url_parse
 */

func PhpUrlParseEx(str *byte, length int) *PhpUrl {
	var has_port zend.ZendBool
	return PhpUrlParseEx2(str, length, &has_port)
}

/* {{{ php_url_parse_ex2
 */

func PhpUrlParseEx2(str *byte, length int, has_port *zend.ZendBool) *PhpUrl {
	var port_buf []byte
	var ret *PhpUrl = zend._ecalloc(1, g.SizeOf("php_url"))
	var s *byte
	var e byte
	var p byte
	var pp byte
	var ue byte
	*has_port = 0
	s = str
	ue = s + length

	/* parse scheme */

	if g.Assign(&e, memchr(s, ':', length)) && e != s {

		/* validate scheme */

		p = s
		for p < e {

			/* scheme = 1*[ lowalpha | digit | "+" | "-" | "." ] */

			if !(isalpha(*p)) && !(isdigit(*p)) && (*p) != '+' && (*p) != '.' && (*p) != '-' {
				if e+1 < ue && e < BinaryStrcspn(s, ue, "?#") {
					goto parse_port
				} else if s+1 < ue && (*s) == '/' && (*(s + 1)) == '/' {
					s += 2
					e = 0
					goto parse_host
				} else {
					goto just_path
				}
			}
			p++
		}
		if e+1 == ue {
			ret.SetScheme(zend.ZendStringInit(s, e-s, 0))
			PhpReplaceControlcharsEx(ret.GetScheme().val, ret.GetScheme().len_)
			return ret
		}

		/*
		 * certain schemas like mailto: and zlib: may not have any / after them
		 * this check ensures we support those.
		 */

		if (*(e + 1)) != '/' {

			/* check if the data we get is a port this allows us to
			 * correctly parse things like a.com:80
			 */

			p = e + 1
			for p < ue && isdigit(*p) {
				p++
			}
			if (p == ue || (*p) == '/') && p-e < 7 {
				goto parse_port
			}
			ret.SetScheme(zend.ZendStringInit(s, e-s, 0))
			PhpReplaceControlcharsEx(ret.GetScheme().val, ret.GetScheme().len_)
			s = e + 1
			goto just_path
		} else {
			ret.SetScheme(zend.ZendStringInit(s, e-s, 0))
			PhpReplaceControlcharsEx(ret.GetScheme().val, ret.GetScheme().len_)
			if e+2 < ue && (*(e + 2)) == '/' {
				s = e + 3
				if ret.GetScheme().len_ == g.SizeOf("\"file\"")-1 && zend.ZendBinaryStrcasecmp(ret.GetScheme().val, ret.GetScheme().len_, "file", g.SizeOf("\"file\"")-1) == 0 {
					if e+3 < ue && (*(e + 3)) == '/' {

						/* support windows drive letters as in:
						   file:///c:/somedir/file.txt
						*/

						if e+5 < ue && (*(e + 5)) == ':' {
							s = e + 4
						}
						goto just_path
					}
				}
			} else {
				s = e + 1
				goto just_path
			}
		}

		/*
		 * certain schemas like mailto: and zlib: may not have any / after them
		 * this check ensures we support those.
		 */

	} else if e {
	parse_port:
		p = e + 1
		pp = p
		for pp < ue && pp-p < 6 && isdigit(*pp) {
			pp++
		}
		if pp-p > 0 && pp-p < 6 && (pp == ue || (*pp) == '/') {
			var port zend.ZendLong
			var end *byte
			memcpy(port_buf, p, pp-p)
			port_buf[pp-p] = '0'
			port = strtoll(port_buf, &end, 10)
			if port >= 0 && port <= 65535 && end != port_buf {
				*has_port = 1
				ret.SetPort(uint16(port))
				if s+1 < ue && (*s) == '/' && (*(s + 1)) == '/' {
					s += 2
				}
			} else {
				PhpUrlFree(ret)
				return nil
			}
		} else if p == pp && pp == ue {
			PhpUrlFree(ret)
			return nil
		} else if s+1 < ue && (*s) == '/' && (*(s + 1)) == '/' {
			s += 2
		} else {
			goto just_path
		}
	} else if s+1 < ue && (*s) == '/' && (*(s + 1)) == '/' {
		s += 2
	} else {
		goto just_path
	}
parse_host:
	e = BinaryStrcspn(s, ue, "/?#")

	/* check for login and password */

	if g.Assign(&p, zend.ZendMemrchr(s, '@', e-s)) {
		if g.Assign(&pp, memchr(s, ':', p-s)) {
			ret.SetUser(zend.ZendStringInit(s, pp-s, 0))
			PhpReplaceControlcharsEx(ret.GetUser().val, ret.GetUser().len_)
			pp++
			ret.SetPass(zend.ZendStringInit(pp, p-pp, 0))
			PhpReplaceControlcharsEx(ret.GetPass().val, ret.GetPass().len_)
		} else {
			ret.SetUser(zend.ZendStringInit(s, p-s, 0))
			PhpReplaceControlcharsEx(ret.GetUser().val, ret.GetUser().len_)
		}
		s = p + 1
	}

	/* check for port */

	if s < ue && (*s) == '[' && (*(e - 1)) == ']' {

		/* Short circuit portscan,
		   we're dealing with an
		   IPv6 embedded address */

		p = nil

		/* Short circuit portscan,
		   we're dealing with an
		   IPv6 embedded address */

	} else {
		p = zend.ZendMemrchr(s, ':', e-s)
	}
	if p {
		if ret.GetPort() == 0 {
			p++
			if e-p > 5 {
				PhpUrlFree(ret)
				return nil
			} else if e-p > 0 {
				var port zend.ZendLong
				var end *byte
				memcpy(port_buf, p, e-p)
				port_buf[e-p] = '0'
				port = strtoll(port_buf, &end, 10)
				if port >= 0 && port <= 65535 && end != port_buf {
					*has_port = 1
					ret.SetPort(uint16(port))
				} else {
					PhpUrlFree(ret)
					return nil
				}
			}
			p--
		}
	} else {
		p = e
	}

	/* check if we have a valid host, if we don't reject the string as url */

	if p-s < 1 {
		PhpUrlFree(ret)
		return nil
	}
	ret.SetHost(zend.ZendStringInit(s, p-s, 0))
	PhpReplaceControlcharsEx(ret.GetHost().val, ret.GetHost().len_)
	if e == ue {
		return ret
	}
	s = e
just_path:
	e = ue
	p = memchr(s, '#', e-s)
	if p {
		p++
		if p < e {
			ret.SetFragment(zend.ZendStringInit(p, e-p, 0))
			PhpReplaceControlcharsEx(ret.GetFragment().val, ret.GetFragment().len_)
		}
		e = p - 1
	}
	p = memchr(s, '?', e-s)
	if p {
		p++
		if p < e {
			ret.SetQuery(zend.ZendStringInit(p, e-p, 0))
			PhpReplaceControlcharsEx(ret.GetQuery().val, ret.GetQuery().len_)
		}
		e = p - 1
	}
	if s < e || s == ue {
		ret.SetPath(zend.ZendStringInit(s, e-s, 0))
		PhpReplaceControlcharsEx(ret.GetPath().val, ret.GetPath().len_)
	}
	return ret
}

/* }}} */

func ZifParseUrl(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var str *byte
	var str_len int
	var resource *PhpUrl
	var key zend.ZendLong = -1
	var tmp zend.Zval
	var has_port zend.ZendBool
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

			if zend.ZendParseArgLong(_arg, &key, &_dummy, 0, 0) == 0 {
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
	resource = PhpUrlParseEx2(str, str_len, &has_port)
	if resource == nil {

		/* @todo Find a method to determine why php_url_parse_ex() failed */

		return_value.u1.type_info = 2
		return
	}
	if key > -1 {
		switch key {
		case 0:
			if resource.GetScheme() != nil {
				var __z *zend.Zval = return_value
				var __s *zend.ZendString = resource.GetScheme()
				__z.value.str = __s
				if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
					__z.u1.type_info = 6
				} else {
					zend.ZendGcAddref(&__s.gc)
					__z.u1.type_info = 6 | 1<<0<<8
				}
			}
			break
		case 1:
			if resource.GetHost() != nil {
				var __z *zend.Zval = return_value
				var __s *zend.ZendString = resource.GetHost()
				__z.value.str = __s
				if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
					__z.u1.type_info = 6
				} else {
					zend.ZendGcAddref(&__s.gc)
					__z.u1.type_info = 6 | 1<<0<<8
				}
			}
			break
		case 2:
			if has_port != 0 {
				var __z *zend.Zval = return_value
				__z.value.lval = resource.GetPort()
				__z.u1.type_info = 4
			}
			break
		case 3:
			if resource.GetUser() != nil {
				var __z *zend.Zval = return_value
				var __s *zend.ZendString = resource.GetUser()
				__z.value.str = __s
				if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
					__z.u1.type_info = 6
				} else {
					zend.ZendGcAddref(&__s.gc)
					__z.u1.type_info = 6 | 1<<0<<8
				}
			}
			break
		case 4:
			if resource.GetPass() != nil {
				var __z *zend.Zval = return_value
				var __s *zend.ZendString = resource.GetPass()
				__z.value.str = __s
				if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
					__z.u1.type_info = 6
				} else {
					zend.ZendGcAddref(&__s.gc)
					__z.u1.type_info = 6 | 1<<0<<8
				}
			}
			break
		case 5:
			if resource.GetPath() != nil {
				var __z *zend.Zval = return_value
				var __s *zend.ZendString = resource.GetPath()
				__z.value.str = __s
				if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
					__z.u1.type_info = 6
				} else {
					zend.ZendGcAddref(&__s.gc)
					__z.u1.type_info = 6 | 1<<0<<8
				}
			}
			break
		case 6:
			if resource.GetQuery() != nil {
				var __z *zend.Zval = return_value
				var __s *zend.ZendString = resource.GetQuery()
				__z.value.str = __s
				if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
					__z.u1.type_info = 6
				} else {
					zend.ZendGcAddref(&__s.gc)
					__z.u1.type_info = 6 | 1<<0<<8
				}
			}
			break
		case 7:
			if resource.GetFragment() != nil {
				var __z *zend.Zval = return_value
				var __s *zend.ZendString = resource.GetFragment()
				__z.value.str = __s
				if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
					__z.u1.type_info = 6
				} else {
					zend.ZendGcAddref(&__s.gc)
					__z.u1.type_info = 6 | 1<<0<<8
				}
			}
			break
		default:
			core.PhpErrorDocref(nil, 1<<1, "Invalid URL component identifier "+"%"+"lld", key)
			return_value.u1.type_info = 2
		}
		goto done
	}

	/* allocate an array for return */

	var __arr *zend.ZendArray = zend._zendNewArray(0)
	var __z *zend.Zval = return_value
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8

	/* add the various elements to the array */

	if resource.GetScheme() != nil {
		var __z *zend.Zval = &tmp
		var __s *zend.ZendString = resource.GetScheme()
		__z.value.str = __s
		if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
			__z.u1.type_info = 6
		} else {
			zend.ZendGcAddref(&__s.gc)
			__z.u1.type_info = 6 | 1<<0<<8
		}
		zend.ZendHashAddNew(return_value.value.arr, zend.ZendKnownStrings[zend.ZEND_STR_SCHEME], &tmp)
	}
	if resource.GetHost() != nil {
		var __z *zend.Zval = &tmp
		var __s *zend.ZendString = resource.GetHost()
		__z.value.str = __s
		if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
			__z.u1.type_info = 6
		} else {
			zend.ZendGcAddref(&__s.gc)
			__z.u1.type_info = 6 | 1<<0<<8
		}
		zend.ZendHashAddNew(return_value.value.arr, zend.ZendKnownStrings[zend.ZEND_STR_HOST], &tmp)
	}
	if has_port != 0 {
		var __z *zend.Zval = &tmp
		__z.value.lval = resource.GetPort()
		__z.u1.type_info = 4
		zend.ZendHashAddNew(return_value.value.arr, zend.ZendKnownStrings[zend.ZEND_STR_PORT], &tmp)
	}
	if resource.GetUser() != nil {
		var __z *zend.Zval = &tmp
		var __s *zend.ZendString = resource.GetUser()
		__z.value.str = __s
		if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
			__z.u1.type_info = 6
		} else {
			zend.ZendGcAddref(&__s.gc)
			__z.u1.type_info = 6 | 1<<0<<8
		}
		zend.ZendHashAddNew(return_value.value.arr, zend.ZendKnownStrings[zend.ZEND_STR_USER], &tmp)
	}
	if resource.GetPass() != nil {
		var __z *zend.Zval = &tmp
		var __s *zend.ZendString = resource.GetPass()
		__z.value.str = __s
		if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
			__z.u1.type_info = 6
		} else {
			zend.ZendGcAddref(&__s.gc)
			__z.u1.type_info = 6 | 1<<0<<8
		}
		zend.ZendHashAddNew(return_value.value.arr, zend.ZendKnownStrings[zend.ZEND_STR_PASS], &tmp)
	}
	if resource.GetPath() != nil {
		var __z *zend.Zval = &tmp
		var __s *zend.ZendString = resource.GetPath()
		__z.value.str = __s
		if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
			__z.u1.type_info = 6
		} else {
			zend.ZendGcAddref(&__s.gc)
			__z.u1.type_info = 6 | 1<<0<<8
		}
		zend.ZendHashAddNew(return_value.value.arr, zend.ZendKnownStrings[zend.ZEND_STR_PATH], &tmp)
	}
	if resource.GetQuery() != nil {
		var __z *zend.Zval = &tmp
		var __s *zend.ZendString = resource.GetQuery()
		__z.value.str = __s
		if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
			__z.u1.type_info = 6
		} else {
			zend.ZendGcAddref(&__s.gc)
			__z.u1.type_info = 6 | 1<<0<<8
		}
		zend.ZendHashAddNew(return_value.value.arr, zend.ZendKnownStrings[zend.ZEND_STR_QUERY], &tmp)
	}
	if resource.GetFragment() != nil {
		var __z *zend.Zval = &tmp
		var __s *zend.ZendString = resource.GetFragment()
		__z.value.str = __s
		if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
			__z.u1.type_info = 6
		} else {
			zend.ZendGcAddref(&__s.gc)
			__z.u1.type_info = 6 | 1<<0<<8
		}
		zend.ZendHashAddNew(return_value.value.arr, zend.ZendKnownStrings[zend.ZEND_STR_FRAGMENT], &tmp)
	}
done:
	PhpUrlFree(resource)
}

/* }}} */

func PhpHtoi(s *byte) int {
	var value int
	var c int
	c = (*uint8)(s)[0]
	if isupper(c) {
		c = tolower(c)
	}
	value = g.Cond(c >= '0' && c <= '9', c-'0', c-'a'+10) * 16
	c = (*uint8)(s)[1]
	if isupper(c) {
		c = tolower(c)
	}
	if c >= '0' && c <= '9' {
		value += c - '0'
	} else {
		value += c - 'a' + 10
	}
	return value
}

/* }}} */

var UrlHexchars []uint8 = "0123456789ABCDEF"

/* {{{ php_url_encode
 */

func PhpUrlEncode(s *byte, len_ int) *zend.ZendString {
	var c uint8
	var to *uint8
	var from *uint8
	var end uint8
	var start *zend.ZendString
	from = (*uint8)(s)
	end = (*uint8)(s + len_)
	start = zend.ZendStringSafeAlloc(3, len_, 0, 0)
	to = (*uint8)(start.val)
	for from < end {
		*from++
		c = (*from) - 1
		if c == ' ' {
			g.PostInc(&(*to)) = '+'
		} else if c < '0' && c != '-' && c != '.' || c < 'A' && c > '9' || c > 'Z' && c < 'a' && c != '_' || c > 'z' {
			to[0] = '%'
			to[1] = Hexchars[c>>4]
			to[2] = Hexchars[c&15]
			to += 3
		} else {
			g.PostInc(&(*to)) = c
		}
	}
	*to = '0'
	start = zend.ZendStringTruncate(start, to-(*uint8)(start.val), 0)
	return start
}

/* }}} */

func ZifUrlencode(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var in_str *zend.ZendString
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

			if zend.ZendParseArgStr(_arg, &in_str, 0) == 0 {
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
	var __s *zend.ZendString = PhpUrlEncode(in_str.val, in_str.len_)
	__z.value.str = __s
	if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
		__z.u1.type_info = 6
	} else {
		__z.u1.type_info = 6 | 1<<0<<8
	}
	return
}

/* }}} */

func ZifUrldecode(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var in_str *zend.ZendString
	var out_str *zend.ZendString
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

			if zend.ZendParseArgStr(_arg, &in_str, 0) == 0 {
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
	out_str = zend.ZendStringInit(in_str.val, in_str.len_, 0)
	out_str.len_ = PhpUrlDecode(out_str.val, out_str.len_)
	var __z *zend.Zval = return_value
	var __s *zend.ZendString = out_str
	__z.value.str = __s
	__z.u1.type_info = 6 | 1<<0<<8
	return
}

/* }}} */

func PhpUrlDecode(str *byte, len_ int) int {
	var dest *byte = str
	var data *byte = str
	for g.PostDec(&len_) {
		if (*data) == '+' {
			*dest = ' '
		} else if (*data) == '%' && len_ >= 2 && isxdigit(int(*(data + 1))) && isxdigit(int(*(data + 2))) {
			*dest = byte(PhpHtoi(data + 1))
			data += 2
			len_ -= 2
		} else {
			*dest = *data
		}
		data++
		dest++
	}
	*dest = '0'
	return dest - str
}

/* }}} */

func PhpRawUrlEncode(s *byte, len_ int) *zend.ZendString {
	var x int
	var y int
	var str *zend.ZendString
	var ret *byte
	str = zend.ZendStringSafeAlloc(3, len_, 0, 0)
	ret = str.val
	x = 0
	y = 0
	for g.PostDec(&len_) {
		var c byte = s[x]
		ret[y] = c
		if c < '0' && c != '-' && c != '.' || c < 'A' && c > '9' || c > 'Z' && c < 'a' && c != '_' || c > 'z' && c != '~' {
			ret[g.PostInc(&y)] = '%'
			ret[g.PostInc(&y)] = Hexchars[uint8(c>>4)]
			ret[y] = Hexchars[uint8(c&15)]
		}
		x++
		y++
	}
	ret[y] = '0'
	str = zend.ZendStringTruncate(str, y, 0)
	return str
}

/* }}} */

func ZifRawurlencode(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var in_str *zend.ZendString
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

			if zend.ZendParseArgStr(_arg, &in_str, 0) == 0 {
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
	var __s *zend.ZendString = PhpRawUrlEncode(in_str.val, in_str.len_)
	__z.value.str = __s
	if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
		__z.u1.type_info = 6
	} else {
		__z.u1.type_info = 6 | 1<<0<<8
	}
	return
}

/* }}} */

func ZifRawurldecode(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var in_str *zend.ZendString
	var out_str *zend.ZendString
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

			if zend.ZendParseArgStr(_arg, &in_str, 0) == 0 {
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
	out_str = zend.ZendStringInit(in_str.val, in_str.len_, 0)
	out_str.len_ = PhpRawUrlDecode(out_str.val, out_str.len_)
	var __z *zend.Zval = return_value
	var __s *zend.ZendString = out_str
	__z.value.str = __s
	__z.u1.type_info = 6 | 1<<0<<8
	return
}

/* }}} */

func PhpRawUrlDecode(str *byte, len_ int) int {
	var dest *byte = str
	var data *byte = str
	for g.PostDec(&len_) {
		if (*data) == '%' && len_ >= 2 && isxdigit(int(*(data + 1))) && isxdigit(int(*(data + 2))) {
			*dest = byte(PhpHtoi(data + 1))
			data += 2
			len_ -= 2
		} else {
			*dest = *data
		}
		data++
		dest++
	}
	*dest = '0'
	return dest - str
}

/* }}} */

func ZifGetHeaders(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var url *byte
	var url_len int
	var stream *core.PhpStream
	var prev_val *zend.Zval
	var hdr *zend.Zval = nil
	var format zend.ZendLong = 0
	var zcontext *zend.Zval = nil
	var context *core.PhpStreamContext
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

			if zend.ZendParseArgPath(_arg, &url, &url_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_PATH
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

			if zend.ZendParseArgLong(_arg, &format, &_dummy, 0, 0) == 0 {
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

			if zend.ZendParseArgResource(_arg, &zcontext, 1) == 0 {
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
	if g.CondF2(g.CondF1(zcontext != nil, func() any { return zend.ZendFetchResourceEx(zcontext, "Stream-Context", PhpLeStreamContext()) }, 0), nil, func() *core.PhpStreamContext { return FileGlobals.GetDefaultContext() }) {
		context = FileGlobals.GetDefaultContext()
	} else {
		FileGlobals.SetDefaultContext(streams.PhpStreamContextAlloc())
		context = FileGlobals.GetDefaultContext()
	}
	if !(g.Assign(&stream, streams._phpStreamOpenWrapperEx(url, "r", 0x8|0x100|0x200, nil, context))) {
		return_value.u1.type_info = 2
		return
	}
	if stream.wrapperdata.u1.v.type_ != 7 {
		streams._phpStreamFree(stream, 1|2)
		return_value.u1.type_info = 2
		return
	}
	var __arr *zend.ZendArray = zend._zendNewArray(0)
	var __z *zend.Zval = return_value
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	for {
		var __ht *zend.HashTable = &stream.wrapperdata.value.arr
		var _p *zend.Bucket = __ht.arData
		var _end *zend.Bucket = _p + __ht.nNumUsed
		for ; _p != _end; _p++ {
			var _z *zend.Zval = &_p.val

			if _z.u1.v.type_ == 0 {
				continue
			}
			hdr = _z
			if hdr.u1.v.type_ != 6 {
				continue
			}
			if format == 0 {
			no_name_header:
				zend.AddNextIndexStr(return_value, zend.ZendStringCopy(hdr.value.str))
			} else {
				var c byte
				var s *byte
				var p *byte
				if g.Assign(&p, strchr(hdr.value.str.val, ':')) {
					c = *p
					*p = '0'
					s = p + 1
					for isspace(int(*((*uint8)(s)))) {
						s++
					}
					if g.Assign(&prev_val, zend.ZendHashStrFind(return_value.value.arr, hdr.value.str.val, p-hdr.value.str.val)) == nil {
						zend.AddAssocStringlEx(return_value, hdr.value.str.val, p-hdr.value.str.val, s, hdr.value.str.len_-(s-hdr.value.str.val))
					} else {
						zend.ConvertToArray(prev_val)
						zend.AddNextIndexStringl(prev_val, s, hdr.value.str.len_-(s-hdr.value.str.val))
					}
					*p = c
				} else {
					goto no_name_header
				}
			}
		}
		break
	}
	streams._phpStreamFree(stream, 1|2)
}

/* }}} */
