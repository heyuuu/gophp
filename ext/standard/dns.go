// <<generate>>

package standard

import (
	"sik/core"
	g "sik/runtime/grammar"
	"sik/zend"
)

// Source: <ext/standard/dns.c>

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
   | Authors: The typical suspects                                        |
   |          Pollita <pollita@php.net>                                   |
   |          Marcus Boerger <helly@php.net>                              |
   +----------------------------------------------------------------------+
*/

// # include "php.h"

// # include "php_network.h"

// # include < sys / socket . h >

// # include < netinet / in . h >

// # include < arpa / inet . h >

// # include < netdb . h >

// #define BIND_8_COMPAT       1

// # include < arpa / nameser . h >

// # include < resolv . h >

// # include < dns . h >

// #define MAXHOSTNAMELEN       255

/* For the local hostname obtained via gethostname which is different from the
   dns-related MAXHOSTNAMELEN constant above */

// #define HOST_NAME_MAX       255

// # include "php_dns.h"

/* type compat */

// #define DNS_T_A       1

// #define DNS_T_NS       2

// #define DNS_T_CNAME       5

// #define DNS_T_SOA       6

// #define DNS_T_PTR       12

// #define DNS_T_HINFO       13

// #define DNS_T_MINFO       14

// #define DNS_T_MX       15

// #define DNS_T_TXT       16

// #define DNS_T_AAAA       28

// #define DNS_T_SRV       33

// #define DNS_T_NAPTR       35

// #define DNS_T_A6       38

// #define DNS_T_CAA       257

// #define DNS_T_ANY       255

/* }}} */

/* {{{ proto string gethostname()
   Get the host name of the current machine */

func ZifGethostname(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var buf []byte
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	if gethostname(buf, g.SizeOf("buf")) {
		core.PhpErrorDocref(nil, 1<<1, "unable to fetch host [%d]: %s", errno, strerror(errno))
		return_value.u1.type_info = 2
		return
	}
	var _s *byte = buf
	var __z *zend.Zval = return_value
	var __s *zend.ZendString = zend.ZendStringInit(_s, strlen(_s), 0)
	__z.value.str = __s
	__z.u1.type_info = 6 | 1<<0<<8
	return
}

/* }}} */

/* TODO: Reimplement the gethostby* functions using the new winxp+ API, in dns_win32.c, then
we can have a dns.c, dns_unix.c and dns_win32.c instead of a messy dns.c full of #ifdef
*/

func ZifGethostbyaddr(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var addr *byte
	var addr_len int
	var hostname *zend.ZendString
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

			if zend.ZendParseArgString(_arg, &addr, &addr_len, 0) == 0 {
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
	hostname = PhpGethostbyaddr(addr)
	if hostname == nil {
		core.PhpErrorDocref(nil, 1<<1, "Address is not a valid IPv4 or IPv6 address")
		return_value.u1.type_info = 2
	} else {
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = hostname
		__z.value.str = __s
		if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
			__z.u1.type_info = 6
		} else {
			__z.u1.type_info = 6 | 1<<0<<8
		}
	}
}

/* }}} */

func PhpGethostbyaddr(ip *byte) *zend.ZendString {
	var addr6 __struct__in6_addr
	var addr __struct__in_addr
	var hp *__struct__hostent
	if inet_pton(AF_INET6, ip, &addr6) {
		hp = gethostbyaddr((*byte)(&addr6), g.SizeOf("addr6"), AF_INET6)
	} else if inet_pton(AF_INET, ip, &addr) {
		hp = gethostbyaddr((*byte)(&addr), g.SizeOf("addr"), AF_INET)
	} else {
		return nil
	}
	if hp == nil || hp.h_name == nil || hp.h_name[0] == '0' {
		return zend.ZendStringInit(ip, strlen(ip), 0)
	}
	return zend.ZendStringInit(hp.h_name, strlen(hp.h_name), 0)
}

/* }}} */

func ZifGethostbyname(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var hostname *byte
	var hostname_len int
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

			if zend.ZendParseArgString(_arg, &hostname, &hostname_len, 0) == 0 {
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
	if hostname_len > 255 {

		/* name too long, protect from CVE-2015-0235 */

		core.PhpErrorDocref(nil, 1<<1, "Host name is too long, the limit is %d characters", 255)
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = zend.ZendStringInit(hostname, hostname_len, 0)
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
		return
	}
	var __z *zend.Zval = return_value
	var __s *zend.ZendString = PhpGethostbyname(hostname)
	__z.value.str = __s
	if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
		__z.u1.type_info = 6
	} else {
		__z.u1.type_info = 6 | 1<<0<<8
	}
	return
}

/* }}} */

func ZifGethostbynamel(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var hostname *byte
	var hostname_len int
	var hp *__struct__hostent
	var in __struct__in_addr
	var i int
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

			if zend.ZendParseArgString(_arg, &hostname, &hostname_len, 0) == 0 {
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
	if hostname_len > 255 {

		/* name too long, protect from CVE-2015-0235 */

		core.PhpErrorDocref(nil, 1<<1, "Host name is too long, the limit is %d characters", 255)
		return_value.u1.type_info = 2
		return
	}
	hp = core.PhpNetworkGethostbyname(hostname)
	if hp == nil {
		return_value.u1.type_info = 2
		return
	}
	var __arr *zend.ZendArray = zend._zendNewArray(0)
	var __z *zend.Zval = return_value
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	for i = 0; ; i++ {

		/* On macos h_addr_list entries may be misaligned. */

		var h_addr_entry *__struct__in_addr
		memcpy(&h_addr_entry, &hp.h_addr_list[i], g.SizeOf("struct in_addr *"))
		if h_addr_entry == nil {
			return
		}
		in = *h_addr_entry
		zend.AddNextIndexString(return_value, inet_ntoa(in))
	}
}

/* }}} */

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

	memcpy(&h_addr_0, &hp.h_addr_list[0], g.SizeOf("struct in_addr *"))
	if h_addr_0 == nil {
		return zend.ZendStringInit(name, strlen(name), 0)
	}
	memcpy(&in.s_addr, h_addr_0, g.SizeOf("in . s_addr"))
	address = inet_ntoa(in)
	return zend.ZendStringInit(address, strlen(address), 0)
}

/* }}} */

// #define PHP_DNS_NUM_TYPES       13

// #define PHP_DNS_A       0x00000001

// #define PHP_DNS_NS       0x00000002

// #define PHP_DNS_CNAME       0x00000010

// #define PHP_DNS_SOA       0x00000020

// #define PHP_DNS_PTR       0x00000800

// #define PHP_DNS_HINFO       0x00001000

// #define PHP_DNS_CAA       0x00002000

// #define PHP_DNS_MX       0x00004000

// #define PHP_DNS_TXT       0x00008000

// #define PHP_DNS_A6       0x01000000

// #define PHP_DNS_SRV       0x02000000

// #define PHP_DNS_NAPTR       0x04000000

// #define PHP_DNS_AAAA       0x08000000

// #define PHP_DNS_ANY       0x10000000

// #define PHP_DNS_ALL       ( PHP_DNS_A | PHP_DNS_NS | PHP_DNS_CNAME | PHP_DNS_SOA | PHP_DNS_PTR | PHP_DNS_HINFO | PHP_DNS_CAA | PHP_DNS_MX | PHP_DNS_TXT | PHP_DNS_A6 | PHP_DNS_SRV | PHP_DNS_NAPTR | PHP_DNS_AAAA )

/* Note: These functions are defined in ext/standard/dns_win32.c for Windows! */

// #define HFIXEDSZ       12

// #define QFIXEDSZ       4

// #define MAXHOSTNAMELEN       1024

// #define MAXRESOURCERECORDS       64

// @type Querybuf struct

/* just a hack to free resources allocated by glibc in __res_nsend()
 * See also:
 *   res_thread_freeres() in glibc/resolv/res_init.c
 *   __libc_res_nsend()   in resolv/res_send.c
 * */

// #define php_dns_free_res(__res__)

/* {{{ proto bool dns_check_record(string host [, string type])
   Check DNS records corresponding to a given Internet host name or IP address */

func ZifDnsCheckRecord(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var hp *HEADER
	var answer Querybuf
	var hostname *byte
	var rectype *byte = nil
	var hostname_len int
	var rectype_len int = 0
	var type_ int = 15
	var i int
	var from __struct__sockaddr_storage
	var fromsize uint32 = g.SizeOf("from")
	var handle dns_handle_t
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
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgString(_arg, &hostname, &hostname_len, 0) == 0 {
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

			if zend.ZendParseArgString(_arg, &rectype, &rectype_len, 0) == 0 {
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
	if hostname_len == 0 {
		core.PhpErrorDocref(nil, 1<<1, "Host cannot be empty")
		return_value.u1.type_info = 2
		return
	}
	if rectype != nil {
		if !(strcasecmp("A", rectype)) {
			type_ = 1
		} else if !(strcasecmp("NS", rectype)) {
			type_ = 2
		} else if !(strcasecmp("MX", rectype)) {
			type_ = 15
		} else if !(strcasecmp("PTR", rectype)) {
			type_ = 12
		} else if !(strcasecmp("ANY", rectype)) {
			type_ = 255
		} else if !(strcasecmp("SOA", rectype)) {
			type_ = 6
		} else if !(strcasecmp("CAA", rectype)) {
			type_ = 257
		} else if !(strcasecmp("TXT", rectype)) {
			type_ = 16
		} else if !(strcasecmp("CNAME", rectype)) {
			type_ = 5
		} else if !(strcasecmp("AAAA", rectype)) {
			type_ = 28
		} else if !(strcasecmp("SRV", rectype)) {
			type_ = 33
		} else if !(strcasecmp("NAPTR", rectype)) {
			type_ = 35
		} else if !(strcasecmp("A6", rectype)) {
			type_ = 38
		} else {
			core.PhpErrorDocref(nil, 1<<1, "Type '%s' not supported", rectype)
			return_value.u1.type_info = 2
			return
		}
	}
	handle = dns_open(nil)
	if handle == nil {
		return_value.u1.type_info = 2
		return
	}
	i = int(dns_search(handle, hostname, C_IN, type_, (*byte)(answer.GetQb2()), g.SizeOf(answer), (*__struct__sockaddr)(&from), &fromsize))
	dns_free(handle)
	if i < 0 {
		return_value.u1.type_info = 2
		return
	}
	hp = (*HEADER)(&answer)
	if ntohs(hp.ancount) != 0 {
		return_value.u1.type_info = 3
	} else {
		return_value.u1.type_info = 2
	}
	return
}

/* }}} */

// #define CHECKCP(n) do { if ( cp + n > end ) { return NULL ; } } while ( 0 )

/* {{{ php_parserr */

func PhpParserr(cp *u_char, end *u_char, answer *Querybuf, type_to_fetch int, store int, raw int, subarray *zend.Zval) *u_char {
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
	subarray.u1.type_info = 0
	n = dn_expand(answer.GetQb2(), end, cp, name, g.SizeOf("name")-2)
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
	if type_to_fetch != 255 && type_ != type_to_fetch {
		cp += dlen
		return cp
	}
	if store == 0 {
		cp += dlen
		return cp
	}
	var __arr *zend.ZendArray = zend._zendNewArray(0)
	var __z *zend.Zval = subarray
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	zend.AddAssocStringEx(subarray, "host", strlen("host"), name)
	zend.AddAssocStringEx(subarray, "class", strlen("class"), "IN")
	zend.AddAssocLongEx(subarray, "ttl", strlen("ttl"), ttl)
	void(class)
	if raw != 0 {
		zend.AddAssocLongEx(subarray, "type", strlen("type"), type_)
		zend.AddAssocStringlEx(subarray, "data", strlen("data"), (*byte)(cp), uint32(dlen))
		cp += dlen
		return cp
	}
	switch type_ {
	case 1:
		if cp+4 > end {
			return nil
		}
		zend.AddAssocStringEx(subarray, "type", strlen("type"), "A")
		core.ApPhpSnprintf(name, g.SizeOf("name"), "%d.%d.%d.%d", cp[0], cp[1], cp[2], cp[3])
		zend.AddAssocStringEx(subarray, "ip", strlen("ip"), name)
		cp += dlen
		break
	case 15:
		if cp+2 > end {
			return nil
		}
		zend.AddAssocStringEx(subarray, "type", strlen("type"), "MX")
		GETSHORT(n, cp)
		zend.AddAssocLongEx(subarray, "pri", strlen("pri"), n)
	case 5:
		if type_ == 5 {
			zend.AddAssocStringEx(subarray, "type", strlen("type"), "CNAME")
		}
	case 2:
		if type_ == 2 {
			zend.AddAssocStringEx(subarray, "type", strlen("type"), "NS")
		}
	case 12:
		if type_ == 12 {
			zend.AddAssocStringEx(subarray, "type", strlen("type"), "PTR")
		}
		n = dn_expand(answer.GetQb2(), end, cp, name, g.SizeOf(name)-2)
		if n < 0 {
			return nil
		}
		cp += n
		zend.AddAssocStringEx(subarray, "target", strlen("target"), name)
		break
	case 13:

		/* See RFC 1010 for values */

		zend.AddAssocStringEx(subarray, "type", strlen("type"), "HINFO")
		if cp+1 > end {
			return nil
		}
		n = (*cp) & 0xff
		cp++
		if cp+n > end {
			return nil
		}
		zend.AddAssocStringlEx(subarray, "cpu", strlen("cpu"), (*byte)(cp), n)
		cp += n
		if cp+1 > end {
			return nil
		}
		n = (*cp) & 0xff
		cp++
		if cp+n > end {
			return nil
		}
		zend.AddAssocStringlEx(subarray, "os", strlen("os"), (*byte)(cp), n)
		cp += n
		break
	case 257:

		/* See RFC 6844 for values https://tools.ietf.org/html/rfc6844 */

		zend.AddAssocStringEx(subarray, "type", strlen("type"), "CAA")

		// 1 flag byte

		if cp+1 > end {
			return nil
		}
		n = (*cp) & 0xff
		zend.AddAssocLongEx(subarray, "flags", strlen("flags"), n)
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
		zend.AddAssocStringlEx(subarray, "tag", strlen("tag"), (*byte)(cp), n)
		cp += n
		if int(dlen < int(n)+2) != 0 {
			return nil
		}
		n = dlen - n - 2
		if cp+n > end {
			return nil
		}
		zend.AddAssocStringlEx(subarray, "value", strlen("value"), (*byte)(cp), n)
		cp += n
		break
	case 16:
		var l1 int = 0
		var l2 int = 0
		var entries zend.Zval
		var tp *zend.ZendString
		zend.AddAssocStringEx(subarray, "type", strlen("type"), "TXT")
		tp = zend.ZendStringAlloc(dlen, 0)
		var __arr *zend.ZendArray = zend._zendNewArray(0)
		var __z *zend.Zval = &entries
		__z.value.arr = __arr
		__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
		for l1 < dlen {
			n = cp[l1]
			if l1+n >= dlen {

				// Invalid chunk length, truncate

				n = dlen - (l1 + 1)

				// Invalid chunk length, truncate

			}
			if n {
				memcpy(tp.val+l2, cp+l1+1, n)
				zend.AddNextIndexStringl(&entries, (*byte)(cp+l1+1), n)
			}
			l1 = l1 + n + 1
			l2 = l2 + n
		}
		tp.val[l2] = '0'
		tp.len_ = l2
		cp += dlen
		zend.AddAssocStrEx(subarray, "txt", strlen("txt"), tp)
		zend.AddAssocZvalEx(subarray, "entries", strlen("entries"), &entries)
		break
	case 6:
		zend.AddAssocStringEx(subarray, "type", strlen("type"), "SOA")
		n = dn_expand(answer.GetQb2(), end, cp, name, g.SizeOf(name)-2)
		if n < 0 {
			return nil
		}
		cp += n
		zend.AddAssocStringEx(subarray, "mname", strlen("mname"), name)
		n = dn_expand(answer.GetQb2(), end, cp, name, g.SizeOf(name)-2)
		if n < 0 {
			return nil
		}
		cp += n
		zend.AddAssocStringEx(subarray, "rname", strlen("rname"), name)
		if cp+5*4 > end {
			return nil
		}
		GETLONG(n, cp)
		zend.AddAssocLongEx(subarray, "serial", strlen("serial"), n)
		GETLONG(n, cp)
		zend.AddAssocLongEx(subarray, "refresh", strlen("refresh"), n)
		GETLONG(n, cp)
		zend.AddAssocLongEx(subarray, "retry", strlen("retry"), n)
		GETLONG(n, cp)
		zend.AddAssocLongEx(subarray, "expire", strlen("expire"), n)
		GETLONG(n, cp)
		zend.AddAssocLongEx(subarray, "minimum-ttl", strlen("minimum-ttl"), n)
		break
	case 28:
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
		zend.AddAssocStringEx(subarray, "type", strlen("type"), "AAAA")
		zend.AddAssocStringEx(subarray, "ipv6", strlen("ipv6"), name)
		break
	case 38:
		p = cp
		zend.AddAssocStringEx(subarray, "type", strlen("type"), "A6")
		if cp+1 > end {
			return nil
		}
		n = int(cp[0]) & 0xff
		cp++
		zend.AddAssocLongEx(subarray, "masklen", strlen("masklen"), n)
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
		zend.AddAssocStringEx(subarray, "ipv6", strlen("ipv6"), name)
		if cp < p+dlen {
			n = dn_expand(answer.GetQb2(), end, cp, name, g.SizeOf(name)-2)
			if n < 0 {
				return nil
			}
			cp += n
			zend.AddAssocStringEx(subarray, "chain", strlen("chain"), name)
		}
		break
	case 33:
		if cp+3*2 > end {
			return nil
		}
		zend.AddAssocStringEx(subarray, "type", strlen("type"), "SRV")
		GETSHORT(n, cp)
		zend.AddAssocLongEx(subarray, "pri", strlen("pri"), n)
		GETSHORT(n, cp)
		zend.AddAssocLongEx(subarray, "weight", strlen("weight"), n)
		GETSHORT(n, cp)
		zend.AddAssocLongEx(subarray, "port", strlen("port"), n)
		n = dn_expand(answer.GetQb2(), end, cp, name, g.SizeOf(name)-2)
		if n < 0 {
			return nil
		}
		cp += n
		zend.AddAssocStringEx(subarray, "target", strlen("target"), name)
		break
	case 35:
		if cp+2*2 > end {
			return nil
		}
		zend.AddAssocStringEx(subarray, "type", strlen("type"), "NAPTR")
		GETSHORT(n, cp)
		zend.AddAssocLongEx(subarray, "order", strlen("order"), n)
		GETSHORT(n, cp)
		zend.AddAssocLongEx(subarray, "pref", strlen("pref"), n)
		if cp+1 > end {
			return nil
		}
		n = cp[0] & 0xff
		cp++
		if cp+n > end {
			return nil
		}
		zend.AddAssocStringlEx(subarray, "flags", strlen("flags"), (*byte)(cp), n)
		cp += n
		if cp+1 > end {
			return nil
		}
		n = cp[0] & 0xff
		cp++
		if cp+n > end {
			return nil
		}
		zend.AddAssocStringlEx(subarray, "services", strlen("services"), (*byte)(cp), n)
		cp += n
		if cp+1 > end {
			return nil
		}
		n = cp[0] & 0xff
		cp++
		if cp+n > end {
			return nil
		}
		zend.AddAssocStringlEx(subarray, "regex", strlen("regex"), (*byte)(cp), n)
		cp += n
		n = dn_expand(answer.GetQb2(), end, cp, name, g.SizeOf(name)-2)
		if n < 0 {
			return nil
		}
		cp += n
		zend.AddAssocStringEx(subarray, "replacement", strlen("replacement"), name)
		break
	default:
		zend.ZvalPtrDtor(subarray)
		subarray.u1.type_info = 0
		cp += dlen
		break
	}
	return cp
}

/* }}} */

func ZifDnsGetRecord(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var hostname *byte
	var hostname_len int
	var type_param zend.ZendLong = 0x10000000
	var authns *zend.Zval = nil
	var addtl *zend.Zval = nil
	var type_to_fetch int
	var dns_errno int
	var from __struct__sockaddr_storage
	var fromsize uint32 = g.SizeOf("from")
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

			if zend.ZendParseArgString(_arg, &hostname, &hostname_len, 0) == 0 {
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

			if zend.ZendParseArgLong(_arg, &type_param, &_dummy, 0, 0) == 0 {
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

			zend.ZendParseArgZvalDeref(_arg, &authns, 0)
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

			zend.ZendParseArgZvalDeref(_arg, &addtl, 0)
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

			if zend.ZendParseArgBool(_arg, &raw, &_dummy, 0) == 0 {
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
		if (type_param & ^(0x1|0x2|0x10|0x20|0x800|0x1000|0x2000|0x4000|0x8000|0x1000000|0x2000000|0x4000000|0x8000000)) != 0 && type_param != 0x10000000 {
			core.PhpErrorDocref(nil, 1<<1, "Type '"+"%"+"lld"+"' not supported", type_param)
			return_value.u1.type_info = 2
			return
		}
	} else {
		if type_param < 1 || type_param > 0xffff {
			core.PhpErrorDocref(nil, 1<<1, "Numeric DNS record type must be between 1 and 65535, '"+"%"+"lld"+"' given", type_param)
			return_value.u1.type_info = 2
			return
		}
	}

	/* Initialize the return array */

	var __arr *zend.ZendArray = zend._zendNewArray(0)
	var __z *zend.Zval = return_value
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8

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
	} else if type_param == 0x10000000 {
		type_ = 13 + 1
	} else {
		type_ = 0
	}
	for ; type_ < g.Cond(addtl != nil, 13+2, 13) || first_query != 0; type_++ {
		first_query = 0
		switch type_ {
		case -1:
			type_to_fetch = type_param

			/* skip over the rest and go directly to additional records */

			type_ = 13 - 1
			break
		case 0:
			if (type_param & 0x1) != 0 {
				type_to_fetch = 1
			} else {
				type_to_fetch = 0
			}
			break
		case 1:
			if (type_param & 0x2) != 0 {
				type_to_fetch = 2
			} else {
				type_to_fetch = 0
			}
			break
		case 2:
			if (type_param & 0x10) != 0 {
				type_to_fetch = 5
			} else {
				type_to_fetch = 0
			}
			break
		case 3:
			if (type_param & 0x20) != 0 {
				type_to_fetch = 6
			} else {
				type_to_fetch = 0
			}
			break
		case 4:
			if (type_param & 0x800) != 0 {
				type_to_fetch = 12
			} else {
				type_to_fetch = 0
			}
			break
		case 5:
			if (type_param & 0x1000) != 0 {
				type_to_fetch = 13
			} else {
				type_to_fetch = 0
			}
			break
		case 6:
			if (type_param & 0x4000) != 0 {
				type_to_fetch = 15
			} else {
				type_to_fetch = 0
			}
			break
		case 7:
			if (type_param & 0x8000) != 0 {
				type_to_fetch = 16
			} else {
				type_to_fetch = 0
			}
			break
		case 8:
			if (type_param & 0x8000000) != 0 {
				type_to_fetch = 28
			} else {
				type_to_fetch = 0
			}
			break
		case 9:
			if (type_param & 0x2000000) != 0 {
				type_to_fetch = 33
			} else {
				type_to_fetch = 0
			}
			break
		case 10:
			if (type_param & 0x4000000) != 0 {
				type_to_fetch = 35
			} else {
				type_to_fetch = 0
			}
			break
		case 11:
			if (type_param & 0x1000000) != 0 {
				type_to_fetch = 38
			} else {
				type_to_fetch = 0
			}
			break
		case 12:
			if (type_param & 0x2000) != 0 {
				type_to_fetch = 257
			} else {
				type_to_fetch = 0
			}
			break
		case 13:
			store_results = 0
			continue
		default:

		case 13 + 1:
			type_to_fetch = 255
			break
		}
		if type_to_fetch != 0 {
			handle = dns_open(nil)
			if handle == nil {
				zend.ZendArrayDestroy(return_value.value.arr)
				return_value.u1.type_info = 2
				return
			}
			n = int(dns_search(handle, hostname, C_IN, type_to_fetch, (*byte)(answer.GetQb2()), g.SizeOf(answer), (*__struct__sockaddr)(&from), &fromsize))
			if n < 0 {
				dns_errno = h_errno
				dns_free(handle)
				switch dns_errno {
				case NO_DATA:

				case HOST_NOT_FOUND:
					continue
				case NO_RECOVERY:
					core.PhpErrorDocref(nil, 1<<1, "An unexpected server failure occurred.")
					break
				case TRY_AGAIN:
					core.PhpErrorDocref(nil, 1<<1, "A temporary server error occurred.")
					break
				default:
					core.PhpErrorDocref(nil, 1<<1, "DNS Query failed")
				}
				zend.ZendArrayDestroy(return_value.value.arr)
				return_value.u1.type_info = 2
				return
			}
			cp = answer.GetQb2() + 12
			end = answer.GetQb2() + n
			hp = (*HEADER)(&answer)
			qd = ntohs(hp.qdcount)
			an = ntohs(hp.ancount)
			ns = ntohs(hp.nscount)
			ar = ntohs(hp.arcount)

			/* Skip QD entries, they're only used by dn_expand later on */

			for g.PostDec(&qd) > 0 {
				n = dn_skipname(cp, end)
				if n < 0 {
					core.PhpErrorDocref(nil, 1<<1, "Unable to parse DNS data received")
					zend.ZendArrayDestroy(return_value.value.arr)
					dns_free(handle)
					return_value.u1.type_info = 2
					return
				}
				cp += n + 4
			}

			/* YAY! Our real answers! */

			for g.PostDec(&an) && cp != nil && cp < end {
				var retval zend.Zval
				cp = PhpParserr(cp, end, &answer, type_to_fetch, store_results, raw, &retval)
				if retval.u1.v.type_ != 0 && store_results != 0 {
					zend.AddNextIndexZval(return_value, &retval)
				}
			}
			if authns != nil || addtl != nil {

				/* List of Authoritative Name Servers
				 * Process when only requesting addtl so that we can skip through the section
				 */

				for g.PostDec(&ns) > 0 && cp != nil && cp < end {
					var retval zend.Zval
					cp = PhpParserr(cp, end, &answer, 255, authns != nil, raw, &retval)
					if retval.u1.v.type_ != 0 {
						zend.AddNextIndexZval(authns, &retval)
					}
				}

				/* List of Authoritative Name Servers
				 * Process when only requesting addtl so that we can skip through the section
				 */

			}
			if addtl != nil {

				/* Additional records associated with authoritative name servers */

				for g.PostDec(&ar) > 0 && cp != nil && cp < end {
					var retval zend.Zval
					cp = PhpParserr(cp, end, &answer, 255, 1, raw, &retval)
					if retval.u1.v.type_ != 0 {
						zend.AddNextIndexZval(addtl, &retval)
					}
				}

				/* Additional records associated with authoritative name servers */

			}
			dns_free(handle)
		}
	}
}

/* }}} */

func ZifDnsGetMx(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
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
	var fromsize uint32 = g.SizeOf("from")
	var handle dns_handle_t
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

			if zend.ZendParseArgString(_arg, &hostname, &hostname_len, 0) == 0 {
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

			zend.ZendParseArgZvalDeref(_arg, &mx_list, 0)
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

			zend.ZendParseArgZvalDeref(_arg, &weight_list, 0)
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
		return_value.u1.type_info = 2
		return
	}
	i = int(dns_search(handle, hostname, C_IN, 15, (*byte)(answer.GetQb2()), g.SizeOf(answer), (*__struct__sockaddr)(&from), &fromsize))
	if i < 0 {
		dns_free(handle)
		return_value.u1.type_info = 2
		return
	}
	hp = (*HEADER)(&answer)
	cp = answer.GetQb2() + 12
	end = answer.GetQb2() + i
	for qdc = ntohs(uint16(hp.qdcount)); g.PostDec(&qdc); cp += i + 4 {
		if g.Assign(&i, dn_skipname(cp, end)) < 0 {
			dns_free(handle)
			return_value.u1.type_info = 2
			return
		}
	}
	count = ntohs(uint16(hp.ancount))
	for g.PreDec(&count) >= 0 && cp < end {
		if g.Assign(&i, dn_skipname(cp, end)) < 0 {
			dns_free(handle)
			return_value.u1.type_info = 2
			return
		}
		cp += i
		GETSHORT(type_, cp)
		cp += 2 + 4
		GETSHORT(i, cp)
		if type_ != 15 {
			cp += i
			continue
		}
		GETSHORT(weight, cp)
		if g.Assign(&i, dn_expand(answer.GetQb2(), end, cp, buf, g.SizeOf("buf")-1)) < 0 {
			dns_free(handle)
			return_value.u1.type_info = 2
			return
		}
		cp += i
		zend.AddNextIndexString(mx_list, buf)
		if weight_list != nil {
			zend.AddNextIndexLong(weight_list, weight)
		}
	}
	dns_free(handle)
	if mx_list.value.arr.nNumOfElements != 0 {
		return_value.u1.type_info = 3
	} else {
		return_value.u1.type_info = 2
	}
	return
}

/* }}} */

func ZmStartupDns(type_ int, module_number int) int {
	zend.ZendRegisterLongConstant("DNS_A", g.SizeOf("\"DNS_A\"")-1, 0x1, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("DNS_NS", g.SizeOf("\"DNS_NS\"")-1, 0x2, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("DNS_CNAME", g.SizeOf("\"DNS_CNAME\"")-1, 0x10, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("DNS_SOA", g.SizeOf("\"DNS_SOA\"")-1, 0x20, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("DNS_PTR", g.SizeOf("\"DNS_PTR\"")-1, 0x800, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("DNS_HINFO", g.SizeOf("\"DNS_HINFO\"")-1, 0x1000, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("DNS_CAA", g.SizeOf("\"DNS_CAA\"")-1, 0x2000, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("DNS_MX", g.SizeOf("\"DNS_MX\"")-1, 0x4000, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("DNS_TXT", g.SizeOf("\"DNS_TXT\"")-1, 0x8000, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("DNS_SRV", g.SizeOf("\"DNS_SRV\"")-1, 0x2000000, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("DNS_NAPTR", g.SizeOf("\"DNS_NAPTR\"")-1, 0x4000000, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("DNS_AAAA", g.SizeOf("\"DNS_AAAA\"")-1, 0x8000000, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("DNS_A6", g.SizeOf("\"DNS_A6\"")-1, 0x1000000, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("DNS_ANY", g.SizeOf("\"DNS_ANY\"")-1, 0x10000000, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("DNS_ALL", g.SizeOf("\"DNS_ALL\"")-1, 0x1|0x2|0x10|0x20|0x800|0x1000|0x2000|0x4000|0x8000|0x1000000|0x2000000|0x4000000|0x8000000, 1<<0|1<<1, module_number)
	return zend.SUCCESS
}
