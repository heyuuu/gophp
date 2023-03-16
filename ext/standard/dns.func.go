// <<generate>>

package standard

import (
	b "sik/builtin"
	"sik/core"
	"sik/zend"
)

func ZifGethostname(executeData *zend.ZendExecuteData, return_value *zend.Zval) {
	var buf []byte
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	if gethostname(buf, b.SizeOf("buf")) {
		core.PhpErrorDocref(nil, zend.E_WARNING, "unable to fetch host [%d]: %s", errno, strerror(errno))
		return_value.SetFalse()
		return
	}
	return_value.SetRawString(b.CastStrAuto(buf))
	return
}
func ZifGethostbyaddr(executeData *zend.ZendExecuteData, return_value *zend.Zval) {
	var addr *byte
	var addr_len int
	var hostname *zend.ZendString
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
						zend.CheckNumArgsException(_min_num_args, _max_num_args)
					} else {
						zend.CheckNumArgsError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(executeData, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgString(_arg, &addr, &addr_len, 0) == 0 {
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
	hostname = PhpGethostbyaddr(addr)
	if hostname == nil {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Address is not a valid IPv4 or IPv6 address")
		return_value.SetFalse()
	} else {
		return_value.SetString(hostname)
	}
}
func PhpGethostbyaddr(ip *byte) *zend.ZendString {
	var addr6 __struct__in6_addr
	var addr __struct__in_addr
	var hp *__struct__hostent
	if inet_pton(AF_INET6, ip, &addr6) {
		hp = gethostbyaddr((*byte)(&addr6), b.SizeOf("addr6"), AF_INET6)
	} else if inet_pton(AF_INET, ip, &addr) {
		hp = gethostbyaddr((*byte)(&addr), b.SizeOf("addr"), AF_INET)
	} else {
		return nil
	}
	if hp == nil || hp.h_name == nil || hp.h_name[0] == '0' {
		return zend.ZendStringInit(ip, strlen(ip), 0)
	}
	return zend.ZendStringInit(hp.h_name, strlen(hp.h_name), 0)
}
func ZifGethostbyname(executeData *zend.ZendExecuteData, return_value *zend.Zval) {
	var hostname *byte
	var hostname_len int
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
						zend.CheckNumArgsException(_min_num_args, _max_num_args)
					} else {
						zend.CheckNumArgsError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(executeData, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgString(_arg, &hostname, &hostname_len, 0) == 0 {
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
	if hostname_len > core.MAXFQDNLEN {

		/* name too long, protect from CVE-2015-0235 */

		core.PhpErrorDocref(nil, zend.E_WARNING, "Host name is too long, the limit is %d characters", core.MAXFQDNLEN)
		return_value.SetRawString(b.CastStr(hostname, hostname_len))
		return
	}
	return_value.SetString(PhpGethostbyname(hostname))
	return
}
func ZifGethostbynamel(executeData *zend.ZendExecuteData, return_value *zend.Zval) {
	var hostname *byte
	var hostname_len int
	var hp *__struct__hostent
	var in __struct__in_addr
	var i int
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
						zend.CheckNumArgsException(_min_num_args, _max_num_args)
					} else {
						zend.CheckNumArgsError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(executeData, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgString(_arg, &hostname, &hostname_len, 0) == 0 {
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
	if hostname_len > core.MAXFQDNLEN {

		/* name too long, protect from CVE-2015-0235 */

		core.PhpErrorDocref(nil, zend.E_WARNING, "Host name is too long, the limit is %d characters", core.MAXFQDNLEN)
		return_value.SetFalse()
		return
	}
	hp = core.PhpNetworkGethostbyname(hostname)
	if hp == nil {
		return_value.SetFalse()
		return
	}
	zend.ArrayInit(return_value)
	for i = 0; ; i++ {

		/* On macos h_addr_list entries may be misaligned. */

		var h_addr_entry *__struct__in_addr
		memcpy(&h_addr_entry, hp.h_addr_list[i], b.SizeOf("struct in_addr *"))
		if h_addr_entry == nil {
			return
		}
		in = *h_addr_entry
		zend.AddNextIndexString(return_value, inet_ntoa(in))
	}
}
func PhpGethostbyname(name *byte) *zend.ZendString {
	var hp *__struct__hostent
	var h_addr_0 *__struct__in_addr
	var in __struct__in_addr
	var address *byte
	hp = core.PhpNetworkGethostbyname(name)
	if hp == nil {
		return zend.ZendStringInit(name, strlen(name), 0)
	}

	/* On macos h_addr_list entries may be misaligned. */

	memcpy(&h_addr_0, hp.h_addr_list[0], b.SizeOf("struct in_addr *"))
	if h_addr_0 == nil {
		return zend.ZendStringInit(name, strlen(name), 0)
	}
	memcpy(in.s_addr, h_addr_0, b.SizeOf("in . s_addr"))
	address = inet_ntoa(in)
	return zend.ZendStringInit(address, strlen(address), 0)
}
func ZifDnsCheckRecord(executeData *zend.ZendExecuteData, return_value *zend.Zval) {
	var hp *HEADER
	var answer Querybuf
	var hostname *byte
	var rectype *byte = nil
	var hostname_len int
	var rectype_len int = 0
	var type_ int = DNS_T_MX
	var i int
	var from __struct__sockaddr_storage
	var fromsize uint32 = b.SizeOf("from")
	var handle dns_handle_t
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
						zend.CheckNumArgsException(_min_num_args, _max_num_args)
					} else {
						zend.CheckNumArgsError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(executeData, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgString(_arg, &hostname, &hostname_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgString(_arg, &rectype, &rectype_len, 0) == 0 {
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
	if hostname_len == 0 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Host cannot be empty")
		return_value.SetFalse()
		return
	}
	if rectype != nil {
		if !(strcasecmp("A", rectype)) {
			type_ = DNS_T_A
		} else if !(strcasecmp("NS", rectype)) {
			type_ = DNS_T_NS
		} else if !(strcasecmp("MX", rectype)) {
			type_ = DNS_T_MX
		} else if !(strcasecmp("PTR", rectype)) {
			type_ = DNS_T_PTR
		} else if !(strcasecmp("ANY", rectype)) {
			type_ = DNS_T_ANY
		} else if !(strcasecmp("SOA", rectype)) {
			type_ = DNS_T_SOA
		} else if !(strcasecmp("CAA", rectype)) {
			type_ = DNS_T_CAA
		} else if !(strcasecmp("TXT", rectype)) {
			type_ = DNS_T_TXT
		} else if !(strcasecmp("CNAME", rectype)) {
			type_ = DNS_T_CNAME
		} else if !(strcasecmp("AAAA", rectype)) {
			type_ = DNS_T_AAAA
		} else if !(strcasecmp("SRV", rectype)) {
			type_ = DNS_T_SRV
		} else if !(strcasecmp("NAPTR", rectype)) {
			type_ = DNS_T_NAPTR
		} else if !(strcasecmp("A6", rectype)) {
			type_ = DNS_T_A6
		} else {
			core.PhpErrorDocref(nil, zend.E_WARNING, "Type '%s' not supported", rectype)
			return_value.SetFalse()
			return
		}
	}
	handle = dns_open(nil)
	if handle == nil {
		return_value.SetFalse()
		return
	}
	i = PhpDnsSearch(handle, hostname, C_IN, type_, answer.GetQb2(), b.SizeOf(answer))
	PhpDnsFreeHandle(handle)
	if i < 0 {
		return_value.SetFalse()
		return
	}
	hp = (*HEADER)(&answer)
	zend.ZVAL_BOOL(return_value, ntohs(hp.ancount) != 0)
	return
}
func PhpParserr(
	cp *u_char,
	end *u_char,
	answer *Querybuf,
	type_to_fetch int,
	store int,
	raw int,
	subarray *zend.Zval,
) *u_char {
	var type_ u_short
	var class u_short
	var dlen u_short
	var ttl u_long
	var n long
	var i long
	var s u_short
	var tp *u_char
	var p *u_char
	var name []byte
	var have_v6_break int = 0
	var in_v6_break int = 0
	subarray.SetUndef()
	n = dn_expand(answer.GetQb2(), end, cp, name, b.SizeOf("name")-2)
	if n < 0 {
		return nil
	}
	cp += n
	if cp+10 > end {
		return nil
	}
	GETSHORT(type_, cp)
	GETSHORT(class, cp)
	GETLONG(ttl, cp)
	GETSHORT(dlen, cp)
	if cp+dlen > end {
		return nil
	}
	if dlen == 0 {

		/* No data in the response - nothing to do */

		return nil

		/* No data in the response - nothing to do */

	}
	if type_to_fetch != DNS_T_ANY && type_ != type_to_fetch {
		cp += dlen
		return cp
	}
	if store == 0 {
		cp += dlen
		return cp
	}
	zend.ArrayInit(subarray)
	zend.AddAssocString(subarray, "host", name)
	zend.AddAssocString(subarray, "class", "IN")
	zend.AddAssocLong(subarray, "ttl", ttl)
	void(class)
	if raw != 0 {
		zend.AddAssocLong(subarray, "type", type_)
		zend.AddAssocStringl(subarray, "data", (*byte)(cp), uint32(dlen))
		cp += dlen
		return cp
	}
	switch type_ {
	case DNS_T_A:
		if cp+4 > end {
			return nil
		}
		zend.AddAssocString(subarray, "type", "A")
		core.Snprintf(name, b.SizeOf("name"), "%d.%d.%d.%d", cp[0], cp[1], cp[2], cp[3])
		zend.AddAssocString(subarray, "ip", name)
		cp += dlen
	case DNS_T_MX:
		if cp+2 > end {
			return nil
		}
		zend.AddAssocString(subarray, "type", "MX")
		GETSHORT(n, cp)
		zend.AddAssocLong(subarray, "pri", n)
		fallthrough
	case DNS_T_CNAME:
		if type_ == DNS_T_CNAME {
			zend.AddAssocString(subarray, "type", "CNAME")
		}
		fallthrough
	case DNS_T_NS:
		if type_ == DNS_T_NS {
			zend.AddAssocString(subarray, "type", "NS")
		}
		fallthrough
	case DNS_T_PTR:
		if type_ == DNS_T_PTR {
			zend.AddAssocString(subarray, "type", "PTR")
		}
		n = dn_expand(answer.GetQb2(), end, cp, name, b.SizeOf(name)-2)
		if n < 0 {
			return nil
		}
		cp += n
		zend.AddAssocString(subarray, "target", name)
	case DNS_T_HINFO:

		/* See RFC 1010 for values */

		zend.AddAssocString(subarray, "type", "HINFO")
		if cp+1 > end {
			return nil
		}
		n = (*cp) & 0xff
		cp++
		if cp+n > end {
			return nil
		}
		zend.AddAssocStringl(subarray, "cpu", (*byte)(cp), n)
		cp += n
		if cp+1 > end {
			return nil
		}
		n = (*cp) & 0xff
		cp++
		if cp+n > end {
			return nil
		}
		zend.AddAssocStringl(subarray, "os", (*byte)(cp), n)
		cp += n
	case DNS_T_CAA:

		/* See RFC 6844 for values https://tools.ietf.org/html/rfc6844 */

		zend.AddAssocString(subarray, "type", "CAA")

		// 1 flag byte

		if cp+1 > end {
			return nil
		}
		n = (*cp) & 0xff
		zend.AddAssocLong(subarray, "flags", n)
		cp++

		// Tag length (1 byte)

		if cp+1 > end {
			return nil
		}
		n = (*cp) & 0xff
		cp++
		if cp+n > end {
			return nil
		}
		zend.AddAssocStringl(subarray, "tag", (*byte)(cp), n)
		cp += n
		if int(dlen < int(n)+2) != 0 {
			return nil
		}
		n = dlen - n - 2
		if cp+n > end {
			return nil
		}
		zend.AddAssocStringl(subarray, "value", (*byte)(cp), n)
		cp += n
	case DNS_T_TXT:
		var l1 int = 0
		var l2 int = 0
		var entries zend.Zval
		var tp *zend.ZendString
		zend.AddAssocString(subarray, "type", "TXT")
		tp = zend.ZendStringAlloc(dlen, 0)
		zend.ArrayInit(&entries)
		for l1 < dlen {
			n = cp[l1]
			if l1+n >= dlen {

				// Invalid chunk length, truncate

				n = dlen - (l1 + 1)

				// Invalid chunk length, truncate

			}
			if n {
				memcpy(tp.GetVal()+l2, cp+l1+1, n)
				zend.AddNextIndexStringl(&entries, (*byte)(cp+l1+1), n)
			}
			l1 = l1 + n + 1
			l2 = l2 + n
		}
		tp.GetVal()[l2] = '0'
		tp.SetLen(l2)
		cp += dlen
		zend.AddAssocStr(subarray, "txt", tp.GetStr())
		zend.AddAssocZval(subarray, "entries", &entries)
	case DNS_T_SOA:
		zend.AddAssocString(subarray, "type", "SOA")
		n = dn_expand(answer.GetQb2(), end, cp, name, b.SizeOf(name)-2)
		if n < 0 {
			return nil
		}
		cp += n
		zend.AddAssocString(subarray, "mname", name)
		n = dn_expand(answer.GetQb2(), end, cp, name, b.SizeOf(name)-2)
		if n < 0 {
			return nil
		}
		cp += n
		zend.AddAssocString(subarray, "rname", name)
		if cp+5*4 > end {
			return nil
		}
		GETLONG(n, cp)
		zend.AddAssocLong(subarray, "serial", n)
		GETLONG(n, cp)
		zend.AddAssocLong(subarray, "refresh", n)
		GETLONG(n, cp)
		zend.AddAssocLong(subarray, "retry", n)
		GETLONG(n, cp)
		zend.AddAssocLong(subarray, "expire", n)
		GETLONG(n, cp)
		zend.AddAssocLong(subarray, "minimum-ttl", n)
	case DNS_T_AAAA:
		tp = (*u_char)(name)
		if cp+8*2 > end {
			return nil
		}
		for i = 0; i < 8; i++ {
			GETSHORT(s, cp)
			if s != 0 {
				if tp > (*u_char)(name) {
					in_v6_break = 0
					tp[0] = ':'
					tp++
				}
				tp += sprintf((*byte)(tp), "%x", s)
			} else {
				if have_v6_break == 0 {
					have_v6_break = 1
					in_v6_break = 1
					tp[0] = ':'
					tp++
				} else if in_v6_break == 0 {
					tp[0] = ':'
					tp++
					tp[0] = '0'
					tp++
				}
			}
		}
		if have_v6_break != 0 && in_v6_break != 0 {
			tp[0] = ':'
			tp++
		}
		tp[0] = '0'
		zend.AddAssocString(subarray, "type", "AAAA")
		zend.AddAssocString(subarray, "ipv6", name)
	case DNS_T_A6:
		p = cp
		zend.AddAssocString(subarray, "type", "A6")
		if cp+1 > end {
			return nil
		}
		n = int(cp[0]) & 0xff
		cp++
		zend.AddAssocLong(subarray, "masklen", n)
		tp = (*u_char)(name)
		if n > 15 {
			have_v6_break = 1
			in_v6_break = 1
			tp[0] = ':'
			tp++
		}
		if n%16 > 8 {

			/* Partial short */

			if cp[0] != 0 {
				if tp > (*u_char)(name) {
					in_v6_break = 0
					tp[0] = ':'
					tp++
				}
				sprintf((*byte)(tp), "%x", cp[0]&0xff)
			} else {
				if have_v6_break == 0 {
					have_v6_break = 1
					in_v6_break = 1
					tp[0] = ':'
					tp++
				} else if in_v6_break == 0 {
					tp[0] = ':'
					tp++
					tp[0] = '0'
					tp++
				}
			}
			cp++
		}
		for i = (n + 8) / 16; i < 8; i++ {
			if cp+2 > end {
				return nil
			}
			GETSHORT(s, cp)
			if s != 0 {
				if tp > (*u_char)(name) {
					in_v6_break = 0
					tp[0] = ':'
					tp++
				}
				tp += sprintf((*byte)(tp), "%x", s)
			} else {
				if have_v6_break == 0 {
					have_v6_break = 1
					in_v6_break = 1
					tp[0] = ':'
					tp++
				} else if in_v6_break == 0 {
					tp[0] = ':'
					tp++
					tp[0] = '0'
					tp++
				}
			}
		}
		if have_v6_break != 0 && in_v6_break != 0 {
			tp[0] = ':'
			tp++
		}
		tp[0] = '0'
		zend.AddAssocString(subarray, "ipv6", name)
		if cp < p+dlen {
			n = dn_expand(answer.GetQb2(), end, cp, name, b.SizeOf(name)-2)
			if n < 0 {
				return nil
			}
			cp += n
			zend.AddAssocString(subarray, "chain", name)
		}
	case DNS_T_SRV:
		if cp+3*2 > end {
			return nil
		}
		zend.AddAssocString(subarray, "type", "SRV")
		GETSHORT(n, cp)
		zend.AddAssocLong(subarray, "pri", n)
		GETSHORT(n, cp)
		zend.AddAssocLong(subarray, "weight", n)
		GETSHORT(n, cp)
		zend.AddAssocLong(subarray, "port", n)
		n = dn_expand(answer.GetQb2(), end, cp, name, b.SizeOf(name)-2)
		if n < 0 {
			return nil
		}
		cp += n
		zend.AddAssocString(subarray, "target", name)
	case DNS_T_NAPTR:
		if cp+2*2 > end {
			return nil
		}
		zend.AddAssocString(subarray, "type", "NAPTR")
		GETSHORT(n, cp)
		zend.AddAssocLong(subarray, "order", n)
		GETSHORT(n, cp)
		zend.AddAssocLong(subarray, "pref", n)
		if cp+1 > end {
			return nil
		}
		n = cp[0] & 0xff
		cp++
		if cp+n > end {
			return nil
		}
		zend.AddAssocStringl(subarray, "flags", (*byte)(cp), n)
		cp += n
		if cp+1 > end {
			return nil
		}
		n = cp[0] & 0xff
		cp++
		if cp+n > end {
			return nil
		}
		zend.AddAssocStringl(subarray, "services", (*byte)(cp), n)
		cp += n
		if cp+1 > end {
			return nil
		}
		n = cp[0] & 0xff
		cp++
		if cp+n > end {
			return nil
		}
		zend.AddAssocStringl(subarray, "regex", (*byte)(cp), n)
		cp += n
		n = dn_expand(answer.GetQb2(), end, cp, name, b.SizeOf(name)-2)
		if n < 0 {
			return nil
		}
		cp += n
		zend.AddAssocString(subarray, "replacement", name)
	default:
		zend.ZvalPtrDtor(subarray)
		subarray.SetUndef()
		cp += dlen
	}
	return cp
}
func ZifDnsGetRecord(executeData *zend.ZendExecuteData, return_value *zend.Zval) {
	var hostname *byte
	var hostname_len int
	var type_param zend.ZendLong = PHP_DNS_ANY
	var authns *zend.Zval = nil
	var addtl *zend.Zval = nil
	var type_to_fetch int
	var dns_errno int
	var from __struct__sockaddr_storage
	var fromsize uint32 = b.SizeOf("from")
	var handle dns_handle_t
	var hp *HEADER
	var answer Querybuf
	var cp *u_char = nil
	var end *u_char = nil
	var n int
	var qd int
	var an int
	var ns int = 0
	var ar int = 0
	var type_ int
	var first_query int = 1
	var store_results int = 1
	var raw zend.ZendBool = 0
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
						zend.CheckNumArgsException(_min_num_args, _max_num_args)
					} else {
						zend.CheckNumArgsError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(executeData, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgString(_arg, &hostname, &hostname_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgLong(_arg, &type_param, &_dummy, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_LONG
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			zend.ZendParseArgZvalDeref(_arg, &authns, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			zend.ZendParseArgZvalDeref(_arg, &addtl, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgBool(_arg, &raw, &_dummy, 0) == 0 {
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
	if authns != nil {
		authns = zend.ZendTryArrayInit(authns)
		if authns == nil {
			return
		}
	}
	if addtl != nil {
		addtl = zend.ZendTryArrayInit(addtl)
		if addtl == nil {
			return
		}
	}
	if raw == 0 {
		if (type_param & ^PHP_DNS_ALL) != 0 && type_param != PHP_DNS_ANY {
			core.PhpErrorDocref(nil, zend.E_WARNING, "Type '"+zend.ZEND_LONG_FMT+"' not supported", type_param)
			return_value.SetFalse()
			return
		}
	} else {
		if type_param < 1 || type_param > 0xffff {
			core.PhpErrorDocref(nil, zend.E_WARNING, "Numeric DNS record type must be between 1 and 65535, '"+zend.ZEND_LONG_FMT+"' given", type_param)
			return_value.SetFalse()
			return
		}
	}

	/* Initialize the return array */

	zend.ArrayInit(return_value)

	/* - We emulate an or'ed type mask by querying type by type. (Steps 0 - NUMTYPES-1 )
	 *   If additional info is wanted we check again with DNS_T_ANY (step NUMTYPES / NUMTYPES+1 )
	 *   store_results is used to skip storing the results retrieved in step
	 *   NUMTYPES+1 when results were already fetched.
	 * - In case of PHP_DNS_ANY we use the directly fetch DNS_T_ANY. (step NUMTYPES+1 )
	 * - In case of raw mode, we query only the requestd type instead of looping type by type
	 *   before going with the additional info stuff.
	 */

	if raw != 0 {
		type_ = -1
	} else if type_param == PHP_DNS_ANY {
		type_ = PHP_DNS_NUM_TYPES + 1
	} else {
		type_ = 0
	}
	for ; type_ < b.Cond(addtl != nil, PHP_DNS_NUM_TYPES+2, PHP_DNS_NUM_TYPES) || first_query != 0; type_++ {
		first_query = 0
		switch type_ {
		case -1:
			type_to_fetch = type_param

			/* skip over the rest and go directly to additional records */

			type_ = PHP_DNS_NUM_TYPES - 1
		case 0:
			if (type_param & PHP_DNS_A) != 0 {
				type_to_fetch = DNS_T_A
			} else {
				type_to_fetch = 0
			}
		case 1:
			if (type_param & PHP_DNS_NS) != 0 {
				type_to_fetch = DNS_T_NS
			} else {
				type_to_fetch = 0
			}
		case 2:
			if (type_param & PHP_DNS_CNAME) != 0 {
				type_to_fetch = DNS_T_CNAME
			} else {
				type_to_fetch = 0
			}
		case 3:
			if (type_param & PHP_DNS_SOA) != 0 {
				type_to_fetch = DNS_T_SOA
			} else {
				type_to_fetch = 0
			}
		case 4:
			if (type_param & PHP_DNS_PTR) != 0 {
				type_to_fetch = DNS_T_PTR
			} else {
				type_to_fetch = 0
			}
		case 5:
			if (type_param & PHP_DNS_HINFO) != 0 {
				type_to_fetch = DNS_T_HINFO
			} else {
				type_to_fetch = 0
			}
		case 6:
			if (type_param & PHP_DNS_MX) != 0 {
				type_to_fetch = DNS_T_MX
			} else {
				type_to_fetch = 0
			}
		case 7:
			if (type_param & PHP_DNS_TXT) != 0 {
				type_to_fetch = DNS_T_TXT
			} else {
				type_to_fetch = 0
			}
		case 8:
			if (type_param & PHP_DNS_AAAA) != 0 {
				type_to_fetch = DNS_T_AAAA
			} else {
				type_to_fetch = 0
			}
		case 9:
			if (type_param & PHP_DNS_SRV) != 0 {
				type_to_fetch = DNS_T_SRV
			} else {
				type_to_fetch = 0
			}
		case 10:
			if (type_param & PHP_DNS_NAPTR) != 0 {
				type_to_fetch = DNS_T_NAPTR
			} else {
				type_to_fetch = 0
			}
		case 11:
			if (type_param & PHP_DNS_A6) != 0 {
				type_to_fetch = DNS_T_A6
			} else {
				type_to_fetch = 0
			}
		case 12:
			if (type_param & PHP_DNS_CAA) != 0 {
				type_to_fetch = DNS_T_CAA
			} else {
				type_to_fetch = 0
			}
		case PHP_DNS_NUM_TYPES:
			store_results = 0
			continue
		default:
			fallthrough
		case PHP_DNS_NUM_TYPES + 1:
			type_to_fetch = DNS_T_ANY
		}
		if type_to_fetch != 0 {
			handle = dns_open(nil)
			if handle == nil {
				return_value.GetArr().DestroyEx()
				return_value.SetFalse()
				return
			}
			n = PhpDnsSearch(handle, hostname, C_IN, type_to_fetch, answer.GetQb2(), b.SizeOf(answer))
			if n < 0 {
				dns_errno = PhpDnsErrno(handle)
				PhpDnsFreeHandle(handle)
				switch dns_errno {
				case NO_DATA:
					fallthrough
				case HOST_NOT_FOUND:
					continue
				case NO_RECOVERY:
					core.PhpErrorDocref(nil, zend.E_WARNING, "An unexpected server failure occurred.")
				case TRY_AGAIN:
					core.PhpErrorDocref(nil, zend.E_WARNING, "A temporary server error occurred.")
				default:
					core.PhpErrorDocref(nil, zend.E_WARNING, "DNS Query failed")
				}
				return_value.GetArr().DestroyEx()
				return_value.SetFalse()
				return
			}
			cp = answer.GetQb2() + HFIXEDSZ
			end = answer.GetQb2() + n
			hp = (*HEADER)(&answer)
			qd = ntohs(hp.qdcount)
			an = ntohs(hp.ancount)
			ns = ntohs(hp.nscount)
			ar = ntohs(hp.arcount)

			/* Skip QD entries, they're only used by dn_expand later on */

			for b.PostDec(&qd) > 0 {
				n = dn_skipname(cp, end)
				if n < 0 {
					core.PhpErrorDocref(nil, zend.E_WARNING, "Unable to parse DNS data received")
					return_value.GetArr().DestroyEx()
					PhpDnsFreeHandle(handle)
					return_value.SetFalse()
					return
				}
				cp += n + QFIXEDSZ
			}

			/* YAY! Our real answers! */

			for b.PostDec(&an) && cp != nil && cp < end {
				var retval zend.Zval
				cp = PhpParserr(cp, end, &answer, type_to_fetch, store_results, raw, &retval)
				if retval.GetType() != zend.IS_UNDEF && store_results != 0 {
					zend.AddNextIndexZval(return_value, &retval)
				}
			}
			if authns != nil || addtl != nil {

				/* List of Authoritative Name Servers
				 * Process when only requesting addtl so that we can skip through the section
				 */

				for b.PostDec(&ns) > 0 && cp != nil && cp < end {
					var retval zend.Zval
					cp = PhpParserr(cp, end, &answer, DNS_T_ANY, authns != nil, raw, &retval)
					if retval.GetType() != zend.IS_UNDEF {
						zend.AddNextIndexZval(authns, &retval)
					}
				}

				/* List of Authoritative Name Servers
				 * Process when only requesting addtl so that we can skip through the section
				 */

			}
			if addtl != nil {

				/* Additional records associated with authoritative name servers */

				for b.PostDec(&ar) > 0 && cp != nil && cp < end {
					var retval zend.Zval
					cp = PhpParserr(cp, end, &answer, DNS_T_ANY, 1, raw, &retval)
					if retval.GetType() != zend.IS_UNDEF {
						zend.AddNextIndexZval(addtl, &retval)
					}
				}

				/* Additional records associated with authoritative name servers */

			}
			PhpDnsFreeHandle(handle)
		}
	}
}
func ZifDnsGetMx(executeData *zend.ZendExecuteData, return_value *zend.Zval) {
	var hostname *byte
	var hostname_len int
	var mx_list *zend.Zval
	var weight_list *zend.Zval = nil
	var count int
	var qdc int
	var type_ u_short
	var weight u_short
	var answer Querybuf
	var buf []byte
	var hp *HEADER
	var cp *u_char
	var end *u_char
	var i int
	var from __struct__sockaddr_storage
	var fromsize uint32 = b.SizeOf("from")
	var handle dns_handle_t
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
						zend.CheckNumArgsException(_min_num_args, _max_num_args)
					} else {
						zend.CheckNumArgsError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(executeData, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgString(_arg, &hostname, &hostname_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			zend.ZendParseArgZvalDeref(_arg, &mx_list, 0)
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			zend.ZendParseArgZvalDeref(_arg, &weight_list, 0)
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
	mx_list = zend.ZendTryArrayInit(mx_list)
	if mx_list == nil {
		return
	}
	if weight_list != nil {
		weight_list = zend.ZendTryArrayInit(weight_list)
		if weight_list == nil {
			return
		}
	}
	handle = dns_open(nil)
	if handle == nil {
		return_value.SetFalse()
		return
	}
	i = PhpDnsSearch(handle, hostname, C_IN, DNS_T_MX, answer.GetQb2(), b.SizeOf(answer))
	if i < 0 {
		PhpDnsFreeHandle(handle)
		return_value.SetFalse()
		return
	}
	hp = (*HEADER)(&answer)
	cp = answer.GetQb2() + HFIXEDSZ
	end = answer.GetQb2() + i
	for qdc = ntohs(uint16(hp.qdcount)); b.PostDec(&qdc); cp += i + QFIXEDSZ {
		if b.Assign(&i, dn_skipname(cp, end)) < 0 {
			PhpDnsFreeHandle(handle)
			return_value.SetFalse()
			return
		}
	}
	count = ntohs(uint16(hp.ancount))
	for b.PreDec(&count) >= 0 && cp < end {
		if b.Assign(&i, dn_skipname(cp, end)) < 0 {
			PhpDnsFreeHandle(handle)
			return_value.SetFalse()
			return
		}
		cp += i
		GETSHORT(type_, cp)
		cp += INT16SZ + INT32SZ
		GETSHORT(i, cp)
		if type_ != DNS_T_MX {
			cp += i
			continue
		}
		GETSHORT(weight, cp)
		if b.Assign(&i, dn_expand(answer.GetQb2(), end, cp, buf, b.SizeOf("buf")-1)) < 0 {
			PhpDnsFreeHandle(handle)
			return_value.SetFalse()
			return
		}
		cp += i
		zend.AddNextIndexString(mx_list, buf)
		if weight_list != nil {
			zend.AddNextIndexLong(weight_list, weight)
		}
	}
	PhpDnsFreeHandle(handle)
	zend.ZVAL_BOOL(return_value, zend.Z_ARRVAL_P(mx_list).GetNNumOfElements() != 0)
	return
}
func ZmStartupDns(type_ int, module_number int) int {
	zend.REGISTER_LONG_CONSTANT("DNS_A", PHP_DNS_A, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("DNS_NS", PHP_DNS_NS, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("DNS_CNAME", PHP_DNS_CNAME, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("DNS_SOA", PHP_DNS_SOA, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("DNS_PTR", PHP_DNS_PTR, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("DNS_HINFO", PHP_DNS_HINFO, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("DNS_CAA", PHP_DNS_CAA, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("DNS_MX", PHP_DNS_MX, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("DNS_TXT", PHP_DNS_TXT, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("DNS_SRV", PHP_DNS_SRV, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("DNS_NAPTR", PHP_DNS_NAPTR, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("DNS_AAAA", PHP_DNS_AAAA, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("DNS_A6", PHP_DNS_A6, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("DNS_ANY", PHP_DNS_ANY, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("DNS_ALL", PHP_DNS_ALL, zend.CONST_CS|zend.CONST_PERSISTENT)
	return zend.SUCCESS
}
