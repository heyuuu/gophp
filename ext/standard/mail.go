// <<generate>>

package standard

import (
	b "sik/builtin"
	"sik/core"
	r "sik/runtime"
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

func SKIP_LONG_HEADER_SEP(str *byte, pos int) {
	if str[pos] == '\r' && str[pos+1] == '\n' && (str[pos+2] == ' ' || str[pos+2] == '\t') {
		pos += 2
		for str[pos+1] == ' ' || str[pos+1] == '\t' {
			pos++
		}
		continue
	}
}
func MAIL_ASCIIZ_CHECK(str __auto__, len_ int) {
	p = str
	e = p + len_
	for b.Assign(&p, memchr(p, '0', e-p)) {
		*p = ' '
	}
}

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
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgString(_arg, &str, &str_len, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
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
	for j = 0; j < str_len; j++ {
		h = h + (h << 5) ^ zend.ZendUlong(uint8(tolower(str[j])))
	}
	h = h % 53
	zend.RETVAL_LONG(zend.ZendLong(h))
	return
}

/* }}} */

func PhpMailBuildHeadersCheckFieldValue(val *zend.Zval) zend.ZendBool {
	var len_ int = 0
	var value *zend.ZendString = zend.Z_STR_P(val)

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
	switch zend.Z_TYPE_P(val) {
	case zend.IS_STRING:
		if PhpMailBuildHeadersCheckFieldName(key) != zend.SUCCESS {
			core.PhpErrorDocref(nil, zend.E_WARNING, "Header field name (%s) contains invalid chars", zend.ZSTR_VAL(key))
			return
		}
		if PhpMailBuildHeadersCheckFieldValue(val) != zend.SUCCESS {
			core.PhpErrorDocref(nil, zend.E_WARNING, "Header field value (%s => %s) contains invalid chars or format", zend.ZSTR_VAL(key), zend.Z_STRVAL_P(val))
			return
		}
		zend.SmartStrAppend(s, key)
		zend.SmartStrAppendl(s, ": ", 2)
		zend.SmartStrAppends(s, zend.Z_STRVAL_P(val))
		zend.SmartStrAppendl(s, "\r\n", 2)
		break
	case zend.IS_ARRAY:
		PhpMailBuildHeadersElems(s, key, val)
		break
	default:
		core.PhpErrorDocref(nil, zend.E_WARNING, "headers array elements must be string or array (%s)", zend.ZSTR_VAL(key))
	}
}
func PhpMailBuildHeadersElems(s *zend.SmartStr, key *zend.ZendString, val *zend.Zval) {
	var tmp_key *zend.ZendString
	var tmp_val *zend.Zval
	for {
		var __ht *zend.HashTable = zend.Z_ARRVAL_P(val)
		var _p *zend.Bucket = __ht.arData
		var _end *zend.Bucket = _p + __ht.nNumUsed
		for ; _p != _end; _p++ {
			var _z *zend.Zval = &_p.val

			if zend.UNEXPECTED(zend.Z_TYPE_P(_z) == zend.IS_UNDEF) {
				continue
			}
			tmp_key = _p.key
			tmp_val = _z
			if tmp_key != nil {
				core.PhpErrorDocref(nil, zend.E_WARNING, "Multiple header key must be numeric index (%s)", zend.ZSTR_VAL(tmp_key))
				continue
			}
			if zend.Z_TYPE_P(tmp_val) != zend.IS_STRING {
				core.PhpErrorDocref(nil, zend.E_WARNING, "Multiple header values must be string (%s)", zend.ZSTR_VAL(key))
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
	zend.ZEND_ASSERT(zend.Z_TYPE_P(headers) == zend.IS_ARRAY)
	for {
		var __ht *zend.HashTable = zend.Z_ARRVAL_P(headers)
		var _p *zend.Bucket = __ht.arData
		var _end *zend.Bucket = _p + __ht.nNumUsed
		for ; _p != _end; _p++ {
			var _z *zend.Zval = &_p.val

			if zend.UNEXPECTED(zend.Z_TYPE_P(_z) == zend.IS_UNDEF) {
				continue
			}
			idx = _p.h
			key = _p.key
			val = _z
			if key == nil {
				core.PhpErrorDocref(nil, zend.E_WARNING, "Found numeric header ("+zend.ZEND_LONG_FMT+")", idx)
				continue
			}

			/* https://tools.ietf.org/html/rfc2822#section-3.6 */

			switch zend.ZSTR_LEN(key) {
			case b.SizeOf("\"orig-date\"") - 1:
				if !(strncasecmp("orig-date", zend.ZSTR_VAL(key), zend.ZSTR_LEN(key))) {
					PHP_MAIL_BUILD_HEADER_CHECK("orig-date", s, key, val)
				} else {
					PHP_MAIL_BUILD_HEADER_DEFAULT(s, key, val)
				}
				break
			case b.SizeOf("\"from\"") - 1:
				if !(strncasecmp("from", zend.ZSTR_VAL(key), zend.ZSTR_LEN(key))) {
					PHP_MAIL_BUILD_HEADER_CHECK("from", s, key, val)
				} else {
					PHP_MAIL_BUILD_HEADER_DEFAULT(s, key, val)
				}
				break
			case b.SizeOf("\"sender\"") - 1:
				if !(strncasecmp("sender", zend.ZSTR_VAL(key), zend.ZSTR_LEN(key))) {
					PHP_MAIL_BUILD_HEADER_CHECK("sender", s, key, val)
				} else {
					PHP_MAIL_BUILD_HEADER_DEFAULT(s, key, val)
				}
				break
			case b.SizeOf("\"reply-to\"") - 1:
				if !(strncasecmp("reply-to", zend.ZSTR_VAL(key), zend.ZSTR_LEN(key))) {
					PHP_MAIL_BUILD_HEADER_CHECK("reply-to", s, key, val)
				} else {
					PHP_MAIL_BUILD_HEADER_DEFAULT(s, key, val)
				}
				break
			case b.SizeOf("\"to\"") - 1:
				if !(strncasecmp("to", zend.ZSTR_VAL(key), zend.ZSTR_LEN(key))) {
					core.PhpErrorDocref(nil, zend.E_WARNING, "Extra header cannot contain 'To' header")
					continue
				}
				if !(strncasecmp("cc", zend.ZSTR_VAL(key), zend.ZSTR_LEN(key))) {
					PHP_MAIL_BUILD_HEADER_CHECK("cc", s, key, val)
				} else {
					PHP_MAIL_BUILD_HEADER_DEFAULT(s, key, val)
				}
				break
			case b.SizeOf("\"bcc\"") - 1:
				if !(strncasecmp("bcc", zend.ZSTR_VAL(key), zend.ZSTR_LEN(key))) {
					PHP_MAIL_BUILD_HEADER_CHECK("bcc", s, key, val)
				} else {
					PHP_MAIL_BUILD_HEADER_DEFAULT(s, key, val)
				}
				break
			case b.SizeOf("\"message-id\"") - 1:
				if !(strncasecmp("message-id", zend.ZSTR_VAL(key), zend.ZSTR_LEN(key))) {
					PHP_MAIL_BUILD_HEADER_CHECK("message-id", s, key, val)
				} else if !(strncasecmp("references", zend.ZSTR_VAL(key), zend.ZSTR_LEN(key))) {
					PHP_MAIL_BUILD_HEADER_CHECK("references", s, key, val)
				} else {
					PHP_MAIL_BUILD_HEADER_DEFAULT(s, key, val)
				}
				break
			case b.SizeOf("\"in-reply-to\"") - 1:
				if !(strncasecmp("in-reply-to", zend.ZSTR_VAL(key), zend.ZSTR_LEN(key))) {
					PHP_MAIL_BUILD_HEADER_CHECK("in-reply-to", s, key, val)
				} else {
					PHP_MAIL_BUILD_HEADER_DEFAULT(s, key, val)
				}
				break
			case b.SizeOf("\"subject\"") - 1:
				if !(strncasecmp("subject", zend.ZSTR_VAL(key), zend.ZSTR_LEN(key))) {
					core.PhpErrorDocref(nil, zend.E_WARNING, "Extra header cannot contain 'Subject' header")
					continue
				}
				PHP_MAIL_BUILD_HEADER_DEFAULT(s, key, val)
				break
			default:
				PHP_MAIL_BUILD_HEADER_DEFAULT(s, key, val)
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
	var force_extra_parameters *byte = zend.INI_STR("mail.force_extra_parameters")
	var to_r *byte
	var subject_r *byte
	var p *byte
	var e *byte
	for {
		var _flags int = 0
		var _min_num_args int = 3
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
			if zend.UNEXPECTED(_num_args < _min_num_args) || zend.UNEXPECTED(_num_args > _max_num_args) && zend.EXPECTED(_max_num_args >= 0) {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgString(_arg, &to, &to_len, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgString(_arg, &subject, &subject_len, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgString(_arg, &message, &message_len, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			zend.ZendParseArgZvalDeref(_arg, &headers, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.UNEXPECTED(zend.ZendParseArgStr(_arg, &extra_cmd, 0) == 0) {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if zend.UNEXPECTED(_error_code != zend.ZPP_ERROR_OK) {
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

	/* ASCIIZ check */

	MAIL_ASCIIZ_CHECK(to, to_len)
	MAIL_ASCIIZ_CHECK(subject, subject_len)
	MAIL_ASCIIZ_CHECK(message, message_len)
	if headers != nil {
		switch zend.Z_TYPE_P(headers) {
		case zend.IS_STRING:
			tmp_headers = zend.ZendStringInit(zend.Z_STRVAL_P(headers), zend.Z_STRLEN_P(headers), 0)
			MAIL_ASCIIZ_CHECK(zend.ZSTR_VAL(tmp_headers), zend.ZSTR_LEN(tmp_headers))
			str_headers = PhpTrim(tmp_headers, nil, 0, 2)
			zend.ZendStringReleaseEx(tmp_headers, 0)
			break
		case zend.IS_ARRAY:
			str_headers = PhpMailBuildHeaders(headers)
			break
		default:
			core.PhpErrorDocref(nil, zend.E_WARNING, "headers parameter must be string or array")
			zend.RETVAL_FALSE
			return
		}
	}
	if extra_cmd != nil {
		MAIL_ASCIIZ_CHECK(zend.ZSTR_VAL(extra_cmd), zend.ZSTR_LEN(extra_cmd))
	}
	if to_len > 0 {
		to_r = zend.Estrndup(to, to_len)
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

				SKIP_LONG_HEADER_SEP(to_r, i)
				to_r[i] = ' '
			}
		}
	} else {
		to_r = to
	}
	if subject_len > 0 {
		subject_r = zend.Estrndup(subject, subject_len)
		for ; subject_len != 0; subject_len-- {
			if !(isspace(uint8(subject_r[subject_len-1]))) {
				break
			}
			subject_r[subject_len-1] = '0'
		}
		for i = 0; subject_r[i]; i++ {
			if iscntrl(uint8(subject_r[i])) {
				SKIP_LONG_HEADER_SEP(subject_r, i)
				subject_r[i] = ' '
			}
		}
	} else {
		subject_r = subject
	}
	if force_extra_parameters != nil {
		extra_cmd = PhpEscapeShellCmd(force_extra_parameters)
	} else if extra_cmd != nil {
		extra_cmd = PhpEscapeShellCmd(zend.ZSTR_VAL(extra_cmd))
	}
	if PhpMail(to_r, subject_r, message, b.CondF1(str_headers != nil && zend.ZSTR_LEN(str_headers) != 0, func() []byte { return zend.ZSTR_VAL(str_headers) }, nil), b.CondF1(extra_cmd != nil, func() []byte { return zend.ZSTR_VAL(extra_cmd) }, nil)) != 0 {
		zend.RETVAL_TRUE
	} else {
		zend.RETVAL_FALSE
	}
	if str_headers != nil {
		zend.ZendStringReleaseEx(str_headers, 0)
	}
	if extra_cmd != nil {
		zend.ZendStringReleaseEx(extra_cmd, 0)
	}
	if to_r != to {
		zend.Efree(to_r)
	}
	if subject_r != subject {
		zend.Efree(subject_r)
	}
}

/* }}} */

func PhpMailLogCrlfToSpaces(message *byte) {
	/* Find all instances of carriage returns or line feeds and
	 * replace them with spaces. Thus, a log line is always one line
	 * long
	 */

	var p *byte = message
	for b.Assign(&p, strpbrk(p, "\r\n")) {
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

	var flags uint32 = core.IGNORE_URL_WIN | core.REPORT_ERRORS | core.STREAM_DISABLE_OPEN_BASEDIR
	var stream *core.PhpStream = core.PhpStreamOpenWrapper(filename, "a", flags, nil)
	if stream != nil {
		core.PhpStreamWrite(stream, message, message_size)
		core.PhpStreamClose(stream)
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
	var sendmail_path *byte = zend.INI_STR("sendmail_path")
	var sendmail_cmd *byte = nil
	var mail_log *byte = zend.INI_STR("mail.log")
	var hdr *byte = headers

	// #define MAIL_RET(val) if ( hdr != headers ) { efree ( hdr ) ; } return val ;

	if mail_log != nil && (*mail_log) {
		var logline *byte
		core.Spprintf(&logline, 0, "mail() on [%s:%d]: To: %s -- Headers: %s -- Subject: %s", zend.ZendGetExecutedFilename(), zend.ZendGetExecutedLineno(), to, b.Cond(hdr != nil, hdr, ""), subject)
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
			len_ = core.Spprintf(&tmp, 0, "[%s] %s%s", date_str.val, logline, core.PHP_EOL)
			PhpMailLogToFile(mail_log, tmp, len_)
			zend.ZendStringFree(date_str)
			zend.Efree(tmp)
		}
		zend.Efree(logline)
	}
	if core.PG(mail_x_header) {
		var tmp *byte = zend.ZendGetExecutedFilename()
		var f *zend.ZendString
		f = PhpBasename(tmp, strlen(tmp), nil, 0)
		if headers != nil && (*headers) {
			core.Spprintf(&hdr, 0, "X-PHP-Originating-Script: "+zend.ZEND_LONG_FMT+":%s\n%s", PhpGetuid(), zend.ZSTR_VAL(f), headers)
		} else {
			core.Spprintf(&hdr, 0, "X-PHP-Originating-Script: "+zend.ZEND_LONG_FMT+":%s", PhpGetuid(), zend.ZSTR_VAL(f))
		}
		zend.ZendStringReleaseEx(f, 0)
	}
	if hdr != nil && PhpMailDetectMultipleCrlf(hdr) != 0 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Multiple or malformed newlines found in additional_header")
		if hdr != headers {
			zend.Efree(hdr)
		}
		return 0
	}
	if sendmail_path == nil {
		if hdr != headers {
			zend.Efree(hdr)
		}
		return 0
	}
	if extra_cmd != nil {
		core.Spprintf(&sendmail_cmd, 0, "%s %s", sendmail_path, extra_cmd)
	} else {
		sendmail_cmd = sendmail_path
	}

	/* Since popen() doesn't indicate if the internal fork() doesn't work
	 * (e.g. the shell can't be executed) we explicitly set it to 0 to be
	 * sure we don't catch any older errno value. */

	errno = 0
	sendmail = popen(sendmail_cmd, "w")
	if extra_cmd != nil {
		zend.Efree(sendmail_cmd)
	}
	if sendmail != nil {
		if EACCES == errno {
			core.PhpErrorDocref(nil, zend.E_WARNING, "Permission denied: unable to execute shell to run mail delivery binary '%s'", sendmail_path)
			pclose(sendmail)
			if hdr != headers {
				zend.Efree(hdr)
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
				zend.Efree(hdr)
			}
			return 0
		} else {
			if hdr != headers {
				zend.Efree(hdr)
			}
			return 1
		}
	} else {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Could not execute mail delivery program '%s'", sendmail_path)
		if hdr != headers {
			zend.Efree(hdr)
		}
		return 0
	}
	if hdr != headers {
		zend.Efree(hdr)
	}
	return 1
}

/* }}} */

func ZmInfoMail(ZEND_MODULE_INFO_FUNC_ARGS) {
	var sendmail_path *byte = zend.INI_STR("sendmail_path")
	PhpInfoPrintTableRow(2, "Path to sendmail", sendmail_path)
}

/* }}} */
