// <<generate>>

package standard

import (
	b "sik/builtin"
	r "sik/builtin/file"
	"sik/core"
	"sik/zend"
	"sik/zend/faults"
	"sik/zend/types"
	"sik/zend/zpp"
)

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
func ZifEzmlmHash(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var str *byte = nil
	var h uint = 5381
	var j int
	var str_len int
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 1, 0)
			str, str_len = fp.ParseString()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	for j = 0; j < str_len; j++ {
		h = h + (h << 5) ^ zend.ZendUlong(uint8(tolower(str[j])))
	}
	h = h % 53
	return_value.SetLong(zend.ZendLong(h))
	return
}
func PhpMailBuildHeadersCheckFieldValue(val *types.Zval) types.ZendBool {
	var len_ int = 0
	var value *types.String = val.GetStr()

	/* https://tools.ietf.org/html/rfc2822#section-2.2.1 */

	for len_ < value.GetLen() {
		if (*(value.GetVal() + len_)) == '\r' {
			if value.GetLen()-len_ >= 3 && (*(value.GetVal() + len_ + 1)) == '\n' && ((*(value.GetVal() + len_ + 2)) == ' ' || (*(value.GetVal() + len_ + 2)) == '\t') {
				len_ += 3
				continue
			}
			return types.FAILURE
		}
		if (*(value.GetVal() + len_)) == '0' {
			return types.FAILURE
		}
		len_++
	}
	return types.SUCCESS
}
func PhpMailBuildHeadersCheckFieldName(key *types.String) types.ZendBool {
	var len_ int = 0

	/* https://tools.ietf.org/html/rfc2822#section-2.2 */

	for len_ < key.GetLen() {
		if (*(key.GetVal() + len_)) < 33 || (*(key.GetVal() + len_)) > 126 || (*(key.GetVal() + len_)) == ':' {
			return types.FAILURE
		}
		len_++
	}
	return types.SUCCESS
}
func PhpMailBuildHeadersElem(s *zend.SmartStr, key *types.String, val *types.Zval) {
	switch val.GetType() {
	case types.IS_STRING:
		if PhpMailBuildHeadersCheckFieldName(key) != types.SUCCESS {
			core.PhpErrorDocref(nil, faults.E_WARNING, "Header field name (%s) contains invalid chars", key.GetVal())
			return
		}
		if PhpMailBuildHeadersCheckFieldValue(val) != types.SUCCESS {
			core.PhpErrorDocref(nil, faults.E_WARNING, "Header field value (%s => %s) contains invalid chars or format", key.GetVal(), val.GetStr().GetVal())
			return
		}
		s.AppendString(key.GetStr())
		s.AppendString(": ")
		s.AppendString(b.CastStrAuto(val.GetStr().GetVal()))
		s.AppendString("\r\n")
	case types.IS_ARRAY:
		PhpMailBuildHeadersElems(s, key, val)
	default:
		core.PhpErrorDocref(nil, faults.E_WARNING, "headers array elements must be string or array (%s)", key.GetVal())
	}
}
func PhpMailBuildHeadersElems(s *zend.SmartStr, key *types.String, val *types.Zval) {
	var tmp_key *types.String
	var tmp_val *types.Zval
	var __ht *types.Array = val.GetArr()
	for _, _p := range __ht.ForeachData() {
		var _z *types.Zval = _p.GetVal()

		tmp_key = _p.GetKey()
		tmp_val = _z
		if tmp_key != nil {
			core.PhpErrorDocref(nil, faults.E_WARNING, "Multiple header key must be numeric index (%s)", tmp_key.GetVal())
			continue
		}
		if tmp_val.GetType() != types.IS_STRING {
			core.PhpErrorDocref(nil, faults.E_WARNING, "Multiple header values must be string (%s)", key.GetVal())
			continue
		}
		PhpMailBuildHeadersElem(s, key, tmp_val)
	}
}
func PhpMailBuildHeaders(headers *types.Zval) *types.String {
	var idx zend.ZendUlong
	var key *types.String
	var val *types.Zval
	var s zend.SmartStr = zend.MakeSmartStr(0)
	b.Assert(headers.IsType(types.IS_ARRAY))
	var __ht *types.Array = headers.GetArr()
	for _, _p := range __ht.ForeachData() {
		var _z *types.Zval = _p.GetVal()

		idx = _p.GetH()
		key = _p.GetKey()
		val = _z
		if key == nil {
			core.PhpErrorDocref(nil, faults.E_WARNING, "Found numeric header ("+zend.ZEND_LONG_FMT+")", idx)
			continue
		}

		/* https://tools.ietf.org/html/rfc2822#section-3.6 */

		switch key.GetLen() {
		case b.SizeOf("\"orig-date\"") - 1:
			if !(strncasecmp("orig-date", key.GetVal(), key.GetLen())) {
				PHP_MAIL_BUILD_HEADER_CHECK("orig-date", s, key, val)
			} else {
				PHP_MAIL_BUILD_HEADER_DEFAULT(s, key, val)
			}
		case b.SizeOf("\"from\"") - 1:
			if !(strncasecmp("from", key.GetVal(), key.GetLen())) {
				PHP_MAIL_BUILD_HEADER_CHECK("from", s, key, val)
			} else {
				PHP_MAIL_BUILD_HEADER_DEFAULT(s, key, val)
			}
		case b.SizeOf("\"sender\"") - 1:
			if !(strncasecmp("sender", key.GetVal(), key.GetLen())) {
				PHP_MAIL_BUILD_HEADER_CHECK("sender", s, key, val)
			} else {
				PHP_MAIL_BUILD_HEADER_DEFAULT(s, key, val)
			}
		case b.SizeOf("\"reply-to\"") - 1:
			if !(strncasecmp("reply-to", key.GetVal(), key.GetLen())) {
				PHP_MAIL_BUILD_HEADER_CHECK("reply-to", s, key, val)
			} else {
				PHP_MAIL_BUILD_HEADER_DEFAULT(s, key, val)
			}
		case b.SizeOf("\"to\"") - 1:
			if !(strncasecmp("to", key.GetVal(), key.GetLen())) {
				core.PhpErrorDocref(nil, faults.E_WARNING, "Extra header cannot contain 'To' header")
				continue
			}
			if !(strncasecmp("cc", key.GetVal(), key.GetLen())) {
				PHP_MAIL_BUILD_HEADER_CHECK("cc", s, key, val)
			} else {
				PHP_MAIL_BUILD_HEADER_DEFAULT(s, key, val)
			}
		case b.SizeOf("\"bcc\"") - 1:
			if !(strncasecmp("bcc", key.GetVal(), key.GetLen())) {
				PHP_MAIL_BUILD_HEADER_CHECK("bcc", s, key, val)
			} else {
				PHP_MAIL_BUILD_HEADER_DEFAULT(s, key, val)
			}
		case b.SizeOf("\"message-id\"") - 1:
			if !(strncasecmp("message-id", key.GetVal(), key.GetLen())) {
				PHP_MAIL_BUILD_HEADER_CHECK("message-id", s, key, val)
			} else if !(strncasecmp("references", key.GetVal(), key.GetLen())) {
				PHP_MAIL_BUILD_HEADER_CHECK("references", s, key, val)
			} else {
				PHP_MAIL_BUILD_HEADER_DEFAULT(s, key, val)
			}
		case b.SizeOf("\"in-reply-to\"") - 1:
			if !(strncasecmp("in-reply-to", key.GetVal(), key.GetLen())) {
				PHP_MAIL_BUILD_HEADER_CHECK("in-reply-to", s, key, val)
			} else {
				PHP_MAIL_BUILD_HEADER_DEFAULT(s, key, val)
			}
		case b.SizeOf("\"subject\"") - 1:
			if !(strncasecmp("subject", key.GetVal(), key.GetLen())) {
				core.PhpErrorDocref(nil, faults.E_WARNING, "Extra header cannot contain 'Subject' header")
				continue
			}
			PHP_MAIL_BUILD_HEADER_DEFAULT(s, key, val)
		default:
			PHP_MAIL_BUILD_HEADER_DEFAULT(s, key, val)
		}

		/* https://tools.ietf.org/html/rfc2822#section-3.6 */

	}

	/* Remove the last \r\n */

	if s.GetS() != nil {
		s.GetS().SetLen(s.GetS().GetLen() - 2)
	}
	s.ZeroTail()
	return s.GetS()
}
func ZifMail(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var to *byte = nil
	var message *byte = nil
	var subject *byte = nil
	var extra_cmd *types.String = nil
	var str_headers *types.String = nil
	var tmp_headers *types.String
	var headers *types.Zval = nil
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
		for {
			fp := zpp.FastParseStart(executeData, 3, 5, 0)
			to, to_len = fp.ParseString()
			subject, subject_len = fp.ParseString()
			message, message_len = fp.ParseString()
			fp.StartOptional()
			headers = fp.ParseZval()
			extra_cmd = fp.ParseStr()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}

	/* ASCIIZ check */

	MAIL_ASCIIZ_CHECK(to, to_len)
	MAIL_ASCIIZ_CHECK(subject, subject_len)
	MAIL_ASCIIZ_CHECK(message, message_len)
	if headers != nil {
		switch headers.GetType() {
		case types.IS_STRING:
			tmp_headers = types.NewString(headers.GetStr().GetStr())
			MAIL_ASCIIZ_CHECK(tmp_headers.GetVal(), tmp_headers.GetLen())
			str_headers = PhpTrim(tmp_headers, nil, 0, 2)
			types.ZendStringReleaseEx(tmp_headers, 0)
		case types.IS_ARRAY:
			str_headers = PhpMailBuildHeaders(headers)
		default:
			core.PhpErrorDocref(nil, faults.E_WARNING, "headers parameter must be string or array")
			return_value.SetFalse()
			return
		}
	}
	if extra_cmd != nil {
		MAIL_ASCIIZ_CHECK(extra_cmd.GetVal(), extra_cmd.GetLen())
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
		extra_cmd = PhpEscapeShellCmd(extra_cmd.GetVal())
	}
	if PhpMail(to_r, subject_r, message, b.CondF1(str_headers != nil && str_headers.GetLen() != 0, func() []byte { return str_headers.GetVal() }, nil), b.CondF1(extra_cmd != nil, func() []byte { return extra_cmd.GetVal() }, nil)) != 0 {
		return_value.SetTrue()
	} else {
		return_value.SetFalse()
	}
	if str_headers != nil {
		types.ZendStringReleaseEx(str_headers, 0)
	}
	if extra_cmd != nil {
		types.ZendStringReleaseEx(extra_cmd, 0)
	}
	if to_r != to {
		zend.Efree(to_r)
	}
	if subject_r != subject {
		zend.Efree(subject_r)
	}
}
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
			var date_str *types.String
			var len_ int
			time(&curtime)
			date_str = php_format_date("d-M-Y H:i:s e", 13, curtime, 1)
			len_ = core.Spprintf(&tmp, 0, "[%s] %s%s", date_str.GetVal(), logline, core.PHP_EOL)
			PhpMailLogToFile(mail_log, tmp, len_)
			types.ZendStringFree(date_str)
			zend.Efree(tmp)
		}
		zend.Efree(logline)
	}
	if core.PG__().mail_x_header {
		var tmp *byte = zend.ZendGetExecutedFilename()
		var f *types.String
		f = PhpBasename(tmp, strlen(tmp), nil, 0)
		if headers != nil && (*headers) {
			core.Spprintf(&hdr, 0, "X-PHP-Originating-Script: "+zend.ZEND_LONG_FMT+":%s\n%s", PhpGetuid(), f.GetVal(), headers)
		} else {
			core.Spprintf(&hdr, 0, "X-PHP-Originating-Script: "+zend.ZEND_LONG_FMT+":%s", PhpGetuid(), f.GetVal())
		}
		types.ZendStringReleaseEx(f, 0)
	}
	if hdr != nil && PhpMailDetectMultipleCrlf(hdr) != 0 {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Multiple or malformed newlines found in additional_header")
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
			core.PhpErrorDocref(nil, faults.E_WARNING, "Permission denied: unable to execute shell to run mail delivery binary '%s'", sendmail_path)
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
		core.PhpErrorDocref(nil, faults.E_WARNING, "Could not execute mail delivery program '%s'", sendmail_path)
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
func ZmInfoMail(zend_module *zend.ZendModuleEntry) {
	var sendmail_path *byte = zend.INI_STR("sendmail_path")
	PhpInfoPrintTableRow(2, "Path to sendmail", sendmail_path)
}
