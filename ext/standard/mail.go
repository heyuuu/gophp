// <<generate>>

package standard

import (
	"sik/core"
	"sik/core/streams"
	r "sik/runtime"
	g "sik/runtime/grammar"
	"sik/zend"
)

// Source: <ext/standard/mail.c>

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
   | Author: Rasmus Lerdorf <rasmus@php.net>                              |
   +----------------------------------------------------------------------+
*/

// # include < stdlib . h >

// # include < ctype . h >

// # include < stdio . h >

// # include < time . h >

// # include "php.h"

// # include "ext/standard/info.h"

// # include "ext/standard/php_string.h"

// # include "ext/standard/basic_functions.h"

// failed # include "ext/date/php_date.h"

// # include "zend_smart_str.h"

// # include < sysexits . h >

// # include "php_syslog.h"

// # include "php_mail.h"

// # include "php_ini.h"

// # include "php_string.h"

// # include "exec.h"

// #define SKIP_LONG_HEADER_SEP(str,pos) if ( str [ pos ] == '\r' && str [ pos + 1 ] == '\n' && ( str [ pos + 2 ] == ' ' || str [ pos + 2 ] == '\t' ) ) { pos += 2 ; while ( str [ pos + 1 ] == ' ' || str [ pos + 1 ] == '\t' ) { pos ++ ; } continue ; }

// #define MAIL_ASCIIZ_CHECK(str,len) p = str ; e = p + len ; while ( ( p = memchr ( p , '\0' , ( e - p ) ) ) ) { * p = ' ' ; }

/* {{{ proto int ezmlm_hash(string addr)
   Calculate EZMLM list hash value. */

func ZifEzmlmHash(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var str *byte = nil
	var h uint = 5381
	var j int
	var str_len int
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
			return
		}
		break
	}
	for j = 0; j < str_len; j++ {
		h = h + (h << 5) ^ zend.ZendUlong(uint8(tolower(str[j])))
	}
	h = h % 53
	var __z *zend.Zval = return_value
	__z.value.lval = zend.ZendLong(h)
	__z.u1.type_info = 4
	return
}

/* }}} */

func PhpMailBuildHeadersCheckFieldValue(val *zend.Zval) zend.ZendBool {
	var len_ int = 0
	var value *zend.ZendString = val.value.str

	/* https://tools.ietf.org/html/rfc2822#section-2.2.1 */

	for len_ < value.len_ {
		if (*(value.val + len_)) == '\r' {
			if value.len_-len_ >= 3 && (*(value.val + len_ + 1)) == '\n' && ((*(value.val + len_ + 2)) == ' ' || (*(value.val + len_ + 2)) == '\t') {
				len_ += 3
				continue
			}
			return zend.FAILURE
		}
		if (*(value.val + len_)) == '0' {
			return zend.FAILURE
		}
		len_++
	}
	return zend.SUCCESS
}
func PhpMailBuildHeadersCheckFieldName(key *zend.ZendString) zend.ZendBool {
	var len_ int = 0

	/* https://tools.ietf.org/html/rfc2822#section-2.2 */

	for len_ < key.len_ {
		if (*(key.val + len_)) < 33 || (*(key.val + len_)) > 126 || (*(key.val + len_)) == ':' {
			return zend.FAILURE
		}
		len_++
	}
	return zend.SUCCESS
}
func PhpMailBuildHeadersElem(s *zend.SmartStr, key *zend.ZendString, val *zend.Zval) {
	switch val.u1.v.type_ {
	case 6:
		if PhpMailBuildHeadersCheckFieldName(key) != zend.SUCCESS {
			core.PhpErrorDocref(nil, 1<<1, "Header field name (%s) contains invalid chars", key.val)
			return
		}
		if PhpMailBuildHeadersCheckFieldValue(val) != zend.SUCCESS {
			core.PhpErrorDocref(nil, 1<<1, "Header field value (%s => %s) contains invalid chars or format", key.val, val.value.str.val)
			return
		}
		zend.SmartStrAppendEx(s, key, 0)
		zend.SmartStrAppendlEx(s, ": ", 2, 0)
		zend.SmartStrAppendlEx(s, val.value.str.val, strlen(val.value.str.val), 0)
		zend.SmartStrAppendlEx(s, "\r\n", 2, 0)
		break
	case 7:
		PhpMailBuildHeadersElems(s, key, val)
		break
	default:
		core.PhpErrorDocref(nil, 1<<1, "headers array elements must be string or array (%s)", key.val)
	}
}
func PhpMailBuildHeadersElems(s *zend.SmartStr, key *zend.ZendString, val *zend.Zval) {
	var tmp_key *zend.ZendString
	var tmp_val *zend.Zval
	for {
		var __ht *zend.HashTable = val.value.arr
		var _p *zend.Bucket = __ht.arData
		var _end *zend.Bucket = _p + __ht.nNumUsed
		for ; _p != _end; _p++ {
			var _z *zend.Zval = &_p.val

			if _z.u1.v.type_ == 0 {
				continue
			}
			tmp_key = _p.key
			tmp_val = _z
			if tmp_key != nil {
				core.PhpErrorDocref(nil, 1<<1, "Multiple header key must be numeric index (%s)", tmp_key.val)
				continue
			}
			if tmp_val.u1.v.type_ != 6 {
				core.PhpErrorDocref(nil, 1<<1, "Multiple header values must be string (%s)", key.val)
				continue
			}
			PhpMailBuildHeadersElem(s, key, tmp_val)
		}
		break
	}
}
func PhpMailBuildHeaders(headers *zend.Zval) *zend.ZendString {
	var idx zend.ZendUlong
	var key *zend.ZendString
	var val *zend.Zval
	var s zend.SmartStr = zend.SmartStr{0}
	r.Assert(headers.u1.v.type_ == 7)
	for {
		var __ht *zend.HashTable = headers.value.arr
		var _p *zend.Bucket = __ht.arData
		var _end *zend.Bucket = _p + __ht.nNumUsed
		for ; _p != _end; _p++ {
			var _z *zend.Zval = &_p.val

			if _z.u1.v.type_ == 0 {
				continue
			}
			idx = _p.h
			key = _p.key
			val = _z
			if key == nil {
				core.PhpErrorDocref(nil, 1<<1, "Found numeric header ("+"%"+"lld"+")", idx)
				continue
			}

			/* https://tools.ietf.org/html/rfc2822#section-3.6 */

			switch key.len_ {
			case g.SizeOf("\"orig-date\"") - 1:
				if !(strncasecmp("orig-date", key.val, key.len_)) {
					for {
						if val.u1.v.type_ == 6 {
							PhpMailBuildHeadersElem(&s, key, val)
						} else if val.u1.v.type_ == 7 {
							if !(strncasecmp("orig-date", key.val, key.len_)) {
								core.PhpErrorDocref(nil, 1<<1, "'%s' header must be at most one header. Array is passed for '%s'", "orig-date", "orig-date")
								continue
							}
							PhpMailBuildHeadersElems(&s, key, val)
						} else {
							core.PhpErrorDocref(nil, 1<<1, "Extra header element '%s' cannot be other than string or array.", key.val)
						}
						break
					}
				} else {
					if val.u1.v.type_ == 6 {
						PhpMailBuildHeadersElem(&s, key, val)
					} else if val.u1.v.type_ == 7 {
						PhpMailBuildHeadersElems(&s, key, val)
					} else {
						core.PhpErrorDocref(nil, 1<<1, "Extra header element '%s' cannot be other than string or array.", key.val)
					}
				}
				break
			case g.SizeOf("\"from\"") - 1:
				if !(strncasecmp("from", key.val, key.len_)) {
					for {
						if val.u1.v.type_ == 6 {
							PhpMailBuildHeadersElem(&s, key, val)
						} else if val.u1.v.type_ == 7 {
							if !(strncasecmp("from", key.val, key.len_)) {
								core.PhpErrorDocref(nil, 1<<1, "'%s' header must be at most one header. Array is passed for '%s'", "from", "from")
								continue
							}
							PhpMailBuildHeadersElems(&s, key, val)
						} else {
							core.PhpErrorDocref(nil, 1<<1, "Extra header element '%s' cannot be other than string or array.", key.val)
						}
						break
					}
				} else {
					if val.u1.v.type_ == 6 {
						PhpMailBuildHeadersElem(&s, key, val)
					} else if val.u1.v.type_ == 7 {
						PhpMailBuildHeadersElems(&s, key, val)
					} else {
						core.PhpErrorDocref(nil, 1<<1, "Extra header element '%s' cannot be other than string or array.", key.val)
					}
				}
				break
			case g.SizeOf("\"sender\"") - 1:
				if !(strncasecmp("sender", key.val, key.len_)) {
					for {
						if val.u1.v.type_ == 6 {
							PhpMailBuildHeadersElem(&s, key, val)
						} else if val.u1.v.type_ == 7 {
							if !(strncasecmp("sender", key.val, key.len_)) {
								core.PhpErrorDocref(nil, 1<<1, "'%s' header must be at most one header. Array is passed for '%s'", "sender", "sender")
								continue
							}
							PhpMailBuildHeadersElems(&s, key, val)
						} else {
							core.PhpErrorDocref(nil, 1<<1, "Extra header element '%s' cannot be other than string or array.", key.val)
						}
						break
					}
				} else {
					if val.u1.v.type_ == 6 {
						PhpMailBuildHeadersElem(&s, key, val)
					} else if val.u1.v.type_ == 7 {
						PhpMailBuildHeadersElems(&s, key, val)
					} else {
						core.PhpErrorDocref(nil, 1<<1, "Extra header element '%s' cannot be other than string or array.", key.val)
					}
				}
				break
			case g.SizeOf("\"reply-to\"") - 1:
				if !(strncasecmp("reply-to", key.val, key.len_)) {
					for {
						if val.u1.v.type_ == 6 {
							PhpMailBuildHeadersElem(&s, key, val)
						} else if val.u1.v.type_ == 7 {
							if !(strncasecmp("reply-to", key.val, key.len_)) {
								core.PhpErrorDocref(nil, 1<<1, "'%s' header must be at most one header. Array is passed for '%s'", "reply-to", "reply-to")
								continue
							}
							PhpMailBuildHeadersElems(&s, key, val)
						} else {
							core.PhpErrorDocref(nil, 1<<1, "Extra header element '%s' cannot be other than string or array.", key.val)
						}
						break
					}
				} else {
					if val.u1.v.type_ == 6 {
						PhpMailBuildHeadersElem(&s, key, val)
					} else if val.u1.v.type_ == 7 {
						PhpMailBuildHeadersElems(&s, key, val)
					} else {
						core.PhpErrorDocref(nil, 1<<1, "Extra header element '%s' cannot be other than string or array.", key.val)
					}
				}
				break
			case g.SizeOf("\"to\"") - 1:
				if !(strncasecmp("to", key.val, key.len_)) {
					core.PhpErrorDocref(nil, 1<<1, "Extra header cannot contain 'To' header")
					continue
				}
				if !(strncasecmp("cc", key.val, key.len_)) {
					for {
						if val.u1.v.type_ == 6 {
							PhpMailBuildHeadersElem(&s, key, val)
						} else if val.u1.v.type_ == 7 {
							if !(strncasecmp("cc", key.val, key.len_)) {
								core.PhpErrorDocref(nil, 1<<1, "'%s' header must be at most one header. Array is passed for '%s'", "cc", "cc")
								continue
							}
							PhpMailBuildHeadersElems(&s, key, val)
						} else {
							core.PhpErrorDocref(nil, 1<<1, "Extra header element '%s' cannot be other than string or array.", key.val)
						}
						break
					}
				} else {
					if val.u1.v.type_ == 6 {
						PhpMailBuildHeadersElem(&s, key, val)
					} else if val.u1.v.type_ == 7 {
						PhpMailBuildHeadersElems(&s, key, val)
					} else {
						core.PhpErrorDocref(nil, 1<<1, "Extra header element '%s' cannot be other than string or array.", key.val)
					}
				}
				break
			case g.SizeOf("\"bcc\"") - 1:
				if !(strncasecmp("bcc", key.val, key.len_)) {
					for {
						if val.u1.v.type_ == 6 {
							PhpMailBuildHeadersElem(&s, key, val)
						} else if val.u1.v.type_ == 7 {
							if !(strncasecmp("bcc", key.val, key.len_)) {
								core.PhpErrorDocref(nil, 1<<1, "'%s' header must be at most one header. Array is passed for '%s'", "bcc", "bcc")
								continue
							}
							PhpMailBuildHeadersElems(&s, key, val)
						} else {
							core.PhpErrorDocref(nil, 1<<1, "Extra header element '%s' cannot be other than string or array.", key.val)
						}
						break
					}
				} else {
					if val.u1.v.type_ == 6 {
						PhpMailBuildHeadersElem(&s, key, val)
					} else if val.u1.v.type_ == 7 {
						PhpMailBuildHeadersElems(&s, key, val)
					} else {
						core.PhpErrorDocref(nil, 1<<1, "Extra header element '%s' cannot be other than string or array.", key.val)
					}
				}
				break
			case g.SizeOf("\"message-id\"") - 1:
				if !(strncasecmp("message-id", key.val, key.len_)) {
					for {
						if val.u1.v.type_ == 6 {
							PhpMailBuildHeadersElem(&s, key, val)
						} else if val.u1.v.type_ == 7 {
							if !(strncasecmp("message-id", key.val, key.len_)) {
								core.PhpErrorDocref(nil, 1<<1, "'%s' header must be at most one header. Array is passed for '%s'", "message-id", "message-id")
								continue
							}
							PhpMailBuildHeadersElems(&s, key, val)
						} else {
							core.PhpErrorDocref(nil, 1<<1, "Extra header element '%s' cannot be other than string or array.", key.val)
						}
						break
					}
				} else if !(strncasecmp("references", key.val, key.len_)) {
					for {
						if val.u1.v.type_ == 6 {
							PhpMailBuildHeadersElem(&s, key, val)
						} else if val.u1.v.type_ == 7 {
							if !(strncasecmp("references", key.val, key.len_)) {
								core.PhpErrorDocref(nil, 1<<1, "'%s' header must be at most one header. Array is passed for '%s'", "references", "references")
								continue
							}
							PhpMailBuildHeadersElems(&s, key, val)
						} else {
							core.PhpErrorDocref(nil, 1<<1, "Extra header element '%s' cannot be other than string or array.", key.val)
						}
						break
					}
				} else {
					if val.u1.v.type_ == 6 {
						PhpMailBuildHeadersElem(&s, key, val)
					} else if val.u1.v.type_ == 7 {
						PhpMailBuildHeadersElems(&s, key, val)
					} else {
						core.PhpErrorDocref(nil, 1<<1, "Extra header element '%s' cannot be other than string or array.", key.val)
					}
				}
				break
			case g.SizeOf("\"in-reply-to\"") - 1:
				if !(strncasecmp("in-reply-to", key.val, key.len_)) {
					for {
						if val.u1.v.type_ == 6 {
							PhpMailBuildHeadersElem(&s, key, val)
						} else if val.u1.v.type_ == 7 {
							if !(strncasecmp("in-reply-to", key.val, key.len_)) {
								core.PhpErrorDocref(nil, 1<<1, "'%s' header must be at most one header. Array is passed for '%s'", "in-reply-to", "in-reply-to")
								continue
							}
							PhpMailBuildHeadersElems(&s, key, val)
						} else {
							core.PhpErrorDocref(nil, 1<<1, "Extra header element '%s' cannot be other than string or array.", key.val)
						}
						break
					}
				} else {
					if val.u1.v.type_ == 6 {
						PhpMailBuildHeadersElem(&s, key, val)
					} else if val.u1.v.type_ == 7 {
						PhpMailBuildHeadersElems(&s, key, val)
					} else {
						core.PhpErrorDocref(nil, 1<<1, "Extra header element '%s' cannot be other than string or array.", key.val)
					}
				}
				break
			case g.SizeOf("\"subject\"") - 1:
				if !(strncasecmp("subject", key.val, key.len_)) {
					core.PhpErrorDocref(nil, 1<<1, "Extra header cannot contain 'Subject' header")
					continue
				}
				if val.u1.v.type_ == 6 {
					PhpMailBuildHeadersElem(&s, key, val)
				} else if val.u1.v.type_ == 7 {
					PhpMailBuildHeadersElems(&s, key, val)
				} else {
					core.PhpErrorDocref(nil, 1<<1, "Extra header element '%s' cannot be other than string or array.", key.val)
				}
				break
			default:
				if val.u1.v.type_ == 6 {
					PhpMailBuildHeadersElem(&s, key, val)
				} else if val.u1.v.type_ == 7 {
					PhpMailBuildHeadersElems(&s, key, val)
				} else {
					core.PhpErrorDocref(nil, 1<<1, "Extra header element '%s' cannot be other than string or array.", key.val)
				}
			}

			/* https://tools.ietf.org/html/rfc2822#section-3.6 */

		}
		break
	}

	/* Remove the last \r\n */

	if s.s != nil {
		s.s.len_ -= 2
	}
	zend.SmartStr0(&s)
	return s.s
}

/* {{{ proto int mail(string to, string subject, string message [, string additional_headers [, string additional_parameters]])
   Send an email message */

func ZifMail(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var to *byte = nil
	var message *byte = nil
	var subject *byte = nil
	var extra_cmd *zend.ZendString = nil
	var str_headers *zend.ZendString = nil
	var tmp_headers *zend.ZendString
	var headers *zend.Zval = nil
	var to_len int
	var message_len int
	var subject_len int
	var i int
	var force_extra_parameters *byte = zend.ZendIniStringEx("mail.force_extra_parameters", g.SizeOf("\"mail.force_extra_parameters\"")-1, 0, nil)
	var to_r *byte
	var subject_r *byte
	var p *byte
	var e *byte
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

			if zend.ZendParseArgString(_arg, &to, &to_len, 0) == 0 {
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

			if zend.ZendParseArgString(_arg, &subject, &subject_len, 0) == 0 {
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

			if zend.ZendParseArgString(_arg, &message, &message_len, 0) == 0 {
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

			zend.ZendParseArgZvalDeref(_arg, &headers, 0)
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

			if zend.ZendParseArgStr(_arg, &extra_cmd, 0) == 0 {
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

	/* ASCIIZ check */

	p = to
	e = p + to_len
	for g.Assign(&p, memchr(p, '0', e-p)) {
		*p = ' '
	}
	p = subject
	e = p + subject_len
	for g.Assign(&p, memchr(p, '0', e-p)) {
		*p = ' '
	}
	p = message
	e = p + message_len
	for g.Assign(&p, memchr(p, '0', e-p)) {
		*p = ' '
	}
	if headers != nil {
		switch headers.u1.v.type_ {
		case 6:
			tmp_headers = zend.ZendStringInit(headers.value.str.val, headers.value.str.len_, 0)
			p = tmp_headers.val
			e = p + tmp_headers.len_
			for g.Assign(&p, memchr(p, '0', e-p)) {
				*p = ' '
			}
			str_headers = PhpTrim(tmp_headers, nil, 0, 2)
			zend.ZendStringReleaseEx(tmp_headers, 0)
			break
		case 7:
			str_headers = PhpMailBuildHeaders(headers)
			break
		default:
			core.PhpErrorDocref(nil, 1<<1, "headers parameter must be string or array")
			return_value.u1.type_info = 2
			return
		}
	}
	if extra_cmd != nil {
		p = extra_cmd.val
		e = p + extra_cmd.len_
		for g.Assign(&p, memchr(p, '0', e-p)) {
			*p = ' '
		}
	}
	if to_len > 0 {
		to_r = zend._estrndup(to, to_len)
		for ; to_len != 0; to_len-- {
			if !(isspace(uint8(to_r[to_len-1]))) {
				break
			}
			to_r[to_len-1] = '0'
		}
		for i = 0; to_r[i]; i++ {
			if iscntrl(uint8(to_r[i])) {

				/* According to RFC 822, section 3.1.1 long headers may be separated into
				 * parts using CRLF followed at least one linear-white-space character ('\t' or ' ').
				 * To prevent these separators from being replaced with a space, we use the
				 * SKIP_LONG_HEADER_SEP to skip over them. */

				if to_r[i] == '\r' && to_r[i+1] == '\n' && (to_r[i+2] == ' ' || to_r[i+2] == '\t') {
					i += 2
					for to_r[i+1] == ' ' || to_r[i+1] == '\t' {
						i++
					}
					continue
				}
				to_r[i] = ' '
			}
		}
	} else {
		to_r = to
	}
	if subject_len > 0 {
		subject_r = zend._estrndup(subject, subject_len)
		for ; subject_len != 0; subject_len-- {
			if !(isspace(uint8(subject_r[subject_len-1]))) {
				break
			}
			subject_r[subject_len-1] = '0'
		}
		for i = 0; subject_r[i]; i++ {
			if iscntrl(uint8(subject_r[i])) {
				if subject_r[i] == '\r' && subject_r[i+1] == '\n' && (subject_r[i+2] == ' ' || subject_r[i+2] == '\t') {
					i += 2
					for subject_r[i+1] == ' ' || subject_r[i+1] == '\t' {
						i++
					}
					continue
				}
				subject_r[i] = ' '
			}
		}
	} else {
		subject_r = subject
	}
	if force_extra_parameters != nil {
		extra_cmd = PhpEscapeShellCmd(force_extra_parameters)
	} else if extra_cmd != nil {
		extra_cmd = PhpEscapeShellCmd(extra_cmd.val)
	}
	if PhpMail(to_r, subject_r, message, g.CondF1(str_headers != nil && str_headers.len_ != 0, func() []byte { return str_headers.val }, nil), g.CondF1(extra_cmd != nil, func() []byte { return extra_cmd.val }, nil)) != 0 {
		return_value.u1.type_info = 3
	} else {
		return_value.u1.type_info = 2
	}
	if str_headers != nil {
		zend.ZendStringReleaseEx(str_headers, 0)
	}
	if extra_cmd != nil {
		zend.ZendStringReleaseEx(extra_cmd, 0)
	}
	if to_r != to {
		zend._efree(to_r)
	}
	if subject_r != subject {
		zend._efree(subject_r)
	}
}

/* }}} */

func PhpMailLogCrlfToSpaces(message *byte) {
	/* Find all instances of carriage returns or line feeds and
	 * replace them with spaces. Thus, a log line is always one line
	 * long
	 */

	var p *byte = message
	for g.Assign(&p, strpbrk(p, "\r\n")) {
		*p = ' '
	}
}
func PhpMailLogToSyslog(message *byte) {
	/* Write 'message' to syslog. */

	core.PhpSyslog(LOG_NOTICE, "%s", message)

	/* Write 'message' to syslog. */
}
func PhpMailLogToFile(filename *byte, message *byte, message_size int) {
	/* Write 'message' to the given file. */

	var flags uint32 = 0 | 0x8 | 0x400
	var stream *core.PhpStream = streams._phpStreamOpenWrapperEx(filename, "a", flags, nil, nil)
	if stream != nil {
		streams._phpStreamWrite(stream, message, message_size)
		streams._phpStreamFree(stream, 1|2)
	}
}
func PhpMailDetectMultipleCrlf(hdr *byte) int {
	/* This function detects multiple/malformed multiple newlines. */

	if hdr == nil || !(strlen(hdr)) {
		return 0
	}

	/* Should not have any newlines at the beginning. */

	if (*hdr) < 33 || (*hdr) > 126 || (*hdr) == ':' {
		return 1
	}
	for *hdr {
		if (*hdr) == '\r' {
			if (*(hdr + 1)) == '0' || (*(hdr + 1)) == '\r' || (*(hdr + 1)) == '\n' && ((*(hdr + 2)) == '0' || (*(hdr + 2)) == '\n' || (*(hdr + 2)) == '\r') {

				/* Malformed or multiple newlines. */

				return 1

				/* Malformed or multiple newlines. */

			} else {
				hdr += 2
			}
		} else if (*hdr) == '\n' {
			if (*(hdr + 1)) == '0' || (*(hdr + 1)) == '\r' || (*(hdr + 1)) == '\n' {

				/* Malformed or multiple newlines. */

				return 1

				/* Malformed or multiple newlines. */

			} else {
				hdr += 2
			}
		} else {
			hdr++
		}
	}
	return 0
}

/* {{{ php_mail
 */

func PhpMail(to *byte, subject *byte, message *byte, headers *byte, extra_cmd *byte) int {
	var sendmail *r.FILE
	var ret int
	var sendmail_path *byte = zend.ZendIniStringEx("sendmail_path", g.SizeOf("\"sendmail_path\"")-1, 0, nil)
	var sendmail_cmd *byte = nil
	var mail_log *byte = zend.ZendIniStringEx("mail.log", g.SizeOf("\"mail.log\"")-1, 0, nil)
	var hdr *byte = headers

	// #define MAIL_RET(val) if ( hdr != headers ) { efree ( hdr ) ; } return val ;

	if mail_log != nil && (*mail_log) {
		var logline *byte
		zend.ZendSpprintf(&logline, 0, "mail() on [%s:%d]: To: %s -- Headers: %s -- Subject: %s", zend.ZendGetExecutedFilename(), zend.ZendGetExecutedLineno(), to, g.Cond(hdr != nil, hdr, ""), subject)
		if hdr != nil {
			PhpMailLogCrlfToSpaces(logline)
		}
		if !(strcmp(mail_log, "syslog")) {
			PhpMailLogToSyslog(logline)
		} else {

			/* Add date when logging to file */

			var tmp *byte
			var curtime int64
			var date_str *zend.ZendString
			var len_ int
			time(&curtime)
			date_str = php_format_date("d-M-Y H:i:s e", 13, curtime, 1)
			len_ = zend.ZendSpprintf(&tmp, 0, "[%s] %s%s", date_str.val, logline, "\n")
			PhpMailLogToFile(mail_log, tmp, len_)
			zend.ZendStringFree(date_str)
			zend._efree(tmp)
		}
		zend._efree(logline)
	}
	if core.CoreGlobals.mail_x_header != 0 {
		var tmp *byte = zend.ZendGetExecutedFilename()
		var f *zend.ZendString
		f = PhpBasename(tmp, strlen(tmp), nil, 0)
		if headers != nil && (*headers) {
			zend.ZendSpprintf(&hdr, 0, "X-PHP-Originating-Script: "+"%"+"lld"+":%s\n%s", PhpGetuid(), f.val, headers)
		} else {
			zend.ZendSpprintf(&hdr, 0, "X-PHP-Originating-Script: "+"%"+"lld"+":%s", PhpGetuid(), f.val)
		}
		zend.ZendStringReleaseEx(f, 0)
	}
	if hdr != nil && PhpMailDetectMultipleCrlf(hdr) != 0 {
		core.PhpErrorDocref(nil, 1<<1, "Multiple or malformed newlines found in additional_header")
		if hdr != headers {
			zend._efree(hdr)
		}
		return 0
	}
	if sendmail_path == nil {
		if hdr != headers {
			zend._efree(hdr)
		}
		return 0
	}
	if extra_cmd != nil {
		zend.ZendSpprintf(&sendmail_cmd, 0, "%s %s", sendmail_path, extra_cmd)
	} else {
		sendmail_cmd = sendmail_path
	}

	/* Since popen() doesn't indicate if the internal fork() doesn't work
	 * (e.g. the shell can't be executed) we explicitly set it to 0 to be
	 * sure we don't catch any older errno value. */

	errno = 0
	sendmail = popen(sendmail_cmd, "w")
	if extra_cmd != nil {
		zend._efree(sendmail_cmd)
	}
	if sendmail != nil {
		if EACCES == errno {
			core.PhpErrorDocref(nil, 1<<1, "Permission denied: unable to execute shell to run mail delivery binary '%s'", sendmail_path)
			pclose(sendmail)
			if hdr != headers {
				zend._efree(hdr)
			}
			return 0
		}
		r.Fprintf(sendmail, "To: %s\n", to)
		r.Fprintf(sendmail, "Subject: %s\n", subject)
		if hdr != nil {
			r.Fprintf(sendmail, "%s\n", hdr)
		}
		r.Fprintf(sendmail, "\n%s\n", message)
		ret = pclose(sendmail)
		if ret != 0 {
			if hdr != headers {
				zend._efree(hdr)
			}
			return 0
		} else {
			if hdr != headers {
				zend._efree(hdr)
			}
			return 1
		}
	} else {
		core.PhpErrorDocref(nil, 1<<1, "Could not execute mail delivery program '%s'", sendmail_path)
		if hdr != headers {
			zend._efree(hdr)
		}
		return 0
	}
	if hdr != headers {
		zend._efree(hdr)
	}
	return 1
}

/* }}} */

func ZmInfoMail(zend_module *zend.ZendModuleEntry) {
	var sendmail_path *byte = zend.ZendIniStringEx("sendmail_path", g.SizeOf("\"sendmail_path\"")-1, 0, nil)
	PhpInfoPrintTableRow(2, "Path to sendmail", sendmail_path)
}

/* }}} */
