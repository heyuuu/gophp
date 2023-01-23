// <<generate>>

package standard

import (
	"sik/core"
	"sik/core/streams"
	r "sik/runtime"
	g "sik/runtime/grammar"
	"sik/zend"
)

// Source: <ext/standard/file.h>

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

// #define FILE_H

// # include "php_network.h"

var ZifFdSet func(execute_data *zend.ZendExecuteData, return_value *zend.Zval)
var ZifFdIsset func(execute_data *zend.ZendExecuteData, return_value *zend.Zval)
var ZmStartupUserStreams func(type_ int, module_number int) int
var PhpSetSockBlocking func(socketd core.PhpSocketT, block int) int

// #define PHP_CSV_NO_ESCAPE       EOF

// #define META_DEF_BUFSIZE       8192

// #define PHP_FILE_USE_INCLUDE_PATH       1

// #define PHP_FILE_IGNORE_NEW_LINES       2

// #define PHP_FILE_SKIP_EMPTY_LINES       4

// #define PHP_FILE_APPEND       8

// #define PHP_FILE_NO_DEFAULT_CONTEXT       16

type PhpMetaTagsToken = int

const (
	TOK_EOF = 0
	TOK_OPENTAG
	TOK_CLOSETAG
	TOK_SLASH
	TOK_EQUAL
	TOK_SPACE
	TOK_ID
	TOK_STRING
	TOK_OTHER
)

// @type PhpMetaTagsData struct

// @type PhpFileGlobals struct

// #define FG(v) ( file_globals . v )

var FileGlobals PhpFileGlobals

// Source: <ext/standard/file.c>

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
   |          Stig Bakken <ssb@php.net>                                   |
   |          Andi Gutmans <andi@php.net>                                 |
   |          Zeev Suraski <zeev@php.net>                                 |
   | PHP 4.0 patches by Thies C. Arntzen (thies@thieso.net)               |
   | PHP streams by Wez Furlong (wez@thebrainroom.com)                    |
   +----------------------------------------------------------------------+
*/

// # include "php.h"

// # include "php_globals.h"

// # include "ext/standard/flock_compat.h"

// # include "ext/standard/exec.h"

// # include "ext/standard/php_filestat.h"

// # include "php_open_temporary_file.h"

// # include "ext/standard/basic_functions.h"

// # include "php_ini.h"

// # include "zend_smart_str.h"

// # include < stdio . h >

// # include < stdlib . h >

// # include < errno . h >

// # include < sys / types . h >

// # include < sys / stat . h >

// # include < fcntl . h >

// # include < sys / param . h >

// # include < sys / select . h >

// # include < sys / socket . h >

// # include < netinet / in . h >

// # include < netdb . h >

// # include < arpa / inet . h >

// # include "ext/standard/head.h"

// # include "php_string.h"

// # include "file.h"

// # include < pwd . h >

// # include < sys / time . h >

// # include "fsock.h"

// # include "fopen_wrappers.h"

// # include "streamsfuncs.h"

// # include "php_globals.h"

// # include < sys / file . h >

// # include < sys / mman . h >

// # include "scanf.h"

// # include "zend_API.h"

// # include < fnmatch . h >

// # include < wchar . h >

/* }}} */

// #define PHP_STREAM_TO_ZVAL(stream,arg) ZEND_ASSERT ( Z_TYPE_P ( arg ) == IS_RESOURCE ) ; php_stream_from_res ( stream , Z_RES_P ( arg ) ) ;

/* {{{ ZTS-stuff / Globals / Prototypes */

var LeStreamContext int = zend.FAILURE

func PhpLeStreamContext() int { return LeStreamContext }

/* }}} */

func FileContextDtor(res *zend.ZendResource) {
	var context *core.PhpStreamContext = (*core.PhpStreamContext)(res.ptr)
	if context.options.u1.v.type_ != 0 {
		zend.ZvalPtrDtor(&context.options)
		&context.options.u1.type_info = 0
	}
	streams.PhpStreamContextFree(context)
}
func FileGlobalsCtor(file_globals_p *PhpFileGlobals) {
	memset(file_globals_p, 0, g.SizeOf("php_file_globals"))
	file_globals_p.SetDefChunkSize(8192)
}
func FileGlobalsDtor(file_globals_p *PhpFileGlobals) {}
func ZmStartupFile(type_ int, module_number int) int {
	LeStreamContext = zend.ZendRegisterListDestructorsEx(FileContextDtor, nil, "stream-context", module_number)
	FileGlobalsCtor(&FileGlobals)
	zend.ZendRegisterIniEntries(IniEntries, module_number)
	zend.ZendRegisterLongConstant("SEEK_SET", g.SizeOf("\"SEEK_SET\"")-1, 0, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("SEEK_CUR", g.SizeOf("\"SEEK_CUR\"")-1, 1, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("SEEK_END", g.SizeOf("\"SEEK_END\"")-1, 2, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("LOCK_SH", g.SizeOf("\"LOCK_SH\"")-1, 1, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("LOCK_EX", g.SizeOf("\"LOCK_EX\"")-1, 2, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("LOCK_UN", g.SizeOf("\"LOCK_UN\"")-1, 3, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("LOCK_NB", g.SizeOf("\"LOCK_NB\"")-1, 4, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("STREAM_NOTIFY_CONNECT", g.SizeOf("\"STREAM_NOTIFY_CONNECT\"")-1, 2, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("STREAM_NOTIFY_AUTH_REQUIRED", g.SizeOf("\"STREAM_NOTIFY_AUTH_REQUIRED\"")-1, 3, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("STREAM_NOTIFY_AUTH_RESULT", g.SizeOf("\"STREAM_NOTIFY_AUTH_RESULT\"")-1, 10, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("STREAM_NOTIFY_MIME_TYPE_IS", g.SizeOf("\"STREAM_NOTIFY_MIME_TYPE_IS\"")-1, 4, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("STREAM_NOTIFY_FILE_SIZE_IS", g.SizeOf("\"STREAM_NOTIFY_FILE_SIZE_IS\"")-1, 5, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("STREAM_NOTIFY_REDIRECTED", g.SizeOf("\"STREAM_NOTIFY_REDIRECTED\"")-1, 6, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("STREAM_NOTIFY_PROGRESS", g.SizeOf("\"STREAM_NOTIFY_PROGRESS\"")-1, 7, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("STREAM_NOTIFY_FAILURE", g.SizeOf("\"STREAM_NOTIFY_FAILURE\"")-1, 9, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("STREAM_NOTIFY_COMPLETED", g.SizeOf("\"STREAM_NOTIFY_COMPLETED\"")-1, 8, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("STREAM_NOTIFY_RESOLVE", g.SizeOf("\"STREAM_NOTIFY_RESOLVE\"")-1, 1, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("STREAM_NOTIFY_SEVERITY_INFO", g.SizeOf("\"STREAM_NOTIFY_SEVERITY_INFO\"")-1, 0, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("STREAM_NOTIFY_SEVERITY_WARN", g.SizeOf("\"STREAM_NOTIFY_SEVERITY_WARN\"")-1, 1, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("STREAM_NOTIFY_SEVERITY_ERR", g.SizeOf("\"STREAM_NOTIFY_SEVERITY_ERR\"")-1, 2, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("STREAM_FILTER_READ", g.SizeOf("\"STREAM_FILTER_READ\"")-1, 0x1, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("STREAM_FILTER_WRITE", g.SizeOf("\"STREAM_FILTER_WRITE\"")-1, 0x2, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("STREAM_FILTER_ALL", g.SizeOf("\"STREAM_FILTER_ALL\"")-1, 0x1|0x2, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("STREAM_CLIENT_PERSISTENT", g.SizeOf("\"STREAM_CLIENT_PERSISTENT\"")-1, 1, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("STREAM_CLIENT_ASYNC_CONNECT", g.SizeOf("\"STREAM_CLIENT_ASYNC_CONNECT\"")-1, 2, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("STREAM_CLIENT_CONNECT", g.SizeOf("\"STREAM_CLIENT_CONNECT\"")-1, 4, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("STREAM_CRYPTO_METHOD_ANY_CLIENT", g.SizeOf("\"STREAM_CRYPTO_METHOD_ANY_CLIENT\"")-1, streams.STREAM_CRYPTO_METHOD_ANY_CLIENT, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("STREAM_CRYPTO_METHOD_SSLv2_CLIENT", g.SizeOf("\"STREAM_CRYPTO_METHOD_SSLv2_CLIENT\"")-1, streams.STREAM_CRYPTO_METHOD_SSLv2_CLIENT, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("STREAM_CRYPTO_METHOD_SSLv3_CLIENT", g.SizeOf("\"STREAM_CRYPTO_METHOD_SSLv3_CLIENT\"")-1, streams.STREAM_CRYPTO_METHOD_SSLv3_CLIENT, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("STREAM_CRYPTO_METHOD_SSLv23_CLIENT", g.SizeOf("\"STREAM_CRYPTO_METHOD_SSLv23_CLIENT\"")-1, streams.STREAM_CRYPTO_METHOD_SSLv23_CLIENT, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("STREAM_CRYPTO_METHOD_TLS_CLIENT", g.SizeOf("\"STREAM_CRYPTO_METHOD_TLS_CLIENT\"")-1, streams.STREAM_CRYPTO_METHOD_TLS_CLIENT, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("STREAM_CRYPTO_METHOD_TLSv1_0_CLIENT", g.SizeOf("\"STREAM_CRYPTO_METHOD_TLSv1_0_CLIENT\"")-1, streams.STREAM_CRYPTO_METHOD_TLSv1_0_CLIENT, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("STREAM_CRYPTO_METHOD_TLSv1_1_CLIENT", g.SizeOf("\"STREAM_CRYPTO_METHOD_TLSv1_1_CLIENT\"")-1, streams.STREAM_CRYPTO_METHOD_TLSv1_1_CLIENT, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("STREAM_CRYPTO_METHOD_TLSv1_2_CLIENT", g.SizeOf("\"STREAM_CRYPTO_METHOD_TLSv1_2_CLIENT\"")-1, streams.STREAM_CRYPTO_METHOD_TLSv1_2_CLIENT, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("STREAM_CRYPTO_METHOD_TLSv1_3_CLIENT", g.SizeOf("\"STREAM_CRYPTO_METHOD_TLSv1_3_CLIENT\"")-1, streams.STREAM_CRYPTO_METHOD_TLSv1_3_CLIENT, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("STREAM_CRYPTO_METHOD_ANY_SERVER", g.SizeOf("\"STREAM_CRYPTO_METHOD_ANY_SERVER\"")-1, streams.STREAM_CRYPTO_METHOD_ANY_SERVER, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("STREAM_CRYPTO_METHOD_SSLv2_SERVER", g.SizeOf("\"STREAM_CRYPTO_METHOD_SSLv2_SERVER\"")-1, streams.STREAM_CRYPTO_METHOD_SSLv2_SERVER, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("STREAM_CRYPTO_METHOD_SSLv3_SERVER", g.SizeOf("\"STREAM_CRYPTO_METHOD_SSLv3_SERVER\"")-1, streams.STREAM_CRYPTO_METHOD_SSLv3_SERVER, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("STREAM_CRYPTO_METHOD_SSLv23_SERVER", g.SizeOf("\"STREAM_CRYPTO_METHOD_SSLv23_SERVER\"")-1, streams.STREAM_CRYPTO_METHOD_SSLv23_SERVER, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("STREAM_CRYPTO_METHOD_TLS_SERVER", g.SizeOf("\"STREAM_CRYPTO_METHOD_TLS_SERVER\"")-1, streams.STREAM_CRYPTO_METHOD_TLS_SERVER, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("STREAM_CRYPTO_METHOD_TLSv1_0_SERVER", g.SizeOf("\"STREAM_CRYPTO_METHOD_TLSv1_0_SERVER\"")-1, streams.STREAM_CRYPTO_METHOD_TLSv1_0_SERVER, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("STREAM_CRYPTO_METHOD_TLSv1_1_SERVER", g.SizeOf("\"STREAM_CRYPTO_METHOD_TLSv1_1_SERVER\"")-1, streams.STREAM_CRYPTO_METHOD_TLSv1_1_SERVER, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("STREAM_CRYPTO_METHOD_TLSv1_2_SERVER", g.SizeOf("\"STREAM_CRYPTO_METHOD_TLSv1_2_SERVER\"")-1, streams.STREAM_CRYPTO_METHOD_TLSv1_2_SERVER, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("STREAM_CRYPTO_METHOD_TLSv1_3_SERVER", g.SizeOf("\"STREAM_CRYPTO_METHOD_TLSv1_3_SERVER\"")-1, streams.STREAM_CRYPTO_METHOD_TLSv1_3_SERVER, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("STREAM_CRYPTO_PROTO_SSLv3", g.SizeOf("\"STREAM_CRYPTO_PROTO_SSLv3\"")-1, streams.STREAM_CRYPTO_METHOD_SSLv3_SERVER, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("STREAM_CRYPTO_PROTO_TLSv1_0", g.SizeOf("\"STREAM_CRYPTO_PROTO_TLSv1_0\"")-1, streams.STREAM_CRYPTO_METHOD_TLSv1_0_SERVER, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("STREAM_CRYPTO_PROTO_TLSv1_1", g.SizeOf("\"STREAM_CRYPTO_PROTO_TLSv1_1\"")-1, streams.STREAM_CRYPTO_METHOD_TLSv1_1_SERVER, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("STREAM_CRYPTO_PROTO_TLSv1_2", g.SizeOf("\"STREAM_CRYPTO_PROTO_TLSv1_2\"")-1, streams.STREAM_CRYPTO_METHOD_TLSv1_2_SERVER, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("STREAM_CRYPTO_PROTO_TLSv1_3", g.SizeOf("\"STREAM_CRYPTO_PROTO_TLSv1_3\"")-1, streams.STREAM_CRYPTO_METHOD_TLSv1_3_SERVER, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("STREAM_SHUT_RD", g.SizeOf("\"STREAM_SHUT_RD\"")-1, streams.STREAM_SHUT_RD, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("STREAM_SHUT_WR", g.SizeOf("\"STREAM_SHUT_WR\"")-1, streams.STREAM_SHUT_WR, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("STREAM_SHUT_RDWR", g.SizeOf("\"STREAM_SHUT_RDWR\"")-1, streams.STREAM_SHUT_RDWR, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("STREAM_SOCK_STREAM", g.SizeOf("\"STREAM_SOCK_STREAM\"")-1, SOCK_STREAM, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("STREAM_SOCK_DGRAM", g.SizeOf("\"STREAM_SOCK_DGRAM\"")-1, SOCK_DGRAM, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("STREAM_PEEK", g.SizeOf("\"STREAM_PEEK\"")-1, streams.STREAM_PEEK, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("STREAM_OOB", g.SizeOf("\"STREAM_OOB\"")-1, streams.STREAM_OOB, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("STREAM_SERVER_BIND", g.SizeOf("\"STREAM_SERVER_BIND\"")-1, 4, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("STREAM_SERVER_LISTEN", g.SizeOf("\"STREAM_SERVER_LISTEN\"")-1, 8, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("FILE_USE_INCLUDE_PATH", g.SizeOf("\"FILE_USE_INCLUDE_PATH\"")-1, 1, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("FILE_IGNORE_NEW_LINES", g.SizeOf("\"FILE_IGNORE_NEW_LINES\"")-1, 2, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("FILE_SKIP_EMPTY_LINES", g.SizeOf("\"FILE_SKIP_EMPTY_LINES\"")-1, 4, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("FILE_APPEND", g.SizeOf("\"FILE_APPEND\"")-1, 8, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("FILE_NO_DEFAULT_CONTEXT", g.SizeOf("\"FILE_NO_DEFAULT_CONTEXT\"")-1, 16, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("FILE_TEXT", g.SizeOf("\"FILE_TEXT\"")-1, 0, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("FILE_BINARY", g.SizeOf("\"FILE_BINARY\"")-1, 0, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("FNM_NOESCAPE", g.SizeOf("\"FNM_NOESCAPE\"")-1, FNM_NOESCAPE, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("FNM_PATHNAME", g.SizeOf("\"FNM_PATHNAME\"")-1, FNM_PATHNAME, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("FNM_PERIOD", g.SizeOf("\"FNM_PERIOD\"")-1, FNM_PERIOD, 1<<0|1<<1, module_number)
	return zend.SUCCESS
}

/* }}} */

func ZmShutdownFile(type_ int, module_number int) int {
	FileGlobalsDtor(&FileGlobals)
	return zend.SUCCESS
}

/* }}} */

var FlockValues []int = []int{LOCK_SH, LOCK_EX, LOCK_UN}

/* {{{ proto bool flock(resource fp, int operation [, int &wouldblock])
   Portable file locking */

func ZifFlock(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var res *zend.Zval
	var wouldblock *zend.Zval = nil
	var act int
	var stream *core.PhpStream
	var operation zend.ZendLong = 0
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

			if zend.ZendParseArgResource(_arg, &res, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_RESOURCE
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

			if zend.ZendParseArgLong(_arg, &operation, &_dummy, 0, 0) == 0 {
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

			zend.ZendParseArgZvalDeref(_arg, &wouldblock, 0)
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
	r.Assert(res.u1.v.type_ == 9)
	if g.Assign(&stream, (*core.PhpStream)(zend.ZendFetchResource2(res.value.res, "stream", streams.PhpFileLeStream(), streams.PhpFileLePstream()))) == nil {
		return_value.u1.type_info = 2
		return
	}
	act = operation & 3
	if act < 1 || act > 3 {
		core.PhpErrorDocref(nil, 1<<1, "Illegal operation argument")
		return_value.u1.type_info = 2
		return
	}
	if wouldblock != nil {
		for {
			r.Assert(wouldblock.u1.v.type_ == 10)
			for {
				var _zv *zend.Zval = wouldblock
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

	/* flock_values contains all possible actions if (operation & 4) we won't block on the lock */

	act = FlockValues[act-1] | g.Cond((operation&4) != 0, LOCK_NB, 0)
	if streams._phpStreamSetOption(stream, 6, act, any(nil)) != 0 {
		if operation != 0 && errno == EAGAIN && wouldblock != nil {
			for {
				r.Assert(wouldblock.u1.v.type_ == 10)
				for {
					var _zv *zend.Zval = wouldblock
					var ref *zend.ZendReference = _zv.value.ref
					if ref.sources.ptr != nil {
						zend.ZendTryAssignTypedRefLong(ref, 1)
						break
					}
					_zv = &ref.val
					zend.ZvalPtrDtor(_zv)
					var __z *zend.Zval = _zv
					__z.value.lval = 1
					__z.u1.type_info = 4
					break
				}
				break
			}
		}
		return_value.u1.type_info = 2
		return
	}
	return_value.u1.type_info = 3
	return
}

/* }}} */

// #define PHP_META_UNSAFE       ".\\+*?[^]$() "

/* {{{ proto array get_meta_tags(string filename [, bool use_include_path])
   Extracts all meta tag content attributes from a file and returns an array */

func ZifGetMetaTags(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var filename *byte
	var filename_len int
	var use_include_path zend.ZendBool = 0
	var in_tag int = 0
	var done int = 0
	var looking_for_val int = 0
	var have_name int = 0
	var have_content int = 0
	var saw_name int = 0
	var saw_content int = 0
	var name *byte = nil
	var value *byte = nil
	var temp *byte = nil
	var tok PhpMetaTagsToken
	var tok_last PhpMetaTagsToken
	var md PhpMetaTagsData

	/* Initiailize our structure */

	memset(&md, 0, g.SizeOf("md"))

	/* Parse arguments */

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

			if zend.ZendParseArgPath(_arg, &filename, &filename_len, 0) == 0 {
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

			if zend.ZendParseArgBool(_arg, &use_include_path, &_dummy, 0) == 0 {
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
	md.SetStream(streams._phpStreamOpenWrapperEx(filename, "rb", g.Cond(use_include_path != 0, 0x1, 0)|0x8, nil, nil))
	if md.GetStream() == nil {
		return_value.u1.type_info = 2
		return
	}
	var __arr *zend.ZendArray = zend._zendNewArray(0)
	var __z *zend.Zval = return_value
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	tok_last = TOK_EOF
	for done == 0 && g.Assign(&tok, PhpNextMetaToken(&md)) != TOK_EOF {
		if tok == TOK_ID {
			if tok_last == TOK_OPENTAG {
				md.SetInMeta(!(strcasecmp("meta", md.GetTokenData())))
			} else if tok_last == TOK_SLASH && in_tag != 0 {
				if strcasecmp("head", md.GetTokenData()) == 0 {

					/* We are done here! */

					done = 1

					/* We are done here! */

				}
			} else if tok_last == TOK_EQUAL && looking_for_val != 0 {
				if saw_name != 0 {
					if name != nil {
						zend._efree(name)
					}

					/* Get the NAME attr (Single word attr, non-quoted) */

					name = zend._estrndup(md.GetTokenData(), md.GetTokenLen())
					temp = name
					for temp != nil && (*temp) {
						if strchr(".\\+*?[^]$() ", *temp) {
							*temp = '_'
						}
						temp++
					}
					have_name = 1
				} else if saw_content != 0 {
					if value != nil {
						zend._efree(value)
					}
					value = zend._estrndup(md.GetTokenData(), md.GetTokenLen())
					have_content = 1
				}
				looking_for_val = 0
			} else {
				if md.GetInMeta() != 0 {
					if strcasecmp("name", md.GetTokenData()) == 0 {
						saw_name = 1
						saw_content = 0
						looking_for_val = 1
					} else if strcasecmp("content", md.GetTokenData()) == 0 {
						saw_name = 0
						saw_content = 1
						looking_for_val = 1
					}
				}
			}
		} else if tok == TOK_STRING && tok_last == TOK_EQUAL && looking_for_val != 0 {
			if saw_name != 0 {
				if name != nil {
					zend._efree(name)
				}

				/* Get the NAME attr (Quoted single/double) */

				name = zend._estrndup(md.GetTokenData(), md.GetTokenLen())
				temp = name
				for temp != nil && (*temp) {
					if strchr(".\\+*?[^]$() ", *temp) {
						*temp = '_'
					}
					temp++
				}
				have_name = 1
			} else if saw_content != 0 {
				if value != nil {
					zend._efree(value)
				}
				value = zend._estrndup(md.GetTokenData(), md.GetTokenLen())
				have_content = 1
			}
			looking_for_val = 0
		} else if tok == TOK_OPENTAG {
			if looking_for_val != 0 {
				looking_for_val = 0
				saw_name = 0
				have_name = saw_name
				saw_content = 0
				have_content = saw_content
			}
			in_tag = 1
		} else if tok == TOK_CLOSETAG {
			if have_name != 0 {

				/* For BC */

				PhpStrtolower(name, strlen(name))
				if have_content != 0 {
					zend.AddAssocStringEx(return_value, name, strlen(name), value)
				} else {
					zend.AddAssocStringEx(return_value, name, strlen(name), "")
				}
				zend._efree(name)
				if value != nil {
					zend._efree(value)
				}
			} else if have_content != 0 {
				zend._efree(value)
			}
			value = nil
			name = value

			/* Reset all of our flags */

			looking_for_val = 0
			in_tag = looking_for_val
			saw_name = 0
			have_name = saw_name
			saw_content = 0
			have_content = saw_content
			md.SetInMeta(0)
		}
		tok_last = tok
		if md.GetTokenData() != nil {
			zend._efree(md.GetTokenData())
		}
		md.SetTokenData(nil)
	}
	if value != nil {
		zend._efree(value)
	}
	if name != nil {
		zend._efree(name)
	}
	streams._phpStreamFree(md.GetStream(), 1|2)
}

/* }}} */

func ZifFileGetContents(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var filename *byte
	var filename_len int
	var use_include_path zend.ZendBool = 0
	var stream *core.PhpStream
	var offset zend.ZendLong = 0
	var maxlen zend.ZendLong = ssize_t(size_t - 1)
	var zcontext *zend.Zval = nil
	var context *core.PhpStreamContext = nil
	var contents *zend.ZendString

	/* Parse arguments */

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

			if zend.ZendParseArgPath(_arg, &filename, &filename_len, 0) == 0 {
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

			if zend.ZendParseArgBool(_arg, &use_include_path, &_dummy, 0) == 0 {
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

			if zend.ZendParseArgResource(_arg, &zcontext, 1) == 0 {
				_expected_type = zend.Z_EXPECTED_RESOURCE
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

			if zend.ZendParseArgLong(_arg, &maxlen, &_dummy, 0, 0) == 0 {
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
	if execute_data.This.u2.num_args == 5 && maxlen < 0 {
		core.PhpErrorDocref(nil, 1<<1, "length must be greater than or equal to zero")
		return_value.u1.type_info = 2
		return
	}
	if g.CondF2(g.CondF1(zcontext != nil, func() any { return zend.ZendFetchResourceEx(zcontext, "Stream-Context", PhpLeStreamContext()) }, 0), nil, func() *core.PhpStreamContext { return FileGlobals.GetDefaultContext() }) {
		context = FileGlobals.GetDefaultContext()
	} else {
		FileGlobals.SetDefaultContext(streams.PhpStreamContextAlloc())
		context = FileGlobals.GetDefaultContext()
	}
	stream = streams._phpStreamOpenWrapperEx(filename, "rb", g.Cond(use_include_path != 0, 0x1, 0)|0x8, nil, context)
	if stream == nil {
		return_value.u1.type_info = 2
		return
	}
	if offset != 0 && streams._phpStreamSeek(stream, offset, g.Cond(offset > 0, 0, 2)) < 0 {
		core.PhpErrorDocref(nil, 1<<1, "Failed to seek to position "+"%"+"lld"+" in the stream", offset)
		streams._phpStreamFree(stream, 1|2)
		return_value.u1.type_info = 2
		return
	}
	if g.Assign(&contents, streams._phpStreamCopyToMem(stream, maxlen, 0)) != nil {
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = contents
		__z.value.str = __s
		if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
			__z.u1.type_info = 6
		} else {
			__z.u1.type_info = 6 | 1<<0<<8
		}
	} else {
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = zend.ZendEmptyString
		__z.value.str = __s
		__z.u1.type_info = 6
	}
	streams._phpStreamFree(stream, 1|2)
}

/* }}} */

func ZifFilePutContents(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var stream *core.PhpStream
	var filename *byte
	var filename_len int
	var data *zend.Zval
	var numbytes ssize_t = 0
	var flags zend.ZendLong = 0
	var zcontext *zend.Zval = nil
	var context *core.PhpStreamContext = nil
	var srcstream *core.PhpStream = nil
	var mode []byte = "wb"
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

			if zend.ZendParseArgPath(_arg, &filename, &filename_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_PATH
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

			zend.ZendParseArgZvalDeref(_arg, &data, 0)
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

			if zend.ZendParseArgLong(_arg, &flags, &_dummy, 0, 0) == 0 {
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
	if data.u1.v.type_ == 9 {
		if g.Assign(&srcstream, (*core.PhpStream)(zend.ZendFetchResource2Ex(data, "stream", streams.PhpFileLeStream(), streams.PhpFileLePstream()))) == nil {
			return_value.u1.type_info = 2
			return
		}
	}
	if g.CondF2(g.CondF1(zcontext != nil, func() any { return zend.ZendFetchResourceEx(zcontext, "Stream-Context", PhpLeStreamContext()) }, flags&16), nil, func() *core.PhpStreamContext { return FileGlobals.GetDefaultContext() }) {
		context = FileGlobals.GetDefaultContext()
	} else {
		FileGlobals.SetDefaultContext(streams.PhpStreamContextAlloc())
		context = FileGlobals.GetDefaultContext()
	}
	if (flags & 8) != 0 {
		mode[0] = 'a'
	} else if (flags & LOCK_EX) != 0 {

		/* check to make sure we are dealing with a regular file */

		if zend.ZendMemnstr(filename, "://", g.SizeOf("\"://\"")-1, filename+filename_len) != nil {
			if strncasecmp(filename, "file://", g.SizeOf("\"file://\"")-1) {
				core.PhpErrorDocref(nil, 1<<1, "Exclusive locks may only be set for regular files")
				return_value.u1.type_info = 2
				return
			}
		}
		mode[0] = 'c'
	}
	mode[2] = '0'
	stream = streams._phpStreamOpenWrapperEx(filename, mode, g.Cond((flags&1) != 0, 0x1, 0)|0x8, nil, context)
	if stream == nil {
		return_value.u1.type_info = 2
		return
	}
	if (flags&LOCK_EX) != 0 && (!(g.Cond(streams._phpStreamSetOption(stream, 6, 0, any(1)) == 0, 1, 0)) || streams._phpStreamSetOption(stream, 6, LOCK_EX, any(nil)) != 0) {
		streams._phpStreamFree(stream, 1|2)
		core.PhpErrorDocref(nil, 1<<1, "Exclusive locks are not supported for this stream")
		return_value.u1.type_info = 2
		return
	}
	if mode[0] == 'c' {
		streams._phpStreamTruncateSetSize(stream, 0)
	}
	switch data.u1.v.type_ {
	case 9:
		var len_ int
		if streams._phpStreamCopyToStreamEx(srcstream, stream, size_t-1, &len_) != zend.SUCCESS {
			numbytes = -1
		} else {
			if len_ > INT64_MAX {
				core.PhpErrorDocref(nil, 1<<1, "content truncated from %zu to "+"%"+"lld"+" bytes", len_, INT64_MAX)
				len_ = INT64_MAX
			}
			numbytes = len_
		}
		break
	case 1:

	case 4:

	case 5:

	case 2:

	case 3:
		if data.u1.v.type_ != 6 {
			if data.u1.v.type_ != 6 {
				zend._convertToString(data)
			}
		}
	case 6:
		if data.value.str.len_ != 0 {
			numbytes = streams._phpStreamWrite(stream, data.value.str.val, data.value.str.len_)
			if numbytes != data.value.str.len_ {
				core.PhpErrorDocref(nil, 1<<1, "Only %zd of %zd bytes written, possibly out of free disk space", numbytes, data.value.str.len_)
				numbytes = -1
			}
		}
		break
	case 7:
		if data.value.arr.nNumOfElements != 0 {
			var bytes_written ssize_t
			var tmp *zend.Zval
			for {
				var __ht *zend.HashTable = data.value.arr
				var _p *zend.Bucket = __ht.arData
				var _end *zend.Bucket = _p + __ht.nNumUsed
				for ; _p != _end; _p++ {
					var _z *zend.Zval = &_p.val

					if _z.u1.v.type_ == 0 {
						continue
					}
					tmp = _z
					var t *zend.ZendString
					var str *zend.ZendString = zend.ZvalGetTmpString(tmp, &t)
					if str.len_ != 0 {
						numbytes += str.len_
						bytes_written = streams._phpStreamWrite(stream, str.val, str.len_)
						if bytes_written != str.len_ {
							core.PhpErrorDocref(nil, 1<<1, "Failed to write %zd bytes to %s", str.len_, filename)
							zend.ZendTmpStringRelease(t)
							numbytes = -1
							break
						}
					}
					zend.ZendTmpStringRelease(t)
				}
				break
			}
		}
		break
	case 8:
		if data.value.obj.handlers != nil {
			var out zend.Zval
			if zend.ZendStdCastObjectTostring(data, &out, 6) == zend.SUCCESS {
				numbytes = streams._phpStreamWrite(stream, out.value.str.val, out.value.str.len_)
				if numbytes != out.value.str.len_ {
					core.PhpErrorDocref(nil, 1<<1, "Only %zd of %zd bytes written, possibly out of free disk space", numbytes, out.value.str.len_)
					numbytes = -1
				}
				zend.ZvalPtrDtorStr(&out)
				break
			}
		}
	default:
		numbytes = -1
		break
	}
	streams._phpStreamFree(stream, 1|2)
	if numbytes < 0 {
		return_value.u1.type_info = 2
		return
	}
	var __z *zend.Zval = return_value
	__z.value.lval = numbytes
	__z.u1.type_info = 4
	return
}

/* }}} */

// #define PHP_FILE_BUF_SIZE       80

/* {{{ proto array file(string filename [, int flags[, resource context]])
   Read entire file into an array */

func ZifFile(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var filename *byte
	var filename_len int
	var p *byte
	var s *byte
	var e *byte
	var i int = 0
	var eol_marker byte = '\n'
	var flags zend.ZendLong = 0
	var use_include_path zend.ZendBool
	var include_new_line zend.ZendBool
	var skip_blank_lines zend.ZendBool
	var stream *core.PhpStream
	var zcontext *zend.Zval = nil
	var context *core.PhpStreamContext = nil
	var target_buf *zend.ZendString

	/* Parse arguments */

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

			if zend.ZendParseArgPath(_arg, &filename, &filename_len, 0) == 0 {
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

			if zend.ZendParseArgLong(_arg, &flags, &_dummy, 0, 0) == 0 {
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
	if flags < 0 || flags > (1|2|4|16) {
		core.PhpErrorDocref(nil, 1<<1, "'"+"%"+"lld"+"' flag is not supported", flags)
		return_value.u1.type_info = 2
		return
	}
	use_include_path = flags & 1
	include_new_line = !(flags & 2)
	skip_blank_lines = flags & 4
	if g.CondF2(g.CondF1(zcontext != nil, func() any { return zend.ZendFetchResourceEx(zcontext, "Stream-Context", PhpLeStreamContext()) }, flags&16), nil, func() *core.PhpStreamContext { return FileGlobals.GetDefaultContext() }) {
		context = FileGlobals.GetDefaultContext()
	} else {
		FileGlobals.SetDefaultContext(streams.PhpStreamContextAlloc())
		context = FileGlobals.GetDefaultContext()
	}
	stream = streams._phpStreamOpenWrapperEx(filename, "rb", g.Cond(use_include_path != 0, 0x1, 0)|0x8, nil, context)
	if stream == nil {
		return_value.u1.type_info = 2
		return
	}

	/* Initialize return array */

	var __arr *zend.ZendArray = zend._zendNewArray(0)
	var __z *zend.Zval = return_value
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	if g.Assign(&target_buf, streams._phpStreamCopyToMem(stream, size_t-1, 0)) != nil {
		s = target_buf.val
		e = target_buf.val + target_buf.len_
		if !(g.Assign(&p, (*byte)(streams.PhpStreamLocateEol(stream, target_buf)))) {
			p = e
			goto parse_eol
		}
		if (stream.flags & 0x8) != 0 {
			eol_marker = '\r'
		}

		/* for performance reasons the code is duplicated, so that the if (include_new_line)
		 * will not need to be done for every single line in the file. */

		if include_new_line != 0 {
			for {
				p++
			parse_eol:
				zend.AddIndexStringl(return_value, g.PostInc(&i), s, p-s)
				s = p
				if !(g.Assign(&p, memchr(p, eol_marker, e-p))) {
					break
				}
			}
		} else {
			for {
				var windows_eol int = 0
				if p != target_buf.val && eol_marker == '\n' && (*(p - 1)) == '\r' {
					windows_eol++
				}
				if skip_blank_lines != 0 && p-s-windows_eol == 0 {
					p++
					s = p
					continue
				}
				zend.AddIndexStringl(return_value, g.PostInc(&i), s, p-s-windows_eol)
				p++
				s = p
				if !(g.Assign(&p, memchr(p, eol_marker, e-p))) {
					break
				}
			}
		}

		/* handle any left overs of files without new lines */

		if s != e {
			p = e
			goto parse_eol
		}

		/* handle any left overs of files without new lines */

	}
	if target_buf != nil {
		zend.ZendStringFree(target_buf)
	}
	streams._phpStreamFree(stream, 1|2)
}

/* }}} */

func ZifTempnam(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var dir *byte
	var prefix *byte
	var dir_len int
	var prefix_len int
	var opened_path *zend.ZendString
	var fd int
	var p *zend.ZendString
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

			if zend.ZendParseArgPath(_arg, &dir, &dir_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_PATH
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

			if zend.ZendParseArgPath(_arg, &prefix, &prefix_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_PATH
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
	p = PhpBasename(prefix, prefix_len, nil, 0)
	if p.len_ > 64 {
		p.val[63] = '0'
	}
	return_value.u1.type_info = 2
	if g.Assign(&fd, core.PhpOpenTemporaryFdEx(dir, p.val, &opened_path, 1<<0|1<<2)) >= 0 {
		close(fd)
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = opened_path
		__z.value.str = __s
		if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
			__z.u1.type_info = 6
		} else {
			__z.u1.type_info = 6 | 1<<0<<8
		}
	}
	zend.ZendStringReleaseEx(p, 0)
}

/* }}} */

func PhpIfTmpfile(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var stream *core.PhpStream
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	stream = streams._phpStreamFopenTmpfile(0)
	if stream != nil {
		var __z *zend.Zval = return_value
		__z.value.res = stream.res
		__z.u1.type_info = 9 | 1<<0<<8
		stream.__exposed = 1
	} else {
		return_value.u1.type_info = 2
		return
	}
}

/* }}} */

func PhpIfFopen(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var filename *byte
	var mode *byte
	var filename_len int
	var mode_len int
	var use_include_path zend.ZendBool = 0
	var zcontext *zend.Zval = nil
	var stream *core.PhpStream
	var context *core.PhpStreamContext = nil
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

			if zend.ZendParseArgPath(_arg, &filename, &filename_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_PATH
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

			if zend.ZendParseArgString(_arg, &mode, &mode_len, 0) == 0 {
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

			if zend.ZendParseArgBool(_arg, &use_include_path, &_dummy, 0) == 0 {
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
			return_value.u1.type_info = 2
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
	stream = streams._phpStreamOpenWrapperEx(filename, mode, g.Cond(use_include_path != 0, 0x1, 0)|0x8, nil, context)
	if stream == nil {
		return_value.u1.type_info = 2
		return
	}
	var __z *zend.Zval = return_value
	__z.value.res = stream.res
	__z.u1.type_info = 9 | 1<<0<<8
	stream.__exposed = 1
}

/* }}} */

func ZifFclose(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var res *zend.Zval
	var stream *core.PhpStream
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

			if zend.ZendParseArgResource(_arg, &res, 0) == 0 {
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
			return_value.u1.type_info = 2
			return
		}
		break
	}
	r.Assert(res.u1.v.type_ == 9)
	if g.Assign(&stream, (*core.PhpStream)(zend.ZendFetchResource2(res.value.res, "stream", streams.PhpFileLeStream(), streams.PhpFileLePstream()))) == nil {
		return_value.u1.type_info = 2
		return
	}
	if (stream.flags & 0x80) != 0 {
		core.PhpErrorDocref(nil, 1<<1, "%d is not a valid stream resource", stream.res.handle)
		return_value.u1.type_info = 2
		return
	}
	streams._phpStreamFree(stream, 64|g.Cond(stream.is_persistent != 0, 1|2|16, 1|2))
	return_value.u1.type_info = 3
	return
}

/* }}} */

func ZifPopen(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var command *byte
	var mode *byte
	var command_len int
	var mode_len int
	var fp *r.FILE
	var stream *core.PhpStream
	var posix_mode *byte
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

			if zend.ZendParseArgPath(_arg, &command, &command_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_PATH
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

			if zend.ZendParseArgString(_arg, &mode, &mode_len, 0) == 0 {
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
	posix_mode = zend._estrndup(mode, mode_len)
	var z *byte = memchr(posix_mode, 'b', mode_len)
	if z != nil {
		memmove(z, z+1, mode_len-(z-posix_mode))
	}
	fp = popen(command, posix_mode)
	if fp == nil {
		core.PhpErrorDocref2(nil, command, posix_mode, 1<<1, "%s", strerror(errno))
		zend._efree(posix_mode)
		return_value.u1.type_info = 2
		return
	}
	stream = streams._phpStreamFopenFromPipe(fp, mode)
	if stream == nil {
		core.PhpErrorDocref2(nil, command, mode, 1<<1, "%s", strerror(errno))
		return_value.u1.type_info = 2
	} else {
		var __z *zend.Zval = return_value
		__z.value.res = stream.res
		__z.u1.type_info = 9 | 1<<0<<8
		stream.__exposed = 1
	}
	zend._efree(posix_mode)
}

/* }}} */

func ZifPclose(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var res *zend.Zval
	var stream *core.PhpStream
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

			if zend.ZendParseArgResource(_arg, &res, 0) == 0 {
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
			return_value.u1.type_info = 2
			return
		}
		break
	}
	r.Assert(res.u1.v.type_ == 9)
	if g.Assign(&stream, (*core.PhpStream)(zend.ZendFetchResource2(res.value.res, "stream", streams.PhpFileLeStream(), streams.PhpFileLePstream()))) == nil {
		return_value.u1.type_info = 2
		return
	}
	FileGlobals.SetPcloseWait(1)
	zend.ZendListClose(stream.res)
	FileGlobals.SetPcloseWait(0)
	var __z *zend.Zval = return_value
	__z.value.lval = FileGlobals.GetPcloseRet()
	__z.u1.type_info = 4
	return
}

/* }}} */

func ZifFeof(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var res *zend.Zval
	var stream *core.PhpStream
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

			if zend.ZendParseArgResource(_arg, &res, 0) == 0 {
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
			return_value.u1.type_info = 2
			return
		}
		break
	}
	r.Assert(res.u1.v.type_ == 9)
	if g.Assign(&stream, (*core.PhpStream)(zend.ZendFetchResource2(res.value.res, "stream", streams.PhpFileLeStream(), streams.PhpFileLePstream()))) == nil {
		return_value.u1.type_info = 2
		return
	}
	if streams._phpStreamEof(stream) != 0 {
		return_value.u1.type_info = 3
		return
	} else {
		return_value.u1.type_info = 2
		return
	}
}

/* }}} */

func ZifFgets(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var res *zend.Zval
	var len_ zend.ZendLong = 1024
	var buf *byte = nil
	var argc int = execute_data.This.u2.num_args
	var line_len int = 0
	var str *zend.ZendString
	var stream *core.PhpStream
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

			if zend.ZendParseArgResource(_arg, &res, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_RESOURCE
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
			return_value.u1.type_info = 2
			return
		}
		break
	}
	r.Assert(res.u1.v.type_ == 9)
	if g.Assign(&stream, (*core.PhpStream)(zend.ZendFetchResource2(res.value.res, "stream", streams.PhpFileLeStream(), streams.PhpFileLePstream()))) == nil {
		return_value.u1.type_info = 2
		return
	}
	if argc == 1 {

		/* ask streams to give us a buffer of an appropriate size */

		buf = streams._phpStreamGetLine(stream, nil, 0, &line_len)
		if buf == nil {
			return_value.u1.type_info = 2
			return
		}

		// TODO: avoid reallocation ???

		var __z *zend.Zval = return_value
		var __s *zend.ZendString = zend.ZendStringInit(buf, line_len, 0)
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
		zend._efree(buf)
	} else if argc > 1 {
		if len_ <= 0 {
			core.PhpErrorDocref(nil, 1<<1, "Length parameter must be greater than 0")
			return_value.u1.type_info = 2
			return
		}
		str = zend.ZendStringAlloc(len_, 0)
		if streams._phpStreamGetLine(stream, str.val, len_, &line_len) == nil {
			zend.ZendStringEfree(str)
			return_value.u1.type_info = 2
			return
		}

		/* resize buffer if it's much larger than the result.
		 * Only needed if the user requested a buffer size. */

		if line_len < int(len_/2) {
			str = zend.ZendStringTruncate(str, line_len, 0)
		} else {
			str.len_ = line_len
		}
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = str
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
		return
	}
}

/* }}} */

func ZifFgetc(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var res *zend.Zval
	var buf []byte
	var result int
	var stream *core.PhpStream
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

			if zend.ZendParseArgResource(_arg, &res, 0) == 0 {
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
			return_value.u1.type_info = 2
			return
		}
		break
	}
	r.Assert(res.u1.v.type_ == 9)
	if g.Assign(&stream, (*core.PhpStream)(zend.ZendFetchResource2(res.value.res, "stream", streams.PhpFileLeStream(), streams.PhpFileLePstream()))) == nil {
		return_value.u1.type_info = 2
		return
	}
	result = streams._phpStreamGetc(stream)
	if result == -1 {
		return_value.u1.type_info = 2
	} else {
		buf[0] = result
		buf[1] = '0'
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = zend.ZendStringInit(buf, 1, 0)
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
		return
	}
}

/* }}} */

func ZifFgetss(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var fd *zend.Zval
	var bytes zend.ZendLong = 0
	var len_ int = 0
	var actual_len int
	var retval_len int
	var buf *byte = nil
	var retval *byte
	var stream *core.PhpStream
	var allowed_tags *byte = nil
	var allowed_tags_len int = 0
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

			if zend.ZendParseArgResource(_arg, &fd, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_RESOURCE
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

			if zend.ZendParseArgLong(_arg, &bytes, &_dummy, 0, 0) == 0 {
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

			if zend.ZendParseArgString(_arg, &allowed_tags, &allowed_tags_len, 0) == 0 {
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
	r.Assert(fd.u1.v.type_ == 9)
	if g.Assign(&stream, (*core.PhpStream)(zend.ZendFetchResource2(fd.value.res, "stream", streams.PhpFileLeStream(), streams.PhpFileLePstream()))) == nil {
		return_value.u1.type_info = 2
		return
	}
	if execute_data.This.u2.num_args >= 2 {
		if bytes <= 0 {
			core.PhpErrorDocref(nil, 1<<1, "Length parameter must be greater than 0")
			return_value.u1.type_info = 2
			return
		}
		len_ = int(bytes)
		buf = zend._safeEmalloc(g.SizeOf("char"), len_+1, 0)

		/*needed because recv doesn't set null char at end*/

		memset(buf, 0, len_+1)

		/*needed because recv doesn't set null char at end*/

	}
	if g.Assign(&retval, streams._phpStreamGetLine(stream, buf, len_, &actual_len)) == nil {
		if buf != nil {
			zend._efree(buf)
		}
		return_value.u1.type_info = 2
		return
	}
	retval_len = PhpStripTags(retval, actual_len, &stream.fgetss_state, allowed_tags, allowed_tags_len)

	// TODO: avoid reallocation ???

	var __z *zend.Zval = return_value
	var __s *zend.ZendString = zend.ZendStringInit(retval, retval_len, 0)
	__z.value.str = __s
	__z.u1.type_info = 6 | 1<<0<<8
	zend._efree(retval)
}

/* }}} */

func ZifFscanf(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var result int
	var argc int = 0
	var format_len int
	var args *zend.Zval = nil
	var file_handle *zend.Zval
	var buf *byte
	var format *byte
	var len_ int
	var what any
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

			if zend.ZendParseArgResource(_arg, &file_handle, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_RESOURCE
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
				argc = _num_varargs
				_i += _num_varargs
				_real_arg += _num_varargs
			} else {
				args = nil
				argc = 0
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
	what = zend.ZendFetchResource2(file_handle.value.res, "File-Handle", streams.PhpFileLeStream(), streams.PhpFileLePstream())

	/* we can't do a ZEND_VERIFY_RESOURCE(what), otherwise we end up
	 * with a leak if we have an invalid filehandle. This needs changing
	 * if the code behind ZEND_VERIFY_RESOURCE changed. - cc */

	if !what {
		return_value.u1.type_info = 2
		return
	}
	buf = streams._phpStreamGetLine((*core.PhpStream)(what), nil, 0, &len_)
	if buf == nil {
		return_value.u1.type_info = 2
		return
	}
	result = PhpSscanfInternal(buf, format, argc, args, 0, return_value)
	zend._efree(buf)
	if -1-1-1-1 == result {
		zend.ZendWrongParamCount()
		return
	}
}

/* }}} */

func ZifFwrite(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var res *zend.Zval
	var input *byte
	var inputlen int
	var ret ssize_t
	var num_bytes int
	var maxlen zend.ZendLong = 0
	var stream *core.PhpStream
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

			if zend.ZendParseArgResource(_arg, &res, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_RESOURCE
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

			if zend.ZendParseArgString(_arg, &input, &inputlen, 0) == 0 {
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

			if zend.ZendParseArgLong(_arg, &maxlen, &_dummy, 0, 0) == 0 {
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
	if execute_data.This.u2.num_args == 2 {
		num_bytes = inputlen
	} else if maxlen <= 0 {
		num_bytes = 0
	} else {
		if int(maxlen) < inputlen {
			num_bytes = int(maxlen)
		} else {
			num_bytes = inputlen
		}
	}
	if num_bytes == 0 {
		var __z *zend.Zval = return_value
		__z.value.lval = 0
		__z.u1.type_info = 4
		return
	}
	r.Assert(res.u1.v.type_ == 9)
	if g.Assign(&stream, (*core.PhpStream)(zend.ZendFetchResource2(res.value.res, "stream", streams.PhpFileLeStream(), streams.PhpFileLePstream()))) == nil {
		return_value.u1.type_info = 2
		return
	}
	ret = streams._phpStreamWrite(stream, input, num_bytes)
	if ret < 0 {
		return_value.u1.type_info = 2
		return
	}
	var __z *zend.Zval = return_value
	__z.value.lval = ret
	__z.u1.type_info = 4
	return
}

/* }}} */

func ZifFflush(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var res *zend.Zval
	var ret int
	var stream *core.PhpStream
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

			if zend.ZendParseArgResource(_arg, &res, 0) == 0 {
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
			return_value.u1.type_info = 2
			return
		}
		break
	}
	r.Assert(res.u1.v.type_ == 9)
	if g.Assign(&stream, (*core.PhpStream)(zend.ZendFetchResource2(res.value.res, "stream", streams.PhpFileLeStream(), streams.PhpFileLePstream()))) == nil {
		return_value.u1.type_info = 2
		return
	}
	ret = streams._phpStreamFlush(stream, 0)
	if ret != 0 {
		return_value.u1.type_info = 2
		return
	}
	return_value.u1.type_info = 3
	return
}

/* }}} */

func ZifRewind(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var res *zend.Zval
	var stream *core.PhpStream
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

			if zend.ZendParseArgResource(_arg, &res, 0) == 0 {
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
			return_value.u1.type_info = 2
			return
		}
		break
	}
	r.Assert(res.u1.v.type_ == 9)
	if g.Assign(&stream, (*core.PhpStream)(zend.ZendFetchResource2(res.value.res, "stream", streams.PhpFileLeStream(), streams.PhpFileLePstream()))) == nil {
		return_value.u1.type_info = 2
		return
	}
	if -1 == streams._phpStreamSeek(stream, 0, 0) {
		return_value.u1.type_info = 2
		return
	}
	return_value.u1.type_info = 3
	return
}

/* }}} */

func ZifFtell(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var res *zend.Zval
	var ret zend.ZendLong
	var stream *core.PhpStream
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

			if zend.ZendParseArgResource(_arg, &res, 0) == 0 {
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
			return_value.u1.type_info = 2
			return
		}
		break
	}
	r.Assert(res.u1.v.type_ == 9)
	if g.Assign(&stream, (*core.PhpStream)(zend.ZendFetchResource2(res.value.res, "stream", streams.PhpFileLeStream(), streams.PhpFileLePstream()))) == nil {
		return_value.u1.type_info = 2
		return
	}
	ret = streams._phpStreamTell(stream)
	if ret == -1 {
		return_value.u1.type_info = 2
		return
	}
	var __z *zend.Zval = return_value
	__z.value.lval = ret
	__z.u1.type_info = 4
	return
}

/* }}} */

func ZifFseek(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var res *zend.Zval
	var offset zend.ZendLong
	var whence zend.ZendLong = 0
	var stream *core.PhpStream
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

			if zend.ZendParseArgResource(_arg, &res, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_RESOURCE
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

			if zend.ZendParseArgLong(_arg, &whence, &_dummy, 0, 0) == 0 {
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
	r.Assert(res.u1.v.type_ == 9)
	if g.Assign(&stream, (*core.PhpStream)(zend.ZendFetchResource2(res.value.res, "stream", streams.PhpFileLeStream(), streams.PhpFileLePstream()))) == nil {
		return_value.u1.type_info = 2
		return
	}
	var __z *zend.Zval = return_value
	__z.value.lval = streams._phpStreamSeek(stream, offset, int(whence))
	__z.u1.type_info = 4
	return
}

/* }}} */

func PhpMkdirEx(dir *byte, mode zend.ZendLong, options int) int {
	var ret int
	if core.PhpCheckOpenBasedir(dir) != 0 {
		return -1
	}
	if g.Assign(&ret, mkdir(dir, mode_t(mode))) < 0 && (options&0x8) != 0 {
		core.PhpErrorDocref(nil, 1<<1, "%s", strerror(errno))
	}
	return ret
}
func PhpMkdir(dir *byte, mode zend.ZendLong) int { return PhpMkdirEx(dir, mode, 0x8) }

/* }}} */

func ZifMkdir(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var dir *byte
	var dir_len int
	var zcontext *zend.Zval = nil
	var mode zend.ZendLong = 0777
	var recursive zend.ZendBool = 0
	var context *core.PhpStreamContext
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

			if zend.ZendParseArgPath(_arg, &dir, &dir_len, 0) == 0 {
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

			if zend.ZendParseArgLong(_arg, &mode, &_dummy, 0, 0) == 0 {
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

			if zend.ZendParseArgBool(_arg, &recursive, &_dummy, 0) == 0 {
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
			return_value.u1.type_info = 2
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
	if streams._phpStreamMkdir(dir, int(mode), g.Cond(recursive != 0, 1, 0)|0x8, context) != 0 {
		return_value.u1.type_info = 3
	} else {
		return_value.u1.type_info = 2
	}
	return
}

/* }}} */

func ZifRmdir(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var dir *byte
	var dir_len int
	var zcontext *zend.Zval = nil
	var context *core.PhpStreamContext
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

			if zend.ZendParseArgPath(_arg, &dir, &dir_len, 0) == 0 {
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
			return_value.u1.type_info = 2
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
	if streams._phpStreamRmdir(dir, 0x8, context) != 0 {
		return_value.u1.type_info = 3
	} else {
		return_value.u1.type_info = 2
	}
	return
}

/* }}} */

func ZifReadfile(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var filename *byte
	var filename_len int
	var size int = 0
	var use_include_path zend.ZendBool = 0
	var zcontext *zend.Zval = nil
	var stream *core.PhpStream
	var context *core.PhpStreamContext = nil
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

			if zend.ZendParseArgPath(_arg, &filename, &filename_len, 0) == 0 {
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

			if zend.ZendParseArgBool(_arg, &use_include_path, &_dummy, 0) == 0 {
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
			return_value.u1.type_info = 2
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
	stream = streams._phpStreamOpenWrapperEx(filename, "rb", g.Cond(use_include_path != 0, 0x1, 0)|0x8, nil, context)
	if stream != nil {
		size = streams._phpStreamPassthru(stream)
		streams._phpStreamFree(stream, 1|2)
		var __z *zend.Zval = return_value
		__z.value.lval = size
		__z.u1.type_info = 4
		return
	}
	return_value.u1.type_info = 2
	return
}

/* }}} */

func ZifUmask(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var mask zend.ZendLong = 0
	var oldumask int
	oldumask = umask(077)
	if BasicGlobals.GetUmask() == -1 {
		BasicGlobals.SetUmask(oldumask)
	}
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

			if zend.ZendParseArgLong(_arg, &mask, &_dummy, 0, 0) == 0 {
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
	if execute_data.This.u2.num_args == 0 {
		umask(oldumask)
	} else {
		umask(int(mask))
	}
	var __z *zend.Zval = return_value
	__z.value.lval = oldumask
	__z.u1.type_info = 4
	return
}

/* }}} */

func ZifFpassthru(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var res *zend.Zval
	var size int
	var stream *core.PhpStream
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

			if zend.ZendParseArgResource(_arg, &res, 0) == 0 {
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
			return_value.u1.type_info = 2
			return
		}
		break
	}
	r.Assert(res.u1.v.type_ == 9)
	if g.Assign(&stream, (*core.PhpStream)(zend.ZendFetchResource2(res.value.res, "stream", streams.PhpFileLeStream(), streams.PhpFileLePstream()))) == nil {
		return_value.u1.type_info = 2
		return
	}
	size = streams._phpStreamPassthru(stream)
	var __z *zend.Zval = return_value
	__z.value.lval = size
	__z.u1.type_info = 4
	return
}

/* }}} */

func ZifRename(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var old_name *byte
	var new_name *byte
	var old_name_len int
	var new_name_len int
	var zcontext *zend.Zval = nil
	var wrapper *core.PhpStreamWrapper
	var context *core.PhpStreamContext
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

			if zend.ZendParseArgPath(_arg, &old_name, &old_name_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_PATH
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

			if zend.ZendParseArgPath(_arg, &new_name, &new_name_len, 0) == 0 {
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
			return_value.u1.type_info = 2
			return
		}
		break
	}
	wrapper = streams.PhpStreamLocateUrlWrapper(old_name, nil, 0)
	if wrapper == nil || wrapper.wops == nil {
		core.PhpErrorDocref(nil, 1<<1, "Unable to locate stream wrapper")
		return_value.u1.type_info = 2
		return
	}
	if wrapper.wops.rename == nil {
		core.PhpErrorDocref(nil, 1<<1, "%s wrapper does not support renaming", g.CondF1(wrapper.wops.label != nil, func() *byte { return wrapper.wops.label }, "Source"))
		return_value.u1.type_info = 2
		return
	}
	if wrapper != streams.PhpStreamLocateUrlWrapper(new_name, nil, 0) {
		core.PhpErrorDocref(nil, 1<<1, "Cannot rename a file across wrapper types")
		return_value.u1.type_info = 2
		return
	}
	if g.CondF2(g.CondF1(zcontext != nil, func() any { return zend.ZendFetchResourceEx(zcontext, "Stream-Context", PhpLeStreamContext()) }, 0), nil, func() *core.PhpStreamContext { return FileGlobals.GetDefaultContext() }) {
		context = FileGlobals.GetDefaultContext()
	} else {
		FileGlobals.SetDefaultContext(streams.PhpStreamContextAlloc())
		context = FileGlobals.GetDefaultContext()
	}
	if wrapper.wops.rename(wrapper, old_name, new_name, 0, context) != 0 {
		return_value.u1.type_info = 3
	} else {
		return_value.u1.type_info = 2
	}
	return
}

/* }}} */

func ZifUnlink(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var filename *byte
	var filename_len int
	var wrapper *core.PhpStreamWrapper
	var zcontext *zend.Zval = nil
	var context *core.PhpStreamContext = nil
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

			if zend.ZendParseArgPath(_arg, &filename, &filename_len, 0) == 0 {
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
			return_value.u1.type_info = 2
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
	wrapper = streams.PhpStreamLocateUrlWrapper(filename, nil, 0)
	if wrapper == nil || wrapper.wops == nil {
		core.PhpErrorDocref(nil, 1<<1, "Unable to locate stream wrapper")
		return_value.u1.type_info = 2
		return
	}
	if wrapper.wops.unlink == nil {
		core.PhpErrorDocref(nil, 1<<1, "%s does not allow unlinking", g.CondF1(wrapper.wops.label != nil, func() *byte { return wrapper.wops.label }, "Wrapper"))
		return_value.u1.type_info = 2
		return
	}
	if wrapper.wops.unlink(wrapper, filename, 0x8, context) != 0 {
		return_value.u1.type_info = 3
	} else {
		return_value.u1.type_info = 2
	}
	return
}

/* }}} */

func PhpIfFtruncate(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var fp *zend.Zval
	var size zend.ZendLong
	var stream *core.PhpStream
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

			if zend.ZendParseArgResource(_arg, &fp, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_RESOURCE
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

			if zend.ZendParseArgLong(_arg, &size, &_dummy, 0, 0) == 0 {
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
	if size < 0 {
		core.PhpErrorDocref(nil, 1<<1, "Negative size is not supported")
		return_value.u1.type_info = 2
		return
	}
	r.Assert(fp.u1.v.type_ == 9)
	if g.Assign(&stream, (*core.PhpStream)(zend.ZendFetchResource2(fp.value.res, "stream", streams.PhpFileLeStream(), streams.PhpFileLePstream()))) == nil {
		return_value.u1.type_info = 2
		return
	}
	if !(g.Cond(streams._phpStreamSetOption(stream, 10, 0, nil) == 0, 1, 0)) {
		core.PhpErrorDocref(nil, 1<<1, "Can't truncate this stream!")
		return_value.u1.type_info = 2
		return
	}
	if 0 == streams._phpStreamTruncateSetSize(stream, size) {
		return_value.u1.type_info = 3
	} else {
		return_value.u1.type_info = 2
	}
	return
}

/* }}} */

func PhpIfFstat(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var fp *zend.Zval
	var stat_dev zend.Zval
	var stat_ino zend.Zval
	var stat_mode zend.Zval
	var stat_nlink zend.Zval
	var stat_uid zend.Zval
	var stat_gid zend.Zval
	var stat_rdev zend.Zval
	var stat_size zend.Zval
	var stat_atime zend.Zval
	var stat_mtime zend.Zval
	var stat_ctime zend.Zval
	var stat_blksize zend.Zval
	var stat_blocks zend.Zval
	var stream *core.PhpStream
	var stat_ssb core.PhpStreamStatbuf
	var stat_sb_names []*byte = []*byte{"dev", "ino", "mode", "nlink", "uid", "gid", "rdev", "size", "atime", "mtime", "ctime", "blksize", "blocks"}
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

			if zend.ZendParseArgResource(_arg, &fp, 0) == 0 {
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
			return_value.u1.type_info = 2
			return
		}
		break
	}
	r.Assert(fp.u1.v.type_ == 9)
	if g.Assign(&stream, (*core.PhpStream)(zend.ZendFetchResource2(fp.value.res, "stream", streams.PhpFileLeStream(), streams.PhpFileLePstream()))) == nil {
		return_value.u1.type_info = 2
		return
	}
	if streams._phpStreamStat(stream, &stat_ssb) != 0 {
		return_value.u1.type_info = 2
		return
	}
	var __arr *zend.ZendArray = zend._zendNewArray(0)
	var __z *zend.Zval = return_value
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	var __z *zval = &stat_dev
	__z.value.lval = stat_ssb.sb.st_dev
	__z.u1.type_info = 4
	var __z *zval = &stat_ino
	__z.value.lval = stat_ssb.sb.st_ino
	__z.u1.type_info = 4
	var __z *zval = &stat_mode
	__z.value.lval = stat_ssb.sb.st_mode
	__z.u1.type_info = 4
	var __z *zval = &stat_nlink
	__z.value.lval = stat_ssb.sb.st_nlink
	__z.u1.type_info = 4
	var __z *zval = &stat_uid
	__z.value.lval = stat_ssb.sb.st_uid
	__z.u1.type_info = 4
	var __z *zval = &stat_gid
	__z.value.lval = stat_ssb.sb.st_gid
	__z.u1.type_info = 4
	var __z *zval = &stat_rdev
	__z.value.lval = stat_ssb.sb.st_rdev
	__z.u1.type_info = 4
	var __z *zval = &stat_size
	__z.value.lval = stat_ssb.sb.st_size
	__z.u1.type_info = 4
	var __z *zval = &stat_atime
	__z.value.lval = stat_ssb.sb.st_atime
	__z.u1.type_info = 4
	var __z *zval = &stat_mtime
	__z.value.lval = stat_ssb.sb.st_mtime
	__z.u1.type_info = 4
	var __z *zval = &stat_ctime
	__z.value.lval = stat_ssb.sb.st_ctime
	__z.u1.type_info = 4
	var __z *zval = &stat_blksize
	__z.value.lval = stat_ssb.sb.st_blksize
	__z.u1.type_info = 4
	var __z *zend.Zval = &stat_blocks
	__z.value.lval = stat_ssb.sb.st_blocks
	__z.u1.type_info = 4

	/* Store numeric indexes in proper order */

	zend.ZendHashNextIndexInsert(return_value.value.arr, &stat_dev)
	zend.ZendHashNextIndexInsert(return_value.value.arr, &stat_ino)
	zend.ZendHashNextIndexInsert(return_value.value.arr, &stat_mode)
	zend.ZendHashNextIndexInsert(return_value.value.arr, &stat_nlink)
	zend.ZendHashNextIndexInsert(return_value.value.arr, &stat_uid)
	zend.ZendHashNextIndexInsert(return_value.value.arr, &stat_gid)
	zend.ZendHashNextIndexInsert(return_value.value.arr, &stat_rdev)
	zend.ZendHashNextIndexInsert(return_value.value.arr, &stat_size)
	zend.ZendHashNextIndexInsert(return_value.value.arr, &stat_atime)
	zend.ZendHashNextIndexInsert(return_value.value.arr, &stat_mtime)
	zend.ZendHashNextIndexInsert(return_value.value.arr, &stat_ctime)
	zend.ZendHashNextIndexInsert(return_value.value.arr, &stat_blksize)
	zend.ZendHashNextIndexInsert(return_value.value.arr, &stat_blocks)

	/* Store string indexes referencing the same zval*/

	zend.ZendHashStrAddNew(return_value.value.arr, stat_sb_names[0], strlen(stat_sb_names[0]), &stat_dev)
	zend.ZendHashStrAddNew(return_value.value.arr, stat_sb_names[1], strlen(stat_sb_names[1]), &stat_ino)
	zend.ZendHashStrAddNew(return_value.value.arr, stat_sb_names[2], strlen(stat_sb_names[2]), &stat_mode)
	zend.ZendHashStrAddNew(return_value.value.arr, stat_sb_names[3], strlen(stat_sb_names[3]), &stat_nlink)
	zend.ZendHashStrAddNew(return_value.value.arr, stat_sb_names[4], strlen(stat_sb_names[4]), &stat_uid)
	zend.ZendHashStrAddNew(return_value.value.arr, stat_sb_names[5], strlen(stat_sb_names[5]), &stat_gid)
	zend.ZendHashStrAddNew(return_value.value.arr, stat_sb_names[6], strlen(stat_sb_names[6]), &stat_rdev)
	zend.ZendHashStrAddNew(return_value.value.arr, stat_sb_names[7], strlen(stat_sb_names[7]), &stat_size)
	zend.ZendHashStrAddNew(return_value.value.arr, stat_sb_names[8], strlen(stat_sb_names[8]), &stat_atime)
	zend.ZendHashStrAddNew(return_value.value.arr, stat_sb_names[9], strlen(stat_sb_names[9]), &stat_mtime)
	zend.ZendHashStrAddNew(return_value.value.arr, stat_sb_names[10], strlen(stat_sb_names[10]), &stat_ctime)
	zend.ZendHashStrAddNew(return_value.value.arr, stat_sb_names[11], strlen(stat_sb_names[11]), &stat_blksize)
	zend.ZendHashStrAddNew(return_value.value.arr, stat_sb_names[12], strlen(stat_sb_names[12]), &stat_blocks)
}

/* }}} */

func ZifCopy(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var source *byte
	var target *byte
	var source_len int
	var target_len int
	var zcontext *zend.Zval = nil
	var context *core.PhpStreamContext
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

			if zend.ZendParseArgPath(_arg, &source, &source_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_PATH
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

			if zend.ZendParseArgPath(_arg, &target, &target_len, 0) == 0 {
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
	if streams.PhpStreamLocateUrlWrapper(source, nil, 0) == &PhpPlainFilesWrapper && core.PhpCheckOpenBasedir(source) != 0 {
		return_value.u1.type_info = 2
		return
	}
	if g.CondF2(g.CondF1(zcontext != nil, func() any { return zend.ZendFetchResourceEx(zcontext, "Stream-Context", PhpLeStreamContext()) }, 0), nil, func() *core.PhpStreamContext { return FileGlobals.GetDefaultContext() }) {
		context = FileGlobals.GetDefaultContext()
	} else {
		FileGlobals.SetDefaultContext(streams.PhpStreamContextAlloc())
		context = FileGlobals.GetDefaultContext()
	}
	if PhpCopyFileCtx(source, target, 0, context) == zend.SUCCESS {
		return_value.u1.type_info = 3
		return
	} else {
		return_value.u1.type_info = 2
		return
	}
}

/* }}} */

func PhpCopyFile(src *byte, dest *byte) int { return PhpCopyFileCtx(src, dest, 0, nil) }

/* }}} */

func PhpCopyFileEx(src *byte, dest *byte, src_flg int) int {
	return PhpCopyFileCtx(src, dest, src_flg, nil)
}

/* }}} */

func PhpCopyFileCtx(src *byte, dest *byte, src_flg int, ctx *core.PhpStreamContext) int {
	var srcstream *core.PhpStream = nil
	var deststream *core.PhpStream = nil
	var ret int = zend.FAILURE
	var src_s core.PhpStreamStatbuf
	var dest_s core.PhpStreamStatbuf
	switch streams._phpStreamStatPath(src, 0, &src_s, ctx) {
	case -1:

		/* non-statable stream */

		goto safe_to_copy
		break
	case 0:
		break
	default:
		return ret
	}
	if (src_s.sb.st_mode & S_IFMT) == S_IFDIR {
		core.PhpErrorDocref(nil, 1<<1, "The first argument to copy() function cannot be a directory")
		return zend.FAILURE
	}
	switch streams._phpStreamStatPath(dest, 2|4, &dest_s, ctx) {
	case -1:

		/* non-statable stream */

		goto safe_to_copy
		break
	case 0:
		break
	default:
		return ret
	}
	if (dest_s.sb.st_mode & S_IFMT) == S_IFDIR {
		core.PhpErrorDocref(nil, 1<<1, "The second argument to copy() function cannot be a directory")
		return zend.FAILURE
	}
	if !(src_s.sb.st_ino) || !(dest_s.sb.st_ino) {
		goto no_stat
	}
	if src_s.sb.st_ino == dest_s.sb.st_ino && src_s.sb.st_dev == dest_s.sb.st_dev {
		return ret
	} else {
		goto safe_to_copy
	}
no_stat:
	var sp *byte
	var dp *byte
	var res int
	if g.Assign(&sp, core.ExpandFilepath(src, nil)) == nil {
		return ret
	}
	if g.Assign(&dp, core.ExpandFilepath(dest, nil)) == nil {
		zend._efree(sp)
		goto safe_to_copy
	}
	res = !(strcmp(sp, dp))
	zend._efree(sp)
	zend._efree(dp)
	if res != 0 {
		return ret
	}
safe_to_copy:
	srcstream = streams._phpStreamOpenWrapperEx(src, "rb", src_flg|0x8, nil, ctx)
	if srcstream == nil {
		return ret
	}
	deststream = streams._phpStreamOpenWrapperEx(dest, "wb", 0x8, nil, ctx)
	if srcstream != nil && deststream != nil {
		ret = streams._phpStreamCopyToStreamEx(srcstream, deststream, size_t-1, nil)
	}
	if srcstream != nil {
		streams._phpStreamFree(srcstream, 1|2)
	}
	if deststream != nil {
		streams._phpStreamFree(deststream, 1|2)
	}
	return ret
}

/* }}} */

func ZifFread(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var res *zend.Zval
	var len_ zend.ZendLong
	var stream *core.PhpStream
	var str *zend.ZendString
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

			if zend.ZendParseArgResource(_arg, &res, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_RESOURCE
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
			return_value.u1.type_info = 2
			return
		}
		break
	}
	r.Assert(res.u1.v.type_ == 9)
	if g.Assign(&stream, (*core.PhpStream)(zend.ZendFetchResource2(res.value.res, "stream", streams.PhpFileLeStream(), streams.PhpFileLePstream()))) == nil {
		return_value.u1.type_info = 2
		return
	}
	if len_ <= 0 {
		core.PhpErrorDocref(nil, 1<<1, "Length parameter must be greater than 0")
		return_value.u1.type_info = 2
		return
	}
	str = streams.PhpStreamReadToStr(stream, len_)
	if str == nil {
		zend.ZvalPtrDtorStr(return_value)
		return_value.u1.type_info = 2
		return
	}
	var __z *zend.Zval = return_value
	var __s *zend.ZendString = str
	__z.value.str = __s
	if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
		__z.u1.type_info = 6
	} else {
		__z.u1.type_info = 6 | 1<<0<<8
	}
	return
}

/* }}} */

func PhpFgetcsvLookupTrailingSpaces(ptr *byte, len_ int, delimiter byte) *byte {
	var inc_len int
	var last_chars []uint8 = []uint8{0, 0}
	for len_ > 0 {
		if (*ptr) == '0' {
			inc_len = 1
		} else {
			inc_len = mblen(ptr, len_)
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

		default:
			last_chars[0] = last_chars[1]
			last_chars[1] = *ptr
			break
		}
		ptr += inc_len
		len_ -= inc_len
	}
quit_loop:
	switch last_chars[1] {
	case '\n':
		if last_chars[0] == '\r' {
			return ptr - 2
		}
	case '\r':
		return ptr - 1
	}
	return ptr
}

/* }}} */

// #define FPUTCSV_FLD_CHK(c) memchr ( ZSTR_VAL ( field_str ) , c , ZSTR_LEN ( field_str ) )

/* {{{ proto int fputcsv(resource fp, array fields [, string delimiter [, string enclosure [, string escape_char]]])
   Format line as CSV and write to file pointer */

func ZifFputcsv(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var delimiter byte = ','
	var enclosure byte = '"'
	var escape_char int = uint8('\\')
	var stream *core.PhpStream
	var fp *zend.Zval = nil
	var fields *zend.Zval = nil
	var ret ssize_t
	var delimiter_str *byte = nil
	var enclosure_str *byte = nil
	var escape_str *byte = nil
	var delimiter_str_len int = 0
	var enclosure_str_len int = 0
	var escape_str_len int = 0
	for {
		var _flags int = 0
		var _min_num_args int = 2
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

			if zend.ZendParseArgResource(_arg, &fp, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_RESOURCE
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

			if zend.ZendParseArgArray(_arg, &fields, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_ARRAY
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

			if zend.ZendParseArgString(_arg, &delimiter_str, &delimiter_str_len, 0) == 0 {
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

			if zend.ZendParseArgString(_arg, &enclosure_str, &enclosure_str_len, 0) == 0 {
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

			if zend.ZendParseArgString(_arg, &escape_str, &escape_str_len, 0) == 0 {
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
	if delimiter_str != nil {

		/* Make sure that there is at least one character in string */

		if delimiter_str_len < 1 {
			core.PhpErrorDocref(nil, 1<<1, "delimiter must be a character")
			return_value.u1.type_info = 2
			return
		} else if delimiter_str_len > 1 {
			core.PhpErrorDocref(nil, 1<<3, "delimiter must be a single character")
		}

		/* use first character from string */

		delimiter = *delimiter_str

		/* use first character from string */

	}
	if enclosure_str != nil {
		if enclosure_str_len < 1 {
			core.PhpErrorDocref(nil, 1<<1, "enclosure must be a character")
			return_value.u1.type_info = 2
			return
		} else if enclosure_str_len > 1 {
			core.PhpErrorDocref(nil, 1<<3, "enclosure must be a single character")
		}

		/* use first character from string */

		enclosure = *enclosure_str

		/* use first character from string */

	}
	if escape_str != nil {
		if escape_str_len > 1 {
			core.PhpErrorDocref(nil, 1<<3, "escape must be empty or a single character")
		}
		if escape_str_len < 1 {
			escape_char = -1
		} else {

			/* use first character from string */

			escape_char = uint8(*escape_str)

			/* use first character from string */

		}
	}
	r.Assert(fp.u1.v.type_ == 9)
	if g.Assign(&stream, (*core.PhpStream)(zend.ZendFetchResource2(fp.value.res, "stream", streams.PhpFileLeStream(), streams.PhpFileLePstream()))) == nil {
		return_value.u1.type_info = 2
		return
	}
	ret = PhpFputcsv(stream, fields, delimiter, enclosure, escape_char)
	if ret < 0 {
		return_value.u1.type_info = 2
		return
	}
	var __z *zend.Zval = return_value
	__z.value.lval = ret
	__z.u1.type_info = 4
	return
}

/* }}} */

func PhpFputcsv(stream *core.PhpStream, fields *zend.Zval, delimiter byte, enclosure byte, escape_char int) ssize_t {
	var count int
	var i int = 0
	var ret int
	var field_tmp *zend.Zval
	var csvline zend.SmartStr = zend.SmartStr{0}
	r.Assert(escape_char >= 0 && escape_char <= UCHAR_MAX || escape_char == -1)
	count = fields.value.arr.nNumOfElements
	for {
		var __ht *zend.HashTable = fields.value.arr
		var _p *zend.Bucket = __ht.arData
		var _end *zend.Bucket = _p + __ht.nNumUsed
		for ; _p != _end; _p++ {
			var _z *zend.Zval = &_p.val

			if _z.u1.v.type_ == 0 {
				continue
			}
			field_tmp = _z
			var tmp_field_str *zend.ZendString
			var field_str *zend.ZendString = zend.ZvalGetTmpString(field_tmp, &tmp_field_str)

			/* enclose a field that contains a delimiter, an enclosure character, or a newline */

			if memchr(field_str.val, delimiter, field_str.len_) || memchr(field_str.val, enclosure, field_str.len_) || escape_char != -1 && memchr(field_str.val, escape_char, field_str.len_) || memchr(field_str.val, '\n', field_str.len_) || memchr(field_str.val, '\r', field_str.len_) || memchr(field_str.val, '\t', field_str.len_) || memchr(field_str.val, ' ', field_str.len_) {
				var ch *byte = field_str.val
				var end *byte = ch + field_str.len_
				var escaped int = 0
				zend.SmartStrAppendcEx(&csvline, enclosure, 0)
				for ch < end {
					if escape_char != -1 && (*ch) == escape_char {
						escaped = 1
					} else if escaped == 0 && (*ch) == enclosure {
						zend.SmartStrAppendcEx(&csvline, enclosure, 0)
					} else {
						escaped = 0
					}
					zend.SmartStrAppendcEx(&csvline, *ch, 0)
					ch++
				}
				zend.SmartStrAppendcEx(&csvline, enclosure, 0)
			} else {
				zend.SmartStrAppendEx(&csvline, field_str, 0)
			}
			if g.PreInc(&i) != count {
				zend.SmartStrAppendlEx(&csvline, &delimiter, 1, 0)
			}
			zend.ZendTmpStringRelease(tmp_field_str)
		}
		break
	}
	zend.SmartStrAppendcEx(&csvline, '\n', 0)
	zend.SmartStr0(&csvline)
	ret = streams._phpStreamWrite(stream, csvline.s.val, csvline.s.len_)
	zend.SmartStrFreeEx(&csvline, 0)
	return ret
}

/* }}} */

func ZifFgetcsv(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var delimiter byte = ','
	var enclosure byte = '"'
	var escape int = uint8('\\')

	/* first section exactly as php_fgetss */

	var len_ zend.ZendLong = 0
	var buf_len int
	var buf *byte
	var stream *core.PhpStream
	var fd *zend.Zval
	var len_zv *zend.Zval = nil
	var delimiter_str *byte = nil
	var delimiter_str_len int = 0
	var enclosure_str *byte = nil
	var enclosure_str_len int = 0
	var escape_str *byte = nil
	var escape_str_len int = 0
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

			if zend.ZendParseArgResource(_arg, &fd, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_RESOURCE
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

			zend.ZendParseArgZvalDeref(_arg, &len_zv, 0)
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

			if zend.ZendParseArgString(_arg, &delimiter_str, &delimiter_str_len, 0) == 0 {
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

			if zend.ZendParseArgString(_arg, &enclosure_str, &enclosure_str_len, 0) == 0 {
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

			if zend.ZendParseArgString(_arg, &escape_str, &escape_str_len, 0) == 0 {
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
	if delimiter_str != nil {

		/* Make sure that there is at least one character in string */

		if delimiter_str_len < 1 {
			core.PhpErrorDocref(nil, 1<<1, "delimiter must be a character")
			return_value.u1.type_info = 2
			return
		} else if delimiter_str_len > 1 {
			core.PhpErrorDocref(nil, 1<<3, "delimiter must be a single character")
		}

		/* use first character from string */

		delimiter = delimiter_str[0]

		/* use first character from string */

	}
	if enclosure_str != nil {
		if enclosure_str_len < 1 {
			core.PhpErrorDocref(nil, 1<<1, "enclosure must be a character")
			return_value.u1.type_info = 2
			return
		} else if enclosure_str_len > 1 {
			core.PhpErrorDocref(nil, 1<<3, "enclosure must be a single character")
		}

		/* use first character from string */

		enclosure = enclosure_str[0]

		/* use first character from string */

	}
	if escape_str != nil {
		if escape_str_len > 1 {
			core.PhpErrorDocref(nil, 1<<3, "escape must be empty or a single character")
		}
		if escape_str_len < 1 {
			escape = -1
		} else {
			escape = uint8(escape_str[0])
		}
	}
	if len_zv != nil && len_zv.u1.v.type_ != 1 {
		len_ = zend.ZvalGetLong(len_zv)
		if len_ < 0 {
			core.PhpErrorDocref(nil, 1<<1, "Length parameter may not be negative")
			return_value.u1.type_info = 2
			return
		} else if len_ == 0 {
			len_ = -1
		}
	} else {
		len_ = -1
	}
	r.Assert(fd.u1.v.type_ == 9)
	if g.Assign(&stream, (*core.PhpStream)(zend.ZendFetchResource2(fd.value.res, "stream", streams.PhpFileLeStream(), streams.PhpFileLePstream()))) == nil {
		return_value.u1.type_info = 2
		return
	}
	if len_ < 0 {
		if g.Assign(&buf, streams._phpStreamGetLine(stream, nil, 0, &buf_len)) == nil {
			return_value.u1.type_info = 2
			return
		}
	} else {
		buf = zend._emalloc(len_ + 1)
		if streams._phpStreamGetLine(stream, buf, len_+1, &buf_len) == nil {
			zend._efree(buf)
			return_value.u1.type_info = 2
			return
		}
	}
	PhpFgetcsv(stream, delimiter, enclosure, escape, buf_len, buf, return_value)
}

/* }}} */

func PhpFgetcsv(stream *core.PhpStream, delimiter byte, enclosure byte, escape_char int, buf_len int, buf *byte, return_value *zend.Zval) {
	var temp *byte
	var tptr *byte
	var bptr *byte
	var line_end *byte
	var limit *byte
	var temp_len int
	var line_end_len int
	var inc_len int
	var first_field zend.ZendBool = 1
	r.Assert(escape_char >= 0 && escape_char <= UCHAR_MAX || escape_char == -1)

	/* initialize internal state */

	void(mblen(nil, 0))

	/* Now into new section that parses buf for delimiter/enclosure fields */

	bptr = buf
	tptr = (*byte)(PhpFgetcsvLookupTrailingSpaces(buf, buf_len, delimiter))
	line_end_len = buf_len - size_t(tptr-buf)
	limit = tptr
	line_end = limit

	/* reserve workspace for building each individual field */

	temp_len = buf_len
	temp = zend._emalloc(temp_len + line_end_len + 1)

	/* Initialize return array */

	var __arr *zend.ZendArray = zend._zendNewArray(0)
	var __z *zend.Zval = return_value
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8

	/* Main loop to read CSV fields */

	for {
		var comp_end *byte
		var hunk_begin *byte
		tptr = temp
		if bptr < limit {
			if (*bptr) == '0' {
				inc_len = 1
			} else {
				inc_len = mblen(bptr, limit-bptr)
			}
		} else {
			inc_len = 0
		}
		if inc_len == 1 {
			var tmp *byte = bptr
			for (*tmp) != delimiter && isspace(int(*((*uint8)(tmp)))) {
				tmp++
			}
			if (*tmp) == enclosure {
				bptr = tmp
			}
		}
		if first_field != 0 && bptr == line_end {
			zend.AddNextIndexNull(return_value)
			break
		}
		first_field = 0

		/* 2. Read field, leaving bptr pointing at start of next field */

		if inc_len != 0 && (*bptr) == enclosure {
			var state int = 0
			bptr++
			hunk_begin = bptr

			/* 2A. handle enclosure delimited field */

			for {
				switch inc_len {
				case 0:
					switch state {
					case 2:
						memcpy(tptr, hunk_begin, bptr-hunk_begin-1)
						tptr += bptr - hunk_begin - 1
						hunk_begin = bptr
						goto quit_loop_2
					case 1:
						memcpy(tptr, hunk_begin, bptr-hunk_begin)
						tptr += bptr - hunk_begin
						hunk_begin = bptr
					case 0:
						var new_buf *byte
						var new_len int
						var new_temp *byte
						if hunk_begin != line_end {
							memcpy(tptr, hunk_begin, bptr-hunk_begin)
							tptr += bptr - hunk_begin
							hunk_begin = bptr
						}

						/* add the embedded line end to the field */

						memcpy(tptr, line_end, line_end_len)
						tptr += line_end_len
						if stream == nil {
							goto quit_loop_2
						} else if g.Assign(&new_buf, streams._phpStreamGetLine(stream, nil, 0, &new_len)) == nil {

							/* we've got an unterminated enclosure,
							 * assign all the data from the start of
							 * the enclosure to end of data to the
							 * last element */

							if int(temp_len > size_t(limit-buf)) != 0 {
								goto quit_loop_2
							}
							zend.ZendArrayDestroy(return_value.value.arr)
							return_value.u1.type_info = 2
							goto out
						}
						temp_len += new_len
						new_temp = zend._erealloc(temp, temp_len)
						tptr = new_temp + size_t(tptr-temp)
						temp = new_temp
						zend._efree(buf)
						buf_len = new_len
						buf = new_buf
						bptr = buf
						hunk_begin = buf
						limit = (*byte)(PhpFgetcsvLookupTrailingSpaces(buf, buf_len, delimiter))
						line_end = limit
						line_end_len = buf_len - size_t(limit-buf)
						state = 0
						break
					}
					break
				case -2:

				case -1:
					void(mblen(nil, 0))
				case 1:

					/* we need to determine if the enclosure is
					 * 'real' or is it escaped */

					switch state {
					case 1:
						bptr++
						state = 0
						break
					case 2:
						if (*bptr) != enclosure {

							/* real enclosure */

							memcpy(tptr, hunk_begin, bptr-hunk_begin-1)
							tptr += bptr - hunk_begin - 1
							hunk_begin = bptr
							goto quit_loop_2
						}
						memcpy(tptr, hunk_begin, bptr-hunk_begin)
						tptr += bptr - hunk_begin
						bptr++
						hunk_begin = bptr
						state = 0
						break
					default:
						if (*bptr) == enclosure {
							state = 2
						} else if escape_char != -1 && (*bptr) == escape_char {
							state = 1
						}
						bptr++
						break
					}
					break
				default:
					switch state {
					case 2:

						/* real enclosure */

						memcpy(tptr, hunk_begin, bptr-hunk_begin-1)
						tptr += bptr - hunk_begin - 1
						hunk_begin = bptr
						goto quit_loop_2
					case 1:
						bptr += inc_len
						memcpy(tptr, hunk_begin, bptr-hunk_begin)
						tptr += bptr - hunk_begin
						hunk_begin = bptr
						state = 0
						break
					default:
						bptr += inc_len
						break
					}
					break
				}
				if bptr < limit {
					if (*bptr) == '0' {
						inc_len = 1
					} else {
						inc_len = mblen(bptr, limit-bptr)
					}
				} else {
					inc_len = 0
				}
			}
		quit_loop_2:

			/* look up for a delimiter */

			for {
				switch inc_len {
				case 0:
					goto quit_loop_3
				case -2:

				case -1:
					inc_len = 1
					void(mblen(nil, 0))
				case 1:
					if (*bptr) == delimiter {
						goto quit_loop_3
					}
					break
				default:
					break
				}
				bptr += inc_len
				if bptr < limit {
					if (*bptr) == '0' {
						inc_len = 1
					} else {
						inc_len = mblen(bptr, limit-bptr)
					}
				} else {
					inc_len = 0
				}
			}
		quit_loop_3:
			memcpy(tptr, hunk_begin, bptr-hunk_begin)
			tptr += bptr - hunk_begin
			bptr += inc_len
			comp_end = tptr
		} else {

			/* 2B. Handle non-enclosure field */

			hunk_begin = bptr
			for {
				switch inc_len {
				case 0:
					goto quit_loop_4
				case -2:

				case -1:
					inc_len = 1
					void(mblen(nil, 0))
				case 1:
					if (*bptr) == delimiter {
						goto quit_loop_4
					}
					break
				default:
					break
				}
				bptr += inc_len
				if bptr < limit {
					if (*bptr) == '0' {
						inc_len = 1
					} else {
						inc_len = mblen(bptr, limit-bptr)
					}
				} else {
					inc_len = 0
				}
			}
		quit_loop_4:
			memcpy(tptr, hunk_begin, bptr-hunk_begin)
			tptr += bptr - hunk_begin
			comp_end = (*byte)(PhpFgetcsvLookupTrailingSpaces(temp, tptr-temp, delimiter))
			if (*bptr) == delimiter {
				bptr++
			}
		}

		/* 3. Now pass our field back to php */

		*comp_end = '0'
		zend.AddNextIndexStringl(return_value, temp, comp_end-temp)
		if inc_len <= 0 {
			break
		}
	}
out:
	zend._efree(temp)
	if stream != nil {
		zend._efree(buf)
	}
}

/* }}} */

/* {{{ proto string realpath(string path)
   Return the resolved path */

func ZifRealpath(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var filename *byte
	var filename_len int
	var resolved_path_buff []byte
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

			if zend.ZendParseArgPath(_arg, &filename, &filename_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_PATH
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
	if zend.TsrmRealpath(filename, resolved_path_buff) != nil {
		if core.PhpCheckOpenBasedir(resolved_path_buff) != 0 {
			return_value.u1.type_info = 2
			return
		}
		var _s *byte = resolved_path_buff
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = zend.ZendStringInit(_s, strlen(_s), 0)
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
		return
	} else {
		return_value.u1.type_info = 2
		return
	}
}

/* }}} */

/* See http://www.w3.org/TR/html4/intro/sgmltut.html#h-3.2.2 */

// #define PHP_META_HTML401_CHARS       "-_.:"

/* {{{ php_next_meta_token
   Tokenizes an HTML file for get_meta_tags */

func PhpNextMetaToken(md *PhpMetaTagsData) PhpMetaTagsToken {
	var ch int = 0
	var compliment int
	var buff []byte
	memset(any(buff), 0, 8192+1)
	for md.GetUlc() != 0 || streams._phpStreamEof(md.GetStream()) == 0 && g.Assign(&ch, streams._phpStreamGetc(md.GetStream())) {
		if streams._phpStreamEof(md.GetStream()) != 0 {
			break
		}
		if md.GetUlc() != 0 {
			ch = md.GetLc()
			md.SetUlc(0)
		}
		switch ch {
		case '<':
			return TOK_OPENTAG
			break
		case '>':
			return TOK_CLOSETAG
			break
		case '=':
			return TOK_EQUAL
			break
		case '/':
			return TOK_SLASH
			break
		case '\'':

		case '"':
			compliment = ch
			md.SetTokenLen(0)
			for streams._phpStreamEof(md.GetStream()) == 0 && g.Assign(&ch, streams._phpStreamGetc(md.GetStream())) && ch != compliment && ch != '<' && ch != '>' {
				buff[g.PostInc(&(md.GetTokenLen()))] = ch
				if md.GetTokenLen() == 8192 {
					break
				}
			}
			if ch == '<' || ch == '>' {

				/* Was just an apostrohpe */

				md.SetUlc(1)
				md.SetLc(ch)
			}

			/* We don't need to alloc unless we are in a meta tag */

			if md.GetInMeta() != 0 {
				md.SetTokenData((*byte)(zend._emalloc(md.GetTokenLen() + 1)))
				memcpy(md.GetTokenData(), buff, md.GetTokenLen()+1)
			}
			return TOK_STRING
			break
		case '\n':

		case '\r':

		case '\t':
			break
		case ' ':
			return TOK_SPACE
			break
		default:
			if isalnum(ch) {
				md.SetTokenLen(0)
				buff[g.PostInc(&(md.GetTokenLen()))] = ch
				for streams._phpStreamEof(md.GetStream()) == 0 && g.Assign(&ch, streams._phpStreamGetc(md.GetStream())) && (isalnum(ch) || strchr("-_.:", ch)) {
					buff[g.PostInc(&(md.GetTokenLen()))] = ch
					if md.GetTokenLen() == 8192 {
						break
					}
				}

				/* This is ugly, but we have to replace ungetc */

				if !(isalpha(ch)) && ch != '-' {
					md.SetUlc(1)
					md.SetLc(ch)
				}
				md.SetTokenData((*byte)(zend._emalloc(md.GetTokenLen() + 1)))
				memcpy(md.GetTokenData(), buff, md.GetTokenLen()+1)
				return TOK_ID
			} else {
				return TOK_OTHER
			}
			break
		}
	}
	return TOK_EOF
}

/* }}} */

/* {{{ proto bool fnmatch(string pattern, string filename [, int flags])
   Match filename against pattern */

func ZifFnmatch(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var pattern *byte
	var filename *byte
	var pattern_len int
	var filename_len int
	var flags zend.ZendLong = 0
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

			if zend.ZendParseArgPath(_arg, &pattern, &pattern_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_PATH
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

			if zend.ZendParseArgPath(_arg, &filename, &filename_len, 0) == 0 {
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

			if zend.ZendParseArgLong(_arg, &flags, &_dummy, 0, 0) == 0 {
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
	if filename_len >= 256 {
		core.PhpErrorDocref(nil, 1<<1, "Filename exceeds the maximum allowed length of %d characters", 256)
		return_value.u1.type_info = 2
		return
	}
	if pattern_len >= 256 {
		core.PhpErrorDocref(nil, 1<<1, "Pattern exceeds the maximum allowed length of %d characters", 256)
		return_value.u1.type_info = 2
		return
	}
	if !(fnmatch(pattern, filename, int(flags))) {
		return_value.u1.type_info = 3
	} else {
		return_value.u1.type_info = 2
	}
	return
}

/* }}} */

/* {{{ proto string sys_get_temp_dir()
   Returns directory path used for temporary files */

func ZifSysGetTempDir(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	var _s *byte = (*byte)(core.PhpGetTemporaryDirectory())
	var __z *zend.Zval = return_value
	var __s *zend.ZendString = zend.ZendStringInit(_s, strlen(_s), 0)
	__z.value.str = __s
	__z.u1.type_info = 6 | 1<<0<<8
	return
}

/* }}} */
