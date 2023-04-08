package core

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/ext/standard"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/types"
)

func DummyEncodingTranslation() int { return 0 }
func PhpRfc1867RegisterConstants() {
	zend.RegisterMainLongConstant("UPLOAD_ERR_OK", UPLOAD_ERROR_OK, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.RegisterMainLongConstant("UPLOAD_ERR_INI_SIZE", UPLOAD_ERROR_A, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.RegisterMainLongConstant("UPLOAD_ERR_FORM_SIZE", UPLOAD_ERROR_B, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.RegisterMainLongConstant("UPLOAD_ERR_PARTIAL", UPLOAD_ERROR_C, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.RegisterMainLongConstant("UPLOAD_ERR_NO_FILE", UPLOAD_ERROR_D, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.RegisterMainLongConstant("UPLOAD_ERR_NO_TMP_DIR", UPLOAD_ERROR_E, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.RegisterMainLongConstant("UPLOAD_ERR_CANT_WRITE", UPLOAD_ERROR_F, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.RegisterMainLongConstant("UPLOAD_ERR_EXTENSION", UPLOAD_ERROR_X, zend.CONST_CS|zend.CONST_PERSISTENT)
}
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
			fallthrough
		case '.':
			*p = '_'
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
func AddProtectedVariable(varname *byte) {
	NormalizeProtectedVariable(varname)
	types.ZendHashStrAddEmptyElement(&(PG__().rfc1867_protected_variables), varname)
}
func IsProtectedVariable(varname *byte) types.ZendBool {
	NormalizeProtectedVariable(varname)
	return types.IntBool(&(PG__().rfc1867_protected_variables).KeyExists(varname))
}
func SafePhpRegisterVariable(var_ *byte, strval *byte, val_len int, track_vars_array *types.Zval, override_protection types.ZendBool) {
	if override_protection != 0 || IsProtectedVariable(var_) == 0 {
		PhpRegisterVariableSafe(var_, strval, val_len, track_vars_array)
	}
}
func SafePhpRegisterVariableEx(var_ *byte, val *types.Zval, track_vars_array *types.Zval, override_protection types.ZendBool) {
	if override_protection != 0 || IsProtectedVariable(var_) == 0 {
		PhpRegisterVariableEx(var_, val, track_vars_array)
	}
}
func RegisterHttpPostFilesVariable(strvar *byte, val *byte, http_post_files *types.Zval, override_protection types.ZendBool) {
	SafePhpRegisterVariable(strvar, val, strlen(val), http_post_files, override_protection)
}
func RegisterHttpPostFilesVariableEx(var_ *byte, val *types.Zval, http_post_files *types.Zval, override_protection types.ZendBool) {
	SafePhpRegisterVariableEx(var_, val, http_post_files, override_protection)
}
func DestroyUploadedFilesHash() {
	for filename, _ := range SG__().rfc1867_uploaded_files {
		zend.VCWD_UNLINK(filename)
	}
	SG__().rfc1867_uploaded_files = nil
}
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
		actual_read = int(SM__().GetReadPost()(buf, bytes_to_read))

		/* update the buffer length */

		if actual_read > 0 {
			self.SetBytesInBuffer(self.GetBytesInBuffer() + actual_read)
			SG__().read_post_bytes += actual_read
			total_read += actual_read
			bytes_to_read -= actual_read
		} else {
			break
		}

		/* update the buffer length */

	}
	return total_read
}
func MultipartBufferEof(self *MultipartBuffer) int {
	return self.GetBytesInBuffer() == 0 && FillBuffer(self) < 1
}
func MultipartBufferNew(boundary *byte, boundary_len int) *MultipartBuffer {
	var self *MultipartBuffer = (*MultipartBuffer)(zend.Ecalloc(1, b.SizeOf("multipart_buffer")))
	var minsize int = boundary_len + 6
	if minsize < FILLUNIT {
		minsize = FILLUNIT
	}
	self.SetBuffer((*byte)(zend.Ecalloc(1, minsize+1)))
	self.SetBufsize(minsize)
	Spprintf(self.GetBoundary(), 0, "--%s", boundary)
	self.SetBoundaryNextLen(int(Spprintf(self.GetBoundaryNext(), 0, "\n--%s", boundary)))
	self.SetBufBegin(self.GetBuffer())
	self.SetBytesInBuffer(0)
	if PhpRfc1867EncodingTranslation() != 0 {
		PhpRfc1867GetDetectOrder(self.GetDetectOrder(), self.GetDetectOrderSize())
	} else {
		self.SetDetectOrder(nil)
		self.SetDetectOrderSize(0)
	}
	self.SetInputEncoding(nil)
	return self
}
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
func GetLine(self *MultipartBuffer) *byte {
	var ptr *byte = NextLine(self)
	if ptr == nil {
		FillBuffer(self)
		ptr = NextLine(self)
	}
	return ptr
}
func PhpFreeHdrEntry(h *MimeHeaderEntry) {
	if h.GetKey() != nil {
		zend.Efree(h.GetKey())
	}
	if h.GetValue() != nil {
		zend.Efree(h.GetValue())
	}
}
func FindBoundary(self *MultipartBuffer, boundary *byte) int {
	var line *byte

	/* loop through lines */

	for b.Assign(&line, GetLine(self)) {

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
func MultipartBufferHeaders(self *MultipartBuffer, header *zend.ZendLlist) int {
	var line *byte
	var entry MimeHeaderEntry = MakeMimeHeaderEntry(0)
	var buf_value zend.SmartString = zend.MakeSmartString(0)
	var key *byte = nil

	/* didn't find boundary, abort */

	if FindBoundary(self, self.GetBoundary()) == 0 {
		return 0
	}

	/* get lines of text, or CRLF_CRLF */

	for b.Assign(&line, GetLine(self)) && line[0] != '0' {

		/* add header to table */

		var value *byte = nil
		/* space in the beginning means same header */

		if !(isspace(line[0])) {
			value = strchr(line, ':')
		}
		if value != nil {
			if buf_value.GetC() != nil && key != nil {

				/* new entry, add the old one to the list */

				buf_value.ZeroTail()
				entry.SetKey(key)
				entry.SetValue(buf_value.GetC())
				header.AddElement(&entry)
				buf_value.SetC(nil)
				key = nil
			}
			*value = '0'
			for {
				value++
				if !(isspace(*value)) {
					break
				}
			}
			key = zend.Estrdup(line)
			buf_value.AppendString(b.CastStrAuto(value))
		} else if buf_value.GetC() != nil {
			buf_value.AppendString(b.CastStrAuto(line))
		} else {
			continue
		}
	}
	if buf_value.GetC() != nil && key != nil {

		/* add the last one to the list */

		buf_value.ZeroTail()
		entry.SetKey(key)
		entry.SetValue(buf_value.GetC())
		header.AddElement(&entry)
	}
	return 1
}
func PhpMimeGetHdrValue(header zend.ZendLlist, key string) *byte {
	var entry *MimeHeaderEntry
	if key == nil {
		return nil
	}
	entry = zend.ZendLlistGetFirst(&header)
	for entry != nil {
		if !(strcasecmp(entry.GetKey(), key)) {
			return entry.GetValue()
		}
		entry = zend.ZendLlistGetNext(&header)
	}
	return nil
}
func PhpApGetword(encoding *zend.ZendEncoding, line **byte, stop byte) *byte {
	var pos *byte = *line
	var quote byte
	var res *byte
	for (*pos) && (*pos) != stop {
		if b.Assign(&quote, *pos) == '"' || quote == '\'' {
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
		res = zend.Estrdup(*line)
		*line += strlen(*line)
		return res
	}
	res = zend.Estrndup(*line, pos-(*line))
	for (*pos) == stop {
		pos++
	}
	*line = pos
	return res
}
func SubstringConf(start *byte, len_ int, quote byte) *byte {
	var result *byte = zend.Emalloc(len_ + 1)
	var resp *byte = result
	var i int
	for i = 0; i < len_ && start[i] != quote; i++ {
		if start[i] == '\\' && (start[i+1] == '\\' || quote && start[i+1] == quote) {
			b.PostInc(&(*resp)) = start[b.PreInc(&i)]
		} else {
			b.PostInc(&(*resp)) = start[i]
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
		return zend.Estrdup("")
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
func PhpApMemstr(haystack *byte, haystacklen int, needle *byte, needlen int, partial int) any {
	var len_ int = haystacklen
	var ptr *byte = haystack

	/* iterate through first character matches */

	for b.Assign(&ptr, memchr(ptr, needle[0], len_)) {

		/* calculate length after match */

		len_ = haystacklen - (ptr - (*byte)(haystack))

		/* done if matches up to capacity of buffer */

		if memcmp(needle, ptr, b.Cond(needlen < len_, needlen, len_)) == 0 && (partial != 0 || len_ >= needlen) {
			break
		}

		/* next character */

		ptr++
		len_--
	}
	return ptr
}
func MultipartBufferRead(self *MultipartBuffer, buf *byte, bytes int, end *int) int {
	var len_ int
	var max int
	var bound *byte

	/* fill buffer if needed */

	if bytes > int(self.GetBytesInBuffer()) {
		FillBuffer(self)
	}

	/* look for a potential boundary match, only read data up to that point */

	if b.Assign(&bound, PhpApMemstr(self.GetBufBegin(), self.GetBytesInBuffer(), self.GetBoundaryNext(), self.GetBoundaryNextLen(), 1)) {
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
			buf[b.PreDec(&len_)] = 0
		}

		/* update the buffer */

		self.SetBytesInBuffer(self.GetBytesInBuffer() - int(len_))
		self.SetBufBegin(self.GetBufBegin() + len_)
	}
	return len_
}
func MultipartBufferReadBody(self *MultipartBuffer, len_ *int) *byte {
	var buf []byte
	var out *byte = nil
	var total_bytes int = 0
	var read_bytes int = 0
	for b.Assign(&read_bytes, MultipartBufferRead(self, buf, b.SizeOf("buf"), nil)) {
		out = zend.Erealloc(out, total_bytes+read_bytes+1)
		memcpy(out+total_bytes, buf, read_bytes)
		total_bytes += read_bytes
	}
	if out != nil {
		out[total_bytes] = '0'
	}
	*len_ = total_bytes
	return out
}
func Rfc1867PostHandler(content_type_dup *byte, arg any) {
	var boundary *byte
	var s *byte = nil
	var boundary_end *byte = nil
	var start_arr *byte = nil
	var array_index *byte = nil
	var lbuf *byte = nil
	var abuf *byte = nil
	var temp_filename *types.String = nil
	var boundary_len int = 0
	var cancel_upload int = 0
	var is_arr_upload int = 0
	var array_len int = 0
	var total_bytes int64 = 0
	var max_file_size int64 = 0
	var skip_upload int = 0
	var anonindex int = 0
	var is_anonymous int
	var mbuff *MultipartBuffer
	var array_ptr *types.Zval = (*types.Zval)(arg)
	var fd int = -1
	var header zend.ZendLlist
	var event_extra_data any = nil
	var llen uint = 0
	var upload_cnt int = zend.INI_INT("max_file_uploads")
	var internal_encoding *zend.ZendEncoding = nil
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
	if SG__().post_max_size > 0 && SG__().request_info.content_length > SG__().post_max_size {
		SM__().SapiError(faults.E_WARNING, "POST Content-Length of "+zend.ZEND_LONG_FMT+" bytes exceeds the limit of "+zend.ZEND_LONG_FMT+" bytes", SG__().request_info.content_length, SG__().post_max_size)
		return
	}

	/* Get the boundary */

	boundary = strstr(content_type_dup, "boundary")
	if boundary == nil {
		var content_type_len int = int(strlen(content_type_dup))
		var content_type_lcase *byte = zend.Estrndup(content_type_dup, content_type_len)
		standard.PhpStrtolower(content_type_lcase, content_type_len)
		boundary = strstr(content_type_lcase, "boundary")
		if boundary != nil {
			boundary = content_type_dup + (boundary - content_type_lcase)
		}
		zend.Efree(content_type_lcase)
	}
	if boundary == nil || !(b.Assign(&boundary, strchr(boundary, '='))) {
		SM__().SapiError(faults.E_WARNING, "Missing boundary in multipart/form-data POST data")
		return
	}
	boundary++
	boundary_len = int(strlen(boundary))
	if boundary[0] == '"' {
		boundary++
		boundary_end = strchr(boundary, '"')
		if boundary_end == nil {
			SM__().SapiError(faults.E_WARNING, "Invalid boundary in multipart/form-data POST data")
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

	if !(b.Assign(&mbuff, MultipartBufferNew(boundary, boundary_len))) {
		SM__().SapiError(faults.E_WARNING, "Unable to initialize the input buffer")
		return
	}

	/* Initialize $_FILES[] */

	&(PG__().rfc1867_protected_variables) = types.MakeArrayEx(8, nil, 0)
	SG__().rfc1867_uploaded_files = make(map[string]bool)
	if PG__().http_globals[TRACK_VARS_FILES].GetType() != types.IS_ARRAY {

		/* php_auto_globals_create_files() might have already done that */

		zend.ArrayInit(&PG__().http_globals[TRACK_VARS_FILES])

		/* php_auto_globals_create_files() might have already done that */

	}
	header.Init(b.SizeOf("mime_header_entry"), zend.LlistDtorFuncT(PhpFreeHdrEntry), 0)
	if PhpRfc1867Callback != nil {
		var event_start MultipartEventStart
		event_start.SetContentLength(SG__().request_info.content_length)
		if PhpRfc1867Callback(MULTIPART_EVENT_START, &event_start, &event_extra_data) == types.FAILURE {
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
		header.Clean()
		if MultipartBufferHeaders(mbuff, &header) == 0 {
			goto fileupload_done
		}
		if b.Assign(&cd, PhpMimeGetHdrValue(header, "Content-Disposition")) {
			var pair *byte = nil
			var end int = 0
			for isspace(*cd) {
				cd++
			}
			for (*cd) && b.Assign(&pair, getword(mbuff.GetInputEncoding(), &cd, ';')) {
				var key *byte = nil
				var word *byte = pair
				for isspace(*cd) {
					cd++
				}
				if strchr(pair, '=') {
					key = getword(mbuff.GetInputEncoding(), &pair, '=')
					if !(strcasecmp(key, "name")) {
						if param != nil {
							zend.Efree(param)
						}
						param = getword_conf(mbuff.GetInputEncoding(), pair)
					} else if !(strcasecmp(key, "filename")) {
						if filename != nil {
							zend.Efree(filename)
						}
						filename = getword_conf(mbuff.GetInputEncoding(), pair)
					}
				}
				if key != nil {
					zend.Efree(key)
				}
				zend.Efree(word)
			}

			/* Normal form variable, safe to read all data into memory */

			if filename == nil && param != nil {
				var value_len int
				var value *byte = MultipartBufferReadBody(mbuff, &value_len)
				var new_val_len int
				if value == nil {
					value = zend.Estrdup("")
					value_len = 0
				}
				if b.PreInc(&count) <= PG__().max_input_vars && SM__().GetInputFilter()(PARSE_POST, param, &value, value_len, &new_val_len) != 0 {
					if PhpRfc1867Callback != nil {
						var event_formdata MultipartEventFormdata
						var newlength int = new_val_len
						event_formdata.SetPostBytesProcessed(SG__().read_post_bytes)
						event_formdata.SetName(param)
						event_formdata.SetValue(&value)
						event_formdata.SetLength(new_val_len)
						event_formdata.SetNewlength(&newlength)
						if PhpRfc1867Callback(MULTIPART_EVENT_FORMDATA, &event_formdata, &event_extra_data) == types.FAILURE {
							zend.Efree(param)
							zend.Efree(value)
							continue
						}
						new_val_len = newlength
					}
					SafePhpRegisterVariable(param, value, new_val_len, array_ptr, 0)
				} else {
					if count == PG__().max_input_vars+1 {
						PhpErrorDocref(nil, faults.E_WARNING, "Input variables exceeded "+zend.ZEND_LONG_FMT+". To increase the limit change max_input_vars in php.ini.", PG__().max_input_vars)
					}
					if PhpRfc1867Callback != nil {
						var event_formdata MultipartEventFormdata
						event_formdata.SetPostBytesProcessed(SG__().read_post_bytes)
						event_formdata.SetName(param)
						event_formdata.SetValue(&value)
						event_formdata.SetLength(value_len)
						event_formdata.SetNewlength(nil)
						PhpRfc1867Callback(MULTIPART_EVENT_FORMDATA, &event_formdata, &event_extra_data)
					}
				}
				if !(strcasecmp(param, "MAX_FILE_SIZE")) {
					max_file_size = atoll(value)
				}
				zend.Efree(param)
				zend.Efree(value)
				continue
			}

			/* If file_uploads=off, skip the file part */

			if !(PG__().file_uploads) {
				skip_upload = 1
			} else if upload_cnt <= 0 {
				skip_upload = 1
				SM__().SapiError(faults.E_WARNING, "Maximum number of allowable file uploads has been exceeded")
			}

			/* Return with an error if the posted data is garbled */

			if param == nil && filename == nil {
				SM__().SapiError(faults.E_WARNING, "File Upload Mime headers garbled")
				goto fileupload_done
			}
			if param == nil {
				is_anonymous = 1
				param = zend.Emalloc(MAX_SIZE_ANONNAME)
				Snprintf(param, MAX_SIZE_ANONNAME, "%u", b.PostInc(&anonindex))
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
				event_file_start.SetPostBytesProcessed(SG__().read_post_bytes)
				event_file_start.SetName(param)
				event_file_start.SetFilename(&filename)
				if PhpRfc1867Callback(MULTIPART_EVENT_FILE_START, &event_file_start, &event_extra_data) == types.FAILURE {
					temp_filename = nil
					zend.Efree(param)
					zend.Efree(filename)
					continue
				}
			}
			if skip_upload != 0 {
				zend.Efree(param)
				zend.Efree(filename)
				continue
			}
			if filename[0] == '0' {
				cancel_upload = UPLOAD_ERROR_D
			}
			offset = 0
			end = 0
			if cancel_upload == 0 {

				/* only bother to open temp file if we have data */

				blen = MultipartBufferRead(mbuff, buff, b.SizeOf("buff"), &end)

				/* in non-debug mode we have no problem with 0-length files */

				fd = PhpOpenTemporaryFdEx(PG__().upload_tmp_dir, "php", &temp_filename, PHP_TMP_FILE_OPEN_BASEDIR_CHECK_ON_FALLBACK)
				upload_cnt--
				if fd == -1 {
					SM__().SapiError(faults.E_WARNING, "File upload error - unable to create a temporary file")
					cancel_upload = UPLOAD_ERROR_E
				}

				/* in non-debug mode we have no problem with 0-length files */

			}
			for cancel_upload == 0 && blen > 0 {
				if PhpRfc1867Callback != nil {
					var event_file_data MultipartEventFileData
					event_file_data.SetPostBytesProcessed(SG__().read_post_bytes)
					event_file_data.SetOffset(offset)
					event_file_data.SetData(buff)
					event_file_data.SetLength(blen)
					event_file_data.SetNewlength(&blen)
					if PhpRfc1867Callback(MULTIPART_EVENT_FILE_DATA, &event_file_data, &event_extra_data) == types.FAILURE {
						cancel_upload = UPLOAD_ERROR_X
						continue
					}
				}
				if PG__().upload_max_filesize > 0 && zend_long(total_bytes+blen) > PG__().upload_max_filesize {
					cancel_upload = UPLOAD_ERROR_A
				} else if max_file_size != 0 && zend_long(total_bytes+blen) > max_file_size {
					cancel_upload = UPLOAD_ERROR_B
				} else if blen > 0 {
					wlen = write(fd, buff, blen)
					if wlen == size_t-1 {

						/* write failed */

						cancel_upload = UPLOAD_ERROR_F

						/* write failed */

					} else if wlen < blen {
						cancel_upload = UPLOAD_ERROR_F
					} else {
						total_bytes += wlen
					}
					offset += wlen
				}

				/* read data for next iteration */

				blen = MultipartBufferRead(mbuff, buff, b.SizeOf("buff"), &end)

				/* read data for next iteration */

			}
			if fd != -1 {
				close(fd)
			}
			if cancel_upload == 0 && end == 0 {
				cancel_upload = UPLOAD_ERROR_C
			}
			if PhpRfc1867Callback != nil {
				var event_file_end MultipartEventFileEnd
				event_file_end.SetPostBytesProcessed(SG__().read_post_bytes)
				if temp_filename != nil {
					event_file_end.SetTempFilename(temp_filename.GetVal())
				} else {
					event_file_end.SetTempFilename(nil)
				}
				event_file_end.SetCancelUpload(cancel_upload)
				if PhpRfc1867Callback(MULTIPART_EVENT_FILE_END, &event_file_end, &event_extra_data) == types.FAILURE {
					cancel_upload = UPLOAD_ERROR_X
				}
			}
			if cancel_upload != 0 {
				if temp_filename != nil {
					if cancel_upload != UPLOAD_ERROR_E {
						unlink(temp_filename.GetVal())
					}
					// types.ZendStringReleaseEx(temp_filename, 0)
				}
				temp_filename = nil
			} else {
				SG__().rfc1867_uploaded_files[temp_filename.GetStr()] = true
			}

			/* is_arr_upload is true when name of file upload field
			 * ends in [.*]
			 * start_arr is set to point to 1st [ */

			is_arr_upload = b.Assign(&start_arr, strchr(param, '[')) && param[strlen(param)-1] == ']'
			if is_arr_upload != 0 {
				array_len = strlen(start_arr)
				if array_index != nil {
					zend.Efree(array_index)
				}
				array_index = zend.Estrndup(start_arr+1, array_len-2)
			}

			/* Add $foo_name */

			if llen < strlen(param)+MAX_SIZE_OF_INDEX+1 {
				llen = int(strlen(param))
				lbuf = (*byte)(zend.SafeErealloc(lbuf, llen, 1, MAX_SIZE_OF_INDEX+1))
				llen += MAX_SIZE_OF_INDEX + 1
			}
			if is_arr_upload != 0 {
				if abuf != nil {
					zend.Efree(abuf)
				}
				abuf = zend.Estrndup(param, strlen(param)-array_len)
				Snprintf(lbuf, llen, "%s_name[%s]", abuf, array_index)
			} else {
				Snprintf(lbuf, llen, "%s_name", param)
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
				Snprintf(lbuf, llen, "%s[name][%s]", abuf, array_index)
			} else {
				Snprintf(lbuf, llen, "%s[name]", param)
			}
			RegisterHttpPostFilesVariable(lbuf, s, &PG__().http_globals[TRACK_VARS_FILES], 0)
			zend.Efree(filename)
			s = nil

			/* Possible Content-Type: */

			if cancel_upload != 0 || !(b.Assign(&cd, PhpMimeGetHdrValue(header, "Content-Type"))) {
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
				Snprintf(lbuf, llen, "%s_type[%s]", abuf, array_index)
			} else {
				Snprintf(lbuf, llen, "%s_type", param)
			}
			if is_anonymous == 0 {
				SafePhpRegisterVariable(lbuf, cd, strlen(cd), nil, 0)
			}

			/* Add $foo[type] */

			if is_arr_upload != 0 {
				Snprintf(lbuf, llen, "%s[type][%s]", abuf, array_index)
			} else {
				Snprintf(lbuf, llen, "%s[type]", param)
			}
			RegisterHttpPostFilesVariable(lbuf, cd, &PG__().http_globals[TRACK_VARS_FILES], 0)

			/* Restore Content-Type Header */

			if s != nil {
				*s = ';'
			}
			s = ""

			/* store temp_filename as-is (in case upload_tmp_dir
			 * contains escapeable characters. escape only the variable name.) */

			var zfilename types.Zval

			/* Initialize variables */

			AddProtectedVariable(param)

			/* if param is of form xxx[.*] this will cut it to xxx */

			if is_anonymous == 0 {
				if temp_filename != nil {
					zfilename.SetStringCopy(temp_filename)
				} else {
					zend.ZVAL_EMPTY_STRING(&zfilename)
				}
				SafePhpRegisterVariableEx(param, &zfilename, nil, 1)
			}

			/* Add $foo[tmp_name] */

			if is_arr_upload != 0 {
				Snprintf(lbuf, llen, "%s[tmp_name][%s]", abuf, array_index)
			} else {
				Snprintf(lbuf, llen, "%s[tmp_name]", param)
			}
			AddProtectedVariable(lbuf)
			if temp_filename != nil {
				zfilename.SetStringCopy(temp_filename)
			} else {
				zend.ZVAL_EMPTY_STRING(&zfilename)
			}
			RegisterHttpPostFilesVariableEx(lbuf, &zfilename, &PG__().http_globals[TRACK_VARS_FILES], 1)
			var file_size types.Zval
			var error_type types.Zval
			var size_overflow int = 0
			var file_size_buf []byte
			error_type.SetLong(cancel_upload)

			/* Add $foo[error] */

			if cancel_upload != 0 {
				file_size.SetLong(0)
			} else {
				if total_bytes > zend.ZEND_LONG_MAX {
					var __len int = Snprintf(file_size_buf, 65, "%"+"lld", total_bytes)
					file_size_buf[__len] = '0'
					size_overflow = 1
				} else {
					file_size.SetLong(total_bytes)
				}
			}
			if is_arr_upload != 0 {
				Snprintf(lbuf, llen, "%s[error][%s]", abuf, array_index)
			} else {
				Snprintf(lbuf, llen, "%s[error]", param)
			}
			RegisterHttpPostFilesVariableEx(lbuf, &error_type, &PG__().http_globals[TRACK_VARS_FILES], 0)

			/* Add $foo_size */

			if is_arr_upload != 0 {
				Snprintf(lbuf, llen, "%s_size[%s]", abuf, array_index)
			} else {
				Snprintf(lbuf, llen, "%s_size", param)
			}
			if is_anonymous == 0 {
				if size_overflow != 0 {
					file_size.SetStringVal(b.CastStrAuto(file_size_buf))
				}
				SafePhpRegisterVariableEx(lbuf, &file_size, nil, size_overflow)
			}

			/* Add $foo[size] */

			if is_arr_upload != 0 {
				Snprintf(lbuf, llen, "%s[size][%s]", abuf, array_index)
			} else {
				Snprintf(lbuf, llen, "%s[size]", param)
			}
			if size_overflow != 0 {
				file_size.SetStringVal(b.CastStrAuto(file_size_buf))
			}
			RegisterHttpPostFilesVariableEx(lbuf, &file_size, &PG__().http_globals[TRACK_VARS_FILES], size_overflow)
			zend.Efree(param)
		}
	}
fileupload_done:
	if PhpRfc1867Callback != nil {
		var event_end MultipartEventEnd
		event_end.SetPostBytesProcessed(SG__().read_post_bytes)
		PhpRfc1867Callback(MULTIPART_EVENT_END, &event_end, &event_extra_data)
	}
	if lbuf != nil {
		zend.Efree(lbuf)
	}
	if abuf != nil {
		zend.Efree(abuf)
	}
	if array_index != nil {
		zend.Efree(array_index)
	}
	PG__().rfc1867_protected_variables.Destroy()
	header.Destroy()
	if mbuff.GetBoundaryNext() != nil {
		zend.Efree(mbuff.GetBoundaryNext())
	}
	if mbuff.GetBoundary() != nil {
		zend.Efree(mbuff.GetBoundary())
	}
	if mbuff.GetBuffer() != nil {
		zend.Efree(mbuff.GetBuffer())
	}
	if mbuff != nil {
		zend.Efree(mbuff)
	}
}
