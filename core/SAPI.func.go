package core

import (
	"fmt"
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/kits/ascii"
	"github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
	"strings"
)

func SG__() *SapiGlobals     { return App().SG() }
func SM__() ISapiModule      { return App().sapiModule }
func SetSM__(sf ISapiModule) { App().sapiModule = sf }
func SapiAddHeader(str string) int {
	ctr := MakeSapiHeaderLineEx(str)
	return SapiHeaderOp(SAPI_HEADER_REPLACE, &ctr)
}

//@alias -old
func ZifHeaderRegisterCallback(callback *types.Zval) bool {
	if !zend.IsCallable(callback, nil, 0) {
		return false
	}

	if SG__().callbackFunc.IsNotUndef() {
		SG__().fciCache = types.EmptyFcallInfoCache()
	}

	types.ZVAL_COPY(&SG__().callbackFunc, callback)

	return true
}

func SapiRunHeaderCallback(callback *types.Zval) {
	var error int
	var fci types.ZendFcallInfo
	var callback_error *byte = nil
	var retval types.Zval
	if zend.ZendFcallInfoInit(callback, 0, &fci, &(SG__().fciCache), nil, &callback_error) == types.SUCCESS {
		fci.SetRetval(&retval)
		error = zend.ZendCallFunction(&fci, &(SG__().fciCache))
		if error == types.FAILURE {
			goto callback_failed
		} else {
			// zend.ZvalPtrDtor(&retval)
		}
	} else {
	callback_failed:
		PhpErrorDocref("", faults.E_WARNING, "Could not call the sapi_header_callback")
	}
	if callback_error != nil {
		zend.Efree(callback_error)
	}
}
func SapiHandlePost(arg any) {
	if SG__().RequestInfo.postEntry != nil && SG__().RequestInfo.ContentTypeDup() != "" {
		SG__().RequestInfo.postEntry.PostHandler(SG__().RequestInfo.ContentTypeDup(), arg)
		zend.Efree(SG__().RequestInfo.contentTypeDup)
		SG__().RequestInfo.SetContentTypeDup("")
	}
}
func SapiReadPostData() {
	var post_entry *SapiPostEntry
	var content_type_length uint32 = uint32(strlen(SG__().RequestInfo.contentType))
	var content_type *byte = zend.Estrndup(SG__().RequestInfo.contentType, content_type_length)
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
			fallthrough
		case ',':
			fallthrough
		case ' ':
			content_type_length = p - content_type
			oldchar = *p
			*p = 0
		default:
			*p = tolower(*p)
		}
	}

	/* now try to find an appropriate POST content handler */

	if lang.Assign(&post_entry, types.ZendHashStrFindPtr(&(SG__().knownPostContentTypes), b.CastStr(content_type, content_type_length))) != nil {

		/* found one, register it for use */
		SG__().RequestInfo.postEntry = post_entry
		post_reader_func = post_entry.PostReader
	} else {

		/* fallback */

		SG__().RequestInfo.postEntry = nil
		if SM__().GetDefaultPostReader() == nil {

			/* no default reader ? */

			SG__().RequestInfo.contentTypeDup = nil
			SM__().SapiError(faults.E_WARNING, "Unsupported content type:  '%s'", content_type)
			return
		}
	}
	if oldchar {
		*(p - 1) = oldchar
	}
	SG__().RequestInfo.contentTypeDup = content_type
	if post_reader_func != nil {
		post_reader_func()
	}
	if SM__().GetDefaultPostReader() != nil {
		SM__().GetDefaultPostReader()()
	}
}
func SapiReadPostBlock(buffer *byte, buflen int) int {
	var read_bytes int
	if SM__().GetReadPost() == nil {
		return 0
	}
	read_bytes = SM__().GetReadPost()(buffer, buflen)
	if read_bytes > 0 {

		/* gogo */

		SG__().readPostBytes += read_bytes

		/* gogo */

	}
	if read_bytes < buflen {

		/* done */

		SG__().postRead = 1

		/* done */

	}
	return read_bytes
}
func SapiReadStandardFormData() {
	if SG__().postMaxSize > 0 && SG__().RequestInfo.ContentLength() > SG__().postMaxSize {
		PhpErrorDocref("", faults.E_WARNING, "POST Content-Length of %d bytes exceeds the limit of %d bytes", SG__().RequestInfo.ContentLength(), SG__().postMaxSize)
		return
	}
	SG__().RequestInfo.requestBody = PhpStreamTempCreateEx(TEMP_STREAM_DEFAULT, SAPI_POST_BLOCK_SIZE, PG__().upload_tmp_dir)
	if SM__().GetReadPost() != nil {
		var read_bytes int
		for {
			var buffer []byte
			read_bytes = SapiReadPostBlock(buffer, SAPI_POST_BLOCK_SIZE)
			if read_bytes > 0 {
				if PhpStreamWrite(SG__().RequestInfo.requestBody, buffer, read_bytes) != read_bytes {

					/* if parts of the stream can't be written, purge it completely */

					PhpStreamTruncateSetSize(SG__().RequestInfo.requestBody, 0)
					PhpErrorDocref("", faults.E_WARNING, "POST data can't be buffered; all data discarded")
					break
				}
			}
			if SG__().postMaxSize > 0 && SG__().readPostBytes > SG__().postMaxSize {
				PhpErrorDocref("", faults.E_WARNING, "Actual POST length does not match Content-Length, and exceeds "+zend.ZEND_LONG_FMT+" bytes", SG__().postMaxSize)
				break
			}
			if read_bytes < SAPI_POST_BLOCK_SIZE {

				/* done */

				break

				/* done */

			}
		}
		PhpStreamRewind(SG__().RequestInfo.requestBody)
	}
}
func GetDefaultContentType(prefix string) string {
	var mimetype string
	var charset string
	if SG__().DefaultMimetype() != "" {
		mimetype = SG__().DefaultMimetype()
	} else {
		mimetype = SAPI_DEFAULT_MIMETYPE
	}

	if SG__().DefaultCharset() != "" {
		charset = SG__().DefaultCharset()
	} else {
		charset = SAPI_DEFAULT_CHARSET
	}

	if charset != "" && len(mimetype) >= 5 && ascii.StrToLower(mimetype) == "text/" {
		return prefix + mimetype + "; charset=" + charset
	} else {
		return prefix + mimetype
	}
}
func SapiGetDefaultContentTypeHeader() *SapiHeader {
	return NewSapiHeader(GetDefaultContentType("Content-type: "))
}
func SapiApplyDefaultCharset(mimetype **byte, len_ int) int {
	var charset *byte
	var newtype *byte
	var newlen int
	if SG__().defaultCharset {
		charset = SG__().defaultCharset
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
func SapiActivate() {
	SG__().Activate()

	/* Handle request method */
	if SG__().serverContext != nil {
		if PG__().enable_post_data_reading && SG__().RequestInfo.ContentType() != "" && SG__().RequestInfo.IsRequestMethod("POST") {
			/* HTTP POST may contain form data to be processed into variables
			 * depending on given content type */
			SapiReadPostData()
		} else {
			SG__().RequestInfo.contentTypeDup = nil
		}

		/* Cookies */
		SG__().RequestInfo.SetCookieData(SM__().GetReadCookies()())
	}
	SM__().Activate()
	SM__().InputFilterInit()
}
func SapiSendHeadersFree() {
	if SG__().SapiHeaders().httpStatusLine {
		zend.Efree(SG__().SapiHeaders().httpStatusLine)
		SG__().SapiHeaders().httpStatusLine = nil
	}
}
func SapiDeactivate() {
	SG__().SapiHeaders().headers.Clean()
	if SG__().RequestInfo.requestBody {
		SG__().RequestInfo.requestBody = nil
	} else if SG__().serverContext {
		if !(SG__().postRead) {

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
	if SG__().RequestInfo.authUser {
		zend.Efree(SG__().RequestInfo.authUser)
	}
	if SG__().RequestInfo.authPassword {
		zend.Efree(SG__().RequestInfo.authPassword)
	}
	if SG__().RequestInfo.authDigest {
		zend.Efree(SG__().RequestInfo.authDigest)
	}
	if SG__().RequestInfo.contentTypeDup {
		zend.Efree(SG__().RequestInfo.contentTypeDup)
	}
	SM__().Deactivate()
	if SG__().rfc1867UploadedFiles != nil {
		DestroyUploadedFilesHash()
	}
	if SG__().SapiHeaders().mimetype {
		zend.Efree(SG__().SapiHeaders().mimetype)
		SG__().SapiHeaders().mimetype = nil
	}
	SapiSendHeadersFree()
	SG__().sapiStarted = 0
	SG__().headersSent = false
	SG__().RequestInfo.headersRead = 0
	SG__().globalRequestTime = 0
}
func SapiInitializeEmptyRequest() {
	SG__().serverContext = nil
	SG__().RequestInfo.InitEmpty()
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

	if SG__().SapiHeaders().HttpResponseCode() == ncode {
		return
	}
	if SG__().SapiHeaders().httpStatusLine {
		zend.Efree(SG__().SapiHeaders().httpStatusLine)
		SG__().SapiHeaders().httpStatusLine = nil
	}
	SG__().SapiHeaders().httpResponseCode = ncode
}
func SapiHeaderAddOp(op SapiHeaderOpEnum, sapi_header *SapiHeader) {
	result := SM__().HeaderHandler(sapi_header, op, &(SG__().SapiHeaders()))
	if (SAPI_HEADER_ADD & result) != 0 {
		if op == SAPI_HEADER_REPLACE {
			if key, ok := sapi_header.GetKey(); ok {
				SG__().SapiHeaders().RemoveHeaderByKey(key)
			}
		}
		SG__().SapiHeaders().AddHeader(sapi_header)
	} else {
		//SapiFreeHeader(sapi_header)
	}
}
func SapiHeaderOp(op SapiHeaderOpEnum, arg any) int {
	var sapi_header SapiHeader
	var colon_offset *byte
	var header_line *byte
	var header_line_len int
	var http_response_code int
	if SG__().headersSent && !(SG__().RequestInfo.noHeaders) {
		var outputStartFilename = OG__().StartFilename()
		var outputStartLineno = OG__().StartLineno()
		if outputStartFilename != "" {
			SM__().SapiError(faults.E_WARNING, "Cannot modify header information - headers already sent by (output started at %s:%d)", outputStartFilename, outputStartLineno)
		} else {
			SM__().SapiError(faults.E_WARNING, "Cannot modify header information - headers already sent")
		}
		return types.FAILURE
	}
	switch op {
	case SAPI_HEADER_SET_STATUS:
		SapiUpdateResponseCode(int(types.ZendIntptrT(arg)))
		return types.SUCCESS
	case SAPI_HEADER_ADD:
		fallthrough
	case SAPI_HEADER_REPLACE:
		fallthrough
	case SAPI_HEADER_DELETE:
		var p *SapiHeaderLine = arg
		if p.GetLine() == nil || p.GetLineLen() == 0 {
			return types.FAILURE
		}
		header_line = p.GetLine()
		header_line_len = p.GetLineLen()
		http_response_code = p.GetResponseCode()
	case SAPI_HEADER_DELETE_ALL:
		SM__().HeaderHandler(&sapi_header, op, &(SG__().SapiHeaders()))
		SG__().SapiHeaders().headers.Clean()
		return types.SUCCESS
	default:
		return types.FAILURE
	}
	header_line = zend.Estrndup(header_line, header_line_len)
	headerLine := b.CastStrAuto(header_line)

	/* cut off trailing spaces, linefeeds and carriage-returns */
	headerLine = strings.TrimRightFunc(headerLine, ascii.IsSpaceRune)
	if op == SAPI_HEADER_DELETE {
		if strings.ContainsRune(headerLine, ':') {
			SM__().SapiError(faults.E_WARNING, "Header to delete may not contain colon.")
			return types.FAILURE
		}

		sapi_header.SetHeader(headerLine)
		SM__().HeaderHandler(&sapi_header, op, SG__().SapiHeaders())
		SG__().SapiHeaders().RemoveHeaderByKey(headerLine)
		return types.SUCCESS
	} else {
		/* new line/NUL character safety check */
		for _, c := range []byte(headerLine) {
			/* RFC 7230 ch. 3.2.4 deprecates folding support */
			if c == '\n' || c == '\r' {
				zend.Efree(header_line)
				SM__().SapiError(faults.E_WARNING, "Header may not contain more than a single header, new line detected")
				return types.FAILURE
			}
			if c == 0 {
				zend.Efree(header_line)
				SM__().SapiError(faults.E_WARNING, "Header may not contain NUL bytes")
				return types.FAILURE
			}
		}
	}
	sapi_header.SetHeader(headerLine)

	/* Check the header for a few cases that we have special support for in SAPI */
	if len(headerLine) >= 5 && ascii.StrCaseEquals(headerLine[:5], "HTTP/") {
		/* filter out the response code */
		SapiUpdateResponseCode(SapiExtractResponseCode(header_line))

		SG__().SapiHeaders().SetHttpStatusLine(headerLine)
		return types.SUCCESS
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
					var key *types.String = types.NewString("zlib.output_compression")
					zend.ZendAlterIniEntryChars(key.GetStr(), "0", PHP_INI_USER, PHP_INI_STAGE_RUNTIME)
					// types.ZendStringReleaseEx(key, 0)
				}
				mimetype = zend.Estrdup(ptr)
				newlen = SapiApplyDefaultCharset(&mimetype, len_)
				if !(SG__().SapiHeaders().mimetype) {
					SG__().SapiHeaders().mimetype = zend.Estrdup(mimetype)
				}
				if newlen != 0 {
					newlen += b.SizeOf("\"Content-type: \"")
					newheader = zend.Emalloc(newlen)
					PHP_STRLCPY(newheader, "Content-type: ", newlen, b.SizeOf("\"Content-type: \"")-1)
					strlcat(newheader, mimetype, newlen)
					sapi_header.SetHeader(b.CastStr(newheader, newlen-1))
					zend.Efree(header_line)
				}
				zend.Efree(mimetype)
				SG__().SapiHeaders().SetSendDefaultContentType(false)
			} else if !(strcasecmp(header_line, "Content-Length")) {

				/* Script is setting Content-length. The script cannot reasonably
				 * know the size of the message body after compression, so it's best
				 * do disable compression altogether. This contributes to making scripts
				 * portable between setups that have and don't have zlib compression
				 * enabled globally. See req #44164 */

				var key *types.String = types.NewString("zlib.output_compression")
				zend.ZendAlterIniEntryChars(key.GetStr(), "0", PHP_INI_USER, PHP_INI_STAGE_RUNTIME)
				// types.ZendStringReleaseEx(key, 0)
			} else if ascii.StrCaseEquals(header_line, "Location") {
				if (SG__().SapiHeaders().httpResponseCode < 300 || SG__().SapiHeaders().httpResponseCode > 399) && SG__().SapiHeaders().httpResponseCode != 201 {

					/* Return a Found Redirect if one is not already specified */

					if http_response_code != 0 {
						SapiUpdateResponseCode(http_response_code)
					} else if SG__().RequestInfo.protoNum > 1000 && !SG__().RequestInfo.IsRequestMethod("HEAD") && !SG__().RequestInfo.IsRequestMethod("GET") {
						SapiUpdateResponseCode(303)
					} else {
						SapiUpdateResponseCode(302)
					}
				}
			} else if ascii.StrCaseEquals(header_line, "WWW-Authenticate") {
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
	return types.SUCCESS
}
func SapiSendHeaders() int {
	var retval int
	var ret int = types.FAILURE
	if SG__().headersSent || SG__().RequestInfo.noHeaders {
		return types.SUCCESS
	}

	/* Success-oriented.  We set headers_sent to 1 here to avoid an infinite loop
	 * in case of an error situation.
	 */

	if SG__().SapiHeaders().SendDefaultContentType() && SM__().GetSendHeaders() != nil {
		var defaultMimetype = GetDefaultContentType("")
		var defaultHeader = NewSapiHeader("Content-type: " + defaultMimetype)
		SG__().SapiHeaders().mimetype = defaultMimetype
		SapiHeaderAddOp(SAPI_HEADER_ADD, defaultHeader)
		SG__().SapiHeaders().SetSendDefaultContentType(false)
	}
	if SG__().callbackFunc.IsNotUndef() {
		var cb types.Zval
		types.ZVAL_COPY_VALUE(&cb, &(SG__().callbackFunc))
		SG__().callbackFunc.SetUndef()
		SapiRunHeaderCallback(&cb)
		// zend.ZvalPtrDtor(&cb)
	}
	SG__().headersSent = true
	if SM__().GetSendHeaders() != nil {
		retval = SM__().GetSendHeaders()(&(SG__().SapiHeaders()))
	} else {
		retval = SAPI_HEADER_DO_SEND
	}
	switch retval {
	case SAPI_HEADER_SENT_SUCCESSFULLY:
		ret = types.SUCCESS
	case SAPI_HEADER_DO_SEND:
		var http_status_line SapiHeader
		if SG__().SapiHeaders().HttpStatusLine() != "" {
			http_status_line.SetHeader(SG__().SapiHeaders().httpStatusLine)
		} else {
			http_status_line.SetHeader(fmt.Sprintf("HTTP/1.0 %d X", SG__().SapiHeaders().HttpResponseCode()))
		}
		SM__().GetSendHeader()(&http_status_line, SG__().serverContext)
		SG__().SapiHeaders().GetHeaders().Each(func(h *SapiHeader) {
			SM__().GetSendHeader()(h, SG__().serverContext)
		})

		if SG__().SapiHeaders().SendDefaultContentType() != 0 {
			defaultHeader := SapiGetDefaultContentTypeHeader()
			SM__().GetSendHeader()(defaultHeader, SG__().serverContext)
		}
		SM__().GetSendHeader()(nil, SG__().serverContext)
		ret = types.SUCCESS
	case SAPI_HEADER_SEND_FAILED:
		SG__().headersSent = false
		ret = types.FAILURE
	}
	SapiSendHeadersFree()
	return ret
}
func SapiRegisterPostEntries(postEntries []SapiPostEntry) int {
	for i := range postEntries {
		p := &postEntries[i]
		if SapiRegisterPostEntry(p) == types.FAILURE {
			return types.FAILURE
		}
	}
	return types.SUCCESS
}
func SapiRegisterPostEntry(postEntry *SapiPostEntry) int {
	var ret int
	if SG__().sapiStarted == types.SUCCESS && zend.CurrEX() != nil {
		return types.FAILURE
	}
	if types.ZendHashAddMem(&(SG__().knownPostContentTypes), postEntry.ContentType(), any(postEntry), b.SizeOf("sapi_post_entry")) {
		ret = types.SUCCESS
	} else {
		ret = types.FAILURE
	}
	return ret
}
func SapiRegisterDefaultPostReader(default_post_reader func()) int {
	if SG__().sapiStarted && zend.CurrEX() != nil {
		return types.FAILURE
	}
	SM__().SetDefaultPostReader(default_post_reader)
	return types.SUCCESS
}
func SapiRegisterTreatData(treat_data func(arg int, str *byte, destArray *types.Zval)) int {
	if SG__().sapiStarted && zend.CurrEX() != nil {
		return types.FAILURE
	}
	SM__().SetTreatData(treat_data)
	return types.SUCCESS
}
func SapiRegisterInputFilter(input_filter func(arg int, var_ *byte, val **byte, val_len int, new_val_len *int) uint, input_filter_init func() uint) int {
	if SG__().sapiStarted && zend.CurrEX() != nil {
		return types.FAILURE
	}
	SM__().SetInputFilter(input_filter)
	SM__().SetInputFilterInit(input_filter_init)
	return types.SUCCESS
}
func SapiFlush() {
	SM__().Flush(SG__().serverContext)
}
func SapiGetenv(name string) *string {
	if strings.EqualFold(name, "HTTP_PROXY") {
		/* Ugly fix for HTTP_PROXY issue, see bug #72573 */
		return nil
	}
	if value, ok := SM__().GetEnv(name); ok {
		value = SM__().InputFilter(PARSE_STRING, name, value)
		return &value
	}

	return nil
}
func SapiGetRequestTime() float64 {
	if SG__().globalRequestTime {
		return SG__().globalRequestTime
	}
	var tp __struct__timeval = __struct__timeval{0}
	if !(gettimeofday(&tp, nil)) {
		SG__().globalRequestTime = float64(tp.tv_sec + tp.tv_usec/1000000.0)
	} else {
		SG__().globalRequestTime = float64(time(0))
	}
	return SG__().globalRequestTime
}
func SapiAddRequestHeader(var_ *byte, var_len uint, val *byte, val_len uint, arg any) {
	var return_value *types.Zval = (*types.Zval)(arg)
	var str *byte = nil
	if var_len > 5 && var_[0] == 'H' && var_[1] == 'T' && var_[2] == 'T' && var_[3] == 'P' && var_[4] == '_' {
		var p *byte
		var_len -= 5
		p = var_ + 5
		str = zend.DoAlloca(var_len+1, use_heap)
		var_ = str
		*p++
		lang.PostInc(&(*str)) = (*p) - 1
		for *p {
			if (*p) == '_' {
				lang.PostInc(&(*str)) = '-'
				p++
				if *p {
					*p++
					lang.PostInc(&(*str)) = (*p) - 1
				}
			} else if (*p) >= 'A' && (*p) <= 'Z' {
				lang.PostInc(&(*str)) = lang.PostInc(&(*p)) - 'A' + 'a'
			} else {
				*p++
				lang.PostInc(&(*str)) = (*p) - 1
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
	zend.AddAssocStringlEx(return_value, b.CastStr(var_, var_len), b.CastStr(val, val_len))
	if str != nil {
		zend.FreeAlloca(var_, use_heap)
	}
}
