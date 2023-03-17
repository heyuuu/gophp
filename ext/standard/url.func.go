// <<generate>>

package standard

import (
	b "sik/builtin"
	"sik/core"
	"sik/core/streams"
	"sik/zend"
)

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
	zend.Efree(theurl)
}
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
func PhpUrlParseEx(str *byte, length int) *PhpUrl {
	var has_port zend.ZendBool
	return PhpUrlParseEx2(str, length, &has_port)
}
func PhpUrlParseEx2(str *byte, length int, has_port *zend.ZendBool) *PhpUrl {
	var port_buf []byte
	var ret *PhpUrl = zend.Ecalloc(1, b.SizeOf("php_url"))
	var s *byte
	var e byte
	var p byte
	var pp byte
	var ue byte
	*has_port = 0
	s = str
	ue = s + length

	/* parse scheme */

	if b.Assign(&e, memchr(s, ':', length)) && e != s {

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
			PhpReplaceControlcharsEx(ret.GetScheme().GetVal(), ret.GetScheme().GetLen())
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
			PhpReplaceControlcharsEx(ret.GetScheme().GetVal(), ret.GetScheme().GetLen())
			s = e + 1
			goto just_path
		} else {
			ret.SetScheme(zend.ZendStringInit(s, e-s, 0))
			PhpReplaceControlcharsEx(ret.GetScheme().GetVal(), ret.GetScheme().GetLen())
			if e+2 < ue && (*(e + 2)) == '/' {
				s = e + 3
				if zend.ZendStringEqualsLiteralCi(ret.GetScheme(), "file") {
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
			port = zend.ZEND_STRTOL(port_buf, &end, 10)
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

	if b.Assign(&p, zend.ZendMemrchr(s, '@', e-s)) {
		if b.Assign(&pp, memchr(s, ':', p-s)) {
			ret.SetUser(zend.ZendStringInit(s, pp-s, 0))
			PhpReplaceControlcharsEx(ret.GetUser().GetVal(), ret.GetUser().GetLen())
			pp++
			ret.SetPass(zend.ZendStringInit(pp, p-pp, 0))
			PhpReplaceControlcharsEx(ret.GetPass().GetVal(), ret.GetPass().GetLen())
		} else {
			ret.SetUser(zend.ZendStringInit(s, p-s, 0))
			PhpReplaceControlcharsEx(ret.GetUser().GetVal(), ret.GetUser().GetLen())
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
				port = zend.ZEND_STRTOL(port_buf, &end, 10)
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
	PhpReplaceControlcharsEx(ret.GetHost().GetVal(), ret.GetHost().GetLen())
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
			PhpReplaceControlcharsEx(ret.GetFragment().GetVal(), ret.GetFragment().GetLen())
		}
		e = p - 1
	}
	p = memchr(s, '?', e-s)
	if p {
		p++
		if p < e {
			ret.SetQuery(zend.ZendStringInit(p, e-p, 0))
			PhpReplaceControlcharsEx(ret.GetQuery().GetVal(), ret.GetQuery().GetLen())
		}
		e = p - 1
	}
	if s < e || s == ue {
		ret.SetPath(zend.ZendStringInit(s, e-s, 0))
		PhpReplaceControlcharsEx(ret.GetPath().GetVal(), ret.GetPath().GetLen())
	}
	return ret
}
func ZifParseUrl(executeData *zend.ZendExecuteData, return_value *zend.Zval) {
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
		var _num_args int = executeData.NumArgs()
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
			if zend.ZendParseArgString(_arg, &str, &str_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgLong(_arg, &key, &_dummy, 0, 0) == 0 {
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
	resource = PhpUrlParseEx2(str, str_len, &has_port)
	if resource == nil {

		/* @todo Find a method to determine why php_url_parse_ex() failed */

		return_value.SetFalse()
		return
	}
	if key > -1 {
		switch key {
		case PHP_URL_SCHEME:
			if resource.GetScheme() != nil {
				return_value.SetStringCopy(resource.GetScheme())
			}
		case PHP_URL_HOST:
			if resource.GetHost() != nil {
				return_value.SetStringCopy(resource.GetHost())
			}
		case PHP_URL_PORT:
			if has_port != 0 {
				return_value.SetLong(resource.GetPort())
			}
		case PHP_URL_USER:
			if resource.GetUser() != nil {
				return_value.SetStringCopy(resource.GetUser())
			}
		case PHP_URL_PASS:
			if resource.GetPass() != nil {
				return_value.SetStringCopy(resource.GetPass())
			}
		case PHP_URL_PATH:
			if resource.GetPath() != nil {
				return_value.SetStringCopy(resource.GetPath())
			}
		case PHP_URL_QUERY:
			if resource.GetQuery() != nil {
				return_value.SetStringCopy(resource.GetQuery())
			}
		case PHP_URL_FRAGMENT:
			if resource.GetFragment() != nil {
				return_value.SetStringCopy(resource.GetFragment())
			}
		default:
			core.PhpErrorDocref(nil, zend.E_WARNING, "Invalid URL component identifier "+zend.ZEND_LONG_FMT, key)
			return_value.SetFalse()
		}
		goto done
	}

	/* allocate an array for return */

	zend.ArrayInit(return_value)

	/* add the various elements to the array */

	if resource.GetScheme() != nil {
		tmp.SetStringCopy(resource.GetScheme())
		return_value.GetArr().KeyAddNew(zend.ZSTR_KNOWN(zend.ZEND_STR_SCHEME).GetStr(), &tmp)
	}
	if resource.GetHost() != nil {
		tmp.SetStringCopy(resource.GetHost())
		return_value.GetArr().KeyAddNew(zend.ZSTR_KNOWN(zend.ZEND_STR_HOST).GetStr(), &tmp)
	}
	if has_port != 0 {
		tmp.SetLong(resource.GetPort())
		return_value.GetArr().KeyAddNew(zend.ZSTR_KNOWN(zend.ZEND_STR_PORT).GetStr(), &tmp)
	}
	if resource.GetUser() != nil {
		tmp.SetStringCopy(resource.GetUser())
		return_value.GetArr().KeyAddNew(zend.ZSTR_KNOWN(zend.ZEND_STR_USER).GetStr(), &tmp)
	}
	if resource.GetPass() != nil {
		tmp.SetStringCopy(resource.GetPass())
		return_value.GetArr().KeyAddNew(zend.ZSTR_KNOWN(zend.ZEND_STR_PASS).GetStr(), &tmp)
	}
	if resource.GetPath() != nil {
		tmp.SetStringCopy(resource.GetPath())
		return_value.GetArr().KeyAddNew(zend.ZSTR_KNOWN(zend.ZEND_STR_PATH).GetStr(), &tmp)
	}
	if resource.GetQuery() != nil {
		tmp.SetStringCopy(resource.GetQuery())
		return_value.GetArr().KeyAddNew(zend.ZSTR_KNOWN(zend.ZEND_STR_QUERY).GetStr(), &tmp)
	}
	if resource.GetFragment() != nil {
		tmp.SetStringCopy(resource.GetFragment())
		return_value.GetArr().KeyAddNew(zend.ZSTR_KNOWN(zend.ZEND_STR_FRAGMENT).GetStr(), &tmp)
	}
done:
	PhpUrlFree(resource)
}
func PhpHtoi(s *byte) int {
	var value int
	var c int
	c = (*uint8)(s)[0]
	if isupper(c) {
		c = tolower(c)
	}
	value = b.Cond(c >= '0' && c <= '9', c-'0', c-'a'+10) * 16
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
func PhpUrlEncode(s *byte, len_ int) *zend.ZendString {
	var c uint8
	var to *uint8
	var from *uint8
	var end uint8
	var start *zend.ZendString
	from = (*uint8)(s)
	end = (*uint8)(s + len_)
	start = zend.ZendStringSafeAlloc(3, len_, 0, 0)
	to = (*uint8)(start.GetVal())
	for from < end {
		*from++
		c = (*from) - 1
		if c == ' ' {
			b.PostInc(&(*to)) = '+'
		} else if c < '0' && c != '-' && c != '.' || c < 'A' && c > '9' || c > 'Z' && c < 'a' && c != '_' || c > 'z' {
			to[0] = '%'
			to[1] = Hexchars[c>>4]
			to[2] = Hexchars[c&15]
			to += 3
		} else {
			b.PostInc(&(*to)) = c
		}
	}
	*to = '0'
	start = zend.ZendStringTruncate(start, to-(*uint8)(start.GetVal()), 0)
	return start
}
func ZifUrlencode(executeData *zend.ZendExecuteData, return_value *zend.Zval) {
	var in_str *zend.ZendString
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = executeData.NumArgs()
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
			if zend.ZendParseArgStr(_arg, &in_str, 0) == 0 {
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
	return_value.SetString(PhpUrlEncode(in_str.GetVal(), in_str.GetLen()))
	return
}
func ZifUrldecode(executeData *zend.ZendExecuteData, return_value *zend.Zval) {
	var in_str *zend.ZendString
	var out_str *zend.ZendString
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = executeData.NumArgs()
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
			if zend.ZendParseArgStr(_arg, &in_str, 0) == 0 {
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
	out_str = zend.ZendStringInit(in_str.GetVal(), in_str.GetLen(), 0)
	out_str.SetLen(PhpUrlDecode(out_str.GetVal(), out_str.GetLen()))
	return_value.SetString(out_str)
	return
}
func PhpUrlDecode(str *byte, len_ int) int {
	var dest *byte = str
	var data *byte = str
	for b.PostDec(&len_) {
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
func PhpRawUrlEncode(s *byte, len_ int) *zend.ZendString {
	var x int
	var y int
	var str *zend.ZendString
	var ret *byte
	str = zend.ZendStringSafeAlloc(3, len_, 0, 0)
	ret = str.GetVal()
	x = 0
	y = 0
	for b.PostDec(&len_) {
		var c byte = s[x]
		ret[y] = c
		if c < '0' && c != '-' && c != '.' || c < 'A' && c > '9' || c > 'Z' && c < 'a' && c != '_' || c > 'z' && c != '~' {
			ret[b.PostInc(&y)] = '%'
			ret[b.PostInc(&y)] = Hexchars[uint8(c>>4)]
			ret[y] = Hexchars[uint8(c&15)]
		}
		x++
		y++
	}
	ret[y] = '0'
	str = zend.ZendStringTruncate(str, y, 0)
	return str
}
func ZifRawurlencode(executeData *zend.ZendExecuteData, return_value *zend.Zval) {
	var in_str *zend.ZendString
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = executeData.NumArgs()
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
			if zend.ZendParseArgStr(_arg, &in_str, 0) == 0 {
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
	return_value.SetString(PhpRawUrlEncode(in_str.GetVal(), in_str.GetLen()))
	return
}
func ZifRawurldecode(executeData *zend.ZendExecuteData, return_value *zend.Zval) {
	var in_str *zend.ZendString
	var out_str *zend.ZendString
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = executeData.NumArgs()
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
			if zend.ZendParseArgStr(_arg, &in_str, 0) == 0 {
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
	out_str = zend.ZendStringInit(in_str.GetVal(), in_str.GetLen(), 0)
	out_str.SetLen(PhpRawUrlDecode(out_str.GetVal(), out_str.GetLen()))
	return_value.SetString(out_str)
	return
}
func PhpRawUrlDecode(str *byte, len_ int) int {
	var dest *byte = str
	var data *byte = str
	for b.PostDec(&len_) {
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
func ZifGetHeaders(executeData *zend.ZendExecuteData, return_value *zend.Zval) {
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
		var _num_args int = executeData.NumArgs()
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
			if zend.ZendParseArgPath(_arg, &url, &url_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_PATH
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgLong(_arg, &format, &_dummy, 0, 0) == 0 {
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
	context = streams.PhpStreamContextFromZval(zcontext, 0)
	if !(b.Assign(&stream, core.PhpStreamOpenWrapperEx(url, "r", core.REPORT_ERRORS|core.STREAM_USE_URL|core.STREAM_ONLY_GET_HEADERS, nil, context))) {
		return_value.SetFalse()
		return
	}
	if stream.GetWrapperdata().GetType() != zend.IS_ARRAY {
		core.PhpStreamClose(stream)
		return_value.SetFalse()
		return
	}
	zend.ArrayInit(return_value)
	var __ht *zend.HashTable = stream.GetWrapperdata().GetArr()
	for _, _p := range __ht.foreachData() {
		var _z *zend.Zval = _p.GetVal()

		hdr = _z
		if hdr.GetType() != zend.IS_STRING {
			continue
		}
		if format == 0 {
		no_name_header:
			zend.AddNextIndexStr(return_value, hdr.GetStr().Copy())
		} else {
			var c byte
			var s *byte
			var p *byte
			if b.Assign(&p, strchr(zend.Z_STRVAL_P(hdr), ':')) {
				c = *p
				*p = '0'
				s = p + 1
				for isspace(int(*((*uint8)(s)))) {
					s++
				}
				if b.Assign(&prev_val, return_value.GetArr().KeyFind(b.CastStr(zend.Z_STRVAL_P(hdr), p-zend.Z_STRVAL_P(hdr)))) == nil {
					zend.AddAssocStringlEx(return_value, b.CastStr(zend.Z_STRVAL_P(hdr), p-zend.Z_STRVAL_P(hdr)), b.CastStr(s, zend.Z_STRLEN_P(hdr)-(s-zend.Z_STRVAL_P(hdr))))
				} else {
					zend.ConvertToArray(prev_val)
					zend.AddNextIndexStringl(prev_val, s, zend.Z_STRLEN_P(hdr)-(s-zend.Z_STRVAL_P(hdr)))
				}
				*p = c
			} else {
				goto no_name_header
			}
		}
	}
	core.PhpStreamClose(stream)
}
