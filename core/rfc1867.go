// <<generate>>

package core

import (
	"sik/ext/standard"
	g "sik/runtime/grammar"
	"sik/zend"
)

// Source: <main/rfc1867.h>

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
  | Author:                                                              |
  +----------------------------------------------------------------------+
*/

// #define RFC1867_H

// # include "SAPI.h"

// #define MULTIPART_CONTENT_TYPE       "multipart/form-data"

// #define MULTIPART_EVENT_START       0

// #define MULTIPART_EVENT_FORMDATA       1

// #define MULTIPART_EVENT_FILE_START       2

// #define MULTIPART_EVENT_FILE_DATA       3

// #define MULTIPART_EVENT_FILE_END       4

// #define MULTIPART_EVENT_END       5

// @type MultipartEventStart struct

// @type MultipartEventFormdata struct

// @type MultipartEventFileStart struct

// @type MultipartEventFileData struct

// @type MultipartEventFileEnd struct

// @type MultipartEventEnd struct

type PhpRfc1867EncodingTranslationT func() int
type PhpRfc1867GetDetectOrderT func(list ***zend.ZendEncoding, list_size *int)
type PhpRfc1867SetInputEncodingT func(encoding *zend.ZendEncoding)
type PhpRfc1867GetwordT func(encoding *zend.ZendEncoding, line **byte, stop byte) *byte
type PhpRfc1867GetwordConfT func(encoding *zend.ZendEncoding, str *byte) *byte
type PhpRfc1867BasenameT func(encoding *zend.ZendEncoding, str *byte) *byte

// Source: <main/rfc1867.c>

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
   |          Jani Taskinen <jani@php.net>                                |
   +----------------------------------------------------------------------+
*/

// # include < stdio . h >

// # include "php.h"

// # include "php_open_temporary_file.h"

// # include "zend_globals.h"

// # include "php_globals.h"

// # include "php_variables.h"

// # include "rfc1867.h"

// # include "ext/standard/php_string.h"

// # include "zend_smart_string.h"

// #define DEBUG_FILE_UPLOAD       0

func DummyEncodingTranslation() int { return 0 }

var PhpRfc1867EncodingTranslation PhpRfc1867EncodingTranslationT = DummyEncodingTranslation
var PhpRfc1867GetDetectOrder PhpRfc1867GetDetectOrderT = nil
var PhpRfc1867SetInputEncoding PhpRfc1867SetInputEncodingT = nil
var PhpRfc1867Getword PhpRfc1867GetwordT = PhpApGetword
var PhpRfc1867GetwordConf PhpRfc1867GetwordConfT = PhpApGetwordConf
var PhpRfc1867Basename PhpRfc1867BasenameT = nil
var PhpRfc1867Callback func(event uint, event_data any, extra *any) int = nil

/* The longest property name we use in an uploaded file array */

// #define MAX_SIZE_OF_INDEX       sizeof ( "[tmp_name]" )

/* The longest anonymous name */

// #define MAX_SIZE_ANONNAME       33

/* Errors */

// #define UPLOAD_ERROR_OK       0

// #define UPLOAD_ERROR_A       1

// #define UPLOAD_ERROR_B       2

// #define UPLOAD_ERROR_C       3

// #define UPLOAD_ERROR_D       4

// #define UPLOAD_ERROR_E       6

// #define UPLOAD_ERROR_F       7

// #define UPLOAD_ERROR_X       8

func PhpRfc1867RegisterConstants() {
	zend.ZendRegisterLongConstant("UPLOAD_ERR_OK", g.SizeOf("\"UPLOAD_ERR_OK\"")-1, 0, 1<<0|1<<1, 0)
	zend.ZendRegisterLongConstant("UPLOAD_ERR_INI_SIZE", g.SizeOf("\"UPLOAD_ERR_INI_SIZE\"")-1, 1, 1<<0|1<<1, 0)
	zend.ZendRegisterLongConstant("UPLOAD_ERR_FORM_SIZE", g.SizeOf("\"UPLOAD_ERR_FORM_SIZE\"")-1, 2, 1<<0|1<<1, 0)
	zend.ZendRegisterLongConstant("UPLOAD_ERR_PARTIAL", g.SizeOf("\"UPLOAD_ERR_PARTIAL\"")-1, 3, 1<<0|1<<1, 0)
	zend.ZendRegisterLongConstant("UPLOAD_ERR_NO_FILE", g.SizeOf("\"UPLOAD_ERR_NO_FILE\"")-1, 4, 1<<0|1<<1, 0)
	zend.ZendRegisterLongConstant("UPLOAD_ERR_NO_TMP_DIR", g.SizeOf("\"UPLOAD_ERR_NO_TMP_DIR\"")-1, 6, 1<<0|1<<1, 0)
	zend.ZendRegisterLongConstant("UPLOAD_ERR_CANT_WRITE", g.SizeOf("\"UPLOAD_ERR_CANT_WRITE\"")-1, 7, 1<<0|1<<1, 0)
	zend.ZendRegisterLongConstant("UPLOAD_ERR_EXTENSION", g.SizeOf("\"UPLOAD_ERR_EXTENSION\"")-1, 8, 1<<0|1<<1, 0)
}

/* }}} */

func NormalizeProtectedVariable(varname *byte) {
	var s *byte = varname
	var index *byte = nil
	var indexend *byte = nil
	var p *byte

	/* overjump leading space */

	for (*s) == ' ' {
		s++
	}

	/* and remove it */

	if s != varname {
		memmove(varname, s, strlen(s)+1)
	}
	for p = varname; (*p) && (*p) != '['; p++ {
		switch *p {
		case ' ':

		case '.':
			*p = '_'
			break
		}
	}

	/* find index */

	index = strchr(varname, '[')
	if index != nil {
		index++
		s = index
	} else {
		return
	}

	/* done? */

	for index != nil {
		for (*index) == ' ' || (*index) == '\r' || (*index) == '\n' || (*index) == '\t' {
			index++
		}
		indexend = strchr(index, ']')
		if indexend != nil {
			indexend = indexend + 1
		} else {
			indexend = index + strlen(index)
		}
		if s != index {
			memmove(s, index, strlen(index)+1)
			s += indexend - index
		} else {
			s = indexend
		}
		if (*s) == '[' {
			s++
			index = s
		} else {
			index = nil
		}
	}
	*s = '0'
}

/* }}} */

func AddProtectedVariable(varname *byte) {
	NormalizeProtectedVariable(varname)
	zend.ZendHashStrAddEmptyElement(&(CoreGlobals.GetRfc1867ProtectedVariables()), varname, strlen(varname))
}

/* }}} */

func IsProtectedVariable(varname *byte) zend.ZendBool {
	NormalizeProtectedVariable(varname)
	return zend.ZendHashStrExists(&(CoreGlobals.GetRfc1867ProtectedVariables()), varname, strlen(varname))
}

/* }}} */

func SafePhpRegisterVariable(var_ *byte, strval *byte, val_len int, track_vars_array *zend.Zval, override_protection zend.ZendBool) {
	if override_protection != 0 || IsProtectedVariable(var_) == 0 {
		PhpRegisterVariableSafe(var_, strval, val_len, track_vars_array)
	}
}

/* }}} */

func SafePhpRegisterVariableEx(var_ *byte, val *zend.Zval, track_vars_array *zend.Zval, override_protection zend.ZendBool) {
	if override_protection != 0 || IsProtectedVariable(var_) == 0 {
		PhpRegisterVariableEx(var_, val, track_vars_array)
	}
}

/* }}} */

func RegisterHttpPostFilesVariable(strvar *byte, val *byte, http_post_files *zend.Zval, override_protection zend.ZendBool) {
	SafePhpRegisterVariable(strvar, val, strlen(val), http_post_files, override_protection)
}

/* }}} */

func RegisterHttpPostFilesVariableEx(var_ *byte, val *zend.Zval, http_post_files *zend.Zval, override_protection zend.ZendBool) {
	SafePhpRegisterVariableEx(var_, val, http_post_files, override_protection)
}

/* }}} */

func FreeFilename(el *zend.Zval) {
	var filename *zend.ZendString = el.value.str
	zend.ZendStringReleaseEx(filename, 0)
}
func DestroyUploadedFilesHash() {
	var el *zend.Zval
	for {
		var __ht *zend.HashTable = sapi_globals.GetRfc1867UploadedFiles()
		var _p *zend.Bucket = __ht.arData
		var _end *zend.Bucket = _p + __ht.nNumUsed
		for ; _p != _end; _p++ {
			var _z *zend.Zval = &_p.val

			if _z.u1.v.type_ == 0 {
				continue
			}
			el = _z
			var filename *zend.ZendString = el.value.str
			unlink(filename.val)
		}
		break
	}
	zend.ZendHashDestroy(sapi_globals.GetRfc1867UploadedFiles())
	zend._efree(sapi_globals.GetRfc1867UploadedFiles())
}

/* }}} */

// #define FILLUNIT       ( 1024 * 5 )

// @type MultipartBuffer struct
// @type MimeHeaderEntry struct

/*
 * Fill up the buffer with client data.
 * Returns number of bytes added to buffer.
 */

func FillBuffer(self *MultipartBuffer) int {
	var bytes_to_read int
	var total_read int = 0
	var actual_read int = 0

	/* shift the existing data if necessary */

	if self.GetBytesInBuffer() > 0 && self.GetBufBegin() != self.GetBuffer() {
		memmove(self.GetBuffer(), self.GetBufBegin(), self.GetBytesInBuffer())
	}
	self.SetBufBegin(self.GetBuffer())

	/* calculate the free space in the buffer */

	bytes_to_read = self.GetBufsize() - self.GetBytesInBuffer()

	/* read the required number of bytes */

	for bytes_to_read > 0 {
		var buf *byte = self.GetBuffer() + self.GetBytesInBuffer()
		actual_read = int(sapi_module.GetReadPost()(buf, bytes_to_read))

		/* update the buffer length */

		if actual_read > 0 {
			self.SetBytesInBuffer(self.GetBytesInBuffer() + actual_read)
			sapi_globals.SetReadPostBytes(sapi_globals.GetReadPostBytes() + actual_read)
			total_read += actual_read
			bytes_to_read -= actual_read
		} else {
			break
		}

		/* update the buffer length */

	}
	return total_read
}

/* eof if we are out of bytes, or if we hit the final boundary */

func MultipartBufferEof(self *MultipartBuffer) int {
	return self.GetBytesInBuffer() == 0 && FillBuffer(self) < 1
}

/* create new multipart_buffer structure */

func MultipartBufferNew(boundary *byte, boundary_len int) *MultipartBuffer {
	var self *MultipartBuffer = (*MultipartBuffer)(zend._ecalloc(1, g.SizeOf("multipart_buffer")))
	var minsize int = boundary_len + 6
	if minsize < 1024*5 {
		minsize = 1024 * 5
	}
	self.SetBuffer((*byte)(zend._ecalloc(1, minsize+1)))
	self.SetBufsize(minsize)
	zend.ZendSpprintf(&self.boundary, 0, "--%s", boundary)
	self.SetBoundaryNextLen(int(zend.ZendSpprintf(&self.boundary_next, 0, "\n--%s", boundary)))
	self.SetBufBegin(self.GetBuffer())
	self.SetBytesInBuffer(0)
	if PhpRfc1867EncodingTranslation() != 0 {
		PhpRfc1867GetDetectOrder(&self.detect_order, &self.detect_order_size)
	} else {
		self.SetDetectOrder(nil)
		self.SetDetectOrderSize(0)
	}
	self.SetInputEncoding(nil)
	return self
}

/*
 * Gets the next CRLF terminated line from the input buffer.
 * If it doesn't find a CRLF, and the buffer isn't completely full, returns
 * NULL; otherwise, returns the beginning of the null-terminated line,
 * minus the CRLF.
 *
 * Note that we really just look for LF terminated lines. This works
 * around a bug in internet explorer for the macintosh which sends mime
 * boundaries that are only LF terminated when you use an image submit
 * button in a multipart/form-data form.
 */

func NextLine(self *MultipartBuffer) *byte {
	/* look for LF in the data */

	var line *byte = self.GetBufBegin()
	var ptr *byte = memchr(self.GetBufBegin(), '\n', self.GetBytesInBuffer())
	if ptr != nil {

		/* terminate the string, remove CRLF */

		if ptr-line > 0 && (*(ptr - 1)) == '\r' {
			*(ptr - 1) = 0
		} else {
			*ptr = 0
		}

		/* bump the pointer */

		self.SetBufBegin(ptr + 1)
		self.SetBytesInBuffer(self.GetBytesInBuffer() - self.GetBufBegin() - line)
	} else {

		/* buffer isn't completely full, fail */

		if self.GetBytesInBuffer() < self.GetBufsize() {
			return nil
		}

		/* return entire buffer as a partial line */

		line[self.GetBufsize()] = 0
		self.SetBufBegin(ptr)
		self.SetBytesInBuffer(0)
	}
	return line
}

/* Returns the next CRLF terminated line from the client */

func GetLine(self *MultipartBuffer) *byte {
	var ptr *byte = NextLine(self)
	if ptr == nil {
		FillBuffer(self)
		ptr = NextLine(self)
	}
	return ptr
}

/* Free header entry */

func PhpFreeHdrEntry(h *MimeHeaderEntry) {
	if h.GetKey() != nil {
		zend._efree(h.GetKey())
	}
	if h.GetValue() != nil {
		zend._efree(h.GetValue())
	}
}

/* finds a boundary */

func FindBoundary(self *MultipartBuffer, boundary *byte) int {
	var line *byte

	/* loop through lines */

	for g.Assign(&line, GetLine(self)) {

		/* finished if we found the boundary */

		if !(strcmp(line, boundary)) {
			return 1
		}

		/* finished if we found the boundary */

	}

	/* didn't find the boundary */

	return 0

	/* didn't find the boundary */
}

/* parse headers */

func MultipartBufferHeaders(self *MultipartBuffer, header *zend.ZendLlist) int {
	var line *byte
	var entry MimeHeaderEntry = MimeHeaderEntry{0}
	var buf_value zend.SmartString = zend.SmartString{0}
	var key *byte = nil

	/* didn't find boundary, abort */

	if FindBoundary(self, self.GetBoundary()) == 0 {
		return 0
	}

	/* get lines of text, or CRLF_CRLF */

	for g.Assign(&line, GetLine(self)) && line[0] != '0' {

		/* add header to table */

		var value *byte = nil
		if PhpRfc1867EncodingTranslation() != 0 {
			self.SetInputEncoding(zend.ZendMultibyteEncodingDetector((*uint8)(line), strlen(line), self.GetDetectOrder(), self.GetDetectOrderSize()))
		}

		/* space in the beginning means same header */

		if !(isspace(line[0])) {
			value = strchr(line, ':')
		}
		if value != nil {
			if buf_value.c != nil && key != nil {

				/* new entry, add the old one to the list */

				zend.SmartString0(&buf_value)
				entry.SetKey(key)
				entry.SetValue(buf_value.c)
				zend.ZendLlistAddElement(header, &entry)
				buf_value.c = nil
				key = nil
			}
			*value = '0'
			for {
				value++
				if !(isspace(*value)) {
					break
				}
			}
			key = zend._estrdup(line)
			zend.SmartStringAppendlEx(&buf_value, value, strlen(value), 0)
		} else if buf_value.c != nil {
			zend.SmartStringAppendlEx(&buf_value, line, strlen(line), 0)
		} else {
			continue
		}
	}
	if buf_value.c != nil && key != nil {

		/* add the last one to the list */

		zend.SmartString0(&buf_value)
		entry.SetKey(key)
		entry.SetValue(buf_value.c)
		zend.ZendLlistAddElement(header, &entry)
	}
	return 1
}
func PhpMimeGetHdrValue(header zend.ZendLlist, key string) *byte {
	var entry *MimeHeaderEntry
	if key == nil {
		return nil
	}
	entry = zend.ZendLlistGetFirstEx(&header, nil)
	for entry != nil {
		if !(strcasecmp(entry.GetKey(), key)) {
			return entry.GetValue()
		}
		entry = zend.ZendLlistGetNextEx(&header, nil)
	}
	return nil
}
func PhpApGetword(encoding *zend.ZendEncoding, line **byte, stop byte) *byte {
	var pos *byte = *line
	var quote byte
	var res *byte
	for (*pos) && (*pos) != stop {
		if g.Assign(&quote, *pos) == '"' || quote == '\'' {
			pos++
			for (*pos) && (*pos) != quote {
				if (*pos) == '\\' && pos[1] && pos[1] == quote {
					pos += 2
				} else {
					pos++
				}
			}
			if *pos {
				pos++
			}
		} else {
			pos++
		}
	}
	if (*pos) == '0' {
		res = zend._estrdup(*line)
		*line += strlen(*line)
		return res
	}
	res = zend._estrndup(*line, pos-(*line))
	for (*pos) == stop {
		pos++
	}
	*line = pos
	return res
}
func SubstringConf(start *byte, len_ int, quote byte) *byte {
	var result *byte = zend._emalloc(len_ + 1)
	var resp *byte = result
	var i int
	for i = 0; i < len_ && start[i] != quote; i++ {
		if start[i] == '\\' && (start[i+1] == '\\' || quote && start[i+1] == quote) {
			g.PostInc(&(*resp)) = start[g.PreInc(&i)]
		} else {
			g.PostInc(&(*resp)) = start[i]
		}
	}
	*resp = '0'
	return result
}
func PhpApGetwordConf(encoding *zend.ZendEncoding, str *byte) *byte {
	for (*str) && isspace(*str) {
		str++
	}
	if !(*str) {
		return zend._estrdup("")
	}
	if (*str) == '"' || (*str) == '\'' {
		var quote byte = *str
		str++
		return SubstringConf(str, int(strlen(str)), quote)
	} else {
		var strend *byte = str
		for (*strend) && !(isspace(*strend)) {
			strend++
		}
		return SubstringConf(str, strend-str, 0)
	}
}
func PhpApBasename(encoding *zend.ZendEncoding, path *byte) *byte {
	var s *byte = strrchr(path, '\\')
	var s2 *byte = strrchr(path, '/')
	if s != nil && s2 != nil {
		if s > s2 {
			s++
		} else {
			s2++
			s = s2
		}
		return s
	} else if s != nil {
		s++
		return s
	} else if s2 != nil {
		s2++
		return s2
	}
	return path
}

/*
 * Search for a string in a fixed-length byte string.
 * If partial is true, partial matches are allowed at the end of the buffer.
 * Returns NULL if not found, or a pointer to the start of the first match.
 */

func PhpApMemstr(haystack *byte, haystacklen int, needle *byte, needlen int, partial int) any {
	var len_ int = haystacklen
	var ptr *byte = haystack

	/* iterate through first character matches */

	for g.Assign(&ptr, memchr(ptr, needle[0], len_)) {

		/* calculate length after match */

		len_ = haystacklen - (ptr - (*byte)(haystack))

		/* done if matches up to capacity of buffer */

		if memcmp(needle, ptr, g.Cond(needlen < len_, needlen, len_)) == 0 && (partial != 0 || len_ >= needlen) {
			break
		}

		/* next character */

		ptr++
		len_--
	}
	return ptr
}

/* read until a boundary condition */

func MultipartBufferRead(self *MultipartBuffer, buf *byte, bytes int, end *int) int {
	var len_ int
	var max int
	var bound *byte

	/* fill buffer if needed */

	if bytes > int(self.GetBytesInBuffer()) {
		FillBuffer(self)
	}

	/* look for a potential boundary match, only read data up to that point */

	if g.Assign(&bound, PhpApMemstr(self.GetBufBegin(), self.GetBytesInBuffer(), self.GetBoundaryNext(), self.GetBoundaryNextLen(), 1)) {
		max = bound - self.GetBufBegin()
		if end != nil && PhpApMemstr(self.GetBufBegin(), self.GetBytesInBuffer(), self.GetBoundaryNext(), self.GetBoundaryNextLen(), 0) {
			*end = 1
		}
	} else {
		max = self.GetBytesInBuffer()
	}

	/* maximum number of bytes we are reading */

	if max < bytes-1 {
		len_ = max
	} else {
		len_ = bytes - 1
	}

	/* if we read any data... */

	if len_ > 0 {

		/* copy the data */

		memcpy(buf, self.GetBufBegin(), len_)
		buf[len_] = 0
		if bound != nil && len_ > 0 && buf[len_-1] == '\r' {
			buf[g.PreDec(&len_)] = 0
		}

		/* update the buffer */

		self.SetBytesInBuffer(self.GetBytesInBuffer() - int(len_))
		self.SetBufBegin(self.GetBufBegin() + len_)
	}
	return len_
}

/*
  XXX: this is horrible memory-usage-wise, but we only expect
  to do this on small pieces of form data.
*/

func MultipartBufferReadBody(self *MultipartBuffer, len_ *int) *byte {
	var buf []byte
	var out *byte = nil
	var total_bytes int = 0
	var read_bytes int = 0
	for g.Assign(&read_bytes, MultipartBufferRead(self, buf, g.SizeOf("buf"), nil)) {
		out = zend._erealloc(out, total_bytes+read_bytes+1)
		memcpy(out+total_bytes, buf, read_bytes)
		total_bytes += read_bytes
	}
	if out != nil {
		out[total_bytes] = '0'
	}
	*len_ = total_bytes
	return out
}

/* }}} */

func Rfc1867PostHandler(content_type_dup *byte, arg any) {
	var boundary *byte
	var s *byte = nil
	var boundary_end *byte = nil
	var start_arr *byte = nil
	var array_index *byte = nil
	var lbuf *byte = nil
	var abuf *byte = nil
	var temp_filename *zend.ZendString = nil
	var boundary_len int = 0
	var cancel_upload int = 0
	var is_arr_upload int = 0
	var array_len int = 0
	var total_bytes int64 = 0
	var max_file_size int64 = 0
	var skip_upload int = 0
	var anonindex int = 0
	var is_anonymous int
	var uploaded_files *zend.HashTable = nil
	var mbuff *MultipartBuffer
	var array_ptr *zend.Zval = (*zend.Zval)(arg)
	var fd int = -1
	var header zend.ZendLlist
	var event_extra_data any = nil
	var llen uint = 0
	var upload_cnt int = zend.ZendIniLong("max_file_uploads", g.SizeOf("\"max_file_uploads\"")-1, 0)
	var internal_encoding *zend.ZendEncoding = zend.ZendMultibyteGetInternalEncoding()
	var getword PhpRfc1867GetwordT
	var getword_conf PhpRfc1867GetwordConfT
	var _basename PhpRfc1867BasenameT
	var count zend.ZendLong = 0
	if PhpRfc1867EncodingTranslation() != 0 && internal_encoding != nil {
		getword = PhpRfc1867Getword
		getword_conf = PhpRfc1867GetwordConf
		_basename = PhpRfc1867Basename
	} else {
		getword = PhpApGetword
		getword_conf = PhpApGetwordConf
		_basename = PhpApBasename
	}
	if sapi_globals.GetPostMaxSize() > 0 && sapi_globals.GetRequestInfo().GetContentLength() > sapi_globals.GetPostMaxSize() {
		sapi_module.GetSapiError()(1<<1, "POST Content-Length of "+"%"+"lld"+" bytes exceeds the limit of "+"%"+"lld"+" bytes", sapi_globals.GetRequestInfo().GetContentLength(), sapi_globals.GetPostMaxSize())
		return
	}

	/* Get the boundary */

	boundary = strstr(content_type_dup, "boundary")
	if boundary == nil {
		var content_type_len int = int(strlen(content_type_dup))
		var content_type_lcase *byte = zend._estrndup(content_type_dup, content_type_len)
		standard.PhpStrtolower(content_type_lcase, content_type_len)
		boundary = strstr(content_type_lcase, "boundary")
		if boundary != nil {
			boundary = content_type_dup + (boundary - content_type_lcase)
		}
		zend._efree(content_type_lcase)
	}
	if boundary == nil || !(g.Assign(&boundary, strchr(boundary, '='))) {
		sapi_module.GetSapiError()(1<<1, "Missing boundary in multipart/form-data POST data")
		return
	}
	boundary++
	boundary_len = int(strlen(boundary))
	if boundary[0] == '"' {
		boundary++
		boundary_end = strchr(boundary, '"')
		if boundary_end == nil {
			sapi_module.GetSapiError()(1<<1, "Invalid boundary in multipart/form-data POST data")
			return
		}
	} else {

		/* search for the end of the boundary */

		boundary_end = strpbrk(boundary, ",;")

		/* search for the end of the boundary */

	}
	if boundary_end != nil {
		boundary_end[0] = '0'
		boundary_len = boundary_end - boundary
	}

	/* Initialize the buffer */

	if !(g.Assign(&mbuff, MultipartBufferNew(boundary, boundary_len))) {
		sapi_module.GetSapiError()(1<<1, "Unable to initialize the input buffer")
		return
	}

	/* Initialize $_FILES[] */

	zend._zendHashInit(&(CoreGlobals.GetRfc1867ProtectedVariables()), 8, nil, 0)
	uploaded_files = (*zend.HashTable)(zend._emalloc(g.SizeOf("HashTable")))
	zend._zendHashInit(uploaded_files, 8, FreeFilename, 0)
	sapi_globals.SetRfc1867UploadedFiles(uploaded_files)
	if CoreGlobals.GetHttpGlobals()[5].u1.v.type_ != 7 {

		/* php_auto_globals_create_files() might have already done that */

		var __arr *zend.ZendArray = zend._zendNewArray(0)
		var __z *zend.Zval = &CoreGlobals.GetHttpGlobals()[5]
		__z.value.arr = __arr
		__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8

		/* php_auto_globals_create_files() might have already done that */

	}
	zend.ZendLlistInit(&header, g.SizeOf("mime_header_entry"), zend.LlistDtorFuncT(PhpFreeHdrEntry), 0)
	if PhpRfc1867Callback != nil {
		var event_start MultipartEventStart
		event_start.SetContentLength(sapi_globals.GetRequestInfo().GetContentLength())
		if PhpRfc1867Callback(0, &event_start, &event_extra_data) == zend.FAILURE {
			goto fileupload_done
		}
	}
	for MultipartBufferEof(mbuff) == 0 {
		var buff []byte
		var cd *byte = nil
		var param *byte = nil
		var filename *byte = nil
		var tmp *byte = nil
		var blen int = 0
		var wlen int = 0
		var offset zend.ZendOffT
		zend.ZendLlistClean(&header)
		if MultipartBufferHeaders(mbuff, &header) == 0 {
			goto fileupload_done
		}
		if g.Assign(&cd, PhpMimeGetHdrValue(header, "Content-Disposition")) {
			var pair *byte = nil
			var end int = 0
			for isspace(*cd) {
				cd++
			}
			for (*cd) && g.Assign(&pair, getword(mbuff.GetInputEncoding(), &cd, ';')) {
				var key *byte = nil
				var word *byte = pair
				for isspace(*cd) {
					cd++
				}
				if strchr(pair, '=') {
					key = getword(mbuff.GetInputEncoding(), &pair, '=')
					if !(strcasecmp(key, "name")) {
						if param != nil {
							zend._efree(param)
						}
						param = getword_conf(mbuff.GetInputEncoding(), pair)
						if mbuff.GetInputEncoding() != nil && internal_encoding != nil {
							var new_param *uint8
							var new_param_len int
							if size_t-1 != zend.ZendMultibyteEncodingConverter(&new_param, &new_param_len, (*uint8)(param), strlen(param), internal_encoding, mbuff.GetInputEncoding()) {
								zend._efree(param)
								param = (*byte)(new_param)
							}
						}
					} else if !(strcasecmp(key, "filename")) {
						if filename != nil {
							zend._efree(filename)
						}
						filename = getword_conf(mbuff.GetInputEncoding(), pair)
						if mbuff.GetInputEncoding() != nil && internal_encoding != nil {
							var new_filename *uint8
							var new_filename_len int
							if size_t-1 != zend.ZendMultibyteEncodingConverter(&new_filename, &new_filename_len, (*uint8)(filename), strlen(filename), internal_encoding, mbuff.GetInputEncoding()) {
								zend._efree(filename)
								filename = (*byte)(new_filename)
							}
						}
					}
				}
				if key != nil {
					zend._efree(key)
				}
				zend._efree(word)
			}

			/* Normal form variable, safe to read all data into memory */

			if filename == nil && param != nil {
				var value_len int
				var value *byte = MultipartBufferReadBody(mbuff, &value_len)
				var new_val_len int
				if value == nil {
					value = zend._estrdup("")
					value_len = 0
				}
				if mbuff.GetInputEncoding() != nil && internal_encoding != nil {
					var new_value *uint8
					var new_value_len int
					if size_t-1 != zend.ZendMultibyteEncodingConverter(&new_value, &new_value_len, (*uint8)(value), value_len, internal_encoding, mbuff.GetInputEncoding()) {
						zend._efree(value)
						value = (*byte)(new_value)
						value_len = new_value_len
					}
				}
				if g.PreInc(&count) <= CoreGlobals.GetMaxInputVars() && sapi_module.GetInputFilter()(0, param, &value, value_len, &new_val_len) != 0 {
					if PhpRfc1867Callback != nil {
						var event_formdata MultipartEventFormdata
						var newlength int = new_val_len
						event_formdata.SetPostBytesProcessed(sapi_globals.GetReadPostBytes())
						event_formdata.SetName(param)
						event_formdata.SetValue(&value)
						event_formdata.SetLength(new_val_len)
						event_formdata.SetNewlength(&newlength)
						if PhpRfc1867Callback(1, &event_formdata, &event_extra_data) == zend.FAILURE {
							zend._efree(param)
							zend._efree(value)
							continue
						}
						new_val_len = newlength
					}
					SafePhpRegisterVariable(param, value, new_val_len, array_ptr, 0)
				} else {
					if count == CoreGlobals.GetMaxInputVars()+1 {
						PhpErrorDocref(nil, 1<<1, "Input variables exceeded "+"%"+"lld"+". To increase the limit change max_input_vars in php.ini.", CoreGlobals.GetMaxInputVars())
					}
					if PhpRfc1867Callback != nil {
						var event_formdata MultipartEventFormdata
						event_formdata.SetPostBytesProcessed(sapi_globals.GetReadPostBytes())
						event_formdata.SetName(param)
						event_formdata.SetValue(&value)
						event_formdata.SetLength(value_len)
						event_formdata.SetNewlength(nil)
						PhpRfc1867Callback(1, &event_formdata, &event_extra_data)
					}
				}
				if !(strcasecmp(param, "MAX_FILE_SIZE")) {
					max_file_size = atoll(value)
				}
				zend._efree(param)
				zend._efree(value)
				continue
			}

			/* If file_uploads=off, skip the file part */

			if CoreGlobals.GetFileUploads() == 0 {
				skip_upload = 1
			} else if upload_cnt <= 0 {
				skip_upload = 1
				sapi_module.GetSapiError()(1<<1, "Maximum number of allowable file uploads has been exceeded")
			}

			/* Return with an error if the posted data is garbled */

			if param == nil && filename == nil {
				sapi_module.GetSapiError()(1<<1, "File Upload Mime headers garbled")
				goto fileupload_done
			}
			if param == nil {
				is_anonymous = 1
				param = zend._emalloc(33)
				ApPhpSnprintf(param, 33, "%u", g.PostInc(&anonindex))
			} else {
				is_anonymous = 0
			}

			/* New Rule: never repair potential malicious user input */

			if skip_upload == 0 {
				var c long = 0
				tmp = param
				for *tmp {
					if (*tmp) == '[' {
						c++
					} else if (*tmp) == ']' {
						c--
						if tmp[1] && tmp[1] != '[' {
							skip_upload = 1
							break
						}
					}
					if c < 0 {
						skip_upload = 1
						break
					}
					tmp++
				}

				/* Brackets should always be closed */

				if c != 0 {
					skip_upload = 1
				}

				/* Brackets should always be closed */

			}
			cancel_upload = 0
			total_bytes = cancel_upload
			temp_filename = nil
			fd = -1
			if skip_upload == 0 && PhpRfc1867Callback != nil {
				var event_file_start MultipartEventFileStart
				event_file_start.SetPostBytesProcessed(sapi_globals.GetReadPostBytes())
				event_file_start.SetName(param)
				event_file_start.SetFilename(&filename)
				if PhpRfc1867Callback(2, &event_file_start, &event_extra_data) == zend.FAILURE {
					temp_filename = nil
					zend._efree(param)
					zend._efree(filename)
					continue
				}
			}
			if skip_upload != 0 {
				zend._efree(param)
				zend._efree(filename)
				continue
			}
			if filename[0] == '0' {
				cancel_upload = 4
			}
			offset = 0
			end = 0
			if cancel_upload == 0 {

				/* only bother to open temp file if we have data */

				blen = MultipartBufferRead(mbuff, buff, g.SizeOf("buff"), &end)

				/* in non-debug mode we have no problem with 0-length files */

				fd = PhpOpenTemporaryFdEx(CoreGlobals.GetUploadTmpDir(), "php", &temp_filename, 1<<0)
				upload_cnt--
				if fd == -1 {
					sapi_module.GetSapiError()(1<<1, "File upload error - unable to create a temporary file")
					cancel_upload = 6
				}

				/* in non-debug mode we have no problem with 0-length files */

			}
			for cancel_upload == 0 && blen > 0 {
				if PhpRfc1867Callback != nil {
					var event_file_data MultipartEventFileData
					event_file_data.SetPostBytesProcessed(sapi_globals.GetReadPostBytes())
					event_file_data.SetOffset(offset)
					event_file_data.SetData(buff)
					event_file_data.SetLength(blen)
					event_file_data.SetNewlength(&blen)
					if PhpRfc1867Callback(3, &event_file_data, &event_extra_data) == zend.FAILURE {
						cancel_upload = 8
						continue
					}
				}
				if CoreGlobals.GetUploadMaxFilesize() > 0 && zend_long(total_bytes+blen) > CoreGlobals.GetUploadMaxFilesize() {
					cancel_upload = 1
				} else if max_file_size != 0 && zend_long(total_bytes+blen) > max_file_size {
					cancel_upload = 2
				} else if blen > 0 {
					wlen = write(fd, buff, blen)
					if wlen == size_t-1 {

						/* write failed */

						cancel_upload = 7

						/* write failed */

					} else if wlen < blen {
						cancel_upload = 7
					} else {
						total_bytes += wlen
					}
					offset += wlen
				}

				/* read data for next iteration */

				blen = MultipartBufferRead(mbuff, buff, g.SizeOf("buff"), &end)

				/* read data for next iteration */

			}
			if fd != -1 {
				close(fd)
			}
			if cancel_upload == 0 && end == 0 {
				cancel_upload = 3
			}
			if PhpRfc1867Callback != nil {
				var event_file_end MultipartEventFileEnd
				event_file_end.SetPostBytesProcessed(sapi_globals.GetReadPostBytes())
				if temp_filename != nil {
					event_file_end.SetTempFilename(temp_filename.val)
				} else {
					event_file_end.SetTempFilename(nil)
				}
				event_file_end.SetCancelUpload(cancel_upload)
				if PhpRfc1867Callback(4, &event_file_end, &event_extra_data) == zend.FAILURE {
					cancel_upload = 8
				}
			}
			if cancel_upload != 0 {
				if temp_filename != nil {
					if cancel_upload != 6 {
						unlink(temp_filename.val)
					}
					zend.ZendStringReleaseEx(temp_filename, 0)
				}
				temp_filename = nil
			} else {
				zend.ZendHashAddPtr(sapi_globals.GetRfc1867UploadedFiles(), temp_filename, temp_filename)
			}

			/* is_arr_upload is true when name of file upload field
			 * ends in [.*]
			 * start_arr is set to point to 1st [ */

			is_arr_upload = g.Assign(&start_arr, strchr(param, '[')) && param[strlen(param)-1] == ']'
			if is_arr_upload != 0 {
				array_len = strlen(start_arr)
				if array_index != nil {
					zend._efree(array_index)
				}
				array_index = zend._estrndup(start_arr+1, array_len-2)
			}

			/* Add $foo_name */

			if llen < strlen(param)+g.SizeOf("\"[tmp_name]\"")+1 {
				llen = int(strlen(param))
				lbuf = (*byte)(zend._safeErealloc(lbuf, llen, 1, g.SizeOf("\"[tmp_name]\"")+1))
				llen += g.SizeOf("\"[tmp_name]\"") + 1
			}
			if is_arr_upload != 0 {
				if abuf != nil {
					zend._efree(abuf)
				}
				abuf = zend._estrndup(param, strlen(param)-array_len)
				ApPhpSnprintf(lbuf, llen, "%s_name[%s]", abuf, array_index)
			} else {
				ApPhpSnprintf(lbuf, llen, "%s_name", param)
			}

			/* The \ check should technically be needed for win32 systems only where
			 * it is a valid path separator. However, IE in all it's wisdom always sends
			 * the full path of the file on the user's filesystem, which means that unless
			 * the user does basename() they get a bogus file name. Until IE's user base drops
			 * to nill or problem is fixed this code must remain enabled for all systems. */

			s = _basename(internal_encoding, filename)
			if s == nil {
				s = filename
			}
			if is_anonymous == 0 {
				SafePhpRegisterVariable(lbuf, s, strlen(s), nil, 0)
			}

			/* Add $foo[name] */

			if is_arr_upload != 0 {
				ApPhpSnprintf(lbuf, llen, "%s[name][%s]", abuf, array_index)
			} else {
				ApPhpSnprintf(lbuf, llen, "%s[name]", param)
			}
			RegisterHttpPostFilesVariable(lbuf, s, &CoreGlobals.GetHttpGlobals()[5], 0)
			zend._efree(filename)
			s = nil

			/* Possible Content-Type: */

			if cancel_upload != 0 || !(g.Assign(&cd, PhpMimeGetHdrValue(header, "Content-Type"))) {
				cd = ""
			} else {

				/* fix for Opera 6.01 */

				s = strchr(cd, ';')
				if s != nil {
					*s = '0'
				}
			}

			/* Add $foo_type */

			if is_arr_upload != 0 {
				ApPhpSnprintf(lbuf, llen, "%s_type[%s]", abuf, array_index)
			} else {
				ApPhpSnprintf(lbuf, llen, "%s_type", param)
			}
			if is_anonymous == 0 {
				SafePhpRegisterVariable(lbuf, cd, strlen(cd), nil, 0)
			}

			/* Add $foo[type] */

			if is_arr_upload != 0 {
				ApPhpSnprintf(lbuf, llen, "%s[type][%s]", abuf, array_index)
			} else {
				ApPhpSnprintf(lbuf, llen, "%s[type]", param)
			}
			RegisterHttpPostFilesVariable(lbuf, cd, &CoreGlobals.GetHttpGlobals()[5], 0)

			/* Restore Content-Type Header */

			if s != nil {
				*s = ';'
			}
			s = ""

			/* store temp_filename as-is (in case upload_tmp_dir
			 * contains escapeable characters. escape only the variable name.) */

			var zfilename zend.Zval

			/* Initialize variables */

			AddProtectedVariable(param)

			/* if param is of form xxx[.*] this will cut it to xxx */

			if is_anonymous == 0 {
				if temp_filename != nil {
					var __z *zend.Zval = &zfilename
					var __s *zend.ZendString = temp_filename
					__z.value.str = __s
					if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
						__z.u1.type_info = 6
					} else {
						zend.ZendGcAddref(&__s.gc)
						__z.u1.type_info = 6 | 1<<0<<8
					}
				} else {
					var __z *zend.Zval = &zfilename
					var __s *zend.ZendString = zend.ZendEmptyString
					__z.value.str = __s
					__z.u1.type_info = 6
				}
				SafePhpRegisterVariableEx(param, &zfilename, nil, 1)
			}

			/* Add $foo[tmp_name] */

			if is_arr_upload != 0 {
				ApPhpSnprintf(lbuf, llen, "%s[tmp_name][%s]", abuf, array_index)
			} else {
				ApPhpSnprintf(lbuf, llen, "%s[tmp_name]", param)
			}
			AddProtectedVariable(lbuf)
			if temp_filename != nil {
				var __z *zend.Zval = &zfilename
				var __s *zend.ZendString = temp_filename
				__z.value.str = __s
				if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
					__z.u1.type_info = 6
				} else {
					zend.ZendGcAddref(&__s.gc)
					__z.u1.type_info = 6 | 1<<0<<8
				}
			} else {
				var __z *zend.Zval = &zfilename
				var __s *zend.ZendString = zend.ZendEmptyString
				__z.value.str = __s
				__z.u1.type_info = 6
			}
			RegisterHttpPostFilesVariableEx(lbuf, &zfilename, &CoreGlobals.GetHttpGlobals()[5], 1)
			var file_size zend.Zval
			var error_type zend.Zval
			var size_overflow int = 0
			var file_size_buf []byte
			var __z *zend.Zval = &error_type
			__z.value.lval = cancel_upload
			__z.u1.type_info = 4

			/* Add $foo[error] */

			if cancel_upload != 0 {
				var __z *zend.Zval = &file_size
				__z.value.lval = 0
				__z.u1.type_info = 4
			} else {
				if total_bytes > INT64_MAX {
					var __len int = ApPhpSnprintf(file_size_buf, 65, "%"+"lld", total_bytes)
					file_size_buf[__len] = '0'
					size_overflow = 1
				} else {
					var __z *zend.Zval = &file_size
					__z.value.lval = total_bytes
					__z.u1.type_info = 4
				}
			}
			if is_arr_upload != 0 {
				ApPhpSnprintf(lbuf, llen, "%s[error][%s]", abuf, array_index)
			} else {
				ApPhpSnprintf(lbuf, llen, "%s[error]", param)
			}
			RegisterHttpPostFilesVariableEx(lbuf, &error_type, &CoreGlobals.GetHttpGlobals()[5], 0)

			/* Add $foo_size */

			if is_arr_upload != 0 {
				ApPhpSnprintf(lbuf, llen, "%s_size[%s]", abuf, array_index)
			} else {
				ApPhpSnprintf(lbuf, llen, "%s_size", param)
			}
			if is_anonymous == 0 {
				if size_overflow != 0 {
					var _s *byte = file_size_buf
					var __z *zend.Zval = &file_size
					var __s *zend.ZendString = zend.ZendStringInit(_s, strlen(_s), 0)
					__z.value.str = __s
					__z.u1.type_info = 6 | 1<<0<<8
				}
				SafePhpRegisterVariableEx(lbuf, &file_size, nil, size_overflow)
			}

			/* Add $foo[size] */

			if is_arr_upload != 0 {
				ApPhpSnprintf(lbuf, llen, "%s[size][%s]", abuf, array_index)
			} else {
				ApPhpSnprintf(lbuf, llen, "%s[size]", param)
			}
			if size_overflow != 0 {
				var _s *byte = file_size_buf
				var __z *zend.Zval = &file_size
				var __s *zend.ZendString = zend.ZendStringInit(_s, strlen(_s), 0)
				__z.value.str = __s
				__z.u1.type_info = 6 | 1<<0<<8
			}
			RegisterHttpPostFilesVariableEx(lbuf, &file_size, &CoreGlobals.GetHttpGlobals()[5], size_overflow)
			zend._efree(param)
		}
	}
fileupload_done:
	if PhpRfc1867Callback != nil {
		var event_end MultipartEventEnd
		event_end.SetPostBytesProcessed(sapi_globals.GetReadPostBytes())
		PhpRfc1867Callback(5, &event_end, &event_extra_data)
	}
	if lbuf != nil {
		zend._efree(lbuf)
	}
	if abuf != nil {
		zend._efree(abuf)
	}
	if array_index != nil {
		zend._efree(array_index)
	}
	zend.ZendHashDestroy(&(CoreGlobals.GetRfc1867ProtectedVariables()))
	zend.ZendLlistDestroy(&header)
	if mbuff.GetBoundaryNext() != nil {
		zend._efree(mbuff.GetBoundaryNext())
	}
	if mbuff.GetBoundary() != nil {
		zend._efree(mbuff.GetBoundary())
	}
	if mbuff.GetBuffer() != nil {
		zend._efree(mbuff.GetBuffer())
	}
	if mbuff != nil {
		zend._efree(mbuff)
	}
}

/* }}} */

func PhpRfc1867SetMultibyteCallbacks(encoding_translation PhpRfc1867EncodingTranslationT, get_detect_order PhpRfc1867GetDetectOrderT, set_input_encoding PhpRfc1867SetInputEncodingT, getword PhpRfc1867GetwordT, getword_conf PhpRfc1867GetwordConfT, basename PhpRfc1867BasenameT) {
	PhpRfc1867EncodingTranslation = encoding_translation
	PhpRfc1867GetDetectOrder = get_detect_order
	PhpRfc1867SetInputEncoding = set_input_encoding
	PhpRfc1867Getword = getword
	PhpRfc1867GetwordConf = getword_conf
	PhpRfc1867Basename = basename
}

/* }}} */
