package standard

import (
	"fmt"
	b "github.com/heyuuu/gophp/builtin"
	r "github.com/heyuuu/gophp/builtin/file"
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/ext/standard/str"
	"github.com/heyuuu/gophp/kits/ascii"
	"github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/zpp"
	"strings"
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
	for lang.Assign(&p, memchr(p, '0', e-p)) {
		*p = ' '
	}
}
func ZifEzmlmHash(executeData zpp.Ex, return_value zpp.Ret, addr *types.Zval) {
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
func PhpMailBuildHeadersCheckFieldValue(value string) bool {
	/* https://tools.ietf.org/html/rfc2822#section-2.2.1 */
	l := len(value)
	for i := 0; i < l; {
		if value[i] == '\r' {
			if i+2 < l && value[i+1] == '\n' && (value[i+2] == ' ' || value[i+2] == '\t') {
				l += 3
				continue
			}
			return false
		}
		if value[i] == 0 {
			return false
		}
		i++
	}
	return true
}
func PhpMailBuildHeadersCheckFieldName(key string) bool {
	/* https://tools.ietf.org/html/rfc2822#section-2.2 */
	for _, c := range []byte(key) {
		if c < 33 || c > 126 || c == ':' {
			return false
		}
	}
	return true
}

func PhpMailBuildHeadersElem(s *strings.Builder, key string, val string) {
	if !PhpMailBuildHeadersCheckFieldName(key) {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Header field name (%s) contains invalid chars", key)
		return
	}
	if !PhpMailBuildHeadersCheckFieldValue(val) {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Header field value (%s => %s) contains invalid chars or format", key, val)
		return
	}
	s.WriteString(key)
	s.WriteString(": ")
	s.WriteString(val)
	s.WriteString("\r\n")
}
func PhpMailBuildHeadersElems(s *strings.Builder, key string, val *types.Array) {
	val.Foreach(func(tmpKey types.ArrayKey, tmpVal *types.Zval) {
		if tmpKey.IsStrKey() {
			core.PhpErrorDocref(nil, faults.E_WARNING, "Multiple header key must be numeric index (%s)", tmpKey.StrKey())
			return
		}
		if !tmpVal.IsString() {
			core.PhpErrorDocref(nil, faults.E_WARNING, "Multiple header values must be string (%s)", key.GetStr())
			return
		}
		PhpMailBuildHeadersElem(s, key, tmpVal.String())
	})
}
func PhpMailBuildHeader(s *strings.Builder, key string, val *types.Zval, check bool) {
	if val.IsString() {
		PhpMailBuildHeadersElem(s, key, val.String())
	} else if val.IsType(types.IsArray) {
		if check {
			core.PhpErrorDocref(nil, faults.E_WARNING, "'%s' header must be at most one header. Array is passed for '%s'", key, key)
			return
		}
		PhpMailBuildHeadersElems(s, key, val.Array())
	} else {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Extra header element '%s' cannot be other than string or array.", key)
	}
}
func PhpMailBuildHeaders(headers *types.Zval) string {
	b.Assert(headers.IsArray())

	var s strings.Builder
	headers.Array().Foreach(func(arrayKey types.ArrayKey, value *types.Zval) {
		if !arrayKey.IsStrKey() {
			core.PhpErrorDocref(nil, faults.E_WARNING, "Found numeric header ("+zend.ZEND_LONG_FMT+")", arrayKey.IdxKey())
			return
		}

		/* https://tools.ietf.org/html/rfc2822#section-3.6 */
		key := arrayKey.StrKey()
		switch ascii.StrToLower(key) {
		case "orig-date", "from", "sender", "reply-to", "cc", "bcc", "message-id", "references", "in-reply-to":
			PhpMailBuildHeader(&s, key, value, true)
		case "to":
			core.PhpErrorDocref(nil, faults.E_WARNING, "Extra header cannot contain 'To' header")
			return
		case "subject":
			core.PhpErrorDocref(nil, faults.E_WARNING, "Extra header cannot contain 'Subject' header")
			return
		default:
			PhpMailBuildHeader(&s, key, value, false)
		}

		/* https://tools.ietf.org/html/rfc2822#section-3.6 */
	})

	/* Remove the last \r\n */
	result := s.String()
	if len(result) > 2 {
		result = result[:len(result)-2]
	}
	return result
}
func ZifMail(executeData zpp.Ex, return_value zpp.Ret, to *types.Zval, subject *types.Zval, message *types.Zval, _ zpp.Opt, additionalHeaders *types.Zval, additionalParameters *types.Zval) {
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
		switch headers.Type() {
		case types.IsString:
			tmp_headers = types.NewString(headers.StringEx().GetStr())
			MAIL_ASCIIZ_CHECK(tmp_headers.GetVal(), tmp_headers.GetLen())
			str_headers = types.NewString(str.PhpTrimRight(tmp_headers.GetStr(), nil))
			// types.ZendStringReleaseEx(tmp_headers, 0)
		case types.IsArray:
			str_headers = types.NewString(PhpMailBuildHeaders(headers))
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
	if PhpMail(to_r, subject_r, message, lang.CondF1(str_headers != nil && str_headers.GetLen() != 0, func() []byte { return str_headers.GetVal() }, nil), lang.CondF1(extra_cmd != nil, func() []byte { return extra_cmd.GetVal() }, nil)) != 0 {
		return_value.SetTrue()
	} else {
		return_value.SetFalse()
	}
	if str_headers != nil {
		// types.ZendStringReleaseEx(str_headers, 0)
	}
	if extra_cmd != nil {
		// types.ZendStringReleaseEx(extra_cmd, 0)
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
	for lang.Assign(&p, strpbrk(p, "\r\n")) {
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
	var sendmail *r.File
	var ret int
	var sendmail_path *byte = zend.INI_STR("sendmail_path")
	var sendmail_cmd *byte = nil
	var mail_log *byte = zend.INI_STR("mail.log")
	var hdr *byte = headers

	// #define MAIL_RET(val) if ( hdr != headers ) { efree ( hdr ) ; } return val ;

	if mail_log != nil && (*mail_log) {
		var logline *byte
		core.Spprintf(&logline, 0, "mail() on [%s:%d]: To: %s -- Headers: %s -- Subject: %s", zend.ZendGetExecutedFilename(), zend.ZendGetExecutedLineno(), to, lang.Cond(hdr != nil, hdr, ""), subject)
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
			//types.ZendStringFree(date_str)
			zend.Efree(tmp)
		}
		zend.Efree(logline)
	}
	if core.PG__().mail_x_header {
		var tmp *byte = zend.ZendGetExecutedFilename()
		var f *types.String
		f = str.PhpBasenameZStr(tmp, "")
		if headers != nil && (*headers) {
			core.Spprintf(&hdr, 0, "X-PHP-Originating-Script: "+zend.ZEND_LONG_FMT+":%s\n%s", PhpGetuid(), f.GetVal(), headers)
		} else {
			core.Spprintf(&hdr, 0, "X-PHP-Originating-Script: "+zend.ZEND_LONG_FMT+":%s", PhpGetuid(), f.GetVal())
		}
		// types.ZendStringReleaseEx(f, 0)
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
		_, _ = sendmail.WriteString(fmt.Sprintf("To: %s\n", to))
		_, _ = sendmail.WriteString(fmt.Sprintf("Subject: %s\n", subject))
		if hdr != nil {
			_, _ = sendmail.WriteString(fmt.Sprintf("%s\n", hdr))
		}
		_, _ = sendmail.WriteString(fmt.Sprintf("\n%s\n", message))
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
func ZmInfoMail(zend_module *zend.ModuleEntry) {
	var sendmail_path *byte = zend.INI_STR("sendmail_path")
	PhpInfoPrintTableRow(2, "Path to sendmail", sendmail_path)
}
