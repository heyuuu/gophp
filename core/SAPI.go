// <<generate>>

package core

import (
	g "sik/runtime/grammar"
	"sik/zend"
)

// Source: <main/SAPI.h>

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
   | Author:  Zeev Suraski <zeev@php.net>                                 |
   +----------------------------------------------------------------------+
*/

// #define SAPI_H

// # include "php.h"

// # include "zend.h"

// # include "zend_API.h"

// # include "zend_llist.h"

// # include "zend_operators.h"

// # include < sys / stat . h >

// #define SAPI_OPTION_NO_CHDIR       1

// #define SAPI_POST_BLOCK_SIZE       0x4000

// #define SAPI_API

// @type SapiHeader struct
// @type SapiHeaders struct

type sapi_module_struct = _sapiModule

var sapi_module sapi_module_struct

/* Some values in this structure needs to be filled in before
 * calling sapi_activate(). We WILL change the `char *' entries,
 * so make sure that you allocate a separate buffer for them
 * and that you free them after sapi_deactivate().
 */

// @type SapiRequestInfo struct
// @type sapi_globals_struct struct
type _sapiGlobals = sapi_globals_struct

// #define SG(v) ( sapi_globals . v )

var sapi_globals sapi_globals_struct

/*
 * This is the preferred and maintained API for
 * operating on HTTP headers.
 */

// @type SapiHeaderLine struct
type SapiHeaderOpEnum = int

const (
	SAPI_HEADER_REPLACE = iota
	SAPI_HEADER_ADD
	SAPI_HEADER_DELETE
	SAPI_HEADER_DELETE_ALL
	SAPI_HEADER_SET_STATUS
)

/* Deprecated functions. Use sapi_header_op instead. */

// #define sapi_add_header(a,b,c) sapi_add_header_ex ( ( a ) , ( b ) , ( c ) , 1 )

// @type _sapiModule struct
// @type SapiPostEntry struct

/* header_handler() constants */

// #define SAPI_HEADER_ADD       ( 1 << 0 )

// #define SAPI_HEADER_SENT_SUCCESSFULLY       1

// #define SAPI_HEADER_DO_SEND       2

// #define SAPI_HEADER_SEND_FAILED       3

// #define SAPI_DEFAULT_MIMETYPE       "text/html"

// #define SAPI_DEFAULT_CHARSET       PHP_DEFAULT_CHARSET

// #define SAPI_PHP_VERSION_HEADER       "X-Powered-By: PHP/" PHP_VERSION

// #define SAPI_POST_READER_FUNC(post_reader) void post_reader ( void )

// #define SAPI_POST_HANDLER_FUNC(post_handler) void post_handler ( char * content_type_dup , void * arg )

// #define SAPI_TREAT_DATA_FUNC(treat_data) void treat_data ( int arg , char * str , zval * destArray )

// #define SAPI_INPUT_FILTER_FUNC(input_filter) unsigned int input_filter ( int arg , char * var , char * * val , size_t val_len , size_t * new_val_len )

// #define STANDARD_SAPI_MODULE_PROPERTIES       NULL , NULL , NULL , NULL , 0 , 0 , NULL , NULL , NULL , NULL , NULL , NULL , 0 , NULL , NULL , NULL

// Source: <main/SAPI.c>

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
   | Original design:  Shane Caraveo <shane@caraveo.com>                  |
   | Authors: Andi Gutmans <andi@php.net>                                 |
   |          Zeev Suraski <zeev@php.net>                                 |
   +----------------------------------------------------------------------+
*/

// # include < ctype . h >

// # include < sys / stat . h >

// # include "php.h"

// # include "SAPI.h"

// # include "php_variables.h"

// # include "php_ini.h"

// # include "ext/standard/php_string.h"

// # include "ext/standard/pageinfo.h"

// failed # include "ext/pcre/php_pcre.h"

// # include < sys / time . h >

// # include "rfc1867.h"

// # include "php_content_types.h"

func _typeDtor(zv *zend.Zval) { zend.Free(zv.value.ptr) }
func SapiGlobalsCtor(sapi_globals *sapi_globals_struct) {
	memset(sapi_globals, 0, g.SizeOf("* sapi_globals"))
	zend._zendHashInit(&sapi_globals.known_post_content_types, 8, _typeDtor, 1)
	PhpSetupSapiContentTypes()
}
func SapiGlobalsDtor(sapi_globals *sapi_globals_struct) {
	zend.ZendHashDestroy(&sapi_globals.known_post_content_types)
}

/* True globals (no need for thread safety) */

func SapiStartup(sf *sapi_module_struct) {
	sf.SetIniEntries(nil)
	sapi_module = *sf
	SapiGlobalsCtor(&sapi_globals)
}
func SapiShutdown()                          { SapiGlobalsDtor(&sapi_globals) }
func SapiFreeHeader(sapi_header *SapiHeader) { zend._efree(sapi_header.GetHeader()) }

/* {{{ proto bool header_register_callback(mixed callback)
   call a header function */

func ZifHeaderRegisterCallback(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var callback_func *zend.Zval
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "z", &callback_func) == zend.FAILURE {
		return
	}
	if zend.ZendIsCallable(callback_func, 0, nil) == 0 {
		return_value.u1.type_info = 2
		return
	}
	if sapi_globals.callback_func.u1.v.type_ != 0 {
		zend.ZvalPtrDtor(&(sapi_globals.GetCallbackFunc()))
		sapi_globals.SetFciCache(zend.EmptyFcallInfoCache)
	}
	var _z1 *zend.Zval = &(sapi_globals.GetCallbackFunc())
	var _z2 *zend.Zval = callback_func
	var _gc *zend.ZendRefcounted = _z2.value.counted
	var _t uint32 = _z2.u1.type_info
	_z1.value.counted = _gc
	_z1.u1.type_info = _t
	if (_t & 0xff00) != 0 {
		zend.ZendGcAddref(&_gc.gc)
	}
	return_value.u1.type_info = 3
	return
}

/* }}} */

func SapiRunHeaderCallback(callback *zend.Zval) {
	var error int
	var fci zend.ZendFcallInfo
	var callback_error *byte = nil
	var retval zend.Zval
	if zend.ZendFcallInfoInit(callback, 0, &fci, &(sapi_globals.GetFciCache()), nil, &callback_error) == zend.SUCCESS {
		fci.retval = &retval
		error = zend.ZendCallFunction(&fci, &(sapi_globals.GetFciCache()))
		if error == zend.FAILURE {
			goto callback_failed
		} else {
			zend.ZvalPtrDtor(&retval)
		}
	} else {
	callback_failed:
		PhpErrorDocref(nil, 1<<1, "Could not call the sapi_header_callback")
	}
	if callback_error != nil {
		zend._efree(callback_error)
	}
}
func SapiHandlePost(arg any) {
	if sapi_globals.GetRequestInfo().GetPostEntry() != nil && sapi_globals.GetRequestInfo().GetContentTypeDup() != nil {
		sapi_globals.GetRequestInfo().GetPostEntry().GetPostHandler()(sapi_globals.GetRequestInfo().GetContentTypeDup(), arg)
		zend._efree(sapi_globals.GetRequestInfo().GetContentTypeDup())
		sapi_globals.GetRequestInfo().SetContentTypeDup(nil)
	}
}
func SapiReadPostData() {
	var post_entry *SapiPostEntry
	var content_type_length uint32 = uint32(strlen(sapi_globals.GetRequestInfo().GetContentType()))
	var content_type *byte = zend._estrndup(sapi_globals.GetRequestInfo().GetContentType(), content_type_length)
	var p *byte
	var oldchar byte = 0
	var post_reader_func func() = nil

	/* dedicated implementation for increased performance:
	 * - Make the content type lowercase
	 * - Trim descriptive data, stay with the content-type only
	 */

	for p = content_type; p < content_type+content_type_length; p++ {
		switch *p {
		case ';':

		case ',':

		case ' ':
			content_type_length = p - content_type
			oldchar = *p
			*p = 0
			break
		default:
			*p = tolower(*p)
			break
		}
	}

	/* now try to find an appropriate POST content handler */

	if g.Assign(&post_entry, zend.ZendHashStrFindPtr(&(sapi_globals.GetKnownPostContentTypes()), content_type, content_type_length)) != nil {

		/* found one, register it for use */

		sapi_globals.GetRequestInfo().SetPostEntry(post_entry)
		post_reader_func = post_entry.GetPostReader()
	} else {

		/* fallback */

		sapi_globals.GetRequestInfo().SetPostEntry(nil)
		if sapi_module.GetDefaultPostReader() == nil {

			/* no default reader ? */

			sapi_globals.GetRequestInfo().SetContentTypeDup(nil)
			sapi_module.GetSapiError()(1<<1, "Unsupported content type:  '%s'", content_type)
			return
		}
	}
	if oldchar {
		*(p - 1) = oldchar
	}
	sapi_globals.GetRequestInfo().SetContentTypeDup(content_type)
	if post_reader_func != nil {
		post_reader_func()
	}
	if sapi_module.GetDefaultPostReader() != nil {
		sapi_module.GetDefaultPostReader()()
	}
}
func SapiReadPostBlock(buffer *byte, buflen int) int {
	var read_bytes int
	if sapi_module.GetReadPost() == nil {
		return 0
	}
	read_bytes = sapi_module.GetReadPost()(buffer, buflen)
	if read_bytes > 0 {

		/* gogo */

		sapi_globals.SetReadPostBytes(sapi_globals.GetReadPostBytes() + read_bytes)

		/* gogo */

	}
	if read_bytes < buflen {

		/* done */

		sapi_globals.SetPostRead(1)

		/* done */

	}
	return read_bytes
}
func SapiReadStandardFormData() {
	if sapi_globals.GetPostMaxSize() > 0 && sapi_globals.GetRequestInfo().GetContentLength() > sapi_globals.GetPostMaxSize() {
		PhpErrorDocref(nil, 1<<1, "POST Content-Length of "+"%"+"lld"+" bytes exceeds the limit of "+"%"+"lld"+" bytes", sapi_globals.GetRequestInfo().GetContentLength(), sapi_globals.GetPostMaxSize())
		return
	}
	sapi_globals.GetRequestInfo().SetRequestBody(_phpStreamTempCreateEx(0x0, 0x4000, CoreGlobals.GetUploadTmpDir()))
	if sapi_module.GetReadPost() != nil {
		var read_bytes int
		for {
			var buffer []byte
			read_bytes = SapiReadPostBlock(buffer, 0x4000)
			if read_bytes > 0 {
				if _phpStreamWrite(sapi_globals.GetRequestInfo().GetRequestBody(), buffer, read_bytes) != read_bytes {

					/* if parts of the stream can't be written, purge it completely */

					_phpStreamTruncateSetSize(sapi_globals.GetRequestInfo().GetRequestBody(), 0)
					PhpErrorDocref(nil, 1<<1, "POST data can't be buffered; all data discarded")
					break
				}
			}
			if sapi_globals.GetPostMaxSize() > 0 && sapi_globals.GetReadPostBytes() > sapi_globals.GetPostMaxSize() {
				PhpErrorDocref(nil, 1<<1, "Actual POST length does not match Content-Length, and exceeds "+"%"+"lld"+" bytes", sapi_globals.GetPostMaxSize())
				break
			}
			if read_bytes < 0x4000 {

				/* done */

				break

				/* done */

			}
		}
		_phpStreamSeek(sapi_globals.GetRequestInfo().GetRequestBody(), 0, SEEK_SET)
	}
}
func GetDefaultContentType(prefix_len uint32, len_ *uint32) *byte {
	var mimetype *byte
	var charset *byte
	var content_type *byte
	var mimetype_len uint32
	var charset_len uint32
	if sapi_globals.GetDefaultMimetype() != nil {
		mimetype = sapi_globals.GetDefaultMimetype()
		mimetype_len = uint32(strlen(sapi_globals.GetDefaultMimetype()))
	} else {
		mimetype = "text/html"
		mimetype_len = g.SizeOf("SAPI_DEFAULT_MIMETYPE") - 1
	}
	if sapi_globals.GetDefaultCharset() != nil {
		charset = sapi_globals.GetDefaultCharset()
		charset_len = uint32(strlen(sapi_globals.GetDefaultCharset()))
	} else {
		charset = "UTF-8"
		charset_len = g.SizeOf("SAPI_DEFAULT_CHARSET") - 1
	}
	if (*charset) && strncasecmp(mimetype, "text/", 5) == 0 {
		var p *byte
		*len_ = prefix_len + mimetype_len + g.SizeOf("\"; charset=\"") - 1 + charset_len
		content_type = (*byte)(zend._emalloc((*len_) + 1))
		p = content_type + prefix_len
		memcpy(p, mimetype, mimetype_len)
		p += mimetype_len
		memcpy(p, "; charset=", g.SizeOf("\"; charset=\"")-1)
		p += g.SizeOf("\"; charset=\"") - 1
		memcpy(p, charset, charset_len+1)
	} else {
		*len_ = prefix_len + mimetype_len
		content_type = (*byte)(zend._emalloc((*len_) + 1))
		memcpy(content_type+prefix_len, mimetype, mimetype_len+1)
	}
	return content_type
}
func SapiGetDefaultContentType() *byte {
	var len_ uint32
	return GetDefaultContentType(0, &len_)
}
func SapiGetDefaultContentTypeHeader(default_header *SapiHeader) {
	var len_ uint32
	default_header.SetHeader(GetDefaultContentType(g.SizeOf("\"Content-type: \"")-1, &len_))
	default_header.SetHeaderLen(len_)
	memcpy(default_header.GetHeader(), "Content-type: ", g.SizeOf("\"Content-type: \"")-1)
}

/*
 * Add charset on content-type header if the MIME type starts with
 * "text/", the default_charset directive is not empty and
 * there is not already a charset option in there.
 *
 * If "mimetype" is non-NULL, it should point to a pointer allocated
 * with emalloc().  If a charset is added, the string will be
 * re-allocated and the new length is returned.  If mimetype is
 * unchanged, 0 is returned.
 *
 */

func SapiApplyDefaultCharset(mimetype **byte, len_ int) int {
	var charset *byte
	var newtype *byte
	var newlen int
	if sapi_globals.GetDefaultCharset() != nil {
		charset = sapi_globals.GetDefaultCharset()
	} else {
		charset = "UTF-8"
	}
	if (*mimetype) != nil {
		if (*charset) && strncmp(*mimetype, "text/", 5) == 0 && strstr(*mimetype, "charset=") == nil {
			newlen = len_ + (g.SizeOf("\";charset=\"") - 1) + strlen(charset)
			newtype = zend._emalloc(newlen + 1)
			var php_str_len int
			if len_ >= newlen+1 {
				php_str_len = newlen + 1 - 1
			} else {
				php_str_len = len_
			}
			memcpy(newtype, *mimetype, php_str_len)
			newtype[php_str_len] = '0'
			strlcat(newtype, ";charset=", newlen+1)
			strlcat(newtype, charset, newlen+1)
			zend._efree(*mimetype)
			*mimetype = newtype
			return newlen
		}
	}
	return 0
}
func SapiActivateHeadersOnly() {
	if sapi_globals.GetRequestInfo().GetHeadersRead() == 1 {
		return
	}
	sapi_globals.GetRequestInfo().SetHeadersRead(1)
	zend.ZendLlistInit(&(sapi_globals.GetSapiHeaders()).headers, g.SizeOf("sapi_header_struct"), (func(any))(SapiFreeHeader), 0)
	sapi_globals.GetSapiHeaders().SetSendDefaultContentType(1)

	/* SG(sapi_headers).http_response_code = 200; */

	sapi_globals.GetSapiHeaders().SetHttpStatusLine(nil)
	sapi_globals.GetSapiHeaders().SetMimetype(nil)
	sapi_globals.SetReadPostBytes(0)
	sapi_globals.GetRequestInfo().SetRequestBody(nil)
	sapi_globals.GetRequestInfo().SetCurrentUser(nil)
	sapi_globals.GetRequestInfo().SetCurrentUserLength(0)
	sapi_globals.GetRequestInfo().SetNoHeaders(0)
	sapi_globals.GetRequestInfo().SetPostEntry(nil)
	sapi_globals.SetGlobalRequestTime(0)

	/*
	 * It's possible to override this general case in the activate() callback,
	 * if necessary.
	 */

	if sapi_globals.GetRequestInfo().GetRequestMethod() != nil && !(strcmp(sapi_globals.GetRequestInfo().GetRequestMethod(), "HEAD")) {
		sapi_globals.GetRequestInfo().SetHeadersOnly(1)
	} else {
		sapi_globals.GetRequestInfo().SetHeadersOnly(0)
	}
	if sapi_globals.GetServerContext() {
		sapi_globals.GetRequestInfo().SetCookieData(sapi_module.GetReadCookies()())
		if sapi_module.GetActivate() != nil {
			sapi_module.GetActivate()()
		}
	}
	if sapi_module.GetInputFilterInit() != nil {
		sapi_module.GetInputFilterInit()()
	}
}

/*
 * Called from php_request_startup() for every request.
 */

func SapiActivate() {
	zend.ZendLlistInit(&(sapi_globals.GetSapiHeaders()).headers, g.SizeOf("sapi_header_struct"), (func(any))(SapiFreeHeader), 0)
	sapi_globals.GetSapiHeaders().SetSendDefaultContentType(1)

	/*
	   SG(sapi_headers).http_response_code = 200;
	*/

	sapi_globals.GetSapiHeaders().SetHttpStatusLine(nil)
	sapi_globals.GetSapiHeaders().SetMimetype(nil)
	sapi_globals.SetHeadersSent(0)
	&(sapi_globals.GetCallbackFunc()).u1.type_info = 0
	sapi_globals.SetReadPostBytes(0)
	sapi_globals.GetRequestInfo().SetRequestBody(nil)
	sapi_globals.GetRequestInfo().SetCurrentUser(nil)
	sapi_globals.GetRequestInfo().SetCurrentUserLength(0)
	sapi_globals.GetRequestInfo().SetNoHeaders(0)
	sapi_globals.GetRequestInfo().SetPostEntry(nil)
	sapi_globals.GetRequestInfo().SetProtoNum(1000)
	sapi_globals.SetGlobalRequestTime(0)
	sapi_globals.SetPostRead(0)

	/* It's possible to override this general case in the activate() callback, if necessary. */

	if sapi_globals.GetRequestInfo().GetRequestMethod() != nil && !(strcmp(sapi_globals.GetRequestInfo().GetRequestMethod(), "HEAD")) {
		sapi_globals.GetRequestInfo().SetHeadersOnly(1)
	} else {
		sapi_globals.GetRequestInfo().SetHeadersOnly(0)
	}
	sapi_globals.SetRfc1867UploadedFiles(nil)

	/* Handle request method */

	if sapi_globals.GetServerContext() {
		if CoreGlobals.GetEnablePostDataReading() != 0 && sapi_globals.GetRequestInfo().GetContentType() != nil && sapi_globals.GetRequestInfo().GetRequestMethod() != nil && !(strcmp(sapi_globals.GetRequestInfo().GetRequestMethod(), "POST")) {

			/* HTTP POST may contain form data to be processed into variables
			 * depending on given content type */

			SapiReadPostData()

			/* HTTP POST may contain form data to be processed into variables
			 * depending on given content type */

		} else {
			sapi_globals.GetRequestInfo().SetContentTypeDup(nil)
		}

		/* Cookies */

		sapi_globals.GetRequestInfo().SetCookieData(sapi_module.GetReadCookies()())

		/* Cookies */

	}
	if sapi_module.GetActivate() != nil {
		sapi_module.GetActivate()()
	}
	if sapi_module.GetInputFilterInit() != nil {
		sapi_module.GetInputFilterInit()()
	}
}
func SapiSendHeadersFree() {
	if sapi_globals.GetSapiHeaders().GetHttpStatusLine() != nil {
		zend._efree(sapi_globals.GetSapiHeaders().GetHttpStatusLine())
		sapi_globals.GetSapiHeaders().SetHttpStatusLine(nil)
	}
}
func SapiDeactivate() {
	zend.ZendLlistDestroy(&(sapi_globals.GetSapiHeaders()).headers)
	if sapi_globals.GetRequestInfo().GetRequestBody() != nil {
		sapi_globals.GetRequestInfo().SetRequestBody(nil)
	} else if sapi_globals.GetServerContext() {
		if sapi_globals.GetPostRead() == 0 {

			/* make sure we've consumed all request input data */

			var dummy []byte
			var read_bytes int
			for {
				read_bytes = SapiReadPostBlock(dummy, 0x4000)
				if 0x4000 != read_bytes {
					break
				}
			}
		}
	}
	if sapi_globals.GetRequestInfo().GetAuthUser() != nil {
		zend._efree(sapi_globals.GetRequestInfo().GetAuthUser())
	}
	if sapi_globals.GetRequestInfo().GetAuthPassword() != nil {
		zend._efree(sapi_globals.GetRequestInfo().GetAuthPassword())
	}
	if sapi_globals.GetRequestInfo().GetAuthDigest() != nil {
		zend._efree(sapi_globals.GetRequestInfo().GetAuthDigest())
	}
	if sapi_globals.GetRequestInfo().GetContentTypeDup() != nil {
		zend._efree(sapi_globals.GetRequestInfo().GetContentTypeDup())
	}
	if sapi_globals.GetRequestInfo().GetCurrentUser() != nil {
		zend._efree(sapi_globals.GetRequestInfo().GetCurrentUser())
	}
	if sapi_module.GetDeactivate() != nil {
		sapi_module.GetDeactivate()()
	}
	if sapi_globals.GetRfc1867UploadedFiles() != nil {
		DestroyUploadedFilesHash()
	}
	if sapi_globals.GetSapiHeaders().GetMimetype() != nil {
		zend._efree(sapi_globals.GetSapiHeaders().GetMimetype())
		sapi_globals.GetSapiHeaders().SetMimetype(nil)
	}
	SapiSendHeadersFree()
	sapi_globals.SetSapiStarted(0)
	sapi_globals.SetHeadersSent(0)
	sapi_globals.GetRequestInfo().SetHeadersRead(0)
	sapi_globals.SetGlobalRequestTime(0)
}
func SapiInitializeEmptyRequest() {
	sapi_globals.SetServerContext(nil)
	sapi_globals.GetRequestInfo().SetRequestMethod(nil)
	sapi_globals.GetRequestInfo().SetAuthPassword(nil)
	sapi_globals.GetRequestInfo().SetAuthUser(sapi_globals.GetRequestInfo().GetAuthPassword())
	sapi_globals.GetRequestInfo().SetAuthDigest(sapi_globals.GetRequestInfo().GetAuthUser())
	sapi_globals.GetRequestInfo().SetContentTypeDup(nil)
}
func SapiExtractResponseCode(header_line *byte) int {
	var code int = 200
	var ptr *byte
	for ptr = header_line; *ptr; ptr++ {
		if (*ptr) == ' ' && (*(ptr + 1)) != ' ' {
			code = atoi(ptr + 1)
			break
		}
	}
	return code
}
func SapiUpdateResponseCode(ncode int) {
	/* if the status code did not change, we do not want
	   to change the status line, and no need to change the code */

	if sapi_globals.GetSapiHeaders().GetHttpResponseCode() == ncode {
		return
	}
	if sapi_globals.GetSapiHeaders().GetHttpStatusLine() != nil {
		zend._efree(sapi_globals.GetSapiHeaders().GetHttpStatusLine())
		sapi_globals.GetSapiHeaders().SetHttpStatusLine(nil)
	}
	sapi_globals.GetSapiHeaders().SetHttpResponseCode(ncode)
}

/*
 * since zend_llist_del_element only remove one matched item once,
 * we should remove them by ourself
 */

func SapiRemoveHeader(l *zend.ZendLlist, name *byte, len_ int) {
	var header *SapiHeader
	var next *zend.ZendLlistElement
	var current *zend.ZendLlistElement = l.head
	for current != nil {
		header = (*SapiHeader)(current.data)
		next = current.next
		if header.GetHeaderLen() > len_ && header.GetHeader()[len_] == ':' && !(strncasecmp(header.GetHeader(), name, len_)) {
			if current.prev != nil {
				current.prev.next = next
			} else {
				l.head = next
			}
			if next != nil {
				next.prev = current.prev
			} else {
				l.tail = current.prev
			}
			SapiFreeHeader(header)
			zend._efree(current)
			l.count--
		}
		current = next
	}
}
func SapiAddHeaderEx(header_line string, header_line_len int, duplicate zend.ZendBool, replace zend.ZendBool) int {
	var ctr SapiHeaderLine = SapiHeaderLine{0}
	var r int
	ctr.SetLine(header_line)
	ctr.SetLineLen(header_line_len)
	r = SapiHeaderOp(g.Cond(replace != 0, SAPI_HEADER_REPLACE, 1<<0), &ctr)
	if duplicate == 0 {
		zend._efree(header_line)
	}
	return r
}
func SapiHeaderAddOp(op SapiHeaderOpEnum, sapi_header *SapiHeader) {
	if sapi_module.GetHeaderHandler() == nil || (1<<0&sapi_module.GetHeaderHandler()(sapi_header, op, &(sapi_globals.GetSapiHeaders()))) != 0 {
		if op == SAPI_HEADER_REPLACE {
			var colon_offset *byte = strchr(sapi_header.GetHeader(), ':')
			if colon_offset != nil {
				var sav byte = *colon_offset
				*colon_offset = 0
				SapiRemoveHeader(&(sapi_globals.GetSapiHeaders()).headers, sapi_header.GetHeader(), strlen(sapi_header.GetHeader()))
				*colon_offset = sav
			}
		}
		zend.ZendLlistAddElement(&(sapi_globals.GetSapiHeaders()).headers, any(sapi_header))
	} else {
		SapiFreeHeader(sapi_header)
	}
}
func SapiHeaderOp(op SapiHeaderOpEnum, arg any) int {
	var sapi_header SapiHeader
	var colon_offset *byte
	var header_line *byte
	var header_line_len int
	var http_response_code int
	if sapi_globals.GetHeadersSent() != 0 && sapi_globals.GetRequestInfo().GetNoHeaders() == 0 {
		var output_start_filename *byte = PhpOutputGetStartFilename()
		var output_start_lineno int = PhpOutputGetStartLineno()
		if output_start_filename != nil {
			sapi_module.GetSapiError()(1<<1, "Cannot modify header information - headers already sent by (output started at %s:%d)", output_start_filename, output_start_lineno)
		} else {
			sapi_module.GetSapiError()(1<<1, "Cannot modify header information - headers already sent")
		}
		return zend.FAILURE
	}
	switch op {
	case SAPI_HEADER_SET_STATUS:
		SapiUpdateResponseCode(int(zend.ZendIntptrT(arg)))
		return zend.SUCCESS
	case 1 << 0:

	case SAPI_HEADER_REPLACE:

	case SAPI_HEADER_DELETE:
		var p *SapiHeaderLine = arg
		if p.GetLine() == nil || p.GetLineLen() == 0 {
			return zend.FAILURE
		}
		header_line = p.GetLine()
		header_line_len = p.GetLineLen()
		http_response_code = p.GetResponseCode()
		break
	case SAPI_HEADER_DELETE_ALL:
		if sapi_module.GetHeaderHandler() != nil {
			sapi_module.GetHeaderHandler()(&sapi_header, op, &(sapi_globals.GetSapiHeaders()))
		}
		zend.ZendLlistClean(&(sapi_globals.GetSapiHeaders()).headers)
		return zend.SUCCESS
	default:
		return zend.FAILURE
	}
	header_line = zend._estrndup(header_line, header_line_len)

	/* cut off trailing spaces, linefeeds and carriage-returns */

	if header_line_len != 0 && isspace(header_line[header_line_len-1]) {
		for {
			header_line_len--
			if !(header_line_len != 0 && isspace(header_line[header_line_len-1])) {
				break
			}
		}
		header_line[header_line_len] = '0'
	}
	if op == SAPI_HEADER_DELETE {
		if strchr(header_line, ':') {
			zend._efree(header_line)
			sapi_module.GetSapiError()(1<<1, "Header to delete may not contain colon.")
			return zend.FAILURE
		}
		if sapi_module.GetHeaderHandler() != nil {
			sapi_header.SetHeader(header_line)
			sapi_header.SetHeaderLen(header_line_len)
			sapi_module.GetHeaderHandler()(&sapi_header, op, &(sapi_globals.GetSapiHeaders()))
		}
		SapiRemoveHeader(&(sapi_globals.GetSapiHeaders()).headers, header_line, header_line_len)
		zend._efree(header_line)
		return zend.SUCCESS
	} else {

		/* new line/NUL character safety check */

		var i uint32
		for i = 0; i < header_line_len; i++ {

			/* RFC 7230 ch. 3.2.4 deprecates folding support */

			if header_line[i] == '\n' || header_line[i] == '\r' {
				zend._efree(header_line)
				sapi_module.GetSapiError()(1<<1, "Header may not contain "+"more than a single header, new line detected")
				return zend.FAILURE
			}
			if header_line[i] == '0' {
				zend._efree(header_line)
				sapi_module.GetSapiError()(1<<1, "Header may not contain NUL bytes")
				return zend.FAILURE
			}
		}
	}
	sapi_header.SetHeader(header_line)
	sapi_header.SetHeaderLen(header_line_len)

	/* Check the header for a few cases that we have special support for in SAPI */

	if header_line_len >= 5 && !(strncasecmp(header_line, "HTTP/", 5)) {

		/* filter out the response code */

		SapiUpdateResponseCode(SapiExtractResponseCode(header_line))

		/* sapi_update_response_code doesn't free the status line if the code didn't change */

		if sapi_globals.GetSapiHeaders().GetHttpStatusLine() != nil {
			zend._efree(sapi_globals.GetSapiHeaders().GetHttpStatusLine())
		}
		sapi_globals.GetSapiHeaders().SetHttpStatusLine(header_line)
		return zend.SUCCESS
	} else {
		colon_offset = strchr(header_line, ':')
		if colon_offset != nil {
			*colon_offset = 0
			if !(strcasecmp(header_line, "Content-Type")) {
				var ptr *byte = colon_offset + 1
				var mimetype *byte = nil
				var newheader *byte
				var len_ int = header_line_len - (ptr - header_line)
				var newlen int
				for (*ptr) == ' ' {
					ptr++
					len_--
				}

				/* Disable possible output compression for images */

				if !(strncmp(ptr, "image/", g.SizeOf("\"image/\"")-1)) {
					var key *zend.ZendString = zend.ZendStringInit("zlib.output_compression", g.SizeOf("\"zlib.output_compression\"")-1, 0)
					zend.ZendAlterIniEntryChars(key, "0", g.SizeOf("\"0\"")-1, 1<<0, 1<<4)
					zend.ZendStringReleaseEx(key, 0)
				}
				mimetype = zend._estrdup(ptr)
				newlen = SapiApplyDefaultCharset(&mimetype, len_)
				if sapi_globals.GetSapiHeaders().GetMimetype() == nil {
					sapi_globals.GetSapiHeaders().SetMimetype(zend._estrdup(mimetype))
				}
				if newlen != 0 {
					newlen += g.SizeOf("\"Content-type: \"")
					newheader = zend._emalloc(newlen)
					var php_str_len int
					if g.SizeOf("\"Content-type: \"")-1 >= newlen {
						php_str_len = newlen - 1
					} else {
						php_str_len = g.SizeOf("\"Content-type: \"") - 1
					}
					memcpy(newheader, "Content-type: ", php_str_len)
					newheader[php_str_len] = '0'
					strlcat(newheader, mimetype, newlen)
					sapi_header.SetHeader(newheader)
					sapi_header.SetHeaderLen(uint32(newlen - 1))
					zend._efree(header_line)
				}
				zend._efree(mimetype)
				sapi_globals.GetSapiHeaders().SetSendDefaultContentType(0)
			} else if !(strcasecmp(header_line, "Content-Length")) {

				/* Script is setting Content-length. The script cannot reasonably
				 * know the size of the message body after compression, so it's best
				 * do disable compression altogether. This contributes to making scripts
				 * portable between setups that have and don't have zlib compression
				 * enabled globally. See req #44164 */

				var key *zend.ZendString = zend.ZendStringInit("zlib.output_compression", g.SizeOf("\"zlib.output_compression\"")-1, 0)
				zend.ZendAlterIniEntryChars(key, "0", g.SizeOf("\"0\"")-1, 1<<0, 1<<4)
				zend.ZendStringReleaseEx(key, 0)
			} else if !(strcasecmp(header_line, "Location")) {
				if (sapi_globals.GetSapiHeaders().GetHttpResponseCode() < 300 || sapi_globals.GetSapiHeaders().GetHttpResponseCode() > 399) && sapi_globals.GetSapiHeaders().GetHttpResponseCode() != 201 {

					/* Return a Found Redirect if one is not already specified */

					if http_response_code != 0 {
						SapiUpdateResponseCode(http_response_code)
					} else if sapi_globals.GetRequestInfo().GetProtoNum() > 1000 && sapi_globals.GetRequestInfo().GetRequestMethod() != nil && strcmp(sapi_globals.GetRequestInfo().GetRequestMethod(), "HEAD") && strcmp(sapi_globals.GetRequestInfo().GetRequestMethod(), "GET") {
						SapiUpdateResponseCode(303)
					} else {
						SapiUpdateResponseCode(302)
					}

					/* Return a Found Redirect if one is not already specified */

				}
			} else if !(strcasecmp(header_line, "WWW-Authenticate")) {
				SapiUpdateResponseCode(401)
			}
			if sapi_header.GetHeader() == header_line {
				*colon_offset = ':'
			}
		}
	}
	if http_response_code != 0 {
		SapiUpdateResponseCode(http_response_code)
	}
	SapiHeaderAddOp(op, &sapi_header)
	return zend.SUCCESS
}
func SapiSendHeaders() int {
	var retval int
	var ret int = zend.FAILURE
	if sapi_globals.GetHeadersSent() != 0 || sapi_globals.GetRequestInfo().GetNoHeaders() != 0 {
		return zend.SUCCESS
	}

	/* Success-oriented.  We set headers_sent to 1 here to avoid an infinite loop
	 * in case of an error situation.
	 */

	if sapi_globals.GetSapiHeaders().GetSendDefaultContentType() != 0 && sapi_module.GetSendHeaders() != nil {
		var len_ uint32 = 0
		var default_mimetype *byte = GetDefaultContentType(0, &len_)
		if default_mimetype != nil && len_ != 0 {
			var default_header SapiHeader
			sapi_globals.GetSapiHeaders().SetMimetype(default_mimetype)
			default_header.SetHeaderLen(g.SizeOf("\"Content-type: \"") - 1 + len_)
			default_header.SetHeader(zend._emalloc(default_header.GetHeaderLen() + 1))
			memcpy(default_header.GetHeader(), "Content-type: ", g.SizeOf("\"Content-type: \"")-1)
			memcpy(default_header.GetHeader()+g.SizeOf("\"Content-type: \"")-1, sapi_globals.GetSapiHeaders().GetMimetype(), len_+1)
			SapiHeaderAddOp(1<<0, &default_header)
		} else {
			zend._efree(default_mimetype)
		}
		sapi_globals.GetSapiHeaders().SetSendDefaultContentType(0)
	}
	if sapi_globals.callback_func.u1.v.type_ != 0 {
		var cb zend.Zval
		var _z1 *zend.Zval = &cb
		var _z2 *zend.Zval = &(sapi_globals.GetCallbackFunc())
		var _gc *zend.ZendRefcounted = _z2.value.counted
		var _t uint32 = _z2.u1.type_info
		_z1.value.counted = _gc
		_z1.u1.type_info = _t
		&(sapi_globals.GetCallbackFunc()).u1.type_info = 0
		SapiRunHeaderCallback(&cb)
		zend.ZvalPtrDtor(&cb)
	}
	sapi_globals.SetHeadersSent(1)
	if sapi_module.GetSendHeaders() != nil {
		retval = sapi_module.GetSendHeaders()(&(sapi_globals.GetSapiHeaders()))
	} else {
		retval = 2
	}
	switch retval {
	case 1:
		ret = zend.SUCCESS
		break
	case 2:
		var http_status_line SapiHeader
		var buf []byte
		if sapi_globals.GetSapiHeaders().GetHttpStatusLine() != nil {
			http_status_line.SetHeader(sapi_globals.GetSapiHeaders().GetHttpStatusLine())
			http_status_line.SetHeaderLen(uint32(strlen(sapi_globals.GetSapiHeaders().GetHttpStatusLine())))
		} else {
			http_status_line.SetHeader(buf)
			http_status_line.SetHeaderLen(ApPhpSlprintf(buf, g.SizeOf("buf"), "HTTP/1.0 %d X", sapi_globals.GetSapiHeaders().GetHttpResponseCode()))
		}
		sapi_module.GetSendHeader()(&http_status_line, sapi_globals.GetServerContext())
		zend.ZendLlistApplyWithArgument(&(sapi_globals.GetSapiHeaders()).headers, zend.LlistApplyWithArgFuncT(sapi_module.GetSendHeader()), sapi_globals.GetServerContext())
		if sapi_globals.GetSapiHeaders().GetSendDefaultContentType() != 0 {
			var default_header SapiHeader
			SapiGetDefaultContentTypeHeader(&default_header)
			sapi_module.GetSendHeader()(&default_header, sapi_globals.GetServerContext())
			SapiFreeHeader(&default_header)
		}
		sapi_module.GetSendHeader()(nil, sapi_globals.GetServerContext())
		ret = zend.SUCCESS
		break
	case 3:
		sapi_globals.SetHeadersSent(0)
		ret = zend.FAILURE
		break
	}
	SapiSendHeadersFree()
	return ret
}
func SapiRegisterPostEntries(post_entries *SapiPostEntry) int {
	var p *SapiPostEntry = post_entries
	for p.GetContentType() != nil {
		if SapiRegisterPostEntry(p) == zend.FAILURE {
			return zend.FAILURE
		}
		p++
	}
	return zend.SUCCESS
}
func SapiRegisterPostEntry(post_entry *SapiPostEntry) int {
	var ret int
	var key *zend.ZendString
	if sapi_globals.GetSapiStarted() != 0 && zend.EG.current_execute_data != nil {
		return zend.FAILURE
	}
	key = zend.ZendStringInit(post_entry.GetContentType(), post_entry.GetContentTypeLen(), 1)

	if zend.ZendHashAddMem(&(sapi_globals.GetKnownPostContentTypes()), key, any(post_entry), g.SizeOf("sapi_post_entry")) {
		ret = zend.SUCCESS
	} else {
		ret = zend.FAILURE
	}
	zend.ZendStringReleaseEx(key, 1)
	return ret
}
func SapiUnregisterPostEntry(post_entry *SapiPostEntry) {
	if sapi_globals.GetSapiStarted() != 0 && zend.EG.current_execute_data != nil {
		return
	}
	zend.ZendHashStrDel(&(sapi_globals.GetKnownPostContentTypes()), post_entry.GetContentType(), post_entry.GetContentTypeLen())
}
func SapiRegisterDefaultPostReader(default_post_reader func()) int {
	if sapi_globals.GetSapiStarted() != 0 && zend.EG.current_execute_data != nil {
		return zend.FAILURE
	}
	sapi_module.SetDefaultPostReader(default_post_reader)
	return zend.SUCCESS
}
func SapiRegisterTreatData(treat_data func(arg int, str *byte, destArray *zend.Zval)) int {
	if sapi_globals.GetSapiStarted() != 0 && zend.EG.current_execute_data != nil {
		return zend.FAILURE
	}
	sapi_module.SetTreatData(treat_data)
	return zend.SUCCESS
}
func SapiRegisterInputFilter(input_filter func(arg int, var_ *byte, val **byte, val_len int, new_val_len *int) uint, input_filter_init func() uint) int {
	if sapi_globals.GetSapiStarted() != 0 && zend.EG.current_execute_data != nil {
		return zend.FAILURE
	}
	sapi_module.SetInputFilter(input_filter)
	sapi_module.SetInputFilterInit(input_filter_init)
	return zend.SUCCESS
}
func SapiFlush() int {
	if sapi_module.GetFlush() != nil {
		sapi_module.GetFlush()(sapi_globals.GetServerContext())
		return zend.SUCCESS
	} else {
		return zend.FAILURE
	}
}
func SapiGetStat() *zend.ZendStatT {
	if sapi_module.GetGetStat() != nil {
		return sapi_module.GetGetStat()()
	} else {
		if sapi_globals.GetRequestInfo().GetPathTranslated() == nil || stat(sapi_globals.GetRequestInfo().GetPathTranslated(), &(sapi_globals.GetGlobalStat())) == -1 {
			return nil
		}
		return &(sapi_globals.GetGlobalStat())
	}
}
func SapiGetenv(name *byte, name_len int) *byte {
	if !(strncasecmp(name, "HTTP_PROXY", name_len)) {

		/* Ugly fix for HTTP_PROXY issue, see bug #72573 */

		return nil

		/* Ugly fix for HTTP_PROXY issue, see bug #72573 */

	}
	if sapi_module.GetGetenv() != nil {
		var value *byte
		var tmp *byte = sapi_module.GetGetenv()(name, name_len)
		if tmp != nil {
			value = zend._estrdup(tmp)
		} else {
			return nil
		}
		if sapi_module.GetInputFilter() != nil {
			sapi_module.GetInputFilter()(3, name, &value, strlen(value), nil)
		}
		return value
	}
	return nil
}
func SapiGetFd(fd *int) int {
	if sapi_module.GetGetFd() != nil {
		return sapi_module.GetGetFd()(fd)
	} else {
		return zend.FAILURE
	}
}
func SapiForceHttp10() int {
	if sapi_module.GetForceHttp10() != nil {
		return sapi_module.GetForceHttp10()()
	} else {
		return zend.FAILURE
	}
}
func SapiGetTargetUid(obj *uid_t) int {
	if sapi_module.GetGetTargetUid() != nil {
		return sapi_module.GetGetTargetUid()(obj)
	} else {
		return zend.FAILURE
	}
}
func SapiGetTargetGid(obj *gid_t) int {
	if sapi_module.GetGetTargetGid() != nil {
		return sapi_module.GetGetTargetGid()(obj)
	} else {
		return zend.FAILURE
	}
}
func SapiGetRequestTime() float64 {
	if sapi_globals.GetGlobalRequestTime() {
		return sapi_globals.GetGlobalRequestTime()
	}
	if sapi_module.GetGetRequestTime() != nil && sapi_globals.GetServerContext() {
		sapi_globals.SetGlobalRequestTime(sapi_module.GetGetRequestTime()())
	} else {
		var tp __struct__timeval = __struct__timeval{0}
		if !(gettimeofday(&tp, nil)) {
			sapi_globals.SetGlobalRequestTime(float64(tp.tv_sec + tp.tv_usec/1000000.0))
		} else {
			sapi_globals.SetGlobalRequestTime(float64(time(0)))
		}
	}
	return sapi_globals.GetGlobalRequestTime()
}
func SapiTerminateProcess() {
	if sapi_module.GetTerminateProcess() != nil {
		sapi_module.GetTerminateProcess()()
	}
}
func SapiAddRequestHeader(var_ *byte, var_len uint, val *byte, val_len uint, arg any) {
	var return_value *zend.Zval = (*zend.Zval)(arg)
	var str *byte = nil
	if var_len > 5 && var_[0] == 'H' && var_[1] == 'T' && var_[2] == 'T' && var_[3] == 'P' && var_[4] == '_' {
		var p *byte
		var_len -= 5
		p = var_ + 5
		str = zend._emalloc(var_len + 1)
		var_ = str
		*p++
		g.PostInc(&(*str)) = (*p) - 1
		for *p {
			if (*p) == '_' {
				g.PostInc(&(*str)) = '-'
				p++
				if *p {
					*p++
					g.PostInc(&(*str)) = (*p) - 1
				}
			} else if (*p) >= 'A' && (*p) <= 'Z' {
				g.PostInc(&(*str)) = g.PostInc(&(*p)) - 'A' + 'a'
			} else {
				*p++
				g.PostInc(&(*str)) = (*p) - 1
			}
		}
		*str = 0
	} else if var_len == g.SizeOf("\"CONTENT_TYPE\"")-1 && memcmp(var_, "CONTENT_TYPE", g.SizeOf("\"CONTENT_TYPE\"")-1) == 0 {
		var_ = "Content-Type"
	} else if var_len == g.SizeOf("\"CONTENT_LENGTH\"")-1 && memcmp(var_, "CONTENT_LENGTH", g.SizeOf("\"CONTENT_LENGTH\"")-1) == 0 {
		var_ = "Content-Length"
	} else {
		return
	}
	zend.AddAssocStringlEx(return_value, var_, var_len, val, val_len)
	if str != nil {
		zend._efree(var_)
	}
}

/* }}} */
