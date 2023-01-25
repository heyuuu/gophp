// <<generate>>

package core

import (
	b "sik/builtin"
	"sik/zend"
)

func SG(v *zend.Zval) __auto__                          { return sapi_globals.v }
func SapiAddHeader(a *byte, b int, c zend.ZendBool) int { return SapiAddHeaderEx(a, b, c, 1) }
func _typeDtor(zv *zend.Zval)                           { zend.Free(zend.Z_PTR_P(zv)) }
func SapiGlobalsCtor(sapi_globals *sapi_globals_struct) {
	memset(sapi_globals, 0, b.SizeOf("* sapi_globals"))
	zend.ZendHashInitEx(&sapi_globals.known_post_content_types, 8, nil, _typeDtor, 1, 0)
	PhpSetupSapiContentTypes()
}
func SapiGlobalsDtor(sapi_globals *sapi_globals_struct) {
	zend.ZendHashDestroy(&sapi_globals.known_post_content_types)
}
func SapiStartup(sf *sapi_module_struct) {
	sf.SetIniEntries(nil)
	sapi_module = *sf
	SapiGlobalsCtor(&sapi_globals)
}
func SapiShutdown()                          { SapiGlobalsDtor(&sapi_globals) }
func SapiFreeHeader(sapi_header *SapiHeader) { zend.Efree(sapi_header.GetHeader()) }
func ZifHeaderRegisterCallback(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var callback_func *zend.Zval
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "z", &callback_func) == zend.FAILURE {
		return
	}
	if zend.ZendIsCallable(callback_func, 0, nil) == 0 {
		zend.RETVAL_FALSE
		return
	}
	if zend.Z_TYPE(SG(callback_func)) != zend.IS_UNDEF {
		zend.ZvalPtrDtor(&SG(callback_func))
		SG(fci_cache) = zend.EmptyFcallInfoCache
	}
	zend.ZVAL_COPY(&SG(callback_func), callback_func)
	zend.RETVAL_TRUE
	return
}
func SapiRunHeaderCallback(callback *zend.Zval) {
	var error int
	var fci zend.ZendFcallInfo
	var callback_error *byte = nil
	var retval zend.Zval
	if zend.ZendFcallInfoInit(callback, 0, &fci, &SG(fci_cache), nil, &callback_error) == zend.SUCCESS {
		fci.retval = &retval
		error = zend.ZendCallFunction(&fci, &SG(fci_cache))
		if error == zend.FAILURE {
			goto callback_failed
		} else {
			zend.ZvalPtrDtor(&retval)
		}
	} else {
	callback_failed:
		PhpErrorDocref(nil, zend.E_WARNING, "Could not call the sapi_header_callback")
	}
	if callback_error != nil {
		zend.Efree(callback_error)
	}
}
func SapiHandlePost(arg any) {
	if SG(request_info).post_entry && SG(request_info).content_type_dup {
		SG(request_info).post_entry.post_handler(SG(request_info).content_type_dup, arg)
		zend.Efree(SG(request_info).content_type_dup)
		SG(request_info).content_type_dup = nil
	}
}
func SapiReadPostData() {
	var post_entry *SapiPostEntry
	var content_type_length uint32 = uint32(strlen(SG(request_info).content_type))
	var content_type *byte = zend.Estrndup(SG(request_info).content_type, content_type_length)
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

	if b.Assign(&post_entry, zend.ZendHashStrFindPtr(&SG(known_post_content_types), content_type, content_type_length)) != nil {

		/* found one, register it for use */

		SG(request_info).post_entry = post_entry
		post_reader_func = post_entry.GetPostReader()
	} else {

		/* fallback */

		SG(request_info).post_entry = nil
		if sapi_module.GetDefaultPostReader() == nil {

			/* no default reader ? */

			SG(request_info).content_type_dup = nil
			sapi_module.GetSapiError()(zend.E_WARNING, "Unsupported content type:  '%s'", content_type)
			return
		}
	}
	if oldchar {
		*(p - 1) = oldchar
	}
	SG(request_info).content_type_dup = content_type
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

		SG(read_post_bytes) += read_bytes

		/* gogo */

	}
	if read_bytes < buflen {

		/* done */

		SG(post_read) = 1

		/* done */

	}
	return read_bytes
}
func SapiReadStandardFormData() {
	if SG(post_max_size) > 0 && SG(request_info).content_length > SG(post_max_size) {
		PhpErrorDocref(nil, zend.E_WARNING, "POST Content-Length of "+zend.ZEND_LONG_FMT+" bytes exceeds the limit of "+zend.ZEND_LONG_FMT+" bytes", SG(request_info).content_length, SG(post_max_size))
		return
	}
	SG(request_info).request_body = PhpStreamTempCreateEx(TEMP_STREAM_DEFAULT, SAPI_POST_BLOCK_SIZE, PG(upload_tmp_dir))
	if sapi_module.GetReadPost() != nil {
		var read_bytes int
		for {
			var buffer []byte
			read_bytes = SapiReadPostBlock(buffer, SAPI_POST_BLOCK_SIZE)
			if read_bytes > 0 {
				if PhpStreamWrite(SG(request_info).request_body, buffer, read_bytes) != read_bytes {

					/* if parts of the stream can't be written, purge it completely */

					PhpStreamTruncateSetSize(SG(request_info).request_body, 0)
					PhpErrorDocref(nil, zend.E_WARNING, "POST data can't be buffered; all data discarded")
					break
				}
			}
			if SG(post_max_size) > 0 && SG(read_post_bytes) > SG(post_max_size) {
				PhpErrorDocref(nil, zend.E_WARNING, "Actual POST length does not match Content-Length, and exceeds "+zend.ZEND_LONG_FMT+" bytes", SG(post_max_size))
				break
			}
			if read_bytes < SAPI_POST_BLOCK_SIZE {

				/* done */

				break

				/* done */

			}
		}
		PhpStreamRewind(SG(request_info).request_body)
	}
}
func GetDefaultContentType(prefix_len uint32, len_ *uint32) *byte {
	var mimetype *byte
	var charset *byte
	var content_type *byte
	var mimetype_len uint32
	var charset_len uint32
	if SG(default_mimetype) {
		mimetype = SG(default_mimetype)
		mimetype_len = uint32(strlen(SG(default_mimetype)))
	} else {
		mimetype = SAPI_DEFAULT_MIMETYPE
		mimetype_len = b.SizeOf("SAPI_DEFAULT_MIMETYPE") - 1
	}
	if SG(default_charset) {
		charset = SG(default_charset)
		charset_len = uint32(strlen(SG(default_charset)))
	} else {
		charset = SAPI_DEFAULT_CHARSET
		charset_len = b.SizeOf("SAPI_DEFAULT_CHARSET") - 1
	}
	if (*charset) && strncasecmp(mimetype, "text/", 5) == 0 {
		var p *byte
		*len_ = prefix_len + mimetype_len + b.SizeOf("\"; charset=\"") - 1 + charset_len
		content_type = (*byte)(zend.Emalloc((*len_) + 1))
		p = content_type + prefix_len
		memcpy(p, mimetype, mimetype_len)
		p += mimetype_len
		memcpy(p, "; charset=", b.SizeOf("\"; charset=\"")-1)
		p += b.SizeOf("\"; charset=\"") - 1
		memcpy(p, charset, charset_len+1)
	} else {
		*len_ = prefix_len + mimetype_len
		content_type = (*byte)(zend.Emalloc((*len_) + 1))
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
	default_header.SetHeader(GetDefaultContentType(b.SizeOf("\"Content-type: \"")-1, &len_))
	default_header.SetHeaderLen(len_)
	memcpy(default_header.GetHeader(), "Content-type: ", b.SizeOf("\"Content-type: \"")-1)
}
func SapiApplyDefaultCharset(mimetype **byte, len_ int) int {
	var charset *byte
	var newtype *byte
	var newlen int
	if SG(default_charset) {
		charset = SG(default_charset)
	} else {
		charset = SAPI_DEFAULT_CHARSET
	}
	if (*mimetype) != nil {
		if (*charset) && strncmp(*mimetype, "text/", 5) == 0 && strstr(*mimetype, "charset=") == nil {
			newlen = len_ + (b.SizeOf("\";charset=\"") - 1) + strlen(charset)
			newtype = zend.Emalloc(newlen + 1)
			PHP_STRLCPY(newtype, *mimetype, newlen+1, len_)
			strlcat(newtype, ";charset=", newlen+1)
			strlcat(newtype, charset, newlen+1)
			zend.Efree(*mimetype)
			*mimetype = newtype
			return newlen
		}
	}
	return 0
}
func SapiActivateHeadersOnly() {
	if SG(request_info).headers_read == 1 {
		return
	}
	SG(request_info).headers_read = 1
	zend.ZendLlistInit(&SG(sapi_headers).headers, b.SizeOf("sapi_header_struct"), (func(any))(SapiFreeHeader), 0)
	SG(sapi_headers).send_default_content_type = 1

	/* SG(sapi_headers).http_response_code = 200; */

	SG(sapi_headers).http_status_line = nil
	SG(sapi_headers).mimetype = nil
	SG(read_post_bytes) = 0
	SG(request_info).request_body = nil
	SG(request_info).current_user = nil
	SG(request_info).current_user_length = 0
	SG(request_info).no_headers = 0
	SG(request_info).post_entry = nil
	SG(global_request_time) = 0

	/*
	 * It's possible to override this general case in the activate() callback,
	 * if necessary.
	 */

	if SG(request_info).request_method && !(strcmp(SG(request_info).request_method, "HEAD")) {
		SG(request_info).headers_only = 1
	} else {
		SG(request_info).headers_only = 0
	}
	if SG(server_context) {
		SG(request_info).cookie_data = sapi_module.GetReadCookies()()
		if sapi_module.GetActivate() != nil {
			sapi_module.GetActivate()()
		}
	}
	if sapi_module.GetInputFilterInit() != nil {
		sapi_module.GetInputFilterInit()()
	}
}
func SapiActivate() {
	zend.ZendLlistInit(&SG(sapi_headers).headers, b.SizeOf("sapi_header_struct"), (func(any))(SapiFreeHeader), 0)
	SG(sapi_headers).send_default_content_type = 1

	/*
	   SG(sapi_headers).http_response_code = 200;
	*/

	SG(sapi_headers).http_status_line = nil
	SG(sapi_headers).mimetype = nil
	SG(headers_sent) = 0
	zend.ZVAL_UNDEF(&SG(callback_func))
	SG(read_post_bytes) = 0
	SG(request_info).request_body = nil
	SG(request_info).current_user = nil
	SG(request_info).current_user_length = 0
	SG(request_info).no_headers = 0
	SG(request_info).post_entry = nil
	SG(request_info).proto_num = 1000
	SG(global_request_time) = 0
	SG(post_read) = 0

	/* It's possible to override this general case in the activate() callback, if necessary. */

	if SG(request_info).request_method && !(strcmp(SG(request_info).request_method, "HEAD")) {
		SG(request_info).headers_only = 1
	} else {
		SG(request_info).headers_only = 0
	}
	SG(rfc1867_uploaded_files) = nil

	/* Handle request method */

	if SG(server_context) {
		if PG(enable_post_data_reading) && SG(request_info).content_type && SG(request_info).request_method && !(strcmp(SG(request_info).request_method, "POST")) {

			/* HTTP POST may contain form data to be processed into variables
			 * depending on given content type */

			SapiReadPostData()

			/* HTTP POST may contain form data to be processed into variables
			 * depending on given content type */

		} else {
			SG(request_info).content_type_dup = nil
		}

		/* Cookies */

		SG(request_info).cookie_data = sapi_module.GetReadCookies()()

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
	if SG(sapi_headers).http_status_line {
		zend.Efree(SG(sapi_headers).http_status_line)
		SG(sapi_headers).http_status_line = nil
	}
}
func SapiDeactivate() {
	zend.ZendLlistDestroy(&SG(sapi_headers).headers)
	if SG(request_info).request_body {
		SG(request_info).request_body = nil
	} else if SG(server_context) {
		if !(SG(post_read)) {

			/* make sure we've consumed all request input data */

			var dummy []byte
			var read_bytes int
			for {
				read_bytes = SapiReadPostBlock(dummy, SAPI_POST_BLOCK_SIZE)
				if SAPI_POST_BLOCK_SIZE != read_bytes {
					break
				}
			}
		}
	}
	if SG(request_info).auth_user {
		zend.Efree(SG(request_info).auth_user)
	}
	if SG(request_info).auth_password {
		zend.Efree(SG(request_info).auth_password)
	}
	if SG(request_info).auth_digest {
		zend.Efree(SG(request_info).auth_digest)
	}
	if SG(request_info).content_type_dup {
		zend.Efree(SG(request_info).content_type_dup)
	}
	if SG(request_info).current_user {
		zend.Efree(SG(request_info).current_user)
	}
	if sapi_module.GetDeactivate() != nil {
		sapi_module.GetDeactivate()()
	}
	if SG(rfc1867_uploaded_files) {
		DestroyUploadedFilesHash()
	}
	if SG(sapi_headers).mimetype {
		zend.Efree(SG(sapi_headers).mimetype)
		SG(sapi_headers).mimetype = nil
	}
	SapiSendHeadersFree()
	SG(sapi_started) = 0
	SG(headers_sent) = 0
	SG(request_info).headers_read = 0
	SG(global_request_time) = 0
}
func SapiInitializeEmptyRequest() {
	SG(server_context) = nil
	SG(request_info).request_method = nil
	SG(request_info).auth_password = nil
	SG(request_info).auth_user = SG(request_info).auth_password
	SG(request_info).auth_digest = SG(request_info).auth_user
	SG(request_info).content_type_dup = nil
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

	if SG(sapi_headers).http_response_code == ncode {
		return
	}
	if SG(sapi_headers).http_status_line {
		zend.Efree(SG(sapi_headers).http_status_line)
		SG(sapi_headers).http_status_line = nil
	}
	SG(sapi_headers).http_response_code = ncode
}
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
			zend.Efree(current)
			l.count--
		}
		current = next
	}
}
func SapiAddHeaderEx(header_line *byte, header_line_len int, duplicate zend.ZendBool, replace zend.ZendBool) int {
	var ctr SapiHeaderLine = SapiHeaderLine{0}
	var r int
	ctr.SetLine(header_line)
	ctr.SetLineLen(header_line_len)
	r = SapiHeaderOp(b.Cond(replace != 0, SAPI_HEADER_REPLACE, SAPI_HEADER_ADD), &ctr)
	if duplicate == 0 {
		zend.Efree(header_line)
	}
	return r
}
func SapiHeaderAddOp(op SapiHeaderOpEnum, sapi_header *SapiHeader) {
	if sapi_module.GetHeaderHandler() == nil || (SAPI_HEADER_ADD&sapi_module.GetHeaderHandler()(sapi_header, op, &SG(sapi_headers))) != 0 {
		if op == SAPI_HEADER_REPLACE {
			var colon_offset *byte = strchr(sapi_header.GetHeader(), ':')
			if colon_offset != nil {
				var sav byte = *colon_offset
				*colon_offset = 0
				SapiRemoveHeader(&SG(sapi_headers).headers, sapi_header.GetHeader(), strlen(sapi_header.GetHeader()))
				*colon_offset = sav
			}
		}
		zend.ZendLlistAddElement(&SG(sapi_headers).headers, any(sapi_header))
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
	if SG(headers_sent) && !(SG(request_info).no_headers) {
		var output_start_filename *byte = PhpOutputGetStartFilename()
		var output_start_lineno int = PhpOutputGetStartLineno()
		if output_start_filename != nil {
			sapi_module.GetSapiError()(zend.E_WARNING, "Cannot modify header information - headers already sent by (output started at %s:%d)", output_start_filename, output_start_lineno)
		} else {
			sapi_module.GetSapiError()(zend.E_WARNING, "Cannot modify header information - headers already sent")
		}
		return zend.FAILURE
	}
	switch op {
	case SAPI_HEADER_SET_STATUS:
		SapiUpdateResponseCode(int(zend.ZendIntptrT(arg)))
		return zend.SUCCESS
	case SAPI_HEADER_ADD:

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
			sapi_module.GetHeaderHandler()(&sapi_header, op, &SG(sapi_headers))
		}
		zend.ZendLlistClean(&SG(sapi_headers).headers)
		return zend.SUCCESS
	default:
		return zend.FAILURE
	}
	header_line = zend.Estrndup(header_line, header_line_len)

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
			zend.Efree(header_line)
			sapi_module.GetSapiError()(zend.E_WARNING, "Header to delete may not contain colon.")
			return zend.FAILURE
		}
		if sapi_module.GetHeaderHandler() != nil {
			sapi_header.SetHeader(header_line)
			sapi_header.SetHeaderLen(header_line_len)
			sapi_module.GetHeaderHandler()(&sapi_header, op, &SG(sapi_headers))
		}
		SapiRemoveHeader(&SG(sapi_headers).headers, header_line, header_line_len)
		zend.Efree(header_line)
		return zend.SUCCESS
	} else {

		/* new line/NUL character safety check */

		var i uint32
		for i = 0; i < header_line_len; i++ {

			/* RFC 7230 ch. 3.2.4 deprecates folding support */

			if header_line[i] == '\n' || header_line[i] == '\r' {
				zend.Efree(header_line)
				sapi_module.GetSapiError()(zend.E_WARNING, "Header may not contain "+"more than a single header, new line detected")
				return zend.FAILURE
			}
			if header_line[i] == '0' {
				zend.Efree(header_line)
				sapi_module.GetSapiError()(zend.E_WARNING, "Header may not contain NUL bytes")
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

		if SG(sapi_headers).http_status_line {
			zend.Efree(SG(sapi_headers).http_status_line)
		}
		SG(sapi_headers).http_status_line = header_line
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

				if !(strncmp(ptr, "image/", b.SizeOf("\"image/\"")-1)) {
					var key *zend.ZendString = zend.ZendStringInit("zlib.output_compression", b.SizeOf("\"zlib.output_compression\"")-1, 0)
					zend.ZendAlterIniEntryChars(key, "0", b.SizeOf("\"0\"")-1, PHP_INI_USER, PHP_INI_STAGE_RUNTIME)
					zend.ZendStringReleaseEx(key, 0)
				}
				mimetype = zend.Estrdup(ptr)
				newlen = SapiApplyDefaultCharset(&mimetype, len_)
				if !(SG(sapi_headers).mimetype) {
					SG(sapi_headers).mimetype = zend.Estrdup(mimetype)
				}
				if newlen != 0 {
					newlen += b.SizeOf("\"Content-type: \"")
					newheader = zend.Emalloc(newlen)
					PHP_STRLCPY(newheader, "Content-type: ", newlen, b.SizeOf("\"Content-type: \"")-1)
					strlcat(newheader, mimetype, newlen)
					sapi_header.SetHeader(newheader)
					sapi_header.SetHeaderLen(uint32_t(newlen - 1))
					zend.Efree(header_line)
				}
				zend.Efree(mimetype)
				SG(sapi_headers).send_default_content_type = 0
			} else if !(strcasecmp(header_line, "Content-Length")) {

				/* Script is setting Content-length. The script cannot reasonably
				 * know the size of the message body after compression, so it's best
				 * do disable compression altogether. This contributes to making scripts
				 * portable between setups that have and don't have zlib compression
				 * enabled globally. See req #44164 */

				var key *zend.ZendString = zend.ZendStringInit("zlib.output_compression", b.SizeOf("\"zlib.output_compression\"")-1, 0)
				zend.ZendAlterIniEntryChars(key, "0", b.SizeOf("\"0\"")-1, PHP_INI_USER, PHP_INI_STAGE_RUNTIME)
				zend.ZendStringReleaseEx(key, 0)
			} else if !(strcasecmp(header_line, "Location")) {
				if (SG(sapi_headers).http_response_code < 300 || SG(sapi_headers).http_response_code > 399) && SG(sapi_headers).http_response_code != 201 {

					/* Return a Found Redirect if one is not already specified */

					if http_response_code != 0 {
						SapiUpdateResponseCode(http_response_code)
					} else if SG(request_info).proto_num > 1000 && SG(request_info).request_method && strcmp(SG(request_info).request_method, "HEAD") && strcmp(SG(request_info).request_method, "GET") {
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
	if SG(headers_sent) || SG(request_info).no_headers {
		return zend.SUCCESS
	}

	/* Success-oriented.  We set headers_sent to 1 here to avoid an infinite loop
	 * in case of an error situation.
	 */

	if SG(sapi_headers).send_default_content_type && sapi_module.GetSendHeaders() != nil {
		var len_ uint32 = 0
		var default_mimetype *byte = GetDefaultContentType(0, &len_)
		if default_mimetype != nil && len_ != 0 {
			var default_header SapiHeader
			SG(sapi_headers).mimetype = default_mimetype
			default_header.SetHeaderLen(b.SizeOf("\"Content-type: \"") - 1 + len_)
			default_header.SetHeader(zend.Emalloc(default_header.GetHeaderLen() + 1))
			memcpy(default_header.GetHeader(), "Content-type: ", b.SizeOf("\"Content-type: \"")-1)
			memcpy(default_header.GetHeader()+b.SizeOf("\"Content-type: \"")-1, SG(sapi_headers).mimetype, len_+1)
			SapiHeaderAddOp(SAPI_HEADER_ADD, &default_header)
		} else {
			zend.Efree(default_mimetype)
		}
		SG(sapi_headers).send_default_content_type = 0
	}
	if zend.Z_TYPE(SG(callback_func)) != zend.IS_UNDEF {
		var cb zend.Zval
		zend.ZVAL_COPY_VALUE(&cb, &SG(callback_func))
		zend.ZVAL_UNDEF(&SG(callback_func))
		SapiRunHeaderCallback(&cb)
		zend.ZvalPtrDtor(&cb)
	}
	SG(headers_sent) = 1
	if sapi_module.GetSendHeaders() != nil {
		retval = sapi_module.GetSendHeaders()(&SG(sapi_headers))
	} else {
		retval = SAPI_HEADER_DO_SEND
	}
	switch retval {
	case SAPI_HEADER_SENT_SUCCESSFULLY:
		ret = zend.SUCCESS
		break
	case SAPI_HEADER_DO_SEND:
		var http_status_line SapiHeader
		var buf []byte
		if SG(sapi_headers).http_status_line {
			http_status_line.SetHeader(SG(sapi_headers).http_status_line)
			http_status_line.SetHeaderLen(uint32(strlen(SG(sapi_headers).http_status_line)))
		} else {
			http_status_line.SetHeader(buf)
			http_status_line.SetHeaderLen(Slprintf(buf, b.SizeOf("buf"), "HTTP/1.0 %d X", SG(sapi_headers).http_response_code))
		}
		sapi_module.GetSendHeader()(&http_status_line, SG(server_context))
		zend.ZendLlistApplyWithArgument(&SG(sapi_headers).headers, zend.LlistApplyWithArgFuncT(sapi_module.GetSendHeader()), SG(server_context))
		if SG(sapi_headers).send_default_content_type {
			var default_header SapiHeader
			SapiGetDefaultContentTypeHeader(&default_header)
			sapi_module.GetSendHeader()(&default_header, SG(server_context))
			SapiFreeHeader(&default_header)
		}
		sapi_module.GetSendHeader()(nil, SG(server_context))
		ret = zend.SUCCESS
		break
	case SAPI_HEADER_SEND_FAILED:
		SG(headers_sent) = 0
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
	if SG(sapi_started) && zend.ExecutorGlobals.current_execute_data != nil {
		return zend.FAILURE
	}
	key = zend.ZendStringInit(post_entry.GetContentType(), post_entry.GetContentTypeLen(), 1)
	zend.GC_MAKE_PERSISTENT_LOCAL(key)
	if zend.ZendHashAddMem(&SG(known_post_content_types), key, any(post_entry), b.SizeOf("sapi_post_entry")) {
		ret = zend.SUCCESS
	} else {
		ret = zend.FAILURE
	}
	zend.ZendStringReleaseEx(key, 1)
	return ret
}
func SapiUnregisterPostEntry(post_entry *SapiPostEntry) {
	if SG(sapi_started) && zend.ExecutorGlobals.current_execute_data != nil {
		return
	}
	zend.ZendHashStrDel(&SG(known_post_content_types), post_entry.GetContentType(), post_entry.GetContentTypeLen())
}
func SapiRegisterDefaultPostReader(default_post_reader func()) int {
	if SG(sapi_started) && zend.ExecutorGlobals.current_execute_data != nil {
		return zend.FAILURE
	}
	sapi_module.SetDefaultPostReader(default_post_reader)
	return zend.SUCCESS
}
func SapiRegisterTreatData(treat_data func(arg int, str *byte, destArray *zend.Zval)) int {
	if SG(sapi_started) && zend.ExecutorGlobals.current_execute_data != nil {
		return zend.FAILURE
	}
	sapi_module.SetTreatData(treat_data)
	return zend.SUCCESS
}
func SapiRegisterInputFilter(input_filter func(arg int, var_ *byte, val **byte, val_len int, new_val_len *int) uint, input_filter_init func() uint) int {
	if SG(sapi_started) && zend.ExecutorGlobals.current_execute_data != nil {
		return zend.FAILURE
	}
	sapi_module.SetInputFilter(input_filter)
	sapi_module.SetInputFilterInit(input_filter_init)
	return zend.SUCCESS
}
func SapiFlush() int {
	if sapi_module.GetFlush() != nil {
		sapi_module.GetFlush()(SG(server_context))
		return zend.SUCCESS
	} else {
		return zend.FAILURE
	}
}
func SapiGetStat() *zend.ZendStatT {
	if sapi_module.GetGetStat() != nil {
		return sapi_module.GetGetStat()()
	} else {
		if !(SG(request_info).path_translated) || zend.VCWD_STAT(SG(request_info).path_translated, &SG(global_stat)) == -1 {
			return nil
		}
		return &SG(global_stat)
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
			value = zend.Estrdup(tmp)
		} else {
			return nil
		}
		if sapi_module.GetInputFilter() != nil {
			sapi_module.GetInputFilter()(PARSE_STRING, name, &value, strlen(value), nil)
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
	if SG(global_request_time) {
		return SG(global_request_time)
	}
	if sapi_module.GetGetRequestTime() != nil && SG(server_context) {
		SG(global_request_time) = sapi_module.GetGetRequestTime()()
	} else {
		var tp __struct__timeval = __struct__timeval{0}
		if !(gettimeofday(&tp, nil)) {
			SG(global_request_time) = float64(tp.tv_sec + tp.tv_usec/1000000.0)
		} else {
			SG(global_request_time) = float64(time(0))
		}
	}
	return SG(global_request_time)
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
		str = zend.DoAlloca(var_len+1, use_heap)
		var_ = str
		*p++
		b.PostInc(&(*str)) = (*p) - 1
		for *p {
			if (*p) == '_' {
				b.PostInc(&(*str)) = '-'
				p++
				if *p {
					*p++
					b.PostInc(&(*str)) = (*p) - 1
				}
			} else if (*p) >= 'A' && (*p) <= 'Z' {
				b.PostInc(&(*str)) = b.PostInc(&(*p)) - 'A' + 'a'
			} else {
				*p++
				b.PostInc(&(*str)) = (*p) - 1
			}
		}
		*str = 0
	} else if var_len == b.SizeOf("\"CONTENT_TYPE\"")-1 && memcmp(var_, "CONTENT_TYPE", b.SizeOf("\"CONTENT_TYPE\"")-1) == 0 {
		var_ = "Content-Type"
	} else if var_len == b.SizeOf("\"CONTENT_LENGTH\"")-1 && memcmp(var_, "CONTENT_LENGTH", b.SizeOf("\"CONTENT_LENGTH\"")-1) == 0 {
		var_ = "Content-Length"
	} else {
		return
	}
	zend.AddAssocStringlEx(return_value, var_, var_len, val, val_len)
	if str != nil {
		zend.FreeAlloca(var_, use_heap)
	}
}
